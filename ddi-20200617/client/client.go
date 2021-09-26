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

type CreateFlowRequest struct {
	RegionId                *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProjectId               *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Name                    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	StartSchedule           *int64  `json:"StartSchedule,omitempty" xml:"StartSchedule,omitempty"`
	EndSchedule             *int64  `json:"EndSchedule,omitempty" xml:"EndSchedule,omitempty"`
	CronExpression          *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	CreateCluster           *bool   `json:"CreateCluster,omitempty" xml:"CreateCluster,omitempty"`
	ClusterId               *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	HostName                *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	Namespace               *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Application             *string `json:"Application,omitempty" xml:"Application,omitempty"`
	AlertConf               *string `json:"AlertConf,omitempty" xml:"AlertConf,omitempty"`
	AlertUserGroupBizId     *string `json:"AlertUserGroupBizId,omitempty" xml:"AlertUserGroupBizId,omitempty"`
	AlertDingDingGroupBizId *string `json:"AlertDingDingGroupBizId,omitempty" xml:"AlertDingDingGroupBizId,omitempty"`
	ParentFlowList          *string `json:"ParentFlowList,omitempty" xml:"ParentFlowList,omitempty"`
	ParentCategory          *string `json:"ParentCategory,omitempty" xml:"ParentCategory,omitempty"`
	ClientToken             *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowRequest) SetRegionId(v string) *CreateFlowRequest {
	s.RegionId = &v
	return s
}

func (s *CreateFlowRequest) SetProjectId(v string) *CreateFlowRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateFlowRequest) SetName(v string) *CreateFlowRequest {
	s.Name = &v
	return s
}

func (s *CreateFlowRequest) SetDescription(v string) *CreateFlowRequest {
	s.Description = &v
	return s
}

func (s *CreateFlowRequest) SetStartSchedule(v int64) *CreateFlowRequest {
	s.StartSchedule = &v
	return s
}

func (s *CreateFlowRequest) SetEndSchedule(v int64) *CreateFlowRequest {
	s.EndSchedule = &v
	return s
}

func (s *CreateFlowRequest) SetCronExpression(v string) *CreateFlowRequest {
	s.CronExpression = &v
	return s
}

func (s *CreateFlowRequest) SetCreateCluster(v bool) *CreateFlowRequest {
	s.CreateCluster = &v
	return s
}

func (s *CreateFlowRequest) SetClusterId(v string) *CreateFlowRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateFlowRequest) SetHostName(v string) *CreateFlowRequest {
	s.HostName = &v
	return s
}

func (s *CreateFlowRequest) SetNamespace(v string) *CreateFlowRequest {
	s.Namespace = &v
	return s
}

func (s *CreateFlowRequest) SetApplication(v string) *CreateFlowRequest {
	s.Application = &v
	return s
}

func (s *CreateFlowRequest) SetAlertConf(v string) *CreateFlowRequest {
	s.AlertConf = &v
	return s
}

func (s *CreateFlowRequest) SetAlertUserGroupBizId(v string) *CreateFlowRequest {
	s.AlertUserGroupBizId = &v
	return s
}

func (s *CreateFlowRequest) SetAlertDingDingGroupBizId(v string) *CreateFlowRequest {
	s.AlertDingDingGroupBizId = &v
	return s
}

func (s *CreateFlowRequest) SetParentFlowList(v string) *CreateFlowRequest {
	s.ParentFlowList = &v
	return s
}

func (s *CreateFlowRequest) SetParentCategory(v string) *CreateFlowRequest {
	s.ParentCategory = &v
	return s
}

func (s *CreateFlowRequest) SetClientToken(v string) *CreateFlowRequest {
	s.ClientToken = &v
	return s
}

type CreateFlowResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlowResponseBody) SetId(v string) *CreateFlowResponseBody {
	s.Id = &v
	return s
}

func (s *CreateFlowResponseBody) SetRequestId(v string) *CreateFlowResponseBody {
	s.RequestId = &v
	return s
}

type CreateFlowResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowResponse) GoString() string {
	return s.String()
}

func (s *CreateFlowResponse) SetHeaders(v map[string]*string) *CreateFlowResponse {
	s.Headers = v
	return s
}

func (s *CreateFlowResponse) SetBody(v *CreateFlowResponseBody) *CreateFlowResponse {
	s.Body = v
	return s
}

type ListUsersRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Type            *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) SetResourceOwnerId(v int64) *ListUsersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListUsersRequest) SetRegionId(v string) *ListUsersRequest {
	s.RegionId = &v
	return s
}

func (s *ListUsersRequest) SetClusterId(v string) *ListUsersRequest {
	s.ClusterId = &v
	return s
}

func (s *ListUsersRequest) SetType(v string) *ListUsersRequest {
	s.Type = &v
	return s
}

type ListUsersResponseBody struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserList  *ListUsersResponseBodyUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Struct"`
}

func (s ListUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersResponseBody) SetUserList(v *ListUsersResponseBodyUserList) *ListUsersResponseBody {
	s.UserList = v
	return s
}

type ListUsersResponseBodyUserList struct {
	User []*ListUsersResponseBodyUserListUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBodyUserList) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyUserList) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUserList) SetUser(v []*ListUsersResponseBodyUserListUser) *ListUsersResponseBodyUserList {
	s.User = v
	return s
}

type ListUsersResponseBodyUserListUser struct {
	LinuxStatus    *string `json:"LinuxStatus,omitempty" xml:"LinuxStatus,omitempty"`
	K8sStatus      *string `json:"K8sStatus,omitempty" xml:"K8sStatus,omitempty"`
	KnoxStatus     *string `json:"KnoxStatus,omitempty" xml:"KnoxStatus,omitempty"`
	GroupName      *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	KerberosStatus *string `json:"KerberosStatus,omitempty" xml:"KerberosStatus,omitempty"`
	UserName       *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListUsersResponseBodyUserListUser) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyUserListUser) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUserListUser) SetLinuxStatus(v string) *ListUsersResponseBodyUserListUser {
	s.LinuxStatus = &v
	return s
}

func (s *ListUsersResponseBodyUserListUser) SetK8sStatus(v string) *ListUsersResponseBodyUserListUser {
	s.K8sStatus = &v
	return s
}

func (s *ListUsersResponseBodyUserListUser) SetKnoxStatus(v string) *ListUsersResponseBodyUserListUser {
	s.KnoxStatus = &v
	return s
}

func (s *ListUsersResponseBodyUserListUser) SetGroupName(v string) *ListUsersResponseBodyUserListUser {
	s.GroupName = &v
	return s
}

func (s *ListUsersResponseBodyUserListUser) SetUserId(v string) *ListUsersResponseBodyUserListUser {
	s.UserId = &v
	return s
}

func (s *ListUsersResponseBodyUserListUser) SetKerberosStatus(v string) *ListUsersResponseBodyUserListUser {
	s.KerberosStatus = &v
	return s
}

func (s *ListUsersResponseBodyUserListUser) SetUserName(v string) *ListUsersResponseBodyUserListUser {
	s.UserName = &v
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

type ModifyFlowProjectRequest struct {
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ModifyFlowProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowProjectRequest) GoString() string {
	return s.String()
}

func (s *ModifyFlowProjectRequest) SetProjectId(v string) *ModifyFlowProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *ModifyFlowProjectRequest) SetRegionId(v string) *ModifyFlowProjectRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyFlowProjectRequest) SetName(v string) *ModifyFlowProjectRequest {
	s.Name = &v
	return s
}

func (s *ModifyFlowProjectRequest) SetDescription(v string) *ModifyFlowProjectRequest {
	s.Description = &v
	return s
}

type ModifyFlowProjectResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyFlowProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowProjectResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFlowProjectResponseBody) SetData(v bool) *ModifyFlowProjectResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyFlowProjectResponseBody) SetRequestId(v string) *ModifyFlowProjectResponseBody {
	s.RequestId = &v
	return s
}

type ModifyFlowProjectResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyFlowProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyFlowProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowProjectResponse) GoString() string {
	return s.String()
}

func (s *ModifyFlowProjectResponse) SetHeaders(v map[string]*string) *ModifyFlowProjectResponse {
	s.Headers = v
	return s
}

func (s *ModifyFlowProjectResponse) SetBody(v *ModifyFlowProjectResponseBody) *ModifyFlowProjectResponse {
	s.Body = v
	return s
}

type QueryKnoxUserPasswordRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserId          *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	OwnerId         *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s QueryKnoxUserPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryKnoxUserPasswordRequest) GoString() string {
	return s.String()
}

func (s *QueryKnoxUserPasswordRequest) SetResourceOwnerId(v int64) *QueryKnoxUserPasswordRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryKnoxUserPasswordRequest) SetRegionId(v string) *QueryKnoxUserPasswordRequest {
	s.RegionId = &v
	return s
}

func (s *QueryKnoxUserPasswordRequest) SetUserId(v string) *QueryKnoxUserPasswordRequest {
	s.UserId = &v
	return s
}

func (s *QueryKnoxUserPasswordRequest) SetOwnerId(v string) *QueryKnoxUserPasswordRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryKnoxUserPasswordRequest) SetClusterId(v string) *QueryKnoxUserPasswordRequest {
	s.ClusterId = &v
	return s
}

type QueryKnoxUserPasswordResponseBody struct {
	RequestId    *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserInfoList *QueryKnoxUserPasswordResponseBodyUserInfoList `json:"UserInfoList,omitempty" xml:"UserInfoList,omitempty" type:"Struct"`
}

func (s QueryKnoxUserPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryKnoxUserPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *QueryKnoxUserPasswordResponseBody) SetRequestId(v string) *QueryKnoxUserPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryKnoxUserPasswordResponseBody) SetUserInfoList(v *QueryKnoxUserPasswordResponseBodyUserInfoList) *QueryKnoxUserPasswordResponseBody {
	s.UserInfoList = v
	return s
}

type QueryKnoxUserPasswordResponseBodyUserInfoList struct {
	UserInfo []*QueryKnoxUserPasswordResponseBodyUserInfoListUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Repeated"`
}

func (s QueryKnoxUserPasswordResponseBodyUserInfoList) String() string {
	return tea.Prettify(s)
}

func (s QueryKnoxUserPasswordResponseBodyUserInfoList) GoString() string {
	return s.String()
}

func (s *QueryKnoxUserPasswordResponseBodyUserInfoList) SetUserInfo(v []*QueryKnoxUserPasswordResponseBodyUserInfoListUserInfo) *QueryKnoxUserPasswordResponseBodyUserInfoList {
	s.UserInfo = v
	return s
}

type QueryKnoxUserPasswordResponseBodyUserInfoListUserInfo struct {
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryKnoxUserPasswordResponseBodyUserInfoListUserInfo) String() string {
	return tea.Prettify(s)
}

func (s QueryKnoxUserPasswordResponseBodyUserInfoListUserInfo) GoString() string {
	return s.String()
}

func (s *QueryKnoxUserPasswordResponseBodyUserInfoListUserInfo) SetPassword(v string) *QueryKnoxUserPasswordResponseBodyUserInfoListUserInfo {
	s.Password = &v
	return s
}

func (s *QueryKnoxUserPasswordResponseBodyUserInfoListUserInfo) SetUserName(v string) *QueryKnoxUserPasswordResponseBodyUserInfoListUserInfo {
	s.UserName = &v
	return s
}

func (s *QueryKnoxUserPasswordResponseBodyUserInfoListUserInfo) SetUserId(v string) *QueryKnoxUserPasswordResponseBodyUserInfoListUserInfo {
	s.UserId = &v
	return s
}

type QueryKnoxUserPasswordResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryKnoxUserPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryKnoxUserPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryKnoxUserPasswordResponse) GoString() string {
	return s.String()
}

func (s *QueryKnoxUserPasswordResponse) SetHeaders(v map[string]*string) *QueryKnoxUserPasswordResponse {
	s.Headers = v
	return s
}

func (s *QueryKnoxUserPasswordResponse) SetBody(v *QueryKnoxUserPasswordResponseBody) *QueryKnoxUserPasswordResponse {
	s.Body = v
	return s
}

type DescribeFlowNodeInstanceLauncherLogRequest struct {
	Start          *int32  `json:"Start,omitempty" xml:"Start,omitempty"`
	Lines          *int32  `json:"Lines,omitempty" xml:"Lines,omitempty"`
	Offset         *int32  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	Length         *int32  `json:"Length,omitempty" xml:"Length,omitempty"`
	Reverse        *bool   `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	StartTime      *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime        *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	NodeInstanceId *string `json:"NodeInstanceId,omitempty" xml:"NodeInstanceId,omitempty"`
	ProjectId      *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeFlowNodeInstanceLauncherLogRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowNodeInstanceLauncherLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowNodeInstanceLauncherLogRequest) SetStart(v int32) *DescribeFlowNodeInstanceLauncherLogRequest {
	s.Start = &v
	return s
}

func (s *DescribeFlowNodeInstanceLauncherLogRequest) SetLines(v int32) *DescribeFlowNodeInstanceLauncherLogRequest {
	s.Lines = &v
	return s
}

func (s *DescribeFlowNodeInstanceLauncherLogRequest) SetOffset(v int32) *DescribeFlowNodeInstanceLauncherLogRequest {
	s.Offset = &v
	return s
}

func (s *DescribeFlowNodeInstanceLauncherLogRequest) SetLength(v int32) *DescribeFlowNodeInstanceLauncherLogRequest {
	s.Length = &v
	return s
}

func (s *DescribeFlowNodeInstanceLauncherLogRequest) SetReverse(v bool) *DescribeFlowNodeInstanceLauncherLogRequest {
	s.Reverse = &v
	return s
}

func (s *DescribeFlowNodeInstanceLauncherLogRequest) SetStartTime(v int64) *DescribeFlowNodeInstanceLauncherLogRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeFlowNodeInstanceLauncherLogRequest) SetEndTime(v int64) *DescribeFlowNodeInstanceLauncherLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeFlowNodeInstanceLauncherLogRequest) SetNodeInstanceId(v string) *DescribeFlowNodeInstanceLauncherLogRequest {
	s.NodeInstanceId = &v
	return s
}

func (s *DescribeFlowNodeInstanceLauncherLogRequest) SetProjectId(v string) *DescribeFlowNodeInstanceLauncherLogRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeFlowNodeInstanceLauncherLogRequest) SetRegionId(v string) *DescribeFlowNodeInstanceLauncherLogRequest {
	s.RegionId = &v
	return s
}

type DescribeFlowNodeInstanceLauncherLogResponseBody struct {
	LogEnd    *bool                                                     `json:"LogEnd,omitempty" xml:"LogEnd,omitempty"`
	RequestId *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	LogEntrys *DescribeFlowNodeInstanceLauncherLogResponseBodyLogEntrys `json:"LogEntrys,omitempty" xml:"LogEntrys,omitempty" type:"Struct"`
}

func (s DescribeFlowNodeInstanceLauncherLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowNodeInstanceLauncherLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowNodeInstanceLauncherLogResponseBody) SetLogEnd(v bool) *DescribeFlowNodeInstanceLauncherLogResponseBody {
	s.LogEnd = &v
	return s
}

func (s *DescribeFlowNodeInstanceLauncherLogResponseBody) SetRequestId(v string) *DescribeFlowNodeInstanceLauncherLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowNodeInstanceLauncherLogResponseBody) SetLogEntrys(v *DescribeFlowNodeInstanceLauncherLogResponseBodyLogEntrys) *DescribeFlowNodeInstanceLauncherLogResponseBody {
	s.LogEntrys = v
	return s
}

type DescribeFlowNodeInstanceLauncherLogResponseBodyLogEntrys struct {
	LogEntry []*DescribeFlowNodeInstanceLauncherLogResponseBodyLogEntrysLogEntry `json:"LogEntry,omitempty" xml:"LogEntry,omitempty" type:"Repeated"`
}

func (s DescribeFlowNodeInstanceLauncherLogResponseBodyLogEntrys) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowNodeInstanceLauncherLogResponseBodyLogEntrys) GoString() string {
	return s.String()
}

func (s *DescribeFlowNodeInstanceLauncherLogResponseBodyLogEntrys) SetLogEntry(v []*DescribeFlowNodeInstanceLauncherLogResponseBodyLogEntrysLogEntry) *DescribeFlowNodeInstanceLauncherLogResponseBodyLogEntrys {
	s.LogEntry = v
	return s
}

type DescribeFlowNodeInstanceLauncherLogResponseBodyLogEntrysLogEntry struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s DescribeFlowNodeInstanceLauncherLogResponseBodyLogEntrysLogEntry) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowNodeInstanceLauncherLogResponseBodyLogEntrysLogEntry) GoString() string {
	return s.String()
}

func (s *DescribeFlowNodeInstanceLauncherLogResponseBodyLogEntrysLogEntry) SetContent(v string) *DescribeFlowNodeInstanceLauncherLogResponseBodyLogEntrysLogEntry {
	s.Content = &v
	return s
}

type DescribeFlowNodeInstanceLauncherLogResponse struct {
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFlowNodeInstanceLauncherLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFlowNodeInstanceLauncherLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowNodeInstanceLauncherLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowNodeInstanceLauncherLogResponse) SetHeaders(v map[string]*string) *DescribeFlowNodeInstanceLauncherLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowNodeInstanceLauncherLogResponse) SetBody(v *DescribeFlowNodeInstanceLauncherLogResponseBody) *DescribeFlowNodeInstanceLauncherLogResponse {
	s.Body = v
	return s
}

type ListFlowRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProjectId  *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	JobId      *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id         *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Periodic   *bool   `json:"Periodic,omitempty" xml:"Periodic,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowRequest) GoString() string {
	return s.String()
}

func (s *ListFlowRequest) SetRegionId(v string) *ListFlowRequest {
	s.RegionId = &v
	return s
}

func (s *ListFlowRequest) SetProjectId(v string) *ListFlowRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFlowRequest) SetJobId(v string) *ListFlowRequest {
	s.JobId = &v
	return s
}

func (s *ListFlowRequest) SetName(v string) *ListFlowRequest {
	s.Name = &v
	return s
}

func (s *ListFlowRequest) SetId(v string) *ListFlowRequest {
	s.Id = &v
	return s
}

func (s *ListFlowRequest) SetClusterId(v string) *ListFlowRequest {
	s.ClusterId = &v
	return s
}

func (s *ListFlowRequest) SetStatus(v string) *ListFlowRequest {
	s.Status = &v
	return s
}

func (s *ListFlowRequest) SetPeriodic(v bool) *ListFlowRequest {
	s.Periodic = &v
	return s
}

func (s *ListFlowRequest) SetPageNumber(v int32) *ListFlowRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFlowRequest) SetPageSize(v int32) *ListFlowRequest {
	s.PageSize = &v
	return s
}

type ListFlowResponseBody struct {
	RequestId  *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total      *int32                    `json:"Total,omitempty" xml:"Total,omitempty"`
	Flow       *ListFlowResponseBodyFlow `json:"Flow,omitempty" xml:"Flow,omitempty" type:"Struct"`
}

func (s ListFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowResponseBody) SetRequestId(v string) *ListFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowResponseBody) SetPageNumber(v int32) *ListFlowResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListFlowResponseBody) SetPageSize(v int32) *ListFlowResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListFlowResponseBody) SetTotal(v int32) *ListFlowResponseBody {
	s.Total = &v
	return s
}

func (s *ListFlowResponseBody) SetFlow(v *ListFlowResponseBodyFlow) *ListFlowResponseBody {
	s.Flow = v
	return s
}

type ListFlowResponseBodyFlow struct {
	Flow []*ListFlowResponseBodyFlowFlow `json:"Flow,omitempty" xml:"Flow,omitempty" type:"Repeated"`
}

func (s ListFlowResponseBodyFlow) String() string {
	return tea.Prettify(s)
}

func (s ListFlowResponseBodyFlow) GoString() string {
	return s.String()
}

func (s *ListFlowResponseBodyFlow) SetFlow(v []*ListFlowResponseBodyFlowFlow) *ListFlowResponseBodyFlow {
	s.Flow = v
	return s
}

type ListFlowResponseBodyFlowFlow struct {
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type                    *string `json:"Type,omitempty" xml:"Type,omitempty"`
	AlertUserGroupBizId     *string `json:"AlertUserGroupBizId,omitempty" xml:"AlertUserGroupBizId,omitempty"`
	Periodic                *bool   `json:"Periodic,omitempty" xml:"Periodic,omitempty"`
	ProjectId               *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	HostName                *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	GmtModified             *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	CreateCluster           *bool   `json:"CreateCluster,omitempty" xml:"CreateCluster,omitempty"`
	StartSchedule           *int64  `json:"StartSchedule,omitempty" xml:"StartSchedule,omitempty"`
	EndSchedule             *int64  `json:"EndSchedule,omitempty" xml:"EndSchedule,omitempty"`
	Graph                   *string `json:"Graph,omitempty" xml:"Graph,omitempty"`
	AlertDingDingGroupBizId *string `json:"AlertDingDingGroupBizId,omitempty" xml:"AlertDingDingGroupBizId,omitempty"`
	GmtCreate               *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	CategoryId              *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CronExpr                *string `json:"CronExpr,omitempty" xml:"CronExpr,omitempty"`
	Name                    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id                      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	AlertConf               *string `json:"AlertConf,omitempty" xml:"AlertConf,omitempty"`
	ClusterId               *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s ListFlowResponseBodyFlowFlow) String() string {
	return tea.Prettify(s)
}

func (s ListFlowResponseBodyFlowFlow) GoString() string {
	return s.String()
}

func (s *ListFlowResponseBodyFlowFlow) SetStatus(v string) *ListFlowResponseBodyFlowFlow {
	s.Status = &v
	return s
}

func (s *ListFlowResponseBodyFlowFlow) SetType(v string) *ListFlowResponseBodyFlowFlow {
	s.Type = &v
	return s
}

func (s *ListFlowResponseBodyFlowFlow) SetAlertUserGroupBizId(v string) *ListFlowResponseBodyFlowFlow {
	s.AlertUserGroupBizId = &v
	return s
}

func (s *ListFlowResponseBodyFlowFlow) SetPeriodic(v bool) *ListFlowResponseBodyFlowFlow {
	s.Periodic = &v
	return s
}

func (s *ListFlowResponseBodyFlowFlow) SetProjectId(v string) *ListFlowResponseBodyFlowFlow {
	s.ProjectId = &v
	return s
}

func (s *ListFlowResponseBodyFlowFlow) SetHostName(v string) *ListFlowResponseBodyFlowFlow {
	s.HostName = &v
	return s
}

func (s *ListFlowResponseBodyFlowFlow) SetGmtModified(v int64) *ListFlowResponseBodyFlowFlow {
	s.GmtModified = &v
	return s
}

func (s *ListFlowResponseBodyFlowFlow) SetDescription(v string) *ListFlowResponseBodyFlowFlow {
	s.Description = &v
	return s
}

func (s *ListFlowResponseBodyFlowFlow) SetCreateCluster(v bool) *ListFlowResponseBodyFlowFlow {
	s.CreateCluster = &v
	return s
}

func (s *ListFlowResponseBodyFlowFlow) SetStartSchedule(v int64) *ListFlowResponseBodyFlowFlow {
	s.StartSchedule = &v
	return s
}

func (s *ListFlowResponseBodyFlowFlow) SetEndSchedule(v int64) *ListFlowResponseBodyFlowFlow {
	s.EndSchedule = &v
	return s
}

func (s *ListFlowResponseBodyFlowFlow) SetGraph(v string) *ListFlowResponseBodyFlowFlow {
	s.Graph = &v
	return s
}

func (s *ListFlowResponseBodyFlowFlow) SetAlertDingDingGroupBizId(v string) *ListFlowResponseBodyFlowFlow {
	s.AlertDingDingGroupBizId = &v
	return s
}

func (s *ListFlowResponseBodyFlowFlow) SetGmtCreate(v int64) *ListFlowResponseBodyFlowFlow {
	s.GmtCreate = &v
	return s
}

func (s *ListFlowResponseBodyFlowFlow) SetCategoryId(v string) *ListFlowResponseBodyFlowFlow {
	s.CategoryId = &v
	return s
}

func (s *ListFlowResponseBodyFlowFlow) SetCronExpr(v string) *ListFlowResponseBodyFlowFlow {
	s.CronExpr = &v
	return s
}

func (s *ListFlowResponseBodyFlowFlow) SetName(v string) *ListFlowResponseBodyFlowFlow {
	s.Name = &v
	return s
}

func (s *ListFlowResponseBodyFlowFlow) SetId(v string) *ListFlowResponseBodyFlowFlow {
	s.Id = &v
	return s
}

func (s *ListFlowResponseBodyFlowFlow) SetAlertConf(v string) *ListFlowResponseBodyFlowFlow {
	s.AlertConf = &v
	return s
}

func (s *ListFlowResponseBodyFlowFlow) SetClusterId(v string) *ListFlowResponseBodyFlowFlow {
	s.ClusterId = &v
	return s
}

type ListFlowResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowResponse) GoString() string {
	return s.String()
}

func (s *ListFlowResponse) SetHeaders(v map[string]*string) *ListFlowResponse {
	s.Headers = v
	return s
}

func (s *ListFlowResponse) SetBody(v *ListFlowResponseBody) *ListFlowResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// 资源类型
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// 集群ID列表
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// 解绑的标签键列表
	TagKeys []*string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
	// 是否解绑资源的所有标签
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetResourceOwnerId(v int64) *UntagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceGroupId(v string) *UntagResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceIds(v []*string) *UntagResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *UntagResourcesRequest) SetTagKeys(v []*string) *UntagResourcesRequest {
	s.TagKeys = v
	return s
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

type UntagResourcesResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 请求是否成功被处理
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 响应码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 错误码
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 响应消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UntagResourcesResponseBody) SetSuccess(v bool) *UntagResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *UntagResourcesResponseBody) SetCode(v string) *UntagResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *UntagResourcesResponseBody) SetErrorCode(v string) *UntagResourcesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UntagResourcesResponseBody) SetMessage(v string) *UntagResourcesResponseBody {
	s.Message = &v
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

type ListFlowClusterHostRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProjectId       *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListFlowClusterHostRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterHostRequest) GoString() string {
	return s.String()
}

func (s *ListFlowClusterHostRequest) SetRegionId(v string) *ListFlowClusterHostRequest {
	s.RegionId = &v
	return s
}

func (s *ListFlowClusterHostRequest) SetProjectId(v string) *ListFlowClusterHostRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFlowClusterHostRequest) SetClusterId(v string) *ListFlowClusterHostRequest {
	s.ClusterId = &v
	return s
}

func (s *ListFlowClusterHostRequest) SetResourceGroupId(v string) *ListFlowClusterHostRequest {
	s.ResourceGroupId = &v
	return s
}

type ListFlowClusterHostResponseBody struct {
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HostList  *ListFlowClusterHostResponseBodyHostList `json:"HostList,omitempty" xml:"HostList,omitempty" type:"Struct"`
}

func (s ListFlowClusterHostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterHostResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowClusterHostResponseBody) SetRequestId(v string) *ListFlowClusterHostResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowClusterHostResponseBody) SetHostList(v *ListFlowClusterHostResponseBodyHostList) *ListFlowClusterHostResponseBody {
	s.HostList = v
	return s
}

type ListFlowClusterHostResponseBodyHostList struct {
	Host []*ListFlowClusterHostResponseBodyHostListHost `json:"Host,omitempty" xml:"Host,omitempty" type:"Repeated"`
}

func (s ListFlowClusterHostResponseBodyHostList) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterHostResponseBodyHostList) GoString() string {
	return s.String()
}

func (s *ListFlowClusterHostResponseBodyHostList) SetHost(v []*ListFlowClusterHostResponseBodyHostListHost) *ListFlowClusterHostResponseBodyHostList {
	s.Host = v
	return s
}

type ListFlowClusterHostResponseBodyHostListHost struct {
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
	SerialNumber   *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	PrivateIp      *string `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	HostName       *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	InstanceType   *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	HostId         *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	HostInstanceId *string `json:"HostInstanceId,omitempty" xml:"HostInstanceId,omitempty"`
	Cpu            *int32  `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	PublicIp       *string `json:"PublicIp,omitempty" xml:"PublicIp,omitempty"`
	Memory         *int32  `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Role           *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s ListFlowClusterHostResponseBodyHostListHost) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterHostResponseBodyHostListHost) GoString() string {
	return s.String()
}

func (s *ListFlowClusterHostResponseBodyHostListHost) SetStatus(v string) *ListFlowClusterHostResponseBodyHostListHost {
	s.Status = &v
	return s
}

func (s *ListFlowClusterHostResponseBodyHostListHost) SetType(v string) *ListFlowClusterHostResponseBodyHostListHost {
	s.Type = &v
	return s
}

func (s *ListFlowClusterHostResponseBodyHostListHost) SetSerialNumber(v string) *ListFlowClusterHostResponseBodyHostListHost {
	s.SerialNumber = &v
	return s
}

func (s *ListFlowClusterHostResponseBodyHostListHost) SetPrivateIp(v string) *ListFlowClusterHostResponseBodyHostListHost {
	s.PrivateIp = &v
	return s
}

func (s *ListFlowClusterHostResponseBodyHostListHost) SetHostName(v string) *ListFlowClusterHostResponseBodyHostListHost {
	s.HostName = &v
	return s
}

func (s *ListFlowClusterHostResponseBodyHostListHost) SetInstanceType(v string) *ListFlowClusterHostResponseBodyHostListHost {
	s.InstanceType = &v
	return s
}

func (s *ListFlowClusterHostResponseBodyHostListHost) SetHostId(v string) *ListFlowClusterHostResponseBodyHostListHost {
	s.HostId = &v
	return s
}

func (s *ListFlowClusterHostResponseBodyHostListHost) SetHostInstanceId(v string) *ListFlowClusterHostResponseBodyHostListHost {
	s.HostInstanceId = &v
	return s
}

func (s *ListFlowClusterHostResponseBodyHostListHost) SetCpu(v int32) *ListFlowClusterHostResponseBodyHostListHost {
	s.Cpu = &v
	return s
}

func (s *ListFlowClusterHostResponseBodyHostListHost) SetPublicIp(v string) *ListFlowClusterHostResponseBodyHostListHost {
	s.PublicIp = &v
	return s
}

func (s *ListFlowClusterHostResponseBodyHostListHost) SetMemory(v int32) *ListFlowClusterHostResponseBodyHostListHost {
	s.Memory = &v
	return s
}

func (s *ListFlowClusterHostResponseBodyHostListHost) SetRole(v string) *ListFlowClusterHostResponseBodyHostListHost {
	s.Role = &v
	return s
}

type ListFlowClusterHostResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFlowClusterHostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowClusterHostResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterHostResponse) GoString() string {
	return s.String()
}

func (s *ListFlowClusterHostResponse) SetHeaders(v map[string]*string) *ListFlowClusterHostResponse {
	s.Headers = v
	return s
}

func (s *ListFlowClusterHostResponse) SetBody(v *ListFlowClusterHostResponseBody) *ListFlowClusterHostResponse {
	s.Body = v
	return s
}

type ListClusterOperationRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ServiceName     *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListClusterOperationRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClusterOperationRequest) GoString() string {
	return s.String()
}

func (s *ListClusterOperationRequest) SetResourceOwnerId(v int64) *ListClusterOperationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListClusterOperationRequest) SetRegionId(v string) *ListClusterOperationRequest {
	s.RegionId = &v
	return s
}

func (s *ListClusterOperationRequest) SetClusterId(v string) *ListClusterOperationRequest {
	s.ClusterId = &v
	return s
}

func (s *ListClusterOperationRequest) SetServiceName(v string) *ListClusterOperationRequest {
	s.ServiceName = &v
	return s
}

func (s *ListClusterOperationRequest) SetStatus(v string) *ListClusterOperationRequest {
	s.Status = &v
	return s
}

func (s *ListClusterOperationRequest) SetPageNumber(v int32) *ListClusterOperationRequest {
	s.PageNumber = &v
	return s
}

func (s *ListClusterOperationRequest) SetPageSize(v int32) *ListClusterOperationRequest {
	s.PageSize = &v
	return s
}

type ListClusterOperationResponseBody struct {
	PageSize             *int32                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId            *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber           *int32                                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount           *int32                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	ClusterOperationList *ListClusterOperationResponseBodyClusterOperationList `json:"ClusterOperationList,omitempty" xml:"ClusterOperationList,omitempty" type:"Struct"`
}

func (s ListClusterOperationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClusterOperationResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterOperationResponseBody) SetPageSize(v int32) *ListClusterOperationResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListClusterOperationResponseBody) SetRequestId(v string) *ListClusterOperationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterOperationResponseBody) SetPageNumber(v int32) *ListClusterOperationResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListClusterOperationResponseBody) SetTotalCount(v int32) *ListClusterOperationResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListClusterOperationResponseBody) SetClusterOperationList(v *ListClusterOperationResponseBodyClusterOperationList) *ListClusterOperationResponseBody {
	s.ClusterOperationList = v
	return s
}

type ListClusterOperationResponseBodyClusterOperationList struct {
	ClusterOperation []*ListClusterOperationResponseBodyClusterOperationListClusterOperation `json:"ClusterOperation,omitempty" xml:"ClusterOperation,omitempty" type:"Repeated"`
}

func (s ListClusterOperationResponseBodyClusterOperationList) String() string {
	return tea.Prettify(s)
}

func (s ListClusterOperationResponseBodyClusterOperationList) GoString() string {
	return s.String()
}

func (s *ListClusterOperationResponseBodyClusterOperationList) SetClusterOperation(v []*ListClusterOperationResponseBodyClusterOperationListClusterOperation) *ListClusterOperationResponseBodyClusterOperationList {
	s.ClusterOperation = v
	return s
}

type ListClusterOperationResponseBodyClusterOperationListClusterOperation struct {
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Comment       *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	OperationName *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	OperationId   *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	Duration      *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Percentage    *string `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
}

func (s ListClusterOperationResponseBodyClusterOperationListClusterOperation) String() string {
	return tea.Prettify(s)
}

func (s ListClusterOperationResponseBodyClusterOperationListClusterOperation) GoString() string {
	return s.String()
}

func (s *ListClusterOperationResponseBodyClusterOperationListClusterOperation) SetStatus(v string) *ListClusterOperationResponseBodyClusterOperationListClusterOperation {
	s.Status = &v
	return s
}

func (s *ListClusterOperationResponseBodyClusterOperationListClusterOperation) SetStartTime(v string) *ListClusterOperationResponseBodyClusterOperationListClusterOperation {
	s.StartTime = &v
	return s
}

func (s *ListClusterOperationResponseBodyClusterOperationListClusterOperation) SetComment(v string) *ListClusterOperationResponseBodyClusterOperationListClusterOperation {
	s.Comment = &v
	return s
}

func (s *ListClusterOperationResponseBodyClusterOperationListClusterOperation) SetOperationName(v string) *ListClusterOperationResponseBodyClusterOperationListClusterOperation {
	s.OperationName = &v
	return s
}

func (s *ListClusterOperationResponseBodyClusterOperationListClusterOperation) SetOperationId(v string) *ListClusterOperationResponseBodyClusterOperationListClusterOperation {
	s.OperationId = &v
	return s
}

func (s *ListClusterOperationResponseBodyClusterOperationListClusterOperation) SetDuration(v string) *ListClusterOperationResponseBodyClusterOperationListClusterOperation {
	s.Duration = &v
	return s
}

func (s *ListClusterOperationResponseBodyClusterOperationListClusterOperation) SetPercentage(v string) *ListClusterOperationResponseBodyClusterOperationListClusterOperation {
	s.Percentage = &v
	return s
}

type ListClusterOperationResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListClusterOperationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClusterOperationResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClusterOperationResponse) GoString() string {
	return s.String()
}

func (s *ListClusterOperationResponse) SetHeaders(v map[string]*string) *ListClusterOperationResponse {
	s.Headers = v
	return s
}

func (s *ListClusterOperationResponse) SetBody(v *ListClusterOperationResponseBody) *ListClusterOperationResponse {
	s.Body = v
	return s
}

type ListFlowEntitySnapshotRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Limit           *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentSize     *int32  `json:"CurrentSize,omitempty" xml:"CurrentSize,omitempty"`
	PageCount       *int32  `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	OrderField      *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	OrderMode       *string `json:"OrderMode,omitempty" xml:"OrderMode,omitempty"`
	CommitterId     *string `json:"CommitterId,omitempty" xml:"CommitterId,omitempty"`
	EntityType      *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	EntityGroupId   *string `json:"EntityGroupId,omitempty" xml:"EntityGroupId,omitempty"`
	EntityId        *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	Revision        *string `json:"Revision,omitempty" xml:"Revision,omitempty"`
}

func (s ListFlowEntitySnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowEntitySnapshotRequest) GoString() string {
	return s.String()
}

func (s *ListFlowEntitySnapshotRequest) SetResourceOwnerId(v int64) *ListFlowEntitySnapshotRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListFlowEntitySnapshotRequest) SetRegionId(v string) *ListFlowEntitySnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *ListFlowEntitySnapshotRequest) SetLimit(v int32) *ListFlowEntitySnapshotRequest {
	s.Limit = &v
	return s
}

func (s *ListFlowEntitySnapshotRequest) SetPageNumber(v int32) *ListFlowEntitySnapshotRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFlowEntitySnapshotRequest) SetPageSize(v int32) *ListFlowEntitySnapshotRequest {
	s.PageSize = &v
	return s
}

func (s *ListFlowEntitySnapshotRequest) SetCurrentSize(v int32) *ListFlowEntitySnapshotRequest {
	s.CurrentSize = &v
	return s
}

func (s *ListFlowEntitySnapshotRequest) SetPageCount(v int32) *ListFlowEntitySnapshotRequest {
	s.PageCount = &v
	return s
}

func (s *ListFlowEntitySnapshotRequest) SetOrderField(v string) *ListFlowEntitySnapshotRequest {
	s.OrderField = &v
	return s
}

func (s *ListFlowEntitySnapshotRequest) SetOrderMode(v string) *ListFlowEntitySnapshotRequest {
	s.OrderMode = &v
	return s
}

func (s *ListFlowEntitySnapshotRequest) SetCommitterId(v string) *ListFlowEntitySnapshotRequest {
	s.CommitterId = &v
	return s
}

func (s *ListFlowEntitySnapshotRequest) SetEntityType(v string) *ListFlowEntitySnapshotRequest {
	s.EntityType = &v
	return s
}

func (s *ListFlowEntitySnapshotRequest) SetEntityGroupId(v string) *ListFlowEntitySnapshotRequest {
	s.EntityGroupId = &v
	return s
}

func (s *ListFlowEntitySnapshotRequest) SetEntityId(v string) *ListFlowEntitySnapshotRequest {
	s.EntityId = &v
	return s
}

func (s *ListFlowEntitySnapshotRequest) SetRevision(v string) *ListFlowEntitySnapshotRequest {
	s.Revision = &v
	return s
}

type ListFlowEntitySnapshotResponseBody struct {
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Items      *ListFlowEntitySnapshotResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s ListFlowEntitySnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowEntitySnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowEntitySnapshotResponseBody) SetRequestId(v string) *ListFlowEntitySnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowEntitySnapshotResponseBody) SetPageNumber(v int32) *ListFlowEntitySnapshotResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListFlowEntitySnapshotResponseBody) SetPageSize(v int32) *ListFlowEntitySnapshotResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListFlowEntitySnapshotResponseBody) SetTotalCount(v int32) *ListFlowEntitySnapshotResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListFlowEntitySnapshotResponseBody) SetItems(v *ListFlowEntitySnapshotResponseBodyItems) *ListFlowEntitySnapshotResponseBody {
	s.Items = v
	return s
}

type ListFlowEntitySnapshotResponseBodyItems struct {
	Item []*ListFlowEntitySnapshotResponseBodyItemsItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s ListFlowEntitySnapshotResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s ListFlowEntitySnapshotResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListFlowEntitySnapshotResponseBodyItems) SetItem(v []*ListFlowEntitySnapshotResponseBodyItemsItem) *ListFlowEntitySnapshotResponseBodyItems {
	s.Item = v
	return s
}

type ListFlowEntitySnapshotResponseBodyItemsItem struct {
	Active        *bool   `json:"Active,omitempty" xml:"Active,omitempty"`
	Data          *string `json:"Data,omitempty" xml:"Data,omitempty"`
	EntityId      *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	EntityType    *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	GmtCreate     *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty"`
	EntityGroupId *string `json:"EntityGroupId,omitempty" xml:"EntityGroupId,omitempty"`
	CommitterId   *string `json:"CommitterId,omitempty" xml:"CommitterId,omitempty"`
	Revision      *string `json:"Revision,omitempty" xml:"Revision,omitempty"`
}

func (s ListFlowEntitySnapshotResponseBodyItemsItem) String() string {
	return tea.Prettify(s)
}

func (s ListFlowEntitySnapshotResponseBodyItemsItem) GoString() string {
	return s.String()
}

func (s *ListFlowEntitySnapshotResponseBodyItemsItem) SetActive(v bool) *ListFlowEntitySnapshotResponseBodyItemsItem {
	s.Active = &v
	return s
}

func (s *ListFlowEntitySnapshotResponseBodyItemsItem) SetData(v string) *ListFlowEntitySnapshotResponseBodyItemsItem {
	s.Data = &v
	return s
}

func (s *ListFlowEntitySnapshotResponseBodyItemsItem) SetEntityId(v string) *ListFlowEntitySnapshotResponseBodyItemsItem {
	s.EntityId = &v
	return s
}

func (s *ListFlowEntitySnapshotResponseBodyItemsItem) SetEntityType(v string) *ListFlowEntitySnapshotResponseBodyItemsItem {
	s.EntityType = &v
	return s
}

func (s *ListFlowEntitySnapshotResponseBodyItemsItem) SetUserId(v string) *ListFlowEntitySnapshotResponseBodyItemsItem {
	s.UserId = &v
	return s
}

func (s *ListFlowEntitySnapshotResponseBodyItemsItem) SetGmtCreate(v int64) *ListFlowEntitySnapshotResponseBodyItemsItem {
	s.GmtCreate = &v
	return s
}

func (s *ListFlowEntitySnapshotResponseBodyItemsItem) SetMessage(v string) *ListFlowEntitySnapshotResponseBodyItemsItem {
	s.Message = &v
	return s
}

func (s *ListFlowEntitySnapshotResponseBodyItemsItem) SetEntityGroupId(v string) *ListFlowEntitySnapshotResponseBodyItemsItem {
	s.EntityGroupId = &v
	return s
}

func (s *ListFlowEntitySnapshotResponseBodyItemsItem) SetCommitterId(v string) *ListFlowEntitySnapshotResponseBodyItemsItem {
	s.CommitterId = &v
	return s
}

func (s *ListFlowEntitySnapshotResponseBodyItemsItem) SetRevision(v string) *ListFlowEntitySnapshotResponseBodyItemsItem {
	s.Revision = &v
	return s
}

type ListFlowEntitySnapshotResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFlowEntitySnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowEntitySnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowEntitySnapshotResponse) GoString() string {
	return s.String()
}

func (s *ListFlowEntitySnapshotResponse) SetHeaders(v map[string]*string) *ListFlowEntitySnapshotResponse {
	s.Headers = v
	return s
}

func (s *ListFlowEntitySnapshotResponse) SetBody(v *ListFlowEntitySnapshotResponseBody) *ListFlowEntitySnapshotResponse {
	s.Body = v
	return s
}

type DeleteClusterTemplateRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	BizId           *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DeleteClusterTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterTemplateRequest) SetResourceOwnerId(v int64) *DeleteClusterTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteClusterTemplateRequest) SetRegionId(v string) *DeleteClusterTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteClusterTemplateRequest) SetBizId(v string) *DeleteClusterTemplateRequest {
	s.BizId = &v
	return s
}

func (s *DeleteClusterTemplateRequest) SetResourceGroupId(v string) *DeleteClusterTemplateRequest {
	s.ResourceGroupId = &v
	return s
}

type DeleteClusterTemplateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteClusterTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClusterTemplateResponseBody) SetRequestId(v string) *DeleteClusterTemplateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteClusterTemplateResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteClusterTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteClusterTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteClusterTemplateResponse) SetHeaders(v map[string]*string) *DeleteClusterTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteClusterTemplateResponse) SetBody(v *DeleteClusterTemplateResponseBody) *DeleteClusterTemplateResponse {
	s.Body = v
	return s
}

type CancelOrderRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s CancelOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelOrderRequest) GoString() string {
	return s.String()
}

func (s *CancelOrderRequest) SetResourceOwnerId(v int64) *CancelOrderRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CancelOrderRequest) SetRegionId(v string) *CancelOrderRequest {
	s.RegionId = &v
	return s
}

func (s *CancelOrderRequest) SetClusterId(v string) *CancelOrderRequest {
	s.ClusterId = &v
	return s
}

type CancelOrderResponseBody struct {
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CancelOrderResponseBody) SetClusterId(v string) *CancelOrderResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CancelOrderResponseBody) SetRequestId(v string) *CancelOrderResponseBody {
	s.RequestId = &v
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

type CloneFlowJobRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CloneFlowJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CloneFlowJobRequest) GoString() string {
	return s.String()
}

func (s *CloneFlowJobRequest) SetProjectId(v string) *CloneFlowJobRequest {
	s.ProjectId = &v
	return s
}

func (s *CloneFlowJobRequest) SetId(v string) *CloneFlowJobRequest {
	s.Id = &v
	return s
}

func (s *CloneFlowJobRequest) SetName(v string) *CloneFlowJobRequest {
	s.Name = &v
	return s
}

func (s *CloneFlowJobRequest) SetRegionId(v string) *CloneFlowJobRequest {
	s.RegionId = &v
	return s
}

type CloneFlowJobResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloneFlowJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloneFlowJobResponseBody) GoString() string {
	return s.String()
}

func (s *CloneFlowJobResponseBody) SetId(v string) *CloneFlowJobResponseBody {
	s.Id = &v
	return s
}

func (s *CloneFlowJobResponseBody) SetRequestId(v string) *CloneFlowJobResponseBody {
	s.RequestId = &v
	return s
}

type CloneFlowJobResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CloneFlowJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CloneFlowJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CloneFlowJobResponse) GoString() string {
	return s.String()
}

func (s *CloneFlowJobResponse) SetHeaders(v map[string]*string) *CloneFlowJobResponse {
	s.Headers = v
	return s
}

func (s *CloneFlowJobResponse) SetBody(v *CloneFlowJobResponseBody) *CloneFlowJobResponse {
	s.Body = v
	return s
}

type StartFlowRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProjectId      *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	FlowInstanceId *string `json:"FlowInstanceId,omitempty" xml:"FlowInstanceId,omitempty"`
}

func (s StartFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s StartFlowRequest) GoString() string {
	return s.String()
}

func (s *StartFlowRequest) SetRegionId(v string) *StartFlowRequest {
	s.RegionId = &v
	return s
}

func (s *StartFlowRequest) SetProjectId(v string) *StartFlowRequest {
	s.ProjectId = &v
	return s
}

func (s *StartFlowRequest) SetFlowInstanceId(v string) *StartFlowRequest {
	s.FlowInstanceId = &v
	return s
}

type StartFlowResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartFlowResponseBody) GoString() string {
	return s.String()
}

func (s *StartFlowResponseBody) SetData(v bool) *StartFlowResponseBody {
	s.Data = &v
	return s
}

func (s *StartFlowResponseBody) SetRequestId(v string) *StartFlowResponseBody {
	s.RequestId = &v
	return s
}

type StartFlowResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s StartFlowResponse) GoString() string {
	return s.String()
}

func (s *StartFlowResponse) SetHeaders(v map[string]*string) *StartFlowResponse {
	s.Headers = v
	return s
}

func (s *StartFlowResponse) SetBody(v *StartFlowResponseBody) *StartFlowResponse {
	s.Body = v
	return s
}

type CreateFlowJobRequest struct {
	RegionId        *string                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProjectId       *string                             `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Name            *string                             `json:"Name,omitempty" xml:"Name,omitempty"`
	Description     *string                             `json:"Description,omitempty" xml:"Description,omitempty"`
	Type            *string                             `json:"Type,omitempty" xml:"Type,omitempty"`
	FailAct         *string                             `json:"FailAct,omitempty" xml:"FailAct,omitempty"`
	RetryPolicy     *string                             `json:"RetryPolicy,omitempty" xml:"RetryPolicy,omitempty"`
	Params          *string                             `json:"Params,omitempty" xml:"Params,omitempty"`
	ParamConf       *string                             `json:"ParamConf,omitempty" xml:"ParamConf,omitempty"`
	CustomVariables *string                             `json:"CustomVariables,omitempty" xml:"CustomVariables,omitempty"`
	EnvConf         *string                             `json:"EnvConf,omitempty" xml:"EnvConf,omitempty"`
	RunConf         *string                             `json:"RunConf,omitempty" xml:"RunConf,omitempty"`
	MonitorConf     *string                             `json:"MonitorConf,omitempty" xml:"MonitorConf,omitempty"`
	Mode            *string                             `json:"Mode,omitempty" xml:"Mode,omitempty"`
	ParentCategory  *string                             `json:"ParentCategory,omitempty" xml:"ParentCategory,omitempty"`
	Adhoc           *bool                               `json:"Adhoc,omitempty" xml:"Adhoc,omitempty"`
	ClusterId       *string                             `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	AlertConf       *string                             `json:"AlertConf,omitempty" xml:"AlertConf,omitempty"`
	ClientToken     *string                             `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ResourceList    []*CreateFlowJobRequestResourceList `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Repeated"`
}

func (s CreateFlowJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowJobRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowJobRequest) SetRegionId(v string) *CreateFlowJobRequest {
	s.RegionId = &v
	return s
}

func (s *CreateFlowJobRequest) SetProjectId(v string) *CreateFlowJobRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateFlowJobRequest) SetName(v string) *CreateFlowJobRequest {
	s.Name = &v
	return s
}

func (s *CreateFlowJobRequest) SetDescription(v string) *CreateFlowJobRequest {
	s.Description = &v
	return s
}

func (s *CreateFlowJobRequest) SetType(v string) *CreateFlowJobRequest {
	s.Type = &v
	return s
}

func (s *CreateFlowJobRequest) SetFailAct(v string) *CreateFlowJobRequest {
	s.FailAct = &v
	return s
}

func (s *CreateFlowJobRequest) SetRetryPolicy(v string) *CreateFlowJobRequest {
	s.RetryPolicy = &v
	return s
}

func (s *CreateFlowJobRequest) SetParams(v string) *CreateFlowJobRequest {
	s.Params = &v
	return s
}

func (s *CreateFlowJobRequest) SetParamConf(v string) *CreateFlowJobRequest {
	s.ParamConf = &v
	return s
}

func (s *CreateFlowJobRequest) SetCustomVariables(v string) *CreateFlowJobRequest {
	s.CustomVariables = &v
	return s
}

func (s *CreateFlowJobRequest) SetEnvConf(v string) *CreateFlowJobRequest {
	s.EnvConf = &v
	return s
}

func (s *CreateFlowJobRequest) SetRunConf(v string) *CreateFlowJobRequest {
	s.RunConf = &v
	return s
}

func (s *CreateFlowJobRequest) SetMonitorConf(v string) *CreateFlowJobRequest {
	s.MonitorConf = &v
	return s
}

func (s *CreateFlowJobRequest) SetMode(v string) *CreateFlowJobRequest {
	s.Mode = &v
	return s
}

func (s *CreateFlowJobRequest) SetParentCategory(v string) *CreateFlowJobRequest {
	s.ParentCategory = &v
	return s
}

func (s *CreateFlowJobRequest) SetAdhoc(v bool) *CreateFlowJobRequest {
	s.Adhoc = &v
	return s
}

func (s *CreateFlowJobRequest) SetClusterId(v string) *CreateFlowJobRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateFlowJobRequest) SetAlertConf(v string) *CreateFlowJobRequest {
	s.AlertConf = &v
	return s
}

func (s *CreateFlowJobRequest) SetClientToken(v string) *CreateFlowJobRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateFlowJobRequest) SetResourceList(v []*CreateFlowJobRequestResourceList) *CreateFlowJobRequest {
	s.ResourceList = v
	return s
}

type CreateFlowJobRequestResourceList struct {
	Path  *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
}

func (s CreateFlowJobRequestResourceList) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowJobRequestResourceList) GoString() string {
	return s.String()
}

func (s *CreateFlowJobRequestResourceList) SetPath(v string) *CreateFlowJobRequestResourceList {
	s.Path = &v
	return s
}

func (s *CreateFlowJobRequestResourceList) SetAlias(v string) *CreateFlowJobRequestResourceList {
	s.Alias = &v
	return s
}

type CreateFlowJobResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFlowJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlowJobResponseBody) SetId(v string) *CreateFlowJobResponseBody {
	s.Id = &v
	return s
}

func (s *CreateFlowJobResponseBody) SetRequestId(v string) *CreateFlowJobResponseBody {
	s.RequestId = &v
	return s
}

type CreateFlowJobResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFlowJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFlowJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowJobResponse) GoString() string {
	return s.String()
}

func (s *CreateFlowJobResponse) SetHeaders(v map[string]*string) *CreateFlowJobResponse {
	s.Headers = v
	return s
}

func (s *CreateFlowJobResponse) SetBody(v *CreateFlowJobResponseBody) *CreateFlowJobResponse {
	s.Body = v
	return s
}

type DeleteFlowCategoryRequest struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteFlowCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowCategoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteFlowCategoryRequest) SetId(v string) *DeleteFlowCategoryRequest {
	s.Id = &v
	return s
}

func (s *DeleteFlowCategoryRequest) SetRegionId(v string) *DeleteFlowCategoryRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteFlowCategoryRequest) SetProjectId(v string) *DeleteFlowCategoryRequest {
	s.ProjectId = &v
	return s
}

type DeleteFlowCategoryResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFlowCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFlowCategoryResponseBody) SetData(v bool) *DeleteFlowCategoryResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteFlowCategoryResponseBody) SetRequestId(v string) *DeleteFlowCategoryResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFlowCategoryResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFlowCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFlowCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowCategoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteFlowCategoryResponse) SetHeaders(v map[string]*string) *DeleteFlowCategoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteFlowCategoryResponse) SetBody(v *DeleteFlowCategoryResponseBody) *DeleteFlowCategoryResponse {
	s.Body = v
	return s
}

type DeleteFlowEditLockRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	EntityId        *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
}

func (s DeleteFlowEditLockRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowEditLockRequest) GoString() string {
	return s.String()
}

func (s *DeleteFlowEditLockRequest) SetResourceOwnerId(v int64) *DeleteFlowEditLockRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteFlowEditLockRequest) SetRegionId(v string) *DeleteFlowEditLockRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteFlowEditLockRequest) SetEntityId(v string) *DeleteFlowEditLockRequest {
	s.EntityId = &v
	return s
}

type DeleteFlowEditLockResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFlowEditLockResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowEditLockResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFlowEditLockResponseBody) SetData(v bool) *DeleteFlowEditLockResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteFlowEditLockResponseBody) SetRequestId(v string) *DeleteFlowEditLockResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFlowEditLockResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFlowEditLockResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFlowEditLockResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowEditLockResponse) GoString() string {
	return s.String()
}

func (s *DeleteFlowEditLockResponse) SetHeaders(v map[string]*string) *DeleteFlowEditLockResponse {
	s.Headers = v
	return s
}

func (s *DeleteFlowEditLockResponse) SetBody(v *DeleteFlowEditLockResponseBody) *DeleteFlowEditLockResponse {
	s.Body = v
	return s
}

type ResizeClusterRequest struct {
	RegionId          *string                                  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClusterId         *string                                  `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	AutoPayOrder      *bool                                    `json:"AutoPayOrder,omitempty" xml:"AutoPayOrder,omitempty"`
	VswitchId         *string                                  `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	IsOpenPublicIp    *bool                                    `json:"IsOpenPublicIp,omitempty" xml:"IsOpenPublicIp,omitempty"`
	HostComponentInfo []*ResizeClusterRequestHostComponentInfo `json:"HostComponentInfo,omitempty" xml:"HostComponentInfo,omitempty" type:"Repeated"`
	HostGroup         []*ResizeClusterRequestHostGroup         `json:"HostGroup,omitempty" xml:"HostGroup,omitempty" type:"Repeated"`
	PromotionInfo     []*ResizeClusterRequestPromotionInfo     `json:"PromotionInfo,omitempty" xml:"PromotionInfo,omitempty" type:"Repeated"`
}

func (s ResizeClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s ResizeClusterRequest) GoString() string {
	return s.String()
}

func (s *ResizeClusterRequest) SetRegionId(v string) *ResizeClusterRequest {
	s.RegionId = &v
	return s
}

func (s *ResizeClusterRequest) SetClusterId(v string) *ResizeClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *ResizeClusterRequest) SetAutoPayOrder(v bool) *ResizeClusterRequest {
	s.AutoPayOrder = &v
	return s
}

func (s *ResizeClusterRequest) SetVswitchId(v string) *ResizeClusterRequest {
	s.VswitchId = &v
	return s
}

func (s *ResizeClusterRequest) SetIsOpenPublicIp(v bool) *ResizeClusterRequest {
	s.IsOpenPublicIp = &v
	return s
}

func (s *ResizeClusterRequest) SetHostComponentInfo(v []*ResizeClusterRequestHostComponentInfo) *ResizeClusterRequest {
	s.HostComponentInfo = v
	return s
}

func (s *ResizeClusterRequest) SetHostGroup(v []*ResizeClusterRequestHostGroup) *ResizeClusterRequest {
	s.HostGroup = v
	return s
}

func (s *ResizeClusterRequest) SetPromotionInfo(v []*ResizeClusterRequestPromotionInfo) *ResizeClusterRequest {
	s.PromotionInfo = v
	return s
}

type ResizeClusterRequestHostComponentInfo struct {
	ComponentNameList []*string `json:"ComponentNameList,omitempty" xml:"ComponentNameList,omitempty" type:"Repeated"`
	HostName          *string   `json:"HostName,omitempty" xml:"HostName,omitempty"`
	ServiceName       *string   `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s ResizeClusterRequestHostComponentInfo) String() string {
	return tea.Prettify(s)
}

func (s ResizeClusterRequestHostComponentInfo) GoString() string {
	return s.String()
}

func (s *ResizeClusterRequestHostComponentInfo) SetComponentNameList(v []*string) *ResizeClusterRequestHostComponentInfo {
	s.ComponentNameList = v
	return s
}

func (s *ResizeClusterRequestHostComponentInfo) SetHostName(v string) *ResizeClusterRequestHostComponentInfo {
	s.HostName = &v
	return s
}

func (s *ResizeClusterRequestHostComponentInfo) SetServiceName(v string) *ResizeClusterRequestHostComponentInfo {
	s.ServiceName = &v
	return s
}

type ResizeClusterRequestHostGroup struct {
	SysDiskCapacity *int32  `json:"SysDiskCapacity,omitempty" xml:"SysDiskCapacity,omitempty"`
	HostGroupType   *string `json:"HostGroupType,omitempty" xml:"HostGroupType,omitempty"`
	Comment         *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	VswitchId       *int32  `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	SysDiskType     *string `json:"SysDiskType,omitempty" xml:"SysDiskType,omitempty"`
	AutoRenew       *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	ChargeType      *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	DiskType        *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	HostGroupId     *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	InstanceType    *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	HostKeyPairName *string `json:"HostKeyPairName,omitempty" xml:"HostKeyPairName,omitempty"`
	DiskCount       *int32  `json:"DiskCount,omitempty" xml:"DiskCount,omitempty"`
	CreateType      *string `json:"CreateType,omitempty" xml:"CreateType,omitempty"`
	Period          *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	HostPassword    *string `json:"HostPassword,omitempty" xml:"HostPassword,omitempty"`
	DiskCapacity    *int32  `json:"DiskCapacity,omitempty" xml:"DiskCapacity,omitempty"`
	NodeCount       *int32  `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	HostGroupName   *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s ResizeClusterRequestHostGroup) String() string {
	return tea.Prettify(s)
}

func (s ResizeClusterRequestHostGroup) GoString() string {
	return s.String()
}

func (s *ResizeClusterRequestHostGroup) SetSysDiskCapacity(v int32) *ResizeClusterRequestHostGroup {
	s.SysDiskCapacity = &v
	return s
}

func (s *ResizeClusterRequestHostGroup) SetHostGroupType(v string) *ResizeClusterRequestHostGroup {
	s.HostGroupType = &v
	return s
}

func (s *ResizeClusterRequestHostGroup) SetComment(v string) *ResizeClusterRequestHostGroup {
	s.Comment = &v
	return s
}

func (s *ResizeClusterRequestHostGroup) SetVswitchId(v int32) *ResizeClusterRequestHostGroup {
	s.VswitchId = &v
	return s
}

func (s *ResizeClusterRequestHostGroup) SetSysDiskType(v string) *ResizeClusterRequestHostGroup {
	s.SysDiskType = &v
	return s
}

func (s *ResizeClusterRequestHostGroup) SetAutoRenew(v bool) *ResizeClusterRequestHostGroup {
	s.AutoRenew = &v
	return s
}

func (s *ResizeClusterRequestHostGroup) SetChargeType(v string) *ResizeClusterRequestHostGroup {
	s.ChargeType = &v
	return s
}

func (s *ResizeClusterRequestHostGroup) SetDiskType(v string) *ResizeClusterRequestHostGroup {
	s.DiskType = &v
	return s
}

func (s *ResizeClusterRequestHostGroup) SetHostGroupId(v string) *ResizeClusterRequestHostGroup {
	s.HostGroupId = &v
	return s
}

func (s *ResizeClusterRequestHostGroup) SetInstanceType(v string) *ResizeClusterRequestHostGroup {
	s.InstanceType = &v
	return s
}

func (s *ResizeClusterRequestHostGroup) SetHostKeyPairName(v string) *ResizeClusterRequestHostGroup {
	s.HostKeyPairName = &v
	return s
}

func (s *ResizeClusterRequestHostGroup) SetDiskCount(v int32) *ResizeClusterRequestHostGroup {
	s.DiskCount = &v
	return s
}

func (s *ResizeClusterRequestHostGroup) SetCreateType(v string) *ResizeClusterRequestHostGroup {
	s.CreateType = &v
	return s
}

func (s *ResizeClusterRequestHostGroup) SetPeriod(v int32) *ResizeClusterRequestHostGroup {
	s.Period = &v
	return s
}

func (s *ResizeClusterRequestHostGroup) SetHostPassword(v string) *ResizeClusterRequestHostGroup {
	s.HostPassword = &v
	return s
}

func (s *ResizeClusterRequestHostGroup) SetDiskCapacity(v int32) *ResizeClusterRequestHostGroup {
	s.DiskCapacity = &v
	return s
}

func (s *ResizeClusterRequestHostGroup) SetNodeCount(v int32) *ResizeClusterRequestHostGroup {
	s.NodeCount = &v
	return s
}

func (s *ResizeClusterRequestHostGroup) SetHostGroupName(v string) *ResizeClusterRequestHostGroup {
	s.HostGroupName = &v
	return s
}

func (s *ResizeClusterRequestHostGroup) SetClusterId(v string) *ResizeClusterRequestHostGroup {
	s.ClusterId = &v
	return s
}

type ResizeClusterRequestPromotionInfo struct {
	PromotionOptionNo   *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
	PromotionOptionCode *string `json:"PromotionOptionCode,omitempty" xml:"PromotionOptionCode,omitempty"`
	ProductCode         *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s ResizeClusterRequestPromotionInfo) String() string {
	return tea.Prettify(s)
}

func (s ResizeClusterRequestPromotionInfo) GoString() string {
	return s.String()
}

func (s *ResizeClusterRequestPromotionInfo) SetPromotionOptionNo(v string) *ResizeClusterRequestPromotionInfo {
	s.PromotionOptionNo = &v
	return s
}

func (s *ResizeClusterRequestPromotionInfo) SetPromotionOptionCode(v string) *ResizeClusterRequestPromotionInfo {
	s.PromotionOptionCode = &v
	return s
}

func (s *ResizeClusterRequestPromotionInfo) SetProductCode(v string) *ResizeClusterRequestPromotionInfo {
	s.ProductCode = &v
	return s
}

type ResizeClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s ResizeClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResizeClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ResizeClusterResponseBody) SetRequestId(v string) *ResizeClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResizeClusterResponseBody) SetClusterId(v string) *ResizeClusterResponseBody {
	s.ClusterId = &v
	return s
}

type ResizeClusterResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResizeClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResizeClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s ResizeClusterResponse) GoString() string {
	return s.String()
}

func (s *ResizeClusterResponse) SetHeaders(v map[string]*string) *ResizeClusterResponse {
	s.Headers = v
	return s
}

func (s *ResizeClusterResponse) SetBody(v *ResizeClusterResponseBody) *ResizeClusterResponse {
	s.Body = v
	return s
}

type DescribeMetaTablePreviewTaskRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TaskId          *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeMetaTablePreviewTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaTablePreviewTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetaTablePreviewTaskRequest) SetResourceOwnerId(v int64) *DescribeMetaTablePreviewTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeMetaTablePreviewTaskRequest) SetRegionId(v string) *DescribeMetaTablePreviewTaskRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMetaTablePreviewTaskRequest) SetTaskId(v string) *DescribeMetaTablePreviewTaskRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeMetaTablePreviewTaskRequest) SetResourceGroupId(v string) *DescribeMetaTablePreviewTaskRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeMetaTablePreviewTaskResponseBody struct {
	EndTime     *int64                                        `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime   *int64                                        `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskStatus  *string                                       `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskProcess *int32                                        `json:"TaskProcess,omitempty" xml:"TaskProcess,omitempty"`
	ExecuteTime *int64                                        `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	TaskId      *string                                       `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Data        *DescribeMetaTablePreviewTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeMetaTablePreviewTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaTablePreviewTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetaTablePreviewTaskResponseBody) SetEndTime(v int64) *DescribeMetaTablePreviewTaskResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeMetaTablePreviewTaskResponseBody) SetStartTime(v int64) *DescribeMetaTablePreviewTaskResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeMetaTablePreviewTaskResponseBody) SetRequestId(v string) *DescribeMetaTablePreviewTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMetaTablePreviewTaskResponseBody) SetTaskStatus(v string) *DescribeMetaTablePreviewTaskResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *DescribeMetaTablePreviewTaskResponseBody) SetTaskProcess(v int32) *DescribeMetaTablePreviewTaskResponseBody {
	s.TaskProcess = &v
	return s
}

func (s *DescribeMetaTablePreviewTaskResponseBody) SetExecuteTime(v int64) *DescribeMetaTablePreviewTaskResponseBody {
	s.ExecuteTime = &v
	return s
}

func (s *DescribeMetaTablePreviewTaskResponseBody) SetTaskId(v string) *DescribeMetaTablePreviewTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *DescribeMetaTablePreviewTaskResponseBody) SetData(v *DescribeMetaTablePreviewTaskResponseBodyData) *DescribeMetaTablePreviewTaskResponseBody {
	s.Data = v
	return s
}

type DescribeMetaTablePreviewTaskResponseBodyData struct {
	Rows    *DescribeMetaTablePreviewTaskResponseBodyDataRows    `json:"Rows,omitempty" xml:"Rows,omitempty" type:"Struct"`
	Headers *DescribeMetaTablePreviewTaskResponseBodyDataHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Struct"`
}

func (s DescribeMetaTablePreviewTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaTablePreviewTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMetaTablePreviewTaskResponseBodyData) SetRows(v *DescribeMetaTablePreviewTaskResponseBodyDataRows) *DescribeMetaTablePreviewTaskResponseBodyData {
	s.Rows = v
	return s
}

func (s *DescribeMetaTablePreviewTaskResponseBodyData) SetHeaders(v *DescribeMetaTablePreviewTaskResponseBodyDataHeaders) *DescribeMetaTablePreviewTaskResponseBodyData {
	s.Headers = v
	return s
}

type DescribeMetaTablePreviewTaskResponseBodyDataRows struct {
	Row []*DescribeMetaTablePreviewTaskResponseBodyDataRowsRow `json:"Row,omitempty" xml:"Row,omitempty" type:"Repeated"`
}

func (s DescribeMetaTablePreviewTaskResponseBodyDataRows) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaTablePreviewTaskResponseBodyDataRows) GoString() string {
	return s.String()
}

func (s *DescribeMetaTablePreviewTaskResponseBodyDataRows) SetRow(v []*DescribeMetaTablePreviewTaskResponseBodyDataRowsRow) *DescribeMetaTablePreviewTaskResponseBodyDataRows {
	s.Row = v
	return s
}

type DescribeMetaTablePreviewTaskResponseBodyDataRowsRow struct {
	Columns *DescribeMetaTablePreviewTaskResponseBodyDataRowsRowColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Struct"`
}

func (s DescribeMetaTablePreviewTaskResponseBodyDataRowsRow) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaTablePreviewTaskResponseBodyDataRowsRow) GoString() string {
	return s.String()
}

func (s *DescribeMetaTablePreviewTaskResponseBodyDataRowsRow) SetColumns(v *DescribeMetaTablePreviewTaskResponseBodyDataRowsRowColumns) *DescribeMetaTablePreviewTaskResponseBodyDataRowsRow {
	s.Columns = v
	return s
}

type DescribeMetaTablePreviewTaskResponseBodyDataRowsRowColumns struct {
	Column []*string `json:"Column,omitempty" xml:"Column,omitempty" type:"Repeated"`
}

func (s DescribeMetaTablePreviewTaskResponseBodyDataRowsRowColumns) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaTablePreviewTaskResponseBodyDataRowsRowColumns) GoString() string {
	return s.String()
}

func (s *DescribeMetaTablePreviewTaskResponseBodyDataRowsRowColumns) SetColumn(v []*string) *DescribeMetaTablePreviewTaskResponseBodyDataRowsRowColumns {
	s.Column = v
	return s
}

type DescribeMetaTablePreviewTaskResponseBodyDataHeaders struct {
	Header []*string `json:"Header,omitempty" xml:"Header,omitempty" type:"Repeated"`
}

func (s DescribeMetaTablePreviewTaskResponseBodyDataHeaders) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaTablePreviewTaskResponseBodyDataHeaders) GoString() string {
	return s.String()
}

func (s *DescribeMetaTablePreviewTaskResponseBodyDataHeaders) SetHeader(v []*string) *DescribeMetaTablePreviewTaskResponseBodyDataHeaders {
	s.Header = v
	return s
}

type DescribeMetaTablePreviewTaskResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMetaTablePreviewTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMetaTablePreviewTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMetaTablePreviewTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeMetaTablePreviewTaskResponse) SetHeaders(v map[string]*string) *DescribeMetaTablePreviewTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeMetaTablePreviewTaskResponse) SetBody(v *DescribeMetaTablePreviewTaskResponseBody) *DescribeMetaTablePreviewTaskResponse {
	s.Body = v
	return s
}

type ListClusterServiceConfigHistoryRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ServiceName     *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ConfigVersion   *string `json:"ConfigVersion,omitempty" xml:"ConfigVersion,omitempty"`
	HostGroupId     *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	HostInstanceId  *string `json:"HostInstanceId,omitempty" xml:"HostInstanceId,omitempty"`
	ConfigFileName  *string `json:"ConfigFileName,omitempty" xml:"ConfigFileName,omitempty"`
	ConfigItemKey   *string `json:"ConfigItemKey,omitempty" xml:"ConfigItemKey,omitempty"`
	Author          *string `json:"Author,omitempty" xml:"Author,omitempty"`
	Comment         *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListClusterServiceConfigHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClusterServiceConfigHistoryRequest) GoString() string {
	return s.String()
}

func (s *ListClusterServiceConfigHistoryRequest) SetResourceOwnerId(v int64) *ListClusterServiceConfigHistoryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListClusterServiceConfigHistoryRequest) SetRegionId(v string) *ListClusterServiceConfigHistoryRequest {
	s.RegionId = &v
	return s
}

func (s *ListClusterServiceConfigHistoryRequest) SetClusterId(v string) *ListClusterServiceConfigHistoryRequest {
	s.ClusterId = &v
	return s
}

func (s *ListClusterServiceConfigHistoryRequest) SetServiceName(v string) *ListClusterServiceConfigHistoryRequest {
	s.ServiceName = &v
	return s
}

func (s *ListClusterServiceConfigHistoryRequest) SetConfigVersion(v string) *ListClusterServiceConfigHistoryRequest {
	s.ConfigVersion = &v
	return s
}

func (s *ListClusterServiceConfigHistoryRequest) SetHostGroupId(v string) *ListClusterServiceConfigHistoryRequest {
	s.HostGroupId = &v
	return s
}

func (s *ListClusterServiceConfigHistoryRequest) SetHostInstanceId(v string) *ListClusterServiceConfigHistoryRequest {
	s.HostInstanceId = &v
	return s
}

func (s *ListClusterServiceConfigHistoryRequest) SetConfigFileName(v string) *ListClusterServiceConfigHistoryRequest {
	s.ConfigFileName = &v
	return s
}

func (s *ListClusterServiceConfigHistoryRequest) SetConfigItemKey(v string) *ListClusterServiceConfigHistoryRequest {
	s.ConfigItemKey = &v
	return s
}

func (s *ListClusterServiceConfigHistoryRequest) SetAuthor(v string) *ListClusterServiceConfigHistoryRequest {
	s.Author = &v
	return s
}

func (s *ListClusterServiceConfigHistoryRequest) SetComment(v string) *ListClusterServiceConfigHistoryRequest {
	s.Comment = &v
	return s
}

func (s *ListClusterServiceConfigHistoryRequest) SetPageNumber(v int32) *ListClusterServiceConfigHistoryRequest {
	s.PageNumber = &v
	return s
}

func (s *ListClusterServiceConfigHistoryRequest) SetPageSize(v int32) *ListClusterServiceConfigHistoryRequest {
	s.PageSize = &v
	return s
}

type ListClusterServiceConfigHistoryResponseBody struct {
	PageSize          *int32                                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId         *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber        *int32                                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount        *int32                                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	ConfigHistoryList *ListClusterServiceConfigHistoryResponseBodyConfigHistoryList `json:"ConfigHistoryList,omitempty" xml:"ConfigHistoryList,omitempty" type:"Struct"`
}

func (s ListClusterServiceConfigHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClusterServiceConfigHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterServiceConfigHistoryResponseBody) SetPageSize(v int32) *ListClusterServiceConfigHistoryResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListClusterServiceConfigHistoryResponseBody) SetRequestId(v string) *ListClusterServiceConfigHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterServiceConfigHistoryResponseBody) SetPageNumber(v int32) *ListClusterServiceConfigHistoryResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListClusterServiceConfigHistoryResponseBody) SetTotalCount(v int32) *ListClusterServiceConfigHistoryResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListClusterServiceConfigHistoryResponseBody) SetConfigHistoryList(v *ListClusterServiceConfigHistoryResponseBodyConfigHistoryList) *ListClusterServiceConfigHistoryResponseBody {
	s.ConfigHistoryList = v
	return s
}

type ListClusterServiceConfigHistoryResponseBodyConfigHistoryList struct {
	ConfigHistory []*ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory `json:"ConfigHistory,omitempty" xml:"ConfigHistory,omitempty" type:"Repeated"`
}

func (s ListClusterServiceConfigHistoryResponseBodyConfigHistoryList) String() string {
	return tea.Prettify(s)
}

func (s ListClusterServiceConfigHistoryResponseBodyConfigHistoryList) GoString() string {
	return s.String()
}

func (s *ListClusterServiceConfigHistoryResponseBodyConfigHistoryList) SetConfigHistory(v []*ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory) *ListClusterServiceConfigHistoryResponseBodyConfigHistoryList {
	s.ConfigHistory = v
	return s
}

type ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory struct {
	OldValue       *string `json:"OldValue,omitempty" xml:"OldValue,omitempty"`
	Comment        *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	CreateTime     *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	HostName       *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	Author         *string `json:"Author,omitempty" xml:"Author,omitempty"`
	ConfigItemName *string `json:"ConfigItemName,omitempty" xml:"ConfigItemName,omitempty"`
	HostGroupId    *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	NewValue       *string `json:"NewValue,omitempty" xml:"NewValue,omitempty"`
	HostInstanceId *string `json:"HostInstanceId,omitempty" xml:"HostInstanceId,omitempty"`
	ConfigFileName *string `json:"ConfigFileName,omitempty" xml:"ConfigFileName,omitempty"`
	Applied        *bool   `json:"Applied,omitempty" xml:"Applied,omitempty"`
	ConfigVersion  *string `json:"ConfigVersion,omitempty" xml:"ConfigVersion,omitempty"`
	ServiceName    *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	HostGroupName  *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
}

func (s ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory) String() string {
	return tea.Prettify(s)
}

func (s ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory) GoString() string {
	return s.String()
}

func (s *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory) SetOldValue(v string) *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory {
	s.OldValue = &v
	return s
}

func (s *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory) SetComment(v string) *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory {
	s.Comment = &v
	return s
}

func (s *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory) SetCreateTime(v int64) *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory {
	s.CreateTime = &v
	return s
}

func (s *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory) SetHostName(v string) *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory {
	s.HostName = &v
	return s
}

func (s *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory) SetAuthor(v string) *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory {
	s.Author = &v
	return s
}

func (s *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory) SetConfigItemName(v string) *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory {
	s.ConfigItemName = &v
	return s
}

func (s *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory) SetHostGroupId(v string) *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory {
	s.HostGroupId = &v
	return s
}

func (s *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory) SetNewValue(v string) *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory {
	s.NewValue = &v
	return s
}

func (s *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory) SetHostInstanceId(v string) *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory {
	s.HostInstanceId = &v
	return s
}

func (s *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory) SetConfigFileName(v string) *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory {
	s.ConfigFileName = &v
	return s
}

func (s *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory) SetApplied(v bool) *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory {
	s.Applied = &v
	return s
}

func (s *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory) SetConfigVersion(v string) *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory {
	s.ConfigVersion = &v
	return s
}

func (s *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory) SetServiceName(v string) *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory {
	s.ServiceName = &v
	return s
}

func (s *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory) SetHostGroupName(v string) *ListClusterServiceConfigHistoryResponseBodyConfigHistoryListConfigHistory {
	s.HostGroupName = &v
	return s
}

type ListClusterServiceConfigHistoryResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListClusterServiceConfigHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClusterServiceConfigHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClusterServiceConfigHistoryResponse) GoString() string {
	return s.String()
}

func (s *ListClusterServiceConfigHistoryResponse) SetHeaders(v map[string]*string) *ListClusterServiceConfigHistoryResponse {
	s.Headers = v
	return s
}

func (s *ListClusterServiceConfigHistoryResponse) SetBody(v *ListClusterServiceConfigHistoryResponseBody) *ListClusterServiceConfigHistoryResponse {
	s.Body = v
	return s
}

type ModifyScalingConfigItemRequest struct {
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId       *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	ConfigItemInformation *string `json:"ConfigItemInformation,omitempty" xml:"ConfigItemInformation,omitempty"`
	ConfigItemType        *string `json:"ConfigItemType,omitempty" xml:"ConfigItemType,omitempty"`
	ConfigItemBizId       *string `json:"ConfigItemBizId,omitempty" xml:"ConfigItemBizId,omitempty"`
	ScalingGroupBizId     *string `json:"ScalingGroupBizId,omitempty" xml:"ScalingGroupBizId,omitempty"`
}

func (s ModifyScalingConfigItemRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingConfigItemRequest) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigItemRequest) SetResourceOwnerId(v int64) *ModifyScalingConfigItemRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyScalingConfigItemRequest) SetRegionId(v string) *ModifyScalingConfigItemRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyScalingConfigItemRequest) SetResourceGroupId(v string) *ModifyScalingConfigItemRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyScalingConfigItemRequest) SetConfigItemInformation(v string) *ModifyScalingConfigItemRequest {
	s.ConfigItemInformation = &v
	return s
}

func (s *ModifyScalingConfigItemRequest) SetConfigItemType(v string) *ModifyScalingConfigItemRequest {
	s.ConfigItemType = &v
	return s
}

func (s *ModifyScalingConfigItemRequest) SetConfigItemBizId(v string) *ModifyScalingConfigItemRequest {
	s.ConfigItemBizId = &v
	return s
}

func (s *ModifyScalingConfigItemRequest) SetScalingGroupBizId(v string) *ModifyScalingConfigItemRequest {
	s.ScalingGroupBizId = &v
	return s
}

type ModifyScalingConfigItemResponseBody struct {
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// data
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s ModifyScalingConfigItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingConfigItemResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigItemResponseBody) SetRequestId(v string) *ModifyScalingConfigItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyScalingConfigItemResponseBody) SetData(v bool) *ModifyScalingConfigItemResponseBody {
	s.Data = &v
	return s
}

type ModifyScalingConfigItemResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyScalingConfigItemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyScalingConfigItemResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingConfigItemResponse) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigItemResponse) SetHeaders(v map[string]*string) *ModifyScalingConfigItemResponse {
	s.Headers = v
	return s
}

func (s *ModifyScalingConfigItemResponse) SetBody(v *ModifyScalingConfigItemResponseBody) *ModifyScalingConfigItemResponse {
	s.Body = v
	return s
}

type ListFlowClusterAllRequest struct {
	ProductType     *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListFlowClusterAllRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterAllRequest) GoString() string {
	return s.String()
}

func (s *ListFlowClusterAllRequest) SetProductType(v string) *ListFlowClusterAllRequest {
	s.ProductType = &v
	return s
}

func (s *ListFlowClusterAllRequest) SetRegionId(v string) *ListFlowClusterAllRequest {
	s.RegionId = &v
	return s
}

func (s *ListFlowClusterAllRequest) SetResourceGroupId(v string) *ListFlowClusterAllRequest {
	s.ResourceGroupId = &v
	return s
}

type ListFlowClusterAllResponseBody struct {
	PageSize   *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Clusters   *ListFlowClusterAllResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Struct"`
}

func (s ListFlowClusterAllResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterAllResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowClusterAllResponseBody) SetPageSize(v int32) *ListFlowClusterAllResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListFlowClusterAllResponseBody) SetRequestId(v string) *ListFlowClusterAllResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowClusterAllResponseBody) SetPageNumber(v int32) *ListFlowClusterAllResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListFlowClusterAllResponseBody) SetTotalCount(v int32) *ListFlowClusterAllResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListFlowClusterAllResponseBody) SetClusters(v *ListFlowClusterAllResponseBodyClusters) *ListFlowClusterAllResponseBody {
	s.Clusters = v
	return s
}

type ListFlowClusterAllResponseBodyClusters struct {
	ClusterInfo []*ListFlowClusterAllResponseBodyClustersClusterInfo `json:"ClusterInfo,omitempty" xml:"ClusterInfo,omitempty" type:"Repeated"`
}

func (s ListFlowClusterAllResponseBodyClusters) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterAllResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *ListFlowClusterAllResponseBodyClusters) SetClusterInfo(v []*ListFlowClusterAllResponseBodyClustersClusterInfo) *ListFlowClusterAllResponseBodyClusters {
	s.ClusterInfo = v
	return s
}

type ListFlowClusterAllResponseBodyClustersClusterInfo struct {
	Type                *string                                                         `json:"Type,omitempty" xml:"Type,omitempty"`
	Status              *string                                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	RunningTime         *int32                                                          `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	OrderList           *string                                                         `json:"OrderList,omitempty" xml:"OrderList,omitempty"`
	CreateTime          *int64                                                          `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ChargeType          *string                                                         `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Period              *int32                                                          `json:"Period,omitempty" xml:"Period,omitempty"`
	K8sClusterId        *string                                                         `json:"K8sClusterId,omitempty" xml:"K8sClusterId,omitempty"`
	ExpiredTime         *int64                                                          `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	HasUncompletedOrder *bool                                                           `json:"HasUncompletedOrder,omitempty" xml:"HasUncompletedOrder,omitempty"`
	Name                *string                                                         `json:"Name,omitempty" xml:"Name,omitempty"`
	Id                  *string                                                         `json:"Id,omitempty" xml:"Id,omitempty"`
	CreateResource      *string                                                         `json:"CreateResource,omitempty" xml:"CreateResource,omitempty"`
	OrderTaskInfo       *ListFlowClusterAllResponseBodyClustersClusterInfoOrderTaskInfo `json:"OrderTaskInfo,omitempty" xml:"OrderTaskInfo,omitempty" type:"Struct"`
	FailReason          *ListFlowClusterAllResponseBodyClustersClusterInfoFailReason    `json:"FailReason,omitempty" xml:"FailReason,omitempty" type:"Struct"`
}

func (s ListFlowClusterAllResponseBodyClustersClusterInfo) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterAllResponseBodyClustersClusterInfo) GoString() string {
	return s.String()
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfo) SetType(v string) *ListFlowClusterAllResponseBodyClustersClusterInfo {
	s.Type = &v
	return s
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfo) SetStatus(v string) *ListFlowClusterAllResponseBodyClustersClusterInfo {
	s.Status = &v
	return s
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfo) SetRunningTime(v int32) *ListFlowClusterAllResponseBodyClustersClusterInfo {
	s.RunningTime = &v
	return s
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfo) SetOrderList(v string) *ListFlowClusterAllResponseBodyClustersClusterInfo {
	s.OrderList = &v
	return s
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfo) SetCreateTime(v int64) *ListFlowClusterAllResponseBodyClustersClusterInfo {
	s.CreateTime = &v
	return s
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfo) SetChargeType(v string) *ListFlowClusterAllResponseBodyClustersClusterInfo {
	s.ChargeType = &v
	return s
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfo) SetPeriod(v int32) *ListFlowClusterAllResponseBodyClustersClusterInfo {
	s.Period = &v
	return s
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfo) SetK8sClusterId(v string) *ListFlowClusterAllResponseBodyClustersClusterInfo {
	s.K8sClusterId = &v
	return s
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfo) SetExpiredTime(v int64) *ListFlowClusterAllResponseBodyClustersClusterInfo {
	s.ExpiredTime = &v
	return s
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfo) SetHasUncompletedOrder(v bool) *ListFlowClusterAllResponseBodyClustersClusterInfo {
	s.HasUncompletedOrder = &v
	return s
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfo) SetName(v string) *ListFlowClusterAllResponseBodyClustersClusterInfo {
	s.Name = &v
	return s
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfo) SetId(v string) *ListFlowClusterAllResponseBodyClustersClusterInfo {
	s.Id = &v
	return s
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfo) SetCreateResource(v string) *ListFlowClusterAllResponseBodyClustersClusterInfo {
	s.CreateResource = &v
	return s
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfo) SetOrderTaskInfo(v *ListFlowClusterAllResponseBodyClustersClusterInfoOrderTaskInfo) *ListFlowClusterAllResponseBodyClustersClusterInfo {
	s.OrderTaskInfo = v
	return s
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfo) SetFailReason(v *ListFlowClusterAllResponseBodyClustersClusterInfoFailReason) *ListFlowClusterAllResponseBodyClustersClusterInfo {
	s.FailReason = v
	return s
}

type ListFlowClusterAllResponseBodyClustersClusterInfoOrderTaskInfo struct {
	TargetCount  *int32  `json:"TargetCount,omitempty" xml:"TargetCount,omitempty"`
	CurrentCount *int32  `json:"CurrentCount,omitempty" xml:"CurrentCount,omitempty"`
	OrderIdList  *string `json:"OrderIdList,omitempty" xml:"OrderIdList,omitempty"`
}

func (s ListFlowClusterAllResponseBodyClustersClusterInfoOrderTaskInfo) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterAllResponseBodyClustersClusterInfoOrderTaskInfo) GoString() string {
	return s.String()
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfoOrderTaskInfo) SetTargetCount(v int32) *ListFlowClusterAllResponseBodyClustersClusterInfoOrderTaskInfo {
	s.TargetCount = &v
	return s
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfoOrderTaskInfo) SetCurrentCount(v int32) *ListFlowClusterAllResponseBodyClustersClusterInfoOrderTaskInfo {
	s.CurrentCount = &v
	return s
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfoOrderTaskInfo) SetOrderIdList(v string) *ListFlowClusterAllResponseBodyClustersClusterInfoOrderTaskInfo {
	s.OrderIdList = &v
	return s
}

type ListFlowClusterAllResponseBodyClustersClusterInfoFailReason struct {
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
}

func (s ListFlowClusterAllResponseBodyClustersClusterInfoFailReason) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterAllResponseBodyClustersClusterInfoFailReason) GoString() string {
	return s.String()
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfoFailReason) SetErrorMsg(v string) *ListFlowClusterAllResponseBodyClustersClusterInfoFailReason {
	s.ErrorMsg = &v
	return s
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfoFailReason) SetRequestId(v string) *ListFlowClusterAllResponseBodyClustersClusterInfoFailReason {
	s.RequestId = &v
	return s
}

func (s *ListFlowClusterAllResponseBodyClustersClusterInfoFailReason) SetErrorCode(v string) *ListFlowClusterAllResponseBodyClustersClusterInfoFailReason {
	s.ErrorCode = &v
	return s
}

type ListFlowClusterAllResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFlowClusterAllResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowClusterAllResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterAllResponse) GoString() string {
	return s.String()
}

func (s *ListFlowClusterAllResponse) SetHeaders(v map[string]*string) *ListFlowClusterAllResponse {
	s.Headers = v
	return s
}

func (s *ListFlowClusterAllResponse) SetBody(v *ListFlowClusterAllResponseBody) *ListFlowClusterAllResponse {
	s.Body = v
	return s
}

type DescribeScalingGroupRequest struct {
	ResourceOwnerId   *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ScalingGroupBizId *string `json:"ScalingGroupBizId,omitempty" xml:"ScalingGroupBizId,omitempty"`
	HostGroupBizId    *string `json:"HostGroupBizId,omitempty" xml:"HostGroupBizId,omitempty"`
}

func (s DescribeScalingGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupRequest) SetResourceOwnerId(v int64) *DescribeScalingGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeScalingGroupRequest) SetRegionId(v string) *DescribeScalingGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeScalingGroupRequest) SetScalingGroupBizId(v string) *DescribeScalingGroupRequest {
	s.ScalingGroupBizId = &v
	return s
}

func (s *DescribeScalingGroupRequest) SetHostGroupBizId(v string) *DescribeScalingGroupRequest {
	s.HostGroupBizId = &v
	return s
}

type DescribeScalingGroupResponseBody struct {
	// Id of the request
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ActiveStatus   *string `json:"ActiveStatus,omitempty" xml:"ActiveStatus,omitempty"`
	HostGroupBizId *string `json:"HostGroupBizId,omitempty" xml:"HostGroupBizId,omitempty"`
	ScalingInMode  *string `json:"ScalingInMode,omitempty" xml:"ScalingInMode,omitempty"`
	ScalingMinSize *int64  `json:"ScalingMinSize,omitempty" xml:"ScalingMinSize,omitempty"`
	ScalingMaxSize *int64  `json:"ScalingMaxSize,omitempty" xml:"ScalingMaxSize,omitempty"`
	ConfigState    *string `json:"ConfigState,omitempty" xml:"ConfigState,omitempty"`
}

func (s DescribeScalingGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupResponseBody) SetRequestId(v string) *DescribeScalingGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScalingGroupResponseBody) SetScalingGroupId(v string) *DescribeScalingGroupResponseBody {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScalingGroupResponseBody) SetName(v string) *DescribeScalingGroupResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeScalingGroupResponseBody) SetDescription(v string) *DescribeScalingGroupResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeScalingGroupResponseBody) SetActiveStatus(v string) *DescribeScalingGroupResponseBody {
	s.ActiveStatus = &v
	return s
}

func (s *DescribeScalingGroupResponseBody) SetHostGroupBizId(v string) *DescribeScalingGroupResponseBody {
	s.HostGroupBizId = &v
	return s
}

func (s *DescribeScalingGroupResponseBody) SetScalingInMode(v string) *DescribeScalingGroupResponseBody {
	s.ScalingInMode = &v
	return s
}

func (s *DescribeScalingGroupResponseBody) SetScalingMinSize(v int64) *DescribeScalingGroupResponseBody {
	s.ScalingMinSize = &v
	return s
}

func (s *DescribeScalingGroupResponseBody) SetScalingMaxSize(v int64) *DescribeScalingGroupResponseBody {
	s.ScalingMaxSize = &v
	return s
}

func (s *DescribeScalingGroupResponseBody) SetConfigState(v string) *DescribeScalingGroupResponseBody {
	s.ConfigState = &v
	return s
}

type DescribeScalingGroupResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeScalingGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScalingGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupResponse) SetHeaders(v map[string]*string) *DescribeScalingGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeScalingGroupResponse) SetBody(v *DescribeScalingGroupResponseBody) *DescribeScalingGroupResponse {
	s.Body = v
	return s
}

type ListScalingGroupRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Limit           *int64  `json:"limit,omitempty" xml:"limit,omitempty"`
	PageNumber      *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentSize     *int64  `json:"CurrentSize,omitempty" xml:"CurrentSize,omitempty"`
	PageCount       *int64  `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	OrderField      *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	OrderMode       *string `json:"OrderMode,omitempty" xml:"OrderMode,omitempty"`
	ClusterBizId    *string `json:"ClusterBizId,omitempty" xml:"ClusterBizId,omitempty"`
}

func (s ListScalingGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListScalingGroupRequest) GoString() string {
	return s.String()
}

func (s *ListScalingGroupRequest) SetResourceOwnerId(v int64) *ListScalingGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListScalingGroupRequest) SetRegionId(v string) *ListScalingGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ListScalingGroupRequest) SetResourceGroupId(v string) *ListScalingGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListScalingGroupRequest) SetLimit(v int64) *ListScalingGroupRequest {
	s.Limit = &v
	return s
}

func (s *ListScalingGroupRequest) SetPageNumber(v int64) *ListScalingGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *ListScalingGroupRequest) SetPageSize(v int64) *ListScalingGroupRequest {
	s.PageSize = &v
	return s
}

func (s *ListScalingGroupRequest) SetCurrentSize(v int64) *ListScalingGroupRequest {
	s.CurrentSize = &v
	return s
}

func (s *ListScalingGroupRequest) SetPageCount(v int64) *ListScalingGroupRequest {
	s.PageCount = &v
	return s
}

func (s *ListScalingGroupRequest) SetOrderField(v string) *ListScalingGroupRequest {
	s.OrderField = &v
	return s
}

func (s *ListScalingGroupRequest) SetOrderMode(v string) *ListScalingGroupRequest {
	s.OrderMode = &v
	return s
}

func (s *ListScalingGroupRequest) SetClusterBizId(v string) *ListScalingGroupRequest {
	s.ClusterBizId = &v
	return s
}

type ListScalingGroupResponseBody struct {
	// Id of the request
	RequestId  *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int64                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int64                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	NextToken  *string                            `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Items      *ListScalingGroupResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s ListScalingGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListScalingGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListScalingGroupResponseBody) SetRequestId(v string) *ListScalingGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScalingGroupResponseBody) SetPageNumber(v int64) *ListScalingGroupResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListScalingGroupResponseBody) SetPageSize(v int64) *ListScalingGroupResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListScalingGroupResponseBody) SetTotalCount(v int64) *ListScalingGroupResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListScalingGroupResponseBody) SetNextToken(v string) *ListScalingGroupResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListScalingGroupResponseBody) SetItems(v *ListScalingGroupResponseBodyItems) *ListScalingGroupResponseBody {
	s.Items = v
	return s
}

type ListScalingGroupResponseBodyItems struct {
	Item []*ListScalingGroupResponseBodyItemsItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s ListScalingGroupResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s ListScalingGroupResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListScalingGroupResponseBodyItems) SetItem(v []*ListScalingGroupResponseBodyItemsItem) *ListScalingGroupResponseBodyItems {
	s.Item = v
	return s
}

type ListScalingGroupResponseBodyItemsItem struct {
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ActiveStatus   *string `json:"ActiveStatus,omitempty" xml:"ActiveStatus,omitempty"`
	HostGroupBizId *string `json:"HostGroupBizId,omitempty" xml:"HostGroupBizId,omitempty"`
	ScalingInMode  *string `json:"ScalingInMode,omitempty" xml:"ScalingInMode,omitempty"`
	ScalingMinSize *string `json:"ScalingMinSize,omitempty" xml:"ScalingMinSize,omitempty"`
	ScalingMaxSize *string `json:"ScalingMaxSize,omitempty" xml:"ScalingMaxSize,omitempty"`
}

func (s ListScalingGroupResponseBodyItemsItem) String() string {
	return tea.Prettify(s)
}

func (s ListScalingGroupResponseBodyItemsItem) GoString() string {
	return s.String()
}

func (s *ListScalingGroupResponseBodyItemsItem) SetScalingGroupId(v string) *ListScalingGroupResponseBodyItemsItem {
	s.ScalingGroupId = &v
	return s
}

func (s *ListScalingGroupResponseBodyItemsItem) SetName(v string) *ListScalingGroupResponseBodyItemsItem {
	s.Name = &v
	return s
}

func (s *ListScalingGroupResponseBodyItemsItem) SetDescription(v string) *ListScalingGroupResponseBodyItemsItem {
	s.Description = &v
	return s
}

func (s *ListScalingGroupResponseBodyItemsItem) SetActiveStatus(v string) *ListScalingGroupResponseBodyItemsItem {
	s.ActiveStatus = &v
	return s
}

func (s *ListScalingGroupResponseBodyItemsItem) SetHostGroupBizId(v string) *ListScalingGroupResponseBodyItemsItem {
	s.HostGroupBizId = &v
	return s
}

func (s *ListScalingGroupResponseBodyItemsItem) SetScalingInMode(v string) *ListScalingGroupResponseBodyItemsItem {
	s.ScalingInMode = &v
	return s
}

func (s *ListScalingGroupResponseBodyItemsItem) SetScalingMinSize(v string) *ListScalingGroupResponseBodyItemsItem {
	s.ScalingMinSize = &v
	return s
}

func (s *ListScalingGroupResponseBodyItemsItem) SetScalingMaxSize(v string) *ListScalingGroupResponseBodyItemsItem {
	s.ScalingMaxSize = &v
	return s
}

type ListScalingGroupResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListScalingGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListScalingGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListScalingGroupResponse) GoString() string {
	return s.String()
}

func (s *ListScalingGroupResponse) SetHeaders(v map[string]*string) *ListScalingGroupResponse {
	s.Headers = v
	return s
}

func (s *ListScalingGroupResponse) SetBody(v *ListScalingGroupResponseBody) *ListScalingGroupResponse {
	s.Body = v
	return s
}

type ModifyFlowCategoryRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentId  *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s ModifyFlowCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowCategoryRequest) GoString() string {
	return s.String()
}

func (s *ModifyFlowCategoryRequest) SetProjectId(v string) *ModifyFlowCategoryRequest {
	s.ProjectId = &v
	return s
}

func (s *ModifyFlowCategoryRequest) SetRegionId(v string) *ModifyFlowCategoryRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyFlowCategoryRequest) SetId(v string) *ModifyFlowCategoryRequest {
	s.Id = &v
	return s
}

func (s *ModifyFlowCategoryRequest) SetName(v string) *ModifyFlowCategoryRequest {
	s.Name = &v
	return s
}

func (s *ModifyFlowCategoryRequest) SetParentId(v string) *ModifyFlowCategoryRequest {
	s.ParentId = &v
	return s
}

type ModifyFlowCategoryResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyFlowCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFlowCategoryResponseBody) SetData(v bool) *ModifyFlowCategoryResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyFlowCategoryResponseBody) SetRequestId(v string) *ModifyFlowCategoryResponseBody {
	s.RequestId = &v
	return s
}

type ModifyFlowCategoryResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyFlowCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyFlowCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowCategoryResponse) GoString() string {
	return s.String()
}

func (s *ModifyFlowCategoryResponse) SetHeaders(v map[string]*string) *ModifyFlowCategoryResponse {
	s.Headers = v
	return s
}

func (s *ModifyFlowCategoryResponse) SetBody(v *ModifyFlowCategoryResponseBody) *ModifyFlowCategoryResponse {
	s.Body = v
	return s
}

type ModifyClusterServiceConfigRequest struct {
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClusterId            *string   `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ServiceName          *string   `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	Comment              *string   `json:"Comment,omitempty" xml:"Comment,omitempty"`
	ConfigParams         *string   `json:"ConfigParams,omitempty" xml:"ConfigParams,omitempty"`
	CustomConfigParams   *string   `json:"CustomConfigParams,omitempty" xml:"CustomConfigParams,omitempty"`
	GroupId              *string   `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	HostInstanceId       *string   `json:"HostInstanceId,omitempty" xml:"HostInstanceId,omitempty"`
	ConfigType           *string   `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	RefreshHostConfig    *bool     `json:"RefreshHostConfig,omitempty" xml:"RefreshHostConfig,omitempty"`
	GatewayClusterIdList []*string `json:"GatewayClusterIdList,omitempty" xml:"GatewayClusterIdList,omitempty" type:"Repeated"`
}

func (s ModifyClusterServiceConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterServiceConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterServiceConfigRequest) SetResourceOwnerId(v int64) *ModifyClusterServiceConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyClusterServiceConfigRequest) SetRegionId(v string) *ModifyClusterServiceConfigRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyClusterServiceConfigRequest) SetClusterId(v string) *ModifyClusterServiceConfigRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyClusterServiceConfigRequest) SetServiceName(v string) *ModifyClusterServiceConfigRequest {
	s.ServiceName = &v
	return s
}

func (s *ModifyClusterServiceConfigRequest) SetComment(v string) *ModifyClusterServiceConfigRequest {
	s.Comment = &v
	return s
}

func (s *ModifyClusterServiceConfigRequest) SetConfigParams(v string) *ModifyClusterServiceConfigRequest {
	s.ConfigParams = &v
	return s
}

func (s *ModifyClusterServiceConfigRequest) SetCustomConfigParams(v string) *ModifyClusterServiceConfigRequest {
	s.CustomConfigParams = &v
	return s
}

func (s *ModifyClusterServiceConfigRequest) SetGroupId(v string) *ModifyClusterServiceConfigRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyClusterServiceConfigRequest) SetHostInstanceId(v string) *ModifyClusterServiceConfigRequest {
	s.HostInstanceId = &v
	return s
}

func (s *ModifyClusterServiceConfigRequest) SetConfigType(v string) *ModifyClusterServiceConfigRequest {
	s.ConfigType = &v
	return s
}

func (s *ModifyClusterServiceConfigRequest) SetRefreshHostConfig(v bool) *ModifyClusterServiceConfigRequest {
	s.RefreshHostConfig = &v
	return s
}

func (s *ModifyClusterServiceConfigRequest) SetGatewayClusterIdList(v []*string) *ModifyClusterServiceConfigRequest {
	s.GatewayClusterIdList = v
	return s
}

type ModifyClusterServiceConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyClusterServiceConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterServiceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterServiceConfigResponseBody) SetRequestId(v string) *ModifyClusterServiceConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyClusterServiceConfigResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyClusterServiceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyClusterServiceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterServiceConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterServiceConfigResponse) SetHeaders(v map[string]*string) *ModifyClusterServiceConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterServiceConfigResponse) SetBody(v *ModifyClusterServiceConfigResponseBody) *ModifyClusterServiceConfigResponse {
	s.Body = v
	return s
}

type CloneFlowRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CloneFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s CloneFlowRequest) GoString() string {
	return s.String()
}

func (s *CloneFlowRequest) SetProjectId(v string) *CloneFlowRequest {
	s.ProjectId = &v
	return s
}

func (s *CloneFlowRequest) SetId(v string) *CloneFlowRequest {
	s.Id = &v
	return s
}

func (s *CloneFlowRequest) SetRegionId(v string) *CloneFlowRequest {
	s.RegionId = &v
	return s
}

type CloneFlowResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloneFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloneFlowResponseBody) GoString() string {
	return s.String()
}

func (s *CloneFlowResponseBody) SetId(v string) *CloneFlowResponseBody {
	s.Id = &v
	return s
}

func (s *CloneFlowResponseBody) SetRequestId(v string) *CloneFlowResponseBody {
	s.RequestId = &v
	return s
}

type CloneFlowResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CloneFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CloneFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s CloneFlowResponse) GoString() string {
	return s.String()
}

func (s *CloneFlowResponse) SetHeaders(v map[string]*string) *CloneFlowResponse {
	s.Headers = v
	return s
}

func (s *CloneFlowResponse) SetBody(v *CloneFlowResponseBody) *CloneFlowResponse {
	s.Body = v
	return s
}

type CreateClusterTemplateRequest struct {
	ResourceOwnerId        *int64                                         `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TemplateName           *string                                        `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	RegionId               *string                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId                 *string                                        `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	LogPath                *string                                        `json:"LogPath,omitempty" xml:"LogPath,omitempty"`
	SecurityGroupId        *string                                        `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	IsOpenPublicIp         *bool                                          `json:"IsOpenPublicIp,omitempty" xml:"IsOpenPublicIp,omitempty"`
	SecurityGroupName      *string                                        `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
	Period                 *int32                                         `json:"Period,omitempty" xml:"Period,omitempty"`
	RenewAuto              *bool                                          `json:"RenewAuto,omitempty" xml:"RenewAuto,omitempty"`
	VpcId                  *string                                        `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VSwitchId              *string                                        `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	NetType                *string                                        `json:"NetType,omitempty" xml:"NetType,omitempty"`
	UserDefinedEmrEcsRole  *string                                        `json:"UserDefinedEmrEcsRole,omitempty" xml:"UserDefinedEmrEcsRole,omitempty"`
	EmrVer                 *string                                        `json:"EmrVer,omitempty" xml:"EmrVer,omitempty"`
	ClusterType            *string                                        `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	EnableHighAvailability *bool                                          `json:"EnableHighAvailability,omitempty" xml:"EnableHighAvailability,omitempty"`
	UseLocalMetaDb         *bool                                          `json:"UseLocalMetaDb,omitempty" xml:"UseLocalMetaDb,omitempty"`
	IoOptimizedOption      *bool                                          `json:"IoOptimizedOption,omitempty" xml:"IoOptimizedOption,omitempty"`
	EnableSsh              *bool                                          `json:"EnableSsh,omitempty" xml:"EnableSsh,omitempty"`
	InstanceGeneration     *string                                        `json:"InstanceGeneration,omitempty" xml:"InstanceGeneration,omitempty"`
	MasterPwd              *string                                        `json:"MasterPwd,omitempty" xml:"MasterPwd,omitempty"`
	KeyPairName            *string                                        `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	MetaStoreType          *string                                        `json:"MetaStoreType,omitempty" xml:"MetaStoreType,omitempty"`
	MetaStoreConf          *string                                        `json:"MetaStoreConf,omitempty" xml:"MetaStoreConf,omitempty"`
	Configurations         *string                                        `json:"Configurations,omitempty" xml:"Configurations,omitempty"`
	EnableEas              *bool                                          `json:"EnableEas,omitempty" xml:"EnableEas,omitempty"`
	DepositType            *string                                        `json:"DepositType,omitempty" xml:"DepositType,omitempty"`
	MachineType            *string                                        `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	UseCustomHiveMetaDb    *bool                                          `json:"UseCustomHiveMetaDb,omitempty" xml:"UseCustomHiveMetaDb,omitempty"`
	InitCustomHiveMetaDb   *bool                                          `json:"InitCustomHiveMetaDb,omitempty" xml:"InitCustomHiveMetaDb,omitempty"`
	ResourceGroupId        *string                                        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ClientToken            *string                                        `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OptionSoftWareList     []*string                                      `json:"OptionSoftWareList,omitempty" xml:"OptionSoftWareList,omitempty" type:"Repeated"`
	HostGroup              []*CreateClusterTemplateRequestHostGroup       `json:"HostGroup,omitempty" xml:"HostGroup,omitempty" type:"Repeated"`
	BootstrapAction        []*CreateClusterTemplateRequestBootstrapAction `json:"BootstrapAction,omitempty" xml:"BootstrapAction,omitempty" type:"Repeated"`
	Config                 []*CreateClusterTemplateRequestConfig          `json:"Config,omitempty" xml:"Config,omitempty" type:"Repeated"`
	Tag                    []*CreateClusterTemplateRequestTag             `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateClusterTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterTemplateRequest) SetResourceOwnerId(v int64) *CreateClusterTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateClusterTemplateRequest) SetTemplateName(v string) *CreateClusterTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateClusterTemplateRequest) SetRegionId(v string) *CreateClusterTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *CreateClusterTemplateRequest) SetZoneId(v string) *CreateClusterTemplateRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateClusterTemplateRequest) SetLogPath(v string) *CreateClusterTemplateRequest {
	s.LogPath = &v
	return s
}

func (s *CreateClusterTemplateRequest) SetSecurityGroupId(v string) *CreateClusterTemplateRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateClusterTemplateRequest) SetIsOpenPublicIp(v bool) *CreateClusterTemplateRequest {
	s.IsOpenPublicIp = &v
	return s
}

func (s *CreateClusterTemplateRequest) SetSecurityGroupName(v string) *CreateClusterTemplateRequest {
	s.SecurityGroupName = &v
	return s
}

func (s *CreateClusterTemplateRequest) SetPeriod(v int32) *CreateClusterTemplateRequest {
	s.Period = &v
	return s
}

func (s *CreateClusterTemplateRequest) SetRenewAuto(v bool) *CreateClusterTemplateRequest {
	s.RenewAuto = &v
	return s
}

func (s *CreateClusterTemplateRequest) SetVpcId(v string) *CreateClusterTemplateRequest {
	s.VpcId = &v
	return s
}

func (s *CreateClusterTemplateRequest) SetVSwitchId(v string) *CreateClusterTemplateRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateClusterTemplateRequest) SetNetType(v string) *CreateClusterTemplateRequest {
	s.NetType = &v
	return s
}

func (s *CreateClusterTemplateRequest) SetUserDefinedEmrEcsRole(v string) *CreateClusterTemplateRequest {
	s.UserDefinedEmrEcsRole = &v
	return s
}

func (s *CreateClusterTemplateRequest) SetEmrVer(v string) *CreateClusterTemplateRequest {
	s.EmrVer = &v
	return s
}

func (s *CreateClusterTemplateRequest) SetClusterType(v string) *CreateClusterTemplateRequest {
	s.ClusterType = &v
	return s
}

func (s *CreateClusterTemplateRequest) SetEnableHighAvailability(v bool) *CreateClusterTemplateRequest {
	s.EnableHighAvailability = &v
	return s
}

func (s *CreateClusterTemplateRequest) SetUseLocalMetaDb(v bool) *CreateClusterTemplateRequest {
	s.UseLocalMetaDb = &v
	return s
}

func (s *CreateClusterTemplateRequest) SetIoOptimizedOption(v bool) *CreateClusterTemplateRequest {
	s.IoOptimizedOption = &v
	return s
}

func (s *CreateClusterTemplateRequest) SetEnableSsh(v bool) *CreateClusterTemplateRequest {
	s.EnableSsh = &v
	return s
}

func (s *CreateClusterTemplateRequest) SetInstanceGeneration(v string) *CreateClusterTemplateRequest {
	s.InstanceGeneration = &v
	return s
}

func (s *CreateClusterTemplateRequest) SetMasterPwd(v string) *CreateClusterTemplateRequest {
	s.MasterPwd = &v
	return s
}

func (s *CreateClusterTemplateRequest) SetKeyPairName(v string) *CreateClusterTemplateRequest {
	s.KeyPairName = &v
	return s
}

func (s *CreateClusterTemplateRequest) SetMetaStoreType(v string) *CreateClusterTemplateRequest {
	s.MetaStoreType = &v
	return s
}

func (s *CreateClusterTemplateRequest) SetMetaStoreConf(v string) *CreateClusterTemplateRequest {
	s.MetaStoreConf = &v
	return s
}

func (s *CreateClusterTemplateRequest) SetConfigurations(v string) *CreateClusterTemplateRequest {
	s.Configurations = &v
	return s
}

func (s *CreateClusterTemplateRequest) SetEnableEas(v bool) *CreateClusterTemplateRequest {
	s.EnableEas = &v
	return s
}

func (s *CreateClusterTemplateRequest) SetDepositType(v string) *CreateClusterTemplateRequest {
	s.DepositType = &v
	return s
}

func (s *CreateClusterTemplateRequest) SetMachineType(v string) *CreateClusterTemplateRequest {
	s.MachineType = &v
	return s
}

func (s *CreateClusterTemplateRequest) SetUseCustomHiveMetaDb(v bool) *CreateClusterTemplateRequest {
	s.UseCustomHiveMetaDb = &v
	return s
}

func (s *CreateClusterTemplateRequest) SetInitCustomHiveMetaDb(v bool) *CreateClusterTemplateRequest {
	s.InitCustomHiveMetaDb = &v
	return s
}

func (s *CreateClusterTemplateRequest) SetResourceGroupId(v string) *CreateClusterTemplateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateClusterTemplateRequest) SetClientToken(v string) *CreateClusterTemplateRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateClusterTemplateRequest) SetOptionSoftWareList(v []*string) *CreateClusterTemplateRequest {
	s.OptionSoftWareList = v
	return s
}

func (s *CreateClusterTemplateRequest) SetHostGroup(v []*CreateClusterTemplateRequestHostGroup) *CreateClusterTemplateRequest {
	s.HostGroup = v
	return s
}

func (s *CreateClusterTemplateRequest) SetBootstrapAction(v []*CreateClusterTemplateRequestBootstrapAction) *CreateClusterTemplateRequest {
	s.BootstrapAction = v
	return s
}

func (s *CreateClusterTemplateRequest) SetConfig(v []*CreateClusterTemplateRequestConfig) *CreateClusterTemplateRequest {
	s.Config = v
	return s
}

func (s *CreateClusterTemplateRequest) SetTag(v []*CreateClusterTemplateRequestTag) *CreateClusterTemplateRequest {
	s.Tag = v
	return s
}

type CreateClusterTemplateRequestHostGroup struct {
	SysDiskCapacity    *int32  `json:"SysDiskCapacity,omitempty" xml:"SysDiskCapacity,omitempty"`
	HostGroupType      *string `json:"HostGroupType,omitempty" xml:"HostGroupType,omitempty"`
	Comment            *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	MultiInstanceTypes *string `json:"MultiInstanceTypes,omitempty" xml:"MultiInstanceTypes,omitempty"`
	SysDiskType        *string `json:"SysDiskType,omitempty" xml:"SysDiskType,omitempty"`
	AutoRenew          *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	ChargeType         *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	DiskType           *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	HostGroupId        *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	InstanceType       *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	DiskCount          *int32  `json:"DiskCount,omitempty" xml:"DiskCount,omitempty"`
	CreateType         *string `json:"CreateType,omitempty" xml:"CreateType,omitempty"`
	Period             *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	DiskCapacity       *int32  `json:"DiskCapacity,omitempty" xml:"DiskCapacity,omitempty"`
	VSwitchId          *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	NodeCount          *int32  `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	HostGroupName      *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
	ClusterId          *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s CreateClusterTemplateRequestHostGroup) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterTemplateRequestHostGroup) GoString() string {
	return s.String()
}

func (s *CreateClusterTemplateRequestHostGroup) SetSysDiskCapacity(v int32) *CreateClusterTemplateRequestHostGroup {
	s.SysDiskCapacity = &v
	return s
}

func (s *CreateClusterTemplateRequestHostGroup) SetHostGroupType(v string) *CreateClusterTemplateRequestHostGroup {
	s.HostGroupType = &v
	return s
}

func (s *CreateClusterTemplateRequestHostGroup) SetComment(v string) *CreateClusterTemplateRequestHostGroup {
	s.Comment = &v
	return s
}

func (s *CreateClusterTemplateRequestHostGroup) SetMultiInstanceTypes(v string) *CreateClusterTemplateRequestHostGroup {
	s.MultiInstanceTypes = &v
	return s
}

func (s *CreateClusterTemplateRequestHostGroup) SetSysDiskType(v string) *CreateClusterTemplateRequestHostGroup {
	s.SysDiskType = &v
	return s
}

func (s *CreateClusterTemplateRequestHostGroup) SetAutoRenew(v bool) *CreateClusterTemplateRequestHostGroup {
	s.AutoRenew = &v
	return s
}

func (s *CreateClusterTemplateRequestHostGroup) SetChargeType(v string) *CreateClusterTemplateRequestHostGroup {
	s.ChargeType = &v
	return s
}

func (s *CreateClusterTemplateRequestHostGroup) SetDiskType(v string) *CreateClusterTemplateRequestHostGroup {
	s.DiskType = &v
	return s
}

func (s *CreateClusterTemplateRequestHostGroup) SetHostGroupId(v string) *CreateClusterTemplateRequestHostGroup {
	s.HostGroupId = &v
	return s
}

func (s *CreateClusterTemplateRequestHostGroup) SetInstanceType(v string) *CreateClusterTemplateRequestHostGroup {
	s.InstanceType = &v
	return s
}

func (s *CreateClusterTemplateRequestHostGroup) SetDiskCount(v int32) *CreateClusterTemplateRequestHostGroup {
	s.DiskCount = &v
	return s
}

func (s *CreateClusterTemplateRequestHostGroup) SetCreateType(v string) *CreateClusterTemplateRequestHostGroup {
	s.CreateType = &v
	return s
}

func (s *CreateClusterTemplateRequestHostGroup) SetPeriod(v int32) *CreateClusterTemplateRequestHostGroup {
	s.Period = &v
	return s
}

func (s *CreateClusterTemplateRequestHostGroup) SetDiskCapacity(v int32) *CreateClusterTemplateRequestHostGroup {
	s.DiskCapacity = &v
	return s
}

func (s *CreateClusterTemplateRequestHostGroup) SetVSwitchId(v string) *CreateClusterTemplateRequestHostGroup {
	s.VSwitchId = &v
	return s
}

func (s *CreateClusterTemplateRequestHostGroup) SetNodeCount(v int32) *CreateClusterTemplateRequestHostGroup {
	s.NodeCount = &v
	return s
}

func (s *CreateClusterTemplateRequestHostGroup) SetHostGroupName(v string) *CreateClusterTemplateRequestHostGroup {
	s.HostGroupName = &v
	return s
}

func (s *CreateClusterTemplateRequestHostGroup) SetClusterId(v string) *CreateClusterTemplateRequestHostGroup {
	s.ClusterId = &v
	return s
}

type CreateClusterTemplateRequestBootstrapAction struct {
	Arg  *string `json:"Arg,omitempty" xml:"Arg,omitempty"`
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateClusterTemplateRequestBootstrapAction) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterTemplateRequestBootstrapAction) GoString() string {
	return s.String()
}

func (s *CreateClusterTemplateRequestBootstrapAction) SetArg(v string) *CreateClusterTemplateRequestBootstrapAction {
	s.Arg = &v
	return s
}

func (s *CreateClusterTemplateRequestBootstrapAction) SetPath(v string) *CreateClusterTemplateRequestBootstrapAction {
	s.Path = &v
	return s
}

func (s *CreateClusterTemplateRequestBootstrapAction) SetName(v string) *CreateClusterTemplateRequestBootstrapAction {
	s.Name = &v
	return s
}

type CreateClusterTemplateRequestConfig struct {
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	Replace     *string `json:"Replace,omitempty" xml:"Replace,omitempty"`
	FileName    *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ConfigKey   *string `json:"ConfigKey,omitempty" xml:"ConfigKey,omitempty"`
	Encrypt     *string `json:"Encrypt,omitempty" xml:"Encrypt,omitempty"`
}

func (s CreateClusterTemplateRequestConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterTemplateRequestConfig) GoString() string {
	return s.String()
}

func (s *CreateClusterTemplateRequestConfig) SetConfigValue(v string) *CreateClusterTemplateRequestConfig {
	s.ConfigValue = &v
	return s
}

func (s *CreateClusterTemplateRequestConfig) SetReplace(v string) *CreateClusterTemplateRequestConfig {
	s.Replace = &v
	return s
}

func (s *CreateClusterTemplateRequestConfig) SetFileName(v string) *CreateClusterTemplateRequestConfig {
	s.FileName = &v
	return s
}

func (s *CreateClusterTemplateRequestConfig) SetServiceName(v string) *CreateClusterTemplateRequestConfig {
	s.ServiceName = &v
	return s
}

func (s *CreateClusterTemplateRequestConfig) SetConfigKey(v string) *CreateClusterTemplateRequestConfig {
	s.ConfigKey = &v
	return s
}

func (s *CreateClusterTemplateRequestConfig) SetEncrypt(v string) *CreateClusterTemplateRequestConfig {
	s.Encrypt = &v
	return s
}

type CreateClusterTemplateRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateClusterTemplateRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterTemplateRequestTag) GoString() string {
	return s.String()
}

func (s *CreateClusterTemplateRequestTag) SetKey(v string) *CreateClusterTemplateRequestTag {
	s.Key = &v
	return s
}

func (s *CreateClusterTemplateRequestTag) SetValue(v string) *CreateClusterTemplateRequestTag {
	s.Value = &v
	return s
}

type CreateClusterTemplateResponseBody struct {
	ClusterTemplateId *string `json:"ClusterTemplateId,omitempty" xml:"ClusterTemplateId,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateClusterTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterTemplateResponseBody) SetClusterTemplateId(v string) *CreateClusterTemplateResponseBody {
	s.ClusterTemplateId = &v
	return s
}

func (s *CreateClusterTemplateResponseBody) SetRequestId(v string) *CreateClusterTemplateResponseBody {
	s.RequestId = &v
	return s
}

type CreateClusterTemplateResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateClusterTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateClusterTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateClusterTemplateResponse) SetHeaders(v map[string]*string) *CreateClusterTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateClusterTemplateResponse) SetBody(v *CreateClusterTemplateResponseBody) *CreateClusterTemplateResponse {
	s.Body = v
	return s
}

type UpdateLibraryInstallTaskStatusRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TaskBizId       *string `json:"TaskBizId,omitempty" xml:"TaskBizId,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateLibraryInstallTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLibraryInstallTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateLibraryInstallTaskStatusRequest) SetResourceOwnerId(v int64) *UpdateLibraryInstallTaskStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateLibraryInstallTaskStatusRequest) SetRegionId(v string) *UpdateLibraryInstallTaskStatusRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateLibraryInstallTaskStatusRequest) SetTaskBizId(v string) *UpdateLibraryInstallTaskStatusRequest {
	s.TaskBizId = &v
	return s
}

func (s *UpdateLibraryInstallTaskStatusRequest) SetStatus(v string) *UpdateLibraryInstallTaskStatusRequest {
	s.Status = &v
	return s
}

type UpdateLibraryInstallTaskStatusResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLibraryInstallTaskStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLibraryInstallTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLibraryInstallTaskStatusResponseBody) SetData(v bool) *UpdateLibraryInstallTaskStatusResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateLibraryInstallTaskStatusResponseBody) SetRequestId(v string) *UpdateLibraryInstallTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

type UpdateLibraryInstallTaskStatusResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateLibraryInstallTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateLibraryInstallTaskStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLibraryInstallTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateLibraryInstallTaskStatusResponse) SetHeaders(v map[string]*string) *UpdateLibraryInstallTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateLibraryInstallTaskStatusResponse) SetBody(v *UpdateLibraryInstallTaskStatusResponseBody) *UpdateLibraryInstallTaskStatusResponse {
	s.Body = v
	return s
}

type ListScalingConfigItemRequest struct {
	ResourceOwnerId   *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Limit             *int64  `json:"limit,omitempty" xml:"limit,omitempty"`
	PageNumber        *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentSize       *int64  `json:"CurrentSize,omitempty" xml:"CurrentSize,omitempty"`
	PageCount         *int64  `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	OrderField        *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	OrderMode         *string `json:"OrderMode,omitempty" xml:"OrderMode,omitempty"`
	ScalingGroupBizId *string `json:"ScalingGroupBizId,omitempty" xml:"ScalingGroupBizId,omitempty"`
	ConfigItemType    *string `json:"ConfigItemType,omitempty" xml:"ConfigItemType,omitempty"`
}

func (s ListScalingConfigItemRequest) String() string {
	return tea.Prettify(s)
}

func (s ListScalingConfigItemRequest) GoString() string {
	return s.String()
}

func (s *ListScalingConfigItemRequest) SetResourceOwnerId(v int64) *ListScalingConfigItemRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListScalingConfigItemRequest) SetRegionId(v string) *ListScalingConfigItemRequest {
	s.RegionId = &v
	return s
}

func (s *ListScalingConfigItemRequest) SetLimit(v int64) *ListScalingConfigItemRequest {
	s.Limit = &v
	return s
}

func (s *ListScalingConfigItemRequest) SetPageNumber(v int64) *ListScalingConfigItemRequest {
	s.PageNumber = &v
	return s
}

func (s *ListScalingConfigItemRequest) SetPageSize(v int64) *ListScalingConfigItemRequest {
	s.PageSize = &v
	return s
}

func (s *ListScalingConfigItemRequest) SetCurrentSize(v int64) *ListScalingConfigItemRequest {
	s.CurrentSize = &v
	return s
}

func (s *ListScalingConfigItemRequest) SetPageCount(v int64) *ListScalingConfigItemRequest {
	s.PageCount = &v
	return s
}

func (s *ListScalingConfigItemRequest) SetOrderField(v string) *ListScalingConfigItemRequest {
	s.OrderField = &v
	return s
}

func (s *ListScalingConfigItemRequest) SetOrderMode(v string) *ListScalingConfigItemRequest {
	s.OrderMode = &v
	return s
}

func (s *ListScalingConfigItemRequest) SetScalingGroupBizId(v string) *ListScalingConfigItemRequest {
	s.ScalingGroupBizId = &v
	return s
}

func (s *ListScalingConfigItemRequest) SetConfigItemType(v string) *ListScalingConfigItemRequest {
	s.ConfigItemType = &v
	return s
}

type ListScalingConfigItemResponseBody struct {
	// Id of the request
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int64                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int64                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	NextToken  *string                                 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Items      *ListScalingConfigItemResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s ListScalingConfigItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListScalingConfigItemResponseBody) GoString() string {
	return s.String()
}

func (s *ListScalingConfigItemResponseBody) SetRequestId(v string) *ListScalingConfigItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScalingConfigItemResponseBody) SetPageNumber(v int64) *ListScalingConfigItemResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListScalingConfigItemResponseBody) SetPageSize(v int64) *ListScalingConfigItemResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListScalingConfigItemResponseBody) SetTotalCount(v int64) *ListScalingConfigItemResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListScalingConfigItemResponseBody) SetNextToken(v string) *ListScalingConfigItemResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListScalingConfigItemResponseBody) SetItems(v *ListScalingConfigItemResponseBodyItems) *ListScalingConfigItemResponseBody {
	s.Items = v
	return s
}

type ListScalingConfigItemResponseBodyItems struct {
	Items []*ListScalingConfigItemResponseBodyItemsItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s ListScalingConfigItemResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s ListScalingConfigItemResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListScalingConfigItemResponseBodyItems) SetItems(v []*ListScalingConfigItemResponseBodyItemsItems) *ListScalingConfigItemResponseBodyItems {
	s.Items = v
	return s
}

type ListScalingConfigItemResponseBodyItemsItems struct {
	ConfigItemType         *string `json:"ConfigItemType,omitempty" xml:"ConfigItemType,omitempty"`
	ScalingGroupBizId      *string `json:"ScalingGroupBizId,omitempty" xml:"ScalingGroupBizId,omitempty"`
	ScalingConfigItemBizId *string `json:"ScalingConfigItemBizId,omitempty" xml:"ScalingConfigItemBizId,omitempty"`
	ConfigItemInformation  *string `json:"ConfigItemInformation,omitempty" xml:"ConfigItemInformation,omitempty"`
}

func (s ListScalingConfigItemResponseBodyItemsItems) String() string {
	return tea.Prettify(s)
}

func (s ListScalingConfigItemResponseBodyItemsItems) GoString() string {
	return s.String()
}

func (s *ListScalingConfigItemResponseBodyItemsItems) SetConfigItemType(v string) *ListScalingConfigItemResponseBodyItemsItems {
	s.ConfigItemType = &v
	return s
}

func (s *ListScalingConfigItemResponseBodyItemsItems) SetScalingGroupBizId(v string) *ListScalingConfigItemResponseBodyItemsItems {
	s.ScalingGroupBizId = &v
	return s
}

func (s *ListScalingConfigItemResponseBodyItemsItems) SetScalingConfigItemBizId(v string) *ListScalingConfigItemResponseBodyItemsItems {
	s.ScalingConfigItemBizId = &v
	return s
}

func (s *ListScalingConfigItemResponseBodyItemsItems) SetConfigItemInformation(v string) *ListScalingConfigItemResponseBodyItemsItems {
	s.ConfigItemInformation = &v
	return s
}

type ListScalingConfigItemResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListScalingConfigItemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListScalingConfigItemResponse) String() string {
	return tea.Prettify(s)
}

func (s ListScalingConfigItemResponse) GoString() string {
	return s.String()
}

func (s *ListScalingConfigItemResponse) SetHeaders(v map[string]*string) *ListScalingConfigItemResponse {
	s.Headers = v
	return s
}

func (s *ListScalingConfigItemResponse) SetBody(v *ListScalingConfigItemResponseBody) *ListScalingConfigItemResponse {
	s.Body = v
	return s
}

type ListFlowInstanceRequest struct {
	Id         *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	FlowId     *string   `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	FlowName   *string   `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	Owner      *string   `json:"Owner,omitempty" xml:"Owner,omitempty"`
	InstanceId *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TimeRange  *string   `json:"TimeRange,omitempty" xml:"TimeRange,omitempty"`
	OrderBy    *string   `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	OrderType  *string   `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProjectId  *string   `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	PageNumber *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StatusList []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
}

func (s ListFlowInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowInstanceRequest) GoString() string {
	return s.String()
}

func (s *ListFlowInstanceRequest) SetId(v string) *ListFlowInstanceRequest {
	s.Id = &v
	return s
}

func (s *ListFlowInstanceRequest) SetFlowId(v string) *ListFlowInstanceRequest {
	s.FlowId = &v
	return s
}

func (s *ListFlowInstanceRequest) SetFlowName(v string) *ListFlowInstanceRequest {
	s.FlowName = &v
	return s
}

func (s *ListFlowInstanceRequest) SetOwner(v string) *ListFlowInstanceRequest {
	s.Owner = &v
	return s
}

func (s *ListFlowInstanceRequest) SetInstanceId(v string) *ListFlowInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ListFlowInstanceRequest) SetTimeRange(v string) *ListFlowInstanceRequest {
	s.TimeRange = &v
	return s
}

func (s *ListFlowInstanceRequest) SetOrderBy(v string) *ListFlowInstanceRequest {
	s.OrderBy = &v
	return s
}

func (s *ListFlowInstanceRequest) SetOrderType(v string) *ListFlowInstanceRequest {
	s.OrderType = &v
	return s
}

func (s *ListFlowInstanceRequest) SetRegionId(v string) *ListFlowInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *ListFlowInstanceRequest) SetProjectId(v string) *ListFlowInstanceRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFlowInstanceRequest) SetPageNumber(v int32) *ListFlowInstanceRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFlowInstanceRequest) SetPageSize(v int32) *ListFlowInstanceRequest {
	s.PageSize = &v
	return s
}

func (s *ListFlowInstanceRequest) SetStatusList(v []*string) *ListFlowInstanceRequest {
	s.StatusList = v
	return s
}

type ListFlowInstanceResponseBody struct {
	RequestId     *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber    *int32                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total         *int32                                     `json:"Total,omitempty" xml:"Total,omitempty"`
	FlowInstances *ListFlowInstanceResponseBodyFlowInstances `json:"FlowInstances,omitempty" xml:"FlowInstances,omitempty" type:"Struct"`
}

func (s ListFlowInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowInstanceResponseBody) SetRequestId(v string) *ListFlowInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowInstanceResponseBody) SetPageNumber(v int32) *ListFlowInstanceResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListFlowInstanceResponseBody) SetPageSize(v int32) *ListFlowInstanceResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListFlowInstanceResponseBody) SetTotal(v int32) *ListFlowInstanceResponseBody {
	s.Total = &v
	return s
}

func (s *ListFlowInstanceResponseBody) SetFlowInstances(v *ListFlowInstanceResponseBodyFlowInstances) *ListFlowInstanceResponseBody {
	s.FlowInstances = v
	return s
}

type ListFlowInstanceResponseBodyFlowInstances struct {
	FlowInstance []*ListFlowInstanceResponseBodyFlowInstancesFlowInstance `json:"FlowInstance,omitempty" xml:"FlowInstance,omitempty" type:"Repeated"`
}

func (s ListFlowInstanceResponseBodyFlowInstances) String() string {
	return tea.Prettify(s)
}

func (s ListFlowInstanceResponseBodyFlowInstances) GoString() string {
	return s.String()
}

func (s *ListFlowInstanceResponseBodyFlowInstances) SetFlowInstance(v []*ListFlowInstanceResponseBodyFlowInstancesFlowInstance) *ListFlowInstanceResponseBodyFlowInstances {
	s.FlowInstance = v
	return s
}

type ListFlowInstanceResponseBodyFlowInstancesFlowInstance struct {
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Owner         *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	FlowName      *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	GmtModified   *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	FlowId        *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	EndTime       *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime     *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	HasNodeFailed *bool   `json:"HasNodeFailed,omitempty" xml:"HasNodeFailed,omitempty"`
	GmtCreate     *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	ScheduleTime  *int64  `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	Duration      *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Id            *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ClusterId     *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s ListFlowInstanceResponseBodyFlowInstancesFlowInstance) String() string {
	return tea.Prettify(s)
}

func (s ListFlowInstanceResponseBodyFlowInstancesFlowInstance) GoString() string {
	return s.String()
}

func (s *ListFlowInstanceResponseBodyFlowInstancesFlowInstance) SetStatus(v string) *ListFlowInstanceResponseBodyFlowInstancesFlowInstance {
	s.Status = &v
	return s
}

func (s *ListFlowInstanceResponseBodyFlowInstancesFlowInstance) SetOwner(v string) *ListFlowInstanceResponseBodyFlowInstancesFlowInstance {
	s.Owner = &v
	return s
}

func (s *ListFlowInstanceResponseBodyFlowInstancesFlowInstance) SetProjectId(v string) *ListFlowInstanceResponseBodyFlowInstancesFlowInstance {
	s.ProjectId = &v
	return s
}

func (s *ListFlowInstanceResponseBodyFlowInstancesFlowInstance) SetFlowName(v string) *ListFlowInstanceResponseBodyFlowInstancesFlowInstance {
	s.FlowName = &v
	return s
}

func (s *ListFlowInstanceResponseBodyFlowInstancesFlowInstance) SetGmtModified(v int64) *ListFlowInstanceResponseBodyFlowInstancesFlowInstance {
	s.GmtModified = &v
	return s
}

func (s *ListFlowInstanceResponseBodyFlowInstancesFlowInstance) SetFlowId(v string) *ListFlowInstanceResponseBodyFlowInstancesFlowInstance {
	s.FlowId = &v
	return s
}

func (s *ListFlowInstanceResponseBodyFlowInstancesFlowInstance) SetEndTime(v int64) *ListFlowInstanceResponseBodyFlowInstancesFlowInstance {
	s.EndTime = &v
	return s
}

func (s *ListFlowInstanceResponseBodyFlowInstancesFlowInstance) SetStartTime(v int64) *ListFlowInstanceResponseBodyFlowInstancesFlowInstance {
	s.StartTime = &v
	return s
}

func (s *ListFlowInstanceResponseBodyFlowInstancesFlowInstance) SetHasNodeFailed(v bool) *ListFlowInstanceResponseBodyFlowInstancesFlowInstance {
	s.HasNodeFailed = &v
	return s
}

func (s *ListFlowInstanceResponseBodyFlowInstancesFlowInstance) SetGmtCreate(v int64) *ListFlowInstanceResponseBodyFlowInstancesFlowInstance {
	s.GmtCreate = &v
	return s
}

func (s *ListFlowInstanceResponseBodyFlowInstancesFlowInstance) SetScheduleTime(v int64) *ListFlowInstanceResponseBodyFlowInstancesFlowInstance {
	s.ScheduleTime = &v
	return s
}

func (s *ListFlowInstanceResponseBodyFlowInstancesFlowInstance) SetDuration(v int64) *ListFlowInstanceResponseBodyFlowInstancesFlowInstance {
	s.Duration = &v
	return s
}

func (s *ListFlowInstanceResponseBodyFlowInstancesFlowInstance) SetId(v string) *ListFlowInstanceResponseBodyFlowInstancesFlowInstance {
	s.Id = &v
	return s
}

func (s *ListFlowInstanceResponseBodyFlowInstancesFlowInstance) SetClusterId(v string) *ListFlowInstanceResponseBodyFlowInstancesFlowInstance {
	s.ClusterId = &v
	return s
}

type ListFlowInstanceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFlowInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowInstanceResponse) GoString() string {
	return s.String()
}

func (s *ListFlowInstanceResponse) SetHeaders(v map[string]*string) *ListFlowInstanceResponse {
	s.Headers = v
	return s
}

func (s *ListFlowInstanceResponse) SetBody(v *ListFlowInstanceResponseBody) *ListFlowInstanceResponse {
	s.Body = v
	return s
}

type DescribeScalingMetricsRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	HostGroupId     *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
}

func (s DescribeScalingMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingMetricsRequest) GoString() string {
	return s.String()
}

func (s *DescribeScalingMetricsRequest) SetResourceOwnerId(v int64) *DescribeScalingMetricsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeScalingMetricsRequest) SetRegionId(v string) *DescribeScalingMetricsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeScalingMetricsRequest) SetClusterId(v string) *DescribeScalingMetricsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeScalingMetricsRequest) SetHostGroupId(v string) *DescribeScalingMetricsRequest {
	s.HostGroupId = &v
	return s
}

type DescribeScalingMetricsResponseBody struct {
	// Id of the request
	RequestId  *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MetricList []*DescribeScalingMetricsResponseBodyMetricList `json:"MetricList,omitempty" xml:"MetricList,omitempty" type:"Repeated"`
}

func (s DescribeScalingMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScalingMetricsResponseBody) SetRequestId(v string) *DescribeScalingMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScalingMetricsResponseBody) SetMetricList(v []*DescribeScalingMetricsResponseBodyMetricList) *DescribeScalingMetricsResponseBody {
	s.MetricList = v
	return s
}

type DescribeScalingMetricsResponseBodyMetricList struct {
	MetricName  *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Unit        *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	MinValue    *int64  `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	MaxValue    *int64  `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
}

func (s DescribeScalingMetricsResponseBodyMetricList) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingMetricsResponseBodyMetricList) GoString() string {
	return s.String()
}

func (s *DescribeScalingMetricsResponseBodyMetricList) SetMetricName(v string) *DescribeScalingMetricsResponseBodyMetricList {
	s.MetricName = &v
	return s
}

func (s *DescribeScalingMetricsResponseBodyMetricList) SetUnit(v string) *DescribeScalingMetricsResponseBodyMetricList {
	s.Unit = &v
	return s
}

func (s *DescribeScalingMetricsResponseBodyMetricList) SetDisplayName(v string) *DescribeScalingMetricsResponseBodyMetricList {
	s.DisplayName = &v
	return s
}

func (s *DescribeScalingMetricsResponseBodyMetricList) SetMinValue(v int64) *DescribeScalingMetricsResponseBodyMetricList {
	s.MinValue = &v
	return s
}

func (s *DescribeScalingMetricsResponseBodyMetricList) SetMaxValue(v int64) *DescribeScalingMetricsResponseBodyMetricList {
	s.MaxValue = &v
	return s
}

type DescribeScalingMetricsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeScalingMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScalingMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingMetricsResponse) GoString() string {
	return s.String()
}

func (s *DescribeScalingMetricsResponse) SetHeaders(v map[string]*string) *DescribeScalingMetricsResponse {
	s.Headers = v
	return s
}

func (s *DescribeScalingMetricsResponse) SetBody(v *DescribeScalingMetricsResponseBody) *DescribeScalingMetricsResponse {
	s.Body = v
	return s
}

type UntagResourcesSystemTagsRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 资源类型：cluster
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagOwnerUid  *int64  `json:"TagOwnerUid,omitempty" xml:"TagOwnerUid,omitempty"`
	// 资源ID
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	TagKeys     []*string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
	All         *bool     `json:"All,omitempty" xml:"All,omitempty"`
}

func (s UntagResourcesSystemTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesSystemTagsRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesSystemTagsRequest) SetResourceOwnerId(v int64) *UntagResourcesSystemTagsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UntagResourcesSystemTagsRequest) SetRegionId(v string) *UntagResourcesSystemTagsRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesSystemTagsRequest) SetResourceType(v string) *UntagResourcesSystemTagsRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesSystemTagsRequest) SetTagOwnerUid(v int64) *UntagResourcesSystemTagsRequest {
	s.TagOwnerUid = &v
	return s
}

func (s *UntagResourcesSystemTagsRequest) SetResourceIds(v []*string) *UntagResourcesSystemTagsRequest {
	s.ResourceIds = v
	return s
}

func (s *UntagResourcesSystemTagsRequest) SetTagKeys(v []*string) *UntagResourcesSystemTagsRequest {
	s.TagKeys = v
	return s
}

func (s *UntagResourcesSystemTagsRequest) SetAll(v bool) *UntagResourcesSystemTagsRequest {
	s.All = &v
	return s
}

type UntagResourcesSystemTagsResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 请求是否成功被处理
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 响应码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 错误码
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 响应消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s UntagResourcesSystemTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesSystemTagsResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesSystemTagsResponseBody) SetRequestId(v string) *UntagResourcesSystemTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UntagResourcesSystemTagsResponseBody) SetSuccess(v bool) *UntagResourcesSystemTagsResponseBody {
	s.Success = &v
	return s
}

func (s *UntagResourcesSystemTagsResponseBody) SetCode(v string) *UntagResourcesSystemTagsResponseBody {
	s.Code = &v
	return s
}

func (s *UntagResourcesSystemTagsResponseBody) SetErrorCode(v string) *UntagResourcesSystemTagsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UntagResourcesSystemTagsResponseBody) SetMessage(v string) *UntagResourcesSystemTagsResponseBody {
	s.Message = &v
	return s
}

type UntagResourcesSystemTagsResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UntagResourcesSystemTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UntagResourcesSystemTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesSystemTagsResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesSystemTagsResponse) SetHeaders(v map[string]*string) *UntagResourcesSystemTagsResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesSystemTagsResponse) SetBody(v *UntagResourcesSystemTagsResponseBody) *UntagResourcesSystemTagsResponse {
	s.Body = v
	return s
}

type DescribeFlowProjectRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeFlowProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowProjectRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowProjectRequest) SetProjectId(v string) *DescribeFlowProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeFlowProjectRequest) SetRegionId(v string) *DescribeFlowProjectRequest {
	s.RegionId = &v
	return s
}

type DescribeFlowProjectResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeFlowProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowProjectResponseBody) SetRequestId(v string) *DescribeFlowProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowProjectResponseBody) SetDescription(v string) *DescribeFlowProjectResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeFlowProjectResponseBody) SetUserId(v string) *DescribeFlowProjectResponseBody {
	s.UserId = &v
	return s
}

func (s *DescribeFlowProjectResponseBody) SetGmtCreate(v int64) *DescribeFlowProjectResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *DescribeFlowProjectResponseBody) SetGmtModified(v int64) *DescribeFlowProjectResponseBody {
	s.GmtModified = &v
	return s
}

func (s *DescribeFlowProjectResponseBody) SetName(v string) *DescribeFlowProjectResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeFlowProjectResponseBody) SetId(v string) *DescribeFlowProjectResponseBody {
	s.Id = &v
	return s
}

type DescribeFlowProjectResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFlowProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFlowProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowProjectResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowProjectResponse) SetHeaders(v map[string]*string) *DescribeFlowProjectResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowProjectResponse) SetBody(v *DescribeFlowProjectResponseBody) *DescribeFlowProjectResponse {
	s.Body = v
	return s
}

type DeleteSecurityWhiteListRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	WhiteIp   *string `json:"WhiteIp,omitempty" xml:"WhiteIp,omitempty"`
}

func (s DeleteSecurityWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecurityWhiteListRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecurityWhiteListRequest) SetClusterId(v string) *DeleteSecurityWhiteListRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteSecurityWhiteListRequest) SetPortRange(v string) *DeleteSecurityWhiteListRequest {
	s.PortRange = &v
	return s
}

func (s *DeleteSecurityWhiteListRequest) SetWhiteIp(v string) *DeleteSecurityWhiteListRequest {
	s.WhiteIp = &v
	return s
}

type DeleteSecurityWhiteListResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSecurityWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecurityWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecurityWhiteListResponseBody) SetRequestId(v string) *DeleteSecurityWhiteListResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSecurityWhiteListResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSecurityWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSecurityWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecurityWhiteListResponse) GoString() string {
	return s.String()
}

func (s *DeleteSecurityWhiteListResponse) SetHeaders(v map[string]*string) *DeleteSecurityWhiteListResponse {
	s.Headers = v
	return s
}

func (s *DeleteSecurityWhiteListResponse) SetBody(v *DeleteSecurityWhiteListResponseBody) *DeleteSecurityWhiteListResponse {
	s.Body = v
	return s
}

type ListScalingActivityRequest struct {
	ResourceOwnerId   *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Limit             *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	PageNumber        *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentSize       *int32  `json:"CurrentSize,omitempty" xml:"CurrentSize,omitempty"`
	PageCount         *int32  `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	OrderField        *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	OrderMode         *string `json:"OrderMode,omitempty" xml:"OrderMode,omitempty"`
	ClusterBizId      *string `json:"ClusterBizId,omitempty" xml:"ClusterBizId,omitempty"`
	HostGroupId       *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	ScalingGroupBizId *string `json:"ScalingGroupBizId,omitempty" xml:"ScalingGroupBizId,omitempty"`
	ScalingRuleName   *string `json:"ScalingRuleName,omitempty" xml:"ScalingRuleName,omitempty"`
	HostGroupName     *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListScalingActivityRequest) String() string {
	return tea.Prettify(s)
}

func (s ListScalingActivityRequest) GoString() string {
	return s.String()
}

func (s *ListScalingActivityRequest) SetResourceOwnerId(v int64) *ListScalingActivityRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListScalingActivityRequest) SetRegionId(v string) *ListScalingActivityRequest {
	s.RegionId = &v
	return s
}

func (s *ListScalingActivityRequest) SetLimit(v int32) *ListScalingActivityRequest {
	s.Limit = &v
	return s
}

func (s *ListScalingActivityRequest) SetPageNumber(v int32) *ListScalingActivityRequest {
	s.PageNumber = &v
	return s
}

func (s *ListScalingActivityRequest) SetPageSize(v int32) *ListScalingActivityRequest {
	s.PageSize = &v
	return s
}

func (s *ListScalingActivityRequest) SetCurrentSize(v int32) *ListScalingActivityRequest {
	s.CurrentSize = &v
	return s
}

func (s *ListScalingActivityRequest) SetPageCount(v int32) *ListScalingActivityRequest {
	s.PageCount = &v
	return s
}

func (s *ListScalingActivityRequest) SetOrderField(v string) *ListScalingActivityRequest {
	s.OrderField = &v
	return s
}

func (s *ListScalingActivityRequest) SetOrderMode(v string) *ListScalingActivityRequest {
	s.OrderMode = &v
	return s
}

func (s *ListScalingActivityRequest) SetClusterBizId(v string) *ListScalingActivityRequest {
	s.ClusterBizId = &v
	return s
}

func (s *ListScalingActivityRequest) SetHostGroupId(v string) *ListScalingActivityRequest {
	s.HostGroupId = &v
	return s
}

func (s *ListScalingActivityRequest) SetScalingGroupBizId(v string) *ListScalingActivityRequest {
	s.ScalingGroupBizId = &v
	return s
}

func (s *ListScalingActivityRequest) SetScalingRuleName(v string) *ListScalingActivityRequest {
	s.ScalingRuleName = &v
	return s
}

func (s *ListScalingActivityRequest) SetHostGroupName(v string) *ListScalingActivityRequest {
	s.HostGroupName = &v
	return s
}

func (s *ListScalingActivityRequest) SetStatus(v string) *ListScalingActivityRequest {
	s.Status = &v
	return s
}

type ListScalingActivityResponseBody struct {
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// pageNumber
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// pageSize
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// totalCount
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// nextToken
	NextToken *int64                                `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Items     *ListScalingActivityResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s ListScalingActivityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListScalingActivityResponseBody) GoString() string {
	return s.String()
}

func (s *ListScalingActivityResponseBody) SetRequestId(v string) *ListScalingActivityResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScalingActivityResponseBody) SetPageNumber(v int64) *ListScalingActivityResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListScalingActivityResponseBody) SetPageSize(v int64) *ListScalingActivityResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListScalingActivityResponseBody) SetTotalCount(v int64) *ListScalingActivityResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListScalingActivityResponseBody) SetNextToken(v int64) *ListScalingActivityResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListScalingActivityResponseBody) SetItems(v *ListScalingActivityResponseBodyItems) *ListScalingActivityResponseBody {
	s.Items = v
	return s
}

type ListScalingActivityResponseBodyItems struct {
	Item []*ListScalingActivityResponseBodyItemsItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s ListScalingActivityResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s ListScalingActivityResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListScalingActivityResponseBodyItems) SetItem(v []*ListScalingActivityResponseBodyItemsItem) *ListScalingActivityResponseBodyItems {
	s.Item = v
	return s
}

type ListScalingActivityResponseBodyItemsItem struct {
	// bizId
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// startTime
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// endTime
	EndTime   *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExpectNum *int64 `json:"ExpectNum,omitempty" xml:"ExpectNum,omitempty"`
	// instanceIds
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// totalCapacity
	TotalCapacity *int64 `json:"TotalCapacity,omitempty" xml:"TotalCapacity,omitempty"`
	// cause
	Cause *string `json:"Cause,omitempty" xml:"Cause,omitempty"`
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// transition
	Transition *string `json:"Transition,omitempty" xml:"Transition,omitempty"`
	// status
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// scalingRuleId
	ScalingRuleId *string `json:"ScalingRuleId,omitempty" xml:"ScalingRuleId,omitempty"`
	// scalingRuleName
	ScalingRuleName *string `json:"ScalingRuleName,omitempty" xml:"ScalingRuleName,omitempty"`
	// hostGroupBizId
	HostGroupBizId *string `json:"HostGroupBizId,omitempty" xml:"HostGroupBizId,omitempty"`
	// hostGroupName
	HostGroupName *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
}

func (s ListScalingActivityResponseBodyItemsItem) String() string {
	return tea.Prettify(s)
}

func (s ListScalingActivityResponseBodyItemsItem) GoString() string {
	return s.String()
}

func (s *ListScalingActivityResponseBodyItemsItem) SetBizId(v string) *ListScalingActivityResponseBodyItemsItem {
	s.BizId = &v
	return s
}

func (s *ListScalingActivityResponseBodyItemsItem) SetStartTime(v int64) *ListScalingActivityResponseBodyItemsItem {
	s.StartTime = &v
	return s
}

func (s *ListScalingActivityResponseBodyItemsItem) SetEndTime(v int64) *ListScalingActivityResponseBodyItemsItem {
	s.EndTime = &v
	return s
}

func (s *ListScalingActivityResponseBodyItemsItem) SetExpectNum(v int64) *ListScalingActivityResponseBodyItemsItem {
	s.ExpectNum = &v
	return s
}

func (s *ListScalingActivityResponseBodyItemsItem) SetInstanceIds(v string) *ListScalingActivityResponseBodyItemsItem {
	s.InstanceIds = &v
	return s
}

func (s *ListScalingActivityResponseBodyItemsItem) SetTotalCapacity(v int64) *ListScalingActivityResponseBodyItemsItem {
	s.TotalCapacity = &v
	return s
}

func (s *ListScalingActivityResponseBodyItemsItem) SetCause(v string) *ListScalingActivityResponseBodyItemsItem {
	s.Cause = &v
	return s
}

func (s *ListScalingActivityResponseBodyItemsItem) SetDescription(v string) *ListScalingActivityResponseBodyItemsItem {
	s.Description = &v
	return s
}

func (s *ListScalingActivityResponseBodyItemsItem) SetTransition(v string) *ListScalingActivityResponseBodyItemsItem {
	s.Transition = &v
	return s
}

func (s *ListScalingActivityResponseBodyItemsItem) SetStatus(v string) *ListScalingActivityResponseBodyItemsItem {
	s.Status = &v
	return s
}

func (s *ListScalingActivityResponseBodyItemsItem) SetScalingRuleId(v string) *ListScalingActivityResponseBodyItemsItem {
	s.ScalingRuleId = &v
	return s
}

func (s *ListScalingActivityResponseBodyItemsItem) SetScalingRuleName(v string) *ListScalingActivityResponseBodyItemsItem {
	s.ScalingRuleName = &v
	return s
}

func (s *ListScalingActivityResponseBodyItemsItem) SetHostGroupBizId(v string) *ListScalingActivityResponseBodyItemsItem {
	s.HostGroupBizId = &v
	return s
}

func (s *ListScalingActivityResponseBodyItemsItem) SetHostGroupName(v string) *ListScalingActivityResponseBodyItemsItem {
	s.HostGroupName = &v
	return s
}

type ListScalingActivityResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListScalingActivityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListScalingActivityResponse) String() string {
	return tea.Prettify(s)
}

func (s ListScalingActivityResponse) GoString() string {
	return s.String()
}

func (s *ListScalingActivityResponse) SetHeaders(v map[string]*string) *ListScalingActivityResponse {
	s.Headers = v
	return s
}

func (s *ListScalingActivityResponse) SetBody(v *ListScalingActivityResponseBody) *ListScalingActivityResponse {
	s.Body = v
	return s
}

type ListTagValuesRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// 资源组类型
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Scope        *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PageSize     *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Key          *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s ListTagValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesRequest) GoString() string {
	return s.String()
}

func (s *ListTagValuesRequest) SetResourceOwnerId(v int64) *ListTagValuesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTagValuesRequest) SetRegionId(v string) *ListTagValuesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagValuesRequest) SetResourceGroupId(v string) *ListTagValuesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTagValuesRequest) SetResourceType(v string) *ListTagValuesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagValuesRequest) SetScope(v string) *ListTagValuesRequest {
	s.Scope = &v
	return s
}

func (s *ListTagValuesRequest) SetNextToken(v string) *ListTagValuesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesRequest) SetPageSize(v int64) *ListTagValuesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagValuesRequest) SetKey(v string) *ListTagValuesRequest {
	s.Key = &v
	return s
}

type ListTagValuesResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 是否成功响应
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 是否分页
	Paging *bool `json:"Paging,omitempty" xml:"Paging,omitempty"`
	// 标签值集合
	Data      []*ListTagValuesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorCode *string                          `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ListTagValuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponseBody) SetRequestId(v string) *ListTagValuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagValuesResponseBody) SetSuccess(v bool) *ListTagValuesResponseBody {
	s.Success = &v
	return s
}

func (s *ListTagValuesResponseBody) SetPaging(v bool) *ListTagValuesResponseBody {
	s.Paging = &v
	return s
}

func (s *ListTagValuesResponseBody) SetData(v []*ListTagValuesResponseBodyData) *ListTagValuesResponseBody {
	s.Data = v
	return s
}

func (s *ListTagValuesResponseBody) SetCode(v string) *ListTagValuesResponseBody {
	s.Code = &v
	return s
}

func (s *ListTagValuesResponseBody) SetErrorCode(v string) *ListTagValuesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListTagValuesResponseBody) SetMessage(v string) *ListTagValuesResponseBody {
	s.Message = &v
	return s
}

type ListTagValuesResponseBodyData struct {
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 标签键集合
	Items []*string `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s ListTagValuesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponseBodyData) SetPageNumber(v int64) *ListTagValuesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListTagValuesResponseBodyData) SetPageSize(v int64) *ListTagValuesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListTagValuesResponseBodyData) SetTotalCount(v int64) *ListTagValuesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListTagValuesResponseBodyData) SetNextToken(v string) *ListTagValuesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesResponseBodyData) SetItems(v []*string) *ListTagValuesResponseBodyData {
	s.Items = v
	return s
}

type ListTagValuesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagValuesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponse) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponse) SetHeaders(v map[string]*string) *ListTagValuesResponse {
	s.Headers = v
	return s
}

func (s *ListTagValuesResponse) SetBody(v *ListTagValuesResponseBody) *ListTagValuesResponse {
	s.Body = v
	return s
}

type ListClusterInstalledServiceRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListClusterInstalledServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClusterInstalledServiceRequest) GoString() string {
	return s.String()
}

func (s *ListClusterInstalledServiceRequest) SetResourceOwnerId(v int64) *ListClusterInstalledServiceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListClusterInstalledServiceRequest) SetRegionId(v string) *ListClusterInstalledServiceRequest {
	s.RegionId = &v
	return s
}

func (s *ListClusterInstalledServiceRequest) SetClusterId(v string) *ListClusterInstalledServiceRequest {
	s.ClusterId = &v
	return s
}

func (s *ListClusterInstalledServiceRequest) SetPageNumber(v int32) *ListClusterInstalledServiceRequest {
	s.PageNumber = &v
	return s
}

func (s *ListClusterInstalledServiceRequest) SetPageSize(v int32) *ListClusterInstalledServiceRequest {
	s.PageSize = &v
	return s
}

type ListClusterInstalledServiceResponseBody struct {
	RequestId                   *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ClusterInstalledServiceList *ListClusterInstalledServiceResponseBodyClusterInstalledServiceList `json:"ClusterInstalledServiceList,omitempty" xml:"ClusterInstalledServiceList,omitempty" type:"Struct"`
}

func (s ListClusterInstalledServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClusterInstalledServiceResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterInstalledServiceResponseBody) SetRequestId(v string) *ListClusterInstalledServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterInstalledServiceResponseBody) SetClusterInstalledServiceList(v *ListClusterInstalledServiceResponseBodyClusterInstalledServiceList) *ListClusterInstalledServiceResponseBody {
	s.ClusterInstalledServiceList = v
	return s
}

type ListClusterInstalledServiceResponseBodyClusterInstalledServiceList struct {
	ClusterInstalledService []*ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledService `json:"ClusterInstalledService,omitempty" xml:"ClusterInstalledService,omitempty" type:"Repeated"`
}

func (s ListClusterInstalledServiceResponseBodyClusterInstalledServiceList) String() string {
	return tea.Prettify(s)
}

func (s ListClusterInstalledServiceResponseBodyClusterInstalledServiceList) GoString() string {
	return s.String()
}

func (s *ListClusterInstalledServiceResponseBodyClusterInstalledServiceList) SetClusterInstalledService(v []*ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledService) *ListClusterInstalledServiceResponseBodyClusterInstalledServiceList {
	s.ClusterInstalledService = v
	return s
}

type ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledService struct {
	ServiceEcmVersion  *string                                                                                                     `json:"ServiceEcmVersion,omitempty" xml:"ServiceEcmVersion,omitempty"`
	ServiceDisplayName *string                                                                                                     `json:"ServiceDisplayName,omitempty" xml:"ServiceDisplayName,omitempty"`
	OnlyClient         *bool                                                                                                       `json:"OnlyClient,omitempty" xml:"OnlyClient,omitempty"`
	Comment            *string                                                                                                     `json:"Comment,omitempty" xml:"Comment,omitempty"`
	NotStartedNum      *int32                                                                                                      `json:"NotStartedNum,omitempty" xml:"NotStartedNum,omitempty"`
	NeedRestartNum     *int32                                                                                                      `json:"NeedRestartNum,omitempty" xml:"NeedRestartNum,omitempty"`
	ServiceVersion     *string                                                                                                     `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	ServiceStatus      *string                                                                                                     `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	ServiceName        *string                                                                                                     `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	AbnormalNum        *int32                                                                                                      `json:"AbnormalNum,omitempty" xml:"AbnormalNum,omitempty"`
	ServiceActionList  *ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledServiceServiceActionList `json:"ServiceActionList,omitempty" xml:"ServiceActionList,omitempty" type:"Struct"`
}

func (s ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledService) String() string {
	return tea.Prettify(s)
}

func (s ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledService) GoString() string {
	return s.String()
}

func (s *ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledService) SetServiceEcmVersion(v string) *ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledService {
	s.ServiceEcmVersion = &v
	return s
}

func (s *ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledService) SetServiceDisplayName(v string) *ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledService {
	s.ServiceDisplayName = &v
	return s
}

func (s *ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledService) SetOnlyClient(v bool) *ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledService {
	s.OnlyClient = &v
	return s
}

func (s *ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledService) SetComment(v string) *ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledService {
	s.Comment = &v
	return s
}

func (s *ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledService) SetNotStartedNum(v int32) *ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledService {
	s.NotStartedNum = &v
	return s
}

func (s *ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledService) SetNeedRestartNum(v int32) *ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledService {
	s.NeedRestartNum = &v
	return s
}

func (s *ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledService) SetServiceVersion(v string) *ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledService {
	s.ServiceVersion = &v
	return s
}

func (s *ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledService) SetServiceStatus(v string) *ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledService {
	s.ServiceStatus = &v
	return s
}

func (s *ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledService) SetServiceName(v string) *ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledService {
	s.ServiceName = &v
	return s
}

func (s *ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledService) SetAbnormalNum(v int32) *ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledService {
	s.AbnormalNum = &v
	return s
}

func (s *ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledService) SetServiceActionList(v *ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledServiceServiceActionList) *ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledService {
	s.ServiceActionList = v
	return s
}

type ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledServiceServiceActionList struct {
	ServiceAction []*ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledServiceServiceActionListServiceAction `json:"ServiceAction,omitempty" xml:"ServiceAction,omitempty" type:"Repeated"`
}

func (s ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledServiceServiceActionList) String() string {
	return tea.Prettify(s)
}

func (s ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledServiceServiceActionList) GoString() string {
	return s.String()
}

func (s *ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledServiceServiceActionList) SetServiceAction(v []*ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledServiceServiceActionListServiceAction) *ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledServiceServiceActionList {
	s.ServiceAction = v
	return s
}

type ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledServiceServiceActionListServiceAction struct {
	DisplayName   *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	ActionName    *string `json:"ActionName,omitempty" xml:"ActionName,omitempty"`
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	ServiceName   *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	Command       *string `json:"Command,omitempty" xml:"Command,omitempty"`
}

func (s ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledServiceServiceActionListServiceAction) String() string {
	return tea.Prettify(s)
}

func (s ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledServiceServiceActionListServiceAction) GoString() string {
	return s.String()
}

func (s *ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledServiceServiceActionListServiceAction) SetDisplayName(v string) *ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledServiceServiceActionListServiceAction {
	s.DisplayName = &v
	return s
}

func (s *ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledServiceServiceActionListServiceAction) SetActionName(v string) *ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledServiceServiceActionListServiceAction {
	s.ActionName = &v
	return s
}

func (s *ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledServiceServiceActionListServiceAction) SetComponentName(v string) *ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledServiceServiceActionListServiceAction {
	s.ComponentName = &v
	return s
}

func (s *ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledServiceServiceActionListServiceAction) SetServiceName(v string) *ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledServiceServiceActionListServiceAction {
	s.ServiceName = &v
	return s
}

func (s *ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledServiceServiceActionListServiceAction) SetCommand(v string) *ListClusterInstalledServiceResponseBodyClusterInstalledServiceListClusterInstalledServiceServiceActionListServiceAction {
	s.Command = &v
	return s
}

type ListClusterInstalledServiceResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListClusterInstalledServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClusterInstalledServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClusterInstalledServiceResponse) GoString() string {
	return s.String()
}

func (s *ListClusterInstalledServiceResponse) SetHeaders(v map[string]*string) *ListClusterInstalledServiceResponse {
	s.Headers = v
	return s
}

func (s *ListClusterInstalledServiceResponse) SetBody(v *ListClusterInstalledServiceResponseBody) *ListClusterInstalledServiceResponse {
	s.Body = v
	return s
}

type RunClusterServiceActionRequest struct {
	ResourceOwnerId             *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId                    *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClusterId                   *string   `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	HostIdList                  *string   `json:"HostIdList,omitempty" xml:"HostIdList,omitempty"`
	ServiceName                 *string   `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ServiceActionName           *string   `json:"ServiceActionName,omitempty" xml:"ServiceActionName,omitempty"`
	CustomCommand               *string   `json:"CustomCommand,omitempty" xml:"CustomCommand,omitempty"`
	ComponentNameList           *string   `json:"ComponentNameList,omitempty" xml:"ComponentNameList,omitempty"`
	Comment                     *string   `json:"Comment,omitempty" xml:"Comment,omitempty"`
	IsRolling                   *bool     `json:"IsRolling,omitempty" xml:"IsRolling,omitempty"`
	ExecuteStrategy             *string   `json:"ExecuteStrategy,omitempty" xml:"ExecuteStrategy,omitempty"`
	CustomParams                *string   `json:"CustomParams,omitempty" xml:"CustomParams,omitempty"`
	Interval                    *int64    `json:"Interval,omitempty" xml:"Interval,omitempty"`
	NodeCountPerBatch           *int32    `json:"NodeCountPerBatch,omitempty" xml:"NodeCountPerBatch,omitempty"`
	TolerateFailCount           *int32    `json:"TolerateFailCount,omitempty" xml:"TolerateFailCount,omitempty"`
	OnlyRestartStaleConfigNodes *bool     `json:"OnlyRestartStaleConfigNodes,omitempty" xml:"OnlyRestartStaleConfigNodes,omitempty"`
	TurnOnMaintenanceMode       *bool     `json:"TurnOnMaintenanceMode,omitempty" xml:"TurnOnMaintenanceMode,omitempty"`
	HostGroupIdList             []*string `json:"HostGroupIdList,omitempty" xml:"HostGroupIdList,omitempty" type:"Repeated"`
}

func (s RunClusterServiceActionRequest) String() string {
	return tea.Prettify(s)
}

func (s RunClusterServiceActionRequest) GoString() string {
	return s.String()
}

func (s *RunClusterServiceActionRequest) SetResourceOwnerId(v int64) *RunClusterServiceActionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RunClusterServiceActionRequest) SetRegionId(v string) *RunClusterServiceActionRequest {
	s.RegionId = &v
	return s
}

func (s *RunClusterServiceActionRequest) SetClusterId(v string) *RunClusterServiceActionRequest {
	s.ClusterId = &v
	return s
}

func (s *RunClusterServiceActionRequest) SetHostIdList(v string) *RunClusterServiceActionRequest {
	s.HostIdList = &v
	return s
}

func (s *RunClusterServiceActionRequest) SetServiceName(v string) *RunClusterServiceActionRequest {
	s.ServiceName = &v
	return s
}

func (s *RunClusterServiceActionRequest) SetServiceActionName(v string) *RunClusterServiceActionRequest {
	s.ServiceActionName = &v
	return s
}

func (s *RunClusterServiceActionRequest) SetCustomCommand(v string) *RunClusterServiceActionRequest {
	s.CustomCommand = &v
	return s
}

func (s *RunClusterServiceActionRequest) SetComponentNameList(v string) *RunClusterServiceActionRequest {
	s.ComponentNameList = &v
	return s
}

func (s *RunClusterServiceActionRequest) SetComment(v string) *RunClusterServiceActionRequest {
	s.Comment = &v
	return s
}

func (s *RunClusterServiceActionRequest) SetIsRolling(v bool) *RunClusterServiceActionRequest {
	s.IsRolling = &v
	return s
}

func (s *RunClusterServiceActionRequest) SetExecuteStrategy(v string) *RunClusterServiceActionRequest {
	s.ExecuteStrategy = &v
	return s
}

func (s *RunClusterServiceActionRequest) SetCustomParams(v string) *RunClusterServiceActionRequest {
	s.CustomParams = &v
	return s
}

func (s *RunClusterServiceActionRequest) SetInterval(v int64) *RunClusterServiceActionRequest {
	s.Interval = &v
	return s
}

func (s *RunClusterServiceActionRequest) SetNodeCountPerBatch(v int32) *RunClusterServiceActionRequest {
	s.NodeCountPerBatch = &v
	return s
}

func (s *RunClusterServiceActionRequest) SetTolerateFailCount(v int32) *RunClusterServiceActionRequest {
	s.TolerateFailCount = &v
	return s
}

func (s *RunClusterServiceActionRequest) SetOnlyRestartStaleConfigNodes(v bool) *RunClusterServiceActionRequest {
	s.OnlyRestartStaleConfigNodes = &v
	return s
}

func (s *RunClusterServiceActionRequest) SetTurnOnMaintenanceMode(v bool) *RunClusterServiceActionRequest {
	s.TurnOnMaintenanceMode = &v
	return s
}

func (s *RunClusterServiceActionRequest) SetHostGroupIdList(v []*string) *RunClusterServiceActionRequest {
	s.HostGroupIdList = v
	return s
}

type RunClusterServiceActionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunClusterServiceActionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunClusterServiceActionResponseBody) GoString() string {
	return s.String()
}

func (s *RunClusterServiceActionResponseBody) SetRequestId(v string) *RunClusterServiceActionResponseBody {
	s.RequestId = &v
	return s
}

type RunClusterServiceActionResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RunClusterServiceActionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RunClusterServiceActionResponse) String() string {
	return tea.Prettify(s)
}

func (s RunClusterServiceActionResponse) GoString() string {
	return s.String()
}

func (s *RunClusterServiceActionResponse) SetHeaders(v map[string]*string) *RunClusterServiceActionResponse {
	s.Headers = v
	return s
}

func (s *RunClusterServiceActionResponse) SetBody(v *RunClusterServiceActionResponseBody) *RunClusterServiceActionResponse {
	s.Body = v
	return s
}

type SuspendFlowRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProjectId      *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	FlowInstanceId *string `json:"FlowInstanceId,omitempty" xml:"FlowInstanceId,omitempty"`
}

func (s SuspendFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s SuspendFlowRequest) GoString() string {
	return s.String()
}

func (s *SuspendFlowRequest) SetRegionId(v string) *SuspendFlowRequest {
	s.RegionId = &v
	return s
}

func (s *SuspendFlowRequest) SetProjectId(v string) *SuspendFlowRequest {
	s.ProjectId = &v
	return s
}

func (s *SuspendFlowRequest) SetFlowInstanceId(v string) *SuspendFlowRequest {
	s.FlowInstanceId = &v
	return s
}

type SuspendFlowResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SuspendFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SuspendFlowResponseBody) GoString() string {
	return s.String()
}

func (s *SuspendFlowResponseBody) SetData(v bool) *SuspendFlowResponseBody {
	s.Data = &v
	return s
}

func (s *SuspendFlowResponseBody) SetRequestId(v string) *SuspendFlowResponseBody {
	s.RequestId = &v
	return s
}

type SuspendFlowResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SuspendFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SuspendFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s SuspendFlowResponse) GoString() string {
	return s.String()
}

func (s *SuspendFlowResponse) SetHeaders(v map[string]*string) *SuspendFlowResponse {
	s.Headers = v
	return s
}

func (s *SuspendFlowResponse) SetBody(v *SuspendFlowResponseBody) *SuspendFlowResponse {
	s.Body = v
	return s
}

type CreateFlowProjectRequest struct {
	ProductType     *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateFlowProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowProjectRequest) SetProductType(v string) *CreateFlowProjectRequest {
	s.ProductType = &v
	return s
}

func (s *CreateFlowProjectRequest) SetRegionId(v string) *CreateFlowProjectRequest {
	s.RegionId = &v
	return s
}

func (s *CreateFlowProjectRequest) SetName(v string) *CreateFlowProjectRequest {
	s.Name = &v
	return s
}

func (s *CreateFlowProjectRequest) SetDescription(v string) *CreateFlowProjectRequest {
	s.Description = &v
	return s
}

func (s *CreateFlowProjectRequest) SetResourceGroupId(v string) *CreateFlowProjectRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateFlowProjectRequest) SetClientToken(v string) *CreateFlowProjectRequest {
	s.ClientToken = &v
	return s
}

type CreateFlowProjectResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFlowProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlowProjectResponseBody) SetId(v string) *CreateFlowProjectResponseBody {
	s.Id = &v
	return s
}

func (s *CreateFlowProjectResponseBody) SetRequestId(v string) *CreateFlowProjectResponseBody {
	s.RequestId = &v
	return s
}

type CreateFlowProjectResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFlowProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFlowProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateFlowProjectResponse) SetHeaders(v map[string]*string) *CreateFlowProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateFlowProjectResponse) SetBody(v *CreateFlowProjectResponseBody) *CreateFlowProjectResponse {
	s.Body = v
	return s
}

type ListFlowNodeInstanceContainerStatusRequest struct {
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	NodeInstanceId *string `json:"NodeInstanceId,omitempty" xml:"NodeInstanceId,omitempty"`
	ProjectId      *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListFlowNodeInstanceContainerStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowNodeInstanceContainerStatusRequest) GoString() string {
	return s.String()
}

func (s *ListFlowNodeInstanceContainerStatusRequest) SetPageNumber(v int32) *ListFlowNodeInstanceContainerStatusRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFlowNodeInstanceContainerStatusRequest) SetPageSize(v int32) *ListFlowNodeInstanceContainerStatusRequest {
	s.PageSize = &v
	return s
}

func (s *ListFlowNodeInstanceContainerStatusRequest) SetNodeInstanceId(v string) *ListFlowNodeInstanceContainerStatusRequest {
	s.NodeInstanceId = &v
	return s
}

func (s *ListFlowNodeInstanceContainerStatusRequest) SetProjectId(v string) *ListFlowNodeInstanceContainerStatusRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFlowNodeInstanceContainerStatusRequest) SetRegionId(v string) *ListFlowNodeInstanceContainerStatusRequest {
	s.RegionId = &v
	return s
}

type ListFlowNodeInstanceContainerStatusResponseBody struct {
	RequestId           *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber          *int32                                                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize            *int32                                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total               *int32                                                              `json:"Total,omitempty" xml:"Total,omitempty"`
	ContainerStatusList *ListFlowNodeInstanceContainerStatusResponseBodyContainerStatusList `json:"ContainerStatusList,omitempty" xml:"ContainerStatusList,omitempty" type:"Struct"`
}

func (s ListFlowNodeInstanceContainerStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowNodeInstanceContainerStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowNodeInstanceContainerStatusResponseBody) SetRequestId(v string) *ListFlowNodeInstanceContainerStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowNodeInstanceContainerStatusResponseBody) SetPageNumber(v int32) *ListFlowNodeInstanceContainerStatusResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListFlowNodeInstanceContainerStatusResponseBody) SetPageSize(v int32) *ListFlowNodeInstanceContainerStatusResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListFlowNodeInstanceContainerStatusResponseBody) SetTotal(v int32) *ListFlowNodeInstanceContainerStatusResponseBody {
	s.Total = &v
	return s
}

func (s *ListFlowNodeInstanceContainerStatusResponseBody) SetContainerStatusList(v *ListFlowNodeInstanceContainerStatusResponseBodyContainerStatusList) *ListFlowNodeInstanceContainerStatusResponseBody {
	s.ContainerStatusList = v
	return s
}

type ListFlowNodeInstanceContainerStatusResponseBodyContainerStatusList struct {
	ContainerStatus []*ListFlowNodeInstanceContainerStatusResponseBodyContainerStatusListContainerStatus `json:"ContainerStatus,omitempty" xml:"ContainerStatus,omitempty" type:"Repeated"`
}

func (s ListFlowNodeInstanceContainerStatusResponseBodyContainerStatusList) String() string {
	return tea.Prettify(s)
}

func (s ListFlowNodeInstanceContainerStatusResponseBodyContainerStatusList) GoString() string {
	return s.String()
}

func (s *ListFlowNodeInstanceContainerStatusResponseBodyContainerStatusList) SetContainerStatus(v []*ListFlowNodeInstanceContainerStatusResponseBodyContainerStatusListContainerStatus) *ListFlowNodeInstanceContainerStatusResponseBodyContainerStatusList {
	s.ContainerStatus = v
	return s
}

type ListFlowNodeInstanceContainerStatusResponseBodyContainerStatusListContainerStatus struct {
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	HostName      *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	ContainerId   *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
}

func (s ListFlowNodeInstanceContainerStatusResponseBodyContainerStatusListContainerStatus) String() string {
	return tea.Prettify(s)
}

func (s ListFlowNodeInstanceContainerStatusResponseBodyContainerStatusListContainerStatus) GoString() string {
	return s.String()
}

func (s *ListFlowNodeInstanceContainerStatusResponseBodyContainerStatusListContainerStatus) SetStatus(v string) *ListFlowNodeInstanceContainerStatusResponseBodyContainerStatusListContainerStatus {
	s.Status = &v
	return s
}

func (s *ListFlowNodeInstanceContainerStatusResponseBodyContainerStatusListContainerStatus) SetHostName(v string) *ListFlowNodeInstanceContainerStatusResponseBodyContainerStatusListContainerStatus {
	s.HostName = &v
	return s
}

func (s *ListFlowNodeInstanceContainerStatusResponseBodyContainerStatusListContainerStatus) SetContainerId(v string) *ListFlowNodeInstanceContainerStatusResponseBodyContainerStatusListContainerStatus {
	s.ContainerId = &v
	return s
}

func (s *ListFlowNodeInstanceContainerStatusResponseBodyContainerStatusListContainerStatus) SetApplicationId(v string) *ListFlowNodeInstanceContainerStatusResponseBodyContainerStatusListContainerStatus {
	s.ApplicationId = &v
	return s
}

type ListFlowNodeInstanceContainerStatusResponse struct {
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFlowNodeInstanceContainerStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowNodeInstanceContainerStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowNodeInstanceContainerStatusResponse) GoString() string {
	return s.String()
}

func (s *ListFlowNodeInstanceContainerStatusResponse) SetHeaders(v map[string]*string) *ListFlowNodeInstanceContainerStatusResponse {
	s.Headers = v
	return s
}

func (s *ListFlowNodeInstanceContainerStatusResponse) SetBody(v *ListFlowNodeInstanceContainerStatusResponseBody) *ListFlowNodeInstanceContainerStatusResponse {
	s.Body = v
	return s
}

type ModifyClusterTemplateRequest struct {
	ResourceOwnerId        *int64                                         `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	BizId                  *string                                        `json:"BizId,omitempty" xml:"BizId,omitempty"`
	TemplateName           *string                                        `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	RegionId               *string                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId                 *string                                        `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	LogPath                *string                                        `json:"LogPath,omitempty" xml:"LogPath,omitempty"`
	SecurityGroupId        *string                                        `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	IsOpenPublicIp         *bool                                          `json:"IsOpenPublicIp,omitempty" xml:"IsOpenPublicIp,omitempty"`
	SecurityGroupName      *string                                        `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
	ChargeType             *string                                        `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Period                 *int32                                         `json:"Period,omitempty" xml:"Period,omitempty"`
	RenewAuto              *bool                                          `json:"RenewAuto,omitempty" xml:"RenewAuto,omitempty"`
	VpcId                  *string                                        `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VSwitchId              *string                                        `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	NetType                *string                                        `json:"NetType,omitempty" xml:"NetType,omitempty"`
	UserDefinedEmrEcsRole  *string                                        `json:"UserDefinedEmrEcsRole,omitempty" xml:"UserDefinedEmrEcsRole,omitempty"`
	EmrVer                 *string                                        `json:"EmrVer,omitempty" xml:"EmrVer,omitempty"`
	ClusterType            *string                                        `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	EnableHighAvailability *bool                                          `json:"EnableHighAvailability,omitempty" xml:"EnableHighAvailability,omitempty"`
	UseLocalMetaDb         *bool                                          `json:"UseLocalMetaDb,omitempty" xml:"UseLocalMetaDb,omitempty"`
	IoOptimizedOption      *bool                                          `json:"IoOptimizedOption,omitempty" xml:"IoOptimizedOption,omitempty"`
	EnableSsh              *bool                                          `json:"EnableSsh,omitempty" xml:"EnableSsh,omitempty"`
	InstanceGeneration     *string                                        `json:"InstanceGeneration,omitempty" xml:"InstanceGeneration,omitempty"`
	MasterPwd              *string                                        `json:"MasterPwd,omitempty" xml:"MasterPwd,omitempty"`
	KeyPairName            *string                                        `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	MetaStoreType          *string                                        `json:"MetaStoreType,omitempty" xml:"MetaStoreType,omitempty"`
	MetaStoreConf          *string                                        `json:"MetaStoreConf,omitempty" xml:"MetaStoreConf,omitempty"`
	Configurations         *string                                        `json:"Configurations,omitempty" xml:"Configurations,omitempty"`
	EnableEas              *bool                                          `json:"EnableEas,omitempty" xml:"EnableEas,omitempty"`
	DepositType            *string                                        `json:"DepositType,omitempty" xml:"DepositType,omitempty"`
	MachineType            *string                                        `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	UseCustomHiveMetaDb    *bool                                          `json:"UseCustomHiveMetaDb,omitempty" xml:"UseCustomHiveMetaDb,omitempty"`
	InitCustomHiveMetaDb   *bool                                          `json:"InitCustomHiveMetaDb,omitempty" xml:"InitCustomHiveMetaDb,omitempty"`
	ResourceGroupId        *string                                        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	OptionSoftWareList     []*string                                      `json:"OptionSoftWareList,omitempty" xml:"OptionSoftWareList,omitempty" type:"Repeated"`
	HostGroup              []*ModifyClusterTemplateRequestHostGroup       `json:"HostGroup,omitempty" xml:"HostGroup,omitempty" type:"Repeated"`
	BootstrapAction        []*ModifyClusterTemplateRequestBootstrapAction `json:"BootstrapAction,omitempty" xml:"BootstrapAction,omitempty" type:"Repeated"`
	Config                 []*ModifyClusterTemplateRequestConfig          `json:"Config,omitempty" xml:"Config,omitempty" type:"Repeated"`
	Tag                    []*ModifyClusterTemplateRequestTag             `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ModifyClusterTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterTemplateRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterTemplateRequest) SetResourceOwnerId(v int64) *ModifyClusterTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyClusterTemplateRequest) SetBizId(v string) *ModifyClusterTemplateRequest {
	s.BizId = &v
	return s
}

func (s *ModifyClusterTemplateRequest) SetTemplateName(v string) *ModifyClusterTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *ModifyClusterTemplateRequest) SetRegionId(v string) *ModifyClusterTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyClusterTemplateRequest) SetZoneId(v string) *ModifyClusterTemplateRequest {
	s.ZoneId = &v
	return s
}

func (s *ModifyClusterTemplateRequest) SetLogPath(v string) *ModifyClusterTemplateRequest {
	s.LogPath = &v
	return s
}

func (s *ModifyClusterTemplateRequest) SetSecurityGroupId(v string) *ModifyClusterTemplateRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *ModifyClusterTemplateRequest) SetIsOpenPublicIp(v bool) *ModifyClusterTemplateRequest {
	s.IsOpenPublicIp = &v
	return s
}

func (s *ModifyClusterTemplateRequest) SetSecurityGroupName(v string) *ModifyClusterTemplateRequest {
	s.SecurityGroupName = &v
	return s
}

func (s *ModifyClusterTemplateRequest) SetChargeType(v string) *ModifyClusterTemplateRequest {
	s.ChargeType = &v
	return s
}

func (s *ModifyClusterTemplateRequest) SetPeriod(v int32) *ModifyClusterTemplateRequest {
	s.Period = &v
	return s
}

func (s *ModifyClusterTemplateRequest) SetRenewAuto(v bool) *ModifyClusterTemplateRequest {
	s.RenewAuto = &v
	return s
}

func (s *ModifyClusterTemplateRequest) SetVpcId(v string) *ModifyClusterTemplateRequest {
	s.VpcId = &v
	return s
}

func (s *ModifyClusterTemplateRequest) SetVSwitchId(v string) *ModifyClusterTemplateRequest {
	s.VSwitchId = &v
	return s
}

func (s *ModifyClusterTemplateRequest) SetNetType(v string) *ModifyClusterTemplateRequest {
	s.NetType = &v
	return s
}

func (s *ModifyClusterTemplateRequest) SetUserDefinedEmrEcsRole(v string) *ModifyClusterTemplateRequest {
	s.UserDefinedEmrEcsRole = &v
	return s
}

func (s *ModifyClusterTemplateRequest) SetEmrVer(v string) *ModifyClusterTemplateRequest {
	s.EmrVer = &v
	return s
}

func (s *ModifyClusterTemplateRequest) SetClusterType(v string) *ModifyClusterTemplateRequest {
	s.ClusterType = &v
	return s
}

func (s *ModifyClusterTemplateRequest) SetEnableHighAvailability(v bool) *ModifyClusterTemplateRequest {
	s.EnableHighAvailability = &v
	return s
}

func (s *ModifyClusterTemplateRequest) SetUseLocalMetaDb(v bool) *ModifyClusterTemplateRequest {
	s.UseLocalMetaDb = &v
	return s
}

func (s *ModifyClusterTemplateRequest) SetIoOptimizedOption(v bool) *ModifyClusterTemplateRequest {
	s.IoOptimizedOption = &v
	return s
}

func (s *ModifyClusterTemplateRequest) SetEnableSsh(v bool) *ModifyClusterTemplateRequest {
	s.EnableSsh = &v
	return s
}

func (s *ModifyClusterTemplateRequest) SetInstanceGeneration(v string) *ModifyClusterTemplateRequest {
	s.InstanceGeneration = &v
	return s
}

func (s *ModifyClusterTemplateRequest) SetMasterPwd(v string) *ModifyClusterTemplateRequest {
	s.MasterPwd = &v
	return s
}

func (s *ModifyClusterTemplateRequest) SetKeyPairName(v string) *ModifyClusterTemplateRequest {
	s.KeyPairName = &v
	return s
}

func (s *ModifyClusterTemplateRequest) SetMetaStoreType(v string) *ModifyClusterTemplateRequest {
	s.MetaStoreType = &v
	return s
}

func (s *ModifyClusterTemplateRequest) SetMetaStoreConf(v string) *ModifyClusterTemplateRequest {
	s.MetaStoreConf = &v
	return s
}

func (s *ModifyClusterTemplateRequest) SetConfigurations(v string) *ModifyClusterTemplateRequest {
	s.Configurations = &v
	return s
}

func (s *ModifyClusterTemplateRequest) SetEnableEas(v bool) *ModifyClusterTemplateRequest {
	s.EnableEas = &v
	return s
}

func (s *ModifyClusterTemplateRequest) SetDepositType(v string) *ModifyClusterTemplateRequest {
	s.DepositType = &v
	return s
}

func (s *ModifyClusterTemplateRequest) SetMachineType(v string) *ModifyClusterTemplateRequest {
	s.MachineType = &v
	return s
}

func (s *ModifyClusterTemplateRequest) SetUseCustomHiveMetaDb(v bool) *ModifyClusterTemplateRequest {
	s.UseCustomHiveMetaDb = &v
	return s
}

func (s *ModifyClusterTemplateRequest) SetInitCustomHiveMetaDb(v bool) *ModifyClusterTemplateRequest {
	s.InitCustomHiveMetaDb = &v
	return s
}

func (s *ModifyClusterTemplateRequest) SetResourceGroupId(v string) *ModifyClusterTemplateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyClusterTemplateRequest) SetOptionSoftWareList(v []*string) *ModifyClusterTemplateRequest {
	s.OptionSoftWareList = v
	return s
}

func (s *ModifyClusterTemplateRequest) SetHostGroup(v []*ModifyClusterTemplateRequestHostGroup) *ModifyClusterTemplateRequest {
	s.HostGroup = v
	return s
}

func (s *ModifyClusterTemplateRequest) SetBootstrapAction(v []*ModifyClusterTemplateRequestBootstrapAction) *ModifyClusterTemplateRequest {
	s.BootstrapAction = v
	return s
}

func (s *ModifyClusterTemplateRequest) SetConfig(v []*ModifyClusterTemplateRequestConfig) *ModifyClusterTemplateRequest {
	s.Config = v
	return s
}

func (s *ModifyClusterTemplateRequest) SetTag(v []*ModifyClusterTemplateRequestTag) *ModifyClusterTemplateRequest {
	s.Tag = v
	return s
}

type ModifyClusterTemplateRequestHostGroup struct {
	SysDiskCapacity    *int32  `json:"SysDiskCapacity,omitempty" xml:"SysDiskCapacity,omitempty"`
	HostGroupType      *string `json:"HostGroupType,omitempty" xml:"HostGroupType,omitempty"`
	Comment            *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	MultiInstanceTypes *string `json:"MultiInstanceTypes,omitempty" xml:"MultiInstanceTypes,omitempty"`
	SysDiskType        *string `json:"SysDiskType,omitempty" xml:"SysDiskType,omitempty"`
	ChargeType         *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	DiskType           *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	HostGroupId        *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	InstanceType       *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	DiskCount          *int32  `json:"DiskCount,omitempty" xml:"DiskCount,omitempty"`
	CreateType         *string `json:"CreateType,omitempty" xml:"CreateType,omitempty"`
	Period             *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	DiskCapacity       *int32  `json:"DiskCapacity,omitempty" xml:"DiskCapacity,omitempty"`
	VSwitchId          *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	NodeCount          *int32  `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	HostGroupName      *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
	RenewAuto          *bool   `json:"RenewAuto,omitempty" xml:"RenewAuto,omitempty"`
	ClusterId          *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s ModifyClusterTemplateRequestHostGroup) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterTemplateRequestHostGroup) GoString() string {
	return s.String()
}

func (s *ModifyClusterTemplateRequestHostGroup) SetSysDiskCapacity(v int32) *ModifyClusterTemplateRequestHostGroup {
	s.SysDiskCapacity = &v
	return s
}

func (s *ModifyClusterTemplateRequestHostGroup) SetHostGroupType(v string) *ModifyClusterTemplateRequestHostGroup {
	s.HostGroupType = &v
	return s
}

func (s *ModifyClusterTemplateRequestHostGroup) SetComment(v string) *ModifyClusterTemplateRequestHostGroup {
	s.Comment = &v
	return s
}

func (s *ModifyClusterTemplateRequestHostGroup) SetMultiInstanceTypes(v string) *ModifyClusterTemplateRequestHostGroup {
	s.MultiInstanceTypes = &v
	return s
}

func (s *ModifyClusterTemplateRequestHostGroup) SetSysDiskType(v string) *ModifyClusterTemplateRequestHostGroup {
	s.SysDiskType = &v
	return s
}

func (s *ModifyClusterTemplateRequestHostGroup) SetChargeType(v string) *ModifyClusterTemplateRequestHostGroup {
	s.ChargeType = &v
	return s
}

func (s *ModifyClusterTemplateRequestHostGroup) SetDiskType(v string) *ModifyClusterTemplateRequestHostGroup {
	s.DiskType = &v
	return s
}

func (s *ModifyClusterTemplateRequestHostGroup) SetHostGroupId(v string) *ModifyClusterTemplateRequestHostGroup {
	s.HostGroupId = &v
	return s
}

func (s *ModifyClusterTemplateRequestHostGroup) SetInstanceType(v string) *ModifyClusterTemplateRequestHostGroup {
	s.InstanceType = &v
	return s
}

func (s *ModifyClusterTemplateRequestHostGroup) SetDiskCount(v int32) *ModifyClusterTemplateRequestHostGroup {
	s.DiskCount = &v
	return s
}

func (s *ModifyClusterTemplateRequestHostGroup) SetCreateType(v string) *ModifyClusterTemplateRequestHostGroup {
	s.CreateType = &v
	return s
}

func (s *ModifyClusterTemplateRequestHostGroup) SetPeriod(v int32) *ModifyClusterTemplateRequestHostGroup {
	s.Period = &v
	return s
}

func (s *ModifyClusterTemplateRequestHostGroup) SetDiskCapacity(v int32) *ModifyClusterTemplateRequestHostGroup {
	s.DiskCapacity = &v
	return s
}

func (s *ModifyClusterTemplateRequestHostGroup) SetVSwitchId(v string) *ModifyClusterTemplateRequestHostGroup {
	s.VSwitchId = &v
	return s
}

func (s *ModifyClusterTemplateRequestHostGroup) SetNodeCount(v int32) *ModifyClusterTemplateRequestHostGroup {
	s.NodeCount = &v
	return s
}

func (s *ModifyClusterTemplateRequestHostGroup) SetHostGroupName(v string) *ModifyClusterTemplateRequestHostGroup {
	s.HostGroupName = &v
	return s
}

func (s *ModifyClusterTemplateRequestHostGroup) SetRenewAuto(v bool) *ModifyClusterTemplateRequestHostGroup {
	s.RenewAuto = &v
	return s
}

func (s *ModifyClusterTemplateRequestHostGroup) SetClusterId(v string) *ModifyClusterTemplateRequestHostGroup {
	s.ClusterId = &v
	return s
}

type ModifyClusterTemplateRequestBootstrapAction struct {
	Arg  *string `json:"Arg,omitempty" xml:"Arg,omitempty"`
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ModifyClusterTemplateRequestBootstrapAction) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterTemplateRequestBootstrapAction) GoString() string {
	return s.String()
}

func (s *ModifyClusterTemplateRequestBootstrapAction) SetArg(v string) *ModifyClusterTemplateRequestBootstrapAction {
	s.Arg = &v
	return s
}

func (s *ModifyClusterTemplateRequestBootstrapAction) SetPath(v string) *ModifyClusterTemplateRequestBootstrapAction {
	s.Path = &v
	return s
}

func (s *ModifyClusterTemplateRequestBootstrapAction) SetName(v string) *ModifyClusterTemplateRequestBootstrapAction {
	s.Name = &v
	return s
}

type ModifyClusterTemplateRequestConfig struct {
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	Replace     *string `json:"Replace,omitempty" xml:"Replace,omitempty"`
	FileName    *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ConfigKey   *string `json:"ConfigKey,omitempty" xml:"ConfigKey,omitempty"`
	Encrypt     *string `json:"Encrypt,omitempty" xml:"Encrypt,omitempty"`
}

func (s ModifyClusterTemplateRequestConfig) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterTemplateRequestConfig) GoString() string {
	return s.String()
}

func (s *ModifyClusterTemplateRequestConfig) SetConfigValue(v string) *ModifyClusterTemplateRequestConfig {
	s.ConfigValue = &v
	return s
}

func (s *ModifyClusterTemplateRequestConfig) SetReplace(v string) *ModifyClusterTemplateRequestConfig {
	s.Replace = &v
	return s
}

func (s *ModifyClusterTemplateRequestConfig) SetFileName(v string) *ModifyClusterTemplateRequestConfig {
	s.FileName = &v
	return s
}

func (s *ModifyClusterTemplateRequestConfig) SetServiceName(v string) *ModifyClusterTemplateRequestConfig {
	s.ServiceName = &v
	return s
}

func (s *ModifyClusterTemplateRequestConfig) SetConfigKey(v string) *ModifyClusterTemplateRequestConfig {
	s.ConfigKey = &v
	return s
}

func (s *ModifyClusterTemplateRequestConfig) SetEncrypt(v string) *ModifyClusterTemplateRequestConfig {
	s.Encrypt = &v
	return s
}

type ModifyClusterTemplateRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyClusterTemplateRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterTemplateRequestTag) GoString() string {
	return s.String()
}

func (s *ModifyClusterTemplateRequestTag) SetKey(v string) *ModifyClusterTemplateRequestTag {
	s.Key = &v
	return s
}

func (s *ModifyClusterTemplateRequestTag) SetValue(v string) *ModifyClusterTemplateRequestTag {
	s.Value = &v
	return s
}

type ModifyClusterTemplateResponseBody struct {
	ClusterTemplateId *string `json:"ClusterTemplateId,omitempty" xml:"ClusterTemplateId,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyClusterTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterTemplateResponseBody) SetClusterTemplateId(v string) *ModifyClusterTemplateResponseBody {
	s.ClusterTemplateId = &v
	return s
}

func (s *ModifyClusterTemplateResponseBody) SetRequestId(v string) *ModifyClusterTemplateResponseBody {
	s.RequestId = &v
	return s
}

type ModifyClusterTemplateResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyClusterTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyClusterTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterTemplateResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterTemplateResponse) SetHeaders(v map[string]*string) *ModifyClusterTemplateResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterTemplateResponse) SetBody(v *ModifyClusterTemplateResponseBody) *ModifyClusterTemplateResponse {
	s.Body = v
	return s
}

type AddSecurityWhiteListRequest struct {
	ClusterId   *string   `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PortRange   *string   `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	WhiteIpList []*string `json:"WhiteIpList,omitempty" xml:"WhiteIpList,omitempty" type:"Repeated"`
	Description *string   `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s AddSecurityWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s AddSecurityWhiteListRequest) GoString() string {
	return s.String()
}

func (s *AddSecurityWhiteListRequest) SetClusterId(v string) *AddSecurityWhiteListRequest {
	s.ClusterId = &v
	return s
}

func (s *AddSecurityWhiteListRequest) SetPortRange(v string) *AddSecurityWhiteListRequest {
	s.PortRange = &v
	return s
}

func (s *AddSecurityWhiteListRequest) SetWhiteIpList(v []*string) *AddSecurityWhiteListRequest {
	s.WhiteIpList = v
	return s
}

func (s *AddSecurityWhiteListRequest) SetDescription(v string) *AddSecurityWhiteListRequest {
	s.Description = &v
	return s
}

type AddSecurityWhiteListResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddSecurityWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddSecurityWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *AddSecurityWhiteListResponseBody) SetRequestId(v string) *AddSecurityWhiteListResponseBody {
	s.RequestId = &v
	return s
}

type AddSecurityWhiteListResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddSecurityWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddSecurityWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s AddSecurityWhiteListResponse) GoString() string {
	return s.String()
}

func (s *AddSecurityWhiteListResponse) SetHeaders(v map[string]*string) *AddSecurityWhiteListResponse {
	s.Headers = v
	return s
}

func (s *AddSecurityWhiteListResponse) SetBody(v *AddSecurityWhiteListResponseBody) *AddSecurityWhiteListResponse {
	s.Body = v
	return s
}

type ListMetaClusterRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SourceType      *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s ListMetaClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMetaClusterRequest) GoString() string {
	return s.String()
}

func (s *ListMetaClusterRequest) SetResourceOwnerId(v int64) *ListMetaClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListMetaClusterRequest) SetRegionId(v string) *ListMetaClusterRequest {
	s.RegionId = &v
	return s
}

func (s *ListMetaClusterRequest) SetPageNumber(v int32) *ListMetaClusterRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMetaClusterRequest) SetPageSize(v int32) *ListMetaClusterRequest {
	s.PageSize = &v
	return s
}

func (s *ListMetaClusterRequest) SetSourceType(v string) *ListMetaClusterRequest {
	s.SourceType = &v
	return s
}

type ListMetaClusterResponseBody struct {
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Items      *ListMetaClusterResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s ListMetaClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMetaClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ListMetaClusterResponseBody) SetRequestId(v string) *ListMetaClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMetaClusterResponseBody) SetPageNumber(v int32) *ListMetaClusterResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListMetaClusterResponseBody) SetPageSize(v int32) *ListMetaClusterResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListMetaClusterResponseBody) SetTotalCount(v int32) *ListMetaClusterResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListMetaClusterResponseBody) SetItems(v *ListMetaClusterResponseBodyItems) *ListMetaClusterResponseBody {
	s.Items = v
	return s
}

type ListMetaClusterResponseBodyItems struct {
	Item []*ListMetaClusterResponseBodyItemsItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s ListMetaClusterResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s ListMetaClusterResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListMetaClusterResponseBodyItems) SetItem(v []*ListMetaClusterResponseBodyItemsItem) *ListMetaClusterResponseBodyItems {
	s.Item = v
	return s
}

type ListMetaClusterResponseBodyItemsItem struct {
	Name         *string                                           `json:"Name,omitempty" xml:"Name,omitempty"`
	Id           *string                                           `json:"Id,omitempty" xml:"Id,omitempty"`
	SoftwareInfo *ListMetaClusterResponseBodyItemsItemSoftwareInfo `json:"SoftwareInfo,omitempty" xml:"SoftwareInfo,omitempty" type:"Struct"`
}

func (s ListMetaClusterResponseBodyItemsItem) String() string {
	return tea.Prettify(s)
}

func (s ListMetaClusterResponseBodyItemsItem) GoString() string {
	return s.String()
}

func (s *ListMetaClusterResponseBodyItemsItem) SetName(v string) *ListMetaClusterResponseBodyItemsItem {
	s.Name = &v
	return s
}

func (s *ListMetaClusterResponseBodyItemsItem) SetId(v string) *ListMetaClusterResponseBodyItemsItem {
	s.Id = &v
	return s
}

func (s *ListMetaClusterResponseBodyItemsItem) SetSoftwareInfo(v *ListMetaClusterResponseBodyItemsItemSoftwareInfo) *ListMetaClusterResponseBodyItemsItem {
	s.SoftwareInfo = v
	return s
}

type ListMetaClusterResponseBodyItemsItemSoftwareInfo struct {
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	EmrVer      *string `json:"EmrVer,omitempty" xml:"EmrVer,omitempty"`
}

func (s ListMetaClusterResponseBodyItemsItemSoftwareInfo) String() string {
	return tea.Prettify(s)
}

func (s ListMetaClusterResponseBodyItemsItemSoftwareInfo) GoString() string {
	return s.String()
}

func (s *ListMetaClusterResponseBodyItemsItemSoftwareInfo) SetClusterType(v string) *ListMetaClusterResponseBodyItemsItemSoftwareInfo {
	s.ClusterType = &v
	return s
}

func (s *ListMetaClusterResponseBodyItemsItemSoftwareInfo) SetEmrVer(v string) *ListMetaClusterResponseBodyItemsItemSoftwareInfo {
	s.EmrVer = &v
	return s
}

type ListMetaClusterResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListMetaClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMetaClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMetaClusterResponse) GoString() string {
	return s.String()
}

func (s *ListMetaClusterResponse) SetHeaders(v map[string]*string) *ListMetaClusterResponse {
	s.Headers = v
	return s
}

func (s *ListMetaClusterResponse) SetBody(v *ListMetaClusterResponseBody) *ListMetaClusterResponse {
	s.Body = v
	return s
}

type ListClusterOperationHostRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OperationId     *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListClusterOperationHostRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClusterOperationHostRequest) GoString() string {
	return s.String()
}

func (s *ListClusterOperationHostRequest) SetResourceOwnerId(v int64) *ListClusterOperationHostRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListClusterOperationHostRequest) SetRegionId(v string) *ListClusterOperationHostRequest {
	s.RegionId = &v
	return s
}

func (s *ListClusterOperationHostRequest) SetClusterId(v string) *ListClusterOperationHostRequest {
	s.ClusterId = &v
	return s
}

func (s *ListClusterOperationHostRequest) SetOperationId(v string) *ListClusterOperationHostRequest {
	s.OperationId = &v
	return s
}

func (s *ListClusterOperationHostRequest) SetStatus(v string) *ListClusterOperationHostRequest {
	s.Status = &v
	return s
}

func (s *ListClusterOperationHostRequest) SetPageNumber(v int32) *ListClusterOperationHostRequest {
	s.PageNumber = &v
	return s
}

func (s *ListClusterOperationHostRequest) SetPageSize(v int32) *ListClusterOperationHostRequest {
	s.PageSize = &v
	return s
}

type ListClusterOperationHostResponseBody struct {
	PageSize                 *int32                                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId                *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber               *int32                                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount               *int32                                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	ClusterOperationHostList *ListClusterOperationHostResponseBodyClusterOperationHostList `json:"ClusterOperationHostList,omitempty" xml:"ClusterOperationHostList,omitempty" type:"Struct"`
}

func (s ListClusterOperationHostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClusterOperationHostResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterOperationHostResponseBody) SetPageSize(v int32) *ListClusterOperationHostResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListClusterOperationHostResponseBody) SetRequestId(v string) *ListClusterOperationHostResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterOperationHostResponseBody) SetPageNumber(v int32) *ListClusterOperationHostResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListClusterOperationHostResponseBody) SetTotalCount(v int32) *ListClusterOperationHostResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListClusterOperationHostResponseBody) SetClusterOperationHostList(v *ListClusterOperationHostResponseBodyClusterOperationHostList) *ListClusterOperationHostResponseBody {
	s.ClusterOperationHostList = v
	return s
}

type ListClusterOperationHostResponseBodyClusterOperationHostList struct {
	ClusterOperationHost []*ListClusterOperationHostResponseBodyClusterOperationHostListClusterOperationHost `json:"ClusterOperationHost,omitempty" xml:"ClusterOperationHost,omitempty" type:"Repeated"`
}

func (s ListClusterOperationHostResponseBodyClusterOperationHostList) String() string {
	return tea.Prettify(s)
}

func (s ListClusterOperationHostResponseBodyClusterOperationHostList) GoString() string {
	return s.String()
}

func (s *ListClusterOperationHostResponseBodyClusterOperationHostList) SetClusterOperationHost(v []*ListClusterOperationHostResponseBodyClusterOperationHostListClusterOperationHost) *ListClusterOperationHostResponseBodyClusterOperationHostList {
	s.ClusterOperationHost = v
	return s
}

type ListClusterOperationHostResponseBodyClusterOperationHostListClusterOperationHost struct {
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	HostName   *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	Percentage *string `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	HostId     *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
}

func (s ListClusterOperationHostResponseBodyClusterOperationHostListClusterOperationHost) String() string {
	return tea.Prettify(s)
}

func (s ListClusterOperationHostResponseBodyClusterOperationHostListClusterOperationHost) GoString() string {
	return s.String()
}

func (s *ListClusterOperationHostResponseBodyClusterOperationHostListClusterOperationHost) SetStatus(v string) *ListClusterOperationHostResponseBodyClusterOperationHostListClusterOperationHost {
	s.Status = &v
	return s
}

func (s *ListClusterOperationHostResponseBodyClusterOperationHostListClusterOperationHost) SetHostName(v string) *ListClusterOperationHostResponseBodyClusterOperationHostListClusterOperationHost {
	s.HostName = &v
	return s
}

func (s *ListClusterOperationHostResponseBodyClusterOperationHostListClusterOperationHost) SetPercentage(v string) *ListClusterOperationHostResponseBodyClusterOperationHostListClusterOperationHost {
	s.Percentage = &v
	return s
}

func (s *ListClusterOperationHostResponseBodyClusterOperationHostListClusterOperationHost) SetHostId(v string) *ListClusterOperationHostResponseBodyClusterOperationHostListClusterOperationHost {
	s.HostId = &v
	return s
}

type ListClusterOperationHostResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListClusterOperationHostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClusterOperationHostResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClusterOperationHostResponse) GoString() string {
	return s.String()
}

func (s *ListClusterOperationHostResponse) SetHeaders(v map[string]*string) *ListClusterOperationHostResponse {
	s.Headers = v
	return s
}

func (s *ListClusterOperationHostResponse) SetBody(v *ListClusterOperationHostResponseBody) *ListClusterOperationHostResponse {
	s.Body = v
	return s
}

type ListClusterTemplatesRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	BizId           *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProductType     *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListClusterTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClusterTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListClusterTemplatesRequest) SetResourceOwnerId(v int64) *ListClusterTemplatesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListClusterTemplatesRequest) SetRegionId(v string) *ListClusterTemplatesRequest {
	s.RegionId = &v
	return s
}

func (s *ListClusterTemplatesRequest) SetBizId(v string) *ListClusterTemplatesRequest {
	s.BizId = &v
	return s
}

func (s *ListClusterTemplatesRequest) SetPageNumber(v int32) *ListClusterTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListClusterTemplatesRequest) SetPageSize(v int32) *ListClusterTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListClusterTemplatesRequest) SetProductType(v string) *ListClusterTemplatesRequest {
	s.ProductType = &v
	return s
}

func (s *ListClusterTemplatesRequest) SetResourceGroupId(v string) *ListClusterTemplatesRequest {
	s.ResourceGroupId = &v
	return s
}

type ListClusterTemplatesResponseBody struct {
	PageSize         *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId        *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber       *int32                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount       *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TemplateInfoList *ListClusterTemplatesResponseBodyTemplateInfoList `json:"TemplateInfoList,omitempty" xml:"TemplateInfoList,omitempty" type:"Struct"`
}

func (s ListClusterTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClusterTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterTemplatesResponseBody) SetPageSize(v int32) *ListClusterTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListClusterTemplatesResponseBody) SetRequestId(v string) *ListClusterTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterTemplatesResponseBody) SetPageNumber(v int32) *ListClusterTemplatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListClusterTemplatesResponseBody) SetTotalCount(v int32) *ListClusterTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListClusterTemplatesResponseBody) SetTemplateInfoList(v *ListClusterTemplatesResponseBodyTemplateInfoList) *ListClusterTemplatesResponseBody {
	s.TemplateInfoList = v
	return s
}

type ListClusterTemplatesResponseBodyTemplateInfoList struct {
	TemplateInfo []*ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo `json:"TemplateInfo,omitempty" xml:"TemplateInfo,omitempty" type:"Repeated"`
}

func (s ListClusterTemplatesResponseBodyTemplateInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListClusterTemplatesResponseBodyTemplateInfoList) GoString() string {
	return s.String()
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoList) SetTemplateInfo(v []*ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo) *ListClusterTemplatesResponseBodyTemplateInfoList {
	s.TemplateInfo = v
	return s
}

type ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo struct {
	VpcId                  *string                                                                          `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	KeyPairName            *string                                                                          `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	LogEnable              *bool                                                                            `json:"LogEnable,omitempty" xml:"LogEnable,omitempty"`
	SshEnable              *bool                                                                            `json:"SshEnable,omitempty" xml:"SshEnable,omitempty"`
	HighAvailabilityEnable *bool                                                                            `json:"HighAvailabilityEnable,omitempty" xml:"HighAvailabilityEnable,omitempty"`
	SecurityGroupId        *string                                                                          `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	UserId                 *string                                                                          `json:"UserId,omitempty" xml:"UserId,omitempty"`
	IsOpenPublicIp         *bool                                                                            `json:"IsOpenPublicIp,omitempty" xml:"IsOpenPublicIp,omitempty"`
	AllowNotebook          *bool                                                                            `json:"AllowNotebook,omitempty" xml:"AllowNotebook,omitempty"`
	GmtModified            *int64                                                                           `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	TemplateName           *string                                                                          `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	DepositType            *string                                                                          `json:"DepositType,omitempty" xml:"DepositType,omitempty"`
	SecurityGroupName      *string                                                                          `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
	CreateSource           *string                                                                          `json:"CreateSource,omitempty" xml:"CreateSource,omitempty"`
	InstanceGeneration     *string                                                                          `json:"InstanceGeneration,omitempty" xml:"InstanceGeneration,omitempty"`
	UseCustomHiveMetaDb    *bool                                                                            `json:"UseCustomHiveMetaDb,omitempty" xml:"UseCustomHiveMetaDb,omitempty"`
	EasEnable              *bool                                                                            `json:"EasEnable,omitempty" xml:"EasEnable,omitempty"`
	UserDefinedEmrEcsRole  *string                                                                          `json:"UserDefinedEmrEcsRole,omitempty" xml:"UserDefinedEmrEcsRole,omitempty"`
	MetaStoreType          *string                                                                          `json:"MetaStoreType,omitempty" xml:"MetaStoreType,omitempty"`
	MachineType            *string                                                                          `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	UseLocalMetaDb         *bool                                                                            `json:"UseLocalMetaDb,omitempty" xml:"UseLocalMetaDb,omitempty"`
	MasterNodeTotal        *int32                                                                           `json:"MasterNodeTotal,omitempty" xml:"MasterNodeTotal,omitempty"`
	InitCustomHiveMetaDb   *bool                                                                            `json:"InitCustomHiveMetaDb,omitempty" xml:"InitCustomHiveMetaDb,omitempty"`
	IoOptimized            *bool                                                                            `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	MetaStoreConf          *string                                                                          `json:"MetaStoreConf,omitempty" xml:"MetaStoreConf,omitempty"`
	VSwitchId              *string                                                                          `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	Configurations         *string                                                                          `json:"Configurations,omitempty" xml:"Configurations,omitempty"`
	LogPath                *string                                                                          `json:"LogPath,omitempty" xml:"LogPath,omitempty"`
	ClusterType            *string                                                                          `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	NetType                *string                                                                          `json:"NetType,omitempty" xml:"NetType,omitempty"`
	ZoneId                 *string                                                                          `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	GmtCreate              *int64                                                                           `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id                     *string                                                                          `json:"Id,omitempty" xml:"Id,omitempty"`
	BootstrapActionList    *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoBootstrapActionList `json:"BootstrapActionList,omitempty" xml:"BootstrapActionList,omitempty" type:"Struct"`
	HostGroupList          *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoHostGroupList       `json:"HostGroupList,omitempty" xml:"HostGroupList,omitempty" type:"Struct"`
	ConfigList             *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoConfigList          `json:"ConfigList,omitempty" xml:"ConfigList,omitempty" type:"Struct"`
	SoftwareInfoList       *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoSoftwareInfoList    `json:"SoftwareInfoList,omitempty" xml:"SoftwareInfoList,omitempty" type:"Struct"`
}

func (s ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo) String() string {
	return tea.Prettify(s)
}

func (s ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo) GoString() string {
	return s.String()
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo) SetVpcId(v string) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo {
	s.VpcId = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo) SetKeyPairName(v string) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo {
	s.KeyPairName = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo) SetLogEnable(v bool) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo {
	s.LogEnable = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo) SetSshEnable(v bool) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo {
	s.SshEnable = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo) SetHighAvailabilityEnable(v bool) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo {
	s.HighAvailabilityEnable = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo) SetSecurityGroupId(v string) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo {
	s.SecurityGroupId = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo) SetUserId(v string) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo {
	s.UserId = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo) SetIsOpenPublicIp(v bool) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo {
	s.IsOpenPublicIp = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo) SetAllowNotebook(v bool) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo {
	s.AllowNotebook = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo) SetGmtModified(v int64) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo {
	s.GmtModified = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo) SetTemplateName(v string) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo {
	s.TemplateName = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo) SetDepositType(v string) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo {
	s.DepositType = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo) SetSecurityGroupName(v string) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo {
	s.SecurityGroupName = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo) SetCreateSource(v string) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo {
	s.CreateSource = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo) SetInstanceGeneration(v string) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo {
	s.InstanceGeneration = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo) SetUseCustomHiveMetaDb(v bool) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo {
	s.UseCustomHiveMetaDb = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo) SetEasEnable(v bool) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo {
	s.EasEnable = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo) SetUserDefinedEmrEcsRole(v string) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo {
	s.UserDefinedEmrEcsRole = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo) SetMetaStoreType(v string) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo {
	s.MetaStoreType = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo) SetMachineType(v string) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo {
	s.MachineType = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo) SetUseLocalMetaDb(v bool) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo {
	s.UseLocalMetaDb = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo) SetMasterNodeTotal(v int32) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo {
	s.MasterNodeTotal = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo) SetInitCustomHiveMetaDb(v bool) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo {
	s.InitCustomHiveMetaDb = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo) SetIoOptimized(v bool) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo {
	s.IoOptimized = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo) SetMetaStoreConf(v string) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo {
	s.MetaStoreConf = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo) SetVSwitchId(v string) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo {
	s.VSwitchId = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo) SetConfigurations(v string) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo {
	s.Configurations = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo) SetLogPath(v string) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo {
	s.LogPath = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo) SetClusterType(v string) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo {
	s.ClusterType = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo) SetNetType(v string) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo {
	s.NetType = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo) SetZoneId(v string) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo {
	s.ZoneId = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo) SetGmtCreate(v int64) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo {
	s.GmtCreate = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo) SetId(v string) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo {
	s.Id = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo) SetBootstrapActionList(v *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoBootstrapActionList) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo {
	s.BootstrapActionList = v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo) SetHostGroupList(v *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoHostGroupList) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo {
	s.HostGroupList = v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo) SetConfigList(v *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoConfigList) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo {
	s.ConfigList = v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo) SetSoftwareInfoList(v *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoSoftwareInfoList) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfo {
	s.SoftwareInfoList = v
	return s
}

type ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoBootstrapActionList struct {
	BootstrapAction []*ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoBootstrapActionListBootstrapAction `json:"BootstrapAction,omitempty" xml:"BootstrapAction,omitempty" type:"Repeated"`
}

func (s ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoBootstrapActionList) String() string {
	return tea.Prettify(s)
}

func (s ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoBootstrapActionList) GoString() string {
	return s.String()
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoBootstrapActionList) SetBootstrapAction(v []*ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoBootstrapActionListBootstrapAction) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoBootstrapActionList {
	s.BootstrapAction = v
	return s
}

type ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoBootstrapActionListBootstrapAction struct {
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Arg  *string `json:"Arg,omitempty" xml:"Arg,omitempty"`
}

func (s ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoBootstrapActionListBootstrapAction) String() string {
	return tea.Prettify(s)
}

func (s ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoBootstrapActionListBootstrapAction) GoString() string {
	return s.String()
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoBootstrapActionListBootstrapAction) SetPath(v string) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoBootstrapActionListBootstrapAction {
	s.Path = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoBootstrapActionListBootstrapAction) SetName(v string) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoBootstrapActionListBootstrapAction {
	s.Name = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoBootstrapActionListBootstrapAction) SetArg(v string) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoBootstrapActionListBootstrapAction {
	s.Arg = &v
	return s
}

type ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoHostGroupList struct {
	HostGroup []*ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoHostGroupListHostGroup `json:"HostGroup,omitempty" xml:"HostGroup,omitempty" type:"Repeated"`
}

func (s ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoHostGroupList) String() string {
	return tea.Prettify(s)
}

func (s ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoHostGroupList) GoString() string {
	return s.String()
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoHostGroupList) SetHostGroup(v []*ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoHostGroupListHostGroup) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoHostGroupList {
	s.HostGroup = v
	return s
}

type ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoHostGroupListHostGroup struct {
	SysDiskCapacity    *int32  `json:"SysDiskCapacity,omitempty" xml:"SysDiskCapacity,omitempty"`
	HostGroupType      *string `json:"HostGroupType,omitempty" xml:"HostGroupType,omitempty"`
	MultiInstanceTypes *string `json:"MultiInstanceTypes,omitempty" xml:"MultiInstanceTypes,omitempty"`
	SysDiskType        *string `json:"SysDiskType,omitempty" xml:"SysDiskType,omitempty"`
	ChargeType         *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	DiskType           *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	HostGroupId        *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	InstanceType       *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	DiskCount          *int32  `json:"DiskCount,omitempty" xml:"DiskCount,omitempty"`
	Period             *string `json:"Period,omitempty" xml:"Period,omitempty"`
	DiskCapacity       *int32  `json:"DiskCapacity,omitempty" xml:"DiskCapacity,omitempty"`
	NodeCount          *int32  `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	HostGroupName      *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
}

func (s ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoHostGroupListHostGroup) String() string {
	return tea.Prettify(s)
}

func (s ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoHostGroupListHostGroup) GoString() string {
	return s.String()
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoHostGroupListHostGroup) SetSysDiskCapacity(v int32) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoHostGroupListHostGroup {
	s.SysDiskCapacity = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoHostGroupListHostGroup) SetHostGroupType(v string) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoHostGroupListHostGroup {
	s.HostGroupType = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoHostGroupListHostGroup) SetMultiInstanceTypes(v string) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoHostGroupListHostGroup {
	s.MultiInstanceTypes = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoHostGroupListHostGroup) SetSysDiskType(v string) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoHostGroupListHostGroup {
	s.SysDiskType = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoHostGroupListHostGroup) SetChargeType(v string) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoHostGroupListHostGroup {
	s.ChargeType = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoHostGroupListHostGroup) SetDiskType(v string) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoHostGroupListHostGroup {
	s.DiskType = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoHostGroupListHostGroup) SetHostGroupId(v string) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoHostGroupListHostGroup {
	s.HostGroupId = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoHostGroupListHostGroup) SetInstanceType(v string) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoHostGroupListHostGroup {
	s.InstanceType = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoHostGroupListHostGroup) SetDiskCount(v int32) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoHostGroupListHostGroup {
	s.DiskCount = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoHostGroupListHostGroup) SetPeriod(v string) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoHostGroupListHostGroup {
	s.Period = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoHostGroupListHostGroup) SetDiskCapacity(v int32) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoHostGroupListHostGroup {
	s.DiskCapacity = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoHostGroupListHostGroup) SetNodeCount(v int32) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoHostGroupListHostGroup {
	s.NodeCount = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoHostGroupListHostGroup) SetHostGroupName(v string) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoHostGroupListHostGroup {
	s.HostGroupName = &v
	return s
}

type ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoConfigList struct {
	Config []*ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoConfigListConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Repeated"`
}

func (s ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoConfigList) String() string {
	return tea.Prettify(s)
}

func (s ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoConfigList) GoString() string {
	return s.String()
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoConfigList) SetConfig(v []*ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoConfigListConfig) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoConfigList {
	s.Config = v
	return s
}

type ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoConfigListConfig struct {
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	Replace     *string `json:"Replace,omitempty" xml:"Replace,omitempty"`
	FileName    *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ConfigKey   *string `json:"ConfigKey,omitempty" xml:"ConfigKey,omitempty"`
	Encrypt     *string `json:"Encrypt,omitempty" xml:"Encrypt,omitempty"`
}

func (s ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoConfigListConfig) String() string {
	return tea.Prettify(s)
}

func (s ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoConfigListConfig) GoString() string {
	return s.String()
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoConfigListConfig) SetConfigValue(v string) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoConfigListConfig {
	s.ConfigValue = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoConfigListConfig) SetReplace(v string) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoConfigListConfig {
	s.Replace = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoConfigListConfig) SetFileName(v string) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoConfigListConfig {
	s.FileName = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoConfigListConfig) SetServiceName(v string) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoConfigListConfig {
	s.ServiceName = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoConfigListConfig) SetConfigKey(v string) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoConfigListConfig {
	s.ConfigKey = &v
	return s
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoConfigListConfig) SetEncrypt(v string) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoConfigListConfig {
	s.Encrypt = &v
	return s
}

type ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoSoftwareInfoList struct {
	SoftwareInfo []*string `json:"SoftwareInfo,omitempty" xml:"SoftwareInfo,omitempty" type:"Repeated"`
}

func (s ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoSoftwareInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoSoftwareInfoList) GoString() string {
	return s.String()
}

func (s *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoSoftwareInfoList) SetSoftwareInfo(v []*string) *ListClusterTemplatesResponseBodyTemplateInfoListTemplateInfoSoftwareInfoList {
	s.SoftwareInfo = v
	return s
}

type ListClusterTemplatesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListClusterTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClusterTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClusterTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListClusterTemplatesResponse) SetHeaders(v map[string]*string) *ListClusterTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListClusterTemplatesResponse) SetBody(v *ListClusterTemplatesResponseBody) *ListClusterTemplatesResponse {
	s.Body = v
	return s
}

type ListClustersRequest struct {
	ResourceOwnerId *int64                    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	CreateType      *string                   `json:"CreateType,omitempty" xml:"CreateType,omitempty"`
	MachineType     *string                   `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	IsDesc          *bool                     `json:"IsDesc,omitempty" xml:"IsDesc,omitempty"`
	DepositType     *string                   `json:"DepositType,omitempty" xml:"DepositType,omitempty"`
	PageNumber      *int32                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	DefaultStatus   *bool                     `json:"DefaultStatus,omitempty" xml:"DefaultStatus,omitempty"`
	Name            *string                   `json:"Name,omitempty" xml:"Name,omitempty"`
	ResourceGroupId *string                   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ClusterTypeList []*string                 `json:"ClusterTypeList,omitempty" xml:"ClusterTypeList,omitempty" type:"Repeated"`
	StatusList      []*string                 `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
	Tag             []*ListClustersRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	ExpiredTagList  []*string                 `json:"ExpiredTagList,omitempty" xml:"ExpiredTagList,omitempty" type:"Repeated"`
}

func (s ListClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClustersRequest) GoString() string {
	return s.String()
}

func (s *ListClustersRequest) SetResourceOwnerId(v int64) *ListClustersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListClustersRequest) SetRegionId(v string) *ListClustersRequest {
	s.RegionId = &v
	return s
}

func (s *ListClustersRequest) SetCreateType(v string) *ListClustersRequest {
	s.CreateType = &v
	return s
}

func (s *ListClustersRequest) SetMachineType(v string) *ListClustersRequest {
	s.MachineType = &v
	return s
}

func (s *ListClustersRequest) SetIsDesc(v bool) *ListClustersRequest {
	s.IsDesc = &v
	return s
}

func (s *ListClustersRequest) SetDepositType(v string) *ListClustersRequest {
	s.DepositType = &v
	return s
}

func (s *ListClustersRequest) SetPageNumber(v int32) *ListClustersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListClustersRequest) SetPageSize(v int32) *ListClustersRequest {
	s.PageSize = &v
	return s
}

func (s *ListClustersRequest) SetDefaultStatus(v bool) *ListClustersRequest {
	s.DefaultStatus = &v
	return s
}

func (s *ListClustersRequest) SetName(v string) *ListClustersRequest {
	s.Name = &v
	return s
}

func (s *ListClustersRequest) SetResourceGroupId(v string) *ListClustersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListClustersRequest) SetClusterTypeList(v []*string) *ListClustersRequest {
	s.ClusterTypeList = v
	return s
}

func (s *ListClustersRequest) SetStatusList(v []*string) *ListClustersRequest {
	s.StatusList = v
	return s
}

func (s *ListClustersRequest) SetTag(v []*ListClustersRequestTag) *ListClustersRequest {
	s.Tag = v
	return s
}

func (s *ListClustersRequest) SetExpiredTagList(v []*string) *ListClustersRequest {
	s.ExpiredTagList = v
	return s
}

type ListClustersRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListClustersRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListClustersRequestTag) GoString() string {
	return s.String()
}

func (s *ListClustersRequestTag) SetKey(v string) *ListClustersRequestTag {
	s.Key = &v
	return s
}

func (s *ListClustersRequestTag) SetValue(v string) *ListClustersRequestTag {
	s.Value = &v
	return s
}

type ListClustersResponseBody struct {
	PageSize   *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Clusters   *ListClustersResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Struct"`
}

func (s ListClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBody) SetPageSize(v int32) *ListClustersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListClustersResponseBody) SetRequestId(v string) *ListClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClustersResponseBody) SetPageNumber(v int32) *ListClustersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListClustersResponseBody) SetTotalCount(v int32) *ListClustersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListClustersResponseBody) SetClusters(v *ListClustersResponseBodyClusters) *ListClustersResponseBody {
	s.Clusters = v
	return s
}

type ListClustersResponseBodyClusters struct {
	ClusterInfo []*ListClustersResponseBodyClustersClusterInfo `json:"ClusterInfo,omitempty" xml:"ClusterInfo,omitempty" type:"Repeated"`
}

func (s ListClustersResponseBodyClusters) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClusters) SetClusterInfo(v []*ListClustersResponseBodyClustersClusterInfo) *ListClustersResponseBodyClusters {
	s.ClusterInfo = v
	return s
}

type ListClustersResponseBodyClustersClusterInfo struct {
	Type                *string                                                   `json:"Type,omitempty" xml:"Type,omitempty"`
	Status              *string                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	RunningTime         *int32                                                    `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	MetaStoreType       *string                                                   `json:"MetaStoreType,omitempty" xml:"MetaStoreType,omitempty"`
	MachineType         *string                                                   `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	OrderList           *string                                                   `json:"OrderList,omitempty" xml:"OrderList,omitempty"`
	CreateTime          *int64                                                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ChargeType          *string                                                   `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	DepositType         *string                                                   `json:"DepositType,omitempty" xml:"DepositType,omitempty"`
	Period              *int32                                                    `json:"Period,omitempty" xml:"Period,omitempty"`
	K8sClusterId        *string                                                   `json:"K8sClusterId,omitempty" xml:"K8sClusterId,omitempty"`
	ExpiredTime         *int64                                                    `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	HasUncompletedOrder *bool                                                     `json:"HasUncompletedOrder,omitempty" xml:"HasUncompletedOrder,omitempty"`
	Name                *string                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	CreateResource      *string                                                   `json:"CreateResource,omitempty" xml:"CreateResource,omitempty"`
	Id                  *string                                                   `json:"Id,omitempty" xml:"Id,omitempty"`
	Tags                *ListClustersResponseBodyClustersClusterInfoTags          `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	OrderTaskInfo       *ListClustersResponseBodyClustersClusterInfoOrderTaskInfo `json:"OrderTaskInfo,omitempty" xml:"OrderTaskInfo,omitempty" type:"Struct"`
	FailReason          *ListClustersResponseBodyClustersClusterInfoFailReason    `json:"FailReason,omitempty" xml:"FailReason,omitempty" type:"Struct"`
}

func (s ListClustersResponseBodyClustersClusterInfo) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersClusterInfo) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersClusterInfo) SetType(v string) *ListClustersResponseBodyClustersClusterInfo {
	s.Type = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfo) SetStatus(v string) *ListClustersResponseBodyClustersClusterInfo {
	s.Status = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfo) SetRunningTime(v int32) *ListClustersResponseBodyClustersClusterInfo {
	s.RunningTime = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfo) SetMetaStoreType(v string) *ListClustersResponseBodyClustersClusterInfo {
	s.MetaStoreType = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfo) SetMachineType(v string) *ListClustersResponseBodyClustersClusterInfo {
	s.MachineType = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfo) SetOrderList(v string) *ListClustersResponseBodyClustersClusterInfo {
	s.OrderList = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfo) SetCreateTime(v int64) *ListClustersResponseBodyClustersClusterInfo {
	s.CreateTime = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfo) SetChargeType(v string) *ListClustersResponseBodyClustersClusterInfo {
	s.ChargeType = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfo) SetDepositType(v string) *ListClustersResponseBodyClustersClusterInfo {
	s.DepositType = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfo) SetPeriod(v int32) *ListClustersResponseBodyClustersClusterInfo {
	s.Period = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfo) SetK8sClusterId(v string) *ListClustersResponseBodyClustersClusterInfo {
	s.K8sClusterId = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfo) SetExpiredTime(v int64) *ListClustersResponseBodyClustersClusterInfo {
	s.ExpiredTime = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfo) SetHasUncompletedOrder(v bool) *ListClustersResponseBodyClustersClusterInfo {
	s.HasUncompletedOrder = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfo) SetName(v string) *ListClustersResponseBodyClustersClusterInfo {
	s.Name = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfo) SetCreateResource(v string) *ListClustersResponseBodyClustersClusterInfo {
	s.CreateResource = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfo) SetId(v string) *ListClustersResponseBodyClustersClusterInfo {
	s.Id = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfo) SetTags(v *ListClustersResponseBodyClustersClusterInfoTags) *ListClustersResponseBodyClustersClusterInfo {
	s.Tags = v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfo) SetOrderTaskInfo(v *ListClustersResponseBodyClustersClusterInfoOrderTaskInfo) *ListClustersResponseBodyClustersClusterInfo {
	s.OrderTaskInfo = v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfo) SetFailReason(v *ListClustersResponseBodyClustersClusterInfoFailReason) *ListClustersResponseBodyClustersClusterInfo {
	s.FailReason = v
	return s
}

type ListClustersResponseBodyClustersClusterInfoTags struct {
	Tag []*ListClustersResponseBodyClustersClusterInfoTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListClustersResponseBodyClustersClusterInfoTags) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersClusterInfoTags) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersClusterInfoTags) SetTag(v []*ListClustersResponseBodyClustersClusterInfoTagsTag) *ListClustersResponseBodyClustersClusterInfoTags {
	s.Tag = v
	return s
}

type ListClustersResponseBodyClustersClusterInfoTagsTag struct {
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListClustersResponseBodyClustersClusterInfoTagsTag) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersClusterInfoTagsTag) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersClusterInfoTagsTag) SetTagValue(v string) *ListClustersResponseBodyClustersClusterInfoTagsTag {
	s.TagValue = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoTagsTag) SetTagKey(v string) *ListClustersResponseBodyClustersClusterInfoTagsTag {
	s.TagKey = &v
	return s
}

type ListClustersResponseBodyClustersClusterInfoOrderTaskInfo struct {
	TargetCount  *int32  `json:"TargetCount,omitempty" xml:"TargetCount,omitempty"`
	CurrentCount *int32  `json:"CurrentCount,omitempty" xml:"CurrentCount,omitempty"`
	OrderIdList  *string `json:"OrderIdList,omitempty" xml:"OrderIdList,omitempty"`
}

func (s ListClustersResponseBodyClustersClusterInfoOrderTaskInfo) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersClusterInfoOrderTaskInfo) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersClusterInfoOrderTaskInfo) SetTargetCount(v int32) *ListClustersResponseBodyClustersClusterInfoOrderTaskInfo {
	s.TargetCount = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoOrderTaskInfo) SetCurrentCount(v int32) *ListClustersResponseBodyClustersClusterInfoOrderTaskInfo {
	s.CurrentCount = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoOrderTaskInfo) SetOrderIdList(v string) *ListClustersResponseBodyClustersClusterInfoOrderTaskInfo {
	s.OrderIdList = &v
	return s
}

type ListClustersResponseBodyClustersClusterInfoFailReason struct {
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
}

func (s ListClustersResponseBodyClustersClusterInfoFailReason) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersClusterInfoFailReason) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersClusterInfoFailReason) SetErrorMsg(v string) *ListClustersResponseBodyClustersClusterInfoFailReason {
	s.ErrorMsg = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoFailReason) SetRequestId(v string) *ListClustersResponseBodyClustersClusterInfoFailReason {
	s.RequestId = &v
	return s
}

func (s *ListClustersResponseBodyClustersClusterInfoFailReason) SetErrorCode(v string) *ListClustersResponseBodyClustersClusterInfoFailReason {
	s.ErrorCode = &v
	return s
}

type ListClustersResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponse) GoString() string {
	return s.String()
}

func (s *ListClustersResponse) SetHeaders(v map[string]*string) *ListClustersResponse {
	s.Headers = v
	return s
}

func (s *ListClustersResponse) SetBody(v *ListClustersResponseBody) *ListClustersResponse {
	s.Body = v
	return s
}

type TagResourcesSystemTagsRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 资源类型：cluster
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// 主账号UID
	TagOwnerUid *int64  `json:"TagOwnerUid,omitempty" xml:"TagOwnerUid,omitempty"`
	Scope       *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// 资源标签
	Tags []*TagResourcesSystemTagsRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// 资源ID
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
}

func (s TagResourcesSystemTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesSystemTagsRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesSystemTagsRequest) SetResourceOwnerId(v int64) *TagResourcesSystemTagsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TagResourcesSystemTagsRequest) SetRegionId(v string) *TagResourcesSystemTagsRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesSystemTagsRequest) SetResourceType(v string) *TagResourcesSystemTagsRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesSystemTagsRequest) SetTagOwnerUid(v int64) *TagResourcesSystemTagsRequest {
	s.TagOwnerUid = &v
	return s
}

func (s *TagResourcesSystemTagsRequest) SetScope(v string) *TagResourcesSystemTagsRequest {
	s.Scope = &v
	return s
}

func (s *TagResourcesSystemTagsRequest) SetTags(v []*TagResourcesSystemTagsRequestTags) *TagResourcesSystemTagsRequest {
	s.Tags = v
	return s
}

func (s *TagResourcesSystemTagsRequest) SetResourceIds(v []*string) *TagResourcesSystemTagsRequest {
	s.ResourceIds = v
	return s
}

type TagResourcesSystemTagsRequestTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesSystemTagsRequestTags) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesSystemTagsRequestTags) GoString() string {
	return s.String()
}

func (s *TagResourcesSystemTagsRequestTags) SetKey(v string) *TagResourcesSystemTagsRequestTags {
	s.Key = &v
	return s
}

func (s *TagResourcesSystemTagsRequestTags) SetValue(v string) *TagResourcesSystemTagsRequestTags {
	s.Value = &v
	return s
}

type TagResourcesSystemTagsResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 请求是否成功被处理
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 响应码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 错误码
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 响应消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s TagResourcesSystemTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesSystemTagsResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesSystemTagsResponseBody) SetRequestId(v string) *TagResourcesSystemTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *TagResourcesSystemTagsResponseBody) SetSuccess(v bool) *TagResourcesSystemTagsResponseBody {
	s.Success = &v
	return s
}

func (s *TagResourcesSystemTagsResponseBody) SetCode(v string) *TagResourcesSystemTagsResponseBody {
	s.Code = &v
	return s
}

func (s *TagResourcesSystemTagsResponseBody) SetErrorCode(v string) *TagResourcesSystemTagsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TagResourcesSystemTagsResponseBody) SetMessage(v string) *TagResourcesSystemTagsResponseBody {
	s.Message = &v
	return s
}

type TagResourcesSystemTagsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TagResourcesSystemTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TagResourcesSystemTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesSystemTagsResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesSystemTagsResponse) SetHeaders(v map[string]*string) *TagResourcesSystemTagsResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesSystemTagsResponse) SetBody(v *TagResourcesSystemTagsResponseBody) *TagResourcesSystemTagsResponse {
	s.Body = v
	return s
}

type ModifyFlowJobRequest struct {
	RegionId        *string                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProjectId       *string                             `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Id              *string                             `json:"Id,omitempty" xml:"Id,omitempty"`
	Name            *string                             `json:"Name,omitempty" xml:"Name,omitempty"`
	Description     *string                             `json:"Description,omitempty" xml:"Description,omitempty"`
	FailAct         *string                             `json:"FailAct,omitempty" xml:"FailAct,omitempty"`
	RetryPolicy     *string                             `json:"RetryPolicy,omitempty" xml:"RetryPolicy,omitempty"`
	Params          *string                             `json:"Params,omitempty" xml:"Params,omitempty"`
	ParamConf       *string                             `json:"ParamConf,omitempty" xml:"ParamConf,omitempty"`
	CustomVariables *string                             `json:"CustomVariables,omitempty" xml:"CustomVariables,omitempty"`
	EnvConf         *string                             `json:"EnvConf,omitempty" xml:"EnvConf,omitempty"`
	RunConf         *string                             `json:"RunConf,omitempty" xml:"RunConf,omitempty"`
	MonitorConf     *string                             `json:"MonitorConf,omitempty" xml:"MonitorConf,omitempty"`
	Mode            *string                             `json:"Mode,omitempty" xml:"Mode,omitempty"`
	ClusterId       *string                             `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	AlertConf       *string                             `json:"AlertConf,omitempty" xml:"AlertConf,omitempty"`
	ResourceList    []*ModifyFlowJobRequestResourceList `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Repeated"`
}

func (s ModifyFlowJobRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowJobRequest) GoString() string {
	return s.String()
}

func (s *ModifyFlowJobRequest) SetRegionId(v string) *ModifyFlowJobRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyFlowJobRequest) SetProjectId(v string) *ModifyFlowJobRequest {
	s.ProjectId = &v
	return s
}

func (s *ModifyFlowJobRequest) SetId(v string) *ModifyFlowJobRequest {
	s.Id = &v
	return s
}

func (s *ModifyFlowJobRequest) SetName(v string) *ModifyFlowJobRequest {
	s.Name = &v
	return s
}

func (s *ModifyFlowJobRequest) SetDescription(v string) *ModifyFlowJobRequest {
	s.Description = &v
	return s
}

func (s *ModifyFlowJobRequest) SetFailAct(v string) *ModifyFlowJobRequest {
	s.FailAct = &v
	return s
}

func (s *ModifyFlowJobRequest) SetRetryPolicy(v string) *ModifyFlowJobRequest {
	s.RetryPolicy = &v
	return s
}

func (s *ModifyFlowJobRequest) SetParams(v string) *ModifyFlowJobRequest {
	s.Params = &v
	return s
}

func (s *ModifyFlowJobRequest) SetParamConf(v string) *ModifyFlowJobRequest {
	s.ParamConf = &v
	return s
}

func (s *ModifyFlowJobRequest) SetCustomVariables(v string) *ModifyFlowJobRequest {
	s.CustomVariables = &v
	return s
}

func (s *ModifyFlowJobRequest) SetEnvConf(v string) *ModifyFlowJobRequest {
	s.EnvConf = &v
	return s
}

func (s *ModifyFlowJobRequest) SetRunConf(v string) *ModifyFlowJobRequest {
	s.RunConf = &v
	return s
}

func (s *ModifyFlowJobRequest) SetMonitorConf(v string) *ModifyFlowJobRequest {
	s.MonitorConf = &v
	return s
}

func (s *ModifyFlowJobRequest) SetMode(v string) *ModifyFlowJobRequest {
	s.Mode = &v
	return s
}

func (s *ModifyFlowJobRequest) SetClusterId(v string) *ModifyFlowJobRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyFlowJobRequest) SetAlertConf(v string) *ModifyFlowJobRequest {
	s.AlertConf = &v
	return s
}

func (s *ModifyFlowJobRequest) SetResourceList(v []*ModifyFlowJobRequestResourceList) *ModifyFlowJobRequest {
	s.ResourceList = v
	return s
}

type ModifyFlowJobRequestResourceList struct {
	Path  *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
}

func (s ModifyFlowJobRequestResourceList) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowJobRequestResourceList) GoString() string {
	return s.String()
}

func (s *ModifyFlowJobRequestResourceList) SetPath(v string) *ModifyFlowJobRequestResourceList {
	s.Path = &v
	return s
}

func (s *ModifyFlowJobRequestResourceList) SetAlias(v string) *ModifyFlowJobRequestResourceList {
	s.Alias = &v
	return s
}

type ModifyFlowJobResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyFlowJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowJobResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFlowJobResponseBody) SetData(v bool) *ModifyFlowJobResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyFlowJobResponseBody) SetRequestId(v string) *ModifyFlowJobResponseBody {
	s.RequestId = &v
	return s
}

type ModifyFlowJobResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyFlowJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyFlowJobResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowJobResponse) GoString() string {
	return s.String()
}

func (s *ModifyFlowJobResponse) SetHeaders(v map[string]*string) *ModifyFlowJobResponse {
	s.Headers = v
	return s
}

func (s *ModifyFlowJobResponse) SetBody(v *ModifyFlowJobResponseBody) *ModifyFlowJobResponse {
	s.Body = v
	return s
}

type DeleteFlowRequest struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowRequest) GoString() string {
	return s.String()
}

func (s *DeleteFlowRequest) SetId(v string) *DeleteFlowRequest {
	s.Id = &v
	return s
}

func (s *DeleteFlowRequest) SetProjectId(v string) *DeleteFlowRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteFlowRequest) SetRegionId(v string) *DeleteFlowRequest {
	s.RegionId = &v
	return s
}

type DeleteFlowResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFlowResponseBody) SetData(v bool) *DeleteFlowResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteFlowResponseBody) SetRequestId(v string) *DeleteFlowResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFlowResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowResponse) GoString() string {
	return s.String()
}

func (s *DeleteFlowResponse) SetHeaders(v map[string]*string) *DeleteFlowResponse {
	s.Headers = v
	return s
}

func (s *DeleteFlowResponse) SetBody(v *DeleteFlowResponseBody) *DeleteFlowResponse {
	s.Body = v
	return s
}

type CreateFlowEditLockRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	EntityId        *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	Force           *bool   `json:"Force,omitempty" xml:"Force,omitempty"`
}

func (s CreateFlowEditLockRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowEditLockRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowEditLockRequest) SetResourceOwnerId(v int64) *CreateFlowEditLockRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateFlowEditLockRequest) SetRegionId(v string) *CreateFlowEditLockRequest {
	s.RegionId = &v
	return s
}

func (s *CreateFlowEditLockRequest) SetEntityId(v string) *CreateFlowEditLockRequest {
	s.EntityId = &v
	return s
}

func (s *CreateFlowEditLockRequest) SetForce(v bool) *CreateFlowEditLockRequest {
	s.Force = &v
	return s
}

type CreateFlowEditLockResponseBody struct {
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	EntityId  *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	UserId    *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	BizId     *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	OwnerId   *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s CreateFlowEditLockResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowEditLockResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlowEditLockResponseBody) SetStatus(v string) *CreateFlowEditLockResponseBody {
	s.Status = &v
	return s
}

func (s *CreateFlowEditLockResponseBody) SetRequestId(v string) *CreateFlowEditLockResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFlowEditLockResponseBody) SetEntityId(v string) *CreateFlowEditLockResponseBody {
	s.EntityId = &v
	return s
}

func (s *CreateFlowEditLockResponseBody) SetUserId(v string) *CreateFlowEditLockResponseBody {
	s.UserId = &v
	return s
}

func (s *CreateFlowEditLockResponseBody) SetBizId(v string) *CreateFlowEditLockResponseBody {
	s.BizId = &v
	return s
}

func (s *CreateFlowEditLockResponseBody) SetOwnerId(v string) *CreateFlowEditLockResponseBody {
	s.OwnerId = &v
	return s
}

type CreateFlowEditLockResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFlowEditLockResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFlowEditLockResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowEditLockResponse) GoString() string {
	return s.String()
}

func (s *CreateFlowEditLockResponse) SetHeaders(v map[string]*string) *CreateFlowEditLockResponse {
	s.Headers = v
	return s
}

func (s *CreateFlowEditLockResponse) SetBody(v *CreateFlowEditLockResponseBody) *CreateFlowEditLockResponse {
	s.Body = v
	return s
}

type DescribeFlowNodeInstanceRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeFlowNodeInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowNodeInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowNodeInstanceRequest) SetProjectId(v string) *DescribeFlowNodeInstanceRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeFlowNodeInstanceRequest) SetId(v string) *DescribeFlowNodeInstanceRequest {
	s.Id = &v
	return s
}

func (s *DescribeFlowNodeInstanceRequest) SetRegionId(v string) *DescribeFlowNodeInstanceRequest {
	s.RegionId = &v
	return s
}

type DescribeFlowNodeInstanceResponseBody struct {
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
	EnvConf          *string `json:"EnvConf,omitempty" xml:"EnvConf,omitempty"`
	ProjectId        *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RetryInterval    *string `json:"RetryInterval,omitempty" xml:"RetryInterval,omitempty"`
	JobType          *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	Mode             *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	ExternalInfo     *string `json:"ExternalInfo,omitempty" xml:"ExternalInfo,omitempty"`
	GmtModified      *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	ExternalChildIds *string `json:"ExternalChildIds,omitempty" xml:"ExternalChildIds,omitempty"`
	MonitorConf      *string `json:"MonitorConf,omitempty" xml:"MonitorConf,omitempty"`
	ExternalStatus   *string `json:"ExternalStatus,omitempty" xml:"ExternalStatus,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	JobName          *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	RetryPolicy      *string `json:"RetryPolicy,omitempty" xml:"RetryPolicy,omitempty"`
	Adhoc            *bool   `json:"Adhoc,omitempty" xml:"Adhoc,omitempty"`
	ExternalId       *string `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	Pending          *bool   `json:"Pending,omitempty" xml:"Pending,omitempty"`
	ClusterId        *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	MaxRetry         *string `json:"MaxRetry,omitempty" xml:"MaxRetry,omitempty"`
	FailAct          *string `json:"FailAct,omitempty" xml:"FailAct,omitempty"`
	JobParams        *string `json:"JobParams,omitempty" xml:"JobParams,omitempty"`
	FlowInstanceId   *string `json:"FlowInstanceId,omitempty" xml:"FlowInstanceId,omitempty"`
	ParamConf        *string `json:"ParamConf,omitempty" xml:"ParamConf,omitempty"`
	HostName         *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	FlowId           *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	Retries          *int32  `json:"Retries,omitempty" xml:"Retries,omitempty"`
	EndTime          *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime        *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	RunConf          *string `json:"RunConf,omitempty" xml:"RunConf,omitempty"`
	ExternalSubId    *string `json:"ExternalSubId,omitempty" xml:"ExternalSubId,omitempty"`
	NodeName         *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	JobId            *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ClusterName      *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	GmtCreate        *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Duration         *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Id               *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeFlowNodeInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowNodeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowNodeInstanceResponseBody) SetStatus(v string) *DescribeFlowNodeInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetType(v string) *DescribeFlowNodeInstanceResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetEnvConf(v string) *DescribeFlowNodeInstanceResponseBody {
	s.EnvConf = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetProjectId(v string) *DescribeFlowNodeInstanceResponseBody {
	s.ProjectId = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetRetryInterval(v string) *DescribeFlowNodeInstanceResponseBody {
	s.RetryInterval = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetJobType(v string) *DescribeFlowNodeInstanceResponseBody {
	s.JobType = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetMode(v string) *DescribeFlowNodeInstanceResponseBody {
	s.Mode = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetExternalInfo(v string) *DescribeFlowNodeInstanceResponseBody {
	s.ExternalInfo = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetGmtModified(v int64) *DescribeFlowNodeInstanceResponseBody {
	s.GmtModified = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetExternalChildIds(v string) *DescribeFlowNodeInstanceResponseBody {
	s.ExternalChildIds = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetMonitorConf(v string) *DescribeFlowNodeInstanceResponseBody {
	s.MonitorConf = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetExternalStatus(v string) *DescribeFlowNodeInstanceResponseBody {
	s.ExternalStatus = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetRequestId(v string) *DescribeFlowNodeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetJobName(v string) *DescribeFlowNodeInstanceResponseBody {
	s.JobName = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetRetryPolicy(v string) *DescribeFlowNodeInstanceResponseBody {
	s.RetryPolicy = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetAdhoc(v bool) *DescribeFlowNodeInstanceResponseBody {
	s.Adhoc = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetExternalId(v string) *DescribeFlowNodeInstanceResponseBody {
	s.ExternalId = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetPending(v bool) *DescribeFlowNodeInstanceResponseBody {
	s.Pending = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetClusterId(v string) *DescribeFlowNodeInstanceResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetMaxRetry(v string) *DescribeFlowNodeInstanceResponseBody {
	s.MaxRetry = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetFailAct(v string) *DescribeFlowNodeInstanceResponseBody {
	s.FailAct = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetJobParams(v string) *DescribeFlowNodeInstanceResponseBody {
	s.JobParams = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetFlowInstanceId(v string) *DescribeFlowNodeInstanceResponseBody {
	s.FlowInstanceId = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetParamConf(v string) *DescribeFlowNodeInstanceResponseBody {
	s.ParamConf = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetHostName(v string) *DescribeFlowNodeInstanceResponseBody {
	s.HostName = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetFlowId(v string) *DescribeFlowNodeInstanceResponseBody {
	s.FlowId = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetRetries(v int32) *DescribeFlowNodeInstanceResponseBody {
	s.Retries = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetEndTime(v int64) *DescribeFlowNodeInstanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetStartTime(v int64) *DescribeFlowNodeInstanceResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetRunConf(v string) *DescribeFlowNodeInstanceResponseBody {
	s.RunConf = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetExternalSubId(v string) *DescribeFlowNodeInstanceResponseBody {
	s.ExternalSubId = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetNodeName(v string) *DescribeFlowNodeInstanceResponseBody {
	s.NodeName = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetJobId(v string) *DescribeFlowNodeInstanceResponseBody {
	s.JobId = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetClusterName(v string) *DescribeFlowNodeInstanceResponseBody {
	s.ClusterName = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetGmtCreate(v int64) *DescribeFlowNodeInstanceResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetDuration(v int64) *DescribeFlowNodeInstanceResponseBody {
	s.Duration = &v
	return s
}

func (s *DescribeFlowNodeInstanceResponseBody) SetId(v string) *DescribeFlowNodeInstanceResponseBody {
	s.Id = &v
	return s
}

type DescribeFlowNodeInstanceResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFlowNodeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFlowNodeInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowNodeInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowNodeInstanceResponse) SetHeaders(v map[string]*string) *DescribeFlowNodeInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowNodeInstanceResponse) SetBody(v *DescribeFlowNodeInstanceResponseBody) *DescribeFlowNodeInstanceResponse {
	s.Body = v
	return s
}

type DetachAndReleaseClusterEniRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TargetClusterId *string `json:"TargetClusterId,omitempty" xml:"TargetClusterId,omitempty"`
	VswitchId       *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s DetachAndReleaseClusterEniRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachAndReleaseClusterEniRequest) GoString() string {
	return s.String()
}

func (s *DetachAndReleaseClusterEniRequest) SetResourceOwnerId(v int64) *DetachAndReleaseClusterEniRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DetachAndReleaseClusterEniRequest) SetRegionId(v string) *DetachAndReleaseClusterEniRequest {
	s.RegionId = &v
	return s
}

func (s *DetachAndReleaseClusterEniRequest) SetTargetClusterId(v string) *DetachAndReleaseClusterEniRequest {
	s.TargetClusterId = &v
	return s
}

func (s *DetachAndReleaseClusterEniRequest) SetVswitchId(v string) *DetachAndReleaseClusterEniRequest {
	s.VswitchId = &v
	return s
}

type DetachAndReleaseClusterEniResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachAndReleaseClusterEniResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachAndReleaseClusterEniResponseBody) GoString() string {
	return s.String()
}

func (s *DetachAndReleaseClusterEniResponseBody) SetRequestId(v string) *DetachAndReleaseClusterEniResponseBody {
	s.RequestId = &v
	return s
}

type DetachAndReleaseClusterEniResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetachAndReleaseClusterEniResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetachAndReleaseClusterEniResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachAndReleaseClusterEniResponse) GoString() string {
	return s.String()
}

func (s *DetachAndReleaseClusterEniResponse) SetHeaders(v map[string]*string) *DetachAndReleaseClusterEniResponse {
	s.Headers = v
	return s
}

func (s *DetachAndReleaseClusterEniResponse) SetBody(v *DetachAndReleaseClusterEniResponseBody) *DetachAndReleaseClusterEniResponse {
	s.Body = v
	return s
}

type DescribeScalingGroupInstanceRequest struct {
	ResourceOwnerId   *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId   *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ScalingGroupBizId *string `json:"ScalingGroupBizId,omitempty" xml:"ScalingGroupBizId,omitempty"`
	HostGroupBizId    *string `json:"HostGroupBizId,omitempty" xml:"HostGroupBizId,omitempty"`
}

func (s DescribeScalingGroupInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingGroupInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupInstanceRequest) SetResourceOwnerId(v int64) *DescribeScalingGroupInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeScalingGroupInstanceRequest) SetRegionId(v string) *DescribeScalingGroupInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeScalingGroupInstanceRequest) SetResourceGroupId(v string) *DescribeScalingGroupInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeScalingGroupInstanceRequest) SetScalingGroupBizId(v string) *DescribeScalingGroupInstanceRequest {
	s.ScalingGroupBizId = &v
	return s
}

func (s *DescribeScalingGroupInstanceRequest) SetHostGroupBizId(v string) *DescribeScalingGroupInstanceRequest {
	s.HostGroupBizId = &v
	return s
}

type DescribeScalingGroupInstanceResponseBody struct {
	// Id of the request
	RequestId                 *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HostGroupId               *string                                                  `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	ScalingGroupId            *string                                                  `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	MinSize                   *int64                                                   `json:"MinSize,omitempty" xml:"MinSize,omitempty"`
	MaxSize                   *int64                                                   `json:"MaxSize,omitempty" xml:"MaxSize,omitempty"`
	DefaultCooldown           *int64                                                   `json:"DefaultCooldown,omitempty" xml:"DefaultCooldown,omitempty"`
	ActiveRuleCategory        *string                                                  `json:"ActiveRuleCategory,omitempty" xml:"ActiveRuleCategory,omitempty"`
	WithGrace                 *bool                                                    `json:"WithGrace,omitempty" xml:"WithGrace,omitempty"`
	TimeoutWithGrace          *int64                                                   `json:"TimeoutWithGrace,omitempty" xml:"TimeoutWithGrace,omitempty"`
	MultiAvailablePolicy      *string                                                  `json:"MultiAvailablePolicy,omitempty" xml:"MultiAvailablePolicy,omitempty"`
	MultiAvailablePolicyParam *string                                                  `json:"MultiAvailablePolicyParam,omitempty" xml:"MultiAvailablePolicyParam,omitempty"`
	ScalingConfig             *DescribeScalingGroupInstanceResponseBodyScalingConfig   `json:"ScalingConfig,omitempty" xml:"ScalingConfig,omitempty" type:"Struct"`
	ScalingRuleList           *DescribeScalingGroupInstanceResponseBodyScalingRuleList `json:"ScalingRuleList,omitempty" xml:"ScalingRuleList,omitempty" type:"Struct"`
}

func (s DescribeScalingGroupInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingGroupInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupInstanceResponseBody) SetRequestId(v string) *DescribeScalingGroupInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBody) SetHostGroupId(v string) *DescribeScalingGroupInstanceResponseBody {
	s.HostGroupId = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBody) SetScalingGroupId(v string) *DescribeScalingGroupInstanceResponseBody {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBody) SetMinSize(v int64) *DescribeScalingGroupInstanceResponseBody {
	s.MinSize = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBody) SetMaxSize(v int64) *DescribeScalingGroupInstanceResponseBody {
	s.MaxSize = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBody) SetDefaultCooldown(v int64) *DescribeScalingGroupInstanceResponseBody {
	s.DefaultCooldown = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBody) SetActiveRuleCategory(v string) *DescribeScalingGroupInstanceResponseBody {
	s.ActiveRuleCategory = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBody) SetWithGrace(v bool) *DescribeScalingGroupInstanceResponseBody {
	s.WithGrace = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBody) SetTimeoutWithGrace(v int64) *DescribeScalingGroupInstanceResponseBody {
	s.TimeoutWithGrace = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBody) SetMultiAvailablePolicy(v string) *DescribeScalingGroupInstanceResponseBody {
	s.MultiAvailablePolicy = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBody) SetMultiAvailablePolicyParam(v string) *DescribeScalingGroupInstanceResponseBody {
	s.MultiAvailablePolicyParam = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBody) SetScalingConfig(v *DescribeScalingGroupInstanceResponseBodyScalingConfig) *DescribeScalingGroupInstanceResponseBody {
	s.ScalingConfig = v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBody) SetScalingRuleList(v *DescribeScalingGroupInstanceResponseBodyScalingRuleList) *DescribeScalingGroupInstanceResponseBody {
	s.ScalingRuleList = v
	return s
}

type DescribeScalingGroupInstanceResponseBodyScalingConfig struct {
	SpotStrategy     *string                                                                `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	PayType          *string                                                                `json:"PayType,omitempty" xml:"PayType,omitempty"`
	DataDiskCategory *string                                                                `json:"DataDiskCategory,omitempty" xml:"DataDiskCategory,omitempty"`
	DataDiskSize     *int64                                                                 `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	DataDiskCount    *int64                                                                 `json:"DataDiskCount,omitempty" xml:"DataDiskCount,omitempty"`
	SysDiskCategory  *string                                                                `json:"SysDiskCategory,omitempty" xml:"SysDiskCategory,omitempty"`
	SysDiskSize      *int64                                                                 `json:"SysDiskSize,omitempty" xml:"SysDiskSize,omitempty"`
	InstanceTypeList *DescribeScalingGroupInstanceResponseBodyScalingConfigInstanceTypeList `json:"InstanceTypeList,omitempty" xml:"InstanceTypeList,omitempty" type:"Struct"`
	SpotPriceLimits  *DescribeScalingGroupInstanceResponseBodyScalingConfigSpotPriceLimits  `json:"SpotPriceLimits,omitempty" xml:"SpotPriceLimits,omitempty" type:"Struct"`
}

func (s DescribeScalingGroupInstanceResponseBodyScalingConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingGroupInstanceResponseBodyScalingConfig) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingConfig) SetSpotStrategy(v string) *DescribeScalingGroupInstanceResponseBodyScalingConfig {
	s.SpotStrategy = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingConfig) SetPayType(v string) *DescribeScalingGroupInstanceResponseBodyScalingConfig {
	s.PayType = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingConfig) SetDataDiskCategory(v string) *DescribeScalingGroupInstanceResponseBodyScalingConfig {
	s.DataDiskCategory = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingConfig) SetDataDiskSize(v int64) *DescribeScalingGroupInstanceResponseBodyScalingConfig {
	s.DataDiskSize = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingConfig) SetDataDiskCount(v int64) *DescribeScalingGroupInstanceResponseBodyScalingConfig {
	s.DataDiskCount = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingConfig) SetSysDiskCategory(v string) *DescribeScalingGroupInstanceResponseBodyScalingConfig {
	s.SysDiskCategory = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingConfig) SetSysDiskSize(v int64) *DescribeScalingGroupInstanceResponseBodyScalingConfig {
	s.SysDiskSize = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingConfig) SetInstanceTypeList(v *DescribeScalingGroupInstanceResponseBodyScalingConfigInstanceTypeList) *DescribeScalingGroupInstanceResponseBodyScalingConfig {
	s.InstanceTypeList = v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingConfig) SetSpotPriceLimits(v *DescribeScalingGroupInstanceResponseBodyScalingConfigSpotPriceLimits) *DescribeScalingGroupInstanceResponseBodyScalingConfig {
	s.SpotPriceLimits = v
	return s
}

type DescribeScalingGroupInstanceResponseBodyScalingConfigInstanceTypeList struct {
	InstanceTypeList []*string `json:"instanceTypeList,omitempty" xml:"instanceTypeList,omitempty" type:"Repeated"`
}

func (s DescribeScalingGroupInstanceResponseBodyScalingConfigInstanceTypeList) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingGroupInstanceResponseBodyScalingConfigInstanceTypeList) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingConfigInstanceTypeList) SetInstanceTypeList(v []*string) *DescribeScalingGroupInstanceResponseBodyScalingConfigInstanceTypeList {
	s.InstanceTypeList = v
	return s
}

type DescribeScalingGroupInstanceResponseBodyScalingConfigSpotPriceLimits struct {
	SpotPriceLimits []*DescribeScalingGroupInstanceResponseBodyScalingConfigSpotPriceLimitsSpotPriceLimits `json:"spotPriceLimits,omitempty" xml:"spotPriceLimits,omitempty" type:"Repeated"`
}

func (s DescribeScalingGroupInstanceResponseBodyScalingConfigSpotPriceLimits) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingGroupInstanceResponseBodyScalingConfigSpotPriceLimits) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingConfigSpotPriceLimits) SetSpotPriceLimits(v []*DescribeScalingGroupInstanceResponseBodyScalingConfigSpotPriceLimitsSpotPriceLimits) *DescribeScalingGroupInstanceResponseBodyScalingConfigSpotPriceLimits {
	s.SpotPriceLimits = v
	return s
}

type DescribeScalingGroupInstanceResponseBodyScalingConfigSpotPriceLimitsSpotPriceLimits struct {
	InstanceType *string  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	PriceLimit   *float32 `json:"PriceLimit,omitempty" xml:"PriceLimit,omitempty"`
}

func (s DescribeScalingGroupInstanceResponseBodyScalingConfigSpotPriceLimitsSpotPriceLimits) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingGroupInstanceResponseBodyScalingConfigSpotPriceLimitsSpotPriceLimits) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingConfigSpotPriceLimitsSpotPriceLimits) SetInstanceType(v string) *DescribeScalingGroupInstanceResponseBodyScalingConfigSpotPriceLimitsSpotPriceLimits {
	s.InstanceType = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingConfigSpotPriceLimitsSpotPriceLimits) SetPriceLimit(v float32) *DescribeScalingGroupInstanceResponseBodyScalingConfigSpotPriceLimitsSpotPriceLimits {
	s.PriceLimit = &v
	return s
}

type DescribeScalingGroupInstanceResponseBodyScalingRuleList struct {
	ScalingRule []*DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRule `json:"ScalingRule,omitempty" xml:"ScalingRule,omitempty" type:"Repeated"`
}

func (s DescribeScalingGroupInstanceResponseBodyScalingRuleList) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingGroupInstanceResponseBodyScalingRuleList) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingRuleList) SetScalingRule(v []*DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRule) *DescribeScalingGroupInstanceResponseBodyScalingRuleList {
	s.ScalingRule = v
	return s
}

type DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRule struct {
	RuleCategory         *string                                                                              `json:"RuleCategory,omitempty" xml:"RuleCategory,omitempty"`
	EssScalingRuleId     *string                                                                              `json:"EssScalingRuleId,omitempty" xml:"EssScalingRuleId,omitempty"`
	ScalingGroupId       *int64                                                                               `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	RuleName             *string                                                                              `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	AdjustmentType       *string                                                                              `json:"AdjustmentType,omitempty" xml:"AdjustmentType,omitempty"`
	AdjustmentValue      *int64                                                                               `json:"AdjustmentValue,omitempty" xml:"AdjustmentValue,omitempty"`
	Cooldown             *int64                                                                               `json:"Cooldown,omitempty" xml:"Cooldown,omitempty"`
	Status               *string                                                                              `json:"Status,omitempty" xml:"Status,omitempty"`
	LaunchTime           *string                                                                              `json:"LaunchTime,omitempty" xml:"LaunchTime,omitempty"`
	LaunchExpirationTime *int64                                                                               `json:"LaunchExpirationTime,omitempty" xml:"LaunchExpirationTime,omitempty"`
	RecurrenceType       *string                                                                              `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	RecurrenceValue      *string                                                                              `json:"RecurrenceValue,omitempty" xml:"RecurrenceValue,omitempty"`
	RecurrenceEndTime    *string                                                                              `json:"RecurrenceEndTime,omitempty" xml:"RecurrenceEndTime,omitempty"`
	WithGrace            *bool                                                                                `json:"WithGrace,omitempty" xml:"WithGrace,omitempty"`
	TimeoutWithGrace     *int64                                                                               `json:"TimeoutWithGrace,omitempty" xml:"TimeoutWithGrace,omitempty"`
	SchedulerTrigger     *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRuleSchedulerTrigger  `json:"SchedulerTrigger,omitempty" xml:"SchedulerTrigger,omitempty" type:"Struct"`
	CloudWatchTrigger    *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRuleCloudWatchTrigger `json:"CloudWatchTrigger,omitempty" xml:"CloudWatchTrigger,omitempty" type:"Struct"`
}

func (s DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRule) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRule) SetRuleCategory(v string) *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRule {
	s.RuleCategory = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRule) SetEssScalingRuleId(v string) *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRule {
	s.EssScalingRuleId = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRule) SetScalingGroupId(v int64) *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRule {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRule) SetRuleName(v string) *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRule {
	s.RuleName = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRule) SetAdjustmentType(v string) *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRule {
	s.AdjustmentType = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRule) SetAdjustmentValue(v int64) *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRule {
	s.AdjustmentValue = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRule) SetCooldown(v int64) *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRule {
	s.Cooldown = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRule) SetStatus(v string) *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRule {
	s.Status = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRule) SetLaunchTime(v string) *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRule {
	s.LaunchTime = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRule) SetLaunchExpirationTime(v int64) *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRule {
	s.LaunchExpirationTime = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRule) SetRecurrenceType(v string) *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRule {
	s.RecurrenceType = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRule) SetRecurrenceValue(v string) *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRule {
	s.RecurrenceValue = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRule) SetRecurrenceEndTime(v string) *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRule {
	s.RecurrenceEndTime = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRule) SetWithGrace(v bool) *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRule {
	s.WithGrace = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRule) SetTimeoutWithGrace(v int64) *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRule {
	s.TimeoutWithGrace = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRule) SetSchedulerTrigger(v *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRuleSchedulerTrigger) *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRule {
	s.SchedulerTrigger = v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRule) SetCloudWatchTrigger(v *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRuleCloudWatchTrigger) *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRule {
	s.CloudWatchTrigger = v
	return s
}

type DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRuleSchedulerTrigger struct {
	LaunchTime           *int64  `json:"LaunchTime,omitempty" xml:"LaunchTime,omitempty"`
	LaunchExpirationTime *int64  `json:"LaunchExpirationTime,omitempty" xml:"LaunchExpirationTime,omitempty"`
	RecurrenceType       *string `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	RecurrenceValue      *string `json:"RecurrenceValue,omitempty" xml:"RecurrenceValue,omitempty"`
	RecurrenceEndTime    *int64  `json:"RecurrenceEndTime,omitempty" xml:"RecurrenceEndTime,omitempty"`
}

func (s DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRuleSchedulerTrigger) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRuleSchedulerTrigger) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRuleSchedulerTrigger) SetLaunchTime(v int64) *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRuleSchedulerTrigger {
	s.LaunchTime = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRuleSchedulerTrigger) SetLaunchExpirationTime(v int64) *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRuleSchedulerTrigger {
	s.LaunchExpirationTime = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRuleSchedulerTrigger) SetRecurrenceType(v string) *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRuleSchedulerTrigger {
	s.RecurrenceType = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRuleSchedulerTrigger) SetRecurrenceValue(v string) *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRuleSchedulerTrigger {
	s.RecurrenceValue = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRuleSchedulerTrigger) SetRecurrenceEndTime(v int64) *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRuleSchedulerTrigger {
	s.RecurrenceEndTime = &v
	return s
}

type DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRuleCloudWatchTrigger struct {
	MetricName         *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	Period             *int64  `json:"Period,omitempty" xml:"Period,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	EvaluationCount    *string `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	Unit               *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	MetricDisplayName  *string `json:"MetricDisplayName,omitempty" xml:"MetricDisplayName,omitempty"`
}

func (s DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRuleCloudWatchTrigger) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRuleCloudWatchTrigger) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRuleCloudWatchTrigger) SetMetricName(v string) *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRuleCloudWatchTrigger {
	s.MetricName = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRuleCloudWatchTrigger) SetPeriod(v int64) *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRuleCloudWatchTrigger {
	s.Period = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRuleCloudWatchTrigger) SetStatistics(v string) *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRuleCloudWatchTrigger {
	s.Statistics = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRuleCloudWatchTrigger) SetComparisonOperator(v string) *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRuleCloudWatchTrigger {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRuleCloudWatchTrigger) SetThreshold(v string) *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRuleCloudWatchTrigger {
	s.Threshold = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRuleCloudWatchTrigger) SetEvaluationCount(v string) *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRuleCloudWatchTrigger {
	s.EvaluationCount = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRuleCloudWatchTrigger) SetUnit(v string) *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRuleCloudWatchTrigger {
	s.Unit = &v
	return s
}

func (s *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRuleCloudWatchTrigger) SetMetricDisplayName(v string) *DescribeScalingGroupInstanceResponseBodyScalingRuleListScalingRuleCloudWatchTrigger {
	s.MetricDisplayName = &v
	return s
}

type DescribeScalingGroupInstanceResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeScalingGroupInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScalingGroupInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingGroupInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupInstanceResponse) SetHeaders(v map[string]*string) *DescribeScalingGroupInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeScalingGroupInstanceResponse) SetBody(v *DescribeScalingGroupInstanceResponseBody) *DescribeScalingGroupInstanceResponse {
	s.Body = v
	return s
}

type CreateClusterHostGroupRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	HostGroupName   *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
	HostGroupType   *string `json:"HostGroupType,omitempty" xml:"HostGroupType,omitempty"`
	Comment         *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	PayType         *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	VswitchId       *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	HostGroupParams *string `json:"HostGroupParams,omitempty" xml:"HostGroupParams,omitempty"`
}

func (s CreateClusterHostGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterHostGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterHostGroupRequest) SetResourceOwnerId(v int64) *CreateClusterHostGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateClusterHostGroupRequest) SetRegionId(v string) *CreateClusterHostGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateClusterHostGroupRequest) SetClusterId(v string) *CreateClusterHostGroupRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateClusterHostGroupRequest) SetHostGroupName(v string) *CreateClusterHostGroupRequest {
	s.HostGroupName = &v
	return s
}

func (s *CreateClusterHostGroupRequest) SetHostGroupType(v string) *CreateClusterHostGroupRequest {
	s.HostGroupType = &v
	return s
}

func (s *CreateClusterHostGroupRequest) SetComment(v string) *CreateClusterHostGroupRequest {
	s.Comment = &v
	return s
}

func (s *CreateClusterHostGroupRequest) SetPayType(v string) *CreateClusterHostGroupRequest {
	s.PayType = &v
	return s
}

func (s *CreateClusterHostGroupRequest) SetVswitchId(v string) *CreateClusterHostGroupRequest {
	s.VswitchId = &v
	return s
}

func (s *CreateClusterHostGroupRequest) SetSecurityGroupId(v string) *CreateClusterHostGroupRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateClusterHostGroupRequest) SetHostGroupParams(v string) *CreateClusterHostGroupRequest {
	s.HostGroupParams = &v
	return s
}

type CreateClusterHostGroupResponseBody struct {
	// Id of the request
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
}

func (s CreateClusterHostGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterHostGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterHostGroupResponseBody) SetRequestId(v string) *CreateClusterHostGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClusterHostGroupResponseBody) SetClusterId(v string) *CreateClusterHostGroupResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateClusterHostGroupResponseBody) SetHostGroupId(v string) *CreateClusterHostGroupResponseBody {
	s.HostGroupId = &v
	return s
}

type CreateClusterHostGroupResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateClusterHostGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateClusterHostGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterHostGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateClusterHostGroupResponse) SetHeaders(v map[string]*string) *CreateClusterHostGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateClusterHostGroupResponse) SetBody(v *CreateClusterHostGroupResponseBody) *CreateClusterHostGroupResponse {
	s.Body = v
	return s
}

type DescribeClusterTemplateRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	BizId           *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeClusterTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterTemplateRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterTemplateRequest) SetResourceOwnerId(v int64) *DescribeClusterTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeClusterTemplateRequest) SetBizId(v string) *DescribeClusterTemplateRequest {
	s.BizId = &v
	return s
}

func (s *DescribeClusterTemplateRequest) SetResourceGroupId(v string) *DescribeClusterTemplateRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeClusterTemplateResponseBody struct {
	RequestId    *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateInfo *DescribeClusterTemplateResponseBodyTemplateInfo `json:"TemplateInfo,omitempty" xml:"TemplateInfo,omitempty" type:"Struct"`
}

func (s DescribeClusterTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterTemplateResponseBody) SetRequestId(v string) *DescribeClusterTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterTemplateResponseBody) SetTemplateInfo(v *DescribeClusterTemplateResponseBodyTemplateInfo) *DescribeClusterTemplateResponseBody {
	s.TemplateInfo = v
	return s
}

type DescribeClusterTemplateResponseBodyTemplateInfo struct {
	VpcId                  *string                                                             `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	KeyPairName            *string                                                             `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	LogEnable              *bool                                                               `json:"LogEnable,omitempty" xml:"LogEnable,omitempty"`
	SshEnable              *bool                                                               `json:"SshEnable,omitempty" xml:"SshEnable,omitempty"`
	HighAvailabilityEnable *bool                                                               `json:"HighAvailabilityEnable,omitempty" xml:"HighAvailabilityEnable,omitempty"`
	SecurityGroupId        *string                                                             `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	MasterPwd              *string                                                             `json:"MasterPwd,omitempty" xml:"MasterPwd,omitempty"`
	UserId                 *string                                                             `json:"UserId,omitempty" xml:"UserId,omitempty"`
	IsOpenPublicIp         *bool                                                               `json:"IsOpenPublicIp,omitempty" xml:"IsOpenPublicIp,omitempty"`
	AllowNotebook          *bool                                                               `json:"AllowNotebook,omitempty" xml:"AllowNotebook,omitempty"`
	GmtModified            *int64                                                              `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	TemplateName           *string                                                             `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	DepositType            *string                                                             `json:"DepositType,omitempty" xml:"DepositType,omitempty"`
	SecurityGroupName      *string                                                             `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
	CreateSource           *string                                                             `json:"CreateSource,omitempty" xml:"CreateSource,omitempty"`
	InstanceGeneration     *string                                                             `json:"InstanceGeneration,omitempty" xml:"InstanceGeneration,omitempty"`
	UseCustomHiveMetaDb    *bool                                                               `json:"UseCustomHiveMetaDb,omitempty" xml:"UseCustomHiveMetaDb,omitempty"`
	EasEnable              *bool                                                               `json:"EasEnable,omitempty" xml:"EasEnable,omitempty"`
	UserDefinedEmrEcsRole  *string                                                             `json:"UserDefinedEmrEcsRole,omitempty" xml:"UserDefinedEmrEcsRole,omitempty"`
	MetaStoreType          *string                                                             `json:"MetaStoreType,omitempty" xml:"MetaStoreType,omitempty"`
	MachineType            *string                                                             `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	UseLocalMetaDb         *bool                                                               `json:"UseLocalMetaDb,omitempty" xml:"UseLocalMetaDb,omitempty"`
	MasterNodeTotal        *int32                                                              `json:"MasterNodeTotal,omitempty" xml:"MasterNodeTotal,omitempty"`
	InitCustomHiveMetaDb   *bool                                                               `json:"InitCustomHiveMetaDb,omitempty" xml:"InitCustomHiveMetaDb,omitempty"`
	IoOptimized            *bool                                                               `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	MetaStoreConf          *string                                                             `json:"MetaStoreConf,omitempty" xml:"MetaStoreConf,omitempty"`
	VSwitchId              *string                                                             `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	Configurations         *string                                                             `json:"Configurations,omitempty" xml:"Configurations,omitempty"`
	EmrVer                 *string                                                             `json:"EmrVer,omitempty" xml:"EmrVer,omitempty"`
	LogPath                *string                                                             `json:"LogPath,omitempty" xml:"LogPath,omitempty"`
	ClusterType            *string                                                             `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	NetType                *string                                                             `json:"NetType,omitempty" xml:"NetType,omitempty"`
	ZoneId                 *string                                                             `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	GmtCreate              *int64                                                              `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id                     *string                                                             `json:"Id,omitempty" xml:"Id,omitempty"`
	BootstrapActionList    *DescribeClusterTemplateResponseBodyTemplateInfoBootstrapActionList `json:"BootstrapActionList,omitempty" xml:"BootstrapActionList,omitempty" type:"Struct"`
	HostGroupList          *DescribeClusterTemplateResponseBodyTemplateInfoHostGroupList       `json:"HostGroupList,omitempty" xml:"HostGroupList,omitempty" type:"Struct"`
	ConfigList             *DescribeClusterTemplateResponseBodyTemplateInfoConfigList          `json:"ConfigList,omitempty" xml:"ConfigList,omitempty" type:"Struct"`
	Tags                   *DescribeClusterTemplateResponseBodyTemplateInfoTags                `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	SoftwareInfoList       *DescribeClusterTemplateResponseBodyTemplateInfoSoftwareInfoList    `json:"SoftwareInfoList,omitempty" xml:"SoftwareInfoList,omitempty" type:"Struct"`
}

func (s DescribeClusterTemplateResponseBodyTemplateInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterTemplateResponseBodyTemplateInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfo) SetVpcId(v string) *DescribeClusterTemplateResponseBodyTemplateInfo {
	s.VpcId = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfo) SetKeyPairName(v string) *DescribeClusterTemplateResponseBodyTemplateInfo {
	s.KeyPairName = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfo) SetLogEnable(v bool) *DescribeClusterTemplateResponseBodyTemplateInfo {
	s.LogEnable = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfo) SetSshEnable(v bool) *DescribeClusterTemplateResponseBodyTemplateInfo {
	s.SshEnable = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfo) SetHighAvailabilityEnable(v bool) *DescribeClusterTemplateResponseBodyTemplateInfo {
	s.HighAvailabilityEnable = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfo) SetSecurityGroupId(v string) *DescribeClusterTemplateResponseBodyTemplateInfo {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfo) SetMasterPwd(v string) *DescribeClusterTemplateResponseBodyTemplateInfo {
	s.MasterPwd = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfo) SetUserId(v string) *DescribeClusterTemplateResponseBodyTemplateInfo {
	s.UserId = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfo) SetIsOpenPublicIp(v bool) *DescribeClusterTemplateResponseBodyTemplateInfo {
	s.IsOpenPublicIp = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfo) SetAllowNotebook(v bool) *DescribeClusterTemplateResponseBodyTemplateInfo {
	s.AllowNotebook = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfo) SetGmtModified(v int64) *DescribeClusterTemplateResponseBodyTemplateInfo {
	s.GmtModified = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfo) SetTemplateName(v string) *DescribeClusterTemplateResponseBodyTemplateInfo {
	s.TemplateName = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfo) SetDepositType(v string) *DescribeClusterTemplateResponseBodyTemplateInfo {
	s.DepositType = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfo) SetSecurityGroupName(v string) *DescribeClusterTemplateResponseBodyTemplateInfo {
	s.SecurityGroupName = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfo) SetCreateSource(v string) *DescribeClusterTemplateResponseBodyTemplateInfo {
	s.CreateSource = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfo) SetInstanceGeneration(v string) *DescribeClusterTemplateResponseBodyTemplateInfo {
	s.InstanceGeneration = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfo) SetUseCustomHiveMetaDb(v bool) *DescribeClusterTemplateResponseBodyTemplateInfo {
	s.UseCustomHiveMetaDb = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfo) SetEasEnable(v bool) *DescribeClusterTemplateResponseBodyTemplateInfo {
	s.EasEnable = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfo) SetUserDefinedEmrEcsRole(v string) *DescribeClusterTemplateResponseBodyTemplateInfo {
	s.UserDefinedEmrEcsRole = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfo) SetMetaStoreType(v string) *DescribeClusterTemplateResponseBodyTemplateInfo {
	s.MetaStoreType = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfo) SetMachineType(v string) *DescribeClusterTemplateResponseBodyTemplateInfo {
	s.MachineType = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfo) SetUseLocalMetaDb(v bool) *DescribeClusterTemplateResponseBodyTemplateInfo {
	s.UseLocalMetaDb = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfo) SetMasterNodeTotal(v int32) *DescribeClusterTemplateResponseBodyTemplateInfo {
	s.MasterNodeTotal = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfo) SetInitCustomHiveMetaDb(v bool) *DescribeClusterTemplateResponseBodyTemplateInfo {
	s.InitCustomHiveMetaDb = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfo) SetIoOptimized(v bool) *DescribeClusterTemplateResponseBodyTemplateInfo {
	s.IoOptimized = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfo) SetMetaStoreConf(v string) *DescribeClusterTemplateResponseBodyTemplateInfo {
	s.MetaStoreConf = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfo) SetVSwitchId(v string) *DescribeClusterTemplateResponseBodyTemplateInfo {
	s.VSwitchId = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfo) SetConfigurations(v string) *DescribeClusterTemplateResponseBodyTemplateInfo {
	s.Configurations = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfo) SetEmrVer(v string) *DescribeClusterTemplateResponseBodyTemplateInfo {
	s.EmrVer = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfo) SetLogPath(v string) *DescribeClusterTemplateResponseBodyTemplateInfo {
	s.LogPath = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfo) SetClusterType(v string) *DescribeClusterTemplateResponseBodyTemplateInfo {
	s.ClusterType = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfo) SetNetType(v string) *DescribeClusterTemplateResponseBodyTemplateInfo {
	s.NetType = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfo) SetZoneId(v string) *DescribeClusterTemplateResponseBodyTemplateInfo {
	s.ZoneId = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfo) SetGmtCreate(v int64) *DescribeClusterTemplateResponseBodyTemplateInfo {
	s.GmtCreate = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfo) SetId(v string) *DescribeClusterTemplateResponseBodyTemplateInfo {
	s.Id = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfo) SetBootstrapActionList(v *DescribeClusterTemplateResponseBodyTemplateInfoBootstrapActionList) *DescribeClusterTemplateResponseBodyTemplateInfo {
	s.BootstrapActionList = v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfo) SetHostGroupList(v *DescribeClusterTemplateResponseBodyTemplateInfoHostGroupList) *DescribeClusterTemplateResponseBodyTemplateInfo {
	s.HostGroupList = v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfo) SetConfigList(v *DescribeClusterTemplateResponseBodyTemplateInfoConfigList) *DescribeClusterTemplateResponseBodyTemplateInfo {
	s.ConfigList = v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfo) SetTags(v *DescribeClusterTemplateResponseBodyTemplateInfoTags) *DescribeClusterTemplateResponseBodyTemplateInfo {
	s.Tags = v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfo) SetSoftwareInfoList(v *DescribeClusterTemplateResponseBodyTemplateInfoSoftwareInfoList) *DescribeClusterTemplateResponseBodyTemplateInfo {
	s.SoftwareInfoList = v
	return s
}

type DescribeClusterTemplateResponseBodyTemplateInfoBootstrapActionList struct {
	BootstrapAction []*DescribeClusterTemplateResponseBodyTemplateInfoBootstrapActionListBootstrapAction `json:"BootstrapAction,omitempty" xml:"BootstrapAction,omitempty" type:"Repeated"`
}

func (s DescribeClusterTemplateResponseBodyTemplateInfoBootstrapActionList) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterTemplateResponseBodyTemplateInfoBootstrapActionList) GoString() string {
	return s.String()
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfoBootstrapActionList) SetBootstrapAction(v []*DescribeClusterTemplateResponseBodyTemplateInfoBootstrapActionListBootstrapAction) *DescribeClusterTemplateResponseBodyTemplateInfoBootstrapActionList {
	s.BootstrapAction = v
	return s
}

type DescribeClusterTemplateResponseBodyTemplateInfoBootstrapActionListBootstrapAction struct {
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Arg  *string `json:"Arg,omitempty" xml:"Arg,omitempty"`
}

func (s DescribeClusterTemplateResponseBodyTemplateInfoBootstrapActionListBootstrapAction) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterTemplateResponseBodyTemplateInfoBootstrapActionListBootstrapAction) GoString() string {
	return s.String()
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfoBootstrapActionListBootstrapAction) SetPath(v string) *DescribeClusterTemplateResponseBodyTemplateInfoBootstrapActionListBootstrapAction {
	s.Path = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfoBootstrapActionListBootstrapAction) SetName(v string) *DescribeClusterTemplateResponseBodyTemplateInfoBootstrapActionListBootstrapAction {
	s.Name = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfoBootstrapActionListBootstrapAction) SetArg(v string) *DescribeClusterTemplateResponseBodyTemplateInfoBootstrapActionListBootstrapAction {
	s.Arg = &v
	return s
}

type DescribeClusterTemplateResponseBodyTemplateInfoHostGroupList struct {
	HostGroup []*DescribeClusterTemplateResponseBodyTemplateInfoHostGroupListHostGroup `json:"HostGroup,omitempty" xml:"HostGroup,omitempty" type:"Repeated"`
}

func (s DescribeClusterTemplateResponseBodyTemplateInfoHostGroupList) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterTemplateResponseBodyTemplateInfoHostGroupList) GoString() string {
	return s.String()
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfoHostGroupList) SetHostGroup(v []*DescribeClusterTemplateResponseBodyTemplateInfoHostGroupListHostGroup) *DescribeClusterTemplateResponseBodyTemplateInfoHostGroupList {
	s.HostGroup = v
	return s
}

type DescribeClusterTemplateResponseBodyTemplateInfoHostGroupListHostGroup struct {
	SysDiskCapacity    *int32  `json:"SysDiskCapacity,omitempty" xml:"SysDiskCapacity,omitempty"`
	HostGroupType      *string `json:"HostGroupType,omitempty" xml:"HostGroupType,omitempty"`
	MultiInstanceTypes *string `json:"MultiInstanceTypes,omitempty" xml:"MultiInstanceTypes,omitempty"`
	SysDiskType        *string `json:"SysDiskType,omitempty" xml:"SysDiskType,omitempty"`
	ChargeType         *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	DiskType           *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	HostGroupId        *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	InstanceType       *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	DiskCount          *int32  `json:"DiskCount,omitempty" xml:"DiskCount,omitempty"`
	Period             *string `json:"Period,omitempty" xml:"Period,omitempty"`
	DiskCapacity       *int32  `json:"DiskCapacity,omitempty" xml:"DiskCapacity,omitempty"`
	NodeCount          *int32  `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	HostGroupName      *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
}

func (s DescribeClusterTemplateResponseBodyTemplateInfoHostGroupListHostGroup) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterTemplateResponseBodyTemplateInfoHostGroupListHostGroup) GoString() string {
	return s.String()
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfoHostGroupListHostGroup) SetSysDiskCapacity(v int32) *DescribeClusterTemplateResponseBodyTemplateInfoHostGroupListHostGroup {
	s.SysDiskCapacity = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfoHostGroupListHostGroup) SetHostGroupType(v string) *DescribeClusterTemplateResponseBodyTemplateInfoHostGroupListHostGroup {
	s.HostGroupType = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfoHostGroupListHostGroup) SetMultiInstanceTypes(v string) *DescribeClusterTemplateResponseBodyTemplateInfoHostGroupListHostGroup {
	s.MultiInstanceTypes = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfoHostGroupListHostGroup) SetSysDiskType(v string) *DescribeClusterTemplateResponseBodyTemplateInfoHostGroupListHostGroup {
	s.SysDiskType = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfoHostGroupListHostGroup) SetChargeType(v string) *DescribeClusterTemplateResponseBodyTemplateInfoHostGroupListHostGroup {
	s.ChargeType = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfoHostGroupListHostGroup) SetDiskType(v string) *DescribeClusterTemplateResponseBodyTemplateInfoHostGroupListHostGroup {
	s.DiskType = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfoHostGroupListHostGroup) SetHostGroupId(v string) *DescribeClusterTemplateResponseBodyTemplateInfoHostGroupListHostGroup {
	s.HostGroupId = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfoHostGroupListHostGroup) SetInstanceType(v string) *DescribeClusterTemplateResponseBodyTemplateInfoHostGroupListHostGroup {
	s.InstanceType = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfoHostGroupListHostGroup) SetDiskCount(v int32) *DescribeClusterTemplateResponseBodyTemplateInfoHostGroupListHostGroup {
	s.DiskCount = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfoHostGroupListHostGroup) SetPeriod(v string) *DescribeClusterTemplateResponseBodyTemplateInfoHostGroupListHostGroup {
	s.Period = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfoHostGroupListHostGroup) SetDiskCapacity(v int32) *DescribeClusterTemplateResponseBodyTemplateInfoHostGroupListHostGroup {
	s.DiskCapacity = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfoHostGroupListHostGroup) SetNodeCount(v int32) *DescribeClusterTemplateResponseBodyTemplateInfoHostGroupListHostGroup {
	s.NodeCount = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfoHostGroupListHostGroup) SetHostGroupName(v string) *DescribeClusterTemplateResponseBodyTemplateInfoHostGroupListHostGroup {
	s.HostGroupName = &v
	return s
}

type DescribeClusterTemplateResponseBodyTemplateInfoConfigList struct {
	Config []*DescribeClusterTemplateResponseBodyTemplateInfoConfigListConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Repeated"`
}

func (s DescribeClusterTemplateResponseBodyTemplateInfoConfigList) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterTemplateResponseBodyTemplateInfoConfigList) GoString() string {
	return s.String()
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfoConfigList) SetConfig(v []*DescribeClusterTemplateResponseBodyTemplateInfoConfigListConfig) *DescribeClusterTemplateResponseBodyTemplateInfoConfigList {
	s.Config = v
	return s
}

type DescribeClusterTemplateResponseBodyTemplateInfoConfigListConfig struct {
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	Replace     *string `json:"Replace,omitempty" xml:"Replace,omitempty"`
	FileName    *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ConfigKey   *string `json:"ConfigKey,omitempty" xml:"ConfigKey,omitempty"`
	Encrypt     *string `json:"Encrypt,omitempty" xml:"Encrypt,omitempty"`
}

func (s DescribeClusterTemplateResponseBodyTemplateInfoConfigListConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterTemplateResponseBodyTemplateInfoConfigListConfig) GoString() string {
	return s.String()
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfoConfigListConfig) SetConfigValue(v string) *DescribeClusterTemplateResponseBodyTemplateInfoConfigListConfig {
	s.ConfigValue = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfoConfigListConfig) SetReplace(v string) *DescribeClusterTemplateResponseBodyTemplateInfoConfigListConfig {
	s.Replace = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfoConfigListConfig) SetFileName(v string) *DescribeClusterTemplateResponseBodyTemplateInfoConfigListConfig {
	s.FileName = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfoConfigListConfig) SetServiceName(v string) *DescribeClusterTemplateResponseBodyTemplateInfoConfigListConfig {
	s.ServiceName = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfoConfigListConfig) SetConfigKey(v string) *DescribeClusterTemplateResponseBodyTemplateInfoConfigListConfig {
	s.ConfigKey = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfoConfigListConfig) SetEncrypt(v string) *DescribeClusterTemplateResponseBodyTemplateInfoConfigListConfig {
	s.Encrypt = &v
	return s
}

type DescribeClusterTemplateResponseBodyTemplateInfoTags struct {
	Tag []*DescribeClusterTemplateResponseBodyTemplateInfoTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeClusterTemplateResponseBodyTemplateInfoTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterTemplateResponseBodyTemplateInfoTags) GoString() string {
	return s.String()
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfoTags) SetTag(v []*DescribeClusterTemplateResponseBodyTemplateInfoTagsTag) *DescribeClusterTemplateResponseBodyTemplateInfoTags {
	s.Tag = v
	return s
}

type DescribeClusterTemplateResponseBodyTemplateInfoTagsTag struct {
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s DescribeClusterTemplateResponseBodyTemplateInfoTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterTemplateResponseBodyTemplateInfoTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfoTagsTag) SetTagValue(v string) *DescribeClusterTemplateResponseBodyTemplateInfoTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfoTagsTag) SetTagKey(v string) *DescribeClusterTemplateResponseBodyTemplateInfoTagsTag {
	s.TagKey = &v
	return s
}

type DescribeClusterTemplateResponseBodyTemplateInfoSoftwareInfoList struct {
	SoftwareInfo []*string `json:"SoftwareInfo,omitempty" xml:"SoftwareInfo,omitempty" type:"Repeated"`
}

func (s DescribeClusterTemplateResponseBodyTemplateInfoSoftwareInfoList) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterTemplateResponseBodyTemplateInfoSoftwareInfoList) GoString() string {
	return s.String()
}

func (s *DescribeClusterTemplateResponseBodyTemplateInfoSoftwareInfoList) SetSoftwareInfo(v []*string) *DescribeClusterTemplateResponseBodyTemplateInfoSoftwareInfoList {
	s.SoftwareInfo = v
	return s
}

type DescribeClusterTemplateResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClusterTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClusterTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterTemplateResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterTemplateResponse) SetHeaders(v map[string]*string) *DescribeClusterTemplateResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterTemplateResponse) SetBody(v *DescribeClusterTemplateResponseBody) *DescribeClusterTemplateResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// 资源类型
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// 资源ID
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// 标签列表
	Tags []*TagResourcesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetResourceOwnerId(v int64) *TagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceGroupId(v string) *TagResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetResourceIds(v []*string) *TagResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *TagResourcesRequest) SetTags(v []*TagResourcesRequestTags) *TagResourcesRequest {
	s.Tags = v
	return s
}

type TagResourcesRequestTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTags) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTags) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTags) SetKey(v string) *TagResourcesRequestTags {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTags) SetValue(v string) *TagResourcesRequestTags {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 请求是否成功被处理
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 响应码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 错误码
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 响应消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *TagResourcesResponseBody) SetSuccess(v bool) *TagResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *TagResourcesResponseBody) SetCode(v string) *TagResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *TagResourcesResponseBody) SetErrorCode(v string) *TagResourcesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TagResourcesResponseBody) SetMessage(v string) *TagResourcesResponseBody {
	s.Message = &v
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

type CommitFlowEntitySnapshotRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	EntityType      *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	EntityId        *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	Message         *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s CommitFlowEntitySnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s CommitFlowEntitySnapshotRequest) GoString() string {
	return s.String()
}

func (s *CommitFlowEntitySnapshotRequest) SetResourceOwnerId(v int64) *CommitFlowEntitySnapshotRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CommitFlowEntitySnapshotRequest) SetRegionId(v string) *CommitFlowEntitySnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *CommitFlowEntitySnapshotRequest) SetEntityType(v string) *CommitFlowEntitySnapshotRequest {
	s.EntityType = &v
	return s
}

func (s *CommitFlowEntitySnapshotRequest) SetEntityId(v string) *CommitFlowEntitySnapshotRequest {
	s.EntityId = &v
	return s
}

func (s *CommitFlowEntitySnapshotRequest) SetMessage(v string) *CommitFlowEntitySnapshotRequest {
	s.Message = &v
	return s
}

type CommitFlowEntitySnapshotResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CommitFlowEntitySnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CommitFlowEntitySnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *CommitFlowEntitySnapshotResponseBody) SetData(v string) *CommitFlowEntitySnapshotResponseBody {
	s.Data = &v
	return s
}

func (s *CommitFlowEntitySnapshotResponseBody) SetRequestId(v string) *CommitFlowEntitySnapshotResponseBody {
	s.RequestId = &v
	return s
}

type CommitFlowEntitySnapshotResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CommitFlowEntitySnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CommitFlowEntitySnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s CommitFlowEntitySnapshotResponse) GoString() string {
	return s.String()
}

func (s *CommitFlowEntitySnapshotResponse) SetHeaders(v map[string]*string) *CommitFlowEntitySnapshotResponse {
	s.Headers = v
	return s
}

func (s *CommitFlowEntitySnapshotResponse) SetBody(v *CommitFlowEntitySnapshotResponseBody) *CommitFlowEntitySnapshotResponse {
	s.Body = v
	return s
}

type SubmitFlowJobRequest struct {
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	HostName  *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	Conf      *string `json:"Conf,omitempty" xml:"Conf,omitempty"`
}

func (s SubmitFlowJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitFlowJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitFlowJobRequest) SetRegionId(v string) *SubmitFlowJobRequest {
	s.RegionId = &v
	return s
}

func (s *SubmitFlowJobRequest) SetProjectId(v string) *SubmitFlowJobRequest {
	s.ProjectId = &v
	return s
}

func (s *SubmitFlowJobRequest) SetJobId(v string) *SubmitFlowJobRequest {
	s.JobId = &v
	return s
}

func (s *SubmitFlowJobRequest) SetClusterId(v string) *SubmitFlowJobRequest {
	s.ClusterId = &v
	return s
}

func (s *SubmitFlowJobRequest) SetHostName(v string) *SubmitFlowJobRequest {
	s.HostName = &v
	return s
}

func (s *SubmitFlowJobRequest) SetConf(v string) *SubmitFlowJobRequest {
	s.Conf = &v
	return s
}

type SubmitFlowJobResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitFlowJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitFlowJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitFlowJobResponseBody) SetId(v string) *SubmitFlowJobResponseBody {
	s.Id = &v
	return s
}

func (s *SubmitFlowJobResponseBody) SetRequestId(v string) *SubmitFlowJobResponseBody {
	s.RequestId = &v
	return s
}

type SubmitFlowJobResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitFlowJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitFlowJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitFlowJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitFlowJobResponse) SetHeaders(v map[string]*string) *SubmitFlowJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitFlowJobResponse) SetBody(v *SubmitFlowJobResponseBody) *SubmitFlowJobResponse {
	s.Body = v
	return s
}

type ListFlowJobHistoryRequest struct {
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProjectId  *string   `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Id         *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	JobType    *string   `json:"JobType,omitempty" xml:"JobType,omitempty"`
	InstanceId *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TimeRange  *string   `json:"TimeRange,omitempty" xml:"TimeRange,omitempty"`
	PageNumber *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StatusList []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
}

func (s ListFlowJobHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowJobHistoryRequest) GoString() string {
	return s.String()
}

func (s *ListFlowJobHistoryRequest) SetRegionId(v string) *ListFlowJobHistoryRequest {
	s.RegionId = &v
	return s
}

func (s *ListFlowJobHistoryRequest) SetProjectId(v string) *ListFlowJobHistoryRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFlowJobHistoryRequest) SetId(v string) *ListFlowJobHistoryRequest {
	s.Id = &v
	return s
}

func (s *ListFlowJobHistoryRequest) SetJobType(v string) *ListFlowJobHistoryRequest {
	s.JobType = &v
	return s
}

func (s *ListFlowJobHistoryRequest) SetInstanceId(v string) *ListFlowJobHistoryRequest {
	s.InstanceId = &v
	return s
}

func (s *ListFlowJobHistoryRequest) SetTimeRange(v string) *ListFlowJobHistoryRequest {
	s.TimeRange = &v
	return s
}

func (s *ListFlowJobHistoryRequest) SetPageNumber(v int32) *ListFlowJobHistoryRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFlowJobHistoryRequest) SetPageSize(v int32) *ListFlowJobHistoryRequest {
	s.PageSize = &v
	return s
}

func (s *ListFlowJobHistoryRequest) SetStatusList(v []*string) *ListFlowJobHistoryRequest {
	s.StatusList = v
	return s
}

type ListFlowJobHistoryResponseBody struct {
	RequestId     *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber    *int32                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total         *int32                                       `json:"Total,omitempty" xml:"Total,omitempty"`
	NodeInstances *ListFlowJobHistoryResponseBodyNodeInstances `json:"NodeInstances,omitempty" xml:"NodeInstances,omitempty" type:"Struct"`
}

func (s ListFlowJobHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowJobHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowJobHistoryResponseBody) SetRequestId(v string) *ListFlowJobHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowJobHistoryResponseBody) SetPageNumber(v int32) *ListFlowJobHistoryResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListFlowJobHistoryResponseBody) SetPageSize(v int32) *ListFlowJobHistoryResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListFlowJobHistoryResponseBody) SetTotal(v int32) *ListFlowJobHistoryResponseBody {
	s.Total = &v
	return s
}

func (s *ListFlowJobHistoryResponseBody) SetNodeInstances(v *ListFlowJobHistoryResponseBodyNodeInstances) *ListFlowJobHistoryResponseBody {
	s.NodeInstances = v
	return s
}

type ListFlowJobHistoryResponseBodyNodeInstances struct {
	NodeInstance []*ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance `json:"NodeInstance,omitempty" xml:"NodeInstance,omitempty" type:"Repeated"`
}

func (s ListFlowJobHistoryResponseBodyNodeInstances) String() string {
	return tea.Prettify(s)
}

func (s ListFlowJobHistoryResponseBodyNodeInstances) GoString() string {
	return s.String()
}

func (s *ListFlowJobHistoryResponseBodyNodeInstances) SetNodeInstance(v []*ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) *ListFlowJobHistoryResponseBodyNodeInstances {
	s.NodeInstance = v
	return s
}

type ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance struct {
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Pending        *bool   `json:"pending,omitempty" xml:"pending,omitempty"`
	EnvConf        *string `json:"EnvConf,omitempty" xml:"EnvConf,omitempty"`
	RetryInterval  *int64  `json:"RetryInterval,omitempty" xml:"RetryInterval,omitempty"`
	ProjectId      *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	JobType        *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	GmtModified    *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	ExternalInfo   *string `json:"ExternalInfo,omitempty" xml:"ExternalInfo,omitempty"`
	ExternalStatus *string `json:"ExternalStatus,omitempty" xml:"ExternalStatus,omitempty"`
	JobName        *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	ExternalId     *string `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	MaxRetry       *int32  `json:"MaxRetry,omitempty" xml:"MaxRetry,omitempty"`
	ClusterId      *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	FailAct        *string `json:"FailAct,omitempty" xml:"FailAct,omitempty"`
	JobParams      *string `json:"JobParams,omitempty" xml:"JobParams,omitempty"`
	ParamConf      *string `json:"ParamConf,omitempty" xml:"ParamConf,omitempty"`
	HostName       *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	Retries        *int32  `json:"Retries,omitempty" xml:"Retries,omitempty"`
	EndTime        *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime      *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	RunConf        *string `json:"RunConf,omitempty" xml:"RunConf,omitempty"`
	NodeName       *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	JobId          *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	GmtCreate      *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) String() string {
	return tea.Prettify(s)
}

func (s ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) GoString() string {
	return s.String()
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetType(v string) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.Type = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetStatus(v string) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.Status = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetPending(v bool) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.Pending = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetEnvConf(v string) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.EnvConf = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetRetryInterval(v int64) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.RetryInterval = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetProjectId(v string) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.ProjectId = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetJobType(v string) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.JobType = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetGmtModified(v int64) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.GmtModified = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetExternalInfo(v string) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.ExternalInfo = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetExternalStatus(v string) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.ExternalStatus = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetJobName(v string) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.JobName = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetExternalId(v string) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.ExternalId = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetMaxRetry(v int32) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.MaxRetry = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetClusterId(v string) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.ClusterId = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetFailAct(v string) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.FailAct = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetJobParams(v string) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.JobParams = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetParamConf(v string) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.ParamConf = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetHostName(v string) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.HostName = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetRetries(v int32) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.Retries = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetEndTime(v int64) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.EndTime = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetStartTime(v int64) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.StartTime = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetRunConf(v string) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.RunConf = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetNodeName(v string) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.NodeName = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetJobId(v string) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.JobId = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetGmtCreate(v int64) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.GmtCreate = &v
	return s
}

func (s *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance) SetId(v string) *ListFlowJobHistoryResponseBodyNodeInstancesNodeInstance {
	s.Id = &v
	return s
}

type ListFlowJobHistoryResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFlowJobHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowJobHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowJobHistoryResponse) GoString() string {
	return s.String()
}

func (s *ListFlowJobHistoryResponse) SetHeaders(v map[string]*string) *ListFlowJobHistoryResponse {
	s.Headers = v
	return s
}

func (s *ListFlowJobHistoryResponse) SetBody(v *ListFlowJobHistoryResponseBody) *ListFlowJobHistoryResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceType    *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// 资源组ID
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// 标签
	Tags      []*ListTagResourcesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	NextToken *string                        `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetResourceOwnerId(v int64) *ListTagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceGroupId(v string) *ListTagResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceIds(v []*string) *ListTagResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *ListTagResourcesRequest) SetTags(v []*ListTagResourcesRequestTags) *ListTagResourcesRequest {
	s.Tags = v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

type ListTagResourcesRequestTags struct {
	// 标签键
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 标签值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTags) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTags) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTags) SetKey(v string) *ListTagResourcesRequestTags {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTags) SetValue(v string) *ListTagResourcesRequestTags {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 请求是否成功被处理
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 响应码
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 错误码
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 响应消息
	Message *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	Data    []*ListTagResourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetSuccess(v bool) *ListTagResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetCode(v string) *ListTagResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetErrorCode(v string) *ListTagResourcesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetMessage(v string) *ListTagResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetData(v []*ListTagResourcesResponseBodyData) *ListTagResourcesResponseBody {
	s.Data = v
	return s
}

type ListTagResourcesResponseBodyData struct {
	// 标签键
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// 标签值
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	// 资源类型
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// 资源ID
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s ListTagResourcesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyData) SetTagKey(v string) *ListTagResourcesResponseBodyData {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyData) SetTagValue(v string) *ListTagResourcesResponseBodyData {
	s.TagValue = &v
	return s
}

func (s *ListTagResourcesResponseBodyData) SetResourceType(v string) *ListTagResourcesResponseBodyData {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyData) SetResourceId(v string) *ListTagResourcesResponseBodyData {
	s.ResourceId = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ListClusterHostComponentRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	HostInstanceId  *string `json:"HostInstanceId,omitempty" xml:"HostInstanceId,omitempty"`
	HostName        *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	ServiceName     *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ComponentName   *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	ComponentStatus *string `json:"ComponentStatus,omitempty" xml:"ComponentStatus,omitempty"`
	HostRole        *string `json:"HostRole,omitempty" xml:"HostRole,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListClusterHostComponentRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClusterHostComponentRequest) GoString() string {
	return s.String()
}

func (s *ListClusterHostComponentRequest) SetResourceOwnerId(v int64) *ListClusterHostComponentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListClusterHostComponentRequest) SetRegionId(v string) *ListClusterHostComponentRequest {
	s.RegionId = &v
	return s
}

func (s *ListClusterHostComponentRequest) SetClusterId(v string) *ListClusterHostComponentRequest {
	s.ClusterId = &v
	return s
}

func (s *ListClusterHostComponentRequest) SetHostInstanceId(v string) *ListClusterHostComponentRequest {
	s.HostInstanceId = &v
	return s
}

func (s *ListClusterHostComponentRequest) SetHostName(v string) *ListClusterHostComponentRequest {
	s.HostName = &v
	return s
}

func (s *ListClusterHostComponentRequest) SetServiceName(v string) *ListClusterHostComponentRequest {
	s.ServiceName = &v
	return s
}

func (s *ListClusterHostComponentRequest) SetComponentName(v string) *ListClusterHostComponentRequest {
	s.ComponentName = &v
	return s
}

func (s *ListClusterHostComponentRequest) SetComponentStatus(v string) *ListClusterHostComponentRequest {
	s.ComponentStatus = &v
	return s
}

func (s *ListClusterHostComponentRequest) SetHostRole(v string) *ListClusterHostComponentRequest {
	s.HostRole = &v
	return s
}

func (s *ListClusterHostComponentRequest) SetPageNumber(v int32) *ListClusterHostComponentRequest {
	s.PageNumber = &v
	return s
}

func (s *ListClusterHostComponentRequest) SetPageSize(v int32) *ListClusterHostComponentRequest {
	s.PageSize = &v
	return s
}

type ListClusterHostComponentResponseBody struct {
	RequestId     *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber    *int32                                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total         *int32                                             `json:"Total,omitempty" xml:"Total,omitempty"`
	ComponentList *ListClusterHostComponentResponseBodyComponentList `json:"ComponentList,omitempty" xml:"ComponentList,omitempty" type:"Struct"`
}

func (s ListClusterHostComponentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClusterHostComponentResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterHostComponentResponseBody) SetRequestId(v string) *ListClusterHostComponentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterHostComponentResponseBody) SetPageNumber(v int32) *ListClusterHostComponentResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListClusterHostComponentResponseBody) SetPageSize(v int32) *ListClusterHostComponentResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListClusterHostComponentResponseBody) SetTotal(v int32) *ListClusterHostComponentResponseBody {
	s.Total = &v
	return s
}

func (s *ListClusterHostComponentResponseBody) SetComponentList(v *ListClusterHostComponentResponseBodyComponentList) *ListClusterHostComponentResponseBody {
	s.ComponentList = v
	return s
}

type ListClusterHostComponentResponseBodyComponentList struct {
	Component []*ListClusterHostComponentResponseBodyComponentListComponent `json:"Component,omitempty" xml:"Component,omitempty" type:"Repeated"`
}

func (s ListClusterHostComponentResponseBodyComponentList) String() string {
	return tea.Prettify(s)
}

func (s ListClusterHostComponentResponseBodyComponentList) GoString() string {
	return s.String()
}

func (s *ListClusterHostComponentResponseBodyComponentList) SetComponent(v []*ListClusterHostComponentResponseBodyComponentListComponent) *ListClusterHostComponentResponseBodyComponentList {
	s.Component = v
	return s
}

type ListClusterHostComponentResponseBodyComponentListComponent struct {
	SerialNumber         *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ServiceDisplayName   *string `json:"ServiceDisplayName,omitempty" xml:"ServiceDisplayName,omitempty"`
	PrivateIp            *string `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	ServerStatus         *string `json:"ServerStatus,omitempty" xml:"ServerStatus,omitempty"`
	ComponentName        *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	CommissionStatus     *string `json:"CommissionStatus,omitempty" xml:"CommissionStatus,omitempty"`
	HostName             *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	InstanceType         *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	NeedRestart          *bool   `json:"NeedRestart,omitempty" xml:"NeedRestart,omitempty"`
	HostId               *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	HostInstanceId       *string `json:"HostInstanceId,omitempty" xml:"HostInstanceId,omitempty"`
	Cpu                  *int32  `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	ComponentDisplayName *string `json:"ComponentDisplayName,omitempty" xml:"ComponentDisplayName,omitempty"`
	PublicIp             *string `json:"PublicIp,omitempty" xml:"PublicIp,omitempty"`
	Memory               *int32  `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Role                 *string `json:"Role,omitempty" xml:"Role,omitempty"`
	ServiceName          *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s ListClusterHostComponentResponseBodyComponentListComponent) String() string {
	return tea.Prettify(s)
}

func (s ListClusterHostComponentResponseBodyComponentListComponent) GoString() string {
	return s.String()
}

func (s *ListClusterHostComponentResponseBodyComponentListComponent) SetSerialNumber(v string) *ListClusterHostComponentResponseBodyComponentListComponent {
	s.SerialNumber = &v
	return s
}

func (s *ListClusterHostComponentResponseBodyComponentListComponent) SetStatus(v string) *ListClusterHostComponentResponseBodyComponentListComponent {
	s.Status = &v
	return s
}

func (s *ListClusterHostComponentResponseBodyComponentListComponent) SetServiceDisplayName(v string) *ListClusterHostComponentResponseBodyComponentListComponent {
	s.ServiceDisplayName = &v
	return s
}

func (s *ListClusterHostComponentResponseBodyComponentListComponent) SetPrivateIp(v string) *ListClusterHostComponentResponseBodyComponentListComponent {
	s.PrivateIp = &v
	return s
}

func (s *ListClusterHostComponentResponseBodyComponentListComponent) SetServerStatus(v string) *ListClusterHostComponentResponseBodyComponentListComponent {
	s.ServerStatus = &v
	return s
}

func (s *ListClusterHostComponentResponseBodyComponentListComponent) SetComponentName(v string) *ListClusterHostComponentResponseBodyComponentListComponent {
	s.ComponentName = &v
	return s
}

func (s *ListClusterHostComponentResponseBodyComponentListComponent) SetCommissionStatus(v string) *ListClusterHostComponentResponseBodyComponentListComponent {
	s.CommissionStatus = &v
	return s
}

func (s *ListClusterHostComponentResponseBodyComponentListComponent) SetHostName(v string) *ListClusterHostComponentResponseBodyComponentListComponent {
	s.HostName = &v
	return s
}

func (s *ListClusterHostComponentResponseBodyComponentListComponent) SetInstanceType(v string) *ListClusterHostComponentResponseBodyComponentListComponent {
	s.InstanceType = &v
	return s
}

func (s *ListClusterHostComponentResponseBodyComponentListComponent) SetNeedRestart(v bool) *ListClusterHostComponentResponseBodyComponentListComponent {
	s.NeedRestart = &v
	return s
}

func (s *ListClusterHostComponentResponseBodyComponentListComponent) SetHostId(v string) *ListClusterHostComponentResponseBodyComponentListComponent {
	s.HostId = &v
	return s
}

func (s *ListClusterHostComponentResponseBodyComponentListComponent) SetHostInstanceId(v string) *ListClusterHostComponentResponseBodyComponentListComponent {
	s.HostInstanceId = &v
	return s
}

func (s *ListClusterHostComponentResponseBodyComponentListComponent) SetCpu(v int32) *ListClusterHostComponentResponseBodyComponentListComponent {
	s.Cpu = &v
	return s
}

func (s *ListClusterHostComponentResponseBodyComponentListComponent) SetComponentDisplayName(v string) *ListClusterHostComponentResponseBodyComponentListComponent {
	s.ComponentDisplayName = &v
	return s
}

func (s *ListClusterHostComponentResponseBodyComponentListComponent) SetPublicIp(v string) *ListClusterHostComponentResponseBodyComponentListComponent {
	s.PublicIp = &v
	return s
}

func (s *ListClusterHostComponentResponseBodyComponentListComponent) SetMemory(v int32) *ListClusterHostComponentResponseBodyComponentListComponent {
	s.Memory = &v
	return s
}

func (s *ListClusterHostComponentResponseBodyComponentListComponent) SetRole(v string) *ListClusterHostComponentResponseBodyComponentListComponent {
	s.Role = &v
	return s
}

func (s *ListClusterHostComponentResponseBodyComponentListComponent) SetServiceName(v string) *ListClusterHostComponentResponseBodyComponentListComponent {
	s.ServiceName = &v
	return s
}

type ListClusterHostComponentResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListClusterHostComponentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClusterHostComponentResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClusterHostComponentResponse) GoString() string {
	return s.String()
}

func (s *ListClusterHostComponentResponse) SetHeaders(v map[string]*string) *ListClusterHostComponentResponse {
	s.Headers = v
	return s
}

func (s *ListClusterHostComponentResponse) SetBody(v *ListClusterHostComponentResponseBody) *ListClusterHostComponentResponse {
	s.Body = v
	return s
}

type ModifyScalingGroupRequest struct {
	ResourceOwnerId   *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId   *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ScalingGroupBizId *string `json:"ScalingGroupBizId,omitempty" xml:"ScalingGroupBizId,omitempty"`
}

func (s ModifyScalingGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyScalingGroupRequest) SetResourceOwnerId(v int64) *ModifyScalingGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetRegionId(v string) *ModifyScalingGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetResourceGroupId(v string) *ModifyScalingGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetName(v string) *ModifyScalingGroupRequest {
	s.Name = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetDescription(v string) *ModifyScalingGroupRequest {
	s.Description = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetScalingGroupBizId(v string) *ModifyScalingGroupRequest {
	s.ScalingGroupBizId = &v
	return s
}

type ModifyScalingGroupResponseBody struct {
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// data
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s ModifyScalingGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyScalingGroupResponseBody) SetRequestId(v string) *ModifyScalingGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyScalingGroupResponseBody) SetData(v bool) *ModifyScalingGroupResponseBody {
	s.Data = &v
	return s
}

type ModifyScalingGroupResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyScalingGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyScalingGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyScalingGroupResponse) SetHeaders(v map[string]*string) *ModifyScalingGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyScalingGroupResponse) SetBody(v *ModifyScalingGroupResponseBody) *ModifyScalingGroupResponse {
	s.Body = v
	return s
}

type DescribeFlowProjectClusterSettingRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeFlowProjectClusterSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowProjectClusterSettingRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowProjectClusterSettingRequest) SetProjectId(v string) *DescribeFlowProjectClusterSettingRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeFlowProjectClusterSettingRequest) SetRegionId(v string) *DescribeFlowProjectClusterSettingRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFlowProjectClusterSettingRequest) SetClusterId(v string) *DescribeFlowProjectClusterSettingRequest {
	s.ClusterId = &v
	return s
}

type DescribeFlowProjectClusterSettingResponseBody struct {
	RequestId    *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	K8sClusterId *string                                                 `json:"K8sClusterId,omitempty" xml:"K8sClusterId,omitempty"`
	DefaultQueue *string                                                 `json:"DefaultQueue,omitempty" xml:"DefaultQueue,omitempty"`
	ProjectId    *string                                                 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	GmtCreate    *int64                                                  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	DefaultUser  *string                                                 `json:"DefaultUser,omitempty" xml:"DefaultUser,omitempty"`
	GmtModified  *int64                                                  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	ClusterId    *string                                                 `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	QueueList    *DescribeFlowProjectClusterSettingResponseBodyQueueList `json:"QueueList,omitempty" xml:"QueueList,omitempty" type:"Struct"`
	UserList     *DescribeFlowProjectClusterSettingResponseBodyUserList  `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Struct"`
	HostList     *DescribeFlowProjectClusterSettingResponseBodyHostList  `json:"HostList,omitempty" xml:"HostList,omitempty" type:"Struct"`
}

func (s DescribeFlowProjectClusterSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowProjectClusterSettingResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowProjectClusterSettingResponseBody) SetRequestId(v string) *DescribeFlowProjectClusterSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowProjectClusterSettingResponseBody) SetK8sClusterId(v string) *DescribeFlowProjectClusterSettingResponseBody {
	s.K8sClusterId = &v
	return s
}

func (s *DescribeFlowProjectClusterSettingResponseBody) SetDefaultQueue(v string) *DescribeFlowProjectClusterSettingResponseBody {
	s.DefaultQueue = &v
	return s
}

func (s *DescribeFlowProjectClusterSettingResponseBody) SetProjectId(v string) *DescribeFlowProjectClusterSettingResponseBody {
	s.ProjectId = &v
	return s
}

func (s *DescribeFlowProjectClusterSettingResponseBody) SetGmtCreate(v int64) *DescribeFlowProjectClusterSettingResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *DescribeFlowProjectClusterSettingResponseBody) SetDefaultUser(v string) *DescribeFlowProjectClusterSettingResponseBody {
	s.DefaultUser = &v
	return s
}

func (s *DescribeFlowProjectClusterSettingResponseBody) SetGmtModified(v int64) *DescribeFlowProjectClusterSettingResponseBody {
	s.GmtModified = &v
	return s
}

func (s *DescribeFlowProjectClusterSettingResponseBody) SetClusterId(v string) *DescribeFlowProjectClusterSettingResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeFlowProjectClusterSettingResponseBody) SetQueueList(v *DescribeFlowProjectClusterSettingResponseBodyQueueList) *DescribeFlowProjectClusterSettingResponseBody {
	s.QueueList = v
	return s
}

func (s *DescribeFlowProjectClusterSettingResponseBody) SetUserList(v *DescribeFlowProjectClusterSettingResponseBodyUserList) *DescribeFlowProjectClusterSettingResponseBody {
	s.UserList = v
	return s
}

func (s *DescribeFlowProjectClusterSettingResponseBody) SetHostList(v *DescribeFlowProjectClusterSettingResponseBodyHostList) *DescribeFlowProjectClusterSettingResponseBody {
	s.HostList = v
	return s
}

type DescribeFlowProjectClusterSettingResponseBodyQueueList struct {
	Queue []*string `json:"Queue,omitempty" xml:"Queue,omitempty" type:"Repeated"`
}

func (s DescribeFlowProjectClusterSettingResponseBodyQueueList) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowProjectClusterSettingResponseBodyQueueList) GoString() string {
	return s.String()
}

func (s *DescribeFlowProjectClusterSettingResponseBodyQueueList) SetQueue(v []*string) *DescribeFlowProjectClusterSettingResponseBodyQueueList {
	s.Queue = v
	return s
}

type DescribeFlowProjectClusterSettingResponseBodyUserList struct {
	User []*string `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s DescribeFlowProjectClusterSettingResponseBodyUserList) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowProjectClusterSettingResponseBodyUserList) GoString() string {
	return s.String()
}

func (s *DescribeFlowProjectClusterSettingResponseBodyUserList) SetUser(v []*string) *DescribeFlowProjectClusterSettingResponseBodyUserList {
	s.User = v
	return s
}

type DescribeFlowProjectClusterSettingResponseBodyHostList struct {
	Host []*string `json:"Host,omitempty" xml:"Host,omitempty" type:"Repeated"`
}

func (s DescribeFlowProjectClusterSettingResponseBodyHostList) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowProjectClusterSettingResponseBodyHostList) GoString() string {
	return s.String()
}

func (s *DescribeFlowProjectClusterSettingResponseBodyHostList) SetHost(v []*string) *DescribeFlowProjectClusterSettingResponseBodyHostList {
	s.Host = v
	return s
}

type DescribeFlowProjectClusterSettingResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFlowProjectClusterSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFlowProjectClusterSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowProjectClusterSettingResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowProjectClusterSettingResponse) SetHeaders(v map[string]*string) *DescribeFlowProjectClusterSettingResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowProjectClusterSettingResponse) SetBody(v *DescribeFlowProjectClusterSettingResponseBody) *DescribeFlowProjectClusterSettingResponse {
	s.Body = v
	return s
}

type ListFlowProjectClusterSettingRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProjectId  *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListFlowProjectClusterSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowProjectClusterSettingRequest) GoString() string {
	return s.String()
}

func (s *ListFlowProjectClusterSettingRequest) SetRegionId(v string) *ListFlowProjectClusterSettingRequest {
	s.RegionId = &v
	return s
}

func (s *ListFlowProjectClusterSettingRequest) SetProjectId(v string) *ListFlowProjectClusterSettingRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFlowProjectClusterSettingRequest) SetPageNumber(v int32) *ListFlowProjectClusterSettingRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFlowProjectClusterSettingRequest) SetPageSize(v int32) *ListFlowProjectClusterSettingRequest {
	s.PageSize = &v
	return s
}

type ListFlowProjectClusterSettingResponseBody struct {
	RequestId       *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber      *int32                                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32                                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total           *int32                                                    `json:"Total,omitempty" xml:"Total,omitempty"`
	ClusterSettings *ListFlowProjectClusterSettingResponseBodyClusterSettings `json:"ClusterSettings,omitempty" xml:"ClusterSettings,omitempty" type:"Struct"`
}

func (s ListFlowProjectClusterSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowProjectClusterSettingResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowProjectClusterSettingResponseBody) SetRequestId(v string) *ListFlowProjectClusterSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowProjectClusterSettingResponseBody) SetPageNumber(v int32) *ListFlowProjectClusterSettingResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListFlowProjectClusterSettingResponseBody) SetPageSize(v int32) *ListFlowProjectClusterSettingResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListFlowProjectClusterSettingResponseBody) SetTotal(v int32) *ListFlowProjectClusterSettingResponseBody {
	s.Total = &v
	return s
}

func (s *ListFlowProjectClusterSettingResponseBody) SetClusterSettings(v *ListFlowProjectClusterSettingResponseBodyClusterSettings) *ListFlowProjectClusterSettingResponseBody {
	s.ClusterSettings = v
	return s
}

type ListFlowProjectClusterSettingResponseBodyClusterSettings struct {
	ClusterSetting []*ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting `json:"ClusterSetting,omitempty" xml:"ClusterSetting,omitempty" type:"Repeated"`
}

func (s ListFlowProjectClusterSettingResponseBodyClusterSettings) String() string {
	return tea.Prettify(s)
}

func (s ListFlowProjectClusterSettingResponseBodyClusterSettings) GoString() string {
	return s.String()
}

func (s *ListFlowProjectClusterSettingResponseBodyClusterSettings) SetClusterSetting(v []*ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting) *ListFlowProjectClusterSettingResponseBodyClusterSettings {
	s.ClusterSetting = v
	return s
}

type ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting struct {
	K8sClusterId *string                                                                          `json:"K8sClusterId,omitempty" xml:"K8sClusterId,omitempty"`
	DefaultQueue *string                                                                          `json:"DefaultQueue,omitempty" xml:"DefaultQueue,omitempty"`
	ProjectId    *string                                                                          `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	GmtCreate    *int64                                                                           `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	ClusterName  *string                                                                          `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	DefaultUser  *string                                                                          `json:"DefaultUser,omitempty" xml:"DefaultUser,omitempty"`
	GmtModified  *int64                                                                           `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	ClusterId    *string                                                                          `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	QueueList    *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingQueueList `json:"QueueList,omitempty" xml:"QueueList,omitempty" type:"Struct"`
	UserList     *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingUserList  `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Struct"`
	HostList     *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingHostList  `json:"HostList,omitempty" xml:"HostList,omitempty" type:"Struct"`
}

func (s ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting) String() string {
	return tea.Prettify(s)
}

func (s ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting) GoString() string {
	return s.String()
}

func (s *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting) SetK8sClusterId(v string) *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting {
	s.K8sClusterId = &v
	return s
}

func (s *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting) SetDefaultQueue(v string) *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting {
	s.DefaultQueue = &v
	return s
}

func (s *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting) SetProjectId(v string) *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting {
	s.ProjectId = &v
	return s
}

func (s *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting) SetGmtCreate(v int64) *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting {
	s.GmtCreate = &v
	return s
}

func (s *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting) SetClusterName(v string) *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting {
	s.ClusterName = &v
	return s
}

func (s *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting) SetDefaultUser(v string) *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting {
	s.DefaultUser = &v
	return s
}

func (s *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting) SetGmtModified(v int64) *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting {
	s.GmtModified = &v
	return s
}

func (s *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting) SetClusterId(v string) *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting {
	s.ClusterId = &v
	return s
}

func (s *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting) SetQueueList(v *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingQueueList) *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting {
	s.QueueList = v
	return s
}

func (s *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting) SetUserList(v *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingUserList) *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting {
	s.UserList = v
	return s
}

func (s *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting) SetHostList(v *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingHostList) *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSetting {
	s.HostList = v
	return s
}

type ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingQueueList struct {
	Queue []*string `json:"Queue,omitempty" xml:"Queue,omitempty" type:"Repeated"`
}

func (s ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingQueueList) String() string {
	return tea.Prettify(s)
}

func (s ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingQueueList) GoString() string {
	return s.String()
}

func (s *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingQueueList) SetQueue(v []*string) *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingQueueList {
	s.Queue = v
	return s
}

type ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingUserList struct {
	User []*string `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingUserList) String() string {
	return tea.Prettify(s)
}

func (s ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingUserList) GoString() string {
	return s.String()
}

func (s *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingUserList) SetUser(v []*string) *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingUserList {
	s.User = v
	return s
}

type ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingHostList struct {
	Host []*string `json:"Host,omitempty" xml:"Host,omitempty" type:"Repeated"`
}

func (s ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingHostList) String() string {
	return tea.Prettify(s)
}

func (s ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingHostList) GoString() string {
	return s.String()
}

func (s *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingHostList) SetHost(v []*string) *ListFlowProjectClusterSettingResponseBodyClusterSettingsClusterSettingHostList {
	s.Host = v
	return s
}

type ListFlowProjectClusterSettingResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFlowProjectClusterSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowProjectClusterSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowProjectClusterSettingResponse) GoString() string {
	return s.String()
}

func (s *ListFlowProjectClusterSettingResponse) SetHeaders(v map[string]*string) *ListFlowProjectClusterSettingResponse {
	s.Headers = v
	return s
}

func (s *ListFlowProjectClusterSettingResponse) SetBody(v *ListFlowProjectClusterSettingResponseBody) *ListFlowProjectClusterSettingResponse {
	s.Body = v
	return s
}

type SubmitFlowRequest struct {
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	FlowId    *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	Conf      *string `json:"Conf,omitempty" xml:"Conf,omitempty"`
}

func (s SubmitFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitFlowRequest) GoString() string {
	return s.String()
}

func (s *SubmitFlowRequest) SetRegionId(v string) *SubmitFlowRequest {
	s.RegionId = &v
	return s
}

func (s *SubmitFlowRequest) SetProjectId(v string) *SubmitFlowRequest {
	s.ProjectId = &v
	return s
}

func (s *SubmitFlowRequest) SetFlowId(v string) *SubmitFlowRequest {
	s.FlowId = &v
	return s
}

func (s *SubmitFlowRequest) SetConf(v string) *SubmitFlowRequest {
	s.Conf = &v
	return s
}

type SubmitFlowResponseBody struct {
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Id         *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SubmitFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitFlowResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitFlowResponseBody) SetData(v string) *SubmitFlowResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitFlowResponseBody) SetRequestId(v string) *SubmitFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitFlowResponseBody) SetInstanceId(v string) *SubmitFlowResponseBody {
	s.InstanceId = &v
	return s
}

func (s *SubmitFlowResponseBody) SetId(v string) *SubmitFlowResponseBody {
	s.Id = &v
	return s
}

type SubmitFlowResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitFlowResponse) GoString() string {
	return s.String()
}

func (s *SubmitFlowResponse) SetHeaders(v map[string]*string) *SubmitFlowResponse {
	s.Headers = v
	return s
}

func (s *SubmitFlowResponse) SetBody(v *SubmitFlowResponseBody) *SubmitFlowResponse {
	s.Body = v
	return s
}

type DescribeScalingCommonConfigRequest struct {
	// ResourceOwnerId
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeScalingCommonConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingCommonConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeScalingCommonConfigRequest) SetResourceOwnerId(v int64) *DescribeScalingCommonConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeScalingCommonConfigRequest) SetRegionId(v string) *DescribeScalingCommonConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeScalingCommonConfigRequest) SetResourceGroupId(v string) *DescribeScalingCommonConfigRequest {
	s.ResourceGroupId = &v
	return s
}

type DescribeScalingCommonConfigResponseBody struct {
	// requestId
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// autoScalingHookHeartbeatDefaultTime
	AutoScalingHookHeartbeatDefaultTime        *int64 `json:"AutoScalingHookHeartbeatDefaultTime,omitempty" xml:"AutoScalingHookHeartbeatDefaultTime,omitempty"`
	AutoScalingCoolDownTime                    *int64 `json:"AutoScalingCoolDownTime,omitempty" xml:"AutoScalingCoolDownTime,omitempty"`
	AutoScalingMNSScalingThreadSleepTime       *int64 `json:"AutoScalingMNSScalingThreadSleepTime,omitempty" xml:"AutoScalingMNSScalingThreadSleepTime,omitempty"`
	AutoScalingGroupMinSizeLimit               *int64 `json:"AutoScalingGroupMinSizeLimit,omitempty" xml:"AutoScalingGroupMinSizeLimit,omitempty"`
	AutoScalingGroupMaxSizeLimit               *int64 `json:"AutoScalingGroupMaxSizeLimit,omitempty" xml:"AutoScalingGroupMaxSizeLimit,omitempty"`
	AutoScalingRuleMinDelayLimit               *int64 `json:"AutoScalingRuleMinDelayLimit,omitempty" xml:"AutoScalingRuleMinDelayLimit,omitempty"`
	AutoScalingRuleAlarmDelayLimit             *int64 `json:"AutoScalingRuleAlarmDelayLimit,omitempty" xml:"AutoScalingRuleAlarmDelayLimit,omitempty"`
	AutoScalingRuleAlarmSilentTime             *int64 `json:"AutoScalingRuleAlarmSilentTime,omitempty" xml:"AutoScalingRuleAlarmSilentTime,omitempty"`
	AutoScalingConfigSystemDiskSize            *int64 `json:"AutoScalingConfigSystemDiskSize,omitempty" xml:"AutoScalingConfigSystemDiskSize,omitempty"`
	AutoScalingConfigDecommissionQueryInterval *int64 `json:"AutoScalingConfigDecommissionQueryInterval,omitempty" xml:"AutoScalingConfigDecommissionQueryInterval,omitempty"`
}

func (s DescribeScalingCommonConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingCommonConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScalingCommonConfigResponseBody) SetRequestId(v string) *DescribeScalingCommonConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScalingCommonConfigResponseBody) SetAutoScalingHookHeartbeatDefaultTime(v int64) *DescribeScalingCommonConfigResponseBody {
	s.AutoScalingHookHeartbeatDefaultTime = &v
	return s
}

func (s *DescribeScalingCommonConfigResponseBody) SetAutoScalingCoolDownTime(v int64) *DescribeScalingCommonConfigResponseBody {
	s.AutoScalingCoolDownTime = &v
	return s
}

func (s *DescribeScalingCommonConfigResponseBody) SetAutoScalingMNSScalingThreadSleepTime(v int64) *DescribeScalingCommonConfigResponseBody {
	s.AutoScalingMNSScalingThreadSleepTime = &v
	return s
}

func (s *DescribeScalingCommonConfigResponseBody) SetAutoScalingGroupMinSizeLimit(v int64) *DescribeScalingCommonConfigResponseBody {
	s.AutoScalingGroupMinSizeLimit = &v
	return s
}

func (s *DescribeScalingCommonConfigResponseBody) SetAutoScalingGroupMaxSizeLimit(v int64) *DescribeScalingCommonConfigResponseBody {
	s.AutoScalingGroupMaxSizeLimit = &v
	return s
}

func (s *DescribeScalingCommonConfigResponseBody) SetAutoScalingRuleMinDelayLimit(v int64) *DescribeScalingCommonConfigResponseBody {
	s.AutoScalingRuleMinDelayLimit = &v
	return s
}

func (s *DescribeScalingCommonConfigResponseBody) SetAutoScalingRuleAlarmDelayLimit(v int64) *DescribeScalingCommonConfigResponseBody {
	s.AutoScalingRuleAlarmDelayLimit = &v
	return s
}

func (s *DescribeScalingCommonConfigResponseBody) SetAutoScalingRuleAlarmSilentTime(v int64) *DescribeScalingCommonConfigResponseBody {
	s.AutoScalingRuleAlarmSilentTime = &v
	return s
}

func (s *DescribeScalingCommonConfigResponseBody) SetAutoScalingConfigSystemDiskSize(v int64) *DescribeScalingCommonConfigResponseBody {
	s.AutoScalingConfigSystemDiskSize = &v
	return s
}

func (s *DescribeScalingCommonConfigResponseBody) SetAutoScalingConfigDecommissionQueryInterval(v int64) *DescribeScalingCommonConfigResponseBody {
	s.AutoScalingConfigDecommissionQueryInterval = &v
	return s
}

type DescribeScalingCommonConfigResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeScalingCommonConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScalingCommonConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingCommonConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeScalingCommonConfigResponse) SetHeaders(v map[string]*string) *DescribeScalingCommonConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeScalingCommonConfigResponse) SetBody(v *DescribeScalingCommonConfigResponseBody) *DescribeScalingCommonConfigResponse {
	s.Body = v
	return s
}

type ResumeFlowRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProjectId      *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	FlowInstanceId *string `json:"FlowInstanceId,omitempty" xml:"FlowInstanceId,omitempty"`
}

func (s ResumeFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s ResumeFlowRequest) GoString() string {
	return s.String()
}

func (s *ResumeFlowRequest) SetRegionId(v string) *ResumeFlowRequest {
	s.RegionId = &v
	return s
}

func (s *ResumeFlowRequest) SetProjectId(v string) *ResumeFlowRequest {
	s.ProjectId = &v
	return s
}

func (s *ResumeFlowRequest) SetFlowInstanceId(v string) *ResumeFlowRequest {
	s.FlowInstanceId = &v
	return s
}

type ResumeFlowResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResumeFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResumeFlowResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeFlowResponseBody) SetData(v bool) *ResumeFlowResponseBody {
	s.Data = &v
	return s
}

func (s *ResumeFlowResponseBody) SetRequestId(v string) *ResumeFlowResponseBody {
	s.RequestId = &v
	return s
}

type ResumeFlowResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResumeFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResumeFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s ResumeFlowResponse) GoString() string {
	return s.String()
}

func (s *ResumeFlowResponse) SetHeaders(v map[string]*string) *ResumeFlowResponse {
	s.Headers = v
	return s
}

func (s *ResumeFlowResponse) SetBody(v *ResumeFlowResponseBody) *ResumeFlowResponse {
	s.Body = v
	return s
}

type RestoreFlowEntitySnapshotRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	OperatorId      *string `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	EntityType      *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	EntityId        *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	Revision        *string `json:"Revision,omitempty" xml:"Revision,omitempty"`
}

func (s RestoreFlowEntitySnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s RestoreFlowEntitySnapshotRequest) GoString() string {
	return s.String()
}

func (s *RestoreFlowEntitySnapshotRequest) SetResourceOwnerId(v int64) *RestoreFlowEntitySnapshotRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RestoreFlowEntitySnapshotRequest) SetRegionId(v string) *RestoreFlowEntitySnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *RestoreFlowEntitySnapshotRequest) SetOperatorId(v string) *RestoreFlowEntitySnapshotRequest {
	s.OperatorId = &v
	return s
}

func (s *RestoreFlowEntitySnapshotRequest) SetEntityType(v string) *RestoreFlowEntitySnapshotRequest {
	s.EntityType = &v
	return s
}

func (s *RestoreFlowEntitySnapshotRequest) SetEntityId(v string) *RestoreFlowEntitySnapshotRequest {
	s.EntityId = &v
	return s
}

func (s *RestoreFlowEntitySnapshotRequest) SetRevision(v string) *RestoreFlowEntitySnapshotRequest {
	s.Revision = &v
	return s
}

type RestoreFlowEntitySnapshotResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestoreFlowEntitySnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestoreFlowEntitySnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *RestoreFlowEntitySnapshotResponseBody) SetData(v bool) *RestoreFlowEntitySnapshotResponseBody {
	s.Data = &v
	return s
}

func (s *RestoreFlowEntitySnapshotResponseBody) SetRequestId(v string) *RestoreFlowEntitySnapshotResponseBody {
	s.RequestId = &v
	return s
}

type RestoreFlowEntitySnapshotResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RestoreFlowEntitySnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RestoreFlowEntitySnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s RestoreFlowEntitySnapshotResponse) GoString() string {
	return s.String()
}

func (s *RestoreFlowEntitySnapshotResponse) SetHeaders(v map[string]*string) *RestoreFlowEntitySnapshotResponse {
	s.Headers = v
	return s
}

func (s *RestoreFlowEntitySnapshotResponse) SetBody(v *RestoreFlowEntitySnapshotResponseBody) *RestoreFlowEntitySnapshotResponse {
	s.Body = v
	return s
}

type CreateLibraryRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Type            *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	LibraryVersion  *string `json:"LibraryVersion,omitempty" xml:"LibraryVersion,omitempty"`
	SourceType      *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	SourceLocation  *string `json:"SourceLocation,omitempty" xml:"SourceLocation,omitempty"`
	Scope           *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	Properties      *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
}

func (s CreateLibraryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLibraryRequest) GoString() string {
	return s.String()
}

func (s *CreateLibraryRequest) SetResourceOwnerId(v int64) *CreateLibraryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateLibraryRequest) SetRegionId(v string) *CreateLibraryRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLibraryRequest) SetClientToken(v string) *CreateLibraryRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateLibraryRequest) SetType(v string) *CreateLibraryRequest {
	s.Type = &v
	return s
}

func (s *CreateLibraryRequest) SetName(v string) *CreateLibraryRequest {
	s.Name = &v
	return s
}

func (s *CreateLibraryRequest) SetLibraryVersion(v string) *CreateLibraryRequest {
	s.LibraryVersion = &v
	return s
}

func (s *CreateLibraryRequest) SetSourceType(v string) *CreateLibraryRequest {
	s.SourceType = &v
	return s
}

func (s *CreateLibraryRequest) SetSourceLocation(v string) *CreateLibraryRequest {
	s.SourceLocation = &v
	return s
}

func (s *CreateLibraryRequest) SetScope(v string) *CreateLibraryRequest {
	s.Scope = &v
	return s
}

func (s *CreateLibraryRequest) SetProperties(v string) *CreateLibraryRequest {
	s.Properties = &v
	return s
}

type CreateLibraryResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLibraryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLibraryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLibraryResponseBody) SetData(v string) *CreateLibraryResponseBody {
	s.Data = &v
	return s
}

func (s *CreateLibraryResponseBody) SetRequestId(v string) *CreateLibraryResponseBody {
	s.RequestId = &v
	return s
}

type CreateLibraryResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateLibraryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLibraryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLibraryResponse) GoString() string {
	return s.String()
}

func (s *CreateLibraryResponse) SetHeaders(v map[string]*string) *CreateLibraryResponse {
	s.Headers = v
	return s
}

func (s *CreateLibraryResponse) SetBody(v *CreateLibraryResponseBody) *CreateLibraryResponse {
	s.Body = v
	return s
}

type ListVswitchRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VpcId       *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId      *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	DepositType *string `json:"DepositType,omitempty" xml:"DepositType,omitempty"`
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s ListVswitchRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVswitchRequest) GoString() string {
	return s.String()
}

func (s *ListVswitchRequest) SetRegionId(v string) *ListVswitchRequest {
	s.RegionId = &v
	return s
}

func (s *ListVswitchRequest) SetVpcId(v string) *ListVswitchRequest {
	s.VpcId = &v
	return s
}

func (s *ListVswitchRequest) SetZoneId(v string) *ListVswitchRequest {
	s.ZoneId = &v
	return s
}

func (s *ListVswitchRequest) SetDepositType(v string) *ListVswitchRequest {
	s.DepositType = &v
	return s
}

func (s *ListVswitchRequest) SetProductType(v string) *ListVswitchRequest {
	s.ProductType = &v
	return s
}

type ListVswitchResponseBody struct {
	RequestId   *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VswitchList *ListVswitchResponseBodyVswitchList `json:"VswitchList,omitempty" xml:"VswitchList,omitempty" type:"Struct"`
}

func (s ListVswitchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVswitchResponseBody) GoString() string {
	return s.String()
}

func (s *ListVswitchResponseBody) SetRequestId(v string) *ListVswitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVswitchResponseBody) SetVswitchList(v *ListVswitchResponseBodyVswitchList) *ListVswitchResponseBody {
	s.VswitchList = v
	return s
}

type ListVswitchResponseBodyVswitchList struct {
	Vswitch []*ListVswitchResponseBodyVswitchListVswitch `json:"Vswitch,omitempty" xml:"Vswitch,omitempty" type:"Repeated"`
}

func (s ListVswitchResponseBodyVswitchList) String() string {
	return tea.Prettify(s)
}

func (s ListVswitchResponseBodyVswitchList) GoString() string {
	return s.String()
}

func (s *ListVswitchResponseBodyVswitchList) SetVswitch(v []*ListVswitchResponseBodyVswitchListVswitch) *ListVswitchResponseBodyVswitchList {
	s.Vswitch = v
	return s
}

type ListVswitchResponseBodyVswitchListVswitch struct {
	CreationTime            *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VpcId                   *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	IsDefault               *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	VSwitchId               *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	CidrBlock               *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	AvailableIpAddressCount *string `json:"AvailableIpAddressCount,omitempty" xml:"AvailableIpAddressCount,omitempty"`
	ResourceGroupId         *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ZoneId                  *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	VSwitchName             *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
}

func (s ListVswitchResponseBodyVswitchListVswitch) String() string {
	return tea.Prettify(s)
}

func (s ListVswitchResponseBodyVswitchListVswitch) GoString() string {
	return s.String()
}

func (s *ListVswitchResponseBodyVswitchListVswitch) SetCreationTime(v string) *ListVswitchResponseBodyVswitchListVswitch {
	s.CreationTime = &v
	return s
}

func (s *ListVswitchResponseBodyVswitchListVswitch) SetStatus(v string) *ListVswitchResponseBodyVswitchListVswitch {
	s.Status = &v
	return s
}

func (s *ListVswitchResponseBodyVswitchListVswitch) SetVpcId(v string) *ListVswitchResponseBodyVswitchListVswitch {
	s.VpcId = &v
	return s
}

func (s *ListVswitchResponseBodyVswitchListVswitch) SetIsDefault(v bool) *ListVswitchResponseBodyVswitchListVswitch {
	s.IsDefault = &v
	return s
}

func (s *ListVswitchResponseBodyVswitchListVswitch) SetVSwitchId(v string) *ListVswitchResponseBodyVswitchListVswitch {
	s.VSwitchId = &v
	return s
}

func (s *ListVswitchResponseBodyVswitchListVswitch) SetCidrBlock(v string) *ListVswitchResponseBodyVswitchListVswitch {
	s.CidrBlock = &v
	return s
}

func (s *ListVswitchResponseBodyVswitchListVswitch) SetDescription(v string) *ListVswitchResponseBodyVswitchListVswitch {
	s.Description = &v
	return s
}

func (s *ListVswitchResponseBodyVswitchListVswitch) SetAvailableIpAddressCount(v string) *ListVswitchResponseBodyVswitchListVswitch {
	s.AvailableIpAddressCount = &v
	return s
}

func (s *ListVswitchResponseBodyVswitchListVswitch) SetResourceGroupId(v string) *ListVswitchResponseBodyVswitchListVswitch {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVswitchResponseBodyVswitchListVswitch) SetZoneId(v string) *ListVswitchResponseBodyVswitchListVswitch {
	s.ZoneId = &v
	return s
}

func (s *ListVswitchResponseBodyVswitchListVswitch) SetVSwitchName(v string) *ListVswitchResponseBodyVswitchListVswitch {
	s.VSwitchName = &v
	return s
}

type ListVswitchResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListVswitchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVswitchResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVswitchResponse) GoString() string {
	return s.String()
}

func (s *ListVswitchResponse) SetHeaders(v map[string]*string) *ListVswitchResponse {
	s.Headers = v
	return s
}

func (s *ListVswitchResponse) SetBody(v *ListVswitchResponseBody) *ListVswitchResponse {
	s.Body = v
	return s
}

type DeleteFlowProjectRequest struct {
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteFlowProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteFlowProjectRequest) SetRegionId(v string) *DeleteFlowProjectRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteFlowProjectRequest) SetProjectId(v string) *DeleteFlowProjectRequest {
	s.ProjectId = &v
	return s
}

type DeleteFlowProjectResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFlowProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFlowProjectResponseBody) SetData(v bool) *DeleteFlowProjectResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteFlowProjectResponseBody) SetRequestId(v string) *DeleteFlowProjectResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFlowProjectResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFlowProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFlowProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowProjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteFlowProjectResponse) SetHeaders(v map[string]*string) *DeleteFlowProjectResponse {
	s.Headers = v
	return s
}

func (s *DeleteFlowProjectResponse) SetBody(v *DeleteFlowProjectResponseBody) *DeleteFlowProjectResponse {
	s.Body = v
	return s
}

type ReleaseClusterRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Id              *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ForceRelease    *bool   `json:"ForceRelease,omitempty" xml:"ForceRelease,omitempty"`
}

func (s ReleaseClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseClusterRequest) GoString() string {
	return s.String()
}

func (s *ReleaseClusterRequest) SetResourceOwnerId(v int64) *ReleaseClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReleaseClusterRequest) SetRegionId(v string) *ReleaseClusterRequest {
	s.RegionId = &v
	return s
}

func (s *ReleaseClusterRequest) SetId(v string) *ReleaseClusterRequest {
	s.Id = &v
	return s
}

func (s *ReleaseClusterRequest) SetForceRelease(v bool) *ReleaseClusterRequest {
	s.ForceRelease = &v
	return s
}

type ReleaseClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseClusterResponseBody) SetRequestId(v string) *ReleaseClusterResponseBody {
	s.RequestId = &v
	return s
}

type ReleaseClusterResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReleaseClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleaseClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseClusterResponse) GoString() string {
	return s.String()
}

func (s *ReleaseClusterResponse) SetHeaders(v map[string]*string) *ReleaseClusterResponse {
	s.Headers = v
	return s
}

func (s *ReleaseClusterResponse) SetBody(v *ReleaseClusterResponseBody) *ReleaseClusterResponse {
	s.Body = v
	return s
}

type AddScalingConfigItemRequest struct {
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId       *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ScalingGroupBizId     *string `json:"ScalingGroupBizId,omitempty" xml:"ScalingGroupBizId,omitempty"`
	ConfigItemType        *string `json:"ConfigItemType,omitempty" xml:"ConfigItemType,omitempty"`
	ConfigItemInformation *string `json:"ConfigItemInformation,omitempty" xml:"ConfigItemInformation,omitempty"`
}

func (s AddScalingConfigItemRequest) String() string {
	return tea.Prettify(s)
}

func (s AddScalingConfigItemRequest) GoString() string {
	return s.String()
}

func (s *AddScalingConfigItemRequest) SetResourceOwnerId(v int64) *AddScalingConfigItemRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddScalingConfigItemRequest) SetRegionId(v string) *AddScalingConfigItemRequest {
	s.RegionId = &v
	return s
}

func (s *AddScalingConfigItemRequest) SetResourceGroupId(v string) *AddScalingConfigItemRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AddScalingConfigItemRequest) SetScalingGroupBizId(v string) *AddScalingConfigItemRequest {
	s.ScalingGroupBizId = &v
	return s
}

func (s *AddScalingConfigItemRequest) SetConfigItemType(v string) *AddScalingConfigItemRequest {
	s.ConfigItemType = &v
	return s
}

func (s *AddScalingConfigItemRequest) SetConfigItemInformation(v string) *AddScalingConfigItemRequest {
	s.ConfigItemInformation = &v
	return s
}

type AddScalingConfigItemResponseBody struct {
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// data
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s AddScalingConfigItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddScalingConfigItemResponseBody) GoString() string {
	return s.String()
}

func (s *AddScalingConfigItemResponseBody) SetRequestId(v string) *AddScalingConfigItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddScalingConfigItemResponseBody) SetData(v string) *AddScalingConfigItemResponseBody {
	s.Data = &v
	return s
}

type AddScalingConfigItemResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddScalingConfigItemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddScalingConfigItemResponse) String() string {
	return tea.Prettify(s)
}

func (s AddScalingConfigItemResponse) GoString() string {
	return s.String()
}

func (s *AddScalingConfigItemResponse) SetHeaders(v map[string]*string) *AddScalingConfigItemResponse {
	s.Headers = v
	return s
}

func (s *AddScalingConfigItemResponse) SetBody(v *AddScalingConfigItemResponseBody) *AddScalingConfigItemResponse {
	s.Body = v
	return s
}

type ResetUserPasswordRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Type            *string `json:"Type,omitempty" xml:"Type,omitempty"`
	OldPassword     *string `json:"OldPassword,omitempty" xml:"OldPassword,omitempty"`
	Password        *string `json:"Password,omitempty" xml:"Password,omitempty"`
	UserName        *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ResetUserPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetUserPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetUserPasswordRequest) SetResourceOwnerId(v int64) *ResetUserPasswordRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ResetUserPasswordRequest) SetRegionId(v string) *ResetUserPasswordRequest {
	s.RegionId = &v
	return s
}

func (s *ResetUserPasswordRequest) SetClusterId(v string) *ResetUserPasswordRequest {
	s.ClusterId = &v
	return s
}

func (s *ResetUserPasswordRequest) SetType(v string) *ResetUserPasswordRequest {
	s.Type = &v
	return s
}

func (s *ResetUserPasswordRequest) SetOldPassword(v string) *ResetUserPasswordRequest {
	s.OldPassword = &v
	return s
}

func (s *ResetUserPasswordRequest) SetPassword(v string) *ResetUserPasswordRequest {
	s.Password = &v
	return s
}

func (s *ResetUserPasswordRequest) SetUserName(v string) *ResetUserPasswordRequest {
	s.UserName = &v
	return s
}

type ResetUserPasswordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetUserPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetUserPasswordResponseBody) GoString() string {
	return s.String()
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

type ListFlowClusterAllHostsRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProjectId       *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListFlowClusterAllHostsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterAllHostsRequest) GoString() string {
	return s.String()
}

func (s *ListFlowClusterAllHostsRequest) SetRegionId(v string) *ListFlowClusterAllHostsRequest {
	s.RegionId = &v
	return s
}

func (s *ListFlowClusterAllHostsRequest) SetProjectId(v string) *ListFlowClusterAllHostsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFlowClusterAllHostsRequest) SetClusterId(v string) *ListFlowClusterAllHostsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListFlowClusterAllHostsRequest) SetResourceGroupId(v string) *ListFlowClusterAllHostsRequest {
	s.ResourceGroupId = &v
	return s
}

type ListFlowClusterAllHostsResponseBody struct {
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HostList  *ListFlowClusterAllHostsResponseBodyHostList `json:"HostList,omitempty" xml:"HostList,omitempty" type:"Struct"`
}

func (s ListFlowClusterAllHostsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterAllHostsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowClusterAllHostsResponseBody) SetRequestId(v string) *ListFlowClusterAllHostsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowClusterAllHostsResponseBody) SetHostList(v *ListFlowClusterAllHostsResponseBodyHostList) *ListFlowClusterAllHostsResponseBody {
	s.HostList = v
	return s
}

type ListFlowClusterAllHostsResponseBodyHostList struct {
	Host []*ListFlowClusterAllHostsResponseBodyHostListHost `json:"Host,omitempty" xml:"Host,omitempty" type:"Repeated"`
}

func (s ListFlowClusterAllHostsResponseBodyHostList) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterAllHostsResponseBodyHostList) GoString() string {
	return s.String()
}

func (s *ListFlowClusterAllHostsResponseBodyHostList) SetHost(v []*ListFlowClusterAllHostsResponseBodyHostListHost) *ListFlowClusterAllHostsResponseBodyHostList {
	s.Host = v
	return s
}

type ListFlowClusterAllHostsResponseBodyHostListHost struct {
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
	SerialNumber   *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	PrivateIp      *string `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	HostName       *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	InstanceType   *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	HostId         *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	HostInstanceId *string `json:"HostInstanceId,omitempty" xml:"HostInstanceId,omitempty"`
	Cpu            *int32  `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	PublicIp       *string `json:"PublicIp,omitempty" xml:"PublicIp,omitempty"`
	Memory         *int32  `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Role           *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s ListFlowClusterAllHostsResponseBodyHostListHost) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterAllHostsResponseBodyHostListHost) GoString() string {
	return s.String()
}

func (s *ListFlowClusterAllHostsResponseBodyHostListHost) SetStatus(v string) *ListFlowClusterAllHostsResponseBodyHostListHost {
	s.Status = &v
	return s
}

func (s *ListFlowClusterAllHostsResponseBodyHostListHost) SetType(v string) *ListFlowClusterAllHostsResponseBodyHostListHost {
	s.Type = &v
	return s
}

func (s *ListFlowClusterAllHostsResponseBodyHostListHost) SetSerialNumber(v string) *ListFlowClusterAllHostsResponseBodyHostListHost {
	s.SerialNumber = &v
	return s
}

func (s *ListFlowClusterAllHostsResponseBodyHostListHost) SetPrivateIp(v string) *ListFlowClusterAllHostsResponseBodyHostListHost {
	s.PrivateIp = &v
	return s
}

func (s *ListFlowClusterAllHostsResponseBodyHostListHost) SetHostName(v string) *ListFlowClusterAllHostsResponseBodyHostListHost {
	s.HostName = &v
	return s
}

func (s *ListFlowClusterAllHostsResponseBodyHostListHost) SetInstanceType(v string) *ListFlowClusterAllHostsResponseBodyHostListHost {
	s.InstanceType = &v
	return s
}

func (s *ListFlowClusterAllHostsResponseBodyHostListHost) SetHostId(v string) *ListFlowClusterAllHostsResponseBodyHostListHost {
	s.HostId = &v
	return s
}

func (s *ListFlowClusterAllHostsResponseBodyHostListHost) SetHostInstanceId(v string) *ListFlowClusterAllHostsResponseBodyHostListHost {
	s.HostInstanceId = &v
	return s
}

func (s *ListFlowClusterAllHostsResponseBodyHostListHost) SetCpu(v int32) *ListFlowClusterAllHostsResponseBodyHostListHost {
	s.Cpu = &v
	return s
}

func (s *ListFlowClusterAllHostsResponseBodyHostListHost) SetPublicIp(v string) *ListFlowClusterAllHostsResponseBodyHostListHost {
	s.PublicIp = &v
	return s
}

func (s *ListFlowClusterAllHostsResponseBodyHostListHost) SetMemory(v int32) *ListFlowClusterAllHostsResponseBodyHostListHost {
	s.Memory = &v
	return s
}

func (s *ListFlowClusterAllHostsResponseBodyHostListHost) SetRole(v string) *ListFlowClusterAllHostsResponseBodyHostListHost {
	s.Role = &v
	return s
}

type ListFlowClusterAllHostsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFlowClusterAllHostsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowClusterAllHostsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterAllHostsResponse) GoString() string {
	return s.String()
}

func (s *ListFlowClusterAllHostsResponse) SetHeaders(v map[string]*string) *ListFlowClusterAllHostsResponse {
	s.Headers = v
	return s
}

func (s *ListFlowClusterAllHostsResponse) SetBody(v *ListFlowClusterAllHostsResponseBody) *ListFlowClusterAllHostsResponse {
	s.Body = v
	return s
}

type DeleteLibrariesRequest struct {
	ResourceOwnerId  *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId         *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	LibraryBizIdList []*string `json:"LibraryBizIdList,omitempty" xml:"LibraryBizIdList,omitempty" type:"Repeated"`
}

func (s DeleteLibrariesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLibrariesRequest) GoString() string {
	return s.String()
}

func (s *DeleteLibrariesRequest) SetResourceOwnerId(v int64) *DeleteLibrariesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteLibrariesRequest) SetRegionId(v string) *DeleteLibrariesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLibrariesRequest) SetLibraryBizIdList(v []*string) *DeleteLibrariesRequest {
	s.LibraryBizIdList = v
	return s
}

type DeleteLibrariesResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLibrariesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLibrariesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLibrariesResponseBody) SetData(v bool) *DeleteLibrariesResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteLibrariesResponseBody) SetRequestId(v string) *DeleteLibrariesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteLibrariesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteLibrariesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteLibrariesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLibrariesResponse) GoString() string {
	return s.String()
}

func (s *DeleteLibrariesResponse) SetHeaders(v map[string]*string) *DeleteLibrariesResponse {
	s.Headers = v
	return s
}

func (s *DeleteLibrariesResponse) SetBody(v *DeleteLibrariesResponseBody) *DeleteLibrariesResponse {
	s.Body = v
	return s
}

type DescribeFlowCategoryTreeRequest struct {
	ProjectId  *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Mode       *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Keyword    *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
}

func (s DescribeFlowCategoryTreeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowCategoryTreeRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowCategoryTreeRequest) SetProjectId(v string) *DescribeFlowCategoryTreeRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeFlowCategoryTreeRequest) SetRegionId(v string) *DescribeFlowCategoryTreeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFlowCategoryTreeRequest) SetType(v string) *DescribeFlowCategoryTreeRequest {
	s.Type = &v
	return s
}

func (s *DescribeFlowCategoryTreeRequest) SetMode(v string) *DescribeFlowCategoryTreeRequest {
	s.Mode = &v
	return s
}

func (s *DescribeFlowCategoryTreeRequest) SetKeyword(v string) *DescribeFlowCategoryTreeRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeFlowCategoryTreeRequest) SetCategoryId(v string) *DescribeFlowCategoryTreeRequest {
	s.CategoryId = &v
	return s
}

type DescribeFlowCategoryTreeResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFlowCategoryTreeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowCategoryTreeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowCategoryTreeResponseBody) SetData(v string) *DescribeFlowCategoryTreeResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeFlowCategoryTreeResponseBody) SetRequestId(v string) *DescribeFlowCategoryTreeResponseBody {
	s.RequestId = &v
	return s
}

type DescribeFlowCategoryTreeResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFlowCategoryTreeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFlowCategoryTreeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowCategoryTreeResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowCategoryTreeResponse) SetHeaders(v map[string]*string) *DescribeFlowCategoryTreeResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowCategoryTreeResponse) SetBody(v *DescribeFlowCategoryTreeResponseBody) *DescribeFlowCategoryTreeResponse {
	s.Body = v
	return s
}

type ListDatasourceInstancesRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TargetClusterId *string `json:"TargetClusterId,omitempty" xml:"TargetClusterId,omitempty"`
	DatasourceType  *string `json:"DatasourceType,omitempty" xml:"DatasourceType,omitempty"`
}

func (s ListDatasourceInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDatasourceInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListDatasourceInstancesRequest) SetResourceOwnerId(v int64) *ListDatasourceInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListDatasourceInstancesRequest) SetRegionId(v string) *ListDatasourceInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *ListDatasourceInstancesRequest) SetTargetClusterId(v string) *ListDatasourceInstancesRequest {
	s.TargetClusterId = &v
	return s
}

func (s *ListDatasourceInstancesRequest) SetDatasourceType(v string) *ListDatasourceInstancesRequest {
	s.DatasourceType = &v
	return s
}

type ListDatasourceInstancesResponseBody struct {
	DdiDatasourceInfoList *ListDatasourceInstancesResponseBodyDdiDatasourceInfoList `json:"DdiDatasourceInfoList,omitempty" xml:"DdiDatasourceInfoList,omitempty" type:"Struct"`
	RequestId             *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDatasourceInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDatasourceInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatasourceInstancesResponseBody) SetDdiDatasourceInfoList(v *ListDatasourceInstancesResponseBodyDdiDatasourceInfoList) *ListDatasourceInstancesResponseBody {
	s.DdiDatasourceInfoList = v
	return s
}

func (s *ListDatasourceInstancesResponseBody) SetRequestId(v string) *ListDatasourceInstancesResponseBody {
	s.RequestId = &v
	return s
}

type ListDatasourceInstancesResponseBodyDdiDatasourceInfoList struct {
	DdiDatasourceInfo []*ListDatasourceInstancesResponseBodyDdiDatasourceInfoListDdiDatasourceInfo `json:"DdiDatasourceInfo,omitempty" xml:"DdiDatasourceInfo,omitempty" type:"Repeated"`
}

func (s ListDatasourceInstancesResponseBodyDdiDatasourceInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListDatasourceInstancesResponseBodyDdiDatasourceInfoList) GoString() string {
	return s.String()
}

func (s *ListDatasourceInstancesResponseBodyDdiDatasourceInfoList) SetDdiDatasourceInfo(v []*ListDatasourceInstancesResponseBodyDdiDatasourceInfoListDdiDatasourceInfo) *ListDatasourceInstancesResponseBodyDdiDatasourceInfoList {
	s.DdiDatasourceInfo = v
	return s
}

type ListDatasourceInstancesResponseBodyDdiDatasourceInfoListDdiDatasourceInfo struct {
	ClusterId          *string                                                                                      `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CreateTime         *int64                                                                                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DatasourceId       *string                                                                                      `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	Descreption        *string                                                                                      `json:"Descreption,omitempty" xml:"Descreption,omitempty"`
	ModifyTime         *int64                                                                                       `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Name               *string                                                                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	Status             *string                                                                                      `json:"Status,omitempty" xml:"Status,omitempty"`
	Type               *string                                                                                      `json:"Type,omitempty" xml:"Type,omitempty"`
	VswitchId          *string                                                                                      `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	PrivateAddressList *ListDatasourceInstancesResponseBodyDdiDatasourceInfoListDdiDatasourceInfoPrivateAddressList `json:"PrivateAddressList,omitempty" xml:"PrivateAddressList,omitempty" type:"Struct"`
	SecurityGroupId    *string                                                                                      `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s ListDatasourceInstancesResponseBodyDdiDatasourceInfoListDdiDatasourceInfo) String() string {
	return tea.Prettify(s)
}

func (s ListDatasourceInstancesResponseBodyDdiDatasourceInfoListDdiDatasourceInfo) GoString() string {
	return s.String()
}

func (s *ListDatasourceInstancesResponseBodyDdiDatasourceInfoListDdiDatasourceInfo) SetClusterId(v string) *ListDatasourceInstancesResponseBodyDdiDatasourceInfoListDdiDatasourceInfo {
	s.ClusterId = &v
	return s
}

func (s *ListDatasourceInstancesResponseBodyDdiDatasourceInfoListDdiDatasourceInfo) SetCreateTime(v int64) *ListDatasourceInstancesResponseBodyDdiDatasourceInfoListDdiDatasourceInfo {
	s.CreateTime = &v
	return s
}

func (s *ListDatasourceInstancesResponseBodyDdiDatasourceInfoListDdiDatasourceInfo) SetDatasourceId(v string) *ListDatasourceInstancesResponseBodyDdiDatasourceInfoListDdiDatasourceInfo {
	s.DatasourceId = &v
	return s
}

func (s *ListDatasourceInstancesResponseBodyDdiDatasourceInfoListDdiDatasourceInfo) SetDescreption(v string) *ListDatasourceInstancesResponseBodyDdiDatasourceInfoListDdiDatasourceInfo {
	s.Descreption = &v
	return s
}

func (s *ListDatasourceInstancesResponseBodyDdiDatasourceInfoListDdiDatasourceInfo) SetModifyTime(v int64) *ListDatasourceInstancesResponseBodyDdiDatasourceInfoListDdiDatasourceInfo {
	s.ModifyTime = &v
	return s
}

func (s *ListDatasourceInstancesResponseBodyDdiDatasourceInfoListDdiDatasourceInfo) SetName(v string) *ListDatasourceInstancesResponseBodyDdiDatasourceInfoListDdiDatasourceInfo {
	s.Name = &v
	return s
}

func (s *ListDatasourceInstancesResponseBodyDdiDatasourceInfoListDdiDatasourceInfo) SetStatus(v string) *ListDatasourceInstancesResponseBodyDdiDatasourceInfoListDdiDatasourceInfo {
	s.Status = &v
	return s
}

func (s *ListDatasourceInstancesResponseBodyDdiDatasourceInfoListDdiDatasourceInfo) SetType(v string) *ListDatasourceInstancesResponseBodyDdiDatasourceInfoListDdiDatasourceInfo {
	s.Type = &v
	return s
}

func (s *ListDatasourceInstancesResponseBodyDdiDatasourceInfoListDdiDatasourceInfo) SetVswitchId(v string) *ListDatasourceInstancesResponseBodyDdiDatasourceInfoListDdiDatasourceInfo {
	s.VswitchId = &v
	return s
}

func (s *ListDatasourceInstancesResponseBodyDdiDatasourceInfoListDdiDatasourceInfo) SetPrivateAddressList(v *ListDatasourceInstancesResponseBodyDdiDatasourceInfoListDdiDatasourceInfoPrivateAddressList) *ListDatasourceInstancesResponseBodyDdiDatasourceInfoListDdiDatasourceInfo {
	s.PrivateAddressList = v
	return s
}

func (s *ListDatasourceInstancesResponseBodyDdiDatasourceInfoListDdiDatasourceInfo) SetSecurityGroupId(v string) *ListDatasourceInstancesResponseBodyDdiDatasourceInfoListDdiDatasourceInfo {
	s.SecurityGroupId = &v
	return s
}

type ListDatasourceInstancesResponseBodyDdiDatasourceInfoListDdiDatasourceInfoPrivateAddressList struct {
	PrivateAddressList []*string `json:"privateAddressList,omitempty" xml:"privateAddressList,omitempty" type:"Repeated"`
}

func (s ListDatasourceInstancesResponseBodyDdiDatasourceInfoListDdiDatasourceInfoPrivateAddressList) String() string {
	return tea.Prettify(s)
}

func (s ListDatasourceInstancesResponseBodyDdiDatasourceInfoListDdiDatasourceInfoPrivateAddressList) GoString() string {
	return s.String()
}

func (s *ListDatasourceInstancesResponseBodyDdiDatasourceInfoListDdiDatasourceInfoPrivateAddressList) SetPrivateAddressList(v []*string) *ListDatasourceInstancesResponseBodyDdiDatasourceInfoListDdiDatasourceInfoPrivateAddressList {
	s.PrivateAddressList = v
	return s
}

type ListDatasourceInstancesResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDatasourceInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDatasourceInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDatasourceInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListDatasourceInstancesResponse) SetHeaders(v map[string]*string) *ListDatasourceInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListDatasourceInstancesResponse) SetBody(v *ListDatasourceInstancesResponseBody) *ListDatasourceInstancesResponse {
	s.Body = v
	return s
}

type ListFlowNodeSqlResultRequest struct {
	NodeInstanceId *string `json:"NodeInstanceId,omitempty" xml:"NodeInstanceId,omitempty"`
	SqlIndex       *int32  `json:"SqlIndex,omitempty" xml:"SqlIndex,omitempty"`
	Offset         *int32  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	Length         *int32  `json:"Length,omitempty" xml:"Length,omitempty"`
	ProjectId      *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListFlowNodeSqlResultRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowNodeSqlResultRequest) GoString() string {
	return s.String()
}

func (s *ListFlowNodeSqlResultRequest) SetNodeInstanceId(v string) *ListFlowNodeSqlResultRequest {
	s.NodeInstanceId = &v
	return s
}

func (s *ListFlowNodeSqlResultRequest) SetSqlIndex(v int32) *ListFlowNodeSqlResultRequest {
	s.SqlIndex = &v
	return s
}

func (s *ListFlowNodeSqlResultRequest) SetOffset(v int32) *ListFlowNodeSqlResultRequest {
	s.Offset = &v
	return s
}

func (s *ListFlowNodeSqlResultRequest) SetLength(v int32) *ListFlowNodeSqlResultRequest {
	s.Length = &v
	return s
}

func (s *ListFlowNodeSqlResultRequest) SetProjectId(v string) *ListFlowNodeSqlResultRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFlowNodeSqlResultRequest) SetRegionId(v string) *ListFlowNodeSqlResultRequest {
	s.RegionId = &v
	return s
}

type ListFlowNodeSqlResultResponseBody struct {
	End        *bool                                        `json:"End,omitempty" xml:"End,omitempty"`
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HeaderList *ListFlowNodeSqlResultResponseBodyHeaderList `json:"HeaderList,omitempty" xml:"HeaderList,omitempty" type:"Struct"`
	RowList    *ListFlowNodeSqlResultResponseBodyRowList    `json:"RowList,omitempty" xml:"RowList,omitempty" type:"Struct"`
}

func (s ListFlowNodeSqlResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowNodeSqlResultResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowNodeSqlResultResponseBody) SetEnd(v bool) *ListFlowNodeSqlResultResponseBody {
	s.End = &v
	return s
}

func (s *ListFlowNodeSqlResultResponseBody) SetRequestId(v string) *ListFlowNodeSqlResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowNodeSqlResultResponseBody) SetHeaderList(v *ListFlowNodeSqlResultResponseBodyHeaderList) *ListFlowNodeSqlResultResponseBody {
	s.HeaderList = v
	return s
}

func (s *ListFlowNodeSqlResultResponseBody) SetRowList(v *ListFlowNodeSqlResultResponseBodyRowList) *ListFlowNodeSqlResultResponseBody {
	s.RowList = v
	return s
}

type ListFlowNodeSqlResultResponseBodyHeaderList struct {
	Header []*string `json:"Header,omitempty" xml:"Header,omitempty" type:"Repeated"`
}

func (s ListFlowNodeSqlResultResponseBodyHeaderList) String() string {
	return tea.Prettify(s)
}

func (s ListFlowNodeSqlResultResponseBodyHeaderList) GoString() string {
	return s.String()
}

func (s *ListFlowNodeSqlResultResponseBodyHeaderList) SetHeader(v []*string) *ListFlowNodeSqlResultResponseBodyHeaderList {
	s.Header = v
	return s
}

type ListFlowNodeSqlResultResponseBodyRowList struct {
	Row []*ListFlowNodeSqlResultResponseBodyRowListRow `json:"Row,omitempty" xml:"Row,omitempty" type:"Repeated"`
}

func (s ListFlowNodeSqlResultResponseBodyRowList) String() string {
	return tea.Prettify(s)
}

func (s ListFlowNodeSqlResultResponseBodyRowList) GoString() string {
	return s.String()
}

func (s *ListFlowNodeSqlResultResponseBodyRowList) SetRow(v []*ListFlowNodeSqlResultResponseBodyRowListRow) *ListFlowNodeSqlResultResponseBodyRowList {
	s.Row = v
	return s
}

type ListFlowNodeSqlResultResponseBodyRowListRow struct {
	RowIndex    *int32                                                  `json:"RowIndex,omitempty" xml:"RowIndex,omitempty"`
	RowItemList *ListFlowNodeSqlResultResponseBodyRowListRowRowItemList `json:"RowItemList,omitempty" xml:"RowItemList,omitempty" type:"Struct"`
}

func (s ListFlowNodeSqlResultResponseBodyRowListRow) String() string {
	return tea.Prettify(s)
}

func (s ListFlowNodeSqlResultResponseBodyRowListRow) GoString() string {
	return s.String()
}

func (s *ListFlowNodeSqlResultResponseBodyRowListRow) SetRowIndex(v int32) *ListFlowNodeSqlResultResponseBodyRowListRow {
	s.RowIndex = &v
	return s
}

func (s *ListFlowNodeSqlResultResponseBodyRowListRow) SetRowItemList(v *ListFlowNodeSqlResultResponseBodyRowListRowRowItemList) *ListFlowNodeSqlResultResponseBodyRowListRow {
	s.RowItemList = v
	return s
}

type ListFlowNodeSqlResultResponseBodyRowListRowRowItemList struct {
	RowItem []*string `json:"RowItem,omitempty" xml:"RowItem,omitempty" type:"Repeated"`
}

func (s ListFlowNodeSqlResultResponseBodyRowListRowRowItemList) String() string {
	return tea.Prettify(s)
}

func (s ListFlowNodeSqlResultResponseBodyRowListRowRowItemList) GoString() string {
	return s.String()
}

func (s *ListFlowNodeSqlResultResponseBodyRowListRowRowItemList) SetRowItem(v []*string) *ListFlowNodeSqlResultResponseBodyRowListRowRowItemList {
	s.RowItem = v
	return s
}

type ListFlowNodeSqlResultResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFlowNodeSqlResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowNodeSqlResultResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowNodeSqlResultResponse) GoString() string {
	return s.String()
}

func (s *ListFlowNodeSqlResultResponse) SetHeaders(v map[string]*string) *ListFlowNodeSqlResultResponse {
	s.Headers = v
	return s
}

func (s *ListFlowNodeSqlResultResponse) SetBody(v *ListFlowNodeSqlResultResponseBody) *ListFlowNodeSqlResultResponse {
	s.Body = v
	return s
}

type DescribeFlowJobRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeFlowJobRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowJobRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowJobRequest) SetProjectId(v string) *DescribeFlowJobRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeFlowJobRequest) SetId(v string) *DescribeFlowJobRequest {
	s.Id = &v
	return s
}

func (s *DescribeFlowJobRequest) SetRegionId(v string) *DescribeFlowJobRequest {
	s.RegionId = &v
	return s
}

type DescribeFlowJobResponseBody struct {
	Type              *string                                  `json:"Type,omitempty" xml:"Type,omitempty"`
	LastInstanceId    *string                                  `json:"LastInstanceId,omitempty" xml:"LastInstanceId,omitempty"`
	EnvConf           *string                                  `json:"EnvConf,omitempty" xml:"EnvConf,omitempty"`
	RetryInterval     *int64                                   `json:"RetryInterval,omitempty" xml:"RetryInterval,omitempty"`
	Mode              *string                                  `json:"Mode,omitempty" xml:"Mode,omitempty"`
	GmtModified       *int64                                   `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	MonitorConf       *string                                  `json:"MonitorConf,omitempty" xml:"MonitorConf,omitempty"`
	Params            *string                                  `json:"Params,omitempty" xml:"Params,omitempty"`
	RequestId         *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Description       *string                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	RetryPolicy       *string                                  `json:"RetryPolicy,omitempty" xml:"RetryPolicy,omitempty"`
	Adhoc             *string                                  `json:"Adhoc,omitempty" xml:"Adhoc,omitempty"`
	Name              *string                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	MaxRetry          *int32                                   `json:"MaxRetry,omitempty" xml:"MaxRetry,omitempty"`
	MaxRunningTimeSec *int64                                   `json:"MaxRunningTimeSec,omitempty" xml:"MaxRunningTimeSec,omitempty"`
	FailAct           *string                                  `json:"FailAct,omitempty" xml:"FailAct,omitempty"`
	CustomVariables   *string                                  `json:"CustomVariables,omitempty" xml:"CustomVariables,omitempty"`
	EditLockDetail    *string                                  `json:"EditLockDetail,omitempty" xml:"EditLockDetail,omitempty"`
	ParamConf         *string                                  `json:"ParamConf,omitempty" xml:"ParamConf,omitempty"`
	RunConf           *string                                  `json:"RunConf,omitempty" xml:"RunConf,omitempty"`
	GmtCreate         *int64                                   `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	CategoryId        *string                                  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	Id                *string                                  `json:"Id,omitempty" xml:"Id,omitempty"`
	AlertConf         *string                                  `json:"AlertConf,omitempty" xml:"AlertConf,omitempty"`
	ResourceList      *DescribeFlowJobResponseBodyResourceList `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Struct"`
}

func (s DescribeFlowJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowJobResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowJobResponseBody) SetType(v string) *DescribeFlowJobResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetLastInstanceId(v string) *DescribeFlowJobResponseBody {
	s.LastInstanceId = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetEnvConf(v string) *DescribeFlowJobResponseBody {
	s.EnvConf = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetRetryInterval(v int64) *DescribeFlowJobResponseBody {
	s.RetryInterval = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetMode(v string) *DescribeFlowJobResponseBody {
	s.Mode = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetGmtModified(v int64) *DescribeFlowJobResponseBody {
	s.GmtModified = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetMonitorConf(v string) *DescribeFlowJobResponseBody {
	s.MonitorConf = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetParams(v string) *DescribeFlowJobResponseBody {
	s.Params = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetRequestId(v string) *DescribeFlowJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetDescription(v string) *DescribeFlowJobResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetRetryPolicy(v string) *DescribeFlowJobResponseBody {
	s.RetryPolicy = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetAdhoc(v string) *DescribeFlowJobResponseBody {
	s.Adhoc = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetName(v string) *DescribeFlowJobResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetMaxRetry(v int32) *DescribeFlowJobResponseBody {
	s.MaxRetry = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetMaxRunningTimeSec(v int64) *DescribeFlowJobResponseBody {
	s.MaxRunningTimeSec = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetFailAct(v string) *DescribeFlowJobResponseBody {
	s.FailAct = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetCustomVariables(v string) *DescribeFlowJobResponseBody {
	s.CustomVariables = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetEditLockDetail(v string) *DescribeFlowJobResponseBody {
	s.EditLockDetail = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetParamConf(v string) *DescribeFlowJobResponseBody {
	s.ParamConf = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetRunConf(v string) *DescribeFlowJobResponseBody {
	s.RunConf = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetGmtCreate(v int64) *DescribeFlowJobResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetCategoryId(v string) *DescribeFlowJobResponseBody {
	s.CategoryId = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetId(v string) *DescribeFlowJobResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetAlertConf(v string) *DescribeFlowJobResponseBody {
	s.AlertConf = &v
	return s
}

func (s *DescribeFlowJobResponseBody) SetResourceList(v *DescribeFlowJobResponseBodyResourceList) *DescribeFlowJobResponseBody {
	s.ResourceList = v
	return s
}

type DescribeFlowJobResponseBodyResourceList struct {
	Resource []*DescribeFlowJobResponseBodyResourceListResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
}

func (s DescribeFlowJobResponseBodyResourceList) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowJobResponseBodyResourceList) GoString() string {
	return s.String()
}

func (s *DescribeFlowJobResponseBodyResourceList) SetResource(v []*DescribeFlowJobResponseBodyResourceListResource) *DescribeFlowJobResponseBodyResourceList {
	s.Resource = v
	return s
}

type DescribeFlowJobResponseBodyResourceListResource struct {
	Path  *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
}

func (s DescribeFlowJobResponseBodyResourceListResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowJobResponseBodyResourceListResource) GoString() string {
	return s.String()
}

func (s *DescribeFlowJobResponseBodyResourceListResource) SetPath(v string) *DescribeFlowJobResponseBodyResourceListResource {
	s.Path = &v
	return s
}

func (s *DescribeFlowJobResponseBodyResourceListResource) SetAlias(v string) *DescribeFlowJobResponseBodyResourceListResource {
	s.Alias = &v
	return s
}

type DescribeFlowJobResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFlowJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFlowJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowJobResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowJobResponse) SetHeaders(v map[string]*string) *DescribeFlowJobResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowJobResponse) SetBody(v *DescribeFlowJobResponseBody) *DescribeFlowJobResponse {
	s.Body = v
	return s
}

type DescribeLibraryInstallTaskDetailRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TaskBizId       *string `json:"TaskBizId,omitempty" xml:"TaskBizId,omitempty"`
}

func (s DescribeLibraryInstallTaskDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLibraryInstallTaskDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeLibraryInstallTaskDetailRequest) SetResourceOwnerId(v int64) *DescribeLibraryInstallTaskDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeLibraryInstallTaskDetailRequest) SetRegionId(v string) *DescribeLibraryInstallTaskDetailRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLibraryInstallTaskDetailRequest) SetTaskBizId(v string) *DescribeLibraryInstallTaskDetailRequest {
	s.TaskBizId = &v
	return s
}

type DescribeLibraryInstallTaskDetailResponseBody struct {
	LibraryBizId *string `json:"LibraryBizId,omitempty" xml:"LibraryBizId,omitempty"`
	Hostname     *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime    *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TaskType     *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskGroupId  *string `json:"TaskGroupId,omitempty" xml:"TaskGroupId,omitempty"`
	TaskStatus   *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskProcess  *int32  `json:"TaskProcess,omitempty" xml:"TaskProcess,omitempty"`
	ExecuteTime  *int64  `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	ClusterBizId *string `json:"ClusterBizId,omitempty" xml:"ClusterBizId,omitempty"`
	TaskId       *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Detail       *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
}

func (s DescribeLibraryInstallTaskDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLibraryInstallTaskDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLibraryInstallTaskDetailResponseBody) SetLibraryBizId(v string) *DescribeLibraryInstallTaskDetailResponseBody {
	s.LibraryBizId = &v
	return s
}

func (s *DescribeLibraryInstallTaskDetailResponseBody) SetHostname(v string) *DescribeLibraryInstallTaskDetailResponseBody {
	s.Hostname = &v
	return s
}

func (s *DescribeLibraryInstallTaskDetailResponseBody) SetEndTime(v int64) *DescribeLibraryInstallTaskDetailResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeLibraryInstallTaskDetailResponseBody) SetStartTime(v int64) *DescribeLibraryInstallTaskDetailResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeLibraryInstallTaskDetailResponseBody) SetTaskType(v string) *DescribeLibraryInstallTaskDetailResponseBody {
	s.TaskType = &v
	return s
}

func (s *DescribeLibraryInstallTaskDetailResponseBody) SetRequestId(v string) *DescribeLibraryInstallTaskDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLibraryInstallTaskDetailResponseBody) SetTaskGroupId(v string) *DescribeLibraryInstallTaskDetailResponseBody {
	s.TaskGroupId = &v
	return s
}

func (s *DescribeLibraryInstallTaskDetailResponseBody) SetTaskStatus(v string) *DescribeLibraryInstallTaskDetailResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *DescribeLibraryInstallTaskDetailResponseBody) SetTaskProcess(v int32) *DescribeLibraryInstallTaskDetailResponseBody {
	s.TaskProcess = &v
	return s
}

func (s *DescribeLibraryInstallTaskDetailResponseBody) SetExecuteTime(v int64) *DescribeLibraryInstallTaskDetailResponseBody {
	s.ExecuteTime = &v
	return s
}

func (s *DescribeLibraryInstallTaskDetailResponseBody) SetClusterBizId(v string) *DescribeLibraryInstallTaskDetailResponseBody {
	s.ClusterBizId = &v
	return s
}

func (s *DescribeLibraryInstallTaskDetailResponseBody) SetTaskId(v string) *DescribeLibraryInstallTaskDetailResponseBody {
	s.TaskId = &v
	return s
}

func (s *DescribeLibraryInstallTaskDetailResponseBody) SetDetail(v string) *DescribeLibraryInstallTaskDetailResponseBody {
	s.Detail = &v
	return s
}

type DescribeLibraryInstallTaskDetailResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeLibraryInstallTaskDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLibraryInstallTaskDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLibraryInstallTaskDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeLibraryInstallTaskDetailResponse) SetHeaders(v map[string]*string) *DescribeLibraryInstallTaskDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeLibraryInstallTaskDetailResponse) SetBody(v *DescribeLibraryInstallTaskDetailResponseBody) *DescribeLibraryInstallTaskDetailResponse {
	s.Body = v
	return s
}

type ModifyFlowForWebRequest struct {
	RegionId                *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProjectId               *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Id                      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name                    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Periodic                *bool   `json:"Periodic,omitempty" xml:"Periodic,omitempty"`
	StartSchedule           *int64  `json:"StartSchedule,omitempty" xml:"StartSchedule,omitempty"`
	EndSchedule             *int64  `json:"EndSchedule,omitempty" xml:"EndSchedule,omitempty"`
	CronExpr                *string `json:"CronExpr,omitempty" xml:"CronExpr,omitempty"`
	CreateCluster           *bool   `json:"CreateCluster,omitempty" xml:"CreateCluster,omitempty"`
	ClusterId               *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	HostName                *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	Namespace               *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Graph                   *string `json:"Graph,omitempty" xml:"Graph,omitempty"`
	AlertConf               *string `json:"AlertConf,omitempty" xml:"AlertConf,omitempty"`
	AlertUserGroupBizId     *string `json:"AlertUserGroupBizId,omitempty" xml:"AlertUserGroupBizId,omitempty"`
	AlertDingDingGroupBizId *string `json:"AlertDingDingGroupBizId,omitempty" xml:"AlertDingDingGroupBizId,omitempty"`
	ParentFlowList          *string `json:"ParentFlowList,omitempty" xml:"ParentFlowList,omitempty"`
	ParentCategory          *string `json:"ParentCategory,omitempty" xml:"ParentCategory,omitempty"`
}

func (s ModifyFlowForWebRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowForWebRequest) GoString() string {
	return s.String()
}

func (s *ModifyFlowForWebRequest) SetRegionId(v string) *ModifyFlowForWebRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyFlowForWebRequest) SetProjectId(v string) *ModifyFlowForWebRequest {
	s.ProjectId = &v
	return s
}

func (s *ModifyFlowForWebRequest) SetId(v string) *ModifyFlowForWebRequest {
	s.Id = &v
	return s
}

func (s *ModifyFlowForWebRequest) SetName(v string) *ModifyFlowForWebRequest {
	s.Name = &v
	return s
}

func (s *ModifyFlowForWebRequest) SetStatus(v string) *ModifyFlowForWebRequest {
	s.Status = &v
	return s
}

func (s *ModifyFlowForWebRequest) SetDescription(v string) *ModifyFlowForWebRequest {
	s.Description = &v
	return s
}

func (s *ModifyFlowForWebRequest) SetPeriodic(v bool) *ModifyFlowForWebRequest {
	s.Periodic = &v
	return s
}

func (s *ModifyFlowForWebRequest) SetStartSchedule(v int64) *ModifyFlowForWebRequest {
	s.StartSchedule = &v
	return s
}

func (s *ModifyFlowForWebRequest) SetEndSchedule(v int64) *ModifyFlowForWebRequest {
	s.EndSchedule = &v
	return s
}

func (s *ModifyFlowForWebRequest) SetCronExpr(v string) *ModifyFlowForWebRequest {
	s.CronExpr = &v
	return s
}

func (s *ModifyFlowForWebRequest) SetCreateCluster(v bool) *ModifyFlowForWebRequest {
	s.CreateCluster = &v
	return s
}

func (s *ModifyFlowForWebRequest) SetClusterId(v string) *ModifyFlowForWebRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyFlowForWebRequest) SetHostName(v string) *ModifyFlowForWebRequest {
	s.HostName = &v
	return s
}

func (s *ModifyFlowForWebRequest) SetNamespace(v string) *ModifyFlowForWebRequest {
	s.Namespace = &v
	return s
}

func (s *ModifyFlowForWebRequest) SetGraph(v string) *ModifyFlowForWebRequest {
	s.Graph = &v
	return s
}

func (s *ModifyFlowForWebRequest) SetAlertConf(v string) *ModifyFlowForWebRequest {
	s.AlertConf = &v
	return s
}

func (s *ModifyFlowForWebRequest) SetAlertUserGroupBizId(v string) *ModifyFlowForWebRequest {
	s.AlertUserGroupBizId = &v
	return s
}

func (s *ModifyFlowForWebRequest) SetAlertDingDingGroupBizId(v string) *ModifyFlowForWebRequest {
	s.AlertDingDingGroupBizId = &v
	return s
}

func (s *ModifyFlowForWebRequest) SetParentFlowList(v string) *ModifyFlowForWebRequest {
	s.ParentFlowList = &v
	return s
}

func (s *ModifyFlowForWebRequest) SetParentCategory(v string) *ModifyFlowForWebRequest {
	s.ParentCategory = &v
	return s
}

type ModifyFlowForWebResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyFlowForWebResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowForWebResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFlowForWebResponseBody) SetData(v bool) *ModifyFlowForWebResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyFlowForWebResponseBody) SetRequestId(v string) *ModifyFlowForWebResponseBody {
	s.RequestId = &v
	return s
}

type ModifyFlowForWebResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyFlowForWebResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyFlowForWebResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowForWebResponse) GoString() string {
	return s.String()
}

func (s *ModifyFlowForWebResponse) SetHeaders(v map[string]*string) *ModifyFlowForWebResponse {
	s.Headers = v
	return s
}

func (s *ModifyFlowForWebResponse) SetBody(v *ModifyFlowForWebResponseBody) *ModifyFlowForWebResponse {
	s.Body = v
	return s
}

type RemoveScalingConfigItemRequest struct {
	ResourceOwnerId   *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId   *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ScalingGroupBizId *string `json:"ScalingGroupBizId,omitempty" xml:"ScalingGroupBizId,omitempty"`
	ConfigItemType    *string `json:"ConfigItemType,omitempty" xml:"ConfigItemType,omitempty"`
	ConfigItemBizId   *string `json:"ConfigItemBizId,omitempty" xml:"ConfigItemBizId,omitempty"`
}

func (s RemoveScalingConfigItemRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveScalingConfigItemRequest) GoString() string {
	return s.String()
}

func (s *RemoveScalingConfigItemRequest) SetResourceOwnerId(v int64) *RemoveScalingConfigItemRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RemoveScalingConfigItemRequest) SetRegionId(v string) *RemoveScalingConfigItemRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveScalingConfigItemRequest) SetResourceGroupId(v string) *RemoveScalingConfigItemRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *RemoveScalingConfigItemRequest) SetScalingGroupBizId(v string) *RemoveScalingConfigItemRequest {
	s.ScalingGroupBizId = &v
	return s
}

func (s *RemoveScalingConfigItemRequest) SetConfigItemType(v string) *RemoveScalingConfigItemRequest {
	s.ConfigItemType = &v
	return s
}

func (s *RemoveScalingConfigItemRequest) SetConfigItemBizId(v string) *RemoveScalingConfigItemRequest {
	s.ConfigItemBizId = &v
	return s
}

type RemoveScalingConfigItemResponseBody struct {
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// data
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s RemoveScalingConfigItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveScalingConfigItemResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveScalingConfigItemResponseBody) SetRequestId(v string) *RemoveScalingConfigItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveScalingConfigItemResponseBody) SetData(v bool) *RemoveScalingConfigItemResponseBody {
	s.Data = &v
	return s
}

type RemoveScalingConfigItemResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveScalingConfigItemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveScalingConfigItemResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveScalingConfigItemResponse) GoString() string {
	return s.String()
}

func (s *RemoveScalingConfigItemResponse) SetHeaders(v map[string]*string) *RemoveScalingConfigItemResponse {
	s.Headers = v
	return s
}

func (s *RemoveScalingConfigItemResponse) SetBody(v *RemoveScalingConfigItemResponseBody) *RemoveScalingConfigItemResponse {
	s.Body = v
	return s
}

type DescribeSecurityWhiteListRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeSecurityWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityWhiteListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityWhiteListRequest) SetClusterId(v string) *DescribeSecurityWhiteListRequest {
	s.ClusterId = &v
	return s
}

type DescribeSecurityWhiteListResponseBody struct {
	DescribeSecurityWhiteList []*DescribeSecurityWhiteListResponseBodyDescribeSecurityWhiteList `json:"DescribeSecurityWhiteList,omitempty" xml:"DescribeSecurityWhiteList,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSecurityWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityWhiteListResponseBody) SetDescribeSecurityWhiteList(v []*DescribeSecurityWhiteListResponseBodyDescribeSecurityWhiteList) *DescribeSecurityWhiteListResponseBody {
	s.DescribeSecurityWhiteList = v
	return s
}

func (s *DescribeSecurityWhiteListResponseBody) SetRequestId(v string) *DescribeSecurityWhiteListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeSecurityWhiteListResponseBodyDescribeSecurityWhiteList struct {
	PortRange   *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	WhiteIp     *string `json:"WhiteIp,omitempty" xml:"WhiteIp,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
}

func (s DescribeSecurityWhiteListResponseBodyDescribeSecurityWhiteList) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityWhiteListResponseBodyDescribeSecurityWhiteList) GoString() string {
	return s.String()
}

func (s *DescribeSecurityWhiteListResponseBodyDescribeSecurityWhiteList) SetPortRange(v string) *DescribeSecurityWhiteListResponseBodyDescribeSecurityWhiteList {
	s.PortRange = &v
	return s
}

func (s *DescribeSecurityWhiteListResponseBodyDescribeSecurityWhiteList) SetWhiteIp(v string) *DescribeSecurityWhiteListResponseBodyDescribeSecurityWhiteList {
	s.WhiteIp = &v
	return s
}

func (s *DescribeSecurityWhiteListResponseBodyDescribeSecurityWhiteList) SetDescription(v string) *DescribeSecurityWhiteListResponseBodyDescribeSecurityWhiteList {
	s.Description = &v
	return s
}

func (s *DescribeSecurityWhiteListResponseBodyDescribeSecurityWhiteList) SetCreateTime(v string) *DescribeSecurityWhiteListResponseBodyDescribeSecurityWhiteList {
	s.CreateTime = &v
	return s
}

type DescribeSecurityWhiteListResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSecurityWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSecurityWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityWhiteListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityWhiteListResponse) SetHeaders(v map[string]*string) *DescribeSecurityWhiteListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecurityWhiteListResponse) SetBody(v *DescribeSecurityWhiteListResponseBody) *DescribeSecurityWhiteListResponse {
	s.Body = v
	return s
}

type DescribeFlowNodeInstanceContainerLogRequest struct {
	Offset         *int32  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	Length         *int32  `json:"Length,omitempty" xml:"Length,omitempty"`
	NodeInstanceId *string `json:"NodeInstanceId,omitempty" xml:"NodeInstanceId,omitempty"`
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ContainerId    *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	LogName        *string `json:"LogName,omitempty" xml:"LogName,omitempty"`
	ProjectId      *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeFlowNodeInstanceContainerLogRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowNodeInstanceContainerLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowNodeInstanceContainerLogRequest) SetOffset(v int32) *DescribeFlowNodeInstanceContainerLogRequest {
	s.Offset = &v
	return s
}

func (s *DescribeFlowNodeInstanceContainerLogRequest) SetLength(v int32) *DescribeFlowNodeInstanceContainerLogRequest {
	s.Length = &v
	return s
}

func (s *DescribeFlowNodeInstanceContainerLogRequest) SetNodeInstanceId(v string) *DescribeFlowNodeInstanceContainerLogRequest {
	s.NodeInstanceId = &v
	return s
}

func (s *DescribeFlowNodeInstanceContainerLogRequest) SetAppId(v string) *DescribeFlowNodeInstanceContainerLogRequest {
	s.AppId = &v
	return s
}

func (s *DescribeFlowNodeInstanceContainerLogRequest) SetContainerId(v string) *DescribeFlowNodeInstanceContainerLogRequest {
	s.ContainerId = &v
	return s
}

func (s *DescribeFlowNodeInstanceContainerLogRequest) SetLogName(v string) *DescribeFlowNodeInstanceContainerLogRequest {
	s.LogName = &v
	return s
}

func (s *DescribeFlowNodeInstanceContainerLogRequest) SetProjectId(v string) *DescribeFlowNodeInstanceContainerLogRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeFlowNodeInstanceContainerLogRequest) SetRegionId(v string) *DescribeFlowNodeInstanceContainerLogRequest {
	s.RegionId = &v
	return s
}

type DescribeFlowNodeInstanceContainerLogResponseBody struct {
	LogEnd    *bool                                                      `json:"LogEnd,omitempty" xml:"LogEnd,omitempty"`
	RequestId *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	LogEntrys *DescribeFlowNodeInstanceContainerLogResponseBodyLogEntrys `json:"LogEntrys,omitempty" xml:"LogEntrys,omitempty" type:"Struct"`
}

func (s DescribeFlowNodeInstanceContainerLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowNodeInstanceContainerLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowNodeInstanceContainerLogResponseBody) SetLogEnd(v bool) *DescribeFlowNodeInstanceContainerLogResponseBody {
	s.LogEnd = &v
	return s
}

func (s *DescribeFlowNodeInstanceContainerLogResponseBody) SetRequestId(v string) *DescribeFlowNodeInstanceContainerLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowNodeInstanceContainerLogResponseBody) SetLogEntrys(v *DescribeFlowNodeInstanceContainerLogResponseBodyLogEntrys) *DescribeFlowNodeInstanceContainerLogResponseBody {
	s.LogEntrys = v
	return s
}

type DescribeFlowNodeInstanceContainerLogResponseBodyLogEntrys struct {
	LogEntry []*DescribeFlowNodeInstanceContainerLogResponseBodyLogEntrysLogEntry `json:"LogEntry,omitempty" xml:"LogEntry,omitempty" type:"Repeated"`
}

func (s DescribeFlowNodeInstanceContainerLogResponseBodyLogEntrys) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowNodeInstanceContainerLogResponseBodyLogEntrys) GoString() string {
	return s.String()
}

func (s *DescribeFlowNodeInstanceContainerLogResponseBodyLogEntrys) SetLogEntry(v []*DescribeFlowNodeInstanceContainerLogResponseBodyLogEntrysLogEntry) *DescribeFlowNodeInstanceContainerLogResponseBodyLogEntrys {
	s.LogEntry = v
	return s
}

type DescribeFlowNodeInstanceContainerLogResponseBodyLogEntrysLogEntry struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s DescribeFlowNodeInstanceContainerLogResponseBodyLogEntrysLogEntry) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowNodeInstanceContainerLogResponseBodyLogEntrysLogEntry) GoString() string {
	return s.String()
}

func (s *DescribeFlowNodeInstanceContainerLogResponseBodyLogEntrysLogEntry) SetContent(v string) *DescribeFlowNodeInstanceContainerLogResponseBodyLogEntrysLogEntry {
	s.Content = &v
	return s
}

type DescribeFlowNodeInstanceContainerLogResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFlowNodeInstanceContainerLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFlowNodeInstanceContainerLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowNodeInstanceContainerLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowNodeInstanceContainerLogResponse) SetHeaders(v map[string]*string) *DescribeFlowNodeInstanceContainerLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowNodeInstanceContainerLogResponse) SetBody(v *DescribeFlowNodeInstanceContainerLogResponseBody) *DescribeFlowNodeInstanceContainerLogResponse {
	s.Body = v
	return s
}

type RerunFlowRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProjectId      *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	FlowInstanceId *string `json:"FlowInstanceId,omitempty" xml:"FlowInstanceId,omitempty"`
	ReRunFail      *bool   `json:"ReRunFail,omitempty" xml:"ReRunFail,omitempty"`
}

func (s RerunFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s RerunFlowRequest) GoString() string {
	return s.String()
}

func (s *RerunFlowRequest) SetRegionId(v string) *RerunFlowRequest {
	s.RegionId = &v
	return s
}

func (s *RerunFlowRequest) SetProjectId(v string) *RerunFlowRequest {
	s.ProjectId = &v
	return s
}

func (s *RerunFlowRequest) SetFlowInstanceId(v string) *RerunFlowRequest {
	s.FlowInstanceId = &v
	return s
}

func (s *RerunFlowRequest) SetReRunFail(v bool) *RerunFlowRequest {
	s.ReRunFail = &v
	return s
}

type RerunFlowResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RerunFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RerunFlowResponseBody) GoString() string {
	return s.String()
}

func (s *RerunFlowResponseBody) SetData(v bool) *RerunFlowResponseBody {
	s.Data = &v
	return s
}

func (s *RerunFlowResponseBody) SetRequestId(v string) *RerunFlowResponseBody {
	s.RequestId = &v
	return s
}

type RerunFlowResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RerunFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RerunFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s RerunFlowResponse) GoString() string {
	return s.String()
}

func (s *RerunFlowResponse) SetHeaders(v map[string]*string) *RerunFlowResponse {
	s.Headers = v
	return s
}

func (s *RerunFlowResponse) SetBody(v *RerunFlowResponseBody) *RerunFlowResponse {
	s.Body = v
	return s
}

type ListTagKeysRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// 资源类型
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Category     *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Scope        *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PageSize     *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) SetResourceOwnerId(v int64) *ListTagKeysRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTagKeysRequest) SetRegionId(v string) *ListTagKeysRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceGroupId(v string) *ListTagKeysRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceType(v string) *ListTagKeysRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagKeysRequest) SetCategory(v string) *ListTagKeysRequest {
	s.Category = &v
	return s
}

func (s *ListTagKeysRequest) SetScope(v string) *ListTagKeysRequest {
	s.Scope = &v
	return s
}

func (s *ListTagKeysRequest) SetNextToken(v string) *ListTagKeysRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysRequest) SetPageSize(v int64) *ListTagKeysRequest {
	s.PageSize = &v
	return s
}

type ListTagKeysResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 是否成功响应
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// 是否分页
	Paging *bool `json:"Paging,omitempty" xml:"Paging,omitempty"`
	// 标签键集合
	Data      []*ListTagKeysResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorCode *string                        `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ListTagKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBody) SetRequestId(v string) *ListTagKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagKeysResponseBody) SetSuccess(v bool) *ListTagKeysResponseBody {
	s.Success = &v
	return s
}

func (s *ListTagKeysResponseBody) SetPaging(v bool) *ListTagKeysResponseBody {
	s.Paging = &v
	return s
}

func (s *ListTagKeysResponseBody) SetData(v []*ListTagKeysResponseBodyData) *ListTagKeysResponseBody {
	s.Data = v
	return s
}

func (s *ListTagKeysResponseBody) SetCode(v string) *ListTagKeysResponseBody {
	s.Code = &v
	return s
}

func (s *ListTagKeysResponseBody) SetErrorCode(v string) *ListTagKeysResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListTagKeysResponseBody) SetMessage(v string) *ListTagKeysResponseBody {
	s.Message = &v
	return s
}

type ListTagKeysResponseBodyData struct {
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 标签键集合
	Items []*string `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
}

func (s ListTagKeysResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBodyData) SetPageNumber(v int64) *ListTagKeysResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListTagKeysResponseBodyData) SetPageSize(v int64) *ListTagKeysResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListTagKeysResponseBodyData) SetTotalCount(v int64) *ListTagKeysResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListTagKeysResponseBodyData) SetNextToken(v string) *ListTagKeysResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysResponseBodyData) SetItems(v []*string) *ListTagKeysResponseBodyData {
	s.Items = v
	return s
}

type ListTagKeysResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponse) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponse) SetHeaders(v map[string]*string) *ListTagKeysResponse {
	s.Headers = v
	return s
}

func (s *ListTagKeysResponse) SetBody(v *ListTagKeysResponseBody) *ListTagKeysResponse {
	s.Body = v
	return s
}

type DescribeClusterOperationHostTaskLogRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OperationId     *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	HostId          *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	TaskId          *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeClusterOperationHostTaskLogRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterOperationHostTaskLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterOperationHostTaskLogRequest) SetResourceOwnerId(v int64) *DescribeClusterOperationHostTaskLogRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeClusterOperationHostTaskLogRequest) SetRegionId(v string) *DescribeClusterOperationHostTaskLogRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeClusterOperationHostTaskLogRequest) SetClusterId(v string) *DescribeClusterOperationHostTaskLogRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterOperationHostTaskLogRequest) SetOperationId(v string) *DescribeClusterOperationHostTaskLogRequest {
	s.OperationId = &v
	return s
}

func (s *DescribeClusterOperationHostTaskLogRequest) SetHostId(v string) *DescribeClusterOperationHostTaskLogRequest {
	s.HostId = &v
	return s
}

func (s *DescribeClusterOperationHostTaskLogRequest) SetTaskId(v string) *DescribeClusterOperationHostTaskLogRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeClusterOperationHostTaskLogRequest) SetStatus(v string) *DescribeClusterOperationHostTaskLogRequest {
	s.Status = &v
	return s
}

type DescribeClusterOperationHostTaskLogResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Stdout    *string `json:"Stdout,omitempty" xml:"Stdout,omitempty"`
	Stderr    *string `json:"Stderr,omitempty" xml:"Stderr,omitempty"`
}

func (s DescribeClusterOperationHostTaskLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterOperationHostTaskLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterOperationHostTaskLogResponseBody) SetRequestId(v string) *DescribeClusterOperationHostTaskLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterOperationHostTaskLogResponseBody) SetStdout(v string) *DescribeClusterOperationHostTaskLogResponseBody {
	s.Stdout = &v
	return s
}

func (s *DescribeClusterOperationHostTaskLogResponseBody) SetStderr(v string) *DescribeClusterOperationHostTaskLogResponseBody {
	s.Stderr = &v
	return s
}

type DescribeClusterOperationHostTaskLogResponse struct {
	Headers map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClusterOperationHostTaskLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClusterOperationHostTaskLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterOperationHostTaskLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterOperationHostTaskLogResponse) SetHeaders(v map[string]*string) *DescribeClusterOperationHostTaskLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterOperationHostTaskLogResponse) SetBody(v *DescribeClusterOperationHostTaskLogResponseBody) *DescribeClusterOperationHostTaskLogResponse {
	s.Body = v
	return s
}

type KillFlowJobRequest struct {
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	JobInstanceId *string `json:"JobInstanceId,omitempty" xml:"JobInstanceId,omitempty"`
}

func (s KillFlowJobRequest) String() string {
	return tea.Prettify(s)
}

func (s KillFlowJobRequest) GoString() string {
	return s.String()
}

func (s *KillFlowJobRequest) SetRegionId(v string) *KillFlowJobRequest {
	s.RegionId = &v
	return s
}

func (s *KillFlowJobRequest) SetProjectId(v string) *KillFlowJobRequest {
	s.ProjectId = &v
	return s
}

func (s *KillFlowJobRequest) SetJobInstanceId(v string) *KillFlowJobRequest {
	s.JobInstanceId = &v
	return s
}

type KillFlowJobResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s KillFlowJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s KillFlowJobResponseBody) GoString() string {
	return s.String()
}

func (s *KillFlowJobResponseBody) SetData(v bool) *KillFlowJobResponseBody {
	s.Data = &v
	return s
}

func (s *KillFlowJobResponseBody) SetRequestId(v string) *KillFlowJobResponseBody {
	s.RequestId = &v
	return s
}

type KillFlowJobResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *KillFlowJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s KillFlowJobResponse) String() string {
	return tea.Prettify(s)
}

func (s KillFlowJobResponse) GoString() string {
	return s.String()
}

func (s *KillFlowJobResponse) SetHeaders(v map[string]*string) *KillFlowJobResponse {
	s.Headers = v
	return s
}

func (s *KillFlowJobResponse) SetBody(v *KillFlowJobResponseBody) *KillFlowJobResponse {
	s.Body = v
	return s
}

type UninstallLibrariesRequest struct {
	ResourceOwnerId  *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId         *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClusterBizIdList []*string `json:"ClusterBizIdList,omitempty" xml:"ClusterBizIdList,omitempty" type:"Repeated"`
	LibraryBizId     *string   `json:"LibraryBizId,omitempty" xml:"LibraryBizId,omitempty"`
}

func (s UninstallLibrariesRequest) String() string {
	return tea.Prettify(s)
}

func (s UninstallLibrariesRequest) GoString() string {
	return s.String()
}

func (s *UninstallLibrariesRequest) SetResourceOwnerId(v int64) *UninstallLibrariesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UninstallLibrariesRequest) SetRegionId(v string) *UninstallLibrariesRequest {
	s.RegionId = &v
	return s
}

func (s *UninstallLibrariesRequest) SetClusterBizIdList(v []*string) *UninstallLibrariesRequest {
	s.ClusterBizIdList = v
	return s
}

func (s *UninstallLibrariesRequest) SetLibraryBizId(v string) *UninstallLibrariesRequest {
	s.LibraryBizId = &v
	return s
}

type UninstallLibrariesResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UninstallLibrariesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UninstallLibrariesResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallLibrariesResponseBody) SetData(v string) *UninstallLibrariesResponseBody {
	s.Data = &v
	return s
}

func (s *UninstallLibrariesResponseBody) SetRequestId(v string) *UninstallLibrariesResponseBody {
	s.RequestId = &v
	return s
}

type UninstallLibrariesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UninstallLibrariesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UninstallLibrariesResponse) String() string {
	return tea.Prettify(s)
}

func (s UninstallLibrariesResponse) GoString() string {
	return s.String()
}

func (s *UninstallLibrariesResponse) SetHeaders(v map[string]*string) *UninstallLibrariesResponse {
	s.Headers = v
	return s
}

func (s *UninstallLibrariesResponse) SetBody(v *UninstallLibrariesResponseBody) *UninstallLibrariesResponse {
	s.Body = v
	return s
}

type DescribeClusterV2Request struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Id              *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeClusterV2Request) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2Request) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2Request) SetResourceOwnerId(v int64) *DescribeClusterV2Request {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeClusterV2Request) SetRegionId(v string) *DescribeClusterV2Request {
	s.RegionId = &v
	return s
}

func (s *DescribeClusterV2Request) SetId(v string) *DescribeClusterV2Request {
	s.Id = &v
	return s
}

type DescribeClusterV2ResponseBody struct {
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ClusterInfo *DescribeClusterV2ResponseBodyClusterInfo `json:"ClusterInfo,omitempty" xml:"ClusterInfo,omitempty" type:"Struct"`
}

func (s DescribeClusterV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBody) SetRequestId(v string) *DescribeClusterV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterV2ResponseBody) SetClusterInfo(v *DescribeClusterV2ResponseBodyClusterInfo) *DescribeClusterV2ResponseBody {
	s.ClusterInfo = v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfo struct {
	VpcId                           *string                                                         `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	LogEnable                       *bool                                                           `json:"LogEnable,omitempty" xml:"LogEnable,omitempty"`
	TaskNodeInService               *int32                                                          `json:"TaskNodeInService,omitempty" xml:"TaskNodeInService,omitempty"`
	AutoScalingSpotWithLimitAllowed *bool                                                           `json:"AutoScalingSpotWithLimitAllowed,omitempty" xml:"AutoScalingSpotWithLimitAllowed,omitempty"`
	UserId                          *string                                                         `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ChargeType                      *string                                                         `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	StopTime                        *int64                                                          `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
	DepositType                     *string                                                         `json:"DepositType,omitempty" xml:"DepositType,omitempty"`
	CreateType                      *string                                                         `json:"CreateType,omitempty" xml:"CreateType,omitempty"`
	RelateClusterId                 *string                                                         `json:"RelateClusterId,omitempty" xml:"RelateClusterId,omitempty"`
	SecurityGroupName               *string                                                         `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
	ResizeDiskEnable                *bool                                                           `json:"ResizeDiskEnable,omitempty" xml:"ResizeDiskEnable,omitempty"`
	ImageId                         *string                                                         `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	UserDefinedEmrEcsRole           *string                                                         `json:"UserDefinedEmrEcsRole,omitempty" xml:"UserDefinedEmrEcsRole,omitempty"`
	MetaStoreType                   *string                                                         `json:"MetaStoreType,omitempty" xml:"MetaStoreType,omitempty"`
	StartTime                       *int64                                                          `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Configurations                  *string                                                         `json:"Configurations,omitempty" xml:"Configurations,omitempty"`
	LogPath                         *string                                                         `json:"LogPath,omitempty" xml:"LogPath,omitempty"`
	AutoScalingVersion              *string                                                         `json:"AutoScalingVersion,omitempty" xml:"AutoScalingVersion,omitempty"`
	NetType                         *string                                                         `json:"NetType,omitempty" xml:"NetType,omitempty"`
	ZoneId                          *string                                                         `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	CreateResource                  *string                                                         `json:"CreateResource,omitempty" xml:"CreateResource,omitempty"`
	Status                          *string                                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	RunningTime                     *int32                                                          `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	HighAvailabilityEnable          *bool                                                           `json:"HighAvailabilityEnable,omitempty" xml:"HighAvailabilityEnable,omitempty"`
	SecurityGroupId                 *string                                                         `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	AutoScalingAllowed              *bool                                                           `json:"AutoScalingAllowed,omitempty" xml:"AutoScalingAllowed,omitempty"`
	MasterNodeInService             *int32                                                          `json:"MasterNodeInService,omitempty" xml:"MasterNodeInService,omitempty"`
	AutoScalingEnable               *bool                                                           `json:"AutoScalingEnable,omitempty" xml:"AutoScalingEnable,omitempty"`
	AutoScalingWithGraceAllowed     *bool                                                           `json:"AutoScalingWithGraceAllowed,omitempty" xml:"AutoScalingWithGraceAllowed,omitempty"`
	CoreNodeInService               *int32                                                          `json:"CoreNodeInService,omitempty" xml:"CoreNodeInService,omitempty"`
	ShowSoftwareInterface           *bool                                                           `json:"ShowSoftwareInterface,omitempty" xml:"ShowSoftwareInterface,omitempty"`
	K8sClusterId                    *string                                                         `json:"K8sClusterId,omitempty" xml:"K8sClusterId,omitempty"`
	AutoScalingByLoadAllowed        *bool                                                           `json:"AutoScalingByLoadAllowed,omitempty" xml:"AutoScalingByLoadAllowed,omitempty"`
	LocalMetaDb                     *bool                                                           `json:"LocalMetaDb,omitempty" xml:"LocalMetaDb,omitempty"`
	InstanceGeneration              *string                                                         `json:"InstanceGeneration,omitempty" xml:"InstanceGeneration,omitempty"`
	Name                            *string                                                         `json:"Name,omitempty" xml:"Name,omitempty"`
	EasEnable                       *bool                                                           `json:"EasEnable,omitempty" xml:"EasEnable,omitempty"`
	MachineType                     *string                                                         `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	MasterNodeTotal                 *int32                                                          `json:"MasterNodeTotal,omitempty" xml:"MasterNodeTotal,omitempty"`
	RegionId                        *string                                                         `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Period                          *int32                                                          `json:"Period,omitempty" xml:"Period,omitempty"`
	ExtraInfo                       *string                                                         `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	IoOptimized                     *bool                                                           `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	VSwitchId                       *string                                                         `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ExpiredTime                     *int64                                                          `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	CoreNodeTotal                   *int32                                                          `json:"CoreNodeTotal,omitempty" xml:"CoreNodeTotal,omitempty"`
	GatewayClusterIds               *string                                                         `json:"GatewayClusterIds,omitempty" xml:"GatewayClusterIds,omitempty"`
	BootstrapFailed                 *bool                                                           `json:"BootstrapFailed,omitempty" xml:"BootstrapFailed,omitempty"`
	Id                              *string                                                         `json:"Id,omitempty" xml:"Id,omitempty"`
	TaskNodeTotal                   *int32                                                          `json:"TaskNodeTotal,omitempty" xml:"TaskNodeTotal,omitempty"`
	GatewayClusterInfoList          *DescribeClusterV2ResponseBodyClusterInfoGatewayClusterInfoList `json:"GatewayClusterInfoList,omitempty" xml:"GatewayClusterInfoList,omitempty" type:"Struct"`
	HostGroupList                   *DescribeClusterV2ResponseBodyClusterInfoHostGroupList          `json:"HostGroupList,omitempty" xml:"HostGroupList,omitempty" type:"Struct"`
	BootstrapActionList             *DescribeClusterV2ResponseBodyClusterInfoBootstrapActionList    `json:"BootstrapActionList,omitempty" xml:"BootstrapActionList,omitempty" type:"Struct"`
	RelateClusterInfo               *DescribeClusterV2ResponseBodyClusterInfoRelateClusterInfo      `json:"RelateClusterInfo,omitempty" xml:"RelateClusterInfo,omitempty" type:"Struct"`
	HostPoolInfo                    *DescribeClusterV2ResponseBodyClusterInfoHostPoolInfo           `json:"HostPoolInfo,omitempty" xml:"HostPoolInfo,omitempty" type:"Struct"`
	FailReason                      *DescribeClusterV2ResponseBodyClusterInfoFailReason             `json:"FailReason,omitempty" xml:"FailReason,omitempty" type:"Struct"`
	SoftwareInfo                    *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfo           `json:"SoftwareInfo,omitempty" xml:"SoftwareInfo,omitempty" type:"Struct"`
	AccessInfo                      *DescribeClusterV2ResponseBodyClusterInfoAccessInfo             `json:"AccessInfo,omitempty" xml:"AccessInfo,omitempty" type:"Struct"`
}

func (s DescribeClusterV2ResponseBodyClusterInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetVpcId(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.VpcId = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetLogEnable(v bool) *DescribeClusterV2ResponseBodyClusterInfo {
	s.LogEnable = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetTaskNodeInService(v int32) *DescribeClusterV2ResponseBodyClusterInfo {
	s.TaskNodeInService = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetAutoScalingSpotWithLimitAllowed(v bool) *DescribeClusterV2ResponseBodyClusterInfo {
	s.AutoScalingSpotWithLimitAllowed = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetUserId(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.UserId = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetChargeType(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.ChargeType = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetStopTime(v int64) *DescribeClusterV2ResponseBodyClusterInfo {
	s.StopTime = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetDepositType(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.DepositType = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetCreateType(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.CreateType = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetRelateClusterId(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.RelateClusterId = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetSecurityGroupName(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.SecurityGroupName = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetResizeDiskEnable(v bool) *DescribeClusterV2ResponseBodyClusterInfo {
	s.ResizeDiskEnable = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetImageId(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.ImageId = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetUserDefinedEmrEcsRole(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.UserDefinedEmrEcsRole = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetMetaStoreType(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.MetaStoreType = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetStartTime(v int64) *DescribeClusterV2ResponseBodyClusterInfo {
	s.StartTime = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetConfigurations(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.Configurations = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetLogPath(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.LogPath = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetAutoScalingVersion(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.AutoScalingVersion = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetNetType(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.NetType = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetZoneId(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.ZoneId = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetCreateResource(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.CreateResource = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetStatus(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.Status = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetRunningTime(v int32) *DescribeClusterV2ResponseBodyClusterInfo {
	s.RunningTime = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetHighAvailabilityEnable(v bool) *DescribeClusterV2ResponseBodyClusterInfo {
	s.HighAvailabilityEnable = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetSecurityGroupId(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetAutoScalingAllowed(v bool) *DescribeClusterV2ResponseBodyClusterInfo {
	s.AutoScalingAllowed = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetMasterNodeInService(v int32) *DescribeClusterV2ResponseBodyClusterInfo {
	s.MasterNodeInService = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetAutoScalingEnable(v bool) *DescribeClusterV2ResponseBodyClusterInfo {
	s.AutoScalingEnable = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetAutoScalingWithGraceAllowed(v bool) *DescribeClusterV2ResponseBodyClusterInfo {
	s.AutoScalingWithGraceAllowed = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetCoreNodeInService(v int32) *DescribeClusterV2ResponseBodyClusterInfo {
	s.CoreNodeInService = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetShowSoftwareInterface(v bool) *DescribeClusterV2ResponseBodyClusterInfo {
	s.ShowSoftwareInterface = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetK8sClusterId(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.K8sClusterId = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetAutoScalingByLoadAllowed(v bool) *DescribeClusterV2ResponseBodyClusterInfo {
	s.AutoScalingByLoadAllowed = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetLocalMetaDb(v bool) *DescribeClusterV2ResponseBodyClusterInfo {
	s.LocalMetaDb = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetInstanceGeneration(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.InstanceGeneration = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetName(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.Name = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetEasEnable(v bool) *DescribeClusterV2ResponseBodyClusterInfo {
	s.EasEnable = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetMachineType(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.MachineType = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetMasterNodeTotal(v int32) *DescribeClusterV2ResponseBodyClusterInfo {
	s.MasterNodeTotal = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetRegionId(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetPeriod(v int32) *DescribeClusterV2ResponseBodyClusterInfo {
	s.Period = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetExtraInfo(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.ExtraInfo = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetIoOptimized(v bool) *DescribeClusterV2ResponseBodyClusterInfo {
	s.IoOptimized = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetVSwitchId(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.VSwitchId = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetExpiredTime(v int64) *DescribeClusterV2ResponseBodyClusterInfo {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetCoreNodeTotal(v int32) *DescribeClusterV2ResponseBodyClusterInfo {
	s.CoreNodeTotal = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetGatewayClusterIds(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.GatewayClusterIds = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetBootstrapFailed(v bool) *DescribeClusterV2ResponseBodyClusterInfo {
	s.BootstrapFailed = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetId(v string) *DescribeClusterV2ResponseBodyClusterInfo {
	s.Id = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetTaskNodeTotal(v int32) *DescribeClusterV2ResponseBodyClusterInfo {
	s.TaskNodeTotal = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetGatewayClusterInfoList(v *DescribeClusterV2ResponseBodyClusterInfoGatewayClusterInfoList) *DescribeClusterV2ResponseBodyClusterInfo {
	s.GatewayClusterInfoList = v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetHostGroupList(v *DescribeClusterV2ResponseBodyClusterInfoHostGroupList) *DescribeClusterV2ResponseBodyClusterInfo {
	s.HostGroupList = v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetBootstrapActionList(v *DescribeClusterV2ResponseBodyClusterInfoBootstrapActionList) *DescribeClusterV2ResponseBodyClusterInfo {
	s.BootstrapActionList = v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetRelateClusterInfo(v *DescribeClusterV2ResponseBodyClusterInfoRelateClusterInfo) *DescribeClusterV2ResponseBodyClusterInfo {
	s.RelateClusterInfo = v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetHostPoolInfo(v *DescribeClusterV2ResponseBodyClusterInfoHostPoolInfo) *DescribeClusterV2ResponseBodyClusterInfo {
	s.HostPoolInfo = v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetFailReason(v *DescribeClusterV2ResponseBodyClusterInfoFailReason) *DescribeClusterV2ResponseBodyClusterInfo {
	s.FailReason = v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetSoftwareInfo(v *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfo) *DescribeClusterV2ResponseBodyClusterInfo {
	s.SoftwareInfo = v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfo) SetAccessInfo(v *DescribeClusterV2ResponseBodyClusterInfoAccessInfo) *DescribeClusterV2ResponseBodyClusterInfo {
	s.AccessInfo = v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoGatewayClusterInfoList struct {
	GatewayClusterInfo []*DescribeClusterV2ResponseBodyClusterInfoGatewayClusterInfoListGatewayClusterInfo `json:"GatewayClusterInfo,omitempty" xml:"GatewayClusterInfo,omitempty" type:"Repeated"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoGatewayClusterInfoList) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoGatewayClusterInfoList) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoGatewayClusterInfoList) SetGatewayClusterInfo(v []*DescribeClusterV2ResponseBodyClusterInfoGatewayClusterInfoListGatewayClusterInfo) *DescribeClusterV2ResponseBodyClusterInfoGatewayClusterInfoList {
	s.GatewayClusterInfo = v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoGatewayClusterInfoListGatewayClusterInfo struct {
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoGatewayClusterInfoListGatewayClusterInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoGatewayClusterInfoListGatewayClusterInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoGatewayClusterInfoListGatewayClusterInfo) SetClusterName(v string) *DescribeClusterV2ResponseBodyClusterInfoGatewayClusterInfoListGatewayClusterInfo {
	s.ClusterName = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoGatewayClusterInfoListGatewayClusterInfo) SetStatus(v string) *DescribeClusterV2ResponseBodyClusterInfoGatewayClusterInfoListGatewayClusterInfo {
	s.Status = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoGatewayClusterInfoListGatewayClusterInfo) SetClusterId(v string) *DescribeClusterV2ResponseBodyClusterInfoGatewayClusterInfoListGatewayClusterInfo {
	s.ClusterId = &v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoHostGroupList struct {
	HostGroup []*DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup `json:"HostGroup,omitempty" xml:"HostGroup,omitempty" type:"Repeated"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoHostGroupList) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoHostGroupList) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupList) SetHostGroup(v []*DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) *DescribeClusterV2ResponseBodyClusterInfoHostGroupList {
	s.HostGroup = v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup struct {
	LockType              *string                                                              `json:"LockType,omitempty" xml:"LockType,omitempty"`
	HostGroupSubType      *string                                                              `json:"HostGroupSubType,omitempty" xml:"HostGroupSubType,omitempty"`
	HostGroupType         *string                                                              `json:"HostGroupType,omitempty" xml:"HostGroupType,omitempty"`
	HostGroupChangeStatus *string                                                              `json:"HostGroupChangeStatus,omitempty" xml:"HostGroupChangeStatus,omitempty"`
	ChargeType            *string                                                              `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	LockReason            *string                                                              `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	DiskType              *string                                                              `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	HostGroupId           *string                                                              `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	BandWidth             *string                                                              `json:"BandWidth,omitempty" xml:"BandWidth,omitempty"`
	InstanceType          *string                                                              `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	DiskCount             *int32                                                               `json:"DiskCount,omitempty" xml:"DiskCount,omitempty"`
	Period                *string                                                              `json:"Period,omitempty" xml:"Period,omitempty"`
	DiskCapacity          *int32                                                               `json:"DiskCapacity,omitempty" xml:"DiskCapacity,omitempty"`
	CpuCore               *int32                                                               `json:"CpuCore,omitempty" xml:"CpuCore,omitempty"`
	MemoryCapacity        *int32                                                               `json:"MemoryCapacity,omitempty" xml:"MemoryCapacity,omitempty"`
	NodeCount             *int32                                                               `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	HostGroupChangeType   *string                                                              `json:"HostGroupChangeType,omitempty" xml:"HostGroupChangeType,omitempty"`
	HostGroupName         *string                                                              `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
	Nodes                 *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Struct"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) SetLockType(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup {
	s.LockType = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) SetHostGroupSubType(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup {
	s.HostGroupSubType = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) SetHostGroupType(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup {
	s.HostGroupType = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) SetHostGroupChangeStatus(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup {
	s.HostGroupChangeStatus = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) SetChargeType(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup {
	s.ChargeType = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) SetLockReason(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup {
	s.LockReason = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) SetDiskType(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup {
	s.DiskType = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) SetHostGroupId(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup {
	s.HostGroupId = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) SetBandWidth(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup {
	s.BandWidth = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) SetInstanceType(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup {
	s.InstanceType = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) SetDiskCount(v int32) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup {
	s.DiskCount = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) SetPeriod(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup {
	s.Period = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) SetDiskCapacity(v int32) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup {
	s.DiskCapacity = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) SetCpuCore(v int32) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup {
	s.CpuCore = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) SetMemoryCapacity(v int32) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup {
	s.MemoryCapacity = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) SetNodeCount(v int32) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup {
	s.NodeCount = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) SetHostGroupChangeType(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup {
	s.HostGroupChangeType = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) SetHostGroupName(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup {
	s.HostGroupName = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup) SetNodes(v *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodes) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroup {
	s.Nodes = v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodes struct {
	Node []*DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode `json:"Node,omitempty" xml:"Node,omitempty" type:"Repeated"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodes) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodes) SetNode(v []*DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodes {
	s.Node = v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode struct {
	Status         *string                                                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	SupportIpV6    *bool                                                                               `json:"SupportIpV6,omitempty" xml:"SupportIpV6,omitempty"`
	InnerIp        *string                                                                             `json:"InnerIp,omitempty" xml:"InnerIp,omitempty"`
	ExpiredTime    *string                                                                             `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	CreateTime     *string                                                                             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ZoneId         *string                                                                             `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	InstanceId     *string                                                                             `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	EmrExpiredTime *string                                                                             `json:"EmrExpiredTime,omitempty" xml:"EmrExpiredTime,omitempty"`
	PubIp          *string                                                                             `json:"PubIp,omitempty" xml:"PubIp,omitempty"`
	DaemonInfos    *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDaemonInfos `json:"DaemonInfos,omitempty" xml:"DaemonInfos,omitempty" type:"Struct"`
	DiskInfos      *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfos   `json:"DiskInfos,omitempty" xml:"DiskInfos,omitempty" type:"Struct"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode) SetStatus(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode {
	s.Status = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode) SetSupportIpV6(v bool) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode {
	s.SupportIpV6 = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode) SetInnerIp(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode {
	s.InnerIp = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode) SetExpiredTime(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode) SetCreateTime(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode {
	s.CreateTime = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode) SetZoneId(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode {
	s.ZoneId = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode) SetInstanceId(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode {
	s.InstanceId = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode) SetEmrExpiredTime(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode {
	s.EmrExpiredTime = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode) SetPubIp(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode {
	s.PubIp = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode) SetDaemonInfos(v *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDaemonInfos) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode {
	s.DaemonInfos = v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode) SetDiskInfos(v *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfos) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNode {
	s.DiskInfos = v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDaemonInfos struct {
	DaemonInfo []*DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDaemonInfosDaemonInfo `json:"DaemonInfo,omitempty" xml:"DaemonInfo,omitempty" type:"Repeated"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDaemonInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDaemonInfos) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDaemonInfos) SetDaemonInfo(v []*DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDaemonInfosDaemonInfo) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDaemonInfos {
	s.DaemonInfo = v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDaemonInfosDaemonInfo struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDaemonInfosDaemonInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDaemonInfosDaemonInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDaemonInfosDaemonInfo) SetName(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDaemonInfosDaemonInfo {
	s.Name = &v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfos struct {
	DiskInfo []*DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfosDiskInfo `json:"DiskInfo,omitempty" xml:"DiskInfo,omitempty" type:"Repeated"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfos) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfos) SetDiskInfo(v []*DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfosDiskInfo) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfos {
	s.DiskInfo = v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfosDiskInfo struct {
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	DiskId   *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	Size     *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Device   *string `json:"Device,omitempty" xml:"Device,omitempty"`
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfosDiskInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfosDiskInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfosDiskInfo) SetType(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfosDiskInfo {
	s.Type = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfosDiskInfo) SetDiskId(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfosDiskInfo {
	s.DiskId = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfosDiskInfo) SetSize(v int32) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfosDiskInfo {
	s.Size = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfosDiskInfo) SetDevice(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfosDiskInfo {
	s.Device = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfosDiskInfo) SetDiskName(v string) *DescribeClusterV2ResponseBodyClusterInfoHostGroupListHostGroupNodesNodeDiskInfosDiskInfo {
	s.DiskName = &v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoBootstrapActionList struct {
	BootstrapAction []*DescribeClusterV2ResponseBodyClusterInfoBootstrapActionListBootstrapAction `json:"BootstrapAction,omitempty" xml:"BootstrapAction,omitempty" type:"Repeated"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoBootstrapActionList) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoBootstrapActionList) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoBootstrapActionList) SetBootstrapAction(v []*DescribeClusterV2ResponseBodyClusterInfoBootstrapActionListBootstrapAction) *DescribeClusterV2ResponseBodyClusterInfoBootstrapActionList {
	s.BootstrapAction = v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoBootstrapActionListBootstrapAction struct {
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Arg  *string `json:"Arg,omitempty" xml:"Arg,omitempty"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoBootstrapActionListBootstrapAction) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoBootstrapActionListBootstrapAction) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoBootstrapActionListBootstrapAction) SetPath(v string) *DescribeClusterV2ResponseBodyClusterInfoBootstrapActionListBootstrapAction {
	s.Path = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoBootstrapActionListBootstrapAction) SetName(v string) *DescribeClusterV2ResponseBodyClusterInfoBootstrapActionListBootstrapAction {
	s.Name = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoBootstrapActionListBootstrapAction) SetArg(v string) *DescribeClusterV2ResponseBodyClusterInfoBootstrapActionListBootstrapAction {
	s.Arg = &v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoRelateClusterInfo struct {
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoRelateClusterInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoRelateClusterInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoRelateClusterInfo) SetClusterName(v string) *DescribeClusterV2ResponseBodyClusterInfoRelateClusterInfo {
	s.ClusterName = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoRelateClusterInfo) SetStatus(v string) *DescribeClusterV2ResponseBodyClusterInfoRelateClusterInfo {
	s.Status = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoRelateClusterInfo) SetClusterId(v string) *DescribeClusterV2ResponseBodyClusterInfoRelateClusterInfo {
	s.ClusterId = &v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoHostPoolInfo struct {
	HpName  *string `json:"HpName,omitempty" xml:"HpName,omitempty"`
	HpBizId *string `json:"HpBizId,omitempty" xml:"HpBizId,omitempty"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoHostPoolInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoHostPoolInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostPoolInfo) SetHpName(v string) *DescribeClusterV2ResponseBodyClusterInfoHostPoolInfo {
	s.HpName = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoHostPoolInfo) SetHpBizId(v string) *DescribeClusterV2ResponseBodyClusterInfoHostPoolInfo {
	s.HpBizId = &v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoFailReason struct {
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoFailReason) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoFailReason) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoFailReason) SetErrorMsg(v string) *DescribeClusterV2ResponseBodyClusterInfoFailReason {
	s.ErrorMsg = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoFailReason) SetRequestId(v string) *DescribeClusterV2ResponseBodyClusterInfoFailReason {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoFailReason) SetErrorCode(v string) *DescribeClusterV2ResponseBodyClusterInfoFailReason {
	s.ErrorCode = &v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoSoftwareInfo struct {
	ClusterType *string                                                        `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	EmrVer      *string                                                        `json:"EmrVer,omitempty" xml:"EmrVer,omitempty"`
	Softwares   *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwares `json:"Softwares,omitempty" xml:"Softwares,omitempty" type:"Struct"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoSoftwareInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoSoftwareInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfo) SetClusterType(v string) *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfo {
	s.ClusterType = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfo) SetEmrVer(v string) *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfo {
	s.EmrVer = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfo) SetSoftwares(v *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwares) *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfo {
	s.Softwares = v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwares struct {
	Software []*DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwaresSoftware `json:"Software,omitempty" xml:"Software,omitempty" type:"Repeated"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwares) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwares) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwares) SetSoftware(v []*DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwaresSoftware) *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwares {
	s.Software = v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwaresSoftware struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	OnlyDisplay *bool   `json:"OnlyDisplay,omitempty" xml:"OnlyDisplay,omitempty"`
	StartTpe    *int32  `json:"StartTpe,omitempty" xml:"StartTpe,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwaresSoftware) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwaresSoftware) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwaresSoftware) SetDisplayName(v string) *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwaresSoftware {
	s.DisplayName = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwaresSoftware) SetOnlyDisplay(v bool) *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwaresSoftware {
	s.OnlyDisplay = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwaresSoftware) SetStartTpe(v int32) *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwaresSoftware {
	s.StartTpe = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwaresSoftware) SetName(v string) *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwaresSoftware {
	s.Name = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwaresSoftware) SetVersion(v string) *DescribeClusterV2ResponseBodyClusterInfoSoftwareInfoSoftwaresSoftware {
	s.Version = &v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoAccessInfo struct {
	ZKLinks *DescribeClusterV2ResponseBodyClusterInfoAccessInfoZKLinks `json:"ZKLinks,omitempty" xml:"ZKLinks,omitempty" type:"Struct"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoAccessInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoAccessInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoAccessInfo) SetZKLinks(v *DescribeClusterV2ResponseBodyClusterInfoAccessInfoZKLinks) *DescribeClusterV2ResponseBodyClusterInfoAccessInfo {
	s.ZKLinks = v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoAccessInfoZKLinks struct {
	ZKLink []*DescribeClusterV2ResponseBodyClusterInfoAccessInfoZKLinksZKLink `json:"ZKLink,omitempty" xml:"ZKLink,omitempty" type:"Repeated"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoAccessInfoZKLinks) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoAccessInfoZKLinks) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoAccessInfoZKLinks) SetZKLink(v []*DescribeClusterV2ResponseBodyClusterInfoAccessInfoZKLinksZKLink) *DescribeClusterV2ResponseBodyClusterInfoAccessInfoZKLinks {
	s.ZKLink = v
	return s
}

type DescribeClusterV2ResponseBodyClusterInfoAccessInfoZKLinksZKLink struct {
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	Link *string `json:"Link,omitempty" xml:"Link,omitempty"`
}

func (s DescribeClusterV2ResponseBodyClusterInfoAccessInfoZKLinksZKLink) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2ResponseBodyClusterInfoAccessInfoZKLinksZKLink) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2ResponseBodyClusterInfoAccessInfoZKLinksZKLink) SetPort(v string) *DescribeClusterV2ResponseBodyClusterInfoAccessInfoZKLinksZKLink {
	s.Port = &v
	return s
}

func (s *DescribeClusterV2ResponseBodyClusterInfoAccessInfoZKLinksZKLink) SetLink(v string) *DescribeClusterV2ResponseBodyClusterInfoAccessInfoZKLinksZKLink {
	s.Link = &v
	return s
}

type DescribeClusterV2Response struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClusterV2ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClusterV2Response) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterV2Response) GoString() string {
	return s.String()
}

func (s *DescribeClusterV2Response) SetHeaders(v map[string]*string) *DescribeClusterV2Response {
	s.Headers = v
	return s
}

func (s *DescribeClusterV2Response) SetBody(v *DescribeClusterV2ResponseBody) *DescribeClusterV2Response {
	s.Body = v
	return s
}

type DescribeFlowRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowRequest) SetProjectId(v string) *DescribeFlowRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeFlowRequest) SetId(v string) *DescribeFlowRequest {
	s.Id = &v
	return s
}

func (s *DescribeFlowRequest) SetRegionId(v string) *DescribeFlowRequest {
	s.RegionId = &v
	return s
}

type DescribeFlowResponseBody struct {
	Status                  *string                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	Type                    *string                                 `json:"Type,omitempty" xml:"Type,omitempty"`
	AlertUserGroupBizId     *string                                 `json:"AlertUserGroupBizId,omitempty" xml:"AlertUserGroupBizId,omitempty"`
	Periodic                *bool                                   `json:"Periodic,omitempty" xml:"Periodic,omitempty"`
	EditLockDetail          *string                                 `json:"EditLockDetail,omitempty" xml:"EditLockDetail,omitempty"`
	Namespace               *string                                 `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	HostName                *string                                 `json:"HostName,omitempty" xml:"HostName,omitempty"`
	Application             *string                                 `json:"Application,omitempty" xml:"Application,omitempty"`
	GmtModified             *int64                                  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	RequestId               *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Description             *string                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	CreateCluster           *bool                                   `json:"CreateCluster,omitempty" xml:"CreateCluster,omitempty"`
	StartSchedule           *int64                                  `json:"StartSchedule,omitempty" xml:"StartSchedule,omitempty"`
	EndSchedule             *int64                                  `json:"EndSchedule,omitempty" xml:"EndSchedule,omitempty"`
	Graph                   *string                                 `json:"Graph,omitempty" xml:"Graph,omitempty"`
	AlertDingDingGroupBizId *string                                 `json:"AlertDingDingGroupBizId,omitempty" xml:"AlertDingDingGroupBizId,omitempty"`
	GmtCreate               *int64                                  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	CategoryId              *string                                 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CronExpr                *string                                 `json:"CronExpr,omitempty" xml:"CronExpr,omitempty"`
	Name                    *string                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	Id                      *string                                 `json:"Id,omitempty" xml:"Id,omitempty"`
	AlertConf               *string                                 `json:"AlertConf,omitempty" xml:"AlertConf,omitempty"`
	ClusterId               *string                                 `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ParentFlowList          *DescribeFlowResponseBodyParentFlowList `json:"ParentFlowList,omitempty" xml:"ParentFlowList,omitempty" type:"Struct"`
}

func (s DescribeFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowResponseBody) SetStatus(v string) *DescribeFlowResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeFlowResponseBody) SetType(v string) *DescribeFlowResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeFlowResponseBody) SetAlertUserGroupBizId(v string) *DescribeFlowResponseBody {
	s.AlertUserGroupBizId = &v
	return s
}

func (s *DescribeFlowResponseBody) SetPeriodic(v bool) *DescribeFlowResponseBody {
	s.Periodic = &v
	return s
}

func (s *DescribeFlowResponseBody) SetEditLockDetail(v string) *DescribeFlowResponseBody {
	s.EditLockDetail = &v
	return s
}

func (s *DescribeFlowResponseBody) SetNamespace(v string) *DescribeFlowResponseBody {
	s.Namespace = &v
	return s
}

func (s *DescribeFlowResponseBody) SetHostName(v string) *DescribeFlowResponseBody {
	s.HostName = &v
	return s
}

func (s *DescribeFlowResponseBody) SetApplication(v string) *DescribeFlowResponseBody {
	s.Application = &v
	return s
}

func (s *DescribeFlowResponseBody) SetGmtModified(v int64) *DescribeFlowResponseBody {
	s.GmtModified = &v
	return s
}

func (s *DescribeFlowResponseBody) SetRequestId(v string) *DescribeFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowResponseBody) SetDescription(v string) *DescribeFlowResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeFlowResponseBody) SetCreateCluster(v bool) *DescribeFlowResponseBody {
	s.CreateCluster = &v
	return s
}

func (s *DescribeFlowResponseBody) SetStartSchedule(v int64) *DescribeFlowResponseBody {
	s.StartSchedule = &v
	return s
}

func (s *DescribeFlowResponseBody) SetEndSchedule(v int64) *DescribeFlowResponseBody {
	s.EndSchedule = &v
	return s
}

func (s *DescribeFlowResponseBody) SetGraph(v string) *DescribeFlowResponseBody {
	s.Graph = &v
	return s
}

func (s *DescribeFlowResponseBody) SetAlertDingDingGroupBizId(v string) *DescribeFlowResponseBody {
	s.AlertDingDingGroupBizId = &v
	return s
}

func (s *DescribeFlowResponseBody) SetGmtCreate(v int64) *DescribeFlowResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *DescribeFlowResponseBody) SetCategoryId(v string) *DescribeFlowResponseBody {
	s.CategoryId = &v
	return s
}

func (s *DescribeFlowResponseBody) SetCronExpr(v string) *DescribeFlowResponseBody {
	s.CronExpr = &v
	return s
}

func (s *DescribeFlowResponseBody) SetName(v string) *DescribeFlowResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeFlowResponseBody) SetId(v string) *DescribeFlowResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeFlowResponseBody) SetAlertConf(v string) *DescribeFlowResponseBody {
	s.AlertConf = &v
	return s
}

func (s *DescribeFlowResponseBody) SetClusterId(v string) *DescribeFlowResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeFlowResponseBody) SetParentFlowList(v *DescribeFlowResponseBodyParentFlowList) *DescribeFlowResponseBody {
	s.ParentFlowList = v
	return s
}

type DescribeFlowResponseBodyParentFlowList struct {
	ParentFlow []*DescribeFlowResponseBodyParentFlowListParentFlow `json:"ParentFlow,omitempty" xml:"ParentFlow,omitempty" type:"Repeated"`
}

func (s DescribeFlowResponseBodyParentFlowList) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowResponseBodyParentFlowList) GoString() string {
	return s.String()
}

func (s *DescribeFlowResponseBodyParentFlowList) SetParentFlow(v []*DescribeFlowResponseBodyParentFlowListParentFlow) *DescribeFlowResponseBodyParentFlowList {
	s.ParentFlow = v
	return s
}

type DescribeFlowResponseBodyParentFlowListParentFlow struct {
	ParentFlowName *string `json:"ParentFlowName,omitempty" xml:"ParentFlowName,omitempty"`
	ParentFlowId   *string `json:"ParentFlowId,omitempty" xml:"ParentFlowId,omitempty"`
	ProjectName    *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ProjectId      *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DescribeFlowResponseBodyParentFlowListParentFlow) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowResponseBodyParentFlowListParentFlow) GoString() string {
	return s.String()
}

func (s *DescribeFlowResponseBodyParentFlowListParentFlow) SetParentFlowName(v string) *DescribeFlowResponseBodyParentFlowListParentFlow {
	s.ParentFlowName = &v
	return s
}

func (s *DescribeFlowResponseBodyParentFlowListParentFlow) SetParentFlowId(v string) *DescribeFlowResponseBodyParentFlowListParentFlow {
	s.ParentFlowId = &v
	return s
}

func (s *DescribeFlowResponseBodyParentFlowListParentFlow) SetProjectName(v string) *DescribeFlowResponseBodyParentFlowListParentFlow {
	s.ProjectName = &v
	return s
}

func (s *DescribeFlowResponseBodyParentFlowListParentFlow) SetProjectId(v string) *DescribeFlowResponseBodyParentFlowListParentFlow {
	s.ProjectId = &v
	return s
}

type DescribeFlowResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowResponse) SetHeaders(v map[string]*string) *DescribeFlowResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowResponse) SetBody(v *DescribeFlowResponseBody) *DescribeFlowResponse {
	s.Body = v
	return s
}

type ListFlowClusterRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProjectId       *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListFlowClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterRequest) GoString() string {
	return s.String()
}

func (s *ListFlowClusterRequest) SetRegionId(v string) *ListFlowClusterRequest {
	s.RegionId = &v
	return s
}

func (s *ListFlowClusterRequest) SetProjectId(v string) *ListFlowClusterRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFlowClusterRequest) SetPageNumber(v int32) *ListFlowClusterRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFlowClusterRequest) SetPageSize(v int32) *ListFlowClusterRequest {
	s.PageSize = &v
	return s
}

func (s *ListFlowClusterRequest) SetResourceGroupId(v string) *ListFlowClusterRequest {
	s.ResourceGroupId = &v
	return s
}

type ListFlowClusterResponseBody struct {
	PageSize   *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount *int32                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Clusters   *ListFlowClusterResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Struct"`
}

func (s ListFlowClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowClusterResponseBody) SetPageSize(v int32) *ListFlowClusterResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListFlowClusterResponseBody) SetRequestId(v string) *ListFlowClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowClusterResponseBody) SetPageNumber(v int32) *ListFlowClusterResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListFlowClusterResponseBody) SetTotalCount(v int32) *ListFlowClusterResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListFlowClusterResponseBody) SetClusters(v *ListFlowClusterResponseBodyClusters) *ListFlowClusterResponseBody {
	s.Clusters = v
	return s
}

type ListFlowClusterResponseBodyClusters struct {
	ClusterInfo []*ListFlowClusterResponseBodyClustersClusterInfo `json:"ClusterInfo,omitempty" xml:"ClusterInfo,omitempty" type:"Repeated"`
}

func (s ListFlowClusterResponseBodyClusters) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *ListFlowClusterResponseBodyClusters) SetClusterInfo(v []*ListFlowClusterResponseBodyClustersClusterInfo) *ListFlowClusterResponseBodyClusters {
	s.ClusterInfo = v
	return s
}

type ListFlowClusterResponseBodyClustersClusterInfo struct {
	Type                *string                                                      `json:"Type,omitempty" xml:"Type,omitempty"`
	Status              *string                                                      `json:"Status,omitempty" xml:"Status,omitempty"`
	RunningTime         *int32                                                       `json:"RunningTime,omitempty" xml:"RunningTime,omitempty"`
	OrderList           *string                                                      `json:"OrderList,omitempty" xml:"OrderList,omitempty"`
	CreateTime          *int64                                                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ChargeType          *string                                                      `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Period              *int32                                                       `json:"Period,omitempty" xml:"Period,omitempty"`
	K8sClusterId        *string                                                      `json:"K8sClusterId,omitempty" xml:"K8sClusterId,omitempty"`
	ExpiredTime         *int64                                                       `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	HasUncompletedOrder *bool                                                        `json:"HasUncompletedOrder,omitempty" xml:"HasUncompletedOrder,omitempty"`
	Name                *string                                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	Id                  *string                                                      `json:"Id,omitempty" xml:"Id,omitempty"`
	CreateResource      *string                                                      `json:"CreateResource,omitempty" xml:"CreateResource,omitempty"`
	OrderTaskInfo       *ListFlowClusterResponseBodyClustersClusterInfoOrderTaskInfo `json:"OrderTaskInfo,omitempty" xml:"OrderTaskInfo,omitempty" type:"Struct"`
	FailReason          *ListFlowClusterResponseBodyClustersClusterInfoFailReason    `json:"FailReason,omitempty" xml:"FailReason,omitempty" type:"Struct"`
}

func (s ListFlowClusterResponseBodyClustersClusterInfo) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterResponseBodyClustersClusterInfo) GoString() string {
	return s.String()
}

func (s *ListFlowClusterResponseBodyClustersClusterInfo) SetType(v string) *ListFlowClusterResponseBodyClustersClusterInfo {
	s.Type = &v
	return s
}

func (s *ListFlowClusterResponseBodyClustersClusterInfo) SetStatus(v string) *ListFlowClusterResponseBodyClustersClusterInfo {
	s.Status = &v
	return s
}

func (s *ListFlowClusterResponseBodyClustersClusterInfo) SetRunningTime(v int32) *ListFlowClusterResponseBodyClustersClusterInfo {
	s.RunningTime = &v
	return s
}

func (s *ListFlowClusterResponseBodyClustersClusterInfo) SetOrderList(v string) *ListFlowClusterResponseBodyClustersClusterInfo {
	s.OrderList = &v
	return s
}

func (s *ListFlowClusterResponseBodyClustersClusterInfo) SetCreateTime(v int64) *ListFlowClusterResponseBodyClustersClusterInfo {
	s.CreateTime = &v
	return s
}

func (s *ListFlowClusterResponseBodyClustersClusterInfo) SetChargeType(v string) *ListFlowClusterResponseBodyClustersClusterInfo {
	s.ChargeType = &v
	return s
}

func (s *ListFlowClusterResponseBodyClustersClusterInfo) SetPeriod(v int32) *ListFlowClusterResponseBodyClustersClusterInfo {
	s.Period = &v
	return s
}

func (s *ListFlowClusterResponseBodyClustersClusterInfo) SetK8sClusterId(v string) *ListFlowClusterResponseBodyClustersClusterInfo {
	s.K8sClusterId = &v
	return s
}

func (s *ListFlowClusterResponseBodyClustersClusterInfo) SetExpiredTime(v int64) *ListFlowClusterResponseBodyClustersClusterInfo {
	s.ExpiredTime = &v
	return s
}

func (s *ListFlowClusterResponseBodyClustersClusterInfo) SetHasUncompletedOrder(v bool) *ListFlowClusterResponseBodyClustersClusterInfo {
	s.HasUncompletedOrder = &v
	return s
}

func (s *ListFlowClusterResponseBodyClustersClusterInfo) SetName(v string) *ListFlowClusterResponseBodyClustersClusterInfo {
	s.Name = &v
	return s
}

func (s *ListFlowClusterResponseBodyClustersClusterInfo) SetId(v string) *ListFlowClusterResponseBodyClustersClusterInfo {
	s.Id = &v
	return s
}

func (s *ListFlowClusterResponseBodyClustersClusterInfo) SetCreateResource(v string) *ListFlowClusterResponseBodyClustersClusterInfo {
	s.CreateResource = &v
	return s
}

func (s *ListFlowClusterResponseBodyClustersClusterInfo) SetOrderTaskInfo(v *ListFlowClusterResponseBodyClustersClusterInfoOrderTaskInfo) *ListFlowClusterResponseBodyClustersClusterInfo {
	s.OrderTaskInfo = v
	return s
}

func (s *ListFlowClusterResponseBodyClustersClusterInfo) SetFailReason(v *ListFlowClusterResponseBodyClustersClusterInfoFailReason) *ListFlowClusterResponseBodyClustersClusterInfo {
	s.FailReason = v
	return s
}

type ListFlowClusterResponseBodyClustersClusterInfoOrderTaskInfo struct {
	TargetCount  *int32  `json:"TargetCount,omitempty" xml:"TargetCount,omitempty"`
	CurrentCount *int32  `json:"CurrentCount,omitempty" xml:"CurrentCount,omitempty"`
	OrderIdList  *string `json:"OrderIdList,omitempty" xml:"OrderIdList,omitempty"`
}

func (s ListFlowClusterResponseBodyClustersClusterInfoOrderTaskInfo) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterResponseBodyClustersClusterInfoOrderTaskInfo) GoString() string {
	return s.String()
}

func (s *ListFlowClusterResponseBodyClustersClusterInfoOrderTaskInfo) SetTargetCount(v int32) *ListFlowClusterResponseBodyClustersClusterInfoOrderTaskInfo {
	s.TargetCount = &v
	return s
}

func (s *ListFlowClusterResponseBodyClustersClusterInfoOrderTaskInfo) SetCurrentCount(v int32) *ListFlowClusterResponseBodyClustersClusterInfoOrderTaskInfo {
	s.CurrentCount = &v
	return s
}

func (s *ListFlowClusterResponseBodyClustersClusterInfoOrderTaskInfo) SetOrderIdList(v string) *ListFlowClusterResponseBodyClustersClusterInfoOrderTaskInfo {
	s.OrderIdList = &v
	return s
}

type ListFlowClusterResponseBodyClustersClusterInfoFailReason struct {
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
}

func (s ListFlowClusterResponseBodyClustersClusterInfoFailReason) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterResponseBodyClustersClusterInfoFailReason) GoString() string {
	return s.String()
}

func (s *ListFlowClusterResponseBodyClustersClusterInfoFailReason) SetErrorMsg(v string) *ListFlowClusterResponseBodyClustersClusterInfoFailReason {
	s.ErrorMsg = &v
	return s
}

func (s *ListFlowClusterResponseBodyClustersClusterInfoFailReason) SetRequestId(v string) *ListFlowClusterResponseBodyClustersClusterInfoFailReason {
	s.RequestId = &v
	return s
}

func (s *ListFlowClusterResponseBodyClustersClusterInfoFailReason) SetErrorCode(v string) *ListFlowClusterResponseBodyClustersClusterInfoFailReason {
	s.ErrorCode = &v
	return s
}

type ListFlowClusterResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFlowClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowClusterResponse) GoString() string {
	return s.String()
}

func (s *ListFlowClusterResponse) SetHeaders(v map[string]*string) *ListFlowClusterResponse {
	s.Headers = v
	return s
}

func (s *ListFlowClusterResponse) SetBody(v *ListFlowClusterResponseBody) *ListFlowClusterResponse {
	s.Body = v
	return s
}

type ListLdapUsersRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	UserNameSearch  *string `json:"UserNameSearch,omitempty" xml:"UserNameSearch,omitempty"`
}

func (s ListLdapUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLdapUsersRequest) GoString() string {
	return s.String()
}

func (s *ListLdapUsersRequest) SetResourceOwnerId(v int64) *ListLdapUsersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListLdapUsersRequest) SetRegionId(v string) *ListLdapUsersRequest {
	s.RegionId = &v
	return s
}

func (s *ListLdapUsersRequest) SetClusterId(v string) *ListLdapUsersRequest {
	s.ClusterId = &v
	return s
}

func (s *ListLdapUsersRequest) SetUserNameSearch(v string) *ListLdapUsersRequest {
	s.UserNameSearch = &v
	return s
}

type ListLdapUsersResponseBody struct {
	IsAdmin   *bool                              `json:"IsAdmin,omitempty" xml:"IsAdmin,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserList  *ListLdapUsersResponseBodyUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Struct"`
}

func (s ListLdapUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLdapUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListLdapUsersResponseBody) SetIsAdmin(v bool) *ListLdapUsersResponseBody {
	s.IsAdmin = &v
	return s
}

func (s *ListLdapUsersResponseBody) SetRequestId(v string) *ListLdapUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLdapUsersResponseBody) SetUserList(v *ListLdapUsersResponseBodyUserList) *ListLdapUsersResponseBody {
	s.UserList = v
	return s
}

type ListLdapUsersResponseBodyUserList struct {
	User []*ListLdapUsersResponseBodyUserListUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s ListLdapUsersResponseBodyUserList) String() string {
	return tea.Prettify(s)
}

func (s ListLdapUsersResponseBodyUserList) GoString() string {
	return s.String()
}

func (s *ListLdapUsersResponseBodyUserList) SetUser(v []*ListLdapUsersResponseBodyUserListUser) *ListLdapUsersResponseBodyUserList {
	s.User = v
	return s
}

type ListLdapUsersResponseBodyUserListUser struct {
	UserCreateTime *int64  `json:"UserCreateTime,omitempty" xml:"UserCreateTime,omitempty"`
	GroupName      *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Note           *string `json:"Note,omitempty" xml:"Note,omitempty"`
	KeytabHex      *string `json:"KeytabHex,omitempty" xml:"KeytabHex,omitempty"`
	UserName       *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ListLdapUsersResponseBodyUserListUser) String() string {
	return tea.Prettify(s)
}

func (s ListLdapUsersResponseBodyUserListUser) GoString() string {
	return s.String()
}

func (s *ListLdapUsersResponseBodyUserListUser) SetUserCreateTime(v int64) *ListLdapUsersResponseBodyUserListUser {
	s.UserCreateTime = &v
	return s
}

func (s *ListLdapUsersResponseBodyUserListUser) SetGroupName(v string) *ListLdapUsersResponseBodyUserListUser {
	s.GroupName = &v
	return s
}

func (s *ListLdapUsersResponseBodyUserListUser) SetUserId(v string) *ListLdapUsersResponseBodyUserListUser {
	s.UserId = &v
	return s
}

func (s *ListLdapUsersResponseBodyUserListUser) SetNote(v string) *ListLdapUsersResponseBodyUserListUser {
	s.Note = &v
	return s
}

func (s *ListLdapUsersResponseBodyUserListUser) SetKeytabHex(v string) *ListLdapUsersResponseBodyUserListUser {
	s.KeytabHex = &v
	return s
}

func (s *ListLdapUsersResponseBodyUserListUser) SetUserName(v string) *ListLdapUsersResponseBodyUserListUser {
	s.UserName = &v
	return s
}

type ListLdapUsersResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListLdapUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListLdapUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLdapUsersResponse) GoString() string {
	return s.String()
}

func (s *ListLdapUsersResponse) SetHeaders(v map[string]*string) *ListLdapUsersResponse {
	s.Headers = v
	return s
}

func (s *ListLdapUsersResponse) SetBody(v *ListLdapUsersResponseBody) *ListLdapUsersResponse {
	s.Body = v
	return s
}

type DeleteUserRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	UserId          *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Type            *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DeleteUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserRequest) SetResourceOwnerId(v int64) *DeleteUserRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteUserRequest) SetRegionId(v string) *DeleteUserRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteUserRequest) SetClusterId(v string) *DeleteUserRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteUserRequest) SetUserId(v string) *DeleteUserRequest {
	s.UserId = &v
	return s
}

func (s *DeleteUserRequest) SetType(v string) *DeleteUserRequest {
	s.Type = &v
	return s
}

type DeleteUserResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserResponseBody) SetRequestId(v string) *DeleteUserResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUserResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserResponse) SetHeaders(v map[string]*string) *DeleteUserResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserResponse) SetBody(v *DeleteUserResponseBody) *DeleteUserResponse {
	s.Body = v
	return s
}

type CreateFlowProjectClusterSettingRequest struct {
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProjectId    *string   `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ClusterId    *string   `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DefaultUser  *string   `json:"DefaultUser,omitempty" xml:"DefaultUser,omitempty"`
	DefaultQueue *string   `json:"DefaultQueue,omitempty" xml:"DefaultQueue,omitempty"`
	ClientToken  *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	UserList     []*string `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
	QueueList    []*string `json:"QueueList,omitempty" xml:"QueueList,omitempty" type:"Repeated"`
	HostList     []*string `json:"HostList,omitempty" xml:"HostList,omitempty" type:"Repeated"`
}

func (s CreateFlowProjectClusterSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowProjectClusterSettingRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowProjectClusterSettingRequest) SetRegionId(v string) *CreateFlowProjectClusterSettingRequest {
	s.RegionId = &v
	return s
}

func (s *CreateFlowProjectClusterSettingRequest) SetProjectId(v string) *CreateFlowProjectClusterSettingRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateFlowProjectClusterSettingRequest) SetClusterId(v string) *CreateFlowProjectClusterSettingRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateFlowProjectClusterSettingRequest) SetDefaultUser(v string) *CreateFlowProjectClusterSettingRequest {
	s.DefaultUser = &v
	return s
}

func (s *CreateFlowProjectClusterSettingRequest) SetDefaultQueue(v string) *CreateFlowProjectClusterSettingRequest {
	s.DefaultQueue = &v
	return s
}

func (s *CreateFlowProjectClusterSettingRequest) SetClientToken(v string) *CreateFlowProjectClusterSettingRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateFlowProjectClusterSettingRequest) SetUserList(v []*string) *CreateFlowProjectClusterSettingRequest {
	s.UserList = v
	return s
}

func (s *CreateFlowProjectClusterSettingRequest) SetQueueList(v []*string) *CreateFlowProjectClusterSettingRequest {
	s.QueueList = v
	return s
}

func (s *CreateFlowProjectClusterSettingRequest) SetHostList(v []*string) *CreateFlowProjectClusterSettingRequest {
	s.HostList = v
	return s
}

type CreateFlowProjectClusterSettingResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFlowProjectClusterSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowProjectClusterSettingResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlowProjectClusterSettingResponseBody) SetData(v bool) *CreateFlowProjectClusterSettingResponseBody {
	s.Data = &v
	return s
}

func (s *CreateFlowProjectClusterSettingResponseBody) SetRequestId(v string) *CreateFlowProjectClusterSettingResponseBody {
	s.RequestId = &v
	return s
}

type CreateFlowProjectClusterSettingResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFlowProjectClusterSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFlowProjectClusterSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowProjectClusterSettingResponse) GoString() string {
	return s.String()
}

func (s *CreateFlowProjectClusterSettingResponse) SetHeaders(v map[string]*string) *CreateFlowProjectClusterSettingResponse {
	s.Headers = v
	return s
}

func (s *CreateFlowProjectClusterSettingResponse) SetBody(v *CreateFlowProjectClusterSettingResponseBody) *CreateFlowProjectClusterSettingResponse {
	s.Body = v
	return s
}

type DescribeFlowInstanceRequest struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeFlowInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowInstanceRequest) SetProjectId(v string) *DescribeFlowInstanceRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeFlowInstanceRequest) SetId(v string) *DescribeFlowInstanceRequest {
	s.Id = &v
	return s
}

func (s *DescribeFlowInstanceRequest) SetRegionId(v string) *DescribeFlowInstanceRequest {
	s.RegionId = &v
	return s
}

type DescribeFlowInstanceResponseBody struct {
	Status             *string                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	ProjectId          *string                                             `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Namespace          *string                                             `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	FlowName           *string                                             `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	GmtModified        *int64                                              `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	FlowId             *string                                             `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	CronExpression     *string                                             `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	EndTime            *int64                                              `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime          *int64                                              `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	HasNodeFailed      *bool                                               `json:"HasNodeFailed,omitempty" xml:"HasNodeFailed,omitempty"`
	RequestId          *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Graph              *string                                             `json:"Graph,omitempty" xml:"Graph,omitempty"`
	GmtCreate          *int64                                              `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	ScheduleTime       *int64                                              `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	Duration           *int64                                              `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Id                 *string                                             `json:"Id,omitempty" xml:"Id,omitempty"`
	ClusterId          *string                                             `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DependencyFlowList *DescribeFlowInstanceResponseBodyDependencyFlowList `json:"DependencyFlowList,omitempty" xml:"DependencyFlowList,omitempty" type:"Struct"`
	NodeInstance       *DescribeFlowInstanceResponseBodyNodeInstance       `json:"NodeInstance,omitempty" xml:"NodeInstance,omitempty" type:"Struct"`
}

func (s DescribeFlowInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowInstanceResponseBody) SetStatus(v string) *DescribeFlowInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeFlowInstanceResponseBody) SetProjectId(v string) *DescribeFlowInstanceResponseBody {
	s.ProjectId = &v
	return s
}

func (s *DescribeFlowInstanceResponseBody) SetNamespace(v string) *DescribeFlowInstanceResponseBody {
	s.Namespace = &v
	return s
}

func (s *DescribeFlowInstanceResponseBody) SetFlowName(v string) *DescribeFlowInstanceResponseBody {
	s.FlowName = &v
	return s
}

func (s *DescribeFlowInstanceResponseBody) SetGmtModified(v int64) *DescribeFlowInstanceResponseBody {
	s.GmtModified = &v
	return s
}

func (s *DescribeFlowInstanceResponseBody) SetFlowId(v string) *DescribeFlowInstanceResponseBody {
	s.FlowId = &v
	return s
}

func (s *DescribeFlowInstanceResponseBody) SetCronExpression(v string) *DescribeFlowInstanceResponseBody {
	s.CronExpression = &v
	return s
}

func (s *DescribeFlowInstanceResponseBody) SetEndTime(v int64) *DescribeFlowInstanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeFlowInstanceResponseBody) SetStartTime(v int64) *DescribeFlowInstanceResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeFlowInstanceResponseBody) SetHasNodeFailed(v bool) *DescribeFlowInstanceResponseBody {
	s.HasNodeFailed = &v
	return s
}

func (s *DescribeFlowInstanceResponseBody) SetRequestId(v string) *DescribeFlowInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowInstanceResponseBody) SetGraph(v string) *DescribeFlowInstanceResponseBody {
	s.Graph = &v
	return s
}

func (s *DescribeFlowInstanceResponseBody) SetGmtCreate(v int64) *DescribeFlowInstanceResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *DescribeFlowInstanceResponseBody) SetScheduleTime(v int64) *DescribeFlowInstanceResponseBody {
	s.ScheduleTime = &v
	return s
}

func (s *DescribeFlowInstanceResponseBody) SetDuration(v int64) *DescribeFlowInstanceResponseBody {
	s.Duration = &v
	return s
}

func (s *DescribeFlowInstanceResponseBody) SetId(v string) *DescribeFlowInstanceResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeFlowInstanceResponseBody) SetClusterId(v string) *DescribeFlowInstanceResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeFlowInstanceResponseBody) SetDependencyFlowList(v *DescribeFlowInstanceResponseBodyDependencyFlowList) *DescribeFlowInstanceResponseBody {
	s.DependencyFlowList = v
	return s
}

func (s *DescribeFlowInstanceResponseBody) SetNodeInstance(v *DescribeFlowInstanceResponseBodyNodeInstance) *DescribeFlowInstanceResponseBody {
	s.NodeInstance = v
	return s
}

type DescribeFlowInstanceResponseBodyDependencyFlowList struct {
	ParentFlow []*DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow `json:"ParentFlow,omitempty" xml:"ParentFlow,omitempty" type:"Repeated"`
}

func (s DescribeFlowInstanceResponseBodyDependencyFlowList) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowInstanceResponseBodyDependencyFlowList) GoString() string {
	return s.String()
}

func (s *DescribeFlowInstanceResponseBodyDependencyFlowList) SetParentFlow(v []*DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow) *DescribeFlowInstanceResponseBodyDependencyFlowList {
	s.ParentFlow = v
	return s
}

type DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow struct {
	ScheduleKey          *string `json:"ScheduleKey,omitempty" xml:"ScheduleKey,omitempty"`
	BizDate              *int64  `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	DependencyInstanceId *string `json:"DependencyInstanceId,omitempty" xml:"DependencyInstanceId,omitempty"`
	DependencyFlowId     *string `json:"DependencyFlowId,omitempty" xml:"DependencyFlowId,omitempty"`
	Meet                 *bool   `json:"Meet,omitempty" xml:"Meet,omitempty"`
	FlowInstanceId       *string `json:"FlowInstanceId,omitempty" xml:"FlowInstanceId,omitempty"`
	ProjectId            *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	FlowId               *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
}

func (s DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow) GoString() string {
	return s.String()
}

func (s *DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow) SetScheduleKey(v string) *DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow {
	s.ScheduleKey = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow) SetBizDate(v int64) *DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow {
	s.BizDate = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow) SetDependencyInstanceId(v string) *DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow {
	s.DependencyInstanceId = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow) SetDependencyFlowId(v string) *DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow {
	s.DependencyFlowId = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow) SetMeet(v bool) *DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow {
	s.Meet = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow) SetFlowInstanceId(v string) *DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow {
	s.FlowInstanceId = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow) SetProjectId(v string) *DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow {
	s.ProjectId = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow) SetFlowId(v string) *DescribeFlowInstanceResponseBodyDependencyFlowListParentFlow {
	s.FlowId = &v
	return s
}

type DescribeFlowInstanceResponseBodyNodeInstance struct {
	NodeInstance []*DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance `json:"NodeInstance,omitempty" xml:"NodeInstance,omitempty" type:"Repeated"`
}

func (s DescribeFlowInstanceResponseBodyNodeInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowInstanceResponseBodyNodeInstance) GoString() string {
	return s.String()
}

func (s *DescribeFlowInstanceResponseBodyNodeInstance) SetNodeInstance(v []*DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) *DescribeFlowInstanceResponseBodyNodeInstance {
	s.NodeInstance = v
	return s
}

type DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance struct {
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
	FailAct        *string `json:"FailAct,omitempty" xml:"FailAct,omitempty"`
	ProjectId      *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RetryInterval  *string `json:"RetryInterval,omitempty" xml:"RetryInterval,omitempty"`
	JobType        *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	HostName       *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	ExternalInfo   *string `json:"ExternalInfo,omitempty" xml:"ExternalInfo,omitempty"`
	GmtModified    *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Retries        *int32  `json:"Retries,omitempty" xml:"Retries,omitempty"`
	EndTime        *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime      *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	ExternalStatus *string `json:"ExternalStatus,omitempty" xml:"ExternalStatus,omitempty"`
	JobName        *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	NodeName       *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	JobId          *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	GmtCreate      *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	ExternalId     *string `json:"ExternalId,omitempty" xml:"ExternalId,omitempty"`
	Duration       *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Id             *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Pending        *bool   `json:"Pending,omitempty" xml:"Pending,omitempty"`
	ClusterId      *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	MaxRetry       *string `json:"MaxRetry,omitempty" xml:"MaxRetry,omitempty"`
}

func (s DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) GoString() string {
	return s.String()
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetStatus(v string) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.Status = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetType(v string) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.Type = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetFailAct(v string) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.FailAct = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetProjectId(v string) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.ProjectId = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetRetryInterval(v string) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.RetryInterval = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetJobType(v string) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.JobType = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetHostName(v string) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.HostName = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetExternalInfo(v string) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.ExternalInfo = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetGmtModified(v int64) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.GmtModified = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetRetries(v int32) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.Retries = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetEndTime(v int64) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.EndTime = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetStartTime(v int64) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.StartTime = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetExternalStatus(v string) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.ExternalStatus = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetJobName(v string) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.JobName = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetNodeName(v string) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.NodeName = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetJobId(v string) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.JobId = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetGmtCreate(v int64) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.GmtCreate = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetExternalId(v string) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.ExternalId = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetDuration(v int64) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.Duration = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetId(v string) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.Id = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetPending(v bool) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.Pending = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetClusterId(v string) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.ClusterId = &v
	return s
}

func (s *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance) SetMaxRetry(v string) *DescribeFlowInstanceResponseBodyNodeInstanceNodeInstance {
	s.MaxRetry = &v
	return s
}

type DescribeFlowInstanceResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeFlowInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFlowInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFlowInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeFlowInstanceResponse) SetHeaders(v map[string]*string) *DescribeFlowInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeFlowInstanceResponse) SetBody(v *DescribeFlowInstanceResponseBody) *DescribeFlowInstanceResponse {
	s.Body = v
	return s
}

type CreateFlowProjectUserRequest struct {
	RegionId    *string                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProjectId   *string                             `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ClientToken *string                             `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	User        []*CreateFlowProjectUserRequestUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s CreateFlowProjectUserRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowProjectUserRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowProjectUserRequest) SetRegionId(v string) *CreateFlowProjectUserRequest {
	s.RegionId = &v
	return s
}

func (s *CreateFlowProjectUserRequest) SetProjectId(v string) *CreateFlowProjectUserRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateFlowProjectUserRequest) SetClientToken(v string) *CreateFlowProjectUserRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateFlowProjectUserRequest) SetUser(v []*CreateFlowProjectUserRequestUser) *CreateFlowProjectUserRequest {
	s.User = v
	return s
}

type CreateFlowProjectUserRequestUser struct {
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateFlowProjectUserRequestUser) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowProjectUserRequestUser) GoString() string {
	return s.String()
}

func (s *CreateFlowProjectUserRequestUser) SetUserId(v string) *CreateFlowProjectUserRequestUser {
	s.UserId = &v
	return s
}

func (s *CreateFlowProjectUserRequestUser) SetUserName(v string) *CreateFlowProjectUserRequestUser {
	s.UserName = &v
	return s
}

type CreateFlowProjectUserResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFlowProjectUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowProjectUserResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlowProjectUserResponseBody) SetData(v bool) *CreateFlowProjectUserResponseBody {
	s.Data = &v
	return s
}

func (s *CreateFlowProjectUserResponseBody) SetRequestId(v string) *CreateFlowProjectUserResponseBody {
	s.RequestId = &v
	return s
}

type CreateFlowProjectUserResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFlowProjectUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFlowProjectUserResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowProjectUserResponse) GoString() string {
	return s.String()
}

func (s *CreateFlowProjectUserResponse) SetHeaders(v map[string]*string) *CreateFlowProjectUserResponse {
	s.Headers = v
	return s
}

func (s *CreateFlowProjectUserResponse) SetBody(v *CreateFlowProjectUserResponseBody) *CreateFlowProjectUserResponse {
	s.Body = v
	return s
}

type CreateFlowCategoryRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	ParentId    *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateFlowCategoryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowCategoryRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowCategoryRequest) SetRegionId(v string) *CreateFlowCategoryRequest {
	s.RegionId = &v
	return s
}

func (s *CreateFlowCategoryRequest) SetProjectId(v string) *CreateFlowCategoryRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateFlowCategoryRequest) SetName(v string) *CreateFlowCategoryRequest {
	s.Name = &v
	return s
}

func (s *CreateFlowCategoryRequest) SetType(v string) *CreateFlowCategoryRequest {
	s.Type = &v
	return s
}

func (s *CreateFlowCategoryRequest) SetParentId(v string) *CreateFlowCategoryRequest {
	s.ParentId = &v
	return s
}

func (s *CreateFlowCategoryRequest) SetClientToken(v string) *CreateFlowCategoryRequest {
	s.ClientToken = &v
	return s
}

type CreateFlowCategoryResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFlowCategoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlowCategoryResponseBody) SetId(v string) *CreateFlowCategoryResponseBody {
	s.Id = &v
	return s
}

func (s *CreateFlowCategoryResponseBody) SetRequestId(v string) *CreateFlowCategoryResponseBody {
	s.RequestId = &v
	return s
}

type CreateFlowCategoryResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateFlowCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFlowCategoryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFlowCategoryResponse) GoString() string {
	return s.String()
}

func (s *CreateFlowCategoryResponse) SetHeaders(v map[string]*string) *CreateFlowCategoryResponse {
	s.Headers = v
	return s
}

func (s *CreateFlowCategoryResponse) SetBody(v *CreateFlowCategoryResponseBody) *CreateFlowCategoryResponse {
	s.Body = v
	return s
}

type DeleteFlowProjectClusterSettingRequest struct {
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DeleteFlowProjectClusterSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowProjectClusterSettingRequest) GoString() string {
	return s.String()
}

func (s *DeleteFlowProjectClusterSettingRequest) SetRegionId(v string) *DeleteFlowProjectClusterSettingRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteFlowProjectClusterSettingRequest) SetProjectId(v string) *DeleteFlowProjectClusterSettingRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteFlowProjectClusterSettingRequest) SetClusterId(v string) *DeleteFlowProjectClusterSettingRequest {
	s.ClusterId = &v
	return s
}

type DeleteFlowProjectClusterSettingResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFlowProjectClusterSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowProjectClusterSettingResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFlowProjectClusterSettingResponseBody) SetData(v bool) *DeleteFlowProjectClusterSettingResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteFlowProjectClusterSettingResponseBody) SetRequestId(v string) *DeleteFlowProjectClusterSettingResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFlowProjectClusterSettingResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFlowProjectClusterSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFlowProjectClusterSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowProjectClusterSettingResponse) GoString() string {
	return s.String()
}

func (s *DeleteFlowProjectClusterSettingResponse) SetHeaders(v map[string]*string) *DeleteFlowProjectClusterSettingResponse {
	s.Headers = v
	return s
}

func (s *DeleteFlowProjectClusterSettingResponse) SetBody(v *DeleteFlowProjectClusterSettingResponseBody) *DeleteFlowProjectClusterSettingResponse {
	s.Body = v
	return s
}

type ListLibrariesRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Limit           *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentSize     *int32  `json:"CurrentSize,omitempty" xml:"CurrentSize,omitempty"`
	PageCount       *int32  `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	OrderField      *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	OrderMode       *string `json:"OrderMode,omitempty" xml:"OrderMode,omitempty"`
	ClusterBizId    *string `json:"ClusterBizId,omitempty" xml:"ClusterBizId,omitempty"`
}

func (s ListLibrariesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLibrariesRequest) GoString() string {
	return s.String()
}

func (s *ListLibrariesRequest) SetResourceOwnerId(v int64) *ListLibrariesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListLibrariesRequest) SetRegionId(v string) *ListLibrariesRequest {
	s.RegionId = &v
	return s
}

func (s *ListLibrariesRequest) SetLimit(v int32) *ListLibrariesRequest {
	s.Limit = &v
	return s
}

func (s *ListLibrariesRequest) SetPageNumber(v int32) *ListLibrariesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLibrariesRequest) SetPageSize(v int32) *ListLibrariesRequest {
	s.PageSize = &v
	return s
}

func (s *ListLibrariesRequest) SetCurrentSize(v int32) *ListLibrariesRequest {
	s.CurrentSize = &v
	return s
}

func (s *ListLibrariesRequest) SetPageCount(v int32) *ListLibrariesRequest {
	s.PageCount = &v
	return s
}

func (s *ListLibrariesRequest) SetOrderField(v string) *ListLibrariesRequest {
	s.OrderField = &v
	return s
}

func (s *ListLibrariesRequest) SetOrderMode(v string) *ListLibrariesRequest {
	s.OrderMode = &v
	return s
}

func (s *ListLibrariesRequest) SetClusterBizId(v string) *ListLibrariesRequest {
	s.ClusterBizId = &v
	return s
}

type ListLibrariesResponseBody struct {
	NextToken  *string                         `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PageSize   *int32                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Items      *ListLibrariesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s ListLibrariesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLibrariesResponseBody) GoString() string {
	return s.String()
}

func (s *ListLibrariesResponseBody) SetNextToken(v string) *ListLibrariesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListLibrariesResponseBody) SetPageSize(v int32) *ListLibrariesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListLibrariesResponseBody) SetPageNumber(v int32) *ListLibrariesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListLibrariesResponseBody) SetRequestId(v string) *ListLibrariesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLibrariesResponseBody) SetTotalCount(v int32) *ListLibrariesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListLibrariesResponseBody) SetItems(v *ListLibrariesResponseBodyItems) *ListLibrariesResponseBody {
	s.Items = v
	return s
}

type ListLibrariesResponseBodyItems struct {
	Item []*ListLibrariesResponseBodyItemsItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s ListLibrariesResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s ListLibrariesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListLibrariesResponseBodyItems) SetItem(v []*ListLibrariesResponseBodyItemsItem) *ListLibrariesResponseBodyItems {
	s.Item = v
	return s
}

type ListLibrariesResponseBodyItemsItem struct {
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
	CreateTime     *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	SourceType     *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	BizId          *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Scope          *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	LibraryVersion *string `json:"LibraryVersion,omitempty" xml:"LibraryVersion,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Properties     *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	SourceLocation *string `json:"SourceLocation,omitempty" xml:"SourceLocation,omitempty"`
}

func (s ListLibrariesResponseBodyItemsItem) String() string {
	return tea.Prettify(s)
}

func (s ListLibrariesResponseBodyItemsItem) GoString() string {
	return s.String()
}

func (s *ListLibrariesResponseBodyItemsItem) SetType(v string) *ListLibrariesResponseBodyItemsItem {
	s.Type = &v
	return s
}

func (s *ListLibrariesResponseBodyItemsItem) SetCreateTime(v int64) *ListLibrariesResponseBodyItemsItem {
	s.CreateTime = &v
	return s
}

func (s *ListLibrariesResponseBodyItemsItem) SetUserId(v string) *ListLibrariesResponseBodyItemsItem {
	s.UserId = &v
	return s
}

func (s *ListLibrariesResponseBodyItemsItem) SetSourceType(v string) *ListLibrariesResponseBodyItemsItem {
	s.SourceType = &v
	return s
}

func (s *ListLibrariesResponseBodyItemsItem) SetBizId(v string) *ListLibrariesResponseBodyItemsItem {
	s.BizId = &v
	return s
}

func (s *ListLibrariesResponseBodyItemsItem) SetScope(v string) *ListLibrariesResponseBodyItemsItem {
	s.Scope = &v
	return s
}

func (s *ListLibrariesResponseBodyItemsItem) SetLibraryVersion(v string) *ListLibrariesResponseBodyItemsItem {
	s.LibraryVersion = &v
	return s
}

func (s *ListLibrariesResponseBodyItemsItem) SetName(v string) *ListLibrariesResponseBodyItemsItem {
	s.Name = &v
	return s
}

func (s *ListLibrariesResponseBodyItemsItem) SetProperties(v string) *ListLibrariesResponseBodyItemsItem {
	s.Properties = &v
	return s
}

func (s *ListLibrariesResponseBodyItemsItem) SetSourceLocation(v string) *ListLibrariesResponseBodyItemsItem {
	s.SourceLocation = &v
	return s
}

type ListLibrariesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListLibrariesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListLibrariesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLibrariesResponse) GoString() string {
	return s.String()
}

func (s *ListLibrariesResponse) SetHeaders(v map[string]*string) *ListLibrariesResponse {
	s.Headers = v
	return s
}

func (s *ListLibrariesResponse) SetBody(v *ListLibrariesResponseBody) *ListLibrariesResponse {
	s.Body = v
	return s
}

type RunScalingActionRequest struct {
	ResourceOwnerId   *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ScalingActionType *string `json:"ScalingActionType,omitempty" xml:"ScalingActionType,omitempty"`
	ScalingGroupBizId *string `json:"ScalingGroupBizId,omitempty" xml:"ScalingGroupBizId,omitempty"`
	ActionParam       *string `json:"ActionParam,omitempty" xml:"ActionParam,omitempty"`
}

func (s RunScalingActionRequest) String() string {
	return tea.Prettify(s)
}

func (s RunScalingActionRequest) GoString() string {
	return s.String()
}

func (s *RunScalingActionRequest) SetResourceOwnerId(v int64) *RunScalingActionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RunScalingActionRequest) SetRegionId(v string) *RunScalingActionRequest {
	s.RegionId = &v
	return s
}

func (s *RunScalingActionRequest) SetScalingActionType(v string) *RunScalingActionRequest {
	s.ScalingActionType = &v
	return s
}

func (s *RunScalingActionRequest) SetScalingGroupBizId(v string) *RunScalingActionRequest {
	s.ScalingGroupBizId = &v
	return s
}

func (s *RunScalingActionRequest) SetActionParam(v string) *RunScalingActionRequest {
	s.ActionParam = &v
	return s
}

type RunScalingActionResponseBody struct {
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// data
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s RunScalingActionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunScalingActionResponseBody) GoString() string {
	return s.String()
}

func (s *RunScalingActionResponseBody) SetRequestId(v string) *RunScalingActionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunScalingActionResponseBody) SetData(v string) *RunScalingActionResponseBody {
	s.Data = &v
	return s
}

type RunScalingActionResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RunScalingActionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RunScalingActionResponse) String() string {
	return tea.Prettify(s)
}

func (s RunScalingActionResponse) GoString() string {
	return s.String()
}

func (s *RunScalingActionResponse) SetHeaders(v map[string]*string) *RunScalingActionResponse {
	s.Headers = v
	return s
}

func (s *RunScalingActionResponse) SetBody(v *RunScalingActionResponseBody) *RunScalingActionResponse {
	s.Body = v
	return s
}

type InstallLibrariesRequest struct {
	ResourceOwnerId  *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId         *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClusterBizIdList []*string `json:"ClusterBizIdList,omitempty" xml:"ClusterBizIdList,omitempty" type:"Repeated"`
	LibraryBizId     *string   `json:"LibraryBizId,omitempty" xml:"LibraryBizId,omitempty"`
}

func (s InstallLibrariesRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallLibrariesRequest) GoString() string {
	return s.String()
}

func (s *InstallLibrariesRequest) SetResourceOwnerId(v int64) *InstallLibrariesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *InstallLibrariesRequest) SetRegionId(v string) *InstallLibrariesRequest {
	s.RegionId = &v
	return s
}

func (s *InstallLibrariesRequest) SetClusterBizIdList(v []*string) *InstallLibrariesRequest {
	s.ClusterBizIdList = v
	return s
}

func (s *InstallLibrariesRequest) SetLibraryBizId(v string) *InstallLibrariesRequest {
	s.LibraryBizId = &v
	return s
}

type InstallLibrariesResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InstallLibrariesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InstallLibrariesResponseBody) GoString() string {
	return s.String()
}

func (s *InstallLibrariesResponseBody) SetData(v string) *InstallLibrariesResponseBody {
	s.Data = &v
	return s
}

func (s *InstallLibrariesResponseBody) SetRequestId(v string) *InstallLibrariesResponseBody {
	s.RequestId = &v
	return s
}

type InstallLibrariesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InstallLibrariesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InstallLibrariesResponse) String() string {
	return tea.Prettify(s)
}

func (s InstallLibrariesResponse) GoString() string {
	return s.String()
}

func (s *InstallLibrariesResponse) SetHeaders(v map[string]*string) *InstallLibrariesResponse {
	s.Headers = v
	return s
}

func (s *InstallLibrariesResponse) SetBody(v *InstallLibrariesResponseBody) *InstallLibrariesResponse {
	s.Body = v
	return s
}

type ListFlowJobsRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProjectId  *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Id         *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Adhoc      *bool   `json:"Adhoc,omitempty" xml:"Adhoc,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListFlowJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowJobsRequest) GoString() string {
	return s.String()
}

func (s *ListFlowJobsRequest) SetRegionId(v string) *ListFlowJobsRequest {
	s.RegionId = &v
	return s
}

func (s *ListFlowJobsRequest) SetProjectId(v string) *ListFlowJobsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFlowJobsRequest) SetId(v string) *ListFlowJobsRequest {
	s.Id = &v
	return s
}

func (s *ListFlowJobsRequest) SetName(v string) *ListFlowJobsRequest {
	s.Name = &v
	return s
}

func (s *ListFlowJobsRequest) SetType(v string) *ListFlowJobsRequest {
	s.Type = &v
	return s
}

func (s *ListFlowJobsRequest) SetAdhoc(v bool) *ListFlowJobsRequest {
	s.Adhoc = &v
	return s
}

func (s *ListFlowJobsRequest) SetPageNumber(v int32) *ListFlowJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFlowJobsRequest) SetPageSize(v int32) *ListFlowJobsRequest {
	s.PageSize = &v
	return s
}

type ListFlowJobsResponseBody struct {
	PageNumber *int32                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total      *int32                           `json:"Total,omitempty" xml:"Total,omitempty"`
	JobList    *ListFlowJobsResponseBodyJobList `json:"JobList,omitempty" xml:"JobList,omitempty" type:"Struct"`
}

func (s ListFlowJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowJobsResponseBody) SetPageNumber(v int32) *ListFlowJobsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListFlowJobsResponseBody) SetPageSize(v int32) *ListFlowJobsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListFlowJobsResponseBody) SetRequestId(v string) *ListFlowJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowJobsResponseBody) SetTotal(v int32) *ListFlowJobsResponseBody {
	s.Total = &v
	return s
}

func (s *ListFlowJobsResponseBody) SetJobList(v *ListFlowJobsResponseBodyJobList) *ListFlowJobsResponseBody {
	s.JobList = v
	return s
}

type ListFlowJobsResponseBodyJobList struct {
	Job []*ListFlowJobsResponseBodyJobListJob `json:"Job,omitempty" xml:"Job,omitempty" type:"Repeated"`
}

func (s ListFlowJobsResponseBodyJobList) String() string {
	return tea.Prettify(s)
}

func (s ListFlowJobsResponseBodyJobList) GoString() string {
	return s.String()
}

func (s *ListFlowJobsResponseBodyJobList) SetJob(v []*ListFlowJobsResponseBodyJobListJob) *ListFlowJobsResponseBodyJobList {
	s.Job = v
	return s
}

type ListFlowJobsResponseBodyJobListJob struct {
	Type               *string                                         `json:"Type,omitempty" xml:"Type,omitempty"`
	FailAct            *string                                         `json:"FailAct,omitempty" xml:"FailAct,omitempty"`
	CustomVariables    *string                                         `json:"CustomVariables,omitempty" xml:"CustomVariables,omitempty"`
	RetryInterval      *int64                                          `json:"RetryInterval,omitempty" xml:"RetryInterval,omitempty"`
	EnvConf            *string                                         `json:"EnvConf,omitempty" xml:"EnvConf,omitempty"`
	ParamConf          *string                                         `json:"ParamConf,omitempty" xml:"ParamConf,omitempty"`
	Mode               *string                                         `json:"Mode,omitempty" xml:"Mode,omitempty"`
	GmtModified        *int64                                          `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	MonitorConf        *string                                         `json:"MonitorConf,omitempty" xml:"MonitorConf,omitempty"`
	LastInstanceDetail *string                                         `json:"LastInstanceDetail,omitempty" xml:"LastInstanceDetail,omitempty"`
	RunConf            *string                                         `json:"RunConf,omitempty" xml:"RunConf,omitempty"`
	Params             *string                                         `json:"Params,omitempty" xml:"Params,omitempty"`
	Description        *string                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate          *int64                                          `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	CategoryId         *string                                         `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	Adhoc              *string                                         `json:"Adhoc,omitempty" xml:"Adhoc,omitempty"`
	Name               *string                                         `json:"Name,omitempty" xml:"Name,omitempty"`
	Id                 *string                                         `json:"Id,omitempty" xml:"Id,omitempty"`
	MaxRetry           *int32                                          `json:"MaxRetry,omitempty" xml:"MaxRetry,omitempty"`
	AlertConf          *string                                         `json:"AlertConf,omitempty" xml:"AlertConf,omitempty"`
	ResourceList       *ListFlowJobsResponseBodyJobListJobResourceList `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Struct"`
}

func (s ListFlowJobsResponseBodyJobListJob) String() string {
	return tea.Prettify(s)
}

func (s ListFlowJobsResponseBodyJobListJob) GoString() string {
	return s.String()
}

func (s *ListFlowJobsResponseBodyJobListJob) SetType(v string) *ListFlowJobsResponseBodyJobListJob {
	s.Type = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJob) SetFailAct(v string) *ListFlowJobsResponseBodyJobListJob {
	s.FailAct = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJob) SetCustomVariables(v string) *ListFlowJobsResponseBodyJobListJob {
	s.CustomVariables = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJob) SetRetryInterval(v int64) *ListFlowJobsResponseBodyJobListJob {
	s.RetryInterval = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJob) SetEnvConf(v string) *ListFlowJobsResponseBodyJobListJob {
	s.EnvConf = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJob) SetParamConf(v string) *ListFlowJobsResponseBodyJobListJob {
	s.ParamConf = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJob) SetMode(v string) *ListFlowJobsResponseBodyJobListJob {
	s.Mode = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJob) SetGmtModified(v int64) *ListFlowJobsResponseBodyJobListJob {
	s.GmtModified = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJob) SetMonitorConf(v string) *ListFlowJobsResponseBodyJobListJob {
	s.MonitorConf = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJob) SetLastInstanceDetail(v string) *ListFlowJobsResponseBodyJobListJob {
	s.LastInstanceDetail = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJob) SetRunConf(v string) *ListFlowJobsResponseBodyJobListJob {
	s.RunConf = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJob) SetParams(v string) *ListFlowJobsResponseBodyJobListJob {
	s.Params = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJob) SetDescription(v string) *ListFlowJobsResponseBodyJobListJob {
	s.Description = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJob) SetGmtCreate(v int64) *ListFlowJobsResponseBodyJobListJob {
	s.GmtCreate = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJob) SetCategoryId(v string) *ListFlowJobsResponseBodyJobListJob {
	s.CategoryId = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJob) SetAdhoc(v string) *ListFlowJobsResponseBodyJobListJob {
	s.Adhoc = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJob) SetName(v string) *ListFlowJobsResponseBodyJobListJob {
	s.Name = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJob) SetId(v string) *ListFlowJobsResponseBodyJobListJob {
	s.Id = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJob) SetMaxRetry(v int32) *ListFlowJobsResponseBodyJobListJob {
	s.MaxRetry = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJob) SetAlertConf(v string) *ListFlowJobsResponseBodyJobListJob {
	s.AlertConf = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJob) SetResourceList(v *ListFlowJobsResponseBodyJobListJobResourceList) *ListFlowJobsResponseBodyJobListJob {
	s.ResourceList = v
	return s
}

type ListFlowJobsResponseBodyJobListJobResourceList struct {
	Resource []*ListFlowJobsResponseBodyJobListJobResourceListResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
}

func (s ListFlowJobsResponseBodyJobListJobResourceList) String() string {
	return tea.Prettify(s)
}

func (s ListFlowJobsResponseBodyJobListJobResourceList) GoString() string {
	return s.String()
}

func (s *ListFlowJobsResponseBodyJobListJobResourceList) SetResource(v []*ListFlowJobsResponseBodyJobListJobResourceListResource) *ListFlowJobsResponseBodyJobListJobResourceList {
	s.Resource = v
	return s
}

type ListFlowJobsResponseBodyJobListJobResourceListResource struct {
	Path  *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
}

func (s ListFlowJobsResponseBodyJobListJobResourceListResource) String() string {
	return tea.Prettify(s)
}

func (s ListFlowJobsResponseBodyJobListJobResourceListResource) GoString() string {
	return s.String()
}

func (s *ListFlowJobsResponseBodyJobListJobResourceListResource) SetPath(v string) *ListFlowJobsResponseBodyJobListJobResourceListResource {
	s.Path = &v
	return s
}

func (s *ListFlowJobsResponseBodyJobListJobResourceListResource) SetAlias(v string) *ListFlowJobsResponseBodyJobListJobResourceListResource {
	s.Alias = &v
	return s
}

type ListFlowJobsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFlowJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowJobsResponse) GoString() string {
	return s.String()
}

func (s *ListFlowJobsResponse) SetHeaders(v map[string]*string) *ListFlowJobsResponse {
	s.Headers = v
	return s
}

func (s *ListFlowJobsResponse) SetBody(v *ListFlowJobsResponseBody) *ListFlowJobsResponse {
	s.Body = v
	return s
}

type ModifyFlowRequest struct {
	RegionId                *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProjectId               *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Id                      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name                    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Periodic                *bool   `json:"Periodic,omitempty" xml:"Periodic,omitempty"`
	StartSchedule           *int64  `json:"StartSchedule,omitempty" xml:"StartSchedule,omitempty"`
	EndSchedule             *int64  `json:"EndSchedule,omitempty" xml:"EndSchedule,omitempty"`
	CronExpr                *string `json:"CronExpr,omitempty" xml:"CronExpr,omitempty"`
	CreateCluster           *bool   `json:"CreateCluster,omitempty" xml:"CreateCluster,omitempty"`
	ClusterId               *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	HostName                *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	Application             *string `json:"Application,omitempty" xml:"Application,omitempty"`
	AlertConf               *string `json:"AlertConf,omitempty" xml:"AlertConf,omitempty"`
	AlertUserGroupBizId     *string `json:"AlertUserGroupBizId,omitempty" xml:"AlertUserGroupBizId,omitempty"`
	AlertDingDingGroupBizId *string `json:"AlertDingDingGroupBizId,omitempty" xml:"AlertDingDingGroupBizId,omitempty"`
	ParentFlowList          *string `json:"ParentFlowList,omitempty" xml:"ParentFlowList,omitempty"`
	ParentCategory          *string `json:"ParentCategory,omitempty" xml:"ParentCategory,omitempty"`
}

func (s ModifyFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowRequest) GoString() string {
	return s.String()
}

func (s *ModifyFlowRequest) SetRegionId(v string) *ModifyFlowRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyFlowRequest) SetProjectId(v string) *ModifyFlowRequest {
	s.ProjectId = &v
	return s
}

func (s *ModifyFlowRequest) SetId(v string) *ModifyFlowRequest {
	s.Id = &v
	return s
}

func (s *ModifyFlowRequest) SetName(v string) *ModifyFlowRequest {
	s.Name = &v
	return s
}

func (s *ModifyFlowRequest) SetStatus(v string) *ModifyFlowRequest {
	s.Status = &v
	return s
}

func (s *ModifyFlowRequest) SetDescription(v string) *ModifyFlowRequest {
	s.Description = &v
	return s
}

func (s *ModifyFlowRequest) SetPeriodic(v bool) *ModifyFlowRequest {
	s.Periodic = &v
	return s
}

func (s *ModifyFlowRequest) SetStartSchedule(v int64) *ModifyFlowRequest {
	s.StartSchedule = &v
	return s
}

func (s *ModifyFlowRequest) SetEndSchedule(v int64) *ModifyFlowRequest {
	s.EndSchedule = &v
	return s
}

func (s *ModifyFlowRequest) SetCronExpr(v string) *ModifyFlowRequest {
	s.CronExpr = &v
	return s
}

func (s *ModifyFlowRequest) SetCreateCluster(v bool) *ModifyFlowRequest {
	s.CreateCluster = &v
	return s
}

func (s *ModifyFlowRequest) SetClusterId(v string) *ModifyFlowRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyFlowRequest) SetHostName(v string) *ModifyFlowRequest {
	s.HostName = &v
	return s
}

func (s *ModifyFlowRequest) SetApplication(v string) *ModifyFlowRequest {
	s.Application = &v
	return s
}

func (s *ModifyFlowRequest) SetAlertConf(v string) *ModifyFlowRequest {
	s.AlertConf = &v
	return s
}

func (s *ModifyFlowRequest) SetAlertUserGroupBizId(v string) *ModifyFlowRequest {
	s.AlertUserGroupBizId = &v
	return s
}

func (s *ModifyFlowRequest) SetAlertDingDingGroupBizId(v string) *ModifyFlowRequest {
	s.AlertDingDingGroupBizId = &v
	return s
}

func (s *ModifyFlowRequest) SetParentFlowList(v string) *ModifyFlowRequest {
	s.ParentFlowList = &v
	return s
}

func (s *ModifyFlowRequest) SetParentCategory(v string) *ModifyFlowRequest {
	s.ParentCategory = &v
	return s
}

type ModifyFlowResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFlowResponseBody) SetData(v bool) *ModifyFlowResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyFlowResponseBody) SetRequestId(v string) *ModifyFlowResponseBody {
	s.RequestId = &v
	return s
}

type ModifyFlowResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowResponse) GoString() string {
	return s.String()
}

func (s *ModifyFlowResponse) SetHeaders(v map[string]*string) *ModifyFlowResponse {
	s.Headers = v
	return s
}

func (s *ModifyFlowResponse) SetBody(v *ModifyFlowResponseBody) *ModifyFlowResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetRegionId(v string) *DescribeRegionsRequest {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponseBody struct {
	RequestId         *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	BusinessLocations []*DescribeRegionsResponseBodyBusinessLocations `json:"BusinessLocations,omitempty" xml:"BusinessLocations,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetBusinessLocations(v []*DescribeRegionsResponseBodyBusinessLocations) *DescribeRegionsResponseBody {
	s.BusinessLocations = v
	return s
}

type DescribeRegionsResponseBodyBusinessLocations struct {
	Ordering      *string                                               `json:"Ordering,omitempty" xml:"Ordering,omitempty"`
	Type          *string                                               `json:"Type,omitempty" xml:"Type,omitempty"`
	ShowName      *string                                               `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
	Description   *string                                               `json:"Description,omitempty" xml:"Description,omitempty"`
	EnDescription *string                                               `json:"EnDescription,omitempty" xml:"EnDescription,omitempty"`
	EnName        *string                                               `json:"EnName,omitempty" xml:"EnName,omitempty"`
	CnName        *string                                               `json:"CnName,omitempty" xml:"CnName,omitempty"`
	Name          *string                                               `json:"Name,omitempty" xml:"Name,omitempty"`
	District      *DescribeRegionsResponseBodyBusinessLocationsDistrict `json:"District,omitempty" xml:"District,omitempty" type:"Struct"`
}

func (s DescribeRegionsResponseBodyBusinessLocations) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyBusinessLocations) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyBusinessLocations) SetOrdering(v string) *DescribeRegionsResponseBodyBusinessLocations {
	s.Ordering = &v
	return s
}

func (s *DescribeRegionsResponseBodyBusinessLocations) SetType(v string) *DescribeRegionsResponseBodyBusinessLocations {
	s.Type = &v
	return s
}

func (s *DescribeRegionsResponseBodyBusinessLocations) SetShowName(v string) *DescribeRegionsResponseBodyBusinessLocations {
	s.ShowName = &v
	return s
}

func (s *DescribeRegionsResponseBodyBusinessLocations) SetDescription(v string) *DescribeRegionsResponseBodyBusinessLocations {
	s.Description = &v
	return s
}

func (s *DescribeRegionsResponseBodyBusinessLocations) SetEnDescription(v string) *DescribeRegionsResponseBodyBusinessLocations {
	s.EnDescription = &v
	return s
}

func (s *DescribeRegionsResponseBodyBusinessLocations) SetEnName(v string) *DescribeRegionsResponseBodyBusinessLocations {
	s.EnName = &v
	return s
}

func (s *DescribeRegionsResponseBodyBusinessLocations) SetCnName(v string) *DescribeRegionsResponseBodyBusinessLocations {
	s.CnName = &v
	return s
}

func (s *DescribeRegionsResponseBodyBusinessLocations) SetName(v string) *DescribeRegionsResponseBodyBusinessLocations {
	s.Name = &v
	return s
}

func (s *DescribeRegionsResponseBodyBusinessLocations) SetDistrict(v *DescribeRegionsResponseBodyBusinessLocationsDistrict) *DescribeRegionsResponseBodyBusinessLocations {
	s.District = v
	return s
}

type DescribeRegionsResponseBodyBusinessLocationsDistrict struct {
	Ordering   *string `json:"Ordering,omitempty" xml:"Ordering,omitempty"`
	CnName     *string `json:"CnName,omitempty" xml:"CnName,omitempty"`
	ShowName   *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
	DistrictId *string `json:"DistrictId,omitempty" xml:"DistrictId,omitempty"`
	EnName     *string `json:"EnName,omitempty" xml:"EnName,omitempty"`
}

func (s DescribeRegionsResponseBodyBusinessLocationsDistrict) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyBusinessLocationsDistrict) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyBusinessLocationsDistrict) SetOrdering(v string) *DescribeRegionsResponseBodyBusinessLocationsDistrict {
	s.Ordering = &v
	return s
}

func (s *DescribeRegionsResponseBodyBusinessLocationsDistrict) SetCnName(v string) *DescribeRegionsResponseBodyBusinessLocationsDistrict {
	s.CnName = &v
	return s
}

func (s *DescribeRegionsResponseBodyBusinessLocationsDistrict) SetShowName(v string) *DescribeRegionsResponseBodyBusinessLocationsDistrict {
	s.ShowName = &v
	return s
}

func (s *DescribeRegionsResponseBodyBusinessLocationsDistrict) SetDistrictId(v string) *DescribeRegionsResponseBodyBusinessLocationsDistrict {
	s.DistrictId = &v
	return s
}

func (s *DescribeRegionsResponseBodyBusinessLocationsDistrict) SetEnName(v string) *DescribeRegionsResponseBodyBusinessLocationsDistrict {
	s.EnName = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type ListLibraryStatusRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Limit           *int32  `json:"Limit,omitempty" xml:"Limit,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentSize     *int32  `json:"CurrentSize,omitempty" xml:"CurrentSize,omitempty"`
	PageCount       *int32  `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	OrderField      *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	OrderMode       *string `json:"OrderMode,omitempty" xml:"OrderMode,omitempty"`
	LibraryBizId    *string `json:"LibraryBizId,omitempty" xml:"LibraryBizId,omitempty"`
	ClusterBizId    *string `json:"ClusterBizId,omitempty" xml:"ClusterBizId,omitempty"`
}

func (s ListLibraryStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLibraryStatusRequest) GoString() string {
	return s.String()
}

func (s *ListLibraryStatusRequest) SetResourceOwnerId(v int64) *ListLibraryStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListLibraryStatusRequest) SetRegionId(v string) *ListLibraryStatusRequest {
	s.RegionId = &v
	return s
}

func (s *ListLibraryStatusRequest) SetLimit(v int32) *ListLibraryStatusRequest {
	s.Limit = &v
	return s
}

func (s *ListLibraryStatusRequest) SetPageNumber(v int32) *ListLibraryStatusRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLibraryStatusRequest) SetPageSize(v int32) *ListLibraryStatusRequest {
	s.PageSize = &v
	return s
}

func (s *ListLibraryStatusRequest) SetCurrentSize(v int32) *ListLibraryStatusRequest {
	s.CurrentSize = &v
	return s
}

func (s *ListLibraryStatusRequest) SetPageCount(v int32) *ListLibraryStatusRequest {
	s.PageCount = &v
	return s
}

func (s *ListLibraryStatusRequest) SetOrderField(v string) *ListLibraryStatusRequest {
	s.OrderField = &v
	return s
}

func (s *ListLibraryStatusRequest) SetOrderMode(v string) *ListLibraryStatusRequest {
	s.OrderMode = &v
	return s
}

func (s *ListLibraryStatusRequest) SetLibraryBizId(v string) *ListLibraryStatusRequest {
	s.LibraryBizId = &v
	return s
}

func (s *ListLibraryStatusRequest) SetClusterBizId(v string) *ListLibraryStatusRequest {
	s.ClusterBizId = &v
	return s
}

type ListLibraryStatusResponseBody struct {
	NextToken  *string                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PageSize   *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Items      *ListLibraryStatusResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
}

func (s ListLibraryStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLibraryStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ListLibraryStatusResponseBody) SetNextToken(v string) *ListLibraryStatusResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListLibraryStatusResponseBody) SetPageSize(v int32) *ListLibraryStatusResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListLibraryStatusResponseBody) SetPageNumber(v int32) *ListLibraryStatusResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListLibraryStatusResponseBody) SetRequestId(v string) *ListLibraryStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLibraryStatusResponseBody) SetTotalCount(v int32) *ListLibraryStatusResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListLibraryStatusResponseBody) SetItems(v *ListLibraryStatusResponseBodyItems) *ListLibraryStatusResponseBody {
	s.Items = v
	return s
}

type ListLibraryStatusResponseBodyItems struct {
	Item []*ListLibraryStatusResponseBodyItemsItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s ListLibraryStatusResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s ListLibraryStatusResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListLibraryStatusResponseBodyItems) SetItem(v []*ListLibraryStatusResponseBodyItemsItem) *ListLibraryStatusResponseBodyItems {
	s.Item = v
	return s
}

type ListLibraryStatusResponseBodyItemsItem struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ClusterName  *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	ClusterBizId *string `json:"ClusterBizId,omitempty" xml:"ClusterBizId,omitempty"`
	LibraryBizId *string `json:"LibraryBizId,omitempty" xml:"LibraryBizId,omitempty"`
	LibraryName  *string `json:"LibraryName,omitempty" xml:"LibraryName,omitempty"`
}

func (s ListLibraryStatusResponseBodyItemsItem) String() string {
	return tea.Prettify(s)
}

func (s ListLibraryStatusResponseBodyItemsItem) GoString() string {
	return s.String()
}

func (s *ListLibraryStatusResponseBodyItemsItem) SetStatus(v string) *ListLibraryStatusResponseBodyItemsItem {
	s.Status = &v
	return s
}

func (s *ListLibraryStatusResponseBodyItemsItem) SetClusterName(v string) *ListLibraryStatusResponseBodyItemsItem {
	s.ClusterName = &v
	return s
}

func (s *ListLibraryStatusResponseBodyItemsItem) SetClusterBizId(v string) *ListLibraryStatusResponseBodyItemsItem {
	s.ClusterBizId = &v
	return s
}

func (s *ListLibraryStatusResponseBodyItemsItem) SetLibraryBizId(v string) *ListLibraryStatusResponseBodyItemsItem {
	s.LibraryBizId = &v
	return s
}

func (s *ListLibraryStatusResponseBodyItemsItem) SetLibraryName(v string) *ListLibraryStatusResponseBodyItemsItem {
	s.LibraryName = &v
	return s
}

type ListLibraryStatusResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListLibraryStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListLibraryStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLibraryStatusResponse) GoString() string {
	return s.String()
}

func (s *ListLibraryStatusResponse) SetHeaders(v map[string]*string) *ListLibraryStatusResponse {
	s.Headers = v
	return s
}

func (s *ListLibraryStatusResponse) SetBody(v *ListLibraryStatusResponseBody) *ListLibraryStatusResponse {
	s.Body = v
	return s
}

type DescribeClusterServiceConfigRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ServiceName     *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ConfigVersion   *string `json:"ConfigVersion,omitempty" xml:"ConfigVersion,omitempty"`
	GroupId         *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	HostInstanceId  *string `json:"HostInstanceId,omitempty" xml:"HostInstanceId,omitempty"`
	TagValue        *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeClusterServiceConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterServiceConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterServiceConfigRequest) SetResourceOwnerId(v int64) *DescribeClusterServiceConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeClusterServiceConfigRequest) SetRegionId(v string) *DescribeClusterServiceConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeClusterServiceConfigRequest) SetClusterId(v string) *DescribeClusterServiceConfigRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterServiceConfigRequest) SetServiceName(v string) *DescribeClusterServiceConfigRequest {
	s.ServiceName = &v
	return s
}

func (s *DescribeClusterServiceConfigRequest) SetConfigVersion(v string) *DescribeClusterServiceConfigRequest {
	s.ConfigVersion = &v
	return s
}

func (s *DescribeClusterServiceConfigRequest) SetGroupId(v string) *DescribeClusterServiceConfigRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeClusterServiceConfigRequest) SetHostInstanceId(v string) *DescribeClusterServiceConfigRequest {
	s.HostInstanceId = &v
	return s
}

func (s *DescribeClusterServiceConfigRequest) SetTagValue(v string) *DescribeClusterServiceConfigRequest {
	s.TagValue = &v
	return s
}

type DescribeClusterServiceConfigResponseBody struct {
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Config    *DescribeClusterServiceConfigResponseBodyConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
}

func (s DescribeClusterServiceConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterServiceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterServiceConfigResponseBody) SetRequestId(v string) *DescribeClusterServiceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterServiceConfigResponseBody) SetConfig(v *DescribeClusterServiceConfigResponseBodyConfig) *DescribeClusterServiceConfigResponseBody {
	s.Config = v
	return s
}

type DescribeClusterServiceConfigResponseBodyConfig struct {
	Applied          *string                                                         `json:"Applied,omitempty" xml:"Applied,omitempty"`
	Comment          *string                                                         `json:"Comment,omitempty" xml:"Comment,omitempty"`
	CreateTime       *string                                                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ConfigVersion    *string                                                         `json:"ConfigVersion,omitempty" xml:"ConfigVersion,omitempty"`
	Author           *string                                                         `json:"Author,omitempty" xml:"Author,omitempty"`
	ServiceName      *string                                                         `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ConfigValueList  *DescribeClusterServiceConfigResponseBodyConfigConfigValueList  `json:"ConfigValueList,omitempty" xml:"ConfigValueList,omitempty" type:"Struct"`
	PropertyInfoList *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoList `json:"PropertyInfoList,omitempty" xml:"PropertyInfoList,omitempty" type:"Struct"`
}

func (s DescribeClusterServiceConfigResponseBodyConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterServiceConfigResponseBodyConfig) GoString() string {
	return s.String()
}

func (s *DescribeClusterServiceConfigResponseBodyConfig) SetApplied(v string) *DescribeClusterServiceConfigResponseBodyConfig {
	s.Applied = &v
	return s
}

func (s *DescribeClusterServiceConfigResponseBodyConfig) SetComment(v string) *DescribeClusterServiceConfigResponseBodyConfig {
	s.Comment = &v
	return s
}

func (s *DescribeClusterServiceConfigResponseBodyConfig) SetCreateTime(v string) *DescribeClusterServiceConfigResponseBodyConfig {
	s.CreateTime = &v
	return s
}

func (s *DescribeClusterServiceConfigResponseBodyConfig) SetConfigVersion(v string) *DescribeClusterServiceConfigResponseBodyConfig {
	s.ConfigVersion = &v
	return s
}

func (s *DescribeClusterServiceConfigResponseBodyConfig) SetAuthor(v string) *DescribeClusterServiceConfigResponseBodyConfig {
	s.Author = &v
	return s
}

func (s *DescribeClusterServiceConfigResponseBodyConfig) SetServiceName(v string) *DescribeClusterServiceConfigResponseBodyConfig {
	s.ServiceName = &v
	return s
}

func (s *DescribeClusterServiceConfigResponseBodyConfig) SetConfigValueList(v *DescribeClusterServiceConfigResponseBodyConfigConfigValueList) *DescribeClusterServiceConfigResponseBodyConfig {
	s.ConfigValueList = v
	return s
}

func (s *DescribeClusterServiceConfigResponseBodyConfig) SetPropertyInfoList(v *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoList) *DescribeClusterServiceConfigResponseBodyConfig {
	s.PropertyInfoList = v
	return s
}

type DescribeClusterServiceConfigResponseBodyConfigConfigValueList struct {
	ConfigValue []*DescribeClusterServiceConfigResponseBodyConfigConfigValueListConfigValue `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty" type:"Repeated"`
}

func (s DescribeClusterServiceConfigResponseBodyConfigConfigValueList) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterServiceConfigResponseBodyConfigConfigValueList) GoString() string {
	return s.String()
}

func (s *DescribeClusterServiceConfigResponseBodyConfigConfigValueList) SetConfigValue(v []*DescribeClusterServiceConfigResponseBodyConfigConfigValueListConfigValue) *DescribeClusterServiceConfigResponseBodyConfigConfigValueList {
	s.ConfigValue = v
	return s
}

type DescribeClusterServiceConfigResponseBodyConfigConfigValueListConfigValue struct {
	ConfigName          *string                                                                                      `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	ScopeId             *int64                                                                                       `json:"ScopeId,omitempty" xml:"ScopeId,omitempty"`
	Scope               *string                                                                                      `json:"Scope,omitempty" xml:"Scope,omitempty"`
	AllowCustom         *bool                                                                                        `json:"AllowCustom,omitempty" xml:"AllowCustom,omitempty"`
	ConfigItemValueList *DescribeClusterServiceConfigResponseBodyConfigConfigValueListConfigValueConfigItemValueList `json:"ConfigItemValueList,omitempty" xml:"ConfigItemValueList,omitempty" type:"Struct"`
}

func (s DescribeClusterServiceConfigResponseBodyConfigConfigValueListConfigValue) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterServiceConfigResponseBodyConfigConfigValueListConfigValue) GoString() string {
	return s.String()
}

func (s *DescribeClusterServiceConfigResponseBodyConfigConfigValueListConfigValue) SetConfigName(v string) *DescribeClusterServiceConfigResponseBodyConfigConfigValueListConfigValue {
	s.ConfigName = &v
	return s
}

func (s *DescribeClusterServiceConfigResponseBodyConfigConfigValueListConfigValue) SetScopeId(v int64) *DescribeClusterServiceConfigResponseBodyConfigConfigValueListConfigValue {
	s.ScopeId = &v
	return s
}

func (s *DescribeClusterServiceConfigResponseBodyConfigConfigValueListConfigValue) SetScope(v string) *DescribeClusterServiceConfigResponseBodyConfigConfigValueListConfigValue {
	s.Scope = &v
	return s
}

func (s *DescribeClusterServiceConfigResponseBodyConfigConfigValueListConfigValue) SetAllowCustom(v bool) *DescribeClusterServiceConfigResponseBodyConfigConfigValueListConfigValue {
	s.AllowCustom = &v
	return s
}

func (s *DescribeClusterServiceConfigResponseBodyConfigConfigValueListConfigValue) SetConfigItemValueList(v *DescribeClusterServiceConfigResponseBodyConfigConfigValueListConfigValueConfigItemValueList) *DescribeClusterServiceConfigResponseBodyConfigConfigValueListConfigValue {
	s.ConfigItemValueList = v
	return s
}

type DescribeClusterServiceConfigResponseBodyConfigConfigValueListConfigValueConfigItemValueList struct {
	ConfigItemValue []*DescribeClusterServiceConfigResponseBodyConfigConfigValueListConfigValueConfigItemValueListConfigItemValue `json:"ConfigItemValue,omitempty" xml:"ConfigItemValue,omitempty" type:"Repeated"`
}

func (s DescribeClusterServiceConfigResponseBodyConfigConfigValueListConfigValueConfigItemValueList) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterServiceConfigResponseBodyConfigConfigValueListConfigValueConfigItemValueList) GoString() string {
	return s.String()
}

func (s *DescribeClusterServiceConfigResponseBodyConfigConfigValueListConfigValueConfigItemValueList) SetConfigItemValue(v []*DescribeClusterServiceConfigResponseBodyConfigConfigValueListConfigValueConfigItemValueListConfigItemValue) *DescribeClusterServiceConfigResponseBodyConfigConfigValueListConfigValueConfigItemValueList {
	s.ConfigItemValue = v
	return s
}

type DescribeClusterServiceConfigResponseBodyConfigConfigValueListConfigValueConfigItemValueListConfigItemValue struct {
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ItemName    *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	IsCustom    *bool   `json:"IsCustom,omitempty" xml:"IsCustom,omitempty"`
}

func (s DescribeClusterServiceConfigResponseBodyConfigConfigValueListConfigValueConfigItemValueListConfigItemValue) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterServiceConfigResponseBodyConfigConfigValueListConfigValueConfigItemValueListConfigItemValue) GoString() string {
	return s.String()
}

func (s *DescribeClusterServiceConfigResponseBodyConfigConfigValueListConfigValueConfigItemValueListConfigItemValue) SetValue(v string) *DescribeClusterServiceConfigResponseBodyConfigConfigValueListConfigValueConfigItemValueListConfigItemValue {
	s.Value = &v
	return s
}

func (s *DescribeClusterServiceConfigResponseBodyConfigConfigValueListConfigValueConfigItemValueListConfigItemValue) SetDescription(v string) *DescribeClusterServiceConfigResponseBodyConfigConfigValueListConfigValueConfigItemValueListConfigItemValue {
	s.Description = &v
	return s
}

func (s *DescribeClusterServiceConfigResponseBodyConfigConfigValueListConfigValueConfigItemValueListConfigItemValue) SetItemName(v string) *DescribeClusterServiceConfigResponseBodyConfigConfigValueListConfigValueConfigItemValueListConfigItemValue {
	s.ItemName = &v
	return s
}

func (s *DescribeClusterServiceConfigResponseBodyConfigConfigValueListConfigValueConfigItemValueListConfigItemValue) SetIsCustom(v bool) *DescribeClusterServiceConfigResponseBodyConfigConfigValueListConfigValueConfigItemValueListConfigItemValue {
	s.IsCustom = &v
	return s
}

type DescribeClusterServiceConfigResponseBodyConfigPropertyInfoList struct {
	PropertyInfo []*DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfo `json:"PropertyInfo,omitempty" xml:"PropertyInfo,omitempty" type:"Repeated"`
}

func (s DescribeClusterServiceConfigResponseBodyConfigPropertyInfoList) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterServiceConfigResponseBodyConfigPropertyInfoList) GoString() string {
	return s.String()
}

func (s *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoList) SetPropertyInfo(v []*DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfo) *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoList {
	s.PropertyInfo = v
	return s
}

type DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfo struct {
	DisplayName             *string                                                                                            `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Value                   *string                                                                                            `json:"Value,omitempty" xml:"Value,omitempty"`
	Description             *string                                                                                            `json:"Description,omitempty" xml:"Description,omitempty"`
	Component               *string                                                                                            `json:"Component,omitempty" xml:"Component,omitempty"`
	FileName                *string                                                                                            `json:"FileName,omitempty" xml:"FileName,omitempty"`
	Name                    *string                                                                                            `json:"Name,omitempty" xml:"Name,omitempty"`
	ServiceName             *string                                                                                            `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	PropertyTypes           *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyTypes           `json:"PropertyTypes,omitempty" xml:"PropertyTypes,omitempty" type:"Struct"`
	PropertyValueAttributes *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyValueAttributes `json:"PropertyValueAttributes,omitempty" xml:"PropertyValueAttributes,omitempty" type:"Struct"`
	EffectWay               *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoEffectWay               `json:"EffectWay,omitempty" xml:"EffectWay,omitempty" type:"Struct"`
}

func (s DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfo) SetDisplayName(v string) *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfo {
	s.DisplayName = &v
	return s
}

func (s *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfo) SetValue(v string) *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfo {
	s.Value = &v
	return s
}

func (s *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfo) SetDescription(v string) *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfo {
	s.Description = &v
	return s
}

func (s *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfo) SetComponent(v string) *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfo {
	s.Component = &v
	return s
}

func (s *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfo) SetFileName(v string) *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfo {
	s.FileName = &v
	return s
}

func (s *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfo) SetName(v string) *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfo {
	s.Name = &v
	return s
}

func (s *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfo) SetServiceName(v string) *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfo {
	s.ServiceName = &v
	return s
}

func (s *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfo) SetPropertyTypes(v *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyTypes) *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfo {
	s.PropertyTypes = v
	return s
}

func (s *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfo) SetPropertyValueAttributes(v *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyValueAttributes) *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfo {
	s.PropertyValueAttributes = v
	return s
}

func (s *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfo) SetEffectWay(v *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoEffectWay) *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfo {
	s.EffectWay = v
	return s
}

type DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyTypes struct {
	PropertyType []*string `json:"PropertyType,omitempty" xml:"PropertyType,omitempty" type:"Repeated"`
}

func (s DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyTypes) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyTypes) GoString() string {
	return s.String()
}

func (s *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyTypes) SetPropertyType(v []*string) *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyTypes {
	s.PropertyType = v
	return s
}

type DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyValueAttributes struct {
	Type          *string                                                                                                   `json:"Type,omitempty" xml:"Type,omitempty"`
	Maximum       *string                                                                                                   `json:"Maximum,omitempty" xml:"Maximum,omitempty"`
	Unit          *string                                                                                                   `json:"Unit,omitempty" xml:"Unit,omitempty"`
	Hidden        *bool                                                                                                     `json:"Hidden,omitempty" xml:"Hidden,omitempty"`
	IncrememtStep *string                                                                                                   `json:"IncrememtStep,omitempty" xml:"IncrememtStep,omitempty"`
	ReadOnly      *bool                                                                                                     `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	Mimimum       *string                                                                                                   `json:"Mimimum,omitempty" xml:"Mimimum,omitempty"`
	Entries       *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyValueAttributesEntries `json:"Entries,omitempty" xml:"Entries,omitempty" type:"Struct"`
}

func (s DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyValueAttributes) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyValueAttributes) GoString() string {
	return s.String()
}

func (s *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyValueAttributes) SetType(v string) *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyValueAttributes {
	s.Type = &v
	return s
}

func (s *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyValueAttributes) SetMaximum(v string) *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyValueAttributes {
	s.Maximum = &v
	return s
}

func (s *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyValueAttributes) SetUnit(v string) *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyValueAttributes {
	s.Unit = &v
	return s
}

func (s *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyValueAttributes) SetHidden(v bool) *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyValueAttributes {
	s.Hidden = &v
	return s
}

func (s *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyValueAttributes) SetIncrememtStep(v string) *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyValueAttributes {
	s.IncrememtStep = &v
	return s
}

func (s *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyValueAttributes) SetReadOnly(v bool) *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyValueAttributes {
	s.ReadOnly = &v
	return s
}

func (s *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyValueAttributes) SetMimimum(v string) *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyValueAttributes {
	s.Mimimum = &v
	return s
}

func (s *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyValueAttributes) SetEntries(v *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyValueAttributesEntries) *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyValueAttributes {
	s.Entries = v
	return s
}

type DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyValueAttributesEntries struct {
	ValueEntryInfo []*DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyValueAttributesEntriesValueEntryInfo `json:"ValueEntryInfo,omitempty" xml:"ValueEntryInfo,omitempty" type:"Repeated"`
}

func (s DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyValueAttributesEntries) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyValueAttributesEntries) GoString() string {
	return s.String()
}

func (s *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyValueAttributesEntries) SetValueEntryInfo(v []*DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyValueAttributesEntriesValueEntryInfo) *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyValueAttributesEntries {
	s.ValueEntryInfo = v
	return s
}

type DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyValueAttributesEntriesValueEntryInfo struct {
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Label       *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyValueAttributesEntriesValueEntryInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyValueAttributesEntriesValueEntryInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyValueAttributesEntriesValueEntryInfo) SetValue(v string) *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyValueAttributesEntriesValueEntryInfo {
	s.Value = &v
	return s
}

func (s *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyValueAttributesEntriesValueEntryInfo) SetLabel(v string) *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyValueAttributesEntriesValueEntryInfo {
	s.Label = &v
	return s
}

func (s *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyValueAttributesEntriesValueEntryInfo) SetDescription(v string) *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoPropertyValueAttributesEntriesValueEntryInfo {
	s.Description = &v
	return s
}

type DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoEffectWay struct {
	EffectType        *string `json:"EffectType,omitempty" xml:"EffectType,omitempty"`
	InvokeServiceName *string `json:"InvokeServiceName,omitempty" xml:"InvokeServiceName,omitempty"`
}

func (s DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoEffectWay) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoEffectWay) GoString() string {
	return s.String()
}

func (s *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoEffectWay) SetEffectType(v string) *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoEffectWay {
	s.EffectType = &v
	return s
}

func (s *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoEffectWay) SetInvokeServiceName(v string) *DescribeClusterServiceConfigResponseBodyConfigPropertyInfoListPropertyInfoEffectWay {
	s.InvokeServiceName = &v
	return s
}

type DescribeClusterServiceConfigResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClusterServiceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClusterServiceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterServiceConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterServiceConfigResponse) SetHeaders(v map[string]*string) *DescribeClusterServiceConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterServiceConfigResponse) SetBody(v *DescribeClusterServiceConfigResponseBody) *DescribeClusterServiceConfigResponse {
	s.Body = v
	return s
}

type ModifyFlowProjectClusterSettingRequest struct {
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProjectId    *string   `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ClusterId    *string   `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DefaultUser  *string   `json:"DefaultUser,omitempty" xml:"DefaultUser,omitempty"`
	DefaultQueue *string   `json:"DefaultQueue,omitempty" xml:"DefaultQueue,omitempty"`
	UserList     []*string `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
	QueueList    []*string `json:"QueueList,omitempty" xml:"QueueList,omitempty" type:"Repeated"`
	HostList     []*string `json:"HostList,omitempty" xml:"HostList,omitempty" type:"Repeated"`
}

func (s ModifyFlowProjectClusterSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowProjectClusterSettingRequest) GoString() string {
	return s.String()
}

func (s *ModifyFlowProjectClusterSettingRequest) SetRegionId(v string) *ModifyFlowProjectClusterSettingRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyFlowProjectClusterSettingRequest) SetProjectId(v string) *ModifyFlowProjectClusterSettingRequest {
	s.ProjectId = &v
	return s
}

func (s *ModifyFlowProjectClusterSettingRequest) SetClusterId(v string) *ModifyFlowProjectClusterSettingRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyFlowProjectClusterSettingRequest) SetDefaultUser(v string) *ModifyFlowProjectClusterSettingRequest {
	s.DefaultUser = &v
	return s
}

func (s *ModifyFlowProjectClusterSettingRequest) SetDefaultQueue(v string) *ModifyFlowProjectClusterSettingRequest {
	s.DefaultQueue = &v
	return s
}

func (s *ModifyFlowProjectClusterSettingRequest) SetUserList(v []*string) *ModifyFlowProjectClusterSettingRequest {
	s.UserList = v
	return s
}

func (s *ModifyFlowProjectClusterSettingRequest) SetQueueList(v []*string) *ModifyFlowProjectClusterSettingRequest {
	s.QueueList = v
	return s
}

func (s *ModifyFlowProjectClusterSettingRequest) SetHostList(v []*string) *ModifyFlowProjectClusterSettingRequest {
	s.HostList = v
	return s
}

type ModifyFlowProjectClusterSettingResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyFlowProjectClusterSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowProjectClusterSettingResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFlowProjectClusterSettingResponseBody) SetData(v bool) *ModifyFlowProjectClusterSettingResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyFlowProjectClusterSettingResponseBody) SetRequestId(v string) *ModifyFlowProjectClusterSettingResponseBody {
	s.RequestId = &v
	return s
}

type ModifyFlowProjectClusterSettingResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyFlowProjectClusterSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyFlowProjectClusterSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyFlowProjectClusterSettingResponse) GoString() string {
	return s.String()
}

func (s *ModifyFlowProjectClusterSettingResponse) SetHeaders(v map[string]*string) *ModifyFlowProjectClusterSettingResponse {
	s.Headers = v
	return s
}

func (s *ModifyFlowProjectClusterSettingResponse) SetBody(v *ModifyFlowProjectClusterSettingResponseBody) *ModifyFlowProjectClusterSettingResponse {
	s.Body = v
	return s
}

type DeleteFlowProjectUserRequest struct {
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	UserName  *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DeleteFlowProjectUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowProjectUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteFlowProjectUserRequest) SetRegionId(v string) *DeleteFlowProjectUserRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteFlowProjectUserRequest) SetProjectId(v string) *DeleteFlowProjectUserRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteFlowProjectUserRequest) SetUserName(v string) *DeleteFlowProjectUserRequest {
	s.UserName = &v
	return s
}

type DeleteFlowProjectUserResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFlowProjectUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowProjectUserResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFlowProjectUserResponseBody) SetData(v bool) *DeleteFlowProjectUserResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteFlowProjectUserResponseBody) SetRequestId(v string) *DeleteFlowProjectUserResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFlowProjectUserResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFlowProjectUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFlowProjectUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFlowProjectUserResponse) GoString() string {
	return s.String()
}

func (s *DeleteFlowProjectUserResponse) SetHeaders(v map[string]*string) *DeleteFlowProjectUserResponse {
	s.Headers = v
	return s
}

func (s *DeleteFlowProjectUserResponse) SetBody(v *DeleteFlowProjectUserResponseBody) *DeleteFlowProjectUserResponse {
	s.Body = v
	return s
}

type CreateClusterV2Request struct {
	ResourceOwnerId        *int64                                     `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Name                   *string                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId               *string                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId                 *string                                    `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	LogPath                *string                                    `json:"LogPath,omitempty" xml:"LogPath,omitempty"`
	SecurityGroupId        *string                                    `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	IsOpenPublicIp         *bool                                      `json:"IsOpenPublicIp,omitempty" xml:"IsOpenPublicIp,omitempty"`
	SecurityGroupName      *string                                    `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
	ChargeType             *string                                    `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Period                 *int32                                     `json:"Period,omitempty" xml:"Period,omitempty"`
	Auto                   *bool                                      `json:"Auto,omitempty" xml:"Auto,omitempty"`
	AutoPayOrder           *bool                                      `json:"AutoPayOrder,omitempty" xml:"AutoPayOrder,omitempty"`
	VpcId                  *string                                    `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VSwitchId              *string                                    `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	NetType                *string                                    `json:"NetType,omitempty" xml:"NetType,omitempty"`
	UserDefinedEmrEcsRole  *string                                    `json:"UserDefinedEmrEcsRole,omitempty" xml:"UserDefinedEmrEcsRole,omitempty"`
	EmrVer                 *string                                    `json:"EmrVer,omitempty" xml:"EmrVer,omitempty"`
	ClusterType            *string                                    `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	EnableHighAvailability *bool                                      `json:"EnableHighAvailability,omitempty" xml:"EnableHighAvailability,omitempty"`
	UseLocalMetaDb         *bool                                      `json:"UseLocalMetaDb,omitempty" xml:"UseLocalMetaDb,omitempty"`
	EnableSsh              *bool                                      `json:"EnableSsh,omitempty" xml:"EnableSsh,omitempty"`
	InstanceGeneration     *string                                    `json:"InstanceGeneration,omitempty" xml:"InstanceGeneration,omitempty"`
	MasterPwd              *string                                    `json:"MasterPwd,omitempty" xml:"MasterPwd,omitempty"`
	KeyPairName            *string                                    `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	MetaStoreType          *string                                    `json:"MetaStoreType,omitempty" xml:"MetaStoreType,omitempty"`
	MetaStoreConf          *string                                    `json:"MetaStoreConf,omitempty" xml:"MetaStoreConf,omitempty"`
	ClickHouseConf         *string                                    `json:"ClickHouseConf,omitempty" xml:"ClickHouseConf,omitempty"`
	ExtraAttributes        *string                                    `json:"ExtraAttributes,omitempty" xml:"ExtraAttributes,omitempty"`
	DepositType            *string                                    `json:"DepositType,omitempty" xml:"DepositType,omitempty"`
	MachineType            *string                                    `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	UseCustomHiveMetaDB    *bool                                      `json:"UseCustomHiveMetaDB,omitempty" xml:"UseCustomHiveMetaDB,omitempty"`
	InitCustomHiveMetaDB   *bool                                      `json:"InitCustomHiveMetaDB,omitempty" xml:"InitCustomHiveMetaDB,omitempty"`
	Configurations         *string                                    `json:"Configurations,omitempty" xml:"Configurations,omitempty"`
	EnableEas              *bool                                      `json:"EnableEas,omitempty" xml:"EnableEas,omitempty"`
	RelatedClusterId       *string                                    `json:"RelatedClusterId,omitempty" xml:"RelatedClusterId,omitempty"`
	WhiteListType          *string                                    `json:"WhiteListType,omitempty" xml:"WhiteListType,omitempty"`
	AuthorizeContent       *string                                    `json:"AuthorizeContent,omitempty" xml:"AuthorizeContent,omitempty"`
	ResourceGroupId        *string                                    `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ClientToken            *string                                    `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	UserInfo               []*CreateClusterV2RequestUserInfo          `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Repeated"`
	HostComponentInfo      []*CreateClusterV2RequestHostComponentInfo `json:"HostComponentInfo,omitempty" xml:"HostComponentInfo,omitempty" type:"Repeated"`
	ServiceInfo            []*CreateClusterV2RequestServiceInfo       `json:"ServiceInfo,omitempty" xml:"ServiceInfo,omitempty" type:"Repeated"`
	PromotionInfo          []*CreateClusterV2RequestPromotionInfo     `json:"PromotionInfo,omitempty" xml:"PromotionInfo,omitempty" type:"Repeated"`
	HostGroup              []*CreateClusterV2RequestHostGroup         `json:"HostGroup,omitempty" xml:"HostGroup,omitempty" type:"Repeated"`
	BootstrapAction        []*CreateClusterV2RequestBootstrapAction   `json:"BootstrapAction,omitempty" xml:"BootstrapAction,omitempty" type:"Repeated"`
	Config                 []*CreateClusterV2RequestConfig            `json:"Config,omitempty" xml:"Config,omitempty" type:"Repeated"`
	Tag                    []*CreateClusterV2RequestTag               `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateClusterV2Request) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterV2Request) GoString() string {
	return s.String()
}

func (s *CreateClusterV2Request) SetResourceOwnerId(v int64) *CreateClusterV2Request {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateClusterV2Request) SetName(v string) *CreateClusterV2Request {
	s.Name = &v
	return s
}

func (s *CreateClusterV2Request) SetRegionId(v string) *CreateClusterV2Request {
	s.RegionId = &v
	return s
}

func (s *CreateClusterV2Request) SetZoneId(v string) *CreateClusterV2Request {
	s.ZoneId = &v
	return s
}

func (s *CreateClusterV2Request) SetLogPath(v string) *CreateClusterV2Request {
	s.LogPath = &v
	return s
}

func (s *CreateClusterV2Request) SetSecurityGroupId(v string) *CreateClusterV2Request {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateClusterV2Request) SetIsOpenPublicIp(v bool) *CreateClusterV2Request {
	s.IsOpenPublicIp = &v
	return s
}

func (s *CreateClusterV2Request) SetSecurityGroupName(v string) *CreateClusterV2Request {
	s.SecurityGroupName = &v
	return s
}

func (s *CreateClusterV2Request) SetChargeType(v string) *CreateClusterV2Request {
	s.ChargeType = &v
	return s
}

func (s *CreateClusterV2Request) SetPeriod(v int32) *CreateClusterV2Request {
	s.Period = &v
	return s
}

func (s *CreateClusterV2Request) SetAuto(v bool) *CreateClusterV2Request {
	s.Auto = &v
	return s
}

func (s *CreateClusterV2Request) SetAutoPayOrder(v bool) *CreateClusterV2Request {
	s.AutoPayOrder = &v
	return s
}

func (s *CreateClusterV2Request) SetVpcId(v string) *CreateClusterV2Request {
	s.VpcId = &v
	return s
}

func (s *CreateClusterV2Request) SetVSwitchId(v string) *CreateClusterV2Request {
	s.VSwitchId = &v
	return s
}

func (s *CreateClusterV2Request) SetNetType(v string) *CreateClusterV2Request {
	s.NetType = &v
	return s
}

func (s *CreateClusterV2Request) SetUserDefinedEmrEcsRole(v string) *CreateClusterV2Request {
	s.UserDefinedEmrEcsRole = &v
	return s
}

func (s *CreateClusterV2Request) SetEmrVer(v string) *CreateClusterV2Request {
	s.EmrVer = &v
	return s
}

func (s *CreateClusterV2Request) SetClusterType(v string) *CreateClusterV2Request {
	s.ClusterType = &v
	return s
}

func (s *CreateClusterV2Request) SetEnableHighAvailability(v bool) *CreateClusterV2Request {
	s.EnableHighAvailability = &v
	return s
}

func (s *CreateClusterV2Request) SetUseLocalMetaDb(v bool) *CreateClusterV2Request {
	s.UseLocalMetaDb = &v
	return s
}

func (s *CreateClusterV2Request) SetEnableSsh(v bool) *CreateClusterV2Request {
	s.EnableSsh = &v
	return s
}

func (s *CreateClusterV2Request) SetInstanceGeneration(v string) *CreateClusterV2Request {
	s.InstanceGeneration = &v
	return s
}

func (s *CreateClusterV2Request) SetMasterPwd(v string) *CreateClusterV2Request {
	s.MasterPwd = &v
	return s
}

func (s *CreateClusterV2Request) SetKeyPairName(v string) *CreateClusterV2Request {
	s.KeyPairName = &v
	return s
}

func (s *CreateClusterV2Request) SetMetaStoreType(v string) *CreateClusterV2Request {
	s.MetaStoreType = &v
	return s
}

func (s *CreateClusterV2Request) SetMetaStoreConf(v string) *CreateClusterV2Request {
	s.MetaStoreConf = &v
	return s
}

func (s *CreateClusterV2Request) SetClickHouseConf(v string) *CreateClusterV2Request {
	s.ClickHouseConf = &v
	return s
}

func (s *CreateClusterV2Request) SetExtraAttributes(v string) *CreateClusterV2Request {
	s.ExtraAttributes = &v
	return s
}

func (s *CreateClusterV2Request) SetDepositType(v string) *CreateClusterV2Request {
	s.DepositType = &v
	return s
}

func (s *CreateClusterV2Request) SetMachineType(v string) *CreateClusterV2Request {
	s.MachineType = &v
	return s
}

func (s *CreateClusterV2Request) SetUseCustomHiveMetaDB(v bool) *CreateClusterV2Request {
	s.UseCustomHiveMetaDB = &v
	return s
}

func (s *CreateClusterV2Request) SetInitCustomHiveMetaDB(v bool) *CreateClusterV2Request {
	s.InitCustomHiveMetaDB = &v
	return s
}

func (s *CreateClusterV2Request) SetConfigurations(v string) *CreateClusterV2Request {
	s.Configurations = &v
	return s
}

func (s *CreateClusterV2Request) SetEnableEas(v bool) *CreateClusterV2Request {
	s.EnableEas = &v
	return s
}

func (s *CreateClusterV2Request) SetRelatedClusterId(v string) *CreateClusterV2Request {
	s.RelatedClusterId = &v
	return s
}

func (s *CreateClusterV2Request) SetWhiteListType(v string) *CreateClusterV2Request {
	s.WhiteListType = &v
	return s
}

func (s *CreateClusterV2Request) SetAuthorizeContent(v string) *CreateClusterV2Request {
	s.AuthorizeContent = &v
	return s
}

func (s *CreateClusterV2Request) SetResourceGroupId(v string) *CreateClusterV2Request {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateClusterV2Request) SetClientToken(v string) *CreateClusterV2Request {
	s.ClientToken = &v
	return s
}

func (s *CreateClusterV2Request) SetUserInfo(v []*CreateClusterV2RequestUserInfo) *CreateClusterV2Request {
	s.UserInfo = v
	return s
}

func (s *CreateClusterV2Request) SetHostComponentInfo(v []*CreateClusterV2RequestHostComponentInfo) *CreateClusterV2Request {
	s.HostComponentInfo = v
	return s
}

func (s *CreateClusterV2Request) SetServiceInfo(v []*CreateClusterV2RequestServiceInfo) *CreateClusterV2Request {
	s.ServiceInfo = v
	return s
}

func (s *CreateClusterV2Request) SetPromotionInfo(v []*CreateClusterV2RequestPromotionInfo) *CreateClusterV2Request {
	s.PromotionInfo = v
	return s
}

func (s *CreateClusterV2Request) SetHostGroup(v []*CreateClusterV2RequestHostGroup) *CreateClusterV2Request {
	s.HostGroup = v
	return s
}

func (s *CreateClusterV2Request) SetBootstrapAction(v []*CreateClusterV2RequestBootstrapAction) *CreateClusterV2Request {
	s.BootstrapAction = v
	return s
}

func (s *CreateClusterV2Request) SetConfig(v []*CreateClusterV2RequestConfig) *CreateClusterV2Request {
	s.Config = v
	return s
}

func (s *CreateClusterV2Request) SetTag(v []*CreateClusterV2RequestTag) *CreateClusterV2Request {
	s.Tag = v
	return s
}

type CreateClusterV2RequestUserInfo struct {
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateClusterV2RequestUserInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterV2RequestUserInfo) GoString() string {
	return s.String()
}

func (s *CreateClusterV2RequestUserInfo) SetPassword(v string) *CreateClusterV2RequestUserInfo {
	s.Password = &v
	return s
}

func (s *CreateClusterV2RequestUserInfo) SetUserId(v string) *CreateClusterV2RequestUserInfo {
	s.UserId = &v
	return s
}

func (s *CreateClusterV2RequestUserInfo) SetUserName(v string) *CreateClusterV2RequestUserInfo {
	s.UserName = &v
	return s
}

type CreateClusterV2RequestHostComponentInfo struct {
	ComponentNameList []*string `json:"ComponentNameList,omitempty" xml:"ComponentNameList,omitempty" type:"Repeated"`
	HostName          *string   `json:"HostName,omitempty" xml:"HostName,omitempty"`
	ServiceName       *string   `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s CreateClusterV2RequestHostComponentInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterV2RequestHostComponentInfo) GoString() string {
	return s.String()
}

func (s *CreateClusterV2RequestHostComponentInfo) SetComponentNameList(v []*string) *CreateClusterV2RequestHostComponentInfo {
	s.ComponentNameList = v
	return s
}

func (s *CreateClusterV2RequestHostComponentInfo) SetHostName(v string) *CreateClusterV2RequestHostComponentInfo {
	s.HostName = &v
	return s
}

func (s *CreateClusterV2RequestHostComponentInfo) SetServiceName(v string) *CreateClusterV2RequestHostComponentInfo {
	s.ServiceName = &v
	return s
}

type CreateClusterV2RequestServiceInfo struct {
	ServiceName    *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
}

func (s CreateClusterV2RequestServiceInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterV2RequestServiceInfo) GoString() string {
	return s.String()
}

func (s *CreateClusterV2RequestServiceInfo) SetServiceName(v string) *CreateClusterV2RequestServiceInfo {
	s.ServiceName = &v
	return s
}

func (s *CreateClusterV2RequestServiceInfo) SetServiceVersion(v string) *CreateClusterV2RequestServiceInfo {
	s.ServiceVersion = &v
	return s
}

type CreateClusterV2RequestPromotionInfo struct {
	ProductCode         *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	PromotionOptionCode *string `json:"PromotionOptionCode,omitempty" xml:"PromotionOptionCode,omitempty"`
	PromotionOptionNo   *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
}

func (s CreateClusterV2RequestPromotionInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterV2RequestPromotionInfo) GoString() string {
	return s.String()
}

func (s *CreateClusterV2RequestPromotionInfo) SetProductCode(v string) *CreateClusterV2RequestPromotionInfo {
	s.ProductCode = &v
	return s
}

func (s *CreateClusterV2RequestPromotionInfo) SetPromotionOptionCode(v string) *CreateClusterV2RequestPromotionInfo {
	s.PromotionOptionCode = &v
	return s
}

func (s *CreateClusterV2RequestPromotionInfo) SetPromotionOptionNo(v string) *CreateClusterV2RequestPromotionInfo {
	s.PromotionOptionNo = &v
	return s
}

type CreateClusterV2RequestHostGroup struct {
	AutoRenew       *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	ChargeType      *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Comment         *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	CreateType      *string `json:"CreateType,omitempty" xml:"CreateType,omitempty"`
	DiskCapacity    *int32  `json:"DiskCapacity,omitempty" xml:"DiskCapacity,omitempty"`
	DiskCount       *int32  `json:"DiskCount,omitempty" xml:"DiskCount,omitempty"`
	DiskType        *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	GpuDriver       *string `json:"GpuDriver,omitempty" xml:"GpuDriver,omitempty"`
	HostGroupId     *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	HostGroupName   *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
	HostGroupType   *string `json:"HostGroupType,omitempty" xml:"HostGroupType,omitempty"`
	InstanceType    *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	NodeCount       *int32  `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	Period          *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	SysDiskCapacity *int32  `json:"SysDiskCapacity,omitempty" xml:"SysDiskCapacity,omitempty"`
	SysDiskType     *string `json:"SysDiskType,omitempty" xml:"SysDiskType,omitempty"`
	VSwitchId       *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateClusterV2RequestHostGroup) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterV2RequestHostGroup) GoString() string {
	return s.String()
}

func (s *CreateClusterV2RequestHostGroup) SetAutoRenew(v bool) *CreateClusterV2RequestHostGroup {
	s.AutoRenew = &v
	return s
}

func (s *CreateClusterV2RequestHostGroup) SetChargeType(v string) *CreateClusterV2RequestHostGroup {
	s.ChargeType = &v
	return s
}

func (s *CreateClusterV2RequestHostGroup) SetClusterId(v string) *CreateClusterV2RequestHostGroup {
	s.ClusterId = &v
	return s
}

func (s *CreateClusterV2RequestHostGroup) SetComment(v string) *CreateClusterV2RequestHostGroup {
	s.Comment = &v
	return s
}

func (s *CreateClusterV2RequestHostGroup) SetCreateType(v string) *CreateClusterV2RequestHostGroup {
	s.CreateType = &v
	return s
}

func (s *CreateClusterV2RequestHostGroup) SetDiskCapacity(v int32) *CreateClusterV2RequestHostGroup {
	s.DiskCapacity = &v
	return s
}

func (s *CreateClusterV2RequestHostGroup) SetDiskCount(v int32) *CreateClusterV2RequestHostGroup {
	s.DiskCount = &v
	return s
}

func (s *CreateClusterV2RequestHostGroup) SetDiskType(v string) *CreateClusterV2RequestHostGroup {
	s.DiskType = &v
	return s
}

func (s *CreateClusterV2RequestHostGroup) SetGpuDriver(v string) *CreateClusterV2RequestHostGroup {
	s.GpuDriver = &v
	return s
}

func (s *CreateClusterV2RequestHostGroup) SetHostGroupId(v string) *CreateClusterV2RequestHostGroup {
	s.HostGroupId = &v
	return s
}

func (s *CreateClusterV2RequestHostGroup) SetHostGroupName(v string) *CreateClusterV2RequestHostGroup {
	s.HostGroupName = &v
	return s
}

func (s *CreateClusterV2RequestHostGroup) SetHostGroupType(v string) *CreateClusterV2RequestHostGroup {
	s.HostGroupType = &v
	return s
}

func (s *CreateClusterV2RequestHostGroup) SetInstanceType(v string) *CreateClusterV2RequestHostGroup {
	s.InstanceType = &v
	return s
}

func (s *CreateClusterV2RequestHostGroup) SetNodeCount(v int32) *CreateClusterV2RequestHostGroup {
	s.NodeCount = &v
	return s
}

func (s *CreateClusterV2RequestHostGroup) SetPeriod(v int32) *CreateClusterV2RequestHostGroup {
	s.Period = &v
	return s
}

func (s *CreateClusterV2RequestHostGroup) SetSysDiskCapacity(v int32) *CreateClusterV2RequestHostGroup {
	s.SysDiskCapacity = &v
	return s
}

func (s *CreateClusterV2RequestHostGroup) SetSysDiskType(v string) *CreateClusterV2RequestHostGroup {
	s.SysDiskType = &v
	return s
}

func (s *CreateClusterV2RequestHostGroup) SetVSwitchId(v string) *CreateClusterV2RequestHostGroup {
	s.VSwitchId = &v
	return s
}

type CreateClusterV2RequestBootstrapAction struct {
	Arg  *string `json:"Arg,omitempty" xml:"Arg,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s CreateClusterV2RequestBootstrapAction) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterV2RequestBootstrapAction) GoString() string {
	return s.String()
}

func (s *CreateClusterV2RequestBootstrapAction) SetArg(v string) *CreateClusterV2RequestBootstrapAction {
	s.Arg = &v
	return s
}

func (s *CreateClusterV2RequestBootstrapAction) SetName(v string) *CreateClusterV2RequestBootstrapAction {
	s.Name = &v
	return s
}

func (s *CreateClusterV2RequestBootstrapAction) SetPath(v string) *CreateClusterV2RequestBootstrapAction {
	s.Path = &v
	return s
}

type CreateClusterV2RequestConfig struct {
	ConfigKey   *string `json:"ConfigKey,omitempty" xml:"ConfigKey,omitempty"`
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	Encrypt     *string `json:"Encrypt,omitempty" xml:"Encrypt,omitempty"`
	FileName    *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	Replace     *string `json:"Replace,omitempty" xml:"Replace,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s CreateClusterV2RequestConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterV2RequestConfig) GoString() string {
	return s.String()
}

func (s *CreateClusterV2RequestConfig) SetConfigKey(v string) *CreateClusterV2RequestConfig {
	s.ConfigKey = &v
	return s
}

func (s *CreateClusterV2RequestConfig) SetConfigValue(v string) *CreateClusterV2RequestConfig {
	s.ConfigValue = &v
	return s
}

func (s *CreateClusterV2RequestConfig) SetEncrypt(v string) *CreateClusterV2RequestConfig {
	s.Encrypt = &v
	return s
}

func (s *CreateClusterV2RequestConfig) SetFileName(v string) *CreateClusterV2RequestConfig {
	s.FileName = &v
	return s
}

func (s *CreateClusterV2RequestConfig) SetReplace(v string) *CreateClusterV2RequestConfig {
	s.Replace = &v
	return s
}

func (s *CreateClusterV2RequestConfig) SetServiceName(v string) *CreateClusterV2RequestConfig {
	s.ServiceName = &v
	return s
}

type CreateClusterV2RequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateClusterV2RequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterV2RequestTag) GoString() string {
	return s.String()
}

func (s *CreateClusterV2RequestTag) SetKey(v string) *CreateClusterV2RequestTag {
	s.Key = &v
	return s
}

func (s *CreateClusterV2RequestTag) SetValue(v string) *CreateClusterV2RequestTag {
	s.Value = &v
	return s
}

type CreateClusterV2ResponseBody struct {
	ClusterId     *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CoreOrderId   *string `json:"CoreOrderId,omitempty" xml:"CoreOrderId,omitempty"`
	EmrOrderId    *string `json:"EmrOrderId,omitempty" xml:"EmrOrderId,omitempty"`
	MasterOrderId *string `json:"MasterOrderId,omitempty" xml:"MasterOrderId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateClusterV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterV2ResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterV2ResponseBody) SetClusterId(v string) *CreateClusterV2ResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateClusterV2ResponseBody) SetCoreOrderId(v string) *CreateClusterV2ResponseBody {
	s.CoreOrderId = &v
	return s
}

func (s *CreateClusterV2ResponseBody) SetEmrOrderId(v string) *CreateClusterV2ResponseBody {
	s.EmrOrderId = &v
	return s
}

func (s *CreateClusterV2ResponseBody) SetMasterOrderId(v string) *CreateClusterV2ResponseBody {
	s.MasterOrderId = &v
	return s
}

func (s *CreateClusterV2ResponseBody) SetRequestId(v string) *CreateClusterV2ResponseBody {
	s.RequestId = &v
	return s
}

type CreateClusterV2Response struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateClusterV2ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateClusterV2Response) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterV2Response) GoString() string {
	return s.String()
}

func (s *CreateClusterV2Response) SetHeaders(v map[string]*string) *CreateClusterV2Response {
	s.Headers = v
	return s
}

func (s *CreateClusterV2Response) SetBody(v *CreateClusterV2ResponseBody) *CreateClusterV2Response {
	s.Body = v
	return s
}

type ModifyClusterNameRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Id              *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ModifyClusterNameRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterNameRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterNameRequest) SetResourceOwnerId(v int64) *ModifyClusterNameRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyClusterNameRequest) SetRegionId(v string) *ModifyClusterNameRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyClusterNameRequest) SetId(v string) *ModifyClusterNameRequest {
	s.Id = &v
	return s
}

func (s *ModifyClusterNameRequest) SetName(v string) *ModifyClusterNameRequest {
	s.Name = &v
	return s
}

type ModifyClusterNameResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyClusterNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterNameResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterNameResponseBody) SetRequestId(v string) *ModifyClusterNameResponseBody {
	s.RequestId = &v
	return s
}

type ModifyClusterNameResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyClusterNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyClusterNameResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterNameResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterNameResponse) SetHeaders(v map[string]*string) *ModifyClusterNameResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterNameResponse) SetBody(v *ModifyClusterNameResponseBody) *ModifyClusterNameResponse {
	s.Body = v
	return s
}

type ListClusterOperationHostTaskRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OperationId     *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	HostId          *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListClusterOperationHostTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClusterOperationHostTaskRequest) GoString() string {
	return s.String()
}

func (s *ListClusterOperationHostTaskRequest) SetResourceOwnerId(v int64) *ListClusterOperationHostTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListClusterOperationHostTaskRequest) SetRegionId(v string) *ListClusterOperationHostTaskRequest {
	s.RegionId = &v
	return s
}

func (s *ListClusterOperationHostTaskRequest) SetClusterId(v string) *ListClusterOperationHostTaskRequest {
	s.ClusterId = &v
	return s
}

func (s *ListClusterOperationHostTaskRequest) SetOperationId(v string) *ListClusterOperationHostTaskRequest {
	s.OperationId = &v
	return s
}

func (s *ListClusterOperationHostTaskRequest) SetHostId(v string) *ListClusterOperationHostTaskRequest {
	s.HostId = &v
	return s
}

func (s *ListClusterOperationHostTaskRequest) SetStatus(v string) *ListClusterOperationHostTaskRequest {
	s.Status = &v
	return s
}

func (s *ListClusterOperationHostTaskRequest) SetPageNumber(v int32) *ListClusterOperationHostTaskRequest {
	s.PageNumber = &v
	return s
}

func (s *ListClusterOperationHostTaskRequest) SetPageSize(v int32) *ListClusterOperationHostTaskRequest {
	s.PageSize = &v
	return s
}

type ListClusterOperationHostTaskResponseBody struct {
	PageSize                     *int32                                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId                    *string                                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber                   *int32                                                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCount                   *int32                                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	ClusterOperationHostTaskList *ListClusterOperationHostTaskResponseBodyClusterOperationHostTaskList `json:"ClusterOperationHostTaskList,omitempty" xml:"ClusterOperationHostTaskList,omitempty" type:"Struct"`
}

func (s ListClusterOperationHostTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClusterOperationHostTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterOperationHostTaskResponseBody) SetPageSize(v int32) *ListClusterOperationHostTaskResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListClusterOperationHostTaskResponseBody) SetRequestId(v string) *ListClusterOperationHostTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterOperationHostTaskResponseBody) SetPageNumber(v int32) *ListClusterOperationHostTaskResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListClusterOperationHostTaskResponseBody) SetTotalCount(v int32) *ListClusterOperationHostTaskResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListClusterOperationHostTaskResponseBody) SetClusterOperationHostTaskList(v *ListClusterOperationHostTaskResponseBodyClusterOperationHostTaskList) *ListClusterOperationHostTaskResponseBody {
	s.ClusterOperationHostTaskList = v
	return s
}

type ListClusterOperationHostTaskResponseBodyClusterOperationHostTaskList struct {
	ClusterOperationHostTask []*ListClusterOperationHostTaskResponseBodyClusterOperationHostTaskListClusterOperationHostTask `json:"ClusterOperationHostTask,omitempty" xml:"ClusterOperationHostTask,omitempty" type:"Repeated"`
}

func (s ListClusterOperationHostTaskResponseBodyClusterOperationHostTaskList) String() string {
	return tea.Prettify(s)
}

func (s ListClusterOperationHostTaskResponseBodyClusterOperationHostTaskList) GoString() string {
	return s.String()
}

func (s *ListClusterOperationHostTaskResponseBodyClusterOperationHostTaskList) SetClusterOperationHostTask(v []*ListClusterOperationHostTaskResponseBodyClusterOperationHostTaskListClusterOperationHostTask) *ListClusterOperationHostTaskResponseBodyClusterOperationHostTaskList {
	s.ClusterOperationHostTask = v
	return s
}

type ListClusterOperationHostTaskResponseBodyClusterOperationHostTaskListClusterOperationHostTask struct {
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Percentage *string `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName   *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s ListClusterOperationHostTaskResponseBodyClusterOperationHostTaskListClusterOperationHostTask) String() string {
	return tea.Prettify(s)
}

func (s ListClusterOperationHostTaskResponseBodyClusterOperationHostTaskListClusterOperationHostTask) GoString() string {
	return s.String()
}

func (s *ListClusterOperationHostTaskResponseBodyClusterOperationHostTaskListClusterOperationHostTask) SetStatus(v string) *ListClusterOperationHostTaskResponseBodyClusterOperationHostTaskListClusterOperationHostTask {
	s.Status = &v
	return s
}

func (s *ListClusterOperationHostTaskResponseBodyClusterOperationHostTaskListClusterOperationHostTask) SetPercentage(v string) *ListClusterOperationHostTaskResponseBodyClusterOperationHostTaskListClusterOperationHostTask {
	s.Percentage = &v
	return s
}

func (s *ListClusterOperationHostTaskResponseBodyClusterOperationHostTaskListClusterOperationHostTask) SetTaskId(v string) *ListClusterOperationHostTaskResponseBodyClusterOperationHostTaskListClusterOperationHostTask {
	s.TaskId = &v
	return s
}

func (s *ListClusterOperationHostTaskResponseBodyClusterOperationHostTaskListClusterOperationHostTask) SetTaskName(v string) *ListClusterOperationHostTaskResponseBodyClusterOperationHostTaskListClusterOperationHostTask {
	s.TaskName = &v
	return s
}

type ListClusterOperationHostTaskResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListClusterOperationHostTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClusterOperationHostTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClusterOperationHostTaskResponse) GoString() string {
	return s.String()
}

func (s *ListClusterOperationHostTaskResponse) SetHeaders(v map[string]*string) *ListClusterOperationHostTaskResponse {
	s.Headers = v
	return s
}

func (s *ListClusterOperationHostTaskResponse) SetBody(v *ListClusterOperationHostTaskResponseBody) *ListClusterOperationHostTaskResponse {
	s.Body = v
	return s
}

type DescribeScalingConfigItemRequest struct {
	ResourceOwnerId     *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ScalingConfigItemId *string `json:"ScalingConfigItemId,omitempty" xml:"ScalingConfigItemId,omitempty"`
	ScalingGroupBizId   *string `json:"ScalingGroupBizId,omitempty" xml:"ScalingGroupBizId,omitempty"`
	ConfigItemType      *string `json:"ConfigItemType,omitempty" xml:"ConfigItemType,omitempty"`
}

func (s DescribeScalingConfigItemRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingConfigItemRequest) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigItemRequest) SetResourceOwnerId(v int64) *DescribeScalingConfigItemRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeScalingConfigItemRequest) SetRegionId(v string) *DescribeScalingConfigItemRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeScalingConfigItemRequest) SetScalingConfigItemId(v string) *DescribeScalingConfigItemRequest {
	s.ScalingConfigItemId = &v
	return s
}

func (s *DescribeScalingConfigItemRequest) SetScalingGroupBizId(v string) *DescribeScalingConfigItemRequest {
	s.ScalingGroupBizId = &v
	return s
}

func (s *DescribeScalingConfigItemRequest) SetConfigItemType(v string) *DescribeScalingConfigItemRequest {
	s.ConfigItemType = &v
	return s
}

type DescribeScalingConfigItemResponseBody struct {
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// data
	ConfigItemType         *string `json:"ConfigItemType,omitempty" xml:"ConfigItemType,omitempty"`
	ScalingGroupBizId      *string `json:"ScalingGroupBizId,omitempty" xml:"ScalingGroupBizId,omitempty"`
	ScalingConfigItemBizId *string `json:"ScalingConfigItemBizId,omitempty" xml:"ScalingConfigItemBizId,omitempty"`
	ConfigItemInformation  *string `json:"ConfigItemInformation,omitempty" xml:"ConfigItemInformation,omitempty"`
}

func (s DescribeScalingConfigItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingConfigItemResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigItemResponseBody) SetRequestId(v string) *DescribeScalingConfigItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScalingConfigItemResponseBody) SetConfigItemType(v string) *DescribeScalingConfigItemResponseBody {
	s.ConfigItemType = &v
	return s
}

func (s *DescribeScalingConfigItemResponseBody) SetScalingGroupBizId(v string) *DescribeScalingConfigItemResponseBody {
	s.ScalingGroupBizId = &v
	return s
}

func (s *DescribeScalingConfigItemResponseBody) SetScalingConfigItemBizId(v string) *DescribeScalingConfigItemResponseBody {
	s.ScalingConfigItemBizId = &v
	return s
}

func (s *DescribeScalingConfigItemResponseBody) SetConfigItemInformation(v string) *DescribeScalingConfigItemResponseBody {
	s.ConfigItemInformation = &v
	return s
}

type DescribeScalingConfigItemResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeScalingConfigItemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScalingConfigItemResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingConfigItemResponse) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigItemResponse) SetHeaders(v map[string]*string) *DescribeScalingConfigItemResponse {
	s.Headers = v
	return s
}

func (s *DescribeScalingConfigItemResponse) SetBody(v *DescribeScalingConfigItemResponseBody) *DescribeScalingConfigItemResponse {
	s.Body = v
	return s
}

type ListClusterHostRequest struct {
	ResourceOwnerId *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClusterId       *string   `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	HostInstanceId  *string   `json:"HostInstanceId,omitempty" xml:"HostInstanceId,omitempty"`
	HostGroupId     *string   `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	HostName        *string   `json:"HostName,omitempty" xml:"HostName,omitempty"`
	PrivateIp       *string   `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	PublicIp        *string   `json:"PublicIp,omitempty" xml:"PublicIp,omitempty"`
	GroupType       *string   `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	ComponentName   *string   `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	PageNumber      *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StatusList      []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
}

func (s ListClusterHostRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClusterHostRequest) GoString() string {
	return s.String()
}

func (s *ListClusterHostRequest) SetResourceOwnerId(v int64) *ListClusterHostRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListClusterHostRequest) SetRegionId(v string) *ListClusterHostRequest {
	s.RegionId = &v
	return s
}

func (s *ListClusterHostRequest) SetClusterId(v string) *ListClusterHostRequest {
	s.ClusterId = &v
	return s
}

func (s *ListClusterHostRequest) SetHostInstanceId(v string) *ListClusterHostRequest {
	s.HostInstanceId = &v
	return s
}

func (s *ListClusterHostRequest) SetHostGroupId(v string) *ListClusterHostRequest {
	s.HostGroupId = &v
	return s
}

func (s *ListClusterHostRequest) SetHostName(v string) *ListClusterHostRequest {
	s.HostName = &v
	return s
}

func (s *ListClusterHostRequest) SetPrivateIp(v string) *ListClusterHostRequest {
	s.PrivateIp = &v
	return s
}

func (s *ListClusterHostRequest) SetPublicIp(v string) *ListClusterHostRequest {
	s.PublicIp = &v
	return s
}

func (s *ListClusterHostRequest) SetGroupType(v string) *ListClusterHostRequest {
	s.GroupType = &v
	return s
}

func (s *ListClusterHostRequest) SetComponentName(v string) *ListClusterHostRequest {
	s.ComponentName = &v
	return s
}

func (s *ListClusterHostRequest) SetPageNumber(v int32) *ListClusterHostRequest {
	s.PageNumber = &v
	return s
}

func (s *ListClusterHostRequest) SetPageSize(v int32) *ListClusterHostRequest {
	s.PageSize = &v
	return s
}

func (s *ListClusterHostRequest) SetStatusList(v []*string) *ListClusterHostRequest {
	s.StatusList = v
	return s
}

type ListClusterHostResponseBody struct {
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total      *int32                               `json:"Total,omitempty" xml:"Total,omitempty"`
	HostList   *ListClusterHostResponseBodyHostList `json:"HostList,omitempty" xml:"HostList,omitempty" type:"Struct"`
}

func (s ListClusterHostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClusterHostResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterHostResponseBody) SetRequestId(v string) *ListClusterHostResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterHostResponseBody) SetPageNumber(v int32) *ListClusterHostResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListClusterHostResponseBody) SetPageSize(v int32) *ListClusterHostResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListClusterHostResponseBody) SetTotal(v int32) *ListClusterHostResponseBody {
	s.Total = &v
	return s
}

func (s *ListClusterHostResponseBody) SetHostList(v *ListClusterHostResponseBodyHostList) *ListClusterHostResponseBody {
	s.HostList = v
	return s
}

type ListClusterHostResponseBodyHostList struct {
	Host []*ListClusterHostResponseBodyHostListHost `json:"Host,omitempty" xml:"Host,omitempty" type:"Repeated"`
}

func (s ListClusterHostResponseBodyHostList) String() string {
	return tea.Prettify(s)
}

func (s ListClusterHostResponseBodyHostList) GoString() string {
	return s.String()
}

func (s *ListClusterHostResponseBodyHostList) SetHost(v []*ListClusterHostResponseBodyHostListHost) *ListClusterHostResponseBodyHostList {
	s.Host = v
	return s
}

type ListClusterHostResponseBodyHostListHost struct {
	SerialNumber   *string                                          `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	Type           *string                                          `json:"Type,omitempty" xml:"Type,omitempty"`
	Status         *string                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	PrivateIp      *string                                          `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	CreateTime     *string                                          `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ChargeType     *string                                          `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	HostName       *string                                          `json:"HostName,omitempty" xml:"HostName,omitempty"`
	HostGroupId    *string                                          `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	InstanceType   *string                                          `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	HostInstanceId *string                                          `json:"HostInstanceId,omitempty" xml:"HostInstanceId,omitempty"`
	SupportIpV6    *bool                                            `json:"SupportIpV6,omitempty" xml:"SupportIpV6,omitempty"`
	Cpu            *int32                                           `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	ExpiredTime    *int64                                           `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	ZoneId         *string                                          `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	PublicIp       *string                                          `json:"PublicIp,omitempty" xml:"PublicIp,omitempty"`
	InstanceStatus *string                                          `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	Memory         *int32                                           `json:"Memory,omitempty" xml:"Memory,omitempty"`
	EmrExpiredTime *string                                          `json:"EmrExpiredTime,omitempty" xml:"EmrExpiredTime,omitempty"`
	Role           *string                                          `json:"Role,omitempty" xml:"Role,omitempty"`
	DiskList       *ListClusterHostResponseBodyHostListHostDiskList `json:"DiskList,omitempty" xml:"DiskList,omitempty" type:"Struct"`
}

func (s ListClusterHostResponseBodyHostListHost) String() string {
	return tea.Prettify(s)
}

func (s ListClusterHostResponseBodyHostListHost) GoString() string {
	return s.String()
}

func (s *ListClusterHostResponseBodyHostListHost) SetSerialNumber(v string) *ListClusterHostResponseBodyHostListHost {
	s.SerialNumber = &v
	return s
}

func (s *ListClusterHostResponseBodyHostListHost) SetType(v string) *ListClusterHostResponseBodyHostListHost {
	s.Type = &v
	return s
}

func (s *ListClusterHostResponseBodyHostListHost) SetStatus(v string) *ListClusterHostResponseBodyHostListHost {
	s.Status = &v
	return s
}

func (s *ListClusterHostResponseBodyHostListHost) SetPrivateIp(v string) *ListClusterHostResponseBodyHostListHost {
	s.PrivateIp = &v
	return s
}

func (s *ListClusterHostResponseBodyHostListHost) SetCreateTime(v string) *ListClusterHostResponseBodyHostListHost {
	s.CreateTime = &v
	return s
}

func (s *ListClusterHostResponseBodyHostListHost) SetChargeType(v string) *ListClusterHostResponseBodyHostListHost {
	s.ChargeType = &v
	return s
}

func (s *ListClusterHostResponseBodyHostListHost) SetHostName(v string) *ListClusterHostResponseBodyHostListHost {
	s.HostName = &v
	return s
}

func (s *ListClusterHostResponseBodyHostListHost) SetHostGroupId(v string) *ListClusterHostResponseBodyHostListHost {
	s.HostGroupId = &v
	return s
}

func (s *ListClusterHostResponseBodyHostListHost) SetInstanceType(v string) *ListClusterHostResponseBodyHostListHost {
	s.InstanceType = &v
	return s
}

func (s *ListClusterHostResponseBodyHostListHost) SetHostInstanceId(v string) *ListClusterHostResponseBodyHostListHost {
	s.HostInstanceId = &v
	return s
}

func (s *ListClusterHostResponseBodyHostListHost) SetSupportIpV6(v bool) *ListClusterHostResponseBodyHostListHost {
	s.SupportIpV6 = &v
	return s
}

func (s *ListClusterHostResponseBodyHostListHost) SetCpu(v int32) *ListClusterHostResponseBodyHostListHost {
	s.Cpu = &v
	return s
}

func (s *ListClusterHostResponseBodyHostListHost) SetExpiredTime(v int64) *ListClusterHostResponseBodyHostListHost {
	s.ExpiredTime = &v
	return s
}

func (s *ListClusterHostResponseBodyHostListHost) SetZoneId(v string) *ListClusterHostResponseBodyHostListHost {
	s.ZoneId = &v
	return s
}

func (s *ListClusterHostResponseBodyHostListHost) SetPublicIp(v string) *ListClusterHostResponseBodyHostListHost {
	s.PublicIp = &v
	return s
}

func (s *ListClusterHostResponseBodyHostListHost) SetInstanceStatus(v string) *ListClusterHostResponseBodyHostListHost {
	s.InstanceStatus = &v
	return s
}

func (s *ListClusterHostResponseBodyHostListHost) SetMemory(v int32) *ListClusterHostResponseBodyHostListHost {
	s.Memory = &v
	return s
}

func (s *ListClusterHostResponseBodyHostListHost) SetEmrExpiredTime(v string) *ListClusterHostResponseBodyHostListHost {
	s.EmrExpiredTime = &v
	return s
}

func (s *ListClusterHostResponseBodyHostListHost) SetRole(v string) *ListClusterHostResponseBodyHostListHost {
	s.Role = &v
	return s
}

func (s *ListClusterHostResponseBodyHostListHost) SetDiskList(v *ListClusterHostResponseBodyHostListHostDiskList) *ListClusterHostResponseBodyHostListHost {
	s.DiskList = v
	return s
}

type ListClusterHostResponseBodyHostListHostDiskList struct {
	Disk []*ListClusterHostResponseBodyHostListHostDiskListDisk `json:"Disk,omitempty" xml:"Disk,omitempty" type:"Repeated"`
}

func (s ListClusterHostResponseBodyHostListHostDiskList) String() string {
	return tea.Prettify(s)
}

func (s ListClusterHostResponseBodyHostListHostDiskList) GoString() string {
	return s.String()
}

func (s *ListClusterHostResponseBodyHostListHostDiskList) SetDisk(v []*ListClusterHostResponseBodyHostListHostDiskListDisk) *ListClusterHostResponseBodyHostListHostDiskList {
	s.Disk = v
	return s
}

type ListClusterHostResponseBodyHostListHostDiskListDisk struct {
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	DiskId   *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	Device   *string `json:"Device,omitempty" xml:"Device,omitempty"`
	DiskSize *int32  `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
}

func (s ListClusterHostResponseBodyHostListHostDiskListDisk) String() string {
	return tea.Prettify(s)
}

func (s ListClusterHostResponseBodyHostListHostDiskListDisk) GoString() string {
	return s.String()
}

func (s *ListClusterHostResponseBodyHostListHostDiskListDisk) SetType(v string) *ListClusterHostResponseBodyHostListHostDiskListDisk {
	s.Type = &v
	return s
}

func (s *ListClusterHostResponseBodyHostListHostDiskListDisk) SetDiskType(v string) *ListClusterHostResponseBodyHostListHostDiskListDisk {
	s.DiskType = &v
	return s
}

func (s *ListClusterHostResponseBodyHostListHostDiskListDisk) SetDiskId(v string) *ListClusterHostResponseBodyHostListHostDiskListDisk {
	s.DiskId = &v
	return s
}

func (s *ListClusterHostResponseBodyHostListHostDiskListDisk) SetDevice(v string) *ListClusterHostResponseBodyHostListHostDiskListDisk {
	s.Device = &v
	return s
}

func (s *ListClusterHostResponseBodyHostListHostDiskListDisk) SetDiskSize(v int32) *ListClusterHostResponseBodyHostListHostDiskListDisk {
	s.DiskSize = &v
	return s
}

type ListClusterHostResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListClusterHostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClusterHostResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClusterHostResponse) GoString() string {
	return s.String()
}

func (s *ListClusterHostResponse) SetHeaders(v map[string]*string) *ListClusterHostResponse {
	s.Headers = v
	return s
}

func (s *ListClusterHostResponse) SetBody(v *ListClusterHostResponseBody) *ListClusterHostResponse {
	s.Body = v
	return s
}

type CreateScalingGroupRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	HostGroupId     *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
}

func (s CreateScalingGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateScalingGroupRequest) SetResourceOwnerId(v int64) *CreateScalingGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateScalingGroupRequest) SetRegionId(v string) *CreateScalingGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateScalingGroupRequest) SetName(v string) *CreateScalingGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateScalingGroupRequest) SetDescription(v string) *CreateScalingGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateScalingGroupRequest) SetHostGroupId(v string) *CreateScalingGroupRequest {
	s.HostGroupId = &v
	return s
}

type CreateScalingGroupResponseBody struct {
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// data
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s CreateScalingGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScalingGroupResponseBody) SetRequestId(v string) *CreateScalingGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScalingGroupResponseBody) SetData(v string) *CreateScalingGroupResponseBody {
	s.Data = &v
	return s
}

type CreateScalingGroupResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateScalingGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateScalingGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateScalingGroupResponse) SetHeaders(v map[string]*string) *CreateScalingGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateScalingGroupResponse) SetBody(v *CreateScalingGroupResponseBody) *CreateScalingGroupResponse {
	s.Body = v
	return s
}

type DescribeClusterServiceRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ServiceName     *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s DescribeClusterServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterServiceRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterServiceRequest) SetResourceOwnerId(v int64) *DescribeClusterServiceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeClusterServiceRequest) SetRegionId(v string) *DescribeClusterServiceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeClusterServiceRequest) SetClusterId(v string) *DescribeClusterServiceRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterServiceRequest) SetServiceName(v string) *DescribeClusterServiceRequest {
	s.ServiceName = &v
	return s
}

type DescribeClusterServiceResponseBody struct {
	RequestId   *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceInfo *DescribeClusterServiceResponseBodyServiceInfo `json:"ServiceInfo,omitempty" xml:"ServiceInfo,omitempty" type:"Struct"`
}

func (s DescribeClusterServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterServiceResponseBody) SetRequestId(v string) *DescribeClusterServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterServiceResponseBody) SetServiceInfo(v *DescribeClusterServiceResponseBodyServiceInfo) *DescribeClusterServiceResponseBody {
	s.ServiceInfo = v
	return s
}

type DescribeClusterServiceResponseBodyServiceInfo struct {
	NeedRestartInfo              *string                                                                    `json:"NeedRestartInfo,omitempty" xml:"NeedRestartInfo,omitempty"`
	ServiceVersion               *string                                                                    `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	ServiceStatus                *string                                                                    `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	ServiceName                  *string                                                                    `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	NeedRestartNum               *int32                                                                     `json:"NeedRestartNum,omitempty" xml:"NeedRestartNum,omitempty"`
	ServiceActionList            *DescribeClusterServiceResponseBodyServiceInfoServiceActionList            `json:"ServiceActionList,omitempty" xml:"ServiceActionList,omitempty" type:"Struct"`
	ClusterServiceSummaryList    *DescribeClusterServiceResponseBodyServiceInfoClusterServiceSummaryList    `json:"ClusterServiceSummaryList,omitempty" xml:"ClusterServiceSummaryList,omitempty" type:"Struct"`
	NeedRestartHostIdList        *DescribeClusterServiceResponseBodyServiceInfoNeedRestartHostIdList        `json:"NeedRestartHostIdList,omitempty" xml:"NeedRestartHostIdList,omitempty" type:"Struct"`
	NeedRestartComponentNameList *DescribeClusterServiceResponseBodyServiceInfoNeedRestartComponentNameList `json:"NeedRestartComponentNameList,omitempty" xml:"NeedRestartComponentNameList,omitempty" type:"Struct"`
}

func (s DescribeClusterServiceResponseBodyServiceInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterServiceResponseBodyServiceInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterServiceResponseBodyServiceInfo) SetNeedRestartInfo(v string) *DescribeClusterServiceResponseBodyServiceInfo {
	s.NeedRestartInfo = &v
	return s
}

func (s *DescribeClusterServiceResponseBodyServiceInfo) SetServiceVersion(v string) *DescribeClusterServiceResponseBodyServiceInfo {
	s.ServiceVersion = &v
	return s
}

func (s *DescribeClusterServiceResponseBodyServiceInfo) SetServiceStatus(v string) *DescribeClusterServiceResponseBodyServiceInfo {
	s.ServiceStatus = &v
	return s
}

func (s *DescribeClusterServiceResponseBodyServiceInfo) SetServiceName(v string) *DescribeClusterServiceResponseBodyServiceInfo {
	s.ServiceName = &v
	return s
}

func (s *DescribeClusterServiceResponseBodyServiceInfo) SetNeedRestartNum(v int32) *DescribeClusterServiceResponseBodyServiceInfo {
	s.NeedRestartNum = &v
	return s
}

func (s *DescribeClusterServiceResponseBodyServiceInfo) SetServiceActionList(v *DescribeClusterServiceResponseBodyServiceInfoServiceActionList) *DescribeClusterServiceResponseBodyServiceInfo {
	s.ServiceActionList = v
	return s
}

func (s *DescribeClusterServiceResponseBodyServiceInfo) SetClusterServiceSummaryList(v *DescribeClusterServiceResponseBodyServiceInfoClusterServiceSummaryList) *DescribeClusterServiceResponseBodyServiceInfo {
	s.ClusterServiceSummaryList = v
	return s
}

func (s *DescribeClusterServiceResponseBodyServiceInfo) SetNeedRestartHostIdList(v *DescribeClusterServiceResponseBodyServiceInfoNeedRestartHostIdList) *DescribeClusterServiceResponseBodyServiceInfo {
	s.NeedRestartHostIdList = v
	return s
}

func (s *DescribeClusterServiceResponseBodyServiceInfo) SetNeedRestartComponentNameList(v *DescribeClusterServiceResponseBodyServiceInfoNeedRestartComponentNameList) *DescribeClusterServiceResponseBodyServiceInfo {
	s.NeedRestartComponentNameList = v
	return s
}

type DescribeClusterServiceResponseBodyServiceInfoServiceActionList struct {
	ServiceAction []*DescribeClusterServiceResponseBodyServiceInfoServiceActionListServiceAction `json:"ServiceAction,omitempty" xml:"ServiceAction,omitempty" type:"Repeated"`
}

func (s DescribeClusterServiceResponseBodyServiceInfoServiceActionList) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterServiceResponseBodyServiceInfoServiceActionList) GoString() string {
	return s.String()
}

func (s *DescribeClusterServiceResponseBodyServiceInfoServiceActionList) SetServiceAction(v []*DescribeClusterServiceResponseBodyServiceInfoServiceActionListServiceAction) *DescribeClusterServiceResponseBodyServiceInfoServiceActionList {
	s.ServiceAction = v
	return s
}

type DescribeClusterServiceResponseBodyServiceInfoServiceActionListServiceAction struct {
	DisplayName   *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	ActionName    *string `json:"ActionName,omitempty" xml:"ActionName,omitempty"`
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	ServiceName   *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	Command       *string `json:"Command,omitempty" xml:"Command,omitempty"`
}

func (s DescribeClusterServiceResponseBodyServiceInfoServiceActionListServiceAction) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterServiceResponseBodyServiceInfoServiceActionListServiceAction) GoString() string {
	return s.String()
}

func (s *DescribeClusterServiceResponseBodyServiceInfoServiceActionListServiceAction) SetDisplayName(v string) *DescribeClusterServiceResponseBodyServiceInfoServiceActionListServiceAction {
	s.DisplayName = &v
	return s
}

func (s *DescribeClusterServiceResponseBodyServiceInfoServiceActionListServiceAction) SetActionName(v string) *DescribeClusterServiceResponseBodyServiceInfoServiceActionListServiceAction {
	s.ActionName = &v
	return s
}

func (s *DescribeClusterServiceResponseBodyServiceInfoServiceActionListServiceAction) SetComponentName(v string) *DescribeClusterServiceResponseBodyServiceInfoServiceActionListServiceAction {
	s.ComponentName = &v
	return s
}

func (s *DescribeClusterServiceResponseBodyServiceInfoServiceActionListServiceAction) SetServiceName(v string) *DescribeClusterServiceResponseBodyServiceInfoServiceActionListServiceAction {
	s.ServiceName = &v
	return s
}

func (s *DescribeClusterServiceResponseBodyServiceInfoServiceActionListServiceAction) SetCommand(v string) *DescribeClusterServiceResponseBodyServiceInfoServiceActionListServiceAction {
	s.Command = &v
	return s
}

type DescribeClusterServiceResponseBodyServiceInfoClusterServiceSummaryList struct {
	ClusterServiceSummary []*DescribeClusterServiceResponseBodyServiceInfoClusterServiceSummaryListClusterServiceSummary `json:"ClusterServiceSummary,omitempty" xml:"ClusterServiceSummary,omitempty" type:"Repeated"`
}

func (s DescribeClusterServiceResponseBodyServiceInfoClusterServiceSummaryList) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterServiceResponseBodyServiceInfoClusterServiceSummaryList) GoString() string {
	return s.String()
}

func (s *DescribeClusterServiceResponseBodyServiceInfoClusterServiceSummaryList) SetClusterServiceSummary(v []*DescribeClusterServiceResponseBodyServiceInfoClusterServiceSummaryListClusterServiceSummary) *DescribeClusterServiceResponseBodyServiceInfoClusterServiceSummaryList {
	s.ClusterServiceSummary = v
	return s
}

type DescribeClusterServiceResponseBodyServiceInfoClusterServiceSummaryListClusterServiceSummary struct {
	Key                 *string `json:"Key,omitempty" xml:"Key,omitempty"`
	DisplayName         *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value               *string `json:"Value,omitempty" xml:"Value,omitempty"`
	DesiredStoppedValue *int32  `json:"DesiredStoppedValue,omitempty" xml:"DesiredStoppedValue,omitempty"`
	AlertInfo           *string `json:"AlertInfo,omitempty" xml:"AlertInfo,omitempty"`
	Category            *string `json:"Category,omitempty" xml:"Category,omitempty"`
}

func (s DescribeClusterServiceResponseBodyServiceInfoClusterServiceSummaryListClusterServiceSummary) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterServiceResponseBodyServiceInfoClusterServiceSummaryListClusterServiceSummary) GoString() string {
	return s.String()
}

func (s *DescribeClusterServiceResponseBodyServiceInfoClusterServiceSummaryListClusterServiceSummary) SetKey(v string) *DescribeClusterServiceResponseBodyServiceInfoClusterServiceSummaryListClusterServiceSummary {
	s.Key = &v
	return s
}

func (s *DescribeClusterServiceResponseBodyServiceInfoClusterServiceSummaryListClusterServiceSummary) SetDisplayName(v string) *DescribeClusterServiceResponseBodyServiceInfoClusterServiceSummaryListClusterServiceSummary {
	s.DisplayName = &v
	return s
}

func (s *DescribeClusterServiceResponseBodyServiceInfoClusterServiceSummaryListClusterServiceSummary) SetStatus(v string) *DescribeClusterServiceResponseBodyServiceInfoClusterServiceSummaryListClusterServiceSummary {
	s.Status = &v
	return s
}

func (s *DescribeClusterServiceResponseBodyServiceInfoClusterServiceSummaryListClusterServiceSummary) SetType(v string) *DescribeClusterServiceResponseBodyServiceInfoClusterServiceSummaryListClusterServiceSummary {
	s.Type = &v
	return s
}

func (s *DescribeClusterServiceResponseBodyServiceInfoClusterServiceSummaryListClusterServiceSummary) SetValue(v string) *DescribeClusterServiceResponseBodyServiceInfoClusterServiceSummaryListClusterServiceSummary {
	s.Value = &v
	return s
}

func (s *DescribeClusterServiceResponseBodyServiceInfoClusterServiceSummaryListClusterServiceSummary) SetDesiredStoppedValue(v int32) *DescribeClusterServiceResponseBodyServiceInfoClusterServiceSummaryListClusterServiceSummary {
	s.DesiredStoppedValue = &v
	return s
}

func (s *DescribeClusterServiceResponseBodyServiceInfoClusterServiceSummaryListClusterServiceSummary) SetAlertInfo(v string) *DescribeClusterServiceResponseBodyServiceInfoClusterServiceSummaryListClusterServiceSummary {
	s.AlertInfo = &v
	return s
}

func (s *DescribeClusterServiceResponseBodyServiceInfoClusterServiceSummaryListClusterServiceSummary) SetCategory(v string) *DescribeClusterServiceResponseBodyServiceInfoClusterServiceSummaryListClusterServiceSummary {
	s.Category = &v
	return s
}

type DescribeClusterServiceResponseBodyServiceInfoNeedRestartHostIdList struct {
	Service []*string `json:"Service,omitempty" xml:"Service,omitempty" type:"Repeated"`
}

func (s DescribeClusterServiceResponseBodyServiceInfoNeedRestartHostIdList) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterServiceResponseBodyServiceInfoNeedRestartHostIdList) GoString() string {
	return s.String()
}

func (s *DescribeClusterServiceResponseBodyServiceInfoNeedRestartHostIdList) SetService(v []*string) *DescribeClusterServiceResponseBodyServiceInfoNeedRestartHostIdList {
	s.Service = v
	return s
}

type DescribeClusterServiceResponseBodyServiceInfoNeedRestartComponentNameList struct {
	Service []*string `json:"Service,omitempty" xml:"Service,omitempty" type:"Repeated"`
}

func (s DescribeClusterServiceResponseBodyServiceInfoNeedRestartComponentNameList) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterServiceResponseBodyServiceInfoNeedRestartComponentNameList) GoString() string {
	return s.String()
}

func (s *DescribeClusterServiceResponseBodyServiceInfoNeedRestartComponentNameList) SetService(v []*string) *DescribeClusterServiceResponseBodyServiceInfoNeedRestartComponentNameList {
	s.Service = v
	return s
}

type DescribeClusterServiceResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClusterServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClusterServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterServiceResponse) SetHeaders(v map[string]*string) *DescribeClusterServiceResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterServiceResponse) SetBody(v *DescribeClusterServiceResponseBody) *DescribeClusterServiceResponse {
	s.Body = v
	return s
}

type ListFlowProjectsRequest struct {
	ProductType     *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProjectId       *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListFlowProjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListFlowProjectsRequest) SetProductType(v string) *ListFlowProjectsRequest {
	s.ProductType = &v
	return s
}

func (s *ListFlowProjectsRequest) SetRegionId(v string) *ListFlowProjectsRequest {
	s.RegionId = &v
	return s
}

func (s *ListFlowProjectsRequest) SetProjectId(v string) *ListFlowProjectsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFlowProjectsRequest) SetName(v string) *ListFlowProjectsRequest {
	s.Name = &v
	return s
}

func (s *ListFlowProjectsRequest) SetPageNumber(v int32) *ListFlowProjectsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFlowProjectsRequest) SetPageSize(v int32) *ListFlowProjectsRequest {
	s.PageSize = &v
	return s
}

func (s *ListFlowProjectsRequest) SetResourceGroupId(v string) *ListFlowProjectsRequest {
	s.ResourceGroupId = &v
	return s
}

type ListFlowProjectsResponseBody struct {
	PageNumber *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total      *int32                                `json:"Total,omitempty" xml:"Total,omitempty"`
	Projects   *ListFlowProjectsResponseBodyProjects `json:"Projects,omitempty" xml:"Projects,omitempty" type:"Struct"`
}

func (s ListFlowProjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowProjectsResponseBody) SetPageNumber(v int32) *ListFlowProjectsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListFlowProjectsResponseBody) SetPageSize(v int32) *ListFlowProjectsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListFlowProjectsResponseBody) SetRequestId(v string) *ListFlowProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowProjectsResponseBody) SetTotal(v int32) *ListFlowProjectsResponseBody {
	s.Total = &v
	return s
}

func (s *ListFlowProjectsResponseBody) SetProjects(v *ListFlowProjectsResponseBodyProjects) *ListFlowProjectsResponseBody {
	s.Projects = v
	return s
}

type ListFlowProjectsResponseBodyProjects struct {
	Project []*ListFlowProjectsResponseBodyProjectsProject `json:"Project,omitempty" xml:"Project,omitempty" type:"Repeated"`
}

func (s ListFlowProjectsResponseBodyProjects) String() string {
	return tea.Prettify(s)
}

func (s ListFlowProjectsResponseBodyProjects) GoString() string {
	return s.String()
}

func (s *ListFlowProjectsResponseBodyProjects) SetProject(v []*ListFlowProjectsResponseBodyProjectsProject) *ListFlowProjectsResponseBodyProjects {
	s.Project = v
	return s
}

type ListFlowProjectsResponseBodyProjectsProject struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	GmtCreate   *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListFlowProjectsResponseBodyProjectsProject) String() string {
	return tea.Prettify(s)
}

func (s ListFlowProjectsResponseBodyProjectsProject) GoString() string {
	return s.String()
}

func (s *ListFlowProjectsResponseBodyProjectsProject) SetDescription(v string) *ListFlowProjectsResponseBodyProjectsProject {
	s.Description = &v
	return s
}

func (s *ListFlowProjectsResponseBodyProjectsProject) SetUserId(v string) *ListFlowProjectsResponseBodyProjectsProject {
	s.UserId = &v
	return s
}

func (s *ListFlowProjectsResponseBodyProjectsProject) SetGmtCreate(v int64) *ListFlowProjectsResponseBodyProjectsProject {
	s.GmtCreate = &v
	return s
}

func (s *ListFlowProjectsResponseBodyProjectsProject) SetGmtModified(v int64) *ListFlowProjectsResponseBodyProjectsProject {
	s.GmtModified = &v
	return s
}

func (s *ListFlowProjectsResponseBodyProjectsProject) SetName(v string) *ListFlowProjectsResponseBodyProjectsProject {
	s.Name = &v
	return s
}

func (s *ListFlowProjectsResponseBodyProjectsProject) SetId(v string) *ListFlowProjectsResponseBodyProjectsProject {
	s.Id = &v
	return s
}

type ListFlowProjectsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFlowProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowProjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowProjectsResponse) GoString() string {
	return s.String()
}

func (s *ListFlowProjectsResponse) SetHeaders(v map[string]*string) *ListFlowProjectsResponse {
	s.Headers = v
	return s
}

func (s *ListFlowProjectsResponse) SetBody(v *ListFlowProjectsResponseBody) *ListFlowProjectsResponse {
	s.Body = v
	return s
}

type CreateMetaTablePreviewTaskRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DatabaseId      *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	TableId         *string `json:"TableId,omitempty" xml:"TableId,omitempty"`
	User            *string `json:"User,omitempty" xml:"User,omitempty"`
	Password        *string `json:"Password,omitempty" xml:"Password,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s CreateMetaTablePreviewTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMetaTablePreviewTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateMetaTablePreviewTaskRequest) SetResourceOwnerId(v int64) *CreateMetaTablePreviewTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateMetaTablePreviewTaskRequest) SetRegionId(v string) *CreateMetaTablePreviewTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateMetaTablePreviewTaskRequest) SetClusterId(v string) *CreateMetaTablePreviewTaskRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateMetaTablePreviewTaskRequest) SetDatabaseId(v string) *CreateMetaTablePreviewTaskRequest {
	s.DatabaseId = &v
	return s
}

func (s *CreateMetaTablePreviewTaskRequest) SetTableId(v string) *CreateMetaTablePreviewTaskRequest {
	s.TableId = &v
	return s
}

func (s *CreateMetaTablePreviewTaskRequest) SetUser(v string) *CreateMetaTablePreviewTaskRequest {
	s.User = &v
	return s
}

func (s *CreateMetaTablePreviewTaskRequest) SetPassword(v string) *CreateMetaTablePreviewTaskRequest {
	s.Password = &v
	return s
}

func (s *CreateMetaTablePreviewTaskRequest) SetResourceGroupId(v string) *CreateMetaTablePreviewTaskRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateMetaTablePreviewTaskRequest) SetClientToken(v string) *CreateMetaTablePreviewTaskRequest {
	s.ClientToken = &v
	return s
}

type CreateMetaTablePreviewTaskResponseBody struct {
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMetaTablePreviewTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMetaTablePreviewTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMetaTablePreviewTaskResponseBody) SetTaskId(v string) *CreateMetaTablePreviewTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateMetaTablePreviewTaskResponseBody) SetRequestId(v string) *CreateMetaTablePreviewTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateMetaTablePreviewTaskResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateMetaTablePreviewTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMetaTablePreviewTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMetaTablePreviewTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateMetaTablePreviewTaskResponse) SetHeaders(v map[string]*string) *CreateMetaTablePreviewTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateMetaTablePreviewTaskResponse) SetBody(v *CreateMetaTablePreviewTaskResponseBody) *CreateMetaTablePreviewTaskResponse {
	s.Body = v
	return s
}

type ListFlowProjectUserRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProjectId  *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListFlowProjectUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowProjectUserRequest) GoString() string {
	return s.String()
}

func (s *ListFlowProjectUserRequest) SetRegionId(v string) *ListFlowProjectUserRequest {
	s.RegionId = &v
	return s
}

func (s *ListFlowProjectUserRequest) SetProjectId(v string) *ListFlowProjectUserRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFlowProjectUserRequest) SetPageNumber(v int32) *ListFlowProjectUserRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFlowProjectUserRequest) SetPageSize(v int32) *ListFlowProjectUserRequest {
	s.PageSize = &v
	return s
}

type ListFlowProjectUserResponseBody struct {
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total      *int32                                `json:"Total,omitempty" xml:"Total,omitempty"`
	Users      *ListFlowProjectUserResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Struct"`
}

func (s ListFlowProjectUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowProjectUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowProjectUserResponseBody) SetRequestId(v string) *ListFlowProjectUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowProjectUserResponseBody) SetPageNumber(v int32) *ListFlowProjectUserResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListFlowProjectUserResponseBody) SetPageSize(v int32) *ListFlowProjectUserResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListFlowProjectUserResponseBody) SetTotal(v int32) *ListFlowProjectUserResponseBody {
	s.Total = &v
	return s
}

func (s *ListFlowProjectUserResponseBody) SetUsers(v *ListFlowProjectUserResponseBodyUsers) *ListFlowProjectUserResponseBody {
	s.Users = v
	return s
}

type ListFlowProjectUserResponseBodyUsers struct {
	User []*ListFlowProjectUserResponseBodyUsersUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s ListFlowProjectUserResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s ListFlowProjectUserResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListFlowProjectUserResponseBodyUsers) SetUser(v []*ListFlowProjectUserResponseBodyUsersUser) *ListFlowProjectUserResponseBodyUsers {
	s.User = v
	return s
}

type ListFlowProjectUserResponseBodyUsersUser struct {
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	GmtCreate     *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified   *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	UserName      *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	AccountUserId *string `json:"AccountUserId,omitempty" xml:"AccountUserId,omitempty"`
	OwnerId       *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s ListFlowProjectUserResponseBodyUsersUser) String() string {
	return tea.Prettify(s)
}

func (s ListFlowProjectUserResponseBodyUsersUser) GoString() string {
	return s.String()
}

func (s *ListFlowProjectUserResponseBodyUsersUser) SetProjectId(v string) *ListFlowProjectUserResponseBodyUsersUser {
	s.ProjectId = &v
	return s
}

func (s *ListFlowProjectUserResponseBodyUsersUser) SetGmtCreate(v int64) *ListFlowProjectUserResponseBodyUsersUser {
	s.GmtCreate = &v
	return s
}

func (s *ListFlowProjectUserResponseBodyUsersUser) SetGmtModified(v int64) *ListFlowProjectUserResponseBodyUsersUser {
	s.GmtModified = &v
	return s
}

func (s *ListFlowProjectUserResponseBodyUsersUser) SetUserName(v string) *ListFlowProjectUserResponseBodyUsersUser {
	s.UserName = &v
	return s
}

func (s *ListFlowProjectUserResponseBodyUsersUser) SetAccountUserId(v string) *ListFlowProjectUserResponseBodyUsersUser {
	s.AccountUserId = &v
	return s
}

func (s *ListFlowProjectUserResponseBodyUsersUser) SetOwnerId(v string) *ListFlowProjectUserResponseBodyUsersUser {
	s.OwnerId = &v
	return s
}

type ListFlowProjectUserResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFlowProjectUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowProjectUserResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowProjectUserResponse) GoString() string {
	return s.String()
}

func (s *ListFlowProjectUserResponse) SetHeaders(v map[string]*string) *ListFlowProjectUserResponse {
	s.Headers = v
	return s
}

func (s *ListFlowProjectUserResponse) SetBody(v *ListFlowProjectUserResponseBody) *ListFlowProjectUserResponse {
	s.Body = v
	return s
}

type DeleteClusterHostGroupRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	HostGroupId     *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	Comment         *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
}

func (s DeleteClusterHostGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterHostGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterHostGroupRequest) SetResourceOwnerId(v int64) *DeleteClusterHostGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteClusterHostGroupRequest) SetRegionId(v string) *DeleteClusterHostGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteClusterHostGroupRequest) SetClusterId(v string) *DeleteClusterHostGroupRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteClusterHostGroupRequest) SetHostGroupId(v string) *DeleteClusterHostGroupRequest {
	s.HostGroupId = &v
	return s
}

func (s *DeleteClusterHostGroupRequest) SetComment(v string) *DeleteClusterHostGroupRequest {
	s.Comment = &v
	return s
}

type DeleteClusterHostGroupResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteClusterHostGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterHostGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClusterHostGroupResponseBody) SetRequestId(v string) *DeleteClusterHostGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteClusterHostGroupResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteClusterHostGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteClusterHostGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterHostGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteClusterHostGroupResponse) SetHeaders(v map[string]*string) *DeleteClusterHostGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteClusterHostGroupResponse) SetBody(v *DeleteClusterHostGroupResponseBody) *DeleteClusterHostGroupResponse {
	s.Body = v
	return s
}

type DescribeLibraryDetailRequest struct {
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	LibraryBizId    *string `json:"LibraryBizId,omitempty" xml:"LibraryBizId,omitempty"`
}

func (s DescribeLibraryDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLibraryDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeLibraryDetailRequest) SetResourceOwnerId(v int64) *DescribeLibraryDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeLibraryDetailRequest) SetRegionId(v string) *DescribeLibraryDetailRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLibraryDetailRequest) SetLibraryBizId(v string) *DescribeLibraryDetailRequest {
	s.LibraryBizId = &v
	return s
}

type DescribeLibraryDetailResponseBody struct {
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CreateTime     *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	SourceType     *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	BizId          *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Scope          *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	LibraryVersion *string `json:"LibraryVersion,omitempty" xml:"LibraryVersion,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Properties     *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	SourceLocation *string `json:"SourceLocation,omitempty" xml:"SourceLocation,omitempty"`
}

func (s DescribeLibraryDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLibraryDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLibraryDetailResponseBody) SetType(v string) *DescribeLibraryDetailResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeLibraryDetailResponseBody) SetRequestId(v string) *DescribeLibraryDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLibraryDetailResponseBody) SetCreateTime(v int64) *DescribeLibraryDetailResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeLibraryDetailResponseBody) SetUserId(v string) *DescribeLibraryDetailResponseBody {
	s.UserId = &v
	return s
}

func (s *DescribeLibraryDetailResponseBody) SetSourceType(v string) *DescribeLibraryDetailResponseBody {
	s.SourceType = &v
	return s
}

func (s *DescribeLibraryDetailResponseBody) SetBizId(v string) *DescribeLibraryDetailResponseBody {
	s.BizId = &v
	return s
}

func (s *DescribeLibraryDetailResponseBody) SetScope(v string) *DescribeLibraryDetailResponseBody {
	s.Scope = &v
	return s
}

func (s *DescribeLibraryDetailResponseBody) SetLibraryVersion(v string) *DescribeLibraryDetailResponseBody {
	s.LibraryVersion = &v
	return s
}

func (s *DescribeLibraryDetailResponseBody) SetName(v string) *DescribeLibraryDetailResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeLibraryDetailResponseBody) SetProperties(v string) *DescribeLibraryDetailResponseBody {
	s.Properties = &v
	return s
}

func (s *DescribeLibraryDetailResponseBody) SetSourceLocation(v string) *DescribeLibraryDetailResponseBody {
	s.SourceLocation = &v
	return s
}

type DescribeLibraryDetailResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeLibraryDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLibraryDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLibraryDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeLibraryDetailResponse) SetHeaders(v map[string]*string) *DescribeLibraryDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeLibraryDetailResponse) SetBody(v *DescribeLibraryDetailResponseBody) *DescribeLibraryDetailResponse {
	s.Body = v
	return s
}

type ListFlowsRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ProjectId  *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	JobId      *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id         *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Periodic   *bool   `json:"Periodic,omitempty" xml:"Periodic,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListFlowsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFlowsRequest) GoString() string {
	return s.String()
}

func (s *ListFlowsRequest) SetRegionId(v string) *ListFlowsRequest {
	s.RegionId = &v
	return s
}

func (s *ListFlowsRequest) SetProjectId(v string) *ListFlowsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFlowsRequest) SetJobId(v string) *ListFlowsRequest {
	s.JobId = &v
	return s
}

func (s *ListFlowsRequest) SetName(v string) *ListFlowsRequest {
	s.Name = &v
	return s
}

func (s *ListFlowsRequest) SetId(v string) *ListFlowsRequest {
	s.Id = &v
	return s
}

func (s *ListFlowsRequest) SetClusterId(v string) *ListFlowsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListFlowsRequest) SetStatus(v string) *ListFlowsRequest {
	s.Status = &v
	return s
}

func (s *ListFlowsRequest) SetPeriodic(v bool) *ListFlowsRequest {
	s.Periodic = &v
	return s
}

func (s *ListFlowsRequest) SetPageNumber(v int32) *ListFlowsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFlowsRequest) SetPageSize(v int32) *ListFlowsRequest {
	s.PageSize = &v
	return s
}

type ListFlowsResponseBody struct {
	PageNumber *int32                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total      *int32                     `json:"Total,omitempty" xml:"Total,omitempty"`
	Flow       *ListFlowsResponseBodyFlow `json:"Flow,omitempty" xml:"Flow,omitempty" type:"Struct"`
}

func (s ListFlowsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFlowsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlowsResponseBody) SetPageNumber(v int32) *ListFlowsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListFlowsResponseBody) SetPageSize(v int32) *ListFlowsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListFlowsResponseBody) SetRequestId(v string) *ListFlowsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlowsResponseBody) SetTotal(v int32) *ListFlowsResponseBody {
	s.Total = &v
	return s
}

func (s *ListFlowsResponseBody) SetFlow(v *ListFlowsResponseBodyFlow) *ListFlowsResponseBody {
	s.Flow = v
	return s
}

type ListFlowsResponseBodyFlow struct {
	Flow []*ListFlowsResponseBodyFlowFlow `json:"Flow,omitempty" xml:"Flow,omitempty" type:"Repeated"`
}

func (s ListFlowsResponseBodyFlow) String() string {
	return tea.Prettify(s)
}

func (s ListFlowsResponseBodyFlow) GoString() string {
	return s.String()
}

func (s *ListFlowsResponseBodyFlow) SetFlow(v []*ListFlowsResponseBodyFlowFlow) *ListFlowsResponseBodyFlow {
	s.Flow = v
	return s
}

type ListFlowsResponseBodyFlowFlow struct {
	Type                    *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Status                  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	AlertUserGroupBizId     *string `json:"AlertUserGroupBizId,omitempty" xml:"AlertUserGroupBizId,omitempty"`
	Periodic                *bool   `json:"Periodic,omitempty" xml:"Periodic,omitempty"`
	ProjectId               *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	HostName                *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	GmtModified             *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	StartSchedule           *int64  `json:"StartSchedule,omitempty" xml:"StartSchedule,omitempty"`
	CreateCluster           *bool   `json:"CreateCluster,omitempty" xml:"CreateCluster,omitempty"`
	EndSchedule             *int64  `json:"EndSchedule,omitempty" xml:"EndSchedule,omitempty"`
	Graph                   *string `json:"Graph,omitempty" xml:"Graph,omitempty"`
	GmtCreate               *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	AlertDingDingGroupBizId *string `json:"AlertDingDingGroupBizId,omitempty" xml:"AlertDingDingGroupBizId,omitempty"`
	CronExpr                *string `json:"CronExpr,omitempty" xml:"CronExpr,omitempty"`
	CategoryId              *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	Name                    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id                      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ClusterId               *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	AlertConf               *string `json:"AlertConf,omitempty" xml:"AlertConf,omitempty"`
}

func (s ListFlowsResponseBodyFlowFlow) String() string {
	return tea.Prettify(s)
}

func (s ListFlowsResponseBodyFlowFlow) GoString() string {
	return s.String()
}

func (s *ListFlowsResponseBodyFlowFlow) SetType(v string) *ListFlowsResponseBodyFlowFlow {
	s.Type = &v
	return s
}

func (s *ListFlowsResponseBodyFlowFlow) SetStatus(v string) *ListFlowsResponseBodyFlowFlow {
	s.Status = &v
	return s
}

func (s *ListFlowsResponseBodyFlowFlow) SetAlertUserGroupBizId(v string) *ListFlowsResponseBodyFlowFlow {
	s.AlertUserGroupBizId = &v
	return s
}

func (s *ListFlowsResponseBodyFlowFlow) SetPeriodic(v bool) *ListFlowsResponseBodyFlowFlow {
	s.Periodic = &v
	return s
}

func (s *ListFlowsResponseBodyFlowFlow) SetProjectId(v string) *ListFlowsResponseBodyFlowFlow {
	s.ProjectId = &v
	return s
}

func (s *ListFlowsResponseBodyFlowFlow) SetHostName(v string) *ListFlowsResponseBodyFlowFlow {
	s.HostName = &v
	return s
}

func (s *ListFlowsResponseBodyFlowFlow) SetGmtModified(v int64) *ListFlowsResponseBodyFlowFlow {
	s.GmtModified = &v
	return s
}

func (s *ListFlowsResponseBodyFlowFlow) SetDescription(v string) *ListFlowsResponseBodyFlowFlow {
	s.Description = &v
	return s
}

func (s *ListFlowsResponseBodyFlowFlow) SetStartSchedule(v int64) *ListFlowsResponseBodyFlowFlow {
	s.StartSchedule = &v
	return s
}

func (s *ListFlowsResponseBodyFlowFlow) SetCreateCluster(v bool) *ListFlowsResponseBodyFlowFlow {
	s.CreateCluster = &v
	return s
}

func (s *ListFlowsResponseBodyFlowFlow) SetEndSchedule(v int64) *ListFlowsResponseBodyFlowFlow {
	s.EndSchedule = &v
	return s
}

func (s *ListFlowsResponseBodyFlowFlow) SetGraph(v string) *ListFlowsResponseBodyFlowFlow {
	s.Graph = &v
	return s
}

func (s *ListFlowsResponseBodyFlowFlow) SetGmtCreate(v int64) *ListFlowsResponseBodyFlowFlow {
	s.GmtCreate = &v
	return s
}

func (s *ListFlowsResponseBodyFlowFlow) SetAlertDingDingGroupBizId(v string) *ListFlowsResponseBodyFlowFlow {
	s.AlertDingDingGroupBizId = &v
	return s
}

func (s *ListFlowsResponseBodyFlowFlow) SetCronExpr(v string) *ListFlowsResponseBodyFlowFlow {
	s.CronExpr = &v
	return s
}

func (s *ListFlowsResponseBodyFlowFlow) SetCategoryId(v string) *ListFlowsResponseBodyFlowFlow {
	s.CategoryId = &v
	return s
}

func (s *ListFlowsResponseBodyFlowFlow) SetName(v string) *ListFlowsResponseBodyFlowFlow {
	s.Name = &v
	return s
}

func (s *ListFlowsResponseBodyFlowFlow) SetId(v string) *ListFlowsResponseBodyFlowFlow {
	s.Id = &v
	return s
}

func (s *ListFlowsResponseBodyFlowFlow) SetClusterId(v string) *ListFlowsResponseBodyFlowFlow {
	s.ClusterId = &v
	return s
}

func (s *ListFlowsResponseBodyFlowFlow) SetAlertConf(v string) *ListFlowsResponseBodyFlowFlow {
	s.AlertConf = &v
	return s
}

type ListFlowsResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFlowsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFlowsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFlowsResponse) GoString() string {
	return s.String()
}

func (s *ListFlowsResponse) SetHeaders(v map[string]*string) *ListFlowsResponse {
	s.Headers = v
	return s
}

func (s *ListFlowsResponse) SetBody(v *ListFlowsResponseBody) *ListFlowsResponse {
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
		"cn-qingdao":     tea.String("ddi.cn-qingdao.aliyuncs.com"),
		"cn-chengdu":     tea.String("ddi.cn-chengdu.aliyuncs.com"),
		"cn-zhangjiakou": tea.String("ddi.cn-zhangjiakou.aliyuncs.com"),
		"cn-huhehaote":   tea.String("ddi.cn-huhehaote.aliyuncs.com"),
		"cn-hongkong":    tea.String("ddi.cn-hongkong.aliyuncs.com"),
		"ap-southeast-2": tea.String("ddi.ap-southeast-2.aliyuncs.com"),
		"ap-southeast-3": tea.String("ddi.ap-southeast-3.aliyuncs.com"),
		"ap-southeast-5": tea.String("ddi.ap-southeast-5.aliyuncs.com"),
		"ap-northeast-1": tea.String("ddi.ap-northeast-1.aliyuncs.com"),
		"eu-west-1":      tea.String("ddi.eu-west-1.aliyuncs.com"),
		"us-east-1":      tea.String("ddi.us-east-1.aliyuncs.com"),
		"eu-central-1":   tea.String("ddi.eu-central-1.aliyuncs.com"),
		"me-east-1":      tea.String("ddi.me-east-1.aliyuncs.com"),
		"ap-south-1":     tea.String("ddi.ap-south-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ddi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateFlowWithOptions(request *CreateFlowRequest, runtime *util.RuntimeOptions) (_result *CreateFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateFlow"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFlow(request *CreateFlowRequest) (_result *CreateFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFlowResponse{}
	_body, _err := client.CreateFlowWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("ListUsers"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ModifyFlowProjectWithOptions(request *ModifyFlowProjectRequest, runtime *util.RuntimeOptions) (_result *ModifyFlowProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyFlowProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyFlowProject"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyFlowProject(request *ModifyFlowProjectRequest) (_result *ModifyFlowProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyFlowProjectResponse{}
	_body, _err := client.ModifyFlowProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryKnoxUserPasswordWithOptions(request *QueryKnoxUserPasswordRequest, runtime *util.RuntimeOptions) (_result *QueryKnoxUserPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryKnoxUserPasswordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryKnoxUserPassword"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryKnoxUserPassword(request *QueryKnoxUserPasswordRequest) (_result *QueryKnoxUserPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryKnoxUserPasswordResponse{}
	_body, _err := client.QueryKnoxUserPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFlowNodeInstanceLauncherLogWithOptions(request *DescribeFlowNodeInstanceLauncherLogRequest, runtime *util.RuntimeOptions) (_result *DescribeFlowNodeInstanceLauncherLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFlowNodeInstanceLauncherLogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFlowNodeInstanceLauncherLog"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlowNodeInstanceLauncherLog(request *DescribeFlowNodeInstanceLauncherLogRequest) (_result *DescribeFlowNodeInstanceLauncherLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlowNodeInstanceLauncherLogResponse{}
	_body, _err := client.DescribeFlowNodeInstanceLauncherLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFlowWithOptions(request *ListFlowRequest, runtime *util.RuntimeOptions) (_result *ListFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFlow"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlow(request *ListFlowRequest) (_result *ListFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowResponse{}
	_body, _err := client.ListFlowWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("UntagResources"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListFlowClusterHostWithOptions(request *ListFlowClusterHostRequest, runtime *util.RuntimeOptions) (_result *ListFlowClusterHostResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListFlowClusterHostResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFlowClusterHost"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlowClusterHost(request *ListFlowClusterHostRequest) (_result *ListFlowClusterHostResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowClusterHostResponse{}
	_body, _err := client.ListFlowClusterHostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClusterOperationWithOptions(request *ListClusterOperationRequest, runtime *util.RuntimeOptions) (_result *ListClusterOperationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListClusterOperationResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListClusterOperation"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClusterOperation(request *ListClusterOperationRequest) (_result *ListClusterOperationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClusterOperationResponse{}
	_body, _err := client.ListClusterOperationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFlowEntitySnapshotWithOptions(request *ListFlowEntitySnapshotRequest, runtime *util.RuntimeOptions) (_result *ListFlowEntitySnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListFlowEntitySnapshotResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFlowEntitySnapshot"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlowEntitySnapshot(request *ListFlowEntitySnapshotRequest) (_result *ListFlowEntitySnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowEntitySnapshotResponse{}
	_body, _err := client.ListFlowEntitySnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteClusterTemplateWithOptions(request *DeleteClusterTemplateRequest, runtime *util.RuntimeOptions) (_result *DeleteClusterTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteClusterTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteClusterTemplate"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteClusterTemplate(request *DeleteClusterTemplateRequest) (_result *DeleteClusterTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteClusterTemplateResponse{}
	_body, _err := client.DeleteClusterTemplateWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("CancelOrder"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CloneFlowJobWithOptions(request *CloneFlowJobRequest, runtime *util.RuntimeOptions) (_result *CloneFlowJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CloneFlowJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CloneFlowJob"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CloneFlowJob(request *CloneFlowJobRequest) (_result *CloneFlowJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CloneFlowJobResponse{}
	_body, _err := client.CloneFlowJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartFlowWithOptions(request *StartFlowRequest, runtime *util.RuntimeOptions) (_result *StartFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartFlow"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartFlow(request *StartFlowRequest) (_result *StartFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartFlowResponse{}
	_body, _err := client.StartFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFlowJobWithOptions(request *CreateFlowJobRequest, runtime *util.RuntimeOptions) (_result *CreateFlowJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateFlowJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateFlowJob"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFlowJob(request *CreateFlowJobRequest) (_result *CreateFlowJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFlowJobResponse{}
	_body, _err := client.CreateFlowJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFlowCategoryWithOptions(request *DeleteFlowCategoryRequest, runtime *util.RuntimeOptions) (_result *DeleteFlowCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteFlowCategoryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteFlowCategory"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFlowCategory(request *DeleteFlowCategoryRequest) (_result *DeleteFlowCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFlowCategoryResponse{}
	_body, _err := client.DeleteFlowCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFlowEditLockWithOptions(request *DeleteFlowEditLockRequest, runtime *util.RuntimeOptions) (_result *DeleteFlowEditLockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteFlowEditLockResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteFlowEditLock"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFlowEditLock(request *DeleteFlowEditLockRequest) (_result *DeleteFlowEditLockResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFlowEditLockResponse{}
	_body, _err := client.DeleteFlowEditLockWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResizeClusterWithOptions(request *ResizeClusterRequest, runtime *util.RuntimeOptions) (_result *ResizeClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ResizeClusterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ResizeCluster"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResizeCluster(request *ResizeClusterRequest) (_result *ResizeClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResizeClusterResponse{}
	_body, _err := client.ResizeClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMetaTablePreviewTaskWithOptions(request *DescribeMetaTablePreviewTaskRequest, runtime *util.RuntimeOptions) (_result *DescribeMetaTablePreviewTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMetaTablePreviewTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMetaTablePreviewTask"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMetaTablePreviewTask(request *DescribeMetaTablePreviewTaskRequest) (_result *DescribeMetaTablePreviewTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMetaTablePreviewTaskResponse{}
	_body, _err := client.DescribeMetaTablePreviewTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClusterServiceConfigHistoryWithOptions(request *ListClusterServiceConfigHistoryRequest, runtime *util.RuntimeOptions) (_result *ListClusterServiceConfigHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListClusterServiceConfigHistoryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListClusterServiceConfigHistory"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClusterServiceConfigHistory(request *ListClusterServiceConfigHistoryRequest) (_result *ListClusterServiceConfigHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClusterServiceConfigHistoryResponse{}
	_body, _err := client.ListClusterServiceConfigHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyScalingConfigItemWithOptions(request *ModifyScalingConfigItemRequest, runtime *util.RuntimeOptions) (_result *ModifyScalingConfigItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyScalingConfigItemResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyScalingConfigItem"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyScalingConfigItem(request *ModifyScalingConfigItemRequest) (_result *ModifyScalingConfigItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyScalingConfigItemResponse{}
	_body, _err := client.ModifyScalingConfigItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFlowClusterAllWithOptions(request *ListFlowClusterAllRequest, runtime *util.RuntimeOptions) (_result *ListFlowClusterAllResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListFlowClusterAllResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFlowClusterAll"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlowClusterAll(request *ListFlowClusterAllRequest) (_result *ListFlowClusterAllResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowClusterAllResponse{}
	_body, _err := client.ListFlowClusterAllWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScalingGroupWithOptions(request *DescribeScalingGroupRequest, runtime *util.RuntimeOptions) (_result *DescribeScalingGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeScalingGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeScalingGroup"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScalingGroup(request *DescribeScalingGroupRequest) (_result *DescribeScalingGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScalingGroupResponse{}
	_body, _err := client.DescribeScalingGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListScalingGroupWithOptions(request *ListScalingGroupRequest, runtime *util.RuntimeOptions) (_result *ListScalingGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListScalingGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListScalingGroup"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListScalingGroup(request *ListScalingGroupRequest) (_result *ListScalingGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListScalingGroupResponse{}
	_body, _err := client.ListScalingGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyFlowCategoryWithOptions(request *ModifyFlowCategoryRequest, runtime *util.RuntimeOptions) (_result *ModifyFlowCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyFlowCategoryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyFlowCategory"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyFlowCategory(request *ModifyFlowCategoryRequest) (_result *ModifyFlowCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyFlowCategoryResponse{}
	_body, _err := client.ModifyFlowCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyClusterServiceConfigWithOptions(request *ModifyClusterServiceConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyClusterServiceConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyClusterServiceConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyClusterServiceConfig"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyClusterServiceConfig(request *ModifyClusterServiceConfigRequest) (_result *ModifyClusterServiceConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyClusterServiceConfigResponse{}
	_body, _err := client.ModifyClusterServiceConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CloneFlowWithOptions(request *CloneFlowRequest, runtime *util.RuntimeOptions) (_result *CloneFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CloneFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CloneFlow"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CloneFlow(request *CloneFlowRequest) (_result *CloneFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CloneFlowResponse{}
	_body, _err := client.CloneFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateClusterTemplateWithOptions(request *CreateClusterTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateClusterTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateClusterTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateClusterTemplate"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateClusterTemplate(request *CreateClusterTemplateRequest) (_result *CreateClusterTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateClusterTemplateResponse{}
	_body, _err := client.CreateClusterTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateLibraryInstallTaskStatusWithOptions(request *UpdateLibraryInstallTaskStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateLibraryInstallTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateLibraryInstallTaskStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateLibraryInstallTaskStatus"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateLibraryInstallTaskStatus(request *UpdateLibraryInstallTaskStatusRequest) (_result *UpdateLibraryInstallTaskStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLibraryInstallTaskStatusResponse{}
	_body, _err := client.UpdateLibraryInstallTaskStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListScalingConfigItemWithOptions(request *ListScalingConfigItemRequest, runtime *util.RuntimeOptions) (_result *ListScalingConfigItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListScalingConfigItemResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListScalingConfigItem"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListScalingConfigItem(request *ListScalingConfigItemRequest) (_result *ListScalingConfigItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListScalingConfigItemResponse{}
	_body, _err := client.ListScalingConfigItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFlowInstanceWithOptions(request *ListFlowInstanceRequest, runtime *util.RuntimeOptions) (_result *ListFlowInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListFlowInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFlowInstance"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlowInstance(request *ListFlowInstanceRequest) (_result *ListFlowInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowInstanceResponse{}
	_body, _err := client.ListFlowInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScalingMetricsWithOptions(request *DescribeScalingMetricsRequest, runtime *util.RuntimeOptions) (_result *DescribeScalingMetricsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeScalingMetricsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeScalingMetrics"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScalingMetrics(request *DescribeScalingMetricsRequest) (_result *DescribeScalingMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScalingMetricsResponse{}
	_body, _err := client.DescribeScalingMetricsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourcesSystemTagsWithOptions(request *UntagResourcesSystemTagsRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesSystemTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UntagResourcesSystemTagsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UntagResourcesSystemTags"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResourcesSystemTags(request *UntagResourcesSystemTagsRequest) (_result *UntagResourcesSystemTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesSystemTagsResponse{}
	_body, _err := client.UntagResourcesSystemTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFlowProjectWithOptions(request *DescribeFlowProjectRequest, runtime *util.RuntimeOptions) (_result *DescribeFlowProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFlowProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFlowProject"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlowProject(request *DescribeFlowProjectRequest) (_result *DescribeFlowProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlowProjectResponse{}
	_body, _err := client.DescribeFlowProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSecurityWhiteListWithOptions(request *DeleteSecurityWhiteListRequest, runtime *util.RuntimeOptions) (_result *DeleteSecurityWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteSecurityWhiteListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteSecurityWhiteList"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSecurityWhiteList(request *DeleteSecurityWhiteListRequest) (_result *DeleteSecurityWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSecurityWhiteListResponse{}
	_body, _err := client.DeleteSecurityWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListScalingActivityWithOptions(request *ListScalingActivityRequest, runtime *util.RuntimeOptions) (_result *ListScalingActivityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListScalingActivityResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListScalingActivity"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListScalingActivity(request *ListScalingActivityRequest) (_result *ListScalingActivityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListScalingActivityResponse{}
	_body, _err := client.ListScalingActivityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagValuesWithOptions(request *ListTagValuesRequest, runtime *util.RuntimeOptions) (_result *ListTagValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListTagValuesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagValues"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagValues(request *ListTagValuesRequest) (_result *ListTagValuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagValuesResponse{}
	_body, _err := client.ListTagValuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClusterInstalledServiceWithOptions(request *ListClusterInstalledServiceRequest, runtime *util.RuntimeOptions) (_result *ListClusterInstalledServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListClusterInstalledServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListClusterInstalledService"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClusterInstalledService(request *ListClusterInstalledServiceRequest) (_result *ListClusterInstalledServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClusterInstalledServiceResponse{}
	_body, _err := client.ListClusterInstalledServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RunClusterServiceActionWithOptions(request *RunClusterServiceActionRequest, runtime *util.RuntimeOptions) (_result *RunClusterServiceActionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RunClusterServiceActionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RunClusterServiceAction"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RunClusterServiceAction(request *RunClusterServiceActionRequest) (_result *RunClusterServiceActionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunClusterServiceActionResponse{}
	_body, _err := client.RunClusterServiceActionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SuspendFlowWithOptions(request *SuspendFlowRequest, runtime *util.RuntimeOptions) (_result *SuspendFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SuspendFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SuspendFlow"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SuspendFlow(request *SuspendFlowRequest) (_result *SuspendFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SuspendFlowResponse{}
	_body, _err := client.SuspendFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFlowProjectWithOptions(request *CreateFlowProjectRequest, runtime *util.RuntimeOptions) (_result *CreateFlowProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateFlowProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateFlowProject"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFlowProject(request *CreateFlowProjectRequest) (_result *CreateFlowProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFlowProjectResponse{}
	_body, _err := client.CreateFlowProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFlowNodeInstanceContainerStatusWithOptions(request *ListFlowNodeInstanceContainerStatusRequest, runtime *util.RuntimeOptions) (_result *ListFlowNodeInstanceContainerStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListFlowNodeInstanceContainerStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFlowNodeInstanceContainerStatus"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlowNodeInstanceContainerStatus(request *ListFlowNodeInstanceContainerStatusRequest) (_result *ListFlowNodeInstanceContainerStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowNodeInstanceContainerStatusResponse{}
	_body, _err := client.ListFlowNodeInstanceContainerStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyClusterTemplateWithOptions(request *ModifyClusterTemplateRequest, runtime *util.RuntimeOptions) (_result *ModifyClusterTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyClusterTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyClusterTemplate"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyClusterTemplate(request *ModifyClusterTemplateRequest) (_result *ModifyClusterTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyClusterTemplateResponse{}
	_body, _err := client.ModifyClusterTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddSecurityWhiteListWithOptions(request *AddSecurityWhiteListRequest, runtime *util.RuntimeOptions) (_result *AddSecurityWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddSecurityWhiteListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddSecurityWhiteList"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddSecurityWhiteList(request *AddSecurityWhiteListRequest) (_result *AddSecurityWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddSecurityWhiteListResponse{}
	_body, _err := client.AddSecurityWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMetaClusterWithOptions(request *ListMetaClusterRequest, runtime *util.RuntimeOptions) (_result *ListMetaClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListMetaClusterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListMetaCluster"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMetaCluster(request *ListMetaClusterRequest) (_result *ListMetaClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMetaClusterResponse{}
	_body, _err := client.ListMetaClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClusterOperationHostWithOptions(request *ListClusterOperationHostRequest, runtime *util.RuntimeOptions) (_result *ListClusterOperationHostResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListClusterOperationHostResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListClusterOperationHost"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClusterOperationHost(request *ListClusterOperationHostRequest) (_result *ListClusterOperationHostResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClusterOperationHostResponse{}
	_body, _err := client.ListClusterOperationHostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClusterTemplatesWithOptions(request *ListClusterTemplatesRequest, runtime *util.RuntimeOptions) (_result *ListClusterTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListClusterTemplatesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListClusterTemplates"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClusterTemplates(request *ListClusterTemplatesRequest) (_result *ListClusterTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClusterTemplatesResponse{}
	_body, _err := client.ListClusterTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClustersWithOptions(request *ListClustersRequest, runtime *util.RuntimeOptions) (_result *ListClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListClustersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListClusters"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClusters(request *ListClustersRequest) (_result *ListClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClustersResponse{}
	_body, _err := client.ListClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesSystemTagsWithOptions(request *TagResourcesSystemTagsRequest, runtime *util.RuntimeOptions) (_result *TagResourcesSystemTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TagResourcesSystemTagsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TagResourcesSystemTags"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResourcesSystemTags(request *TagResourcesSystemTagsRequest) (_result *TagResourcesSystemTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesSystemTagsResponse{}
	_body, _err := client.TagResourcesSystemTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyFlowJobWithOptions(request *ModifyFlowJobRequest, runtime *util.RuntimeOptions) (_result *ModifyFlowJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyFlowJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyFlowJob"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyFlowJob(request *ModifyFlowJobRequest) (_result *ModifyFlowJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyFlowJobResponse{}
	_body, _err := client.ModifyFlowJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFlowWithOptions(request *DeleteFlowRequest, runtime *util.RuntimeOptions) (_result *DeleteFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteFlow"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFlow(request *DeleteFlowRequest) (_result *DeleteFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFlowResponse{}
	_body, _err := client.DeleteFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFlowEditLockWithOptions(request *CreateFlowEditLockRequest, runtime *util.RuntimeOptions) (_result *CreateFlowEditLockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateFlowEditLockResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateFlowEditLock"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFlowEditLock(request *CreateFlowEditLockRequest) (_result *CreateFlowEditLockResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFlowEditLockResponse{}
	_body, _err := client.CreateFlowEditLockWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFlowNodeInstanceWithOptions(request *DescribeFlowNodeInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeFlowNodeInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFlowNodeInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFlowNodeInstance"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlowNodeInstance(request *DescribeFlowNodeInstanceRequest) (_result *DescribeFlowNodeInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlowNodeInstanceResponse{}
	_body, _err := client.DescribeFlowNodeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetachAndReleaseClusterEniWithOptions(request *DetachAndReleaseClusterEniRequest, runtime *util.RuntimeOptions) (_result *DetachAndReleaseClusterEniResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetachAndReleaseClusterEniResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetachAndReleaseClusterEni"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetachAndReleaseClusterEni(request *DetachAndReleaseClusterEniRequest) (_result *DetachAndReleaseClusterEniResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachAndReleaseClusterEniResponse{}
	_body, _err := client.DetachAndReleaseClusterEniWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScalingGroupInstanceWithOptions(request *DescribeScalingGroupInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeScalingGroupInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeScalingGroupInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeScalingGroupInstance"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScalingGroupInstance(request *DescribeScalingGroupInstanceRequest) (_result *DescribeScalingGroupInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScalingGroupInstanceResponse{}
	_body, _err := client.DescribeScalingGroupInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateClusterHostGroupWithOptions(request *CreateClusterHostGroupRequest, runtime *util.RuntimeOptions) (_result *CreateClusterHostGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateClusterHostGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateClusterHostGroup"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateClusterHostGroup(request *CreateClusterHostGroupRequest) (_result *CreateClusterHostGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateClusterHostGroupResponse{}
	_body, _err := client.CreateClusterHostGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterTemplateWithOptions(request *DescribeClusterTemplateRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeClusterTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeClusterTemplate"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterTemplate(request *DescribeClusterTemplateRequest) (_result *DescribeClusterTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterTemplateResponse{}
	_body, _err := client.DescribeClusterTemplateWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("TagResources"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CommitFlowEntitySnapshotWithOptions(request *CommitFlowEntitySnapshotRequest, runtime *util.RuntimeOptions) (_result *CommitFlowEntitySnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CommitFlowEntitySnapshotResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CommitFlowEntitySnapshot"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CommitFlowEntitySnapshot(request *CommitFlowEntitySnapshotRequest) (_result *CommitFlowEntitySnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CommitFlowEntitySnapshotResponse{}
	_body, _err := client.CommitFlowEntitySnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitFlowJobWithOptions(request *SubmitFlowJobRequest, runtime *util.RuntimeOptions) (_result *SubmitFlowJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitFlowJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitFlowJob"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitFlowJob(request *SubmitFlowJobRequest) (_result *SubmitFlowJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitFlowJobResponse{}
	_body, _err := client.SubmitFlowJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFlowJobHistoryWithOptions(request *ListFlowJobHistoryRequest, runtime *util.RuntimeOptions) (_result *ListFlowJobHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListFlowJobHistoryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFlowJobHistory"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlowJobHistory(request *ListFlowJobHistoryRequest) (_result *ListFlowJobHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowJobHistoryResponse{}
	_body, _err := client.ListFlowJobHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagResources"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClusterHostComponentWithOptions(request *ListClusterHostComponentRequest, runtime *util.RuntimeOptions) (_result *ListClusterHostComponentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListClusterHostComponentResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListClusterHostComponent"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClusterHostComponent(request *ListClusterHostComponentRequest) (_result *ListClusterHostComponentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClusterHostComponentResponse{}
	_body, _err := client.ListClusterHostComponentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyScalingGroupWithOptions(request *ModifyScalingGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyScalingGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyScalingGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyScalingGroup"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyScalingGroup(request *ModifyScalingGroupRequest) (_result *ModifyScalingGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyScalingGroupResponse{}
	_body, _err := client.ModifyScalingGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFlowProjectClusterSettingWithOptions(request *DescribeFlowProjectClusterSettingRequest, runtime *util.RuntimeOptions) (_result *DescribeFlowProjectClusterSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFlowProjectClusterSettingResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFlowProjectClusterSetting"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlowProjectClusterSetting(request *DescribeFlowProjectClusterSettingRequest) (_result *DescribeFlowProjectClusterSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlowProjectClusterSettingResponse{}
	_body, _err := client.DescribeFlowProjectClusterSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFlowProjectClusterSettingWithOptions(request *ListFlowProjectClusterSettingRequest, runtime *util.RuntimeOptions) (_result *ListFlowProjectClusterSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListFlowProjectClusterSettingResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFlowProjectClusterSetting"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlowProjectClusterSetting(request *ListFlowProjectClusterSettingRequest) (_result *ListFlowProjectClusterSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowProjectClusterSettingResponse{}
	_body, _err := client.ListFlowProjectClusterSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitFlowWithOptions(request *SubmitFlowRequest, runtime *util.RuntimeOptions) (_result *SubmitFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SubmitFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitFlow"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitFlow(request *SubmitFlowRequest) (_result *SubmitFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitFlowResponse{}
	_body, _err := client.SubmitFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScalingCommonConfigWithOptions(request *DescribeScalingCommonConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeScalingCommonConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeScalingCommonConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeScalingCommonConfig"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScalingCommonConfig(request *DescribeScalingCommonConfigRequest) (_result *DescribeScalingCommonConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScalingCommonConfigResponse{}
	_body, _err := client.DescribeScalingCommonConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResumeFlowWithOptions(request *ResumeFlowRequest, runtime *util.RuntimeOptions) (_result *ResumeFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ResumeFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ResumeFlow"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResumeFlow(request *ResumeFlowRequest) (_result *ResumeFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResumeFlowResponse{}
	_body, _err := client.ResumeFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RestoreFlowEntitySnapshotWithOptions(request *RestoreFlowEntitySnapshotRequest, runtime *util.RuntimeOptions) (_result *RestoreFlowEntitySnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RestoreFlowEntitySnapshotResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RestoreFlowEntitySnapshot"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestoreFlowEntitySnapshot(request *RestoreFlowEntitySnapshotRequest) (_result *RestoreFlowEntitySnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestoreFlowEntitySnapshotResponse{}
	_body, _err := client.RestoreFlowEntitySnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLibraryWithOptions(request *CreateLibraryRequest, runtime *util.RuntimeOptions) (_result *CreateLibraryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateLibraryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateLibrary"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLibrary(request *CreateLibraryRequest) (_result *CreateLibraryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLibraryResponse{}
	_body, _err := client.CreateLibraryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVswitchWithOptions(request *ListVswitchRequest, runtime *util.RuntimeOptions) (_result *ListVswitchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListVswitchResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListVswitch"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVswitch(request *ListVswitchRequest) (_result *ListVswitchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVswitchResponse{}
	_body, _err := client.ListVswitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFlowProjectWithOptions(request *DeleteFlowProjectRequest, runtime *util.RuntimeOptions) (_result *DeleteFlowProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteFlowProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteFlowProject"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFlowProject(request *DeleteFlowProjectRequest) (_result *DeleteFlowProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFlowProjectResponse{}
	_body, _err := client.DeleteFlowProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseClusterWithOptions(request *ReleaseClusterRequest, runtime *util.RuntimeOptions) (_result *ReleaseClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReleaseClusterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReleaseCluster"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseCluster(request *ReleaseClusterRequest) (_result *ReleaseClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseClusterResponse{}
	_body, _err := client.ReleaseClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddScalingConfigItemWithOptions(request *AddScalingConfigItemRequest, runtime *util.RuntimeOptions) (_result *AddScalingConfigItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddScalingConfigItemResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddScalingConfigItem"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddScalingConfigItem(request *AddScalingConfigItemRequest) (_result *AddScalingConfigItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddScalingConfigItemResponse{}
	_body, _err := client.AddScalingConfigItemWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("ResetUserPassword"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListFlowClusterAllHostsWithOptions(request *ListFlowClusterAllHostsRequest, runtime *util.RuntimeOptions) (_result *ListFlowClusterAllHostsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListFlowClusterAllHostsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFlowClusterAllHosts"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlowClusterAllHosts(request *ListFlowClusterAllHostsRequest) (_result *ListFlowClusterAllHostsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowClusterAllHostsResponse{}
	_body, _err := client.ListFlowClusterAllHostsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLibrariesWithOptions(request *DeleteLibrariesRequest, runtime *util.RuntimeOptions) (_result *DeleteLibrariesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteLibrariesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteLibraries"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLibraries(request *DeleteLibrariesRequest) (_result *DeleteLibrariesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLibrariesResponse{}
	_body, _err := client.DeleteLibrariesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFlowCategoryTreeWithOptions(request *DescribeFlowCategoryTreeRequest, runtime *util.RuntimeOptions) (_result *DescribeFlowCategoryTreeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFlowCategoryTreeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFlowCategoryTree"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlowCategoryTree(request *DescribeFlowCategoryTreeRequest) (_result *DescribeFlowCategoryTreeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlowCategoryTreeResponse{}
	_body, _err := client.DescribeFlowCategoryTreeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDatasourceInstancesWithOptions(request *ListDatasourceInstancesRequest, runtime *util.RuntimeOptions) (_result *ListDatasourceInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListDatasourceInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListDatasourceInstances"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDatasourceInstances(request *ListDatasourceInstancesRequest) (_result *ListDatasourceInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDatasourceInstancesResponse{}
	_body, _err := client.ListDatasourceInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFlowNodeSqlResultWithOptions(request *ListFlowNodeSqlResultRequest, runtime *util.RuntimeOptions) (_result *ListFlowNodeSqlResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListFlowNodeSqlResultResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFlowNodeSqlResult"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlowNodeSqlResult(request *ListFlowNodeSqlResultRequest) (_result *ListFlowNodeSqlResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowNodeSqlResultResponse{}
	_body, _err := client.ListFlowNodeSqlResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFlowJobWithOptions(request *DescribeFlowJobRequest, runtime *util.RuntimeOptions) (_result *DescribeFlowJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFlowJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFlowJob"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlowJob(request *DescribeFlowJobRequest) (_result *DescribeFlowJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlowJobResponse{}
	_body, _err := client.DescribeFlowJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLibraryInstallTaskDetailWithOptions(request *DescribeLibraryInstallTaskDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeLibraryInstallTaskDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeLibraryInstallTaskDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeLibraryInstallTaskDetail"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLibraryInstallTaskDetail(request *DescribeLibraryInstallTaskDetailRequest) (_result *DescribeLibraryInstallTaskDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLibraryInstallTaskDetailResponse{}
	_body, _err := client.DescribeLibraryInstallTaskDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyFlowForWebWithOptions(request *ModifyFlowForWebRequest, runtime *util.RuntimeOptions) (_result *ModifyFlowForWebResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyFlowForWebResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyFlowForWeb"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyFlowForWeb(request *ModifyFlowForWebRequest) (_result *ModifyFlowForWebResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyFlowForWebResponse{}
	_body, _err := client.ModifyFlowForWebWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveScalingConfigItemWithOptions(request *RemoveScalingConfigItemRequest, runtime *util.RuntimeOptions) (_result *RemoveScalingConfigItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveScalingConfigItemResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveScalingConfigItem"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveScalingConfigItem(request *RemoveScalingConfigItemRequest) (_result *RemoveScalingConfigItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveScalingConfigItemResponse{}
	_body, _err := client.RemoveScalingConfigItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSecurityWhiteListWithOptions(request *DescribeSecurityWhiteListRequest, runtime *util.RuntimeOptions) (_result *DescribeSecurityWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSecurityWhiteListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSecurityWhiteList"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSecurityWhiteList(request *DescribeSecurityWhiteListRequest) (_result *DescribeSecurityWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSecurityWhiteListResponse{}
	_body, _err := client.DescribeSecurityWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFlowNodeInstanceContainerLogWithOptions(request *DescribeFlowNodeInstanceContainerLogRequest, runtime *util.RuntimeOptions) (_result *DescribeFlowNodeInstanceContainerLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFlowNodeInstanceContainerLogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFlowNodeInstanceContainerLog"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlowNodeInstanceContainerLog(request *DescribeFlowNodeInstanceContainerLogRequest) (_result *DescribeFlowNodeInstanceContainerLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlowNodeInstanceContainerLogResponse{}
	_body, _err := client.DescribeFlowNodeInstanceContainerLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RerunFlowWithOptions(request *RerunFlowRequest, runtime *util.RuntimeOptions) (_result *RerunFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RerunFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RerunFlow"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RerunFlow(request *RerunFlowRequest) (_result *RerunFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RerunFlowResponse{}
	_body, _err := client.RerunFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagKeysWithOptions(request *ListTagKeysRequest, runtime *util.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagKeysResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagKeys"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagKeys(request *ListTagKeysRequest) (_result *ListTagKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagKeysResponse{}
	_body, _err := client.ListTagKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterOperationHostTaskLogWithOptions(request *DescribeClusterOperationHostTaskLogRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterOperationHostTaskLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeClusterOperationHostTaskLogResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeClusterOperationHostTaskLog"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterOperationHostTaskLog(request *DescribeClusterOperationHostTaskLogRequest) (_result *DescribeClusterOperationHostTaskLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterOperationHostTaskLogResponse{}
	_body, _err := client.DescribeClusterOperationHostTaskLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) KillFlowJobWithOptions(request *KillFlowJobRequest, runtime *util.RuntimeOptions) (_result *KillFlowJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &KillFlowJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("KillFlowJob"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) KillFlowJob(request *KillFlowJobRequest) (_result *KillFlowJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &KillFlowJobResponse{}
	_body, _err := client.KillFlowJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UninstallLibrariesWithOptions(request *UninstallLibrariesRequest, runtime *util.RuntimeOptions) (_result *UninstallLibrariesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UninstallLibrariesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UninstallLibraries"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UninstallLibraries(request *UninstallLibrariesRequest) (_result *UninstallLibrariesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UninstallLibrariesResponse{}
	_body, _err := client.UninstallLibrariesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterV2WithOptions(request *DescribeClusterV2Request, runtime *util.RuntimeOptions) (_result *DescribeClusterV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeClusterV2Response{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeClusterV2"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterV2(request *DescribeClusterV2Request) (_result *DescribeClusterV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterV2Response{}
	_body, _err := client.DescribeClusterV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFlowWithOptions(request *DescribeFlowRequest, runtime *util.RuntimeOptions) (_result *DescribeFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFlow"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlow(request *DescribeFlowRequest) (_result *DescribeFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlowResponse{}
	_body, _err := client.DescribeFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFlowClusterWithOptions(request *ListFlowClusterRequest, runtime *util.RuntimeOptions) (_result *ListFlowClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListFlowClusterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFlowCluster"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlowCluster(request *ListFlowClusterRequest) (_result *ListFlowClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowClusterResponse{}
	_body, _err := client.ListFlowClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListLdapUsersWithOptions(request *ListLdapUsersRequest, runtime *util.RuntimeOptions) (_result *ListLdapUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListLdapUsersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListLdapUsers"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListLdapUsers(request *ListLdapUsersRequest) (_result *ListLdapUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListLdapUsersResponse{}
	_body, _err := client.ListLdapUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUserWithOptions(request *DeleteUserRequest, runtime *util.RuntimeOptions) (_result *DeleteUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteUser"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUser(request *DeleteUserRequest) (_result *DeleteUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserResponse{}
	_body, _err := client.DeleteUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFlowProjectClusterSettingWithOptions(request *CreateFlowProjectClusterSettingRequest, runtime *util.RuntimeOptions) (_result *CreateFlowProjectClusterSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateFlowProjectClusterSettingResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateFlowProjectClusterSetting"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFlowProjectClusterSetting(request *CreateFlowProjectClusterSettingRequest) (_result *CreateFlowProjectClusterSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFlowProjectClusterSettingResponse{}
	_body, _err := client.CreateFlowProjectClusterSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFlowInstanceWithOptions(request *DescribeFlowInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeFlowInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeFlowInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeFlowInstance"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFlowInstance(request *DescribeFlowInstanceRequest) (_result *DescribeFlowInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFlowInstanceResponse{}
	_body, _err := client.DescribeFlowInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFlowProjectUserWithOptions(request *CreateFlowProjectUserRequest, runtime *util.RuntimeOptions) (_result *CreateFlowProjectUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateFlowProjectUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateFlowProjectUser"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFlowProjectUser(request *CreateFlowProjectUserRequest) (_result *CreateFlowProjectUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFlowProjectUserResponse{}
	_body, _err := client.CreateFlowProjectUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFlowCategoryWithOptions(request *CreateFlowCategoryRequest, runtime *util.RuntimeOptions) (_result *CreateFlowCategoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateFlowCategoryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateFlowCategory"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFlowCategory(request *CreateFlowCategoryRequest) (_result *CreateFlowCategoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFlowCategoryResponse{}
	_body, _err := client.CreateFlowCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFlowProjectClusterSettingWithOptions(request *DeleteFlowProjectClusterSettingRequest, runtime *util.RuntimeOptions) (_result *DeleteFlowProjectClusterSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteFlowProjectClusterSettingResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteFlowProjectClusterSetting"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFlowProjectClusterSetting(request *DeleteFlowProjectClusterSettingRequest) (_result *DeleteFlowProjectClusterSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFlowProjectClusterSettingResponse{}
	_body, _err := client.DeleteFlowProjectClusterSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListLibrariesWithOptions(request *ListLibrariesRequest, runtime *util.RuntimeOptions) (_result *ListLibrariesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListLibrariesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListLibraries"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListLibraries(request *ListLibrariesRequest) (_result *ListLibrariesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListLibrariesResponse{}
	_body, _err := client.ListLibrariesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RunScalingActionWithOptions(request *RunScalingActionRequest, runtime *util.RuntimeOptions) (_result *RunScalingActionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RunScalingActionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RunScalingAction"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RunScalingAction(request *RunScalingActionRequest) (_result *RunScalingActionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunScalingActionResponse{}
	_body, _err := client.RunScalingActionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InstallLibrariesWithOptions(request *InstallLibrariesRequest, runtime *util.RuntimeOptions) (_result *InstallLibrariesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &InstallLibrariesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("InstallLibraries"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InstallLibraries(request *InstallLibrariesRequest) (_result *InstallLibrariesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InstallLibrariesResponse{}
	_body, _err := client.InstallLibrariesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFlowJobsWithOptions(request *ListFlowJobsRequest, runtime *util.RuntimeOptions) (_result *ListFlowJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListFlowJobsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFlowJobs"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlowJobs(request *ListFlowJobsRequest) (_result *ListFlowJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowJobsResponse{}
	_body, _err := client.ListFlowJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyFlowWithOptions(request *ModifyFlowRequest, runtime *util.RuntimeOptions) (_result *ModifyFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyFlowResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyFlow"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyFlow(request *ModifyFlowRequest) (_result *ModifyFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyFlowResponse{}
	_body, _err := client.ModifyFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRegions"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListLibraryStatusWithOptions(request *ListLibraryStatusRequest, runtime *util.RuntimeOptions) (_result *ListLibraryStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListLibraryStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListLibraryStatus"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListLibraryStatus(request *ListLibraryStatusRequest) (_result *ListLibraryStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListLibraryStatusResponse{}
	_body, _err := client.ListLibraryStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterServiceConfigWithOptions(request *DescribeClusterServiceConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterServiceConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeClusterServiceConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeClusterServiceConfig"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterServiceConfig(request *DescribeClusterServiceConfigRequest) (_result *DescribeClusterServiceConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterServiceConfigResponse{}
	_body, _err := client.DescribeClusterServiceConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyFlowProjectClusterSettingWithOptions(request *ModifyFlowProjectClusterSettingRequest, runtime *util.RuntimeOptions) (_result *ModifyFlowProjectClusterSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyFlowProjectClusterSettingResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyFlowProjectClusterSetting"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyFlowProjectClusterSetting(request *ModifyFlowProjectClusterSettingRequest) (_result *ModifyFlowProjectClusterSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyFlowProjectClusterSettingResponse{}
	_body, _err := client.ModifyFlowProjectClusterSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFlowProjectUserWithOptions(request *DeleteFlowProjectUserRequest, runtime *util.RuntimeOptions) (_result *DeleteFlowProjectUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteFlowProjectUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteFlowProjectUser"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFlowProjectUser(request *DeleteFlowProjectUserRequest) (_result *DeleteFlowProjectUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFlowProjectUserResponse{}
	_body, _err := client.DeleteFlowProjectUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateClusterV2WithOptions(request *CreateClusterV2Request, runtime *util.RuntimeOptions) (_result *CreateClusterV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateClusterV2Response{}
	_body, _err := client.DoRPCRequest(tea.String("CreateClusterV2"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateClusterV2(request *CreateClusterV2Request) (_result *CreateClusterV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateClusterV2Response{}
	_body, _err := client.CreateClusterV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyClusterNameWithOptions(request *ModifyClusterNameRequest, runtime *util.RuntimeOptions) (_result *ModifyClusterNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyClusterNameResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyClusterName"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyClusterName(request *ModifyClusterNameRequest) (_result *ModifyClusterNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyClusterNameResponse{}
	_body, _err := client.ModifyClusterNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClusterOperationHostTaskWithOptions(request *ListClusterOperationHostTaskRequest, runtime *util.RuntimeOptions) (_result *ListClusterOperationHostTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListClusterOperationHostTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListClusterOperationHostTask"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClusterOperationHostTask(request *ListClusterOperationHostTaskRequest) (_result *ListClusterOperationHostTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClusterOperationHostTaskResponse{}
	_body, _err := client.ListClusterOperationHostTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeScalingConfigItemWithOptions(request *DescribeScalingConfigItemRequest, runtime *util.RuntimeOptions) (_result *DescribeScalingConfigItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeScalingConfigItemResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeScalingConfigItem"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScalingConfigItem(request *DescribeScalingConfigItemRequest) (_result *DescribeScalingConfigItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScalingConfigItemResponse{}
	_body, _err := client.DescribeScalingConfigItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClusterHostWithOptions(request *ListClusterHostRequest, runtime *util.RuntimeOptions) (_result *ListClusterHostResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListClusterHostResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListClusterHost"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClusterHost(request *ListClusterHostRequest) (_result *ListClusterHostResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClusterHostResponse{}
	_body, _err := client.ListClusterHostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateScalingGroupWithOptions(request *CreateScalingGroupRequest, runtime *util.RuntimeOptions) (_result *CreateScalingGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateScalingGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateScalingGroup"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateScalingGroup(request *CreateScalingGroupRequest) (_result *CreateScalingGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateScalingGroupResponse{}
	_body, _err := client.CreateScalingGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterServiceWithOptions(request *DescribeClusterServiceRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeClusterServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeClusterService"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterService(request *DescribeClusterServiceRequest) (_result *DescribeClusterServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterServiceResponse{}
	_body, _err := client.DescribeClusterServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFlowProjectsWithOptions(request *ListFlowProjectsRequest, runtime *util.RuntimeOptions) (_result *ListFlowProjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListFlowProjectsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFlowProjects"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlowProjects(request *ListFlowProjectsRequest) (_result *ListFlowProjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowProjectsResponse{}
	_body, _err := client.ListFlowProjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMetaTablePreviewTaskWithOptions(request *CreateMetaTablePreviewTaskRequest, runtime *util.RuntimeOptions) (_result *CreateMetaTablePreviewTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateMetaTablePreviewTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateMetaTablePreviewTask"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMetaTablePreviewTask(request *CreateMetaTablePreviewTaskRequest) (_result *CreateMetaTablePreviewTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMetaTablePreviewTaskResponse{}
	_body, _err := client.CreateMetaTablePreviewTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFlowProjectUserWithOptions(request *ListFlowProjectUserRequest, runtime *util.RuntimeOptions) (_result *ListFlowProjectUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListFlowProjectUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFlowProjectUser"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlowProjectUser(request *ListFlowProjectUserRequest) (_result *ListFlowProjectUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowProjectUserResponse{}
	_body, _err := client.ListFlowProjectUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteClusterHostGroupWithOptions(request *DeleteClusterHostGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteClusterHostGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteClusterHostGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteClusterHostGroup"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteClusterHostGroup(request *DeleteClusterHostGroupRequest) (_result *DeleteClusterHostGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteClusterHostGroupResponse{}
	_body, _err := client.DeleteClusterHostGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLibraryDetailWithOptions(request *DescribeLibraryDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeLibraryDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeLibraryDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeLibraryDetail"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLibraryDetail(request *DescribeLibraryDetailRequest) (_result *DescribeLibraryDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLibraryDetailResponse{}
	_body, _err := client.DescribeLibraryDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFlowsWithOptions(request *ListFlowsRequest, runtime *util.RuntimeOptions) (_result *ListFlowsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListFlowsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFlows"), tea.String("2020-06-17"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFlows(request *ListFlowsRequest) (_result *ListFlowsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFlowsResponse{}
	_body, _err := client.ListFlowsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
