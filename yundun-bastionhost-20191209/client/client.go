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

type AddHostsToGroupRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	HostIds     *string `json:"HostIds,omitempty" xml:"HostIds,omitempty"`
}

func (s AddHostsToGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddHostsToGroupRequest) GoString() string {
	return s.String()
}

func (s *AddHostsToGroupRequest) SetSourceIp(v string) *AddHostsToGroupRequest {
	s.SourceIp = &v
	return s
}

func (s *AddHostsToGroupRequest) SetInstanceId(v string) *AddHostsToGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *AddHostsToGroupRequest) SetRegionId(v string) *AddHostsToGroupRequest {
	s.RegionId = &v
	return s
}

func (s *AddHostsToGroupRequest) SetHostGroupId(v string) *AddHostsToGroupRequest {
	s.HostGroupId = &v
	return s
}

func (s *AddHostsToGroupRequest) SetHostIds(v string) *AddHostsToGroupRequest {
	s.HostIds = &v
	return s
}

type AddHostsToGroupResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*AddHostsToGroupResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s AddHostsToGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddHostsToGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddHostsToGroupResponseBody) SetRequestId(v string) *AddHostsToGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddHostsToGroupResponseBody) SetResults(v []*AddHostsToGroupResponseBodyResults) *AddHostsToGroupResponseBody {
	s.Results = v
	return s
}

type AddHostsToGroupResponseBodyResults struct {
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty"`
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	HostId      *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
}

func (s AddHostsToGroupResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s AddHostsToGroupResponseBodyResults) GoString() string {
	return s.String()
}

func (s *AddHostsToGroupResponseBodyResults) SetCode(v string) *AddHostsToGroupResponseBodyResults {
	s.Code = &v
	return s
}

func (s *AddHostsToGroupResponseBodyResults) SetMessage(v string) *AddHostsToGroupResponseBodyResults {
	s.Message = &v
	return s
}

func (s *AddHostsToGroupResponseBodyResults) SetHostGroupId(v string) *AddHostsToGroupResponseBodyResults {
	s.HostGroupId = &v
	return s
}

func (s *AddHostsToGroupResponseBodyResults) SetHostId(v string) *AddHostsToGroupResponseBodyResults {
	s.HostId = &v
	return s
}

type AddHostsToGroupResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddHostsToGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddHostsToGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddHostsToGroupResponse) GoString() string {
	return s.String()
}

func (s *AddHostsToGroupResponse) SetHeaders(v map[string]*string) *AddHostsToGroupResponse {
	s.Headers = v
	return s
}

func (s *AddHostsToGroupResponse) SetBody(v *AddHostsToGroupResponseBody) *AddHostsToGroupResponse {
	s.Body = v
	return s
}

type AddUsersToGroupRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	UserIds     *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
}

func (s AddUsersToGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddUsersToGroupRequest) GoString() string {
	return s.String()
}

func (s *AddUsersToGroupRequest) SetSourceIp(v string) *AddUsersToGroupRequest {
	s.SourceIp = &v
	return s
}

func (s *AddUsersToGroupRequest) SetInstanceId(v string) *AddUsersToGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *AddUsersToGroupRequest) SetRegionId(v string) *AddUsersToGroupRequest {
	s.RegionId = &v
	return s
}

func (s *AddUsersToGroupRequest) SetUserGroupId(v string) *AddUsersToGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *AddUsersToGroupRequest) SetUserIds(v string) *AddUsersToGroupRequest {
	s.UserIds = &v
	return s
}

type AddUsersToGroupResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*AddUsersToGroupResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s AddUsersToGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddUsersToGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddUsersToGroupResponseBody) SetRequestId(v string) *AddUsersToGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddUsersToGroupResponseBody) SetResults(v []*AddUsersToGroupResponseBodyResults) *AddUsersToGroupResponseBody {
	s.Results = v
	return s
}

type AddUsersToGroupResponseBodyResults struct {
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s AddUsersToGroupResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s AddUsersToGroupResponseBodyResults) GoString() string {
	return s.String()
}

func (s *AddUsersToGroupResponseBodyResults) SetUserGroupId(v string) *AddUsersToGroupResponseBodyResults {
	s.UserGroupId = &v
	return s
}

func (s *AddUsersToGroupResponseBodyResults) SetUserId(v string) *AddUsersToGroupResponseBodyResults {
	s.UserId = &v
	return s
}

func (s *AddUsersToGroupResponseBodyResults) SetCode(v string) *AddUsersToGroupResponseBodyResults {
	s.Code = &v
	return s
}

func (s *AddUsersToGroupResponseBodyResults) SetMessage(v string) *AddUsersToGroupResponseBodyResults {
	s.Message = &v
	return s
}

type AddUsersToGroupResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddUsersToGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddUsersToGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddUsersToGroupResponse) GoString() string {
	return s.String()
}

func (s *AddUsersToGroupResponse) SetHeaders(v map[string]*string) *AddUsersToGroupResponse {
	s.Headers = v
	return s
}

func (s *AddUsersToGroupResponse) SetBody(v *AddUsersToGroupResponseBody) *AddUsersToGroupResponse {
	s.Body = v
	return s
}

type AttachHostAccountsToUserRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Hosts      *string `json:"Hosts,omitempty" xml:"Hosts,omitempty"`
}

func (s AttachHostAccountsToUserRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachHostAccountsToUserRequest) GoString() string {
	return s.String()
}

func (s *AttachHostAccountsToUserRequest) SetSourceIp(v string) *AttachHostAccountsToUserRequest {
	s.SourceIp = &v
	return s
}

func (s *AttachHostAccountsToUserRequest) SetInstanceId(v string) *AttachHostAccountsToUserRequest {
	s.InstanceId = &v
	return s
}

func (s *AttachHostAccountsToUserRequest) SetRegionId(v string) *AttachHostAccountsToUserRequest {
	s.RegionId = &v
	return s
}

func (s *AttachHostAccountsToUserRequest) SetUserId(v string) *AttachHostAccountsToUserRequest {
	s.UserId = &v
	return s
}

func (s *AttachHostAccountsToUserRequest) SetHosts(v string) *AttachHostAccountsToUserRequest {
	s.Hosts = &v
	return s
}

type AttachHostAccountsToUserResponseBody struct {
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*AttachHostAccountsToUserResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s AttachHostAccountsToUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachHostAccountsToUserResponseBody) GoString() string {
	return s.String()
}

func (s *AttachHostAccountsToUserResponseBody) SetRequestId(v string) *AttachHostAccountsToUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachHostAccountsToUserResponseBody) SetResults(v []*AttachHostAccountsToUserResponseBodyResults) *AttachHostAccountsToUserResponseBody {
	s.Results = v
	return s
}

type AttachHostAccountsToUserResponseBodyResults struct {
	HostAccounts []*AttachHostAccountsToUserResponseBodyResultsHostAccounts `json:"HostAccounts,omitempty" xml:"HostAccounts,omitempty" type:"Repeated"`
	UserId       *string                                                    `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Code         *string                                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	HostId       *string                                                    `json:"HostId,omitempty" xml:"HostId,omitempty"`
}

func (s AttachHostAccountsToUserResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s AttachHostAccountsToUserResponseBodyResults) GoString() string {
	return s.String()
}

func (s *AttachHostAccountsToUserResponseBodyResults) SetHostAccounts(v []*AttachHostAccountsToUserResponseBodyResultsHostAccounts) *AttachHostAccountsToUserResponseBodyResults {
	s.HostAccounts = v
	return s
}

func (s *AttachHostAccountsToUserResponseBodyResults) SetUserId(v string) *AttachHostAccountsToUserResponseBodyResults {
	s.UserId = &v
	return s
}

func (s *AttachHostAccountsToUserResponseBodyResults) SetCode(v string) *AttachHostAccountsToUserResponseBodyResults {
	s.Code = &v
	return s
}

func (s *AttachHostAccountsToUserResponseBodyResults) SetMessage(v string) *AttachHostAccountsToUserResponseBodyResults {
	s.Message = &v
	return s
}

func (s *AttachHostAccountsToUserResponseBodyResults) SetHostId(v string) *AttachHostAccountsToUserResponseBodyResults {
	s.HostId = &v
	return s
}

type AttachHostAccountsToUserResponseBodyResultsHostAccounts struct {
	Code          *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty"`
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
}

func (s AttachHostAccountsToUserResponseBodyResultsHostAccounts) String() string {
	return tea.Prettify(s)
}

func (s AttachHostAccountsToUserResponseBodyResultsHostAccounts) GoString() string {
	return s.String()
}

func (s *AttachHostAccountsToUserResponseBodyResultsHostAccounts) SetCode(v string) *AttachHostAccountsToUserResponseBodyResultsHostAccounts {
	s.Code = &v
	return s
}

func (s *AttachHostAccountsToUserResponseBodyResultsHostAccounts) SetMessage(v string) *AttachHostAccountsToUserResponseBodyResultsHostAccounts {
	s.Message = &v
	return s
}

func (s *AttachHostAccountsToUserResponseBodyResultsHostAccounts) SetHostAccountId(v string) *AttachHostAccountsToUserResponseBodyResultsHostAccounts {
	s.HostAccountId = &v
	return s
}

type AttachHostAccountsToUserResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AttachHostAccountsToUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachHostAccountsToUserResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachHostAccountsToUserResponse) GoString() string {
	return s.String()
}

func (s *AttachHostAccountsToUserResponse) SetHeaders(v map[string]*string) *AttachHostAccountsToUserResponse {
	s.Headers = v
	return s
}

func (s *AttachHostAccountsToUserResponse) SetBody(v *AttachHostAccountsToUserResponseBody) *AttachHostAccountsToUserResponse {
	s.Body = v
	return s
}

type AttachHostAccountsToUserGroupRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	Hosts       *string `json:"Hosts,omitempty" xml:"Hosts,omitempty"`
}

func (s AttachHostAccountsToUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachHostAccountsToUserGroupRequest) GoString() string {
	return s.String()
}

func (s *AttachHostAccountsToUserGroupRequest) SetSourceIp(v string) *AttachHostAccountsToUserGroupRequest {
	s.SourceIp = &v
	return s
}

func (s *AttachHostAccountsToUserGroupRequest) SetInstanceId(v string) *AttachHostAccountsToUserGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *AttachHostAccountsToUserGroupRequest) SetRegionId(v string) *AttachHostAccountsToUserGroupRequest {
	s.RegionId = &v
	return s
}

func (s *AttachHostAccountsToUserGroupRequest) SetUserGroupId(v string) *AttachHostAccountsToUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *AttachHostAccountsToUserGroupRequest) SetHosts(v string) *AttachHostAccountsToUserGroupRequest {
	s.Hosts = &v
	return s
}

type AttachHostAccountsToUserGroupResponseBody struct {
	RequestId *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*AttachHostAccountsToUserGroupResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s AttachHostAccountsToUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachHostAccountsToUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AttachHostAccountsToUserGroupResponseBody) SetRequestId(v string) *AttachHostAccountsToUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachHostAccountsToUserGroupResponseBody) SetResults(v []*AttachHostAccountsToUserGroupResponseBodyResults) *AttachHostAccountsToUserGroupResponseBody {
	s.Results = v
	return s
}

type AttachHostAccountsToUserGroupResponseBodyResults struct {
	UserGroupId  *string                                                         `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	HostAccounts []*AttachHostAccountsToUserGroupResponseBodyResultsHostAccounts `json:"HostAccounts,omitempty" xml:"HostAccounts,omitempty" type:"Repeated"`
	Code         *string                                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	HostId       *string                                                         `json:"HostId,omitempty" xml:"HostId,omitempty"`
}

func (s AttachHostAccountsToUserGroupResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s AttachHostAccountsToUserGroupResponseBodyResults) GoString() string {
	return s.String()
}

func (s *AttachHostAccountsToUserGroupResponseBodyResults) SetUserGroupId(v string) *AttachHostAccountsToUserGroupResponseBodyResults {
	s.UserGroupId = &v
	return s
}

func (s *AttachHostAccountsToUserGroupResponseBodyResults) SetHostAccounts(v []*AttachHostAccountsToUserGroupResponseBodyResultsHostAccounts) *AttachHostAccountsToUserGroupResponseBodyResults {
	s.HostAccounts = v
	return s
}

func (s *AttachHostAccountsToUserGroupResponseBodyResults) SetCode(v string) *AttachHostAccountsToUserGroupResponseBodyResults {
	s.Code = &v
	return s
}

func (s *AttachHostAccountsToUserGroupResponseBodyResults) SetMessage(v string) *AttachHostAccountsToUserGroupResponseBodyResults {
	s.Message = &v
	return s
}

func (s *AttachHostAccountsToUserGroupResponseBodyResults) SetHostId(v string) *AttachHostAccountsToUserGroupResponseBodyResults {
	s.HostId = &v
	return s
}

type AttachHostAccountsToUserGroupResponseBodyResultsHostAccounts struct {
	Code          *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty"`
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
}

func (s AttachHostAccountsToUserGroupResponseBodyResultsHostAccounts) String() string {
	return tea.Prettify(s)
}

func (s AttachHostAccountsToUserGroupResponseBodyResultsHostAccounts) GoString() string {
	return s.String()
}

func (s *AttachHostAccountsToUserGroupResponseBodyResultsHostAccounts) SetCode(v string) *AttachHostAccountsToUserGroupResponseBodyResultsHostAccounts {
	s.Code = &v
	return s
}

func (s *AttachHostAccountsToUserGroupResponseBodyResultsHostAccounts) SetMessage(v string) *AttachHostAccountsToUserGroupResponseBodyResultsHostAccounts {
	s.Message = &v
	return s
}

func (s *AttachHostAccountsToUserGroupResponseBodyResultsHostAccounts) SetHostAccountId(v string) *AttachHostAccountsToUserGroupResponseBodyResultsHostAccounts {
	s.HostAccountId = &v
	return s
}

type AttachHostAccountsToUserGroupResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AttachHostAccountsToUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachHostAccountsToUserGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachHostAccountsToUserGroupResponse) GoString() string {
	return s.String()
}

func (s *AttachHostAccountsToUserGroupResponse) SetHeaders(v map[string]*string) *AttachHostAccountsToUserGroupResponse {
	s.Headers = v
	return s
}

func (s *AttachHostAccountsToUserGroupResponse) SetBody(v *AttachHostAccountsToUserGroupResponseBody) *AttachHostAccountsToUserGroupResponse {
	s.Body = v
	return s
}

type AttachHostGroupAccountsToUserRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	HostGroups *string `json:"HostGroups,omitempty" xml:"HostGroups,omitempty"`
}

func (s AttachHostGroupAccountsToUserRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachHostGroupAccountsToUserRequest) GoString() string {
	return s.String()
}

func (s *AttachHostGroupAccountsToUserRequest) SetSourceIp(v string) *AttachHostGroupAccountsToUserRequest {
	s.SourceIp = &v
	return s
}

func (s *AttachHostGroupAccountsToUserRequest) SetInstanceId(v string) *AttachHostGroupAccountsToUserRequest {
	s.InstanceId = &v
	return s
}

func (s *AttachHostGroupAccountsToUserRequest) SetRegionId(v string) *AttachHostGroupAccountsToUserRequest {
	s.RegionId = &v
	return s
}

func (s *AttachHostGroupAccountsToUserRequest) SetUserId(v string) *AttachHostGroupAccountsToUserRequest {
	s.UserId = &v
	return s
}

func (s *AttachHostGroupAccountsToUserRequest) SetHostGroups(v string) *AttachHostGroupAccountsToUserRequest {
	s.HostGroups = &v
	return s
}

type AttachHostGroupAccountsToUserResponseBody struct {
	RequestId *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*AttachHostGroupAccountsToUserResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s AttachHostGroupAccountsToUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachHostGroupAccountsToUserResponseBody) GoString() string {
	return s.String()
}

func (s *AttachHostGroupAccountsToUserResponseBody) SetRequestId(v string) *AttachHostGroupAccountsToUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachHostGroupAccountsToUserResponseBody) SetResults(v []*AttachHostGroupAccountsToUserResponseBodyResults) *AttachHostGroupAccountsToUserResponseBody {
	s.Results = v
	return s
}

type AttachHostGroupAccountsToUserResponseBodyResults struct {
	HostAccountNames []*AttachHostGroupAccountsToUserResponseBodyResultsHostAccountNames `json:"HostAccountNames,omitempty" xml:"HostAccountNames,omitempty" type:"Repeated"`
	UserId           *string                                                             `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Code             *string                                                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Message          *string                                                             `json:"Message,omitempty" xml:"Message,omitempty"`
	HostGroupId      *string                                                             `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
}

func (s AttachHostGroupAccountsToUserResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s AttachHostGroupAccountsToUserResponseBodyResults) GoString() string {
	return s.String()
}

func (s *AttachHostGroupAccountsToUserResponseBodyResults) SetHostAccountNames(v []*AttachHostGroupAccountsToUserResponseBodyResultsHostAccountNames) *AttachHostGroupAccountsToUserResponseBodyResults {
	s.HostAccountNames = v
	return s
}

func (s *AttachHostGroupAccountsToUserResponseBodyResults) SetUserId(v string) *AttachHostGroupAccountsToUserResponseBodyResults {
	s.UserId = &v
	return s
}

func (s *AttachHostGroupAccountsToUserResponseBodyResults) SetCode(v string) *AttachHostGroupAccountsToUserResponseBodyResults {
	s.Code = &v
	return s
}

func (s *AttachHostGroupAccountsToUserResponseBodyResults) SetMessage(v string) *AttachHostGroupAccountsToUserResponseBodyResults {
	s.Message = &v
	return s
}

func (s *AttachHostGroupAccountsToUserResponseBodyResults) SetHostGroupId(v string) *AttachHostGroupAccountsToUserResponseBodyResults {
	s.HostGroupId = &v
	return s
}

type AttachHostGroupAccountsToUserResponseBodyResultsHostAccountNames struct {
	Code            *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message         *string `json:"Message,omitempty" xml:"Message,omitempty"`
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
}

func (s AttachHostGroupAccountsToUserResponseBodyResultsHostAccountNames) String() string {
	return tea.Prettify(s)
}

func (s AttachHostGroupAccountsToUserResponseBodyResultsHostAccountNames) GoString() string {
	return s.String()
}

func (s *AttachHostGroupAccountsToUserResponseBodyResultsHostAccountNames) SetCode(v string) *AttachHostGroupAccountsToUserResponseBodyResultsHostAccountNames {
	s.Code = &v
	return s
}

func (s *AttachHostGroupAccountsToUserResponseBodyResultsHostAccountNames) SetMessage(v string) *AttachHostGroupAccountsToUserResponseBodyResultsHostAccountNames {
	s.Message = &v
	return s
}

func (s *AttachHostGroupAccountsToUserResponseBodyResultsHostAccountNames) SetHostAccountName(v string) *AttachHostGroupAccountsToUserResponseBodyResultsHostAccountNames {
	s.HostAccountName = &v
	return s
}

type AttachHostGroupAccountsToUserResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AttachHostGroupAccountsToUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachHostGroupAccountsToUserResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachHostGroupAccountsToUserResponse) GoString() string {
	return s.String()
}

func (s *AttachHostGroupAccountsToUserResponse) SetHeaders(v map[string]*string) *AttachHostGroupAccountsToUserResponse {
	s.Headers = v
	return s
}

func (s *AttachHostGroupAccountsToUserResponse) SetBody(v *AttachHostGroupAccountsToUserResponseBody) *AttachHostGroupAccountsToUserResponse {
	s.Body = v
	return s
}

type AttachHostGroupAccountsToUserGroupRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	HostGroups  *string `json:"HostGroups,omitempty" xml:"HostGroups,omitempty"`
}

func (s AttachHostGroupAccountsToUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachHostGroupAccountsToUserGroupRequest) GoString() string {
	return s.String()
}

func (s *AttachHostGroupAccountsToUserGroupRequest) SetSourceIp(v string) *AttachHostGroupAccountsToUserGroupRequest {
	s.SourceIp = &v
	return s
}

func (s *AttachHostGroupAccountsToUserGroupRequest) SetInstanceId(v string) *AttachHostGroupAccountsToUserGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *AttachHostGroupAccountsToUserGroupRequest) SetRegionId(v string) *AttachHostGroupAccountsToUserGroupRequest {
	s.RegionId = &v
	return s
}

func (s *AttachHostGroupAccountsToUserGroupRequest) SetUserGroupId(v string) *AttachHostGroupAccountsToUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *AttachHostGroupAccountsToUserGroupRequest) SetHostGroups(v string) *AttachHostGroupAccountsToUserGroupRequest {
	s.HostGroups = &v
	return s
}

type AttachHostGroupAccountsToUserGroupResponseBody struct {
	RequestId *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*AttachHostGroupAccountsToUserGroupResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s AttachHostGroupAccountsToUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachHostGroupAccountsToUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AttachHostGroupAccountsToUserGroupResponseBody) SetRequestId(v string) *AttachHostGroupAccountsToUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachHostGroupAccountsToUserGroupResponseBody) SetResults(v []*AttachHostGroupAccountsToUserGroupResponseBodyResults) *AttachHostGroupAccountsToUserGroupResponseBody {
	s.Results = v
	return s
}

type AttachHostGroupAccountsToUserGroupResponseBodyResults struct {
	HostAccountNames []*AttachHostGroupAccountsToUserGroupResponseBodyResultsHostAccountNames `json:"HostAccountNames,omitempty" xml:"HostAccountNames,omitempty" type:"Repeated"`
	UserGroupId      *string                                                                  `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	Code             *string                                                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message          *string                                                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	HostGroupId      *string                                                                  `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
}

func (s AttachHostGroupAccountsToUserGroupResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s AttachHostGroupAccountsToUserGroupResponseBodyResults) GoString() string {
	return s.String()
}

func (s *AttachHostGroupAccountsToUserGroupResponseBodyResults) SetHostAccountNames(v []*AttachHostGroupAccountsToUserGroupResponseBodyResultsHostAccountNames) *AttachHostGroupAccountsToUserGroupResponseBodyResults {
	s.HostAccountNames = v
	return s
}

func (s *AttachHostGroupAccountsToUserGroupResponseBodyResults) SetUserGroupId(v string) *AttachHostGroupAccountsToUserGroupResponseBodyResults {
	s.UserGroupId = &v
	return s
}

func (s *AttachHostGroupAccountsToUserGroupResponseBodyResults) SetCode(v string) *AttachHostGroupAccountsToUserGroupResponseBodyResults {
	s.Code = &v
	return s
}

func (s *AttachHostGroupAccountsToUserGroupResponseBodyResults) SetMessage(v string) *AttachHostGroupAccountsToUserGroupResponseBodyResults {
	s.Message = &v
	return s
}

func (s *AttachHostGroupAccountsToUserGroupResponseBodyResults) SetHostGroupId(v string) *AttachHostGroupAccountsToUserGroupResponseBodyResults {
	s.HostGroupId = &v
	return s
}

type AttachHostGroupAccountsToUserGroupResponseBodyResultsHostAccountNames struct {
	Code            *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message         *string `json:"Message,omitempty" xml:"Message,omitempty"`
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
}

func (s AttachHostGroupAccountsToUserGroupResponseBodyResultsHostAccountNames) String() string {
	return tea.Prettify(s)
}

func (s AttachHostGroupAccountsToUserGroupResponseBodyResultsHostAccountNames) GoString() string {
	return s.String()
}

func (s *AttachHostGroupAccountsToUserGroupResponseBodyResultsHostAccountNames) SetCode(v string) *AttachHostGroupAccountsToUserGroupResponseBodyResultsHostAccountNames {
	s.Code = &v
	return s
}

func (s *AttachHostGroupAccountsToUserGroupResponseBodyResultsHostAccountNames) SetMessage(v string) *AttachHostGroupAccountsToUserGroupResponseBodyResultsHostAccountNames {
	s.Message = &v
	return s
}

func (s *AttachHostGroupAccountsToUserGroupResponseBodyResultsHostAccountNames) SetHostAccountName(v string) *AttachHostGroupAccountsToUserGroupResponseBodyResultsHostAccountNames {
	s.HostAccountName = &v
	return s
}

type AttachHostGroupAccountsToUserGroupResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AttachHostGroupAccountsToUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachHostGroupAccountsToUserGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachHostGroupAccountsToUserGroupResponse) GoString() string {
	return s.String()
}

func (s *AttachHostGroupAccountsToUserGroupResponse) SetHeaders(v map[string]*string) *AttachHostGroupAccountsToUserGroupResponse {
	s.Headers = v
	return s
}

func (s *AttachHostGroupAccountsToUserGroupResponse) SetBody(v *AttachHostGroupAccountsToUserGroupResponseBody) *AttachHostGroupAccountsToUserGroupResponse {
	s.Body = v
	return s
}

type ConfigInstanceSecurityGroupsRequest struct {
	SourceIp                 *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang                     *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	AuthorizedSecurityGroups []*string `json:"AuthorizedSecurityGroups,omitempty" xml:"AuthorizedSecurityGroups,omitempty" type:"Repeated"`
	InstanceId               *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId                 *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ConfigInstanceSecurityGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigInstanceSecurityGroupsRequest) GoString() string {
	return s.String()
}

func (s *ConfigInstanceSecurityGroupsRequest) SetSourceIp(v string) *ConfigInstanceSecurityGroupsRequest {
	s.SourceIp = &v
	return s
}

func (s *ConfigInstanceSecurityGroupsRequest) SetLang(v string) *ConfigInstanceSecurityGroupsRequest {
	s.Lang = &v
	return s
}

func (s *ConfigInstanceSecurityGroupsRequest) SetAuthorizedSecurityGroups(v []*string) *ConfigInstanceSecurityGroupsRequest {
	s.AuthorizedSecurityGroups = v
	return s
}

func (s *ConfigInstanceSecurityGroupsRequest) SetInstanceId(v string) *ConfigInstanceSecurityGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *ConfigInstanceSecurityGroupsRequest) SetRegionId(v string) *ConfigInstanceSecurityGroupsRequest {
	s.RegionId = &v
	return s
}

type ConfigInstanceSecurityGroupsResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ConfigInstanceSecurityGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigInstanceSecurityGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigInstanceSecurityGroupsResponseBody) SetRequestId(v string) *ConfigInstanceSecurityGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigInstanceSecurityGroupsResponseBody) SetInstanceId(v string) *ConfigInstanceSecurityGroupsResponseBody {
	s.InstanceId = &v
	return s
}

type ConfigInstanceSecurityGroupsResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfigInstanceSecurityGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfigInstanceSecurityGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigInstanceSecurityGroupsResponse) GoString() string {
	return s.String()
}

func (s *ConfigInstanceSecurityGroupsResponse) SetHeaders(v map[string]*string) *ConfigInstanceSecurityGroupsResponse {
	s.Headers = v
	return s
}

func (s *ConfigInstanceSecurityGroupsResponse) SetBody(v *ConfigInstanceSecurityGroupsResponseBody) *ConfigInstanceSecurityGroupsResponse {
	s.Body = v
	return s
}

type ConfigInstanceWhiteListRequest struct {
	SourceIp   *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang       *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	InstanceId *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	WhiteList  []*string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty" type:"Repeated"`
}

func (s ConfigInstanceWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigInstanceWhiteListRequest) GoString() string {
	return s.String()
}

func (s *ConfigInstanceWhiteListRequest) SetSourceIp(v string) *ConfigInstanceWhiteListRequest {
	s.SourceIp = &v
	return s
}

func (s *ConfigInstanceWhiteListRequest) SetLang(v string) *ConfigInstanceWhiteListRequest {
	s.Lang = &v
	return s
}

func (s *ConfigInstanceWhiteListRequest) SetInstanceId(v string) *ConfigInstanceWhiteListRequest {
	s.InstanceId = &v
	return s
}

func (s *ConfigInstanceWhiteListRequest) SetRegionId(v string) *ConfigInstanceWhiteListRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigInstanceWhiteListRequest) SetWhiteList(v []*string) *ConfigInstanceWhiteListRequest {
	s.WhiteList = v
	return s
}

type ConfigInstanceWhiteListResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ConfigInstanceWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigInstanceWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigInstanceWhiteListResponseBody) SetRequestId(v string) *ConfigInstanceWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigInstanceWhiteListResponseBody) SetInstanceId(v string) *ConfigInstanceWhiteListResponseBody {
	s.InstanceId = &v
	return s
}

type ConfigInstanceWhiteListResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfigInstanceWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfigInstanceWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigInstanceWhiteListResponse) GoString() string {
	return s.String()
}

func (s *ConfigInstanceWhiteListResponse) SetHeaders(v map[string]*string) *ConfigInstanceWhiteListResponse {
	s.Headers = v
	return s
}

func (s *ConfigInstanceWhiteListResponse) SetBody(v *ConfigInstanceWhiteListResponseBody) *ConfigInstanceWhiteListResponse {
	s.Body = v
	return s
}

type CreateHostRequest struct {
	SourceIp           *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	HostPrivateAddress *string `json:"HostPrivateAddress,omitempty" xml:"HostPrivateAddress,omitempty"`
	HostPublicAddress  *string `json:"HostPublicAddress,omitempty" xml:"HostPublicAddress,omitempty"`
	ActiveAddressType  *string `json:"ActiveAddressType,omitempty" xml:"ActiveAddressType,omitempty"`
	HostName           *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	Source             *string `json:"Source,omitempty" xml:"Source,omitempty"`
	OSType             *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	Comment            *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	SourceInstanceId   *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
	InstanceRegionId   *string `json:"InstanceRegionId,omitempty" xml:"InstanceRegionId,omitempty"`
}

func (s CreateHostRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateHostRequest) GoString() string {
	return s.String()
}

func (s *CreateHostRequest) SetSourceIp(v string) *CreateHostRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateHostRequest) SetInstanceId(v string) *CreateHostRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateHostRequest) SetRegionId(v string) *CreateHostRequest {
	s.RegionId = &v
	return s
}

func (s *CreateHostRequest) SetHostPrivateAddress(v string) *CreateHostRequest {
	s.HostPrivateAddress = &v
	return s
}

func (s *CreateHostRequest) SetHostPublicAddress(v string) *CreateHostRequest {
	s.HostPublicAddress = &v
	return s
}

func (s *CreateHostRequest) SetActiveAddressType(v string) *CreateHostRequest {
	s.ActiveAddressType = &v
	return s
}

func (s *CreateHostRequest) SetHostName(v string) *CreateHostRequest {
	s.HostName = &v
	return s
}

func (s *CreateHostRequest) SetSource(v string) *CreateHostRequest {
	s.Source = &v
	return s
}

func (s *CreateHostRequest) SetOSType(v string) *CreateHostRequest {
	s.OSType = &v
	return s
}

func (s *CreateHostRequest) SetComment(v string) *CreateHostRequest {
	s.Comment = &v
	return s
}

func (s *CreateHostRequest) SetSourceInstanceId(v string) *CreateHostRequest {
	s.SourceInstanceId = &v
	return s
}

func (s *CreateHostRequest) SetInstanceRegionId(v string) *CreateHostRequest {
	s.InstanceRegionId = &v
	return s
}

type CreateHostResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HostId    *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
}

func (s CreateHostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateHostResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHostResponseBody) SetRequestId(v string) *CreateHostResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHostResponseBody) SetHostId(v string) *CreateHostResponseBody {
	s.HostId = &v
	return s
}

type CreateHostResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateHostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateHostResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateHostResponse) GoString() string {
	return s.String()
}

func (s *CreateHostResponse) SetHeaders(v map[string]*string) *CreateHostResponse {
	s.Headers = v
	return s
}

func (s *CreateHostResponse) SetBody(v *CreateHostResponseBody) *CreateHostResponse {
	s.Body = v
	return s
}

type CreateHostAccountRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	HostId          *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	ProtocolName    *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	Password        *string `json:"Password,omitempty" xml:"Password,omitempty"`
	PrivateKey      *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	PassPhrase      *string `json:"PassPhrase,omitempty" xml:"PassPhrase,omitempty"`
}

func (s CreateHostAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateHostAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateHostAccountRequest) SetSourceIp(v string) *CreateHostAccountRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateHostAccountRequest) SetInstanceId(v string) *CreateHostAccountRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateHostAccountRequest) SetRegionId(v string) *CreateHostAccountRequest {
	s.RegionId = &v
	return s
}

func (s *CreateHostAccountRequest) SetHostId(v string) *CreateHostAccountRequest {
	s.HostId = &v
	return s
}

func (s *CreateHostAccountRequest) SetProtocolName(v string) *CreateHostAccountRequest {
	s.ProtocolName = &v
	return s
}

func (s *CreateHostAccountRequest) SetHostAccountName(v string) *CreateHostAccountRequest {
	s.HostAccountName = &v
	return s
}

func (s *CreateHostAccountRequest) SetPassword(v string) *CreateHostAccountRequest {
	s.Password = &v
	return s
}

func (s *CreateHostAccountRequest) SetPrivateKey(v string) *CreateHostAccountRequest {
	s.PrivateKey = &v
	return s
}

func (s *CreateHostAccountRequest) SetPassPhrase(v string) *CreateHostAccountRequest {
	s.PassPhrase = &v
	return s
}

type CreateHostAccountResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
}

func (s CreateHostAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateHostAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHostAccountResponseBody) SetRequestId(v string) *CreateHostAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHostAccountResponseBody) SetHostAccountId(v string) *CreateHostAccountResponseBody {
	s.HostAccountId = &v
	return s
}

type CreateHostAccountResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateHostAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateHostAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateHostAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateHostAccountResponse) SetHeaders(v map[string]*string) *CreateHostAccountResponse {
	s.Headers = v
	return s
}

func (s *CreateHostAccountResponse) SetBody(v *CreateHostAccountResponseBody) *CreateHostAccountResponse {
	s.Body = v
	return s
}

type CreateHostGroupRequest struct {
	SourceIp      *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	HostGroupName *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
	Comment       *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
}

func (s CreateHostGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateHostGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateHostGroupRequest) SetSourceIp(v string) *CreateHostGroupRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateHostGroupRequest) SetInstanceId(v string) *CreateHostGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateHostGroupRequest) SetRegionId(v string) *CreateHostGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateHostGroupRequest) SetHostGroupName(v string) *CreateHostGroupRequest {
	s.HostGroupName = &v
	return s
}

func (s *CreateHostGroupRequest) SetComment(v string) *CreateHostGroupRequest {
	s.Comment = &v
	return s
}

type CreateHostGroupResponseBody struct {
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateHostGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateHostGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHostGroupResponseBody) SetHostGroupId(v string) *CreateHostGroupResponseBody {
	s.HostGroupId = &v
	return s
}

func (s *CreateHostGroupResponseBody) SetRequestId(v string) *CreateHostGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateHostGroupResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateHostGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateHostGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateHostGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateHostGroupResponse) SetHeaders(v map[string]*string) *CreateHostGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateHostGroupResponse) SetBody(v *CreateHostGroupResponseBody) *CreateHostGroupResponse {
	s.Body = v
	return s
}

type CreateUserRequest struct {
	SourceIp          *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Source            *string `json:"Source,omitempty" xml:"Source,omitempty"`
	UserName          *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	Password          *string `json:"Password,omitempty" xml:"Password,omitempty"`
	DisplayName       *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Comment           *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	Email             *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Mobile            *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	SourceUserId      *string `json:"SourceUserId,omitempty" xml:"SourceUserId,omitempty"`
	MobileCountryCode *string `json:"MobileCountryCode,omitempty" xml:"MobileCountryCode,omitempty"`
}

func (s CreateUserRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserRequest) GoString() string {
	return s.String()
}

func (s *CreateUserRequest) SetSourceIp(v string) *CreateUserRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateUserRequest) SetInstanceId(v string) *CreateUserRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateUserRequest) SetRegionId(v string) *CreateUserRequest {
	s.RegionId = &v
	return s
}

func (s *CreateUserRequest) SetSource(v string) *CreateUserRequest {
	s.Source = &v
	return s
}

func (s *CreateUserRequest) SetUserName(v string) *CreateUserRequest {
	s.UserName = &v
	return s
}

func (s *CreateUserRequest) SetPassword(v string) *CreateUserRequest {
	s.Password = &v
	return s
}

func (s *CreateUserRequest) SetDisplayName(v string) *CreateUserRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateUserRequest) SetComment(v string) *CreateUserRequest {
	s.Comment = &v
	return s
}

func (s *CreateUserRequest) SetEmail(v string) *CreateUserRequest {
	s.Email = &v
	return s
}

func (s *CreateUserRequest) SetMobile(v string) *CreateUserRequest {
	s.Mobile = &v
	return s
}

func (s *CreateUserRequest) SetSourceUserId(v string) *CreateUserRequest {
	s.SourceUserId = &v
	return s
}

func (s *CreateUserRequest) SetMobileCountryCode(v string) *CreateUserRequest {
	s.MobileCountryCode = &v
	return s
}

type CreateUserResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserId    *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserResponseBody) SetRequestId(v string) *CreateUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserResponseBody) SetUserId(v string) *CreateUserResponseBody {
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

type CreateUserGroupRequest struct {
	SourceIp      *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
	Comment       *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
}

func (s CreateUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateUserGroupRequest) SetSourceIp(v string) *CreateUserGroupRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateUserGroupRequest) SetInstanceId(v string) *CreateUserGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateUserGroupRequest) SetRegionId(v string) *CreateUserGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateUserGroupRequest) SetUserGroupName(v string) *CreateUserGroupRequest {
	s.UserGroupName = &v
	return s
}

func (s *CreateUserGroupRequest) SetComment(v string) *CreateUserGroupRequest {
	s.Comment = &v
	return s
}

type CreateUserGroupResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s CreateUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserGroupResponseBody) SetRequestId(v string) *CreateUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserGroupResponseBody) SetUserGroupId(v string) *CreateUserGroupResponseBody {
	s.UserGroupId = &v
	return s
}

type CreateUserGroupResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateUserGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateUserGroupResponse) SetHeaders(v map[string]*string) *CreateUserGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateUserGroupResponse) SetBody(v *CreateUserGroupResponseBody) *CreateUserGroupResponse {
	s.Body = v
	return s
}

type DeleteHostRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	HostId     *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
}

func (s DeleteHostRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteHostRequest) GoString() string {
	return s.String()
}

func (s *DeleteHostRequest) SetSourceIp(v string) *DeleteHostRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteHostRequest) SetInstanceId(v string) *DeleteHostRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteHostRequest) SetRegionId(v string) *DeleteHostRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteHostRequest) SetHostId(v string) *DeleteHostRequest {
	s.HostId = &v
	return s
}

type DeleteHostResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteHostResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHostResponseBody) SetRequestId(v string) *DeleteHostResponseBody {
	s.RequestId = &v
	return s
}

type DeleteHostResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteHostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteHostResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteHostResponse) GoString() string {
	return s.String()
}

func (s *DeleteHostResponse) SetHeaders(v map[string]*string) *DeleteHostResponse {
	s.Headers = v
	return s
}

func (s *DeleteHostResponse) SetBody(v *DeleteHostResponseBody) *DeleteHostResponse {
	s.Body = v
	return s
}

type DeleteHostAccountRequest struct {
	SourceIp      *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
}

func (s DeleteHostAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteHostAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteHostAccountRequest) SetSourceIp(v string) *DeleteHostAccountRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteHostAccountRequest) SetInstanceId(v string) *DeleteHostAccountRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteHostAccountRequest) SetRegionId(v string) *DeleteHostAccountRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteHostAccountRequest) SetHostAccountId(v string) *DeleteHostAccountRequest {
	s.HostAccountId = &v
	return s
}

type DeleteHostAccountResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHostAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteHostAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHostAccountResponseBody) SetRequestId(v string) *DeleteHostAccountResponseBody {
	s.RequestId = &v
	return s
}

type DeleteHostAccountResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteHostAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteHostAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteHostAccountResponse) GoString() string {
	return s.String()
}

func (s *DeleteHostAccountResponse) SetHeaders(v map[string]*string) *DeleteHostAccountResponse {
	s.Headers = v
	return s
}

func (s *DeleteHostAccountResponse) SetBody(v *DeleteHostAccountResponseBody) *DeleteHostAccountResponse {
	s.Body = v
	return s
}

type DeleteHostGroupRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
}

func (s DeleteHostGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteHostGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteHostGroupRequest) SetSourceIp(v string) *DeleteHostGroupRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteHostGroupRequest) SetInstanceId(v string) *DeleteHostGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteHostGroupRequest) SetRegionId(v string) *DeleteHostGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteHostGroupRequest) SetHostGroupId(v string) *DeleteHostGroupRequest {
	s.HostGroupId = &v
	return s
}

type DeleteHostGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHostGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteHostGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHostGroupResponseBody) SetRequestId(v string) *DeleteHostGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteHostGroupResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteHostGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteHostGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteHostGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteHostGroupResponse) SetHeaders(v map[string]*string) *DeleteHostGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteHostGroupResponse) SetBody(v *DeleteHostGroupResponseBody) *DeleteHostGroupResponse {
	s.Body = v
	return s
}

type DeleteUserRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserRequest) SetSourceIp(v string) *DeleteUserRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteUserRequest) SetInstanceId(v string) *DeleteUserRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteUserRequest) SetRegionId(v string) *DeleteUserRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteUserRequest) SetUserId(v string) *DeleteUserRequest {
	s.UserId = &v
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

type DeleteUserGroupRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s DeleteUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupRequest) SetSourceIp(v string) *DeleteUserGroupRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteUserGroupRequest) SetInstanceId(v string) *DeleteUserGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteUserGroupRequest) SetRegionId(v string) *DeleteUserGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteUserGroupRequest) SetUserGroupId(v string) *DeleteUserGroupRequest {
	s.UserGroupId = &v
	return s
}

type DeleteUserGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupResponseBody) SetRequestId(v string) *DeleteUserGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUserGroupResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUserGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupResponse) SetHeaders(v map[string]*string) *DeleteUserGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserGroupResponse) SetBody(v *DeleteUserGroupResponseBody) *DeleteUserGroupResponse {
	s.Body = v
	return s
}

type DescribeInstanceAttributeRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeInstanceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeRequest) SetSourceIp(v string) *DescribeInstanceAttributeRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInstanceAttributeRequest) SetLang(v string) *DescribeInstanceAttributeRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInstanceAttributeRequest) SetRegionId(v string) *DescribeInstanceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceAttributeRequest) SetInstanceId(v string) *DescribeInstanceAttributeRequest {
	s.InstanceId = &v
	return s
}

type DescribeInstanceAttributeResponseBody struct {
	RequestId         *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceAttribute *DescribeInstanceAttributeResponseBodyInstanceAttribute `json:"InstanceAttribute,omitempty" xml:"InstanceAttribute,omitempty" type:"Struct"`
}

func (s DescribeInstanceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBody) SetRequestId(v string) *DescribeInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetInstanceAttribute(v *DescribeInstanceAttributeResponseBodyInstanceAttribute) *DescribeInstanceAttributeResponseBody {
	s.InstanceAttribute = v
	return s
}

type DescribeInstanceAttributeResponseBodyInstanceAttribute struct {
	VpcId                    *string                                                        `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchId                *string                                                        `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	Ports                    []*DescribeInstanceAttributeResponseBodyInstanceAttributePorts `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	AuthorizedSecurityGroups []*string                                                      `json:"AuthorizedSecurityGroups,omitempty" xml:"AuthorizedSecurityGroups,omitempty" type:"Repeated"`
	Description              *string                                                        `json:"Description,omitempty" xml:"Description,omitempty"`
	PrivateExportIps         []*string                                                      `json:"PrivateExportIps,omitempty" xml:"PrivateExportIps,omitempty" type:"Repeated"`
	PrivateWhiteList         []*string                                                      `json:"PrivateWhiteList,omitempty" xml:"PrivateWhiteList,omitempty" type:"Repeated"`
	ExpireTime               *int64                                                         `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	EniInstanceId            *string                                                        `json:"EniInstanceId,omitempty" xml:"EniInstanceId,omitempty"`
	ModifyPasswordModule     *string                                                        `json:"ModifyPasswordModule,omitempty" xml:"ModifyPasswordModule,omitempty"`
	InternetEndpoint         *string                                                        `json:"InternetEndpoint,omitempty" xml:"InternetEndpoint,omitempty"`
	InstanceId               *string                                                        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SecurityGroupIds         []*string                                                      `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	RegionId                 *string                                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	IntranetEndpoint         *string                                                        `json:"IntranetEndpoint,omitempty" xml:"IntranetEndpoint,omitempty"`
	PublicExportIps          []*string                                                      `json:"PublicExportIps,omitempty" xml:"PublicExportIps,omitempty" type:"Repeated"`
	StartTime                *int64                                                         `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	PublicWhiteList          []*string                                                      `json:"PublicWhiteList,omitempty" xml:"PublicWhiteList,omitempty" type:"Repeated"`
	WebTerminalModule        *string                                                        `json:"WebTerminalModule,omitempty" xml:"WebTerminalModule,omitempty"`
	InstanceStatus           *string                                                        `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	LicenseCode              *string                                                        `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
	PublicIps                []*string                                                      `json:"PublicIps,omitempty" xml:"PublicIps,omitempty" type:"Repeated"`
	PublicNetworkAccess      *bool                                                          `json:"PublicNetworkAccess,omitempty" xml:"PublicNetworkAccess,omitempty"`
	Storage                  *int64                                                         `json:"Storage,omitempty" xml:"Storage,omitempty"`
}

func (s DescribeInstanceAttributeResponseBodyInstanceAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBodyInstanceAttribute) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetVpcId(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.VpcId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetVswitchId(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.VswitchId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetPorts(v []*DescribeInstanceAttributeResponseBodyInstanceAttributePorts) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.Ports = v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetAuthorizedSecurityGroups(v []*string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.AuthorizedSecurityGroups = v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetDescription(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.Description = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetPrivateExportIps(v []*string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.PrivateExportIps = v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetPrivateWhiteList(v []*string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.PrivateWhiteList = v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetExpireTime(v int64) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetEniInstanceId(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.EniInstanceId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetModifyPasswordModule(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.ModifyPasswordModule = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetInternetEndpoint(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.InternetEndpoint = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetInstanceId(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetSecurityGroupIds(v []*string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.SecurityGroupIds = v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetRegionId(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetIntranetEndpoint(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.IntranetEndpoint = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetPublicExportIps(v []*string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.PublicExportIps = v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetStartTime(v int64) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.StartTime = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetPublicWhiteList(v []*string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.PublicWhiteList = v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetWebTerminalModule(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.WebTerminalModule = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetInstanceStatus(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetLicenseCode(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.LicenseCode = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetPublicIps(v []*string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.PublicIps = v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetPublicNetworkAccess(v bool) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.PublicNetworkAccess = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetStorage(v int64) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.Storage = &v
	return s
}

type DescribeInstanceAttributeResponseBodyInstanceAttributePorts struct {
	StandardPort *int32 `json:"StandardPort,omitempty" xml:"StandardPort,omitempty"`
	CustomPort   *int32 `json:"CustomPort,omitempty" xml:"CustomPort,omitempty"`
}

func (s DescribeInstanceAttributeResponseBodyInstanceAttributePorts) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBodyInstanceAttributePorts) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttributePorts) SetStandardPort(v int32) *DescribeInstanceAttributeResponseBodyInstanceAttributePorts {
	s.StandardPort = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttributePorts) SetCustomPort(v int32) *DescribeInstanceAttributeResponseBodyInstanceAttributePorts {
	s.CustomPort = &v
	return s
}

type DescribeInstanceAttributeResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponse) SetHeaders(v map[string]*string) *DescribeInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceAttributeResponse) SetBody(v *DescribeInstanceAttributeResponseBody) *DescribeInstanceAttributeResponse {
	s.Body = v
	return s
}

type DescribeInstancesRequest struct {
	SourceIp        *string                        `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string                        `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageNumber      *int32                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId        *string                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceStatus  *string                        `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	ResourceGroupId *string                        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	InstanceId      []*string                      `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	Tag             []*DescribeInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequest) SetSourceIp(v string) *DescribeInstancesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInstancesRequest) SetLang(v string) *DescribeInstancesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageNumber(v int32) *DescribeInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageSize(v int32) *DescribeInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesRequest) SetRegionId(v string) *DescribeInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceStatus(v string) *DescribeInstancesRequest {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeInstancesRequest) SetResourceGroupId(v string) *DescribeInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceId(v []*string) *DescribeInstancesRequest {
	s.InstanceId = v
	return s
}

func (s *DescribeInstancesRequest) SetTag(v []*DescribeInstancesRequestTag) *DescribeInstancesRequest {
	s.Tag = v
	return s
}

type DescribeInstancesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstancesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequestTag) SetKey(v string) *DescribeInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeInstancesRequestTag) SetValue(v string) *DescribeInstancesRequestTag {
	s.Value = &v
	return s
}

type DescribeInstancesResponseBody struct {
	Instances  []*DescribeInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	TotalCount *int64                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBody) SetInstances(v []*DescribeInstancesResponseBodyInstances) *DescribeInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeInstancesResponseBody) SetTotalCount(v int64) *DescribeInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetRequestId(v string) *DescribeInstancesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInstancesResponseBodyInstances struct {
	VpcId               *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchId           *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	ImageVersion        *string `json:"ImageVersion,omitempty" xml:"ImageVersion,omitempty"`
	PlanCode            *string `json:"PlanCode,omitempty" xml:"PlanCode,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ExpireTime          *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Legacy              *bool   `json:"Legacy,omitempty" xml:"Legacy,omitempty"`
	InternetEndpoint    *string `json:"InternetEndpoint,omitempty" xml:"InternetEndpoint,omitempty"`
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	IntranetEndpoint    *string `json:"IntranetEndpoint,omitempty" xml:"IntranetEndpoint,omitempty"`
	StartTime           *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	InstanceStatus      *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	LicenseCode         *string `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
	PublicNetworkAccess *bool   `json:"PublicNetworkAccess,omitempty" xml:"PublicNetworkAccess,omitempty"`
}

func (s DescribeInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstances) SetVpcId(v string) *DescribeInstancesResponseBodyInstances {
	s.VpcId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetVswitchId(v string) *DescribeInstancesResponseBodyInstances {
	s.VswitchId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetImageVersion(v string) *DescribeInstancesResponseBodyInstances {
	s.ImageVersion = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetPlanCode(v string) *DescribeInstancesResponseBodyInstances {
	s.PlanCode = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetDescription(v string) *DescribeInstancesResponseBodyInstances {
	s.Description = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetExpireTime(v int64) *DescribeInstancesResponseBodyInstances {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetLegacy(v bool) *DescribeInstancesResponseBodyInstances {
	s.Legacy = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInternetEndpoint(v string) *DescribeInstancesResponseBodyInstances {
	s.InternetEndpoint = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInstanceId(v string) *DescribeInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetRegionId(v string) *DescribeInstancesResponseBodyInstances {
	s.RegionId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetIntranetEndpoint(v string) *DescribeInstancesResponseBodyInstances {
	s.IntranetEndpoint = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetStartTime(v int64) *DescribeInstancesResponseBodyInstances {
	s.StartTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInstanceStatus(v string) *DescribeInstancesResponseBodyInstances {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetLicenseCode(v string) *DescribeInstancesResponseBodyInstances {
	s.LicenseCode = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetPublicNetworkAccess(v bool) *DescribeInstancesResponseBodyInstances {
	s.PublicNetworkAccess = &v
	return s
}

type DescribeInstancesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponse) SetHeaders(v map[string]*string) *DescribeInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstancesResponse) SetBody(v *DescribeInstancesResponseBody) *DescribeInstancesResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	SourceIp       *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang           *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetSourceIp(v string) *DescribeRegionsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeRegionsRequest) SetLang(v string) *DescribeRegionsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

type DescribeRegionsResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Regions   []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
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

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	LocalName      *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetLocalName(v string) *DescribeRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionId = &v
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

type DetachHostAccountsFromUserRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Hosts      *string `json:"Hosts,omitempty" xml:"Hosts,omitempty"`
}

func (s DetachHostAccountsFromUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachHostAccountsFromUserRequest) GoString() string {
	return s.String()
}

func (s *DetachHostAccountsFromUserRequest) SetSourceIp(v string) *DetachHostAccountsFromUserRequest {
	s.SourceIp = &v
	return s
}

func (s *DetachHostAccountsFromUserRequest) SetInstanceId(v string) *DetachHostAccountsFromUserRequest {
	s.InstanceId = &v
	return s
}

func (s *DetachHostAccountsFromUserRequest) SetRegionId(v string) *DetachHostAccountsFromUserRequest {
	s.RegionId = &v
	return s
}

func (s *DetachHostAccountsFromUserRequest) SetUserId(v string) *DetachHostAccountsFromUserRequest {
	s.UserId = &v
	return s
}

func (s *DetachHostAccountsFromUserRequest) SetHosts(v string) *DetachHostAccountsFromUserRequest {
	s.Hosts = &v
	return s
}

type DetachHostAccountsFromUserResponseBody struct {
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*DetachHostAccountsFromUserResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s DetachHostAccountsFromUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachHostAccountsFromUserResponseBody) GoString() string {
	return s.String()
}

func (s *DetachHostAccountsFromUserResponseBody) SetRequestId(v string) *DetachHostAccountsFromUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachHostAccountsFromUserResponseBody) SetResults(v []*DetachHostAccountsFromUserResponseBodyResults) *DetachHostAccountsFromUserResponseBody {
	s.Results = v
	return s
}

type DetachHostAccountsFromUserResponseBodyResults struct {
	HostAccounts []*DetachHostAccountsFromUserResponseBodyResultsHostAccounts `json:"HostAccounts,omitempty" xml:"HostAccounts,omitempty" type:"Repeated"`
	UserId       *string                                                      `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Code         *string                                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	HostId       *string                                                      `json:"HostId,omitempty" xml:"HostId,omitempty"`
}

func (s DetachHostAccountsFromUserResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s DetachHostAccountsFromUserResponseBodyResults) GoString() string {
	return s.String()
}

func (s *DetachHostAccountsFromUserResponseBodyResults) SetHostAccounts(v []*DetachHostAccountsFromUserResponseBodyResultsHostAccounts) *DetachHostAccountsFromUserResponseBodyResults {
	s.HostAccounts = v
	return s
}

func (s *DetachHostAccountsFromUserResponseBodyResults) SetUserId(v string) *DetachHostAccountsFromUserResponseBodyResults {
	s.UserId = &v
	return s
}

func (s *DetachHostAccountsFromUserResponseBodyResults) SetCode(v string) *DetachHostAccountsFromUserResponseBodyResults {
	s.Code = &v
	return s
}

func (s *DetachHostAccountsFromUserResponseBodyResults) SetMessage(v string) *DetachHostAccountsFromUserResponseBodyResults {
	s.Message = &v
	return s
}

func (s *DetachHostAccountsFromUserResponseBodyResults) SetHostId(v string) *DetachHostAccountsFromUserResponseBodyResults {
	s.HostId = &v
	return s
}

type DetachHostAccountsFromUserResponseBodyResultsHostAccounts struct {
	Code          *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty"`
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
}

func (s DetachHostAccountsFromUserResponseBodyResultsHostAccounts) String() string {
	return tea.Prettify(s)
}

func (s DetachHostAccountsFromUserResponseBodyResultsHostAccounts) GoString() string {
	return s.String()
}

func (s *DetachHostAccountsFromUserResponseBodyResultsHostAccounts) SetCode(v string) *DetachHostAccountsFromUserResponseBodyResultsHostAccounts {
	s.Code = &v
	return s
}

func (s *DetachHostAccountsFromUserResponseBodyResultsHostAccounts) SetMessage(v string) *DetachHostAccountsFromUserResponseBodyResultsHostAccounts {
	s.Message = &v
	return s
}

func (s *DetachHostAccountsFromUserResponseBodyResultsHostAccounts) SetHostAccountId(v string) *DetachHostAccountsFromUserResponseBodyResultsHostAccounts {
	s.HostAccountId = &v
	return s
}

type DetachHostAccountsFromUserResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetachHostAccountsFromUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetachHostAccountsFromUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachHostAccountsFromUserResponse) GoString() string {
	return s.String()
}

func (s *DetachHostAccountsFromUserResponse) SetHeaders(v map[string]*string) *DetachHostAccountsFromUserResponse {
	s.Headers = v
	return s
}

func (s *DetachHostAccountsFromUserResponse) SetBody(v *DetachHostAccountsFromUserResponseBody) *DetachHostAccountsFromUserResponse {
	s.Body = v
	return s
}

type DetachHostAccountsFromUserGroupRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	Hosts       *string `json:"Hosts,omitempty" xml:"Hosts,omitempty"`
}

func (s DetachHostAccountsFromUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachHostAccountsFromUserGroupRequest) GoString() string {
	return s.String()
}

func (s *DetachHostAccountsFromUserGroupRequest) SetSourceIp(v string) *DetachHostAccountsFromUserGroupRequest {
	s.SourceIp = &v
	return s
}

func (s *DetachHostAccountsFromUserGroupRequest) SetInstanceId(v string) *DetachHostAccountsFromUserGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *DetachHostAccountsFromUserGroupRequest) SetRegionId(v string) *DetachHostAccountsFromUserGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DetachHostAccountsFromUserGroupRequest) SetUserGroupId(v string) *DetachHostAccountsFromUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *DetachHostAccountsFromUserGroupRequest) SetHosts(v string) *DetachHostAccountsFromUserGroupRequest {
	s.Hosts = &v
	return s
}

type DetachHostAccountsFromUserGroupResponseBody struct {
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*DetachHostAccountsFromUserGroupResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s DetachHostAccountsFromUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachHostAccountsFromUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DetachHostAccountsFromUserGroupResponseBody) SetRequestId(v string) *DetachHostAccountsFromUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachHostAccountsFromUserGroupResponseBody) SetResults(v []*DetachHostAccountsFromUserGroupResponseBodyResults) *DetachHostAccountsFromUserGroupResponseBody {
	s.Results = v
	return s
}

type DetachHostAccountsFromUserGroupResponseBodyResults struct {
	UserGroupId  *string                                                           `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	HostAccounts []*DetachHostAccountsFromUserGroupResponseBodyResultsHostAccounts `json:"HostAccounts,omitempty" xml:"HostAccounts,omitempty" type:"Repeated"`
	Code         *string                                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	HostId       *string                                                           `json:"HostId,omitempty" xml:"HostId,omitempty"`
}

func (s DetachHostAccountsFromUserGroupResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s DetachHostAccountsFromUserGroupResponseBodyResults) GoString() string {
	return s.String()
}

func (s *DetachHostAccountsFromUserGroupResponseBodyResults) SetUserGroupId(v string) *DetachHostAccountsFromUserGroupResponseBodyResults {
	s.UserGroupId = &v
	return s
}

func (s *DetachHostAccountsFromUserGroupResponseBodyResults) SetHostAccounts(v []*DetachHostAccountsFromUserGroupResponseBodyResultsHostAccounts) *DetachHostAccountsFromUserGroupResponseBodyResults {
	s.HostAccounts = v
	return s
}

func (s *DetachHostAccountsFromUserGroupResponseBodyResults) SetCode(v string) *DetachHostAccountsFromUserGroupResponseBodyResults {
	s.Code = &v
	return s
}

func (s *DetachHostAccountsFromUserGroupResponseBodyResults) SetMessage(v string) *DetachHostAccountsFromUserGroupResponseBodyResults {
	s.Message = &v
	return s
}

func (s *DetachHostAccountsFromUserGroupResponseBodyResults) SetHostId(v string) *DetachHostAccountsFromUserGroupResponseBodyResults {
	s.HostId = &v
	return s
}

type DetachHostAccountsFromUserGroupResponseBodyResultsHostAccounts struct {
	Code          *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty"`
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
}

func (s DetachHostAccountsFromUserGroupResponseBodyResultsHostAccounts) String() string {
	return tea.Prettify(s)
}

func (s DetachHostAccountsFromUserGroupResponseBodyResultsHostAccounts) GoString() string {
	return s.String()
}

func (s *DetachHostAccountsFromUserGroupResponseBodyResultsHostAccounts) SetCode(v string) *DetachHostAccountsFromUserGroupResponseBodyResultsHostAccounts {
	s.Code = &v
	return s
}

func (s *DetachHostAccountsFromUserGroupResponseBodyResultsHostAccounts) SetMessage(v string) *DetachHostAccountsFromUserGroupResponseBodyResultsHostAccounts {
	s.Message = &v
	return s
}

func (s *DetachHostAccountsFromUserGroupResponseBodyResultsHostAccounts) SetHostAccountId(v string) *DetachHostAccountsFromUserGroupResponseBodyResultsHostAccounts {
	s.HostAccountId = &v
	return s
}

type DetachHostAccountsFromUserGroupResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetachHostAccountsFromUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetachHostAccountsFromUserGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachHostAccountsFromUserGroupResponse) GoString() string {
	return s.String()
}

func (s *DetachHostAccountsFromUserGroupResponse) SetHeaders(v map[string]*string) *DetachHostAccountsFromUserGroupResponse {
	s.Headers = v
	return s
}

func (s *DetachHostAccountsFromUserGroupResponse) SetBody(v *DetachHostAccountsFromUserGroupResponseBody) *DetachHostAccountsFromUserGroupResponse {
	s.Body = v
	return s
}

type DetachHostGroupAccountsFromUserRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	HostGroups *string `json:"HostGroups,omitempty" xml:"HostGroups,omitempty"`
}

func (s DetachHostGroupAccountsFromUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachHostGroupAccountsFromUserRequest) GoString() string {
	return s.String()
}

func (s *DetachHostGroupAccountsFromUserRequest) SetSourceIp(v string) *DetachHostGroupAccountsFromUserRequest {
	s.SourceIp = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserRequest) SetInstanceId(v string) *DetachHostGroupAccountsFromUserRequest {
	s.InstanceId = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserRequest) SetRegionId(v string) *DetachHostGroupAccountsFromUserRequest {
	s.RegionId = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserRequest) SetUserId(v string) *DetachHostGroupAccountsFromUserRequest {
	s.UserId = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserRequest) SetHostGroups(v string) *DetachHostGroupAccountsFromUserRequest {
	s.HostGroups = &v
	return s
}

type DetachHostGroupAccountsFromUserResponseBody struct {
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*DetachHostGroupAccountsFromUserResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s DetachHostGroupAccountsFromUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachHostGroupAccountsFromUserResponseBody) GoString() string {
	return s.String()
}

func (s *DetachHostGroupAccountsFromUserResponseBody) SetRequestId(v string) *DetachHostGroupAccountsFromUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserResponseBody) SetResults(v []*DetachHostGroupAccountsFromUserResponseBodyResults) *DetachHostGroupAccountsFromUserResponseBody {
	s.Results = v
	return s
}

type DetachHostGroupAccountsFromUserResponseBodyResults struct {
	HostAccountNames []*DetachHostGroupAccountsFromUserResponseBodyResultsHostAccountNames `json:"HostAccountNames,omitempty" xml:"HostAccountNames,omitempty" type:"Repeated"`
	UserId           *string                                                               `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Code             *string                                                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Message          *string                                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	HostGroupId      *string                                                               `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
}

func (s DetachHostGroupAccountsFromUserResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s DetachHostGroupAccountsFromUserResponseBodyResults) GoString() string {
	return s.String()
}

func (s *DetachHostGroupAccountsFromUserResponseBodyResults) SetHostAccountNames(v []*DetachHostGroupAccountsFromUserResponseBodyResultsHostAccountNames) *DetachHostGroupAccountsFromUserResponseBodyResults {
	s.HostAccountNames = v
	return s
}

func (s *DetachHostGroupAccountsFromUserResponseBodyResults) SetUserId(v string) *DetachHostGroupAccountsFromUserResponseBodyResults {
	s.UserId = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserResponseBodyResults) SetCode(v string) *DetachHostGroupAccountsFromUserResponseBodyResults {
	s.Code = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserResponseBodyResults) SetMessage(v string) *DetachHostGroupAccountsFromUserResponseBodyResults {
	s.Message = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserResponseBodyResults) SetHostGroupId(v string) *DetachHostGroupAccountsFromUserResponseBodyResults {
	s.HostGroupId = &v
	return s
}

type DetachHostGroupAccountsFromUserResponseBodyResultsHostAccountNames struct {
	Code            *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message         *string `json:"Message,omitempty" xml:"Message,omitempty"`
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
}

func (s DetachHostGroupAccountsFromUserResponseBodyResultsHostAccountNames) String() string {
	return tea.Prettify(s)
}

func (s DetachHostGroupAccountsFromUserResponseBodyResultsHostAccountNames) GoString() string {
	return s.String()
}

func (s *DetachHostGroupAccountsFromUserResponseBodyResultsHostAccountNames) SetCode(v string) *DetachHostGroupAccountsFromUserResponseBodyResultsHostAccountNames {
	s.Code = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserResponseBodyResultsHostAccountNames) SetMessage(v string) *DetachHostGroupAccountsFromUserResponseBodyResultsHostAccountNames {
	s.Message = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserResponseBodyResultsHostAccountNames) SetHostAccountName(v string) *DetachHostGroupAccountsFromUserResponseBodyResultsHostAccountNames {
	s.HostAccountName = &v
	return s
}

type DetachHostGroupAccountsFromUserResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetachHostGroupAccountsFromUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetachHostGroupAccountsFromUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachHostGroupAccountsFromUserResponse) GoString() string {
	return s.String()
}

func (s *DetachHostGroupAccountsFromUserResponse) SetHeaders(v map[string]*string) *DetachHostGroupAccountsFromUserResponse {
	s.Headers = v
	return s
}

func (s *DetachHostGroupAccountsFromUserResponse) SetBody(v *DetachHostGroupAccountsFromUserResponseBody) *DetachHostGroupAccountsFromUserResponse {
	s.Body = v
	return s
}

type DetachHostGroupAccountsFromUserGroupRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	HostGroups  *string `json:"HostGroups,omitempty" xml:"HostGroups,omitempty"`
}

func (s DetachHostGroupAccountsFromUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachHostGroupAccountsFromUserGroupRequest) GoString() string {
	return s.String()
}

func (s *DetachHostGroupAccountsFromUserGroupRequest) SetSourceIp(v string) *DetachHostGroupAccountsFromUserGroupRequest {
	s.SourceIp = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserGroupRequest) SetInstanceId(v string) *DetachHostGroupAccountsFromUserGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserGroupRequest) SetRegionId(v string) *DetachHostGroupAccountsFromUserGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserGroupRequest) SetUserGroupId(v string) *DetachHostGroupAccountsFromUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserGroupRequest) SetHostGroups(v string) *DetachHostGroupAccountsFromUserGroupRequest {
	s.HostGroups = &v
	return s
}

type DetachHostGroupAccountsFromUserGroupResponseBody struct {
	RequestId *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*DetachHostGroupAccountsFromUserGroupResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s DetachHostGroupAccountsFromUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachHostGroupAccountsFromUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DetachHostGroupAccountsFromUserGroupResponseBody) SetRequestId(v string) *DetachHostGroupAccountsFromUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserGroupResponseBody) SetResults(v []*DetachHostGroupAccountsFromUserGroupResponseBodyResults) *DetachHostGroupAccountsFromUserGroupResponseBody {
	s.Results = v
	return s
}

type DetachHostGroupAccountsFromUserGroupResponseBodyResults struct {
	HostAccountNames []*DetachHostGroupAccountsFromUserGroupResponseBodyResultsHostAccountNames `json:"HostAccountNames,omitempty" xml:"HostAccountNames,omitempty" type:"Repeated"`
	UserGroupId      *string                                                                    `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	Code             *string                                                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Message          *string                                                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	HostGroupId      *string                                                                    `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
}

func (s DetachHostGroupAccountsFromUserGroupResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s DetachHostGroupAccountsFromUserGroupResponseBodyResults) GoString() string {
	return s.String()
}

func (s *DetachHostGroupAccountsFromUserGroupResponseBodyResults) SetHostAccountNames(v []*DetachHostGroupAccountsFromUserGroupResponseBodyResultsHostAccountNames) *DetachHostGroupAccountsFromUserGroupResponseBodyResults {
	s.HostAccountNames = v
	return s
}

func (s *DetachHostGroupAccountsFromUserGroupResponseBodyResults) SetUserGroupId(v string) *DetachHostGroupAccountsFromUserGroupResponseBodyResults {
	s.UserGroupId = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserGroupResponseBodyResults) SetCode(v string) *DetachHostGroupAccountsFromUserGroupResponseBodyResults {
	s.Code = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserGroupResponseBodyResults) SetMessage(v string) *DetachHostGroupAccountsFromUserGroupResponseBodyResults {
	s.Message = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserGroupResponseBodyResults) SetHostGroupId(v string) *DetachHostGroupAccountsFromUserGroupResponseBodyResults {
	s.HostGroupId = &v
	return s
}

type DetachHostGroupAccountsFromUserGroupResponseBodyResultsHostAccountNames struct {
	Code            *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message         *string `json:"Message,omitempty" xml:"Message,omitempty"`
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
}

func (s DetachHostGroupAccountsFromUserGroupResponseBodyResultsHostAccountNames) String() string {
	return tea.Prettify(s)
}

func (s DetachHostGroupAccountsFromUserGroupResponseBodyResultsHostAccountNames) GoString() string {
	return s.String()
}

func (s *DetachHostGroupAccountsFromUserGroupResponseBodyResultsHostAccountNames) SetCode(v string) *DetachHostGroupAccountsFromUserGroupResponseBodyResultsHostAccountNames {
	s.Code = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserGroupResponseBodyResultsHostAccountNames) SetMessage(v string) *DetachHostGroupAccountsFromUserGroupResponseBodyResultsHostAccountNames {
	s.Message = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserGroupResponseBodyResultsHostAccountNames) SetHostAccountName(v string) *DetachHostGroupAccountsFromUserGroupResponseBodyResultsHostAccountNames {
	s.HostAccountName = &v
	return s
}

type DetachHostGroupAccountsFromUserGroupResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DetachHostGroupAccountsFromUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetachHostGroupAccountsFromUserGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachHostGroupAccountsFromUserGroupResponse) GoString() string {
	return s.String()
}

func (s *DetachHostGroupAccountsFromUserGroupResponse) SetHeaders(v map[string]*string) *DetachHostGroupAccountsFromUserGroupResponse {
	s.Headers = v
	return s
}

func (s *DetachHostGroupAccountsFromUserGroupResponse) SetBody(v *DetachHostGroupAccountsFromUserGroupResponseBody) *DetachHostGroupAccountsFromUserGroupResponse {
	s.Body = v
	return s
}

type DisableInstancePublicAccessRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DisableInstancePublicAccessRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableInstancePublicAccessRequest) GoString() string {
	return s.String()
}

func (s *DisableInstancePublicAccessRequest) SetSourceIp(v string) *DisableInstancePublicAccessRequest {
	s.SourceIp = &v
	return s
}

func (s *DisableInstancePublicAccessRequest) SetLang(v string) *DisableInstancePublicAccessRequest {
	s.Lang = &v
	return s
}

func (s *DisableInstancePublicAccessRequest) SetInstanceId(v string) *DisableInstancePublicAccessRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableInstancePublicAccessRequest) SetRegionId(v string) *DisableInstancePublicAccessRequest {
	s.RegionId = &v
	return s
}

type DisableInstancePublicAccessResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DisableInstancePublicAccessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableInstancePublicAccessResponseBody) GoString() string {
	return s.String()
}

func (s *DisableInstancePublicAccessResponseBody) SetRequestId(v string) *DisableInstancePublicAccessResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableInstancePublicAccessResponseBody) SetInstanceId(v string) *DisableInstancePublicAccessResponseBody {
	s.InstanceId = &v
	return s
}

type DisableInstancePublicAccessResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableInstancePublicAccessResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableInstancePublicAccessResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableInstancePublicAccessResponse) GoString() string {
	return s.String()
}

func (s *DisableInstancePublicAccessResponse) SetHeaders(v map[string]*string) *DisableInstancePublicAccessResponse {
	s.Headers = v
	return s
}

func (s *DisableInstancePublicAccessResponse) SetBody(v *DisableInstancePublicAccessResponseBody) *DisableInstancePublicAccessResponse {
	s.Body = v
	return s
}

type EnableInstancePublicAccessRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s EnableInstancePublicAccessRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableInstancePublicAccessRequest) GoString() string {
	return s.String()
}

func (s *EnableInstancePublicAccessRequest) SetSourceIp(v string) *EnableInstancePublicAccessRequest {
	s.SourceIp = &v
	return s
}

func (s *EnableInstancePublicAccessRequest) SetLang(v string) *EnableInstancePublicAccessRequest {
	s.Lang = &v
	return s
}

func (s *EnableInstancePublicAccessRequest) SetInstanceId(v string) *EnableInstancePublicAccessRequest {
	s.InstanceId = &v
	return s
}

func (s *EnableInstancePublicAccessRequest) SetRegionId(v string) *EnableInstancePublicAccessRequest {
	s.RegionId = &v
	return s
}

type EnableInstancePublicAccessResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s EnableInstancePublicAccessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableInstancePublicAccessResponseBody) GoString() string {
	return s.String()
}

func (s *EnableInstancePublicAccessResponseBody) SetRequestId(v string) *EnableInstancePublicAccessResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableInstancePublicAccessResponseBody) SetInstanceId(v string) *EnableInstancePublicAccessResponseBody {
	s.InstanceId = &v
	return s
}

type EnableInstancePublicAccessResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableInstancePublicAccessResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableInstancePublicAccessResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableInstancePublicAccessResponse) GoString() string {
	return s.String()
}

func (s *EnableInstancePublicAccessResponse) SetHeaders(v map[string]*string) *EnableInstancePublicAccessResponse {
	s.Headers = v
	return s
}

func (s *EnableInstancePublicAccessResponse) SetBody(v *EnableInstancePublicAccessResponseBody) *EnableInstancePublicAccessResponse {
	s.Body = v
	return s
}

type GetHostRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	HostId     *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
}

func (s GetHostRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHostRequest) GoString() string {
	return s.String()
}

func (s *GetHostRequest) SetSourceIp(v string) *GetHostRequest {
	s.SourceIp = &v
	return s
}

func (s *GetHostRequest) SetInstanceId(v string) *GetHostRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHostRequest) SetRegionId(v string) *GetHostRequest {
	s.RegionId = &v
	return s
}

func (s *GetHostRequest) SetHostId(v string) *GetHostRequest {
	s.HostId = &v
	return s
}

type GetHostResponseBody struct {
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Host      *GetHostResponseBodyHost `json:"Host,omitempty" xml:"Host,omitempty" type:"Struct"`
}

func (s GetHostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHostResponseBody) GoString() string {
	return s.String()
}

func (s *GetHostResponseBody) SetRequestId(v string) *GetHostResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHostResponseBody) SetHost(v *GetHostResponseBodyHost) *GetHostResponseBody {
	s.Host = v
	return s
}

type GetHostResponseBodyHost struct {
	Comment             *string                             `json:"Comment,omitempty" xml:"Comment,omitempty"`
	ActiveAddressType   *string                             `json:"ActiveAddressType,omitempty" xml:"ActiveAddressType,omitempty"`
	HostPublicAddress   *string                             `json:"HostPublicAddress,omitempty" xml:"HostPublicAddress,omitempty"`
	Protocols           []*GetHostResponseBodyHostProtocols `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
	HostName            *string                             `json:"HostName,omitempty" xml:"HostName,omitempty"`
	Source              *string                             `json:"Source,omitempty" xml:"Source,omitempty"`
	HostPrivateAddress  *string                             `json:"HostPrivateAddress,omitempty" xml:"HostPrivateAddress,omitempty"`
	OSType              *string                             `json:"OSType,omitempty" xml:"OSType,omitempty"`
	HostId              *string                             `json:"HostId,omitempty" xml:"HostId,omitempty"`
	SourceInstanceState *string                             `json:"SourceInstanceState,omitempty" xml:"SourceInstanceState,omitempty"`
	SourceInstanceId    *string                             `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
}

func (s GetHostResponseBodyHost) String() string {
	return tea.Prettify(s)
}

func (s GetHostResponseBodyHost) GoString() string {
	return s.String()
}

func (s *GetHostResponseBodyHost) SetComment(v string) *GetHostResponseBodyHost {
	s.Comment = &v
	return s
}

func (s *GetHostResponseBodyHost) SetActiveAddressType(v string) *GetHostResponseBodyHost {
	s.ActiveAddressType = &v
	return s
}

func (s *GetHostResponseBodyHost) SetHostPublicAddress(v string) *GetHostResponseBodyHost {
	s.HostPublicAddress = &v
	return s
}

func (s *GetHostResponseBodyHost) SetProtocols(v []*GetHostResponseBodyHostProtocols) *GetHostResponseBodyHost {
	s.Protocols = v
	return s
}

func (s *GetHostResponseBodyHost) SetHostName(v string) *GetHostResponseBodyHost {
	s.HostName = &v
	return s
}

func (s *GetHostResponseBodyHost) SetSource(v string) *GetHostResponseBodyHost {
	s.Source = &v
	return s
}

func (s *GetHostResponseBodyHost) SetHostPrivateAddress(v string) *GetHostResponseBodyHost {
	s.HostPrivateAddress = &v
	return s
}

func (s *GetHostResponseBodyHost) SetOSType(v string) *GetHostResponseBodyHost {
	s.OSType = &v
	return s
}

func (s *GetHostResponseBodyHost) SetHostId(v string) *GetHostResponseBodyHost {
	s.HostId = &v
	return s
}

func (s *GetHostResponseBodyHost) SetSourceInstanceState(v string) *GetHostResponseBodyHost {
	s.SourceInstanceState = &v
	return s
}

func (s *GetHostResponseBodyHost) SetSourceInstanceId(v string) *GetHostResponseBodyHost {
	s.SourceInstanceId = &v
	return s
}

type GetHostResponseBodyHostProtocols struct {
	ProtocolName    *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
	HostFingerPrint *string `json:"HostFingerPrint,omitempty" xml:"HostFingerPrint,omitempty"`
	Port            *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s GetHostResponseBodyHostProtocols) String() string {
	return tea.Prettify(s)
}

func (s GetHostResponseBodyHostProtocols) GoString() string {
	return s.String()
}

func (s *GetHostResponseBodyHostProtocols) SetProtocolName(v string) *GetHostResponseBodyHostProtocols {
	s.ProtocolName = &v
	return s
}

func (s *GetHostResponseBodyHostProtocols) SetHostFingerPrint(v string) *GetHostResponseBodyHostProtocols {
	s.HostFingerPrint = &v
	return s
}

func (s *GetHostResponseBodyHostProtocols) SetPort(v int32) *GetHostResponseBodyHostProtocols {
	s.Port = &v
	return s
}

type GetHostResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetHostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHostResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHostResponse) GoString() string {
	return s.String()
}

func (s *GetHostResponse) SetHeaders(v map[string]*string) *GetHostResponse {
	s.Headers = v
	return s
}

func (s *GetHostResponse) SetBody(v *GetHostResponseBody) *GetHostResponse {
	s.Body = v
	return s
}

type GetHostAccountRequest struct {
	SourceIp      *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
}

func (s GetHostAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHostAccountRequest) GoString() string {
	return s.String()
}

func (s *GetHostAccountRequest) SetSourceIp(v string) *GetHostAccountRequest {
	s.SourceIp = &v
	return s
}

func (s *GetHostAccountRequest) SetInstanceId(v string) *GetHostAccountRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHostAccountRequest) SetRegionId(v string) *GetHostAccountRequest {
	s.RegionId = &v
	return s
}

func (s *GetHostAccountRequest) SetHostAccountId(v string) *GetHostAccountRequest {
	s.HostAccountId = &v
	return s
}

type GetHostAccountResponseBody struct {
	RequestId   *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HostAccount *GetHostAccountResponseBodyHostAccount `json:"HostAccount,omitempty" xml:"HostAccount,omitempty" type:"Struct"`
}

func (s GetHostAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHostAccountResponseBody) GoString() string {
	return s.String()
}

func (s *GetHostAccountResponseBody) SetRequestId(v string) *GetHostAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHostAccountResponseBody) SetHostAccount(v *GetHostAccountResponseBodyHostAccount) *GetHostAccountResponseBody {
	s.HostAccount = v
	return s
}

type GetHostAccountResponseBodyHostAccount struct {
	HasPassword           *bool   `json:"HasPassword,omitempty" xml:"HasPassword,omitempty"`
	PrivateKeyFingerprint *string `json:"PrivateKeyFingerprint,omitempty" xml:"PrivateKeyFingerprint,omitempty"`
	ProtocolName          *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
	HostAccountName       *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	HostAccountId         *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	HostId                *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
}

func (s GetHostAccountResponseBodyHostAccount) String() string {
	return tea.Prettify(s)
}

func (s GetHostAccountResponseBodyHostAccount) GoString() string {
	return s.String()
}

func (s *GetHostAccountResponseBodyHostAccount) SetHasPassword(v bool) *GetHostAccountResponseBodyHostAccount {
	s.HasPassword = &v
	return s
}

func (s *GetHostAccountResponseBodyHostAccount) SetPrivateKeyFingerprint(v string) *GetHostAccountResponseBodyHostAccount {
	s.PrivateKeyFingerprint = &v
	return s
}

func (s *GetHostAccountResponseBodyHostAccount) SetProtocolName(v string) *GetHostAccountResponseBodyHostAccount {
	s.ProtocolName = &v
	return s
}

func (s *GetHostAccountResponseBodyHostAccount) SetHostAccountName(v string) *GetHostAccountResponseBodyHostAccount {
	s.HostAccountName = &v
	return s
}

func (s *GetHostAccountResponseBodyHostAccount) SetHostAccountId(v string) *GetHostAccountResponseBodyHostAccount {
	s.HostAccountId = &v
	return s
}

func (s *GetHostAccountResponseBodyHostAccount) SetHostId(v string) *GetHostAccountResponseBodyHostAccount {
	s.HostId = &v
	return s
}

type GetHostAccountResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetHostAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHostAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHostAccountResponse) GoString() string {
	return s.String()
}

func (s *GetHostAccountResponse) SetHeaders(v map[string]*string) *GetHostAccountResponse {
	s.Headers = v
	return s
}

func (s *GetHostAccountResponse) SetBody(v *GetHostAccountResponseBody) *GetHostAccountResponse {
	s.Body = v
	return s
}

type GetHostGroupRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
}

func (s GetHostGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHostGroupRequest) GoString() string {
	return s.String()
}

func (s *GetHostGroupRequest) SetSourceIp(v string) *GetHostGroupRequest {
	s.SourceIp = &v
	return s
}

func (s *GetHostGroupRequest) SetInstanceId(v string) *GetHostGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHostGroupRequest) SetRegionId(v string) *GetHostGroupRequest {
	s.RegionId = &v
	return s
}

func (s *GetHostGroupRequest) SetHostGroupId(v string) *GetHostGroupRequest {
	s.HostGroupId = &v
	return s
}

type GetHostGroupResponseBody struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HostGroup *GetHostGroupResponseBodyHostGroup `json:"HostGroup,omitempty" xml:"HostGroup,omitempty" type:"Struct"`
}

func (s GetHostGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHostGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetHostGroupResponseBody) SetRequestId(v string) *GetHostGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHostGroupResponseBody) SetHostGroup(v *GetHostGroupResponseBodyHostGroup) *GetHostGroupResponseBody {
	s.HostGroup = v
	return s
}

type GetHostGroupResponseBodyHostGroup struct {
	Comment       *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	HostGroupId   *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	HostGroupName *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
}

func (s GetHostGroupResponseBodyHostGroup) String() string {
	return tea.Prettify(s)
}

func (s GetHostGroupResponseBodyHostGroup) GoString() string {
	return s.String()
}

func (s *GetHostGroupResponseBodyHostGroup) SetComment(v string) *GetHostGroupResponseBodyHostGroup {
	s.Comment = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroup) SetHostGroupId(v string) *GetHostGroupResponseBodyHostGroup {
	s.HostGroupId = &v
	return s
}

func (s *GetHostGroupResponseBodyHostGroup) SetHostGroupName(v string) *GetHostGroupResponseBodyHostGroup {
	s.HostGroupName = &v
	return s
}

type GetHostGroupResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetHostGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHostGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHostGroupResponse) GoString() string {
	return s.String()
}

func (s *GetHostGroupResponse) SetHeaders(v map[string]*string) *GetHostGroupResponse {
	s.Headers = v
	return s
}

func (s *GetHostGroupResponse) SetBody(v *GetHostGroupResponseBody) *GetHostGroupResponse {
	s.Body = v
	return s
}

type GetUserRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetUserRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserRequest) GoString() string {
	return s.String()
}

func (s *GetUserRequest) SetSourceIp(v string) *GetUserRequest {
	s.SourceIp = &v
	return s
}

func (s *GetUserRequest) SetInstanceId(v string) *GetUserRequest {
	s.InstanceId = &v
	return s
}

func (s *GetUserRequest) SetRegionId(v string) *GetUserRequest {
	s.RegionId = &v
	return s
}

func (s *GetUserRequest) SetUserId(v string) *GetUserRequest {
	s.UserId = &v
	return s
}

type GetUserResponseBody struct {
	User      *GetUserResponseBodyUser `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResponseBody) SetUser(v *GetUserResponseBodyUser) *GetUserResponseBody {
	s.User = v
	return s
}

func (s *GetUserResponseBody) SetRequestId(v string) *GetUserResponseBody {
	s.RequestId = &v
	return s
}

type GetUserResponseBodyUser struct {
	DisplayName       *string   `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Comment           *string   `json:"Comment,omitempty" xml:"Comment,omitempty"`
	Email             *string   `json:"Email,omitempty" xml:"Email,omitempty"`
	Mobile            *string   `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	MobileCountryCode *string   `json:"MobileCountryCode,omitempty" xml:"MobileCountryCode,omitempty"`
	UserId            *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Source            *string   `json:"Source,omitempty" xml:"Source,omitempty"`
	UserName          *string   `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserState         []*string `json:"UserState,omitempty" xml:"UserState,omitempty" type:"Repeated"`
	SourceUserId      *string   `json:"SourceUserId,omitempty" xml:"SourceUserId,omitempty"`
}

func (s GetUserResponseBodyUser) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyUser) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyUser) SetDisplayName(v string) *GetUserResponseBodyUser {
	s.DisplayName = &v
	return s
}

func (s *GetUserResponseBodyUser) SetComment(v string) *GetUserResponseBodyUser {
	s.Comment = &v
	return s
}

func (s *GetUserResponseBodyUser) SetEmail(v string) *GetUserResponseBodyUser {
	s.Email = &v
	return s
}

func (s *GetUserResponseBodyUser) SetMobile(v string) *GetUserResponseBodyUser {
	s.Mobile = &v
	return s
}

func (s *GetUserResponseBodyUser) SetMobileCountryCode(v string) *GetUserResponseBodyUser {
	s.MobileCountryCode = &v
	return s
}

func (s *GetUserResponseBodyUser) SetUserId(v string) *GetUserResponseBodyUser {
	s.UserId = &v
	return s
}

func (s *GetUserResponseBodyUser) SetSource(v string) *GetUserResponseBodyUser {
	s.Source = &v
	return s
}

func (s *GetUserResponseBodyUser) SetUserName(v string) *GetUserResponseBodyUser {
	s.UserName = &v
	return s
}

func (s *GetUserResponseBodyUser) SetUserState(v []*string) *GetUserResponseBodyUser {
	s.UserState = v
	return s
}

func (s *GetUserResponseBodyUser) SetSourceUserId(v string) *GetUserResponseBodyUser {
	s.SourceUserId = &v
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

type GetUserGroupRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s GetUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserGroupRequest) GoString() string {
	return s.String()
}

func (s *GetUserGroupRequest) SetSourceIp(v string) *GetUserGroupRequest {
	s.SourceIp = &v
	return s
}

func (s *GetUserGroupRequest) SetInstanceId(v string) *GetUserGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *GetUserGroupRequest) SetRegionId(v string) *GetUserGroupRequest {
	s.RegionId = &v
	return s
}

func (s *GetUserGroupRequest) SetUserGroupId(v string) *GetUserGroupRequest {
	s.UserGroupId = &v
	return s
}

type GetUserGroupResponseBody struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserGroup *GetUserGroupResponseBodyUserGroup `json:"UserGroup,omitempty" xml:"UserGroup,omitempty" type:"Struct"`
}

func (s GetUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserGroupResponseBody) SetRequestId(v string) *GetUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserGroupResponseBody) SetUserGroup(v *GetUserGroupResponseBodyUserGroup) *GetUserGroupResponseBody {
	s.UserGroup = v
	return s
}

type GetUserGroupResponseBodyUserGroup struct {
	Comment       *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	UserGroupId   *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
}

func (s GetUserGroupResponseBodyUserGroup) String() string {
	return tea.Prettify(s)
}

func (s GetUserGroupResponseBodyUserGroup) GoString() string {
	return s.String()
}

func (s *GetUserGroupResponseBodyUserGroup) SetComment(v string) *GetUserGroupResponseBodyUserGroup {
	s.Comment = &v
	return s
}

func (s *GetUserGroupResponseBodyUserGroup) SetUserGroupId(v string) *GetUserGroupResponseBodyUserGroup {
	s.UserGroupId = &v
	return s
}

func (s *GetUserGroupResponseBodyUserGroup) SetUserGroupName(v string) *GetUserGroupResponseBodyUserGroup {
	s.UserGroupName = &v
	return s
}

type GetUserGroupResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserGroupResponse) GoString() string {
	return s.String()
}

func (s *GetUserGroupResponse) SetHeaders(v map[string]*string) *GetUserGroupResponse {
	s.Headers = v
	return s
}

func (s *GetUserGroupResponse) SetBody(v *GetUserGroupResponseBody) *GetUserGroupResponse {
	s.Body = v
	return s
}

type ListHostAccountsRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	HostId          *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	PageNumber      *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	ProtocolName    *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
}

func (s ListHostAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHostAccountsRequest) GoString() string {
	return s.String()
}

func (s *ListHostAccountsRequest) SetSourceIp(v string) *ListHostAccountsRequest {
	s.SourceIp = &v
	return s
}

func (s *ListHostAccountsRequest) SetInstanceId(v string) *ListHostAccountsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHostAccountsRequest) SetRegionId(v string) *ListHostAccountsRequest {
	s.RegionId = &v
	return s
}

func (s *ListHostAccountsRequest) SetHostId(v string) *ListHostAccountsRequest {
	s.HostId = &v
	return s
}

func (s *ListHostAccountsRequest) SetPageNumber(v string) *ListHostAccountsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHostAccountsRequest) SetPageSize(v string) *ListHostAccountsRequest {
	s.PageSize = &v
	return s
}

func (s *ListHostAccountsRequest) SetHostAccountName(v string) *ListHostAccountsRequest {
	s.HostAccountName = &v
	return s
}

func (s *ListHostAccountsRequest) SetProtocolName(v string) *ListHostAccountsRequest {
	s.ProtocolName = &v
	return s
}

type ListHostAccountsResponseBody struct {
	TotalCount   *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HostAccounts []*ListHostAccountsResponseBodyHostAccounts `json:"HostAccounts,omitempty" xml:"HostAccounts,omitempty" type:"Repeated"`
}

func (s ListHostAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHostAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *ListHostAccountsResponseBody) SetTotalCount(v int32) *ListHostAccountsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListHostAccountsResponseBody) SetRequestId(v string) *ListHostAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHostAccountsResponseBody) SetHostAccounts(v []*ListHostAccountsResponseBodyHostAccounts) *ListHostAccountsResponseBody {
	s.HostAccounts = v
	return s
}

type ListHostAccountsResponseBodyHostAccounts struct {
	HasPassword           *bool   `json:"HasPassword,omitempty" xml:"HasPassword,omitempty"`
	PrivateKeyFingerprint *string `json:"PrivateKeyFingerprint,omitempty" xml:"PrivateKeyFingerprint,omitempty"`
	ProtocolName          *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
	HostAccountName       *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	HostAccountId         *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	HostId                *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
}

func (s ListHostAccountsResponseBodyHostAccounts) String() string {
	return tea.Prettify(s)
}

func (s ListHostAccountsResponseBodyHostAccounts) GoString() string {
	return s.String()
}

func (s *ListHostAccountsResponseBodyHostAccounts) SetHasPassword(v bool) *ListHostAccountsResponseBodyHostAccounts {
	s.HasPassword = &v
	return s
}

func (s *ListHostAccountsResponseBodyHostAccounts) SetPrivateKeyFingerprint(v string) *ListHostAccountsResponseBodyHostAccounts {
	s.PrivateKeyFingerprint = &v
	return s
}

func (s *ListHostAccountsResponseBodyHostAccounts) SetProtocolName(v string) *ListHostAccountsResponseBodyHostAccounts {
	s.ProtocolName = &v
	return s
}

func (s *ListHostAccountsResponseBodyHostAccounts) SetHostAccountName(v string) *ListHostAccountsResponseBodyHostAccounts {
	s.HostAccountName = &v
	return s
}

func (s *ListHostAccountsResponseBodyHostAccounts) SetHostAccountId(v string) *ListHostAccountsResponseBodyHostAccounts {
	s.HostAccountId = &v
	return s
}

func (s *ListHostAccountsResponseBodyHostAccounts) SetHostId(v string) *ListHostAccountsResponseBodyHostAccounts {
	s.HostId = &v
	return s
}

type ListHostAccountsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListHostAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListHostAccountsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHostAccountsResponse) GoString() string {
	return s.String()
}

func (s *ListHostAccountsResponse) SetHeaders(v map[string]*string) *ListHostAccountsResponse {
	s.Headers = v
	return s
}

func (s *ListHostAccountsResponse) SetBody(v *ListHostAccountsResponseBody) *ListHostAccountsResponse {
	s.Body = v
	return s
}

type ListHostAccountsForUserRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserId          *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	HostId          *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	PageNumber      *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
}

func (s ListHostAccountsForUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHostAccountsForUserRequest) GoString() string {
	return s.String()
}

func (s *ListHostAccountsForUserRequest) SetSourceIp(v string) *ListHostAccountsForUserRequest {
	s.SourceIp = &v
	return s
}

func (s *ListHostAccountsForUserRequest) SetInstanceId(v string) *ListHostAccountsForUserRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHostAccountsForUserRequest) SetRegionId(v string) *ListHostAccountsForUserRequest {
	s.RegionId = &v
	return s
}

func (s *ListHostAccountsForUserRequest) SetUserId(v string) *ListHostAccountsForUserRequest {
	s.UserId = &v
	return s
}

func (s *ListHostAccountsForUserRequest) SetHostId(v string) *ListHostAccountsForUserRequest {
	s.HostId = &v
	return s
}

func (s *ListHostAccountsForUserRequest) SetPageNumber(v string) *ListHostAccountsForUserRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHostAccountsForUserRequest) SetPageSize(v string) *ListHostAccountsForUserRequest {
	s.PageSize = &v
	return s
}

func (s *ListHostAccountsForUserRequest) SetHostAccountName(v string) *ListHostAccountsForUserRequest {
	s.HostAccountName = &v
	return s
}

type ListHostAccountsForUserResponseBody struct {
	TotalCount   *int32                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId    *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HostAccounts []*ListHostAccountsForUserResponseBodyHostAccounts `json:"HostAccounts,omitempty" xml:"HostAccounts,omitempty" type:"Repeated"`
}

func (s ListHostAccountsForUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHostAccountsForUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListHostAccountsForUserResponseBody) SetTotalCount(v int32) *ListHostAccountsForUserResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListHostAccountsForUserResponseBody) SetRequestId(v string) *ListHostAccountsForUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHostAccountsForUserResponseBody) SetHostAccounts(v []*ListHostAccountsForUserResponseBodyHostAccounts) *ListHostAccountsForUserResponseBody {
	s.HostAccounts = v
	return s
}

type ListHostAccountsForUserResponseBodyHostAccounts struct {
	IsAuthorized    *bool   `json:"IsAuthorized,omitempty" xml:"IsAuthorized,omitempty"`
	ProtocolName    *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	HostAccountId   *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	HostId          *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
}

func (s ListHostAccountsForUserResponseBodyHostAccounts) String() string {
	return tea.Prettify(s)
}

func (s ListHostAccountsForUserResponseBodyHostAccounts) GoString() string {
	return s.String()
}

func (s *ListHostAccountsForUserResponseBodyHostAccounts) SetIsAuthorized(v bool) *ListHostAccountsForUserResponseBodyHostAccounts {
	s.IsAuthorized = &v
	return s
}

func (s *ListHostAccountsForUserResponseBodyHostAccounts) SetProtocolName(v string) *ListHostAccountsForUserResponseBodyHostAccounts {
	s.ProtocolName = &v
	return s
}

func (s *ListHostAccountsForUserResponseBodyHostAccounts) SetHostAccountName(v string) *ListHostAccountsForUserResponseBodyHostAccounts {
	s.HostAccountName = &v
	return s
}

func (s *ListHostAccountsForUserResponseBodyHostAccounts) SetHostAccountId(v string) *ListHostAccountsForUserResponseBodyHostAccounts {
	s.HostAccountId = &v
	return s
}

func (s *ListHostAccountsForUserResponseBodyHostAccounts) SetHostId(v string) *ListHostAccountsForUserResponseBodyHostAccounts {
	s.HostId = &v
	return s
}

type ListHostAccountsForUserResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListHostAccountsForUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListHostAccountsForUserResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHostAccountsForUserResponse) GoString() string {
	return s.String()
}

func (s *ListHostAccountsForUserResponse) SetHeaders(v map[string]*string) *ListHostAccountsForUserResponse {
	s.Headers = v
	return s
}

func (s *ListHostAccountsForUserResponse) SetBody(v *ListHostAccountsForUserResponseBody) *ListHostAccountsForUserResponse {
	s.Body = v
	return s
}

type ListHostAccountsForUserGroupRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserGroupId     *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	HostId          *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	PageNumber      *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
}

func (s ListHostAccountsForUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHostAccountsForUserGroupRequest) GoString() string {
	return s.String()
}

func (s *ListHostAccountsForUserGroupRequest) SetSourceIp(v string) *ListHostAccountsForUserGroupRequest {
	s.SourceIp = &v
	return s
}

func (s *ListHostAccountsForUserGroupRequest) SetInstanceId(v string) *ListHostAccountsForUserGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHostAccountsForUserGroupRequest) SetRegionId(v string) *ListHostAccountsForUserGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ListHostAccountsForUserGroupRequest) SetUserGroupId(v string) *ListHostAccountsForUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *ListHostAccountsForUserGroupRequest) SetHostId(v string) *ListHostAccountsForUserGroupRequest {
	s.HostId = &v
	return s
}

func (s *ListHostAccountsForUserGroupRequest) SetPageNumber(v string) *ListHostAccountsForUserGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHostAccountsForUserGroupRequest) SetPageSize(v string) *ListHostAccountsForUserGroupRequest {
	s.PageSize = &v
	return s
}

func (s *ListHostAccountsForUserGroupRequest) SetHostAccountName(v string) *ListHostAccountsForUserGroupRequest {
	s.HostAccountName = &v
	return s
}

type ListHostAccountsForUserGroupResponseBody struct {
	TotalCount   *int32                                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId    *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HostAccounts []*ListHostAccountsForUserGroupResponseBodyHostAccounts `json:"HostAccounts,omitempty" xml:"HostAccounts,omitempty" type:"Repeated"`
}

func (s ListHostAccountsForUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHostAccountsForUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListHostAccountsForUserGroupResponseBody) SetTotalCount(v int32) *ListHostAccountsForUserGroupResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListHostAccountsForUserGroupResponseBody) SetRequestId(v string) *ListHostAccountsForUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHostAccountsForUserGroupResponseBody) SetHostAccounts(v []*ListHostAccountsForUserGroupResponseBodyHostAccounts) *ListHostAccountsForUserGroupResponseBody {
	s.HostAccounts = v
	return s
}

type ListHostAccountsForUserGroupResponseBodyHostAccounts struct {
	IsAuthorized    *bool   `json:"IsAuthorized,omitempty" xml:"IsAuthorized,omitempty"`
	ProtocolName    *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	HostAccountId   *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	HostId          *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
}

func (s ListHostAccountsForUserGroupResponseBodyHostAccounts) String() string {
	return tea.Prettify(s)
}

func (s ListHostAccountsForUserGroupResponseBodyHostAccounts) GoString() string {
	return s.String()
}

func (s *ListHostAccountsForUserGroupResponseBodyHostAccounts) SetIsAuthorized(v bool) *ListHostAccountsForUserGroupResponseBodyHostAccounts {
	s.IsAuthorized = &v
	return s
}

func (s *ListHostAccountsForUserGroupResponseBodyHostAccounts) SetProtocolName(v string) *ListHostAccountsForUserGroupResponseBodyHostAccounts {
	s.ProtocolName = &v
	return s
}

func (s *ListHostAccountsForUserGroupResponseBodyHostAccounts) SetHostAccountName(v string) *ListHostAccountsForUserGroupResponseBodyHostAccounts {
	s.HostAccountName = &v
	return s
}

func (s *ListHostAccountsForUserGroupResponseBodyHostAccounts) SetHostAccountId(v string) *ListHostAccountsForUserGroupResponseBodyHostAccounts {
	s.HostAccountId = &v
	return s
}

func (s *ListHostAccountsForUserGroupResponseBodyHostAccounts) SetHostId(v string) *ListHostAccountsForUserGroupResponseBodyHostAccounts {
	s.HostId = &v
	return s
}

type ListHostAccountsForUserGroupResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListHostAccountsForUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListHostAccountsForUserGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHostAccountsForUserGroupResponse) GoString() string {
	return s.String()
}

func (s *ListHostAccountsForUserGroupResponse) SetHeaders(v map[string]*string) *ListHostAccountsForUserGroupResponse {
	s.Headers = v
	return s
}

func (s *ListHostAccountsForUserGroupResponse) SetBody(v *ListHostAccountsForUserGroupResponseBody) *ListHostAccountsForUserGroupResponse {
	s.Body = v
	return s
}

type ListHostGroupAccountNamesForUserRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
}

func (s ListHostGroupAccountNamesForUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHostGroupAccountNamesForUserRequest) GoString() string {
	return s.String()
}

func (s *ListHostGroupAccountNamesForUserRequest) SetSourceIp(v string) *ListHostGroupAccountNamesForUserRequest {
	s.SourceIp = &v
	return s
}

func (s *ListHostGroupAccountNamesForUserRequest) SetInstanceId(v string) *ListHostGroupAccountNamesForUserRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHostGroupAccountNamesForUserRequest) SetRegionId(v string) *ListHostGroupAccountNamesForUserRequest {
	s.RegionId = &v
	return s
}

func (s *ListHostGroupAccountNamesForUserRequest) SetUserId(v string) *ListHostGroupAccountNamesForUserRequest {
	s.UserId = &v
	return s
}

func (s *ListHostGroupAccountNamesForUserRequest) SetHostGroupId(v string) *ListHostGroupAccountNamesForUserRequest {
	s.HostGroupId = &v
	return s
}

type ListHostGroupAccountNamesForUserResponseBody struct {
	RequestId        *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HostAccountNames []*string `json:"HostAccountNames,omitempty" xml:"HostAccountNames,omitempty" type:"Repeated"`
}

func (s ListHostGroupAccountNamesForUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHostGroupAccountNamesForUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListHostGroupAccountNamesForUserResponseBody) SetRequestId(v string) *ListHostGroupAccountNamesForUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHostGroupAccountNamesForUserResponseBody) SetHostAccountNames(v []*string) *ListHostGroupAccountNamesForUserResponseBody {
	s.HostAccountNames = v
	return s
}

type ListHostGroupAccountNamesForUserResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListHostGroupAccountNamesForUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListHostGroupAccountNamesForUserResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHostGroupAccountNamesForUserResponse) GoString() string {
	return s.String()
}

func (s *ListHostGroupAccountNamesForUserResponse) SetHeaders(v map[string]*string) *ListHostGroupAccountNamesForUserResponse {
	s.Headers = v
	return s
}

func (s *ListHostGroupAccountNamesForUserResponse) SetBody(v *ListHostGroupAccountNamesForUserResponseBody) *ListHostGroupAccountNamesForUserResponse {
	s.Body = v
	return s
}

type ListHostGroupAccountNamesForUserGroupRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
}

func (s ListHostGroupAccountNamesForUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHostGroupAccountNamesForUserGroupRequest) GoString() string {
	return s.String()
}

func (s *ListHostGroupAccountNamesForUserGroupRequest) SetSourceIp(v string) *ListHostGroupAccountNamesForUserGroupRequest {
	s.SourceIp = &v
	return s
}

func (s *ListHostGroupAccountNamesForUserGroupRequest) SetInstanceId(v string) *ListHostGroupAccountNamesForUserGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHostGroupAccountNamesForUserGroupRequest) SetRegionId(v string) *ListHostGroupAccountNamesForUserGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ListHostGroupAccountNamesForUserGroupRequest) SetUserGroupId(v string) *ListHostGroupAccountNamesForUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *ListHostGroupAccountNamesForUserGroupRequest) SetHostGroupId(v string) *ListHostGroupAccountNamesForUserGroupRequest {
	s.HostGroupId = &v
	return s
}

type ListHostGroupAccountNamesForUserGroupResponseBody struct {
	RequestId        *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HostAccountNames []*string `json:"HostAccountNames,omitempty" xml:"HostAccountNames,omitempty" type:"Repeated"`
}

func (s ListHostGroupAccountNamesForUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHostGroupAccountNamesForUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListHostGroupAccountNamesForUserGroupResponseBody) SetRequestId(v string) *ListHostGroupAccountNamesForUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHostGroupAccountNamesForUserGroupResponseBody) SetHostAccountNames(v []*string) *ListHostGroupAccountNamesForUserGroupResponseBody {
	s.HostAccountNames = v
	return s
}

type ListHostGroupAccountNamesForUserGroupResponse struct {
	Headers map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListHostGroupAccountNamesForUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListHostGroupAccountNamesForUserGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHostGroupAccountNamesForUserGroupResponse) GoString() string {
	return s.String()
}

func (s *ListHostGroupAccountNamesForUserGroupResponse) SetHeaders(v map[string]*string) *ListHostGroupAccountNamesForUserGroupResponse {
	s.Headers = v
	return s
}

func (s *ListHostGroupAccountNamesForUserGroupResponse) SetBody(v *ListHostGroupAccountNamesForUserGroupResponseBody) *ListHostGroupAccountNamesForUserGroupResponse {
	s.Body = v
	return s
}

type ListHostGroupsRequest struct {
	SourceIp      *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PageNumber    *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	HostGroupName *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
}

func (s ListHostGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHostGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListHostGroupsRequest) SetSourceIp(v string) *ListHostGroupsRequest {
	s.SourceIp = &v
	return s
}

func (s *ListHostGroupsRequest) SetInstanceId(v string) *ListHostGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHostGroupsRequest) SetRegionId(v string) *ListHostGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *ListHostGroupsRequest) SetPageNumber(v string) *ListHostGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHostGroupsRequest) SetPageSize(v string) *ListHostGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListHostGroupsRequest) SetHostGroupName(v string) *ListHostGroupsRequest {
	s.HostGroupName = &v
	return s
}

type ListHostGroupsResponseBody struct {
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HostGroups []*ListHostGroupsResponseBodyHostGroups `json:"HostGroups,omitempty" xml:"HostGroups,omitempty" type:"Repeated"`
}

func (s ListHostGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHostGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListHostGroupsResponseBody) SetTotalCount(v int32) *ListHostGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListHostGroupsResponseBody) SetRequestId(v string) *ListHostGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHostGroupsResponseBody) SetHostGroups(v []*ListHostGroupsResponseBodyHostGroups) *ListHostGroupsResponseBody {
	s.HostGroups = v
	return s
}

type ListHostGroupsResponseBodyHostGroups struct {
	MemberCount   *int32  `json:"MemberCount,omitempty" xml:"MemberCount,omitempty"`
	Comment       *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	HostGroupId   *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	HostGroupName *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
}

func (s ListHostGroupsResponseBodyHostGroups) String() string {
	return tea.Prettify(s)
}

func (s ListHostGroupsResponseBodyHostGroups) GoString() string {
	return s.String()
}

func (s *ListHostGroupsResponseBodyHostGroups) SetMemberCount(v int32) *ListHostGroupsResponseBodyHostGroups {
	s.MemberCount = &v
	return s
}

func (s *ListHostGroupsResponseBodyHostGroups) SetComment(v string) *ListHostGroupsResponseBodyHostGroups {
	s.Comment = &v
	return s
}

func (s *ListHostGroupsResponseBodyHostGroups) SetHostGroupId(v string) *ListHostGroupsResponseBodyHostGroups {
	s.HostGroupId = &v
	return s
}

func (s *ListHostGroupsResponseBodyHostGroups) SetHostGroupName(v string) *ListHostGroupsResponseBodyHostGroups {
	s.HostGroupName = &v
	return s
}

type ListHostGroupsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListHostGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListHostGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHostGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListHostGroupsResponse) SetHeaders(v map[string]*string) *ListHostGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListHostGroupsResponse) SetBody(v *ListHostGroupsResponseBody) *ListHostGroupsResponse {
	s.Body = v
	return s
}

type ListHostGroupsForUserRequest struct {
	SourceIp      *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Mode          *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	PageNumber    *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	HostGroupName *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
}

func (s ListHostGroupsForUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHostGroupsForUserRequest) GoString() string {
	return s.String()
}

func (s *ListHostGroupsForUserRequest) SetSourceIp(v string) *ListHostGroupsForUserRequest {
	s.SourceIp = &v
	return s
}

func (s *ListHostGroupsForUserRequest) SetInstanceId(v string) *ListHostGroupsForUserRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHostGroupsForUserRequest) SetRegionId(v string) *ListHostGroupsForUserRequest {
	s.RegionId = &v
	return s
}

func (s *ListHostGroupsForUserRequest) SetMode(v string) *ListHostGroupsForUserRequest {
	s.Mode = &v
	return s
}

func (s *ListHostGroupsForUserRequest) SetUserId(v string) *ListHostGroupsForUserRequest {
	s.UserId = &v
	return s
}

func (s *ListHostGroupsForUserRequest) SetPageNumber(v string) *ListHostGroupsForUserRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHostGroupsForUserRequest) SetPageSize(v string) *ListHostGroupsForUserRequest {
	s.PageSize = &v
	return s
}

func (s *ListHostGroupsForUserRequest) SetHostGroupName(v string) *ListHostGroupsForUserRequest {
	s.HostGroupName = &v
	return s
}

type ListHostGroupsForUserResponseBody struct {
	TotalCount *int32                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HostGroups []*ListHostGroupsForUserResponseBodyHostGroups `json:"HostGroups,omitempty" xml:"HostGroups,omitempty" type:"Repeated"`
}

func (s ListHostGroupsForUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHostGroupsForUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListHostGroupsForUserResponseBody) SetTotalCount(v int32) *ListHostGroupsForUserResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListHostGroupsForUserResponseBody) SetRequestId(v string) *ListHostGroupsForUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHostGroupsForUserResponseBody) SetHostGroups(v []*ListHostGroupsForUserResponseBodyHostGroups) *ListHostGroupsForUserResponseBody {
	s.HostGroups = v
	return s
}

type ListHostGroupsForUserResponseBodyHostGroups struct {
	Comment       *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	HostGroupId   *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	HostGroupName *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
}

func (s ListHostGroupsForUserResponseBodyHostGroups) String() string {
	return tea.Prettify(s)
}

func (s ListHostGroupsForUserResponseBodyHostGroups) GoString() string {
	return s.String()
}

func (s *ListHostGroupsForUserResponseBodyHostGroups) SetComment(v string) *ListHostGroupsForUserResponseBodyHostGroups {
	s.Comment = &v
	return s
}

func (s *ListHostGroupsForUserResponseBodyHostGroups) SetHostGroupId(v string) *ListHostGroupsForUserResponseBodyHostGroups {
	s.HostGroupId = &v
	return s
}

func (s *ListHostGroupsForUserResponseBodyHostGroups) SetHostGroupName(v string) *ListHostGroupsForUserResponseBodyHostGroups {
	s.HostGroupName = &v
	return s
}

type ListHostGroupsForUserResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListHostGroupsForUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListHostGroupsForUserResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHostGroupsForUserResponse) GoString() string {
	return s.String()
}

func (s *ListHostGroupsForUserResponse) SetHeaders(v map[string]*string) *ListHostGroupsForUserResponse {
	s.Headers = v
	return s
}

func (s *ListHostGroupsForUserResponse) SetBody(v *ListHostGroupsForUserResponseBody) *ListHostGroupsForUserResponse {
	s.Body = v
	return s
}

type ListHostGroupsForUserGroupRequest struct {
	SourceIp      *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Mode          *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	UserGroupId   *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	PageNumber    *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	HostGroupName *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
}

func (s ListHostGroupsForUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHostGroupsForUserGroupRequest) GoString() string {
	return s.String()
}

func (s *ListHostGroupsForUserGroupRequest) SetSourceIp(v string) *ListHostGroupsForUserGroupRequest {
	s.SourceIp = &v
	return s
}

func (s *ListHostGroupsForUserGroupRequest) SetInstanceId(v string) *ListHostGroupsForUserGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHostGroupsForUserGroupRequest) SetRegionId(v string) *ListHostGroupsForUserGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ListHostGroupsForUserGroupRequest) SetMode(v string) *ListHostGroupsForUserGroupRequest {
	s.Mode = &v
	return s
}

func (s *ListHostGroupsForUserGroupRequest) SetUserGroupId(v string) *ListHostGroupsForUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *ListHostGroupsForUserGroupRequest) SetPageNumber(v string) *ListHostGroupsForUserGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHostGroupsForUserGroupRequest) SetPageSize(v string) *ListHostGroupsForUserGroupRequest {
	s.PageSize = &v
	return s
}

func (s *ListHostGroupsForUserGroupRequest) SetHostGroupName(v string) *ListHostGroupsForUserGroupRequest {
	s.HostGroupName = &v
	return s
}

type ListHostGroupsForUserGroupResponseBody struct {
	TotalCount *int32                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HostGroups []*ListHostGroupsForUserGroupResponseBodyHostGroups `json:"HostGroups,omitempty" xml:"HostGroups,omitempty" type:"Repeated"`
}

func (s ListHostGroupsForUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHostGroupsForUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListHostGroupsForUserGroupResponseBody) SetTotalCount(v int32) *ListHostGroupsForUserGroupResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListHostGroupsForUserGroupResponseBody) SetRequestId(v string) *ListHostGroupsForUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHostGroupsForUserGroupResponseBody) SetHostGroups(v []*ListHostGroupsForUserGroupResponseBodyHostGroups) *ListHostGroupsForUserGroupResponseBody {
	s.HostGroups = v
	return s
}

type ListHostGroupsForUserGroupResponseBodyHostGroups struct {
	Comment       *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	HostGroupId   *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	HostGroupName *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
}

func (s ListHostGroupsForUserGroupResponseBodyHostGroups) String() string {
	return tea.Prettify(s)
}

func (s ListHostGroupsForUserGroupResponseBodyHostGroups) GoString() string {
	return s.String()
}

func (s *ListHostGroupsForUserGroupResponseBodyHostGroups) SetComment(v string) *ListHostGroupsForUserGroupResponseBodyHostGroups {
	s.Comment = &v
	return s
}

func (s *ListHostGroupsForUserGroupResponseBodyHostGroups) SetHostGroupId(v string) *ListHostGroupsForUserGroupResponseBodyHostGroups {
	s.HostGroupId = &v
	return s
}

func (s *ListHostGroupsForUserGroupResponseBodyHostGroups) SetHostGroupName(v string) *ListHostGroupsForUserGroupResponseBodyHostGroups {
	s.HostGroupName = &v
	return s
}

type ListHostGroupsForUserGroupResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListHostGroupsForUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListHostGroupsForUserGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHostGroupsForUserGroupResponse) GoString() string {
	return s.String()
}

func (s *ListHostGroupsForUserGroupResponse) SetHeaders(v map[string]*string) *ListHostGroupsForUserGroupResponse {
	s.Headers = v
	return s
}

func (s *ListHostGroupsForUserGroupResponse) SetBody(v *ListHostGroupsForUserGroupResponseBody) *ListHostGroupsForUserGroupResponse {
	s.Body = v
	return s
}

type ListHostsRequest struct {
	SourceIp            *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PageNumber          *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize            *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	OSType              *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	HostName            *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	HostAddress         *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	Source              *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceInstanceId    *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
	SourceInstanceState *string `json:"SourceInstanceState,omitempty" xml:"SourceInstanceState,omitempty"`
	HostGroupId         *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
}

func (s ListHostsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHostsRequest) GoString() string {
	return s.String()
}

func (s *ListHostsRequest) SetSourceIp(v string) *ListHostsRequest {
	s.SourceIp = &v
	return s
}

func (s *ListHostsRequest) SetInstanceId(v string) *ListHostsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHostsRequest) SetRegionId(v string) *ListHostsRequest {
	s.RegionId = &v
	return s
}

func (s *ListHostsRequest) SetPageNumber(v string) *ListHostsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHostsRequest) SetPageSize(v string) *ListHostsRequest {
	s.PageSize = &v
	return s
}

func (s *ListHostsRequest) SetOSType(v string) *ListHostsRequest {
	s.OSType = &v
	return s
}

func (s *ListHostsRequest) SetHostName(v string) *ListHostsRequest {
	s.HostName = &v
	return s
}

func (s *ListHostsRequest) SetHostAddress(v string) *ListHostsRequest {
	s.HostAddress = &v
	return s
}

func (s *ListHostsRequest) SetSource(v string) *ListHostsRequest {
	s.Source = &v
	return s
}

func (s *ListHostsRequest) SetSourceInstanceId(v string) *ListHostsRequest {
	s.SourceInstanceId = &v
	return s
}

func (s *ListHostsRequest) SetSourceInstanceState(v string) *ListHostsRequest {
	s.SourceInstanceState = &v
	return s
}

func (s *ListHostsRequest) SetHostGroupId(v string) *ListHostsRequest {
	s.HostGroupId = &v
	return s
}

type ListHostsResponseBody struct {
	Hosts      []*ListHostsResponseBodyHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
	TotalCount *int32                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListHostsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHostsResponseBody) GoString() string {
	return s.String()
}

func (s *ListHostsResponseBody) SetHosts(v []*ListHostsResponseBodyHosts) *ListHostsResponseBody {
	s.Hosts = v
	return s
}

func (s *ListHostsResponseBody) SetTotalCount(v int32) *ListHostsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListHostsResponseBody) SetRequestId(v string) *ListHostsResponseBody {
	s.RequestId = &v
	return s
}

type ListHostsResponseBodyHosts struct {
	Comment             *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	ActiveAddressType   *string `json:"ActiveAddressType,omitempty" xml:"ActiveAddressType,omitempty"`
	HostPublicAddress   *string `json:"HostPublicAddress,omitempty" xml:"HostPublicAddress,omitempty"`
	HostName            *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	HostAccountCount    *int32  `json:"HostAccountCount,omitempty" xml:"HostAccountCount,omitempty"`
	Source              *string `json:"Source,omitempty" xml:"Source,omitempty"`
	HostPrivateAddress  *string `json:"HostPrivateAddress,omitempty" xml:"HostPrivateAddress,omitempty"`
	OSType              *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	HostId              *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	SourceInstanceState *string `json:"SourceInstanceState,omitempty" xml:"SourceInstanceState,omitempty"`
	SourceInstanceId    *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
}

func (s ListHostsResponseBodyHosts) String() string {
	return tea.Prettify(s)
}

func (s ListHostsResponseBodyHosts) GoString() string {
	return s.String()
}

func (s *ListHostsResponseBodyHosts) SetComment(v string) *ListHostsResponseBodyHosts {
	s.Comment = &v
	return s
}

func (s *ListHostsResponseBodyHosts) SetActiveAddressType(v string) *ListHostsResponseBodyHosts {
	s.ActiveAddressType = &v
	return s
}

func (s *ListHostsResponseBodyHosts) SetHostPublicAddress(v string) *ListHostsResponseBodyHosts {
	s.HostPublicAddress = &v
	return s
}

func (s *ListHostsResponseBodyHosts) SetHostName(v string) *ListHostsResponseBodyHosts {
	s.HostName = &v
	return s
}

func (s *ListHostsResponseBodyHosts) SetHostAccountCount(v int32) *ListHostsResponseBodyHosts {
	s.HostAccountCount = &v
	return s
}

func (s *ListHostsResponseBodyHosts) SetSource(v string) *ListHostsResponseBodyHosts {
	s.Source = &v
	return s
}

func (s *ListHostsResponseBodyHosts) SetHostPrivateAddress(v string) *ListHostsResponseBodyHosts {
	s.HostPrivateAddress = &v
	return s
}

func (s *ListHostsResponseBodyHosts) SetOSType(v string) *ListHostsResponseBodyHosts {
	s.OSType = &v
	return s
}

func (s *ListHostsResponseBodyHosts) SetHostId(v string) *ListHostsResponseBodyHosts {
	s.HostId = &v
	return s
}

func (s *ListHostsResponseBodyHosts) SetSourceInstanceState(v string) *ListHostsResponseBodyHosts {
	s.SourceInstanceState = &v
	return s
}

func (s *ListHostsResponseBodyHosts) SetSourceInstanceId(v string) *ListHostsResponseBodyHosts {
	s.SourceInstanceId = &v
	return s
}

type ListHostsResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListHostsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListHostsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHostsResponse) GoString() string {
	return s.String()
}

func (s *ListHostsResponse) SetHeaders(v map[string]*string) *ListHostsResponse {
	s.Headers = v
	return s
}

func (s *ListHostsResponse) SetBody(v *ListHostsResponseBody) *ListHostsResponse {
	s.Body = v
	return s
}

type ListHostsForUserRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Mode        *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	PageNumber  *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	HostAddress *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	HostName    *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	OSType      *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
}

func (s ListHostsForUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHostsForUserRequest) GoString() string {
	return s.String()
}

func (s *ListHostsForUserRequest) SetSourceIp(v string) *ListHostsForUserRequest {
	s.SourceIp = &v
	return s
}

func (s *ListHostsForUserRequest) SetInstanceId(v string) *ListHostsForUserRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHostsForUserRequest) SetRegionId(v string) *ListHostsForUserRequest {
	s.RegionId = &v
	return s
}

func (s *ListHostsForUserRequest) SetMode(v string) *ListHostsForUserRequest {
	s.Mode = &v
	return s
}

func (s *ListHostsForUserRequest) SetUserId(v string) *ListHostsForUserRequest {
	s.UserId = &v
	return s
}

func (s *ListHostsForUserRequest) SetPageNumber(v string) *ListHostsForUserRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHostsForUserRequest) SetPageSize(v string) *ListHostsForUserRequest {
	s.PageSize = &v
	return s
}

func (s *ListHostsForUserRequest) SetHostAddress(v string) *ListHostsForUserRequest {
	s.HostAddress = &v
	return s
}

func (s *ListHostsForUserRequest) SetHostName(v string) *ListHostsForUserRequest {
	s.HostName = &v
	return s
}

func (s *ListHostsForUserRequest) SetOSType(v string) *ListHostsForUserRequest {
	s.OSType = &v
	return s
}

type ListHostsForUserResponseBody struct {
	Hosts      []*ListHostsForUserResponseBodyHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
	TotalCount *int32                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListHostsForUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHostsForUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListHostsForUserResponseBody) SetHosts(v []*ListHostsForUserResponseBodyHosts) *ListHostsForUserResponseBody {
	s.Hosts = v
	return s
}

func (s *ListHostsForUserResponseBody) SetTotalCount(v int32) *ListHostsForUserResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListHostsForUserResponseBody) SetRequestId(v string) *ListHostsForUserResponseBody {
	s.RequestId = &v
	return s
}

type ListHostsForUserResponseBodyHosts struct {
	Comment            *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	ActiveAddressType  *string `json:"ActiveAddressType,omitempty" xml:"ActiveAddressType,omitempty"`
	HostPublicAddress  *string `json:"HostPublicAddress,omitempty" xml:"HostPublicAddress,omitempty"`
	HostName           *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	HostPrivateAddress *string `json:"HostPrivateAddress,omitempty" xml:"HostPrivateAddress,omitempty"`
	OSType             *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	HostId             *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
}

func (s ListHostsForUserResponseBodyHosts) String() string {
	return tea.Prettify(s)
}

func (s ListHostsForUserResponseBodyHosts) GoString() string {
	return s.String()
}

func (s *ListHostsForUserResponseBodyHosts) SetComment(v string) *ListHostsForUserResponseBodyHosts {
	s.Comment = &v
	return s
}

func (s *ListHostsForUserResponseBodyHosts) SetActiveAddressType(v string) *ListHostsForUserResponseBodyHosts {
	s.ActiveAddressType = &v
	return s
}

func (s *ListHostsForUserResponseBodyHosts) SetHostPublicAddress(v string) *ListHostsForUserResponseBodyHosts {
	s.HostPublicAddress = &v
	return s
}

func (s *ListHostsForUserResponseBodyHosts) SetHostName(v string) *ListHostsForUserResponseBodyHosts {
	s.HostName = &v
	return s
}

func (s *ListHostsForUserResponseBodyHosts) SetHostPrivateAddress(v string) *ListHostsForUserResponseBodyHosts {
	s.HostPrivateAddress = &v
	return s
}

func (s *ListHostsForUserResponseBodyHosts) SetOSType(v string) *ListHostsForUserResponseBodyHosts {
	s.OSType = &v
	return s
}

func (s *ListHostsForUserResponseBodyHosts) SetHostId(v string) *ListHostsForUserResponseBodyHosts {
	s.HostId = &v
	return s
}

type ListHostsForUserResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListHostsForUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListHostsForUserResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHostsForUserResponse) GoString() string {
	return s.String()
}

func (s *ListHostsForUserResponse) SetHeaders(v map[string]*string) *ListHostsForUserResponse {
	s.Headers = v
	return s
}

func (s *ListHostsForUserResponse) SetBody(v *ListHostsForUserResponseBody) *ListHostsForUserResponse {
	s.Body = v
	return s
}

type ListHostsForUserGroupRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Mode        *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	PageNumber  *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	HostAddress *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	HostName    *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	OSType      *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
}

func (s ListHostsForUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHostsForUserGroupRequest) GoString() string {
	return s.String()
}

func (s *ListHostsForUserGroupRequest) SetSourceIp(v string) *ListHostsForUserGroupRequest {
	s.SourceIp = &v
	return s
}

func (s *ListHostsForUserGroupRequest) SetInstanceId(v string) *ListHostsForUserGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHostsForUserGroupRequest) SetRegionId(v string) *ListHostsForUserGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ListHostsForUserGroupRequest) SetMode(v string) *ListHostsForUserGroupRequest {
	s.Mode = &v
	return s
}

func (s *ListHostsForUserGroupRequest) SetUserGroupId(v string) *ListHostsForUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *ListHostsForUserGroupRequest) SetPageNumber(v string) *ListHostsForUserGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHostsForUserGroupRequest) SetPageSize(v string) *ListHostsForUserGroupRequest {
	s.PageSize = &v
	return s
}

func (s *ListHostsForUserGroupRequest) SetHostAddress(v string) *ListHostsForUserGroupRequest {
	s.HostAddress = &v
	return s
}

func (s *ListHostsForUserGroupRequest) SetHostName(v string) *ListHostsForUserGroupRequest {
	s.HostName = &v
	return s
}

func (s *ListHostsForUserGroupRequest) SetOSType(v string) *ListHostsForUserGroupRequest {
	s.OSType = &v
	return s
}

type ListHostsForUserGroupResponseBody struct {
	Hosts      []*ListHostsForUserGroupResponseBodyHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
	TotalCount *int32                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListHostsForUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHostsForUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListHostsForUserGroupResponseBody) SetHosts(v []*ListHostsForUserGroupResponseBodyHosts) *ListHostsForUserGroupResponseBody {
	s.Hosts = v
	return s
}

func (s *ListHostsForUserGroupResponseBody) SetTotalCount(v int32) *ListHostsForUserGroupResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListHostsForUserGroupResponseBody) SetRequestId(v string) *ListHostsForUserGroupResponseBody {
	s.RequestId = &v
	return s
}

type ListHostsForUserGroupResponseBodyHosts struct {
	Comment            *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	ActiveAddressType  *string `json:"ActiveAddressType,omitempty" xml:"ActiveAddressType,omitempty"`
	HostPublicAddress  *string `json:"HostPublicAddress,omitempty" xml:"HostPublicAddress,omitempty"`
	HostName           *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	HostPrivateAddress *string `json:"HostPrivateAddress,omitempty" xml:"HostPrivateAddress,omitempty"`
	OSType             *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	HostId             *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
}

func (s ListHostsForUserGroupResponseBodyHosts) String() string {
	return tea.Prettify(s)
}

func (s ListHostsForUserGroupResponseBodyHosts) GoString() string {
	return s.String()
}

func (s *ListHostsForUserGroupResponseBodyHosts) SetComment(v string) *ListHostsForUserGroupResponseBodyHosts {
	s.Comment = &v
	return s
}

func (s *ListHostsForUserGroupResponseBodyHosts) SetActiveAddressType(v string) *ListHostsForUserGroupResponseBodyHosts {
	s.ActiveAddressType = &v
	return s
}

func (s *ListHostsForUserGroupResponseBodyHosts) SetHostPublicAddress(v string) *ListHostsForUserGroupResponseBodyHosts {
	s.HostPublicAddress = &v
	return s
}

func (s *ListHostsForUserGroupResponseBodyHosts) SetHostName(v string) *ListHostsForUserGroupResponseBodyHosts {
	s.HostName = &v
	return s
}

func (s *ListHostsForUserGroupResponseBodyHosts) SetHostPrivateAddress(v string) *ListHostsForUserGroupResponseBodyHosts {
	s.HostPrivateAddress = &v
	return s
}

func (s *ListHostsForUserGroupResponseBodyHosts) SetOSType(v string) *ListHostsForUserGroupResponseBodyHosts {
	s.OSType = &v
	return s
}

func (s *ListHostsForUserGroupResponseBodyHosts) SetHostId(v string) *ListHostsForUserGroupResponseBodyHosts {
	s.HostId = &v
	return s
}

type ListHostsForUserGroupResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListHostsForUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListHostsForUserGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHostsForUserGroupResponse) GoString() string {
	return s.String()
}

func (s *ListHostsForUserGroupResponse) SetHeaders(v map[string]*string) *ListHostsForUserGroupResponse {
	s.Headers = v
	return s
}

func (s *ListHostsForUserGroupResponse) SetBody(v *ListHostsForUserGroupResponseBody) *ListHostsForUserGroupResponse {
	s.Body = v
	return s
}

type ListTagKeysRequest struct {
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) SetSourceIp(v string) *ListTagKeysRequest {
	s.SourceIp = &v
	return s
}

func (s *ListTagKeysRequest) SetLang(v string) *ListTagKeysRequest {
	s.Lang = &v
	return s
}

func (s *ListTagKeysRequest) SetRegionId(v string) *ListTagKeysRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceType(v string) *ListTagKeysRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagKeysRequest) SetPageSize(v int32) *ListTagKeysRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagKeysRequest) SetPageNumber(v int32) *ListTagKeysRequest {
	s.PageNumber = &v
	return s
}

type ListTagKeysResponseBody struct {
	TotalCount *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize   *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber *int32                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TagKeys    []*ListTagKeysResponseBodyTagKeys `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
}

func (s ListTagKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBody) SetTotalCount(v int32) *ListTagKeysResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTagKeysResponseBody) SetRequestId(v string) *ListTagKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagKeysResponseBody) SetPageSize(v int32) *ListTagKeysResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTagKeysResponseBody) SetPageNumber(v int32) *ListTagKeysResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTagKeysResponseBody) SetTagKeys(v []*ListTagKeysResponseBodyTagKeys) *ListTagKeysResponseBody {
	s.TagKeys = v
	return s
}

type ListTagKeysResponseBodyTagKeys struct {
	TagCount *int32  `json:"TagCount,omitempty" xml:"TagCount,omitempty"`
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagKeysResponseBodyTagKeys) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBodyTagKeys) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBodyTagKeys) SetTagCount(v int32) *ListTagKeysResponseBodyTagKeys {
	s.TagCount = &v
	return s
}

func (s *ListTagKeysResponseBodyTagKeys) SetTagKey(v string) *ListTagKeysResponseBodyTagKeys {
	s.TagKey = &v
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

type ListTagResourcesRequest struct {
	SourceIp     *string                       `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang         *string                       `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RegionId     *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	NextToken    *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceId   []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag          []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetSourceIp(v string) *ListTagResourcesRequest {
	s.SourceIp = &v
	return s
}

func (s *ListTagResourcesRequest) SetLang(v string) *ListTagResourcesRequest {
	s.Lang = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	NextToken    *string                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagValue = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagKey = &v
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

type ListUserGroupsRequest struct {
	SourceIp      *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PageNumber    *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
}

func (s ListUserGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListUserGroupsRequest) SetSourceIp(v string) *ListUserGroupsRequest {
	s.SourceIp = &v
	return s
}

func (s *ListUserGroupsRequest) SetInstanceId(v string) *ListUserGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListUserGroupsRequest) SetRegionId(v string) *ListUserGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *ListUserGroupsRequest) SetPageNumber(v string) *ListUserGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUserGroupsRequest) SetPageSize(v string) *ListUserGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserGroupsRequest) SetUserGroupName(v string) *ListUserGroupsRequest {
	s.UserGroupName = &v
	return s
}

type ListUserGroupsResponseBody struct {
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserGroups []*ListUserGroupsResponseBodyUserGroups `json:"UserGroups,omitempty" xml:"UserGroups,omitempty" type:"Repeated"`
}

func (s ListUserGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserGroupsResponseBody) SetTotalCount(v int32) *ListUserGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUserGroupsResponseBody) SetRequestId(v string) *ListUserGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserGroupsResponseBody) SetUserGroups(v []*ListUserGroupsResponseBodyUserGroups) *ListUserGroupsResponseBody {
	s.UserGroups = v
	return s
}

type ListUserGroupsResponseBodyUserGroups struct {
	MemberCount   *int32  `json:"MemberCount,omitempty" xml:"MemberCount,omitempty"`
	Comment       *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	UserGroupId   *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
}

func (s ListUserGroupsResponseBodyUserGroups) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsResponseBodyUserGroups) GoString() string {
	return s.String()
}

func (s *ListUserGroupsResponseBodyUserGroups) SetMemberCount(v int32) *ListUserGroupsResponseBodyUserGroups {
	s.MemberCount = &v
	return s
}

func (s *ListUserGroupsResponseBodyUserGroups) SetComment(v string) *ListUserGroupsResponseBodyUserGroups {
	s.Comment = &v
	return s
}

func (s *ListUserGroupsResponseBodyUserGroups) SetUserGroupId(v string) *ListUserGroupsResponseBodyUserGroups {
	s.UserGroupId = &v
	return s
}

func (s *ListUserGroupsResponseBodyUserGroups) SetUserGroupName(v string) *ListUserGroupsResponseBodyUserGroups {
	s.UserGroupName = &v
	return s
}

type ListUserGroupsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUserGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUserGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListUserGroupsResponse) SetHeaders(v map[string]*string) *ListUserGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListUserGroupsResponse) SetBody(v *ListUserGroupsResponseBody) *ListUserGroupsResponse {
	s.Body = v
	return s
}

type ListUsersRequest struct {
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PageNumber   *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	DisplayName  *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Source       *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Mobile       *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	UserState    *string `json:"UserState,omitempty" xml:"UserState,omitempty"`
	SourceUserId *string `json:"SourceUserId,omitempty" xml:"SourceUserId,omitempty"`
	UserGroupId  *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s ListUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) SetSourceIp(v string) *ListUsersRequest {
	s.SourceIp = &v
	return s
}

func (s *ListUsersRequest) SetInstanceId(v string) *ListUsersRequest {
	s.InstanceId = &v
	return s
}

func (s *ListUsersRequest) SetRegionId(v string) *ListUsersRequest {
	s.RegionId = &v
	return s
}

func (s *ListUsersRequest) SetPageNumber(v string) *ListUsersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUsersRequest) SetPageSize(v string) *ListUsersRequest {
	s.PageSize = &v
	return s
}

func (s *ListUsersRequest) SetUserName(v string) *ListUsersRequest {
	s.UserName = &v
	return s
}

func (s *ListUsersRequest) SetDisplayName(v string) *ListUsersRequest {
	s.DisplayName = &v
	return s
}

func (s *ListUsersRequest) SetSource(v string) *ListUsersRequest {
	s.Source = &v
	return s
}

func (s *ListUsersRequest) SetMobile(v string) *ListUsersRequest {
	s.Mobile = &v
	return s
}

func (s *ListUsersRequest) SetUserState(v string) *ListUsersRequest {
	s.UserState = &v
	return s
}

func (s *ListUsersRequest) SetSourceUserId(v string) *ListUsersRequest {
	s.SourceUserId = &v
	return s
}

func (s *ListUsersRequest) SetUserGroupId(v string) *ListUsersRequest {
	s.UserGroupId = &v
	return s
}

type ListUsersResponseBody struct {
	TotalCount *int32                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Users      []*ListUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) SetTotalCount(v int32) *ListUsersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersResponseBody) SetUsers(v []*ListUsersResponseBodyUsers) *ListUsersResponseBody {
	s.Users = v
	return s
}

type ListUsersResponseBodyUsers struct {
	DisplayName       *string   `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Email             *string   `json:"Email,omitempty" xml:"Email,omitempty"`
	Comment           *string   `json:"Comment,omitempty" xml:"Comment,omitempty"`
	MobileCountryCode *string   `json:"MobileCountryCode,omitempty" xml:"MobileCountryCode,omitempty"`
	Mobile            *string   `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	UserId            *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Source            *string   `json:"Source,omitempty" xml:"Source,omitempty"`
	UserName          *string   `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserState         []*string `json:"UserState,omitempty" xml:"UserState,omitempty" type:"Repeated"`
	SourceUserId      *string   `json:"SourceUserId,omitempty" xml:"SourceUserId,omitempty"`
}

func (s ListUsersResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUsers) SetDisplayName(v string) *ListUsersResponseBodyUsers {
	s.DisplayName = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetEmail(v string) *ListUsersResponseBodyUsers {
	s.Email = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetComment(v string) *ListUsersResponseBodyUsers {
	s.Comment = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetMobileCountryCode(v string) *ListUsersResponseBodyUsers {
	s.MobileCountryCode = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetMobile(v string) *ListUsersResponseBodyUsers {
	s.Mobile = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetUserId(v string) *ListUsersResponseBodyUsers {
	s.UserId = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetSource(v string) *ListUsersResponseBodyUsers {
	s.Source = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetUserName(v string) *ListUsersResponseBodyUsers {
	s.UserName = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetUserState(v []*string) *ListUsersResponseBodyUsers {
	s.UserState = v
	return s
}

func (s *ListUsersResponseBodyUsers) SetSourceUserId(v string) *ListUsersResponseBodyUsers {
	s.SourceUserId = &v
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

type LockUsersRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserIds    *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
}

func (s LockUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s LockUsersRequest) GoString() string {
	return s.String()
}

func (s *LockUsersRequest) SetSourceIp(v string) *LockUsersRequest {
	s.SourceIp = &v
	return s
}

func (s *LockUsersRequest) SetInstanceId(v string) *LockUsersRequest {
	s.InstanceId = &v
	return s
}

func (s *LockUsersRequest) SetRegionId(v string) *LockUsersRequest {
	s.RegionId = &v
	return s
}

func (s *LockUsersRequest) SetUserIds(v string) *LockUsersRequest {
	s.UserIds = &v
	return s
}

type LockUsersResponseBody struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*LockUsersResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s LockUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LockUsersResponseBody) GoString() string {
	return s.String()
}

func (s *LockUsersResponseBody) SetRequestId(v string) *LockUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *LockUsersResponseBody) SetResults(v []*LockUsersResponseBodyResults) *LockUsersResponseBody {
	s.Results = v
	return s
}

type LockUsersResponseBodyResults struct {
	UserId  *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s LockUsersResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s LockUsersResponseBodyResults) GoString() string {
	return s.String()
}

func (s *LockUsersResponseBodyResults) SetUserId(v string) *LockUsersResponseBodyResults {
	s.UserId = &v
	return s
}

func (s *LockUsersResponseBodyResults) SetCode(v string) *LockUsersResponseBodyResults {
	s.Code = &v
	return s
}

func (s *LockUsersResponseBodyResults) SetMessage(v string) *LockUsersResponseBodyResults {
	s.Message = &v
	return s
}

type LockUsersResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *LockUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LockUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s LockUsersResponse) GoString() string {
	return s.String()
}

func (s *LockUsersResponse) SetHeaders(v map[string]*string) *LockUsersResponse {
	s.Headers = v
	return s
}

func (s *LockUsersResponse) SetBody(v *LockUsersResponseBody) *LockUsersResponse {
	s.Body = v
	return s
}

type ModifyHostRequest struct {
	SourceIp           *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	HostId             *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	HostPrivateAddress *string `json:"HostPrivateAddress,omitempty" xml:"HostPrivateAddress,omitempty"`
	HostPublicAddress  *string `json:"HostPublicAddress,omitempty" xml:"HostPublicAddress,omitempty"`
	OSType             *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	HostName           *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	Comment            *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
}

func (s ModifyHostRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyHostRequest) GoString() string {
	return s.String()
}

func (s *ModifyHostRequest) SetSourceIp(v string) *ModifyHostRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyHostRequest) SetInstanceId(v string) *ModifyHostRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyHostRequest) SetRegionId(v string) *ModifyHostRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyHostRequest) SetHostId(v string) *ModifyHostRequest {
	s.HostId = &v
	return s
}

func (s *ModifyHostRequest) SetHostPrivateAddress(v string) *ModifyHostRequest {
	s.HostPrivateAddress = &v
	return s
}

func (s *ModifyHostRequest) SetHostPublicAddress(v string) *ModifyHostRequest {
	s.HostPublicAddress = &v
	return s
}

func (s *ModifyHostRequest) SetOSType(v string) *ModifyHostRequest {
	s.OSType = &v
	return s
}

func (s *ModifyHostRequest) SetHostName(v string) *ModifyHostRequest {
	s.HostName = &v
	return s
}

func (s *ModifyHostRequest) SetComment(v string) *ModifyHostRequest {
	s.Comment = &v
	return s
}

type ModifyHostResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyHostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyHostResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHostResponseBody) SetRequestId(v string) *ModifyHostResponseBody {
	s.RequestId = &v
	return s
}

type ModifyHostResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyHostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyHostResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyHostResponse) GoString() string {
	return s.String()
}

func (s *ModifyHostResponse) SetHeaders(v map[string]*string) *ModifyHostResponse {
	s.Headers = v
	return s
}

func (s *ModifyHostResponse) SetBody(v *ModifyHostResponseBody) *ModifyHostResponse {
	s.Body = v
	return s
}

type ModifyHostAccountRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	HostAccountId   *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	Password        *string `json:"Password,omitempty" xml:"Password,omitempty"`
	PrivateKey      *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	PassPhrase      *string `json:"PassPhrase,omitempty" xml:"PassPhrase,omitempty"`
}

func (s ModifyHostAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyHostAccountRequest) GoString() string {
	return s.String()
}

func (s *ModifyHostAccountRequest) SetSourceIp(v string) *ModifyHostAccountRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyHostAccountRequest) SetInstanceId(v string) *ModifyHostAccountRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyHostAccountRequest) SetRegionId(v string) *ModifyHostAccountRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyHostAccountRequest) SetHostAccountId(v string) *ModifyHostAccountRequest {
	s.HostAccountId = &v
	return s
}

func (s *ModifyHostAccountRequest) SetHostAccountName(v string) *ModifyHostAccountRequest {
	s.HostAccountName = &v
	return s
}

func (s *ModifyHostAccountRequest) SetPassword(v string) *ModifyHostAccountRequest {
	s.Password = &v
	return s
}

func (s *ModifyHostAccountRequest) SetPrivateKey(v string) *ModifyHostAccountRequest {
	s.PrivateKey = &v
	return s
}

func (s *ModifyHostAccountRequest) SetPassPhrase(v string) *ModifyHostAccountRequest {
	s.PassPhrase = &v
	return s
}

type ModifyHostAccountResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyHostAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyHostAccountResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHostAccountResponseBody) SetRequestId(v string) *ModifyHostAccountResponseBody {
	s.RequestId = &v
	return s
}

type ModifyHostAccountResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyHostAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyHostAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyHostAccountResponse) GoString() string {
	return s.String()
}

func (s *ModifyHostAccountResponse) SetHeaders(v map[string]*string) *ModifyHostAccountResponse {
	s.Headers = v
	return s
}

func (s *ModifyHostAccountResponse) SetBody(v *ModifyHostAccountResponseBody) *ModifyHostAccountResponse {
	s.Body = v
	return s
}

type ModifyHostGroupRequest struct {
	SourceIp      *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	HostGroupId   *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	HostGroupName *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
	Comment       *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
}

func (s ModifyHostGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyHostGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyHostGroupRequest) SetSourceIp(v string) *ModifyHostGroupRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyHostGroupRequest) SetInstanceId(v string) *ModifyHostGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyHostGroupRequest) SetRegionId(v string) *ModifyHostGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyHostGroupRequest) SetHostGroupId(v string) *ModifyHostGroupRequest {
	s.HostGroupId = &v
	return s
}

func (s *ModifyHostGroupRequest) SetHostGroupName(v string) *ModifyHostGroupRequest {
	s.HostGroupName = &v
	return s
}

func (s *ModifyHostGroupRequest) SetComment(v string) *ModifyHostGroupRequest {
	s.Comment = &v
	return s
}

type ModifyHostGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyHostGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyHostGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHostGroupResponseBody) SetRequestId(v string) *ModifyHostGroupResponseBody {
	s.RequestId = &v
	return s
}

type ModifyHostGroupResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyHostGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyHostGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyHostGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyHostGroupResponse) SetHeaders(v map[string]*string) *ModifyHostGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyHostGroupResponse) SetBody(v *ModifyHostGroupResponseBody) *ModifyHostGroupResponse {
	s.Body = v
	return s
}

type ModifyHostsActiveAddressTypeRequest struct {
	SourceIp          *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	HostIds           *string `json:"HostIds,omitempty" xml:"HostIds,omitempty"`
	ActiveAddressType *string `json:"ActiveAddressType,omitempty" xml:"ActiveAddressType,omitempty"`
}

func (s ModifyHostsActiveAddressTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyHostsActiveAddressTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyHostsActiveAddressTypeRequest) SetSourceIp(v string) *ModifyHostsActiveAddressTypeRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyHostsActiveAddressTypeRequest) SetInstanceId(v string) *ModifyHostsActiveAddressTypeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyHostsActiveAddressTypeRequest) SetRegionId(v string) *ModifyHostsActiveAddressTypeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyHostsActiveAddressTypeRequest) SetHostIds(v string) *ModifyHostsActiveAddressTypeRequest {
	s.HostIds = &v
	return s
}

func (s *ModifyHostsActiveAddressTypeRequest) SetActiveAddressType(v string) *ModifyHostsActiveAddressTypeRequest {
	s.ActiveAddressType = &v
	return s
}

type ModifyHostsActiveAddressTypeResponseBody struct {
	RequestId *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*ModifyHostsActiveAddressTypeResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s ModifyHostsActiveAddressTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyHostsActiveAddressTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHostsActiveAddressTypeResponseBody) SetRequestId(v string) *ModifyHostsActiveAddressTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHostsActiveAddressTypeResponseBody) SetResults(v []*ModifyHostsActiveAddressTypeResponseBodyResults) *ModifyHostsActiveAddressTypeResponseBody {
	s.Results = v
	return s
}

type ModifyHostsActiveAddressTypeResponseBodyResults struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	HostId  *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
}

func (s ModifyHostsActiveAddressTypeResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s ModifyHostsActiveAddressTypeResponseBodyResults) GoString() string {
	return s.String()
}

func (s *ModifyHostsActiveAddressTypeResponseBodyResults) SetCode(v string) *ModifyHostsActiveAddressTypeResponseBodyResults {
	s.Code = &v
	return s
}

func (s *ModifyHostsActiveAddressTypeResponseBodyResults) SetMessage(v string) *ModifyHostsActiveAddressTypeResponseBodyResults {
	s.Message = &v
	return s
}

func (s *ModifyHostsActiveAddressTypeResponseBodyResults) SetHostId(v string) *ModifyHostsActiveAddressTypeResponseBodyResults {
	s.HostId = &v
	return s
}

type ModifyHostsActiveAddressTypeResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyHostsActiveAddressTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyHostsActiveAddressTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyHostsActiveAddressTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyHostsActiveAddressTypeResponse) SetHeaders(v map[string]*string) *ModifyHostsActiveAddressTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyHostsActiveAddressTypeResponse) SetBody(v *ModifyHostsActiveAddressTypeResponseBody) *ModifyHostsActiveAddressTypeResponse {
	s.Body = v
	return s
}

type ModifyHostsPortRequest struct {
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	HostIds      *string `json:"HostIds,omitempty" xml:"HostIds,omitempty"`
	ProtocolName *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
	Port         *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s ModifyHostsPortRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyHostsPortRequest) GoString() string {
	return s.String()
}

func (s *ModifyHostsPortRequest) SetSourceIp(v string) *ModifyHostsPortRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyHostsPortRequest) SetInstanceId(v string) *ModifyHostsPortRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyHostsPortRequest) SetRegionId(v string) *ModifyHostsPortRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyHostsPortRequest) SetHostIds(v string) *ModifyHostsPortRequest {
	s.HostIds = &v
	return s
}

func (s *ModifyHostsPortRequest) SetProtocolName(v string) *ModifyHostsPortRequest {
	s.ProtocolName = &v
	return s
}

func (s *ModifyHostsPortRequest) SetPort(v string) *ModifyHostsPortRequest {
	s.Port = &v
	return s
}

type ModifyHostsPortResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*ModifyHostsPortResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s ModifyHostsPortResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyHostsPortResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHostsPortResponseBody) SetRequestId(v string) *ModifyHostsPortResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHostsPortResponseBody) SetResults(v []*ModifyHostsPortResponseBodyResults) *ModifyHostsPortResponseBody {
	s.Results = v
	return s
}

type ModifyHostsPortResponseBodyResults struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	HostId  *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
}

func (s ModifyHostsPortResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s ModifyHostsPortResponseBodyResults) GoString() string {
	return s.String()
}

func (s *ModifyHostsPortResponseBodyResults) SetCode(v string) *ModifyHostsPortResponseBodyResults {
	s.Code = &v
	return s
}

func (s *ModifyHostsPortResponseBodyResults) SetMessage(v string) *ModifyHostsPortResponseBodyResults {
	s.Message = &v
	return s
}

func (s *ModifyHostsPortResponseBodyResults) SetHostId(v string) *ModifyHostsPortResponseBodyResults {
	s.HostId = &v
	return s
}

type ModifyHostsPortResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyHostsPortResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyHostsPortResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyHostsPortResponse) GoString() string {
	return s.String()
}

func (s *ModifyHostsPortResponse) SetHeaders(v map[string]*string) *ModifyHostsPortResponse {
	s.Headers = v
	return s
}

func (s *ModifyHostsPortResponse) SetBody(v *ModifyHostsPortResponseBody) *ModifyHostsPortResponse {
	s.Body = v
	return s
}

type ModifyInstanceAttributeRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyInstanceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttributeRequest) SetSourceIp(v string) *ModifyInstanceAttributeRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetLang(v string) *ModifyInstanceAttributeRequest {
	s.Lang = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetInstanceId(v string) *ModifyInstanceAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetDescription(v string) *ModifyInstanceAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetRegionId(v string) *ModifyInstanceAttributeRequest {
	s.RegionId = &v
	return s
}

type ModifyInstanceAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttributeResponseBody) SetRequestId(v string) *ModifyInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyInstanceAttributeResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyInstanceAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttributeResponse) SetHeaders(v map[string]*string) *ModifyInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceAttributeResponse) SetBody(v *ModifyInstanceAttributeResponseBody) *ModifyInstanceAttributeResponse {
	s.Body = v
	return s
}

type ModifyUserRequest struct {
	SourceIp          *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserId            *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Password          *string `json:"Password,omitempty" xml:"Password,omitempty"`
	DisplayName       *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Comment           *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	Email             *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Mobile            *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	MobileCountryCode *string `json:"MobileCountryCode,omitempty" xml:"MobileCountryCode,omitempty"`
}

func (s ModifyUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserRequest) SetSourceIp(v string) *ModifyUserRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyUserRequest) SetInstanceId(v string) *ModifyUserRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyUserRequest) SetRegionId(v string) *ModifyUserRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyUserRequest) SetUserId(v string) *ModifyUserRequest {
	s.UserId = &v
	return s
}

func (s *ModifyUserRequest) SetPassword(v string) *ModifyUserRequest {
	s.Password = &v
	return s
}

func (s *ModifyUserRequest) SetDisplayName(v string) *ModifyUserRequest {
	s.DisplayName = &v
	return s
}

func (s *ModifyUserRequest) SetComment(v string) *ModifyUserRequest {
	s.Comment = &v
	return s
}

func (s *ModifyUserRequest) SetEmail(v string) *ModifyUserRequest {
	s.Email = &v
	return s
}

func (s *ModifyUserRequest) SetMobile(v string) *ModifyUserRequest {
	s.Mobile = &v
	return s
}

func (s *ModifyUserRequest) SetMobileCountryCode(v string) *ModifyUserRequest {
	s.MobileCountryCode = &v
	return s
}

type ModifyUserResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUserResponseBody) SetRequestId(v string) *ModifyUserResponseBody {
	s.RequestId = &v
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

type ModifyUserGroupRequest struct {
	SourceIp      *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserGroupId   *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
	Comment       *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
}

func (s ModifyUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserGroupRequest) SetSourceIp(v string) *ModifyUserGroupRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyUserGroupRequest) SetInstanceId(v string) *ModifyUserGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyUserGroupRequest) SetRegionId(v string) *ModifyUserGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyUserGroupRequest) SetUserGroupId(v string) *ModifyUserGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *ModifyUserGroupRequest) SetUserGroupName(v string) *ModifyUserGroupRequest {
	s.UserGroupName = &v
	return s
}

func (s *ModifyUserGroupRequest) SetComment(v string) *ModifyUserGroupRequest {
	s.Comment = &v
	return s
}

type ModifyUserGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUserGroupResponseBody) SetRequestId(v string) *ModifyUserGroupResponseBody {
	s.RequestId = &v
	return s
}

type ModifyUserGroupResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyUserGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyUserGroupResponse) SetHeaders(v map[string]*string) *ModifyUserGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyUserGroupResponse) SetBody(v *ModifyUserGroupResponseBody) *ModifyUserGroupResponse {
	s.Body = v
	return s
}

type MoveResourceGroupRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ResourceId      *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceType    *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s MoveResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupRequest) SetSourceIp(v string) *MoveResourceGroupRequest {
	s.SourceIp = &v
	return s
}

func (s *MoveResourceGroupRequest) SetLang(v string) *MoveResourceGroupRequest {
	s.Lang = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceId(v string) *MoveResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceGroupId(v string) *MoveResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceType(v string) *MoveResourceGroupRequest {
	s.ResourceType = &v
	return s
}

func (s *MoveResourceGroupRequest) SetRegionId(v string) *MoveResourceGroupRequest {
	s.RegionId = &v
	return s
}

type MoveResourceGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MoveResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponseBody) SetRequestId(v string) *MoveResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type MoveResourceGroupResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MoveResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MoveResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponse) SetHeaders(v map[string]*string) *MoveResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *MoveResourceGroupResponse) SetBody(v *MoveResourceGroupResponseBody) *MoveResourceGroupResponse {
	s.Body = v
	return s
}

type RemoveHostsFromGroupRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	HostIds     *string `json:"HostIds,omitempty" xml:"HostIds,omitempty"`
}

func (s RemoveHostsFromGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveHostsFromGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveHostsFromGroupRequest) SetSourceIp(v string) *RemoveHostsFromGroupRequest {
	s.SourceIp = &v
	return s
}

func (s *RemoveHostsFromGroupRequest) SetInstanceId(v string) *RemoveHostsFromGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *RemoveHostsFromGroupRequest) SetRegionId(v string) *RemoveHostsFromGroupRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveHostsFromGroupRequest) SetHostGroupId(v string) *RemoveHostsFromGroupRequest {
	s.HostGroupId = &v
	return s
}

func (s *RemoveHostsFromGroupRequest) SetHostIds(v string) *RemoveHostsFromGroupRequest {
	s.HostIds = &v
	return s
}

type RemoveHostsFromGroupResponseBody struct {
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*RemoveHostsFromGroupResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s RemoveHostsFromGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveHostsFromGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveHostsFromGroupResponseBody) SetRequestId(v string) *RemoveHostsFromGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveHostsFromGroupResponseBody) SetResults(v []*RemoveHostsFromGroupResponseBodyResults) *RemoveHostsFromGroupResponseBody {
	s.Results = v
	return s
}

type RemoveHostsFromGroupResponseBodyResults struct {
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty"`
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	HostId      *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
}

func (s RemoveHostsFromGroupResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s RemoveHostsFromGroupResponseBodyResults) GoString() string {
	return s.String()
}

func (s *RemoveHostsFromGroupResponseBodyResults) SetCode(v string) *RemoveHostsFromGroupResponseBodyResults {
	s.Code = &v
	return s
}

func (s *RemoveHostsFromGroupResponseBodyResults) SetMessage(v string) *RemoveHostsFromGroupResponseBodyResults {
	s.Message = &v
	return s
}

func (s *RemoveHostsFromGroupResponseBodyResults) SetHostGroupId(v string) *RemoveHostsFromGroupResponseBodyResults {
	s.HostGroupId = &v
	return s
}

func (s *RemoveHostsFromGroupResponseBodyResults) SetHostId(v string) *RemoveHostsFromGroupResponseBodyResults {
	s.HostId = &v
	return s
}

type RemoveHostsFromGroupResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveHostsFromGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveHostsFromGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveHostsFromGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveHostsFromGroupResponse) SetHeaders(v map[string]*string) *RemoveHostsFromGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveHostsFromGroupResponse) SetBody(v *RemoveHostsFromGroupResponseBody) *RemoveHostsFromGroupResponse {
	s.Body = v
	return s
}

type RemoveUsersFromGroupRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	UserIds     *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
}

func (s RemoveUsersFromGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveUsersFromGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveUsersFromGroupRequest) SetSourceIp(v string) *RemoveUsersFromGroupRequest {
	s.SourceIp = &v
	return s
}

func (s *RemoveUsersFromGroupRequest) SetInstanceId(v string) *RemoveUsersFromGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *RemoveUsersFromGroupRequest) SetRegionId(v string) *RemoveUsersFromGroupRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveUsersFromGroupRequest) SetUserGroupId(v string) *RemoveUsersFromGroupRequest {
	s.UserGroupId = &v
	return s
}

func (s *RemoveUsersFromGroupRequest) SetUserIds(v string) *RemoveUsersFromGroupRequest {
	s.UserIds = &v
	return s
}

type RemoveUsersFromGroupResponseBody struct {
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*RemoveUsersFromGroupResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s RemoveUsersFromGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveUsersFromGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveUsersFromGroupResponseBody) SetRequestId(v string) *RemoveUsersFromGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveUsersFromGroupResponseBody) SetResults(v []*RemoveUsersFromGroupResponseBodyResults) *RemoveUsersFromGroupResponseBody {
	s.Results = v
	return s
}

type RemoveUsersFromGroupResponseBodyResults struct {
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RemoveUsersFromGroupResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s RemoveUsersFromGroupResponseBodyResults) GoString() string {
	return s.String()
}

func (s *RemoveUsersFromGroupResponseBodyResults) SetUserGroupId(v string) *RemoveUsersFromGroupResponseBodyResults {
	s.UserGroupId = &v
	return s
}

func (s *RemoveUsersFromGroupResponseBodyResults) SetUserId(v string) *RemoveUsersFromGroupResponseBodyResults {
	s.UserId = &v
	return s
}

func (s *RemoveUsersFromGroupResponseBodyResults) SetCode(v string) *RemoveUsersFromGroupResponseBodyResults {
	s.Code = &v
	return s
}

func (s *RemoveUsersFromGroupResponseBodyResults) SetMessage(v string) *RemoveUsersFromGroupResponseBodyResults {
	s.Message = &v
	return s
}

type RemoveUsersFromGroupResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveUsersFromGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveUsersFromGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveUsersFromGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveUsersFromGroupResponse) SetHeaders(v map[string]*string) *RemoveUsersFromGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveUsersFromGroupResponse) SetBody(v *RemoveUsersFromGroupResponseBody) *RemoveUsersFromGroupResponse {
	s.Body = v
	return s
}

type ResetHostAccountCredentialRequest struct {
	SourceIp       *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	HostAccountId  *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	CredentialType *string `json:"CredentialType,omitempty" xml:"CredentialType,omitempty"`
}

func (s ResetHostAccountCredentialRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetHostAccountCredentialRequest) GoString() string {
	return s.String()
}

func (s *ResetHostAccountCredentialRequest) SetSourceIp(v string) *ResetHostAccountCredentialRequest {
	s.SourceIp = &v
	return s
}

func (s *ResetHostAccountCredentialRequest) SetInstanceId(v string) *ResetHostAccountCredentialRequest {
	s.InstanceId = &v
	return s
}

func (s *ResetHostAccountCredentialRequest) SetRegionId(v string) *ResetHostAccountCredentialRequest {
	s.RegionId = &v
	return s
}

func (s *ResetHostAccountCredentialRequest) SetHostAccountId(v string) *ResetHostAccountCredentialRequest {
	s.HostAccountId = &v
	return s
}

func (s *ResetHostAccountCredentialRequest) SetCredentialType(v string) *ResetHostAccountCredentialRequest {
	s.CredentialType = &v
	return s
}

type ResetHostAccountCredentialResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetHostAccountCredentialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetHostAccountCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *ResetHostAccountCredentialResponseBody) SetRequestId(v string) *ResetHostAccountCredentialResponseBody {
	s.RequestId = &v
	return s
}

type ResetHostAccountCredentialResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResetHostAccountCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetHostAccountCredentialResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetHostAccountCredentialResponse) GoString() string {
	return s.String()
}

func (s *ResetHostAccountCredentialResponse) SetHeaders(v map[string]*string) *ResetHostAccountCredentialResponse {
	s.Headers = v
	return s
}

func (s *ResetHostAccountCredentialResponse) SetBody(v *ResetHostAccountCredentialResponseBody) *ResetHostAccountCredentialResponse {
	s.Body = v
	return s
}

type StartInstanceRequest struct {
	SourceIp               *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang                   *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	InstanceId             *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	VswitchId              *string   `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	RegionId               *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	InstanceSecurityGroups []*string `json:"InstanceSecurityGroups,omitempty" xml:"InstanceSecurityGroups,omitempty" type:"Repeated"`
	SecurityGroupIds       []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
}

func (s StartInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartInstanceRequest) SetSourceIp(v string) *StartInstanceRequest {
	s.SourceIp = &v
	return s
}

func (s *StartInstanceRequest) SetLang(v string) *StartInstanceRequest {
	s.Lang = &v
	return s
}

func (s *StartInstanceRequest) SetInstanceId(v string) *StartInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *StartInstanceRequest) SetVswitchId(v string) *StartInstanceRequest {
	s.VswitchId = &v
	return s
}

func (s *StartInstanceRequest) SetRegionId(v string) *StartInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *StartInstanceRequest) SetInstanceSecurityGroups(v []*string) *StartInstanceRequest {
	s.InstanceSecurityGroups = v
	return s
}

func (s *StartInstanceRequest) SetSecurityGroupIds(v []*string) *StartInstanceRequest {
	s.SecurityGroupIds = v
	return s
}

type StartInstanceResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s StartInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartInstanceResponseBody) SetRequestId(v string) *StartInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartInstanceResponseBody) SetInstanceId(v string) *StartInstanceResponseBody {
	s.InstanceId = &v
	return s
}

type StartInstanceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartInstanceResponse) SetHeaders(v map[string]*string) *StartInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartInstanceResponse) SetBody(v *StartInstanceResponseBody) *StartInstanceResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	SourceIp     *string                   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang         *string                   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RegionId     *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *TagResourcesRequest) SetSourceIp(v string) *TagResourcesRequest {
	s.SourceIp = &v
	return s
}

func (s *TagResourcesRequest) SetLang(v string) *TagResourcesRequest {
	s.Lang = &v
	return s
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

type UnlockUsersRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserIds    *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
}

func (s UnlockUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s UnlockUsersRequest) GoString() string {
	return s.String()
}

func (s *UnlockUsersRequest) SetSourceIp(v string) *UnlockUsersRequest {
	s.SourceIp = &v
	return s
}

func (s *UnlockUsersRequest) SetInstanceId(v string) *UnlockUsersRequest {
	s.InstanceId = &v
	return s
}

func (s *UnlockUsersRequest) SetRegionId(v string) *UnlockUsersRequest {
	s.RegionId = &v
	return s
}

func (s *UnlockUsersRequest) SetUserIds(v string) *UnlockUsersRequest {
	s.UserIds = &v
	return s
}

type UnlockUsersResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*UnlockUsersResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s UnlockUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnlockUsersResponseBody) GoString() string {
	return s.String()
}

func (s *UnlockUsersResponseBody) SetRequestId(v string) *UnlockUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnlockUsersResponseBody) SetResults(v []*UnlockUsersResponseBodyResults) *UnlockUsersResponseBody {
	s.Results = v
	return s
}

type UnlockUsersResponseBodyResults struct {
	UserId  *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s UnlockUsersResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s UnlockUsersResponseBodyResults) GoString() string {
	return s.String()
}

func (s *UnlockUsersResponseBodyResults) SetUserId(v string) *UnlockUsersResponseBodyResults {
	s.UserId = &v
	return s
}

func (s *UnlockUsersResponseBodyResults) SetCode(v string) *UnlockUsersResponseBodyResults {
	s.Code = &v
	return s
}

func (s *UnlockUsersResponseBodyResults) SetMessage(v string) *UnlockUsersResponseBodyResults {
	s.Message = &v
	return s
}

type UnlockUsersResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnlockUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnlockUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s UnlockUsersResponse) GoString() string {
	return s.String()
}

func (s *UnlockUsersResponse) SetHeaders(v map[string]*string) *UnlockUsersResponse {
	s.Headers = v
	return s
}

func (s *UnlockUsersResponse) SetBody(v *UnlockUsersResponseBody) *UnlockUsersResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	SourceIp     *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang         *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	All          *bool     `json:"All,omitempty" xml:"All,omitempty"`
	ResourceId   []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetSourceIp(v string) *UntagResourcesRequest {
	s.SourceIp = &v
	return s
}

func (s *UntagResourcesRequest) SetLang(v string) *UntagResourcesRequest {
	s.Lang = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
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

type UntagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("yundun-bastionhost"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddHostsToGroupWithOptions(request *AddHostsToGroupRequest, runtime *util.RuntimeOptions) (_result *AddHostsToGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddHostsToGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddHostsToGroup"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddHostsToGroup(request *AddHostsToGroupRequest) (_result *AddHostsToGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddHostsToGroupResponse{}
	_body, _err := client.AddHostsToGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddUsersToGroupWithOptions(request *AddUsersToGroupRequest, runtime *util.RuntimeOptions) (_result *AddUsersToGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddUsersToGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddUsersToGroup"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddUsersToGroup(request *AddUsersToGroupRequest) (_result *AddUsersToGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddUsersToGroupResponse{}
	_body, _err := client.AddUsersToGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachHostAccountsToUserWithOptions(request *AttachHostAccountsToUserRequest, runtime *util.RuntimeOptions) (_result *AttachHostAccountsToUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AttachHostAccountsToUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AttachHostAccountsToUser"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachHostAccountsToUser(request *AttachHostAccountsToUserRequest) (_result *AttachHostAccountsToUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachHostAccountsToUserResponse{}
	_body, _err := client.AttachHostAccountsToUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachHostAccountsToUserGroupWithOptions(request *AttachHostAccountsToUserGroupRequest, runtime *util.RuntimeOptions) (_result *AttachHostAccountsToUserGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AttachHostAccountsToUserGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AttachHostAccountsToUserGroup"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachHostAccountsToUserGroup(request *AttachHostAccountsToUserGroupRequest) (_result *AttachHostAccountsToUserGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachHostAccountsToUserGroupResponse{}
	_body, _err := client.AttachHostAccountsToUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachHostGroupAccountsToUserWithOptions(request *AttachHostGroupAccountsToUserRequest, runtime *util.RuntimeOptions) (_result *AttachHostGroupAccountsToUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AttachHostGroupAccountsToUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AttachHostGroupAccountsToUser"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachHostGroupAccountsToUser(request *AttachHostGroupAccountsToUserRequest) (_result *AttachHostGroupAccountsToUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachHostGroupAccountsToUserResponse{}
	_body, _err := client.AttachHostGroupAccountsToUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachHostGroupAccountsToUserGroupWithOptions(request *AttachHostGroupAccountsToUserGroupRequest, runtime *util.RuntimeOptions) (_result *AttachHostGroupAccountsToUserGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AttachHostGroupAccountsToUserGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AttachHostGroupAccountsToUserGroup"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachHostGroupAccountsToUserGroup(request *AttachHostGroupAccountsToUserGroupRequest) (_result *AttachHostGroupAccountsToUserGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachHostGroupAccountsToUserGroupResponse{}
	_body, _err := client.AttachHostGroupAccountsToUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigInstanceSecurityGroupsWithOptions(request *ConfigInstanceSecurityGroupsRequest, runtime *util.RuntimeOptions) (_result *ConfigInstanceSecurityGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConfigInstanceSecurityGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConfigInstanceSecurityGroups"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigInstanceSecurityGroups(request *ConfigInstanceSecurityGroupsRequest) (_result *ConfigInstanceSecurityGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigInstanceSecurityGroupsResponse{}
	_body, _err := client.ConfigInstanceSecurityGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigInstanceWhiteListWithOptions(request *ConfigInstanceWhiteListRequest, runtime *util.RuntimeOptions) (_result *ConfigInstanceWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConfigInstanceWhiteListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConfigInstanceWhiteList"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigInstanceWhiteList(request *ConfigInstanceWhiteListRequest) (_result *ConfigInstanceWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigInstanceWhiteListResponse{}
	_body, _err := client.ConfigInstanceWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateHostWithOptions(request *CreateHostRequest, runtime *util.RuntimeOptions) (_result *CreateHostResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateHostResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateHost"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateHost(request *CreateHostRequest) (_result *CreateHostResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateHostResponse{}
	_body, _err := client.CreateHostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateHostAccountWithOptions(request *CreateHostAccountRequest, runtime *util.RuntimeOptions) (_result *CreateHostAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateHostAccountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateHostAccount"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateHostAccount(request *CreateHostAccountRequest) (_result *CreateHostAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateHostAccountResponse{}
	_body, _err := client.CreateHostAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateHostGroupWithOptions(request *CreateHostGroupRequest, runtime *util.RuntimeOptions) (_result *CreateHostGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateHostGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateHostGroup"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateHostGroup(request *CreateHostGroupRequest) (_result *CreateHostGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateHostGroupResponse{}
	_body, _err := client.CreateHostGroupWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("CreateUser"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CreateUserGroupWithOptions(request *CreateUserGroupRequest, runtime *util.RuntimeOptions) (_result *CreateUserGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateUserGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateUserGroup"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUserGroup(request *CreateUserGroupRequest) (_result *CreateUserGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUserGroupResponse{}
	_body, _err := client.CreateUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteHostWithOptions(request *DeleteHostRequest, runtime *util.RuntimeOptions) (_result *DeleteHostResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteHostResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteHost"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteHost(request *DeleteHostRequest) (_result *DeleteHostResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteHostResponse{}
	_body, _err := client.DeleteHostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteHostAccountWithOptions(request *DeleteHostAccountRequest, runtime *util.RuntimeOptions) (_result *DeleteHostAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteHostAccountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteHostAccount"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteHostAccount(request *DeleteHostAccountRequest) (_result *DeleteHostAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteHostAccountResponse{}
	_body, _err := client.DeleteHostAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteHostGroupWithOptions(request *DeleteHostGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteHostGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteHostGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteHostGroup"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteHostGroup(request *DeleteHostGroupRequest) (_result *DeleteHostGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteHostGroupResponse{}
	_body, _err := client.DeleteHostGroupWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DeleteUser"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DeleteUserGroupWithOptions(request *DeleteUserGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteUserGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteUserGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteUserGroup"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUserGroup(request *DeleteUserGroupRequest) (_result *DeleteUserGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserGroupResponse{}
	_body, _err := client.DeleteUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceAttributeWithOptions(request *DescribeInstanceAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstanceAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstanceAttribute"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceAttribute(request *DescribeInstanceAttributeRequest) (_result *DescribeInstanceAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceAttributeResponse{}
	_body, _err := client.DescribeInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstancesWithOptions(request *DescribeInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInstances"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstances(request *DescribeInstancesRequest) (_result *DescribeInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.DescribeInstancesWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("DescribeRegions"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DetachHostAccountsFromUserWithOptions(request *DetachHostAccountsFromUserRequest, runtime *util.RuntimeOptions) (_result *DetachHostAccountsFromUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetachHostAccountsFromUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetachHostAccountsFromUser"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetachHostAccountsFromUser(request *DetachHostAccountsFromUserRequest) (_result *DetachHostAccountsFromUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachHostAccountsFromUserResponse{}
	_body, _err := client.DetachHostAccountsFromUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetachHostAccountsFromUserGroupWithOptions(request *DetachHostAccountsFromUserGroupRequest, runtime *util.RuntimeOptions) (_result *DetachHostAccountsFromUserGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetachHostAccountsFromUserGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetachHostAccountsFromUserGroup"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetachHostAccountsFromUserGroup(request *DetachHostAccountsFromUserGroupRequest) (_result *DetachHostAccountsFromUserGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachHostAccountsFromUserGroupResponse{}
	_body, _err := client.DetachHostAccountsFromUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetachHostGroupAccountsFromUserWithOptions(request *DetachHostGroupAccountsFromUserRequest, runtime *util.RuntimeOptions) (_result *DetachHostGroupAccountsFromUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetachHostGroupAccountsFromUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetachHostGroupAccountsFromUser"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetachHostGroupAccountsFromUser(request *DetachHostGroupAccountsFromUserRequest) (_result *DetachHostGroupAccountsFromUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachHostGroupAccountsFromUserResponse{}
	_body, _err := client.DetachHostGroupAccountsFromUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetachHostGroupAccountsFromUserGroupWithOptions(request *DetachHostGroupAccountsFromUserGroupRequest, runtime *util.RuntimeOptions) (_result *DetachHostGroupAccountsFromUserGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DetachHostGroupAccountsFromUserGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DetachHostGroupAccountsFromUserGroup"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetachHostGroupAccountsFromUserGroup(request *DetachHostGroupAccountsFromUserGroupRequest) (_result *DetachHostGroupAccountsFromUserGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachHostGroupAccountsFromUserGroupResponse{}
	_body, _err := client.DetachHostGroupAccountsFromUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableInstancePublicAccessWithOptions(request *DisableInstancePublicAccessRequest, runtime *util.RuntimeOptions) (_result *DisableInstancePublicAccessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DisableInstancePublicAccessResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisableInstancePublicAccess"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableInstancePublicAccess(request *DisableInstancePublicAccessRequest) (_result *DisableInstancePublicAccessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableInstancePublicAccessResponse{}
	_body, _err := client.DisableInstancePublicAccessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableInstancePublicAccessWithOptions(request *EnableInstancePublicAccessRequest, runtime *util.RuntimeOptions) (_result *EnableInstancePublicAccessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableInstancePublicAccessResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableInstancePublicAccess"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableInstancePublicAccess(request *EnableInstancePublicAccessRequest) (_result *EnableInstancePublicAccessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableInstancePublicAccessResponse{}
	_body, _err := client.EnableInstancePublicAccessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHostWithOptions(request *GetHostRequest, runtime *util.RuntimeOptions) (_result *GetHostResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetHostResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetHost"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHost(request *GetHostRequest) (_result *GetHostResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHostResponse{}
	_body, _err := client.GetHostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHostAccountWithOptions(request *GetHostAccountRequest, runtime *util.RuntimeOptions) (_result *GetHostAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetHostAccountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetHostAccount"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHostAccount(request *GetHostAccountRequest) (_result *GetHostAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHostAccountResponse{}
	_body, _err := client.GetHostAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHostGroupWithOptions(request *GetHostGroupRequest, runtime *util.RuntimeOptions) (_result *GetHostGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetHostGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetHostGroup"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHostGroup(request *GetHostGroupRequest) (_result *GetHostGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHostGroupResponse{}
	_body, _err := client.GetHostGroupWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("GetUser"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) GetUserGroupWithOptions(request *GetUserGroupRequest, runtime *util.RuntimeOptions) (_result *GetUserGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetUserGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetUserGroup"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserGroup(request *GetUserGroupRequest) (_result *GetUserGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserGroupResponse{}
	_body, _err := client.GetUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListHostAccountsWithOptions(request *ListHostAccountsRequest, runtime *util.RuntimeOptions) (_result *ListHostAccountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListHostAccountsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListHostAccounts"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListHostAccounts(request *ListHostAccountsRequest) (_result *ListHostAccountsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHostAccountsResponse{}
	_body, _err := client.ListHostAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListHostAccountsForUserWithOptions(request *ListHostAccountsForUserRequest, runtime *util.RuntimeOptions) (_result *ListHostAccountsForUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListHostAccountsForUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListHostAccountsForUser"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListHostAccountsForUser(request *ListHostAccountsForUserRequest) (_result *ListHostAccountsForUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHostAccountsForUserResponse{}
	_body, _err := client.ListHostAccountsForUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListHostAccountsForUserGroupWithOptions(request *ListHostAccountsForUserGroupRequest, runtime *util.RuntimeOptions) (_result *ListHostAccountsForUserGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListHostAccountsForUserGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListHostAccountsForUserGroup"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListHostAccountsForUserGroup(request *ListHostAccountsForUserGroupRequest) (_result *ListHostAccountsForUserGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHostAccountsForUserGroupResponse{}
	_body, _err := client.ListHostAccountsForUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListHostGroupAccountNamesForUserWithOptions(request *ListHostGroupAccountNamesForUserRequest, runtime *util.RuntimeOptions) (_result *ListHostGroupAccountNamesForUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListHostGroupAccountNamesForUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListHostGroupAccountNamesForUser"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListHostGroupAccountNamesForUser(request *ListHostGroupAccountNamesForUserRequest) (_result *ListHostGroupAccountNamesForUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHostGroupAccountNamesForUserResponse{}
	_body, _err := client.ListHostGroupAccountNamesForUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListHostGroupAccountNamesForUserGroupWithOptions(request *ListHostGroupAccountNamesForUserGroupRequest, runtime *util.RuntimeOptions) (_result *ListHostGroupAccountNamesForUserGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListHostGroupAccountNamesForUserGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListHostGroupAccountNamesForUserGroup"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListHostGroupAccountNamesForUserGroup(request *ListHostGroupAccountNamesForUserGroupRequest) (_result *ListHostGroupAccountNamesForUserGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHostGroupAccountNamesForUserGroupResponse{}
	_body, _err := client.ListHostGroupAccountNamesForUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListHostGroupsWithOptions(request *ListHostGroupsRequest, runtime *util.RuntimeOptions) (_result *ListHostGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListHostGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListHostGroups"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListHostGroups(request *ListHostGroupsRequest) (_result *ListHostGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHostGroupsResponse{}
	_body, _err := client.ListHostGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListHostGroupsForUserWithOptions(request *ListHostGroupsForUserRequest, runtime *util.RuntimeOptions) (_result *ListHostGroupsForUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListHostGroupsForUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListHostGroupsForUser"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListHostGroupsForUser(request *ListHostGroupsForUserRequest) (_result *ListHostGroupsForUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHostGroupsForUserResponse{}
	_body, _err := client.ListHostGroupsForUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListHostGroupsForUserGroupWithOptions(request *ListHostGroupsForUserGroupRequest, runtime *util.RuntimeOptions) (_result *ListHostGroupsForUserGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListHostGroupsForUserGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListHostGroupsForUserGroup"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListHostGroupsForUserGroup(request *ListHostGroupsForUserGroupRequest) (_result *ListHostGroupsForUserGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHostGroupsForUserGroupResponse{}
	_body, _err := client.ListHostGroupsForUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListHostsWithOptions(request *ListHostsRequest, runtime *util.RuntimeOptions) (_result *ListHostsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListHostsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListHosts"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListHosts(request *ListHostsRequest) (_result *ListHostsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHostsResponse{}
	_body, _err := client.ListHostsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListHostsForUserWithOptions(request *ListHostsForUserRequest, runtime *util.RuntimeOptions) (_result *ListHostsForUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListHostsForUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListHostsForUser"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListHostsForUser(request *ListHostsForUserRequest) (_result *ListHostsForUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHostsForUserResponse{}
	_body, _err := client.ListHostsForUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListHostsForUserGroupWithOptions(request *ListHostsForUserGroupRequest, runtime *util.RuntimeOptions) (_result *ListHostsForUserGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListHostsForUserGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListHostsForUserGroup"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListHostsForUserGroup(request *ListHostsForUserGroupRequest) (_result *ListHostsForUserGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHostsForUserGroupResponse{}
	_body, _err := client.ListHostsForUserGroupWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("ListTagKeys"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagResources"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ListUserGroupsWithOptions(request *ListUserGroupsRequest, runtime *util.RuntimeOptions) (_result *ListUserGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListUserGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListUserGroups"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserGroups(request *ListUserGroupsRequest) (_result *ListUserGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserGroupsResponse{}
	_body, _err := client.ListUserGroupsWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("ListUsers"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) LockUsersWithOptions(request *LockUsersRequest, runtime *util.RuntimeOptions) (_result *LockUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &LockUsersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("LockUsers"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LockUsers(request *LockUsersRequest) (_result *LockUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LockUsersResponse{}
	_body, _err := client.LockUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyHostWithOptions(request *ModifyHostRequest, runtime *util.RuntimeOptions) (_result *ModifyHostResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyHostResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyHost"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyHost(request *ModifyHostRequest) (_result *ModifyHostResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyHostResponse{}
	_body, _err := client.ModifyHostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyHostAccountWithOptions(request *ModifyHostAccountRequest, runtime *util.RuntimeOptions) (_result *ModifyHostAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyHostAccountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyHostAccount"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyHostAccount(request *ModifyHostAccountRequest) (_result *ModifyHostAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyHostAccountResponse{}
	_body, _err := client.ModifyHostAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyHostGroupWithOptions(request *ModifyHostGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyHostGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyHostGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyHostGroup"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyHostGroup(request *ModifyHostGroupRequest) (_result *ModifyHostGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyHostGroupResponse{}
	_body, _err := client.ModifyHostGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyHostsActiveAddressTypeWithOptions(request *ModifyHostsActiveAddressTypeRequest, runtime *util.RuntimeOptions) (_result *ModifyHostsActiveAddressTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyHostsActiveAddressTypeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyHostsActiveAddressType"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyHostsActiveAddressType(request *ModifyHostsActiveAddressTypeRequest) (_result *ModifyHostsActiveAddressTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyHostsActiveAddressTypeResponse{}
	_body, _err := client.ModifyHostsActiveAddressTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyHostsPortWithOptions(request *ModifyHostsPortRequest, runtime *util.RuntimeOptions) (_result *ModifyHostsPortResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyHostsPortResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyHostsPort"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyHostsPort(request *ModifyHostsPortRequest) (_result *ModifyHostsPortResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyHostsPortResponse{}
	_body, _err := client.ModifyHostsPortWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyInstanceAttributeWithOptions(request *ModifyInstanceAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyInstanceAttributeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyInstanceAttribute"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyInstanceAttribute(request *ModifyInstanceAttributeRequest) (_result *ModifyInstanceAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceAttributeResponse{}
	_body, _err := client.ModifyInstanceAttributeWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("ModifyUser"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ModifyUserGroupWithOptions(request *ModifyUserGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyUserGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyUserGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyUserGroup"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyUserGroup(request *ModifyUserGroupRequest) (_result *ModifyUserGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyUserGroupResponse{}
	_body, _err := client.ModifyUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MoveResourceGroupWithOptions(request *MoveResourceGroupRequest, runtime *util.RuntimeOptions) (_result *MoveResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("MoveResourceGroup"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MoveResourceGroup(request *MoveResourceGroupRequest) (_result *MoveResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.MoveResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveHostsFromGroupWithOptions(request *RemoveHostsFromGroupRequest, runtime *util.RuntimeOptions) (_result *RemoveHostsFromGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveHostsFromGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveHostsFromGroup"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveHostsFromGroup(request *RemoveHostsFromGroupRequest) (_result *RemoveHostsFromGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveHostsFromGroupResponse{}
	_body, _err := client.RemoveHostsFromGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveUsersFromGroupWithOptions(request *RemoveUsersFromGroupRequest, runtime *util.RuntimeOptions) (_result *RemoveUsersFromGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveUsersFromGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveUsersFromGroup"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveUsersFromGroup(request *RemoveUsersFromGroupRequest) (_result *RemoveUsersFromGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveUsersFromGroupResponse{}
	_body, _err := client.RemoveUsersFromGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetHostAccountCredentialWithOptions(request *ResetHostAccountCredentialRequest, runtime *util.RuntimeOptions) (_result *ResetHostAccountCredentialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ResetHostAccountCredentialResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ResetHostAccountCredential"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetHostAccountCredential(request *ResetHostAccountCredentialRequest) (_result *ResetHostAccountCredentialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetHostAccountCredentialResponse{}
	_body, _err := client.ResetHostAccountCredentialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartInstanceWithOptions(request *StartInstanceRequest, runtime *util.RuntimeOptions) (_result *StartInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartInstance"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartInstance(request *StartInstanceRequest) (_result *StartInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartInstanceResponse{}
	_body, _err := client.StartInstanceWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("TagResources"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) UnlockUsersWithOptions(request *UnlockUsersRequest, runtime *util.RuntimeOptions) (_result *UnlockUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UnlockUsersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UnlockUsers"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnlockUsers(request *UnlockUsersRequest) (_result *UnlockUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnlockUsersResponse{}
	_body, _err := client.UnlockUsersWithOptions(request, runtime)
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
	_body, _err := client.DoRPCRequest(tea.String("UntagResources"), tea.String("2019-12-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
