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

type AcceptApproveCommandRequest struct {
	CommandId  *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AcceptApproveCommandRequest) String() string {
	return tea.Prettify(s)
}

func (s AcceptApproveCommandRequest) GoString() string {
	return s.String()
}

func (s *AcceptApproveCommandRequest) SetCommandId(v string) *AcceptApproveCommandRequest {
	s.CommandId = &v
	return s
}

func (s *AcceptApproveCommandRequest) SetInstanceId(v string) *AcceptApproveCommandRequest {
	s.InstanceId = &v
	return s
}

func (s *AcceptApproveCommandRequest) SetRegionId(v string) *AcceptApproveCommandRequest {
	s.RegionId = &v
	return s
}

type AcceptApproveCommandResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AcceptApproveCommandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AcceptApproveCommandResponseBody) GoString() string {
	return s.String()
}

func (s *AcceptApproveCommandResponseBody) SetRequestId(v string) *AcceptApproveCommandResponseBody {
	s.RequestId = &v
	return s
}

type AcceptApproveCommandResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AcceptApproveCommandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AcceptApproveCommandResponse) String() string {
	return tea.Prettify(s)
}

func (s AcceptApproveCommandResponse) GoString() string {
	return s.String()
}

func (s *AcceptApproveCommandResponse) SetHeaders(v map[string]*string) *AcceptApproveCommandResponse {
	s.Headers = v
	return s
}

func (s *AcceptApproveCommandResponse) SetStatusCode(v int32) *AcceptApproveCommandResponse {
	s.StatusCode = &v
	return s
}

func (s *AcceptApproveCommandResponse) SetBody(v *AcceptApproveCommandResponseBody) *AcceptApproveCommandResponse {
	s.Body = v
	return s
}

type AcceptOperationTicketRequest struct {
	// The maximum number of logons allowed. Valid values:
	//
	// *   0: The number of logons is unlimited. The O\&M engineer can log on to the specified asset for an unlimited number of times during the validity period.
	// *   1: The O\&M engineer can log on to the specified asset only once during the validity period.
	//
	// >  You can set this parameter only to 0 if you review an O\&M application on a database.
	EffectCount *string `json:"EffectCount,omitempty" xml:"EffectCount,omitempty"`
	// The end time of the validity period. The value is a UNIX timestamp. Unit: seconds.
	EffectEndTime *string `json:"EffectEndTime,omitempty" xml:"EffectEndTime,omitempty"`
	// The start time of the validity period. The value is a UNIX timestamp. Unit: seconds.
	EffectStartTime *string `json:"EffectStartTime,omitempty" xml:"EffectStartTime,omitempty"`
	// The ID of the bastion host.
	//
	// > You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the O\&M application that you want to approve. You can call the ListOperationTickets operation to query the IDs of all O\&M applications that require review.
	OperationTicketId *string `json:"OperationTicketId,omitempty" xml:"OperationTicketId,omitempty"`
	// The region ID of the bastion host.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AcceptOperationTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s AcceptOperationTicketRequest) GoString() string {
	return s.String()
}

func (s *AcceptOperationTicketRequest) SetEffectCount(v string) *AcceptOperationTicketRequest {
	s.EffectCount = &v
	return s
}

func (s *AcceptOperationTicketRequest) SetEffectEndTime(v string) *AcceptOperationTicketRequest {
	s.EffectEndTime = &v
	return s
}

func (s *AcceptOperationTicketRequest) SetEffectStartTime(v string) *AcceptOperationTicketRequest {
	s.EffectStartTime = &v
	return s
}

func (s *AcceptOperationTicketRequest) SetInstanceId(v string) *AcceptOperationTicketRequest {
	s.InstanceId = &v
	return s
}

func (s *AcceptOperationTicketRequest) SetOperationTicketId(v string) *AcceptOperationTicketRequest {
	s.OperationTicketId = &v
	return s
}

func (s *AcceptOperationTicketRequest) SetRegionId(v string) *AcceptOperationTicketRequest {
	s.RegionId = &v
	return s
}

type AcceptOperationTicketResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AcceptOperationTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AcceptOperationTicketResponseBody) GoString() string {
	return s.String()
}

func (s *AcceptOperationTicketResponseBody) SetRequestId(v string) *AcceptOperationTicketResponseBody {
	s.RequestId = &v
	return s
}

type AcceptOperationTicketResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AcceptOperationTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AcceptOperationTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s AcceptOperationTicketResponse) GoString() string {
	return s.String()
}

func (s *AcceptOperationTicketResponse) SetHeaders(v map[string]*string) *AcceptOperationTicketResponse {
	s.Headers = v
	return s
}

func (s *AcceptOperationTicketResponse) SetStatusCode(v int32) *AcceptOperationTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *AcceptOperationTicketResponse) SetBody(v *AcceptOperationTicketResponseBody) *AcceptOperationTicketResponse {
	s.Body = v
	return s
}

type AddHostsToGroupRequest struct {
	// The ID of the host group to which you want to add hosts.
	//
	// > You can call the [ListHostGroups](~~201307~~) operation to query the ID of the host group.
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// The ID of the host that you want to add to the host group. The value is a JSON string. You can add up to 100 host IDs.
	//
	// > You can call the [ListHosts](~~200665~~) operation to query the IDs of hosts.
	HostIds *string `json:"HostIds,omitempty" xml:"HostIds,omitempty"`
	// The ID of the bastion host for which you want to add hosts to the host group.
	//
	// > You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host for which you want to add hosts to the host group.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the call.
	Results []*AddHostsToGroupResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
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
	// The return code that indicates whether the call was successful. Valid values:
	//
	// *   **OK**: The call was successful.
	//
	// *   **UNEXPECTED**: An unknown error occurred.
	//
	// *   **INVALID_ARGUMENT**: A request parameter is invalid.
	//
	// >Make sure that the request parameters are valid and call the operation again.
	//
	// *   **OBJECT_NOT_FOUND**: The specified object on which you want to perform the operation does not exist.
	//
	// >Check whether the specified ID of the bastion host exists, whether the specified hosts exist, and whether the specified host IDs are valid. Then, call the operation again.
	//
	// *   **OBJECT_AlREADY_EXISTS**: The specified object on which you want to perform the operation already exists.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the host group.
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// The ID of the host.
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// This parameter is deprecated.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddHostsToGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the bastion host for which you want to add users to the user group.
	//
	// > You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host for which you want to add users to the user group.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user group to which you want to add users.
	//
	// > You can call the [ListUserGroups](~~204509~~) operation to query the ID of the user group.
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// The ID of the user that you want to add to the user group. The value is a JSON string. You can add up to 100 user IDs. If you specify multiple IDs, separate the IDs with commas (,).
	//
	// > You can call the [ListUsers](~~204522~~) operation to query the ID of the user.
	UserIds *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
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
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the call.
	Results []*AddUsersToGroupResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
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
	// The return code that indicates whether the call was successful. Valid values:
	//
	// *   **OK**: The call was successful.
	//
	// *   **UNEXPECTED**: An unknown error occurred.
	//
	// *   **INVALID_ARGUMENT**: A request parameter is invalid.
	//
	//     **
	//
	//     **Note**Make sure that the request parameters are valid and call the operation again.
	//
	// *   **OBJECT_NOT_FOUND**: The specified object on which you want to perform the operation does not exist.
	//
	//     **
	//
	//     **Note**Check whether the specified ID of the bastion host exists, whether the specified hosts exist, and whether the specified host IDs are valid. Then, call the operation again.
	//
	// *   **OBJECT_AlREADY_EXISTS**: The specified object on which you want to perform the operation already exists.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// This parameter is deprecated.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the group.
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// The ID of the user.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddUsersToGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The IDs of the host accounts.
	//
	// > You must specify this parameter.
	HostAccountIds *string `json:"HostAccountIds,omitempty" xml:"HostAccountIds,omitempty"`
	// The ID of the shared key.
	//
	// > You must specify this parameter.
	HostShareKeyId *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	// The ID of the bastion host. You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host. For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the call.
	Results []*AttachHostAccountsToHostShareKeyResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
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
	// The error code returned. If **OK** is returned, the association was successful. If another error code is returned, the association failed.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the host account.
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	// The ID of the shared key.
	HostShareKeyId *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	// The error message returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachHostAccountsToHostShareKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The IDs of the host and host account that you want to authorize the user to manage. You can specify up to 10 host IDs and up to 10 host account IDs for each host. You can specify only host IDs. In this case, the user is authorized to manage only the specified hosts. For more information about this parameter, see the "Description of the Hosts parameter" section of this topic.
	//
	// >  You can call the [ListHosts](~~200665~~) operation to query the ID of the host and the [ListHostAccounts](~~204372~~) operation to query the ID of the host account.
	Hosts *string `json:"Hosts,omitempty" xml:"Hosts,omitempty"`
	// The ID of the bastion host for which you want to authorize the user to manage the specified hosts and host accounts.
	//
	// >  You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host for which you want to authorize the user to manage the specified hosts and host accounts.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user that you want to authorize to manage the specified hosts and host accounts.
	//
	// >  You can call the [ListUsers](~~204522~~) operation to query the ID of the user.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the call.
	Results []*AttachHostAccountsToUserResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
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
	// The return code that indicates whether the call was successful. Valid values:
	//
	// *   **OK**: The call was successful.
	//
	// *   **UNEXPECTED**: An unknown error occurred.
	//
	// *   **INVALID_ARGUMENT**: A request parameter is invalid.
	//
	// > Make sure that the request parameters are valid and call the operation again.
	//
	// *   **OBJECT_NOT_FOUND**: The specified object on which you want to perform the operation does not exist.
	//
	// > Check whether the specified ID of the bastion host exists, whether the specified hosts exist, and whether the specified host IDs are valid. Then, call the operation again.
	//
	// *   **OBJECT_AlREADY_EXISTS**: The specified object on which you want to perform the operation already exists.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of authorizing the specified user to manage the specified host accounts.
	HostAccounts []*AttachHostAccountsToUserResponseBodyResultsHostAccounts `json:"HostAccounts,omitempty" xml:"HostAccounts,omitempty" type:"Repeated"`
	// The ID of the host.
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// This parameter is deprecated.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the user.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// The return code that indicates whether the user was authorized to manage the specified host account. Valid values:
	//
	// *   **OK**: The call was successful.
	// *   **UNEXPECTED**: An unknown error occurred.
	// *   **INVALID_ARGUMENT**: A request parameter is invalid.
	// *   **OBJECT_NOT_FOUND**: The specified object on which you want to perform the operation does not exist.
	// *   **OBJECT_AlREADY_EXISTS**: The specified object on which you want to perform the operation already exists.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the host account.
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	// This parameter is deprecated.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachHostAccountsToUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The IDs of the host and host account that you want to authorize the user group to manage. You can specify up to 10 host IDs and up to 10 host account IDs for each host. You can specify only host IDs. In this case, the user group is authorized to manage only the specified hosts. For more information about this parameter, see the "Description of the Hosts parameter" section of this topic.
	//
	// > You can call the [ListHosts](~~200665~~) operation to query the ID of the host and the [ListHostAccounts](~~204372~~) operation to query the ID of the host account.
	Hosts *string `json:"Hosts,omitempty" xml:"Hosts,omitempty"`
	// The ID of the bastion host in which you want to authorize the user group to manage the specified hosts and host accounts.
	//
	// > You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host in which you want to authorize the user group to manage the specified hosts and host accounts.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user group that you want to authorize to manage the specified hosts and host accounts.
	//
	// > You can call the [ListUserGroups](~~204509~~) operation to query the ID of the user group.
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
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the call.
	Results []*AttachHostAccountsToUserGroupResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
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
	// The return code that indicates whether the call was successful. Valid values:
	//
	// *   **OK**: The call was successful.
	// *   **UNEXPECTED**: An unknown error occurred.
	// *   **INVALID_ARGUMENT**: A request parameter is invalid.
	// *   **OBJECT_NOT_FOUND**: The specified object on which you want to perform the operation does not exist.
	// *   **OBJECT_AlREADY_EXISTS**: The specified object on which you want to perform the operation already exists.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of authorizing the specified user group to manage the specified host accounts.
	HostAccounts []*AttachHostAccountsToUserGroupResponseBodyResultsHostAccounts `json:"HostAccounts,omitempty" xml:"HostAccounts,omitempty" type:"Repeated"`
	// The ID of the host.
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// This parameter is deprecated.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the user group.
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
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
	// The return code that indicates whether the user group was authorized to manage the specified host account. Valid values:
	//
	// *   **OK**: The call was successful.
	// *   **UNEXPECTED**: An unknown error occurred.
	// *   **INVALID_ARGUMENT**: A request parameter is invalid.
	// *   **OBJECT_NOT_FOUND**: The specified object on which you want to perform the operation does not exist.
	// *   **OBJECT_AlREADY_EXISTS**: The specified object on which you want to perform the operation already exists.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the host account.
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	// This parameter is deprecated.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachHostAccountsToUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the host group and the name of the host account that you want to authorize the user to manage. You can specify up to 10 host group IDs and up to 10 host account names for each host group. You can specify only host group IDs. In this case, the user is authorized to manage only the specified host groups. For more information about this parameter, see the "Description of the HostGroups parameter" section of this topic.
	//
	// >  You can call the [ListHostGroups](~~201307~~) operation to query the ID of the host group and the [ListHostAccounts](~~204372~~) operation to query the name of the host account.
	HostGroups *string `json:"HostGroups,omitempty" xml:"HostGroups,omitempty"`
	// The ID of the bastion host for which you want to authorize the user to manage the specified host groups and host accounts.
	//
	// >  You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host for which you want to authorize the user to manage the specified host groups and host accounts.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user that you want to authorize to manage the specified host groups and host accounts.
	//
	// >  You can call the [ListUsers](~~204522~~) operation to query the ID of the user.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the call.
	Results []*AttachHostGroupAccountsToUserResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
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
	// The return code that indicates whether the call was successful. Valid values:
	//
	// *   **OK**: The call was successful.
	// *   **UNEXPECTED**: An unknown error occurred.
	// *   **INVALID_ARGUMENT**: A request parameter is invalid.
	// *   **OBJECT_NOT_FOUND**: The specified object on which you want to perform the operation does not exist.
	// *   **OBJECT_AlREADY_EXISTS**: The specified object on which you want to perform the operation already exists.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of authorizing the user to manage the specified host accounts.
	HostAccountNames []*AttachHostGroupAccountsToUserResponseBodyResultsHostAccountNames `json:"HostAccountNames,omitempty" xml:"HostAccountNames,omitempty" type:"Repeated"`
	// The ID of the host group.
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// This parameter is deprecated.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the user.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// The return code that indicates whether the user was authorized to manage the specified host account. Valid values:
	//
	// *   **OK**: The call was successful.
	// *   **UNEXPECTED**: An unknown error occurred.
	// *   **INVALID_ARGUMENT**: A request parameter is invalid.
	// *   **OBJECT_NOT_FOUND**: The specified object on which you want to perform the operation does not exist.
	// *   **OBJECT_AlREADY_EXISTS**: The specified object on which you want to perform the operation already exists.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The name of the host account.
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	// This parameter is deprecated.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachHostGroupAccountsToUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the host group and the name of the host account that you want to authorize the user group to manage. You can specify up to 10 host group IDs and up to 10 host account names for each host group. You can specify only host group IDs. In this case, the user group is authorized to manage only the specified host groups. For more information about this parameter, see the "Description of the HostGroups parameter" section of this topic.
	//
	// >  You can call the [ListHostGroups](~~201307~~) operation to query the ID of the host group and the [ListHostAccounts](~~204372~~) operation to query the name of the host account.
	HostGroups *string `json:"HostGroups,omitempty" xml:"HostGroups,omitempty"`
	// The ID of the bastion host for which you want to authorize the user group to manage the specified host groups and host accounts.
	//
	// >  You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host for which you want to authorize the user group to manage the specified host groups and host accounts.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user group that you want to authorize to manage the specified host groups and host accounts.
	//
	// >  You can call the [ListUserGroups](~~204509~~) operation to query the ID of the user group.
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
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the call.
	Results []*AttachHostGroupAccountsToUserGroupResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
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
	// The return code that indicates whether the call was successful. Valid values:
	//
	// *   **OK**: The call was successful.
	// *   **UNEXPECTED**: An unknown error occurred.
	// *   **INVALID_ARGUMENT**: A request parameter is invalid.
	// *   **OBJECT_NOT_FOUND**: The specified object on which you want to perform the operation does not exist.
	// *   **OBJECT_AlREADY_EXISTS**: The specified object on which you want to perform the operation already exists.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of authorizing the user group to manage the specified host accounts.
	HostAccountNames []*AttachHostGroupAccountsToUserGroupResponseBodyResultsHostAccountNames `json:"HostAccountNames,omitempty" xml:"HostAccountNames,omitempty" type:"Repeated"`
	// The ID of the host group.
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// This parameter is deprecated.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the group.
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
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
	// The return code that indicates whether the user group was authorized to manage the specified host account. Valid values:
	//
	// *   **OK**: The call was successful.
	// *   **UNEXPECTED**: An unknown error occurred.
	// *   **INVALID_ARGUMENT**: A request parameter is invalid.
	// *   **OBJECT_NOT_FOUND**: The specified object on which you want to perform the operation does not exist.
	// *   **OBJECT_AlREADY_EXISTS**: The specified object on which you want to perform the operation already exists.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The name of the host account.
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	// This parameter is deprecated.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachHostGroupAccountsToUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// An array that consists of the IDs of authorized security groups.
	AuthorizedSecurityGroups []*string `json:"AuthorizedSecurityGroups,omitempty" xml:"AuthorizedSecurityGroups,omitempty" type:"Repeated"`
	// The ID of the bastion host.
	//
	// > You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// *   **zh**: Chinese
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region ID of the bastion host.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The ID of the bastion host for which security groups were configured.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigInstanceSecurityGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the bastion host for which a whitelist of public IP addresses is configured.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Configures a whitelist of public IP addresses for a bastion host.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// ConfigInstanceWhiteList
	WhiteList []*string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty" type:"Repeated"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigInstanceWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The endpoint type of the host that you want to create. Valid values:
	//
	// *   **Public**: public endpoint
	// *   **Private**: internal endpoint
	ActiveAddressType *string `json:"ActiveAddressType,omitempty" xml:"ActiveAddressType,omitempty"`
	// The description of the host that you want to create. The value can be up to 500 characters in length.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The name of the host that you want to create. The name can be up to 128 characters in length.
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The internal endpoint of the host that you want to create. You can set this parameter to a domain name or an IP address.
	//
	// > This parameter is required if the **ActiveAddressType** parameter is set to **Private**.
	HostPrivateAddress *string `json:"HostPrivateAddress,omitempty" xml:"HostPrivateAddress,omitempty"`
	// The public endpoint of the host that you want to create. You can set this parameter to a domain name or an IP address.
	//
	// > This parameter is required if the **ActiveAddressType** parameter is set to **Public**.
	HostPublicAddress *string `json:"HostPublicAddress,omitempty" xml:"HostPublicAddress,omitempty"`
	// The ID of the bastion host in which you want to create the host.
	//
	// > You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region to which the ECS instance or the host in an ApsaraDB MyBase dedicated cluster belongs.
	//
	// > This parameter is required if the **Source** parameter is set to **Ecs** or **Rds**.
	InstanceRegionId *string `json:"InstanceRegionId,omitempty" xml:"InstanceRegionId,omitempty"`
	// The ID of the network domain to which the host belongs.
	NetworkDomainId *string `json:"NetworkDomainId,omitempty" xml:"NetworkDomainId,omitempty"`
	// The operating system of the host that you want to create. Valid values:
	//
	// *   **Linux**
	// *   **Windows**
	OSType *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	// The region ID of the bastion host in which you want to create the host.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The source of the host that you want to create. Valid values:
	//
	// *   **Local**: a host in a data center
	// *   **Ecs**: an Elastic Compute Service (ECS) instance
	// *   **Rds**: a host in an ApsaraDB MyBase dedicated cluster
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The ID of the ECS instance or the host in an ApsaraDB MyBase dedicated cluster.
	//
	// > This parameter is required if the **Source** parameter is set to **Ecs** or **Rds**.
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
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

func (s *CreateHostRequest) SetNetworkDomainId(v string) *CreateHostRequest {
	s.NetworkDomainId = &v
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
	// The ID of the host.
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The ID of the request.
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHostResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The passphrase of the private key for the host account.
	//
	// >  You can specify this parameter when the ProtocolName parameter is set to SSH. If the ProtocolName parameter is set to RDP, you do not need to specify this parameter.
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	// The ID of the shared key.
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The protocol of the host to which you want to add a host account.
	//
	// Valid values:
	//
	// *   SSH
	// *   RDP
	HostShareKeyId *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	// master
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The private key of the host account. The value is a Base64-encoded string.
	//
	// >  This parameter takes effect only when the ProtocolName parameter is set to SSH. If the ProtocolName parameter is set to RDP, you do not need to specify this parameter. You can configure a password and a private key for the host account at the same time. If both a password and a private key are configured for the host account, Bastionhost preferentially uses the private key to log on to the host.
	PassPhrase *string `json:"PassPhrase,omitempty" xml:"PassPhrase,omitempty"`
	// The region ID of the bastion host in which you want to add a host account to the host.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The ID of the host to which you want to add a host account.
	//
	// >  You can call the [ListHosts](~~200665~~) operation to query the ID of the host.
	PrivateKey *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	// The ID of the bastion host in which you want to add a host account to the host.
	//
	// >  You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	ProtocolName *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
	// The password of the host account.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The ID of the request.
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	// The operation that you want to perform. Set the value to **CreateHostAccount**.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHostAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The description of the host group. The description can be up to 500 characters in length.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The name of the host group. The name can be up to 128 characters in length.
	HostGroupName *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
	// The ID of the bastion host on which you want to create a host group.
	//
	// > You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host on which you want to create a host group.
	//
	// **
	//
	// **For more information about the mapping between region IDs and region names, see **Regions and zones[.](~~40654~~)
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The ID of the host group.
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHostGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The name of the shared key that you want to create. The name can be a maximum of 128 characters in length.
	HostShareKeyName *string `json:"HostShareKeyName,omitempty" xml:"HostShareKeyName,omitempty"`
	// The ID of the bastion host. You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The password of the private key. The value is a Base64-encoded string.
	PassPhrase *string `json:"PassPhrase,omitempty" xml:"PassPhrase,omitempty"`
	// The private key. The value is a Base64-encoded string.
	PrivateKey *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	// The region ID of the bastion host. For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The ID of the shared key.
	HostShareKeyId *int64 `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHostShareKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The remarks of the user that you want to add. The remarks can be up to 500 characters in length.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The display name of the user that you want to add. This display name can be up to 128 characters in length.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The end of the validity period of the user. The value is a UNIX timestamp. Unit: seconds.
	EffectiveEndTime *int64 `json:"EffectiveEndTime,omitempty" xml:"EffectiveEndTime,omitempty"`
	// The beginning of the validity period of the user. The value is a UNIX timestamp. Unit: seconds.
	EffectiveStartTime *int64 `json:"EffectiveStartTime,omitempty" xml:"EffectiveStartTime,omitempty"`
	// The email address of the user that you want to add.
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The ID of the bastion host to which you want to add a user.
	//
	// >  You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Language       *string `json:"Language,omitempty" xml:"Language,omitempty"`
	LanguageStatus *string `json:"LanguageStatus,omitempty" xml:"LanguageStatus,omitempty"`
	// The mobile phone number of the user that you want to add.
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The country where the mobile number of the user is registered. Default value: **CN**. Valid values:
	//
	// *   **CN**: the Chinese mainland, whose country calling code is +86
	// *   **HK**: Hong Kong (China), whose country calling code is +852
	// *   **MO**: Macau (China), whose country calling code is +853
	// *   **TW**: Taiwan (China), whose country calling code is +886
	// *   **RU**: Russia, whose country calling code is +7
	// *   **SG**: Singapore, whose country calling code is +65
	// *   **MY**: Malaysia, whose country calling code is +60
	// *   **ID**: Indonesia, whose country calling code is +62
	// *   **DE**: Germany, whose country calling code is +49
	// *   **AU**: Australia, whose country calling code is +61
	// *   **US**: US, whose country calling code is +1
	// *   **AE**: United Arab Emirates, whose country calling code is +971
	// *   **JP**: Japan, whose country calling code is +81
	// *   **GB**: UK, whose country calling code is +44
	// *   **IN**: India, whose country calling code is +91
	// *   **KR**: Republic of Korea, whose country calling code is +82
	// *   **PH**: Philippines, whose country calling code is +63
	// *   **CH**: Switzerland, whose country calling code is +41
	// *   **SE**: Sweden, whose country calling code is +46
	MobileCountryCode *string `json:"MobileCountryCode,omitempty" xml:"MobileCountryCode,omitempty"`
	// Specifies whether password reset is required upon the next logon. Valid values:
	//
	// - true: yes
	//
	// - false: no
	NeedResetPassword *bool `json:"NeedResetPassword,omitempty" xml:"NeedResetPassword,omitempty"`
	// The logon password of the user that you want to add. The logon password can be up to 128 characters in length.
	//
	// >  This parameter is required if the **Source** parameter is set to **Local**.
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The region ID of the bastion host to which you want to add a user.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The source of the user that you want to add. Valid values:
	//
	// - **Local**: a local user
	// - **Ram**: a RAM user
	// - **AD** : an AD-authenticated user
	// - **LDAP** : an LDAP-authenticated user
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The unique identifier of the user that you want to add.
	//
	// >  This parameter uniquely identifies a RAM user of the bastion host. This parameter is required if the **Source** parameter is set to **Ram**. You can call the [ListUsers](~~28684~~) operation to obtain the unique identifier of the user from the **UserId** response parameter.
	SourceUserId *string `json:"SourceUserId,omitempty" xml:"SourceUserId,omitempty"`
	// The two-factor authentication method. You can select only one method. Valid values:
	//
	// *   **sms:** text message
	// *   **email:** email
	// *   **dingtalk:** DingTalk
	// *   **totp OTP:** time-based one-time password (TOTP) app
	//
	// > *   When the TwoFactorStatus parameter is set to Enable, you must specify one of the preceding values.
	TwoFactorMethods *string `json:"TwoFactorMethods,omitempty" xml:"TwoFactorMethods,omitempty"`
	// The two-factor authentication status of the user. Valid values:
	//
	// - Global: follows the global settings
	// - Disable: disables two-factor authentication
	// - Enable: enable two-factor authentication and follows settings of the single user
	TwoFactorStatus *string `json:"TwoFactorStatus,omitempty" xml:"TwoFactorStatus,omitempty"`
	// The logon name of the user that you want to add. The logon name can contain only letters, digits, and underscores (\_) and can be up to 128 characters in length.
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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

func (s *CreateUserRequest) SetEffectiveEndTime(v int64) *CreateUserRequest {
	s.EffectiveEndTime = &v
	return s
}

func (s *CreateUserRequest) SetEffectiveStartTime(v int64) *CreateUserRequest {
	s.EffectiveStartTime = &v
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

func (s *CreateUserRequest) SetLanguage(v string) *CreateUserRequest {
	s.Language = &v
	return s
}

func (s *CreateUserRequest) SetLanguageStatus(v string) *CreateUserRequest {
	s.LanguageStatus = &v
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

func (s *CreateUserRequest) SetNeedResetPassword(v bool) *CreateUserRequest {
	s.NeedResetPassword = &v
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

func (s *CreateUserRequest) SetTwoFactorMethods(v string) *CreateUserRequest {
	s.TwoFactorMethods = &v
	return s
}

func (s *CreateUserRequest) SetTwoFactorStatus(v string) *CreateUserRequest {
	s.TwoFactorStatus = &v
	return s
}

func (s *CreateUserRequest) SetUserName(v string) *CreateUserRequest {
	s.UserName = &v
	return s
}

type CreateUserResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the user.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The description of the user group. The description can be up to 500 characters in length.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the bastion host for which you want to create a user group.
	//
	// > You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host for which you want to create a user group.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the user group that you want to create. This name can be a up to 128 characters in length.
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
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the user group.
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type CreateUserPublicKeyRequest struct {
	// The description of the public key. The description can be up to 500 characters in length.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the bastion host on which you want to create a public key for the user.
	//
	// > You can call the [listinstances](~~204522~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The public key. Encode the value by using the Base64 algorithm.
	PublicKey *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
	// The name of the public key.
	PublicKeyName *string `json:"PublicKeyName,omitempty" xml:"PublicKeyName,omitempty"`
	// Specifies the region ID of the bastion host on which you want to create a public key for the user.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies the ID of the user for whom you want to create a public key.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateUserPublicKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserPublicKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateUserPublicKeyRequest) SetComment(v string) *CreateUserPublicKeyRequest {
	s.Comment = &v
	return s
}

func (s *CreateUserPublicKeyRequest) SetInstanceId(v string) *CreateUserPublicKeyRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateUserPublicKeyRequest) SetPublicKey(v string) *CreateUserPublicKeyRequest {
	s.PublicKey = &v
	return s
}

func (s *CreateUserPublicKeyRequest) SetPublicKeyName(v string) *CreateUserPublicKeyRequest {
	s.PublicKeyName = &v
	return s
}

func (s *CreateUserPublicKeyRequest) SetRegionId(v string) *CreateUserPublicKeyRequest {
	s.RegionId = &v
	return s
}

func (s *CreateUserPublicKeyRequest) SetUserId(v string) *CreateUserPublicKeyRequest {
	s.UserId = &v
	return s
}

type CreateUserPublicKeyResponseBody struct {
	// The ID of the public key.
	PublicKeyId *string `json:"PublicKeyId,omitempty" xml:"PublicKeyId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateUserPublicKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserPublicKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserPublicKeyResponseBody) SetPublicKeyId(v string) *CreateUserPublicKeyResponseBody {
	s.PublicKeyId = &v
	return s
}

func (s *CreateUserPublicKeyResponseBody) SetRequestId(v string) *CreateUserPublicKeyResponseBody {
	s.RequestId = &v
	return s
}

type CreateUserPublicKeyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserPublicKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUserPublicKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserPublicKeyResponse) GoString() string {
	return s.String()
}

func (s *CreateUserPublicKeyResponse) SetHeaders(v map[string]*string) *CreateUserPublicKeyResponse {
	s.Headers = v
	return s
}

func (s *CreateUserPublicKeyResponse) SetStatusCode(v int32) *CreateUserPublicKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserPublicKeyResponse) SetBody(v *CreateUserPublicKeyResponseBody) *CreateUserPublicKeyResponse {
	s.Body = v
	return s
}

type DeleteHostRequest struct {
	// The ID of the host that you want to delete.
	//
	// > You can call the [ListHosts](~~200665~~) operation to query the ID of the host.
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The ID of the bastion host on which you want to delete the host.
	//
	// > You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host on which you want to delete the host.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The ID of the request.
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHostResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the host account that you want to remove.
	//
	// >  You can call the [ListHostAccounts](~~204372~~) operation to query the ID of the host account.
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	// The ID of the bastion host from which you want to remove the host account.
	//
	// >  You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host from which you want to remove the host account.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The ID of the request, which is used to locate and troubleshoot issues.
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHostAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the host group that you want to delete.
	//
	// > You can call the [ListHostGroups](~~201307~~) operation to query the ID of the host group.
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// The ID of the bastion host from which you want to delete the host group.
	//
	// > You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host from which you want to delete the host group.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The ID of the request, which is used to locate and troubleshoot issues.
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHostGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the shared key.
	//
	// > You must specify this parameter.
	HostShareKeyId *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	// The ID of the bastion host. You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host. For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The ID of the request, which is used to locate and troubleshoot issues.
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHostShareKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the bastion host to which the user to be deleted belongs.
	//
	// > You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host to which the user to be deleted belongs.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user to be deleted.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// The ID of the request, which is used to locate and troubleshoot issues.
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the bastion host on which you want to delete the user group.
	//
	// > You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host on which you want to delete the user group.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user group that you want to delete.
	//
	// > You can call the [ListUserGroups](~~204509~~) operation to query the ID of the user group.
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
	// The ID of the request.
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DeleteUserPublicKeyRequest struct {
	// The region ID of the bastion host on which you want to delete the public key from the user.
	//
	// > You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the public key.
	PublicKeyId *string `json:"PublicKeyId,omitempty" xml:"PublicKeyId,omitempty"`
	// The region ID of the bastion host. For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteUserPublicKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserPublicKeyRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserPublicKeyRequest) SetInstanceId(v string) *DeleteUserPublicKeyRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteUserPublicKeyRequest) SetPublicKeyId(v string) *DeleteUserPublicKeyRequest {
	s.PublicKeyId = &v
	return s
}

func (s *DeleteUserPublicKeyRequest) SetRegionId(v string) *DeleteUserPublicKeyRequest {
	s.RegionId = &v
	return s
}

type DeleteUserPublicKeyResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserPublicKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserPublicKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserPublicKeyResponseBody) SetRequestId(v string) *DeleteUserPublicKeyResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUserPublicKeyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserPublicKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserPublicKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserPublicKeyResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserPublicKeyResponse) SetHeaders(v map[string]*string) *DeleteUserPublicKeyResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserPublicKeyResponse) SetStatusCode(v int32) *DeleteUserPublicKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserPublicKeyResponse) SetBody(v *DeleteUserPublicKeyResponseBody) *DeleteUserPublicKeyResponse {
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
	// The attribute information about the bastion host.
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
	AuthorizedSecurityGroups []*string `json:"AuthorizedSecurityGroups,omitempty" xml:"AuthorizedSecurityGroups,omitempty" type:"Repeated"`
	// The total bandwidth of the bastion host.
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The extra bandwidth plan of the bastion host.
	BandwidthPackage     *string                                                        `json:"BandwidthPackage,omitempty" xml:"BandwidthPackage,omitempty"`
	DbOperationModule    *string                                                        `json:"DbOperationModule,omitempty" xml:"DbOperationModule,omitempty"`
	Description          *string                                                        `json:"Description,omitempty" xml:"Description,omitempty"`
	EniInstanceId        *string                                                        `json:"EniInstanceId,omitempty" xml:"EniInstanceId,omitempty"`
	ExpireTime           *int64                                                         `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	InstanceId           *string                                                        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceStatus       *string                                                        `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	InternetEndpoint     *string                                                        `json:"InternetEndpoint,omitempty" xml:"InternetEndpoint,omitempty"`
	IntranetEndpoint     *string                                                        `json:"IntranetEndpoint,omitempty" xml:"IntranetEndpoint,omitempty"`
	LicenseCode          *string                                                        `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
	ModifyPasswordModule *string                                                        `json:"ModifyPasswordModule,omitempty" xml:"ModifyPasswordModule,omitempty"`
	NetworkProxyModule   *string                                                        `json:"NetworkProxyModule,omitempty" xml:"NetworkProxyModule,omitempty"`
	Ports                []*DescribeInstanceAttributeResponseBodyInstanceAttributePorts `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	PrivateExportIps     []*string                                                      `json:"PrivateExportIps,omitempty" xml:"PrivateExportIps,omitempty" type:"Repeated"`
	PrivateWhiteList     []*string                                                      `json:"PrivateWhiteList,omitempty" xml:"PrivateWhiteList,omitempty" type:"Repeated"`
	PublicExportIps      []*string                                                      `json:"PublicExportIps,omitempty" xml:"PublicExportIps,omitempty" type:"Repeated"`
	PublicIps            []*string                                                      `json:"PublicIps,omitempty" xml:"PublicIps,omitempty" type:"Repeated"`
	PublicNetworkAccess  *bool                                                          `json:"PublicNetworkAccess,omitempty" xml:"PublicNetworkAccess,omitempty"`
	PublicWhiteList      []*string                                                      `json:"PublicWhiteList,omitempty" xml:"PublicWhiteList,omitempty" type:"Repeated"`
	RegionId             *string                                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string                                                        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityGroupIds     []*string                                                      `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	StartTime            *int64                                                         `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Storage              *int64                                                         `json:"Storage,omitempty" xml:"Storage,omitempty"`
	VpcId                *string                                                        `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchId            *string                                                        `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	WebTerminalModule    *string                                                        `json:"WebTerminalModule,omitempty" xml:"WebTerminalModule,omitempty"`
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

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetBandwidth(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.Bandwidth = &v
	return s
}

func (s *DescribeInstanceAttributeResponseBodyInstanceAttribute) SetBandwidthPackage(v string) *DescribeInstanceAttributeResponseBodyInstanceAttribute {
	s.BandwidthPackage = &v
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
	// The custom port.
	//
	// >  You can change only the SSH and RDP ports. If O\&M ports are not specified, the value of the StandardPort parameter is returned.
	CustomPort *int32 `json:"CustomPort,omitempty" xml:"CustomPort,omitempty"`
	// The standard port of the bastion host. Valid values:
	//
	// *   **SSH**: 60022
	// *   **RDP**: 63389
	// *   **HTTPS**: 443
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// An array that consists of the IDs of the bastion hosts.
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	// The status of the bastion host. Valid values:
	//
	// *   **PENDING**: The bastion host is not initialized.
	// *   **CREATING**: The bastion host is being created.
	// *   **RUNNING**: The bastion host is running.
	// *   **EXPIRED**: The bastion host expired.
	// *   **CREATE_FAILED**: The bastion host fails to be created.
	// *   **UPGRADING**: The configurations of the bastion host are being changed.
	// *   **UPGRADE_FAILED**: The configurations of the bastion host fail to be changed.
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The number of the page to return. Default value: **1**.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the bastion host.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the bastion host belongs.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// An array consisting of the tags that are added to the bastion hosts.
	Tag []*DescribeInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	// The key of the tag.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
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
	// An array that consists of the queried bastion hosts.
	Instances []*DescribeInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of bastion hosts that are queried.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// The description of the bastion host.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The timestamp when the bastion host expires. Unit: milliseconds.
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The image version of the bastion host.
	ImageVersion *string `json:"ImageVersion,omitempty" xml:"ImageVersion,omitempty"`
	// The ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of the bastion host. Valid values:
	//
	// *   **PENDING**: The bastion host is not initialized.
	// *   **CREATING**: The bastion host is being created.
	// *   **RUNNING**: The bastion host is running.
	// *   **EXPIRED**: The bastion host expired.
	// *   **CREATE_FAILED**: The bastion host fails to be created.
	// *   **UPGRADING**: The configurations of the bastion host are being changed.
	// *   **UPGRADE_FAILED**: The configurations of the bastion host fail to be changed.
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The public O\&M address of the bastion host.
	InternetEndpoint *string `json:"InternetEndpoint,omitempty" xml:"InternetEndpoint,omitempty"`
	// The private O\&M address of the bastion host.
	IntranetEndpoint *string `json:"IntranetEndpoint,omitempty" xml:"IntranetEndpoint,omitempty"`
	// Indicates whether the bastion host runs an earlier version. Valid values:
	//
	// *   **true**: indicates that the bastion host runs V2 or V3.1.
	// *   **false**:indicates that the bastion host runs V3.2.
	Legacy *bool `json:"Legacy,omitempty" xml:"Legacy,omitempty"`
	// The license code of the bastion host.
	LicenseCode *string `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
	// The edition of the bastion host. Valid values:
	//
	// *   **cloudbastion**: Basic
	// *   **cloudbastion_ha**: Enterprise
	PlanCode *string `json:"PlanCode,omitempty" xml:"PlanCode,omitempty"`
	// Indicates whether the bastion host can be accessed from the Internet. Valid values:
	//
	// *   **true**: The bastion host can be accessed from the Internet.
	// *   **false**: The bastion host cannot be accessed from the Internet.
	PublicNetworkAccess *bool `json:"PublicNetworkAccess,omitempty" xml:"PublicNetworkAccess,omitempty"`
	// The region ID of the bastion host.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the bastion host belongs.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The timestamp when the bastion host is purchased or renewed. Unit: milliseconds.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the virtual private cloud (VPC) to which the bastion host belongs.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the vSwitch to which the bastion host belongs.
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the region.
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The ID of request.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// DescribeRegions
	Regions []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// Queries available regions where you can create bastion hosts.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The IDs of the host accounts.
	HostAccountIds *string `json:"HostAccountIds,omitempty" xml:"HostAccountIds,omitempty"`
	// The ID of the shared key.
	HostShareKeyId *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	// The ID of the bastion host. You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host. For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the call.
	Results []*DetachHostAccountsFromHostShareKeyResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
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
	// The error code returned. If **OK** is returned, the disassociation was successful. If a different error code is returned, the disassociation failed.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the host account.
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	// The ID of the shared key.
	HostShareKeyId *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	// The error message returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachHostAccountsFromHostShareKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The IDs of the host and host account on which you want to revoke permissions from the user. You can specify up to 10 host IDs and up to 10 host account IDs for each host. You can specify only host IDs. In this case, the permissions on both the specified hosts and all host accounts of the hosts are revoked from the user. For more information about this parameter, see the "Description of the Hosts parameter" section of this topic.
	//
	// >  You can call the [ListHosts](~~200665~~) operation to query the ID of the host and the [ListHostAccounts](~~204372~~) operation to query the ID of the host account.
	Hosts *string `json:"Hosts,omitempty" xml:"Hosts,omitempty"`
	// The ID of the bastion host in which you want to revoke permissions on the specified hosts and host accounts from the user.
	//
	// >  You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host in which you want to revoke permissions on the specified hosts and host accounts from the user.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user from which you want to revoke permissions on the specified hosts and host accounts.
	//
	// >  You can call the [ListUsers](~~204522~~) operation to query the ID of the user.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the call.
	Results []*DetachHostAccountsFromUserResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
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
	// The return code that indicates whether the call was successful. Valid values:
	//
	// *   **OK**: The call was successful.
	// *   **UNEXPECTED**: An unknown error occurred.
	// *   **INVALID_ARGUMENT**: A request parameter is invalid.
	// *   **OBJECT_NOT_FOUND**: The specified object on which you want to perform the operation does not exist.
	// *   **OBJECT_AlREADY_EXISTS**: The specified object on which you want to perform the operation already exists.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of revoking permissions on the specified host accounts from the user.
	HostAccounts []*DetachHostAccountsFromUserResponseBodyResultsHostAccounts `json:"HostAccounts,omitempty" xml:"HostAccounts,omitempty" type:"Repeated"`
	// The ID of the host.
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// This parameter is deprecated.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the user.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// The return code that indicates whether permissions on the specified host account were revoked from the user. Valid values:
	//
	// *   **OK**: The call was successful.
	// *   **UNEXPECTED**: An unknown error occurred.
	// *   **INVALID_ARGUMENT**: A request parameter is invalid.
	// *   **OBJECT_NOT_FOUND**: The specified object on which you want to perform the operation does not exist.
	// *   **OBJECT_AlREADY_EXISTS**: The specified object on which you want to perform the operation already exists.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the host account.
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	// This parameter is deprecated.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachHostAccountsFromUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The IDs of the host and host account on which you want to revoke permissions from the user group.
	//
	// You can specify up to 10 host IDs and up to 10 host account IDs for each host. You can specify only host IDs. In this case, the permissions on both the specified hosts and all host accounts of the hosts are revoked from the user group. For more information about this parameter, see the "Description of the Hosts parameter" section of this topic.
	//
	// >  You can call the [ListHosts](~~200665~~) operation to query the ID of the host and the [ListHostAccounts](~~204372~~) operation to query the ID of the host account.
	Hosts *string `json:"Hosts,omitempty" xml:"Hosts,omitempty"`
	// The ID of the bastion host in which you want to revoke permissions on the specified hosts and host accounts from the user group.
	//
	// >  You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host in which you want to revoke permissions on the specified hosts and host accounts from the user group.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user group from which you want to revoke permissions on the specified hosts and host accounts.
	//
	// >  You can call the [ListUserGroups](~~204509~~) operation to query the ID of the user group.
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
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the call.
	Results []*DetachHostAccountsFromUserGroupResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
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
	// The return code that indicates whether the call was successful. Valid values:
	//
	// *   **OK**: The call was successful.
	// *   **UNEXPECTED**: An unknown error occurred.
	// *   **INVALID_ARGUMENT**: A request parameter is invalid.
	// *   **OBJECT_NOT_FOUND**: The specified object on which you want to perform the operation does not exist.
	// *   **OBJECT_AlREADY_EXISTS**: The specified object on which you want to perform the operation already exists.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of revoking permissions on the specified host accounts from the user group.
	HostAccounts []*DetachHostAccountsFromUserGroupResponseBodyResultsHostAccounts `json:"HostAccounts,omitempty" xml:"HostAccounts,omitempty" type:"Repeated"`
	// The ID of the host.
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// This parameter is deprecated.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the group.
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
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
	// The return code that indicates whether permissions on the specified host account were revoked from the user group. Valid values:
	//
	// *   **OK**: The call was successful.
	// *   **UNEXPECTED**: An unknown error occurred.
	// *   **INVALID_ARGUMENT**: A request parameter is invalid.
	// *   **OBJECT_NOT_FOUND**: The specified object on which you want to perform the operation does not exist.
	// *   **OBJECT_AlREADY_EXISTS**: The specified object on which you want to perform the operation already exists.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the host account.
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	// This parameter is deprecated.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachHostAccountsFromUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the host group and the name of the host account on which you want to revoke permissions from the user. You can specify up to 10 host group IDs and up to 10 host account names for each host group. You can specify only host group IDs. In this case, the permissions on the specified host groups and all host accounts in the host groups are revoked from the user. For more information about this parameter, see the "Description of the HostGroups parameter" section of this topic.
	//
	// > You can call the [ListHostGroups](~~201307~~) operation to query the ID of the host group and the [ListHostAccounts](~~204372~~) operation to query the name of the host account.
	HostGroups *string `json:"HostGroups,omitempty" xml:"HostGroups,omitempty"`
	// The ID of the bastion host for which you want to revoke permissions on the specified host groups and host accounts from the user.
	//
	// > You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host for which you want to revoke permissions on the specified host groups and host accounts from the user.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user from which you want to revoke permissions on the specified host groups and host accounts.
	//
	// > You can call the [ListUsers](~~204522~~) operation to query the ID of the user.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the call.
	Results []*DetachHostGroupAccountsFromUserResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
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
	// The return code that indicates whether the call was successful. Valid values:
	//
	// *   **OK**: The call was successful.
	// *   **UNEXPECTED**: An unknown error occurred.
	// *   **INVALID_ARGUMENT**: A request parameter is invalid.
	// *   **OBJECT_NOT_FOUND**: The specified object on which you want to perform the operation does not exist.
	// *   **OBJECT_AlREADY_EXISTS**: The specified object on which you want to perform the operation already exists.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of revoking permissions on the specified host accounts from the user.
	HostAccountNames []*DetachHostGroupAccountsFromUserResponseBodyResultsHostAccountNames `json:"HostAccountNames,omitempty" xml:"HostAccountNames,omitempty" type:"Repeated"`
	// The ID of the host group.
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// This parameter is deprecated.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the user.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// The return code that indicates whether permissions on the specified host account were revoked from the user. Valid values:
	//
	// *   **OK**: The call was successful.
	// *   **UNEXPECTED**: An unknown error occurred.
	// *   **INVALID_ARGUMENT**: A request parameter is invalid.
	// *   **OBJECT_NOT_FOUND**: The specified object on which you want to perform the operation does not exist.
	// *   **OBJECT_AlREADY_EXISTS**: The specified object on which you want to perform the operation already exists.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The name of the host account.
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	// This parameter is deprecated.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachHostGroupAccountsFromUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the host group and the name of host account on which you want to revoke permissions from the user group. You can specify up to 10 host group IDs and up to 10 host account names for each host group. You can specify only host group IDs. In this case, the permissions on the specified host groups and all host accounts in the host groups are revoked from the user group. For more information about this parameter, see the "Description of the HostGroups parameter" section of this topic.
	//
	// >  You can call the [ListHostGroups](~~201307~~) operation to query the ID of the host group and the [ListHostAccounts](~~204372~~) operation to query the name of the host account.
	HostGroups *string `json:"HostGroups,omitempty" xml:"HostGroups,omitempty"`
	// The ID of the bastion host for which you want to revoke permissions on the specified host groups and host accounts from the user group.
	//
	// >  You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host for which you want to revoke permissions on the specified host groups and host accounts from the user group.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user group from which you want to revoke permissions on the specified host groups and host accounts.
	//
	// >  You can call the [ListUserGroups](~~204509~~) operation to query the ID of the user group.
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
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the call.
	Results []*DetachHostGroupAccountsFromUserGroupResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
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
	// The return code that indicates whether the call was successful. Valid values:
	//
	// *   **OK**: The call was successful.
	// *   **UNEXPECTED**: An unknown error occurred.
	// *   **INVALID_ARGUMENT**: A request parameter is invalid.
	// *   **OBJECT_NOT_FOUND**: The specified object on which you want to perform the operation does not exist.
	// *   **OBJECT_AlREADY_EXISTS**: The specified object on which you want to perform the operation already exists.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of revoking permissions on the specified host accounts from the user group.
	HostAccountNames []*DetachHostGroupAccountsFromUserGroupResponseBodyResultsHostAccountNames `json:"HostAccountNames,omitempty" xml:"HostAccountNames,omitempty" type:"Repeated"`
	// The ID of the host group.
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// This parameter is deprecated.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the group.
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
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
	// The return code that indicates whether permissions on the specified host account were revoked from the specified user group. Valid values:
	//
	// *   **OK**: The call was successful.
	// *   **UNEXPECTED**: An unknown error occurred.
	// *   **INVALID_ARGUMENT**: A request parameter is invalid.
	// *   **OBJECT_NOT_FOUND**: The specified object on which you want to perform the operation does not exist.
	// *   **OBJECT_AlREADY_EXISTS**: The specified object on which you want to perform the operation already exists.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The name of the host account.
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	// This parameter is deprecated.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachHostGroupAccountsFromUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the bastion host whose Internet access you want to disable.
	//
	// > You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The ID of the bastion host whose Internet access is disabled.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableInstancePublicAccessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the bastion host.
	//
	// >  You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The ID of the bastion host whose Internet access is enabled.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableInstancePublicAccessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the host that you want to query. You can specify only one host ID.
	//
	// >  You can call the [ListHosts](~~200665~~) operation to query the ID of the host.
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The ID of the bastion host in which you want to query the host.
	//
	// >  You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host in which you want to query the host.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The information about the host that was queried.
	Host *GetHostResponseBodyHost `json:"Host,omitempty" xml:"Host,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The address type of the host. Valid values:
	//
	// *   **Public**: a public address
	// *   **Private**: a private address
	ActiveAddressType *string `json:"ActiveAddressType,omitempty" xml:"ActiveAddressType,omitempty"`
	// The description of the host.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the host.
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The hostname.
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The internal endpoint of the host. The value is a domain name or an IP address.
	HostPrivateAddress *string `json:"HostPrivateAddress,omitempty" xml:"HostPrivateAddress,omitempty"`
	// The public address of the host. The value is a domain name or an IP address.
	HostPublicAddress *string `json:"HostPublicAddress,omitempty" xml:"HostPublicAddress,omitempty"`
	// The ID of the new network domain to which the host belongs.
	NetworkDomainId *string `json:"NetworkDomainId,omitempty" xml:"NetworkDomainId,omitempty"`
	// The operating system of the host. Valid values:
	//
	// *   **Linux**
	// *   **Windows**
	OSType *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	// The protocol information about the host.
	Protocols []*GetHostResponseBodyHostProtocols `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
	// The source of the host. Valid values:
	//
	// *   **Local**: a host in a data center
	// *   **Ecs**: an Elastic Compute Service (ECS) instance
	// *   **Rds**: a host in an ApsaraDB MyBase dedicated cluster
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The ID of the ECS instance or the host in an ApsaraDB MyBase dedicated cluster.
	//
	// >  If **Local** is returned for the **Source** parameter, no value is returned for this parameter.
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
	// The status of the host. Valid values:
	//
	// *   **Normal**: normal
	// *   **Release**: released
	SourceInstanceState *string `json:"SourceInstanceState,omitempty" xml:"SourceInstanceState,omitempty"`
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

func (s *GetHostResponseBodyHost) SetNetworkDomainId(v string) *GetHostResponseBodyHost {
	s.NetworkDomainId = &v
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
	// The fingerprint of the host. This parameter uniquely identifies a host.
	HostFingerPrint *string `json:"HostFingerPrint,omitempty" xml:"HostFingerPrint,omitempty"`
	// The service port of the host.
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol that is used to connect to the host. Valid values:
	//
	// *   **SSH**
	// *   **RDP**
	ProtocolName *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
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
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHostResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the host account that you want to query.
	//
	// > You can call the [ListHostAccounts](~~204372~~) operation to query the ID of the host account.
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	// The ID of the bastion host in which you want to query the details of the host account.
	//
	// > You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host in which you want to query the details of the host account.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The details of the host account that was queried.
	HostAccount *GetHostAccountResponseBodyHostAccount `json:"HostAccount,omitempty" xml:"HostAccount,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// Indicates whether a password is configured for the host account. Valid values:
	//
	// *   **true**: yes
	// *   **false**: no
	HasPassword *bool `json:"HasPassword,omitempty" xml:"HasPassword,omitempty"`
	// The ID of the host account.
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	// The name of the host account.
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	// The ID of the host to which the host account belongs.
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The ID of the shared key.
	HostShareKeyId *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	// The name of the shared key.
	HostShareKeyName *string `json:"HostShareKeyName,omitempty" xml:"HostShareKeyName,omitempty"`
	// The fingerprint of the private key.
	PrivateKeyFingerprint *string `json:"PrivateKeyFingerprint,omitempty" xml:"PrivateKeyFingerprint,omitempty"`
	// The protocol that is used by the host. Valid values:
	//
	// *   **SSH**
	// *   **RDP**
	ProtocolName *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHostAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The region ID of the Bastionhost instance where you want to query the host group.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// MyHostGroup
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the Bastionhost instance where you want to query the host group.
	//
	// >  You can call the [DescribeInstances](~~153281~~) operation to query the ID of the Bastionhost instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The ID of the host group that you want to query.
	//
	// >  You can call the [ListHostGroups](~~201307~~) operation to query the ID of the host group.
	HostGroup *GetHostGroupResponseBodyHostGroup `json:"HostGroup,omitempty" xml:"HostGroup,omitempty" type:"Struct"`
	// my host group.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The details of the host group returned.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The description of the host group.
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// The ID of the host group.
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHostGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The time when the information about the shared key was last modified.
	HostShareKeyId *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	// The ID of the shared key whose details you want to query.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the shared key.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The operation that you want to perform. Set the value to **GetHostShareKey**.
	HostShareKey *GetHostShareKeyResponseBodyHostShareKey `json:"HostShareKey,omitempty" xml:"HostShareKey,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The fingerprint of the private key.
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHostShareKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The field that is used to indicate the email address of a user on the AD server.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether passwords are required. Valid values:
	//
	// *   **true**: required
	// *   **false**: not required
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The operation that you want to perform. Set the value to **GetInstanceADAuthServer**.
	AD *GetInstanceADAuthServerResponseBodyAD `json:"AD,omitempty" xml:"AD,omitempty" type:"Struct"`
	// Indicates whether SSL is supported. Valid values:
	//
	// *   **true**: supported
	// *   **false**: not supported
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The port that is used to access the AD server.
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The ID of the bastion host to query.
	//
	// You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	BaseDN *string `json:"BaseDN,omitempty" xml:"BaseDN,omitempty"`
	// The settings of AD authentication.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The address of the secondary AD server.
	EmailMapping *string `json:"EmailMapping,omitempty" xml:"EmailMapping,omitempty"`
	// The field that is used to indicate the mobile phone number of a user on the AD server.
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The address of the AD server.
	HasPassword *bool `json:"HasPassword,omitempty" xml:"HasPassword,omitempty"`
	// The Base DN of the AD server.
	IsSSL *bool `json:"IsSSL,omitempty" xml:"IsSSL,omitempty"`
	// The field that is used to indicate the name of a user on the AD server.
	MobileMapping *string `json:"MobileMapping,omitempty" xml:"MobileMapping,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	NameMapping *string `json:"NameMapping,omitempty" xml:"NameMapping,omitempty"`
	// Queries the settings of Active Directory (AD) authentication on a bastion host.
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceADAuthServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// Indicates whether passwords are required. Valid values:
	//
	// *   **true**: required
	// *   **false**: not required
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The operation that you want to perform. Set the value to **GetInstanceLDAPAuthServer**.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// Indicates whether SSL is supported. Valid values:
	//
	// *   **true**: supported
	// *   **false**: not supported
	LDAP *GetInstanceLDAPAuthServerResponseBodyLDAP `json:"LDAP,omitempty" xml:"LDAP,omitempty" type:"Struct"`
	// The settings of LDAP authentication.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The ID of the bastion host.
	//
	// >  You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The field that is used to indicate the logon name of a user on the LDAP server.
	BaseDN *string `json:"BaseDN,omitempty" xml:"BaseDN,omitempty"`
	// The address of the secondary LDAP server.
	EmailMapping *string `json:"EmailMapping,omitempty" xml:"EmailMapping,omitempty"`
	// The Base distinguished name (DN).
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	HasPassword *string `json:"HasPassword,omitempty" xml:"HasPassword,omitempty"`
	// The condition that is used to filter users.
	IsSSL *bool `json:"IsSSL,omitempty" xml:"IsSSL,omitempty"`
	// The port that is used to access the LDAP server.
	LoginNameMapping *string `json:"LoginNameMapping,omitempty" xml:"LoginNameMapping,omitempty"`
	// The field that is used to indicate the email address of a user on the LDAP server.
	MobileMapping *string `json:"MobileMapping,omitempty" xml:"MobileMapping,omitempty"`
	// The field that is used to indicate the mobile phone number of a user on the LDAP server.
	NameMapping   *string `json:"NameMapping,omitempty" xml:"NameMapping,omitempty"`
	Port          *int64  `json:"Port,omitempty" xml:"Port,omitempty"`
	Server        *string `json:"Server,omitempty" xml:"Server,omitempty"`
	StandbyServer *string `json:"StandbyServer,omitempty" xml:"StandbyServer,omitempty"`
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceLDAPAuthServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the request, which is used to locate and troubleshoot issues.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The operation that you want to perform. Set the value to **GetInstanceTwoFactor**.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The settings of two-factor authentication.
	Config *GetInstanceTwoFactorResponseBodyConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// The duration within which two-factor authentication is not required after a local user passes two-factor authentication. Valid values: `0 to 168`. Unit: hours.
	//
	// >  If 0 is returned, a local user must pass two-factor authentication every time the local user logs on to the bastion host.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// Queries the settings of two-factor authentication on a bastion host.
	EnableTwoFactor   *bool     `json:"EnableTwoFactor,omitempty" xml:"EnableTwoFactor,omitempty"`
	SkipTwoFactorTime *int64    `json:"SkipTwoFactorTime,omitempty" xml:"SkipTwoFactorTime,omitempty"`
	TwoFactorMethods  []*string `json:"TwoFactorMethods,omitempty" xml:"TwoFactorMethods,omitempty" type:"Repeated"`
}

func (s GetInstanceTwoFactorResponseBodyConfig) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceTwoFactorResponseBodyConfig) GoString() string {
	return s.String()
}

func (s *GetInstanceTwoFactorResponseBodyConfig) SetEnableTwoFactor(v bool) *GetInstanceTwoFactorResponseBodyConfig {
	s.EnableTwoFactor = &v
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

type GetInstanceTwoFactorResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceTwoFactorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the bastion host on which you want to query the user.
	//
	// > You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host on which you want to query the user.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user.
	//
	// > You can call the [ListUsers](~~204522~~) operation to query the ID of the user.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the user that was queried.
	User *GetUserResponseBodyUser `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
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
	// The description of the user.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The display name of the user.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The end of the validity period of the user. The value is a UNIX timestamp. Unit: seconds.
	EffectiveEndTime *int64 `json:"EffectiveEndTime,omitempty" xml:"EffectiveEndTime,omitempty"`
	// The beginning of the validity period of the user. The value is a UNIX timestamp. Unit: seconds.
	EffectiveStartTime *int64 `json:"EffectiveStartTime,omitempty" xml:"EffectiveStartTime,omitempty"`
	// The email address of the user.
	Email          *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Language       *string `json:"Language,omitempty" xml:"Language,omitempty"`
	LanguageStatus *string `json:"LanguageStatus,omitempty" xml:"LanguageStatus,omitempty"`
	// The mobile phone number of the user.
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The location in which the mobile number of the user is registered. Valid values:
	//
	// *   **CN**: the Chinese mainland, whose country calling code is +86
	// *   **HK**: Hong Kong (China), whose country calling code is +852
	// *   **MO**: Macao (China), whose country calling code is +853
	// *   **TW**: Taiwan (China), whose country calling code is +886
	// *   **RU**: Russia, whose country calling code is +7
	// *   **SG**: Singapore, whose country calling code is +65
	// *   **MY**: Malaysia, whose country calling code is +60
	// *   **ID**: Indonesia, whose country calling code is +62
	// *   **DE**: Germany, whose country calling code is +49
	// *   **AU**: Australia, whose country calling code is +61
	// *   **US**: US, whose country calling code is +1
	// *   **AE**: United Arab Emirates, whose country calling code is +971
	// *   **JP:** Japan, whose country calling code is +81
	// *   **GB**: UK, whose country calling code is +44
	// *   **IN**: India, whose country calling code is +91
	// *   **KR**: Republic of Korea, whose country calling code is +82
	// *   **PH**: Philippines, whose country calling code is +63
	// *   **CH**: Switzerland, whose country calling code is +41
	// *   **SE**: Sweden, whose country calling code is +46
	MobileCountryCode *string `json:"MobileCountryCode,omitempty" xml:"MobileCountryCode,omitempty"`
	// Specifies whether password reset is required upon the next logon. Valid values:
	//
	// *   **true**: yes
	// *   **false**: no
	NeedResetPassword *bool `json:"NeedResetPassword,omitempty" xml:"NeedResetPassword,omitempty"`
	// The source of the user. Valid values:
	//
	// *   **Local**: a local user
	// *   **Ram**: a RAM user
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The unique ID of the user.
	//
	// > This parameter uniquely identifies a RAM user of the bastion host. A value is returned for this parameter if the **Source** parameter is set to **Ram**. No value is returned for this parameter if the **Source** parameter is set to **Local**.
	SourceUserId *string `json:"SourceUserId,omitempty" xml:"SourceUserId,omitempty"`
	// An array that consists of the details of the two-factor authentication method.
	TwoFactorMethods []*string `json:"TwoFactorMethods,omitempty" xml:"TwoFactorMethods,omitempty" type:"Repeated"`
	// The two-factor authentication status of the user. Valid values:
	//
	// *   **Global**: The global settings are used.
	// *   **Disable**: The two-factor authentication is disabled.
	// *   **Enable**: The two-factor authentication is enabled and the user-specific setting is used.
	TwoFactorStatus *string `json:"TwoFactorStatus,omitempty" xml:"TwoFactorStatus,omitempty"`
	// The ID of the user.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The logon name of the user.
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// An array that consists of the details of the user status.
	UserState []*string `json:"UserState,omitempty" xml:"UserState,omitempty" type:"Repeated"`
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

func (s *GetUserResponseBodyUser) SetEffectiveEndTime(v int64) *GetUserResponseBodyUser {
	s.EffectiveEndTime = &v
	return s
}

func (s *GetUserResponseBodyUser) SetEffectiveStartTime(v int64) *GetUserResponseBodyUser {
	s.EffectiveStartTime = &v
	return s
}

func (s *GetUserResponseBodyUser) SetEmail(v string) *GetUserResponseBodyUser {
	s.Email = &v
	return s
}

func (s *GetUserResponseBodyUser) SetLanguage(v string) *GetUserResponseBodyUser {
	s.Language = &v
	return s
}

func (s *GetUserResponseBodyUser) SetLanguageStatus(v string) *GetUserResponseBodyUser {
	s.LanguageStatus = &v
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

func (s *GetUserResponseBodyUser) SetNeedResetPassword(v bool) *GetUserResponseBodyUser {
	s.NeedResetPassword = &v
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

func (s *GetUserResponseBodyUser) SetTwoFactorMethods(v []*string) *GetUserResponseBodyUser {
	s.TwoFactorMethods = v
	return s
}

func (s *GetUserResponseBodyUser) SetTwoFactorStatus(v string) *GetUserResponseBodyUser {
	s.TwoFactorStatus = &v
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
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the request.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the user group.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// All Bastionhost API requests must include common request parameters. For more information about common request parameters, see [Common parameters](~~315526~~).
	//
	// For more information about sample requests, see the "Examples" section of this topic.
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
	// Queries the details of a specified user group in a specified Bastionhost instance.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// GetUserGroup
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
	// GetUserGroup
	Comment     *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// WB662865
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ListApproveCommandsRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListApproveCommandsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListApproveCommandsRequest) GoString() string {
	return s.String()
}

func (s *ListApproveCommandsRequest) SetInstanceId(v string) *ListApproveCommandsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListApproveCommandsRequest) SetPageNumber(v string) *ListApproveCommandsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListApproveCommandsRequest) SetPageSize(v string) *ListApproveCommandsRequest {
	s.PageSize = &v
	return s
}

func (s *ListApproveCommandsRequest) SetRegionId(v string) *ListApproveCommandsRequest {
	s.RegionId = &v
	return s
}

type ListApproveCommandsResponseBody struct {
	ApproveCommands []*ListApproveCommandsResponseBodyApproveCommands `json:"ApproveCommands,omitempty" xml:"ApproveCommands,omitempty" type:"Repeated"`
	RequestId       *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount      *int64                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApproveCommandsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListApproveCommandsResponseBody) GoString() string {
	return s.String()
}

func (s *ListApproveCommandsResponseBody) SetApproveCommands(v []*ListApproveCommandsResponseBodyApproveCommands) *ListApproveCommandsResponseBody {
	s.ApproveCommands = v
	return s
}

func (s *ListApproveCommandsResponseBody) SetRequestId(v string) *ListApproveCommandsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApproveCommandsResponseBody) SetTotalCount(v int64) *ListApproveCommandsResponseBody {
	s.TotalCount = &v
	return s
}

type ListApproveCommandsResponseBodyApproveCommands struct {
	ApproveCommandId *string `json:"ApproveCommandId,omitempty" xml:"ApproveCommandId,omitempty"`
	AssetAccountName *string `json:"AssetAccountName,omitempty" xml:"AssetAccountName,omitempty"`
	AssetIp          *string `json:"AssetIp,omitempty" xml:"AssetIp,omitempty"`
	AssetName        *string `json:"AssetName,omitempty" xml:"AssetName,omitempty"`
	ClientIp         *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	ClientUser       *string `json:"ClientUser,omitempty" xml:"ClientUser,omitempty"`
	Command          *string `json:"Command,omitempty" xml:"Command,omitempty"`
	CreateTime       *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ProtocolName     *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
	SessionId        *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	State            *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListApproveCommandsResponseBodyApproveCommands) String() string {
	return tea.Prettify(s)
}

func (s ListApproveCommandsResponseBodyApproveCommands) GoString() string {
	return s.String()
}

func (s *ListApproveCommandsResponseBodyApproveCommands) SetApproveCommandId(v string) *ListApproveCommandsResponseBodyApproveCommands {
	s.ApproveCommandId = &v
	return s
}

func (s *ListApproveCommandsResponseBodyApproveCommands) SetAssetAccountName(v string) *ListApproveCommandsResponseBodyApproveCommands {
	s.AssetAccountName = &v
	return s
}

func (s *ListApproveCommandsResponseBodyApproveCommands) SetAssetIp(v string) *ListApproveCommandsResponseBodyApproveCommands {
	s.AssetIp = &v
	return s
}

func (s *ListApproveCommandsResponseBodyApproveCommands) SetAssetName(v string) *ListApproveCommandsResponseBodyApproveCommands {
	s.AssetName = &v
	return s
}

func (s *ListApproveCommandsResponseBodyApproveCommands) SetClientIp(v string) *ListApproveCommandsResponseBodyApproveCommands {
	s.ClientIp = &v
	return s
}

func (s *ListApproveCommandsResponseBodyApproveCommands) SetClientUser(v string) *ListApproveCommandsResponseBodyApproveCommands {
	s.ClientUser = &v
	return s
}

func (s *ListApproveCommandsResponseBodyApproveCommands) SetCommand(v string) *ListApproveCommandsResponseBodyApproveCommands {
	s.Command = &v
	return s
}

func (s *ListApproveCommandsResponseBodyApproveCommands) SetCreateTime(v string) *ListApproveCommandsResponseBodyApproveCommands {
	s.CreateTime = &v
	return s
}

func (s *ListApproveCommandsResponseBodyApproveCommands) SetProtocolName(v string) *ListApproveCommandsResponseBodyApproveCommands {
	s.ProtocolName = &v
	return s
}

func (s *ListApproveCommandsResponseBodyApproveCommands) SetSessionId(v string) *ListApproveCommandsResponseBodyApproveCommands {
	s.SessionId = &v
	return s
}

func (s *ListApproveCommandsResponseBodyApproveCommands) SetState(v string) *ListApproveCommandsResponseBodyApproveCommands {
	s.State = &v
	return s
}

type ListApproveCommandsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApproveCommandsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApproveCommandsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListApproveCommandsResponse) GoString() string {
	return s.String()
}

func (s *ListApproveCommandsResponse) SetHeaders(v map[string]*string) *ListApproveCommandsResponse {
	s.Headers = v
	return s
}

func (s *ListApproveCommandsResponse) SetStatusCode(v int32) *ListApproveCommandsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApproveCommandsResponse) SetBody(v *ListApproveCommandsResponseBody) *ListApproveCommandsResponse {
	s.Body = v
	return s
}

type ListHostAccountsRequest struct {
	// Indicates whether a password is configured for the host account.
	//
	// Valid values:
	//
	// *   true: A password is configured for the host account.
	// *   false: No passwords are configured for the host account.
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	// The protocol used by the host whose accounts you want to query.
	//
	// Valid values:
	//
	// *   SSH
	// *   RDP
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The ID of the shared key.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The operation that you want to perform.
	//
	// Set the value to **ListHostAccounts**.
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Maximum value: 100. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// >  We recommend that you do not leave this parameter empty.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the host account that you want to query. The name can be up to 128 characters in length. Only exact match is supported.
	ProtocolName *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
	// The ID of the specified host whose accounts you want to query.
	//
	// >  You can call the [ListHosts](~~200665~~) operation to query the ID of the host.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The ID of the host account.
	HostAccounts []*ListHostAccountsResponseBodyHostAccounts `json:"HostAccounts,omitempty" xml:"HostAccounts,omitempty" type:"Repeated"`
	// An array that consists of the queried host accounts.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the bastion host in which you want to query accounts of the specified host.
	//
	// >  You can call the DescribeInstances operation to query the ID of the bastion host.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// The fingerprint of the private key for the host account.
	HasPassword *bool `json:"HasPassword,omitempty" xml:"HasPassword,omitempty"`
	// The ID of the request.
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	// The name of the shared key.
	HostAccountName  *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	HostId           *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	HostShareKeyId   *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	HostShareKeyName *string `json:"HostShareKeyName,omitempty" xml:"HostShareKeyName,omitempty"`
	// The protocol that is used by the host.
	//
	// Valid values:
	//
	// *   SSH
	// *   RDP
	PrivateKeyFingerprint *string `json:"PrivateKeyFingerprint,omitempty" xml:"PrivateKeyFingerprint,omitempty"`
	// The number of the page to return. Default value: **1**.
	ProtocolName *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHostAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the shared key.
	HostShareKeyId *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	// The ID of the bastion host. You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return. Default value: **1**.
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the bastion host. For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// An array that consists of the host accounts that are associated with the shared key.
	HostAccounts []*ListHostAccountsForHostShareKeyResponseBodyHostAccounts `json:"HostAccounts,omitempty" xml:"HostAccounts,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the host accounts that are associated with the shared key.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// The name of the host account.
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	// The ID of the host.
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The ID of the host account.
	HostsAccountId *string `json:"HostsAccountId,omitempty" xml:"HostsAccountId,omitempty"`
	// The O\&M protocol.
	ProtocolName *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
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
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHostAccountsForHostShareKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The number of the page to return. Default value: **1**.
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	// The ID of the host for which the host accounts were queried.
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The total number of host accounts returned.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the user for which you want to query authorized host accounts.
	//
	// >  You can call the [ListUsers](~~204522~~) operation to query the ID of the user ID.
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The name of the host account that you want to query. Exact match is supported.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the host account.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The region ID of the Bastionhost instance where you want to query the host accounts that the user is authorized to manage on the host.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// The host accounts returned.
	HostAccounts []*ListHostAccountsForUserResponseBodyHostAccounts `json:"HostAccounts,omitempty" xml:"HostAccounts,omitempty" type:"Repeated"`
	// The ID of the Bastionhost instance where you want to query the host accounts that the user is authorized to manage on the host.
	//
	// >  You can call the [DescribeInstances](~~153281~~) operation to query the ID of the Bastionhost instance.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries to return on each page.
	//
	// The value of the PageSize parameter must not exceed 100. Default value: 20. If you leave the PageSize parameter empty, 20 entries are returned on each page.
	//
	// >  We recommend that you do not leave the PageSize parameter empty.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// The protocol that is used by the host account. Valid values:
	//
	// *   **SSH**
	// *   **RDP**
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	// The ID of the host account.
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	// The ID of the request.
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The ID of the host for which you want to query the host accounts that the user is authorized to manage.
	//
	// >  You can call the [ListHosts](~~200665~~) operation to query the ID of the host.
	IsAuthorized *bool `json:"IsAuthorized,omitempty" xml:"IsAuthorized,omitempty"`
	// Indicates whether the user is authorized to manage the host account. Valid values:
	//
	// *   **true**: The user is authorized to manage the host account.
	// *   **false**: The user is not authorized to manage the host account.
	ProtocolName *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHostAccountsForUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The name of the host account that you want to query. Exact match is supported.
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	// The ID of the host to query.
	//
	// > You can call the [ListHosts](~~200665~~) operation to query the ID of the host.
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The ID of the bastion host on which you want to query the host accounts to be managed by the specified user group on the specified host.
	//
	// > You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return. Default value: **1**.
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.\
	// Maximum value: 100. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// > We recommend that you do not leave this parameter empty.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the bastion host on which you want to query the host accounts to be managed by the specified user group on the specified host.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user group for which you want to query authorized host accounts.
	//
	// > You can call the [ListUserGroups](~~204509~~) operation to query the ID of the user group.
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
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
	// An array that consists of the queried host accounts.
	HostAccounts []*ListHostAccountsForUserGroupResponseBodyHostAccounts `json:"HostAccounts,omitempty" xml:"HostAccounts,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of host accounts that were queried.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// The ID of the host account.
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	// The name of the host account.
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	// The ID of the host for which the host accounts were queried.
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// Indicates whether the user group is authorized to manage the host account. Valid values:
	//
	// *   **true**: yes
	// *   **false**: no
	IsAuthorized *bool `json:"IsAuthorized,omitempty" xml:"IsAuthorized,omitempty"`
	// The protocol that is used by the host. Valid values:
	//
	// *   **SSH**
	// *   **RDP**
	ProtocolName *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHostAccountsForUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the host group.
	//
	// > You can call the [ListHostGroups](~~201307~~) operation to query the ID of the host group.
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// The ID of the bastion host to which the user belongs.
	//
	// > You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host to which the user belongs.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user.
	//
	// > You can call the [ListUsers](~~204522~~) operation to query the ID of the user.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// An array that consists of the names of host accounts.
	HostAccountNames []*string `json:"HostAccountNames,omitempty" xml:"HostAccountNames,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHostGroupAccountNamesForUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the host group.
	//
	// > You can call the [ListHostGroups](~~201307~~) operation to query the ID of the host group.
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// The ID of the bastion host on which you want to query the host account names the user group is authorized to manage in a host group.
	//
	// > You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host on which you want to query the host account names the user group is authorized to manage in a host group.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user group.
	//
	// > You can call the [ListUserGroups](~~204509~~) operation to query the ID of the user group.
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
	// The names of host accounts returned.
	HostAccountNames []*string `json:"HostAccountNames,omitempty" xml:"HostAccountNames,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHostGroupAccountNamesForUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The name of the host group that you want to query. Only exact match is supported.
	HostGroupName *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
	// The ID of the bastion host in which you want to query the host group.
	//
	// > You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return. Default value: **1**.
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.\
	// Maximum value: 100. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// > We recommend that you do not leave this parameter empty.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the bastion host in which you want to query the host group.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// An array that consists of the host groups.
	HostGroups []*ListHostGroupsResponseBodyHostGroups `json:"HostGroups,omitempty" xml:"HostGroups,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of host groups returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// The description of the host group.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the host group.
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// The name of the host group.
	HostGroupName *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
	// The number of hosts in the host group.
	MemberCount *int32 `json:"MemberCount,omitempty" xml:"MemberCount,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHostGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the request.
	HostGroupName *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
	// The host groups returned.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries to return on each page.
	//
	// The value of the PageSize parameter must not exceed 100. Default value: 20. If you leave the PageSize parameter empty, 20 entries are returned on each page.
	//
	// >  We recommend that you do not leave the PageSize parameter empty.
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The ID of the host group.
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The ID of the user.
	//
	// >  You can call the [ListUsers](~~204522~~) operation to query the ID of the user.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The number of the page to return. Default value: **1**.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Bastionhost instance where you want to query the host groups that the user is authorized or not authorized to manage.
	//
	// >  You can call the [DescribeInstances](~~153281~~) operation to query the ID of the Bastionhost instance.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// ListHostGroupsForUser
	HostGroups []*ListHostGroupsForUserResponseBodyHostGroups `json:"HostGroups,omitempty" xml:"HostGroups,omitempty" type:"Repeated"`
	// Queries the host groups that a specified user is authorized or not authorized to manage.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// All Bastionhost API requests must include common request parameters. For more information about common request parameters, see [Common parameters](~~148139~~).
	//
	// For more information about sample requests, see the "Examples" section of this topic.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// ListHostGroupsForUser
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// WB662865
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHostGroupsForUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The name of the host group that you want to query. Only exact match is supported.
	HostGroupName *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
	// The ID of the bastion host to which the user group belongs.
	//
	// > You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies the category of the host group that you want to query. Valid values:
	//
	// *   **Authorized**: queries the host groups that the user group is authorized to manage. This is the default value.
	// *   **Unauthorized**: queries the host groups that the user group is not authorized to manage.
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The number of the page to return. Default value: **1**.
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.\
	// Maximum value: 100. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// > We recommend that you do not leave this parameter empty.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the bastion host to which the user group belongs.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user group.
	//
	// > You can call the [ListUserGroups](~~204509~~) operation to query the ID of the user group.
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
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
	// The host groups returned.
	HostGroups []*ListHostGroupsForUserGroupResponseBodyHostGroups `json:"HostGroups,omitempty" xml:"HostGroups,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of host groups returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// The description of the host group.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the host group.
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// The name of the host group.
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHostGroupsForUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the bastion host. You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return. Default value: **1**.
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **20**.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the bastion host. For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// An array that consists of the shared keys.
	HostShareKeys []*ListHostShareKeysResponseBodyHostShareKeys `json:"HostShareKeys,omitempty" xml:"HostShareKeys,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the shared keys.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// The number of the associated host accounts.
	HostAccountCount *int64 `json:"HostAccountCount,omitempty" xml:"HostAccountCount,omitempty"`
	// The ID of the host account.
	HostShareKeyId *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	// The name of the shared key.
	HostShareKeyName *string `json:"HostShareKeyName,omitempty" xml:"HostShareKeyName,omitempty"`
	// The time when the shared key was last modified.
	LastModifyKeyAt *int64 `json:"LastModifyKeyAt,omitempty" xml:"LastModifyKeyAt,omitempty"`
	// The fingerprint of the private key.
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHostShareKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The address of the host that you want to query. You can set this parameter to a domain name or an IP address. Only exact match is supported.
	HostAddress *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	// The ID of the host group to which the host to be queried belongs.
	//
	// > You can call the [ListHostGroups](~~201307~~) operation to query the ID of the host group.
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// The name of the host that you want to query. Only exact match is supported.
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The ID of the bastion host on which you want to query hosts.
	//
	// > You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The operating system of the host that you want to query. Valid values:
	//
	// *   **Linux**
	// *   **Windows**
	OSType *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	// The number of the page to return. Default value: **1**.
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
	//
	// > We recommend that you do not leave this parameter empty.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the bastion host on which you want to query hosts.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The source of the host that you want to query. Valid values:
	//
	// *   **Local**: a host in a data center
	// *   **Ecs**: an Elastic Compute Service (ECS) instance
	// *   **Rds**: a host in an ApsaraDB MyBase dedicated cluster
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The ID of the ECS instance or the host in an ApsaraDB MyBase dedicated cluster that you want to query. Only exact match is supported.
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
	// The status of the host that you want to query. Valid values:
	//
	// *   **Normal**: normal
	// *   **Release**: released
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
	// An array that consists of the hosts returned.
	Hosts []*ListHostsResponseBodyHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of hosts returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// The address type of the host. Valid values:
	//
	// *   **Public**: a public address
	// *   **Private**: a private address
	ActiveAddressType *string `json:"ActiveAddressType,omitempty" xml:"ActiveAddressType,omitempty"`
	// The description of the host.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The number of host accounts.
	HostAccountCount *int32 `json:"HostAccountCount,omitempty" xml:"HostAccountCount,omitempty"`
	// The ID of the host.
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The name of the host.
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The private address of the host. The value is a domain name or an IP address.
	HostPrivateAddress *string `json:"HostPrivateAddress,omitempty" xml:"HostPrivateAddress,omitempty"`
	// The public address of the host. The value is a domain name or an IP address.
	HostPublicAddress *string `json:"HostPublicAddress,omitempty" xml:"HostPublicAddress,omitempty"`
	// The operating system of the host. Valid values:
	//
	// *   **Linux**
	// *   **Windows**
	OSType *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	// The source of the host. Valid values:
	//
	// *   **Local**: a host in a data center
	// *   **Ecs**: an ECS instance
	// *   **Rds**: a host in an ApsaraDB MyBase dedicated cluster
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The ID of the ECS instance or the host in an ApsaraDB MyBase dedicated cluster.
	//
	// > No value is returned for this parameter if the **Source** parameter is set to **Local**.
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
	// The status of the host. Valid values:
	//
	// *   **Normal**: normal
	// *   **Release**: released
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHostsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The endpoint of the host that you want to query. You can set this parameter to a domain name or an IP address. Only exact match is supported.
	HostAddress *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	// The name of the host that you want to query. Only exact match is supported.
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The ID of the bastion host on which you want to query the hosts that the user is authorized or not authorized to manage.
	//
	// > You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies the category of the hosts that you want to query. Valid values:
	//
	// *   **Authorized**: queries the hosts that the user is authorized to manage. This is the default value.
	// *   **Unauthorized**: queries the hosts that the user is not authorized to manage.
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The operating system of the host that you want to query. Valid values:
	//
	// *   **Linux**
	// *   **Windows**
	OSType *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	// The number of the page. Default value: 1.
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.\
	// Maximum value: 100. Default value: 20. If you leave this parameter empty, 20 entries are returned per page.
	//
	// > We recommend that you do not leave this parameter empty.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the bastion host on which you want to query the hosts that the user is authorized or not authorized to manage.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user.
	//
	// > You can call the [ListUsers](~~204522~~) operation to query the ID of the user.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// The hosts returned.
	Hosts []*ListHostsForUserResponseBodyHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of hosts returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// The endpoint type of the host. Valid values:
	//
	// *   **Public**: public endpoint
	// *   **Private**: internal endpoint
	ActiveAddressType *string `json:"ActiveAddressType,omitempty" xml:"ActiveAddressType,omitempty"`
	// The description of the host.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the host.
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The name of the host.
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The internal endpoint of the host. The value is a domain name or an IP address.
	HostPrivateAddress *string `json:"HostPrivateAddress,omitempty" xml:"HostPrivateAddress,omitempty"`
	// The public endpoint of the host. The value is a domain name or an IP address.
	HostPublicAddress *string `json:"HostPublicAddress,omitempty" xml:"HostPublicAddress,omitempty"`
	// The operating system of the host. Valid values:
	//
	// *   **Linux**
	// *   **Windows**
	OSType *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHostsForUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The operating system of the host that you want to query. Valid values:
	//
	// *   **Linux**
	// *   **Windows**
	HostAddress *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	// The ID of the Bastionhost instance where you want to query the hosts that the user group is authorized or not authorized to manage.
	//
	// >  You can call the [DescribeInstances](~~153281~~) operation to query the ID of the Bastionhost instance.
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The category of the host that you want to query. Valid values:
	//
	// *   **Authorized**: Query the hosts that the user group is authorized to manage. This is the default value.
	// *   **Unauthorized**: Query the hosts that the user group is not authorized to manage.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The operating system of the host. Valid values:
	//
	// *   **Linux**
	// *   **Windows**
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The internal endpoint of the host. The value is a domain name or an IP address.
	OSType *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	// The endpoint type of the host. Valid values:
	//
	// *   **Public**: a public endpoint
	// *   **Private**: an internal endpoint
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// The value of the PageSize parameter must not exceed 100. Default value: 20. If you leave the PageSize parameter empty, 20 entries are returned on each page.
	//
	// >  We recommend that you do not leave the PageSize parameter empty.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The endpoint of the host that you want to query. You can set this parameter to a domain name or an IP address. Only exact match is supported.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of the page to return. Default value: 1.
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
	// The ID of the user group for which you want to query hosts.
	//
	// >  You can call the [ListUserGroups](~~204509~~) operation to query the ID of the user group.
	Hosts []*ListHostsForUserGroupResponseBodyHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
	// The hosts returned.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The public endpoint of the host. The value is a domain name or an IP address.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// All Bastionhost API requests must include common request parameters. For more information about common request parameters, see [Common parameters](~~148139~~).
	//
	// For more information about sample requests, see the "Examples" section of this topic.
	ActiveAddressType *string `json:"ActiveAddressType,omitempty" xml:"ActiveAddressType,omitempty"`
	// The ID of the request.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	HostId  *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// ListHostsForUserGroup
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// WB662865
	HostPrivateAddress *string `json:"HostPrivateAddress,omitempty" xml:"HostPrivateAddress,omitempty"`
	// Queries the hosts that a specified user group is authorized or not authorized to manage.
	HostPublicAddress *string `json:"HostPublicAddress,omitempty" xml:"HostPublicAddress,omitempty"`
	// ListHostsForUserGroup
	OSType *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHostsForUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ListOperationTicketsRequest struct {
	AssetAddress *string `json:"AssetAddress,omitempty" xml:"AssetAddress,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber   *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListOperationTicketsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOperationTicketsRequest) GoString() string {
	return s.String()
}

func (s *ListOperationTicketsRequest) SetAssetAddress(v string) *ListOperationTicketsRequest {
	s.AssetAddress = &v
	return s
}

func (s *ListOperationTicketsRequest) SetInstanceId(v string) *ListOperationTicketsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListOperationTicketsRequest) SetPageNumber(v string) *ListOperationTicketsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOperationTicketsRequest) SetPageSize(v string) *ListOperationTicketsRequest {
	s.PageSize = &v
	return s
}

func (s *ListOperationTicketsRequest) SetRegionId(v string) *ListOperationTicketsRequest {
	s.RegionId = &v
	return s
}

type ListOperationTicketsResponseBody struct {
	OperationTickets []*ListOperationTicketsResponseBodyOperationTickets `json:"OperationTickets,omitempty" xml:"OperationTickets,omitempty" type:"Repeated"`
	RequestId        *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount       *int64                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOperationTicketsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOperationTicketsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOperationTicketsResponseBody) SetOperationTickets(v []*ListOperationTicketsResponseBodyOperationTickets) *ListOperationTicketsResponseBody {
	s.OperationTickets = v
	return s
}

func (s *ListOperationTicketsResponseBody) SetRequestId(v string) *ListOperationTicketsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOperationTicketsResponseBody) SetTotalCount(v int64) *ListOperationTicketsResponseBody {
	s.TotalCount = &v
	return s
}

type ListOperationTicketsResponseBodyOperationTickets struct {
	ApplyUserId           *string `json:"ApplyUserId,omitempty" xml:"ApplyUserId,omitempty"`
	ApplyUsername         *string `json:"ApplyUsername,omitempty" xml:"ApplyUsername,omitempty"`
	AssetAccountId        *string `json:"AssetAccountId,omitempty" xml:"AssetAccountId,omitempty"`
	AssetAccountName      *string `json:"AssetAccountName,omitempty" xml:"AssetAccountName,omitempty"`
	AssetAddress          *string `json:"AssetAddress,omitempty" xml:"AssetAddress,omitempty"`
	AssetId               *string `json:"AssetId,omitempty" xml:"AssetId,omitempty"`
	AssetName             *string `json:"AssetName,omitempty" xml:"AssetName,omitempty"`
	AssetNetworkDomainId  *string `json:"AssetNetworkDomainId,omitempty" xml:"AssetNetworkDomainId,omitempty"`
	AssetOs               *string `json:"AssetOs,omitempty" xml:"AssetOs,omitempty"`
	AssetSource           *string `json:"AssetSource,omitempty" xml:"AssetSource,omitempty"`
	AssetSourceInstanceId *string `json:"AssetSourceInstanceId,omitempty" xml:"AssetSourceInstanceId,omitempty"`
	CreatedTime           *int64  `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	OperationTicketId     *string `json:"OperationTicketId,omitempty" xml:"OperationTicketId,omitempty"`
	ProtocolName          *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
	State                 *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListOperationTicketsResponseBodyOperationTickets) String() string {
	return tea.Prettify(s)
}

func (s ListOperationTicketsResponseBodyOperationTickets) GoString() string {
	return s.String()
}

func (s *ListOperationTicketsResponseBodyOperationTickets) SetApplyUserId(v string) *ListOperationTicketsResponseBodyOperationTickets {
	s.ApplyUserId = &v
	return s
}

func (s *ListOperationTicketsResponseBodyOperationTickets) SetApplyUsername(v string) *ListOperationTicketsResponseBodyOperationTickets {
	s.ApplyUsername = &v
	return s
}

func (s *ListOperationTicketsResponseBodyOperationTickets) SetAssetAccountId(v string) *ListOperationTicketsResponseBodyOperationTickets {
	s.AssetAccountId = &v
	return s
}

func (s *ListOperationTicketsResponseBodyOperationTickets) SetAssetAccountName(v string) *ListOperationTicketsResponseBodyOperationTickets {
	s.AssetAccountName = &v
	return s
}

func (s *ListOperationTicketsResponseBodyOperationTickets) SetAssetAddress(v string) *ListOperationTicketsResponseBodyOperationTickets {
	s.AssetAddress = &v
	return s
}

func (s *ListOperationTicketsResponseBodyOperationTickets) SetAssetId(v string) *ListOperationTicketsResponseBodyOperationTickets {
	s.AssetId = &v
	return s
}

func (s *ListOperationTicketsResponseBodyOperationTickets) SetAssetName(v string) *ListOperationTicketsResponseBodyOperationTickets {
	s.AssetName = &v
	return s
}

func (s *ListOperationTicketsResponseBodyOperationTickets) SetAssetNetworkDomainId(v string) *ListOperationTicketsResponseBodyOperationTickets {
	s.AssetNetworkDomainId = &v
	return s
}

func (s *ListOperationTicketsResponseBodyOperationTickets) SetAssetOs(v string) *ListOperationTicketsResponseBodyOperationTickets {
	s.AssetOs = &v
	return s
}

func (s *ListOperationTicketsResponseBodyOperationTickets) SetAssetSource(v string) *ListOperationTicketsResponseBodyOperationTickets {
	s.AssetSource = &v
	return s
}

func (s *ListOperationTicketsResponseBodyOperationTickets) SetAssetSourceInstanceId(v string) *ListOperationTicketsResponseBodyOperationTickets {
	s.AssetSourceInstanceId = &v
	return s
}

func (s *ListOperationTicketsResponseBodyOperationTickets) SetCreatedTime(v int64) *ListOperationTicketsResponseBodyOperationTickets {
	s.CreatedTime = &v
	return s
}

func (s *ListOperationTicketsResponseBodyOperationTickets) SetOperationTicketId(v string) *ListOperationTicketsResponseBodyOperationTickets {
	s.OperationTicketId = &v
	return s
}

func (s *ListOperationTicketsResponseBodyOperationTickets) SetProtocolName(v string) *ListOperationTicketsResponseBodyOperationTickets {
	s.ProtocolName = &v
	return s
}

func (s *ListOperationTicketsResponseBodyOperationTickets) SetState(v string) *ListOperationTicketsResponseBodyOperationTickets {
	s.State = &v
	return s
}

type ListOperationTicketsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOperationTicketsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOperationTicketsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOperationTicketsResponse) GoString() string {
	return s.String()
}

func (s *ListOperationTicketsResponse) SetHeaders(v map[string]*string) *ListOperationTicketsResponse {
	s.Headers = v
	return s
}

func (s *ListOperationTicketsResponse) SetStatusCode(v int32) *ListOperationTicketsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOperationTicketsResponse) SetBody(v *ListOperationTicketsResponseBody) *ListOperationTicketsResponse {
	s.Body = v
	return s
}

type ListTagKeysRequest struct {
	// The number of the page to return.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the bastion host.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the resource.
	//
	// Set the value to INSTANCE, which indicates that the resource is a bastion host.
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
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of tags.
	TagKeys []*ListTagKeysResponseBodyTagKeys `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Repeated"`
	// The total number of tags returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// The total number of tag keys.
	TagCount *int32 `json:"TagCount,omitempty" xml:"TagCount,omitempty"`
	// The name of the tag key.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The region ID of the Bastionhost instance.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The value of the tag.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The operation that you want to perform.
	//
	// Set the value to **ListTagResources**.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The key of tag N.
	//
	// Valid values of N: 1 to 20.
	Tag []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	// The ID of the request.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The type of the resource.
	//
	// The returned value is INSTANCE, which indicates that the resource is a Bastionhost instance.
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
	// Queries the tags bound to one or more Bastionhost instances.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// ListTagResources
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 58928
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ListUserPublicKeysRequest struct {
	// The ID of the bastion host on which you want to query all public keys of the user.
	//
	// > You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return. Default value: 1.
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.\
	// Maximum value: 100. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// > We recommend that you do not leave this parameter empty.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the bastion host on which you want to query all public keys of the user.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user whose public keys you want to query.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListUserPublicKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserPublicKeysRequest) GoString() string {
	return s.String()
}

func (s *ListUserPublicKeysRequest) SetInstanceId(v string) *ListUserPublicKeysRequest {
	s.InstanceId = &v
	return s
}

func (s *ListUserPublicKeysRequest) SetPageNumber(v string) *ListUserPublicKeysRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUserPublicKeysRequest) SetPageSize(v string) *ListUserPublicKeysRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserPublicKeysRequest) SetRegionId(v string) *ListUserPublicKeysRequest {
	s.RegionId = &v
	return s
}

func (s *ListUserPublicKeysRequest) SetUserId(v string) *ListUserPublicKeysRequest {
	s.UserId = &v
	return s
}

type ListUserPublicKeysResponseBody struct {
	// An array that consists of the public keys of the user.
	PublicKeys []*ListUserPublicKeysResponseBodyPublicKeys `json:"PublicKeys,omitempty" xml:"PublicKeys,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of public keys.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListUserPublicKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserPublicKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserPublicKeysResponseBody) SetPublicKeys(v []*ListUserPublicKeysResponseBodyPublicKeys) *ListUserPublicKeysResponseBody {
	s.PublicKeys = v
	return s
}

func (s *ListUserPublicKeysResponseBody) SetRequestId(v string) *ListUserPublicKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserPublicKeysResponseBody) SetTotalCount(v int64) *ListUserPublicKeysResponseBody {
	s.TotalCount = &v
	return s
}

type ListUserPublicKeysResponseBodyPublicKeys struct {
	// The description of the public key.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The fingerprint of the public key.
	FingerPrint *string `json:"FingerPrint,omitempty" xml:"FingerPrint,omitempty"`
	// The ID of the public key.
	PublicKeyId *string `json:"PublicKeyId,omitempty" xml:"PublicKeyId,omitempty"`
	// The name of the public key.
	PublicKeyName *string `json:"PublicKeyName,omitempty" xml:"PublicKeyName,omitempty"`
	// The ID of the user to which the public key belongs.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListUserPublicKeysResponseBodyPublicKeys) String() string {
	return tea.Prettify(s)
}

func (s ListUserPublicKeysResponseBodyPublicKeys) GoString() string {
	return s.String()
}

func (s *ListUserPublicKeysResponseBodyPublicKeys) SetComment(v string) *ListUserPublicKeysResponseBodyPublicKeys {
	s.Comment = &v
	return s
}

func (s *ListUserPublicKeysResponseBodyPublicKeys) SetFingerPrint(v string) *ListUserPublicKeysResponseBodyPublicKeys {
	s.FingerPrint = &v
	return s
}

func (s *ListUserPublicKeysResponseBodyPublicKeys) SetPublicKeyId(v string) *ListUserPublicKeysResponseBodyPublicKeys {
	s.PublicKeyId = &v
	return s
}

func (s *ListUserPublicKeysResponseBodyPublicKeys) SetPublicKeyName(v string) *ListUserPublicKeysResponseBodyPublicKeys {
	s.PublicKeyName = &v
	return s
}

func (s *ListUserPublicKeysResponseBodyPublicKeys) SetUserId(v string) *ListUserPublicKeysResponseBodyPublicKeys {
	s.UserId = &v
	return s
}

type ListUserPublicKeysResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserPublicKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserPublicKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserPublicKeysResponse) GoString() string {
	return s.String()
}

func (s *ListUserPublicKeysResponse) SetHeaders(v map[string]*string) *ListUserPublicKeysResponse {
	s.Headers = v
	return s
}

func (s *ListUserPublicKeysResponse) SetStatusCode(v int32) *ListUserPublicKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserPublicKeysResponse) SetBody(v *ListUserPublicKeysResponseBody) *ListUserPublicKeysResponse {
	s.Body = v
	return s
}

type ListUsersRequest struct {
	// The display name of the user to be queried. Only exact match is supported.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the Bastionhost instance to which the users to be queried belong.
	//
	// >  You can call the [DescribeInstances](~~153281~~) operation to query the ID of the Bastionhost instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The mobile number of the user to be queried. Only exact match is supported.
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The number of the page to return. Default value: **1**.
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// The value of the PageSize parameter must not exceed 100. By default, the number of entries on each page is 20. If you do not set the PageSize parameter, 20 entries are returned per page by default.
	//
	// >  We recommend that you do not leave this parameter empty.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the Bastionhost instance to which the users to be queried belong.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The source of the user to be queried. Valid values:
	//
	// *   **Local**: a local user
	// *   **Ram**: a RAM user
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The unique ID of the user to be queried. Only exact match is supported.
	//
	// >  This parameter uniquely identifies a RAM user of the Bastionhost instance. This parameter takes effect only when the **Source** parameter is set to **Ram**. You can call the [ListUsers](~~28684~~) operation to obtain the unique ID of the user from the **UserId** response parameter.
	SourceUserId *string `json:"SourceUserId,omitempty" xml:"SourceUserId,omitempty"`
	// The ID of the user group to be queried.
	//
	// >  You can call the [ListUserGroups](~~204509~~) operation to query the ID of the user group.
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// The logon name of the user to be queried. Only exact match is supported.
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The status of the user to be queried. Valid values:
	//
	// *   **Normal**: The user can access the Bastionhost instance.
	// *   **Frozen**: The user is locked and cannot access the Bastionhost instance.
	// *   **Expired**: The user has expired and cannot access the Bastionhost instance.
	UserState *string `json:"UserState,omitempty" xml:"UserState,omitempty"`
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
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of users that were queried.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of users that were queried.
	Users []*ListUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
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
	// The description of the user.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The display name of the user.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The end of the validity period of the user. The value is a UNIX timestamp. Unit: seconds.
	EffectiveEndTime *int64 `json:"EffectiveEndTime,omitempty" xml:"EffectiveEndTime,omitempty"`
	// The beginning of the validity period of the user. The value is a UNIX timestamp. Unit: seconds.
	EffectiveStartTime *int64 `json:"EffectiveStartTime,omitempty" xml:"EffectiveStartTime,omitempty"`
	// The email address of the user.
	Email          *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Language       *string `json:"Language,omitempty" xml:"Language,omitempty"`
	LanguageStatus *string `json:"LanguageStatus,omitempty" xml:"LanguageStatus,omitempty"`
	// The mobile number of the user.
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The country where the mobile number of the user is registered. Valid values:
	//
	// *   **CN**: the Chinese mainland, whose country calling code is +86
	// *   **HK**: Hong Kong (China), whose country calling code is +852
	// *   **MO**: Macau (China), whose country calling code is +853
	// *   **TW**: Taiwan (China), whose country calling code is +886
	// *   **RU**: Russia, whose country calling code is +7
	// *   **SG**: Singapore, whose country calling code is +65
	// *   **MY**: Malaysia, whose country calling code is +60
	// *   **ID**: Indonesia, whose country calling code is +62
	// *   **DE**: Germany, whose country calling code is +49
	// *   **AU**: Australia, whose country calling code is +61
	// *   **US**: United States, whose country calling code is +1
	// *   **AE**: United Arab Emirates, whose country calling code is +971
	// *   **JP**: Japan, whose country calling code is +81
	// *   **GB**: United Kingdom, whose country calling code is +44
	// *   **IN**: India, whose country calling code is +91
	// *   **KR**: South Korea, whose country calling code is +82
	// *   **PH**: Philippines, whose country calling code is +63
	// *   **CH**: Switzerland, whose country calling code is +41
	// *   **SE**: Sweden, whose country calling code is +46
	MobileCountryCode *string `json:"MobileCountryCode,omitempty" xml:"MobileCountryCode,omitempty"`
	// Specifies whether password reset is required upon the next logon. Valid values:
	//
	// - true: yes
	// - false: no
	NeedResetPassword *bool `json:"NeedResetPassword,omitempty" xml:"NeedResetPassword,omitempty"`
	// The source of the user. Valid values:
	//
	// *   **Local**: a local user
	// *   **Ram**: a RAM user
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The unique ID of the user.
	//
	// >  This parameter uniquely identifies a RAM user of the Bastionhost instance. A value is returned for this parameter if the **Source** parameter is set to **Ram**. No value is returned for this parameter if the **Source** parameter is set to **Local**.
	SourceUserId *string `json:"SourceUserId,omitempty" xml:"SourceUserId,omitempty"`
	// The two-factor authentication method.
	TwoFactorMethods []*string `json:"TwoFactorMethods,omitempty" xml:"TwoFactorMethods,omitempty" type:"Repeated"`
	// The two-factor authentication status of the user. Valid values:
	//
	// *   **Global:** follows the global settings
	// *   **Disable:** disables two-factor authentication
	// *   **Enable:** enable two-factor authentication and follows settings of the single user
	TwoFactorStatus *string `json:"TwoFactorStatus,omitempty" xml:"TwoFactorStatus,omitempty"`
	// The ID of the user.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The logon name of the user.
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The statuses of the user.
	UserState []*string `json:"UserState,omitempty" xml:"UserState,omitempty" type:"Repeated"`
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

func (s *ListUsersResponseBodyUsers) SetEffectiveEndTime(v int64) *ListUsersResponseBodyUsers {
	s.EffectiveEndTime = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetEffectiveStartTime(v int64) *ListUsersResponseBodyUsers {
	s.EffectiveStartTime = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetEmail(v string) *ListUsersResponseBodyUsers {
	s.Email = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetLanguage(v string) *ListUsersResponseBodyUsers {
	s.Language = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetLanguageStatus(v string) *ListUsersResponseBodyUsers {
	s.LanguageStatus = &v
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

func (s *ListUsersResponseBodyUsers) SetNeedResetPassword(v bool) *ListUsersResponseBodyUsers {
	s.NeedResetPassword = &v
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

func (s *ListUsersResponseBodyUsers) SetTwoFactorMethods(v []*string) *ListUsersResponseBodyUsers {
	s.TwoFactorMethods = v
	return s
}

func (s *ListUsersResponseBodyUsers) SetTwoFactorStatus(v string) *ListUsersResponseBodyUsers {
	s.TwoFactorStatus = &v
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the bastion host to which the users to be locked belong.
	//
	// > You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host to which the users to be locked belong.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user to be locked. The value is a JSON string. You can add up to 100 user IDs. If you specify multiple IDs, separate the IDs with commas (,).
	//
	// > You can call the [ListUsers](~~204522~~) operation to query the ID of the user.
	UserIds *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
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
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the call.
	Results []*LockUsersResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
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
	// The return code that indicates whether the call was successful. Valid values:
	//
	// *   **OK**: The call was successful.
	//
	// *   **UNEXPECTED**: An unknown error occurred.
	//
	// *   **INVALID_ARGUMENT**: A request parameter is invalid.
	//
	// >Make sure that the request parameters are valid and call the operation again.
	//
	// *   **OBJECT_NOT_FOUND**: The specified object on which you want to perform the operation does not exist.
	//
	// >Check whether the specified ID of the bastion host exists, whether the specified hosts exist, and whether the specified host IDs are valid. Then, call the operation again.
	//
	// *   **OBJECT_AlREADY_EXISTS**: The specified object on which you want to perform the operation already exists.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// This parameter is deprecated.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the user.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LockUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The new description of the host. The description can be up to 500 characters in length.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the host.
	//
	// > You can call the [ListHosts](~~200665~~) operation to query the ID of the host.
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The new name of the host. The name can be up to 128 characters.
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The new internal endpoint of the host. You can set this parameter to a domain name or an IP address.
	HostPrivateAddress *string `json:"HostPrivateAddress,omitempty" xml:"HostPrivateAddress,omitempty"`
	// The new public endpoint of the host. You can set this parameter to a domain name or an IP address.
	HostPublicAddress *string `json:"HostPublicAddress,omitempty" xml:"HostPublicAddress,omitempty"`
	// The ID of the bastion host on which you want to modify the information about the host.
	//
	// > You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the new network domain to which the host belongs.
	NetworkDomainId *string `json:"NetworkDomainId,omitempty" xml:"NetworkDomainId,omitempty"`
	// The new operating system of the host. Valid values:
	//
	// *   **Linux**
	// *   **Windows**
	OSType *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	// The region ID of the bastion host on which you want to modify the information about the host.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *ModifyHostRequest) SetNetworkDomainId(v string) *ModifyHostRequest {
	s.NetworkDomainId = &v
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
	// The ID of the request.
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHostResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the host account whose information you want to modify.
	//
	// > You can call the [ListHostAccounts](~~204372~~) operation to query the ID of the host account.
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	// The new name of the host account. The name can be up to 128 characters in length.
	HostAccountName *string `json:"HostAccountName,omitempty" xml:"HostAccountName,omitempty"`
	// The ID of the shared key.
	HostShareKeyId *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	// The ID of the bastion host in which you want to modify the information about the host account.
	//
	// > You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The passphrase of the new private key for the host account.
	//
	// > This parameter takes effect only when the protocol of the host is set to SSH. If the protocol of the host is set to RDP, this parameter is not required.
	PassPhrase *string `json:"PassPhrase,omitempty" xml:"PassPhrase,omitempty"`
	// The new password of the host account.
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The new private key of the host account. The value is a Base64-encoded string.
	//
	// > This parameter takes effect only when the protocol of the host is set to SSH. If the protocol of the host is set to RDP, this parameter is not required. You can call the [GetHostAccount](~~204391~~) operation to query the protocol used by the host. You can configure a password and a private key for the host account at the same time. If both a password and a private key are configured for the host account, Bastionhost preferentially uses the private key for logon.
	PrivateKey *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	// The region ID of the bastion host in which you want to query the details of the host account.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The ID of the request.
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHostAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The new name of the host group. The name can be up to 128 characters in length.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The region ID of the Bastionhost instance where you want to modify the information of the host group.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// The ID of the host group that you want to modify.
	//
	// >  You can call the [ListHostGroups](~~201307~~) operation to query the ID of the host group.
	HostGroupName *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
	// The ID of the request.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the Bastionhost instance where you want to modify the information of the host group.
	//
	// >  You can call the [DescribeInstances](~~153281~~) operation to query the ID of the Bastionhost instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The ID of the request.
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHostGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the shared key whose information you want to modify.
	HostShareKeyId *string `json:"HostShareKeyId,omitempty" xml:"HostShareKeyId,omitempty"`
	// The name of the shared key.
	HostShareKeyName *string `json:"HostShareKeyName,omitempty" xml:"HostShareKeyName,omitempty"`
	// The ID of the bastion host. You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The passphrase of the private key. The value is a Base64-encoded string.
	PassPhrase *string `json:"PassPhrase,omitempty" xml:"PassPhrase,omitempty"`
	// The private key. The value is a Base64-encoded string.
	PrivateKey *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	// The region ID of the bastion host. For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The ID of the request, which is used to locate and troubleshoot issues.
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHostShareKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The new portal type of the host. Valid values:
	//
	// *   **Public**: public portal
	// *   **Private**: internal portal
	ActiveAddressType *string `json:"ActiveAddressType,omitempty" xml:"ActiveAddressType,omitempty"`
	// The ID of the host for which you want to change the portal type. The value is a JSON string. You can add up to 100 host IDs.
	//
	// >  You can call the [ListHosts](~~200665~~) operation to query the ID of the host.
	HostIds *string `json:"HostIds,omitempty" xml:"HostIds,omitempty"`
	// The ID of the bastion host for which you want to change the portal type of the host.
	//
	// >  You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host for which you want to change the portal type of the host.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the call.
	Results []*ModifyHostsActiveAddressTypeResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
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
	// The return code that indicates whether the call was successful. Valid values:
	//
	// *   **OK**: The call was successful.
	// *   **UNEXPECTED**: An unknown error occurred.
	// *   **INVALID_ARGUMENT**: A request parameter is invalid.
	// *   **OBJECT_NOT_FOUND**: The specified object on which you want to perform the operation does not exist.
	// *   **OBJECT_AlREADY_EXISTS**: The specified object on which you want to perform the operation already exists.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the host.
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// This parameter is deprecated.
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHostsActiveAddressTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the host for which you want to change the port. The value is a JSON string. You can add up to 100 host IDs. If you specify multiple IDs, separate the IDs with commas (,).
	//
	// >  You can call the [ListHosts](~~200665~~) operation to query the IDs of hosts.
	HostIds *string `json:"HostIds,omitempty" xml:"HostIds,omitempty"`
	// The ID of the bastion host for which you want to change the port of the host.
	//
	// >  You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The new port of the host. The port number must be an integer. Valid values: 22 to 65535.
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol that is used to connect to the host. Valid values:
	//
	// *   **SSH**
	// *   **RDP**
	ProtocolName *string `json:"ProtocolName,omitempty" xml:"ProtocolName,omitempty"`
	// The region ID of the bastion host for which you want to change the port of the host.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the call.
	Results []*ModifyHostsPortResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
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
	// The return code that indicates whether the call was successful. Valid values:
	//
	// *   **OK**: The call was successful.
	// *   **UNEXPECTED**: An unknown error occurred.
	// *   **INVALID_ARGUMENT**: A request parameter is invalid.
	//     > Make sure that the request parameters are valid and call the operation again.
	//
	// *   **OBJECT_NOT_FOUND**: The specified object on which you want to perform the operation does not exist.
	//
	//     > Check whether the specified ID of the bastion host exists, whether the specified hosts exist, and whether the specified host IDs are valid. Then, call the operation again.
	//
	// *   **OBJECT\_AlREADY\_EXISTS**: The specified object on which you want to perform the operation already exists.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the host.
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// This parameter is deprecated.
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHostsPortResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The username of the account that is used for the AD server.
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The Base distinguished name (DN).
	BaseDN *string `json:"BaseDN,omitempty" xml:"BaseDN,omitempty"`
	// The domain on the AD server.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The field that is used to indicate the email address of a user on the AD server.
	EmailMapping *string `json:"EmailMapping,omitempty" xml:"EmailMapping,omitempty"`
	// The condition that is used to filter users.
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The ID of the bastion host. You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether to support SSL. Valid values:
	//
	// *   **true**: yes
	// *   **false**: no
	IsSSL *string `json:"IsSSL,omitempty" xml:"IsSSL,omitempty"`
	// The field that is used to indicate the mobile phone number of a user on the AD server.
	MobileMapping *string `json:"MobileMapping,omitempty" xml:"MobileMapping,omitempty"`
	// The field that is used to indicate the name of a user on the AD server.
	NameMapping *string `json:"NameMapping,omitempty" xml:"NameMapping,omitempty"`
	// The password of the account that is used for the AD server.
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The port that is used to access the AD server.
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The region ID of the bastion host. For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The address of the AD server.
	Server *string `json:"Server,omitempty" xml:"Server,omitempty"`
	// The address of the secondary AD server.
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
	// The ID of the request, which is used to locate and troubleshoot issues.
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceADAuthServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The description of the bastion host.
	//
	// > The description can contain only letters, digits, underscores (\_), and hyphens (-). The description can be up to 30 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the bastion host.
	//
	// > You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The ID of the request, which is used to locate and troubleshoot issues.
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The username of the account that is used for the LDAP server.
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The Base distinguished name (DN).
	BaseDN *string `json:"BaseDN,omitempty" xml:"BaseDN,omitempty"`
	// The field that is used to indicate the email address of a user on the LDAP server.
	EmailMapping *string `json:"EmailMapping,omitempty" xml:"EmailMapping,omitempty"`
	// The condition that is used to filter users.
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The ID of the bastion host. You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether to support SSL. Valid values:
	//
	// *   **true**
	// *   **false**
	IsSSL *string `json:"IsSSL,omitempty" xml:"IsSSL,omitempty"`
	// The field that is used to indicate the logon name of a user on the LDAP server.
	LoginNameMapping *string `json:"LoginNameMapping,omitempty" xml:"LoginNameMapping,omitempty"`
	// The field that is used to indicate the mobile phone number of a user on the LDAP server.
	MobileMapping *string `json:"MobileMapping,omitempty" xml:"MobileMapping,omitempty"`
	// The field that is used to indicate the name of a user on the LDAP server.
	NameMapping *string `json:"NameMapping,omitempty" xml:"NameMapping,omitempty"`
	// The password of the account that is used for the LDAP server. You must configure a password when you configure LDAP authentication. If you leave this parameter empty when you modify the settings of LDAP authentication, the current password is used.
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The port that is used to access the LDAP server.
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The region ID of the bastion host. For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The address of the LDAP server.
	Server *string `json:"Server,omitempty" xml:"Server,omitempty"`
	// The address of the secondary LDAP server.
	StandbyServer *string `json:"StandbyServer,omitempty" xml:"StandbyServer,omitempty"`
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
	// The ID of the request, which is used to locate and troubleshoot issues.
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceLDAPAuthServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// Specifies whether to enable two-factor authentication. Valid values:
	//
	// *   **true**: enables two-factor authentication.
	// *   **false**: disables two-factor authentication.
	EnableTwoFactor *string `json:"EnableTwoFactor,omitempty" xml:"EnableTwoFactor,omitempty"`
	// The ID of the bastion host.
	//
	// >  You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	SkipTwoFactorTime *string `json:"SkipTwoFactorTime,omitempty" xml:"SkipTwoFactorTime,omitempty"`
	// One or more methods that are used to send a verification code if two-factor authentication is enabled. If you set the EnableTwoFactor parameter to true, you must specify at least one method. Valid values:
	//
	// *   **sms**: text message
	// *   **email**: email
	// *   **dingtalk**: Notice in DingTalk
	TwoFactorMethods *string `json:"TwoFactorMethods,omitempty" xml:"TwoFactorMethods,omitempty"`
}

func (s ModifyInstanceTwoFactorRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceTwoFactorRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceTwoFactorRequest) SetEnableTwoFactor(v string) *ModifyInstanceTwoFactorRequest {
	s.EnableTwoFactor = &v
	return s
}

func (s *ModifyInstanceTwoFactorRequest) SetInstanceId(v string) *ModifyInstanceTwoFactorRequest {
	s.InstanceId = &v
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
	// The duration within which two-factor authentication is not required after a local user passes two-factor authentication. Valid values: 0 to 168. Unit: hours. If you set this parameter to 0, the local user must pass two-factor authentication every time the local user logs on to the bastion host.
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceTwoFactorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The new description of the user. The description can be up to 500 characters in length.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The new display name of the user. This display name can be up to 128 characters in length.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The end of the validity period of the user. The value is a UNIX timestamp. Unit: seconds.
	EffectiveEndTime *int64 `json:"EffectiveEndTime,omitempty" xml:"EffectiveEndTime,omitempty"`
	// The beginning of the validity period of the user. The value is a UNIX timestamp. Unit: seconds.
	EffectiveStartTime *int64 `json:"EffectiveStartTime,omitempty" xml:"EffectiveStartTime,omitempty"`
	// The new email address of the user.
	//
	// > This parameter is required when the TwoFactorStatus parameter is set to Enable and the TwoFactorMethods parameter is set to email.
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The ID of the bastion host where you want to modify user information.
	//
	// > You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Language       *string `json:"Language,omitempty" xml:"Language,omitempty"`
	LanguageStatus *string `json:"LanguageStatus,omitempty" xml:"LanguageStatus,omitempty"`
	// The new mobile number of the user.
	//
	// > This parameter is required when the TwoFactorStatus parameter is set to Enable and the TwoFactorMethods parameter is set to sms or dingtalk.
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The country where the new mobile number of the user is registered. Valid values:
	//
	// *   **CN:** the Chinese mainland, whose country calling code is +86
	// *   **HK:** Hong Kong (China), whose country calling code is +852
	// *   **MO:** Macao (China), whose country calling code is +853
	// *   **TW:** Taiwan (China), whose country calling code is +886
	// *   **RU:** Russia, whose country calling code is +7
	// *   **SG:** Singapore, whose country calling code is +65
	// *   **MY:** Malaysia, whose country calling code is +60
	// *   **ID:** Indonesia, whose country calling code is +62
	// *   **DE:** Germany, whose country calling code is +49
	// *   **AU:** Australia, whose country calling code is +61
	// *   **US:** US, whose country calling code is +1
	// *   **AE:** United Arab Emirates, whose country calling code is +971
	// *   **JP:** Japan, whose country calling code is +81
	// *   **GB:** UK, whose country calling code is +44
	// *   **IN:** India, whose country calling code is +91
	// *   **KR:** Republic of Korea, whose country calling code is +82
	// *   **PH:** Philippines, whose country calling code is +63
	// *   **CH:** Switzerland, whose country calling code is +41
	// *   **SE:** Sweden, whose country calling code is +46
	// *   **SA:** Saudi Arabia, whose country calling code is +966
	MobileCountryCode *string `json:"MobileCountryCode,omitempty" xml:"MobileCountryCode,omitempty"`
	// Specifies whether password reset is required upon the next logon. Valid values:
	//
	// - true: yes
	// - false: no
	NeedResetPassword *bool `json:"NeedResetPassword,omitempty" xml:"NeedResetPassword,omitempty"`
	// The new password of the user. The password must be 8 to 128 characters in length and must contain lowercase letters, uppercase letters, digits, and special characters.
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The region ID of the bastion host where you want to modify user information.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The two-factor authentication method. You can select only one method. Valid values:
	//
	// *   **sms:** text message
	// *   **email:** email
	// *   **dingtalk:** DingTalk
	// *   **totp OTP:** time-based one-time password (TOTP) app
	//
	// > *   When the TwoFactorStatus parameter is set to Enable, you must specify one of the preceding values.
	TwoFactorMethods *string `json:"TwoFactorMethods,omitempty" xml:"TwoFactorMethods,omitempty"`
	// The two-factor authentication status of the user. Valid values:
	//
	// *   **Global:** follows the global settings
	// *   **Disable:** disables two-factor authentication
	// *   **Enable:** enable two-factor authentication and follows settings of the single user
	TwoFactorStatus *string `json:"TwoFactorStatus,omitempty" xml:"TwoFactorStatus,omitempty"`
	// The ID of the user whose information you want to modify.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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

func (s *ModifyUserRequest) SetEffectiveEndTime(v int64) *ModifyUserRequest {
	s.EffectiveEndTime = &v
	return s
}

func (s *ModifyUserRequest) SetEffectiveStartTime(v int64) *ModifyUserRequest {
	s.EffectiveStartTime = &v
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

func (s *ModifyUserRequest) SetLanguage(v string) *ModifyUserRequest {
	s.Language = &v
	return s
}

func (s *ModifyUserRequest) SetLanguageStatus(v string) *ModifyUserRequest {
	s.LanguageStatus = &v
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

func (s *ModifyUserRequest) SetNeedResetPassword(v bool) *ModifyUserRequest {
	s.NeedResetPassword = &v
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

func (s *ModifyUserRequest) SetTwoFactorMethods(v string) *ModifyUserRequest {
	s.TwoFactorMethods = &v
	return s
}

func (s *ModifyUserRequest) SetTwoFactorStatus(v string) *ModifyUserRequest {
	s.TwoFactorStatus = &v
	return s
}

func (s *ModifyUserRequest) SetUserId(v string) *ModifyUserRequest {
	s.UserId = &v
	return s
}

type ModifyUserResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The new description of the user group. The description can be up to 500 characters in length.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the bastion host in which you want to modify the information about the user group.
	//
	// > You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host in which you want to modify the information about the user group.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user group that you want to modify.
	//
	// > You can call the [ListUserGroups](~~204509~~) operation to query the ID of the user group.
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// The new name of the user group. This name can be up to 128 characters in length.
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
	// The ID of the request.
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The region ID of the bastion host.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the bastion host is moved.
	//
	// > You can call the [DescribeInstances](~~153281~~) operation to query the resource group ID of the bastion host.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the bastion host for which you want to change the resource group.
	//
	// > You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource. Set the value to **INSTANCE**, which indicates that the resource is a bastion host.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
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
	// The ID of the request, which is used to locate and troubleshoot issues.
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type RejectApproveCommandRequest struct {
	CommandId  *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RejectApproveCommandRequest) String() string {
	return tea.Prettify(s)
}

func (s RejectApproveCommandRequest) GoString() string {
	return s.String()
}

func (s *RejectApproveCommandRequest) SetCommandId(v string) *RejectApproveCommandRequest {
	s.CommandId = &v
	return s
}

func (s *RejectApproveCommandRequest) SetInstanceId(v string) *RejectApproveCommandRequest {
	s.InstanceId = &v
	return s
}

func (s *RejectApproveCommandRequest) SetRegionId(v string) *RejectApproveCommandRequest {
	s.RegionId = &v
	return s
}

type RejectApproveCommandResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RejectApproveCommandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RejectApproveCommandResponseBody) GoString() string {
	return s.String()
}

func (s *RejectApproveCommandResponseBody) SetRequestId(v string) *RejectApproveCommandResponseBody {
	s.RequestId = &v
	return s
}

type RejectApproveCommandResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RejectApproveCommandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RejectApproveCommandResponse) String() string {
	return tea.Prettify(s)
}

func (s RejectApproveCommandResponse) GoString() string {
	return s.String()
}

func (s *RejectApproveCommandResponse) SetHeaders(v map[string]*string) *RejectApproveCommandResponse {
	s.Headers = v
	return s
}

func (s *RejectApproveCommandResponse) SetStatusCode(v int32) *RejectApproveCommandResponse {
	s.StatusCode = &v
	return s
}

func (s *RejectApproveCommandResponse) SetBody(v *RejectApproveCommandResponseBody) *RejectApproveCommandResponse {
	s.Body = v
	return s
}

type RejectOperationTicketRequest struct {
	// The ID of the bastion host.
	//
	// >  You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the O\&M application that you want to reject. You can call the ListOperationTickets operation to query the IDs of all O\&M applications that require review.
	OperationTicketId *string `json:"OperationTicketId,omitempty" xml:"OperationTicketId,omitempty"`
	// The region ID of the bastion host.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RejectOperationTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s RejectOperationTicketRequest) GoString() string {
	return s.String()
}

func (s *RejectOperationTicketRequest) SetInstanceId(v string) *RejectOperationTicketRequest {
	s.InstanceId = &v
	return s
}

func (s *RejectOperationTicketRequest) SetOperationTicketId(v string) *RejectOperationTicketRequest {
	s.OperationTicketId = &v
	return s
}

func (s *RejectOperationTicketRequest) SetRegionId(v string) *RejectOperationTicketRequest {
	s.RegionId = &v
	return s
}

type RejectOperationTicketResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RejectOperationTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RejectOperationTicketResponseBody) GoString() string {
	return s.String()
}

func (s *RejectOperationTicketResponseBody) SetRequestId(v string) *RejectOperationTicketResponseBody {
	s.RequestId = &v
	return s
}

type RejectOperationTicketResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RejectOperationTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RejectOperationTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s RejectOperationTicketResponse) GoString() string {
	return s.String()
}

func (s *RejectOperationTicketResponse) SetHeaders(v map[string]*string) *RejectOperationTicketResponse {
	s.Headers = v
	return s
}

func (s *RejectOperationTicketResponse) SetStatusCode(v int32) *RejectOperationTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *RejectOperationTicketResponse) SetBody(v *RejectOperationTicketResponseBody) *RejectOperationTicketResponse {
	s.Body = v
	return s
}

type RemoveHostsFromGroupRequest struct {
	// The ID of the host group from which you want to remove hosts.
	//
	// >  You can call the [ListHostGroups](~~201307~~) operation to query the ID of the host group.
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// The ID of the host that you want to remove from the host group. The value is a JSON string. You can add up to 100 host IDs.
	//
	// >  You can call the [ListHosts](~~200665~~) operation to query the IDs of hosts.
	HostIds *string `json:"HostIds,omitempty" xml:"HostIds,omitempty"`
	// The ID of the bastion host for which you want to remove hosts from the host group.
	//
	// >  You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host for which you want to remove hosts from the host group.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the call.
	Results []*RemoveHostsFromGroupResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
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
	// The return code that indicates whether the call was successful. Valid values:
	//
	// *   **OK**: The call was successful.
	// *   **UNEXPECTED**: An unknown error occurred.
	// *   **INVALID_ARGUMENT**: A request parameter is invalid.
	// *   **OBJECT_NOT_FOUND**: The specified object on which you want to perform the operation does not exist.
	// *   **OBJECT_AlREADY_EXISTS**: The specified object on which you want to perform the operation already exists.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the host group.
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// The ID of the host.
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// This parameter is deprecated.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveHostsFromGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the bastion host for which you want to remove users from the user group.
	//
	// >  You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host for which you want to remove users from the user group.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user group from which you want to remove users.
	//
	// >  You can call the [ListUserGroups](~~204509~~) operation to query the ID of the user group.
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// The ID of the user who you want to remove. The value is a JSON string. You can add up to 100 user IDs. If you specify multiple IDs, separate the IDs with commas (,).
	//
	// >  You can call the [ListUsers](~~204522~~) operation to query the IDs of users.
	UserIds *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
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
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the call.
	Results []*RemoveUsersFromGroupResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
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
	// The return code that indicates whether the call was successful. Valid values:
	//
	// *   **OK**: The call was successful.
	//
	// *   **UNEXPECTED**: An unknown error occurred.
	//
	// *   **INVALID_ARGUMENT**: A request parameter is invalid.
	//
	// > Make sure that the request parameters are valid and call the operation again.
	//
	// *   **OBJECT_NOT_FOUND**: The specified object on which you want to perform the operation does not exist.
	//
	// > Check whether the specified ID of the bastion host exists, whether the specified hosts exist, and whether the specified host IDs are valid. Then, call the operation again.
	//
	// *   **OBJECT_AlREADY_EXISTS**: The specified object on which you want to perform the operation already exists.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// This parameter is deprecated.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the group.
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	// The ID of the user.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveUsersFromGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The type of the logon credential that you want to delete. Valid values:
	//
	// *   **Password**
	// *   **PrivateKey**
	CredentialType *string `json:"CredentialType,omitempty" xml:"CredentialType,omitempty"`
	// The ID of the host account for which the logon credential is to be deleted.
	//
	// >  You can call the [ListHostAccounts](~~204372~~) operation to query the ID of the host account.
	HostAccountId *string `json:"HostAccountId,omitempty" xml:"HostAccountId,omitempty"`
	// The ID of the bastion host from which you want to delete the logon credential for the host account.
	//
	// >  You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host from which you want to delete the logon credential for the host account.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// The request ID.
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetHostAccountCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the bastion host that you want to enable.
	//
	// > You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// An array consisting of the IDs of security groups to which the bastion host is added.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// The ID of the vSwitch to which the bastion host belongs.
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
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
	// The ID of the bastion host that you enable.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The region ID of the bastion hosts to which you want to create and add tags.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// An array that consists of IDs of bastion hosts.
	//
	// Valid values: 1 to 20.
	//
	// > You can call the [DescribeInstances](~~153281~~) operation to query IDs of bastion hosts.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resource.
	//
	// Set the value to **INSTANCE**, which indicates that the resource is a bastion host.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// An array that consists of tags.
	Tag []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	// The key of tag N. Valid values of N: 1 to 20.
	//
	// >
	//
	// *   The value can be up to 128 characters in length but cannot be an empty string.
	//
	// *   The value cannot start with **aliyun** or **acs:**. The value cannot contain **http://** or **https://**.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N.\
	// Valid values of N: 1 to 20.
	//
	// >
	//
	// *   The value can be up to 128 characters in length or an empty string.
	//
	// *   The value cannot start with **aliyun** or **acs:**. The value cannot contain **http://** or **https://**.
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
	// The ID of the request, which is used to locate and troubleshoot issues.
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the bastion host to which the users to be unlocked belong.
	//
	// > You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host to which the users to be unlocked belong.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the user that you want to unlock. The value is a JSON string. You can add up to 100 user IDs. If you specify multiple IDs, separate the IDs with commas (,).
	//
	// > You can call the [ListUsers](~~204522~~) operation to query the ID of the user.
	UserIds *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
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
	// The ID of the request, which is used to locate and troubleshoot issues.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of information about the result of the call.
	Results []*UnlockUsersResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
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
	// The result of the call. Valid values:
	//
	// *   **OK**: The call was successful.
	//
	// *   **UNEXPECTED**: An unknown error occurred.
	//
	// *   **INVALID_ARGUMENT**: A request parameter is invalid.
	//
	//     **
	//
	//     **Note**Make sure that the request parameters are valid and call the operation again.
	//
	// *   **OBJECT_NOT_FOUND**: The specified object on which you want to perform the operation does not exist.
	//
	//     **
	//
	//     **Note**Check whether the specified ID of the bastion host exists, whether the specified hosts exist, and whether the specified host IDs are valid. Then, call the operation again.
	//
	// *   **OBJECT_AlREADY_EXISTS**: The specified object on which you want to perform the operation already exists.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// This parameter is deprecated.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the user.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnlockUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// Specifies whether to delete all tags that are added to the bastion host.
	//
	// *   If you specify TagKey.N, the value of this parameter can only be **false**, which indicates that only a specified tag is deleted.
	// *   If you do not specify TagKey.N and the value of this parameter is **true**, all tags are deleted. If you do not specify TagKey.N and the value of this parameter is **false**, no tags are deleted.
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The region ID of the bastion host to query.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](~~40654~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// An array that consists of IDs of bastion hosts.
	//
	// Valid values: 1 to 20.
	//
	// > You can call the [DescribeInstances](~~153281~~) operation to query the ID of the bastion host.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resource.
	//
	// Set the value to **INSTANCE**, which indicates that the resource is a bastion host.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The key of tag N.
	//
	// Valid values of N: 1 to 20.
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
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
	// The ID of the request, which is used to locate and troubleshoot issues.
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (client *Client) AcceptApproveCommandWithOptions(request *AcceptApproveCommandRequest, runtime *util.RuntimeOptions) (_result *AcceptApproveCommandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CommandId)) {
		query["CommandId"] = request.CommandId
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
		Action:      tea.String("AcceptApproveCommand"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AcceptApproveCommandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AcceptApproveCommand(request *AcceptApproveCommandRequest) (_result *AcceptApproveCommandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AcceptApproveCommandResponse{}
	_body, _err := client.AcceptApproveCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation as a Bastionhost administrator to approve an O\\&M application of an O\\&M engineer.
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request AcceptOperationTicketRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AcceptOperationTicketResponse
 */
func (client *Client) AcceptOperationTicketWithOptions(request *AcceptOperationTicketRequest, runtime *util.RuntimeOptions) (_result *AcceptOperationTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EffectCount)) {
		query["EffectCount"] = request.EffectCount
	}

	if !tea.BoolValue(util.IsUnset(request.EffectEndTime)) {
		query["EffectEndTime"] = request.EffectEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.EffectStartTime)) {
		query["EffectStartTime"] = request.EffectStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OperationTicketId)) {
		query["OperationTicketId"] = request.OperationTicketId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AcceptOperationTicket"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AcceptOperationTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation as a Bastionhost administrator to approve an O\\&M application of an O\\&M engineer.
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request AcceptOperationTicketRequest
 * @return AcceptOperationTicketResponse
 */
func (client *Client) AcceptOperationTicket(request *AcceptOperationTicketRequest) (_result *AcceptOperationTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AcceptOperationTicketResponse{}
	_body, _err := client.AcceptOperationTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to add one or more hosts to a host group. You can add multiple hosts to a host group to manage and grant permissions on the hosts in a centralized manner.
 * # Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds a limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limits when you call this operation.
 *
 * @param request AddHostsToGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AddHostsToGroupResponse
 */
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

/**
 * You can call this operation to add one or more hosts to a host group. You can add multiple hosts to a host group to manage and grant permissions on the hosts in a centralized manner.
 * # Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds a limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limits when you call this operation.
 *
 * @param request AddHostsToGroupRequest
 * @return AddHostsToGroupResponse
 */
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

/**
 * #
 * You can call this operation to add one or more users to a user group. After you call the [CreateUserGroup](~~204596~~) operation to create a user group, you can call the AddUsersToGroup operation to add multiple users to the user group. Then, you can manage and grant permissions to the users at a time.
 * # Limit
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request AddUsersToGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AddUsersToGroupResponse
 */
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

/**
 * #
 * You can call this operation to add one or more users to a user group. After you call the [CreateUserGroup](~~204596~~) operation to create a user group, you can call the AddUsersToGroup operation to add multiple users to the user group. Then, you can manage and grant permissions to the users at a time.
 * # Limit
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request AddUsersToGroupRequest
 * @return AddUsersToGroupResponse
 */
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

/**
 * After you authorize a user group to manage specific hosts and host accounts, all the users in the user group have access to the authorized hosts and host accounts.
 *
 * @param request AttachHostAccountsToUserGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AttachHostAccountsToUserGroupResponse
 */
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

/**
 * After you authorize a user group to manage specific hosts and host accounts, all the users in the user group have access to the authorized hosts and host accounts.
 *
 * @param request AttachHostAccountsToUserGroupRequest
 * @return AttachHostAccountsToUserGroupResponse
 */
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

/**
 * The ID of the request, which is used to locate and troubleshoot issues.
 *
 * @param request ConfigInstanceWhiteListRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ConfigInstanceWhiteListResponse
 */
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

/**
 * The ID of the request, which is used to locate and troubleshoot issues.
 *
 * @param request ConfigInstanceWhiteListRequest
 * @return ConfigInstanceWhiteListResponse
 */
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

	if !tea.BoolValue(util.IsUnset(request.NetworkDomainId)) {
		query["NetworkDomainId"] = request.NetworkDomainId
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

/**
 * ## Usage notes
 * You can call this operation to add a user to a bastion host. You can add local users and Resource Access Management (RAM) users. After a Bastionhost administrator adds a user to a bastion host, the O&M personnel can log on to the bastion host as the user to perform O&M operations on the host on which they have permissions.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request CreateUserRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateUserResponse
 */
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

	if !tea.BoolValue(util.IsUnset(request.EffectiveEndTime)) {
		query["EffectiveEndTime"] = request.EffectiveEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.EffectiveStartTime)) {
		query["EffectiveStartTime"] = request.EffectiveStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.LanguageStatus)) {
		query["LanguageStatus"] = request.LanguageStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		query["Mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.MobileCountryCode)) {
		query["MobileCountryCode"] = request.MobileCountryCode
	}

	if !tea.BoolValue(util.IsUnset(request.NeedResetPassword)) {
		query["NeedResetPassword"] = request.NeedResetPassword
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

	if !tea.BoolValue(util.IsUnset(request.TwoFactorMethods)) {
		query["TwoFactorMethods"] = request.TwoFactorMethods
	}

	if !tea.BoolValue(util.IsUnset(request.TwoFactorStatus)) {
		query["TwoFactorStatus"] = request.TwoFactorStatus
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

/**
 * ## Usage notes
 * You can call this operation to add a user to a bastion host. You can add local users and Resource Access Management (RAM) users. After a Bastionhost administrator adds a user to a bastion host, the O&M personnel can log on to the bastion host as the user to perform O&M operations on the host on which they have permissions.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request CreateUserRequest
 * @return CreateUserResponse
 */
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

/**
 * You can call this operation to create a user group for a bastion host as an administrator. Then, you can call the [AddUsersToGroup](~~204600~~) operation to add users to the user group at a time. After you add the users to the user group, you can authorize and manage the users in a centralized manner.
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request CreateUserGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateUserGroupResponse
 */
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

/**
 * You can call this operation to create a user group for a bastion host as an administrator. Then, you can call the [AddUsersToGroup](~~204600~~) operation to add users to the user group at a time. After you add the users to the user group, you can authorize and manage the users in a centralized manner.
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request CreateUserGroupRequest
 * @return CreateUserGroupResponse
 */
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

/**
 * You can call the CreateUserPublicKey operation to create a public key for the specified user of a bastion host.
 *
 * @param request CreateUserPublicKeyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateUserPublicKeyResponse
 */
func (client *Client) CreateUserPublicKeyWithOptions(request *CreateUserPublicKeyRequest, runtime *util.RuntimeOptions) (_result *CreateUserPublicKeyResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.PublicKey)) {
		query["PublicKey"] = request.PublicKey
	}

	if !tea.BoolValue(util.IsUnset(request.PublicKeyName)) {
		query["PublicKeyName"] = request.PublicKeyName
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
		Action:      tea.String("CreateUserPublicKey"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUserPublicKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the CreateUserPublicKey operation to create a public key for the specified user of a bastion host.
 *
 * @param request CreateUserPublicKeyRequest
 * @return CreateUserPublicKeyResponse
 */
func (client *Client) CreateUserPublicKey(request *CreateUserPublicKeyRequest) (_result *CreateUserPublicKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUserPublicKeyResponse{}
	_body, _err := client.CreateUserPublicKeyWithOptions(request, runtime)
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

/**
 * ## Usage notes
 * You can call this operation to remove a single host account. If you no longer use a host account that is added to a host in Bastionhost, you can call this operation to remove the host account from the host.
 * >  After you remove the host account, you must enter the username and password of the host when you log on to the host in Bastionhost.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DeleteHostAccountRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteHostAccountResponse
 */
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

/**
 * ## Usage notes
 * You can call this operation to remove a single host account. If you no longer use a host account that is added to a host in Bastionhost, you can call this operation to remove the host account from the host.
 * >  After you remove the host account, you must enter the username and password of the host when you log on to the host in Bastionhost.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DeleteHostAccountRequest
 * @return DeleteHostAccountResponse
 */
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

/**
 * You can call this operation to delete a single host group. If you no longer need to perform O\\&M operations on all hosts in a host group, you can call this operation to delete the host group.
 * ### Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DeleteHostGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteHostGroupResponse
 */
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

/**
 * You can call this operation to delete a single host group. If you no longer need to perform O\\&M operations on all hosts in a host group, you can call this operation to delete the host group.
 * ### Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DeleteHostGroupRequest
 * @return DeleteHostGroupResponse
 */
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

/**
 * You can call the DeleteUserPublicKey operation to delete a public key from the specified user of a bastion host.
 *
 * @param request DeleteUserPublicKeyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteUserPublicKeyResponse
 */
func (client *Client) DeleteUserPublicKeyWithOptions(request *DeleteUserPublicKeyRequest, runtime *util.RuntimeOptions) (_result *DeleteUserPublicKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PublicKeyId)) {
		query["PublicKeyId"] = request.PublicKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUserPublicKey"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUserPublicKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DeleteUserPublicKey operation to delete a public key from the specified user of a bastion host.
 *
 * @param request DeleteUserPublicKeyRequest
 * @return DeleteUserPublicKeyResponse
 */
func (client *Client) DeleteUserPublicKey(request *DeleteUserPublicKeyRequest) (_result *DeleteUserPublicKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserPublicKeyResponse{}
	_body, _err := client.DeleteUserPublicKeyWithOptions(request, runtime)
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

/**
 * ****
 *
 * @param request DetachHostGroupAccountsFromUserGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DetachHostGroupAccountsFromUserGroupResponse
 */
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

/**
 * ****
 *
 * @param request DetachHostGroupAccountsFromUserGroupRequest
 * @return DetachHostGroupAccountsFromUserGroupResponse
 */
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

/**
 * The condition that is used to filter users.
 *
 * @param request GetInstanceADAuthServerRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetInstanceADAuthServerResponse
 */
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

/**
 * The condition that is used to filter users.
 *
 * @param request GetInstanceADAuthServerRequest
 * @return GetInstanceADAuthServerResponse
 */
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

/**
 * Indicates whether two-factor authentication is enabled. Valid values:
 * *   **true**: enabled
 * *   **false**: disabled
 *
 * @param request GetInstanceTwoFactorRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetInstanceTwoFactorResponse
 */
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

/**
 * Indicates whether two-factor authentication is enabled. Valid values:
 * *   **true**: enabled
 * *   **false**: disabled
 *
 * @param request GetInstanceTwoFactorRequest
 * @return GetInstanceTwoFactorResponse
 */
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

func (client *Client) ListApproveCommandsWithOptions(request *ListApproveCommandsRequest, runtime *util.RuntimeOptions) (_result *ListApproveCommandsResponse, _err error) {
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
		Action:      tea.String("ListApproveCommands"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListApproveCommandsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListApproveCommands(request *ListApproveCommandsRequest) (_result *ListApproveCommandsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListApproveCommandsResponse{}
	_body, _err := client.ListApproveCommandsWithOptions(request, runtime)
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

func (client *Client) ListOperationTicketsWithOptions(request *ListOperationTicketsRequest, runtime *util.RuntimeOptions) (_result *ListOperationTicketsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssetAddress)) {
		query["AssetAddress"] = request.AssetAddress
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
		Action:      tea.String("ListOperationTickets"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOperationTicketsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOperationTickets(request *ListOperationTicketsRequest) (_result *ListOperationTicketsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOperationTicketsResponse{}
	_body, _err := client.ListOperationTicketsWithOptions(request, runtime)
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

func (client *Client) ListUserPublicKeysWithOptions(request *ListUserPublicKeysRequest, runtime *util.RuntimeOptions) (_result *ListUserPublicKeysResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserPublicKeys"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserPublicKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserPublicKeys(request *ListUserPublicKeysRequest) (_result *ListUserPublicKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserPublicKeysResponse{}
	_body, _err := client.ListUserPublicKeysWithOptions(request, runtime)
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

/**
 * # Description
 * You can call this operation to lock one or more users of a bastion host. If a user does not need to use a bastion host to perform O\\&M operations within a specific period of time, you can lock the user. The locked user can no longer log on to or perform O\\&M operations on the hosts on which the user is granted permissions. If you want to unlock the user later, you can call the [UnlockUsers](~~204590~~) operation.
 * # Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request LockUsersRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return LockUsersResponse
 */
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

/**
 * # Description
 * You can call this operation to lock one or more users of a bastion host. If a user does not need to use a bastion host to perform O\\&M operations within a specific period of time, you can lock the user. The locked user can no longer log on to or perform O\\&M operations on the hosts on which the user is granted permissions. If you want to unlock the user later, you can call the [UnlockUsers](~~204590~~) operation.
 * # Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request LockUsersRequest
 * @return LockUsersResponse
 */
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

/**
 * You can call the ModifyHost operation to modify the basic information about a host in a data center, an Elastic Compute Service (ECS) instance, or a host in an ApsaraDB MyBase dedicated cluster.
 * > The basic information about ECS instances and hosts in ApsaraDB MyBase dedicated clusters within your Alibaba Cloud account is synchronized to Bastionhost on a regular basis. After you modify the basic information about an ECS instance or a host in an ApsaraDB MyBase dedicated cluster, the modification result may be overwritten by the synchronized information.
 *
 * @param request ModifyHostRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyHostResponse
 */
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

	if !tea.BoolValue(util.IsUnset(request.NetworkDomainId)) {
		query["NetworkDomainId"] = request.NetworkDomainId
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

/**
 * You can call the ModifyHost operation to modify the basic information about a host in a data center, an Elastic Compute Service (ECS) instance, or a host in an ApsaraDB MyBase dedicated cluster.
 * > The basic information about ECS instances and hosts in ApsaraDB MyBase dedicated clusters within your Alibaba Cloud account is synchronized to Bastionhost on a regular basis. After you modify the basic information about an ECS instance or a host in an ApsaraDB MyBase dedicated cluster, the modification result may be overwritten by the synchronized information.
 *
 * @param request ModifyHostRequest
 * @return ModifyHostResponse
 */
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

/**
 * ## Usage notes
 * You can call this operation to change the port for the O&M protocol on one or more hosts. If the standard port for the O&M protocol on your host is vulnerable to attacks, you can call this operation to specify a custom port. For example, the standard port for SSH is port 22.
 * >  Ports 0 to 1024 are reserved for Bastionhost. Do not change the port for the O&M protocol to a reserved port.
 * ## QPS limit
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyHostsPortRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyHostsPortResponse
 */
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

/**
 * ## Usage notes
 * You can call this operation to change the port for the O&M protocol on one or more hosts. If the standard port for the O&M protocol on your host is vulnerable to attacks, you can call this operation to specify a custom port. For example, the standard port for SSH is port 22.
 * >  Ports 0 to 1024 are reserved for Bastionhost. Do not change the port for the O&M protocol to a reserved port.
 * ## QPS limit
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyHostsPortRequest
 * @return ModifyHostsPortResponse
 */
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
	if !tea.BoolValue(util.IsUnset(request.EnableTwoFactor)) {
		query["EnableTwoFactor"] = request.EnableTwoFactor
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
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

	if !tea.BoolValue(util.IsUnset(request.EffectiveEndTime)) {
		query["EffectiveEndTime"] = request.EffectiveEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.EffectiveStartTime)) {
		query["EffectiveStartTime"] = request.EffectiveStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.LanguageStatus)) {
		query["LanguageStatus"] = request.LanguageStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Mobile)) {
		query["Mobile"] = request.Mobile
	}

	if !tea.BoolValue(util.IsUnset(request.MobileCountryCode)) {
		query["MobileCountryCode"] = request.MobileCountryCode
	}

	if !tea.BoolValue(util.IsUnset(request.NeedResetPassword)) {
		query["NeedResetPassword"] = request.NeedResetPassword
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TwoFactorMethods)) {
		query["TwoFactorMethods"] = request.TwoFactorMethods
	}

	if !tea.BoolValue(util.IsUnset(request.TwoFactorStatus)) {
		query["TwoFactorStatus"] = request.TwoFactorStatus
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

func (client *Client) RejectApproveCommandWithOptions(request *RejectApproveCommandRequest, runtime *util.RuntimeOptions) (_result *RejectApproveCommandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CommandId)) {
		query["CommandId"] = request.CommandId
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
		Action:      tea.String("RejectApproveCommand"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RejectApproveCommandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RejectApproveCommand(request *RejectApproveCommandRequest) (_result *RejectApproveCommandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RejectApproveCommandResponse{}
	_body, _err := client.RejectApproveCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to reject an O\\&M application of an O\\&M engineer as a Bastionhost administrator.
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request RejectOperationTicketRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return RejectOperationTicketResponse
 */
func (client *Client) RejectOperationTicketWithOptions(request *RejectOperationTicketRequest, runtime *util.RuntimeOptions) (_result *RejectOperationTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OperationTicketId)) {
		query["OperationTicketId"] = request.OperationTicketId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RejectOperationTicket"),
		Version:     tea.String("2019-12-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RejectOperationTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to reject an O\\&M application of an O\\&M engineer as a Bastionhost administrator.
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request RejectOperationTicketRequest
 * @return RejectOperationTicketResponse
 */
func (client *Client) RejectOperationTicket(request *RejectOperationTicketRequest) (_result *RejectOperationTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RejectOperationTicketResponse{}
	_body, _err := client.RejectOperationTicketWithOptions(request, runtime)
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

/**
 * You can call this operation to remove one or more users from a user group. When users in a user group are transferred to a new position, resign, or are switched to another user group, you can call this operation to remove the users from the current user group at a time.
 * ## QPS limit
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request RemoveUsersFromGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return RemoveUsersFromGroupResponse
 */
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

/**
 * You can call this operation to remove one or more users from a user group. When users in a user group are transferred to a new position, resign, or are switched to another user group, you can call this operation to remove the users from the current user group at a time.
 * ## QPS limit
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request RemoveUsersFromGroupRequest
 * @return RemoveUsersFromGroupResponse
 */
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

/**
 * After you call the [LockUsers](~~204591~~) operation to lock one or more users of a bastion host, you can call this operation to unlock the users. After the users are unlocked, the users can perform O\\&M operations by using the bastion host.
 * # Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request UnlockUsersRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UnlockUsersResponse
 */
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

/**
 * After you call the [LockUsers](~~204591~~) operation to lock one or more users of a bastion host, you can call this operation to unlock the users. After the users are unlocked, the users can perform O\\&M operations by using the bastion host.
 * # Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request UnlockUsersRequest
 * @return UnlockUsersResponse
 */
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
