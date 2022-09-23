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

type AddHostsToGroupRequest struct {
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	HostIds     *string `json:"HostIds,omitempty" xml:"HostIds,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddHostsToGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddHostsToGroupRequest) GoString() string {
	return s.String()
}

func (s *AddHostsToGroupRequest) SetHostGroupId(v string) *AddHostsToGroupRequest {
	s.HostGroupId = &v
	return s
}

func (s *AddHostsToGroupRequest) SetHostIds(v string) *AddHostsToGroupRequest {
	s.HostIds = &v
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
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	HostId      *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty"`
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

func (s *AddHostsToGroupResponseBodyResults) SetHostGroupId(v string) *AddHostsToGroupResponseBodyResults {
	s.HostGroupId = &v
	return s
}

func (s *AddHostsToGroupResponseBodyResults) SetHostId(v string) *AddHostsToGroupResponseBodyResults {
	s.HostId = &v
	return s
}

func (s *AddHostsToGroupResponseBodyResults) SetMessage(v string) *AddHostsToGroupResponseBodyResults {
	s.Message = &v
	return s
}

type AddHostsToGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddHostsToGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddHostsToGroupResponse) SetStatusCode(v int32) *AddHostsToGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddHostsToGroupResponse) SetBody(v *AddHostsToGroupResponseBody) *AddHostsToGroupResponse {
	s.Body = v
	return s
}

type AddUsersToGroupRequest struct {
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
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty"`
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AddUsersToGroupResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s AddUsersToGroupResponseBodyResults) GoString() string {
	return s.String()
}

func (s *AddUsersToGroupResponseBodyResults) SetCode(v string) *AddUsersToGroupResponseBodyResults {
	s.Code = &v
	return s
}

func (s *AddUsersToGroupResponseBodyResults) SetMessage(v string) *AddUsersToGroupResponseBodyResults {
	s.Message = &v
	return s
}

func (s *AddUsersToGroupResponseBodyResults) SetUserGroupId(v string) *AddUsersToGroupResponseBodyResults {
	s.UserGroupId = &v
	return s
}

func (s *AddUsersToGroupResponseBodyResults) SetUserId(v string) *AddUsersToGroupResponseBodyResults {
	s.UserId = &v
	return s
}

type AddUsersToGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddUsersToGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddUsersToGroupResponse) SetStatusCode(v int32) *AddUsersToGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUsersToGroupResponse) SetBody(v *AddUsersToGroupResponseBody) *AddUsersToGroupResponse {
	s.Body = v
	return s
}

type AttachHostAccountsToHostShareKeyRequest struct {
	HostAccountIds *string `json:"HostAccountIds,omitempty" xml:"HostAccountIds,omitempty"`
	HostShareKeyId *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AttachHostAccountsToHostShareKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachHostAccountsToHostShareKeyRequest) GoString() string {
	return s.String()
}

func (s *AttachHostAccountsToHostShareKeyRequest) SetHostAccountIds(v string) *AttachHostAccountsToHostShareKeyRequest {
	s.HostAccountIds = &v
	return s
}

func (s *AttachHostAccountsToHostShareKeyRequest) SetHostShareKeyId(v string) *AttachHostAccountsToHostShareKeyRequest {
	s.HostShareKeyId = &v
	return s
}

func (s *AttachHostAccountsToHostShareKeyRequest) SetInstanceId(v string) *AttachHostAccountsToHostShareKeyRequest {
	s.InstanceId = &v
	return s
}

func (s *AttachHostAccountsToHostShareKeyRequest) SetRegionId(v string) *AttachHostAccountsToHostShareKeyRequest {
	s.RegionId = &v
	return s
}

type AttachHostAccountsToHostShareKeyResponseBody struct {
	RequestId *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*AttachHostAccountsToHostShareKeyResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s AttachHostAccountsToHostShareKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachHostAccountsToHostShareKeyResponseBody) GoString() string {
	return s.String()
}

func (s *AttachHostAccountsToHostShareKeyResponseBody) SetRequestId(v string) *AttachHostAccountsToHostShareKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachHostAccountsToHostShareKeyResponseBody) SetResults(v []*AttachHostAccountsToHostShareKeyResponseBodyResults) *AttachHostAccountsToHostShareKeyResponseBody {
	s.Results = v
	return s
}

type AttachHostAccountsToHostShareKeyResponseBodyResults struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HostAccountId  *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	HostShareKeyId *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s AttachHostAccountsToHostShareKeyResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s AttachHostAccountsToHostShareKeyResponseBodyResults) GoString() string {
	return s.String()
}

func (s *AttachHostAccountsToHostShareKeyResponseBodyResults) SetCode(v string) *AttachHostAccountsToHostShareKeyResponseBodyResults {
	s.Code = &v
	return s
}

func (s *AttachHostAccountsToHostShareKeyResponseBodyResults) SetHostAccountId(v string) *AttachHostAccountsToHostShareKeyResponseBodyResults {
	s.HostAccountId = &v
	return s
}

func (s *AttachHostAccountsToHostShareKeyResponseBodyResults) SetHostShareKeyId(v string) *AttachHostAccountsToHostShareKeyResponseBodyResults {
	s.HostShareKeyId = &v
	return s
}

func (s *AttachHostAccountsToHostShareKeyResponseBodyResults) SetMessage(v string) *AttachHostAccountsToHostShareKeyResponseBodyResults {
	s.Message = &v
	return s
}

type AttachHostAccountsToHostShareKeyResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AttachHostAccountsToHostShareKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachHostAccountsToHostShareKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachHostAccountsToHostShareKeyResponse) GoString() string {
	return s.String()
}

func (s *AttachHostAccountsToHostShareKeyResponse) SetHeaders(v map[string]*string) *AttachHostAccountsToHostShareKeyResponse {
	s.Headers = v
	return s
}

func (s *AttachHostAccountsToHostShareKeyResponse) SetStatusCode(v int32) *AttachHostAccountsToHostShareKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachHostAccountsToHostShareKeyResponse) SetBody(v *AttachHostAccountsToHostShareKeyResponseBody) *AttachHostAccountsToHostShareKeyResponse {
	s.Body = v
	return s
}

type AttachHostAccountsToUserRequest struct {
	Hosts      *string `json:"Hosts,omitempty" xml:"Hosts,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AttachHostAccountsToUserRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachHostAccountsToUserRequest) GoString() string {
	return s.String()
}

func (s *AttachHostAccountsToUserRequest) SetHosts(v string) *AttachHostAccountsToUserRequest {
	s.Hosts = &v
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
	Code         *string                                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	HostAccounts []*AttachHostAccountsToUserResponseBodyResultsHostAccounts `json:"HostAccounts,omitempty" xml:"HostAccounts,omitempty" type:"Repeated"`
	HostId       *string                                                    `json:"HostId,omitempty" xml:"HostId,omitempty"`
	Message      *string                                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	UserId       *string                                                    `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AttachHostAccountsToUserResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s AttachHostAccountsToUserResponseBodyResults) GoString() string {
	return s.String()
}

func (s *AttachHostAccountsToUserResponseBodyResults) SetCode(v string) *AttachHostAccountsToUserResponseBodyResults {
	s.Code = &v
	return s
}

func (s *AttachHostAccountsToUserResponseBodyResults) SetHostAccounts(v []*AttachHostAccountsToUserResponseBodyResultsHostAccounts) *AttachHostAccountsToUserResponseBodyResults {
	s.HostAccounts = v
	return s
}

func (s *AttachHostAccountsToUserResponseBodyResults) SetHostId(v string) *AttachHostAccountsToUserResponseBodyResults {
	s.HostId = &v
	return s
}

func (s *AttachHostAccountsToUserResponseBodyResults) SetMessage(v string) *AttachHostAccountsToUserResponseBodyResults {
	s.Message = &v
	return s
}

func (s *AttachHostAccountsToUserResponseBodyResults) SetUserId(v string) *AttachHostAccountsToUserResponseBodyResults {
	s.UserId = &v
	return s
}

type AttachHostAccountsToUserResponseBodyResultsHostAccounts struct {
	Code          *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty"`
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

func (s *AttachHostAccountsToUserResponseBodyResultsHostAccounts) SetHostAccountId(v string) *AttachHostAccountsToUserResponseBodyResultsHostAccounts {
	s.HostAccountId = &v
	return s
}

func (s *AttachHostAccountsToUserResponseBodyResultsHostAccounts) SetMessage(v string) *AttachHostAccountsToUserResponseBodyResultsHostAccounts {
	s.Message = &v
	return s
}

type AttachHostAccountsToUserResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AttachHostAccountsToUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AttachHostAccountsToUserResponse) SetStatusCode(v int32) *AttachHostAccountsToUserResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachHostAccountsToUserResponse) SetBody(v *AttachHostAccountsToUserResponseBody) *AttachHostAccountsToUserResponse {
	s.Body = v
	return s
}

type AttachHostAccountsToUserGroupRequest struct {
	Hosts       *string `json:"Hosts,omitempty" xml:"Hosts,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s AttachHostAccountsToUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachHostAccountsToUserGroupRequest) GoString() string {
	return s.String()
}

func (s *AttachHostAccountsToUserGroupRequest) SetHosts(v string) *AttachHostAccountsToUserGroupRequest {
	s.Hosts = &v
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
	Code         *string                                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	HostAccounts []*AttachHostAccountsToUserGroupResponseBodyResultsHostAccounts `json:"HostAccounts,omitempty" xml:"HostAccounts,omitempty" type:"Repeated"`
	HostId       *string                                                         `json:"HostId,omitempty" xml:"HostId,omitempty"`
	Message      *string                                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	UserGroupId  *string                                                         `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s AttachHostAccountsToUserGroupResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s AttachHostAccountsToUserGroupResponseBodyResults) GoString() string {
	return s.String()
}

func (s *AttachHostAccountsToUserGroupResponseBodyResults) SetCode(v string) *AttachHostAccountsToUserGroupResponseBodyResults {
	s.Code = &v
	return s
}

func (s *AttachHostAccountsToUserGroupResponseBodyResults) SetHostAccounts(v []*AttachHostAccountsToUserGroupResponseBodyResultsHostAccounts) *AttachHostAccountsToUserGroupResponseBodyResults {
	s.HostAccounts = v
	return s
}

func (s *AttachHostAccountsToUserGroupResponseBodyResults) SetHostId(v string) *AttachHostAccountsToUserGroupResponseBodyResults {
	s.HostId = &v
	return s
}

func (s *AttachHostAccountsToUserGroupResponseBodyResults) SetMessage(v string) *AttachHostAccountsToUserGroupResponseBodyResults {
	s.Message = &v
	return s
}

func (s *AttachHostAccountsToUserGroupResponseBodyResults) SetUserGroupId(v string) *AttachHostAccountsToUserGroupResponseBodyResults {
	s.UserGroupId = &v
	return s
}

type AttachHostAccountsToUserGroupResponseBodyResultsHostAccounts struct {
	Code          *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty"`
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

func (s *AttachHostAccountsToUserGroupResponseBodyResultsHostAccounts) SetHostAccountId(v string) *AttachHostAccountsToUserGroupResponseBodyResultsHostAccounts {
	s.HostAccountId = &v
	return s
}

func (s *AttachHostAccountsToUserGroupResponseBodyResultsHostAccounts) SetMessage(v string) *AttachHostAccountsToUserGroupResponseBodyResultsHostAccounts {
	s.Message = &v
	return s
}

type AttachHostAccountsToUserGroupResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AttachHostAccountsToUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AttachHostAccountsToUserGroupResponse) SetStatusCode(v int32) *AttachHostAccountsToUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachHostAccountsToUserGroupResponse) SetBody(v *AttachHostAccountsToUserGroupResponseBody) *AttachHostAccountsToUserGroupResponse {
	s.Body = v
	return s
}

type AttachHostGroupAccountsToUserRequest struct {
	HostGroups *string `json:"HostGroups,omitempty" xml:"HostGroups,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AttachHostGroupAccountsToUserRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachHostGroupAccountsToUserRequest) GoString() string {
	return s.String()
}

func (s *AttachHostGroupAccountsToUserRequest) SetHostGroups(v string) *AttachHostGroupAccountsToUserRequest {
	s.HostGroups = &v
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
	Code             *string                                                             `json:"Code,omitempty" xml:"Code,omitempty"`
	HostAccountNames []*AttachHostGroupAccountsToUserResponseBodyResultsHostAccountNames `json:"HostAccountNames,omitempty" xml:"HostAccountNames,omitempty" type:"Repeated"`
	HostGroupId      *string                                                             `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	Message          *string                                                             `json:"Message,omitempty" xml:"Message,omitempty"`
	UserId           *string                                                             `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AttachHostGroupAccountsToUserResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s AttachHostGroupAccountsToUserResponseBodyResults) GoString() string {
	return s.String()
}

func (s *AttachHostGroupAccountsToUserResponseBodyResults) SetCode(v string) *AttachHostGroupAccountsToUserResponseBodyResults {
	s.Code = &v
	return s
}

func (s *AttachHostGroupAccountsToUserResponseBodyResults) SetHostAccountNames(v []*AttachHostGroupAccountsToUserResponseBodyResultsHostAccountNames) *AttachHostGroupAccountsToUserResponseBodyResults {
	s.HostAccountNames = v
	return s
}

func (s *AttachHostGroupAccountsToUserResponseBodyResults) SetHostGroupId(v string) *AttachHostGroupAccountsToUserResponseBodyResults {
	s.HostGroupId = &v
	return s
}

func (s *AttachHostGroupAccountsToUserResponseBodyResults) SetMessage(v string) *AttachHostGroupAccountsToUserResponseBodyResults {
	s.Message = &v
	return s
}

func (s *AttachHostGroupAccountsToUserResponseBodyResults) SetUserId(v string) *AttachHostGroupAccountsToUserResponseBodyResults {
	s.UserId = &v
	return s
}

type AttachHostGroupAccountsToUserResponseBodyResultsHostAccountNames struct {
	Code            *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	Message         *string `json:"Message,omitempty" xml:"Message,omitempty"`
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

func (s *AttachHostGroupAccountsToUserResponseBodyResultsHostAccountNames) SetHostAccountName(v string) *AttachHostGroupAccountsToUserResponseBodyResultsHostAccountNames {
	s.HostAccountName = &v
	return s
}

func (s *AttachHostGroupAccountsToUserResponseBodyResultsHostAccountNames) SetMessage(v string) *AttachHostGroupAccountsToUserResponseBodyResultsHostAccountNames {
	s.Message = &v
	return s
}

type AttachHostGroupAccountsToUserResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AttachHostGroupAccountsToUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AttachHostGroupAccountsToUserResponse) SetStatusCode(v int32) *AttachHostGroupAccountsToUserResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachHostGroupAccountsToUserResponse) SetBody(v *AttachHostGroupAccountsToUserResponseBody) *AttachHostGroupAccountsToUserResponse {
	s.Body = v
	return s
}

type AttachHostGroupAccountsToUserGroupRequest struct {
	HostGroups  *string `json:"HostGroups,omitempty" xml:"HostGroups,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s AttachHostGroupAccountsToUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachHostGroupAccountsToUserGroupRequest) GoString() string {
	return s.String()
}

func (s *AttachHostGroupAccountsToUserGroupRequest) SetHostGroups(v string) *AttachHostGroupAccountsToUserGroupRequest {
	s.HostGroups = &v
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
	Code             *string                                                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	HostAccountNames []*AttachHostGroupAccountsToUserGroupResponseBodyResultsHostAccountNames `json:"HostAccountNames,omitempty" xml:"HostAccountNames,omitempty" type:"Repeated"`
	HostGroupId      *string                                                                  `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	Message          *string                                                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	UserGroupId      *string                                                                  `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s AttachHostGroupAccountsToUserGroupResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s AttachHostGroupAccountsToUserGroupResponseBodyResults) GoString() string {
	return s.String()
}

func (s *AttachHostGroupAccountsToUserGroupResponseBodyResults) SetCode(v string) *AttachHostGroupAccountsToUserGroupResponseBodyResults {
	s.Code = &v
	return s
}

func (s *AttachHostGroupAccountsToUserGroupResponseBodyResults) SetHostAccountNames(v []*AttachHostGroupAccountsToUserGroupResponseBodyResultsHostAccountNames) *AttachHostGroupAccountsToUserGroupResponseBodyResults {
	s.HostAccountNames = v
	return s
}

func (s *AttachHostGroupAccountsToUserGroupResponseBodyResults) SetHostGroupId(v string) *AttachHostGroupAccountsToUserGroupResponseBodyResults {
	s.HostGroupId = &v
	return s
}

func (s *AttachHostGroupAccountsToUserGroupResponseBodyResults) SetMessage(v string) *AttachHostGroupAccountsToUserGroupResponseBodyResults {
	s.Message = &v
	return s
}

func (s *AttachHostGroupAccountsToUserGroupResponseBodyResults) SetUserGroupId(v string) *AttachHostGroupAccountsToUserGroupResponseBodyResults {
	s.UserGroupId = &v
	return s
}

type AttachHostGroupAccountsToUserGroupResponseBodyResultsHostAccountNames struct {
	Code            *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	Message         *string `json:"Message,omitempty" xml:"Message,omitempty"`
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

func (s *AttachHostGroupAccountsToUserGroupResponseBodyResultsHostAccountNames) SetHostAccountName(v string) *AttachHostGroupAccountsToUserGroupResponseBodyResultsHostAccountNames {
	s.HostAccountName = &v
	return s
}

func (s *AttachHostGroupAccountsToUserGroupResponseBodyResultsHostAccountNames) SetMessage(v string) *AttachHostGroupAccountsToUserGroupResponseBodyResultsHostAccountNames {
	s.Message = &v
	return s
}

type AttachHostGroupAccountsToUserGroupResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AttachHostGroupAccountsToUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AttachHostGroupAccountsToUserGroupResponse) SetStatusCode(v int32) *AttachHostGroupAccountsToUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachHostGroupAccountsToUserGroupResponse) SetBody(v *AttachHostGroupAccountsToUserGroupResponseBody) *AttachHostGroupAccountsToUserGroupResponse {
	s.Body = v
	return s
}

type ConfigInstanceSecurityGroupsRequest struct {
	AuthorizedSecurityGroups []*string `json:"AuthorizedSecurityGroups,omitempty" xml:"AuthorizedSecurityGroups,omitempty" type:"Repeated"`
	InstanceId               *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang                     *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RegionId                 *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ConfigInstanceSecurityGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigInstanceSecurityGroupsRequest) GoString() string {
	return s.String()
}

func (s *ConfigInstanceSecurityGroupsRequest) SetAuthorizedSecurityGroups(v []*string) *ConfigInstanceSecurityGroupsRequest {
	s.AuthorizedSecurityGroups = v
	return s
}

func (s *ConfigInstanceSecurityGroupsRequest) SetInstanceId(v string) *ConfigInstanceSecurityGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *ConfigInstanceSecurityGroupsRequest) SetLang(v string) *ConfigInstanceSecurityGroupsRequest {
	s.Lang = &v
	return s
}

func (s *ConfigInstanceSecurityGroupsRequest) SetRegionId(v string) *ConfigInstanceSecurityGroupsRequest {
	s.RegionId = &v
	return s
}

type ConfigInstanceSecurityGroupsResponseBody struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigInstanceSecurityGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigInstanceSecurityGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigInstanceSecurityGroupsResponseBody) SetInstanceId(v string) *ConfigInstanceSecurityGroupsResponseBody {
	s.InstanceId = &v
	return s
}

func (s *ConfigInstanceSecurityGroupsResponseBody) SetRequestId(v string) *ConfigInstanceSecurityGroupsResponseBody {
	s.RequestId = &v
	return s
}

type ConfigInstanceSecurityGroupsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ConfigInstanceSecurityGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ConfigInstanceSecurityGroupsResponse) SetStatusCode(v int32) *ConfigInstanceSecurityGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigInstanceSecurityGroupsResponse) SetBody(v *ConfigInstanceSecurityGroupsResponseBody) *ConfigInstanceSecurityGroupsResponse {
	s.Body = v
	return s
}

type ConfigInstanceWhiteListRequest struct {
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigInstanceWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigInstanceWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigInstanceWhiteListResponseBody) SetInstanceId(v string) *ConfigInstanceWhiteListResponseBody {
	s.InstanceId = &v
	return s
}

func (s *ConfigInstanceWhiteListResponseBody) SetRequestId(v string) *ConfigInstanceWhiteListResponseBody {
	s.RequestId = &v
	return s
}

type ConfigInstanceWhiteListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ConfigInstanceWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ConfigInstanceWhiteListResponse) SetStatusCode(v int32) *ConfigInstanceWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigInstanceWhiteListResponse) SetBody(v *ConfigInstanceWhiteListResponseBody) *ConfigInstanceWhiteListResponse {
	s.Body = v
	return s
}

type CreateHostRequest struct {
	ActiveAddressType  *string `json:"ActiveAddressType,omitempty" xml:"ActiveAddressType,omitempty"`
	Comment            *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	HostName           *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	HostPrivateAddress *string `json:"HostPrivateAddress,omitempty" xml:"HostPrivateAddress,omitempty"`
	HostPublicAddress  *string `json:"HostPublicAddress,omitempty" xml:"HostPublicAddress,omitempty"`
	InstanceId         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceRegionId   *string `json:"InstanceRegionId,omitempty" xml:"InstanceRegionId,omitempty"`
	OSType             *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Source             *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceInstanceId   *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
}

func (s CreateHostRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateHostRequest) GoString() string {
	return s.String()
}

func (s *CreateHostRequest) SetActiveAddressType(v string) *CreateHostRequest {
	s.ActiveAddressType = &v
	return s
}

func (s *CreateHostRequest) SetComment(v string) *CreateHostRequest {
	s.Comment = &v
	return s
}

func (s *CreateHostRequest) SetHostName(v string) *CreateHostRequest {
	s.HostName = &v
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

func (s *CreateHostRequest) SetInstanceId(v string) *CreateHostRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateHostRequest) SetInstanceRegionId(v string) *CreateHostRequest {
	s.InstanceRegionId = &v
	return s
}

func (s *CreateHostRequest) SetOSType(v string) *CreateHostRequest {
	s.OSType = &v
	return s
}

func (s *CreateHostRequest) SetRegionId(v string) *CreateHostRequest {
	s.RegionId = &v
	return s
}

func (s *CreateHostRequest) SetSource(v string) *CreateHostRequest {
	s.Source = &v
	return s
}

func (s *CreateHostRequest) SetSourceInstanceId(v string) *CreateHostRequest {
	s.SourceInstanceId = &v
	return s
}

type CreateHostResponseBody struct {
	HostId    *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateHostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateHostResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHostResponseBody) SetHostId(v string) *CreateHostResponseBody {
	s.HostId = &v
	return s
}

func (s *CreateHostResponseBody) SetRequestId(v string) *CreateHostResponseBody {
	s.RequestId = &v
	return s
}

type CreateHostResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateHostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateHostResponse) SetStatusCode(v int32) *CreateHostResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHostResponse) SetBody(v *CreateHostResponseBody) *CreateHostResponse {
	s.Body = v
	return s
}

type CreateHostAccountRequest struct {
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	HostId          *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	HostShareKeyId  *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PassPhrase      *string `json:"PassPhrase,omitempty" xml:"PassPhrase,omitempty"`
	Password        *string `json:"Password,omitempty" xml:"Password,omitempty"`
	PrivateKey      *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	ProtocolName    *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateHostAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateHostAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateHostAccountRequest) SetHostAccountName(v string) *CreateHostAccountRequest {
	s.HostAccountName = &v
	return s
}

func (s *CreateHostAccountRequest) SetHostId(v string) *CreateHostAccountRequest {
	s.HostId = &v
	return s
}

func (s *CreateHostAccountRequest) SetHostShareKeyId(v string) *CreateHostAccountRequest {
	s.HostShareKeyId = &v
	return s
}

func (s *CreateHostAccountRequest) SetInstanceId(v string) *CreateHostAccountRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateHostAccountRequest) SetPassPhrase(v string) *CreateHostAccountRequest {
	s.PassPhrase = &v
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

func (s *CreateHostAccountRequest) SetProtocolName(v string) *CreateHostAccountRequest {
	s.ProtocolName = &v
	return s
}

func (s *CreateHostAccountRequest) SetRegionId(v string) *CreateHostAccountRequest {
	s.RegionId = &v
	return s
}

type CreateHostAccountResponseBody struct {
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateHostAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateHostAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHostAccountResponseBody) SetHostAccountId(v string) *CreateHostAccountResponseBody {
	s.HostAccountId = &v
	return s
}

func (s *CreateHostAccountResponseBody) SetRequestId(v string) *CreateHostAccountResponseBody {
	s.RequestId = &v
	return s
}

type CreateHostAccountResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateHostAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateHostAccountResponse) SetStatusCode(v int32) *CreateHostAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHostAccountResponse) SetBody(v *CreateHostAccountResponseBody) *CreateHostAccountResponse {
	s.Body = v
	return s
}

type CreateHostGroupRequest struct {
	Comment       *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	HostGroupName *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateHostGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateHostGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateHostGroupRequest) SetComment(v string) *CreateHostGroupRequest {
	s.Comment = &v
	return s
}

func (s *CreateHostGroupRequest) SetHostGroupName(v string) *CreateHostGroupRequest {
	s.HostGroupName = &v
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateHostGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateHostGroupResponse) SetStatusCode(v int32) *CreateHostGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHostGroupResponse) SetBody(v *CreateHostGroupResponseBody) *CreateHostGroupResponse {
	s.Body = v
	return s
}

type CreateHostShareKeyRequest struct {
	HostShareKeyName *string `json:"HostShareKeyName,omitempty" xml:"HostShareKeyName,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PassPhrase       *string `json:"PassPhrase,omitempty" xml:"PassPhrase,omitempty"`
	PrivateKey       *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateHostShareKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateHostShareKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateHostShareKeyRequest) SetHostShareKeyName(v string) *CreateHostShareKeyRequest {
	s.HostShareKeyName = &v
	return s
}

func (s *CreateHostShareKeyRequest) SetInstanceId(v string) *CreateHostShareKeyRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateHostShareKeyRequest) SetPassPhrase(v string) *CreateHostShareKeyRequest {
	s.PassPhrase = &v
	return s
}

func (s *CreateHostShareKeyRequest) SetPrivateKey(v string) *CreateHostShareKeyRequest {
	s.PrivateKey = &v
	return s
}

func (s *CreateHostShareKeyRequest) SetRegionId(v string) *CreateHostShareKeyRequest {
	s.RegionId = &v
	return s
}

type CreateHostShareKeyResponseBody struct {
	HostShareKeyId *int64  `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateHostShareKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateHostShareKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHostShareKeyResponseBody) SetHostShareKeyId(v int64) *CreateHostShareKeyResponseBody {
	s.HostShareKeyId = &v
	return s
}

func (s *CreateHostShareKeyResponseBody) SetRequestId(v string) *CreateHostShareKeyResponseBody {
	s.RequestId = &v
	return s
}

type CreateHostShareKeyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateHostShareKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateHostShareKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateHostShareKeyResponse) GoString() string {
	return s.String()
}

func (s *CreateHostShareKeyResponse) SetHeaders(v map[string]*string) *CreateHostShareKeyResponse {
	s.Headers = v
	return s
}

func (s *CreateHostShareKeyResponse) SetStatusCode(v int32) *CreateHostShareKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHostShareKeyResponse) SetBody(v *CreateHostShareKeyResponseBody) *CreateHostShareKeyResponse {
	s.Body = v
	return s
}

type CreateUserRequest struct {
	Comment           *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	DisplayName       *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Email             *string `json:"Email,omitempty" xml:"Email,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Mobile            *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	MobileCountryCode *string `json:"MobileCountryCode,omitempty" xml:"MobileCountryCode,omitempty"`
	Password          *string `json:"Password,omitempty" xml:"Password,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Source            *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceUserId      *string `json:"SourceUserId,omitempty" xml:"SourceUserId,omitempty"`
	UserName          *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateUserRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserRequest) GoString() string {
	return s.String()
}

func (s *CreateUserRequest) SetComment(v string) *CreateUserRequest {
	s.Comment = &v
	return s
}

func (s *CreateUserRequest) SetDisplayName(v string) *CreateUserRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateUserRequest) SetEmail(v string) *CreateUserRequest {
	s.Email = &v
	return s
}

func (s *CreateUserRequest) SetInstanceId(v string) *CreateUserRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateUserRequest) SetMobile(v string) *CreateUserRequest {
	s.Mobile = &v
	return s
}

func (s *CreateUserRequest) SetMobileCountryCode(v string) *CreateUserRequest {
	s.MobileCountryCode = &v
	return s
}

func (s *CreateUserRequest) SetPassword(v string) *CreateUserRequest {
	s.Password = &v
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

func (s *CreateUserRequest) SetSourceUserId(v string) *CreateUserRequest {
	s.SourceUserId = &v
	return s
}

func (s *CreateUserRequest) SetUserName(v string) *CreateUserRequest {
	s.UserName = &v
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateUserResponse) SetStatusCode(v int32) *CreateUserResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserResponse) SetBody(v *CreateUserResponseBody) *CreateUserResponse {
	s.Body = v
	return s
}

type CreateUserGroupRequest struct {
	Comment       *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
}

func (s CreateUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateUserGroupRequest) SetComment(v string) *CreateUserGroupRequest {
	s.Comment = &v
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateUserGroupResponse) SetStatusCode(v int32) *CreateUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserGroupResponse) SetBody(v *CreateUserGroupResponseBody) *CreateUserGroupResponse {
	s.Body = v
	return s
}

type DeleteHostRequest struct {
	HostId     *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteHostRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteHostRequest) GoString() string {
	return s.String()
}

func (s *DeleteHostRequest) SetHostId(v string) *DeleteHostRequest {
	s.HostId = &v
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteHostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteHostResponse) SetStatusCode(v int32) *DeleteHostResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHostResponse) SetBody(v *DeleteHostResponseBody) *DeleteHostResponse {
	s.Body = v
	return s
}

type DeleteHostAccountRequest struct {
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteHostAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteHostAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteHostAccountRequest) SetHostAccountId(v string) *DeleteHostAccountRequest {
	s.HostAccountId = &v
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteHostAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteHostAccountResponse) SetStatusCode(v int32) *DeleteHostAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHostAccountResponse) SetBody(v *DeleteHostAccountResponseBody) *DeleteHostAccountResponse {
	s.Body = v
	return s
}

type DeleteHostGroupRequest struct {
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteHostGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteHostGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteHostGroupRequest) SetHostGroupId(v string) *DeleteHostGroupRequest {
	s.HostGroupId = &v
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteHostGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteHostGroupResponse) SetStatusCode(v int32) *DeleteHostGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHostGroupResponse) SetBody(v *DeleteHostGroupResponseBody) *DeleteHostGroupResponse {
	s.Body = v
	return s
}

type DeleteHostShareKeyRequest struct {
	HostShareKeyId *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteHostShareKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteHostShareKeyRequest) GoString() string {
	return s.String()
}

func (s *DeleteHostShareKeyRequest) SetHostShareKeyId(v string) *DeleteHostShareKeyRequest {
	s.HostShareKeyId = &v
	return s
}

func (s *DeleteHostShareKeyRequest) SetInstanceId(v string) *DeleteHostShareKeyRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteHostShareKeyRequest) SetRegionId(v string) *DeleteHostShareKeyRequest {
	s.RegionId = &v
	return s
}

type DeleteHostShareKeyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHostShareKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteHostShareKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHostShareKeyResponseBody) SetRequestId(v string) *DeleteHostShareKeyResponseBody {
	s.RequestId = &v
	return s
}

type DeleteHostShareKeyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteHostShareKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteHostShareKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteHostShareKeyResponse) GoString() string {
	return s.String()
}

func (s *DeleteHostShareKeyResponse) SetHeaders(v map[string]*string) *DeleteHostShareKeyResponse {
	s.Headers = v
	return s
}

func (s *DeleteHostShareKeyResponse) SetStatusCode(v int32) *DeleteHostShareKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHostShareKeyResponse) SetBody(v *DeleteHostShareKeyResponseBody) *DeleteHostShareKeyResponse {
	s.Body = v
	return s
}

type DeleteUserRequest struct {
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteUserResponse) SetStatusCode(v int32) *DeleteUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserResponse) SetBody(v *DeleteUserResponseBody) *DeleteUserResponse {
	s.Body = v
	return s
}

type DeleteUserGroupRequest struct {
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteUserGroupResponse) SetStatusCode(v int32) *DeleteUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserGroupResponse) SetBody(v *DeleteUserGroupResponseBody) *DeleteUserGroupResponse {
	s.Body = v
	return s
}

type DescribeInstanceAttributeRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeInstanceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeRequest) SetInstanceId(v string) *DescribeInstanceAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceAttributeRequest) SetRegionId(v string) *DescribeInstanceAttributeRequest {
	s.RegionId = &v
	return s
}

type DescribeInstanceAttributeResponseBody struct {
	InstanceAttribute *DescribeInstanceAttributeResponseBodyInstanceAttribute `json:"InstanceAttribute,omitempty" xml:"InstanceAttribute,omitempty" type:"Struct"`
	RequestId         *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBody) SetInstanceAttribute(v *DescribeInstanceAttributeResponseBodyInstanceAttribute) *DescribeInstanceAttributeResponseBody {
	s.InstanceAttribute = v
	return s
}

func (s *DescribeInstanceAttributeResponseBody) SetRequestId(v string) *DescribeInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInstanceAttributeResponseBodyInstanceAttribute struct {
	AuthorizedSecurityGroups []*string                                                      `json:"AuthorizedSecurityGroups,omitempty" xml:"AuthorizedSecurityGroups,omitempty" type:"Repeated"`
	DbOperationModule        *string                                                        `json:"DbOperationModule,omitempty" xml:"DbOperationModule,omitempty"`
	Description              *string                                                        `json:"Description,omitempty" xml:"Description,omitempty"`
	EniInstanceId            *string                                                        `json:"EniInstanceId,omitempty" xml:"EniInstanceId,omitempty"`
	ExpireTime               *int64                                                         `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	InstanceId               *string                                                        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceStatus           *string                                                        `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	InternetEndpoint         *string                                                        `json:"InternetEndpoint,omitempty" xml:"InternetEndpoint,omitempty"`
	IntranetEndpoint         *string                                                        `json:"IntranetEndpoint,omitempty" xml:"IntranetEndpoint,omitempty"`
	LicenseCode              *string                                                        `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
	ModifyPasswordModule     *string                                                        `json:"ModifyPasswordModule,omitempty" xml:"ModifyPasswordModule,omitempty"`
	NetworkProxyModule       *string                                                        `json:"NetworkProxyModule,omitempty" xml:"NetworkProxyModule,omitempty"`
	Ports                    []*DescribeInstanceAttributeResponseBodyInstanceAttributePorts `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	PrivateExportIps         []*string                                                      `json:"PrivateExportIps,omitempty" xml:"PrivateExportIps,omitempty" type:"Repeated"`
	PrivateWhiteList         []*string                                                      `json:"PrivateWhiteList,omitempty" xml:"PrivateWhiteList,omitempty" type:"Repeated"`
	PublicExportIps          []*string                                                      `json:"PublicExportIps,omitempty" xml:"PublicExportIps,omitempty" type:"Repeated"`
	PublicIps                []*string                                                      `json:"PublicIps,omitempty" xml:"PublicIps,omitempty" type:"Repeated"`
	PublicNetworkAccess      *bool                                                          `json:"PublicNetworkAccess,omitempty" xml:"PublicNetworkAccess,omitempty"`
	PublicWhiteList          []*string                                                      `json:"PublicWhiteList,omitempty" xml:"PublicWhiteList,omitempty" type:"Repeated"`
	RegionId                 *string                                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId          *string                                                        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityGroupIds         []*string                                                      `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	StartTime                *int64                                                         `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Storage                  *int64                                                         `json:"Storage,omitempty" xml:"Storage,omitempty"`
	VpcId                    *string                                                        `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchId                *string                                                        `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	WebTerminalModule        *string                                                        `json:"WebTerminalModule,omitempty" xml:"WebTerminalModule,omitempty"`
}

func (s DescribeInstanceAttributeResponseBodyInstanceAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBodyInstanceAttribute) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetAuthorizedSecurityGroups(v []*string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.AuthorizedSecurityGroups = v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetDbOperationModule(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.DbOperationModule = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetDescription(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.Description = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetEniInstanceId(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.EniInstanceId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetExpireTime(v int64) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetInstanceId(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetInstanceStatus(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetInternetEndpoint(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.InternetEndpoint = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetIntranetEndpoint(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.IntranetEndpoint = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetLicenseCode(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.LicenseCode = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetModifyPasswordModule(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.ModifyPasswordModule = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetNetworkProxyModule(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.NetworkProxyModule = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetPorts(v []*DescribeInstanceAttributeResponseBodyInstanceAttributePorts) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.Ports = v
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

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetPublicExportIps(v []*string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.PublicExportIps = v
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

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetPublicWhiteList(v []*string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.PublicWhiteList = v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetRegionId(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetResourceGroupId(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetSecurityGroupIds(v []*string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.SecurityGroupIds = v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetStartTime(v int64) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.StartTime = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetStorage(v int64) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.Storage = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetVpcId(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.VpcId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetVswitchId(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.VswitchId = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetWebTerminalModule(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.WebTerminalModule = &v
	return s
}

type DescribeInstanceAttributeResponseBodyInstanceAttributePorts struct {
	CustomPort   *int32 `json:"CustomPort,omitempty" xml:"CustomPort,omitempty"`
	StandardPort *int32 `json:"StandardPort,omitempty" xml:"StandardPort,omitempty"`
}

func (s DescribeInstanceAttributeResponseBodyInstanceAttributePorts) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceAttributeResponseBodyInstanceAttributePorts) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttributePorts) SetCustomPort(v int32) *DescribeInstanceAttributeResponseBodyInstanceAttributePorts {
	s.CustomPort = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttributePorts) SetStandardPort(v int32) *DescribeInstanceAttributeResponseBodyInstanceAttributePorts {
	s.StandardPort = &v
	return s
}

type DescribeInstanceAttributeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeInstanceAttributeResponse) SetStatusCode(v int32) *DescribeInstanceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceAttributeResponse) SetBody(v *DescribeInstanceAttributeResponseBody) *DescribeInstanceAttributeResponse {
	s.Body = v
	return s
}

type DescribeInstancesRequest struct {
	InstanceId      []*string                      `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	InstanceStatus  *string                        `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	PageNumber      *int32                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId        *string                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string                        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag             []*DescribeInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequest) SetInstanceId(v []*string) *DescribeInstancesRequest {
	s.InstanceId = v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceStatus(v string) *DescribeInstancesRequest {
	s.InstanceStatus = &v
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

func (s *DescribeInstancesRequest) SetResourceGroupId(v string) *DescribeInstancesRequest {
	s.ResourceGroupId = &v
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
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *DescribeInstancesResponseBody) SetRequestId(v string) *DescribeInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetTotalCount(v int64) *DescribeInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeInstancesResponseBodyInstances struct {
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ExpireTime          *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	ImageVersion        *string `json:"ImageVersion,omitempty" xml:"ImageVersion,omitempty"`
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceStatus      *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	InternetEndpoint    *string `json:"InternetEndpoint,omitempty" xml:"InternetEndpoint,omitempty"`
	IntranetEndpoint    *string `json:"IntranetEndpoint,omitempty" xml:"IntranetEndpoint,omitempty"`
	Legacy              *bool   `json:"Legacy,omitempty" xml:"Legacy,omitempty"`
	LicenseCode         *string `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
	PlanCode            *string `json:"PlanCode,omitempty" xml:"PlanCode,omitempty"`
	PublicNetworkAccess *bool   `json:"PublicNetworkAccess,omitempty" xml:"PublicNetworkAccess,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId     *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	StartTime           *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	VpcId               *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchId           *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s DescribeInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstances) SetDescription(v string) *DescribeInstancesResponseBodyInstances {
	s.Description = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetExpireTime(v int64) *DescribeInstancesResponseBodyInstances {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetImageVersion(v string) *DescribeInstancesResponseBodyInstances {
	s.ImageVersion = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInstanceId(v string) *DescribeInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInstanceStatus(v string) *DescribeInstancesResponseBodyInstances {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInternetEndpoint(v string) *DescribeInstancesResponseBodyInstances {
	s.InternetEndpoint = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetIntranetEndpoint(v string) *DescribeInstancesResponseBodyInstances {
	s.IntranetEndpoint = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetLegacy(v bool) *DescribeInstancesResponseBodyInstances {
	s.Legacy = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetLicenseCode(v string) *DescribeInstancesResponseBodyInstances {
	s.LicenseCode = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetPlanCode(v string) *DescribeInstancesResponseBodyInstances {
	s.PlanCode = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetPublicNetworkAccess(v bool) *DescribeInstancesResponseBodyInstances {
	s.PublicNetworkAccess = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetRegionId(v string) *DescribeInstancesResponseBodyInstances {
	s.RegionId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetResourceGroupId(v string) *DescribeInstancesResponseBodyInstances {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetStartTime(v int64) *DescribeInstancesResponseBodyInstances {
	s.StartTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetVpcId(v string) *DescribeInstancesResponseBodyInstances {
	s.VpcId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetVswitchId(v string) *DescribeInstancesResponseBodyInstances {
	s.VswitchId = &v
	return s
}

type DescribeInstancesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeInstancesResponse) SetStatusCode(v int32) *DescribeInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstancesResponse) SetBody(v *DescribeInstancesResponseBody) *DescribeInstancesResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeRegionsRequest) SetRegionId(v string) *DescribeRegionsRequest {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponseBody struct {
	Regions   []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DetachHostAccountsFromHostShareKeyRequest struct {
	HostAccountIds *string `json:"HostAccountIds,omitempty" xml:"HostAccountIds,omitempty"`
	HostShareKeyId *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DetachHostAccountsFromHostShareKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachHostAccountsFromHostShareKeyRequest) GoString() string {
	return s.String()
}

func (s *DetachHostAccountsFromHostShareKeyRequest) SetHostAccountIds(v string) *DetachHostAccountsFromHostShareKeyRequest {
	s.HostAccountIds = &v
	return s
}

func (s *DetachHostAccountsFromHostShareKeyRequest) SetHostShareKeyId(v string) *DetachHostAccountsFromHostShareKeyRequest {
	s.HostShareKeyId = &v
	return s
}

func (s *DetachHostAccountsFromHostShareKeyRequest) SetInstanceId(v string) *DetachHostAccountsFromHostShareKeyRequest {
	s.InstanceId = &v
	return s
}

func (s *DetachHostAccountsFromHostShareKeyRequest) SetRegionId(v string) *DetachHostAccountsFromHostShareKeyRequest {
	s.RegionId = &v
	return s
}

type DetachHostAccountsFromHostShareKeyResponseBody struct {
	RequestId *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*DetachHostAccountsFromHostShareKeyResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s DetachHostAccountsFromHostShareKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachHostAccountsFromHostShareKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DetachHostAccountsFromHostShareKeyResponseBody) SetRequestId(v string) *DetachHostAccountsFromHostShareKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachHostAccountsFromHostShareKeyResponseBody) SetResults(v []*DetachHostAccountsFromHostShareKeyResponseBodyResults) *DetachHostAccountsFromHostShareKeyResponseBody {
	s.Results = v
	return s
}

type DetachHostAccountsFromHostShareKeyResponseBodyResults struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HostAccountId  *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	HostShareKeyId *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DetachHostAccountsFromHostShareKeyResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s DetachHostAccountsFromHostShareKeyResponseBodyResults) GoString() string {
	return s.String()
}

func (s *DetachHostAccountsFromHostShareKeyResponseBodyResults) SetCode(v string) *DetachHostAccountsFromHostShareKeyResponseBodyResults {
	s.Code = &v
	return s
}

func (s *DetachHostAccountsFromHostShareKeyResponseBodyResults) SetHostAccountId(v string) *DetachHostAccountsFromHostShareKeyResponseBodyResults {
	s.HostAccountId = &v
	return s
}

func (s *DetachHostAccountsFromHostShareKeyResponseBodyResults) SetHostShareKeyId(v string) *DetachHostAccountsFromHostShareKeyResponseBodyResults {
	s.HostShareKeyId = &v
	return s
}

func (s *DetachHostAccountsFromHostShareKeyResponseBodyResults) SetMessage(v string) *DetachHostAccountsFromHostShareKeyResponseBodyResults {
	s.Message = &v
	return s
}

type DetachHostAccountsFromHostShareKeyResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetachHostAccountsFromHostShareKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetachHostAccountsFromHostShareKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachHostAccountsFromHostShareKeyResponse) GoString() string {
	return s.String()
}

func (s *DetachHostAccountsFromHostShareKeyResponse) SetHeaders(v map[string]*string) *DetachHostAccountsFromHostShareKeyResponse {
	s.Headers = v
	return s
}

func (s *DetachHostAccountsFromHostShareKeyResponse) SetStatusCode(v int32) *DetachHostAccountsFromHostShareKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachHostAccountsFromHostShareKeyResponse) SetBody(v *DetachHostAccountsFromHostShareKeyResponseBody) *DetachHostAccountsFromHostShareKeyResponse {
	s.Body = v
	return s
}

type DetachHostAccountsFromUserRequest struct {
	Hosts      *string `json:"Hosts,omitempty" xml:"Hosts,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DetachHostAccountsFromUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachHostAccountsFromUserRequest) GoString() string {
	return s.String()
}

func (s *DetachHostAccountsFromUserRequest) SetHosts(v string) *DetachHostAccountsFromUserRequest {
	s.Hosts = &v
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
	Code         *string                                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	HostAccounts []*DetachHostAccountsFromUserResponseBodyResultsHostAccounts `json:"HostAccounts,omitempty" xml:"HostAccounts,omitempty" type:"Repeated"`
	HostId       *string                                                      `json:"HostId,omitempty" xml:"HostId,omitempty"`
	Message      *string                                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	UserId       *string                                                      `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DetachHostAccountsFromUserResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s DetachHostAccountsFromUserResponseBodyResults) GoString() string {
	return s.String()
}

func (s *DetachHostAccountsFromUserResponseBodyResults) SetCode(v string) *DetachHostAccountsFromUserResponseBodyResults {
	s.Code = &v
	return s
}

func (s *DetachHostAccountsFromUserResponseBodyResults) SetHostAccounts(v []*DetachHostAccountsFromUserResponseBodyResultsHostAccounts) *DetachHostAccountsFromUserResponseBodyResults {
	s.HostAccounts = v
	return s
}

func (s *DetachHostAccountsFromUserResponseBodyResults) SetHostId(v string) *DetachHostAccountsFromUserResponseBodyResults {
	s.HostId = &v
	return s
}

func (s *DetachHostAccountsFromUserResponseBodyResults) SetMessage(v string) *DetachHostAccountsFromUserResponseBodyResults {
	s.Message = &v
	return s
}

func (s *DetachHostAccountsFromUserResponseBodyResults) SetUserId(v string) *DetachHostAccountsFromUserResponseBodyResults {
	s.UserId = &v
	return s
}

type DetachHostAccountsFromUserResponseBodyResultsHostAccounts struct {
	Code          *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty"`
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

func (s *DetachHostAccountsFromUserResponseBodyResultsHostAccounts) SetHostAccountId(v string) *DetachHostAccountsFromUserResponseBodyResultsHostAccounts {
	s.HostAccountId = &v
	return s
}

func (s *DetachHostAccountsFromUserResponseBodyResultsHostAccounts) SetMessage(v string) *DetachHostAccountsFromUserResponseBodyResultsHostAccounts {
	s.Message = &v
	return s
}

type DetachHostAccountsFromUserResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetachHostAccountsFromUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DetachHostAccountsFromUserResponse) SetStatusCode(v int32) *DetachHostAccountsFromUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachHostAccountsFromUserResponse) SetBody(v *DetachHostAccountsFromUserResponseBody) *DetachHostAccountsFromUserResponse {
	s.Body = v
	return s
}

type DetachHostAccountsFromUserGroupRequest struct {
	Hosts       *string `json:"Hosts,omitempty" xml:"Hosts,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s DetachHostAccountsFromUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachHostAccountsFromUserGroupRequest) GoString() string {
	return s.String()
}

func (s *DetachHostAccountsFromUserGroupRequest) SetHosts(v string) *DetachHostAccountsFromUserGroupRequest {
	s.Hosts = &v
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
	Code         *string                                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	HostAccounts []*DetachHostAccountsFromUserGroupResponseBodyResultsHostAccounts `json:"HostAccounts,omitempty" xml:"HostAccounts,omitempty" type:"Repeated"`
	HostId       *string                                                           `json:"HostId,omitempty" xml:"HostId,omitempty"`
	Message      *string                                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	UserGroupId  *string                                                           `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s DetachHostAccountsFromUserGroupResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s DetachHostAccountsFromUserGroupResponseBodyResults) GoString() string {
	return s.String()
}

func (s *DetachHostAccountsFromUserGroupResponseBodyResults) SetCode(v string) *DetachHostAccountsFromUserGroupResponseBodyResults {
	s.Code = &v
	return s
}

func (s *DetachHostAccountsFromUserGroupResponseBodyResults) SetHostAccounts(v []*DetachHostAccountsFromUserGroupResponseBodyResultsHostAccounts) *DetachHostAccountsFromUserGroupResponseBodyResults {
	s.HostAccounts = v
	return s
}

func (s *DetachHostAccountsFromUserGroupResponseBodyResults) SetHostId(v string) *DetachHostAccountsFromUserGroupResponseBodyResults {
	s.HostId = &v
	return s
}

func (s *DetachHostAccountsFromUserGroupResponseBodyResults) SetMessage(v string) *DetachHostAccountsFromUserGroupResponseBodyResults {
	s.Message = &v
	return s
}

func (s *DetachHostAccountsFromUserGroupResponseBodyResults) SetUserGroupId(v string) *DetachHostAccountsFromUserGroupResponseBodyResults {
	s.UserGroupId = &v
	return s
}

type DetachHostAccountsFromUserGroupResponseBodyResultsHostAccounts struct {
	Code          *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty"`
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

func (s *DetachHostAccountsFromUserGroupResponseBodyResultsHostAccounts) SetHostAccountId(v string) *DetachHostAccountsFromUserGroupResponseBodyResultsHostAccounts {
	s.HostAccountId = &v
	return s
}

func (s *DetachHostAccountsFromUserGroupResponseBodyResultsHostAccounts) SetMessage(v string) *DetachHostAccountsFromUserGroupResponseBodyResultsHostAccounts {
	s.Message = &v
	return s
}

type DetachHostAccountsFromUserGroupResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetachHostAccountsFromUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DetachHostAccountsFromUserGroupResponse) SetStatusCode(v int32) *DetachHostAccountsFromUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachHostAccountsFromUserGroupResponse) SetBody(v *DetachHostAccountsFromUserGroupResponseBody) *DetachHostAccountsFromUserGroupResponse {
	s.Body = v
	return s
}

type DetachHostGroupAccountsFromUserRequest struct {
	HostGroups *string `json:"HostGroups,omitempty" xml:"HostGroups,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DetachHostGroupAccountsFromUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachHostGroupAccountsFromUserRequest) GoString() string {
	return s.String()
}

func (s *DetachHostGroupAccountsFromUserRequest) SetHostGroups(v string) *DetachHostGroupAccountsFromUserRequest {
	s.HostGroups = &v
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
	Code             *string                                                               `json:"Code,omitempty" xml:"Code,omitempty"`
	HostAccountNames []*DetachHostGroupAccountsFromUserResponseBodyResultsHostAccountNames `json:"HostAccountNames,omitempty" xml:"HostAccountNames,omitempty" type:"Repeated"`
	HostGroupId      *string                                                               `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	Message          *string                                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	UserId           *string                                                               `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DetachHostGroupAccountsFromUserResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s DetachHostGroupAccountsFromUserResponseBodyResults) GoString() string {
	return s.String()
}

func (s *DetachHostGroupAccountsFromUserResponseBodyResults) SetCode(v string) *DetachHostGroupAccountsFromUserResponseBodyResults {
	s.Code = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserResponseBodyResults) SetHostAccountNames(v []*DetachHostGroupAccountsFromUserResponseBodyResultsHostAccountNames) *DetachHostGroupAccountsFromUserResponseBodyResults {
	s.HostAccountNames = v
	return s
}

func (s *DetachHostGroupAccountsFromUserResponseBodyResults) SetHostGroupId(v string) *DetachHostGroupAccountsFromUserResponseBodyResults {
	s.HostGroupId = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserResponseBodyResults) SetMessage(v string) *DetachHostGroupAccountsFromUserResponseBodyResults {
	s.Message = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserResponseBodyResults) SetUserId(v string) *DetachHostGroupAccountsFromUserResponseBodyResults {
	s.UserId = &v
	return s
}

type DetachHostGroupAccountsFromUserResponseBodyResultsHostAccountNames struct {
	Code            *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	Message         *string `json:"Message,omitempty" xml:"Message,omitempty"`
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

func (s *DetachHostGroupAccountsFromUserResponseBodyResultsHostAccountNames) SetHostAccountName(v string) *DetachHostGroupAccountsFromUserResponseBodyResultsHostAccountNames {
	s.HostAccountName = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserResponseBodyResultsHostAccountNames) SetMessage(v string) *DetachHostGroupAccountsFromUserResponseBodyResultsHostAccountNames {
	s.Message = &v
	return s
}

type DetachHostGroupAccountsFromUserResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetachHostGroupAccountsFromUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DetachHostGroupAccountsFromUserResponse) SetStatusCode(v int32) *DetachHostGroupAccountsFromUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserResponse) SetBody(v *DetachHostGroupAccountsFromUserResponseBody) *DetachHostGroupAccountsFromUserResponse {
	s.Body = v
	return s
}

type DetachHostGroupAccountsFromUserGroupRequest struct {
	HostGroups  *string `json:"HostGroups,omitempty" xml:"HostGroups,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s DetachHostGroupAccountsFromUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachHostGroupAccountsFromUserGroupRequest) GoString() string {
	return s.String()
}

func (s *DetachHostGroupAccountsFromUserGroupRequest) SetHostGroups(v string) *DetachHostGroupAccountsFromUserGroupRequest {
	s.HostGroups = &v
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
	Code             *string                                                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	HostAccountNames []*DetachHostGroupAccountsFromUserGroupResponseBodyResultsHostAccountNames `json:"HostAccountNames,omitempty" xml:"HostAccountNames,omitempty" type:"Repeated"`
	HostGroupId      *string                                                                    `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	Message          *string                                                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	UserGroupId      *string                                                                    `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s DetachHostGroupAccountsFromUserGroupResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s DetachHostGroupAccountsFromUserGroupResponseBodyResults) GoString() string {
	return s.String()
}

func (s *DetachHostGroupAccountsFromUserGroupResponseBodyResults) SetCode(v string) *DetachHostGroupAccountsFromUserGroupResponseBodyResults {
	s.Code = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserGroupResponseBodyResults) SetHostAccountNames(v []*DetachHostGroupAccountsFromUserGroupResponseBodyResultsHostAccountNames) *DetachHostGroupAccountsFromUserGroupResponseBodyResults {
	s.HostAccountNames = v
	return s
}

func (s *DetachHostGroupAccountsFromUserGroupResponseBodyResults) SetHostGroupId(v string) *DetachHostGroupAccountsFromUserGroupResponseBodyResults {
	s.HostGroupId = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserGroupResponseBodyResults) SetMessage(v string) *DetachHostGroupAccountsFromUserGroupResponseBodyResults {
	s.Message = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserGroupResponseBodyResults) SetUserGroupId(v string) *DetachHostGroupAccountsFromUserGroupResponseBodyResults {
	s.UserGroupId = &v
	return s
}

type DetachHostGroupAccountsFromUserGroupResponseBodyResultsHostAccountNames struct {
	Code            *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	Message         *string `json:"Message,omitempty" xml:"Message,omitempty"`
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

func (s *DetachHostGroupAccountsFromUserGroupResponseBodyResultsHostAccountNames) SetHostAccountName(v string) *DetachHostGroupAccountsFromUserGroupResponseBodyResultsHostAccountNames {
	s.HostAccountName = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserGroupResponseBodyResultsHostAccountNames) SetMessage(v string) *DetachHostGroupAccountsFromUserGroupResponseBodyResultsHostAccountNames {
	s.Message = &v
	return s
}

type DetachHostGroupAccountsFromUserGroupResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetachHostGroupAccountsFromUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DetachHostGroupAccountsFromUserGroupResponse) SetStatusCode(v int32) *DetachHostGroupAccountsFromUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachHostGroupAccountsFromUserGroupResponse) SetBody(v *DetachHostGroupAccountsFromUserGroupResponseBody) *DetachHostGroupAccountsFromUserGroupResponse {
	s.Body = v
	return s
}

type DisableInstancePublicAccessRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DisableInstancePublicAccessRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableInstancePublicAccessRequest) GoString() string {
	return s.String()
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableInstancePublicAccessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableInstancePublicAccessResponseBody) GoString() string {
	return s.String()
}

func (s *DisableInstancePublicAccessResponseBody) SetInstanceId(v string) *DisableInstancePublicAccessResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DisableInstancePublicAccessResponseBody) SetRequestId(v string) *DisableInstancePublicAccessResponseBody {
	s.RequestId = &v
	return s
}

type DisableInstancePublicAccessResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableInstancePublicAccessResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DisableInstancePublicAccessResponse) SetStatusCode(v int32) *DisableInstancePublicAccessResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableInstancePublicAccessResponse) SetBody(v *DisableInstancePublicAccessResponseBody) *DisableInstancePublicAccessResponse {
	s.Body = v
	return s
}

type EnableInstancePublicAccessRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s EnableInstancePublicAccessRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableInstancePublicAccessRequest) GoString() string {
	return s.String()
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableInstancePublicAccessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableInstancePublicAccessResponseBody) GoString() string {
	return s.String()
}

func (s *EnableInstancePublicAccessResponseBody) SetInstanceId(v string) *EnableInstancePublicAccessResponseBody {
	s.InstanceId = &v
	return s
}

func (s *EnableInstancePublicAccessResponseBody) SetRequestId(v string) *EnableInstancePublicAccessResponseBody {
	s.RequestId = &v
	return s
}

type EnableInstancePublicAccessResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableInstancePublicAccessResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *EnableInstancePublicAccessResponse) SetStatusCode(v int32) *EnableInstancePublicAccessResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableInstancePublicAccessResponse) SetBody(v *EnableInstancePublicAccessResponseBody) *EnableInstancePublicAccessResponse {
	s.Body = v
	return s
}

type GetHostRequest struct {
	HostId     *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetHostRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHostRequest) GoString() string {
	return s.String()
}

func (s *GetHostRequest) SetHostId(v string) *GetHostRequest {
	s.HostId = &v
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

type GetHostResponseBody struct {
	Host      *GetHostResponseBodyHost `json:"Host,omitempty" xml:"Host,omitempty" type:"Struct"`
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetHostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHostResponseBody) GoString() string {
	return s.String()
}

func (s *GetHostResponseBody) SetHost(v *GetHostResponseBodyHost) *GetHostResponseBody {
	s.Host = v
	return s
}

func (s *GetHostResponseBody) SetRequestId(v string) *GetHostResponseBody {
	s.RequestId = &v
	return s
}

type GetHostResponseBodyHost struct {
	ActiveAddressType   *string                             `json:"ActiveAddressType,omitempty" xml:"ActiveAddressType,omitempty"`
	Comment             *string                             `json:"Comment,omitempty" xml:"Comment,omitempty"`
	HostId              *string                             `json:"HostId,omitempty" xml:"HostId,omitempty"`
	HostName            *string                             `json:"HostName,omitempty" xml:"HostName,omitempty"`
	HostPrivateAddress  *string                             `json:"HostPrivateAddress,omitempty" xml:"HostPrivateAddress,omitempty"`
	HostPublicAddress   *string                             `json:"HostPublicAddress,omitempty" xml:"HostPublicAddress,omitempty"`
	OSType              *string                             `json:"OSType,omitempty" xml:"OSType,omitempty"`
	Protocols           []*GetHostResponseBodyHostProtocols `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
	Source              *string                             `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceInstanceId    *string                             `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
	SourceInstanceState *string                             `json:"SourceInstanceState,omitempty" xml:"SourceInstanceState,omitempty"`
}

func (s GetHostResponseBodyHost) String() string {
	return tea.Prettify(s)
}

func (s GetHostResponseBodyHost) GoString() string {
	return s.String()
}

func (s *GetHostResponseBodyHost) SetActiveAddressType(v string) *GetHostResponseBodyHost {
	s.ActiveAddressType = &v
	return s
}

func (s *GetHostResponseBodyHost) SetComment(v string) *GetHostResponseBodyHost {
	s.Comment = &v
	return s
}

func (s *GetHostResponseBodyHost) SetHostId(v string) *GetHostResponseBodyHost {
	s.HostId = &v
	return s
}

func (s *GetHostResponseBodyHost) SetHostName(v string) *GetHostResponseBodyHost {
	s.HostName = &v
	return s
}

func (s *GetHostResponseBodyHost) SetHostPrivateAddress(v string) *GetHostResponseBodyHost {
	s.HostPrivateAddress = &v
	return s
}

func (s *GetHostResponseBodyHost) SetHostPublicAddress(v string) *GetHostResponseBodyHost {
	s.HostPublicAddress = &v
	return s
}

func (s *GetHostResponseBodyHost) SetOSType(v string) *GetHostResponseBodyHost {
	s.OSType = &v
	return s
}

func (s *GetHostResponseBodyHost) SetProtocols(v []*GetHostResponseBodyHostProtocols) *GetHostResponseBodyHost {
	s.Protocols = v
	return s
}

func (s *GetHostResponseBodyHost) SetSource(v string) *GetHostResponseBodyHost {
	s.Source = &v
	return s
}

func (s *GetHostResponseBodyHost) SetSourceInstanceId(v string) *GetHostResponseBodyHost {
	s.SourceInstanceId = &v
	return s
}

func (s *GetHostResponseBodyHost) SetSourceInstanceState(v string) *GetHostResponseBodyHost {
	s.SourceInstanceState = &v
	return s
}

type GetHostResponseBodyHostProtocols struct {
	HostFingerPrint *string `json:"HostFingerPrint,omitempty" xml:"HostFingerPrint,omitempty"`
	Port            *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	ProtocolName    *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
}

func (s GetHostResponseBodyHostProtocols) String() string {
	return tea.Prettify(s)
}

func (s GetHostResponseBodyHostProtocols) GoString() string {
	return s.String()
}

func (s *GetHostResponseBodyHostProtocols) SetHostFingerPrint(v string) *GetHostResponseBodyHostProtocols {
	s.HostFingerPrint = &v
	return s
}

func (s *GetHostResponseBodyHostProtocols) SetPort(v int32) *GetHostResponseBodyHostProtocols {
	s.Port = &v
	return s
}

func (s *GetHostResponseBodyHostProtocols) SetProtocolName(v string) *GetHostResponseBodyHostProtocols {
	s.ProtocolName = &v
	return s
}

type GetHostResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetHostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetHostResponse) SetStatusCode(v int32) *GetHostResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHostResponse) SetBody(v *GetHostResponseBody) *GetHostResponse {
	s.Body = v
	return s
}

type GetHostAccountRequest struct {
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetHostAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHostAccountRequest) GoString() string {
	return s.String()
}

func (s *GetHostAccountRequest) SetHostAccountId(v string) *GetHostAccountRequest {
	s.HostAccountId = &v
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

type GetHostAccountResponseBody struct {
	HostAccount *GetHostAccountResponseBodyHostAccount `json:"HostAccount,omitempty" xml:"HostAccount,omitempty" type:"Struct"`
	RequestId   *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetHostAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHostAccountResponseBody) GoString() string {
	return s.String()
}

func (s *GetHostAccountResponseBody) SetHostAccount(v *GetHostAccountResponseBodyHostAccount) *GetHostAccountResponseBody {
	s.HostAccount = v
	return s
}

func (s *GetHostAccountResponseBody) SetRequestId(v string) *GetHostAccountResponseBody {
	s.RequestId = &v
	return s
}

type GetHostAccountResponseBodyHostAccount struct {
	HasPassword           *bool   `json:"HasPassword,omitempty" xml:"HasPassword,omitempty"`
	HostAccountId         *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	HostAccountName       *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	HostId                *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	HostShareKeyId        *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	HostShareKeyName      *string `json:"HostShareKeyName,omitempty" xml:"HostShareKeyName,omitempty"`
	PrivateKeyFingerprint *string `json:"PrivateKeyFingerprint,omitempty" xml:"PrivateKeyFingerprint,omitempty"`
	ProtocolName          *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
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

func (s *GetHostAccountResponseBodyHostAccount) SetHostAccountId(v string) *GetHostAccountResponseBodyHostAccount {
	s.HostAccountId = &v
	return s
}

func (s *GetHostAccountResponseBodyHostAccount) SetHostAccountName(v string) *GetHostAccountResponseBodyHostAccount {
	s.HostAccountName = &v
	return s
}

func (s *GetHostAccountResponseBodyHostAccount) SetHostId(v string) *GetHostAccountResponseBodyHostAccount {
	s.HostId = &v
	return s
}

func (s *GetHostAccountResponseBodyHostAccount) SetHostShareKeyId(v string) *GetHostAccountResponseBodyHostAccount {
	s.HostShareKeyId = &v
	return s
}

func (s *GetHostAccountResponseBodyHostAccount) SetHostShareKeyName(v string) *GetHostAccountResponseBodyHostAccount {
	s.HostShareKeyName = &v
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

type GetHostAccountResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetHostAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetHostAccountResponse) SetStatusCode(v int32) *GetHostAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHostAccountResponse) SetBody(v *GetHostAccountResponseBody) *GetHostAccountResponse {
	s.Body = v
	return s
}

type GetHostGroupRequest struct {
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetHostGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHostGroupRequest) GoString() string {
	return s.String()
}

func (s *GetHostGroupRequest) SetHostGroupId(v string) *GetHostGroupRequest {
	s.HostGroupId = &v
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

type GetHostGroupResponseBody struct {
	HostGroup *GetHostGroupResponseBodyHostGroup `json:"HostGroup,omitempty" xml:"HostGroup,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetHostGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHostGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetHostGroupResponseBody) SetHostGroup(v *GetHostGroupResponseBodyHostGroup) *GetHostGroupResponseBody {
	s.HostGroup = v
	return s
}

func (s *GetHostGroupResponseBody) SetRequestId(v string) *GetHostGroupResponseBody {
	s.RequestId = &v
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetHostGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetHostGroupResponse) SetStatusCode(v int32) *GetHostGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHostGroupResponse) SetBody(v *GetHostGroupResponseBody) *GetHostGroupResponse {
	s.Body = v
	return s
}

type GetHostShareKeyRequest struct {
	HostShareKeyId *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetHostShareKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHostShareKeyRequest) GoString() string {
	return s.String()
}

func (s *GetHostShareKeyRequest) SetHostShareKeyId(v string) *GetHostShareKeyRequest {
	s.HostShareKeyId = &v
	return s
}

func (s *GetHostShareKeyRequest) SetInstanceId(v string) *GetHostShareKeyRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHostShareKeyRequest) SetRegionId(v string) *GetHostShareKeyRequest {
	s.RegionId = &v
	return s
}

type GetHostShareKeyResponseBody struct {
	HostShareKey *GetHostShareKeyResponseBodyHostShareKey `json:"HostShareKey,omitempty" xml:"HostShareKey,omitempty" type:"Struct"`
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetHostShareKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHostShareKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GetHostShareKeyResponseBody) SetHostShareKey(v *GetHostShareKeyResponseBodyHostShareKey) *GetHostShareKeyResponseBody {
	s.HostShareKey = v
	return s
}

func (s *GetHostShareKeyResponseBody) SetRequestId(v string) *GetHostShareKeyResponseBody {
	s.RequestId = &v
	return s
}

type GetHostShareKeyResponseBodyHostShareKey struct {
	HostShareKeyId        *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	HostShareKeyName      *string `json:"HostShareKeyName,omitempty" xml:"HostShareKeyName,omitempty"`
	LastModifyKeyAt       *int64  `json:"LastModifyKeyAt,omitempty" xml:"LastModifyKeyAt,omitempty"`
	PrivateKeyFingerPrint *string `json:"PrivateKeyFingerPrint,omitempty" xml:"PrivateKeyFingerPrint,omitempty"`
}

func (s GetHostShareKeyResponseBodyHostShareKey) String() string {
	return tea.Prettify(s)
}

func (s GetHostShareKeyResponseBodyHostShareKey) GoString() string {
	return s.String()
}

func (s *GetHostShareKeyResponseBodyHostShareKey) SetHostShareKeyId(v string) *GetHostShareKeyResponseBodyHostShareKey {
	s.HostShareKeyId = &v
	return s
}

func (s *GetHostShareKeyResponseBodyHostShareKey) SetHostShareKeyName(v string) *GetHostShareKeyResponseBodyHostShareKey {
	s.HostShareKeyName = &v
	return s
}

func (s *GetHostShareKeyResponseBodyHostShareKey) SetLastModifyKeyAt(v int64) *GetHostShareKeyResponseBodyHostShareKey {
	s.LastModifyKeyAt = &v
	return s
}

func (s *GetHostShareKeyResponseBodyHostShareKey) SetPrivateKeyFingerPrint(v string) *GetHostShareKeyResponseBodyHostShareKey {
	s.PrivateKeyFingerPrint = &v
	return s
}

type GetHostShareKeyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetHostShareKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHostShareKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHostShareKeyResponse) GoString() string {
	return s.String()
}

func (s *GetHostShareKeyResponse) SetHeaders(v map[string]*string) *GetHostShareKeyResponse {
	s.Headers = v
	return s
}

func (s *GetHostShareKeyResponse) SetStatusCode(v int32) *GetHostShareKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHostShareKeyResponse) SetBody(v *GetHostShareKeyResponseBody) *GetHostShareKeyResponse {
	s.Body = v
	return s
}

type GetInstanceADAuthServerRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetInstanceADAuthServerRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceADAuthServerRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceADAuthServerRequest) SetInstanceId(v string) *GetInstanceADAuthServerRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceADAuthServerRequest) SetRegionId(v string) *GetInstanceADAuthServerRequest {
	s.RegionId = &v
	return s
}

type GetInstanceADAuthServerResponseBody struct {
	AD        *GetInstanceADAuthServerResponseBodyAD `json:"AD,omitempty" xml:"AD,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInstanceADAuthServerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceADAuthServerResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceADAuthServerResponseBody) SetAD(v *GetInstanceADAuthServerResponseBodyAD) *GetInstanceADAuthServerResponseBody {
	s.AD = v
	return s
}

func (s *GetInstanceADAuthServerResponseBody) SetRequestId(v string) *GetInstanceADAuthServerResponseBody {
	s.RequestId = &v
	return s
}

type GetInstanceADAuthServerResponseBodyAD struct {
	Account       *string `json:"Account,omitempty" xml:"Account,omitempty"`
	BaseDN        *string `json:"BaseDN,omitempty" xml:"BaseDN,omitempty"`
	Domain        *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	EmailMapping  *string `json:"EmailMapping,omitempty" xml:"EmailMapping,omitempty"`
	Filter        *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	HasPassword   *bool   `json:"HasPassword,omitempty" xml:"HasPassword,omitempty"`
	IsSSL         *bool   `json:"IsSSL,omitempty" xml:"IsSSL,omitempty"`
	MobileMapping *string `json:"MobileMapping,omitempty" xml:"MobileMapping,omitempty"`
	NameMapping   *string `json:"NameMapping,omitempty" xml:"NameMapping,omitempty"`
	Port          *int64  `json:"Port,omitempty" xml:"Port,omitempty"`
	Server        *string `json:"Server,omitempty" xml:"Server,omitempty"`
	StandbyServer *string `json:"StandbyServer,omitempty" xml:"StandbyServer,omitempty"`
}

func (s GetInstanceADAuthServerResponseBodyAD) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceADAuthServerResponseBodyAD) GoString() string {
	return s.String()
}

func (s *GetInstanceADAuthServerResponseBodyAD) SetAccount(v string) *GetInstanceADAuthServerResponseBodyAD {
	s.Account = &v
	return s
}

func (s *GetInstanceADAuthServerResponseBodyAD) SetBaseDN(v string) *GetInstanceADAuthServerResponseBodyAD {
	s.BaseDN = &v
	return s
}

func (s *GetInstanceADAuthServerResponseBodyAD) SetDomain(v string) *GetInstanceADAuthServerResponseBodyAD {
	s.Domain = &v
	return s
}

func (s *GetInstanceADAuthServerResponseBodyAD) SetEmailMapping(v string) *GetInstanceADAuthServerResponseBodyAD {
	s.EmailMapping = &v
	return s
}

func (s *GetInstanceADAuthServerResponseBodyAD) SetFilter(v string) *GetInstanceADAuthServerResponseBodyAD {
	s.Filter = &v
	return s
}

func (s *GetInstanceADAuthServerResponseBodyAD) SetHasPassword(v bool) *GetInstanceADAuthServerResponseBodyAD {
	s.HasPassword = &v
	return s
}

func (s *GetInstanceADAuthServerResponseBodyAD) SetIsSSL(v bool) *GetInstanceADAuthServerResponseBodyAD {
	s.IsSSL = &v
	return s
}

func (s *GetInstanceADAuthServerResponseBodyAD) SetMobileMapping(v string) *GetInstanceADAuthServerResponseBodyAD {
	s.MobileMapping = &v
	return s
}

func (s *GetInstanceADAuthServerResponseBodyAD) SetNameMapping(v string) *GetInstanceADAuthServerResponseBodyAD {
	s.NameMapping = &v
	return s
}

func (s *GetInstanceADAuthServerResponseBodyAD) SetPort(v int64) *GetInstanceADAuthServerResponseBodyAD {
	s.Port = &v
	return s
}

func (s *GetInstanceADAuthServerResponseBodyAD) SetServer(v string) *GetInstanceADAuthServerResponseBodyAD {
	s.Server = &v
	return s
}

func (s *GetInstanceADAuthServerResponseBodyAD) SetStandbyServer(v string) *GetInstanceADAuthServerResponseBodyAD {
	s.StandbyServer = &v
	return s
}

type GetInstanceADAuthServerResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetInstanceADAuthServerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInstanceADAuthServerResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceADAuthServerResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceADAuthServerResponse) SetHeaders(v map[string]*string) *GetInstanceADAuthServerResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceADAuthServerResponse) SetStatusCode(v int32) *GetInstanceADAuthServerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceADAuthServerResponse) SetBody(v *GetInstanceADAuthServerResponseBody) *GetInstanceADAuthServerResponse {
	s.Body = v
	return s
}

type GetInstanceLDAPAuthServerRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetInstanceLDAPAuthServerRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceLDAPAuthServerRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceLDAPAuthServerRequest) SetInstanceId(v string) *GetInstanceLDAPAuthServerRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceLDAPAuthServerRequest) SetRegionId(v string) *GetInstanceLDAPAuthServerRequest {
	s.RegionId = &v
	return s
}

type GetInstanceLDAPAuthServerResponseBody struct {
	LDAP      *GetInstanceLDAPAuthServerResponseBodyLDAP `json:"LDAP,omitempty" xml:"LDAP,omitempty" type:"Struct"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInstanceLDAPAuthServerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceLDAPAuthServerResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceLDAPAuthServerResponseBody) SetLDAP(v *GetInstanceLDAPAuthServerResponseBodyLDAP) *GetInstanceLDAPAuthServerResponseBody {
	s.LDAP = v
	return s
}

func (s *GetInstanceLDAPAuthServerResponseBody) SetRequestId(v string) *GetInstanceLDAPAuthServerResponseBody {
	s.RequestId = &v
	return s
}

type GetInstanceLDAPAuthServerResponseBodyLDAP struct {
	Account          *string `json:"Account,omitempty" xml:"Account,omitempty"`
	BaseDN           *string `json:"BaseDN,omitempty" xml:"BaseDN,omitempty"`
	EmailMapping     *string `json:"EmailMapping,omitempty" xml:"EmailMapping,omitempty"`
	Filter           *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	HasPassword      *string `json:"HasPassword,omitempty" xml:"HasPassword,omitempty"`
	IsSSL            *bool   `json:"IsSSL,omitempty" xml:"IsSSL,omitempty"`
	LoginNameMapping *string `json:"LoginNameMapping,omitempty" xml:"LoginNameMapping,omitempty"`
	MobileMapping    *string `json:"MobileMapping,omitempty" xml:"MobileMapping,omitempty"`
	NameMapping      *string `json:"NameMapping,omitempty" xml:"NameMapping,omitempty"`
	Port             *int64  `json:"Port,omitempty" xml:"Port,omitempty"`
	Server           *string `json:"Server,omitempty" xml:"Server,omitempty"`
	StandbyServer    *string `json:"StandbyServer,omitempty" xml:"StandbyServer,omitempty"`
}

func (s GetInstanceLDAPAuthServerResponseBodyLDAP) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceLDAPAuthServerResponseBodyLDAP) GoString() string {
	return s.String()
}

func (s *GetInstanceLDAPAuthServerResponseBodyLDAP) SetAccount(v string) *GetInstanceLDAPAuthServerResponseBodyLDAP {
	s.Account = &v
	return s
}

func (s *GetInstanceLDAPAuthServerResponseBodyLDAP) SetBaseDN(v string) *GetInstanceLDAPAuthServerResponseBodyLDAP {
	s.BaseDN = &v
	return s
}

func (s *GetInstanceLDAPAuthServerResponseBodyLDAP) SetEmailMapping(v string) *GetInstanceLDAPAuthServerResponseBodyLDAP {
	s.EmailMapping = &v
	return s
}

func (s *GetInstanceLDAPAuthServerResponseBodyLDAP) SetFilter(v string) *GetInstanceLDAPAuthServerResponseBodyLDAP {
	s.Filter = &v
	return s
}

func (s *GetInstanceLDAPAuthServerResponseBodyLDAP) SetHasPassword(v string) *GetInstanceLDAPAuthServerResponseBodyLDAP {
	s.HasPassword = &v
	return s
}

func (s *GetInstanceLDAPAuthServerResponseBodyLDAP) SetIsSSL(v bool) *GetInstanceLDAPAuthServerResponseBodyLDAP {
	s.IsSSL = &v
	return s
}

func (s *GetInstanceLDAPAuthServerResponseBodyLDAP) SetLoginNameMapping(v string) *GetInstanceLDAPAuthServerResponseBodyLDAP {
	s.LoginNameMapping = &v
	return s
}

func (s *GetInstanceLDAPAuthServerResponseBodyLDAP) SetMobileMapping(v string) *GetInstanceLDAPAuthServerResponseBodyLDAP {
	s.MobileMapping = &v
	return s
}

func (s *GetInstanceLDAPAuthServerResponseBodyLDAP) SetNameMapping(v string) *GetInstanceLDAPAuthServerResponseBodyLDAP {
	s.NameMapping = &v
	return s
}

func (s *GetInstanceLDAPAuthServerResponseBodyLDAP) SetPort(v int64) *GetInstanceLDAPAuthServerResponseBodyLDAP {
	s.Port = &v
	return s
}

func (s *GetInstanceLDAPAuthServerResponseBodyLDAP) SetServer(v string) *GetInstanceLDAPAuthServerResponseBodyLDAP {
	s.Server = &v
	return s
}

func (s *GetInstanceLDAPAuthServerResponseBodyLDAP) SetStandbyServer(v string) *GetInstanceLDAPAuthServerResponseBodyLDAP {
	s.StandbyServer = &v
	return s
}

type GetInstanceLDAPAuthServerResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetInstanceLDAPAuthServerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInstanceLDAPAuthServerResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceLDAPAuthServerResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceLDAPAuthServerResponse) SetHeaders(v map[string]*string) *GetInstanceLDAPAuthServerResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceLDAPAuthServerResponse) SetStatusCode(v int32) *GetInstanceLDAPAuthServerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceLDAPAuthServerResponse) SetBody(v *GetInstanceLDAPAuthServerResponseBody) *GetInstanceLDAPAuthServerResponse {
	s.Body = v
	return s
}

type GetInstanceTwoFactorRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetInstanceTwoFactorRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceTwoFactorRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceTwoFactorRequest) SetInstanceId(v string) *GetInstanceTwoFactorRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceTwoFactorRequest) SetRegionId(v string) *GetInstanceTwoFactorRequest {
	s.RegionId = &v
	return s
}

type GetInstanceTwoFactorResponseBody struct {
	Config    *GetInstanceTwoFactorResponseBodyConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInstanceTwoFactorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceTwoFactorResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceTwoFactorResponseBody) SetConfig(v *GetInstanceTwoFactorResponseBodyConfig) *GetInstanceTwoFactorResponseBody {
	s.Config = v
	return s
}

func (s *GetInstanceTwoFactorResponseBody) SetRequestId(v string) *GetInstanceTwoFactorResponseBody {
	s.RequestId = &v
	return s
}

type GetInstanceTwoFactorResponseBodyConfig struct {
	DingTalkConfig    *GetInstanceTwoFactorResponseBodyConfigDingTalkConfig `json:"DingTalkConfig,omitempty" xml:"DingTalkConfig,omitempty" type:"Struct"`
	EnableTwoFactor   *bool                                                 `json:"EnableTwoFactor,omitempty" xml:"EnableTwoFactor,omitempty"`
	MessageLanguage   *string                                               `json:"MessageLanguage,omitempty" xml:"MessageLanguage,omitempty"`
	SkipTwoFactorTime *int64                                                `json:"SkipTwoFactorTime,omitempty" xml:"SkipTwoFactorTime,omitempty"`
	TwoFactorMethods  []*string                                             `json:"TwoFactorMethods,omitempty" xml:"TwoFactorMethods,omitempty" type:"Repeated"`
}

func (s GetInstanceTwoFactorResponseBodyConfig) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceTwoFactorResponseBodyConfig) GoString() string {
	return s.String()
}

func (s *GetInstanceTwoFactorResponseBodyConfig) SetDingTalkConfig(v *GetInstanceTwoFactorResponseBodyConfigDingTalkConfig) *GetInstanceTwoFactorResponseBodyConfig {
	s.DingTalkConfig = v
	return s
}

func (s *GetInstanceTwoFactorResponseBodyConfig) SetEnableTwoFactor(v bool) *GetInstanceTwoFactorResponseBodyConfig {
	s.EnableTwoFactor = &v
	return s
}

func (s *GetInstanceTwoFactorResponseBodyConfig) SetMessageLanguage(v string) *GetInstanceTwoFactorResponseBodyConfig {
	s.MessageLanguage = &v
	return s
}

func (s *GetInstanceTwoFactorResponseBodyConfig) SetSkipTwoFactorTime(v int64) *GetInstanceTwoFactorResponseBodyConfig {
	s.SkipTwoFactorTime = &v
	return s
}

func (s *GetInstanceTwoFactorResponseBodyConfig) SetTwoFactorMethods(v []*string) *GetInstanceTwoFactorResponseBodyConfig {
	s.TwoFactorMethods = v
	return s
}

type GetInstanceTwoFactorResponseBodyConfigDingTalkConfig struct {
	AgentId      *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	AppKey       *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	HasAppSecret *bool   `json:"HasAppSecret,omitempty" xml:"HasAppSecret,omitempty"`
}

func (s GetInstanceTwoFactorResponseBodyConfigDingTalkConfig) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceTwoFactorResponseBodyConfigDingTalkConfig) GoString() string {
	return s.String()
}

func (s *GetInstanceTwoFactorResponseBodyConfigDingTalkConfig) SetAgentId(v string) *GetInstanceTwoFactorResponseBodyConfigDingTalkConfig {
	s.AgentId = &v
	return s
}

func (s *GetInstanceTwoFactorResponseBodyConfigDingTalkConfig) SetAppKey(v string) *GetInstanceTwoFactorResponseBodyConfigDingTalkConfig {
	s.AppKey = &v
	return s
}

func (s *GetInstanceTwoFactorResponseBodyConfigDingTalkConfig) SetHasAppSecret(v bool) *GetInstanceTwoFactorResponseBodyConfigDingTalkConfig {
	s.HasAppSecret = &v
	return s
}

type GetInstanceTwoFactorResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetInstanceTwoFactorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInstanceTwoFactorResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceTwoFactorResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceTwoFactorResponse) SetHeaders(v map[string]*string) *GetInstanceTwoFactorResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceTwoFactorResponse) SetStatusCode(v int32) *GetInstanceTwoFactorResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceTwoFactorResponse) SetBody(v *GetInstanceTwoFactorResponseBody) *GetInstanceTwoFactorResponse {
	s.Body = v
	return s
}

type GetUserRequest struct {
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
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	User      *GetUserResponseBodyUser `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
}

func (s GetUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResponseBody) SetRequestId(v string) *GetUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserResponseBody) SetUser(v *GetUserResponseBodyUser) *GetUserResponseBody {
	s.User = v
	return s
}

type GetUserResponseBodyUser struct {
	Comment           *string   `json:"Comment,omitempty" xml:"Comment,omitempty"`
	DisplayName       *string   `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Email             *string   `json:"Email,omitempty" xml:"Email,omitempty"`
	Mobile            *string   `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	MobileCountryCode *string   `json:"MobileCountryCode,omitempty" xml:"MobileCountryCode,omitempty"`
	Source            *string   `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceUserId      *string   `json:"SourceUserId,omitempty" xml:"SourceUserId,omitempty"`
	UserId            *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName          *string   `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserState         []*string `json:"UserState,omitempty" xml:"UserState,omitempty" type:"Repeated"`
}

func (s GetUserResponseBodyUser) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseBodyUser) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyUser) SetComment(v string) *GetUserResponseBodyUser {
	s.Comment = &v
	return s
}

func (s *GetUserResponseBodyUser) SetDisplayName(v string) *GetUserResponseBodyUser {
	s.DisplayName = &v
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

func (s *GetUserResponseBodyUser) SetSource(v string) *GetUserResponseBodyUser {
	s.Source = &v
	return s
}

func (s *GetUserResponseBodyUser) SetSourceUserId(v string) *GetUserResponseBodyUser {
	s.SourceUserId = &v
	return s
}

func (s *GetUserResponseBodyUser) SetUserId(v string) *GetUserResponseBodyUser {
	s.UserId = &v
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

type GetUserResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetUserResponse) SetStatusCode(v int32) *GetUserResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserResponse) SetBody(v *GetUserResponseBody) *GetUserResponse {
	s.Body = v
	return s
}

type GetUserGroupRequest struct {
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetUserGroupResponse) SetStatusCode(v int32) *GetUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserGroupResponse) SetBody(v *GetUserGroupResponseBody) *GetUserGroupResponse {
	s.Body = v
	return s
}

type ListHostAccountsRequest struct {
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	HostId          *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber      *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProtocolName    *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListHostAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHostAccountsRequest) GoString() string {
	return s.String()
}

func (s *ListHostAccountsRequest) SetHostAccountName(v string) *ListHostAccountsRequest {
	s.HostAccountName = &v
	return s
}

func (s *ListHostAccountsRequest) SetHostId(v string) *ListHostAccountsRequest {
	s.HostId = &v
	return s
}

func (s *ListHostAccountsRequest) SetInstanceId(v string) *ListHostAccountsRequest {
	s.InstanceId = &v
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

func (s *ListHostAccountsRequest) SetProtocolName(v string) *ListHostAccountsRequest {
	s.ProtocolName = &v
	return s
}

func (s *ListHostAccountsRequest) SetRegionId(v string) *ListHostAccountsRequest {
	s.RegionId = &v
	return s
}

type ListHostAccountsResponseBody struct {
	HostAccounts []*ListHostAccountsResponseBodyHostAccounts `json:"HostAccounts,omitempty" xml:"HostAccounts,omitempty" type:"Repeated"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHostAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHostAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *ListHostAccountsResponseBody) SetHostAccounts(v []*ListHostAccountsResponseBodyHostAccounts) *ListHostAccountsResponseBody {
	s.HostAccounts = v
	return s
}

func (s *ListHostAccountsResponseBody) SetRequestId(v string) *ListHostAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHostAccountsResponseBody) SetTotalCount(v int32) *ListHostAccountsResponseBody {
	s.TotalCount = &v
	return s
}

type ListHostAccountsResponseBodyHostAccounts struct {
	HasPassword           *bool   `json:"HasPassword,omitempty" xml:"HasPassword,omitempty"`
	HostAccountId         *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	HostAccountName       *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	HostId                *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	HostShareKeyId        *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	HostShareKeyName      *string `json:"HostShareKeyName,omitempty" xml:"HostShareKeyName,omitempty"`
	PrivateKeyFingerprint *string `json:"PrivateKeyFingerprint,omitempty" xml:"PrivateKeyFingerprint,omitempty"`
	ProtocolName          *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
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

func (s *ListHostAccountsResponseBodyHostAccounts) SetHostAccountId(v string) *ListHostAccountsResponseBodyHostAccounts {
	s.HostAccountId = &v
	return s
}

func (s *ListHostAccountsResponseBodyHostAccounts) SetHostAccountName(v string) *ListHostAccountsResponseBodyHostAccounts {
	s.HostAccountName = &v
	return s
}

func (s *ListHostAccountsResponseBodyHostAccounts) SetHostId(v string) *ListHostAccountsResponseBodyHostAccounts {
	s.HostId = &v
	return s
}

func (s *ListHostAccountsResponseBodyHostAccounts) SetHostShareKeyId(v string) *ListHostAccountsResponseBodyHostAccounts {
	s.HostShareKeyId = &v
	return s
}

func (s *ListHostAccountsResponseBodyHostAccounts) SetHostShareKeyName(v string) *ListHostAccountsResponseBodyHostAccounts {
	s.HostShareKeyName = &v
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

type ListHostAccountsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListHostAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListHostAccountsResponse) SetStatusCode(v int32) *ListHostAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHostAccountsResponse) SetBody(v *ListHostAccountsResponseBody) *ListHostAccountsResponse {
	s.Body = v
	return s
}

type ListHostAccountsForHostShareKeyRequest struct {
	HostShareKeyId *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber     *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListHostAccountsForHostShareKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHostAccountsForHostShareKeyRequest) GoString() string {
	return s.String()
}

func (s *ListHostAccountsForHostShareKeyRequest) SetHostShareKeyId(v string) *ListHostAccountsForHostShareKeyRequest {
	s.HostShareKeyId = &v
	return s
}

func (s *ListHostAccountsForHostShareKeyRequest) SetInstanceId(v string) *ListHostAccountsForHostShareKeyRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHostAccountsForHostShareKeyRequest) SetPageNumber(v string) *ListHostAccountsForHostShareKeyRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHostAccountsForHostShareKeyRequest) SetPageSize(v string) *ListHostAccountsForHostShareKeyRequest {
	s.PageSize = &v
	return s
}

func (s *ListHostAccountsForHostShareKeyRequest) SetRegionId(v string) *ListHostAccountsForHostShareKeyRequest {
	s.RegionId = &v
	return s
}

type ListHostAccountsForHostShareKeyResponseBody struct {
	HostAccounts []*ListHostAccountsForHostShareKeyResponseBodyHostAccounts `json:"HostAccounts,omitempty" xml:"HostAccounts,omitempty" type:"Repeated"`
	RequestId    *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int64                                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHostAccountsForHostShareKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHostAccountsForHostShareKeyResponseBody) GoString() string {
	return s.String()
}

func (s *ListHostAccountsForHostShareKeyResponseBody) SetHostAccounts(v []*ListHostAccountsForHostShareKeyResponseBodyHostAccounts) *ListHostAccountsForHostShareKeyResponseBody {
	s.HostAccounts = v
	return s
}

func (s *ListHostAccountsForHostShareKeyResponseBody) SetRequestId(v string) *ListHostAccountsForHostShareKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHostAccountsForHostShareKeyResponseBody) SetTotalCount(v int64) *ListHostAccountsForHostShareKeyResponseBody {
	s.TotalCount = &v
	return s
}

type ListHostAccountsForHostShareKeyResponseBodyHostAccounts struct {
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	HostId          *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	HostsAccountId  *string `json:"HostsAccountId,omitempty" xml:"HostsAccountId,omitempty"`
	ProtocolName    *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
}

func (s ListHostAccountsForHostShareKeyResponseBodyHostAccounts) String() string {
	return tea.Prettify(s)
}

func (s ListHostAccountsForHostShareKeyResponseBodyHostAccounts) GoString() string {
	return s.String()
}

func (s *ListHostAccountsForHostShareKeyResponseBodyHostAccounts) SetHostAccountName(v string) *ListHostAccountsForHostShareKeyResponseBodyHostAccounts {
	s.HostAccountName = &v
	return s
}

func (s *ListHostAccountsForHostShareKeyResponseBodyHostAccounts) SetHostId(v string) *ListHostAccountsForHostShareKeyResponseBodyHostAccounts {
	s.HostId = &v
	return s
}

func (s *ListHostAccountsForHostShareKeyResponseBodyHostAccounts) SetHostsAccountId(v string) *ListHostAccountsForHostShareKeyResponseBodyHostAccounts {
	s.HostsAccountId = &v
	return s
}

func (s *ListHostAccountsForHostShareKeyResponseBodyHostAccounts) SetProtocolName(v string) *ListHostAccountsForHostShareKeyResponseBodyHostAccounts {
	s.ProtocolName = &v
	return s
}

type ListHostAccountsForHostShareKeyResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListHostAccountsForHostShareKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListHostAccountsForHostShareKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHostAccountsForHostShareKeyResponse) GoString() string {
	return s.String()
}

func (s *ListHostAccountsForHostShareKeyResponse) SetHeaders(v map[string]*string) *ListHostAccountsForHostShareKeyResponse {
	s.Headers = v
	return s
}

func (s *ListHostAccountsForHostShareKeyResponse) SetStatusCode(v int32) *ListHostAccountsForHostShareKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHostAccountsForHostShareKeyResponse) SetBody(v *ListHostAccountsForHostShareKeyResponseBody) *ListHostAccountsForHostShareKeyResponse {
	s.Body = v
	return s
}

type ListHostAccountsForUserRequest struct {
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	HostId          *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber      *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserId          *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListHostAccountsForUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHostAccountsForUserRequest) GoString() string {
	return s.String()
}

func (s *ListHostAccountsForUserRequest) SetHostAccountName(v string) *ListHostAccountsForUserRequest {
	s.HostAccountName = &v
	return s
}

func (s *ListHostAccountsForUserRequest) SetHostId(v string) *ListHostAccountsForUserRequest {
	s.HostId = &v
	return s
}

func (s *ListHostAccountsForUserRequest) SetInstanceId(v string) *ListHostAccountsForUserRequest {
	s.InstanceId = &v
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

func (s *ListHostAccountsForUserRequest) SetRegionId(v string) *ListHostAccountsForUserRequest {
	s.RegionId = &v
	return s
}

func (s *ListHostAccountsForUserRequest) SetUserId(v string) *ListHostAccountsForUserRequest {
	s.UserId = &v
	return s
}

type ListHostAccountsForUserResponseBody struct {
	HostAccounts []*ListHostAccountsForUserResponseBodyHostAccounts `json:"HostAccounts,omitempty" xml:"HostAccounts,omitempty" type:"Repeated"`
	RequestId    *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int32                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHostAccountsForUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHostAccountsForUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListHostAccountsForUserResponseBody) SetHostAccounts(v []*ListHostAccountsForUserResponseBodyHostAccounts) *ListHostAccountsForUserResponseBody {
	s.HostAccounts = v
	return s
}

func (s *ListHostAccountsForUserResponseBody) SetRequestId(v string) *ListHostAccountsForUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHostAccountsForUserResponseBody) SetTotalCount(v int32) *ListHostAccountsForUserResponseBody {
	s.TotalCount = &v
	return s
}

type ListHostAccountsForUserResponseBodyHostAccounts struct {
	HostAccountId   *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	HostId          *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	IsAuthorized    *bool   `json:"IsAuthorized,omitempty" xml:"IsAuthorized,omitempty"`
	ProtocolName    *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
}

func (s ListHostAccountsForUserResponseBodyHostAccounts) String() string {
	return tea.Prettify(s)
}

func (s ListHostAccountsForUserResponseBodyHostAccounts) GoString() string {
	return s.String()
}

func (s *ListHostAccountsForUserResponseBodyHostAccounts) SetHostAccountId(v string) *ListHostAccountsForUserResponseBodyHostAccounts {
	s.HostAccountId = &v
	return s
}

func (s *ListHostAccountsForUserResponseBodyHostAccounts) SetHostAccountName(v string) *ListHostAccountsForUserResponseBodyHostAccounts {
	s.HostAccountName = &v
	return s
}

func (s *ListHostAccountsForUserResponseBodyHostAccounts) SetHostId(v string) *ListHostAccountsForUserResponseBodyHostAccounts {
	s.HostId = &v
	return s
}

func (s *ListHostAccountsForUserResponseBodyHostAccounts) SetIsAuthorized(v bool) *ListHostAccountsForUserResponseBodyHostAccounts {
	s.IsAuthorized = &v
	return s
}

func (s *ListHostAccountsForUserResponseBodyHostAccounts) SetProtocolName(v string) *ListHostAccountsForUserResponseBodyHostAccounts {
	s.ProtocolName = &v
	return s
}

type ListHostAccountsForUserResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListHostAccountsForUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListHostAccountsForUserResponse) SetStatusCode(v int32) *ListHostAccountsForUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHostAccountsForUserResponse) SetBody(v *ListHostAccountsForUserResponseBody) *ListHostAccountsForUserResponse {
	s.Body = v
	return s
}

type ListHostAccountsForUserGroupRequest struct {
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	HostId          *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber      *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserGroupId     *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s ListHostAccountsForUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHostAccountsForUserGroupRequest) GoString() string {
	return s.String()
}

func (s *ListHostAccountsForUserGroupRequest) SetHostAccountName(v string) *ListHostAccountsForUserGroupRequest {
	s.HostAccountName = &v
	return s
}

func (s *ListHostAccountsForUserGroupRequest) SetHostId(v string) *ListHostAccountsForUserGroupRequest {
	s.HostId = &v
	return s
}

func (s *ListHostAccountsForUserGroupRequest) SetInstanceId(v string) *ListHostAccountsForUserGroupRequest {
	s.InstanceId = &v
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

func (s *ListHostAccountsForUserGroupRequest) SetRegionId(v string) *ListHostAccountsForUserGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ListHostAccountsForUserGroupRequest) SetUserGroupId(v string) *ListHostAccountsForUserGroupRequest {
	s.UserGroupId = &v
	return s
}

type ListHostAccountsForUserGroupResponseBody struct {
	HostAccounts []*ListHostAccountsForUserGroupResponseBodyHostAccounts `json:"HostAccounts,omitempty" xml:"HostAccounts,omitempty" type:"Repeated"`
	RequestId    *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int32                                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHostAccountsForUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHostAccountsForUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListHostAccountsForUserGroupResponseBody) SetHostAccounts(v []*ListHostAccountsForUserGroupResponseBodyHostAccounts) *ListHostAccountsForUserGroupResponseBody {
	s.HostAccounts = v
	return s
}

func (s *ListHostAccountsForUserGroupResponseBody) SetRequestId(v string) *ListHostAccountsForUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHostAccountsForUserGroupResponseBody) SetTotalCount(v int32) *ListHostAccountsForUserGroupResponseBody {
	s.TotalCount = &v
	return s
}

type ListHostAccountsForUserGroupResponseBodyHostAccounts struct {
	HostAccountId   *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	HostId          *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	IsAuthorized    *bool   `json:"IsAuthorized,omitempty" xml:"IsAuthorized,omitempty"`
	ProtocolName    *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
}

func (s ListHostAccountsForUserGroupResponseBodyHostAccounts) String() string {
	return tea.Prettify(s)
}

func (s ListHostAccountsForUserGroupResponseBodyHostAccounts) GoString() string {
	return s.String()
}

func (s *ListHostAccountsForUserGroupResponseBodyHostAccounts) SetHostAccountId(v string) *ListHostAccountsForUserGroupResponseBodyHostAccounts {
	s.HostAccountId = &v
	return s
}

func (s *ListHostAccountsForUserGroupResponseBodyHostAccounts) SetHostAccountName(v string) *ListHostAccountsForUserGroupResponseBodyHostAccounts {
	s.HostAccountName = &v
	return s
}

func (s *ListHostAccountsForUserGroupResponseBodyHostAccounts) SetHostId(v string) *ListHostAccountsForUserGroupResponseBodyHostAccounts {
	s.HostId = &v
	return s
}

func (s *ListHostAccountsForUserGroupResponseBodyHostAccounts) SetIsAuthorized(v bool) *ListHostAccountsForUserGroupResponseBodyHostAccounts {
	s.IsAuthorized = &v
	return s
}

func (s *ListHostAccountsForUserGroupResponseBodyHostAccounts) SetProtocolName(v string) *ListHostAccountsForUserGroupResponseBodyHostAccounts {
	s.ProtocolName = &v
	return s
}

type ListHostAccountsForUserGroupResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListHostAccountsForUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListHostAccountsForUserGroupResponse) SetStatusCode(v int32) *ListHostAccountsForUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHostAccountsForUserGroupResponse) SetBody(v *ListHostAccountsForUserGroupResponseBody) *ListHostAccountsForUserGroupResponse {
	s.Body = v
	return s
}

type ListHostGroupAccountNamesForUserRequest struct {
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListHostGroupAccountNamesForUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHostGroupAccountNamesForUserRequest) GoString() string {
	return s.String()
}

func (s *ListHostGroupAccountNamesForUserRequest) SetHostGroupId(v string) *ListHostGroupAccountNamesForUserRequest {
	s.HostGroupId = &v
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

type ListHostGroupAccountNamesForUserResponseBody struct {
	HostAccountNames []*string `json:"HostAccountNames,omitempty" xml:"HostAccountNames,omitempty" type:"Repeated"`
	RequestId        *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListHostGroupAccountNamesForUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHostGroupAccountNamesForUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListHostGroupAccountNamesForUserResponseBody) SetHostAccountNames(v []*string) *ListHostGroupAccountNamesForUserResponseBody {
	s.HostAccountNames = v
	return s
}

func (s *ListHostGroupAccountNamesForUserResponseBody) SetRequestId(v string) *ListHostGroupAccountNamesForUserResponseBody {
	s.RequestId = &v
	return s
}

type ListHostGroupAccountNamesForUserResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListHostGroupAccountNamesForUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListHostGroupAccountNamesForUserResponse) SetStatusCode(v int32) *ListHostGroupAccountNamesForUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHostGroupAccountNamesForUserResponse) SetBody(v *ListHostGroupAccountNamesForUserResponseBody) *ListHostGroupAccountNamesForUserResponse {
	s.Body = v
	return s
}

type ListHostGroupAccountNamesForUserGroupRequest struct {
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s ListHostGroupAccountNamesForUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHostGroupAccountNamesForUserGroupRequest) GoString() string {
	return s.String()
}

func (s *ListHostGroupAccountNamesForUserGroupRequest) SetHostGroupId(v string) *ListHostGroupAccountNamesForUserGroupRequest {
	s.HostGroupId = &v
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

type ListHostGroupAccountNamesForUserGroupResponseBody struct {
	HostAccountNames []*string `json:"HostAccountNames,omitempty" xml:"HostAccountNames,omitempty" type:"Repeated"`
	RequestId        *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListHostGroupAccountNamesForUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHostGroupAccountNamesForUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListHostGroupAccountNamesForUserGroupResponseBody) SetHostAccountNames(v []*string) *ListHostGroupAccountNamesForUserGroupResponseBody {
	s.HostAccountNames = v
	return s
}

func (s *ListHostGroupAccountNamesForUserGroupResponseBody) SetRequestId(v string) *ListHostGroupAccountNamesForUserGroupResponseBody {
	s.RequestId = &v
	return s
}

type ListHostGroupAccountNamesForUserGroupResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListHostGroupAccountNamesForUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListHostGroupAccountNamesForUserGroupResponse) SetStatusCode(v int32) *ListHostGroupAccountNamesForUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHostGroupAccountNamesForUserGroupResponse) SetBody(v *ListHostGroupAccountNamesForUserGroupResponseBody) *ListHostGroupAccountNamesForUserGroupResponse {
	s.Body = v
	return s
}

type ListHostGroupsRequest struct {
	HostGroupName *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber    *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListHostGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHostGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListHostGroupsRequest) SetHostGroupName(v string) *ListHostGroupsRequest {
	s.HostGroupName = &v
	return s
}

func (s *ListHostGroupsRequest) SetInstanceId(v string) *ListHostGroupsRequest {
	s.InstanceId = &v
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

func (s *ListHostGroupsRequest) SetRegionId(v string) *ListHostGroupsRequest {
	s.RegionId = &v
	return s
}

type ListHostGroupsResponseBody struct {
	HostGroups []*ListHostGroupsResponseBodyHostGroups `json:"HostGroups,omitempty" xml:"HostGroups,omitempty" type:"Repeated"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHostGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHostGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListHostGroupsResponseBody) SetHostGroups(v []*ListHostGroupsResponseBodyHostGroups) *ListHostGroupsResponseBody {
	s.HostGroups = v
	return s
}

func (s *ListHostGroupsResponseBody) SetRequestId(v string) *ListHostGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHostGroupsResponseBody) SetTotalCount(v int32) *ListHostGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type ListHostGroupsResponseBodyHostGroups struct {
	Comment       *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	HostGroupId   *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	HostGroupName *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
	MemberCount   *int32  `json:"MemberCount,omitempty" xml:"MemberCount,omitempty"`
}

func (s ListHostGroupsResponseBodyHostGroups) String() string {
	return tea.Prettify(s)
}

func (s ListHostGroupsResponseBodyHostGroups) GoString() string {
	return s.String()
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

func (s *ListHostGroupsResponseBodyHostGroups) SetMemberCount(v int32) *ListHostGroupsResponseBodyHostGroups {
	s.MemberCount = &v
	return s
}

type ListHostGroupsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListHostGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListHostGroupsResponse) SetStatusCode(v int32) *ListHostGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHostGroupsResponse) SetBody(v *ListHostGroupsResponseBody) *ListHostGroupsResponse {
	s.Body = v
	return s
}

type ListHostGroupsForUserRequest struct {
	HostGroupName *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Mode          *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	PageNumber    *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserId        *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListHostGroupsForUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHostGroupsForUserRequest) GoString() string {
	return s.String()
}

func (s *ListHostGroupsForUserRequest) SetHostGroupName(v string) *ListHostGroupsForUserRequest {
	s.HostGroupName = &v
	return s
}

func (s *ListHostGroupsForUserRequest) SetInstanceId(v string) *ListHostGroupsForUserRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHostGroupsForUserRequest) SetMode(v string) *ListHostGroupsForUserRequest {
	s.Mode = &v
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

func (s *ListHostGroupsForUserRequest) SetRegionId(v string) *ListHostGroupsForUserRequest {
	s.RegionId = &v
	return s
}

func (s *ListHostGroupsForUserRequest) SetUserId(v string) *ListHostGroupsForUserRequest {
	s.UserId = &v
	return s
}

type ListHostGroupsForUserResponseBody struct {
	HostGroups []*ListHostGroupsForUserResponseBodyHostGroups `json:"HostGroups,omitempty" xml:"HostGroups,omitempty" type:"Repeated"`
	RequestId  *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHostGroupsForUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHostGroupsForUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListHostGroupsForUserResponseBody) SetHostGroups(v []*ListHostGroupsForUserResponseBodyHostGroups) *ListHostGroupsForUserResponseBody {
	s.HostGroups = v
	return s
}

func (s *ListHostGroupsForUserResponseBody) SetRequestId(v string) *ListHostGroupsForUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHostGroupsForUserResponseBody) SetTotalCount(v int32) *ListHostGroupsForUserResponseBody {
	s.TotalCount = &v
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListHostGroupsForUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListHostGroupsForUserResponse) SetStatusCode(v int32) *ListHostGroupsForUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHostGroupsForUserResponse) SetBody(v *ListHostGroupsForUserResponseBody) *ListHostGroupsForUserResponse {
	s.Body = v
	return s
}

type ListHostGroupsForUserGroupRequest struct {
	HostGroupName *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Mode          *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	PageNumber    *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserGroupId   *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s ListHostGroupsForUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHostGroupsForUserGroupRequest) GoString() string {
	return s.String()
}

func (s *ListHostGroupsForUserGroupRequest) SetHostGroupName(v string) *ListHostGroupsForUserGroupRequest {
	s.HostGroupName = &v
	return s
}

func (s *ListHostGroupsForUserGroupRequest) SetInstanceId(v string) *ListHostGroupsForUserGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHostGroupsForUserGroupRequest) SetMode(v string) *ListHostGroupsForUserGroupRequest {
	s.Mode = &v
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

func (s *ListHostGroupsForUserGroupRequest) SetRegionId(v string) *ListHostGroupsForUserGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ListHostGroupsForUserGroupRequest) SetUserGroupId(v string) *ListHostGroupsForUserGroupRequest {
	s.UserGroupId = &v
	return s
}

type ListHostGroupsForUserGroupResponseBody struct {
	HostGroups []*ListHostGroupsForUserGroupResponseBodyHostGroups `json:"HostGroups,omitempty" xml:"HostGroups,omitempty" type:"Repeated"`
	RequestId  *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHostGroupsForUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHostGroupsForUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListHostGroupsForUserGroupResponseBody) SetHostGroups(v []*ListHostGroupsForUserGroupResponseBodyHostGroups) *ListHostGroupsForUserGroupResponseBody {
	s.HostGroups = v
	return s
}

func (s *ListHostGroupsForUserGroupResponseBody) SetRequestId(v string) *ListHostGroupsForUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHostGroupsForUserGroupResponseBody) SetTotalCount(v int32) *ListHostGroupsForUserGroupResponseBody {
	s.TotalCount = &v
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListHostGroupsForUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListHostGroupsForUserGroupResponse) SetStatusCode(v int32) *ListHostGroupsForUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHostGroupsForUserGroupResponse) SetBody(v *ListHostGroupsForUserGroupResponseBody) *ListHostGroupsForUserGroupResponse {
	s.Body = v
	return s
}

type ListHostShareKeysRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListHostShareKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHostShareKeysRequest) GoString() string {
	return s.String()
}

func (s *ListHostShareKeysRequest) SetInstanceId(v string) *ListHostShareKeysRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHostShareKeysRequest) SetPageNumber(v string) *ListHostShareKeysRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHostShareKeysRequest) SetPageSize(v string) *ListHostShareKeysRequest {
	s.PageSize = &v
	return s
}

func (s *ListHostShareKeysRequest) SetRegionId(v string) *ListHostShareKeysRequest {
	s.RegionId = &v
	return s
}

type ListHostShareKeysResponseBody struct {
	HostShareKeys []*ListHostShareKeysResponseBodyHostShareKeys `json:"HostShareKeys,omitempty" xml:"HostShareKeys,omitempty" type:"Repeated"`
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount    *int64                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHostShareKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHostShareKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListHostShareKeysResponseBody) SetHostShareKeys(v []*ListHostShareKeysResponseBodyHostShareKeys) *ListHostShareKeysResponseBody {
	s.HostShareKeys = v
	return s
}

func (s *ListHostShareKeysResponseBody) SetRequestId(v string) *ListHostShareKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHostShareKeysResponseBody) SetTotalCount(v int64) *ListHostShareKeysResponseBody {
	s.TotalCount = &v
	return s
}

type ListHostShareKeysResponseBodyHostShareKeys struct {
	HostAccountCount      *int64  `json:"HostAccountCount,omitempty" xml:"HostAccountCount,omitempty"`
	HostShareKeyId        *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	HostShareKeyName      *string `json:"HostShareKeyName,omitempty" xml:"HostShareKeyName,omitempty"`
	LastModifyKeyAt       *int64  `json:"LastModifyKeyAt,omitempty" xml:"LastModifyKeyAt,omitempty"`
	PrivateKeyFingerPrint *string `json:"PrivateKeyFingerPrint,omitempty" xml:"PrivateKeyFingerPrint,omitempty"`
}

func (s ListHostShareKeysResponseBodyHostShareKeys) String() string {
	return tea.Prettify(s)
}

func (s ListHostShareKeysResponseBodyHostShareKeys) GoString() string {
	return s.String()
}

func (s *ListHostShareKeysResponseBodyHostShareKeys) SetHostAccountCount(v int64) *ListHostShareKeysResponseBodyHostShareKeys {
	s.HostAccountCount = &v
	return s
}

func (s *ListHostShareKeysResponseBodyHostShareKeys) SetHostShareKeyId(v string) *ListHostShareKeysResponseBodyHostShareKeys {
	s.HostShareKeyId = &v
	return s
}

func (s *ListHostShareKeysResponseBodyHostShareKeys) SetHostShareKeyName(v string) *ListHostShareKeysResponseBodyHostShareKeys {
	s.HostShareKeyName = &v
	return s
}

func (s *ListHostShareKeysResponseBodyHostShareKeys) SetLastModifyKeyAt(v int64) *ListHostShareKeysResponseBodyHostShareKeys {
	s.LastModifyKeyAt = &v
	return s
}

func (s *ListHostShareKeysResponseBodyHostShareKeys) SetPrivateKeyFingerPrint(v string) *ListHostShareKeysResponseBodyHostShareKeys {
	s.PrivateKeyFingerPrint = &v
	return s
}

type ListHostShareKeysResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListHostShareKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListHostShareKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHostShareKeysResponse) GoString() string {
	return s.String()
}

func (s *ListHostShareKeysResponse) SetHeaders(v map[string]*string) *ListHostShareKeysResponse {
	s.Headers = v
	return s
}

func (s *ListHostShareKeysResponse) SetStatusCode(v int32) *ListHostShareKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHostShareKeysResponse) SetBody(v *ListHostShareKeysResponseBody) *ListHostShareKeysResponse {
	s.Body = v
	return s
}

type ListHostsRequest struct {
	HostAddress         *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	HostGroupId         *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	HostName            *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OSType              *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	PageNumber          *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize            *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Source              *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceInstanceId    *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
	SourceInstanceState *string `json:"SourceInstanceState,omitempty" xml:"SourceInstanceState,omitempty"`
}

func (s ListHostsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHostsRequest) GoString() string {
	return s.String()
}

func (s *ListHostsRequest) SetHostAddress(v string) *ListHostsRequest {
	s.HostAddress = &v
	return s
}

func (s *ListHostsRequest) SetHostGroupId(v string) *ListHostsRequest {
	s.HostGroupId = &v
	return s
}

func (s *ListHostsRequest) SetHostName(v string) *ListHostsRequest {
	s.HostName = &v
	return s
}

func (s *ListHostsRequest) SetInstanceId(v string) *ListHostsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHostsRequest) SetOSType(v string) *ListHostsRequest {
	s.OSType = &v
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

func (s *ListHostsRequest) SetRegionId(v string) *ListHostsRequest {
	s.RegionId = &v
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

type ListHostsResponseBody struct {
	Hosts      []*ListHostsResponseBodyHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
	RequestId  *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *ListHostsResponseBody) SetRequestId(v string) *ListHostsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHostsResponseBody) SetTotalCount(v int32) *ListHostsResponseBody {
	s.TotalCount = &v
	return s
}

type ListHostsResponseBodyHosts struct {
	ActiveAddressType   *string `json:"ActiveAddressType,omitempty" xml:"ActiveAddressType,omitempty"`
	Comment             *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	HostAccountCount    *int32  `json:"HostAccountCount,omitempty" xml:"HostAccountCount,omitempty"`
	HostId              *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	HostName            *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	HostPrivateAddress  *string `json:"HostPrivateAddress,omitempty" xml:"HostPrivateAddress,omitempty"`
	HostPublicAddress   *string `json:"HostPublicAddress,omitempty" xml:"HostPublicAddress,omitempty"`
	OSType              *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	Source              *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceInstanceId    *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
	SourceInstanceState *string `json:"SourceInstanceState,omitempty" xml:"SourceInstanceState,omitempty"`
}

func (s ListHostsResponseBodyHosts) String() string {
	return tea.Prettify(s)
}

func (s ListHostsResponseBodyHosts) GoString() string {
	return s.String()
}

func (s *ListHostsResponseBodyHosts) SetActiveAddressType(v string) *ListHostsResponseBodyHosts {
	s.ActiveAddressType = &v
	return s
}

func (s *ListHostsResponseBodyHosts) SetComment(v string) *ListHostsResponseBodyHosts {
	s.Comment = &v
	return s
}

func (s *ListHostsResponseBodyHosts) SetHostAccountCount(v int32) *ListHostsResponseBodyHosts {
	s.HostAccountCount = &v
	return s
}

func (s *ListHostsResponseBodyHosts) SetHostId(v string) *ListHostsResponseBodyHosts {
	s.HostId = &v
	return s
}

func (s *ListHostsResponseBodyHosts) SetHostName(v string) *ListHostsResponseBodyHosts {
	s.HostName = &v
	return s
}

func (s *ListHostsResponseBodyHosts) SetHostPrivateAddress(v string) *ListHostsResponseBodyHosts {
	s.HostPrivateAddress = &v
	return s
}

func (s *ListHostsResponseBodyHosts) SetHostPublicAddress(v string) *ListHostsResponseBodyHosts {
	s.HostPublicAddress = &v
	return s
}

func (s *ListHostsResponseBodyHosts) SetOSType(v string) *ListHostsResponseBodyHosts {
	s.OSType = &v
	return s
}

func (s *ListHostsResponseBodyHosts) SetSource(v string) *ListHostsResponseBodyHosts {
	s.Source = &v
	return s
}

func (s *ListHostsResponseBodyHosts) SetSourceInstanceId(v string) *ListHostsResponseBodyHosts {
	s.SourceInstanceId = &v
	return s
}

func (s *ListHostsResponseBodyHosts) SetSourceInstanceState(v string) *ListHostsResponseBodyHosts {
	s.SourceInstanceState = &v
	return s
}

type ListHostsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListHostsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListHostsResponse) SetStatusCode(v int32) *ListHostsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHostsResponse) SetBody(v *ListHostsResponseBody) *ListHostsResponse {
	s.Body = v
	return s
}

type ListHostsForUserRequest struct {
	HostAddress *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	HostName    *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Mode        *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	OSType      *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	PageNumber  *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListHostsForUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHostsForUserRequest) GoString() string {
	return s.String()
}

func (s *ListHostsForUserRequest) SetHostAddress(v string) *ListHostsForUserRequest {
	s.HostAddress = &v
	return s
}

func (s *ListHostsForUserRequest) SetHostName(v string) *ListHostsForUserRequest {
	s.HostName = &v
	return s
}

func (s *ListHostsForUserRequest) SetInstanceId(v string) *ListHostsForUserRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHostsForUserRequest) SetMode(v string) *ListHostsForUserRequest {
	s.Mode = &v
	return s
}

func (s *ListHostsForUserRequest) SetOSType(v string) *ListHostsForUserRequest {
	s.OSType = &v
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

func (s *ListHostsForUserRequest) SetRegionId(v string) *ListHostsForUserRequest {
	s.RegionId = &v
	return s
}

func (s *ListHostsForUserRequest) SetUserId(v string) *ListHostsForUserRequest {
	s.UserId = &v
	return s
}

type ListHostsForUserResponseBody struct {
	Hosts      []*ListHostsForUserResponseBodyHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *ListHostsForUserResponseBody) SetRequestId(v string) *ListHostsForUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHostsForUserResponseBody) SetTotalCount(v int32) *ListHostsForUserResponseBody {
	s.TotalCount = &v
	return s
}

type ListHostsForUserResponseBodyHosts struct {
	ActiveAddressType  *string `json:"ActiveAddressType,omitempty" xml:"ActiveAddressType,omitempty"`
	Comment            *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	HostId             *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	HostName           *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	HostPrivateAddress *string `json:"HostPrivateAddress,omitempty" xml:"HostPrivateAddress,omitempty"`
	HostPublicAddress  *string `json:"HostPublicAddress,omitempty" xml:"HostPublicAddress,omitempty"`
	OSType             *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
}

func (s ListHostsForUserResponseBodyHosts) String() string {
	return tea.Prettify(s)
}

func (s ListHostsForUserResponseBodyHosts) GoString() string {
	return s.String()
}

func (s *ListHostsForUserResponseBodyHosts) SetActiveAddressType(v string) *ListHostsForUserResponseBodyHosts {
	s.ActiveAddressType = &v
	return s
}

func (s *ListHostsForUserResponseBodyHosts) SetComment(v string) *ListHostsForUserResponseBodyHosts {
	s.Comment = &v
	return s
}

func (s *ListHostsForUserResponseBodyHosts) SetHostId(v string) *ListHostsForUserResponseBodyHosts {
	s.HostId = &v
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

func (s *ListHostsForUserResponseBodyHosts) SetHostPublicAddress(v string) *ListHostsForUserResponseBodyHosts {
	s.HostPublicAddress = &v
	return s
}

func (s *ListHostsForUserResponseBodyHosts) SetOSType(v string) *ListHostsForUserResponseBodyHosts {
	s.OSType = &v
	return s
}

type ListHostsForUserResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListHostsForUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListHostsForUserResponse) SetStatusCode(v int32) *ListHostsForUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHostsForUserResponse) SetBody(v *ListHostsForUserResponseBody) *ListHostsForUserResponse {
	s.Body = v
	return s
}

type ListHostsForUserGroupRequest struct {
	HostAddress *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	HostName    *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Mode        *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	OSType      *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	PageNumber  *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s ListHostsForUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHostsForUserGroupRequest) GoString() string {
	return s.String()
}

func (s *ListHostsForUserGroupRequest) SetHostAddress(v string) *ListHostsForUserGroupRequest {
	s.HostAddress = &v
	return s
}

func (s *ListHostsForUserGroupRequest) SetHostName(v string) *ListHostsForUserGroupRequest {
	s.HostName = &v
	return s
}

func (s *ListHostsForUserGroupRequest) SetInstanceId(v string) *ListHostsForUserGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHostsForUserGroupRequest) SetMode(v string) *ListHostsForUserGroupRequest {
	s.Mode = &v
	return s
}

func (s *ListHostsForUserGroupRequest) SetOSType(v string) *ListHostsForUserGroupRequest {
	s.OSType = &v
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

func (s *ListHostsForUserGroupRequest) SetRegionId(v string) *ListHostsForUserGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ListHostsForUserGroupRequest) SetUserGroupId(v string) *ListHostsForUserGroupRequest {
	s.UserGroupId = &v
	return s
}

type ListHostsForUserGroupResponseBody struct {
	Hosts      []*ListHostsForUserGroupResponseBodyHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *ListHostsForUserGroupResponseBody) SetRequestId(v string) *ListHostsForUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHostsForUserGroupResponseBody) SetTotalCount(v int32) *ListHostsForUserGroupResponseBody {
	s.TotalCount = &v
	return s
}

type ListHostsForUserGroupResponseBodyHosts struct {
	ActiveAddressType  *string `json:"ActiveAddressType,omitempty" xml:"ActiveAddressType,omitempty"`
	Comment            *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	HostId             *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	HostName           *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	HostPrivateAddress *string `json:"HostPrivateAddress,omitempty" xml:"HostPrivateAddress,omitempty"`
	HostPublicAddress  *string `json:"HostPublicAddress,omitempty" xml:"HostPublicAddress,omitempty"`
	OSType             *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
}

func (s ListHostsForUserGroupResponseBodyHosts) String() string {
	return tea.Prettify(s)
}

func (s ListHostsForUserGroupResponseBodyHosts) GoString() string {
	return s.String()
}

func (s *ListHostsForUserGroupResponseBodyHosts) SetActiveAddressType(v string) *ListHostsForUserGroupResponseBodyHosts {
	s.ActiveAddressType = &v
	return s
}

func (s *ListHostsForUserGroupResponseBodyHosts) SetComment(v string) *ListHostsForUserGroupResponseBodyHosts {
	s.Comment = &v
	return s
}

func (s *ListHostsForUserGroupResponseBodyHosts) SetHostId(v string) *ListHostsForUserGroupResponseBodyHosts {
	s.HostId = &v
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

func (s *ListHostsForUserGroupResponseBodyHosts) SetHostPublicAddress(v string) *ListHostsForUserGroupResponseBodyHosts {
	s.HostPublicAddress = &v
	return s
}

func (s *ListHostsForUserGroupResponseBodyHosts) SetOSType(v string) *ListHostsForUserGroupResponseBodyHosts {
	s.OSType = &v
	return s
}

type ListHostsForUserGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListHostsForUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListHostsForUserGroupResponse) SetStatusCode(v int32) *ListHostsForUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHostsForUserGroupResponse) SetBody(v *ListHostsForUserGroupResponseBody) *ListHostsForUserGroupResponse {
	s.Body = v
	return s
}

type ListTagKeysRequest struct {
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) SetPageNumber(v int32) *ListTagKeysRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTagKeysRequest) SetPageSize(v int32) *ListTagKeysRequest {
	s.PageSize = &v
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

type ListTagKeysResponseBody struct {
	PageNumber *int32                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagKeys    []*ListTagKeysResponseBodyTagKeys `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
	TotalCount *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTagKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBody) SetPageNumber(v int32) *ListTagKeysResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTagKeysResponseBody) SetPageSize(v int32) *ListTagKeysResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTagKeysResponseBody) SetRequestId(v string) *ListTagKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagKeysResponseBody) SetTagKeys(v []*ListTagKeysResponseBodyTagKeys) *ListTagKeysResponseBody {
	s.TagKeys = v
	return s
}

func (s *ListTagKeysResponseBody) SetTotalCount(v int32) *ListTagKeysResponseBody {
	s.TotalCount = &v
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListTagKeysResponse) SetStatusCode(v int32) *ListTagKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagKeysResponse) SetBody(v *ListTagKeysResponseBody) *ListTagKeysResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	NextToken    *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId     *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId   []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
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
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ListUserGroupsRequest struct {
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber    *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
}

func (s ListUserGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListUserGroupsRequest) SetInstanceId(v string) *ListUserGroupsRequest {
	s.InstanceId = &v
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

func (s *ListUserGroupsRequest) SetRegionId(v string) *ListUserGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *ListUserGroupsRequest) SetUserGroupName(v string) *ListUserGroupsRequest {
	s.UserGroupName = &v
	return s
}

type ListUserGroupsResponseBody struct {
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	UserGroups []*ListUserGroupsResponseBodyUserGroups `json:"UserGroups,omitempty" xml:"UserGroups,omitempty" type:"Repeated"`
}

func (s ListUserGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserGroupsResponseBody) SetRequestId(v string) *ListUserGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserGroupsResponseBody) SetTotalCount(v int32) *ListUserGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUserGroupsResponseBody) SetUserGroups(v []*ListUserGroupsResponseBodyUserGroups) *ListUserGroupsResponseBody {
	s.UserGroups = v
	return s
}

type ListUserGroupsResponseBodyUserGroups struct {
	Comment       *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	MemberCount   *int32  `json:"MemberCount,omitempty" xml:"MemberCount,omitempty"`
	UserGroupId   *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
}

func (s ListUserGroupsResponseBodyUserGroups) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsResponseBodyUserGroups) GoString() string {
	return s.String()
}

func (s *ListUserGroupsResponseBodyUserGroups) SetComment(v string) *ListUserGroupsResponseBodyUserGroups {
	s.Comment = &v
	return s
}

func (s *ListUserGroupsResponseBodyUserGroups) SetMemberCount(v int32) *ListUserGroupsResponseBodyUserGroups {
	s.MemberCount = &v
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListUserGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListUserGroupsResponse) SetStatusCode(v int32) *ListUserGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserGroupsResponse) SetBody(v *ListUserGroupsResponseBody) *ListUserGroupsResponse {
	s.Body = v
	return s
}

type ListUsersRequest struct {
	DisplayName  *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Mobile       *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	PageNumber   *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Source       *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceUserId *string `json:"SourceUserId,omitempty" xml:"SourceUserId,omitempty"`
	UserGroupId  *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserState    *string `json:"UserState,omitempty" xml:"UserState,omitempty"`
}

func (s ListUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) SetDisplayName(v string) *ListUsersRequest {
	s.DisplayName = &v
	return s
}

func (s *ListUsersRequest) SetInstanceId(v string) *ListUsersRequest {
	s.InstanceId = &v
	return s
}

func (s *ListUsersRequest) SetMobile(v string) *ListUsersRequest {
	s.Mobile = &v
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

func (s *ListUsersRequest) SetRegionId(v string) *ListUsersRequest {
	s.RegionId = &v
	return s
}

func (s *ListUsersRequest) SetSource(v string) *ListUsersRequest {
	s.Source = &v
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

func (s *ListUsersRequest) SetUserName(v string) *ListUsersRequest {
	s.UserName = &v
	return s
}

func (s *ListUsersRequest) SetUserState(v string) *ListUsersRequest {
	s.UserState = &v
	return s
}

type ListUsersResponseBody struct {
	RequestId  *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Users      []*ListUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
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

func (s *ListUsersResponseBody) SetTotalCount(v int32) *ListUsersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUsersResponseBody) SetUsers(v []*ListUsersResponseBodyUsers) *ListUsersResponseBody {
	s.Users = v
	return s
}

type ListUsersResponseBodyUsers struct {
	Comment           *string   `json:"Comment,omitempty" xml:"Comment,omitempty"`
	DisplayName       *string   `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Email             *string   `json:"Email,omitempty" xml:"Email,omitempty"`
	Mobile            *string   `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	MobileCountryCode *string   `json:"MobileCountryCode,omitempty" xml:"MobileCountryCode,omitempty"`
	Source            *string   `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceUserId      *string   `json:"SourceUserId,omitempty" xml:"SourceUserId,omitempty"`
	UserId            *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName          *string   `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserState         []*string `json:"UserState,omitempty" xml:"UserState,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUsers) SetComment(v string) *ListUsersResponseBodyUsers {
	s.Comment = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetDisplayName(v string) *ListUsersResponseBodyUsers {
	s.DisplayName = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetEmail(v string) *ListUsersResponseBodyUsers {
	s.Email = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetMobile(v string) *ListUsersResponseBodyUsers {
	s.Mobile = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetMobileCountryCode(v string) *ListUsersResponseBodyUsers {
	s.MobileCountryCode = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetSource(v string) *ListUsersResponseBodyUsers {
	s.Source = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetSourceUserId(v string) *ListUsersResponseBodyUsers {
	s.SourceUserId = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetUserId(v string) *ListUsersResponseBodyUsers {
	s.UserId = &v
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

type ListUsersResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListUsersResponse) SetStatusCode(v int32) *ListUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUsersResponse) SetBody(v *ListUsersResponseBody) *ListUsersResponse {
	s.Body = v
	return s
}

type LockUsersRequest struct {
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
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	UserId  *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s LockUsersResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s LockUsersResponseBodyResults) GoString() string {
	return s.String()
}

func (s *LockUsersResponseBodyResults) SetCode(v string) *LockUsersResponseBodyResults {
	s.Code = &v
	return s
}

func (s *LockUsersResponseBodyResults) SetMessage(v string) *LockUsersResponseBodyResults {
	s.Message = &v
	return s
}

func (s *LockUsersResponseBodyResults) SetUserId(v string) *LockUsersResponseBodyResults {
	s.UserId = &v
	return s
}

type LockUsersResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *LockUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *LockUsersResponse) SetStatusCode(v int32) *LockUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *LockUsersResponse) SetBody(v *LockUsersResponseBody) *LockUsersResponse {
	s.Body = v
	return s
}

type ModifyHostRequest struct {
	Comment            *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	HostId             *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	HostName           *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	HostPrivateAddress *string `json:"HostPrivateAddress,omitempty" xml:"HostPrivateAddress,omitempty"`
	HostPublicAddress  *string `json:"HostPublicAddress,omitempty" xml:"HostPublicAddress,omitempty"`
	InstanceId         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OSType             *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyHostRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyHostRequest) GoString() string {
	return s.String()
}

func (s *ModifyHostRequest) SetComment(v string) *ModifyHostRequest {
	s.Comment = &v
	return s
}

func (s *ModifyHostRequest) SetHostId(v string) *ModifyHostRequest {
	s.HostId = &v
	return s
}

func (s *ModifyHostRequest) SetHostName(v string) *ModifyHostRequest {
	s.HostName = &v
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

func (s *ModifyHostRequest) SetInstanceId(v string) *ModifyHostRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyHostRequest) SetOSType(v string) *ModifyHostRequest {
	s.OSType = &v
	return s
}

func (s *ModifyHostRequest) SetRegionId(v string) *ModifyHostRequest {
	s.RegionId = &v
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyHostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyHostResponse) SetStatusCode(v int32) *ModifyHostResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHostResponse) SetBody(v *ModifyHostResponseBody) *ModifyHostResponse {
	s.Body = v
	return s
}

type ModifyHostAccountRequest struct {
	HostAccountId   *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	HostShareKeyId  *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PassPhrase      *string `json:"PassPhrase,omitempty" xml:"PassPhrase,omitempty"`
	Password        *string `json:"Password,omitempty" xml:"Password,omitempty"`
	PrivateKey      *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyHostAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyHostAccountRequest) GoString() string {
	return s.String()
}

func (s *ModifyHostAccountRequest) SetHostAccountId(v string) *ModifyHostAccountRequest {
	s.HostAccountId = &v
	return s
}

func (s *ModifyHostAccountRequest) SetHostAccountName(v string) *ModifyHostAccountRequest {
	s.HostAccountName = &v
	return s
}

func (s *ModifyHostAccountRequest) SetHostShareKeyId(v string) *ModifyHostAccountRequest {
	s.HostShareKeyId = &v
	return s
}

func (s *ModifyHostAccountRequest) SetInstanceId(v string) *ModifyHostAccountRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyHostAccountRequest) SetPassPhrase(v string) *ModifyHostAccountRequest {
	s.PassPhrase = &v
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

func (s *ModifyHostAccountRequest) SetRegionId(v string) *ModifyHostAccountRequest {
	s.RegionId = &v
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyHostAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyHostAccountResponse) SetStatusCode(v int32) *ModifyHostAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHostAccountResponse) SetBody(v *ModifyHostAccountResponseBody) *ModifyHostAccountResponse {
	s.Body = v
	return s
}

type ModifyHostGroupRequest struct {
	Comment       *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	HostGroupId   *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	HostGroupName *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyHostGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyHostGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyHostGroupRequest) SetComment(v string) *ModifyHostGroupRequest {
	s.Comment = &v
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

func (s *ModifyHostGroupRequest) SetInstanceId(v string) *ModifyHostGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyHostGroupRequest) SetRegionId(v string) *ModifyHostGroupRequest {
	s.RegionId = &v
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyHostGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyHostGroupResponse) SetStatusCode(v int32) *ModifyHostGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHostGroupResponse) SetBody(v *ModifyHostGroupResponseBody) *ModifyHostGroupResponse {
	s.Body = v
	return s
}

type ModifyHostShareKeyRequest struct {
	HostShareKeyId   *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	HostShareKeyName *string `json:"HostShareKeyName,omitempty" xml:"HostShareKeyName,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PassPhrase       *string `json:"PassPhrase,omitempty" xml:"PassPhrase,omitempty"`
	PrivateKey       *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyHostShareKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyHostShareKeyRequest) GoString() string {
	return s.String()
}

func (s *ModifyHostShareKeyRequest) SetHostShareKeyId(v string) *ModifyHostShareKeyRequest {
	s.HostShareKeyId = &v
	return s
}

func (s *ModifyHostShareKeyRequest) SetHostShareKeyName(v string) *ModifyHostShareKeyRequest {
	s.HostShareKeyName = &v
	return s
}

func (s *ModifyHostShareKeyRequest) SetInstanceId(v string) *ModifyHostShareKeyRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyHostShareKeyRequest) SetPassPhrase(v string) *ModifyHostShareKeyRequest {
	s.PassPhrase = &v
	return s
}

func (s *ModifyHostShareKeyRequest) SetPrivateKey(v string) *ModifyHostShareKeyRequest {
	s.PrivateKey = &v
	return s
}

func (s *ModifyHostShareKeyRequest) SetRegionId(v string) *ModifyHostShareKeyRequest {
	s.RegionId = &v
	return s
}

type ModifyHostShareKeyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyHostShareKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyHostShareKeyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHostShareKeyResponseBody) SetRequestId(v string) *ModifyHostShareKeyResponseBody {
	s.RequestId = &v
	return s
}

type ModifyHostShareKeyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyHostShareKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyHostShareKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyHostShareKeyResponse) GoString() string {
	return s.String()
}

func (s *ModifyHostShareKeyResponse) SetHeaders(v map[string]*string) *ModifyHostShareKeyResponse {
	s.Headers = v
	return s
}

func (s *ModifyHostShareKeyResponse) SetStatusCode(v int32) *ModifyHostShareKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHostShareKeyResponse) SetBody(v *ModifyHostShareKeyResponseBody) *ModifyHostShareKeyResponse {
	s.Body = v
	return s
}

type ModifyHostsActiveAddressTypeRequest struct {
	ActiveAddressType *string `json:"ActiveAddressType,omitempty" xml:"ActiveAddressType,omitempty"`
	HostIds           *string `json:"HostIds,omitempty" xml:"HostIds,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyHostsActiveAddressTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyHostsActiveAddressTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyHostsActiveAddressTypeRequest) SetActiveAddressType(v string) *ModifyHostsActiveAddressTypeRequest {
	s.ActiveAddressType = &v
	return s
}

func (s *ModifyHostsActiveAddressTypeRequest) SetHostIds(v string) *ModifyHostsActiveAddressTypeRequest {
	s.HostIds = &v
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
	HostId  *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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

func (s *ModifyHostsActiveAddressTypeResponseBodyResults) SetHostId(v string) *ModifyHostsActiveAddressTypeResponseBodyResults {
	s.HostId = &v
	return s
}

func (s *ModifyHostsActiveAddressTypeResponseBodyResults) SetMessage(v string) *ModifyHostsActiveAddressTypeResponseBodyResults {
	s.Message = &v
	return s
}

type ModifyHostsActiveAddressTypeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyHostsActiveAddressTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyHostsActiveAddressTypeResponse) SetStatusCode(v int32) *ModifyHostsActiveAddressTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHostsActiveAddressTypeResponse) SetBody(v *ModifyHostsActiveAddressTypeResponseBody) *ModifyHostsActiveAddressTypeResponse {
	s.Body = v
	return s
}

type ModifyHostsPortRequest struct {
	HostIds      *string `json:"HostIds,omitempty" xml:"HostIds,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Port         *string `json:"Port,omitempty" xml:"Port,omitempty"`
	ProtocolName *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyHostsPortRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyHostsPortRequest) GoString() string {
	return s.String()
}

func (s *ModifyHostsPortRequest) SetHostIds(v string) *ModifyHostsPortRequest {
	s.HostIds = &v
	return s
}

func (s *ModifyHostsPortRequest) SetInstanceId(v string) *ModifyHostsPortRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyHostsPortRequest) SetPort(v string) *ModifyHostsPortRequest {
	s.Port = &v
	return s
}

func (s *ModifyHostsPortRequest) SetProtocolName(v string) *ModifyHostsPortRequest {
	s.ProtocolName = &v
	return s
}

func (s *ModifyHostsPortRequest) SetRegionId(v string) *ModifyHostsPortRequest {
	s.RegionId = &v
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
	HostId  *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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

func (s *ModifyHostsPortResponseBodyResults) SetHostId(v string) *ModifyHostsPortResponseBodyResults {
	s.HostId = &v
	return s
}

func (s *ModifyHostsPortResponseBodyResults) SetMessage(v string) *ModifyHostsPortResponseBodyResults {
	s.Message = &v
	return s
}

type ModifyHostsPortResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyHostsPortResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyHostsPortResponse) SetStatusCode(v int32) *ModifyHostsPortResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHostsPortResponse) SetBody(v *ModifyHostsPortResponseBody) *ModifyHostsPortResponse {
	s.Body = v
	return s
}

type ModifyInstanceADAuthServerRequest struct {
	Account       *string `json:"Account,omitempty" xml:"Account,omitempty"`
	BaseDN        *string `json:"BaseDN,omitempty" xml:"BaseDN,omitempty"`
	Domain        *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	EmailMapping  *string `json:"EmailMapping,omitempty" xml:"EmailMapping,omitempty"`
	Filter        *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IsSSL         *string `json:"IsSSL,omitempty" xml:"IsSSL,omitempty"`
	MobileMapping *string `json:"MobileMapping,omitempty" xml:"MobileMapping,omitempty"`
	NameMapping   *string `json:"NameMapping,omitempty" xml:"NameMapping,omitempty"`
	Password      *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Port          *string `json:"Port,omitempty" xml:"Port,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Server        *string `json:"Server,omitempty" xml:"Server,omitempty"`
	StandbyServer *string `json:"StandbyServer,omitempty" xml:"StandbyServer,omitempty"`
}

func (s ModifyInstanceADAuthServerRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceADAuthServerRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceADAuthServerRequest) SetAccount(v string) *ModifyInstanceADAuthServerRequest {
	s.Account = &v
	return s
}

func (s *ModifyInstanceADAuthServerRequest) SetBaseDN(v string) *ModifyInstanceADAuthServerRequest {
	s.BaseDN = &v
	return s
}

func (s *ModifyInstanceADAuthServerRequest) SetDomain(v string) *ModifyInstanceADAuthServerRequest {
	s.Domain = &v
	return s
}

func (s *ModifyInstanceADAuthServerRequest) SetEmailMapping(v string) *ModifyInstanceADAuthServerRequest {
	s.EmailMapping = &v
	return s
}

func (s *ModifyInstanceADAuthServerRequest) SetFilter(v string) *ModifyInstanceADAuthServerRequest {
	s.Filter = &v
	return s
}

func (s *ModifyInstanceADAuthServerRequest) SetInstanceId(v string) *ModifyInstanceADAuthServerRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceADAuthServerRequest) SetIsSSL(v string) *ModifyInstanceADAuthServerRequest {
	s.IsSSL = &v
	return s
}

func (s *ModifyInstanceADAuthServerRequest) SetMobileMapping(v string) *ModifyInstanceADAuthServerRequest {
	s.MobileMapping = &v
	return s
}

func (s *ModifyInstanceADAuthServerRequest) SetNameMapping(v string) *ModifyInstanceADAuthServerRequest {
	s.NameMapping = &v
	return s
}

func (s *ModifyInstanceADAuthServerRequest) SetPassword(v string) *ModifyInstanceADAuthServerRequest {
	s.Password = &v
	return s
}

func (s *ModifyInstanceADAuthServerRequest) SetPort(v string) *ModifyInstanceADAuthServerRequest {
	s.Port = &v
	return s
}

func (s *ModifyInstanceADAuthServerRequest) SetRegionId(v string) *ModifyInstanceADAuthServerRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceADAuthServerRequest) SetServer(v string) *ModifyInstanceADAuthServerRequest {
	s.Server = &v
	return s
}

func (s *ModifyInstanceADAuthServerRequest) SetStandbyServer(v string) *ModifyInstanceADAuthServerRequest {
	s.StandbyServer = &v
	return s
}

type ModifyInstanceADAuthServerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceADAuthServerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceADAuthServerResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceADAuthServerResponseBody) SetRequestId(v string) *ModifyInstanceADAuthServerResponseBody {
	s.RequestId = &v
	return s
}

type ModifyInstanceADAuthServerResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyInstanceADAuthServerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyInstanceADAuthServerResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceADAuthServerResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceADAuthServerResponse) SetHeaders(v map[string]*string) *ModifyInstanceADAuthServerResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceADAuthServerResponse) SetStatusCode(v int32) *ModifyInstanceADAuthServerResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceADAuthServerResponse) SetBody(v *ModifyInstanceADAuthServerResponseBody) *ModifyInstanceADAuthServerResponse {
	s.Body = v
	return s
}

type ModifyInstanceAttributeRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyInstanceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttributeRequest) SetDescription(v string) *ModifyInstanceAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetInstanceId(v string) *ModifyInstanceAttributeRequest {
	s.InstanceId = &v
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyInstanceAttributeResponse) SetStatusCode(v int32) *ModifyInstanceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceAttributeResponse) SetBody(v *ModifyInstanceAttributeResponseBody) *ModifyInstanceAttributeResponse {
	s.Body = v
	return s
}

type ModifyInstanceLDAPAuthServerRequest struct {
	Account          *string `json:"Account,omitempty" xml:"Account,omitempty"`
	BaseDN           *string `json:"BaseDN,omitempty" xml:"BaseDN,omitempty"`
	EmailMapping     *string `json:"EmailMapping,omitempty" xml:"EmailMapping,omitempty"`
	Filter           *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IsSSL            *string `json:"IsSSL,omitempty" xml:"IsSSL,omitempty"`
	LoginNameMapping *string `json:"LoginNameMapping,omitempty" xml:"LoginNameMapping,omitempty"`
	MobileMapping    *string `json:"MobileMapping,omitempty" xml:"MobileMapping,omitempty"`
	NameMapping      *string `json:"NameMapping,omitempty" xml:"NameMapping,omitempty"`
	Password         *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Port             *string `json:"Port,omitempty" xml:"Port,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Server           *string `json:"Server,omitempty" xml:"Server,omitempty"`
	StandbyServer    *string `json:"StandbyServer,omitempty" xml:"StandbyServer,omitempty"`
}

func (s ModifyInstanceLDAPAuthServerRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceLDAPAuthServerRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceLDAPAuthServerRequest) SetAccount(v string) *ModifyInstanceLDAPAuthServerRequest {
	s.Account = &v
	return s
}

func (s *ModifyInstanceLDAPAuthServerRequest) SetBaseDN(v string) *ModifyInstanceLDAPAuthServerRequest {
	s.BaseDN = &v
	return s
}

func (s *ModifyInstanceLDAPAuthServerRequest) SetEmailMapping(v string) *ModifyInstanceLDAPAuthServerRequest {
	s.EmailMapping = &v
	return s
}

func (s *ModifyInstanceLDAPAuthServerRequest) SetFilter(v string) *ModifyInstanceLDAPAuthServerRequest {
	s.Filter = &v
	return s
}

func (s *ModifyInstanceLDAPAuthServerRequest) SetInstanceId(v string) *ModifyInstanceLDAPAuthServerRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceLDAPAuthServerRequest) SetIsSSL(v string) *ModifyInstanceLDAPAuthServerRequest {
	s.IsSSL = &v
	return s
}

func (s *ModifyInstanceLDAPAuthServerRequest) SetLoginNameMapping(v string) *ModifyInstanceLDAPAuthServerRequest {
	s.LoginNameMapping = &v
	return s
}

func (s *ModifyInstanceLDAPAuthServerRequest) SetMobileMapping(v string) *ModifyInstanceLDAPAuthServerRequest {
	s.MobileMapping = &v
	return s
}

func (s *ModifyInstanceLDAPAuthServerRequest) SetNameMapping(v string) *ModifyInstanceLDAPAuthServerRequest {
	s.NameMapping = &v
	return s
}

func (s *ModifyInstanceLDAPAuthServerRequest) SetPassword(v string) *ModifyInstanceLDAPAuthServerRequest {
	s.Password = &v
	return s
}

func (s *ModifyInstanceLDAPAuthServerRequest) SetPort(v string) *ModifyInstanceLDAPAuthServerRequest {
	s.Port = &v
	return s
}

func (s *ModifyInstanceLDAPAuthServerRequest) SetRegionId(v string) *ModifyInstanceLDAPAuthServerRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceLDAPAuthServerRequest) SetServer(v string) *ModifyInstanceLDAPAuthServerRequest {
	s.Server = &v
	return s
}

func (s *ModifyInstanceLDAPAuthServerRequest) SetStandbyServer(v string) *ModifyInstanceLDAPAuthServerRequest {
	s.StandbyServer = &v
	return s
}

type ModifyInstanceLDAPAuthServerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceLDAPAuthServerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceLDAPAuthServerResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceLDAPAuthServerResponseBody) SetRequestId(v string) *ModifyInstanceLDAPAuthServerResponseBody {
	s.RequestId = &v
	return s
}

type ModifyInstanceLDAPAuthServerResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyInstanceLDAPAuthServerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyInstanceLDAPAuthServerResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceLDAPAuthServerResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceLDAPAuthServerResponse) SetHeaders(v map[string]*string) *ModifyInstanceLDAPAuthServerResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceLDAPAuthServerResponse) SetStatusCode(v int32) *ModifyInstanceLDAPAuthServerResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceLDAPAuthServerResponse) SetBody(v *ModifyInstanceLDAPAuthServerResponseBody) *ModifyInstanceLDAPAuthServerResponse {
	s.Body = v
	return s
}

type ModifyInstanceTwoFactorRequest struct {
	DingTalkConfig    *string `json:"DingTalkConfig,omitempty" xml:"DingTalkConfig,omitempty"`
	EnableTwoFactor   *string `json:"EnableTwoFactor,omitempty" xml:"EnableTwoFactor,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MessageLanguage   *string `json:"MessageLanguage,omitempty" xml:"MessageLanguage,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SkipTwoFactorTime *string `json:"SkipTwoFactorTime,omitempty" xml:"SkipTwoFactorTime,omitempty"`
	TwoFactorMethods  *string `json:"TwoFactorMethods,omitempty" xml:"TwoFactorMethods,omitempty"`
}

func (s ModifyInstanceTwoFactorRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceTwoFactorRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceTwoFactorRequest) SetDingTalkConfig(v string) *ModifyInstanceTwoFactorRequest {
	s.DingTalkConfig = &v
	return s
}

func (s *ModifyInstanceTwoFactorRequest) SetEnableTwoFactor(v string) *ModifyInstanceTwoFactorRequest {
	s.EnableTwoFactor = &v
	return s
}

func (s *ModifyInstanceTwoFactorRequest) SetInstanceId(v string) *ModifyInstanceTwoFactorRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceTwoFactorRequest) SetMessageLanguage(v string) *ModifyInstanceTwoFactorRequest {
	s.MessageLanguage = &v
	return s
}

func (s *ModifyInstanceTwoFactorRequest) SetRegionId(v string) *ModifyInstanceTwoFactorRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceTwoFactorRequest) SetSkipTwoFactorTime(v string) *ModifyInstanceTwoFactorRequest {
	s.SkipTwoFactorTime = &v
	return s
}

func (s *ModifyInstanceTwoFactorRequest) SetTwoFactorMethods(v string) *ModifyInstanceTwoFactorRequest {
	s.TwoFactorMethods = &v
	return s
}

type ModifyInstanceTwoFactorResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceTwoFactorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceTwoFactorResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceTwoFactorResponseBody) SetRequestId(v string) *ModifyInstanceTwoFactorResponseBody {
	s.RequestId = &v
	return s
}

type ModifyInstanceTwoFactorResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyInstanceTwoFactorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyInstanceTwoFactorResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceTwoFactorResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceTwoFactorResponse) SetHeaders(v map[string]*string) *ModifyInstanceTwoFactorResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceTwoFactorResponse) SetStatusCode(v int32) *ModifyInstanceTwoFactorResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceTwoFactorResponse) SetBody(v *ModifyInstanceTwoFactorResponseBody) *ModifyInstanceTwoFactorResponse {
	s.Body = v
	return s
}

type ModifyUserRequest struct {
	Comment           *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	DisplayName       *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Email             *string `json:"Email,omitempty" xml:"Email,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Mobile            *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	MobileCountryCode *string `json:"MobileCountryCode,omitempty" xml:"MobileCountryCode,omitempty"`
	Password          *string `json:"Password,omitempty" xml:"Password,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserId            *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ModifyUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserRequest) SetComment(v string) *ModifyUserRequest {
	s.Comment = &v
	return s
}

func (s *ModifyUserRequest) SetDisplayName(v string) *ModifyUserRequest {
	s.DisplayName = &v
	return s
}

func (s *ModifyUserRequest) SetEmail(v string) *ModifyUserRequest {
	s.Email = &v
	return s
}

func (s *ModifyUserRequest) SetInstanceId(v string) *ModifyUserRequest {
	s.InstanceId = &v
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

func (s *ModifyUserRequest) SetPassword(v string) *ModifyUserRequest {
	s.Password = &v
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyUserResponse) SetStatusCode(v int32) *ModifyUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyUserResponse) SetBody(v *ModifyUserResponseBody) *ModifyUserResponse {
	s.Body = v
	return s
}

type ModifyUserGroupRequest struct {
	Comment       *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UserGroupId   *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
}

func (s ModifyUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserGroupRequest) SetComment(v string) *ModifyUserGroupRequest {
	s.Comment = &v
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyUserGroupResponse) SetStatusCode(v int32) *ModifyUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyUserGroupResponse) SetBody(v *ModifyUserGroupResponseBody) *ModifyUserGroupResponse {
	s.Body = v
	return s
}

type MoveResourceGroupRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceId      *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType    *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s MoveResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupRequest) SetRegionId(v string) *MoveResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceGroupId(v string) *MoveResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceId(v string) *MoveResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceType(v string) *MoveResourceGroupRequest {
	s.ResourceType = &v
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *MoveResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *MoveResourceGroupResponse) SetStatusCode(v int32) *MoveResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveResourceGroupResponse) SetBody(v *MoveResourceGroupResponseBody) *MoveResourceGroupResponse {
	s.Body = v
	return s
}

type RemoveHostsFromGroupRequest struct {
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	HostIds     *string `json:"HostIds,omitempty" xml:"HostIds,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RemoveHostsFromGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveHostsFromGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveHostsFromGroupRequest) SetHostGroupId(v string) *RemoveHostsFromGroupRequest {
	s.HostGroupId = &v
	return s
}

func (s *RemoveHostsFromGroupRequest) SetHostIds(v string) *RemoveHostsFromGroupRequest {
	s.HostIds = &v
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
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	HostId      *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty"`
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

func (s *RemoveHostsFromGroupResponseBodyResults) SetHostGroupId(v string) *RemoveHostsFromGroupResponseBodyResults {
	s.HostGroupId = &v
	return s
}

func (s *RemoveHostsFromGroupResponseBodyResults) SetHostId(v string) *RemoveHostsFromGroupResponseBodyResults {
	s.HostId = &v
	return s
}

func (s *RemoveHostsFromGroupResponseBodyResults) SetMessage(v string) *RemoveHostsFromGroupResponseBodyResults {
	s.Message = &v
	return s
}

type RemoveHostsFromGroupResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveHostsFromGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RemoveHostsFromGroupResponse) SetStatusCode(v int32) *RemoveHostsFromGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveHostsFromGroupResponse) SetBody(v *RemoveHostsFromGroupResponseBody) *RemoveHostsFromGroupResponse {
	s.Body = v
	return s
}

type RemoveUsersFromGroupRequest struct {
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
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty"`
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s RemoveUsersFromGroupResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s RemoveUsersFromGroupResponseBodyResults) GoString() string {
	return s.String()
}

func (s *RemoveUsersFromGroupResponseBodyResults) SetCode(v string) *RemoveUsersFromGroupResponseBodyResults {
	s.Code = &v
	return s
}

func (s *RemoveUsersFromGroupResponseBodyResults) SetMessage(v string) *RemoveUsersFromGroupResponseBodyResults {
	s.Message = &v
	return s
}

func (s *RemoveUsersFromGroupResponseBodyResults) SetUserGroupId(v string) *RemoveUsersFromGroupResponseBodyResults {
	s.UserGroupId = &v
	return s
}

func (s *RemoveUsersFromGroupResponseBodyResults) SetUserId(v string) *RemoveUsersFromGroupResponseBodyResults {
	s.UserId = &v
	return s
}

type RemoveUsersFromGroupResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveUsersFromGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RemoveUsersFromGroupResponse) SetStatusCode(v int32) *RemoveUsersFromGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveUsersFromGroupResponse) SetBody(v *RemoveUsersFromGroupResponseBody) *RemoveUsersFromGroupResponse {
	s.Body = v
	return s
}

type ResetHostAccountCredentialRequest struct {
	CredentialType *string `json:"CredentialType,omitempty" xml:"CredentialType,omitempty"`
	HostAccountId  *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ResetHostAccountCredentialRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetHostAccountCredentialRequest) GoString() string {
	return s.String()
}

func (s *ResetHostAccountCredentialRequest) SetCredentialType(v string) *ResetHostAccountCredentialRequest {
	s.CredentialType = &v
	return s
}

func (s *ResetHostAccountCredentialRequest) SetHostAccountId(v string) *ResetHostAccountCredentialRequest {
	s.HostAccountId = &v
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResetHostAccountCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ResetHostAccountCredentialResponse) SetStatusCode(v int32) *ResetHostAccountCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetHostAccountCredentialResponse) SetBody(v *ResetHostAccountCredentialResponseBody) *ResetHostAccountCredentialResponse {
	s.Body = v
	return s
}

type StartInstanceRequest struct {
	InstanceId       *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId         *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	VswitchId        *string   `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s StartInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartInstanceRequest) SetInstanceId(v string) *StartInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *StartInstanceRequest) SetRegionId(v string) *StartInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *StartInstanceRequest) SetSecurityGroupIds(v []*string) *StartInstanceRequest {
	s.SecurityGroupIds = v
	return s
}

func (s *StartInstanceRequest) SetVswitchId(v string) *StartInstanceRequest {
	s.VswitchId = &v
	return s
}

type StartInstanceResponseBody struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartInstanceResponseBody) SetInstanceId(v string) *StartInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *StartInstanceResponseBody) SetRequestId(v string) *StartInstanceResponseBody {
	s.RequestId = &v
	return s
}

type StartInstanceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *StartInstanceResponse) SetStatusCode(v int32) *StartInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartInstanceResponse) SetBody(v *StartInstanceResponseBody) *StartInstanceResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	RegionId     *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId   []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UnlockUsersRequest struct {
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
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	UserId  *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UnlockUsersResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s UnlockUsersResponseBodyResults) GoString() string {
	return s.String()
}

func (s *UnlockUsersResponseBodyResults) SetCode(v string) *UnlockUsersResponseBodyResults {
	s.Code = &v
	return s
}

func (s *UnlockUsersResponseBodyResults) SetMessage(v string) *UnlockUsersResponseBodyResults {
	s.Message = &v
	return s
}

func (s *UnlockUsersResponseBodyResults) SetUserId(v string) *UnlockUsersResponseBodyResults {
	s.UserId = &v
	return s
}

type UnlockUsersResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnlockUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UnlockUsersResponse) SetStatusCode(v int32) *UnlockUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *UnlockUsersResponse) SetBody(v *UnlockUsersResponseBody) *UnlockUsersResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	All          *bool     `json:"All,omitempty" xml:"All,omitempty"`
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId   []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type VerifyInstanceADAuthServerRequest struct {
	Account       *string `json:"Account,omitempty" xml:"Account,omitempty"`
	BaseDN        *string `json:"BaseDN,omitempty" xml:"BaseDN,omitempty"`
	Domain        *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Filter        *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IsSSL         *string `json:"IsSSL,omitempty" xml:"IsSSL,omitempty"`
	Password      *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Port          *string `json:"Port,omitempty" xml:"Port,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Server        *string `json:"Server,omitempty" xml:"Server,omitempty"`
	StandbyServer *string `json:"StandbyServer,omitempty" xml:"StandbyServer,omitempty"`
}

func (s VerifyInstanceADAuthServerRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyInstanceADAuthServerRequest) GoString() string {
	return s.String()
}

func (s *VerifyInstanceADAuthServerRequest) SetAccount(v string) *VerifyInstanceADAuthServerRequest {
	s.Account = &v
	return s
}

func (s *VerifyInstanceADAuthServerRequest) SetBaseDN(v string) *VerifyInstanceADAuthServerRequest {
	s.BaseDN = &v
	return s
}

func (s *VerifyInstanceADAuthServerRequest) SetDomain(v string) *VerifyInstanceADAuthServerRequest {
	s.Domain = &v
	return s
}

func (s *VerifyInstanceADAuthServerRequest) SetFilter(v string) *VerifyInstanceADAuthServerRequest {
	s.Filter = &v
	return s
}

func (s *VerifyInstanceADAuthServerRequest) SetInstanceId(v string) *VerifyInstanceADAuthServerRequest {
	s.InstanceId = &v
	return s
}

func (s *VerifyInstanceADAuthServerRequest) SetIsSSL(v string) *VerifyInstanceADAuthServerRequest {
	s.IsSSL = &v
	return s
}

func (s *VerifyInstanceADAuthServerRequest) SetPassword(v string) *VerifyInstanceADAuthServerRequest {
	s.Password = &v
	return s
}

func (s *VerifyInstanceADAuthServerRequest) SetPort(v string) *VerifyInstanceADAuthServerRequest {
	s.Port = &v
	return s
}

func (s *VerifyInstanceADAuthServerRequest) SetRegionId(v string) *VerifyInstanceADAuthServerRequest {
	s.RegionId = &v
	return s
}

func (s *VerifyInstanceADAuthServerRequest) SetServer(v string) *VerifyInstanceADAuthServerRequest {
	s.Server = &v
	return s
}

func (s *VerifyInstanceADAuthServerRequest) SetStandbyServer(v string) *VerifyInstanceADAuthServerRequest {
	s.StandbyServer = &v
	return s
}

type VerifyInstanceADAuthServerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyInstanceADAuthServerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyInstanceADAuthServerResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyInstanceADAuthServerResponseBody) SetRequestId(v string) *VerifyInstanceADAuthServerResponseBody {
	s.RequestId = &v
	return s
}

type VerifyInstanceADAuthServerResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *VerifyInstanceADAuthServerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VerifyInstanceADAuthServerResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyInstanceADAuthServerResponse) GoString() string {
	return s.String()
}

func (s *VerifyInstanceADAuthServerResponse) SetHeaders(v map[string]*string) *VerifyInstanceADAuthServerResponse {
	s.Headers = v
	return s
}

func (s *VerifyInstanceADAuthServerResponse) SetStatusCode(v int32) *VerifyInstanceADAuthServerResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyInstanceADAuthServerResponse) SetBody(v *VerifyInstanceADAuthServerResponseBody) *VerifyInstanceADAuthServerResponse {
	s.Body = v
	return s
}

type VerifyInstanceLDAPAuthServerRequest struct {
	Account       *string `json:"Account,omitempty" xml:"Account,omitempty"`
	BaseDN        *string `json:"BaseDN,omitempty" xml:"BaseDN,omitempty"`
	Filter        *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IsSSL         *string `json:"IsSSL,omitempty" xml:"IsSSL,omitempty"`
	Password      *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Port          *string `json:"Port,omitempty" xml:"Port,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Server        *string `json:"Server,omitempty" xml:"Server,omitempty"`
	StandbyServer *string `json:"StandbyServer,omitempty" xml:"StandbyServer,omitempty"`
}

func (s VerifyInstanceLDAPAuthServerRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyInstanceLDAPAuthServerRequest) GoString() string {
	return s.String()
}

func (s *VerifyInstanceLDAPAuthServerRequest) SetAccount(v string) *VerifyInstanceLDAPAuthServerRequest {
	s.Account = &v
	return s
}

func (s *VerifyInstanceLDAPAuthServerRequest) SetBaseDN(v string) *VerifyInstanceLDAPAuthServerRequest {
	s.BaseDN = &v
	return s
}

func (s *VerifyInstanceLDAPAuthServerRequest) SetFilter(v string) *VerifyInstanceLDAPAuthServerRequest {
	s.Filter = &v
	return s
}

func (s *VerifyInstanceLDAPAuthServerRequest) SetInstanceId(v string) *VerifyInstanceLDAPAuthServerRequest {
	s.InstanceId = &v
	return s
}

func (s *VerifyInstanceLDAPAuthServerRequest) SetIsSSL(v string) *VerifyInstanceLDAPAuthServerRequest {
	s.IsSSL = &v
	return s
}

func (s *VerifyInstanceLDAPAuthServerRequest) SetPassword(v string) *VerifyInstanceLDAPAuthServerRequest {
	s.Password = &v
	return s
}

func (s *VerifyInstanceLDAPAuthServerRequest) SetPort(v string) *VerifyInstanceLDAPAuthServerRequest {
	s.Port = &v
	return s
}

func (s *VerifyInstanceLDAPAuthServerRequest) SetRegionId(v string) *VerifyInstanceLDAPAuthServerRequest {
	s.RegionId = &v
	return s
}

func (s *VerifyInstanceLDAPAuthServerRequest) SetServer(v string) *VerifyInstanceLDAPAuthServerRequest {
	s.Server = &v
	return s
}

func (s *VerifyInstanceLDAPAuthServerRequest) SetStandbyServer(v string) *VerifyInstanceLDAPAuthServerRequest {
	s.StandbyServer = &v
	return s
}

type VerifyInstanceLDAPAuthServerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyInstanceLDAPAuthServerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyInstanceLDAPAuthServerResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyInstanceLDAPAuthServerResponseBody) SetRequestId(v string) *VerifyInstanceLDAPAuthServerResponseBody {
	s.RequestId = &v
	return s
}

type VerifyInstanceLDAPAuthServerResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *VerifyInstanceLDAPAuthServerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VerifyInstanceLDAPAuthServerResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyInstanceLDAPAuthServerResponse) GoString() string {
	return s.String()
}

func (s *VerifyInstanceLDAPAuthServerResponse) SetHeaders(v map[string]*string) *VerifyInstanceLDAPAuthServerResponse {
	s.Headers = v
	return s
}

func (s *VerifyInstanceLDAPAuthServerResponse) SetStatusCode(v int32) *VerifyInstanceLDAPAuthServerResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyInstanceLDAPAuthServerResponse) SetBody(v *VerifyInstanceLDAPAuthServerResponseBody) *VerifyInstanceLDAPAuthServerResponse {
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HostGroupId)) {
		query["HostGroupId"] = request.HostGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.HostIds)) {
		query["HostIds"] = request.HostIds
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddHostsToGroup"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddHostsToGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupId)) {
		query["UserGroupId"] = request.UserGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		query["UserIds"] = request.UserIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddUsersToGroup"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddUsersToGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) AttachHostAccountsToHostShareKeyWithOptions(request *AttachHostAccountsToHostShareKeyRequest, runtime *util.RuntimeOptions) (_result *AttachHostAccountsToHostShareKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HostAccountIds)) {
		query["HostAccountIds"] = request.HostAccountIds
	}

	if !tea.BoolValue(util.IsUnset(request.HostShareKeyId)) {
		query["HostShareKeyId"] = request.HostShareKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachHostAccountsToHostShareKey"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachHostAccountsToHostShareKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachHostAccountsToHostShareKey(request *AttachHostAccountsToHostShareKeyRequest) (_result *AttachHostAccountsToHostShareKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachHostAccountsToHostShareKeyResponse{}
	_body, _err := client.AttachHostAccountsToHostShareKeyWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Hosts)) {
		query["Hosts"] = request.Hosts
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachHostAccountsToUser"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachHostAccountsToUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Hosts)) {
		query["Hosts"] = request.Hosts
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupId)) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachHostAccountsToUserGroup"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachHostAccountsToUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HostGroups)) {
		query["HostGroups"] = request.HostGroups
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachHostGroupAccountsToUser"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachHostGroupAccountsToUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HostGroups)) {
		query["HostGroups"] = request.HostGroups
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupId)) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachHostGroupAccountsToUserGroup"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachHostGroupAccountsToUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthorizedSecurityGroups)) {
		query["AuthorizedSecurityGroups"] = request.AuthorizedSecurityGroups
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfigInstanceSecurityGroups"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfigInstanceSecurityGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.WhiteList)) {
		query["WhiteList"] = request.WhiteList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfigInstanceWhiteList"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfigInstanceWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActiveAddressType)) {
		query["ActiveAddressType"] = request.ActiveAddressType
	}

	if !tea.BoolValue(util.IsUnset(request.Comment)) {
		query["Comment"] = request.Comment
	}

	if !tea.BoolValue(util.IsUnset(request.HostName)) {
		query["HostName"] = request.HostName
	}

	if !tea.BoolValue(util.IsUnset(request.HostPrivateAddress)) {
		query["HostPrivateAddress"] = request.HostPrivateAddress
	}

	if !tea.BoolValue(util.IsUnset(request.HostPublicAddress)) {
		query["HostPublicAddress"] = request.HostPublicAddress
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceRegionId)) {
		query["InstanceRegionId"] = request.InstanceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.OSType)) {
		query["OSType"] = request.OSType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SourceInstanceId)) {
		query["SourceInstanceId"] = request.SourceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateHost"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateHostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HostAccountName)) {
		query["HostAccountName"] = request.HostAccountName
	}

	if !tea.BoolValue(util.IsUnset(request.HostId)) {
		query["HostId"] = request.HostId
	}

	if !tea.BoolValue(util.IsUnset(request.HostShareKeyId)) {
		query["HostShareKeyId"] = request.HostShareKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PassPhrase)) {
		query["PassPhrase"] = request.PassPhrase
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateKey)) {
		query["PrivateKey"] = request.PrivateKey
	}

	if !tea.BoolValue(util.IsUnset(request.ProtocolName)) {
		query["ProtocolName"] = request.ProtocolName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateHostAccount"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateHostAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Comment)) {
		query["Comment"] = request.Comment
	}

	if !tea.BoolValue(util.IsUnset(request.HostGroupName)) {
		query["HostGroupName"] = request.HostGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateHostGroup"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateHostGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) CreateHostShareKeyWithOptions(request *CreateHostShareKeyRequest, runtime *util.RuntimeOptions) (_result *CreateHostShareKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HostShareKeyName)) {
		query["HostShareKeyName"] = request.HostShareKeyName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PassPhrase)) {
		query["PassPhrase"] = request.PassPhrase
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateKey)) {
		query["PrivateKey"] = request.PrivateKey
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateHostShareKey"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateHostShareKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateHostShareKey(request *CreateHostShareKeyRequest) (_result *CreateHostShareKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateHostShareKeyResponse{}
	_body, _err := client.CreateHostShareKeyWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Comment)) {
		query["Comment"] = request.Comment
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		query["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		query["Mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.MobileCountryCode)) {
		query["MobileCountryCode"] = request.MobileCountryCode
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SourceUserId)) {
		query["SourceUserId"] = request.SourceUserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUser"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Comment)) {
		query["Comment"] = request.Comment
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupName)) {
		query["UserGroupName"] = request.UserGroupName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUserGroup"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HostId)) {
		query["HostId"] = request.HostId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteHost"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteHostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HostAccountId)) {
		query["HostAccountId"] = request.HostAccountId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteHostAccount"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteHostAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HostGroupId)) {
		query["HostGroupId"] = request.HostGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteHostGroup"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteHostGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DeleteHostShareKeyWithOptions(request *DeleteHostShareKeyRequest, runtime *util.RuntimeOptions) (_result *DeleteHostShareKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HostShareKeyId)) {
		query["HostShareKeyId"] = request.HostShareKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteHostShareKey"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteHostShareKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteHostShareKey(request *DeleteHostShareKeyRequest) (_result *DeleteHostShareKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteHostShareKeyResponse{}
	_body, _err := client.DeleteHostShareKeyWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUser"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupId)) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUserGroup"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceAttribute"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceStatus)) {
		query["InstanceStatus"] = request.InstanceStatus
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstances"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DetachHostAccountsFromHostShareKeyWithOptions(request *DetachHostAccountsFromHostShareKeyRequest, runtime *util.RuntimeOptions) (_result *DetachHostAccountsFromHostShareKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HostAccountIds)) {
		query["HostAccountIds"] = request.HostAccountIds
	}

	if !tea.BoolValue(util.IsUnset(request.HostShareKeyId)) {
		query["HostShareKeyId"] = request.HostShareKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachHostAccountsFromHostShareKey"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachHostAccountsFromHostShareKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetachHostAccountsFromHostShareKey(request *DetachHostAccountsFromHostShareKeyRequest) (_result *DetachHostAccountsFromHostShareKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachHostAccountsFromHostShareKeyResponse{}
	_body, _err := client.DetachHostAccountsFromHostShareKeyWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Hosts)) {
		query["Hosts"] = request.Hosts
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachHostAccountsFromUser"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachHostAccountsFromUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Hosts)) {
		query["Hosts"] = request.Hosts
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupId)) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachHostAccountsFromUserGroup"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachHostAccountsFromUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HostGroups)) {
		query["HostGroups"] = request.HostGroups
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachHostGroupAccountsFromUser"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachHostGroupAccountsFromUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HostGroups)) {
		query["HostGroups"] = request.HostGroups
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupId)) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachHostGroupAccountsFromUserGroup"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachHostGroupAccountsFromUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableInstancePublicAccess"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableInstancePublicAccessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableInstancePublicAccess"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableInstancePublicAccessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HostId)) {
		query["HostId"] = request.HostId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHost"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HostAccountId)) {
		query["HostAccountId"] = request.HostAccountId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHostAccount"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHostAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HostGroupId)) {
		query["HostGroupId"] = request.HostGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHostGroup"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHostGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) GetHostShareKeyWithOptions(request *GetHostShareKeyRequest, runtime *util.RuntimeOptions) (_result *GetHostShareKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HostShareKeyId)) {
		query["HostShareKeyId"] = request.HostShareKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHostShareKey"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHostShareKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHostShareKey(request *GetHostShareKeyRequest) (_result *GetHostShareKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHostShareKeyResponse{}
	_body, _err := client.GetHostShareKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstanceADAuthServerWithOptions(request *GetInstanceADAuthServerRequest, runtime *util.RuntimeOptions) (_result *GetInstanceADAuthServerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstanceADAuthServer"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceADAuthServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstanceADAuthServer(request *GetInstanceADAuthServerRequest) (_result *GetInstanceADAuthServerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInstanceADAuthServerResponse{}
	_body, _err := client.GetInstanceADAuthServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstanceLDAPAuthServerWithOptions(request *GetInstanceLDAPAuthServerRequest, runtime *util.RuntimeOptions) (_result *GetInstanceLDAPAuthServerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstanceLDAPAuthServer"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceLDAPAuthServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstanceLDAPAuthServer(request *GetInstanceLDAPAuthServerRequest) (_result *GetInstanceLDAPAuthServerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInstanceLDAPAuthServerResponse{}
	_body, _err := client.GetInstanceLDAPAuthServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstanceTwoFactorWithOptions(request *GetInstanceTwoFactorRequest, runtime *util.RuntimeOptions) (_result *GetInstanceTwoFactorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstanceTwoFactor"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceTwoFactorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstanceTwoFactor(request *GetInstanceTwoFactorRequest) (_result *GetInstanceTwoFactorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInstanceTwoFactorResponse{}
	_body, _err := client.GetInstanceTwoFactorWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUser"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupId)) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserGroup"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HostAccountName)) {
		query["HostAccountName"] = request.HostAccountName
	}

	if !tea.BoolValue(util.IsUnset(request.HostId)) {
		query["HostId"] = request.HostId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProtocolName)) {
		query["ProtocolName"] = request.ProtocolName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHostAccounts"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHostAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ListHostAccountsForHostShareKeyWithOptions(request *ListHostAccountsForHostShareKeyRequest, runtime *util.RuntimeOptions) (_result *ListHostAccountsForHostShareKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HostShareKeyId)) {
		query["HostShareKeyId"] = request.HostShareKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHostAccountsForHostShareKey"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHostAccountsForHostShareKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListHostAccountsForHostShareKey(request *ListHostAccountsForHostShareKeyRequest) (_result *ListHostAccountsForHostShareKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHostAccountsForHostShareKeyResponse{}
	_body, _err := client.ListHostAccountsForHostShareKeyWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HostAccountName)) {
		query["HostAccountName"] = request.HostAccountName
	}

	if !tea.BoolValue(util.IsUnset(request.HostId)) {
		query["HostId"] = request.HostId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHostAccountsForUser"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHostAccountsForUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HostAccountName)) {
		query["HostAccountName"] = request.HostAccountName
	}

	if !tea.BoolValue(util.IsUnset(request.HostId)) {
		query["HostId"] = request.HostId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupId)) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHostAccountsForUserGroup"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHostAccountsForUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HostGroupId)) {
		query["HostGroupId"] = request.HostGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHostGroupAccountNamesForUser"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHostGroupAccountNamesForUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HostGroupId)) {
		query["HostGroupId"] = request.HostGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupId)) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHostGroupAccountNamesForUserGroup"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHostGroupAccountNamesForUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HostGroupName)) {
		query["HostGroupName"] = request.HostGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHostGroups"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHostGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HostGroupName)) {
		query["HostGroupName"] = request.HostGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		query["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHostGroupsForUser"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHostGroupsForUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HostGroupName)) {
		query["HostGroupName"] = request.HostGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		query["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupId)) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHostGroupsForUserGroup"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHostGroupsForUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ListHostShareKeysWithOptions(request *ListHostShareKeysRequest, runtime *util.RuntimeOptions) (_result *ListHostShareKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHostShareKeys"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHostShareKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListHostShareKeys(request *ListHostShareKeysRequest) (_result *ListHostShareKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHostShareKeysResponse{}
	_body, _err := client.ListHostShareKeysWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HostAddress)) {
		query["HostAddress"] = request.HostAddress
	}

	if !tea.BoolValue(util.IsUnset(request.HostGroupId)) {
		query["HostGroupId"] = request.HostGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.HostName)) {
		query["HostName"] = request.HostName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OSType)) {
		query["OSType"] = request.OSType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SourceInstanceId)) {
		query["SourceInstanceId"] = request.SourceInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceInstanceState)) {
		query["SourceInstanceState"] = request.SourceInstanceState
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHosts"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHostsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HostAddress)) {
		query["HostAddress"] = request.HostAddress
	}

	if !tea.BoolValue(util.IsUnset(request.HostName)) {
		query["HostName"] = request.HostName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		query["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.OSType)) {
		query["OSType"] = request.OSType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHostsForUser"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHostsForUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HostAddress)) {
		query["HostAddress"] = request.HostAddress
	}

	if !tea.BoolValue(util.IsUnset(request.HostName)) {
		query["HostName"] = request.HostName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		query["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.OSType)) {
		query["OSType"] = request.OSType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupId)) {
		query["UserGroupId"] = request.UserGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHostsForUserGroup"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHostsForUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagKeys"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupName)) {
		query["UserGroupName"] = request.UserGroupName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserGroups"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		query["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		query["Mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SourceUserId)) {
		query["SourceUserId"] = request.SourceUserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupId)) {
		query["UserGroupId"] = request.UserGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		query["UserName"] = request.UserName
	}

	if !tea.BoolValue(util.IsUnset(request.UserState)) {
		query["UserState"] = request.UserState
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUsers"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		query["UserIds"] = request.UserIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("LockUsers"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &LockUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Comment)) {
		query["Comment"] = request.Comment
	}

	if !tea.BoolValue(util.IsUnset(request.HostId)) {
		query["HostId"] = request.HostId
	}

	if !tea.BoolValue(util.IsUnset(request.HostName)) {
		query["HostName"] = request.HostName
	}

	if !tea.BoolValue(util.IsUnset(request.HostPrivateAddress)) {
		query["HostPrivateAddress"] = request.HostPrivateAddress
	}

	if !tea.BoolValue(util.IsUnset(request.HostPublicAddress)) {
		query["HostPublicAddress"] = request.HostPublicAddress
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OSType)) {
		query["OSType"] = request.OSType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyHost"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyHostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HostAccountId)) {
		query["HostAccountId"] = request.HostAccountId
	}

	if !tea.BoolValue(util.IsUnset(request.HostAccountName)) {
		query["HostAccountName"] = request.HostAccountName
	}

	if !tea.BoolValue(util.IsUnset(request.HostShareKeyId)) {
		query["HostShareKeyId"] = request.HostShareKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PassPhrase)) {
		query["PassPhrase"] = request.PassPhrase
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateKey)) {
		query["PrivateKey"] = request.PrivateKey
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyHostAccount"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyHostAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Comment)) {
		query["Comment"] = request.Comment
	}

	if !tea.BoolValue(util.IsUnset(request.HostGroupId)) {
		query["HostGroupId"] = request.HostGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.HostGroupName)) {
		query["HostGroupName"] = request.HostGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyHostGroup"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyHostGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ModifyHostShareKeyWithOptions(request *ModifyHostShareKeyRequest, runtime *util.RuntimeOptions) (_result *ModifyHostShareKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HostShareKeyId)) {
		query["HostShareKeyId"] = request.HostShareKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.HostShareKeyName)) {
		query["HostShareKeyName"] = request.HostShareKeyName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PassPhrase)) {
		query["PassPhrase"] = request.PassPhrase
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateKey)) {
		query["PrivateKey"] = request.PrivateKey
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyHostShareKey"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyHostShareKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyHostShareKey(request *ModifyHostShareKeyRequest) (_result *ModifyHostShareKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyHostShareKeyResponse{}
	_body, _err := client.ModifyHostShareKeyWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActiveAddressType)) {
		query["ActiveAddressType"] = request.ActiveAddressType
	}

	if !tea.BoolValue(util.IsUnset(request.HostIds)) {
		query["HostIds"] = request.HostIds
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyHostsActiveAddressType"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyHostsActiveAddressTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HostIds)) {
		query["HostIds"] = request.HostIds
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		query["Port"] = request.Port
	}

	if !tea.BoolValue(util.IsUnset(request.ProtocolName)) {
		query["ProtocolName"] = request.ProtocolName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyHostsPort"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyHostsPortResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ModifyInstanceADAuthServerWithOptions(request *ModifyInstanceADAuthServerRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceADAuthServerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Account)) {
		query["Account"] = request.Account
	}

	if !tea.BoolValue(util.IsUnset(request.BaseDN)) {
		query["BaseDN"] = request.BaseDN
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.EmailMapping)) {
		query["EmailMapping"] = request.EmailMapping
	}

	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsSSL)) {
		query["IsSSL"] = request.IsSSL
	}

	if !tea.BoolValue(util.IsUnset(request.MobileMapping)) {
		query["MobileMapping"] = request.MobileMapping
	}

	if !tea.BoolValue(util.IsUnset(request.NameMapping)) {
		query["NameMapping"] = request.NameMapping
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		query["Port"] = request.Port
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Server)) {
		query["Server"] = request.Server
	}

	if !tea.BoolValue(util.IsUnset(request.StandbyServer)) {
		query["StandbyServer"] = request.StandbyServer
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyInstanceADAuthServer"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyInstanceADAuthServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyInstanceADAuthServer(request *ModifyInstanceADAuthServerRequest) (_result *ModifyInstanceADAuthServerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceADAuthServerResponse{}
	_body, _err := client.ModifyInstanceADAuthServerWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyInstanceAttribute"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyInstanceAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ModifyInstanceLDAPAuthServerWithOptions(request *ModifyInstanceLDAPAuthServerRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceLDAPAuthServerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Account)) {
		query["Account"] = request.Account
	}

	if !tea.BoolValue(util.IsUnset(request.BaseDN)) {
		query["BaseDN"] = request.BaseDN
	}

	if !tea.BoolValue(util.IsUnset(request.EmailMapping)) {
		query["EmailMapping"] = request.EmailMapping
	}

	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsSSL)) {
		query["IsSSL"] = request.IsSSL
	}

	if !tea.BoolValue(util.IsUnset(request.LoginNameMapping)) {
		query["LoginNameMapping"] = request.LoginNameMapping
	}

	if !tea.BoolValue(util.IsUnset(request.MobileMapping)) {
		query["MobileMapping"] = request.MobileMapping
	}

	if !tea.BoolValue(util.IsUnset(request.NameMapping)) {
		query["NameMapping"] = request.NameMapping
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		query["Port"] = request.Port
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Server)) {
		query["Server"] = request.Server
	}

	if !tea.BoolValue(util.IsUnset(request.StandbyServer)) {
		query["StandbyServer"] = request.StandbyServer
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyInstanceLDAPAuthServer"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyInstanceLDAPAuthServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyInstanceLDAPAuthServer(request *ModifyInstanceLDAPAuthServerRequest) (_result *ModifyInstanceLDAPAuthServerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceLDAPAuthServerResponse{}
	_body, _err := client.ModifyInstanceLDAPAuthServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyInstanceTwoFactorWithOptions(request *ModifyInstanceTwoFactorRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceTwoFactorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DingTalkConfig)) {
		query["DingTalkConfig"] = request.DingTalkConfig
	}

	if !tea.BoolValue(util.IsUnset(request.EnableTwoFactor)) {
		query["EnableTwoFactor"] = request.EnableTwoFactor
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MessageLanguage)) {
		query["MessageLanguage"] = request.MessageLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SkipTwoFactorTime)) {
		query["SkipTwoFactorTime"] = request.SkipTwoFactorTime
	}

	if !tea.BoolValue(util.IsUnset(request.TwoFactorMethods)) {
		query["TwoFactorMethods"] = request.TwoFactorMethods
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyInstanceTwoFactor"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyInstanceTwoFactorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyInstanceTwoFactor(request *ModifyInstanceTwoFactorRequest) (_result *ModifyInstanceTwoFactorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceTwoFactorResponse{}
	_body, _err := client.ModifyInstanceTwoFactorWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Comment)) {
		query["Comment"] = request.Comment
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		query["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		query["Mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.MobileCountryCode)) {
		query["MobileCountryCode"] = request.MobileCountryCode
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyUser"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Comment)) {
		query["Comment"] = request.Comment
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupId)) {
		query["UserGroupId"] = request.UserGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupName)) {
		query["UserGroupName"] = request.UserGroupName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyUserGroup"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MoveResourceGroup"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HostGroupId)) {
		query["HostGroupId"] = request.HostGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.HostIds)) {
		query["HostIds"] = request.HostIds
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveHostsFromGroup"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveHostsFromGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupId)) {
		query["UserGroupId"] = request.UserGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		query["UserIds"] = request.UserIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveUsersFromGroup"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveUsersFromGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CredentialType)) {
		query["CredentialType"] = request.CredentialType
	}

	if !tea.BoolValue(util.IsUnset(request.HostAccountId)) {
		query["HostAccountId"] = request.HostAccountId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetHostAccountCredential"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetHostAccountCredentialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupIds)) {
		query["SecurityGroupIds"] = request.SecurityGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.VswitchId)) {
		query["VswitchId"] = request.VswitchId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartInstance"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UserIds)) {
		query["UserIds"] = request.UserIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnlockUsers"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnlockUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) VerifyInstanceADAuthServerWithOptions(request *VerifyInstanceADAuthServerRequest, runtime *util.RuntimeOptions) (_result *VerifyInstanceADAuthServerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Account)) {
		query["Account"] = request.Account
	}

	if !tea.BoolValue(util.IsUnset(request.BaseDN)) {
		query["BaseDN"] = request.BaseDN
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsSSL)) {
		query["IsSSL"] = request.IsSSL
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		query["Port"] = request.Port
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Server)) {
		query["Server"] = request.Server
	}

	if !tea.BoolValue(util.IsUnset(request.StandbyServer)) {
		query["StandbyServer"] = request.StandbyServer
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyInstanceADAuthServer"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyInstanceADAuthServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyInstanceADAuthServer(request *VerifyInstanceADAuthServerRequest) (_result *VerifyInstanceADAuthServerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyInstanceADAuthServerResponse{}
	_body, _err := client.VerifyInstanceADAuthServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyInstanceLDAPAuthServerWithOptions(request *VerifyInstanceLDAPAuthServerRequest, runtime *util.RuntimeOptions) (_result *VerifyInstanceLDAPAuthServerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Account)) {
		query["Account"] = request.Account
	}

	if !tea.BoolValue(util.IsUnset(request.BaseDN)) {
		query["BaseDN"] = request.BaseDN
	}

	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsSSL)) {
		query["IsSSL"] = request.IsSSL
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		query["Port"] = request.Port
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Server)) {
		query["Server"] = request.Server
	}

	if !tea.BoolValue(util.IsUnset(request.StandbyServer)) {
		query["StandbyServer"] = request.StandbyServer
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyInstanceLDAPAuthServer"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyInstanceLDAPAuthServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyInstanceLDAPAuthServer(request *VerifyInstanceLDAPAuthServerRequest) (_result *VerifyInstanceLDAPAuthServerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyInstanceLDAPAuthServerResponse{}
	_body, _err := client.VerifyInstanceLDAPAuthServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
