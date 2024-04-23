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

type AddAddressBookRequest struct {
	// The addresses that you want to add to the address book. Separate multiple addresses with commas (,).
	//
	// >  If you set GroupType to `ip`, `port`, or `domain`, you must specify the AddressList parameter.
	//
	// *   If you set GroupType to `ip`, you must add IP addresses to the address book. Example: 192.0.XX.XX/32, 192.0.XX.XX/24.
	// *   If you set GroupType to `port`, you must add port numbers or port ranges to the address book. Example: 80, 100/200.
	// *   If you set GroupType to `domain`, you must add domain names to the address book. Example: example.com, aliyundoc.com.
	AddressList *string `json:"AddressList,omitempty" xml:"AddressList,omitempty"`
	// Specifies whether to automatically add public IP addresses of ECS instances to the address book if the instances match the specified tags. Valid values:
	//
	// *   **1**: yes
	// *   **0** (default): no
	AutoAddTagEcs *string `json:"AutoAddTagEcs,omitempty" xml:"AutoAddTagEcs,omitempty"`
	// The description of the address book.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the address book.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the address book. Valid values:
	//
	// *   **ip**: IP address book
	// *   **domain**: domain address book
	// *   **port**: port address book
	// *   **tag**: ECS tag-based address book
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// *   **zh** (default): Chinese
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Deprecated
	// The source IP address of the request.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The ECS tags that you want to match.
	TagList []*AddAddressBookRequestTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The logical relation among the ECS tags that you want to match. Valid values:
	//
	// *   **and** (default): Only the public IP addresses of ECS instances that match all the specified tags can be added to the address book.
	// *   **or**: The public IP addresses of ECS instances that match one of the specified tags can be added to the address book.
	TagRelation *string `json:"TagRelation,omitempty" xml:"TagRelation,omitempty"`
}

func (s AddAddressBookRequest) String() string {
	return tea.Prettify(s)
}

func (s AddAddressBookRequest) GoString() string {
	return s.String()
}

func (s *AddAddressBookRequest) SetAddressList(v string) *AddAddressBookRequest {
	s.AddressList = &v
	return s
}

func (s *AddAddressBookRequest) SetAutoAddTagEcs(v string) *AddAddressBookRequest {
	s.AutoAddTagEcs = &v
	return s
}

func (s *AddAddressBookRequest) SetDescription(v string) *AddAddressBookRequest {
	s.Description = &v
	return s
}

func (s *AddAddressBookRequest) SetGroupName(v string) *AddAddressBookRequest {
	s.GroupName = &v
	return s
}

func (s *AddAddressBookRequest) SetGroupType(v string) *AddAddressBookRequest {
	s.GroupType = &v
	return s
}

func (s *AddAddressBookRequest) SetLang(v string) *AddAddressBookRequest {
	s.Lang = &v
	return s
}

func (s *AddAddressBookRequest) SetSourceIp(v string) *AddAddressBookRequest {
	s.SourceIp = &v
	return s
}

func (s *AddAddressBookRequest) SetTagList(v []*AddAddressBookRequestTagList) *AddAddressBookRequest {
	s.TagList = v
	return s
}

func (s *AddAddressBookRequest) SetTagRelation(v string) *AddAddressBookRequest {
	s.TagRelation = &v
	return s
}

type AddAddressBookRequestTagList struct {
	// The key of the ECS tag.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the ECS tag.
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s AddAddressBookRequestTagList) String() string {
	return tea.Prettify(s)
}

func (s AddAddressBookRequestTagList) GoString() string {
	return s.String()
}

func (s *AddAddressBookRequestTagList) SetTagKey(v string) *AddAddressBookRequestTagList {
	s.TagKey = &v
	return s
}

func (s *AddAddressBookRequestTagList) SetTagValue(v string) *AddAddressBookRequestTagList {
	s.TagValue = &v
	return s
}

type AddAddressBookResponseBody struct {
	// The UUID of the returned address book.
	GroupUuid *string `json:"GroupUuid,omitempty" xml:"GroupUuid,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddAddressBookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddAddressBookResponseBody) GoString() string {
	return s.String()
}

func (s *AddAddressBookResponseBody) SetGroupUuid(v string) *AddAddressBookResponseBody {
	s.GroupUuid = &v
	return s
}

func (s *AddAddressBookResponseBody) SetRequestId(v string) *AddAddressBookResponseBody {
	s.RequestId = &v
	return s
}

type AddAddressBookResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddAddressBookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddAddressBookResponse) String() string {
	return tea.Prettify(s)
}

func (s AddAddressBookResponse) GoString() string {
	return s.String()
}

func (s *AddAddressBookResponse) SetHeaders(v map[string]*string) *AddAddressBookResponse {
	s.Headers = v
	return s
}

func (s *AddAddressBookResponse) SetStatusCode(v int32) *AddAddressBookResponse {
	s.StatusCode = &v
	return s
}

func (s *AddAddressBookResponse) SetBody(v *AddAddressBookResponseBody) *AddAddressBookResponse {
	s.Body = v
	return s
}

type AddControlPolicyRequest struct {
	// The action that Cloud Firewall performs on the traffic. Valid values:
	//
	// *   **accept**: allows the traffic.
	// *   **drop**: denies the traffic.
	// *   **log**: monitors the traffic.
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The application type supported by the access control policy. Valid values:
	//
	// *   **FTP**
	// *   **HTTP**
	// *   **HTTPS**
	// *   **Memcache**
	// *   **MongoDB**
	// *   **MQTT**
	// *   **MySQL**
	// *   **RDP**
	// *   **Redis**
	// *   **SMTP**
	// *   **SMTPS**
	// *   **SSH**
	// *   **SSL_No_Cert**
	// *   **SSL**
	// *   **VNC**
	// *   **ANY**
	//
	// > The value of this parameter is based on the value of Proto. If Proto is set to TCP, you can set ApplicationName to any valid value. If Proto is set to UDP, ICMP, or ANY, you can set ApplicationName only to ANY. You must specify at least one of the ApplicationNameList and ApplicationName parameters.
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The application types supported by the access control policy.
	ApplicationNameList []*string `json:"ApplicationNameList,omitempty" xml:"ApplicationNameList,omitempty" type:"Repeated"`
	// The description of the access control policy.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination port in the access control policy. Valid values:
	//
	// *   If Proto is set to ICMP, DestPort is automatically left empty.
	//
	// > If Proto is set to ICMP, access control does not take effect on the destination port.
	//
	// *   If Proto is set to TCP, UDP, or ANY and DestPortType is set to group, DestPort is empty.
	//
	// > If DestPortType is set to group, you do not need to specify the destination port number. All ports on which the access control policy takes effect are included in the destination port address book.
	//
	// *   If Proto is set to TCP, UDP, or ANY and DestPortType is set to port, the value of DestPort is the destination port number.
	DestPort *string `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	// The name of the destination port address book in the access control policy.
	//
	// > If DestPortType is set to group, you must specify the name of the destination port address book.
	DestPortGroup *string `json:"DestPortGroup,omitempty" xml:"DestPortGroup,omitempty"`
	// The type of the destination port in the access control policy.
	//
	// Valid values:
	//
	// *   **port**: port
	// *   **group**: port address book
	DestPortType *string `json:"DestPortType,omitempty" xml:"DestPortType,omitempty"`
	// The destination address in the access control policy.
	//
	// Valid values:
	//
	// *   If DestinationType is set to net, the value of this parameter is a CIDR block.
	//
	//     Example: 1.2.XX.XX/24
	//
	// *   If DestinationType is set to group, the value of this parameter is an address book name.
	//
	//     Example: db_group
	//
	// *   If DestinationType is set to domain, the value of this parameter is a domain name.
	//
	//     Example: \*.aliyuncs.com
	//
	// *   If DestinationType is set to location, the value of this parameter is a location.
	//
	//     Example: \["BJ11", "ZB"]
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The type of the destination address in the access control policy. Valid values:
	//
	// *   **net**: CIDR block
	// *   **group**: address book
	// *   **domain**: domain name
	// *   **location**: location
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The direction of the traffic to which the access control policy applies. Valid values:
	//
	// *   **in**: inbound traffic
	// *   **out**: outbound traffic
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The time when the access control policy stops taking effect. The value is a UNIX timestamp. Unit: seconds. The value must be on the hour or on the half hour, and at least 30 minutes later than the start time.
	//
	// >  If you set RepeatType to Permanent, leave this parameter empty. If you set RepeatType to None, Daily, Weekly, or Monthly, you must specify this parameter.
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The IP version supported by the access control policy.
	//
	// Valid values:
	//
	// *   **4**: IPv4
	// *   **6**: IPv6
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The priority of the access control policy. The priority value starts from 1. A smaller priority value indicates a higher priority.
	NewOrder *string `json:"NewOrder,omitempty" xml:"NewOrder,omitempty"`
	// The protocol type supported by the access control policy. Valid values:
	//
	// *   **ANY**
	// *   **TCP**
	// *   **UDP**
	// *   **ICMP**
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// Specifies whether to enable the access control policy. By default, an access control policy is enabled after the policy is created. Valid values:
	//
	// *   **true**: enables the access control policy.
	// *   **false**: disables the access control policy.
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The days of a week or of a month on which the access control policy takes effect.
	//
	// *   If you set RepeatType to `Permanent`, `None`, or `Daily`, leave this parameter empty. Example: \[].
	// *   If you set RepeatType to Weekly, you must specify this parameter. Example: \[0, 6].
	//
	// >  If you set RepeatType to Weekly, the fields in the value of this parameter cannot be repeated.
	//
	// *   If you set RepeatType to `Monthly`, you must specify this parameter. Example: \[1, 31].
	//
	// >  If you set RepeatType to Monthly, the fields in the value of this parameter cannot be repeated.
	RepeatDays []*int64 `json:"RepeatDays,omitempty" xml:"RepeatDays,omitempty" type:"Repeated"`
	// The point in time when the recurrence ends. Example: 23:30. The end time must be on the hour or on the half hour, and at least 30 minutes later than the start time.
	//
	// >  If you set RepeatType to Permanent or None, leave this parameter empty. If you set RepeatType to Daily, Weekly, or Monthly, you must specify this parameter.
	RepeatEndTime *string `json:"RepeatEndTime,omitempty" xml:"RepeatEndTime,omitempty"`
	// The point in time when the recurrence starts. Example: 08:00. The start time must be on the hour or on the half hour, and at least 30 minutes earlier than the end time.
	//
	// >  If you set RepeatType to Permanent or None, leave this parameter empty. If you set RepeatType to Daily, Weekly, or Monthly, you must specify this parameter.
	RepeatStartTime *string `json:"RepeatStartTime,omitempty" xml:"RepeatStartTime,omitempty"`
	// The recurrence type for the access control policy to take effect. Valid values:
	//
	// *   **Permanent** (default): The policy always takes effect.
	// *   **None**: The policy takes effect for only once.
	// *   **Daily**: The policy takes effect on a daily basis.
	// *   **Weekly**: The policy takes effect on a weekly basis.
	// *   **Monthly**: The policy takes effect on a monthly basis.
	RepeatType *string `json:"RepeatType,omitempty" xml:"RepeatType,omitempty"`
	// The source address in the access control policy. Valid values:
	//
	// *   If SourceType is set to net, the value of this parameter is a CIDR block.
	//
	//     Example: 1.1.XX.XX/24
	//
	// *   If SourceType is set to group, the value of this parameter is an address book name.
	//
	//     Example: db_group
	//
	// *   If SourceType is set to location, the value of this parameter is a location.
	//
	//     Example: \["BJ11", "ZB"]
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// Deprecated
	// The source IP address of the request.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The type of the source address in the access control policy. Valid values:
	//
	// *   **net**: CIDR block
	// *   **group**: address book
	// *   **location**: location
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The time when the access control policy starts to take effect. The value is a UNIX timestamp. Unit: seconds. The value must be on the hour or on the half hour, and at least 30 minutes earlier than the end time.
	//
	// >  If you set RepeatType to Permanent, leave this parameter empty. If you set RepeatType to None, Daily, Weekly, or Monthly, you must specify this parameter.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s AddControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s AddControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *AddControlPolicyRequest) SetAclAction(v string) *AddControlPolicyRequest {
	s.AclAction = &v
	return s
}

func (s *AddControlPolicyRequest) SetApplicationName(v string) *AddControlPolicyRequest {
	s.ApplicationName = &v
	return s
}

func (s *AddControlPolicyRequest) SetApplicationNameList(v []*string) *AddControlPolicyRequest {
	s.ApplicationNameList = v
	return s
}

func (s *AddControlPolicyRequest) SetDescription(v string) *AddControlPolicyRequest {
	s.Description = &v
	return s
}

func (s *AddControlPolicyRequest) SetDestPort(v string) *AddControlPolicyRequest {
	s.DestPort = &v
	return s
}

func (s *AddControlPolicyRequest) SetDestPortGroup(v string) *AddControlPolicyRequest {
	s.DestPortGroup = &v
	return s
}

func (s *AddControlPolicyRequest) SetDestPortType(v string) *AddControlPolicyRequest {
	s.DestPortType = &v
	return s
}

func (s *AddControlPolicyRequest) SetDestination(v string) *AddControlPolicyRequest {
	s.Destination = &v
	return s
}

func (s *AddControlPolicyRequest) SetDestinationType(v string) *AddControlPolicyRequest {
	s.DestinationType = &v
	return s
}

func (s *AddControlPolicyRequest) SetDirection(v string) *AddControlPolicyRequest {
	s.Direction = &v
	return s
}

func (s *AddControlPolicyRequest) SetEndTime(v int64) *AddControlPolicyRequest {
	s.EndTime = &v
	return s
}

func (s *AddControlPolicyRequest) SetIpVersion(v string) *AddControlPolicyRequest {
	s.IpVersion = &v
	return s
}

func (s *AddControlPolicyRequest) SetLang(v string) *AddControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *AddControlPolicyRequest) SetNewOrder(v string) *AddControlPolicyRequest {
	s.NewOrder = &v
	return s
}

func (s *AddControlPolicyRequest) SetProto(v string) *AddControlPolicyRequest {
	s.Proto = &v
	return s
}

func (s *AddControlPolicyRequest) SetRelease(v string) *AddControlPolicyRequest {
	s.Release = &v
	return s
}

func (s *AddControlPolicyRequest) SetRepeatDays(v []*int64) *AddControlPolicyRequest {
	s.RepeatDays = v
	return s
}

func (s *AddControlPolicyRequest) SetRepeatEndTime(v string) *AddControlPolicyRequest {
	s.RepeatEndTime = &v
	return s
}

func (s *AddControlPolicyRequest) SetRepeatStartTime(v string) *AddControlPolicyRequest {
	s.RepeatStartTime = &v
	return s
}

func (s *AddControlPolicyRequest) SetRepeatType(v string) *AddControlPolicyRequest {
	s.RepeatType = &v
	return s
}

func (s *AddControlPolicyRequest) SetSource(v string) *AddControlPolicyRequest {
	s.Source = &v
	return s
}

func (s *AddControlPolicyRequest) SetSourceIp(v string) *AddControlPolicyRequest {
	s.SourceIp = &v
	return s
}

func (s *AddControlPolicyRequest) SetSourceType(v string) *AddControlPolicyRequest {
	s.SourceType = &v
	return s
}

func (s *AddControlPolicyRequest) SetStartTime(v int64) *AddControlPolicyRequest {
	s.StartTime = &v
	return s
}

type AddControlPolicyResponseBody struct {
	// The ID of the access control policy that is created on the Internet firewall.
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *AddControlPolicyResponseBody) SetAclUuid(v string) *AddControlPolicyResponseBody {
	s.AclUuid = &v
	return s
}

func (s *AddControlPolicyResponseBody) SetRequestId(v string) *AddControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

type AddControlPolicyResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s AddControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *AddControlPolicyResponse) SetHeaders(v map[string]*string) *AddControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *AddControlPolicyResponse) SetStatusCode(v int32) *AddControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *AddControlPolicyResponse) SetBody(v *AddControlPolicyResponseBody) *AddControlPolicyResponse {
	s.Body = v
	return s
}

type AddInstanceMembersRequest struct {
	// The members that you want to add to Cloud Firewall.
	Members []*AddInstanceMembersRequestMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
}

func (s AddInstanceMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s AddInstanceMembersRequest) GoString() string {
	return s.String()
}

func (s *AddInstanceMembersRequest) SetMembers(v []*AddInstanceMembersRequestMembers) *AddInstanceMembersRequest {
	s.Members = v
	return s
}

type AddInstanceMembersRequestMembers struct {
	// The remarks of member that you want to add to Cloud Firewall. The remarks must be 1 to 256 characters in length.
	MemberDesc *string `json:"MemberDesc,omitempty" xml:"MemberDesc,omitempty"`
	// The UID of member that you want to add to Cloud Firewall.
	MemberUid *int64 `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
}

func (s AddInstanceMembersRequestMembers) String() string {
	return tea.Prettify(s)
}

func (s AddInstanceMembersRequestMembers) GoString() string {
	return s.String()
}

func (s *AddInstanceMembersRequestMembers) SetMemberDesc(v string) *AddInstanceMembersRequestMembers {
	s.MemberDesc = &v
	return s
}

func (s *AddInstanceMembersRequestMembers) SetMemberUid(v int64) *AddInstanceMembersRequestMembers {
	s.MemberUid = &v
	return s
}

type AddInstanceMembersResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddInstanceMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddInstanceMembersResponseBody) GoString() string {
	return s.String()
}

func (s *AddInstanceMembersResponseBody) SetRequestId(v string) *AddInstanceMembersResponseBody {
	s.RequestId = &v
	return s
}

type AddInstanceMembersResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddInstanceMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddInstanceMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s AddInstanceMembersResponse) GoString() string {
	return s.String()
}

func (s *AddInstanceMembersResponse) SetHeaders(v map[string]*string) *AddInstanceMembersResponse {
	s.Headers = v
	return s
}

func (s *AddInstanceMembersResponse) SetStatusCode(v int32) *AddInstanceMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *AddInstanceMembersResponse) SetBody(v *AddInstanceMembersResponseBody) *AddInstanceMembersResponse {
	s.Body = v
	return s
}

type BatchCopyVpcFirewallControlPolicyRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Deprecated
	// The source IP address of the request.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The ID of the policy group of the source VPC firewall. Valid values:
	//
	// *   If the VPC firewall protects mutual access traffic between a VPC and a specified network instance that is attached to a Cloud Enterprise Network (CEN) instance, the value of this parameter is the ID of the CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance.
	// *   If the VPC firewall protects traffic between two VPCs that are connected by using an Express Connect circuit, the value of this parameter is the instance ID of the VPC firewall.
	//
	// >  You can call the [DescribeVpcFirewallAclGroupList](~~159760~~) operation to query the IDs of policy groups.
	SourceVpcFirewallId *string `json:"SourceVpcFirewallId,omitempty" xml:"SourceVpcFirewallId,omitempty"`
	// The ID of the policy group of the destination VPC firewall. Valid values:
	//
	// *   If the VPC firewall protects mutual access traffic between a VPC and a specified network instance that is attached to a CEN instance, the value of this parameter is the ID of the CEN instance. The network instance can be a VPC, a VBR, or a CCN instance.
	// *   If the VPC firewall protects traffic between two VPCs that are connected by using an Express Connect circuit, the value of this parameter is the instance ID of the VPC firewall.
	//
	// >  You can call the [DescribeVpcFirewallAclGroupList](~~159760~~) operation to query the IDs of policy groups.
	TargetVpcFirewallId *string `json:"TargetVpcFirewallId,omitempty" xml:"TargetVpcFirewallId,omitempty"`
}

func (s BatchCopyVpcFirewallControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchCopyVpcFirewallControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *BatchCopyVpcFirewallControlPolicyRequest) SetLang(v string) *BatchCopyVpcFirewallControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *BatchCopyVpcFirewallControlPolicyRequest) SetSourceIp(v string) *BatchCopyVpcFirewallControlPolicyRequest {
	s.SourceIp = &v
	return s
}

func (s *BatchCopyVpcFirewallControlPolicyRequest) SetSourceVpcFirewallId(v string) *BatchCopyVpcFirewallControlPolicyRequest {
	s.SourceVpcFirewallId = &v
	return s
}

func (s *BatchCopyVpcFirewallControlPolicyRequest) SetTargetVpcFirewallId(v string) *BatchCopyVpcFirewallControlPolicyRequest {
	s.TargetVpcFirewallId = &v
	return s
}

type BatchCopyVpcFirewallControlPolicyResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchCopyVpcFirewallControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchCopyVpcFirewallControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *BatchCopyVpcFirewallControlPolicyResponseBody) SetRequestId(v string) *BatchCopyVpcFirewallControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

type BatchCopyVpcFirewallControlPolicyResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchCopyVpcFirewallControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchCopyVpcFirewallControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchCopyVpcFirewallControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *BatchCopyVpcFirewallControlPolicyResponse) SetHeaders(v map[string]*string) *BatchCopyVpcFirewallControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *BatchCopyVpcFirewallControlPolicyResponse) SetStatusCode(v int32) *BatchCopyVpcFirewallControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchCopyVpcFirewallControlPolicyResponse) SetBody(v *BatchCopyVpcFirewallControlPolicyResponseBody) *BatchCopyVpcFirewallControlPolicyResponse {
	s.Body = v
	return s
}

type BatchDeleteVpcFirewallControlPolicyRequest struct {
	// The UUIDs of access control policies.
	AclUuidList []*string `json:"AclUuidList,omitempty" xml:"AclUuidList,omitempty" type:"Repeated"`
	// The instance ID of the VPC firewall.
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s BatchDeleteVpcFirewallControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteVpcFirewallControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteVpcFirewallControlPolicyRequest) SetAclUuidList(v []*string) *BatchDeleteVpcFirewallControlPolicyRequest {
	s.AclUuidList = v
	return s
}

func (s *BatchDeleteVpcFirewallControlPolicyRequest) SetVpcFirewallId(v string) *BatchDeleteVpcFirewallControlPolicyRequest {
	s.VpcFirewallId = &v
	return s
}

type BatchDeleteVpcFirewallControlPolicyResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchDeleteVpcFirewallControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteVpcFirewallControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteVpcFirewallControlPolicyResponseBody) SetRequestId(v string) *BatchDeleteVpcFirewallControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

type BatchDeleteVpcFirewallControlPolicyResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeleteVpcFirewallControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDeleteVpcFirewallControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteVpcFirewallControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteVpcFirewallControlPolicyResponse) SetHeaders(v map[string]*string) *BatchDeleteVpcFirewallControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteVpcFirewallControlPolicyResponse) SetStatusCode(v int32) *BatchDeleteVpcFirewallControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteVpcFirewallControlPolicyResponse) SetBody(v *BatchDeleteVpcFirewallControlPolicyResponseBody) *BatchDeleteVpcFirewallControlPolicyResponse {
	s.Body = v
	return s
}

type CreateDownloadTaskRequest struct {
	// The language of the content within the response.
	//
	// Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The query condition of the download task.
	TaskData *string `json:"TaskData,omitempty" xml:"TaskData,omitempty"`
}

func (s CreateDownloadTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDownloadTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDownloadTaskRequest) SetLang(v string) *CreateDownloadTaskRequest {
	s.Lang = &v
	return s
}

func (s *CreateDownloadTaskRequest) SetTaskData(v string) *CreateDownloadTaskRequest {
	s.TaskData = &v
	return s
}

type CreateDownloadTaskResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the task. Valid values:
	//
	// finish: You can query the task to obtain the download link of the file.
	//
	// start
	//
	// error
	//
	// expire: The task file is invalid and cannot be downloaded.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The unique ID of the task.
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the file download task.
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s CreateDownloadTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDownloadTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDownloadTaskResponseBody) SetRequestId(v string) *CreateDownloadTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDownloadTaskResponseBody) SetStatus(v string) *CreateDownloadTaskResponseBody {
	s.Status = &v
	return s
}

func (s *CreateDownloadTaskResponseBody) SetTaskId(v int64) *CreateDownloadTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateDownloadTaskResponseBody) SetTaskName(v string) *CreateDownloadTaskResponseBody {
	s.TaskName = &v
	return s
}

type CreateDownloadTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDownloadTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDownloadTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDownloadTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateDownloadTaskResponse) SetHeaders(v map[string]*string) *CreateDownloadTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateDownloadTaskResponse) SetStatusCode(v int32) *CreateDownloadTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDownloadTaskResponse) SetBody(v *CreateDownloadTaskResponseBody) *CreateDownloadTaskResponse {
	s.Body = v
	return s
}

type CreateNatFirewallControlPolicyRequest struct {
	// The action that Cloud Firewall performs on the traffic.
	//
	// Valid values:
	//
	// *   **accept**: allows the traffic.
	// *   **drop**: denies the traffic.
	// *   **log**: monitors the traffic.
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The application types supported by the access control policy.
	ApplicationNameList []*string `json:"ApplicationNameList,omitempty" xml:"ApplicationNameList,omitempty" type:"Repeated"`
	// The description of the access control policy.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination port in the access control policy. Valid values:
	//
	// *   If Proto is set to ICMP, DestPort is automatically left empty.
	//
	// > If Proto is set to ICMP, access control does not take effect on the destination port.
	//
	// *   If Proto is set to TCP, UDP, or ANY and DestPortType is set to group, DestPort is empty.
	//
	// > If DestPortType is set to group, you do not need to specify the destination port number. All ports on which the access control policy takes effect are included in the destination port address book.
	//
	// *   If Proto is set to TCP, UDP, or ANY and DestPortType is set to port, the value of DestPort is the destination port number.
	DestPort *string `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	// The name of the destination port address book in the access control policy.
	//
	// > If DestPortType is set to group, you must specify the name of the destination port address book.
	DestPortGroup *string `json:"DestPortGroup,omitempty" xml:"DestPortGroup,omitempty"`
	// The type of the destination port in the access control policy. Valid values:
	//
	// *   **port**: port
	// *   **group**: port address book
	DestPortType *string `json:"DestPortType,omitempty" xml:"DestPortType,omitempty"`
	// The destination address in the access control policy.
	//
	// Valid values:
	//
	// *   If DestinationType is set to net, the value of this parameter is a CIDR block.
	//
	//     Example: 1.2.XX.XX/24
	//
	// *   If DestinationType is set to group, the value of this parameter is an address book.
	//
	//     Example: db_group
	//
	// *   If DestinationType is set to domain, the value of this parameter is a domain name.
	//
	//     Example: \*.aliyuncs.com
	//
	// *   If DestinationType is set to location, the value of this parameter is a location.
	//
	//     Example: \["BJ11", "ZB"]
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The type of the destination address in the access control policy.
	//
	// Valid values:
	//
	// *   **net**: CIDR block
	// *   **group**: address book
	// *   **domain**: domain name
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The direction of the traffic to which the access control policy applies. Valid values:
	//
	// *   **out**: outbound traffic
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The domain name resolution method of the access control policy. Valid values:
	//
	// *   **0**: fully qualified domain name (FQDN)-based resolution
	// *   **1**: Domain Name System (DNS)-based dynamic resolution
	// *   **2**: FQDN and DNS-based dynamic resolution
	DomainResolveType *int32 `json:"DomainResolveType,omitempty" xml:"DomainResolveType,omitempty"`
	// The time when the access control policy stops taking effect. The value is a UNIX timestamp. Unit: seconds. The value must be on the hour or on the half hour, and at least 30 minutes later than the value of StartTime.
	//
	// >  If RepeatType is set to Permanent, EndTime is left empty. If RepeatType is set to None, Daily, Weekly, or Monthly, this parameter must be specified.
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The IP version supported by the access control policy. Valid values:
	//
	// *   **4**: IPv4 (default)
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The language of the content within the response.
	//
	// Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the NAT gateway.
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The priority of the access control policy. The priority value starts from 1. A small priority value indicates a high priority.
	NewOrder *string `json:"NewOrder,omitempty" xml:"NewOrder,omitempty"`
	// The protocol type in the access control policy.
	//
	// Valid values:
	//
	// *   ANY: all types of protocols
	// *   TCP
	// *   UDP
	// *   ICMP
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// Specifies whether to enable the access control policy. By default, an access control policy is enabled after it is created. Valid values:
	//
	// *   **true**
	// *   **false**
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The days of a week or of a month on which the access control policy takes effect.
	//
	// *   If RepeatType is set to `Permanent`, `None`, or `Daily`, RepeatDays is left empty. Example: \[].
	// *   If RepeatType is set to Weekly, RepeatDays must be specified. Example: \[0, 6].
	//
	// >  If RepeatType is set to Weekly, the fields in the value of RepeatDays cannot be repeated.
	//
	// *   If RepeatType is set to `Monthly`, RepeatDays must be specified. Example: \[1, 31].
	//
	// >  If RepeatType is set to Monthly, the fields in the value of RepeatDays cannot be repeated.
	RepeatDays []*int64 `json:"RepeatDays,omitempty" xml:"RepeatDays,omitempty" type:"Repeated"`
	// The point in time when the recurrence ends. Example: 23:30. The value must be on the hour or on the half hour, and at least 30 minutes later than the value of RepeatStartTime.
	//
	// >  If RepeatType is set to Permanent or None, RepeatEndTime is left empty. If RepeatType is set to Daily, Weekly, or Monthly, this parameter must be specified.
	RepeatEndTime *string `json:"RepeatEndTime,omitempty" xml:"RepeatEndTime,omitempty"`
	// The point in time when the recurrence starts. Example: 08:00. The value must be on the hour or on the half hour, and at least 30 minutes earlier than the value of RepeatEndTime.
	//
	// >  If RepeatType is set to Permanent or None, RepeatStartTime is left empty. If RepeatType is set to Daily, Weekly, or Monthly, this parameter must be specified.
	RepeatStartTime *string `json:"RepeatStartTime,omitempty" xml:"RepeatStartTime,omitempty"`
	// The recurrence type for the access control policy to take effect. Valid values:
	//
	// *   **Permanent** (default): The policy always takes effect.
	// *   **None**: The policy takes effect for only once.
	// *   **Daily**: The policy takes effect on a daily basis.
	// *   **Weekly**: The policy takes effect on a weekly basis.
	// *   **Monthly**: The policy takes effect on a monthly basis.
	RepeatType *string `json:"RepeatType,omitempty" xml:"RepeatType,omitempty"`
	// The source address in the access control policy.
	//
	// Valid values:
	//
	// *   If **SourceType** is set to `net`, the value of Source is a CIDR block.
	//
	//     Example: 10.2.4.0/24
	//
	// *   If **SourceType** is set to `group`, the value of this parameter must be an address book name.
	//
	//     Example: db_group
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The type of the source address in the access control policy.
	//
	// Valid values:
	//
	// *   **net**: source CIDR block
	// *   **group**: source address book
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The time when the access control policy starts to take effect. The value is a UNIX timestamp. Unit: seconds. The value must be on the hour or on the half hour, and at least 30 minutes earlier than the value of EndTime.
	//
	// >  If RepeatType is set to Permanent, StartTime is left empty. If RepeatType is set to None, Daily, Weekly, or Monthly, this parameter must be specified.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateNatFirewallControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNatFirewallControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateNatFirewallControlPolicyRequest) SetAclAction(v string) *CreateNatFirewallControlPolicyRequest {
	s.AclAction = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetApplicationNameList(v []*string) *CreateNatFirewallControlPolicyRequest {
	s.ApplicationNameList = v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetDescription(v string) *CreateNatFirewallControlPolicyRequest {
	s.Description = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetDestPort(v string) *CreateNatFirewallControlPolicyRequest {
	s.DestPort = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetDestPortGroup(v string) *CreateNatFirewallControlPolicyRequest {
	s.DestPortGroup = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetDestPortType(v string) *CreateNatFirewallControlPolicyRequest {
	s.DestPortType = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetDestination(v string) *CreateNatFirewallControlPolicyRequest {
	s.Destination = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetDestinationType(v string) *CreateNatFirewallControlPolicyRequest {
	s.DestinationType = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetDirection(v string) *CreateNatFirewallControlPolicyRequest {
	s.Direction = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetDomainResolveType(v int32) *CreateNatFirewallControlPolicyRequest {
	s.DomainResolveType = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetEndTime(v int64) *CreateNatFirewallControlPolicyRequest {
	s.EndTime = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetIpVersion(v string) *CreateNatFirewallControlPolicyRequest {
	s.IpVersion = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetLang(v string) *CreateNatFirewallControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetNatGatewayId(v string) *CreateNatFirewallControlPolicyRequest {
	s.NatGatewayId = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetNewOrder(v string) *CreateNatFirewallControlPolicyRequest {
	s.NewOrder = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetProto(v string) *CreateNatFirewallControlPolicyRequest {
	s.Proto = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetRelease(v string) *CreateNatFirewallControlPolicyRequest {
	s.Release = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetRepeatDays(v []*int64) *CreateNatFirewallControlPolicyRequest {
	s.RepeatDays = v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetRepeatEndTime(v string) *CreateNatFirewallControlPolicyRequest {
	s.RepeatEndTime = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetRepeatStartTime(v string) *CreateNatFirewallControlPolicyRequest {
	s.RepeatStartTime = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetRepeatType(v string) *CreateNatFirewallControlPolicyRequest {
	s.RepeatType = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetSource(v string) *CreateNatFirewallControlPolicyRequest {
	s.Source = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetSourceType(v string) *CreateNatFirewallControlPolicyRequest {
	s.SourceType = &v
	return s
}

func (s *CreateNatFirewallControlPolicyRequest) SetStartTime(v int64) *CreateNatFirewallControlPolicyRequest {
	s.StartTime = &v
	return s
}

type CreateNatFirewallControlPolicyResponseBody struct {
	// The unique ID of the access control policy.
	//
	// >  To modify an access control policy, you must specify the unique ID of the policy. You can call the DescribeNatFirewallControlPolicy operation to obtain the ID.
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNatFirewallControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateNatFirewallControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNatFirewallControlPolicyResponseBody) SetAclUuid(v string) *CreateNatFirewallControlPolicyResponseBody {
	s.AclUuid = &v
	return s
}

func (s *CreateNatFirewallControlPolicyResponseBody) SetRequestId(v string) *CreateNatFirewallControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

type CreateNatFirewallControlPolicyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNatFirewallControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNatFirewallControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateNatFirewallControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateNatFirewallControlPolicyResponse) SetHeaders(v map[string]*string) *CreateNatFirewallControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateNatFirewallControlPolicyResponse) SetStatusCode(v int32) *CreateNatFirewallControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNatFirewallControlPolicyResponse) SetBody(v *CreateNatFirewallControlPolicyResponseBody) *CreateNatFirewallControlPolicyResponse {
	s.Body = v
	return s
}

type CreateSecurityProxyRequest struct {
	FirewallSwitch    *string                                        `json:"FirewallSwitch,omitempty" xml:"FirewallSwitch,omitempty"`
	Lang              *string                                        `json:"Lang,omitempty" xml:"Lang,omitempty"`
	NatGatewayId      *string                                        `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	NatRouteEntryList []*CreateSecurityProxyRequestNatRouteEntryList `json:"NatRouteEntryList,omitempty" xml:"NatRouteEntryList,omitempty" type:"Repeated"`
	ProxyName         *string                                        `json:"ProxyName,omitempty" xml:"ProxyName,omitempty"`
	RegionNo          *string                                        `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	StrictMode        *int32                                         `json:"StrictMode,omitempty" xml:"StrictMode,omitempty"`
	VpcId             *string                                        `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchAuto       *string                                        `json:"VswitchAuto,omitempty" xml:"VswitchAuto,omitempty"`
	VswitchCidr       *string                                        `json:"VswitchCidr,omitempty" xml:"VswitchCidr,omitempty"`
	VswitchId         *string                                        `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s CreateSecurityProxyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSecurityProxyRequest) GoString() string {
	return s.String()
}

func (s *CreateSecurityProxyRequest) SetFirewallSwitch(v string) *CreateSecurityProxyRequest {
	s.FirewallSwitch = &v
	return s
}

func (s *CreateSecurityProxyRequest) SetLang(v string) *CreateSecurityProxyRequest {
	s.Lang = &v
	return s
}

func (s *CreateSecurityProxyRequest) SetNatGatewayId(v string) *CreateSecurityProxyRequest {
	s.NatGatewayId = &v
	return s
}

func (s *CreateSecurityProxyRequest) SetNatRouteEntryList(v []*CreateSecurityProxyRequestNatRouteEntryList) *CreateSecurityProxyRequest {
	s.NatRouteEntryList = v
	return s
}

func (s *CreateSecurityProxyRequest) SetProxyName(v string) *CreateSecurityProxyRequest {
	s.ProxyName = &v
	return s
}

func (s *CreateSecurityProxyRequest) SetRegionNo(v string) *CreateSecurityProxyRequest {
	s.RegionNo = &v
	return s
}

func (s *CreateSecurityProxyRequest) SetStrictMode(v int32) *CreateSecurityProxyRequest {
	s.StrictMode = &v
	return s
}

func (s *CreateSecurityProxyRequest) SetVpcId(v string) *CreateSecurityProxyRequest {
	s.VpcId = &v
	return s
}

func (s *CreateSecurityProxyRequest) SetVswitchAuto(v string) *CreateSecurityProxyRequest {
	s.VswitchAuto = &v
	return s
}

func (s *CreateSecurityProxyRequest) SetVswitchCidr(v string) *CreateSecurityProxyRequest {
	s.VswitchCidr = &v
	return s
}

func (s *CreateSecurityProxyRequest) SetVswitchId(v string) *CreateSecurityProxyRequest {
	s.VswitchId = &v
	return s
}

type CreateSecurityProxyRequestNatRouteEntryList struct {
	DestinationCidr *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
	NextHopId       *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	NextHopType     *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	RouteTableId    *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s CreateSecurityProxyRequestNatRouteEntryList) String() string {
	return tea.Prettify(s)
}

func (s CreateSecurityProxyRequestNatRouteEntryList) GoString() string {
	return s.String()
}

func (s *CreateSecurityProxyRequestNatRouteEntryList) SetDestinationCidr(v string) *CreateSecurityProxyRequestNatRouteEntryList {
	s.DestinationCidr = &v
	return s
}

func (s *CreateSecurityProxyRequestNatRouteEntryList) SetNextHopId(v string) *CreateSecurityProxyRequestNatRouteEntryList {
	s.NextHopId = &v
	return s
}

func (s *CreateSecurityProxyRequestNatRouteEntryList) SetNextHopType(v string) *CreateSecurityProxyRequestNatRouteEntryList {
	s.NextHopType = &v
	return s
}

func (s *CreateSecurityProxyRequestNatRouteEntryList) SetRouteTableId(v string) *CreateSecurityProxyRequestNatRouteEntryList {
	s.RouteTableId = &v
	return s
}

type CreateSecurityProxyResponseBody struct {
	ProxyId   *string `json:"ProxyId,omitempty" xml:"ProxyId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSecurityProxyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSecurityProxyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSecurityProxyResponseBody) SetProxyId(v string) *CreateSecurityProxyResponseBody {
	s.ProxyId = &v
	return s
}

func (s *CreateSecurityProxyResponseBody) SetRequestId(v string) *CreateSecurityProxyResponseBody {
	s.RequestId = &v
	return s
}

type CreateSecurityProxyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSecurityProxyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSecurityProxyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSecurityProxyResponse) GoString() string {
	return s.String()
}

func (s *CreateSecurityProxyResponse) SetHeaders(v map[string]*string) *CreateSecurityProxyResponse {
	s.Headers = v
	return s
}

func (s *CreateSecurityProxyResponse) SetStatusCode(v int32) *CreateSecurityProxyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSecurityProxyResponse) SetBody(v *CreateSecurityProxyResponseBody) *CreateSecurityProxyResponse {
	s.Body = v
	return s
}

type CreateTrFirewallV2Request struct {
	// The ID of the Cloud Enterprise Network (CEN) instance.
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The description of the firewall.
	FirewallDescription *string `json:"FirewallDescription,omitempty" xml:"FirewallDescription,omitempty"`
	// The name of the firewall.
	FirewallName *string `json:"FirewallName,omitempty" xml:"FirewallName,omitempty"`
	// The subnet CIDR block of the VPC in which the ENI of the firewall is stored in automatic mode.
	FirewallSubnetCidr *string `json:"FirewallSubnetCidr,omitempty" xml:"FirewallSubnetCidr,omitempty"`
	// The CIDR block that is allocated to the VPC created for the VPC firewall in automatic mode.
	FirewallVpcCidr *string `json:"FirewallVpcCidr,omitempty" xml:"FirewallVpcCidr,omitempty"`
	// The ID of the VPC in which the ENI associated with the VPC firewall is created in manual mode.
	FirewallVpcId *string `json:"FirewallVpcId,omitempty" xml:"FirewallVpcId,omitempty"`
	// The ID of the vSwitch that is used to create the ENI in manual mode.
	FirewallVswitchId *string `json:"FirewallVswitchId,omitempty" xml:"FirewallVswitchId,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region ID of the route router.
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The routing mode of the VPC firewall. Valid values:
	//
	// *   **managed**: automatic mode
	// *   **manual**: manual mode
	RouteMode *string `json:"RouteMode,omitempty" xml:"RouteMode,omitempty"`
	// The primary subnet CIDR block that the VPC uses to connect to the transit router in automatic mode.
	TrAttachmentMasterCidr *string `json:"TrAttachmentMasterCidr,omitempty" xml:"TrAttachmentMasterCidr,omitempty"`
	// The primary zone for the vSwitch.
	TrAttachmentMasterZone *string `json:"TrAttachmentMasterZone,omitempty" xml:"TrAttachmentMasterZone,omitempty"`
	// The secondary subnet CIDR block that the VPC uses to connect to the transit router in automatic mode.
	TrAttachmentSlaveCidr *string `json:"TrAttachmentSlaveCidr,omitempty" xml:"TrAttachmentSlaveCidr,omitempty"`
	// The secondary zone for the vSwitch.
	TrAttachmentSlaveZone *string `json:"TrAttachmentSlaveZone,omitempty" xml:"TrAttachmentSlaveZone,omitempty"`
	// The ID of the transit router.
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s CreateTrFirewallV2Request) String() string {
	return tea.Prettify(s)
}

func (s CreateTrFirewallV2Request) GoString() string {
	return s.String()
}

func (s *CreateTrFirewallV2Request) SetCenId(v string) *CreateTrFirewallV2Request {
	s.CenId = &v
	return s
}

func (s *CreateTrFirewallV2Request) SetFirewallDescription(v string) *CreateTrFirewallV2Request {
	s.FirewallDescription = &v
	return s
}

func (s *CreateTrFirewallV2Request) SetFirewallName(v string) *CreateTrFirewallV2Request {
	s.FirewallName = &v
	return s
}

func (s *CreateTrFirewallV2Request) SetFirewallSubnetCidr(v string) *CreateTrFirewallV2Request {
	s.FirewallSubnetCidr = &v
	return s
}

func (s *CreateTrFirewallV2Request) SetFirewallVpcCidr(v string) *CreateTrFirewallV2Request {
	s.FirewallVpcCidr = &v
	return s
}

func (s *CreateTrFirewallV2Request) SetFirewallVpcId(v string) *CreateTrFirewallV2Request {
	s.FirewallVpcId = &v
	return s
}

func (s *CreateTrFirewallV2Request) SetFirewallVswitchId(v string) *CreateTrFirewallV2Request {
	s.FirewallVswitchId = &v
	return s
}

func (s *CreateTrFirewallV2Request) SetLang(v string) *CreateTrFirewallV2Request {
	s.Lang = &v
	return s
}

func (s *CreateTrFirewallV2Request) SetRegionNo(v string) *CreateTrFirewallV2Request {
	s.RegionNo = &v
	return s
}

func (s *CreateTrFirewallV2Request) SetRouteMode(v string) *CreateTrFirewallV2Request {
	s.RouteMode = &v
	return s
}

func (s *CreateTrFirewallV2Request) SetTrAttachmentMasterCidr(v string) *CreateTrFirewallV2Request {
	s.TrAttachmentMasterCidr = &v
	return s
}

func (s *CreateTrFirewallV2Request) SetTrAttachmentMasterZone(v string) *CreateTrFirewallV2Request {
	s.TrAttachmentMasterZone = &v
	return s
}

func (s *CreateTrFirewallV2Request) SetTrAttachmentSlaveCidr(v string) *CreateTrFirewallV2Request {
	s.TrAttachmentSlaveCidr = &v
	return s
}

func (s *CreateTrFirewallV2Request) SetTrAttachmentSlaveZone(v string) *CreateTrFirewallV2Request {
	s.TrAttachmentSlaveZone = &v
	return s
}

func (s *CreateTrFirewallV2Request) SetTransitRouterId(v string) *CreateTrFirewallV2Request {
	s.TransitRouterId = &v
	return s
}

type CreateTrFirewallV2ResponseBody struct {
	// The instance ID of the VPC firewall.
	FirewallId *string `json:"FirewallId,omitempty" xml:"FirewallId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTrFirewallV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTrFirewallV2ResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTrFirewallV2ResponseBody) SetFirewallId(v string) *CreateTrFirewallV2ResponseBody {
	s.FirewallId = &v
	return s
}

func (s *CreateTrFirewallV2ResponseBody) SetRequestId(v string) *CreateTrFirewallV2ResponseBody {
	s.RequestId = &v
	return s
}

type CreateTrFirewallV2Response struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTrFirewallV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTrFirewallV2Response) String() string {
	return tea.Prettify(s)
}

func (s CreateTrFirewallV2Response) GoString() string {
	return s.String()
}

func (s *CreateTrFirewallV2Response) SetHeaders(v map[string]*string) *CreateTrFirewallV2Response {
	s.Headers = v
	return s
}

func (s *CreateTrFirewallV2Response) SetStatusCode(v int32) *CreateTrFirewallV2Response {
	s.StatusCode = &v
	return s
}

func (s *CreateTrFirewallV2Response) SetBody(v *CreateTrFirewallV2ResponseBody) *CreateTrFirewallV2Response {
	s.Body = v
	return s
}

type CreateTrFirewallV2RoutePolicyRequest struct {
	DestCandidateList []*CreateTrFirewallV2RoutePolicyRequestDestCandidateList `json:"DestCandidateList,omitempty" xml:"DestCandidateList,omitempty" type:"Repeated"`
	FirewallId        *string                                                  `json:"FirewallId,omitempty" xml:"FirewallId,omitempty"`
	Lang              *string                                                  `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PolicyDescription *string                                                  `json:"PolicyDescription,omitempty" xml:"PolicyDescription,omitempty"`
	PolicyName        *string                                                  `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	PolicyType        *string                                                  `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	SrcCandidateList  []*CreateTrFirewallV2RoutePolicyRequestSrcCandidateList  `json:"SrcCandidateList,omitempty" xml:"SrcCandidateList,omitempty" type:"Repeated"`
}

func (s CreateTrFirewallV2RoutePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTrFirewallV2RoutePolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateTrFirewallV2RoutePolicyRequest) SetDestCandidateList(v []*CreateTrFirewallV2RoutePolicyRequestDestCandidateList) *CreateTrFirewallV2RoutePolicyRequest {
	s.DestCandidateList = v
	return s
}

func (s *CreateTrFirewallV2RoutePolicyRequest) SetFirewallId(v string) *CreateTrFirewallV2RoutePolicyRequest {
	s.FirewallId = &v
	return s
}

func (s *CreateTrFirewallV2RoutePolicyRequest) SetLang(v string) *CreateTrFirewallV2RoutePolicyRequest {
	s.Lang = &v
	return s
}

func (s *CreateTrFirewallV2RoutePolicyRequest) SetPolicyDescription(v string) *CreateTrFirewallV2RoutePolicyRequest {
	s.PolicyDescription = &v
	return s
}

func (s *CreateTrFirewallV2RoutePolicyRequest) SetPolicyName(v string) *CreateTrFirewallV2RoutePolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *CreateTrFirewallV2RoutePolicyRequest) SetPolicyType(v string) *CreateTrFirewallV2RoutePolicyRequest {
	s.PolicyType = &v
	return s
}

func (s *CreateTrFirewallV2RoutePolicyRequest) SetSrcCandidateList(v []*CreateTrFirewallV2RoutePolicyRequestSrcCandidateList) *CreateTrFirewallV2RoutePolicyRequest {
	s.SrcCandidateList = v
	return s
}

type CreateTrFirewallV2RoutePolicyRequestDestCandidateList struct {
	CandidateId   *string `json:"CandidateId,omitempty" xml:"CandidateId,omitempty"`
	CandidateType *string `json:"CandidateType,omitempty" xml:"CandidateType,omitempty"`
}

func (s CreateTrFirewallV2RoutePolicyRequestDestCandidateList) String() string {
	return tea.Prettify(s)
}

func (s CreateTrFirewallV2RoutePolicyRequestDestCandidateList) GoString() string {
	return s.String()
}

func (s *CreateTrFirewallV2RoutePolicyRequestDestCandidateList) SetCandidateId(v string) *CreateTrFirewallV2RoutePolicyRequestDestCandidateList {
	s.CandidateId = &v
	return s
}

func (s *CreateTrFirewallV2RoutePolicyRequestDestCandidateList) SetCandidateType(v string) *CreateTrFirewallV2RoutePolicyRequestDestCandidateList {
	s.CandidateType = &v
	return s
}

type CreateTrFirewallV2RoutePolicyRequestSrcCandidateList struct {
	CandidateId   *string `json:"CandidateId,omitempty" xml:"CandidateId,omitempty"`
	CandidateType *string `json:"CandidateType,omitempty" xml:"CandidateType,omitempty"`
}

func (s CreateTrFirewallV2RoutePolicyRequestSrcCandidateList) String() string {
	return tea.Prettify(s)
}

func (s CreateTrFirewallV2RoutePolicyRequestSrcCandidateList) GoString() string {
	return s.String()
}

func (s *CreateTrFirewallV2RoutePolicyRequestSrcCandidateList) SetCandidateId(v string) *CreateTrFirewallV2RoutePolicyRequestSrcCandidateList {
	s.CandidateId = &v
	return s
}

func (s *CreateTrFirewallV2RoutePolicyRequestSrcCandidateList) SetCandidateType(v string) *CreateTrFirewallV2RoutePolicyRequestSrcCandidateList {
	s.CandidateType = &v
	return s
}

type CreateTrFirewallV2RoutePolicyShrinkRequest struct {
	DestCandidateListShrink *string `json:"DestCandidateList,omitempty" xml:"DestCandidateList,omitempty"`
	FirewallId              *string `json:"FirewallId,omitempty" xml:"FirewallId,omitempty"`
	Lang                    *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PolicyDescription       *string `json:"PolicyDescription,omitempty" xml:"PolicyDescription,omitempty"`
	PolicyName              *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	PolicyType              *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	SrcCandidateListShrink  *string `json:"SrcCandidateList,omitempty" xml:"SrcCandidateList,omitempty"`
}

func (s CreateTrFirewallV2RoutePolicyShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTrFirewallV2RoutePolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateTrFirewallV2RoutePolicyShrinkRequest) SetDestCandidateListShrink(v string) *CreateTrFirewallV2RoutePolicyShrinkRequest {
	s.DestCandidateListShrink = &v
	return s
}

func (s *CreateTrFirewallV2RoutePolicyShrinkRequest) SetFirewallId(v string) *CreateTrFirewallV2RoutePolicyShrinkRequest {
	s.FirewallId = &v
	return s
}

func (s *CreateTrFirewallV2RoutePolicyShrinkRequest) SetLang(v string) *CreateTrFirewallV2RoutePolicyShrinkRequest {
	s.Lang = &v
	return s
}

func (s *CreateTrFirewallV2RoutePolicyShrinkRequest) SetPolicyDescription(v string) *CreateTrFirewallV2RoutePolicyShrinkRequest {
	s.PolicyDescription = &v
	return s
}

func (s *CreateTrFirewallV2RoutePolicyShrinkRequest) SetPolicyName(v string) *CreateTrFirewallV2RoutePolicyShrinkRequest {
	s.PolicyName = &v
	return s
}

func (s *CreateTrFirewallV2RoutePolicyShrinkRequest) SetPolicyType(v string) *CreateTrFirewallV2RoutePolicyShrinkRequest {
	s.PolicyType = &v
	return s
}

func (s *CreateTrFirewallV2RoutePolicyShrinkRequest) SetSrcCandidateListShrink(v string) *CreateTrFirewallV2RoutePolicyShrinkRequest {
	s.SrcCandidateListShrink = &v
	return s
}

type CreateTrFirewallV2RoutePolicyResponseBody struct {
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TrFirewallRoutePolicyId *string `json:"TrFirewallRoutePolicyId,omitempty" xml:"TrFirewallRoutePolicyId,omitempty"`
}

func (s CreateTrFirewallV2RoutePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTrFirewallV2RoutePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTrFirewallV2RoutePolicyResponseBody) SetRequestId(v string) *CreateTrFirewallV2RoutePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTrFirewallV2RoutePolicyResponseBody) SetTrFirewallRoutePolicyId(v string) *CreateTrFirewallV2RoutePolicyResponseBody {
	s.TrFirewallRoutePolicyId = &v
	return s
}

type CreateTrFirewallV2RoutePolicyResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTrFirewallV2RoutePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTrFirewallV2RoutePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTrFirewallV2RoutePolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateTrFirewallV2RoutePolicyResponse) SetHeaders(v map[string]*string) *CreateTrFirewallV2RoutePolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateTrFirewallV2RoutePolicyResponse) SetStatusCode(v int32) *CreateTrFirewallV2RoutePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTrFirewallV2RoutePolicyResponse) SetBody(v *CreateTrFirewallV2RoutePolicyResponseBody) *CreateTrFirewallV2RoutePolicyResponse {
	s.Body = v
	return s
}

type CreateVpcFirewallCenConfigureRequest struct {
	// The ID of the CEN instance.
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// Specifies whether to enable the VPC firewall. Valid values:
	//
	// *   **open**: After you create the VPC firewall, the VPC firewall is automatically enabled. This is the default value.
	// *   **close**: After you create the VPC firewall, the VPC firewall is disabled. You can call the [ModifyVpcFirewallCenSwitchStatus](~~345780~~) operation to manually enable the VPC firewall.
	FirewallSwitch *string `json:"FirewallSwitch,omitempty" xml:"FirewallSwitch,omitempty"`
	// The CIDR block of the vSwitch that is automatically created for the VPC firewall. You must specify a CIDR block for the Cloud_Firewall_VSWITCH VPC that is automatically created for the VPC firewall for traffic redirection. The CIDR block does not conflict with your network plan. The subnet mask of the CIDR block must be less than or equal to 29 bits in length. The CIDR block of the vSwitch must be within the network segment of the VPC.
	//
	// If you do not specify a value, the CIDR block 10.219.219.216/29 is automatically allocated.
	//
	// >  This parameter takes effect only when you create a VPC firewall for the first time in the current CEN instance and region.
	FirewallVSwitchCidrBlock *string `json:"FirewallVSwitchCidrBlock,omitempty" xml:"FirewallVSwitchCidrBlock,omitempty"`
	// The CIDR block of the VPC that is automatically created for the VPC firewall. You must specify a CIDR block for the Cloud_Firewall_VPC VPC that is automatically created for the VPC firewall for traffic redirection. The subnet mask of the CIDR block must be less than or equal to 28 bits in length.
	//
	// If you do not specify a value, the CIDR block 10.0.0.0/8 is automatically allocated.
	//
	// >  This parameter takes effect only when you create a VPC firewall for the first time in the current CEN instance and region.
	FirewallVpcCidrBlock *string `json:"FirewallVpcCidrBlock,omitempty" xml:"FirewallVpcCidrBlock,omitempty"`
	// The ID of the zone to which the vSwitch belongs. If your service is latency-sensitive, you can specify the same zone for the vSwitch of the firewall and the vSwitch of your business VPC to minimize latency.
	//
	// If you do not specify a value, a zone is automatically assigned for the vSwitch.
	//
	// >  This parameter takes effect only when you create a VPC firewall for the first time in the current CEN instance and region. For more information about zones that are supported by each region, see [Query zones](~~36064~~).
	FirewallVpcZoneId *string `json:"FirewallVpcZoneId,omitempty" xml:"FirewallVpcZoneId,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The ID of the VPC for which you want to create the VPC firewall.
	NetworkInstanceId *string `json:"NetworkInstanceId,omitempty" xml:"NetworkInstanceId,omitempty"`
	// The ID of the vSwitch that is used to associate with the elastic network interface (ENI) required by the VPC firewall.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The instance name of the VPC firewall.
	VpcFirewallName *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
	// The ID of the region to which the VPC belongs.
	//
	// > For more information about the regions, see [Supported regions](~~195657~~).
	VpcRegion *string `json:"VpcRegion,omitempty" xml:"VpcRegion,omitempty"`
}

func (s CreateVpcFirewallCenConfigureRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcFirewallCenConfigureRequest) GoString() string {
	return s.String()
}

func (s *CreateVpcFirewallCenConfigureRequest) SetCenId(v string) *CreateVpcFirewallCenConfigureRequest {
	s.CenId = &v
	return s
}

func (s *CreateVpcFirewallCenConfigureRequest) SetFirewallSwitch(v string) *CreateVpcFirewallCenConfigureRequest {
	s.FirewallSwitch = &v
	return s
}

func (s *CreateVpcFirewallCenConfigureRequest) SetFirewallVSwitchCidrBlock(v string) *CreateVpcFirewallCenConfigureRequest {
	s.FirewallVSwitchCidrBlock = &v
	return s
}

func (s *CreateVpcFirewallCenConfigureRequest) SetFirewallVpcCidrBlock(v string) *CreateVpcFirewallCenConfigureRequest {
	s.FirewallVpcCidrBlock = &v
	return s
}

func (s *CreateVpcFirewallCenConfigureRequest) SetFirewallVpcZoneId(v string) *CreateVpcFirewallCenConfigureRequest {
	s.FirewallVpcZoneId = &v
	return s
}

func (s *CreateVpcFirewallCenConfigureRequest) SetLang(v string) *CreateVpcFirewallCenConfigureRequest {
	s.Lang = &v
	return s
}

func (s *CreateVpcFirewallCenConfigureRequest) SetMemberUid(v string) *CreateVpcFirewallCenConfigureRequest {
	s.MemberUid = &v
	return s
}

func (s *CreateVpcFirewallCenConfigureRequest) SetNetworkInstanceId(v string) *CreateVpcFirewallCenConfigureRequest {
	s.NetworkInstanceId = &v
	return s
}

func (s *CreateVpcFirewallCenConfigureRequest) SetVSwitchId(v string) *CreateVpcFirewallCenConfigureRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateVpcFirewallCenConfigureRequest) SetVpcFirewallName(v string) *CreateVpcFirewallCenConfigureRequest {
	s.VpcFirewallName = &v
	return s
}

func (s *CreateVpcFirewallCenConfigureRequest) SetVpcRegion(v string) *CreateVpcFirewallCenConfigureRequest {
	s.VpcRegion = &v
	return s
}

type CreateVpcFirewallCenConfigureResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The instance ID of the VPC firewall.
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s CreateVpcFirewallCenConfigureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcFirewallCenConfigureResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVpcFirewallCenConfigureResponseBody) SetRequestId(v string) *CreateVpcFirewallCenConfigureResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVpcFirewallCenConfigureResponseBody) SetVpcFirewallId(v string) *CreateVpcFirewallCenConfigureResponseBody {
	s.VpcFirewallId = &v
	return s
}

type CreateVpcFirewallCenConfigureResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVpcFirewallCenConfigureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVpcFirewallCenConfigureResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcFirewallCenConfigureResponse) GoString() string {
	return s.String()
}

func (s *CreateVpcFirewallCenConfigureResponse) SetHeaders(v map[string]*string) *CreateVpcFirewallCenConfigureResponse {
	s.Headers = v
	return s
}

func (s *CreateVpcFirewallCenConfigureResponse) SetStatusCode(v int32) *CreateVpcFirewallCenConfigureResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVpcFirewallCenConfigureResponse) SetBody(v *CreateVpcFirewallCenConfigureResponseBody) *CreateVpcFirewallCenConfigureResponse {
	s.Body = v
	return s
}

type CreateVpcFirewallConfigureRequest struct {
	// The status of the VPC firewall after you create the firewall. Valid values:
	//
	// *   **open**: After you create the VPC firewall, the VPC firewall is automatically enabled. This is the default value.
	// *   **close**: After you create the VPC firewall, the VPC firewall is disabled. To enable the firewall, you can call the [ModifyVpcFirewallSwitchStatus](~~342935~~) operation.
	FirewallSwitch *string `json:"FirewallSwitch,omitempty" xml:"FirewallSwitch,omitempty"`
	// The language of the content within the request and the response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English.
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The CIDR blocks of the local VPC. The value is a JSON string that contains the following parameters:
	//
	// *   **RouteTableId**: the ID of the route table for the local VPC.
	// *   **RouteEntryList**: The value is a JSON string that contains the DestinationCidr and NextHopInstanceId parameters. The DestinationCidr parameter indicates the destination CIDR block of the local VPC. The NextHopInstanceId parameter indicates the instance ID of the next hop for the local VPC.
	LocalVpcCidrTableList *string `json:"LocalVpcCidrTableList,omitempty" xml:"LocalVpcCidrTableList,omitempty"`
	// The ID of the local VPC.
	LocalVpcId *string `json:"LocalVpcId,omitempty" xml:"LocalVpcId,omitempty"`
	// The region ID of the local VPC.
	//
	// >  For more information about the regions in which Cloud Firewall is available, see [Supported regions](~~195657~~).
	LocalVpcRegion *string `json:"LocalVpcRegion,omitempty" xml:"LocalVpcRegion,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The CIDR blocks of the peer VPC. The value is a JSON string that contains the following parameters:
	//
	// *   **RouteTableId**: the ID of the route table for the peer VPC.
	// *   **RouteEntryList**: The value is a JSON string that contains the DestinationCidr and NextHopInstanceId parameters. The DestinationCidr parameter indicates the destination CIDR block of the peer VPC. The NextHopInstanceId parameter indicates the instance ID of the next hop for the peer VPC.
	PeerVpcCidrTableList *string `json:"PeerVpcCidrTableList,omitempty" xml:"PeerVpcCidrTableList,omitempty"`
	// The ID of the peer VPC.
	PeerVpcId *string `json:"PeerVpcId,omitempty" xml:"PeerVpcId,omitempty"`
	// The region ID of the peer VPC.
	//
	// >  For more information about Cloud Firewall supported regions, see [Supported regions](~~195657~~).
	PeerVpcRegion *string `json:"PeerVpcRegion,omitempty" xml:"PeerVpcRegion,omitempty"`
	// The instance name of the VPC firewall.
	VpcFirewallName *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
}

func (s CreateVpcFirewallConfigureRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcFirewallConfigureRequest) GoString() string {
	return s.String()
}

func (s *CreateVpcFirewallConfigureRequest) SetFirewallSwitch(v string) *CreateVpcFirewallConfigureRequest {
	s.FirewallSwitch = &v
	return s
}

func (s *CreateVpcFirewallConfigureRequest) SetLang(v string) *CreateVpcFirewallConfigureRequest {
	s.Lang = &v
	return s
}

func (s *CreateVpcFirewallConfigureRequest) SetLocalVpcCidrTableList(v string) *CreateVpcFirewallConfigureRequest {
	s.LocalVpcCidrTableList = &v
	return s
}

func (s *CreateVpcFirewallConfigureRequest) SetLocalVpcId(v string) *CreateVpcFirewallConfigureRequest {
	s.LocalVpcId = &v
	return s
}

func (s *CreateVpcFirewallConfigureRequest) SetLocalVpcRegion(v string) *CreateVpcFirewallConfigureRequest {
	s.LocalVpcRegion = &v
	return s
}

func (s *CreateVpcFirewallConfigureRequest) SetMemberUid(v string) *CreateVpcFirewallConfigureRequest {
	s.MemberUid = &v
	return s
}

func (s *CreateVpcFirewallConfigureRequest) SetPeerVpcCidrTableList(v string) *CreateVpcFirewallConfigureRequest {
	s.PeerVpcCidrTableList = &v
	return s
}

func (s *CreateVpcFirewallConfigureRequest) SetPeerVpcId(v string) *CreateVpcFirewallConfigureRequest {
	s.PeerVpcId = &v
	return s
}

func (s *CreateVpcFirewallConfigureRequest) SetPeerVpcRegion(v string) *CreateVpcFirewallConfigureRequest {
	s.PeerVpcRegion = &v
	return s
}

func (s *CreateVpcFirewallConfigureRequest) SetVpcFirewallName(v string) *CreateVpcFirewallConfigureRequest {
	s.VpcFirewallName = &v
	return s
}

type CreateVpcFirewallConfigureResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The instance ID of the VPC firewall.
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s CreateVpcFirewallConfigureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcFirewallConfigureResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVpcFirewallConfigureResponseBody) SetRequestId(v string) *CreateVpcFirewallConfigureResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVpcFirewallConfigureResponseBody) SetVpcFirewallId(v string) *CreateVpcFirewallConfigureResponseBody {
	s.VpcFirewallId = &v
	return s
}

type CreateVpcFirewallConfigureResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVpcFirewallConfigureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVpcFirewallConfigureResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcFirewallConfigureResponse) GoString() string {
	return s.String()
}

func (s *CreateVpcFirewallConfigureResponse) SetHeaders(v map[string]*string) *CreateVpcFirewallConfigureResponse {
	s.Headers = v
	return s
}

func (s *CreateVpcFirewallConfigureResponse) SetStatusCode(v int32) *CreateVpcFirewallConfigureResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVpcFirewallConfigureResponse) SetBody(v *CreateVpcFirewallConfigureResponseBody) *CreateVpcFirewallConfigureResponse {
	s.Body = v
	return s
}

type CreateVpcFirewallControlPolicyRequest struct {
	// The action that Cloud Firewall performs on the traffic. Valid values:
	//
	// - **accept**: allows the traffic.
	// - **drop**: blocks the traffic.
	// - **log**: monitors the traffic.
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The type of the applications that the access control policy supports. Valid values:
	//
	// - **FTP**
	// - **HTTP**
	// - **HTTPS**
	// - **MySQL**
	// - **SMTP**
	// - **SMTPS**
	// - **RDP**
	// - **VNC**
	// - **SSH**
	// - **Redis**
	// - **MQTT**
	// - **MongoDB**
	// - **Memcache**
	// - **SSL**
	// - **ANY**: all types of applications
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The application types supported by the access control policy.
	ApplicationNameList []*string `json:"ApplicationNameList,omitempty" xml:"ApplicationNameList,omitempty" type:"Repeated"`
	// The description of the access control policy.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination port in the access control policy.
	//
	// >  If **DestPortType** is set to `port`, you must specify this parameter.
	DestPort *string `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	// The name of the destination port address book in the access control policy.
	//
	// >  If **DestPortType** is set to `group`, you must specify this parameter.
	DestPortGroup *string `json:"DestPortGroup,omitempty" xml:"DestPortGroup,omitempty"`
	// The type of the destination port in the access control policy. Valid values:
	//
	// - **port**: port
	// - **group**: port address book
	DestPortType *string `json:"DestPortType,omitempty" xml:"DestPortType,omitempty"`
	// The destination address in the access control policy. Valid values:
	//
	// - If **DestinationType** is set to `net`, the value of **Destination** must be a CIDR block.
	// - If **DestinationType** is set to `group`, the value of **Destination** must be an address book.
	// - If **DestinationType** is set to `domain`, the value of **Destination** must be a domain name.
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The type of the destination address in the access control policy. Valid values:
	//
	// - **net**: CIDR block
	// - **group**: address book
	// - **domain**: domain name
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The time when the access control policy stops taking effect. The value is a UNIX timestamp. Unit: seconds. The value must be on the hour or on the half hour, and at least 30 minutes later than the start time.
	//
	// >  If you set RepeatType to Permanent, leave this parameter empty. If you set RepeatType to None, Daily, Weekly, or Monthly, you must specify this parameter.
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// - **zh**: Chinese (default)
	// - **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The priority of the access control policy.
	//
	// The priority value starts from 1. A smaller priority value indicates a higher priority.
	NewOrder *string `json:"NewOrder,omitempty" xml:"NewOrder,omitempty"`
	// The type of the protocol in the access control policy. Valid values:
	//
	// - **ANY** (If you are not sure about the protocol type, you can set this parameter to ANY.)
	// - **TCP**
	// - **UDP**
	// - **ICMP**
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// Specifies whether to enable the access control policy. By default, an access control policy is enabled after the policy is created. Valid values:
	//
	// - **true**: enables the access control policy.
	// - **false**: disables the access control policy.
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The days of a week or of a month on which the access control policy takes effect.
	//
	// *   If you set RepeatType to `Permanent`, `None`, or `Daily`, leave this parameter empty. Example: \[].
	// *   If you set RepeatType to Weekly, you must specify this parameter. Example: \[0, 6].
	//
	// >  If you set RepeatType to Weekly, the fields in the value of this parameter cannot be repeated.
	//
	// *   If you set RepeatType to `Monthly`, you must specify this parameter. Example: \[1, 31].
	//
	// >  If you set RepeatType to Monthly, the fields in the value of this parameter cannot be repeated.
	RepeatDays []*int64 `json:"RepeatDays,omitempty" xml:"RepeatDays,omitempty" type:"Repeated"`
	// The point in time when the recurrence ends. Example: 23:30. The value must be on the hour or on the half hour, and at least 30 minutes later than the start time.
	//
	// >  If you set RepeatType to Permanent or None, leave this parameter empty. If you set RepeatType to Daily, Weekly, or Monthly, you must specify this parameter.
	RepeatEndTime *string `json:"RepeatEndTime,omitempty" xml:"RepeatEndTime,omitempty"`
	// The point in time when the recurrence starts. Example: 08:00. The value must be on the hour or on the half hour, and at least 30 minutes earlier than the end time.
	//
	// >  If you set RepeatType to Permanent or None, leave this parameter empty. If you set RepeatType to Daily, Weekly, or Monthly, you must specify this parameter.
	RepeatStartTime *string `json:"RepeatStartTime,omitempty" xml:"RepeatStartTime,omitempty"`
	// The recurrence type for the access control policy to take effect. Valid values:
	//
	// *   **Permanent** (default): The policy always takes effect.
	// *   **None**: The policy takes effect for only once.
	// *   **Daily**: The policy takes effect on a daily basis.
	// *   **Weekly**: The policy takes effect on a weekly basis.
	// *   **Monthly**: The policy takes effect on a monthly basis.
	RepeatType *string `json:"RepeatType,omitempty" xml:"RepeatType,omitempty"`
	// The source address in the access control policy.
	//
	// - If SourceType is set to `net`, the value of Source must be a CIDR block.
	// - If SourceType is set to `group`, the value of Source must be an address book.
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The type of the source address in the access control policy. Valid values:
	//
	// - **net**: CIDR block
	// - **group**: address book
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The time when the access control policy starts to take effect. The value is a UNIX timestamp. Unit: seconds. The value must be on the hour or on the half hour, and at least 30 minutes earlier than the end time.
	//
	// >  If you set RepeatType to Permanent, leave this parameter empty. If you set RepeatType to None, Daily, Weekly, or Monthly, you must specify this parameter.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the policy group in which you want to create the access control policy.
	//
	// - If a VPC firewall protects the traffic between two VPCs that are connected by using a CEN instance, the value of this parameter must be the ID of the CEN instance.
	// - If a VPC firewall protects the traffic between two VPCs that are connected by using an Express Connect circuit, the value of this parameter must be the instance ID of the VPC firewall.
	//
	// >  You can call the [DescribeVpcFirewallAclGroupList](https://www.alibabacloud.com/help/en/cloud-firewall/latest/describevpcfirewallaclgrouplist) operation to query the IDs.
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s CreateVpcFirewallControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcFirewallControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateVpcFirewallControlPolicyRequest) SetAclAction(v string) *CreateVpcFirewallControlPolicyRequest {
	s.AclAction = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetApplicationName(v string) *CreateVpcFirewallControlPolicyRequest {
	s.ApplicationName = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetApplicationNameList(v []*string) *CreateVpcFirewallControlPolicyRequest {
	s.ApplicationNameList = v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetDescription(v string) *CreateVpcFirewallControlPolicyRequest {
	s.Description = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetDestPort(v string) *CreateVpcFirewallControlPolicyRequest {
	s.DestPort = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetDestPortGroup(v string) *CreateVpcFirewallControlPolicyRequest {
	s.DestPortGroup = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetDestPortType(v string) *CreateVpcFirewallControlPolicyRequest {
	s.DestPortType = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetDestination(v string) *CreateVpcFirewallControlPolicyRequest {
	s.Destination = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetDestinationType(v string) *CreateVpcFirewallControlPolicyRequest {
	s.DestinationType = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetEndTime(v int64) *CreateVpcFirewallControlPolicyRequest {
	s.EndTime = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetLang(v string) *CreateVpcFirewallControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetMemberUid(v string) *CreateVpcFirewallControlPolicyRequest {
	s.MemberUid = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetNewOrder(v string) *CreateVpcFirewallControlPolicyRequest {
	s.NewOrder = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetProto(v string) *CreateVpcFirewallControlPolicyRequest {
	s.Proto = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetRelease(v string) *CreateVpcFirewallControlPolicyRequest {
	s.Release = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetRepeatDays(v []*int64) *CreateVpcFirewallControlPolicyRequest {
	s.RepeatDays = v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetRepeatEndTime(v string) *CreateVpcFirewallControlPolicyRequest {
	s.RepeatEndTime = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetRepeatStartTime(v string) *CreateVpcFirewallControlPolicyRequest {
	s.RepeatStartTime = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetRepeatType(v string) *CreateVpcFirewallControlPolicyRequest {
	s.RepeatType = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetSource(v string) *CreateVpcFirewallControlPolicyRequest {
	s.Source = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetSourceType(v string) *CreateVpcFirewallControlPolicyRequest {
	s.SourceType = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetStartTime(v int64) *CreateVpcFirewallControlPolicyRequest {
	s.StartTime = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetVpcFirewallId(v string) *CreateVpcFirewallControlPolicyRequest {
	s.VpcFirewallId = &v
	return s
}

type CreateVpcFirewallControlPolicyResponseBody struct {
	// The ID of the access control policy.
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVpcFirewallControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcFirewallControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVpcFirewallControlPolicyResponseBody) SetAclUuid(v string) *CreateVpcFirewallControlPolicyResponseBody {
	s.AclUuid = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyResponseBody) SetRequestId(v string) *CreateVpcFirewallControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

type CreateVpcFirewallControlPolicyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVpcFirewallControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVpcFirewallControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcFirewallControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateVpcFirewallControlPolicyResponse) SetHeaders(v map[string]*string) *CreateVpcFirewallControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateVpcFirewallControlPolicyResponse) SetStatusCode(v int32) *CreateVpcFirewallControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyResponse) SetBody(v *CreateVpcFirewallControlPolicyResponseBody) *CreateVpcFirewallControlPolicyResponse {
	s.Body = v
	return s
}

type DeleteAddressBookRequest struct {
	// The ID of the address book.
	//
	// To delete the address book, you must provide the ID of the address book. You can call the DescribeAddressBook operation to query the ID.
	GroupUuid *string `json:"GroupUuid,omitempty" xml:"GroupUuid,omitempty"`
	// The natural language of the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Deprecated
	// The source IP address of the request.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DeleteAddressBookRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAddressBookRequest) GoString() string {
	return s.String()
}

func (s *DeleteAddressBookRequest) SetGroupUuid(v string) *DeleteAddressBookRequest {
	s.GroupUuid = &v
	return s
}

func (s *DeleteAddressBookRequest) SetLang(v string) *DeleteAddressBookRequest {
	s.Lang = &v
	return s
}

func (s *DeleteAddressBookRequest) SetSourceIp(v string) *DeleteAddressBookRequest {
	s.SourceIp = &v
	return s
}

type DeleteAddressBookResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAddressBookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAddressBookResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAddressBookResponseBody) SetRequestId(v string) *DeleteAddressBookResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAddressBookResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAddressBookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAddressBookResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAddressBookResponse) GoString() string {
	return s.String()
}

func (s *DeleteAddressBookResponse) SetHeaders(v map[string]*string) *DeleteAddressBookResponse {
	s.Headers = v
	return s
}

func (s *DeleteAddressBookResponse) SetStatusCode(v int32) *DeleteAddressBookResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAddressBookResponse) SetBody(v *DeleteAddressBookResponseBody) *DeleteAddressBookResponse {
	s.Body = v
	return s
}

type DeleteControlPolicyRequest struct {
	// The ID of the access control policy.
	//
	// To delete an access control policy, you must provide the ID of the policy. You can call the [DescribeControlPolicy](~~138866~~) operation to query the ID.
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The direction of the traffic to which the access control policy applies.
	//
	// Valid values:
	//
	// *   **in**: inbound traffic
	// *   **out**: outbound traffic
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The natural language of the request and response.
	//
	// Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Deprecated
	// The source IP address of the traffic.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DeleteControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteControlPolicyRequest) SetAclUuid(v string) *DeleteControlPolicyRequest {
	s.AclUuid = &v
	return s
}

func (s *DeleteControlPolicyRequest) SetDirection(v string) *DeleteControlPolicyRequest {
	s.Direction = &v
	return s
}

func (s *DeleteControlPolicyRequest) SetLang(v string) *DeleteControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *DeleteControlPolicyRequest) SetSourceIp(v string) *DeleteControlPolicyRequest {
	s.SourceIp = &v
	return s
}

type DeleteControlPolicyResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteControlPolicyResponseBody) SetRequestId(v string) *DeleteControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

type DeleteControlPolicyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteControlPolicyResponse) SetHeaders(v map[string]*string) *DeleteControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteControlPolicyResponse) SetStatusCode(v int32) *DeleteControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteControlPolicyResponse) SetBody(v *DeleteControlPolicyResponseBody) *DeleteControlPolicyResponse {
	s.Body = v
	return s
}

type DeleteControlPolicyTemplateRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// *   **zh** (default): Chinese
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The source IP address of the request.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The ID of the access control policy template.
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteControlPolicyTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteControlPolicyTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteControlPolicyTemplateRequest) SetLang(v string) *DeleteControlPolicyTemplateRequest {
	s.Lang = &v
	return s
}

func (s *DeleteControlPolicyTemplateRequest) SetSourceIp(v string) *DeleteControlPolicyTemplateRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteControlPolicyTemplateRequest) SetTemplateId(v string) *DeleteControlPolicyTemplateRequest {
	s.TemplateId = &v
	return s
}

type DeleteControlPolicyTemplateResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteControlPolicyTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteControlPolicyTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteControlPolicyTemplateResponseBody) SetRequestId(v string) *DeleteControlPolicyTemplateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteControlPolicyTemplateResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteControlPolicyTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteControlPolicyTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteControlPolicyTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteControlPolicyTemplateResponse) SetHeaders(v map[string]*string) *DeleteControlPolicyTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteControlPolicyTemplateResponse) SetStatusCode(v int32) *DeleteControlPolicyTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteControlPolicyTemplateResponse) SetBody(v *DeleteControlPolicyTemplateResponseBody) *DeleteControlPolicyTemplateResponse {
	s.Body = v
	return s
}

type DeleteDownloadTaskRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// *   **zh** (default): Chinese
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the file download task.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteDownloadTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDownloadTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteDownloadTaskRequest) SetLang(v string) *DeleteDownloadTaskRequest {
	s.Lang = &v
	return s
}

func (s *DeleteDownloadTaskRequest) SetTaskId(v string) *DeleteDownloadTaskRequest {
	s.TaskId = &v
	return s
}

type DeleteDownloadTaskResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDownloadTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDownloadTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDownloadTaskResponseBody) SetRequestId(v string) *DeleteDownloadTaskResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDownloadTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDownloadTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDownloadTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDownloadTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteDownloadTaskResponse) SetHeaders(v map[string]*string) *DeleteDownloadTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteDownloadTaskResponse) SetStatusCode(v int32) *DeleteDownloadTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDownloadTaskResponse) SetBody(v *DeleteDownloadTaskResponseBody) *DeleteDownloadTaskResponse {
	s.Body = v
	return s
}

type DeleteFirewallV2RoutePoliciesRequest struct {
	FirewallId              *string `json:"FirewallId,omitempty" xml:"FirewallId,omitempty"`
	Lang                    *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	TrFirewallRoutePolicyId *string `json:"TrFirewallRoutePolicyId,omitempty" xml:"TrFirewallRoutePolicyId,omitempty"`
}

func (s DeleteFirewallV2RoutePoliciesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFirewallV2RoutePoliciesRequest) GoString() string {
	return s.String()
}

func (s *DeleteFirewallV2RoutePoliciesRequest) SetFirewallId(v string) *DeleteFirewallV2RoutePoliciesRequest {
	s.FirewallId = &v
	return s
}

func (s *DeleteFirewallV2RoutePoliciesRequest) SetLang(v string) *DeleteFirewallV2RoutePoliciesRequest {
	s.Lang = &v
	return s
}

func (s *DeleteFirewallV2RoutePoliciesRequest) SetTrFirewallRoutePolicyId(v string) *DeleteFirewallV2RoutePoliciesRequest {
	s.TrFirewallRoutePolicyId = &v
	return s
}

type DeleteFirewallV2RoutePoliciesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFirewallV2RoutePoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFirewallV2RoutePoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFirewallV2RoutePoliciesResponseBody) SetRequestId(v string) *DeleteFirewallV2RoutePoliciesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFirewallV2RoutePoliciesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFirewallV2RoutePoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFirewallV2RoutePoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFirewallV2RoutePoliciesResponse) GoString() string {
	return s.String()
}

func (s *DeleteFirewallV2RoutePoliciesResponse) SetHeaders(v map[string]*string) *DeleteFirewallV2RoutePoliciesResponse {
	s.Headers = v
	return s
}

func (s *DeleteFirewallV2RoutePoliciesResponse) SetStatusCode(v int32) *DeleteFirewallV2RoutePoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFirewallV2RoutePoliciesResponse) SetBody(v *DeleteFirewallV2RoutePoliciesResponseBody) *DeleteFirewallV2RoutePoliciesResponse {
	s.Body = v
	return s
}

type DeleteInstanceMembersRequest struct {
	// The unique identifiers (UID) of members that you want to remove from Cloud Firewall.
	MemberUids []*int64 `json:"MemberUids,omitempty" xml:"MemberUids,omitempty" type:"Repeated"`
}

func (s DeleteInstanceMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceMembersRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceMembersRequest) SetMemberUids(v []*int64) *DeleteInstanceMembersRequest {
	s.MemberUids = v
	return s
}

type DeleteInstanceMembersResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInstanceMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceMembersResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceMembersResponseBody) SetRequestId(v string) *DeleteInstanceMembersResponseBody {
	s.RequestId = &v
	return s
}

type DeleteInstanceMembersResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInstanceMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInstanceMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceMembersResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceMembersResponse) SetHeaders(v map[string]*string) *DeleteInstanceMembersResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceMembersResponse) SetStatusCode(v int32) *DeleteInstanceMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstanceMembersResponse) SetBody(v *DeleteInstanceMembersResponseBody) *DeleteInstanceMembersResponse {
	s.Body = v
	return s
}

type DeleteNatFirewallControlPolicyRequest struct {
	// The UUID of the access control policy.
	//
	// To delete an access control policy, you must provide the ID of the policy. You can call the DescribeNatFirewallControlPolicy operation to query the UUIDs of access control policies.
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The direction of the traffic to which the access control policy applies.
	//
	// Valid values:
	//
	// *   **out**: outbound traffic
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the NAT gateway.
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
}

func (s DeleteNatFirewallControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNatFirewallControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteNatFirewallControlPolicyRequest) SetAclUuid(v string) *DeleteNatFirewallControlPolicyRequest {
	s.AclUuid = &v
	return s
}

func (s *DeleteNatFirewallControlPolicyRequest) SetDirection(v string) *DeleteNatFirewallControlPolicyRequest {
	s.Direction = &v
	return s
}

func (s *DeleteNatFirewallControlPolicyRequest) SetLang(v string) *DeleteNatFirewallControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *DeleteNatFirewallControlPolicyRequest) SetNatGatewayId(v string) *DeleteNatFirewallControlPolicyRequest {
	s.NatGatewayId = &v
	return s
}

type DeleteNatFirewallControlPolicyResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNatFirewallControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteNatFirewallControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNatFirewallControlPolicyResponseBody) SetRequestId(v string) *DeleteNatFirewallControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

type DeleteNatFirewallControlPolicyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNatFirewallControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNatFirewallControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNatFirewallControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteNatFirewallControlPolicyResponse) SetHeaders(v map[string]*string) *DeleteNatFirewallControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteNatFirewallControlPolicyResponse) SetStatusCode(v int32) *DeleteNatFirewallControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNatFirewallControlPolicyResponse) SetBody(v *DeleteNatFirewallControlPolicyResponseBody) *DeleteNatFirewallControlPolicyResponse {
	s.Body = v
	return s
}

type DeleteNatFirewallControlPolicyBatchRequest struct {
	// The UUIDs of access control policies.
	AclUuidList []*string `json:"AclUuidList,omitempty" xml:"AclUuidList,omitempty" type:"Repeated"`
	// The direction of the traffic to which the access control policy applies. Valid values:
	//
	// *   **out**: outbound traffic
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// *   **zh** (default): Chinese
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the NAT gateway.
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
}

func (s DeleteNatFirewallControlPolicyBatchRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNatFirewallControlPolicyBatchRequest) GoString() string {
	return s.String()
}

func (s *DeleteNatFirewallControlPolicyBatchRequest) SetAclUuidList(v []*string) *DeleteNatFirewallControlPolicyBatchRequest {
	s.AclUuidList = v
	return s
}

func (s *DeleteNatFirewallControlPolicyBatchRequest) SetDirection(v string) *DeleteNatFirewallControlPolicyBatchRequest {
	s.Direction = &v
	return s
}

func (s *DeleteNatFirewallControlPolicyBatchRequest) SetLang(v string) *DeleteNatFirewallControlPolicyBatchRequest {
	s.Lang = &v
	return s
}

func (s *DeleteNatFirewallControlPolicyBatchRequest) SetNatGatewayId(v string) *DeleteNatFirewallControlPolicyBatchRequest {
	s.NatGatewayId = &v
	return s
}

type DeleteNatFirewallControlPolicyBatchResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNatFirewallControlPolicyBatchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteNatFirewallControlPolicyBatchResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNatFirewallControlPolicyBatchResponseBody) SetRequestId(v string) *DeleteNatFirewallControlPolicyBatchResponseBody {
	s.RequestId = &v
	return s
}

type DeleteNatFirewallControlPolicyBatchResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNatFirewallControlPolicyBatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNatFirewallControlPolicyBatchResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNatFirewallControlPolicyBatchResponse) GoString() string {
	return s.String()
}

func (s *DeleteNatFirewallControlPolicyBatchResponse) SetHeaders(v map[string]*string) *DeleteNatFirewallControlPolicyBatchResponse {
	s.Headers = v
	return s
}

func (s *DeleteNatFirewallControlPolicyBatchResponse) SetStatusCode(v int32) *DeleteNatFirewallControlPolicyBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNatFirewallControlPolicyBatchResponse) SetBody(v *DeleteNatFirewallControlPolicyBatchResponseBody) *DeleteNatFirewallControlPolicyBatchResponse {
	s.Body = v
	return s
}

type DeleteTrFirewallV2Request struct {
	// The instance ID of the VPC firewall.
	FirewallId *string `json:"FirewallId,omitempty" xml:"FirewallId,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DeleteTrFirewallV2Request) String() string {
	return tea.Prettify(s)
}

func (s DeleteTrFirewallV2Request) GoString() string {
	return s.String()
}

func (s *DeleteTrFirewallV2Request) SetFirewallId(v string) *DeleteTrFirewallV2Request {
	s.FirewallId = &v
	return s
}

func (s *DeleteTrFirewallV2Request) SetLang(v string) *DeleteTrFirewallV2Request {
	s.Lang = &v
	return s
}

type DeleteTrFirewallV2ResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTrFirewallV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTrFirewallV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTrFirewallV2ResponseBody) SetRequestId(v string) *DeleteTrFirewallV2ResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTrFirewallV2Response struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTrFirewallV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTrFirewallV2Response) String() string {
	return tea.Prettify(s)
}

func (s DeleteTrFirewallV2Response) GoString() string {
	return s.String()
}

func (s *DeleteTrFirewallV2Response) SetHeaders(v map[string]*string) *DeleteTrFirewallV2Response {
	s.Headers = v
	return s
}

func (s *DeleteTrFirewallV2Response) SetStatusCode(v int32) *DeleteTrFirewallV2Response {
	s.StatusCode = &v
	return s
}

func (s *DeleteTrFirewallV2Response) SetBody(v *DeleteTrFirewallV2ResponseBody) *DeleteTrFirewallV2Response {
	s.Body = v
	return s
}

type DeleteVpcFirewallCenConfigureRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The instance IDs of VPC firewalls.
	VpcFirewallIdList []*string `json:"VpcFirewallIdList,omitempty" xml:"VpcFirewallIdList,omitempty" type:"Repeated"`
}

func (s DeleteVpcFirewallCenConfigureRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpcFirewallCenConfigureRequest) GoString() string {
	return s.String()
}

func (s *DeleteVpcFirewallCenConfigureRequest) SetLang(v string) *DeleteVpcFirewallCenConfigureRequest {
	s.Lang = &v
	return s
}

func (s *DeleteVpcFirewallCenConfigureRequest) SetMemberUid(v string) *DeleteVpcFirewallCenConfigureRequest {
	s.MemberUid = &v
	return s
}

func (s *DeleteVpcFirewallCenConfigureRequest) SetVpcFirewallIdList(v []*string) *DeleteVpcFirewallCenConfigureRequest {
	s.VpcFirewallIdList = v
	return s
}

type DeleteVpcFirewallCenConfigureResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVpcFirewallCenConfigureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpcFirewallCenConfigureResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVpcFirewallCenConfigureResponseBody) SetRequestId(v string) *DeleteVpcFirewallCenConfigureResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVpcFirewallCenConfigureResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVpcFirewallCenConfigureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVpcFirewallCenConfigureResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpcFirewallCenConfigureResponse) GoString() string {
	return s.String()
}

func (s *DeleteVpcFirewallCenConfigureResponse) SetHeaders(v map[string]*string) *DeleteVpcFirewallCenConfigureResponse {
	s.Headers = v
	return s
}

func (s *DeleteVpcFirewallCenConfigureResponse) SetStatusCode(v int32) *DeleteVpcFirewallCenConfigureResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVpcFirewallCenConfigureResponse) SetBody(v *DeleteVpcFirewallCenConfigureResponseBody) *DeleteVpcFirewallCenConfigureResponse {
	s.Body = v
	return s
}

type DeleteVpcFirewallConfigureRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The instance IDs of VPC firewalls.
	VpcFirewallIdList []*string `json:"VpcFirewallIdList,omitempty" xml:"VpcFirewallIdList,omitempty" type:"Repeated"`
}

func (s DeleteVpcFirewallConfigureRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpcFirewallConfigureRequest) GoString() string {
	return s.String()
}

func (s *DeleteVpcFirewallConfigureRequest) SetLang(v string) *DeleteVpcFirewallConfigureRequest {
	s.Lang = &v
	return s
}

func (s *DeleteVpcFirewallConfigureRequest) SetMemberUid(v string) *DeleteVpcFirewallConfigureRequest {
	s.MemberUid = &v
	return s
}

func (s *DeleteVpcFirewallConfigureRequest) SetVpcFirewallIdList(v []*string) *DeleteVpcFirewallConfigureRequest {
	s.VpcFirewallIdList = v
	return s
}

type DeleteVpcFirewallConfigureResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVpcFirewallConfigureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpcFirewallConfigureResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVpcFirewallConfigureResponseBody) SetRequestId(v string) *DeleteVpcFirewallConfigureResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVpcFirewallConfigureResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVpcFirewallConfigureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVpcFirewallConfigureResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpcFirewallConfigureResponse) GoString() string {
	return s.String()
}

func (s *DeleteVpcFirewallConfigureResponse) SetHeaders(v map[string]*string) *DeleteVpcFirewallConfigureResponse {
	s.Headers = v
	return s
}

func (s *DeleteVpcFirewallConfigureResponse) SetStatusCode(v int32) *DeleteVpcFirewallConfigureResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVpcFirewallConfigureResponse) SetBody(v *DeleteVpcFirewallConfigureResponseBody) *DeleteVpcFirewallConfigureResponse {
	s.Body = v
	return s
}

type DeleteVpcFirewallControlPolicyRequest struct {
	// The ID of the access control policy.
	//
	// To delete an access control policy, you must provide the ID of the policy. You can call the **DescribeVpcFirewallControlPolicy** operation to query the ID.
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The natural language of the request and response. Valid values:
	//
	// - **zh**: Chinese
	// - **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the group to which the access control policy belongs. You can call the **DescribeVpcFirewallAclGroupList** operation to query the ID.
	//
	// Valid values:
	//
	// - If the VPC firewall is used to protect a CEN instance, the value of this parameter is the ID of the CEN instance.
	//
	// Example: cen-ervw0g12b5jbw****
	// - If the VPC firewall is used to protect an Express Connect circuit, the value of this parameter is the ID of the VPC firewall.
	//
	// Example: vfw-a42bbb7b887148c9****
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s DeleteVpcFirewallControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpcFirewallControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteVpcFirewallControlPolicyRequest) SetAclUuid(v string) *DeleteVpcFirewallControlPolicyRequest {
	s.AclUuid = &v
	return s
}

func (s *DeleteVpcFirewallControlPolicyRequest) SetLang(v string) *DeleteVpcFirewallControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *DeleteVpcFirewallControlPolicyRequest) SetVpcFirewallId(v string) *DeleteVpcFirewallControlPolicyRequest {
	s.VpcFirewallId = &v
	return s
}

type DeleteVpcFirewallControlPolicyResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVpcFirewallControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpcFirewallControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVpcFirewallControlPolicyResponseBody) SetRequestId(v string) *DeleteVpcFirewallControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVpcFirewallControlPolicyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVpcFirewallControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVpcFirewallControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpcFirewallControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteVpcFirewallControlPolicyResponse) SetHeaders(v map[string]*string) *DeleteVpcFirewallControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteVpcFirewallControlPolicyResponse) SetStatusCode(v int32) *DeleteVpcFirewallControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVpcFirewallControlPolicyResponse) SetBody(v *DeleteVpcFirewallControlPolicyResponseBody) *DeleteVpcFirewallControlPolicyResponse {
	s.Body = v
	return s
}

type DescribeACLProtectTrendRequest struct {
	// The end of the time range to query. The value is a UNIX timestamp that is accurate to seconds.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the content within the request and the response. Valid values:
	//
	// *   **zh** (default): Chinese
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is deprecated.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp that is accurate to seconds.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeACLProtectTrendRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeACLProtectTrendRequest) GoString() string {
	return s.String()
}

func (s *DescribeACLProtectTrendRequest) SetEndTime(v string) *DescribeACLProtectTrendRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeACLProtectTrendRequest) SetLang(v string) *DescribeACLProtectTrendRequest {
	s.Lang = &v
	return s
}

func (s *DescribeACLProtectTrendRequest) SetSourceIp(v string) *DescribeACLProtectTrendRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeACLProtectTrendRequest) SetStartTime(v string) *DescribeACLProtectTrendRequest {
	s.StartTime = &v
	return s
}

type DescribeACLProtectTrendResponseBody struct {
	// The number of internal requests that are blocked by the ACL feature.
	InProtectCnt *int64 `json:"InProtectCnt,omitempty" xml:"InProtectCnt,omitempty"`
	// This parameter is deprecated.
	InterVPCProtectCnt *int64 `json:"InterVPCProtectCnt,omitempty" xml:"InterVPCProtectCnt,omitempty"`
	// The interval for returning data. Unit: seconds.
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The number of external requests that are blocked by the ACL feature.
	OutProtectCnt *int64 `json:"OutProtectCnt,omitempty" xml:"OutProtectCnt,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of requests that are blocked by the ACL feature.
	TotalProtectCnt *int64 `json:"TotalProtectCnt,omitempty" xml:"TotalProtectCnt,omitempty"`
	// The statistics on the requests that are blocked by the ACL feature.
	TrendList []*DescribeACLProtectTrendResponseBodyTrendList `json:"TrendList,omitempty" xml:"TrendList,omitempty" type:"Repeated"`
}

func (s DescribeACLProtectTrendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeACLProtectTrendResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeACLProtectTrendResponseBody) SetInProtectCnt(v int64) *DescribeACLProtectTrendResponseBody {
	s.InProtectCnt = &v
	return s
}

func (s *DescribeACLProtectTrendResponseBody) SetInterVPCProtectCnt(v int64) *DescribeACLProtectTrendResponseBody {
	s.InterVPCProtectCnt = &v
	return s
}

func (s *DescribeACLProtectTrendResponseBody) SetInterval(v int32) *DescribeACLProtectTrendResponseBody {
	s.Interval = &v
	return s
}

func (s *DescribeACLProtectTrendResponseBody) SetOutProtectCnt(v int64) *DescribeACLProtectTrendResponseBody {
	s.OutProtectCnt = &v
	return s
}

func (s *DescribeACLProtectTrendResponseBody) SetRequestId(v string) *DescribeACLProtectTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeACLProtectTrendResponseBody) SetTotalProtectCnt(v int64) *DescribeACLProtectTrendResponseBody {
	s.TotalProtectCnt = &v
	return s
}

func (s *DescribeACLProtectTrendResponseBody) SetTrendList(v []*DescribeACLProtectTrendResponseBodyTrendList) *DescribeACLProtectTrendResponseBody {
	s.TrendList = v
	return s
}

type DescribeACLProtectTrendResponseBodyTrendList struct {
	// The number of requests that are blocked by ACL on the current day.
	ProtectCnt *int32 `json:"ProtectCnt,omitempty" xml:"ProtectCnt,omitempty"`
	// The UNIX timestamp at midnight (00:00:00) of each day, which indicates the date of the current day. Unit: seconds.
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DescribeACLProtectTrendResponseBodyTrendList) String() string {
	return tea.Prettify(s)
}

func (s DescribeACLProtectTrendResponseBodyTrendList) GoString() string {
	return s.String()
}

func (s *DescribeACLProtectTrendResponseBodyTrendList) SetProtectCnt(v int32) *DescribeACLProtectTrendResponseBodyTrendList {
	s.ProtectCnt = &v
	return s
}

func (s *DescribeACLProtectTrendResponseBodyTrendList) SetTime(v int64) *DescribeACLProtectTrendResponseBodyTrendList {
	s.Time = &v
	return s
}

type DescribeACLProtectTrendResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeACLProtectTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeACLProtectTrendResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeACLProtectTrendResponse) GoString() string {
	return s.String()
}

func (s *DescribeACLProtectTrendResponse) SetHeaders(v map[string]*string) *DescribeACLProtectTrendResponse {
	s.Headers = v
	return s
}

func (s *DescribeACLProtectTrendResponse) SetStatusCode(v int32) *DescribeACLProtectTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeACLProtectTrendResponse) SetBody(v *DescribeACLProtectTrendResponseBody) *DescribeACLProtectTrendResponse {
	s.Body = v
	return s
}

type DescribeAddressBookRequest struct {
	// The port that is included in the address book. This parameter takes effect only when the **GroupType** parameter is set to **port**.
	ContainPort *string `json:"ContainPort,omitempty" xml:"ContainPort,omitempty"`
	// The page number.
	//
	// Pages start from page 1. Default value: 1.
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The type of the address book. Valid values:
	//
	// *   **ip**: IP address book
	// *   **domain**: domain address book
	// *   **port**: port address book
	// *   **tag**: Elastic Compute Service (ECS) tag-based address book
	// *   **allCloud**: cloud service address book
	// *   **threat**: threat intelligence address book
	//
	// >  If you do not specify a type, the domain address books and ECS tag-based address books are queried.
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The language of the content within the request. Valid values:
	//
	// *   **zh** (default): Chinese
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries per page.
	//
	// Default value: 10. Maximum value: 50.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The query condition that is used to search for the address book.
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s DescribeAddressBookRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddressBookRequest) GoString() string {
	return s.String()
}

func (s *DescribeAddressBookRequest) SetContainPort(v string) *DescribeAddressBookRequest {
	s.ContainPort = &v
	return s
}

func (s *DescribeAddressBookRequest) SetCurrentPage(v string) *DescribeAddressBookRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAddressBookRequest) SetGroupType(v string) *DescribeAddressBookRequest {
	s.GroupType = &v
	return s
}

func (s *DescribeAddressBookRequest) SetLang(v string) *DescribeAddressBookRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAddressBookRequest) SetPageSize(v string) *DescribeAddressBookRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAddressBookRequest) SetQuery(v string) *DescribeAddressBookRequest {
	s.Query = &v
	return s
}

type DescribeAddressBookResponseBody struct {
	// The information about the address book.
	Acls []*DescribeAddressBookResponseBodyAcls `json:"Acls,omitempty" xml:"Acls,omitempty" type:"Repeated"`
	// The page number.
	PageNo *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the returned address books.
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAddressBookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddressBookResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAddressBookResponseBody) SetAcls(v []*DescribeAddressBookResponseBodyAcls) *DescribeAddressBookResponseBody {
	s.Acls = v
	return s
}

func (s *DescribeAddressBookResponseBody) SetPageNo(v string) *DescribeAddressBookResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribeAddressBookResponseBody) SetPageSize(v string) *DescribeAddressBookResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAddressBookResponseBody) SetRequestId(v string) *DescribeAddressBookResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAddressBookResponseBody) SetTotalCount(v string) *DescribeAddressBookResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeAddressBookResponseBodyAcls struct {
	// The addresses in the address book.
	AddressList []*string `json:"AddressList,omitempty" xml:"AddressList,omitempty" type:"Repeated"`
	// The number of addresses in the address book.
	AddressListCount *int32 `json:"AddressListCount,omitempty" xml:"AddressListCount,omitempty"`
	// Indicates whether the public IP addresses of ECS instances are automatically added to the address book if the instances match the specified tags. The setting takes effect on both newly purchased ECS instances whose tag settings are complete and ECS instances whose tag settings are modified. Valid values:
	//
	// *   **1**: yes
	// *   **0**: no
	AutoAddTagEcs *int32 `json:"AutoAddTagEcs,omitempty" xml:"AutoAddTagEcs,omitempty"`
	// The description of the address book.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the address book.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The type of the address book. Valid values:
	//
	// *   **ip**: IP address book
	// *   **domain**: domain address book
	// *   **port**: port address book
	// *   **tag**: ECS tag-based address book
	// *   **allCloud**: cloud service address book
	// *   **threat**: threat intelligence address book
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The UUID of the address book.
	GroupUuid *string `json:"GroupUuid,omitempty" xml:"GroupUuid,omitempty"`
	// The number of times that the address book is referenced.
	ReferenceCount *int32 `json:"ReferenceCount,omitempty" xml:"ReferenceCount,omitempty"`
	// The details about the ECS tags that can be automatically added to the address book.
	TagList []*DescribeAddressBookResponseBodyAclsTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The logical relationship among ECS tags. Valid values:
	//
	// *   **and**: Only the public IP addresses of ECS instances that match all the specified tags can be added to the address book.
	// *   **or**: The public IP addresses of ECS instances that match any of the specified tags can be added to the address book.
	TagRelation *string `json:"TagRelation,omitempty" xml:"TagRelation,omitempty"`
}

func (s DescribeAddressBookResponseBodyAcls) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddressBookResponseBodyAcls) GoString() string {
	return s.String()
}

func (s *DescribeAddressBookResponseBodyAcls) SetAddressList(v []*string) *DescribeAddressBookResponseBodyAcls {
	s.AddressList = v
	return s
}

func (s *DescribeAddressBookResponseBodyAcls) SetAddressListCount(v int32) *DescribeAddressBookResponseBodyAcls {
	s.AddressListCount = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAcls) SetAutoAddTagEcs(v int32) *DescribeAddressBookResponseBodyAcls {
	s.AutoAddTagEcs = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAcls) SetDescription(v string) *DescribeAddressBookResponseBodyAcls {
	s.Description = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAcls) SetGroupName(v string) *DescribeAddressBookResponseBodyAcls {
	s.GroupName = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAcls) SetGroupType(v string) *DescribeAddressBookResponseBodyAcls {
	s.GroupType = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAcls) SetGroupUuid(v string) *DescribeAddressBookResponseBodyAcls {
	s.GroupUuid = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAcls) SetReferenceCount(v int32) *DescribeAddressBookResponseBodyAcls {
	s.ReferenceCount = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAcls) SetTagList(v []*DescribeAddressBookResponseBodyAclsTagList) *DescribeAddressBookResponseBodyAcls {
	s.TagList = v
	return s
}

func (s *DescribeAddressBookResponseBodyAcls) SetTagRelation(v string) *DescribeAddressBookResponseBodyAcls {
	s.TagRelation = &v
	return s
}

type DescribeAddressBookResponseBodyAclsTagList struct {
	// The key of the ECS tag.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the ECS tag.
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeAddressBookResponseBodyAclsTagList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddressBookResponseBodyAclsTagList) GoString() string {
	return s.String()
}

func (s *DescribeAddressBookResponseBodyAclsTagList) SetTagKey(v string) *DescribeAddressBookResponseBodyAclsTagList {
	s.TagKey = &v
	return s
}

func (s *DescribeAddressBookResponseBodyAclsTagList) SetTagValue(v string) *DescribeAddressBookResponseBodyAclsTagList {
	s.TagValue = &v
	return s
}

type DescribeAddressBookResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAddressBookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAddressBookResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAddressBookResponse) GoString() string {
	return s.String()
}

func (s *DescribeAddressBookResponse) SetHeaders(v map[string]*string) *DescribeAddressBookResponse {
	s.Headers = v
	return s
}

func (s *DescribeAddressBookResponse) SetStatusCode(v int32) *DescribeAddressBookResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAddressBookResponse) SetBody(v *DescribeAddressBookResponseBody) *DescribeAddressBookResponse {
	s.Body = v
	return s
}

type DescribeAssetListRequest struct {
	// The page number. Valid values: 1 to 50.
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The IP version of the asset that is protected by Cloud Firewall. Valid values:
	//
	// *   **4**: IPv4 (default)
	// *   **6**: IPv6
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member that is added to Cloud Firewall.
	MemberUid *int64 `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The time when the asset was added. Valid values:
	//
	// *   **discovered in 1 hour**: within one hour.
	// *   **discovered in 1 day**: within one day.
	// *   **discovered in 7 days**: within seven days.
	NewResourceTag *string `json:"NewResourceTag,omitempty" xml:"NewResourceTag,omitempty"`
	// The number of entries per page. Valid values: 1 to 50.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of your Cloud Firewall.
	//
	// > For more information about the regions, see [Supported regions](~~195657~~).
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The type of the asset. Valid values:
	//
	// *   **BastionHostEgressIP**: the egress IP address of a bastion host
	// *   **BastionHostIngressIP**: the ingress IP address of a bastion host
	// *   **EcsEIP**: the elastic IP address (EIP) of an Elastic Compute Service (ECS) instance
	// *   **EcsPublicIP**: the public IP address of an ECS instance
	// *   **EIP**: the EIP
	// *   **EniEIP**: the EIP of an elastic network interface (ENI)
	// *   **NatEIP**: the EIP of a NAT gateway
	// *   **SlbEIP**: the EIP of a Server Load Balancer (SLB) instance or a Classic Load Balancer (CLB) instance
	// *   **SlbPublicIP**: the public IP address of an SLB instance or a CLB instance
	// *   **NatPublicIP**: the public IP address of a NAT gateway
	// *   **HAVIP**: the high-availability virtual IP address (HAVIP)
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The instance ID or IP address of the asset.
	SearchItem *string `json:"SearchItem,omitempty" xml:"SearchItem,omitempty"`
	// The status of the security group policy. Valid values:
	//
	// *   **pass**: delivered
	// *   **block**: undelivered
	// *   **unsupport**: unsupported
	//
	// > If you do not specify this parameter, the assets on which security group policies in all states take effect are queried.
	SgStatus *string `json:"SgStatus,omitempty" xml:"SgStatus,omitempty"`
	// The status of the firewall. Valid values:
	//
	// *   **open**: The firewall is enabled.
	// *   **opening**: The firewall is being enabled.
	// *   **closed**: The firewall is disabled.
	// *   **closing**: The firewall is being disabled.
	//
	// > If you do not specify this parameter, the assets that are configured for firewalls in all states are queried.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// This parameter is deprecated.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The edition of Cloud Firewall. Valid values:
	//
	// *   **buy**: a paid edition (default)
	// *   **free**: Free Edition
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s DescribeAssetListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssetListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAssetListRequest) SetCurrentPage(v string) *DescribeAssetListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAssetListRequest) SetIpVersion(v string) *DescribeAssetListRequest {
	s.IpVersion = &v
	return s
}

func (s *DescribeAssetListRequest) SetLang(v string) *DescribeAssetListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAssetListRequest) SetMemberUid(v int64) *DescribeAssetListRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeAssetListRequest) SetNewResourceTag(v string) *DescribeAssetListRequest {
	s.NewResourceTag = &v
	return s
}

func (s *DescribeAssetListRequest) SetPageSize(v string) *DescribeAssetListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAssetListRequest) SetRegionNo(v string) *DescribeAssetListRequest {
	s.RegionNo = &v
	return s
}

func (s *DescribeAssetListRequest) SetResourceType(v string) *DescribeAssetListRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeAssetListRequest) SetSearchItem(v string) *DescribeAssetListRequest {
	s.SearchItem = &v
	return s
}

func (s *DescribeAssetListRequest) SetSgStatus(v string) *DescribeAssetListRequest {
	s.SgStatus = &v
	return s
}

func (s *DescribeAssetListRequest) SetStatus(v string) *DescribeAssetListRequest {
	s.Status = &v
	return s
}

func (s *DescribeAssetListRequest) SetType(v string) *DescribeAssetListRequest {
	s.Type = &v
	return s
}

func (s *DescribeAssetListRequest) SetUserType(v string) *DescribeAssetListRequest {
	s.UserType = &v
	return s
}

type DescribeAssetListResponseBody struct {
	// The details about the assets that are protected by Cloud Firewall.
	Assets []*DescribeAssetListResponseBodyAssets `json:"Assets,omitempty" xml:"Assets,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the assets that are protected by Cloud Firewall.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAssetListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssetListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAssetListResponseBody) SetAssets(v []*DescribeAssetListResponseBodyAssets) *DescribeAssetListResponseBody {
	s.Assets = v
	return s
}

func (s *DescribeAssetListResponseBody) SetRequestId(v string) *DescribeAssetListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAssetListResponseBody) SetTotalCount(v int32) *DescribeAssetListResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeAssetListResponseBodyAssets struct {
	// The UID of the Alibaba Cloud account.
	//
	// >  The value of this parameter indicates the management account to which the member is added.
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The ID of the cloud resource with which the asset is associated.
	BindInstanceId *string `json:"BindInstanceId,omitempty" xml:"BindInstanceId,omitempty"`
	// The instance name of the asset.
	BindInstanceName *string `json:"BindInstanceName,omitempty" xml:"BindInstanceName,omitempty"`
	// The timestamp when the asset is added to Cloud Firewall.
	CreateTimeStamp *string `json:"CreateTimeStamp,omitempty" xml:"CreateTimeStamp,omitempty"`
	// The public IP address of the server.
	InternetAddress *string `json:"InternetAddress,omitempty" xml:"InternetAddress,omitempty"`
	// The internal IP address of the server.
	IntranetAddress *string `json:"IntranetAddress,omitempty" xml:"IntranetAddress,omitempty"`
	// The IP version of the asset that is protected by Cloud Firewall.
	//
	// Valid values:
	//
	// *   **4**: IPv4
	// *   **6**: IPv6
	IpVersion *int32 `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The UID of the member.
	MemberUid *int64 `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The instance name of the asset that is protected by Cloud Firewall.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The time when the asset was added. Valid values:
	//
	// *   **discovered in 1 hour**: within one hour.
	// *   **discovered in 1 day**: within one day.
	// *   **discovered in 7 days**: within seven days.
	NewResourceTag *string `json:"NewResourceTag,omitempty" xml:"NewResourceTag,omitempty"`
	// The remarks of the asset. Valid values:
	//
	// *   **REGION_NOT_SUPPORT**: The region is not supported.
	// *   **NETWORK_NOT_SUPPORT**: The network is not supported.
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The status of the firewall. Valid values:
	//
	// *   **open**: enabled.
	// *   **opening**: being enabled.
	// *   **closed**: disabled.
	// *   **closing**: being disabled.
	ProtectStatus *string `json:"ProtectStatus,omitempty" xml:"ProtectStatus,omitempty"`
	// The ID of the region in which the asset resides.
	RegionID *string `json:"RegionID,omitempty" xml:"RegionID,omitempty"`
	// Indicates whether the firewall is supported in the region in which the asset resides. Valid values:
	//
	// *   **enable**: yes
	// *   **disable**: no
	RegionStatus *string `json:"RegionStatus,omitempty" xml:"RegionStatus,omitempty"`
	// The instance ID of the asset.
	ResourceInstanceId *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
	// The type of the asset. Valid values:
	//
	// *   **BastionHostEgressIP**: the egress IP address of a bastion host
	// *   **BastionHostIngressIP**: the ingress IP address of a bastion host
	// *   **EcsEIP**: the elastic IP address (EIP) of an Elastic Compute Service (ECS) instance
	// *   **EcsPublicIP**: the public IP address of an ECS instance
	// *   **EIP**: the EIP
	// *   **EniEIP**: the EIP of an elastic network interface (ENI)
	// *   **NatEIP**: the EIP of a NAT gateway
	// *   **SlbEIP**: the EIP of a Server Load Balancer (SLB) instance
	// *   **SlbPublicIP**: the public IP address of an SLB instance
	// *   **NatPublicIP**: the public IP address of a NAT gateway
	// *   **HAVIP**: the high-availability virtual IP address (HAVIP)
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The risk level of the asset. Valid values:
	//
	// *   **low**: low
	// *   **middle**: medium
	// *   **hight**: high
	//
	// >  The value of this parameter is returned only when the UserType parameter is set to free.
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The status of the security group policy. Valid values:
	//
	// *   **pass**: applied
	// *   **block**: not applied
	// *   **unsupport**: unsupported
	SgStatus *string `json:"SgStatus,omitempty" xml:"SgStatus,omitempty"`
	// The time when the status of the security group was last checked. The value is a UNIX timestamp. Unit: seconds.
	SgStatusTime *int64 `json:"SgStatusTime,omitempty" xml:"SgStatusTime,omitempty"`
	// Indicates whether traffic redirection is supported for the asset. Valid values:
	//
	// *   **enable**: yes
	// *   **disable**: no
	SyncStatus *string `json:"SyncStatus,omitempty" xml:"SyncStatus,omitempty"`
	// This parameter is deprecated.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAssetListResponseBodyAssets) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssetListResponseBodyAssets) GoString() string {
	return s.String()
}

func (s *DescribeAssetListResponseBodyAssets) SetAliUid(v int64) *DescribeAssetListResponseBodyAssets {
	s.AliUid = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetBindInstanceId(v string) *DescribeAssetListResponseBodyAssets {
	s.BindInstanceId = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetBindInstanceName(v string) *DescribeAssetListResponseBodyAssets {
	s.BindInstanceName = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetCreateTimeStamp(v string) *DescribeAssetListResponseBodyAssets {
	s.CreateTimeStamp = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetInternetAddress(v string) *DescribeAssetListResponseBodyAssets {
	s.InternetAddress = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetIntranetAddress(v string) *DescribeAssetListResponseBodyAssets {
	s.IntranetAddress = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetIpVersion(v int32) *DescribeAssetListResponseBodyAssets {
	s.IpVersion = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetMemberUid(v int64) *DescribeAssetListResponseBodyAssets {
	s.MemberUid = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetName(v string) *DescribeAssetListResponseBodyAssets {
	s.Name = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetNewResourceTag(v string) *DescribeAssetListResponseBodyAssets {
	s.NewResourceTag = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetNote(v string) *DescribeAssetListResponseBodyAssets {
	s.Note = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetProtectStatus(v string) *DescribeAssetListResponseBodyAssets {
	s.ProtectStatus = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetRegionID(v string) *DescribeAssetListResponseBodyAssets {
	s.RegionID = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetRegionStatus(v string) *DescribeAssetListResponseBodyAssets {
	s.RegionStatus = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetResourceInstanceId(v string) *DescribeAssetListResponseBodyAssets {
	s.ResourceInstanceId = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetResourceType(v string) *DescribeAssetListResponseBodyAssets {
	s.ResourceType = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetRiskLevel(v string) *DescribeAssetListResponseBodyAssets {
	s.RiskLevel = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetSgStatus(v string) *DescribeAssetListResponseBodyAssets {
	s.SgStatus = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetSgStatusTime(v int64) *DescribeAssetListResponseBodyAssets {
	s.SgStatusTime = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetSyncStatus(v string) *DescribeAssetListResponseBodyAssets {
	s.SyncStatus = &v
	return s
}

func (s *DescribeAssetListResponseBodyAssets) SetType(v string) *DescribeAssetListResponseBodyAssets {
	s.Type = &v
	return s
}

type DescribeAssetListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAssetListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAssetListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssetListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAssetListResponse) SetHeaders(v map[string]*string) *DescribeAssetListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAssetListResponse) SetStatusCode(v int32) *DescribeAssetListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAssetListResponse) SetBody(v *DescribeAssetListResponseBody) *DescribeAssetListResponse {
	s.Body = v
	return s
}

type DescribeAssetRiskListRequest struct {
	// The IP addresses to query. Separate the IP addresses with commas (,). You can specify up to 20 IP addresses at a time.
	//
	// >
	//
	// *   Example of an IPv4 address: 47.97.221.164
	//
	// *   Example of an IPv6 address: 2001:db8:ffff:ffff:ffff:\*\*\*\*:ffff
	IpAddrList []*string `json:"IpAddrList,omitempty" xml:"IpAddrList,omitempty" type:"Repeated"`
	// The IP version of the asset that is protected by Cloud Firewall.
	//
	// Valid values:
	//
	// *   **4** (default): IPv4
	// *   **6**: IPv6
	IpVersion *int32 `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// *   **zh** (default): Chinese
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The source IP address of the request.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeAssetRiskListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssetRiskListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAssetRiskListRequest) SetIpAddrList(v []*string) *DescribeAssetRiskListRequest {
	s.IpAddrList = v
	return s
}

func (s *DescribeAssetRiskListRequest) SetIpVersion(v int32) *DescribeAssetRiskListRequest {
	s.IpVersion = &v
	return s
}

func (s *DescribeAssetRiskListRequest) SetLang(v string) *DescribeAssetRiskListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAssetRiskListRequest) SetSourceIp(v string) *DescribeAssetRiskListRequest {
	s.SourceIp = &v
	return s
}

type DescribeAssetRiskListResponseBody struct {
	// The details of the asset.
	AssetList []*DescribeAssetRiskListResponseBodyAssetList `json:"AssetList,omitempty" xml:"AssetList,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAssetRiskListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssetRiskListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAssetRiskListResponseBody) SetAssetList(v []*DescribeAssetRiskListResponseBodyAssetList) *DescribeAssetRiskListResponseBody {
	s.AssetList = v
	return s
}

func (s *DescribeAssetRiskListResponseBody) SetRequestId(v string) *DescribeAssetRiskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAssetRiskListResponseBody) SetTotalCount(v int64) *DescribeAssetRiskListResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeAssetRiskListResponseBodyAssetList struct {
	// The IP address of the server.
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The IP version of the asset that is protected by Cloud Firewall.
	//
	// Valid values:
	//
	// *   **4**: IPv4
	// *   **6**: IPv6
	IpVersion *int64 `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The reason for the risk.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The risk level. Valid values:
	//
	// *   **low**
	// *   **middle**
	// *   **high**
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s DescribeAssetRiskListResponseBodyAssetList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssetRiskListResponseBodyAssetList) GoString() string {
	return s.String()
}

func (s *DescribeAssetRiskListResponseBodyAssetList) SetIp(v string) *DescribeAssetRiskListResponseBodyAssetList {
	s.Ip = &v
	return s
}

func (s *DescribeAssetRiskListResponseBodyAssetList) SetIpVersion(v int64) *DescribeAssetRiskListResponseBodyAssetList {
	s.IpVersion = &v
	return s
}

func (s *DescribeAssetRiskListResponseBodyAssetList) SetReason(v string) *DescribeAssetRiskListResponseBodyAssetList {
	s.Reason = &v
	return s
}

func (s *DescribeAssetRiskListResponseBodyAssetList) SetRiskLevel(v string) *DescribeAssetRiskListResponseBodyAssetList {
	s.RiskLevel = &v
	return s
}

type DescribeAssetRiskListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAssetRiskListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAssetRiskListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAssetRiskListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAssetRiskListResponse) SetHeaders(v map[string]*string) *DescribeAssetRiskListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAssetRiskListResponse) SetStatusCode(v int32) *DescribeAssetRiskListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAssetRiskListResponse) SetBody(v *DescribeAssetRiskListResponseBody) *DescribeAssetRiskListResponse {
	s.Body = v
	return s
}

type DescribeCfwRiskLevelSummaryRequest struct {
	// The instance type.
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The language of the content within the response.
	//
	// Valid values:
	//
	// *   **zh** (default): Chinese
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region ID of your Cloud Firewall.
	//
	// >  For more information about Cloud Firewall supported regions, see [Supported regions](~~195657~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCfwRiskLevelSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCfwRiskLevelSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeCfwRiskLevelSummaryRequest) SetInstanceType(v string) *DescribeCfwRiskLevelSummaryRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeCfwRiskLevelSummaryRequest) SetLang(v string) *DescribeCfwRiskLevelSummaryRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCfwRiskLevelSummaryRequest) SetRegionId(v string) *DescribeCfwRiskLevelSummaryRequest {
	s.RegionId = &v
	return s
}

type DescribeCfwRiskLevelSummaryResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of risks.
	RiskList []*DescribeCfwRiskLevelSummaryResponseBodyRiskList `json:"RiskList,omitempty" xml:"RiskList,omitempty" type:"Repeated"`
}

func (s DescribeCfwRiskLevelSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCfwRiskLevelSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCfwRiskLevelSummaryResponseBody) SetRequestId(v string) *DescribeCfwRiskLevelSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCfwRiskLevelSummaryResponseBody) SetRiskList(v []*DescribeCfwRiskLevelSummaryResponseBodyRiskList) *DescribeCfwRiskLevelSummaryResponseBody {
	s.RiskList = v
	return s
}

type DescribeCfwRiskLevelSummaryResponseBodyRiskList struct {
	// The risk levels. Valid values:
	//
	// *   **medium**
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The number of at-risk Elastic Compute Service (ECS) instances.
	Num *string `json:"Num,omitempty" xml:"Num,omitempty"`
	// The type.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeCfwRiskLevelSummaryResponseBodyRiskList) String() string {
	return tea.Prettify(s)
}

func (s DescribeCfwRiskLevelSummaryResponseBodyRiskList) GoString() string {
	return s.String()
}

func (s *DescribeCfwRiskLevelSummaryResponseBodyRiskList) SetLevel(v string) *DescribeCfwRiskLevelSummaryResponseBodyRiskList {
	s.Level = &v
	return s
}

func (s *DescribeCfwRiskLevelSummaryResponseBodyRiskList) SetNum(v string) *DescribeCfwRiskLevelSummaryResponseBodyRiskList {
	s.Num = &v
	return s
}

func (s *DescribeCfwRiskLevelSummaryResponseBodyRiskList) SetType(v string) *DescribeCfwRiskLevelSummaryResponseBodyRiskList {
	s.Type = &v
	return s
}

type DescribeCfwRiskLevelSummaryResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCfwRiskLevelSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCfwRiskLevelSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCfwRiskLevelSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeCfwRiskLevelSummaryResponse) SetHeaders(v map[string]*string) *DescribeCfwRiskLevelSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeCfwRiskLevelSummaryResponse) SetStatusCode(v int32) *DescribeCfwRiskLevelSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCfwRiskLevelSummaryResponse) SetBody(v *DescribeCfwRiskLevelSummaryResponseBody) *DescribeCfwRiskLevelSummaryResponse {
	s.Body = v
	return s
}

type DescribeControlPolicyRequest struct {
	// The action that Cloud Firewall performs on the traffic. Valid values:
	//
	// *   **accept**: allows the traffic.
	// *   **drop**: denies the traffic.
	// *   **log**: monitors the traffic.
	//
	// >  If you do not specify this parameter, access control policies of all action types are queried.
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The unique ID of the access control policy.
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The number of the page to return.
	//
	// Default value: 1.
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The description of the access control policy. Fuzzy match is supported.
	//
	// >  If you do not specify this parameter, access control policies that have descriptions are queried.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination address in the access control policy. Fuzzy match is supported. The value of this parameter depends on the value of the DestinationType parameter.
	//
	// *   If DestinationType is set to `net`, the value of Destination must be a CIDR block. Example: 10.0.3.0/24.
	// *   If DestinationType is set to `domain`, the value of Destination must be a domain name. Example: aliyun.
	// *   If DestinationType is set to `group`, the value of Destination must be the name of an address book. Example: db_group.
	// *   If DestinationType is set to `location`, the value of Destination must be a location. Example: beijing.
	//
	// >  If you do not specify this parameter, access control policies of all destination address types are queried.
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The direction of the traffic to which the access control policy applies. Valid values:
	//
	// *   **in**: inbound traffic
	// *   **out**: outbound traffic
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The IP version of the address in the access control policy. Valid values:
	//
	// *   **4**: IPv4 (default)
	// *   **6**: IPv6
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries to return on each page.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the protocol in the access control policy. Valid values:
	//
	// * **TCP**
	// * **UDP**
	// * **ICMP**
	// * **ANY**: all types of protocols
	//
	// >  If you do not specify this parameter, access control policies of all protocol types are queried.
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// Specifies whether the access control policy is enabled. By default, an access control policy is enabled after it is created. Valid values:
	//
	// *   **true**: The access control policy is enabled.
	// *   **false**: The access control policy is disabled.
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The recurrence type for the access control policy to take effect. Valid values:
	//
	// *   **Permanent** (default): The policy always takes effect.
	// *   **None**: The policy takes effect for only once.
	// *   **Daily**: The policy takes effect on a daily basis.
	// *   **Weekly**: The policy takes effect on a weekly basis.
	// *   **Monthly**: The policy takes effect on a monthly basis.
	RepeatType *string `json:"RepeatType,omitempty" xml:"RepeatType,omitempty"`
	// The source address in the access control policy. Fuzzy match is supported. The value of this parameter depends on the value of the SourceType parameter.
	//
	// *   If SourceType is set to `net`, the value of Source must be a CIDR block. Example: 192.0.XX.XX/24.
	// *   If SourceType is set to `group`, the value of Source must be the name of an address book. Example: db_group. If the db_group address book does not contain addresses, all source addresses are queried.
	// *   If SourceType is set to `location`, the value of Source must be a location. Example: beijing.
	//
	// >  If you do not specify this parameter, access control policies of all source address types are queried.
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s DescribeControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeControlPolicyRequest) SetAclAction(v string) *DescribeControlPolicyRequest {
	s.AclAction = &v
	return s
}

func (s *DescribeControlPolicyRequest) SetAclUuid(v string) *DescribeControlPolicyRequest {
	s.AclUuid = &v
	return s
}

func (s *DescribeControlPolicyRequest) SetCurrentPage(v string) *DescribeControlPolicyRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeControlPolicyRequest) SetDescription(v string) *DescribeControlPolicyRequest {
	s.Description = &v
	return s
}

func (s *DescribeControlPolicyRequest) SetDestination(v string) *DescribeControlPolicyRequest {
	s.Destination = &v
	return s
}

func (s *DescribeControlPolicyRequest) SetDirection(v string) *DescribeControlPolicyRequest {
	s.Direction = &v
	return s
}

func (s *DescribeControlPolicyRequest) SetIpVersion(v string) *DescribeControlPolicyRequest {
	s.IpVersion = &v
	return s
}

func (s *DescribeControlPolicyRequest) SetLang(v string) *DescribeControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *DescribeControlPolicyRequest) SetPageSize(v string) *DescribeControlPolicyRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeControlPolicyRequest) SetProto(v string) *DescribeControlPolicyRequest {
	s.Proto = &v
	return s
}

func (s *DescribeControlPolicyRequest) SetRelease(v string) *DescribeControlPolicyRequest {
	s.Release = &v
	return s
}

func (s *DescribeControlPolicyRequest) SetRepeatType(v string) *DescribeControlPolicyRequest {
	s.RepeatType = &v
	return s
}

func (s *DescribeControlPolicyRequest) SetSource(v string) *DescribeControlPolicyRequest {
	s.Source = &v
	return s
}

type DescribeControlPolicyResponseBody struct {
	// The page number of the returned page.
	PageNo *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries returned per page.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The information about the access control policies.
	Policys []*DescribeControlPolicyResponseBodyPolicys `json:"Policys,omitempty" xml:"Policys,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the returned access control policies.
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeControlPolicyResponseBody) SetPageNo(v string) *DescribeControlPolicyResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribeControlPolicyResponseBody) SetPageSize(v string) *DescribeControlPolicyResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeControlPolicyResponseBody) SetPolicys(v []*DescribeControlPolicyResponseBodyPolicys) *DescribeControlPolicyResponseBody {
	s.Policys = v
	return s
}

func (s *DescribeControlPolicyResponseBody) SetRequestId(v string) *DescribeControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeControlPolicyResponseBody) SetTotalCount(v string) *DescribeControlPolicyResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeControlPolicyResponseBodyPolicys struct {
	// The action that Cloud Firewall performs on the traffic. Valid values:
	//
	// *   **accept**: allows the traffic.
	// *   **drop**: denies the traffic.
	// *   **log**: monitors the traffic.
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The UUID of the access control policy.
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The application ID in the access control policy.
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The application type supported by the access control policy. We recommend that you specify ApplicationNameList. Valid values:
	//
	// *   **FTP**
	// *   **HTTP**
	// *   **HTTPS**
	// *   **Memcache**
	// *   **MongoDB**
	// *   **MQTT**
	// *   **MySQL**
	// *   **RDP**
	// *   **Redis**
	// *   **SMTP**
	// *   **SMTPS**
	// *   **SSH**
	// *   **SSL**
	// *   **VNC**
	// *   **ANY**: all types of applications
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The application names.
	ApplicationNameList []*string `json:"ApplicationNameList,omitempty" xml:"ApplicationNameList,omitempty" type:"Repeated"`
	// The time when the access control policy was created.
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the access control policy.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination port in the access control policy.
	DestPort *string `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	// The name of the destination port address book in the access control policy.
	DestPortGroup *string `json:"DestPortGroup,omitempty" xml:"DestPortGroup,omitempty"`
	// The ports in the destination port address book.
	DestPortGroupPorts []*string `json:"DestPortGroupPorts,omitempty" xml:"DestPortGroupPorts,omitempty" type:"Repeated"`
	// The type of the destination port in the access control policy. Valid values:
	//
	// *   **port**: port
	// *   **group**: port address book
	DestPortType *string `json:"DestPortType,omitempty" xml:"DestPortType,omitempty"`
	// The destination address in the access control policy. The value of this parameter varies based on the value of DestinationType. Valid values:
	//
	// *   If **DestinationType** is set to **net**, the value of Destination is a CIDR block. Example: 192.0.XX.XX/24.
	// *   If **DestinationType** is set to **domain**, the value of Destination is a domain name. Example: aliyuncs.com.
	// *   If **DestinationType** is set to **group**, the value of Destination is the name of an address book. Example: db_group.
	// *   If **DestinationType** is set to **location**, the value of Destination is a location. For more information about location codes, see [AddControlPolicy](~~138867~~). Example: \["BJ11", "ZB"].
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The CIDR blocks in the destination address book.
	DestinationGroupCidrs []*string `json:"DestinationGroupCidrs,omitempty" xml:"DestinationGroupCidrs,omitempty" type:"Repeated"`
	// The type of the destination address book in the access control policy. Valid values:
	//
	// *   **ip**: an address book that includes one or more IP addresses
	// *   **tag**: an ECS tag-based address book that includes the IP addresses of the ECS instances with one or more specific tags
	// *   **domain**: an address book that includes one or more domain names
	// *   **threat**: an address book that includes one or more malicious IP addresses or domain names
	// *   **backsrc**: an address book that includes one or more back-to-origin addresses of Anti-DDoS Pro or Anti-DDoS Premium instances or WAF instances
	DestinationGroupType *string `json:"DestinationGroupType,omitempty" xml:"DestinationGroupType,omitempty"`
	// The type of the destination address in the access control policy. Valid values:
	//
	// *   **net**: CIDR block
	// *   **group**: address book
	// *   **domain**: domain name
	// *   **location**: location
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The direction of the traffic to which the access control policy applies. Valid values:
	//
	// *   **in**: inbound traffic
	// *   **out**: outbound traffic
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The DNS resolution result.
	DnsResult *string `json:"DnsResult,omitempty" xml:"DnsResult,omitempty"`
	// The time when the Domain Name System (DNS) resolution was performed. The value is a timestamp. Unit: seconds.
	DnsResultTime *int64 `json:"DnsResultTime,omitempty" xml:"DnsResultTime,omitempty"`
	// The time when the access control policy stops taking effect. The value is a timestamp. Unit: seconds. The end time must be on the hour or on the half hour, and at least 30 minutes later than the start time.
	//
	// >  If RepeatType is set to Permanent, this parameter is left empty. If RepeatType is set to None, Daily, Weekly, or Monthly, this parameter must be specified.
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time when the access control policy was last hit. The value is a timestamp. Unit: seconds.
	HitLastTime *int64 `json:"HitLastTime,omitempty" xml:"HitLastTime,omitempty"`
	// The number of hits for the access control policy.
	HitTimes *int64 `json:"HitTimes,omitempty" xml:"HitTimes,omitempty"`
	// The IP version used in the access control policy. Valid values:
	//
	// *   **4**: IPv4
	// *   **6**: IPv6
	IpVersion *int32 `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The time when the access control policy was modified.
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The priority of the access control policy.
	//
	// The priority value starts from 1. A smaller priority value indicates a higher priority.
	Order *int32 `json:"Order,omitempty" xml:"Order,omitempty"`
	// The protocol type in the access control policy. Valid values:
	//
	// *   **ANY**
	// *   **TCP**
	// *   **UDP**
	// *   **ICMP**
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// The status of the access control policy. By default, an access control policy is enabled after it is created. Valid values:
	//
	// *   **true**: enabled
	// *   **false**: disabled
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The days of a week or of a month on which the access control policy takes effect.
	//
	// *   If RepeatType is set to `Permanent`, `None`, or `Daily`, this parameter is left empty. Example: \[].
	// *   If RepeatType is set to Weekly, this parameter must be specified. Example: \[0, 6].
	//
	// >  If RepeatType is set to Weekly, the fields in the value of RepeatDays cannot be repeated.
	//
	// *   If RepeatType is set to `Monthly`, this parameter must be specified. Example: \[1, 31].
	//
	// >  If RepeatType is set to Monthly, the fields in the value of RepeatDays cannot be repeated.
	RepeatDays []*int64 `json:"RepeatDays,omitempty" xml:"RepeatDays,omitempty" type:"Repeated"`
	// The point in time when the recurrence ends. Example: 23:30. The value must be on the hour or on the half hour, and at least 30 minutes later than the start time.
	//
	// >  If RepeatType is set to Permanent or None, this parameter is left empty. If RepeatType is set to Daily, Weekly, or Monthly, this parameter must be specified.
	RepeatEndTime *string `json:"RepeatEndTime,omitempty" xml:"RepeatEndTime,omitempty"`
	// The point in time when the recurrence starts. Example: 08:00. The value must be on the hour or on the half hour, and at least 30 minutes earlier than the end time.
	//
	// >  If RepeatType is set to Permanent or None, this parameter is left empty. If RepeatType is set to Daily, Weekly, or Monthly, this parameter must be specified.
	RepeatStartTime *string `json:"RepeatStartTime,omitempty" xml:"RepeatStartTime,omitempty"`
	// The recurrence type based on which the access control policy takes effect. Valid values:
	//
	// *   **Permanent** (default): The policy always takes effect.
	// *   **None**: The policy takes effect for only once.
	// *   **Daily**: The policy takes effect on a daily basis.
	// *   **Weekly**: The policy takes effect on a weekly basis.
	// *   **Monthly**: The policy takes effect on a monthly basis.
	RepeatType *string `json:"RepeatType,omitempty" xml:"RepeatType,omitempty"`
	// The source address in the access control policy. Valid values:
	//
	// *   If **SourceType** is set to `net`, the value of Source is a CIDR block. Example: 192.0.XX.XX/24.
	// *   If **SourceType** is set to `group`, the value of Source is the name of an address book. Example: db_group.
	// *   If **SourceType** is set to `location`, the value of Source is a location. For more information about location codes, see [AddControlPolicy](~~138867~~). Example: \["BJ11", "ZB"].
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The CIDR blocks in the source address book.
	SourceGroupCidrs []*string `json:"SourceGroupCidrs,omitempty" xml:"SourceGroupCidrs,omitempty" type:"Repeated"`
	// The type of the source address book in the access control policy. Valid values:
	//
	// *   **ip**: an address book that includes one or more IP addresses
	// *   **tag**: an Elastic Compute Service (ECS) tag-based address book that includes the IP addresses of the ECS instances with one or more specific tags
	// *   **domain**: an address book that includes one or more domain names
	// *   **threat**: an address book that includes one or more malicious IP addresses or domain names
	// *   **backsrc**: an address book that includes one or more back-to-origin addresses of Anti-DDoS Pro or Anti-DDoS Premium instances or Web Application Firewall (WAF) instances
	SourceGroupType *string `json:"SourceGroupType,omitempty" xml:"SourceGroupType,omitempty"`
	// The type of the source address in the access control policy. Valid values:
	//
	// *   **net**: CIDR block
	// *   **group**: address book
	// *   **location**: location
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The total quota consumed by the returned access control policies, which is the sum of the quota consumed by each policy. The quota that is consumed by an access control policy is calculated by using the following formula: Quota that is consumed by an access control policy = Number of source addresses (number of CIDR blocks or regions) × Number of destination addresses (number of CIDR blocks, regions, or domain names) × Number of port ranges × Number of applications.
	SpreadCnt *int32 `json:"SpreadCnt,omitempty" xml:"SpreadCnt,omitempty"`
	// The time when the access control policy starts to take effect. The value is a timestamp. Unit: seconds. The start time must be on the hour or on the half hour, and at least 30 minutes earlier than the end time.
	//
	// >  If RepeatType is set to Permanent, this parameter is left empty. If RepeatType is set to None, Daily, Weekly, or Monthly, this parameter must be specified.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeControlPolicyResponseBodyPolicys) String() string {
	return tea.Prettify(s)
}

func (s DescribeControlPolicyResponseBodyPolicys) GoString() string {
	return s.String()
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetAclAction(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.AclAction = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetAclUuid(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.AclUuid = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetApplicationId(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.ApplicationId = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetApplicationName(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.ApplicationName = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetApplicationNameList(v []*string) *DescribeControlPolicyResponseBodyPolicys {
	s.ApplicationNameList = v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetCreateTime(v int64) *DescribeControlPolicyResponseBodyPolicys {
	s.CreateTime = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetDescription(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.Description = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetDestPort(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.DestPort = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetDestPortGroup(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.DestPortGroup = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetDestPortGroupPorts(v []*string) *DescribeControlPolicyResponseBodyPolicys {
	s.DestPortGroupPorts = v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetDestPortType(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.DestPortType = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetDestination(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.Destination = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetDestinationGroupCidrs(v []*string) *DescribeControlPolicyResponseBodyPolicys {
	s.DestinationGroupCidrs = v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetDestinationGroupType(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.DestinationGroupType = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetDestinationType(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.DestinationType = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetDirection(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.Direction = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetDnsResult(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.DnsResult = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetDnsResultTime(v int64) *DescribeControlPolicyResponseBodyPolicys {
	s.DnsResultTime = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetEndTime(v int64) *DescribeControlPolicyResponseBodyPolicys {
	s.EndTime = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetHitLastTime(v int64) *DescribeControlPolicyResponseBodyPolicys {
	s.HitLastTime = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetHitTimes(v int64) *DescribeControlPolicyResponseBodyPolicys {
	s.HitTimes = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetIpVersion(v int32) *DescribeControlPolicyResponseBodyPolicys {
	s.IpVersion = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetModifyTime(v int64) *DescribeControlPolicyResponseBodyPolicys {
	s.ModifyTime = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetOrder(v int32) *DescribeControlPolicyResponseBodyPolicys {
	s.Order = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetProto(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.Proto = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetRelease(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.Release = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetRepeatDays(v []*int64) *DescribeControlPolicyResponseBodyPolicys {
	s.RepeatDays = v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetRepeatEndTime(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.RepeatEndTime = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetRepeatStartTime(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.RepeatStartTime = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetRepeatType(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.RepeatType = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetSource(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.Source = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetSourceGroupCidrs(v []*string) *DescribeControlPolicyResponseBodyPolicys {
	s.SourceGroupCidrs = v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetSourceGroupType(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.SourceGroupType = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetSourceType(v string) *DescribeControlPolicyResponseBodyPolicys {
	s.SourceType = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetSpreadCnt(v int32) *DescribeControlPolicyResponseBodyPolicys {
	s.SpreadCnt = &v
	return s
}

func (s *DescribeControlPolicyResponseBodyPolicys) SetStartTime(v int64) *DescribeControlPolicyResponseBodyPolicys {
	s.StartTime = &v
	return s
}

type DescribeControlPolicyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeControlPolicyResponse) SetHeaders(v map[string]*string) *DescribeControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeControlPolicyResponse) SetStatusCode(v int32) *DescribeControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeControlPolicyResponse) SetBody(v *DescribeControlPolicyResponseBody) *DescribeControlPolicyResponse {
	s.Body = v
	return s
}

type DescribeDefaultIPSConfigRequest struct {
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeDefaultIPSConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefaultIPSConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefaultIPSConfigRequest) SetLang(v string) *DescribeDefaultIPSConfigRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDefaultIPSConfigRequest) SetSourceIp(v string) *DescribeDefaultIPSConfigRequest {
	s.SourceIp = &v
	return s
}

type DescribeDefaultIPSConfigResponseBody struct {
	AiRules        *int32  `json:"AiRules,omitempty" xml:"AiRules,omitempty"`
	BasicRules     *int32  `json:"BasicRules,omitempty" xml:"BasicRules,omitempty"`
	CtiRules       *int32  `json:"CtiRules,omitempty" xml:"CtiRules,omitempty"`
	EnableAllPatch *int32  `json:"EnableAllPatch,omitempty" xml:"EnableAllPatch,omitempty"`
	EnableDefault  *int32  `json:"EnableDefault,omitempty" xml:"EnableDefault,omitempty"`
	PatchRules     *int32  `json:"PatchRules,omitempty" xml:"PatchRules,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RuleClass      *int32  `json:"RuleClass,omitempty" xml:"RuleClass,omitempty"`
	RunMode        *int32  `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
}

func (s DescribeDefaultIPSConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefaultIPSConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefaultIPSConfigResponseBody) SetAiRules(v int32) *DescribeDefaultIPSConfigResponseBody {
	s.AiRules = &v
	return s
}

func (s *DescribeDefaultIPSConfigResponseBody) SetBasicRules(v int32) *DescribeDefaultIPSConfigResponseBody {
	s.BasicRules = &v
	return s
}

func (s *DescribeDefaultIPSConfigResponseBody) SetCtiRules(v int32) *DescribeDefaultIPSConfigResponseBody {
	s.CtiRules = &v
	return s
}

func (s *DescribeDefaultIPSConfigResponseBody) SetEnableAllPatch(v int32) *DescribeDefaultIPSConfigResponseBody {
	s.EnableAllPatch = &v
	return s
}

func (s *DescribeDefaultIPSConfigResponseBody) SetEnableDefault(v int32) *DescribeDefaultIPSConfigResponseBody {
	s.EnableDefault = &v
	return s
}

func (s *DescribeDefaultIPSConfigResponseBody) SetPatchRules(v int32) *DescribeDefaultIPSConfigResponseBody {
	s.PatchRules = &v
	return s
}

func (s *DescribeDefaultIPSConfigResponseBody) SetRequestId(v string) *DescribeDefaultIPSConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefaultIPSConfigResponseBody) SetRuleClass(v int32) *DescribeDefaultIPSConfigResponseBody {
	s.RuleClass = &v
	return s
}

func (s *DescribeDefaultIPSConfigResponseBody) SetRunMode(v int32) *DescribeDefaultIPSConfigResponseBody {
	s.RunMode = &v
	return s
}

type DescribeDefaultIPSConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDefaultIPSConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDefaultIPSConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDefaultIPSConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefaultIPSConfigResponse) SetHeaders(v map[string]*string) *DescribeDefaultIPSConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefaultIPSConfigResponse) SetStatusCode(v int32) *DescribeDefaultIPSConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefaultIPSConfigResponse) SetBody(v *DescribeDefaultIPSConfigResponseBody) *DescribeDefaultIPSConfigResponse {
	s.Body = v
	return s
}

type DescribeDomainResolveRequest struct {
	// The domain name whose DNS record you want to query.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The IP version of the asset that is protected by Cloud Firewall. Valid values:
	//
	// *   **4**: IPv4 (default)
	// *   **6**: IPv6
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The language of the content within the response.
	//
	// Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Deprecated
	// The source IP address of the request.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeDomainResolveRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainResolveRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainResolveRequest) SetDomain(v string) *DescribeDomainResolveRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainResolveRequest) SetIpVersion(v string) *DescribeDomainResolveRequest {
	s.IpVersion = &v
	return s
}

func (s *DescribeDomainResolveRequest) SetLang(v string) *DescribeDomainResolveRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDomainResolveRequest) SetSourceIp(v string) *DescribeDomainResolveRequest {
	s.SourceIp = &v
	return s
}

type DescribeDomainResolveResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details about the DNS record of the domain name.
	ResolveResult *DescribeDomainResolveResponseBodyResolveResult `json:"ResolveResult,omitempty" xml:"ResolveResult,omitempty" type:"Struct"`
}

func (s DescribeDomainResolveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainResolveResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainResolveResponseBody) SetRequestId(v string) *DescribeDomainResolveResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainResolveResponseBody) SetResolveResult(v *DescribeDomainResolveResponseBodyResolveResult) *DescribeDomainResolveResponseBody {
	s.ResolveResult = v
	return s
}

type DescribeDomainResolveResponseBodyResolveResult struct {
	// The IP address to which the domain name is resolved. Multiple IP addresses are separated by commas (,).
	IpAddrs *string `json:"IpAddrs,omitempty" xml:"IpAddrs,omitempty"`
	// The time when the domain name was resolved. The value of this parameter is a timestamp. Unit: seconds.
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeDomainResolveResponseBodyResolveResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainResolveResponseBodyResolveResult) GoString() string {
	return s.String()
}

func (s *DescribeDomainResolveResponseBodyResolveResult) SetIpAddrs(v string) *DescribeDomainResolveResponseBodyResolveResult {
	s.IpAddrs = &v
	return s
}

func (s *DescribeDomainResolveResponseBodyResolveResult) SetUpdateTime(v int64) *DescribeDomainResolveResponseBodyResolveResult {
	s.UpdateTime = &v
	return s
}

type DescribeDomainResolveResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainResolveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainResolveResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainResolveResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainResolveResponse) SetHeaders(v map[string]*string) *DescribeDomainResolveResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainResolveResponse) SetStatusCode(v int32) *DescribeDomainResolveResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainResolveResponse) SetBody(v *DescribeDomainResolveResponseBody) *DescribeDomainResolveResponse {
	s.Body = v
	return s
}

type DescribeDownloadTaskRequest struct {
	// The page number.
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// *   **zh** (default): Chinese
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 50.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the task. For more information about task types, see the descriptions in the "DescribeDownloadTaskType" topic. If you do not specify this parameter, all files are queried by default.
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeDownloadTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeDownloadTaskRequest) SetCurrentPage(v string) *DescribeDownloadTaskRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDownloadTaskRequest) SetLang(v string) *DescribeDownloadTaskRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDownloadTaskRequest) SetPageSize(v string) *DescribeDownloadTaskRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDownloadTaskRequest) SetTaskType(v string) *DescribeDownloadTaskRequest {
	s.TaskType = &v
	return s
}

type DescribeDownloadTaskResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tasks.
	Tasks []*DescribeDownloadTaskResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	// The total number of tasks.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDownloadTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDownloadTaskResponseBody) SetRequestId(v string) *DescribeDownloadTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDownloadTaskResponseBody) SetTasks(v []*DescribeDownloadTaskResponseBodyTasks) *DescribeDownloadTaskResponseBody {
	s.Tasks = v
	return s
}

func (s *DescribeDownloadTaskResponseBody) SetTotalCount(v int32) *DescribeDownloadTaskResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDownloadTaskResponseBodyTasks struct {
	// The time when the task was created. The value is a UNIX timestamp. Unit: seconds.
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the task expires. The value is a UNIX timestamp. Unit: seconds.
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The size of the file.
	FileSize *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The URL of the OSS file.
	FileURL *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	// The status of the task. Valid values:
	//
	// *   **finish**
	// *   **start**
	// *   **error**
	// *   **expire**: The task file is invalid and cannot be downloaded.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the task.
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The type of the task.
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeDownloadTaskResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadTaskResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *DescribeDownloadTaskResponseBodyTasks) SetCreateTime(v int64) *DescribeDownloadTaskResponseBodyTasks {
	s.CreateTime = &v
	return s
}

func (s *DescribeDownloadTaskResponseBodyTasks) SetExpireTime(v int64) *DescribeDownloadTaskResponseBodyTasks {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDownloadTaskResponseBodyTasks) SetFileSize(v string) *DescribeDownloadTaskResponseBodyTasks {
	s.FileSize = &v
	return s
}

func (s *DescribeDownloadTaskResponseBodyTasks) SetFileURL(v string) *DescribeDownloadTaskResponseBodyTasks {
	s.FileURL = &v
	return s
}

func (s *DescribeDownloadTaskResponseBodyTasks) SetStatus(v string) *DescribeDownloadTaskResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *DescribeDownloadTaskResponseBodyTasks) SetTaskId(v string) *DescribeDownloadTaskResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *DescribeDownloadTaskResponseBodyTasks) SetTaskName(v string) *DescribeDownloadTaskResponseBodyTasks {
	s.TaskName = &v
	return s
}

func (s *DescribeDownloadTaskResponseBodyTasks) SetTaskType(v string) *DescribeDownloadTaskResponseBodyTasks {
	s.TaskType = &v
	return s
}

type DescribeDownloadTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDownloadTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDownloadTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeDownloadTaskResponse) SetHeaders(v map[string]*string) *DescribeDownloadTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeDownloadTaskResponse) SetStatusCode(v int32) *DescribeDownloadTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDownloadTaskResponse) SetBody(v *DescribeDownloadTaskResponseBody) *DescribeDownloadTaskResponse {
	s.Body = v
	return s
}

type DescribeDownloadTaskTypeRequest struct {
	// The page number. Pages start from page 1. Default value: **1**.
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// *   **zh** (default): Chinese
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 50.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the task.
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeDownloadTaskTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadTaskTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDownloadTaskTypeRequest) SetCurrentPage(v string) *DescribeDownloadTaskTypeRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDownloadTaskTypeRequest) SetLang(v string) *DescribeDownloadTaskTypeRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDownloadTaskTypeRequest) SetPageSize(v string) *DescribeDownloadTaskTypeRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDownloadTaskTypeRequest) SetTaskType(v string) *DescribeDownloadTaskTypeRequest {
	s.TaskType = &v
	return s
}

type DescribeDownloadTaskTypeResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task types.
	TaskTypeArray []*DescribeDownloadTaskTypeResponseBodyTaskTypeArray `json:"TaskTypeArray,omitempty" xml:"TaskTypeArray,omitempty" type:"Repeated"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDownloadTaskTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadTaskTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDownloadTaskTypeResponseBody) SetRequestId(v string) *DescribeDownloadTaskTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDownloadTaskTypeResponseBody) SetTaskTypeArray(v []*DescribeDownloadTaskTypeResponseBodyTaskTypeArray) *DescribeDownloadTaskTypeResponseBody {
	s.TaskTypeArray = v
	return s
}

func (s *DescribeDownloadTaskTypeResponseBody) SetTotalCount(v int32) *DescribeDownloadTaskTypeResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDownloadTaskTypeResponseBodyTaskTypeArray struct {
	// The name of the task type.
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The type of the task.
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeDownloadTaskTypeResponseBodyTaskTypeArray) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadTaskTypeResponseBodyTaskTypeArray) GoString() string {
	return s.String()
}

func (s *DescribeDownloadTaskTypeResponseBodyTaskTypeArray) SetTaskName(v string) *DescribeDownloadTaskTypeResponseBodyTaskTypeArray {
	s.TaskName = &v
	return s
}

func (s *DescribeDownloadTaskTypeResponseBodyTaskTypeArray) SetTaskType(v string) *DescribeDownloadTaskTypeResponseBodyTaskTypeArray {
	s.TaskType = &v
	return s
}

type DescribeDownloadTaskTypeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDownloadTaskTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDownloadTaskTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDownloadTaskTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDownloadTaskTypeResponse) SetHeaders(v map[string]*string) *DescribeDownloadTaskTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeDownloadTaskTypeResponse) SetStatusCode(v int32) *DescribeDownloadTaskTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDownloadTaskTypeResponse) SetBody(v *DescribeDownloadTaskTypeResponseBody) *DescribeDownloadTaskTypeResponse {
	s.Body = v
	return s
}

type DescribeInstanceMembersRequest struct {
	// The page number. Default value: **1**.
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The remarks of the member. The remarks must be 1 to 256 characters in length.
	MemberDesc *string `json:"MemberDesc,omitempty" xml:"MemberDesc,omitempty"`
	// The name of the member.
	MemberDisplayName *string `json:"MemberDisplayName,omitempty" xml:"MemberDisplayName,omitempty"`
	// The UID of the member.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The number of entries per page.
	//
	// Default value: **20**.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeInstanceMembersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceMembersRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMembersRequest) SetCurrentPage(v string) *DescribeInstanceMembersRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInstanceMembersRequest) SetMemberDesc(v string) *DescribeInstanceMembersRequest {
	s.MemberDesc = &v
	return s
}

func (s *DescribeInstanceMembersRequest) SetMemberDisplayName(v string) *DescribeInstanceMembersRequest {
	s.MemberDisplayName = &v
	return s
}

func (s *DescribeInstanceMembersRequest) SetMemberUid(v string) *DescribeInstanceMembersRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeInstanceMembersRequest) SetPageSize(v string) *DescribeInstanceMembersRequest {
	s.PageSize = &v
	return s
}

type DescribeInstanceMembersResponseBody struct {
	// The information about the member.
	Members []*DescribeInstanceMembersResponseBodyMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeInstanceMembersResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceMembersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceMembersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMembersResponseBody) SetMembers(v []*DescribeInstanceMembersResponseBodyMembers) *DescribeInstanceMembersResponseBody {
	s.Members = v
	return s
}

func (s *DescribeInstanceMembersResponseBody) SetPageInfo(v *DescribeInstanceMembersResponseBodyPageInfo) *DescribeInstanceMembersResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeInstanceMembersResponseBody) SetRequestId(v string) *DescribeInstanceMembersResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInstanceMembersResponseBodyMembers struct {
	// The time when the member was added to Cloud Firewall. The value is a timestamp. Unit: seconds.
	CreateTime *int32 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The remarks of the member.
	MemberDesc *string `json:"MemberDesc,omitempty" xml:"MemberDesc,omitempty"`
	// The name of the member.
	MemberDisplayName *string `json:"MemberDisplayName,omitempty" xml:"MemberDisplayName,omitempty"`
	// The status of the member. Valid values:
	//
	// *   **normal**
	// *   **deleting**
	MemberStatus *string `json:"MemberStatus,omitempty" xml:"MemberStatus,omitempty"`
	// The UID of the member.
	MemberUid *int64 `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The time when the member was last modified. The value is a timestamp. Unit: seconds.
	ModifyTime *int32 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
}

func (s DescribeInstanceMembersResponseBodyMembers) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceMembersResponseBodyMembers) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMembersResponseBodyMembers) SetCreateTime(v int32) *DescribeInstanceMembersResponseBodyMembers {
	s.CreateTime = &v
	return s
}

func (s *DescribeInstanceMembersResponseBodyMembers) SetMemberDesc(v string) *DescribeInstanceMembersResponseBodyMembers {
	s.MemberDesc = &v
	return s
}

func (s *DescribeInstanceMembersResponseBodyMembers) SetMemberDisplayName(v string) *DescribeInstanceMembersResponseBodyMembers {
	s.MemberDisplayName = &v
	return s
}

func (s *DescribeInstanceMembersResponseBodyMembers) SetMemberStatus(v string) *DescribeInstanceMembersResponseBodyMembers {
	s.MemberStatus = &v
	return s
}

func (s *DescribeInstanceMembersResponseBodyMembers) SetMemberUid(v int64) *DescribeInstanceMembersResponseBodyMembers {
	s.MemberUid = &v
	return s
}

func (s *DescribeInstanceMembersResponseBodyMembers) SetModifyTime(v int32) *DescribeInstanceMembersResponseBodyMembers {
	s.ModifyTime = &v
	return s
}

type DescribeInstanceMembersResponseBodyPageInfo struct {
	// The page number.
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of the members.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInstanceMembersResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceMembersResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMembersResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeInstanceMembersResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInstanceMembersResponseBodyPageInfo) SetPageSize(v int32) *DescribeInstanceMembersResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceMembersResponseBodyPageInfo) SetTotalCount(v int32) *DescribeInstanceMembersResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

type DescribeInstanceMembersResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceMembersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceMembersResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMembersResponse) SetHeaders(v map[string]*string) *DescribeInstanceMembersResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceMembersResponse) SetStatusCode(v int32) *DescribeInstanceMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceMembersResponse) SetBody(v *DescribeInstanceMembersResponseBody) *DescribeInstanceMembersResponse {
	s.Body = v
	return s
}

type DescribeInstanceRiskLevelsRequest struct {
	// The information about the instances.
	Instances []*DescribeInstanceRiskLevelsRequestInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The language of the content within the response. Valid values:
	//
	// *   **zh** (default): Chinese
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeInstanceRiskLevelsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceRiskLevelsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRiskLevelsRequest) SetInstances(v []*DescribeInstanceRiskLevelsRequestInstances) *DescribeInstanceRiskLevelsRequest {
	s.Instances = v
	return s
}

func (s *DescribeInstanceRiskLevelsRequest) SetLang(v string) *DescribeInstanceRiskLevelsRequest {
	s.Lang = &v
	return s
}

type DescribeInstanceRiskLevelsRequestInstances struct {
	// The instance ID of your Cloud Firewall.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The public IP addresses of instances.
	InternetIp []*string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty" type:"Repeated"`
	// The private IP address of the instance.
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The UUID of the instance.
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeInstanceRiskLevelsRequestInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceRiskLevelsRequestInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRiskLevelsRequestInstances) SetInstanceId(v string) *DescribeInstanceRiskLevelsRequestInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceRiskLevelsRequestInstances) SetInternetIp(v []*string) *DescribeInstanceRiskLevelsRequestInstances {
	s.InternetIp = v
	return s
}

func (s *DescribeInstanceRiskLevelsRequestInstances) SetIntranetIp(v string) *DescribeInstanceRiskLevelsRequestInstances {
	s.IntranetIp = &v
	return s
}

func (s *DescribeInstanceRiskLevelsRequestInstances) SetUuid(v string) *DescribeInstanceRiskLevelsRequestInstances {
	s.Uuid = &v
	return s
}

type DescribeInstanceRiskLevelsResponseBody struct {
	// The information about the instances.
	InstanceRisks []*DescribeInstanceRiskLevelsResponseBodyInstanceRisks `json:"InstanceRisks,omitempty" xml:"InstanceRisks,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceRiskLevelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceRiskLevelsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRiskLevelsResponseBody) SetInstanceRisks(v []*DescribeInstanceRiskLevelsResponseBodyInstanceRisks) *DescribeInstanceRiskLevelsResponseBody {
	s.InstanceRisks = v
	return s
}

func (s *DescribeInstanceRiskLevelsResponseBody) SetRequestId(v string) *DescribeInstanceRiskLevelsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInstanceRiskLevelsResponseBodyInstanceRisks struct {
	// The risk levels of the Elastic Compute Service (ECS) instance.
	Details []*DescribeInstanceRiskLevelsResponseBodyInstanceRisksDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	// The instance ID of your Cloud Firewall.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The risk levels. Valid values:
	//
	// *   **medium**
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s DescribeInstanceRiskLevelsResponseBodyInstanceRisks) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceRiskLevelsResponseBodyInstanceRisks) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRiskLevelsResponseBodyInstanceRisks) SetDetails(v []*DescribeInstanceRiskLevelsResponseBodyInstanceRisksDetails) *DescribeInstanceRiskLevelsResponseBodyInstanceRisks {
	s.Details = v
	return s
}

func (s *DescribeInstanceRiskLevelsResponseBodyInstanceRisks) SetInstanceId(v string) *DescribeInstanceRiskLevelsResponseBodyInstanceRisks {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceRiskLevelsResponseBodyInstanceRisks) SetLevel(v string) *DescribeInstanceRiskLevelsResponseBodyInstanceRisks {
	s.Level = &v
	return s
}

type DescribeInstanceRiskLevelsResponseBodyInstanceRisksDetails struct {
	// The IP addresses of servers.
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The risk levels. Valid values:
	//
	// *   **medium**
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The type.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeInstanceRiskLevelsResponseBodyInstanceRisksDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceRiskLevelsResponseBodyInstanceRisksDetails) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRiskLevelsResponseBodyInstanceRisksDetails) SetIp(v string) *DescribeInstanceRiskLevelsResponseBodyInstanceRisksDetails {
	s.Ip = &v
	return s
}

func (s *DescribeInstanceRiskLevelsResponseBodyInstanceRisksDetails) SetLevel(v string) *DescribeInstanceRiskLevelsResponseBodyInstanceRisksDetails {
	s.Level = &v
	return s
}

func (s *DescribeInstanceRiskLevelsResponseBodyInstanceRisksDetails) SetType(v string) *DescribeInstanceRiskLevelsResponseBodyInstanceRisksDetails {
	s.Type = &v
	return s
}

type DescribeInstanceRiskLevelsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceRiskLevelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceRiskLevelsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceRiskLevelsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRiskLevelsResponse) SetHeaders(v map[string]*string) *DescribeInstanceRiskLevelsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceRiskLevelsResponse) SetStatusCode(v int32) *DescribeInstanceRiskLevelsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceRiskLevelsResponse) SetBody(v *DescribeInstanceRiskLevelsResponseBody) *DescribeInstanceRiskLevelsResponse {
	s.Body = v
	return s
}

type DescribeInternetOpenIpRequest struct {
	// The instance ID.
	AssetsInstanceId *string `json:"AssetsInstanceId,omitempty" xml:"AssetsInstanceId,omitempty"`
	// The instance name.
	AssetsInstanceName *string `json:"AssetsInstanceName,omitempty" xml:"AssetsInstanceName,omitempty"`
	// The asset type of the instance.
	AssetsType *string `json:"AssetsType,omitempty" xml:"AssetsType,omitempty"`
	// The page number.
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries per page.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The port number.
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The public IP address of the instance.
	PublicIp *string `json:"PublicIp,omitempty" xml:"PublicIp,omitempty"`
	// The region ID of the instance.
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The risk level. If you leave this parameter empty, all risk levels are queried. Valid values:
	//
	// *   **3**: high risk
	// *   **2**: medium risk
	// *   **1**: low risk
	// *   **0**: no risk
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The application.
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeInternetOpenIpRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInternetOpenIpRequest) GoString() string {
	return s.String()
}

func (s *DescribeInternetOpenIpRequest) SetAssetsInstanceId(v string) *DescribeInternetOpenIpRequest {
	s.AssetsInstanceId = &v
	return s
}

func (s *DescribeInternetOpenIpRequest) SetAssetsInstanceName(v string) *DescribeInternetOpenIpRequest {
	s.AssetsInstanceName = &v
	return s
}

func (s *DescribeInternetOpenIpRequest) SetAssetsType(v string) *DescribeInternetOpenIpRequest {
	s.AssetsType = &v
	return s
}

func (s *DescribeInternetOpenIpRequest) SetCurrentPage(v string) *DescribeInternetOpenIpRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInternetOpenIpRequest) SetEndTime(v string) *DescribeInternetOpenIpRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeInternetOpenIpRequest) SetLang(v string) *DescribeInternetOpenIpRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInternetOpenIpRequest) SetPageSize(v string) *DescribeInternetOpenIpRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInternetOpenIpRequest) SetPort(v string) *DescribeInternetOpenIpRequest {
	s.Port = &v
	return s
}

func (s *DescribeInternetOpenIpRequest) SetPublicIp(v string) *DescribeInternetOpenIpRequest {
	s.PublicIp = &v
	return s
}

func (s *DescribeInternetOpenIpRequest) SetRegionNo(v string) *DescribeInternetOpenIpRequest {
	s.RegionNo = &v
	return s
}

func (s *DescribeInternetOpenIpRequest) SetRiskLevel(v string) *DescribeInternetOpenIpRequest {
	s.RiskLevel = &v
	return s
}

func (s *DescribeInternetOpenIpRequest) SetServiceName(v string) *DescribeInternetOpenIpRequest {
	s.ServiceName = &v
	return s
}

func (s *DescribeInternetOpenIpRequest) SetStartTime(v string) *DescribeInternetOpenIpRequest {
	s.StartTime = &v
	return s
}

type DescribeInternetOpenIpResponseBody struct {
	// The data returned.
	DataList []*DescribeInternetOpenIpResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeInternetOpenIpResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInternetOpenIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInternetOpenIpResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInternetOpenIpResponseBody) SetDataList(v []*DescribeInternetOpenIpResponseBodyDataList) *DescribeInternetOpenIpResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeInternetOpenIpResponseBody) SetPageInfo(v *DescribeInternetOpenIpResponseBodyPageInfo) *DescribeInternetOpenIpResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeInternetOpenIpResponseBody) SetRequestId(v string) *DescribeInternetOpenIpResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInternetOpenIpResponseBodyDataList struct {
	// The reason why recommended intelligent policies are unavailable. Valid values:
	//
	// *   No recommended intelligent policies are available.
	// *   This feature is available only to some users.
	// *   The policy configuration has been modified. No recommended intelligent policies are available.
	// *   The recommended intelligent policies have been configured. No new recommended intelligent policies are available.
	AclRecommendDetail *string `json:"AclRecommendDetail,omitempty" xml:"AclRecommendDetail,omitempty"`
	// The instance ID.
	AssetsInstanceId *string `json:"AssetsInstanceId,omitempty" xml:"AssetsInstanceId,omitempty"`
	// The instance name.
	AssetsName *string `json:"AssetsName,omitempty" xml:"AssetsName,omitempty"`
	// The asset type of the instance.
	AssetsType *string `json:"AssetsType,omitempty" xml:"AssetsType,omitempty"`
	// The total number of ports.
	DetailNum *int32 `json:"DetailNum,omitempty" xml:"DetailNum,omitempty"`
	// Specifies whether an access control policy is recommended. Valid values:
	//
	// *   **true**
	// *   **false**
	HasAclRecommend *bool `json:"HasAclRecommend,omitempty" xml:"HasAclRecommend,omitempty"`
	// The list of ports.
	PortList []*string `json:"PortList,omitempty" xml:"PortList,omitempty" type:"Repeated"`
	// The public IP address of the instance.
	PublicIp *string `json:"PublicIp,omitempty" xml:"PublicIp,omitempty"`
	// The region ID of the instance.
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The risk level. Valid values:
	//
	// *   **3**: high risk
	// *   **2**: medium risk
	// *   **1**: low risk
	// *   **0**: no risk
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The reason for the risk.
	RiskReason *string `json:"RiskReason,omitempty" xml:"RiskReason,omitempty"`
	// The list of applications.
	ServiceNameList []*string `json:"ServiceNameList,omitempty" xml:"ServiceNameList,omitempty" type:"Repeated"`
	// The percentage of traffic of a day. Unit: percent (%).
	TrafficPercent1Day *string `json:"TrafficPercent1Day,omitempty" xml:"TrafficPercent1Day,omitempty"`
	// The percentage of traffic of 30 days. Unit: percent (%).
	TrafficPercent30Day *string `json:"TrafficPercent30Day,omitempty" xml:"TrafficPercent30Day,omitempty"`
	// The percentage of traffic of seven days. Unit: percent (%).
	TrafficPercent7Day *string `json:"TrafficPercent7Day,omitempty" xml:"TrafficPercent7Day,omitempty"`
}

func (s DescribeInternetOpenIpResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s DescribeInternetOpenIpResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeInternetOpenIpResponseBodyDataList) SetAclRecommendDetail(v string) *DescribeInternetOpenIpResponseBodyDataList {
	s.AclRecommendDetail = &v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyDataList) SetAssetsInstanceId(v string) *DescribeInternetOpenIpResponseBodyDataList {
	s.AssetsInstanceId = &v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyDataList) SetAssetsName(v string) *DescribeInternetOpenIpResponseBodyDataList {
	s.AssetsName = &v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyDataList) SetAssetsType(v string) *DescribeInternetOpenIpResponseBodyDataList {
	s.AssetsType = &v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyDataList) SetDetailNum(v int32) *DescribeInternetOpenIpResponseBodyDataList {
	s.DetailNum = &v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyDataList) SetHasAclRecommend(v bool) *DescribeInternetOpenIpResponseBodyDataList {
	s.HasAclRecommend = &v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyDataList) SetPortList(v []*string) *DescribeInternetOpenIpResponseBodyDataList {
	s.PortList = v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyDataList) SetPublicIp(v string) *DescribeInternetOpenIpResponseBodyDataList {
	s.PublicIp = &v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyDataList) SetRegionNo(v string) *DescribeInternetOpenIpResponseBodyDataList {
	s.RegionNo = &v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyDataList) SetRiskLevel(v int32) *DescribeInternetOpenIpResponseBodyDataList {
	s.RiskLevel = &v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyDataList) SetRiskReason(v string) *DescribeInternetOpenIpResponseBodyDataList {
	s.RiskReason = &v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyDataList) SetServiceNameList(v []*string) *DescribeInternetOpenIpResponseBodyDataList {
	s.ServiceNameList = v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyDataList) SetTrafficPercent1Day(v string) *DescribeInternetOpenIpResponseBodyDataList {
	s.TrafficPercent1Day = &v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyDataList) SetTrafficPercent30Day(v string) *DescribeInternetOpenIpResponseBodyDataList {
	s.TrafficPercent30Day = &v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyDataList) SetTrafficPercent7Day(v string) *DescribeInternetOpenIpResponseBodyDataList {
	s.TrafficPercent7Day = &v
	return s
}

type DescribeInternetOpenIpResponseBodyPageInfo struct {
	// The page number.
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInternetOpenIpResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeInternetOpenIpResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeInternetOpenIpResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeInternetOpenIpResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyPageInfo) SetPageSize(v int32) *DescribeInternetOpenIpResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeInternetOpenIpResponseBodyPageInfo) SetTotalCount(v int32) *DescribeInternetOpenIpResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

type DescribeInternetOpenIpResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInternetOpenIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInternetOpenIpResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInternetOpenIpResponse) GoString() string {
	return s.String()
}

func (s *DescribeInternetOpenIpResponse) SetHeaders(v map[string]*string) *DescribeInternetOpenIpResponse {
	s.Headers = v
	return s
}

func (s *DescribeInternetOpenIpResponse) SetStatusCode(v int32) *DescribeInternetOpenIpResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInternetOpenIpResponse) SetBody(v *DescribeInternetOpenIpResponseBody) *DescribeInternetOpenIpResponse {
	s.Body = v
	return s
}

type DescribeInternetTrafficTrendRequest struct {
	// The direction of the internet traffic.
	//
	// Valid values:
	//
	// *   **in**: inbound traffic
	// *   **out**: outbound traffic
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the content in the request and response. Valid values:
	//
	// *   **zh** (default): Chinese
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The source code.
	SourceCode *string `json:"SourceCode,omitempty" xml:"SourceCode,omitempty"`
	// Deprecated
	// The IP address of the access source.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The private IP address of the source.
	SrcPrivateIP *string `json:"SrcPrivateIP,omitempty" xml:"SrcPrivateIP,omitempty"`
	// The public IP address of the source.
	SrcPublicIP *string `json:"SrcPublicIP,omitempty" xml:"SrcPublicIP,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The type of the traffic that is captured. Valid values:
	//
	// *   **max** (default): peak traffic
	// *   **avg**: average traffic
	TrafficType *string `json:"TrafficType,omitempty" xml:"TrafficType,omitempty"`
}

func (s DescribeInternetTrafficTrendRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInternetTrafficTrendRequest) GoString() string {
	return s.String()
}

func (s *DescribeInternetTrafficTrendRequest) SetDirection(v string) *DescribeInternetTrafficTrendRequest {
	s.Direction = &v
	return s
}

func (s *DescribeInternetTrafficTrendRequest) SetEndTime(v string) *DescribeInternetTrafficTrendRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeInternetTrafficTrendRequest) SetLang(v string) *DescribeInternetTrafficTrendRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInternetTrafficTrendRequest) SetSourceCode(v string) *DescribeInternetTrafficTrendRequest {
	s.SourceCode = &v
	return s
}

func (s *DescribeInternetTrafficTrendRequest) SetSourceIp(v string) *DescribeInternetTrafficTrendRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInternetTrafficTrendRequest) SetSrcPrivateIP(v string) *DescribeInternetTrafficTrendRequest {
	s.SrcPrivateIP = &v
	return s
}

func (s *DescribeInternetTrafficTrendRequest) SetSrcPublicIP(v string) *DescribeInternetTrafficTrendRequest {
	s.SrcPublicIP = &v
	return s
}

func (s *DescribeInternetTrafficTrendRequest) SetStartTime(v string) *DescribeInternetTrafficTrendRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeInternetTrafficTrendRequest) SetTrafficType(v string) *DescribeInternetTrafficTrendRequest {
	s.TrafficType = &v
	return s
}

type DescribeInternetTrafficTrendResponseBody struct {
	// The average inbound network throughput, which indicates the average number of bits that are sent inbound per second. Unit: bit/s.
	AvgInBps *int64 `json:"AvgInBps,omitempty" xml:"AvgInBps,omitempty"`
	// The average outbound network throughput, which indicates the average number of bits that are sent outbound per second. Unit: bit/s.
	AvgOutBps *int64 `json:"AvgOutBps,omitempty" xml:"AvgOutBps,omitempty"`
	// The average number of requests.
	AvgSession *int64 `json:"AvgSession,omitempty" xml:"AvgSession,omitempty"`
	// The total average inbound and outbound network throughput, which indicates the average number of bits that are sent inbound and outbound per second. Unit: bit/s.
	AvgTotalBps *int64 `json:"AvgTotalBps,omitempty" xml:"AvgTotalBps,omitempty"`
	// The statistics on traffic.
	DataList []*DescribeInternetTrafficTrendResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// The timestamp generated when the bandwidth reaches the peak value. The value is a UNIX timestamp. Unit: seconds.
	MaxBandwidthTime *int64 `json:"MaxBandwidthTime,omitempty" xml:"MaxBandwidthTime,omitempty"`
	// The maximum inbound network throughput, which indicates the maximum number of bits that are sent inbound per second. Unit: bit/s.
	MaxInBps *int64 `json:"MaxInBps,omitempty" xml:"MaxInBps,omitempty"`
	// The maximum outbound network throughput, which indicates the maximum number of bits that are sent outbound per second. Unit: bit/s.
	MaxOutBps *int64 `json:"MaxOutBps,omitempty" xml:"MaxOutBps,omitempty"`
	// The number of requests during the peak hour of the network throughout.
	MaxSession *int64 `json:"MaxSession,omitempty" xml:"MaxSession,omitempty"`
	// The total maximum inbound and outbound network throughput, which indicates the maximum number of bits that are sent inbound and outbound per second. Unit: bit/s.
	MaxTotalBps *int64 `json:"MaxTotalBps,omitempty" xml:"MaxTotalBps,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total inbound and outbound network throughput, which indicates the total number of bytes that are sent inbound and outbound. Unit: bytes.
	TotalBytes *int64 `json:"TotalBytes,omitempty" xml:"TotalBytes,omitempty"`
	// The inbound network throughput, which indicates the total number of bytes that are sent inbound. Unit: bytes.
	TotalInBytes *int64 `json:"TotalInBytes,omitempty" xml:"TotalInBytes,omitempty"`
	// The outbound network throughput, which indicates the total number of bytes that are sent outbound. Unit: bytes.
	TotalOutBytes *int64 `json:"TotalOutBytes,omitempty" xml:"TotalOutBytes,omitempty"`
	// The total number of requests.
	TotalSession *int64 `json:"TotalSession,omitempty" xml:"TotalSession,omitempty"`
}

func (s DescribeInternetTrafficTrendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInternetTrafficTrendResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInternetTrafficTrendResponseBody) SetAvgInBps(v int64) *DescribeInternetTrafficTrendResponseBody {
	s.AvgInBps = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBody) SetAvgOutBps(v int64) *DescribeInternetTrafficTrendResponseBody {
	s.AvgOutBps = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBody) SetAvgSession(v int64) *DescribeInternetTrafficTrendResponseBody {
	s.AvgSession = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBody) SetAvgTotalBps(v int64) *DescribeInternetTrafficTrendResponseBody {
	s.AvgTotalBps = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBody) SetDataList(v []*DescribeInternetTrafficTrendResponseBodyDataList) *DescribeInternetTrafficTrendResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBody) SetMaxBandwidthTime(v int64) *DescribeInternetTrafficTrendResponseBody {
	s.MaxBandwidthTime = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBody) SetMaxInBps(v int64) *DescribeInternetTrafficTrendResponseBody {
	s.MaxInBps = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBody) SetMaxOutBps(v int64) *DescribeInternetTrafficTrendResponseBody {
	s.MaxOutBps = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBody) SetMaxSession(v int64) *DescribeInternetTrafficTrendResponseBody {
	s.MaxSession = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBody) SetMaxTotalBps(v int64) *DescribeInternetTrafficTrendResponseBody {
	s.MaxTotalBps = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBody) SetRequestId(v string) *DescribeInternetTrafficTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBody) SetTotalBytes(v int64) *DescribeInternetTrafficTrendResponseBody {
	s.TotalBytes = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBody) SetTotalInBytes(v int64) *DescribeInternetTrafficTrendResponseBody {
	s.TotalInBytes = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBody) SetTotalOutBytes(v int64) *DescribeInternetTrafficTrendResponseBody {
	s.TotalOutBytes = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBody) SetTotalSession(v int64) *DescribeInternetTrafficTrendResponseBody {
	s.TotalSession = &v
	return s
}

type DescribeInternetTrafficTrendResponseBodyDataList struct {
	// The inbound network throughput, which indicates the number of bits that are sent inbound per second. Unit: bit/s.
	InBps *int64 `json:"InBps,omitempty" xml:"InBps,omitempty"`
	// The inbound network throughput, which indicates the total number of bytes that are sent inbound. Unit: bytes.
	InBytes *int64 `json:"InBytes,omitempty" xml:"InBytes,omitempty"`
	// The inbound network throughput, which indicates the number of packets that are sent inbound per second. Unit: packets per second (pps).
	InPps *int64 `json:"InPps,omitempty" xml:"InPps,omitempty"`
	// The number of new connections.
	NewConn *int64 `json:"NewConn,omitempty" xml:"NewConn,omitempty"`
	// The outbound network throughput, which indicates the number of bits that are sent inbound per second. Unit: bit/s.
	OutBps *int64 `json:"OutBps,omitempty" xml:"OutBps,omitempty"`
	// The outbound network throughput, which indicates the total number of bytes that are sent outbound. Unit: bytes.
	OutBytes *int64 `json:"OutBytes,omitempty" xml:"OutBytes,omitempty"`
	// The outbound network throughput, which indicates the number of packets that are sent outbound per second. Unit: pps.
	OutPps *int64 `json:"OutPps,omitempty" xml:"OutPps,omitempty"`
	// The number of requests.
	SessionCount *int64 `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
	// The time when traffic is generated. The value is a UNIX timestamp. Unit: seconds.
	Time *int32 `json:"Time,omitempty" xml:"Time,omitempty"`
	// The total outbound and inbound network throughput, which indicates the total number of bits that are sent inbound and outbound per second. Unit: bit/s.
	TotalBps *int64 `json:"TotalBps,omitempty" xml:"TotalBps,omitempty"`
}

func (s DescribeInternetTrafficTrendResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s DescribeInternetTrafficTrendResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeInternetTrafficTrendResponseBodyDataList) SetInBps(v int64) *DescribeInternetTrafficTrendResponseBodyDataList {
	s.InBps = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBodyDataList) SetInBytes(v int64) *DescribeInternetTrafficTrendResponseBodyDataList {
	s.InBytes = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBodyDataList) SetInPps(v int64) *DescribeInternetTrafficTrendResponseBodyDataList {
	s.InPps = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBodyDataList) SetNewConn(v int64) *DescribeInternetTrafficTrendResponseBodyDataList {
	s.NewConn = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBodyDataList) SetOutBps(v int64) *DescribeInternetTrafficTrendResponseBodyDataList {
	s.OutBps = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBodyDataList) SetOutBytes(v int64) *DescribeInternetTrafficTrendResponseBodyDataList {
	s.OutBytes = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBodyDataList) SetOutPps(v int64) *DescribeInternetTrafficTrendResponseBodyDataList {
	s.OutPps = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBodyDataList) SetSessionCount(v int64) *DescribeInternetTrafficTrendResponseBodyDataList {
	s.SessionCount = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBodyDataList) SetTime(v int32) *DescribeInternetTrafficTrendResponseBodyDataList {
	s.Time = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponseBodyDataList) SetTotalBps(v int64) *DescribeInternetTrafficTrendResponseBodyDataList {
	s.TotalBps = &v
	return s
}

type DescribeInternetTrafficTrendResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInternetTrafficTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInternetTrafficTrendResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInternetTrafficTrendResponse) GoString() string {
	return s.String()
}

func (s *DescribeInternetTrafficTrendResponse) SetHeaders(v map[string]*string) *DescribeInternetTrafficTrendResponse {
	s.Headers = v
	return s
}

func (s *DescribeInternetTrafficTrendResponse) SetStatusCode(v int32) *DescribeInternetTrafficTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInternetTrafficTrendResponse) SetBody(v *DescribeInternetTrafficTrendResponseBody) *DescribeInternetTrafficTrendResponse {
	s.Body = v
	return s
}

type DescribeInvadeEventListRequest struct {
	// The IP address of the affected asset.
	AssetsIP *string `json:"AssetsIP,omitempty" xml:"AssetsIP,omitempty"`
	// The ID of the instance.
	AssetsInstanceId *string `json:"AssetsInstanceId,omitempty" xml:"AssetsInstanceId,omitempty"`
	// The name of the instance.
	AssetsInstanceName *string `json:"AssetsInstanceName,omitempty" xml:"AssetsInstanceName,omitempty"`
	// The number of the page to return.
	//
	// Default value: 1.
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds. If you do not specify this parameter, the query ends at the current time.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the breach awareness event.
	EventKey *string `json:"EventKey,omitempty" xml:"EventKey,omitempty"`
	// The name of the breach awareness event.
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The UUID of the breach awareness event.
	EventUuid *string `json:"EventUuid,omitempty" xml:"EventUuid,omitempty"`
	// Specifies whether the breach awareness event is ignored. Valid values:
	//
	// *   **true**: The breach awareness event is ignored.
	// *   **false**: The breach awareness event is not ignored.
	IsIgnore *string `json:"IsIgnore,omitempty" xml:"IsIgnore,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the member.
	MemberUid *int64 `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The number of entries to return on each page.
	//
	// Default value: 6. Maximum value: 10.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The handling status of breach awareness events.
	ProcessStatusList []*int32 `json:"ProcessStatusList,omitempty" xml:"ProcessStatusList,omitempty" type:"Repeated"`
	// The risk levels.
	RiskLevel []*int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty" type:"Repeated"`
	// Deprecated
	// The source IP address of the request.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds. If you do not specify this parameter, the query starts from 30 days before the current time.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeInvadeEventListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvadeEventListRequest) GoString() string {
	return s.String()
}

func (s *DescribeInvadeEventListRequest) SetAssetsIP(v string) *DescribeInvadeEventListRequest {
	s.AssetsIP = &v
	return s
}

func (s *DescribeInvadeEventListRequest) SetAssetsInstanceId(v string) *DescribeInvadeEventListRequest {
	s.AssetsInstanceId = &v
	return s
}

func (s *DescribeInvadeEventListRequest) SetAssetsInstanceName(v string) *DescribeInvadeEventListRequest {
	s.AssetsInstanceName = &v
	return s
}

func (s *DescribeInvadeEventListRequest) SetCurrentPage(v string) *DescribeInvadeEventListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInvadeEventListRequest) SetEndTime(v string) *DescribeInvadeEventListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeInvadeEventListRequest) SetEventKey(v string) *DescribeInvadeEventListRequest {
	s.EventKey = &v
	return s
}

func (s *DescribeInvadeEventListRequest) SetEventName(v string) *DescribeInvadeEventListRequest {
	s.EventName = &v
	return s
}

func (s *DescribeInvadeEventListRequest) SetEventUuid(v string) *DescribeInvadeEventListRequest {
	s.EventUuid = &v
	return s
}

func (s *DescribeInvadeEventListRequest) SetIsIgnore(v string) *DescribeInvadeEventListRequest {
	s.IsIgnore = &v
	return s
}

func (s *DescribeInvadeEventListRequest) SetLang(v string) *DescribeInvadeEventListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeInvadeEventListRequest) SetMemberUid(v int64) *DescribeInvadeEventListRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeInvadeEventListRequest) SetPageSize(v string) *DescribeInvadeEventListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInvadeEventListRequest) SetProcessStatusList(v []*int32) *DescribeInvadeEventListRequest {
	s.ProcessStatusList = v
	return s
}

func (s *DescribeInvadeEventListRequest) SetRiskLevel(v []*int32) *DescribeInvadeEventListRequest {
	s.RiskLevel = v
	return s
}

func (s *DescribeInvadeEventListRequest) SetSourceIp(v string) *DescribeInvadeEventListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInvadeEventListRequest) SetStartTime(v string) *DescribeInvadeEventListRequest {
	s.StartTime = &v
	return s
}

type DescribeInvadeEventListResponseBody struct {
	// An array that consists of breach awareness events.
	EventList []*DescribeInvadeEventListResponseBodyEventList `json:"EventList,omitempty" xml:"EventList,omitempty" type:"Repeated"`
	// The percentage of high-risk events.
	HighLevelPercent *int32 `json:"HighLevelPercent,omitempty" xml:"HighLevelPercent,omitempty"`
	// The percentage of low-risk events.
	LowLevelPercent *int32 `json:"LowLevelPercent,omitempty" xml:"LowLevelPercent,omitempty"`
	// The percentage of medium-risk events.
	MiddleLevelPercent *int32 `json:"MiddleLevelPercent,omitempty" xml:"MiddleLevelPercent,omitempty"`
	// The pagination information.
	PageInfo *DescribeInvadeEventListResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInvadeEventListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvadeEventListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInvadeEventListResponseBody) SetEventList(v []*DescribeInvadeEventListResponseBodyEventList) *DescribeInvadeEventListResponseBody {
	s.EventList = v
	return s
}

func (s *DescribeInvadeEventListResponseBody) SetHighLevelPercent(v int32) *DescribeInvadeEventListResponseBody {
	s.HighLevelPercent = &v
	return s
}

func (s *DescribeInvadeEventListResponseBody) SetLowLevelPercent(v int32) *DescribeInvadeEventListResponseBody {
	s.LowLevelPercent = &v
	return s
}

func (s *DescribeInvadeEventListResponseBody) SetMiddleLevelPercent(v int32) *DescribeInvadeEventListResponseBody {
	s.MiddleLevelPercent = &v
	return s
}

func (s *DescribeInvadeEventListResponseBody) SetPageInfo(v *DescribeInvadeEventListResponseBodyPageInfo) *DescribeInvadeEventListResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeInvadeEventListResponseBody) SetRequestId(v string) *DescribeInvadeEventListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInvadeEventListResponseBodyEventList struct {
	// The ID of the affected asset.
	AssetsInstanceId *string `json:"AssetsInstanceId,omitempty" xml:"AssetsInstanceId,omitempty"`
	// The name of the affected asset.
	AssetsInstanceName *string `json:"AssetsInstanceName,omitempty" xml:"AssetsInstanceName,omitempty"`
	// The type of the affected asset. Valid values:
	//
	// *   **BastionHostIP**: the egress IP address of a bastion host
	// *   **BastionHostIngressIP**: the ingress IP address of a bastion host
	// *   **EcsEIP**: the elastic IP address (EIP) of an Elastic Compute Service (ECS) instance
	// *   **EcsPublicIP**: the public IP address of an ECS instance
	// *   **EIP**: the EIP
	// *   **EniEIP**: the EIP of an elastic network interface (ENI)
	// *   **NatEIP**: the EIP of a NAT gateway
	// *   **SlbEIP**: the EIP of a Server Load Balancer (SLB) instance
	// *   **SlbPublicIP**: the public IP address of an SLB instance
	// *   **NatPublicIP**: the public IP address of a NAT gateway
	// *   **HAVIP**: the high-availability virtual IP address (HAVIP)
	AssetsType *string `json:"AssetsType,omitempty" xml:"AssetsType,omitempty"`
	// The ID of the breach awareness event.
	EventKey *string `json:"EventKey,omitempty" xml:"EventKey,omitempty"`
	// The name of the breach awareness event.
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The type of the breach awareness event. Valid values:
	//
	// *   **IPS**: intrusion prevention event
	// *   **offline**: disconnection event
	EventSrc *string `json:"EventSrc,omitempty" xml:"EventSrc,omitempty"`
	// The UUID of the breach awareness event.
	EventUuid *string `json:"EventUuid,omitempty" xml:"EventUuid,omitempty"`
	// The time when the breach awareness event first occurred. The value is a UNIX timestamp. Unit: seconds.
	FirstTime *int32 `json:"FirstTime,omitempty" xml:"FirstTime,omitempty"`
	// Indicates whether the breach awareness event is ignored. Valid values:
	//
	// *   **true**: The breach awareness event is ignored.
	// *   **false**: The breach awareness event is not ignored.
	IsIgnore *bool `json:"IsIgnore,omitempty" xml:"IsIgnore,omitempty"`
	// The time when the breach awareness event last occurred. The value is a UNIX timestamp. Unit: seconds.
	LastTime *int32 `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	// The ID of the member.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The private IP address of the affected asset.
	PrivateIP *string `json:"PrivateIP,omitempty" xml:"PrivateIP,omitempty"`
	// The handling status of the breach awareness event. Valid values:
	//
	// *   **0**: unhandled
	// *   **20**: handled
	ProcessStatus *int32 `json:"ProcessStatus,omitempty" xml:"ProcessStatus,omitempty"`
	// The public IP address of the affected asset.
	PublicIP *string `json:"PublicIP,omitempty" xml:"PublicIP,omitempty"`
	// The type of the affected asset. Valid values:
	//
	// *   **BastionHostIP**: the egress IP address of a bastion host
	// *   **BastionHostIngressIP**: the ingress IP address of a bastion host
	// *   **EcsEIP**: the EIP of an ECS instance
	// *   **EcsPublicIP**: the public IP address of an ECS instance
	// *   **EIP**: the EIP
	// *   **EniEIP**: the EIP of an ENI
	// *   **NatEIP**: the EIP of a NAT gateway
	// *   **SlbEIP**: the EIP of an SLB instance
	// *   **SlbPublicIP**: the public IP address of an SLB instance
	// *   **NatPublicIP**: the public IP address of a NAT gateway
	// *   **HAVIP**: the HAVIP
	PublicIpType *string `json:"PublicIpType,omitempty" xml:"PublicIpType,omitempty"`
	// The risk level. Valid values:
	//
	// *   **1**: low
	// *   **2**: medium
	// *   **3**: high
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s DescribeInvadeEventListResponseBodyEventList) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvadeEventListResponseBodyEventList) GoString() string {
	return s.String()
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetAssetsInstanceId(v string) *DescribeInvadeEventListResponseBodyEventList {
	s.AssetsInstanceId = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetAssetsInstanceName(v string) *DescribeInvadeEventListResponseBodyEventList {
	s.AssetsInstanceName = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetAssetsType(v string) *DescribeInvadeEventListResponseBodyEventList {
	s.AssetsType = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetEventKey(v string) *DescribeInvadeEventListResponseBodyEventList {
	s.EventKey = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetEventName(v string) *DescribeInvadeEventListResponseBodyEventList {
	s.EventName = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetEventSrc(v string) *DescribeInvadeEventListResponseBodyEventList {
	s.EventSrc = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetEventUuid(v string) *DescribeInvadeEventListResponseBodyEventList {
	s.EventUuid = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetFirstTime(v int32) *DescribeInvadeEventListResponseBodyEventList {
	s.FirstTime = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetIsIgnore(v bool) *DescribeInvadeEventListResponseBodyEventList {
	s.IsIgnore = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetLastTime(v int32) *DescribeInvadeEventListResponseBodyEventList {
	s.LastTime = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetMemberUid(v string) *DescribeInvadeEventListResponseBodyEventList {
	s.MemberUid = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetPrivateIP(v string) *DescribeInvadeEventListResponseBodyEventList {
	s.PrivateIP = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetProcessStatus(v int32) *DescribeInvadeEventListResponseBodyEventList {
	s.ProcessStatus = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetPublicIP(v string) *DescribeInvadeEventListResponseBodyEventList {
	s.PublicIP = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetPublicIpType(v string) *DescribeInvadeEventListResponseBodyEventList {
	s.PublicIpType = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyEventList) SetRiskLevel(v int32) *DescribeInvadeEventListResponseBodyEventList {
	s.RiskLevel = &v
	return s
}

type DescribeInvadeEventListResponseBodyPageInfo struct {
	// The page number of the returned page.
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of breach awareness events.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInvadeEventListResponseBodyPageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvadeEventListResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeInvadeEventListResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeInvadeEventListResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyPageInfo) SetPageSize(v int32) *DescribeInvadeEventListResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeInvadeEventListResponseBodyPageInfo) SetTotalCount(v int32) *DescribeInvadeEventListResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

type DescribeInvadeEventListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInvadeEventListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInvadeEventListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvadeEventListResponse) GoString() string {
	return s.String()
}

func (s *DescribeInvadeEventListResponse) SetHeaders(v map[string]*string) *DescribeInvadeEventListResponse {
	s.Headers = v
	return s
}

func (s *DescribeInvadeEventListResponse) SetStatusCode(v int32) *DescribeInvadeEventListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInvadeEventListResponse) SetBody(v *DescribeInvadeEventListResponseBody) *DescribeInvadeEventListResponse {
	s.Body = v
	return s
}

type DescribeNatAclPageStatusRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// *   **zh** (default): Chinese
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeNatAclPageStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNatAclPageStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeNatAclPageStatusRequest) SetLang(v string) *DescribeNatAclPageStatusRequest {
	s.Lang = &v
	return s
}

type DescribeNatAclPageStatusResponseBody struct {
	// Indicates whether pagination for access control policies for NAT firewalls is supported.
	NatAclPageEnable *bool `json:"NatAclPageEnable,omitempty" xml:"NatAclPageEnable,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNatAclPageStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNatAclPageStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNatAclPageStatusResponseBody) SetNatAclPageEnable(v bool) *DescribeNatAclPageStatusResponseBody {
	s.NatAclPageEnable = &v
	return s
}

func (s *DescribeNatAclPageStatusResponseBody) SetRequestId(v string) *DescribeNatAclPageStatusResponseBody {
	s.RequestId = &v
	return s
}

type DescribeNatAclPageStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNatAclPageStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNatAclPageStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNatAclPageStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeNatAclPageStatusResponse) SetHeaders(v map[string]*string) *DescribeNatAclPageStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeNatAclPageStatusResponse) SetStatusCode(v int32) *DescribeNatAclPageStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNatAclPageStatusResponse) SetBody(v *DescribeNatAclPageStatusResponseBody) *DescribeNatAclPageStatusResponse {
	s.Body = v
	return s
}

type DescribeNatFirewallControlPolicyRequest struct {
	// The action that Cloud Firewall performs on the traffic.
	//
	// Valid values:
	//
	// *   **accept**: allows the traffic.
	// *   **drop**: denies the traffic.
	// *   **log**: monitors the traffic.
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The UUID of the access control policy.
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The page number.
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The description of the access control policy. Fuzzy match is supported.
	//
	// > If you do not specify this parameter, the descriptions of all policies are queried.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination address in the access control policy. Fuzzy match is supported. The value of this parameter varies based on the value of the DestinationType parameter.
	//
	// *   If DestinationType is set to `net`, the value of Destination must be a CIDR block. Example: 10.0.3.0/24.
	// *   If DestinationType is set to `domain`, the value of Destination must be a domain name. Example: aliyun.
	// *   If DestinationType is set to `group`, the value of Destination must be the name of an address book. Example: db_group.
	// *   If DestinationType is set to `location`, the value of Destination is a location. For more information about location codes, see [AddControlPolicy](~~474128~~). Example: \["BJ11", "ZB"].
	//
	// > If you do not specify this parameter, all types of destination addresses are queried.
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The direction of the traffic to which the access control policy applies. Valid values:
	//
	// *   **out**: outbound traffic
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the NAT gateway.
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The number of entries per page. Default value: 10.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the protocol in the access control policy. Valid values:
	//
	// *   **TCP**
	// *   **UDP**
	// *   **ICMP**
	// *   **ANY**: all types of protocols
	//
	// > If you do not specify this parameter, access control policies of all protocol types are queried.
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// Specifies whether the access control policy is enabled. By default, an access control policy is enabled after it is created. Valid values:
	//
	// *   **true**
	// *   **false**
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The recurrence type for the access control policy to take effect. Valid values:
	//
	// *   **Permanent** (default): The policy always takes effect.
	// *   **None**: The policy takes effect for only once.
	// *   **Daily**: The policy takes effect on a daily basis.
	// *   **Weekly**: The policy takes effect on a weekly basis.
	// *   **Monthly**: The policy takes effect on a monthly basis.
	RepeatType *string `json:"RepeatType,omitempty" xml:"RepeatType,omitempty"`
	// The source address in the access control policy. Fuzzy match is supported. The value of this parameter varies based on the value of the SourceType parameter.
	//
	// *   If SourceType is set to `net`, the value of Source must be a CIDR block. Example: 192.0.XX.XX/24.
	// *   If SourceType is set to `group`, the value of Source must be the name of an address book. Example: db_group. If the db_group address book does not contain addresses, all source addresses are queried.
	// *   If SourceType is set to `location`, the value of Source must be a location. Example: beijing.
	//
	// > If you do not specify this parameter, all types of source addresses are queried.
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s DescribeNatFirewallControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNatFirewallControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallControlPolicyRequest) SetAclAction(v string) *DescribeNatFirewallControlPolicyRequest {
	s.AclAction = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyRequest) SetAclUuid(v string) *DescribeNatFirewallControlPolicyRequest {
	s.AclUuid = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyRequest) SetCurrentPage(v string) *DescribeNatFirewallControlPolicyRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyRequest) SetDescription(v string) *DescribeNatFirewallControlPolicyRequest {
	s.Description = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyRequest) SetDestination(v string) *DescribeNatFirewallControlPolicyRequest {
	s.Destination = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyRequest) SetDirection(v string) *DescribeNatFirewallControlPolicyRequest {
	s.Direction = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyRequest) SetLang(v string) *DescribeNatFirewallControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyRequest) SetNatGatewayId(v string) *DescribeNatFirewallControlPolicyRequest {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyRequest) SetPageSize(v string) *DescribeNatFirewallControlPolicyRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyRequest) SetProto(v string) *DescribeNatFirewallControlPolicyRequest {
	s.Proto = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyRequest) SetRelease(v string) *DescribeNatFirewallControlPolicyRequest {
	s.Release = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyRequest) SetRepeatType(v string) *DescribeNatFirewallControlPolicyRequest {
	s.RepeatType = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyRequest) SetSource(v string) *DescribeNatFirewallControlPolicyRequest {
	s.Source = &v
	return s
}

type DescribeNatFirewallControlPolicyResponseBody struct {
	// The information about the access control policies.
	Policys []*DescribeNatFirewallControlPolicyResponseBodyPolicys `json:"Policys,omitempty" xml:"Policys,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeNatFirewallControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNatFirewallControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallControlPolicyResponseBody) SetPolicys(v []*DescribeNatFirewallControlPolicyResponseBodyPolicys) *DescribeNatFirewallControlPolicyResponseBody {
	s.Policys = v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBody) SetRequestId(v string) *DescribeNatFirewallControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBody) SetTotalCount(v string) *DescribeNatFirewallControlPolicyResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeNatFirewallControlPolicyResponseBodyPolicys struct {
	// The action that Cloud Firewall performs on the traffic. Valid values:
	//
	// *   **accept**: allows the traffic.
	// *   **drop**: denies the traffic.
	// *   **log**: monitors the traffic.
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The UUID of the access control policy.
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The application names.
	ApplicationNameList []*string `json:"ApplicationNameList,omitempty" xml:"ApplicationNameList,omitempty" type:"Repeated"`
	// The time when the access control policy was created.
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the access control policy.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination port in the access control policy.
	DestPort *string `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	// The name of the destination port address book in the access control policy.
	DestPortGroup *string `json:"DestPortGroup,omitempty" xml:"DestPortGroup,omitempty"`
	// The ports in the destination port address book.
	DestPortGroupPorts []*string `json:"DestPortGroupPorts,omitempty" xml:"DestPortGroupPorts,omitempty" type:"Repeated"`
	// The type of the destination port in the access control policy. Valid values:
	//
	// *   **port**: port
	// *   **group**: port address book
	DestPortType *string `json:"DestPortType,omitempty" xml:"DestPortType,omitempty"`
	// The destination address in the access control policy. The value of this parameter varies based on the value of DestinationType. Valid values:
	//
	// *   If the value of **DestinationType** is **net**, the value of this parameter is a CIDR block. Example: 192.0.XX.XX/24.
	// *   If the value of **DestinationType** is **domain**, the value of this parameter is a domain name. Example: aliyuncs.com.
	// *   If the value of **DestinationType** is **group**, the value of this parameter is the name of an address book. Example: db_group.
	// *   If the value of **DestinationType** is **location**, the value of this parameter is a location. For more information about location codes, see [AddControlPolicy](~~138867~~). Example: \["BJ11", "ZB"].
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The CIDR blocks in the destination address book.
	DestinationGroupCidrs []*string `json:"DestinationGroupCidrs,omitempty" xml:"DestinationGroupCidrs,omitempty" type:"Repeated"`
	// The type of the destination address book in the access control policy. Valid values:
	//
	// *   **ip**: an address book that includes one or more CIDR blocks
	// *   **domain**: an address book that includes one or more domain names
	DestinationGroupType *string `json:"DestinationGroupType,omitempty" xml:"DestinationGroupType,omitempty"`
	// The type of the destination address in the access control policy. Valid values:
	//
	// *   **net**: CIDR block
	// *   **group**: address book
	// *   **domain**: domain name
	// *   **location**: location
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The DNS resolution result.
	DnsResult *string `json:"DnsResult,omitempty" xml:"DnsResult,omitempty"`
	// The time when the Domain Name System (DNS) resolution was performed. The value is a UNIX timestamp. Unit: seconds.
	DnsResultTime *int64 `json:"DnsResultTime,omitempty" xml:"DnsResultTime,omitempty"`
	// The domain name resolution method of the access control policy. By default, an access control policy is enabled after the policy is created. Valid values:
	//
	// *   **0**: fully qualified domain name (FQDN)-based resolution
	// *   **1**: DNS-based dynamic resolution
	// *   **2**: FQDN and DNS-based dynamic resolution
	DomainResolveType *int32 `json:"DomainResolveType,omitempty" xml:"DomainResolveType,omitempty"`
	// The time when the access control policy stops taking effect. The value is a UNIX timestamp. Unit: seconds. The end time must be on the hour or on the half hour, and at least 30 minutes later than the start time.
	//
	// >  If RepeatType is set to Permanent, this parameter is left empty. If RepeatType is set to None, Daily, Weekly, or Monthly, this parameter must be specified.
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time when the access control policy was last hit. The value is a UNIX timestamp. Unit: seconds.
	HitLastTime *int64 `json:"HitLastTime,omitempty" xml:"HitLastTime,omitempty"`
	// The number of hits for the access control policy.
	HitTimes *int32 `json:"HitTimes,omitempty" xml:"HitTimes,omitempty"`
	// The time when the access control policy was modified.
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The ID of the NAT gateway.
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The priority of the access control policy.
	//
	// The priority value starts from 1. A smaller priority value indicates a higher priority.
	Order *int32 `json:"Order,omitempty" xml:"Order,omitempty"`
	// The protocol type in the access control policy. Valid values:
	//
	// *   **ANY**
	// *   **TCP**
	// *   **UDP**
	// *   **ICMP**
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// The status of the access control policy. By default, an access control policy is enabled after it is created. Valid values:
	//
	// *   **true**: enabled
	// *   **false**: disabled
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The days of a week or of a month on which the access control policy takes effect.
	//
	// *   If RepeatType is set to `Permanent`, `None`, or `Daily`, the value of this parameter is an empty array. Example: \[].
	// *   If RepeatType is set to Weekly, this parameter must be specified. Example: \[0, 6].
	//
	// >  If RepeatType is set to Weekly, the fields in the value of this parameter cannot be repeated.
	//
	// *   If RepeatType is set to `Monthly`, this parameter must be specified. Example: \[1, 31].
	//
	// >  If RepeatType is set to Monthly, the fields in the value of this parameter cannot be repeated.
	RepeatDays []*int64 `json:"RepeatDays,omitempty" xml:"RepeatDays,omitempty" type:"Repeated"`
	// The point in time when the recurrence ends. Example: 23:30. The end time must be on the hour or on the half hour, and at least 30 minutes later than the start time.
	//
	// >  If RepeatType is set to Permanent or None, this parameter is left empty. If RepeatType is set to Daily, Weekly, or Monthly, this parameter must be specified.
	RepeatEndTime *string `json:"RepeatEndTime,omitempty" xml:"RepeatEndTime,omitempty"`
	// The point in time when the recurrence starts. Example: 08:00. The start time must be on the hour or on the half hour, and at least 30 minutes earlier than the end time.
	//
	// >  If RepeatType is set to Permanent or None, this parameter is left empty. If RepeatType is set to Daily, Weekly, or Monthly, this parameter must be specified.
	RepeatStartTime *string `json:"RepeatStartTime,omitempty" xml:"RepeatStartTime,omitempty"`
	// The recurrence type for the access control policy to take effect. Valid values:
	//
	// *   **Permanent** (default): The policy always takes effect.
	// *   **None**: The policy takes effect for only once.
	// *   **Daily**: The policy takes effect on a daily basis.
	// *   **Weekly**: The policy takes effect on a weekly basis.
	// *   **Monthly**: The policy takes effect on a monthly basis.
	RepeatType *string `json:"RepeatType,omitempty" xml:"RepeatType,omitempty"`
	// The source address in the access control policy. Valid values:
	//
	// *   If the value of **SourceType** is `net`, the value of this parameter is a CIDR block. Example: 192.0.XX.XX/24.
	// *   If the value of **SourceType** is `group`, the value of this parameter is the name of an address book. Example: db_group.
	// *   If the value of **SourceType** is `location`, the value of this parameter is a location. For more information about location codes, see [AddControlPolicy](~~138867~~). Example: \["BJ11", "ZB"].
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The CIDR blocks in the source address book.
	SourceGroupCidrs []*string `json:"SourceGroupCidrs,omitempty" xml:"SourceGroupCidrs,omitempty" type:"Repeated"`
	// The type of the source address book in the access control policy. The value is fixed as **ip**. The value indicates an address book that includes one or more CIDR blocks.
	SourceGroupType *string `json:"SourceGroupType,omitempty" xml:"SourceGroupType,omitempty"`
	// The type of the source address in the access control policy. Valid values:
	//
	// *   **net**: CIDR block
	// *   **group**: address book
	// *   **location**: location
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The total quota consumed by the returned access control policies, which is the sum of the quota consumed by each policy. The quota that is consumed by an access control policy is calculated by using the following formula: Quota that is consumed by an access control policy = Number of source addresses (number of CIDR blocks or regions) × Number of destination addresses (number of CIDR blocks, regions, or domain names) × Number of port ranges × Number of applications.
	SpreadCnt *string `json:"SpreadCnt,omitempty" xml:"SpreadCnt,omitempty"`
	// The time when the access control policy starts to take effect. The value is a UNIX timestamp. Unit: seconds. The start time must be on the hour or on the half hour, and at least 30 minutes earlier than the end time.
	//
	// >  If RepeatType is set to Permanent, this parameter is left empty. If RepeatType is set to None, Daily, Weekly, or Monthly, this parameter must be specified.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeNatFirewallControlPolicyResponseBodyPolicys) String() string {
	return tea.Prettify(s)
}

func (s DescribeNatFirewallControlPolicyResponseBodyPolicys) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetAclAction(v string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.AclAction = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetAclUuid(v string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.AclUuid = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetApplicationNameList(v []*string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.ApplicationNameList = v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetCreateTime(v int64) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.CreateTime = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetDescription(v string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.Description = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetDestPort(v string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.DestPort = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetDestPortGroup(v string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.DestPortGroup = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetDestPortGroupPorts(v []*string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.DestPortGroupPorts = v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetDestPortType(v string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.DestPortType = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetDestination(v string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.Destination = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetDestinationGroupCidrs(v []*string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.DestinationGroupCidrs = v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetDestinationGroupType(v string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.DestinationGroupType = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetDestinationType(v string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.DestinationType = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetDnsResult(v string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.DnsResult = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetDnsResultTime(v int64) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.DnsResultTime = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetDomainResolveType(v int32) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.DomainResolveType = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetEndTime(v int64) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.EndTime = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetHitLastTime(v int64) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.HitLastTime = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetHitTimes(v int32) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.HitTimes = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetModifyTime(v int64) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.ModifyTime = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetNatGatewayId(v string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetOrder(v int32) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.Order = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetProto(v string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.Proto = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetRelease(v string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.Release = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetRepeatDays(v []*int64) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.RepeatDays = v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetRepeatEndTime(v string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.RepeatEndTime = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetRepeatStartTime(v string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.RepeatStartTime = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetRepeatType(v string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.RepeatType = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetSource(v string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.Source = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetSourceGroupCidrs(v []*string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.SourceGroupCidrs = v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetSourceGroupType(v string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.SourceGroupType = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetSourceType(v string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.SourceType = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetSpreadCnt(v string) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.SpreadCnt = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponseBodyPolicys) SetStartTime(v int64) *DescribeNatFirewallControlPolicyResponseBodyPolicys {
	s.StartTime = &v
	return s
}

type DescribeNatFirewallControlPolicyResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNatFirewallControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNatFirewallControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNatFirewallControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallControlPolicyResponse) SetHeaders(v map[string]*string) *DescribeNatFirewallControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponse) SetStatusCode(v int32) *DescribeNatFirewallControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNatFirewallControlPolicyResponse) SetBody(v *DescribeNatFirewallControlPolicyResponseBody) *DescribeNatFirewallControlPolicyResponse {
	s.Body = v
	return s
}

type DescribeNatFirewallPolicyPriorUsedRequest struct {
	// The direction of the traffic to which the access control policy applies.
	//
	// Valid values:
	//
	// *   **out**: outbound traffic
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The IP version supported by the access control policy. Valid values:
	//
	// *   **4**: IPv4 (default)
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The language of the content within the request and the response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the NAT gateway.
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
}

func (s DescribeNatFirewallPolicyPriorUsedRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNatFirewallPolicyPriorUsedRequest) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallPolicyPriorUsedRequest) SetDirection(v string) *DescribeNatFirewallPolicyPriorUsedRequest {
	s.Direction = &v
	return s
}

func (s *DescribeNatFirewallPolicyPriorUsedRequest) SetIpVersion(v string) *DescribeNatFirewallPolicyPriorUsedRequest {
	s.IpVersion = &v
	return s
}

func (s *DescribeNatFirewallPolicyPriorUsedRequest) SetLang(v string) *DescribeNatFirewallPolicyPriorUsedRequest {
	s.Lang = &v
	return s
}

func (s *DescribeNatFirewallPolicyPriorUsedRequest) SetNatGatewayId(v string) *DescribeNatFirewallPolicyPriorUsedRequest {
	s.NatGatewayId = &v
	return s
}

type DescribeNatFirewallPolicyPriorUsedResponseBody struct {
	// The lowest priority for the access control policy.
	End *int32 `json:"End,omitempty" xml:"End,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The highest priority for the access control policy.
	Start *int32 `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s DescribeNatFirewallPolicyPriorUsedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNatFirewallPolicyPriorUsedResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallPolicyPriorUsedResponseBody) SetEnd(v int32) *DescribeNatFirewallPolicyPriorUsedResponseBody {
	s.End = &v
	return s
}

func (s *DescribeNatFirewallPolicyPriorUsedResponseBody) SetRequestId(v string) *DescribeNatFirewallPolicyPriorUsedResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNatFirewallPolicyPriorUsedResponseBody) SetStart(v int32) *DescribeNatFirewallPolicyPriorUsedResponseBody {
	s.Start = &v
	return s
}

type DescribeNatFirewallPolicyPriorUsedResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNatFirewallPolicyPriorUsedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNatFirewallPolicyPriorUsedResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNatFirewallPolicyPriorUsedResponse) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallPolicyPriorUsedResponse) SetHeaders(v map[string]*string) *DescribeNatFirewallPolicyPriorUsedResponse {
	s.Headers = v
	return s
}

func (s *DescribeNatFirewallPolicyPriorUsedResponse) SetStatusCode(v int32) *DescribeNatFirewallPolicyPriorUsedResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNatFirewallPolicyPriorUsedResponse) SetBody(v *DescribeNatFirewallPolicyPriorUsedResponseBody) *DescribeNatFirewallPolicyPriorUsedResponse {
	s.Body = v
	return s
}

type DescribeOutgoingDestinationIPRequest struct {
	// The application type in the access control policy. Valid values:
	//
	// *   **FTP**
	// *   **HTTP**
	// *   **HTTPS**
	// *   **Memcache**
	// *   **MongoDB**
	// *   **MQTT**
	// *   **MySQL**
	// *   **RDP**
	// *   **Redis**
	// *   **SMTP**
	// *   **SMTPS**
	// *   **SSH**
	// *   **SSL_No_Cert**
	// *   **SSL**
	// *   **VNC**
	//
	// >  The value of this parameter depends on the value of Proto. If you set Proto to TCP, you can set ApplicationNameList to any valid value. If you specify both ApplicationNameList and ApplicationName, only the value of ApplicationNameList is used.
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The ID of the service to which the destination IP address belongs. This parameter is left empty by default. Valid values:
	//
	// *   **All**: all services
	// *   **RiskDomain**: risky domain names
	// *   **RiskIP**: risky IP addresses
	// *   **AliYun**: Alibaba Cloud services
	// *   **NotAliYun**: third-party services
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The number of the page to return.
	//
	// Default value: 1.
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The destination IP address in the outbound connection that is initiated to access a domain name.
	DstIP *string `json:"DstIP,omitempty" xml:"DstIP,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// *   **zh** (default): Chinese
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The method that is used to sort the results. Valid values:
	//
	// *   **asc**: the ascending order.
	// *   **desc** (default): the descending order.
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The number of entries to return on each page.
	//
	// Default value: 6. Maximum value: 10.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The port number.
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The private IP address of the ECS instance that initiates the outbound connection.
	PrivateIP *string `json:"PrivateIP,omitempty" xml:"PrivateIP,omitempty"`
	// The public IP address of the Elastic Compute Service (ECS) instance that initiates the outbound connection.
	PublicIP *string `json:"PublicIP,omitempty" xml:"PublicIP,omitempty"`
	// The field based on which you want to sort the query results. Valid values:
	//
	// *   **SessionCount** (default): the number of requests.
	// *   **TotalBytes**: the total volume of traffic.
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the tag. Valid values:
	//
	// *   **AliYun**: Alibaba Cloud service
	// *   **RiskDomain**: risky domain name
	// *   **RiskIP**: risky IP address
	// *   **TrustedDomain**: trusted website
	// *   **AliPay**: Alipay
	// *   **DingDing**: DingTalk
	// *   **WeChat**: WeChat
	// *   **QQ**: Tencent QQ
	// *   **SecurityService**: security service
	// *   **Microsoft**: Microsoft
	// *   **Amazon**: Amazon Web Services (AWS)
	// *   **Pan**: cloud disk
	// *   **Map**: map
	// *   **Code**: code hosting
	// *   **SystemService**: system service
	// *   **Taobao**: Taobao
	// *   **Google**: Google
	// *   **ThirdPartyService**: third-party service
	// *   **FirstFlow**: the first time
	// *   **Downloader**: malicious download
	// *   **Alexa Top1M**: popular website
	// *   **Miner**: mining pool
	// *   **Intelligence**: threat intelligence
	// *   **DDoS**: DDoS trojan
	// *   **Ransomware**: ransomware
	// *   **Spyware**: spyware
	// *   **Rogue**: rogue software
	// *   **Botnet**: botnet
	// *   **Suspicious**: suspicious website
	// *   **C\&C**: command and control (C\&C)
	// *   **Gang**: gang
	// *   **CVE**: Common Vulnerabilities and Exposures (CVE)
	// *   **Backdoor**: webshell
	// *   **Phishing**: phishing website
	// *   **APT**: advanced persistent threat (APT) attack
	// *   **Supply Chain Attack**: supply chain attack
	// *   **Malicious software**: malware
	TagIdNew *string `json:"TagIdNew,omitempty" xml:"TagIdNew,omitempty"`
}

func (s DescribeOutgoingDestinationIPRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOutgoingDestinationIPRequest) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDestinationIPRequest) SetApplicationName(v string) *DescribeOutgoingDestinationIPRequest {
	s.ApplicationName = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetCategoryId(v string) *DescribeOutgoingDestinationIPRequest {
	s.CategoryId = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetCurrentPage(v string) *DescribeOutgoingDestinationIPRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetDstIP(v string) *DescribeOutgoingDestinationIPRequest {
	s.DstIP = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetEndTime(v string) *DescribeOutgoingDestinationIPRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetLang(v string) *DescribeOutgoingDestinationIPRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetOrder(v string) *DescribeOutgoingDestinationIPRequest {
	s.Order = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetPageSize(v string) *DescribeOutgoingDestinationIPRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetPort(v string) *DescribeOutgoingDestinationIPRequest {
	s.Port = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetPrivateIP(v string) *DescribeOutgoingDestinationIPRequest {
	s.PrivateIP = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetPublicIP(v string) *DescribeOutgoingDestinationIPRequest {
	s.PublicIP = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetSort(v string) *DescribeOutgoingDestinationIPRequest {
	s.Sort = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetStartTime(v string) *DescribeOutgoingDestinationIPRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetTagIdNew(v string) *DescribeOutgoingDestinationIPRequest {
	s.TagIdNew = &v
	return s
}

type DescribeOutgoingDestinationIPResponseBody struct {
	// The IP addresses in outbound connections.
	DstIPList []*DescribeOutgoingDestinationIPResponseBodyDstIPList `json:"DstIPList,omitempty" xml:"DstIPList,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of destination IP addresses in outbound connections.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeOutgoingDestinationIPResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOutgoingDestinationIPResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDestinationIPResponseBody) SetDstIPList(v []*DescribeOutgoingDestinationIPResponseBodyDstIPList) *DescribeOutgoingDestinationIPResponseBody {
	s.DstIPList = v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBody) SetRequestId(v string) *DescribeOutgoingDestinationIPResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBody) SetTotalCount(v int32) *DescribeOutgoingDestinationIPResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeOutgoingDestinationIPResponseBodyDstIPList struct {
	// Indicates whether an access control policy is configured. Valid values:
	//
	// *   **Uncovered**: no
	// *   **FullCoverage**: yes
	AclCoverage *string `json:"AclCoverage,omitempty" xml:"AclCoverage,omitempty"`
	// The suggestion to configure an access control policy.
	AclRecommendDetail *string `json:"AclRecommendDetail,omitempty" xml:"AclRecommendDetail,omitempty"`
	// The status of the access control policy. Valid values:
	//
	// *   **normal**: healthy
	// *   **Abnormal**: unhealthy
	AclStatus *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	// The information about the address book.
	AddressGroupList []*DescribeOutgoingDestinationIPResponseBodyDstIPListAddressGroupList `json:"AddressGroupList,omitempty" xml:"AddressGroupList,omitempty" type:"Repeated"`
	// The application ports.
	ApplicationPortList []*DescribeOutgoingDestinationIPResponseBodyDstIPListApplicationPortList `json:"ApplicationPortList,omitempty" xml:"ApplicationPortList,omitempty" type:"Repeated"`
	// The type of the tag. Valid values:
	//
	// *   **Suspicious**
	// *   **Malicious**
	// *   **Trusted**
	CategoryClassId *string `json:"CategoryClassId,omitempty" xml:"CategoryClassId,omitempty"`
	// The ID of the service type. Valid values:
	//
	// *   **Aliyun**: Alibaba Cloud services
	// *   **NotAliyun**: third-party services
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The type of the service to which the destination IP address belongs. Valid values:
	//
	// *   **Alibaba Cloud services**
	// *   **Third-party services**
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The destination IP addresses in outbound connections.
	DstIP *string `json:"DstIP,omitempty" xml:"DstIP,omitempty"`
	// The name of the group to which the access control policy belongs.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// Indicates whether an access control policy is configured. Valid values:
	//
	// *   **true**
	// *   **false**
	HasAcl *string `json:"HasAcl,omitempty" xml:"HasAcl,omitempty"`
	// Indicates whether an access control policy is recommended. Valid values:
	//
	// *   **true**
	// *   **false**
	HasAclRecommend *bool `json:"HasAclRecommend,omitempty" xml:"HasAclRecommend,omitempty"`
	// The inbound traffic. Unit: bytes.
	InBytes *int64 `json:"InBytes,omitempty" xml:"InBytes,omitempty"`
	// Indicates whether the destination IP address is added to a whitelist. Valid values:
	//
	// *   **true**
	// *   **false**
	IsMarkNormal *bool `json:"IsMarkNormal,omitempty" xml:"IsMarkNormal,omitempty"`
	// The outbound traffic. Unit: bytes.
	OutBytes *int64 `json:"OutBytes,omitempty" xml:"OutBytes,omitempty"`
	// The UUID of the access control policy.
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the access control policy.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The reason why the domain name is secure.
	SecurityReason *string `json:"SecurityReason,omitempty" xml:"SecurityReason,omitempty"`
	// The suggestion to handle the traffic of the domain name in outbound connections. Valid values:
	//
	// *   **pass**: allow
	// *   **alert**: deny
	// *   **drop**: monitor
	SecuritySuggest *string `json:"SecuritySuggest,omitempty" xml:"SecuritySuggest,omitempty"`
	// The number of requests.
	SessionCount *int64 `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
	// The tags.
	TagList []*DescribeOutgoingDestinationIPResponseBodyDstIPListTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The total traffic. Unit: bytes
	TotalBytes *string `json:"TotalBytes,omitempty" xml:"TotalBytes,omitempty"`
}

func (s DescribeOutgoingDestinationIPResponseBodyDstIPList) String() string {
	return tea.Prettify(s)
}

func (s DescribeOutgoingDestinationIPResponseBodyDstIPList) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetAclCoverage(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.AclCoverage = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetAclRecommendDetail(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.AclRecommendDetail = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetAclStatus(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.AclStatus = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetAddressGroupList(v []*DescribeOutgoingDestinationIPResponseBodyDstIPListAddressGroupList) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.AddressGroupList = v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetApplicationPortList(v []*DescribeOutgoingDestinationIPResponseBodyDstIPListApplicationPortList) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.ApplicationPortList = v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetCategoryClassId(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.CategoryClassId = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetCategoryId(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.CategoryId = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetCategoryName(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.CategoryName = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetDstIP(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.DstIP = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetGroupName(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.GroupName = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetHasAcl(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.HasAcl = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetHasAclRecommend(v bool) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.HasAclRecommend = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetInBytes(v int64) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.InBytes = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetIsMarkNormal(v bool) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.IsMarkNormal = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetOutBytes(v int64) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.OutBytes = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetRuleId(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.RuleId = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetRuleName(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.RuleName = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetSecurityReason(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.SecurityReason = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetSecuritySuggest(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.SecuritySuggest = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetSessionCount(v int64) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.SessionCount = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetTagList(v []*DescribeOutgoingDestinationIPResponseBodyDstIPListTagList) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.TagList = v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetTotalBytes(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.TotalBytes = &v
	return s
}

type DescribeOutgoingDestinationIPResponseBodyDstIPListAddressGroupList struct {
	// The name of the address book.
	AddressGroupName *string `json:"AddressGroupName,omitempty" xml:"AddressGroupName,omitempty"`
	// The UUID of the address book.
	AddressGroupUUID *string `json:"AddressGroupUUID,omitempty" xml:"AddressGroupUUID,omitempty"`
}

func (s DescribeOutgoingDestinationIPResponseBodyDstIPListAddressGroupList) String() string {
	return tea.Prettify(s)
}

func (s DescribeOutgoingDestinationIPResponseBodyDstIPListAddressGroupList) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListAddressGroupList) SetAddressGroupName(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPListAddressGroupList {
	s.AddressGroupName = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListAddressGroupList) SetAddressGroupUUID(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPListAddressGroupList {
	s.AddressGroupUUID = &v
	return s
}

type DescribeOutgoingDestinationIPResponseBodyDstIPListApplicationPortList struct {
	// The application type used in the access control policy. Valid values:
	//
	// *   **FTP**
	// *   **HTTP**
	// *   **HTTPS**
	// *   **Memcache**
	// *   **MongoDB**
	// *   **MQTT**
	// *   **MySQL**
	// *   **RDP**
	// *   **Redis**
	// *   **SMTP**
	// *   **SMTPS**
	// *   **SSH**
	// *   **SSL_No_Cert**
	// *   **SSL**
	// *   **VNC**
	//
	// >  The value of this parameter depends on the value of the Proto parameter. If you set Proto to TCP, you can set ApplicationNameList to any valid value. If you configure both ApplicationNameList and ApplicationName, only the value of ApplicationNameList is used.
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The application port.
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeOutgoingDestinationIPResponseBodyDstIPListApplicationPortList) String() string {
	return tea.Prettify(s)
}

func (s DescribeOutgoingDestinationIPResponseBodyDstIPListApplicationPortList) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListApplicationPortList) SetApplicationName(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPListApplicationPortList {
	s.ApplicationName = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListApplicationPortList) SetPort(v int32) *DescribeOutgoingDestinationIPResponseBodyDstIPListApplicationPortList {
	s.Port = &v
	return s
}

type DescribeOutgoingDestinationIPResponseBodyDstIPListTagList struct {
	// The type of the tag. Valid values:
	//
	// *   **Suspicious**
	// *   **Malicious**
	// *   **Trusted**
	ClassId *string `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	// The risk level. Valid values:
	//
	// *   **1**: low
	// *   **2**: medium
	// *   **3**: high
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The description of the tag.
	TagDescribe *string `json:"TagDescribe,omitempty" xml:"TagDescribe,omitempty"`
	// The ID of the tag.
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// The name of the tag.
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s DescribeOutgoingDestinationIPResponseBodyDstIPListTagList) String() string {
	return tea.Prettify(s)
}

func (s DescribeOutgoingDestinationIPResponseBodyDstIPListTagList) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListTagList) SetClassId(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPListTagList {
	s.ClassId = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListTagList) SetRiskLevel(v int32) *DescribeOutgoingDestinationIPResponseBodyDstIPListTagList {
	s.RiskLevel = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListTagList) SetTagDescribe(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPListTagList {
	s.TagDescribe = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListTagList) SetTagId(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPListTagList {
	s.TagId = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPListTagList) SetTagName(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPListTagList {
	s.TagName = &v
	return s
}

type DescribeOutgoingDestinationIPResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOutgoingDestinationIPResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOutgoingDestinationIPResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOutgoingDestinationIPResponse) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDestinationIPResponse) SetHeaders(v map[string]*string) *DescribeOutgoingDestinationIPResponse {
	s.Headers = v
	return s
}

func (s *DescribeOutgoingDestinationIPResponse) SetStatusCode(v int32) *DescribeOutgoingDestinationIPResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponse) SetBody(v *DescribeOutgoingDestinationIPResponseBody) *DescribeOutgoingDestinationIPResponse {
	s.Body = v
	return s
}

type DescribeOutgoingDomainRequest struct {
	// The type of the service. This parameter is empty by default. Valid values:
	//
	// *   **All**: all services
	// *   **RiskDomain**: risky domain names
	// *   **RiskIP**: risky IP addresses
	// *   **AliYun**: Alibaba Cloud services
	// *   **NotAliYun**: third-party services
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The number of the page to return.
	//
	// Default value: 1.
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The domain name in outbound connections.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the content within the request. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The method that is used to sort the results. Valid values:
	//
	// *   **asc**: the ascending order.
	// *   **desc** (default): the descending order.
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The number of entries to return on each page.
	//
	// Default value: 6. Maximum value: 100.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The public IP address of the Elastic Compute Service (ECS) instance that initiates outbound connections.
	PublicIP *string `json:"PublicIP,omitempty" xml:"PublicIP,omitempty"`
	// The field based on which you want to sort the query results. Valid values:
	//
	// *   **SessionCount** (default): the number of requests.
	// *   **TotalBytes**: the total volume of traffic.
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the tag. Valid values:
	//
	// *   **AliYun**: Alibaba Cloud service
	// *   **RiskDomain**: risky domain name
	// *   **RiskIP**: risky IP address
	// *   **TrustedDomain**: trusted website
	// *   **AliPay**: Alipay
	// *   **DingDing**: DingTalk
	// *   **WeChat**: WeChat
	// *   **QQ**: Tencent QQ
	// *   **SecurityService**: security service
	// *   **Microsoft**: Microsoft
	// *   **Amazon**: Amazon Web Services (AWS)
	// *   **Pan**: cloud disk
	// *   **Map**: map
	// *   **Code**: code hosting
	// *   **SystemService**: system service
	// *   **Taobao**: Taobao
	// *   **Google**: Google
	// *   **ThirdPartyService**: third-party service
	// *   **FirstFlow**: the first time when an outbound connection is initiated
	// *   **Downloader**: malicious download
	// *   **Alexa Top1M**: popular website
	// *   **Miner**: mining pool
	// *   **Intelligence**: threat intelligence
	// *   **DDoS**: DDoS trojan
	// *   **Ransomware**: ransomware
	// *   **Spyware**: spyware
	// *   **Rogue**: rogue software
	// *   **Botnet**: botnet
	// *   **Suspicious**: suspicious website
	// *   **C\&C**: command and control (C\&C)
	// *   **Gang**: gang
	// *   **CVE**: Common Vulnerabilities and Exposures (CVE)
	// *   **Backdoor**: webshell
	// *   **Phishing**: phishing website
	// *   **APT**: advanced persistent threat (APT) attack
	// *   **Supply Chain Attack**: supply chain attack
	// *   **Malicious software**: malware
	TagIdNew *string `json:"TagIdNew,omitempty" xml:"TagIdNew,omitempty"`
}

func (s DescribeOutgoingDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOutgoingDomainRequest) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDomainRequest) SetCategoryId(v string) *DescribeOutgoingDomainRequest {
	s.CategoryId = &v
	return s
}

func (s *DescribeOutgoingDomainRequest) SetCurrentPage(v string) *DescribeOutgoingDomainRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeOutgoingDomainRequest) SetDomain(v string) *DescribeOutgoingDomainRequest {
	s.Domain = &v
	return s
}

func (s *DescribeOutgoingDomainRequest) SetEndTime(v string) *DescribeOutgoingDomainRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeOutgoingDomainRequest) SetLang(v string) *DescribeOutgoingDomainRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOutgoingDomainRequest) SetOrder(v string) *DescribeOutgoingDomainRequest {
	s.Order = &v
	return s
}

func (s *DescribeOutgoingDomainRequest) SetPageSize(v string) *DescribeOutgoingDomainRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeOutgoingDomainRequest) SetPublicIP(v string) *DescribeOutgoingDomainRequest {
	s.PublicIP = &v
	return s
}

func (s *DescribeOutgoingDomainRequest) SetSort(v string) *DescribeOutgoingDomainRequest {
	s.Sort = &v
	return s
}

func (s *DescribeOutgoingDomainRequest) SetStartTime(v string) *DescribeOutgoingDomainRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeOutgoingDomainRequest) SetTagIdNew(v string) *DescribeOutgoingDomainRequest {
	s.TagIdNew = &v
	return s
}

type DescribeOutgoingDomainResponseBody struct {
	// An array that consists of the domain names in outbound connections.
	DomainList []*DescribeOutgoingDomainResponseBodyDomainList `json:"DomainList,omitempty" xml:"DomainList,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the domain names in outbound connections.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeOutgoingDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOutgoingDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDomainResponseBody) SetDomainList(v []*DescribeOutgoingDomainResponseBodyDomainList) *DescribeOutgoingDomainResponseBody {
	s.DomainList = v
	return s
}

func (s *DescribeOutgoingDomainResponseBody) SetRequestId(v string) *DescribeOutgoingDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBody) SetTotalCount(v int32) *DescribeOutgoingDomainResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeOutgoingDomainResponseBodyDomainList struct {
	// Indicates whether an access control policy is configured. Valid values:
	//
	// *   **Uncovered**: no
	// *   **FullCoverage**: yes
	AclCoverage *string `json:"AclCoverage,omitempty" xml:"AclCoverage,omitempty"`
	// The suggestion in an access control policy.
	AclRecommendDetail *string `json:"AclRecommendDetail,omitempty" xml:"AclRecommendDetail,omitempty"`
	// The state of the access control policy. Valid values:
	//
	// *   **normal**: healthy
	// *   **abnormal**: unhealthy
	AclStatus *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	// The name of the address book.
	AddressGroupName *string `json:"AddressGroupName,omitempty" xml:"AddressGroupName,omitempty"`
	// The UUID of the address book.
	AddressGroupUUID *string `json:"AddressGroupUUID,omitempty" xml:"AddressGroupUUID,omitempty"`
	// The website service.
	Business *string `json:"Business,omitempty" xml:"Business,omitempty"`
	// The type of the tag. Valid values:
	//
	// *   **Suspicious**
	// *   **Malicious**
	// *   **Trusted**
	CategoryClassId *string `json:"CategoryClassId,omitempty" xml:"CategoryClassId,omitempty"`
	// The type ID of the service to which the domain name belongs. Valid values:
	//
	// *   **Aliyun**: Alibaba Cloud services
	// *   **NotAliyun**: third-party services
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The type of the service to which the domain name belongs. Valid values:
	//
	// *   **Alibaba Cloud services**
	// *   **Third-party services**
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The domain name in outbound connections.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The name of the group to which the access control policy belongs.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// Indicates whether an `access control policy` is configured for the domain name. Valid values:
	//
	// *   **true**: yes
	// *   **false**: no
	HasAcl *string `json:"HasAcl,omitempty" xml:"HasAcl,omitempty"`
	// Indicates whether an access control policy is recommended. Valid values:
	//
	// *   **true**: yes
	// *   **false**: no
	HasAclRecommend *bool `json:"HasAclRecommend,omitempty" xml:"HasAclRecommend,omitempty"`
	// The volume of inbound traffic.
	InBytes *int64 `json:"InBytes,omitempty" xml:"InBytes,omitempty"`
	// Indicates whether the domain name is marked as normal. Valid values:
	//
	// *   **true**: normal
	// *   **false**: abnormal
	IsMarkNormal *bool `json:"IsMarkNormal,omitempty" xml:"IsMarkNormal,omitempty"`
	// The name of the organization.
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The volume of outbound traffic.
	OutBytes *int64 `json:"OutBytes,omitempty" xml:"OutBytes,omitempty"`
	// The ID of the access control policy.
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the access control policy.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The reason why the domain name is secure.
	SecurityReason *string `json:"SecurityReason,omitempty" xml:"SecurityReason,omitempty"`
	// The suggestion to handle the traffic of the domain name in outbound connections. Valid values:
	//
	// *   **pass**: allow
	// *   **alert**: monitor
	// *   **drop**: deny
	SecuritySuggest *string `json:"SecuritySuggest,omitempty" xml:"SecuritySuggest,omitempty"`
	// The number of requests.
	SessionCount *int64 `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
	// An array that consists of tags.
	TagList []*DescribeOutgoingDomainResponseBodyDomainListTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The total volume of traffic. Unit: bytes.
	TotalBytes *string `json:"TotalBytes,omitempty" xml:"TotalBytes,omitempty"`
}

func (s DescribeOutgoingDomainResponseBodyDomainList) String() string {
	return tea.Prettify(s)
}

func (s DescribeOutgoingDomainResponseBodyDomainList) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetAclCoverage(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.AclCoverage = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetAclRecommendDetail(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.AclRecommendDetail = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetAclStatus(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.AclStatus = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetAddressGroupName(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.AddressGroupName = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetAddressGroupUUID(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.AddressGroupUUID = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetBusiness(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.Business = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetCategoryClassId(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.CategoryClassId = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetCategoryId(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.CategoryId = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetCategoryName(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.CategoryName = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetDomain(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.Domain = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetGroupName(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.GroupName = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetHasAcl(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.HasAcl = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetHasAclRecommend(v bool) *DescribeOutgoingDomainResponseBodyDomainList {
	s.HasAclRecommend = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetInBytes(v int64) *DescribeOutgoingDomainResponseBodyDomainList {
	s.InBytes = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetIsMarkNormal(v bool) *DescribeOutgoingDomainResponseBodyDomainList {
	s.IsMarkNormal = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetOrganization(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.Organization = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetOutBytes(v int64) *DescribeOutgoingDomainResponseBodyDomainList {
	s.OutBytes = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetRuleId(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.RuleId = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetRuleName(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.RuleName = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetSecurityReason(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.SecurityReason = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetSecuritySuggest(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.SecuritySuggest = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetSessionCount(v int64) *DescribeOutgoingDomainResponseBodyDomainList {
	s.SessionCount = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetTagList(v []*DescribeOutgoingDomainResponseBodyDomainListTagList) *DescribeOutgoingDomainResponseBodyDomainList {
	s.TagList = v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainList) SetTotalBytes(v string) *DescribeOutgoingDomainResponseBodyDomainList {
	s.TotalBytes = &v
	return s
}

type DescribeOutgoingDomainResponseBodyDomainListTagList struct {
	// The type of the tag. Valid values:
	//
	// *   **Suspicious**
	// *   **Malicious**
	// *   **Trusted**
	ClassId *string `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
	// The risk level. Valid values:
	//
	// *   **1**: low
	// *   **2**: medium
	// *   **3**: high
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The description of the tag.
	TagDescribe *string `json:"TagDescribe,omitempty" xml:"TagDescribe,omitempty"`
	// The ID of the tag.
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// The name of the tag.
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s DescribeOutgoingDomainResponseBodyDomainListTagList) String() string {
	return tea.Prettify(s)
}

func (s DescribeOutgoingDomainResponseBodyDomainListTagList) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDomainResponseBodyDomainListTagList) SetClassId(v string) *DescribeOutgoingDomainResponseBodyDomainListTagList {
	s.ClassId = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainListTagList) SetRiskLevel(v int32) *DescribeOutgoingDomainResponseBodyDomainListTagList {
	s.RiskLevel = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainListTagList) SetTagDescribe(v string) *DescribeOutgoingDomainResponseBodyDomainListTagList {
	s.TagDescribe = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainListTagList) SetTagId(v string) *DescribeOutgoingDomainResponseBodyDomainListTagList {
	s.TagId = &v
	return s
}

func (s *DescribeOutgoingDomainResponseBodyDomainListTagList) SetTagName(v string) *DescribeOutgoingDomainResponseBodyDomainListTagList {
	s.TagName = &v
	return s
}

type DescribeOutgoingDomainResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOutgoingDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOutgoingDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOutgoingDomainResponse) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDomainResponse) SetHeaders(v map[string]*string) *DescribeOutgoingDomainResponse {
	s.Headers = v
	return s
}

func (s *DescribeOutgoingDomainResponse) SetStatusCode(v int32) *DescribeOutgoingDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOutgoingDomainResponse) SetBody(v *DescribeOutgoingDomainResponseBody) *DescribeOutgoingDomainResponse {
	s.Body = v
	return s
}

type DescribePolicyAdvancedConfigRequest struct {
	// The natural language of the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Deprecated
	// The source IP address of the request.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribePolicyAdvancedConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyAdvancedConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribePolicyAdvancedConfigRequest) SetLang(v string) *DescribePolicyAdvancedConfigRequest {
	s.Lang = &v
	return s
}

func (s *DescribePolicyAdvancedConfigRequest) SetSourceIp(v string) *DescribePolicyAdvancedConfigRequest {
	s.SourceIp = &v
	return s
}

type DescribePolicyAdvancedConfigResponseBody struct {
	// Indicates whether the strict mode is enabled for the access control policy. Valid values:
	//
	// *   **on**: The strict mode is enabled.
	// *   **off**: The strict mode is disabled.
	InternetSwitch *string `json:"InternetSwitch,omitempty" xml:"InternetSwitch,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePolicyAdvancedConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyAdvancedConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolicyAdvancedConfigResponseBody) SetInternetSwitch(v string) *DescribePolicyAdvancedConfigResponseBody {
	s.InternetSwitch = &v
	return s
}

func (s *DescribePolicyAdvancedConfigResponseBody) SetRequestId(v string) *DescribePolicyAdvancedConfigResponseBody {
	s.RequestId = &v
	return s
}

type DescribePolicyAdvancedConfigResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolicyAdvancedConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolicyAdvancedConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyAdvancedConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribePolicyAdvancedConfigResponse) SetHeaders(v map[string]*string) *DescribePolicyAdvancedConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribePolicyAdvancedConfigResponse) SetStatusCode(v int32) *DescribePolicyAdvancedConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolicyAdvancedConfigResponse) SetBody(v *DescribePolicyAdvancedConfigResponseBody) *DescribePolicyAdvancedConfigResponse {
	s.Body = v
	return s
}

type DescribePolicyPriorUsedRequest struct {
	// The direction of the traffic to which the access control policy applies.
	//
	// Valid values:
	//
	// *   **in**: inbound traffic
	// *   **out**: outbound traffic
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The IP version of the asset that is protected by Cloud Firewall.
	//
	// Valid values:
	//
	// *   **4**: IPv4 (default)
	// *   **6**: IPv6
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The natural language of the request and response.
	//
	// Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Deprecated
	// The source IP address of the request.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribePolicyPriorUsedRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyPriorUsedRequest) GoString() string {
	return s.String()
}

func (s *DescribePolicyPriorUsedRequest) SetDirection(v string) *DescribePolicyPriorUsedRequest {
	s.Direction = &v
	return s
}

func (s *DescribePolicyPriorUsedRequest) SetIpVersion(v string) *DescribePolicyPriorUsedRequest {
	s.IpVersion = &v
	return s
}

func (s *DescribePolicyPriorUsedRequest) SetLang(v string) *DescribePolicyPriorUsedRequest {
	s.Lang = &v
	return s
}

func (s *DescribePolicyPriorUsedRequest) SetSourceIp(v string) *DescribePolicyPriorUsedRequest {
	s.SourceIp = &v
	return s
}

type DescribePolicyPriorUsedResponseBody struct {
	// The lowest priority of existing access control policies.
	//
	// >  The value -1 indicates the lowest priority.
	End *int32 `json:"End,omitempty" xml:"End,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The highest priority of existing access control policies.
	//
	// >  The value 0 indicates the highest priority.
	Start *int32 `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s DescribePolicyPriorUsedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyPriorUsedResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolicyPriorUsedResponseBody) SetEnd(v int32) *DescribePolicyPriorUsedResponseBody {
	s.End = &v
	return s
}

func (s *DescribePolicyPriorUsedResponseBody) SetRequestId(v string) *DescribePolicyPriorUsedResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePolicyPriorUsedResponseBody) SetStart(v int32) *DescribePolicyPriorUsedResponseBody {
	s.Start = &v
	return s
}

type DescribePolicyPriorUsedResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolicyPriorUsedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolicyPriorUsedResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyPriorUsedResponse) GoString() string {
	return s.String()
}

func (s *DescribePolicyPriorUsedResponse) SetHeaders(v map[string]*string) *DescribePolicyPriorUsedResponse {
	s.Headers = v
	return s
}

func (s *DescribePolicyPriorUsedResponse) SetStatusCode(v int32) *DescribePolicyPriorUsedResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolicyPriorUsedResponse) SetBody(v *DescribePolicyPriorUsedResponseBody) *DescribePolicyPriorUsedResponse {
	s.Body = v
	return s
}

type DescribePostpayTrafficDetailRequest struct {
	// The page number. Default value: 1.
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// *   **zh** (default): Chinese
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The field based on which you want to sort the query results. Valid values:
	//
	// *   **resourceId**
	// *   **trafficDay**
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 50.
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The instance ID or the IP address of the asset.
	SearchItem *string `json:"SearchItem,omitempty" xml:"SearchItem,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The traffic type. This parameter is required. Valid values:
	//
	// *   **EIP_TRAFFIC**: traffic for the Internet firewall
	// *   **NatGateway_TRAFFIC**: traffic for the NAT firewall
	TrafficType *string `json:"TrafficType,omitempty" xml:"TrafficType,omitempty"`
}

func (s DescribePostpayTrafficDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePostpayTrafficDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribePostpayTrafficDetailRequest) SetCurrentPage(v int64) *DescribePostpayTrafficDetailRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePostpayTrafficDetailRequest) SetEndTime(v string) *DescribePostpayTrafficDetailRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePostpayTrafficDetailRequest) SetLang(v string) *DescribePostpayTrafficDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribePostpayTrafficDetailRequest) SetOrder(v string) *DescribePostpayTrafficDetailRequest {
	s.Order = &v
	return s
}

func (s *DescribePostpayTrafficDetailRequest) SetPageSize(v int64) *DescribePostpayTrafficDetailRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePostpayTrafficDetailRequest) SetSearchItem(v string) *DescribePostpayTrafficDetailRequest {
	s.SearchItem = &v
	return s
}

func (s *DescribePostpayTrafficDetailRequest) SetStartTime(v string) *DescribePostpayTrafficDetailRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePostpayTrafficDetailRequest) SetTrafficType(v string) *DescribePostpayTrafficDetailRequest {
	s.TrafficType = &v
	return s
}

type DescribePostpayTrafficDetailResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The traffic statistics.
	TrafficList []*DescribePostpayTrafficDetailResponseBodyTrafficList `json:"TrafficList,omitempty" xml:"TrafficList,omitempty" type:"Repeated"`
}

func (s DescribePostpayTrafficDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePostpayTrafficDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePostpayTrafficDetailResponseBody) SetRequestId(v string) *DescribePostpayTrafficDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePostpayTrafficDetailResponseBody) SetTotalCount(v int32) *DescribePostpayTrafficDetailResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePostpayTrafficDetailResponseBody) SetTrafficList(v []*DescribePostpayTrafficDetailResponseBodyTrafficList) *DescribePostpayTrafficDetailResponseBody {
	s.TrafficList = v
	return s
}

type DescribePostpayTrafficDetailResponseBodyTrafficList struct {
	// The inbound network throughput, which indicates the total number of bytes that are received. Unit: bytes.
	InBytes *int64 `json:"InBytes,omitempty" xml:"InBytes,omitempty"`
	// The instance ID of the asset.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the asset. This value takes effect only for the Internet firewall.
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The outbound network throughput, which indicates the total number of bytes that are sent. Unit: bytes.
	OutBytes *int64 `json:"OutBytes,omitempty" xml:"OutBytes,omitempty"`
	// The resource ID. The resource ID for the Internet firewall is the public IP address that is protected the Internet firewall, and the resource ID for a NAT firewall is the instance ID of the NAT firewall.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The total inbound and outbound network throughput, which indicates the total number of bytes that are sent and received. Unit: bytes.
	TotalBytes *int64 `json:"TotalBytes,omitempty" xml:"TotalBytes,omitempty"`
	// The date on which the statistics are collected.
	TrafficDay *string `json:"TrafficDay,omitempty" xml:"TrafficDay,omitempty"`
	// The traffic type. Valid values:
	//
	// *   **EIP_TRAFFIC**: traffic for the Internet firewall
	// *   **NatGateway_TRAFFIC**: traffic for the NAT firewall
	TrafficType *string `json:"TrafficType,omitempty" xml:"TrafficType,omitempty"`
}

func (s DescribePostpayTrafficDetailResponseBodyTrafficList) String() string {
	return tea.Prettify(s)
}

func (s DescribePostpayTrafficDetailResponseBodyTrafficList) GoString() string {
	return s.String()
}

func (s *DescribePostpayTrafficDetailResponseBodyTrafficList) SetInBytes(v int64) *DescribePostpayTrafficDetailResponseBodyTrafficList {
	s.InBytes = &v
	return s
}

func (s *DescribePostpayTrafficDetailResponseBodyTrafficList) SetInstanceId(v string) *DescribePostpayTrafficDetailResponseBodyTrafficList {
	s.InstanceId = &v
	return s
}

func (s *DescribePostpayTrafficDetailResponseBodyTrafficList) SetInstanceType(v string) *DescribePostpayTrafficDetailResponseBodyTrafficList {
	s.InstanceType = &v
	return s
}

func (s *DescribePostpayTrafficDetailResponseBodyTrafficList) SetOutBytes(v int64) *DescribePostpayTrafficDetailResponseBodyTrafficList {
	s.OutBytes = &v
	return s
}

func (s *DescribePostpayTrafficDetailResponseBodyTrafficList) SetResourceId(v string) *DescribePostpayTrafficDetailResponseBodyTrafficList {
	s.ResourceId = &v
	return s
}

func (s *DescribePostpayTrafficDetailResponseBodyTrafficList) SetTotalBytes(v int64) *DescribePostpayTrafficDetailResponseBodyTrafficList {
	s.TotalBytes = &v
	return s
}

func (s *DescribePostpayTrafficDetailResponseBodyTrafficList) SetTrafficDay(v string) *DescribePostpayTrafficDetailResponseBodyTrafficList {
	s.TrafficDay = &v
	return s
}

func (s *DescribePostpayTrafficDetailResponseBodyTrafficList) SetTrafficType(v string) *DescribePostpayTrafficDetailResponseBodyTrafficList {
	s.TrafficType = &v
	return s
}

type DescribePostpayTrafficDetailResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePostpayTrafficDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePostpayTrafficDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePostpayTrafficDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribePostpayTrafficDetailResponse) SetHeaders(v map[string]*string) *DescribePostpayTrafficDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribePostpayTrafficDetailResponse) SetStatusCode(v int32) *DescribePostpayTrafficDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePostpayTrafficDetailResponse) SetBody(v *DescribePostpayTrafficDetailResponseBody) *DescribePostpayTrafficDetailResponse {
	s.Body = v
	return s
}

type DescribePostpayTrafficTotalRequest struct {
	// The language of the content within the response. Valid values:
	//
	// *   **zh** (default): Chinese
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribePostpayTrafficTotalRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePostpayTrafficTotalRequest) GoString() string {
	return s.String()
}

func (s *DescribePostpayTrafficTotalRequest) SetLang(v string) *DescribePostpayTrafficTotalRequest {
	s.Lang = &v
	return s
}

type DescribePostpayTrafficTotalResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the assets that are protected by the Internet firewall.
	TotalAssets *int64 `json:"TotalAssets,omitempty" xml:"TotalAssets,omitempty"`
	// The total number of the assets that are protected by the NAT firewall.
	TotalNatAssets *int64 `json:"TotalNatAssets,omitempty" xml:"TotalNatAssets,omitempty"`
	// The total traffic for the NAT firewall. Unit: bytes.
	TotalNatTraffic *int64 `json:"TotalNatTraffic,omitempty" xml:"TotalNatTraffic,omitempty"`
	// The total traffic for the Internet firewall. Unit: bytes.
	TotalTraffic *int64 `json:"TotalTraffic,omitempty" xml:"TotalTraffic,omitempty"`
}

func (s DescribePostpayTrafficTotalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePostpayTrafficTotalResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePostpayTrafficTotalResponseBody) SetRequestId(v string) *DescribePostpayTrafficTotalResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePostpayTrafficTotalResponseBody) SetTotalAssets(v int64) *DescribePostpayTrafficTotalResponseBody {
	s.TotalAssets = &v
	return s
}

func (s *DescribePostpayTrafficTotalResponseBody) SetTotalNatAssets(v int64) *DescribePostpayTrafficTotalResponseBody {
	s.TotalNatAssets = &v
	return s
}

func (s *DescribePostpayTrafficTotalResponseBody) SetTotalNatTraffic(v int64) *DescribePostpayTrafficTotalResponseBody {
	s.TotalNatTraffic = &v
	return s
}

func (s *DescribePostpayTrafficTotalResponseBody) SetTotalTraffic(v int64) *DescribePostpayTrafficTotalResponseBody {
	s.TotalTraffic = &v
	return s
}

type DescribePostpayTrafficTotalResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePostpayTrafficTotalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePostpayTrafficTotalResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePostpayTrafficTotalResponse) GoString() string {
	return s.String()
}

func (s *DescribePostpayTrafficTotalResponse) SetHeaders(v map[string]*string) *DescribePostpayTrafficTotalResponse {
	s.Headers = v
	return s
}

func (s *DescribePostpayTrafficTotalResponse) SetStatusCode(v int32) *DescribePostpayTrafficTotalResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePostpayTrafficTotalResponse) SetBody(v *DescribePostpayTrafficTotalResponseBody) *DescribePostpayTrafficTotalResponse {
	s.Body = v
	return s
}

type DescribePrefixListsRequest struct {
	// The region ID of the instance.
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The source IP address of the request.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribePrefixListsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePrefixListsRequest) GoString() string {
	return s.String()
}

func (s *DescribePrefixListsRequest) SetRegionNo(v string) *DescribePrefixListsRequest {
	s.RegionNo = &v
	return s
}

func (s *DescribePrefixListsRequest) SetSourceIp(v string) *DescribePrefixListsRequest {
	s.SourceIp = &v
	return s
}

type DescribePrefixListsResponseBody struct {
	// Details about the prefix lists.
	PrefixList []*DescribePrefixListsResponseBodyPrefixList `json:"PrefixList,omitempty" xml:"PrefixList,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePrefixListsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePrefixListsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePrefixListsResponseBody) SetPrefixList(v []*DescribePrefixListsResponseBodyPrefixList) *DescribePrefixListsResponseBody {
	s.PrefixList = v
	return s
}

func (s *DescribePrefixListsResponseBody) SetRequestId(v string) *DescribePrefixListsResponseBody {
	s.RequestId = &v
	return s
}

type DescribePrefixListsResponseBodyPrefixList struct {
	// The IP address family of the prefix list. Valid values:
	//
	// *   IPv4
	// *   IPv6
	AddressFamily *string `json:"AddressFamily,omitempty" xml:"AddressFamily,omitempty"`
	// The number of associated resources.
	AssociationCount *int32 `json:"AssociationCount,omitempty" xml:"AssociationCount,omitempty"`
	// The creation time.
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The maximum number of entries in the prefix list.
	MaxEntries *int32 `json:"MaxEntries,omitempty" xml:"MaxEntries,omitempty"`
	// The ID of the prefix list.
	PrefixListId *string `json:"PrefixListId,omitempty" xml:"PrefixListId,omitempty"`
	// The name of the prefix list.
	PrefixListName *string `json:"PrefixListName,omitempty" xml:"PrefixListName,omitempty"`
}

func (s DescribePrefixListsResponseBodyPrefixList) String() string {
	return tea.Prettify(s)
}

func (s DescribePrefixListsResponseBodyPrefixList) GoString() string {
	return s.String()
}

func (s *DescribePrefixListsResponseBodyPrefixList) SetAddressFamily(v string) *DescribePrefixListsResponseBodyPrefixList {
	s.AddressFamily = &v
	return s
}

func (s *DescribePrefixListsResponseBodyPrefixList) SetAssociationCount(v int32) *DescribePrefixListsResponseBodyPrefixList {
	s.AssociationCount = &v
	return s
}

func (s *DescribePrefixListsResponseBodyPrefixList) SetCreationTime(v string) *DescribePrefixListsResponseBodyPrefixList {
	s.CreationTime = &v
	return s
}

func (s *DescribePrefixListsResponseBodyPrefixList) SetDescription(v string) *DescribePrefixListsResponseBodyPrefixList {
	s.Description = &v
	return s
}

func (s *DescribePrefixListsResponseBodyPrefixList) SetMaxEntries(v int32) *DescribePrefixListsResponseBodyPrefixList {
	s.MaxEntries = &v
	return s
}

func (s *DescribePrefixListsResponseBodyPrefixList) SetPrefixListId(v string) *DescribePrefixListsResponseBodyPrefixList {
	s.PrefixListId = &v
	return s
}

func (s *DescribePrefixListsResponseBodyPrefixList) SetPrefixListName(v string) *DescribePrefixListsResponseBodyPrefixList {
	s.PrefixListName = &v
	return s
}

type DescribePrefixListsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePrefixListsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePrefixListsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePrefixListsResponse) GoString() string {
	return s.String()
}

func (s *DescribePrefixListsResponse) SetHeaders(v map[string]*string) *DescribePrefixListsResponse {
	s.Headers = v
	return s
}

func (s *DescribePrefixListsResponse) SetStatusCode(v int32) *DescribePrefixListsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePrefixListsResponse) SetBody(v *DescribePrefixListsResponseBody) *DescribePrefixListsResponse {
	s.Body = v
	return s
}

type DescribeRiskEventGroupRequest struct {
	// The names of attacked applications. Set the value in the `["AttackApp1","AttackApp2"]` format.
	AttackApp []*string `json:"AttackApp,omitempty" xml:"AttackApp,omitempty" type:"Repeated"`
	// The attack type of the intrusion events. Valid values:
	//
	// *   **1**: suspicious connection
	// *   **2**: command execution
	// *   **3**: brute-force attack
	// *   **4**: scanning
	// *   **5**: others
	// *   **6**: information leak
	// *   **7**: DoS attack
	// *   **8**: buffer overflow attack
	// *   **9**: web attack
	// *   **10**: trojan backdoor
	// *   **11**: computer worm
	// *   **12**: mining
	// *   **13**: reverse shell
	//
	// > If you do not specify this parameter, the intrusion events of all attack types are queried.
	AttackType *string `json:"AttackType,omitempty" xml:"AttackType,omitempty"`
	// The edition of Cloud Firewall that you purchase. Valid values:
	//
	// *   **2**: Premium Edition
	// *   **3**: Enterprise Edition
	// *   **4**: Ultimate Edition
	// *   **10**: Cloud Firewall that uses the pay-as-you-go billing method
	BuyVersion *int64 `json:"BuyVersion,omitempty" xml:"BuyVersion,omitempty"`
	// The number of the page to return. Default value: **1**.
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The type of the risk events.\
	// Set the value to **session**, which indicates intrusion events.
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The direction of the traffic for the intrusion events. Valid values:
	//
	// *   **in**: inbound
	// *   **out**: outbound
	//
	// > If you do not specify this parameter, the intrusion events that are recorded for both inbound and outbound traffic are queried.
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The destination IP address to query. If you specify this parameter, all intrusion events with the specified destination IP address are queried.
	DstIP *string `json:"DstIP,omitempty" xml:"DstIP,omitempty"`
	// The ID of the destination VPC.
	//
	// > If the FirewallType parameter is set to VpcFirewall, you must specify this parameter.
	DstNetworkInstanceId *string `json:"DstNetworkInstanceId,omitempty" xml:"DstNetworkInstanceId,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the intrusion event.
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The type of the firewall. Valid values:
	//
	// *   **VpcFirewall**: virtual private cloud (VPC) firewall
	// *   **InternetFirewall**: Internet firewall (default)
	FirewallType *string `json:"FirewallType,omitempty" xml:"FirewallType,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Specifies whether to query the information about the geographical locations of IP addresses.
	//
	// *   **true**: does not query the information about the geographical locations of IP addresses.
	// *   **false**: queries the information about the geographical locations of IP addresses. This is the default value.
	NoLocation *string `json:"NoLocation,omitempty" xml:"NoLocation,omitempty"`
	// The order in which you want to sort the results. Valid values:
	//
	// *   **asc**: the ascending order.
	// *   **desc**: the descending order. This is the default value.
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The number of entries to return on each page.
	//
	// Default value: **6**. Maximum value: **10**.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The status of the firewall. Valid values:
	//
	// *   **1**: alerting
	// *   **2**: blocking
	//
	// > If you do not specify this parameter, all intrusion events that are detected by the firewall are queried, regardless of the firewall status.
	RuleResult *string `json:"RuleResult,omitempty" xml:"RuleResult,omitempty"`
	// The module of the rule that is used to detect the intrusion events. Valid values:
	//
	// *   **1**: basic protection
	// *   **2**: virtual patching
	// *   **4**: threat intelligence
	//
	// > If you do not specify this parameter, the intrusion events that are detected by all rules are queried.
	RuleSource *string `json:"RuleSource,omitempty" xml:"RuleSource,omitempty"`
	// The field based on which you want to sort the results. Valid values:
	//
	// *   **VulLevel**: The results are sorted based on the risk level field. This is the default value.
	// *   **LastTime**: The results are sorted based on the most recent occurrence time.
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The source IP address to query. If you specify this parameter, all intrusion events with the specified source IP address are queried.
	SrcIP *string `json:"SrcIP,omitempty" xml:"SrcIP,omitempty"`
	// The ID of the source VPC.
	//
	// > If the FirewallType parameter is set to VpcFirewall, you must specify this parameter.
	SrcNetworkInstanceId *string `json:"SrcNetworkInstanceId,omitempty" xml:"SrcNetworkInstanceId,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The risk level of the intrusion events. Valid values:
	//
	// *   **1**: low
	// *   **2**: medium
	// *   **3**: high
	//
	// > If you do not specify this parameter, the intrusion events that are at all risk levels are queried.
	VulLevel *string `json:"VulLevel,omitempty" xml:"VulLevel,omitempty"`
}

func (s DescribeRiskEventGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskEventGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventGroupRequest) SetAttackApp(v []*string) *DescribeRiskEventGroupRequest {
	s.AttackApp = v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetAttackType(v string) *DescribeRiskEventGroupRequest {
	s.AttackType = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetBuyVersion(v int64) *DescribeRiskEventGroupRequest {
	s.BuyVersion = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetCurrentPage(v string) *DescribeRiskEventGroupRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetDataType(v string) *DescribeRiskEventGroupRequest {
	s.DataType = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetDirection(v string) *DescribeRiskEventGroupRequest {
	s.Direction = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetDstIP(v string) *DescribeRiskEventGroupRequest {
	s.DstIP = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetDstNetworkInstanceId(v string) *DescribeRiskEventGroupRequest {
	s.DstNetworkInstanceId = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetEndTime(v string) *DescribeRiskEventGroupRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetEventName(v string) *DescribeRiskEventGroupRequest {
	s.EventName = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetFirewallType(v string) *DescribeRiskEventGroupRequest {
	s.FirewallType = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetLang(v string) *DescribeRiskEventGroupRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetNoLocation(v string) *DescribeRiskEventGroupRequest {
	s.NoLocation = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetOrder(v string) *DescribeRiskEventGroupRequest {
	s.Order = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetPageSize(v string) *DescribeRiskEventGroupRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetRuleResult(v string) *DescribeRiskEventGroupRequest {
	s.RuleResult = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetRuleSource(v string) *DescribeRiskEventGroupRequest {
	s.RuleSource = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetSort(v string) *DescribeRiskEventGroupRequest {
	s.Sort = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetSrcIP(v string) *DescribeRiskEventGroupRequest {
	s.SrcIP = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetSrcNetworkInstanceId(v string) *DescribeRiskEventGroupRequest {
	s.SrcNetworkInstanceId = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetStartTime(v string) *DescribeRiskEventGroupRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRiskEventGroupRequest) SetVulLevel(v string) *DescribeRiskEventGroupRequest {
	s.VulLevel = &v
	return s
}

type DescribeRiskEventGroupResponseBody struct {
	// An array that consists of the details of the intrusion events.
	DataList []*DescribeRiskEventGroupResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of risk events.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRiskEventGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskEventGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventGroupResponseBody) SetDataList(v []*DescribeRiskEventGroupResponseBodyDataList) *DescribeRiskEventGroupResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeRiskEventGroupResponseBody) SetRequestId(v string) *DescribeRiskEventGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBody) SetTotalCount(v int32) *DescribeRiskEventGroupResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeRiskEventGroupResponseBodyDataList struct {
	// The name of the attacked application.
	AttackApp *string `json:"AttackApp,omitempty" xml:"AttackApp,omitempty"`
	// The attack type of the intrusion event. Valid values:
	//
	// *   **1**: suspicious connection
	// *   **2**: command execution
	// *   **3**: brute-force attack
	// *   **4**: scanning
	// *   **5**: others
	// *   **6**: information leak
	// *   **7**: DoS attack
	// *   **8**: buffer overflow attack
	// *   **9**: web attack
	// *   **10**: trojan backdoor
	// *   **11**: computer worm
	// *   **12**: mining
	// *   **13**: reverse shell
	AttackType *int32 `json:"AttackType,omitempty" xml:"AttackType,omitempty"`
	// The description of the intrusion event.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The direction of the traffic for the intrusion event. Valid values:
	//
	// *   **in**: inbound
	// *   **out**: outbound
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The destination IP address that is included in the intrusion event.
	DstIP *string `json:"DstIP,omitempty" xml:"DstIP,omitempty"`
	// The number of intrusion events.
	EventCount *int32 `json:"EventCount,omitempty" xml:"EventCount,omitempty"`
	// The ID of the intrusion event.
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The name of the intrusion event.
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The time when the intrusion event was first detected. The value is a UNIX timestamp. Unit: seconds.
	FirstEventTime *int32 `json:"FirstEventTime,omitempty" xml:"FirstEventTime,omitempty"`
	// The geographical information about the IP address. The value is a struct that contains the following parameters: **CityId**, **CityName**, **CountryId**, and **CountryName**.\
	// ****************
	IPLocationInfo *DescribeRiskEventGroupResponseBodyDataListIPLocationInfo `json:"IPLocationInfo,omitempty" xml:"IPLocationInfo,omitempty" type:"Struct"`
	// The time when the intrusion event was last detected. The value is a UNIX timestamp. Unit: seconds.
	LastEventTime *int32 `json:"LastEventTime,omitempty" xml:"LastEventTime,omitempty"`
	// The information about the private IP address in the intrusion event. The value is an array that contains the following parameters: **RegionNo**, **ResourceInstanceId**, **ResourceInstanceName**, and **ResourcePrivateIP**.\
	// ****************
	ResourcePrivateIPList []*DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList `json:"ResourcePrivateIPList,omitempty" xml:"ResourcePrivateIPList,omitempty" type:"Repeated"`
	// The type of the public IP address in the intrusion event. Valid values:
	//
	// *   **EIP**: the elastic IP address (EIP)
	// *   **EcsPublicIP**: the public IP address of an Elastic Compute Service (ECS) instance
	// *   **EcsEIP**: the EIP of an ECS instance
	// *   **NatPublicIP**: the public IP address of a NAT gateway
	// *   **NatEIP**: the EIP of a NAT gateway
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The ID of the rule that is used to detect the intrusion event.
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The status of the firewall. Valid values:
	//
	// *   **1**: alerting
	// *   **2**: blocking
	RuleResult *int32 `json:"RuleResult,omitempty" xml:"RuleResult,omitempty"`
	// The module of the rule that is used to detect the intrusion event. Valid values:
	//
	// *   **1**: basic protection
	// *   **2**: virtual patching
	// *   **4**: threat intelligence
	RuleSource *int32 `json:"RuleSource,omitempty" xml:"RuleSource,omitempty"`
	// The source IP address that is included in the intrusion event.
	SrcIP *string `json:"SrcIP,omitempty" xml:"SrcIP,omitempty"`
	// The tag added to the source IP address. The tag helps identify whether the source IP address is a back-to-origin IP address for a cloud service.
	SrcIPTag *string `json:"SrcIPTag,omitempty" xml:"SrcIPTag,omitempty"`
	// An array that consists of the source private IP addresses in the intrusion event.
	SrcPrivateIPList []*string `json:"SrcPrivateIPList,omitempty" xml:"SrcPrivateIPList,omitempty" type:"Repeated"`
	// The tag added to the threat intelligence that is provided for major events.
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The information about the destination VPC of the intrusion event. The value is a struct that contains the following parameters: **EcsInstanceId**, **EcsInstanceName**, **NetworkInstanceId**, **NetworkInstanceName**, and **RegionNo**.\
	// ********************
	VpcDstInfo *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo `json:"VpcDstInfo,omitempty" xml:"VpcDstInfo,omitempty" type:"Struct"`
	// The information about the source VPC of the intrusion event. The value is a struct that contains the following parameters: **EcsInstanceId**, **EcsInstanceName**, **NetworkInstanceId**, **NetworkInstanceName**, and **RegionNo**.\
	// ********************
	VpcSrcInfo *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo `json:"VpcSrcInfo,omitempty" xml:"VpcSrcInfo,omitempty" type:"Struct"`
	// The risk level of the intrusion event. Valid values:
	//
	// *   **1**: low
	// *   **2**: medium
	// *   **3**: high
	VulLevel *int32 `json:"VulLevel,omitempty" xml:"VulLevel,omitempty"`
}

func (s DescribeRiskEventGroupResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskEventGroupResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetAttackApp(v string) *DescribeRiskEventGroupResponseBodyDataList {
	s.AttackApp = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetAttackType(v int32) *DescribeRiskEventGroupResponseBodyDataList {
	s.AttackType = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetDescription(v string) *DescribeRiskEventGroupResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetDirection(v string) *DescribeRiskEventGroupResponseBodyDataList {
	s.Direction = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetDstIP(v string) *DescribeRiskEventGroupResponseBodyDataList {
	s.DstIP = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetEventCount(v int32) *DescribeRiskEventGroupResponseBodyDataList {
	s.EventCount = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetEventId(v string) *DescribeRiskEventGroupResponseBodyDataList {
	s.EventId = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetEventName(v string) *DescribeRiskEventGroupResponseBodyDataList {
	s.EventName = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetFirstEventTime(v int32) *DescribeRiskEventGroupResponseBodyDataList {
	s.FirstEventTime = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetIPLocationInfo(v *DescribeRiskEventGroupResponseBodyDataListIPLocationInfo) *DescribeRiskEventGroupResponseBodyDataList {
	s.IPLocationInfo = v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetLastEventTime(v int32) *DescribeRiskEventGroupResponseBodyDataList {
	s.LastEventTime = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetResourcePrivateIPList(v []*DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList) *DescribeRiskEventGroupResponseBodyDataList {
	s.ResourcePrivateIPList = v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetResourceType(v string) *DescribeRiskEventGroupResponseBodyDataList {
	s.ResourceType = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetRuleId(v string) *DescribeRiskEventGroupResponseBodyDataList {
	s.RuleId = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetRuleResult(v int32) *DescribeRiskEventGroupResponseBodyDataList {
	s.RuleResult = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetRuleSource(v int32) *DescribeRiskEventGroupResponseBodyDataList {
	s.RuleSource = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetSrcIP(v string) *DescribeRiskEventGroupResponseBodyDataList {
	s.SrcIP = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetSrcIPTag(v string) *DescribeRiskEventGroupResponseBodyDataList {
	s.SrcIPTag = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetSrcPrivateIPList(v []*string) *DescribeRiskEventGroupResponseBodyDataList {
	s.SrcPrivateIPList = v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetTag(v string) *DescribeRiskEventGroupResponseBodyDataList {
	s.Tag = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetVpcDstInfo(v *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo) *DescribeRiskEventGroupResponseBodyDataList {
	s.VpcDstInfo = v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetVpcSrcInfo(v *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo) *DescribeRiskEventGroupResponseBodyDataList {
	s.VpcSrcInfo = v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataList) SetVulLevel(v int32) *DescribeRiskEventGroupResponseBodyDataList {
	s.VulLevel = &v
	return s
}

type DescribeRiskEventGroupResponseBodyDataListIPLocationInfo struct {
	// The ID of the city to which the IP address belongs.
	CityId *string `json:"CityId,omitempty" xml:"CityId,omitempty"`
	// The name of the city to which the IP address belongs.
	CityName *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	// The ID of the country to which the IP address belongs.
	CountryId *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
	// The name of the country to which the IP address belongs.
	CountryName *string `json:"CountryName,omitempty" xml:"CountryName,omitempty"`
}

func (s DescribeRiskEventGroupResponseBodyDataListIPLocationInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskEventGroupResponseBodyDataListIPLocationInfo) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventGroupResponseBodyDataListIPLocationInfo) SetCityId(v string) *DescribeRiskEventGroupResponseBodyDataListIPLocationInfo {
	s.CityId = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListIPLocationInfo) SetCityName(v string) *DescribeRiskEventGroupResponseBodyDataListIPLocationInfo {
	s.CityName = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListIPLocationInfo) SetCountryId(v string) *DescribeRiskEventGroupResponseBodyDataListIPLocationInfo {
	s.CountryId = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListIPLocationInfo) SetCountryName(v string) *DescribeRiskEventGroupResponseBodyDataListIPLocationInfo {
	s.CountryName = &v
	return s
}

type DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList struct {
	// The ID of the region to which the private IP address belongs.
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The ID of the instance that uses the private IP address.
	ResourceInstanceId *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
	// The name of the instance that uses the private IP address.
	ResourceInstanceName *string `json:"ResourceInstanceName,omitempty" xml:"ResourceInstanceName,omitempty"`
	// The private IP address.
	ResourcePrivateIP *string `json:"ResourcePrivateIP,omitempty" xml:"ResourcePrivateIP,omitempty"`
}

func (s DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList) SetRegionNo(v string) *DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList {
	s.RegionNo = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList) SetResourceInstanceId(v string) *DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList {
	s.ResourceInstanceId = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList) SetResourceInstanceName(v string) *DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList {
	s.ResourceInstanceName = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList) SetResourcePrivateIP(v string) *DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList {
	s.ResourcePrivateIP = &v
	return s
}

type DescribeRiskEventGroupResponseBodyDataListVpcDstInfo struct {
	// The ID of the ECS instance.
	EcsInstanceId *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	// The name of the ECS instance.
	EcsInstanceName *string `json:"EcsInstanceName,omitempty" xml:"EcsInstanceName,omitempty"`
	// The ID of the VPC.
	NetworkInstanceId *string `json:"NetworkInstanceId,omitempty" xml:"NetworkInstanceId,omitempty"`
	// The name of the VPC.
	NetworkInstanceName *string `json:"NetworkInstanceName,omitempty" xml:"NetworkInstanceName,omitempty"`
	// The ID of the region in which the destination VPC resides.
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
}

func (s DescribeRiskEventGroupResponseBodyDataListVpcDstInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskEventGroupResponseBodyDataListVpcDstInfo) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo) SetEcsInstanceId(v string) *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo {
	s.EcsInstanceId = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo) SetEcsInstanceName(v string) *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo {
	s.EcsInstanceName = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo) SetNetworkInstanceId(v string) *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo {
	s.NetworkInstanceId = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo) SetNetworkInstanceName(v string) *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo {
	s.NetworkInstanceName = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo) SetRegionNo(v string) *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo {
	s.RegionNo = &v
	return s
}

type DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo struct {
	// The ID of the ECS instance.
	EcsInstanceId *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	// The name of the ECS instance.
	EcsInstanceName *string `json:"EcsInstanceName,omitempty" xml:"EcsInstanceName,omitempty"`
	// The ID of the VPC.
	NetworkInstanceId *string `json:"NetworkInstanceId,omitempty" xml:"NetworkInstanceId,omitempty"`
	// The name of the VPC.
	NetworkInstanceName *string `json:"NetworkInstanceName,omitempty" xml:"NetworkInstanceName,omitempty"`
	// The ID of the region in which the source VPC resides.
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
}

func (s DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo) SetEcsInstanceId(v string) *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo {
	s.EcsInstanceId = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo) SetEcsInstanceName(v string) *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo {
	s.EcsInstanceName = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo) SetNetworkInstanceId(v string) *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo {
	s.NetworkInstanceId = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo) SetNetworkInstanceName(v string) *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo {
	s.NetworkInstanceName = &v
	return s
}

func (s *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo) SetRegionNo(v string) *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo {
	s.RegionNo = &v
	return s
}

type DescribeRiskEventGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRiskEventGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRiskEventGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskEventGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventGroupResponse) SetHeaders(v map[string]*string) *DescribeRiskEventGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeRiskEventGroupResponse) SetStatusCode(v int32) *DescribeRiskEventGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRiskEventGroupResponse) SetBody(v *DescribeRiskEventGroupResponseBody) *DescribeRiskEventGroupResponse {
	s.Body = v
	return s
}

type DescribeRiskEventPayloadRequest struct {
	// The destination IP address to query. If you specify this parameter, all intrusion events with the specified destination IP address are queried.
	DstIP *string `json:"DstIP,omitempty" xml:"DstIP,omitempty"`
	// The ID of the destination VPC to query. If you specify this parameter, all intrusion events that contain the specified ID of the destination VPC are queried.
	DstVpcId *string `json:"DstVpcId,omitempty" xml:"DstVpcId,omitempty"`
	// The end of the time range to query. The value is a timestamp. Unit: seconds.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The type of the firewall. Valid values:
	//
	// *   **VpcFirewall**: virtual private cloud (VPC) firewall
	// *   **InternetFirewall** (default): Internet firewall
	FirewallType *string `json:"FirewallType,omitempty" xml:"FirewallType,omitempty"`
	// The public IP address. If you specify this parameter, all intrusion events that contain the specified public IP address are queried.
	PublicIP *string `json:"PublicIP,omitempty" xml:"PublicIP,omitempty"`
	// The source IP address to query. If you specify this parameter, all intrusion events from the specified source IP address are queried.
	SrcIP *string `json:"SrcIP,omitempty" xml:"SrcIP,omitempty"`
	// The ID of the source VPC to query. If you specify this parameter, all intrusion events that contain the specified ID of the source VPC are queried.
	SrcVpcId *string `json:"SrcVpcId,omitempty" xml:"SrcVpcId,omitempty"`
	// The beginning of the time range to query. The value is a timestamp. Unit: seconds.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The UUID of the intrusion event.
	UUID *string `json:"UUID,omitempty" xml:"UUID,omitempty"`
}

func (s DescribeRiskEventPayloadRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskEventPayloadRequest) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventPayloadRequest) SetDstIP(v string) *DescribeRiskEventPayloadRequest {
	s.DstIP = &v
	return s
}

func (s *DescribeRiskEventPayloadRequest) SetDstVpcId(v string) *DescribeRiskEventPayloadRequest {
	s.DstVpcId = &v
	return s
}

func (s *DescribeRiskEventPayloadRequest) SetEndTime(v string) *DescribeRiskEventPayloadRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRiskEventPayloadRequest) SetFirewallType(v string) *DescribeRiskEventPayloadRequest {
	s.FirewallType = &v
	return s
}

func (s *DescribeRiskEventPayloadRequest) SetPublicIP(v string) *DescribeRiskEventPayloadRequest {
	s.PublicIP = &v
	return s
}

func (s *DescribeRiskEventPayloadRequest) SetSrcIP(v string) *DescribeRiskEventPayloadRequest {
	s.SrcIP = &v
	return s
}

func (s *DescribeRiskEventPayloadRequest) SetSrcVpcId(v string) *DescribeRiskEventPayloadRequest {
	s.SrcVpcId = &v
	return s
}

func (s *DescribeRiskEventPayloadRequest) SetStartTime(v string) *DescribeRiskEventPayloadRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRiskEventPayloadRequest) SetUUID(v string) *DescribeRiskEventPayloadRequest {
	s.UUID = &v
	return s
}

type DescribeRiskEventPayloadResponseBody struct {
	// The destination IP address of the intrusion event.
	DstIP *string `json:"DstIP,omitempty" xml:"DstIP,omitempty"`
	// The destination port of the intrusion event.
	DstPort *int32 `json:"DstPort,omitempty" xml:"DstPort,omitempty"`
	// The destination VPC ID of the intrusion event.
	DstVpcId *string `json:"DstVpcId,omitempty" xml:"DstVpcId,omitempty"`
	// The attack payload of the intrusion event.
	Payload *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	// The length of the attack payload of the intrusion event.
	PayloadLen *int32 `json:"PayloadLen,omitempty" xml:"PayloadLen,omitempty"`
	// The protocol type of intrusion event. Valid values:
	//
	// *   **TCP**
	// *   **UDP**
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// The HTTP X-Real-IP field.
	RealIp *string `json:"RealIp,omitempty" xml:"RealIp,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The source IP address of the intrusion event.
	SrcIP *string `json:"SrcIP,omitempty" xml:"SrcIP,omitempty"`
	// The source port of the intrusion event.
	SrcPort *int32 `json:"SrcPort,omitempty" xml:"SrcPort,omitempty"`
	// The source VPC ID of the intrusion event.
	SrcVpcId *string `json:"SrcVpcId,omitempty" xml:"SrcVpcId,omitempty"`
	// The HTTP X-Forwarded-For field.
	XForwardFor *string `json:"XForwardFor,omitempty" xml:"XForwardFor,omitempty"`
}

func (s DescribeRiskEventPayloadResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskEventPayloadResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventPayloadResponseBody) SetDstIP(v string) *DescribeRiskEventPayloadResponseBody {
	s.DstIP = &v
	return s
}

func (s *DescribeRiskEventPayloadResponseBody) SetDstPort(v int32) *DescribeRiskEventPayloadResponseBody {
	s.DstPort = &v
	return s
}

func (s *DescribeRiskEventPayloadResponseBody) SetDstVpcId(v string) *DescribeRiskEventPayloadResponseBody {
	s.DstVpcId = &v
	return s
}

func (s *DescribeRiskEventPayloadResponseBody) SetPayload(v string) *DescribeRiskEventPayloadResponseBody {
	s.Payload = &v
	return s
}

func (s *DescribeRiskEventPayloadResponseBody) SetPayloadLen(v int32) *DescribeRiskEventPayloadResponseBody {
	s.PayloadLen = &v
	return s
}

func (s *DescribeRiskEventPayloadResponseBody) SetProto(v string) *DescribeRiskEventPayloadResponseBody {
	s.Proto = &v
	return s
}

func (s *DescribeRiskEventPayloadResponseBody) SetRealIp(v string) *DescribeRiskEventPayloadResponseBody {
	s.RealIp = &v
	return s
}

func (s *DescribeRiskEventPayloadResponseBody) SetRequestId(v string) *DescribeRiskEventPayloadResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRiskEventPayloadResponseBody) SetSrcIP(v string) *DescribeRiskEventPayloadResponseBody {
	s.SrcIP = &v
	return s
}

func (s *DescribeRiskEventPayloadResponseBody) SetSrcPort(v int32) *DescribeRiskEventPayloadResponseBody {
	s.SrcPort = &v
	return s
}

func (s *DescribeRiskEventPayloadResponseBody) SetSrcVpcId(v string) *DescribeRiskEventPayloadResponseBody {
	s.SrcVpcId = &v
	return s
}

func (s *DescribeRiskEventPayloadResponseBody) SetXForwardFor(v string) *DescribeRiskEventPayloadResponseBody {
	s.XForwardFor = &v
	return s
}

type DescribeRiskEventPayloadResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRiskEventPayloadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRiskEventPayloadResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRiskEventPayloadResponse) GoString() string {
	return s.String()
}

func (s *DescribeRiskEventPayloadResponse) SetHeaders(v map[string]*string) *DescribeRiskEventPayloadResponse {
	s.Headers = v
	return s
}

func (s *DescribeRiskEventPayloadResponse) SetStatusCode(v int32) *DescribeRiskEventPayloadResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRiskEventPayloadResponse) SetBody(v *DescribeRiskEventPayloadResponseBody) *DescribeRiskEventPayloadResponse {
	s.Body = v
	return s
}

type DescribeSignatureLibVersionResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The information about the versions.
	Version []*DescribeSignatureLibVersionResponseBodyVersion `json:"Version,omitempty" xml:"Version,omitempty" type:"Repeated"`
}

func (s DescribeSignatureLibVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSignatureLibVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSignatureLibVersionResponseBody) SetRequestId(v string) *DescribeSignatureLibVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSignatureLibVersionResponseBody) SetTotalCount(v int32) *DescribeSignatureLibVersionResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSignatureLibVersionResponseBody) SetVersion(v []*DescribeSignatureLibVersionResponseBodyVersion) *DescribeSignatureLibVersionResponseBody {
	s.Version = v
	return s
}

type DescribeSignatureLibVersionResponseBodyVersion struct {
	// The type.
	//
	// Valid values:
	//
	// *   ips
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     Basic Rules and Virtual Patching
	//
	//     <!-- -->
	//
	//     .
	//
	// *   intelligence
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     Threat Intelligence
	//
	//     <!-- -->
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The version number.
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeSignatureLibVersionResponseBodyVersion) String() string {
	return tea.Prettify(s)
}

func (s DescribeSignatureLibVersionResponseBodyVersion) GoString() string {
	return s.String()
}

func (s *DescribeSignatureLibVersionResponseBodyVersion) SetType(v string) *DescribeSignatureLibVersionResponseBodyVersion {
	s.Type = &v
	return s
}

func (s *DescribeSignatureLibVersionResponseBodyVersion) SetVersion(v string) *DescribeSignatureLibVersionResponseBodyVersion {
	s.Version = &v
	return s
}

type DescribeSignatureLibVersionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSignatureLibVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSignatureLibVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSignatureLibVersionResponse) GoString() string {
	return s.String()
}

func (s *DescribeSignatureLibVersionResponse) SetHeaders(v map[string]*string) *DescribeSignatureLibVersionResponse {
	s.Headers = v
	return s
}

func (s *DescribeSignatureLibVersionResponse) SetStatusCode(v int32) *DescribeSignatureLibVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSignatureLibVersionResponse) SetBody(v *DescribeSignatureLibVersionResponseBody) *DescribeSignatureLibVersionResponse {
	s.Body = v
	return s
}

type DescribeTrFirewallPolicyBackUpAssociationListRequest struct {
	// An array that consists of the details about the traffic redirection instance.
	CandidateList []*DescribeTrFirewallPolicyBackUpAssociationListRequestCandidateList `json:"CandidateList,omitempty" xml:"CandidateList,omitempty" type:"Repeated"`
	// The instance ID of the VPC firewall.
	FirewallId *string `json:"FirewallId,omitempty" xml:"FirewallId,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// *   **zh** (default): Chinese
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the routing policy.
	TrFirewallRoutePolicyId *string `json:"TrFirewallRoutePolicyId,omitempty" xml:"TrFirewallRoutePolicyId,omitempty"`
}

func (s DescribeTrFirewallPolicyBackUpAssociationListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrFirewallPolicyBackUpAssociationListRequest) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListRequest) SetCandidateList(v []*DescribeTrFirewallPolicyBackUpAssociationListRequestCandidateList) *DescribeTrFirewallPolicyBackUpAssociationListRequest {
	s.CandidateList = v
	return s
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListRequest) SetFirewallId(v string) *DescribeTrFirewallPolicyBackUpAssociationListRequest {
	s.FirewallId = &v
	return s
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListRequest) SetLang(v string) *DescribeTrFirewallPolicyBackUpAssociationListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListRequest) SetTrFirewallRoutePolicyId(v string) *DescribeTrFirewallPolicyBackUpAssociationListRequest {
	s.TrFirewallRoutePolicyId = &v
	return s
}

type DescribeTrFirewallPolicyBackUpAssociationListRequestCandidateList struct {
	// The ID of the traffic redirection instance.
	CandidateId *string `json:"CandidateId,omitempty" xml:"CandidateId,omitempty"`
	// The type of the traffic redirection instance.
	CandidateType *string `json:"CandidateType,omitempty" xml:"CandidateType,omitempty"`
}

func (s DescribeTrFirewallPolicyBackUpAssociationListRequestCandidateList) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrFirewallPolicyBackUpAssociationListRequestCandidateList) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListRequestCandidateList) SetCandidateId(v string) *DescribeTrFirewallPolicyBackUpAssociationListRequestCandidateList {
	s.CandidateId = &v
	return s
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListRequestCandidateList) SetCandidateType(v string) *DescribeTrFirewallPolicyBackUpAssociationListRequestCandidateList {
	s.CandidateType = &v
	return s
}

type DescribeTrFirewallPolicyBackUpAssociationListShrinkRequest struct {
	// An array that consists of the details about the traffic redirection instance.
	CandidateListShrink *string `json:"CandidateList,omitempty" xml:"CandidateList,omitempty"`
	// The instance ID of the VPC firewall.
	FirewallId *string `json:"FirewallId,omitempty" xml:"FirewallId,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// *   **zh** (default): Chinese
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the routing policy.
	TrFirewallRoutePolicyId *string `json:"TrFirewallRoutePolicyId,omitempty" xml:"TrFirewallRoutePolicyId,omitempty"`
}

func (s DescribeTrFirewallPolicyBackUpAssociationListShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrFirewallPolicyBackUpAssociationListShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListShrinkRequest) SetCandidateListShrink(v string) *DescribeTrFirewallPolicyBackUpAssociationListShrinkRequest {
	s.CandidateListShrink = &v
	return s
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListShrinkRequest) SetFirewallId(v string) *DescribeTrFirewallPolicyBackUpAssociationListShrinkRequest {
	s.FirewallId = &v
	return s
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListShrinkRequest) SetLang(v string) *DescribeTrFirewallPolicyBackUpAssociationListShrinkRequest {
	s.Lang = &v
	return s
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListShrinkRequest) SetTrFirewallRoutePolicyId(v string) *DescribeTrFirewallPolicyBackUpAssociationListShrinkRequest {
	s.TrFirewallRoutePolicyId = &v
	return s
}

type DescribeTrFirewallPolicyBackUpAssociationListResponseBody struct {
	// The route tables.
	PolicyAssociationBackupConfigs []*DescribeTrFirewallPolicyBackUpAssociationListResponseBodyPolicyAssociationBackupConfigs `json:"PolicyAssociationBackupConfigs,omitempty" xml:"PolicyAssociationBackupConfigs,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTrFirewallPolicyBackUpAssociationListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrFirewallPolicyBackUpAssociationListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListResponseBody) SetPolicyAssociationBackupConfigs(v []*DescribeTrFirewallPolicyBackUpAssociationListResponseBodyPolicyAssociationBackupConfigs) *DescribeTrFirewallPolicyBackUpAssociationListResponseBody {
	s.PolicyAssociationBackupConfigs = v
	return s
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListResponseBody) SetRequestId(v string) *DescribeTrFirewallPolicyBackUpAssociationListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeTrFirewallPolicyBackUpAssociationListResponseBodyPolicyAssociationBackupConfigs struct {
	// The ID of the traffic redirection instance.
	CandidateId *string `json:"CandidateId,omitempty" xml:"CandidateId,omitempty"`
	// The name of the traffic redirection instance.
	CandidateName *string `json:"CandidateName,omitempty" xml:"CandidateName,omitempty"`
	// The type of the traffic redirection instance.
	CandidateType *string `json:"CandidateType,omitempty" xml:"CandidateType,omitempty"`
	// The route table that is used after traffic redirection.
	CurrentRouteTableId *string `json:"CurrentRouteTableId,omitempty" xml:"CurrentRouteTableId,omitempty"`
	// The ID of the route table.
	OriginalRouteTableId *string `json:"OriginalRouteTableId,omitempty" xml:"OriginalRouteTableId,omitempty"`
}

func (s DescribeTrFirewallPolicyBackUpAssociationListResponseBodyPolicyAssociationBackupConfigs) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrFirewallPolicyBackUpAssociationListResponseBodyPolicyAssociationBackupConfigs) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListResponseBodyPolicyAssociationBackupConfigs) SetCandidateId(v string) *DescribeTrFirewallPolicyBackUpAssociationListResponseBodyPolicyAssociationBackupConfigs {
	s.CandidateId = &v
	return s
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListResponseBodyPolicyAssociationBackupConfigs) SetCandidateName(v string) *DescribeTrFirewallPolicyBackUpAssociationListResponseBodyPolicyAssociationBackupConfigs {
	s.CandidateName = &v
	return s
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListResponseBodyPolicyAssociationBackupConfigs) SetCandidateType(v string) *DescribeTrFirewallPolicyBackUpAssociationListResponseBodyPolicyAssociationBackupConfigs {
	s.CandidateType = &v
	return s
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListResponseBodyPolicyAssociationBackupConfigs) SetCurrentRouteTableId(v string) *DescribeTrFirewallPolicyBackUpAssociationListResponseBodyPolicyAssociationBackupConfigs {
	s.CurrentRouteTableId = &v
	return s
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListResponseBodyPolicyAssociationBackupConfigs) SetOriginalRouteTableId(v string) *DescribeTrFirewallPolicyBackUpAssociationListResponseBodyPolicyAssociationBackupConfigs {
	s.OriginalRouteTableId = &v
	return s
}

type DescribeTrFirewallPolicyBackUpAssociationListResponse struct {
	Headers    map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTrFirewallPolicyBackUpAssociationListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTrFirewallPolicyBackUpAssociationListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrFirewallPolicyBackUpAssociationListResponse) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListResponse) SetHeaders(v map[string]*string) *DescribeTrFirewallPolicyBackUpAssociationListResponse {
	s.Headers = v
	return s
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListResponse) SetStatusCode(v int32) *DescribeTrFirewallPolicyBackUpAssociationListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTrFirewallPolicyBackUpAssociationListResponse) SetBody(v *DescribeTrFirewallPolicyBackUpAssociationListResponseBody) *DescribeTrFirewallPolicyBackUpAssociationListResponse {
	s.Body = v
	return s
}

type DescribeTrFirewallV2RoutePolicyListRequest struct {
	// The page number. Default value: 1.
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The instance ID of the VPC firewall.
	FirewallId *string `json:"FirewallId,omitempty" xml:"FirewallId,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries per page. Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the routing policy.
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s DescribeTrFirewallV2RoutePolicyListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrFirewallV2RoutePolicyListRequest) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallV2RoutePolicyListRequest) SetCurrentPage(v int32) *DescribeTrFirewallV2RoutePolicyListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeTrFirewallV2RoutePolicyListRequest) SetFirewallId(v string) *DescribeTrFirewallV2RoutePolicyListRequest {
	s.FirewallId = &v
	return s
}

func (s *DescribeTrFirewallV2RoutePolicyListRequest) SetLang(v string) *DescribeTrFirewallV2RoutePolicyListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeTrFirewallV2RoutePolicyListRequest) SetPageSize(v int32) *DescribeTrFirewallV2RoutePolicyListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTrFirewallV2RoutePolicyListRequest) SetPolicyId(v string) *DescribeTrFirewallV2RoutePolicyListRequest {
	s.PolicyId = &v
	return s
}

type DescribeTrFirewallV2RoutePolicyListResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The routing policies.
	TrFirewallRoutePolicies []*DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies `json:"TrFirewallRoutePolicies,omitempty" xml:"TrFirewallRoutePolicies,omitempty" type:"Repeated"`
}

func (s DescribeTrFirewallV2RoutePolicyListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrFirewallV2RoutePolicyListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBody) SetRequestId(v string) *DescribeTrFirewallV2RoutePolicyListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBody) SetTotalCount(v string) *DescribeTrFirewallV2RoutePolicyListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBody) SetTrFirewallRoutePolicies(v []*DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies) *DescribeTrFirewallV2RoutePolicyListResponseBody {
	s.TrFirewallRoutePolicies = v
	return s
}

type DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies struct {
	// The secondary traffic redirection instances.
	DestCandidateList []*DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesDestCandidateList `json:"DestCandidateList,omitempty" xml:"DestCandidateList,omitempty" type:"Repeated"`
	// The description of the routing policy.
	PolicyDescription *string `json:"PolicyDescription,omitempty" xml:"PolicyDescription,omitempty"`
	// The name of the routing policy.
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The status of the routing policy. Valid values:
	//
	// *   creating: The policy is being created.
	// *   deleting: The policy is being deleted.
	// *   opening: The policy is being enabled.
	// *   opened: The policy is enabled.
	// *   closing: The policy is being disabled.
	// *   closed: The policy is disabled.
	PolicyStatus *string `json:"PolicyStatus,omitempty" xml:"PolicyStatus,omitempty"`
	// The type of the traffic redirection scenario of the VPC firewall. Valid values:
	//
	// *   **fullmesh**: interconnected instances
	// *   **one_to_one**: instance to instance
	// *   **end_to_end**: instance to instances
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The primary traffic redirection instances.
	SrcCandidateList []*DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesSrcCandidateList `json:"SrcCandidateList,omitempty" xml:"SrcCandidateList,omitempty" type:"Repeated"`
	// The ID of the routing policy.
	TrFirewallRoutePolicyId *string `json:"TrFirewallRoutePolicyId,omitempty" xml:"TrFirewallRoutePolicyId,omitempty"`
}

func (s DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies) SetDestCandidateList(v []*DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesDestCandidateList) *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies {
	s.DestCandidateList = v
	return s
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies) SetPolicyDescription(v string) *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies {
	s.PolicyDescription = &v
	return s
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies) SetPolicyName(v string) *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies {
	s.PolicyName = &v
	return s
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies) SetPolicyStatus(v string) *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies {
	s.PolicyStatus = &v
	return s
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies) SetPolicyType(v string) *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies {
	s.PolicyType = &v
	return s
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies) SetSrcCandidateList(v []*DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesSrcCandidateList) *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies {
	s.SrcCandidateList = v
	return s
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies) SetTrFirewallRoutePolicyId(v string) *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePolicies {
	s.TrFirewallRoutePolicyId = &v
	return s
}

type DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesDestCandidateList struct {
	// The ID of the secondary traffic redirection instance.
	CandidateId *string `json:"CandidateId,omitempty" xml:"CandidateId,omitempty"`
	// The type of the secondary traffic redirection instance.
	CandidateType *string `json:"CandidateType,omitempty" xml:"CandidateType,omitempty"`
}

func (s DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesDestCandidateList) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesDestCandidateList) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesDestCandidateList) SetCandidateId(v string) *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesDestCandidateList {
	s.CandidateId = &v
	return s
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesDestCandidateList) SetCandidateType(v string) *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesDestCandidateList {
	s.CandidateType = &v
	return s
}

type DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesSrcCandidateList struct {
	// The ID of the primary traffic redirection instance.
	CandidateId *string `json:"CandidateId,omitempty" xml:"CandidateId,omitempty"`
	// The type of the primary traffic redirection instance.
	CandidateType *string `json:"CandidateType,omitempty" xml:"CandidateType,omitempty"`
}

func (s DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesSrcCandidateList) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesSrcCandidateList) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesSrcCandidateList) SetCandidateId(v string) *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesSrcCandidateList {
	s.CandidateId = &v
	return s
}

func (s *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesSrcCandidateList) SetCandidateType(v string) *DescribeTrFirewallV2RoutePolicyListResponseBodyTrFirewallRoutePoliciesSrcCandidateList {
	s.CandidateType = &v
	return s
}

type DescribeTrFirewallV2RoutePolicyListResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTrFirewallV2RoutePolicyListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTrFirewallV2RoutePolicyListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrFirewallV2RoutePolicyListResponse) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallV2RoutePolicyListResponse) SetHeaders(v map[string]*string) *DescribeTrFirewallV2RoutePolicyListResponse {
	s.Headers = v
	return s
}

func (s *DescribeTrFirewallV2RoutePolicyListResponse) SetStatusCode(v int32) *DescribeTrFirewallV2RoutePolicyListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTrFirewallV2RoutePolicyListResponse) SetBody(v *DescribeTrFirewallV2RoutePolicyListResponseBody) *DescribeTrFirewallV2RoutePolicyListResponse {
	s.Body = v
	return s
}

type DescribeTrFirewallsV2DetailRequest struct {
	// The instance ID of the VPC firewall.
	FirewallId *string `json:"FirewallId,omitempty" xml:"FirewallId,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeTrFirewallsV2DetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrFirewallsV2DetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallsV2DetailRequest) SetFirewallId(v string) *DescribeTrFirewallsV2DetailRequest {
	s.FirewallId = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailRequest) SetLang(v string) *DescribeTrFirewallsV2DetailRequest {
	s.Lang = &v
	return s
}

type DescribeTrFirewallsV2DetailResponseBody struct {
	// The ID of the Cloud Enterprise Network (CEN) instance.
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The description of the VPC firewall.
	FirewallDescription *string `json:"FirewallDescription,omitempty" xml:"FirewallDescription,omitempty"`
	// The ID of the Elastic Network Interface (ENI) with which the VPC firewall is associated.
	FirewallEniId *string `json:"FirewallEniId,omitempty" xml:"FirewallEniId,omitempty"`
	// The ID of the VPC to which the ENI is attached.
	FirewallEniVpcId *string `json:"FirewallEniVpcId,omitempty" xml:"FirewallEniVpcId,omitempty"`
	// The ID of the vSwitch with which the ENI is associated.
	FirewallEniVswitchId *string `json:"FirewallEniVswitchId,omitempty" xml:"FirewallEniVswitchId,omitempty"`
	// The instance ID of the VPC firewall.
	FirewallId *string `json:"FirewallId,omitempty" xml:"FirewallId,omitempty"`
	// The name of the VPC firewall.
	FirewallName *string `json:"FirewallName,omitempty" xml:"FirewallName,omitempty"`
	// The status of the VPC firewall. Valid values:
	//
	// *   Creating
	// *   Deleting
	// *   Ready
	FirewallStatus *string `json:"FirewallStatus,omitempty" xml:"FirewallStatus,omitempty"`
	// The subnet CIDR block of the VPC in which the ENI of the firewall is stored in automatic mode.
	FirewallSubnetCidr *string `json:"FirewallSubnetCidr,omitempty" xml:"FirewallSubnetCidr,omitempty"`
	// The status of the VPC firewall. Valid values:
	//
	// *   **opened**: The VPC firewall is enabled.
	// *   **closed**: The VPC firewall is disabled.
	// *   **notconfigured**: The VPC firewall is not created.
	// *   **configured**: The VPC firewall is created but is not enabled.
	// *   **creating**: The VPC firewall is being created.
	// *   **opening**: The VPC firewall is being enabled.
	// *   **deleting**: The VPC firewall is being deleted.
	//
	// > If you do not specify this parameter, VPC firewalls in all states are queried.
	FirewallSwitchStatus *string `json:"FirewallSwitchStatus,omitempty" xml:"FirewallSwitchStatus,omitempty"`
	// The CIDR block that is allocated to the VPC created for the VPC firewall in automatic mode.
	FirewallVpcCidr *string `json:"FirewallVpcCidr,omitempty" xml:"FirewallVpcCidr,omitempty"`
	// The region ID of the transit router.
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The routing mode of the VPC firewall. Valid values:
	//
	// *   **managed**: automatic mode
	// *   **manual**: manual mode
	RouteMode *string `json:"RouteMode,omitempty" xml:"RouteMode,omitempty"`
	// The primary subnet CIDR block that the VPC uses to connect to the transit router in automatic mode.
	TrAttachmentMasterCidr *string `json:"TrAttachmentMasterCidr,omitempty" xml:"TrAttachmentMasterCidr,omitempty"`
	// The secondary subnet CIDR block that the VPC uses to connect to the transit router in automatic mode.
	TrAttachmentSlaveCidr *string `json:"TrAttachmentSlaveCidr,omitempty" xml:"TrAttachmentSlaveCidr,omitempty"`
	// The ID of the transit router.
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s DescribeTrFirewallsV2DetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrFirewallsV2DetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallsV2DetailResponseBody) SetCenId(v string) *DescribeTrFirewallsV2DetailResponseBody {
	s.CenId = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponseBody) SetFirewallDescription(v string) *DescribeTrFirewallsV2DetailResponseBody {
	s.FirewallDescription = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponseBody) SetFirewallEniId(v string) *DescribeTrFirewallsV2DetailResponseBody {
	s.FirewallEniId = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponseBody) SetFirewallEniVpcId(v string) *DescribeTrFirewallsV2DetailResponseBody {
	s.FirewallEniVpcId = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponseBody) SetFirewallEniVswitchId(v string) *DescribeTrFirewallsV2DetailResponseBody {
	s.FirewallEniVswitchId = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponseBody) SetFirewallId(v string) *DescribeTrFirewallsV2DetailResponseBody {
	s.FirewallId = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponseBody) SetFirewallName(v string) *DescribeTrFirewallsV2DetailResponseBody {
	s.FirewallName = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponseBody) SetFirewallStatus(v string) *DescribeTrFirewallsV2DetailResponseBody {
	s.FirewallStatus = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponseBody) SetFirewallSubnetCidr(v string) *DescribeTrFirewallsV2DetailResponseBody {
	s.FirewallSubnetCidr = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponseBody) SetFirewallSwitchStatus(v string) *DescribeTrFirewallsV2DetailResponseBody {
	s.FirewallSwitchStatus = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponseBody) SetFirewallVpcCidr(v string) *DescribeTrFirewallsV2DetailResponseBody {
	s.FirewallVpcCidr = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponseBody) SetRegionNo(v string) *DescribeTrFirewallsV2DetailResponseBody {
	s.RegionNo = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponseBody) SetRequestId(v string) *DescribeTrFirewallsV2DetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponseBody) SetRouteMode(v string) *DescribeTrFirewallsV2DetailResponseBody {
	s.RouteMode = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponseBody) SetTrAttachmentMasterCidr(v string) *DescribeTrFirewallsV2DetailResponseBody {
	s.TrAttachmentMasterCidr = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponseBody) SetTrAttachmentSlaveCidr(v string) *DescribeTrFirewallsV2DetailResponseBody {
	s.TrAttachmentSlaveCidr = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponseBody) SetTransitRouterId(v string) *DescribeTrFirewallsV2DetailResponseBody {
	s.TransitRouterId = &v
	return s
}

type DescribeTrFirewallsV2DetailResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTrFirewallsV2DetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTrFirewallsV2DetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrFirewallsV2DetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallsV2DetailResponse) SetHeaders(v map[string]*string) *DescribeTrFirewallsV2DetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponse) SetStatusCode(v int32) *DescribeTrFirewallsV2DetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponse) SetBody(v *DescribeTrFirewallsV2DetailResponseBody) *DescribeTrFirewallsV2DetailResponse {
	s.Body = v
	return s
}

type DescribeTrFirewallsV2ListRequest struct {
	// The ID of the Cloud Enterprise Network (CEN) instance.
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The page number. Default value: **1**.
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The instance ID of the VPC firewall.
	FirewallId *string `json:"FirewallId,omitempty" xml:"FirewallId,omitempty"`
	// The name of the VPC firewall.
	FirewallName *string `json:"FirewallName,omitempty" xml:"FirewallName,omitempty"`
	// The status of the VPC firewall. Valid values:
	//
	// *   **opened**: The VPC firewall is enabled.
	// *   **closed**: The VPC firewall is disabled.
	// *   **notconfigured**: The VPC firewall is not created.
	// *   **configured**: The VPC firewall is created but is not enabled.
	// *   **creating**: The VPC firewall is being created.
	// *   **opening**: The VPC firewall is being enabled.
	// *   **deleting**: The VPC firewall is being deleted.
	//
	// >  If you do not specify this parameter, VPC firewalls in all states are queried.
	FirewallSwitchStatus *string `json:"FirewallSwitchStatus,omitempty" xml:"FirewallSwitchStatus,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang    *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of entries per page. Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the transit router.
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The routing mode of the VPC firewall. Valid values:
	//
	// *   **managed**: automatic mode
	// *   **manual**: manual mode
	//
	// >  If you do not specify this parameter, VPC firewalls in all routing modes are queried.
	RouteMode *string `json:"RouteMode,omitempty" xml:"RouteMode,omitempty"`
	// The ID of the transit router.
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s DescribeTrFirewallsV2ListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrFirewallsV2ListRequest) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallsV2ListRequest) SetCenId(v string) *DescribeTrFirewallsV2ListRequest {
	s.CenId = &v
	return s
}

func (s *DescribeTrFirewallsV2ListRequest) SetCurrentPage(v int32) *DescribeTrFirewallsV2ListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeTrFirewallsV2ListRequest) SetFirewallId(v string) *DescribeTrFirewallsV2ListRequest {
	s.FirewallId = &v
	return s
}

func (s *DescribeTrFirewallsV2ListRequest) SetFirewallName(v string) *DescribeTrFirewallsV2ListRequest {
	s.FirewallName = &v
	return s
}

func (s *DescribeTrFirewallsV2ListRequest) SetFirewallSwitchStatus(v string) *DescribeTrFirewallsV2ListRequest {
	s.FirewallSwitchStatus = &v
	return s
}

func (s *DescribeTrFirewallsV2ListRequest) SetLang(v string) *DescribeTrFirewallsV2ListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeTrFirewallsV2ListRequest) SetOwnerId(v string) *DescribeTrFirewallsV2ListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTrFirewallsV2ListRequest) SetPageSize(v int32) *DescribeTrFirewallsV2ListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTrFirewallsV2ListRequest) SetRegionNo(v string) *DescribeTrFirewallsV2ListRequest {
	s.RegionNo = &v
	return s
}

func (s *DescribeTrFirewallsV2ListRequest) SetRouteMode(v string) *DescribeTrFirewallsV2ListRequest {
	s.RouteMode = &v
	return s
}

func (s *DescribeTrFirewallsV2ListRequest) SetTransitRouterId(v string) *DescribeTrFirewallsV2ListRequest {
	s.TransitRouterId = &v
	return s
}

type DescribeTrFirewallsV2ListResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The VPC firewalls.
	VpcTrFirewalls []*DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls `json:"VpcTrFirewalls,omitempty" xml:"VpcTrFirewalls,omitempty" type:"Repeated"`
}

func (s DescribeTrFirewallsV2ListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrFirewallsV2ListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallsV2ListResponseBody) SetRequestId(v string) *DescribeTrFirewallsV2ListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBody) SetTotalCount(v string) *DescribeTrFirewallsV2ListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBody) SetVpcTrFirewalls(v []*DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) *DescribeTrFirewallsV2ListResponseBody {
	s.VpcTrFirewalls = v
	return s
}

type DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls struct {
	// The ID of the CEN instance.
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The name of the CEN instance.
	CenName *string `json:"CenName,omitempty" xml:"CenName,omitempty"`
	// The instance ID of the VPC firewall.
	FirewallId *string `json:"FirewallId,omitempty" xml:"FirewallId,omitempty"`
	// The status of the VPC firewall. Valid values:
	//
	// *   **opened**: The VPC firewall is enabled.
	// *   **closed**: The VPC firewall is disabled.
	// *   **notconfigured**: The VPC firewall is not created.
	// *   **configured**: The VPC firewall is created but is not enabled.
	// *   **creating**: The VPC firewall is being created.
	// *   **opening**: The VPC firewall is being enabled.
	// *   **deleting**: The VPC firewall is being deleted.
	//
	// >  If you do not specify this parameter, VPC firewalls in all states are queried.
	FirewallSwitchStatus *string `json:"FirewallSwitchStatus,omitempty" xml:"FirewallSwitchStatus,omitempty"`
	// The information about the intrusion prevention system (IPS) configuration.
	IpsConfig *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsIpsConfig `json:"IpsConfig,omitempty" xml:"IpsConfig,omitempty" type:"Struct"`
	// The ID of the Alibaba Cloud account to which the VPC belongs.
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Indicates whether the VPC firewall can be automatically enabled. Valid values:
	//
	// *   **passed**: yes
	// *   **failed**: no
	// *   **unknown**
	PrecheckStatus *string `json:"PrecheckStatus,omitempty" xml:"PrecheckStatus,omitempty"`
	// The protected resources.
	ProtectedResource *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsProtectedResource `json:"ProtectedResource,omitempty" xml:"ProtectedResource,omitempty" type:"Struct"`
	// The region ID of the transit router.
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// Indicates whether you can create a VPC firewall in a specified region. Valid values:
	//
	// *   **enable**: yes
	// *   **disable**: no
	RegionStatus *string `json:"RegionStatus,omitempty" xml:"RegionStatus,omitempty"`
	// The result code of the operation that creates the VPC firewall. Valid values:
	//
	// *   **RegionDisable**: VPC Firewall is not supported in the region of the network instance. You cannot create a VPC firewall for the network instance.
	// *   **Empty string**: You can create a VPC firewall for the network instance.
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// The routing mode of the VPC firewall. Valid values:
	//
	// *   **managed**: automatic mode
	// *   **manual**: manual mode
	RouteMode *string `json:"RouteMode,omitempty" xml:"RouteMode,omitempty"`
	// The ID of the transit router.
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// The unprotected resources.
	UnprotectedResource *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsUnprotectedResource `json:"UnprotectedResource,omitempty" xml:"UnprotectedResource,omitempty" type:"Struct"`
	// The instance name of the VPC firewall.
	VpcFirewallName *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
}

func (s DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) SetCenId(v string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls {
	s.CenId = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) SetCenName(v string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls {
	s.CenName = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) SetFirewallId(v string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls {
	s.FirewallId = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) SetFirewallSwitchStatus(v string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls {
	s.FirewallSwitchStatus = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) SetIpsConfig(v *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsIpsConfig) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls {
	s.IpsConfig = v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) SetOwnerId(v int64) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls {
	s.OwnerId = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) SetPrecheckStatus(v string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls {
	s.PrecheckStatus = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) SetProtectedResource(v *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsProtectedResource) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls {
	s.ProtectedResource = v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) SetRegionNo(v string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls {
	s.RegionNo = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) SetRegionStatus(v string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls {
	s.RegionStatus = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) SetResultCode(v string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls {
	s.ResultCode = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) SetRouteMode(v string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls {
	s.RouteMode = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) SetTransitRouterId(v string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls {
	s.TransitRouterId = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) SetUnprotectedResource(v *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsUnprotectedResource) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls {
	s.UnprotectedResource = v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) SetVpcFirewallName(v string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls {
	s.VpcFirewallName = &v
	return s
}

type DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsIpsConfig struct {
	// Indicates whether basic protection is enabled. Valid values:
	//
	// *   **1**: yes
	// *   **0**: no
	BasicRules *int32 `json:"BasicRules,omitempty" xml:"BasicRules,omitempty"`
	// Indicates whether virtual patching is enabled. Valid values:
	//
	// *   **1**: yes
	// *   **0**: no
	EnableAllPatch *int32 `json:"EnableAllPatch,omitempty" xml:"EnableAllPatch,omitempty"`
	// The mode of the IPS. Valid values:
	//
	// *   **1**: block mode
	// *   **0**: monitor mode
	RunMode *int32 `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
}

func (s DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsIpsConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsIpsConfig) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsIpsConfig) SetBasicRules(v int32) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsIpsConfig {
	s.BasicRules = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsIpsConfig) SetEnableAllPatch(v int32) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsIpsConfig {
	s.EnableAllPatch = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsIpsConfig) SetRunMode(v int32) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsIpsConfig {
	s.RunMode = &v
	return s
}

type DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsProtectedResource struct {
	// The number of protected resources.
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The protected peer transit routers.
	PeerTrList []*string `json:"PeerTrList,omitempty" xml:"PeerTrList,omitempty" type:"Repeated"`
	// The protected virtual border routers (VBRs).
	VbrList []*string `json:"VbrList,omitempty" xml:"VbrList,omitempty" type:"Repeated"`
	// The protected VPCs.
	VpcList []*string `json:"VpcList,omitempty" xml:"VpcList,omitempty" type:"Repeated"`
	// The protected VPN gateways.
	VpnList []*string `json:"VpnList,omitempty" xml:"VpnList,omitempty" type:"Repeated"`
}

func (s DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsProtectedResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsProtectedResource) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsProtectedResource) SetCount(v int32) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsProtectedResource {
	s.Count = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsProtectedResource) SetPeerTrList(v []*string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsProtectedResource {
	s.PeerTrList = v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsProtectedResource) SetVbrList(v []*string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsProtectedResource {
	s.VbrList = v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsProtectedResource) SetVpcList(v []*string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsProtectedResource {
	s.VpcList = v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsProtectedResource) SetVpnList(v []*string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsProtectedResource {
	s.VpnList = v
	return s
}

type DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsUnprotectedResource struct {
	// The number of unprotected resources.
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The unprotected peer transit routers.
	PeerTrList []*string `json:"PeerTrList,omitempty" xml:"PeerTrList,omitempty" type:"Repeated"`
	// The unprotected VBRs.
	VbrList []*string `json:"VbrList,omitempty" xml:"VbrList,omitempty" type:"Repeated"`
	// The unprotected VPCs.
	VpcList []*string `json:"VpcList,omitempty" xml:"VpcList,omitempty" type:"Repeated"`
	// The unprotected VPN gateways.
	VpnList []*string `json:"VpnList,omitempty" xml:"VpnList,omitempty" type:"Repeated"`
}

func (s DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsUnprotectedResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsUnprotectedResource) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsUnprotectedResource) SetCount(v int32) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsUnprotectedResource {
	s.Count = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsUnprotectedResource) SetPeerTrList(v []*string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsUnprotectedResource {
	s.PeerTrList = v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsUnprotectedResource) SetVbrList(v []*string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsUnprotectedResource {
	s.VbrList = v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsUnprotectedResource) SetVpcList(v []*string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsUnprotectedResource {
	s.VpcList = v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsUnprotectedResource) SetVpnList(v []*string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsUnprotectedResource {
	s.VpnList = v
	return s
}

type DescribeTrFirewallsV2ListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTrFirewallsV2ListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTrFirewallsV2ListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrFirewallsV2ListResponse) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallsV2ListResponse) SetHeaders(v map[string]*string) *DescribeTrFirewallsV2ListResponse {
	s.Headers = v
	return s
}

func (s *DescribeTrFirewallsV2ListResponse) SetStatusCode(v int32) *DescribeTrFirewallsV2ListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponse) SetBody(v *DescribeTrFirewallsV2ListResponseBody) *DescribeTrFirewallsV2ListResponse {
	s.Body = v
	return s
}

type DescribeTrFirewallsV2RouteListRequest struct {
	CurrentPage             *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	FirewallId              *string `json:"FirewallId,omitempty" xml:"FirewallId,omitempty"`
	Lang                    *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageSize                *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TrFirewallRoutePolicyId *string `json:"TrFirewallRoutePolicyId,omitempty" xml:"TrFirewallRoutePolicyId,omitempty"`
}

func (s DescribeTrFirewallsV2RouteListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrFirewallsV2RouteListRequest) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallsV2RouteListRequest) SetCurrentPage(v string) *DescribeTrFirewallsV2RouteListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeTrFirewallsV2RouteListRequest) SetFirewallId(v string) *DescribeTrFirewallsV2RouteListRequest {
	s.FirewallId = &v
	return s
}

func (s *DescribeTrFirewallsV2RouteListRequest) SetLang(v string) *DescribeTrFirewallsV2RouteListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeTrFirewallsV2RouteListRequest) SetPageSize(v string) *DescribeTrFirewallsV2RouteListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTrFirewallsV2RouteListRequest) SetTrFirewallRoutePolicyId(v string) *DescribeTrFirewallsV2RouteListRequest {
	s.TrFirewallRoutePolicyId = &v
	return s
}

type DescribeTrFirewallsV2RouteListResponseBody struct {
	FirewallRouteDetailList []*DescribeTrFirewallsV2RouteListResponseBodyFirewallRouteDetailList `json:"FirewallRouteDetailList,omitempty" xml:"FirewallRouteDetailList,omitempty" type:"Repeated"`
	RequestId               *string                                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTrFirewallsV2RouteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrFirewallsV2RouteListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallsV2RouteListResponseBody) SetFirewallRouteDetailList(v []*DescribeTrFirewallsV2RouteListResponseBodyFirewallRouteDetailList) *DescribeTrFirewallsV2RouteListResponseBody {
	s.FirewallRouteDetailList = v
	return s
}

func (s *DescribeTrFirewallsV2RouteListResponseBody) SetRequestId(v string) *DescribeTrFirewallsV2RouteListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeTrFirewallsV2RouteListResponseBodyFirewallRouteDetailList struct {
	TrFirewallRouteDestination *string `json:"TrFirewallRouteDestination,omitempty" xml:"TrFirewallRouteDestination,omitempty"`
	TrFirewallRouteNexthop     *string `json:"TrFirewallRouteNexthop,omitempty" xml:"TrFirewallRouteNexthop,omitempty"`
	TrFirewallRoutePolicyId    *string `json:"TrFirewallRoutePolicyId,omitempty" xml:"TrFirewallRoutePolicyId,omitempty"`
	TrFirewallRouteTableId     *string `json:"TrFirewallRouteTableId,omitempty" xml:"TrFirewallRouteTableId,omitempty"`
}

func (s DescribeTrFirewallsV2RouteListResponseBodyFirewallRouteDetailList) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrFirewallsV2RouteListResponseBodyFirewallRouteDetailList) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallsV2RouteListResponseBodyFirewallRouteDetailList) SetTrFirewallRouteDestination(v string) *DescribeTrFirewallsV2RouteListResponseBodyFirewallRouteDetailList {
	s.TrFirewallRouteDestination = &v
	return s
}

func (s *DescribeTrFirewallsV2RouteListResponseBodyFirewallRouteDetailList) SetTrFirewallRouteNexthop(v string) *DescribeTrFirewallsV2RouteListResponseBodyFirewallRouteDetailList {
	s.TrFirewallRouteNexthop = &v
	return s
}

func (s *DescribeTrFirewallsV2RouteListResponseBodyFirewallRouteDetailList) SetTrFirewallRoutePolicyId(v string) *DescribeTrFirewallsV2RouteListResponseBodyFirewallRouteDetailList {
	s.TrFirewallRoutePolicyId = &v
	return s
}

func (s *DescribeTrFirewallsV2RouteListResponseBodyFirewallRouteDetailList) SetTrFirewallRouteTableId(v string) *DescribeTrFirewallsV2RouteListResponseBodyFirewallRouteDetailList {
	s.TrFirewallRouteTableId = &v
	return s
}

type DescribeTrFirewallsV2RouteListResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTrFirewallsV2RouteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTrFirewallsV2RouteListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrFirewallsV2RouteListResponse) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallsV2RouteListResponse) SetHeaders(v map[string]*string) *DescribeTrFirewallsV2RouteListResponse {
	s.Headers = v
	return s
}

func (s *DescribeTrFirewallsV2RouteListResponse) SetStatusCode(v int32) *DescribeTrFirewallsV2RouteListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTrFirewallsV2RouteListResponse) SetBody(v *DescribeTrFirewallsV2RouteListResponseBody) *DescribeTrFirewallsV2RouteListResponse {
	s.Body = v
	return s
}

type DescribeUserAssetIPTrafficInfoRequest struct {
	// The IP address of the asset.
	AssetIP *string `json:"AssetIP,omitempty" xml:"AssetIP,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// *   **zh** (default): Chinese
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The time range to query. The value is a UNIX timestamp. Unit: seconds.
	TrafficTime *string `json:"TrafficTime,omitempty" xml:"TrafficTime,omitempty"`
}

func (s DescribeUserAssetIPTrafficInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserAssetIPTrafficInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserAssetIPTrafficInfoRequest) SetAssetIP(v string) *DescribeUserAssetIPTrafficInfoRequest {
	s.AssetIP = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoRequest) SetLang(v string) *DescribeUserAssetIPTrafficInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoRequest) SetTrafficTime(v string) *DescribeUserAssetIPTrafficInfoRequest {
	s.TrafficTime = &v
	return s
}

type DescribeUserAssetIPTrafficInfoResponseBody struct {
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The network throughput, which indicates the inbound traffic rate. Unit: bit/s.
	InBps *int64 `json:"InBps,omitempty" xml:"InBps,omitempty"`
	// The inbound network throughput, which indicates the number of packets that are sent inbound per second. Unit: packets per second (pps).
	InPps *int64 `json:"InPps,omitempty" xml:"InPps,omitempty"`
	// The new connection creation rate.
	NewConn *int64 `json:"NewConn,omitempty" xml:"NewConn,omitempty"`
	// The network throughput, which indicates the outbound traffic rate. Unit: bit/s.
	OutBps *int64 `json:"OutBps,omitempty" xml:"OutBps,omitempty"`
	// The outbound network throughput, which indicates the number of packets that are sent outbound per second. Unit: pps.
	OutPps *int64 `json:"OutPps,omitempty" xml:"OutPps,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of requests.
	SessionCount *int64 `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeUserAssetIPTrafficInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserAssetIPTrafficInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserAssetIPTrafficInfoResponseBody) SetEndTime(v int64) *DescribeUserAssetIPTrafficInfoResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponseBody) SetInBps(v int64) *DescribeUserAssetIPTrafficInfoResponseBody {
	s.InBps = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponseBody) SetInPps(v int64) *DescribeUserAssetIPTrafficInfoResponseBody {
	s.InPps = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponseBody) SetNewConn(v int64) *DescribeUserAssetIPTrafficInfoResponseBody {
	s.NewConn = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponseBody) SetOutBps(v int64) *DescribeUserAssetIPTrafficInfoResponseBody {
	s.OutBps = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponseBody) SetOutPps(v int64) *DescribeUserAssetIPTrafficInfoResponseBody {
	s.OutPps = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponseBody) SetRequestId(v string) *DescribeUserAssetIPTrafficInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponseBody) SetSessionCount(v int64) *DescribeUserAssetIPTrafficInfoResponseBody {
	s.SessionCount = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponseBody) SetStartTime(v int64) *DescribeUserAssetIPTrafficInfoResponseBody {
	s.StartTime = &v
	return s
}

type DescribeUserAssetIPTrafficInfoResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserAssetIPTrafficInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserAssetIPTrafficInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserAssetIPTrafficInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserAssetIPTrafficInfoResponse) SetHeaders(v map[string]*string) *DescribeUserAssetIPTrafficInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponse) SetStatusCode(v int32) *DescribeUserAssetIPTrafficInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponse) SetBody(v *DescribeUserAssetIPTrafficInfoResponseBody) *DescribeUserAssetIPTrafficInfoResponse {
	s.Body = v
	return s
}

type DescribeUserIPSWhitelistRequest struct {
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeUserIPSWhitelistRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserIPSWhitelistRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserIPSWhitelistRequest) SetLang(v string) *DescribeUserIPSWhitelistRequest {
	s.Lang = &v
	return s
}

func (s *DescribeUserIPSWhitelistRequest) SetSourceIp(v string) *DescribeUserIPSWhitelistRequest {
	s.SourceIp = &v
	return s
}

type DescribeUserIPSWhitelistResponseBody struct {
	Ipv6Whitelists []*DescribeUserIPSWhitelistResponseBodyIpv6Whitelists `json:"Ipv6Whitelists,omitempty" xml:"Ipv6Whitelists,omitempty" type:"Repeated"`
	RequestId      *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Whitelists     []*DescribeUserIPSWhitelistResponseBodyWhitelists     `json:"Whitelists,omitempty" xml:"Whitelists,omitempty" type:"Repeated"`
}

func (s DescribeUserIPSWhitelistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserIPSWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserIPSWhitelistResponseBody) SetIpv6Whitelists(v []*DescribeUserIPSWhitelistResponseBodyIpv6Whitelists) *DescribeUserIPSWhitelistResponseBody {
	s.Ipv6Whitelists = v
	return s
}

func (s *DescribeUserIPSWhitelistResponseBody) SetRequestId(v string) *DescribeUserIPSWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserIPSWhitelistResponseBody) SetWhitelists(v []*DescribeUserIPSWhitelistResponseBodyWhitelists) *DescribeUserIPSWhitelistResponseBody {
	s.Whitelists = v
	return s
}

type DescribeUserIPSWhitelistResponseBodyIpv6Whitelists struct {
	Direction      *int64    `json:"Direction,omitempty" xml:"Direction,omitempty"`
	ListType       *int64    `json:"ListType,omitempty" xml:"ListType,omitempty"`
	ListValue      *string   `json:"ListValue,omitempty" xml:"ListValue,omitempty"`
	WhiteListValue []*string `json:"WhiteListValue,omitempty" xml:"WhiteListValue,omitempty" type:"Repeated"`
	WhiteType      *int64    `json:"WhiteType,omitempty" xml:"WhiteType,omitempty"`
}

func (s DescribeUserIPSWhitelistResponseBodyIpv6Whitelists) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserIPSWhitelistResponseBodyIpv6Whitelists) GoString() string {
	return s.String()
}

func (s *DescribeUserIPSWhitelistResponseBodyIpv6Whitelists) SetDirection(v int64) *DescribeUserIPSWhitelistResponseBodyIpv6Whitelists {
	s.Direction = &v
	return s
}

func (s *DescribeUserIPSWhitelistResponseBodyIpv6Whitelists) SetListType(v int64) *DescribeUserIPSWhitelistResponseBodyIpv6Whitelists {
	s.ListType = &v
	return s
}

func (s *DescribeUserIPSWhitelistResponseBodyIpv6Whitelists) SetListValue(v string) *DescribeUserIPSWhitelistResponseBodyIpv6Whitelists {
	s.ListValue = &v
	return s
}

func (s *DescribeUserIPSWhitelistResponseBodyIpv6Whitelists) SetWhiteListValue(v []*string) *DescribeUserIPSWhitelistResponseBodyIpv6Whitelists {
	s.WhiteListValue = v
	return s
}

func (s *DescribeUserIPSWhitelistResponseBodyIpv6Whitelists) SetWhiteType(v int64) *DescribeUserIPSWhitelistResponseBodyIpv6Whitelists {
	s.WhiteType = &v
	return s
}

type DescribeUserIPSWhitelistResponseBodyWhitelists struct {
	Direction      *int64    `json:"Direction,omitempty" xml:"Direction,omitempty"`
	ListType       *int64    `json:"ListType,omitempty" xml:"ListType,omitempty"`
	ListValue      *string   `json:"ListValue,omitempty" xml:"ListValue,omitempty"`
	WhiteListValue []*string `json:"WhiteListValue,omitempty" xml:"WhiteListValue,omitempty" type:"Repeated"`
	WhiteType      *int64    `json:"WhiteType,omitempty" xml:"WhiteType,omitempty"`
}

func (s DescribeUserIPSWhitelistResponseBodyWhitelists) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserIPSWhitelistResponseBodyWhitelists) GoString() string {
	return s.String()
}

func (s *DescribeUserIPSWhitelistResponseBodyWhitelists) SetDirection(v int64) *DescribeUserIPSWhitelistResponseBodyWhitelists {
	s.Direction = &v
	return s
}

func (s *DescribeUserIPSWhitelistResponseBodyWhitelists) SetListType(v int64) *DescribeUserIPSWhitelistResponseBodyWhitelists {
	s.ListType = &v
	return s
}

func (s *DescribeUserIPSWhitelistResponseBodyWhitelists) SetListValue(v string) *DescribeUserIPSWhitelistResponseBodyWhitelists {
	s.ListValue = &v
	return s
}

func (s *DescribeUserIPSWhitelistResponseBodyWhitelists) SetWhiteListValue(v []*string) *DescribeUserIPSWhitelistResponseBodyWhitelists {
	s.WhiteListValue = v
	return s
}

func (s *DescribeUserIPSWhitelistResponseBodyWhitelists) SetWhiteType(v int64) *DescribeUserIPSWhitelistResponseBodyWhitelists {
	s.WhiteType = &v
	return s
}

type DescribeUserIPSWhitelistResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserIPSWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserIPSWhitelistResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserIPSWhitelistResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserIPSWhitelistResponse) SetHeaders(v map[string]*string) *DescribeUserIPSWhitelistResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserIPSWhitelistResponse) SetStatusCode(v int32) *DescribeUserIPSWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserIPSWhitelistResponse) SetBody(v *DescribeUserIPSWhitelistResponseBody) *DescribeUserIPSWhitelistResponse {
	s.Body = v
	return s
}

type DescribeVpcFirewallAclGroupListRequest struct {
	// The number of the page to return. Default value: 1.
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Specifies whether VPC firewalls are configured. Valid values:
	//
	// *   **notconfigured**: VPC firewalls are not configured.
	// *   **configured**: VPC firewalls are configured.
	// *   If you do not specify this parameter, the access control policies of all VPC firewalls are queried.
	FirewallConfigureStatus *string `json:"FirewallConfigureStatus,omitempty" xml:"FirewallConfigureStatus,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries to return on each page. Maximum value: 50.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeVpcFirewallAclGroupListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallAclGroupListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallAclGroupListRequest) SetCurrentPage(v string) *DescribeVpcFirewallAclGroupListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVpcFirewallAclGroupListRequest) SetFirewallConfigureStatus(v string) *DescribeVpcFirewallAclGroupListRequest {
	s.FirewallConfigureStatus = &v
	return s
}

func (s *DescribeVpcFirewallAclGroupListRequest) SetLang(v string) *DescribeVpcFirewallAclGroupListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVpcFirewallAclGroupListRequest) SetPageSize(v string) *DescribeVpcFirewallAclGroupListRequest {
	s.PageSize = &v
	return s
}

type DescribeVpcFirewallAclGroupListResponseBody struct {
	// The information about the policy groups.
	AclGroupList []*DescribeVpcFirewallAclGroupListResponseBodyAclGroupList `json:"AclGroupList,omitempty" xml:"AclGroupList,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the policy groups that are returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeVpcFirewallAclGroupListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallAclGroupListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallAclGroupListResponseBody) SetAclGroupList(v []*DescribeVpcFirewallAclGroupListResponseBodyAclGroupList) *DescribeVpcFirewallAclGroupListResponseBody {
	s.AclGroupList = v
	return s
}

func (s *DescribeVpcFirewallAclGroupListResponseBody) SetRequestId(v string) *DescribeVpcFirewallAclGroupListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallAclGroupListResponseBody) SetTotalCount(v int32) *DescribeVpcFirewallAclGroupListResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeVpcFirewallAclGroupListResponseBodyAclGroupList struct {
	// The ID of the policy group.
	//
	// Valid values:
	//
	// *   If the VPC firewall is used to protect a Cloud Enterprise Network (CEN) instance, the value of this parameter is the ID of the CEN instance.
	//
	//     Example: cen-ervw0g12b5jbw\*\*\*\*
	//
	// *   If the VPC firewall is used to protect an Express Connect circuit, the value of this parameter is the instance ID of the VPC firewall.
	//
	//     Example: vfw-a42bbb7b887148c9\*\*\*\*
	AclGroupId *string `json:"AclGroupId,omitempty" xml:"AclGroupId,omitempty"`
	// The name of the policy group. Valid values:
	//
	// *   If the VPC firewall is used to protect a CEN instance, the value of this parameter is the name of the CEN instance.
	// *   If the VPC firewall is used to protect an Express Connect circuit, the value of this parameter is the instance name of the VPC firewall.
	AclGroupName *string `json:"AclGroupName,omitempty" xml:"AclGroupName,omitempty"`
	// The number of access control policies in the policy group.
	AclRuleCount *int32 `json:"AclRuleCount,omitempty" xml:"AclRuleCount,omitempty"`
	// 是否是默认防火墙。取值：
	// - **true**：是默认防火墙。
	// - **false**：不是默认防火墙。
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
}

func (s DescribeVpcFirewallAclGroupListResponseBodyAclGroupList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallAclGroupListResponseBodyAclGroupList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallAclGroupListResponseBodyAclGroupList) SetAclGroupId(v string) *DescribeVpcFirewallAclGroupListResponseBodyAclGroupList {
	s.AclGroupId = &v
	return s
}

func (s *DescribeVpcFirewallAclGroupListResponseBodyAclGroupList) SetAclGroupName(v string) *DescribeVpcFirewallAclGroupListResponseBodyAclGroupList {
	s.AclGroupName = &v
	return s
}

func (s *DescribeVpcFirewallAclGroupListResponseBodyAclGroupList) SetAclRuleCount(v int32) *DescribeVpcFirewallAclGroupListResponseBodyAclGroupList {
	s.AclRuleCount = &v
	return s
}

func (s *DescribeVpcFirewallAclGroupListResponseBodyAclGroupList) SetIsDefault(v bool) *DescribeVpcFirewallAclGroupListResponseBodyAclGroupList {
	s.IsDefault = &v
	return s
}

func (s *DescribeVpcFirewallAclGroupListResponseBodyAclGroupList) SetMemberUid(v string) *DescribeVpcFirewallAclGroupListResponseBodyAclGroupList {
	s.MemberUid = &v
	return s
}

type DescribeVpcFirewallAclGroupListResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcFirewallAclGroupListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcFirewallAclGroupListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallAclGroupListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallAclGroupListResponse) SetHeaders(v map[string]*string) *DescribeVpcFirewallAclGroupListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcFirewallAclGroupListResponse) SetStatusCode(v int32) *DescribeVpcFirewallAclGroupListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcFirewallAclGroupListResponse) SetBody(v *DescribeVpcFirewallAclGroupListResponseBody) *DescribeVpcFirewallAclGroupListResponse {
	s.Body = v
	return s
}

type DescribeVpcFirewallCenDetailRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the VPC for which the VPC firewall is created.
	NetworkInstanceId *string `json:"NetworkInstanceId,omitempty" xml:"NetworkInstanceId,omitempty"`
	// The instance ID of the VPC firewall.
	//
	// > You can call the [DescribeVpcFirewallCenList](~~345777~~) operation to query the instance IDs of VPC firewalls.
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s DescribeVpcFirewallCenDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallCenDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenDetailRequest) SetLang(v string) *DescribeVpcFirewallCenDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailRequest) SetNetworkInstanceId(v string) *DescribeVpcFirewallCenDetailRequest {
	s.NetworkInstanceId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailRequest) SetVpcFirewallId(v string) *DescribeVpcFirewallCenDetailRequest {
	s.VpcFirewallId = &v
	return s
}

type DescribeVpcFirewallCenDetailResponseBody struct {
	// The connection type of the VPC firewall. The value is fixed as **cen**, which indicates CEN instances.
	ConnectType *string `json:"ConnectType,omitempty" xml:"ConnectType,omitempty"`
	// The status of the VPC firewall. Valid values:
	//
	// *   **opened**: enabled
	// *   **closed**: disabled
	// *   **notconfigured**: not configured
	FirewallSwitchStatus *string `json:"FirewallSwitchStatus,omitempty" xml:"FirewallSwitchStatus,omitempty"`
	// The VPC that is automatically created for the firewall.
	FirewallVpc *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc `json:"FirewallVpc,omitempty" xml:"FirewallVpc,omitempty" type:"Struct"`
	// The details about the VPC.
	LocalVpc *DescribeVpcFirewallCenDetailResponseBodyLocalVpc `json:"LocalVpc,omitempty" xml:"LocalVpc,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The instance ID of the VPC firewall.
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
	// The instance name of the VPC firewall.
	VpcFirewallName *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
}

func (s DescribeVpcFirewallCenDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallCenDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenDetailResponseBody) SetConnectType(v string) *DescribeVpcFirewallCenDetailResponseBody {
	s.ConnectType = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBody) SetFirewallSwitchStatus(v string) *DescribeVpcFirewallCenDetailResponseBody {
	s.FirewallSwitchStatus = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBody) SetFirewallVpc(v *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc) *DescribeVpcFirewallCenDetailResponseBody {
	s.FirewallVpc = v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBody) SetLocalVpc(v *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) *DescribeVpcFirewallCenDetailResponseBody {
	s.LocalVpc = v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBody) SetRequestId(v string) *DescribeVpcFirewallCenDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBody) SetVpcFirewallId(v string) *DescribeVpcFirewallCenDetailResponseBody {
	s.VpcFirewallId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBody) SetVpcFirewallName(v string) *DescribeVpcFirewallCenDetailResponseBody {
	s.VpcFirewallName = &v
	return s
}

type DescribeVpcFirewallCenDetailResponseBodyFirewallVpc struct {
	// Indicates whether you can specify a CIDR block when you create a VPC firewall for a Basic Edition transit router of a CEN instance. Valid values:
	//
	// *   **1**: yes
	// *   **0**: no
	AllowConfiguration *int32 `json:"AllowConfiguration,omitempty" xml:"AllowConfiguration,omitempty"`
	// The CIDR block of the VPC.
	VpcCidr *string `json:"VpcCidr,omitempty" xml:"VpcCidr,omitempty"`
	// The VPC ID.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The CIDR block of the vSwitch.
	VswitchCidr *string `json:"VswitchCidr,omitempty" xml:"VswitchCidr,omitempty"`
	// The vSwitch ID.
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	// The zone ID.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeVpcFirewallCenDetailResponseBodyFirewallVpc) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallCenDetailResponseBodyFirewallVpc) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc) SetAllowConfiguration(v int32) *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc {
	s.AllowConfiguration = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc) SetVpcCidr(v string) *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc {
	s.VpcCidr = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc) SetVpcId(v string) *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc) SetVswitchCidr(v string) *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc {
	s.VswitchCidr = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc) SetVswitchId(v string) *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc {
	s.VswitchId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc) SetZoneId(v string) *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc {
	s.ZoneId = &v
	return s
}

type DescribeVpcFirewallCenDetailResponseBodyLocalVpc struct {
	// The ID of the connection between two network instances.
	AttachmentId *string `json:"AttachmentId,omitempty" xml:"AttachmentId,omitempty"`
	// The name of the connection between two network instances.
	AttachmentName *string `json:"AttachmentName,omitempty" xml:"AttachmentName,omitempty"`
	// An array consisting of the CIDR blocks that are protected by the VPC firewall.
	DefendCidrList []*string `json:"DefendCidrList,omitempty" xml:"DefendCidrList,omitempty" type:"Repeated"`
	// The Elastic Network Interfaces (ENIs).
	EniList []*DescribeVpcFirewallCenDetailResponseBodyLocalVpcEniList `json:"EniList,omitempty" xml:"EniList,omitempty" type:"Repeated"`
	// The ID of the specified vSwitch when the routing mode is manual.
	ManualVSwitchId *string `json:"ManualVSwitchId,omitempty" xml:"ManualVSwitchId,omitempty"`
	// The ID of the VPC for which the VPC firewall is created.
	NetworkInstanceId *string `json:"NetworkInstanceId,omitempty" xml:"NetworkInstanceId,omitempty"`
	// The name of the network instance.
	NetworkInstanceName *string `json:"NetworkInstanceName,omitempty" xml:"NetworkInstanceName,omitempty"`
	// The type of the network instance. The value is fixed as **VPC**.
	NetworkInstanceType *string `json:"NetworkInstanceType,omitempty" xml:"NetworkInstanceType,omitempty"`
	// The UID of the Alibaba Cloud account to which the VPC belongs.
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region in which the VPC resides.
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The routing mode. Valid values:
	//
	// *   auto: automatic mode
	// *   manual: manual mode
	RouteMode *string `json:"RouteMode,omitempty" xml:"RouteMode,omitempty"`
	// Indicates whether the manual routing mode is supported. Valid values:
	//
	// *   **1**: yes
	// *   **0**: no
	SupportManualMode *string `json:"SupportManualMode,omitempty" xml:"SupportManualMode,omitempty"`
	// The instance ID of the CEN transit router.
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// The edition of the CEN transit router. Valid values:
	//
	// *   **Basic**: Basic Edition
	// *   **Enterprise**: Enterprise Edition
	TransitRouterType *string `json:"TransitRouterType,omitempty" xml:"TransitRouterType,omitempty"`
	// An array that consists of the CIDR blocks of the VPC.
	VpcCidrTableList []*DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableList `json:"VpcCidrTableList,omitempty" xml:"VpcCidrTableList,omitempty" type:"Repeated"`
	// The ID of the VPC.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the VPC.
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeVpcFirewallCenDetailResponseBodyLocalVpc) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallCenDetailResponseBodyLocalVpc) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetAttachmentId(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.AttachmentId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetAttachmentName(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.AttachmentName = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetDefendCidrList(v []*string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.DefendCidrList = v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetEniList(v []*DescribeVpcFirewallCenDetailResponseBodyLocalVpcEniList) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.EniList = v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetManualVSwitchId(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.ManualVSwitchId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetNetworkInstanceId(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.NetworkInstanceId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetNetworkInstanceName(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.NetworkInstanceName = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetNetworkInstanceType(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.NetworkInstanceType = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetOwnerId(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.OwnerId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetRegionNo(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.RegionNo = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetRouteMode(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.RouteMode = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetSupportManualMode(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.SupportManualMode = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetTransitRouterId(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.TransitRouterId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetTransitRouterType(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.TransitRouterType = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetVpcCidrTableList(v []*DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableList) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.VpcCidrTableList = v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetVpcId(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetVpcName(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.VpcName = &v
	return s
}

type DescribeVpcFirewallCenDetailResponseBodyLocalVpcEniList struct {
	// The ID of the ENI that belongs to the VPC.
	EniId *string `json:"EniId,omitempty" xml:"EniId,omitempty"`
	// The private IP address of the ENI that belongs to the VPC.
	EniPrivateIpAddress *string `json:"EniPrivateIpAddress,omitempty" xml:"EniPrivateIpAddress,omitempty"`
	// The ID of the vSwitch to which the ENI is connected.
	EniVSwitchId *string `json:"EniVSwitchId,omitempty" xml:"EniVSwitchId,omitempty"`
}

func (s DescribeVpcFirewallCenDetailResponseBodyLocalVpcEniList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallCenDetailResponseBodyLocalVpcEniList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpcEniList) SetEniId(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpcEniList {
	s.EniId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpcEniList) SetEniPrivateIpAddress(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpcEniList {
	s.EniPrivateIpAddress = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpcEniList) SetEniVSwitchId(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpcEniList {
	s.EniVSwitchId = &v
	return s
}

type DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableList struct {
	// The route entries for the VPC.
	RouteEntryList []*DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList `json:"RouteEntryList,omitempty" xml:"RouteEntryList,omitempty" type:"Repeated"`
	// The route table ID of the VPC.
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableList) SetRouteEntryList(v []*DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList) *DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableList {
	s.RouteEntryList = v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableList) SetRouteTableId(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableList {
	s.RouteTableId = &v
	return s
}

type DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList struct {
	// The destination CIDR block of the VPC.
	DestinationCidr *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
	// The instance ID of the next hop for the VPC.
	NextHopInstanceId *string `json:"NextHopInstanceId,omitempty" xml:"NextHopInstanceId,omitempty"`
}

func (s DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList) SetDestinationCidr(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList {
	s.DestinationCidr = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList) SetNextHopInstanceId(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList {
	s.NextHopInstanceId = &v
	return s
}

type DescribeVpcFirewallCenDetailResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcFirewallCenDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcFirewallCenDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallCenDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenDetailResponse) SetHeaders(v map[string]*string) *DescribeVpcFirewallCenDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponse) SetStatusCode(v int32) *DescribeVpcFirewallCenDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponse) SetBody(v *DescribeVpcFirewallCenDetailResponseBody) *DescribeVpcFirewallCenDetailResponse {
	s.Body = v
	return s
}

type DescribeVpcFirewallCenListRequest struct {
	// The ID of the CEN instance.
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The status of the VPC firewall. Valid values:
	//
	// *   **opened**: The VPC firewall is enabled.
	// *   **closed**: The VPC firewall is disabled.
	// *   **notconfigured**: The VPC firewall is not configured.
	// *   **configured**: The VPC firewall is configured but is not enabled.
	//
	// > If you do not specify this parameter, VPC firewalls in all states are queried.
	FirewallSwitchStatus *string `json:"FirewallSwitchStatus,omitempty" xml:"FirewallSwitchStatus,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account. The member is also an Alibaba Cloud account.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The ID of the network instance.
	NetworkInstanceId *string `json:"NetworkInstanceId,omitempty" xml:"NetworkInstanceId,omitempty"`
	OwnerId           *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of entries to return on each page.
	//
	// Default value: 10. Maximum value: 50.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the VPC.
	//
	// > For more information about the regions, see [Supported regions](~~195657~~).
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The routing mode of the VPC firewall. Valid values:
	//
	// *   **auto**: automatic mode
	// *   **manual**: manual mode
	//
	// > If you do not specify this parameter, VPC firewalls in all routing modes are queried.
	RouteMode *string `json:"RouteMode,omitempty" xml:"RouteMode,omitempty"`
	// The type of the transit router. Valid values:
	//
	// *   **Basic**: Basic Edition transit router
	// *   **Enterprise**: Enterprise Edition transit router
	TransitRouterType *string `json:"TransitRouterType,omitempty" xml:"TransitRouterType,omitempty"`
	// The instance ID of the VPC firewall.
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
	// The instance name of the VPC firewall.
	VpcFirewallName *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
}

func (s DescribeVpcFirewallCenListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallCenListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenListRequest) SetCenId(v string) *DescribeVpcFirewallCenListRequest {
	s.CenId = &v
	return s
}

func (s *DescribeVpcFirewallCenListRequest) SetCurrentPage(v string) *DescribeVpcFirewallCenListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVpcFirewallCenListRequest) SetFirewallSwitchStatus(v string) *DescribeVpcFirewallCenListRequest {
	s.FirewallSwitchStatus = &v
	return s
}

func (s *DescribeVpcFirewallCenListRequest) SetLang(v string) *DescribeVpcFirewallCenListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVpcFirewallCenListRequest) SetMemberUid(v string) *DescribeVpcFirewallCenListRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeVpcFirewallCenListRequest) SetNetworkInstanceId(v string) *DescribeVpcFirewallCenListRequest {
	s.NetworkInstanceId = &v
	return s
}

func (s *DescribeVpcFirewallCenListRequest) SetOwnerId(v string) *DescribeVpcFirewallCenListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVpcFirewallCenListRequest) SetPageSize(v string) *DescribeVpcFirewallCenListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVpcFirewallCenListRequest) SetRegionNo(v string) *DescribeVpcFirewallCenListRequest {
	s.RegionNo = &v
	return s
}

func (s *DescribeVpcFirewallCenListRequest) SetRouteMode(v string) *DescribeVpcFirewallCenListRequest {
	s.RouteMode = &v
	return s
}

func (s *DescribeVpcFirewallCenListRequest) SetTransitRouterType(v string) *DescribeVpcFirewallCenListRequest {
	s.TransitRouterType = &v
	return s
}

func (s *DescribeVpcFirewallCenListRequest) SetVpcFirewallId(v string) *DescribeVpcFirewallCenListRequest {
	s.VpcFirewallId = &v
	return s
}

func (s *DescribeVpcFirewallCenListRequest) SetVpcFirewallName(v string) *DescribeVpcFirewallCenListRequest {
	s.VpcFirewallName = &v
	return s
}

type DescribeVpcFirewallCenListResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of VPC firewalls.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// An array that consists of the details about the VPC firewall.
	VpcFirewalls []*DescribeVpcFirewallCenListResponseBodyVpcFirewalls `json:"VpcFirewalls,omitempty" xml:"VpcFirewalls,omitempty" type:"Repeated"`
}

func (s DescribeVpcFirewallCenListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallCenListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenListResponseBody) SetRequestId(v string) *DescribeVpcFirewallCenListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBody) SetTotalCount(v int32) *DescribeVpcFirewallCenListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBody) SetVpcFirewalls(v []*DescribeVpcFirewallCenListResponseBodyVpcFirewalls) *DescribeVpcFirewallCenListResponseBody {
	s.VpcFirewalls = v
	return s
}

type DescribeVpcFirewallCenListResponseBodyVpcFirewalls struct {
	// The ID of the CEN instance.
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The name of the CEN instance.
	CenName *string `json:"CenName,omitempty" xml:"CenName,omitempty"`
	// The connection type of the VPC firewall. The value is fixed as cen, which indicates a CEN instance.
	ConnectType *string `json:"ConnectType,omitempty" xml:"ConnectType,omitempty"`
	// The status of the VPC firewall. Valid values:
	//
	// *   **opened**: The VPC firewall is enabled.
	// *   **closed**: The VPC firewall is disabled.
	// *   **notconfigured**: The VPC firewall is not configured.
	FirewallSwitchStatus *string `json:"FirewallSwitchStatus,omitempty" xml:"FirewallSwitchStatus,omitempty"`
	// The information about the intrusion prevention system (IPS) configuration.
	IpsConfig *DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig `json:"IpsConfig,omitempty" xml:"IpsConfig,omitempty" type:"Struct"`
	// The details about the VPC.
	LocalVpc *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc `json:"LocalVpc,omitempty" xml:"LocalVpc,omitempty" type:"Struct"`
	// The UID of the member that is manged by your Alibaba Cloud account. The member is also an Alibaba Cloud account.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// Indicates whether the VPC firewall can be automatically enabled to protect VPC traffic based on route learning. Valid values:
	//
	// *   **passed**: The VPC firewall can be automatically enabled.
	// *   **failed**: The VPC firewall cannot be automatically enabled.
	// *   **unknown**: The VPC firewall is in an unknown state.
	PrecheckStatus *string `json:"PrecheckStatus,omitempty" xml:"PrecheckStatus,omitempty"`
	// Indicates whether you can create a VPC firewall in a specified region. Valid values:
	//
	// *   **enable**: yes
	// *   **disable**: no
	RegionStatus *string `json:"RegionStatus,omitempty" xml:"RegionStatus,omitempty"`
	// The result code of the operation that creates the VPC firewall. Valid values:
	//
	// *   **Unauthorized**: Cloud Firewall is not authorized to access the VPC for which the VPC firewall is created, and the VPC firewall cannot be created.
	// *   **RegionDisable**: VPC Firewall is not supported in the region of the VPC for which the VPC firewall is created, and the VPC firewall cannot be created.
	// *   **OpsDisable**: You are not allowed to create the VPC firewall.
	// *   **VbrNotSupport**: The VPC firewall cannot be created for a VBR that is attached to the CEN instance.
	// *   Empty string: You can create a VPC firewall for the network instance.
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// The instance ID of the VPC firewall.
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
	// The instance name of the VPC firewall.
	VpcFirewallName *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
}

func (s DescribeVpcFirewallCenListResponseBodyVpcFirewalls) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallCenListResponseBodyVpcFirewalls) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) SetCenId(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewalls {
	s.CenId = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) SetCenName(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewalls {
	s.CenName = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) SetConnectType(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewalls {
	s.ConnectType = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) SetFirewallSwitchStatus(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewalls {
	s.FirewallSwitchStatus = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) SetIpsConfig(v *DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig) *DescribeVpcFirewallCenListResponseBodyVpcFirewalls {
	s.IpsConfig = v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) SetLocalVpc(v *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) *DescribeVpcFirewallCenListResponseBodyVpcFirewalls {
	s.LocalVpc = v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) SetMemberUid(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewalls {
	s.MemberUid = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) SetPrecheckStatus(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewalls {
	s.PrecheckStatus = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) SetRegionStatus(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewalls {
	s.RegionStatus = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) SetResultCode(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewalls {
	s.ResultCode = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) SetVpcFirewallId(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewalls {
	s.VpcFirewallId = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) SetVpcFirewallName(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewalls {
	s.VpcFirewallName = &v
	return s
}

type DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig struct {
	// Indicates whether basic protection is enabled. Valid values:
	//
	// *   **1**: yes
	// *   **0**: no
	BasicRules *int32 `json:"BasicRules,omitempty" xml:"BasicRules,omitempty"`
	// Indicates whether virtual patching is enabled. Valid values:
	//
	// *   **1**: yes
	// *   **0**: no
	EnableAllPatch *int32 `json:"EnableAllPatch,omitempty" xml:"EnableAllPatch,omitempty"`
	// The mode of the IPS. Valid values:
	//
	// *   **1**: block mode
	// *   **0**: monitor mode
	RunMode *int32 `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
}

func (s DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig) SetBasicRules(v int32) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig {
	s.BasicRules = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig) SetEnableAllPatch(v int32) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig {
	s.EnableAllPatch = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig) SetRunMode(v int32) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig {
	s.RunMode = &v
	return s
}

type DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc struct {
	// Indicates whether the VPC is granted the required permissions. The value is fixed as **authorized**, which indicates that the VPC is granted the required permissions.
	AuthorizationStatus *string `json:"AuthorizationStatus,omitempty" xml:"AuthorizationStatus,omitempty"`
	// An array consisting of the CIDR blocks that are protected by the VPC firewall.
	DefendCidrList []*string `json:"DefendCidrList,omitempty" xml:"DefendCidrList,omitempty" type:"Repeated"`
	// The ID of the specified vSwitch when the routing mode is manual.
	ManualVSwitchId *string `json:"ManualVSwitchId,omitempty" xml:"ManualVSwitchId,omitempty"`
	// The ID of the network instance.
	NetworkInstanceId *string `json:"NetworkInstanceId,omitempty" xml:"NetworkInstanceId,omitempty"`
	// The name of the network instance.
	NetworkInstanceName *string `json:"NetworkInstanceName,omitempty" xml:"NetworkInstanceName,omitempty"`
	// The type of the network instance. Valid values:
	//
	// *   **VPC**
	// *   **VBR**
	// *   **CCN**
	NetworkInstanceType *string `json:"NetworkInstanceType,omitempty" xml:"NetworkInstanceType,omitempty"`
	// The ID of the Alibaba Cloud account to which the VPC belongs.
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the VPC.
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The routing mode of the VPC firewall. Valid values:
	//
	// *   **auto**: automatic mode
	// *   **manual**: manual mode
	RouteMode *string `json:"RouteMode,omitempty" xml:"RouteMode,omitempty"`
	// Indicates whether the manual routing mode is supported. Valid values:
	//
	// *   **1**: yes
	// *   **0**: no
	SupportManualMode *string `json:"SupportManualMode,omitempty" xml:"SupportManualMode,omitempty"`
	// The edition of the CEN transit router. Valid values:
	//
	// *   **Basic**: Basic Edition transit router
	// *   **Enterprise**: Enterprise Edition transit router
	TransitRouterType *string `json:"TransitRouterType,omitempty" xml:"TransitRouterType,omitempty"`
	// An array that consists of the CIDR blocks of the VPC.
	VpcCidrTableList []*DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList `json:"VpcCidrTableList,omitempty" xml:"VpcCidrTableList,omitempty" type:"Repeated"`
	// The ID of the VPC.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the VPC.
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetAuthorizationStatus(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.AuthorizationStatus = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetDefendCidrList(v []*string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.DefendCidrList = v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetManualVSwitchId(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.ManualVSwitchId = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetNetworkInstanceId(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.NetworkInstanceId = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetNetworkInstanceName(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.NetworkInstanceName = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetNetworkInstanceType(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.NetworkInstanceType = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetOwnerId(v int64) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.OwnerId = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetRegionNo(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.RegionNo = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetRouteMode(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.RouteMode = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetSupportManualMode(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.SupportManualMode = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetTransitRouterType(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.TransitRouterType = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetVpcCidrTableList(v []*DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.VpcCidrTableList = v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetVpcId(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetVpcName(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.VpcName = &v
	return s
}

type DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList struct {
	// An array that consists of the route entries for the VPC.
	RouteEntryList []*DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList `json:"RouteEntryList,omitempty" xml:"RouteEntryList,omitempty" type:"Repeated"`
	// The route table ID of the VPC.
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList) SetRouteEntryList(v []*DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList {
	s.RouteEntryList = v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList) SetRouteTableId(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList {
	s.RouteTableId = &v
	return s
}

type DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList struct {
	// The destination CIDR block of the VPC.
	DestinationCidr *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
	// The instance ID of the next hop for the VPC.
	NextHopInstanceId *string `json:"NextHopInstanceId,omitempty" xml:"NextHopInstanceId,omitempty"`
}

func (s DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList) SetDestinationCidr(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList {
	s.DestinationCidr = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList) SetNextHopInstanceId(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList {
	s.NextHopInstanceId = &v
	return s
}

type DescribeVpcFirewallCenListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcFirewallCenListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcFirewallCenListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallCenListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenListResponse) SetHeaders(v map[string]*string) *DescribeVpcFirewallCenListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcFirewallCenListResponse) SetStatusCode(v int32) *DescribeVpcFirewallCenListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponse) SetBody(v *DescribeVpcFirewallCenListResponseBody) *DescribeVpcFirewallCenListResponse {
	s.Body = v
	return s
}

type DescribeVpcFirewallControlPolicyRequest struct {
	// The action that Cloud Firewall performs on the traffic. Valid values:
	//
	// *   **accept**: allows the traffic.
	// *   **drop**: blocks the traffic.
	// *   **log**: monitors the traffic.
	//
	// > If you do not specify this parameter, access control policies are queried based on all actions.
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The unique ID of the access control policy.
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The number of the page to return.
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The description of the access control policy. Fuzzy match is supported.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination address in the access control policy. Fuzzy match is supported.
	//
	// > The value of this parameter can be a CIDR block or an address book name.
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The language of the content within the request and response.
	//
	// Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The number of entries to return on each page.
	//
	// Maximum value: 50.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The protocol type in the access control policy. Valid values:
	//
	// *   **TCP**
	// *   **UDP**
	// *   **ICMP**
	// *   **ANY**: all protocol types
	//
	// > If you do not specify this parameter, access control policies of all protocol types are queried.
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// The status of the access control policy. Valid values:
	//
	// *   **true**: enabled
	// *   **false**: disabled
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The recurrence type for the access control policy to take effect. Valid values:
	//
	// *   **Permanent** (default): The policy always takes effect.
	// *   **None**: The policy takes effect for only once.
	// *   **Daily**: The policy takes effect on a daily basis.
	// *   **Weekly**: The policy takes effect on a weekly basis.
	// *   **Monthly**: The policy takes effect on a monthly basis.
	RepeatType *string `json:"RepeatType,omitempty" xml:"RepeatType,omitempty"`
	// The source address in the access control policy. Fuzzy match is supported.
	//
	// > The value of this parameter can be a CIDR block or an address book name.
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The instance ID of the VPC firewall. Valid values:
	//
	// *   If the VPC firewall protects the traffic between two VPCs that are connected by using a CEN instance, the value of this parameter must be the ID of the CEN instance.
	// *   If the VPC firewall protects the traffic between two VPCs that are connected by using an Express Connect circuit, the value of this parameter must be the instance ID of the VPC firewall.
	//
	// > You can call the [DescribeVpcFirewallAclGroupList](~~159760~~) operation to query the ID.
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s DescribeVpcFirewallControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetAclAction(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.AclAction = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetAclUuid(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.AclUuid = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetCurrentPage(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetDescription(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.Description = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetDestination(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.Destination = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetLang(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetMemberUid(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetPageSize(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetProto(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.Proto = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetRelease(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.Release = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetRepeatType(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.RepeatType = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetSource(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.Source = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetVpcFirewallId(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.VpcFirewallId = &v
	return s
}

type DescribeVpcFirewallControlPolicyResponseBody struct {
	// The access control policies.
	Policys []*DescribeVpcFirewallControlPolicyResponseBodyPolicys `json:"Policys,omitempty" xml:"Policys,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of access control policies returned.
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeVpcFirewallControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallControlPolicyResponseBody) SetPolicys(v []*DescribeVpcFirewallControlPolicyResponseBodyPolicys) *DescribeVpcFirewallControlPolicyResponseBody {
	s.Policys = v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBody) SetRequestId(v string) *DescribeVpcFirewallControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBody) SetTotalCount(v string) *DescribeVpcFirewallControlPolicyResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeVpcFirewallControlPolicyResponseBodyPolicys struct {
	// The action that Cloud Firewall performs on the traffic. Valid values:
	//
	// *   **accept**: allows the traffic.
	// *   **drop**: denies the traffic.
	// *   **log**: monitors the traffic.
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The UUID of the access control policy.
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The application ID in the access control policy.
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The application types supported by the access control policy. We recommend that you specify ApplicationNameList. Valid values:
	//
	// *   **HTTP**
	// *   **HTTPS**
	// *   **MySQL**
	// *   **SMTP**
	// *   **SMTPS**
	// *   **RDP**
	// *   **VNC**
	// *   **SSH**
	// *   **Redis**
	// *   **MQTT**
	// *   **MongoDB**
	// *   **Memcache**
	// *   **SSL**
	// *   **ANY**: all application types
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The application types supported by the access control policy.
	ApplicationNameList []*string `json:"ApplicationNameList,omitempty" xml:"ApplicationNameList,omitempty" type:"Repeated"`
	// The time when the access control policy was created. The value is a UNIX timestamp. Unit: seconds.
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the access control policy.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination port in the access control policy.
	DestPort *string `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	// The name of the destination port address book in the access control policy.
	DestPortGroup *string `json:"DestPortGroup,omitempty" xml:"DestPortGroup,omitempty"`
	// The ports in the destination port address book of the access control policy.
	DestPortGroupPorts []*string `json:"DestPortGroupPorts,omitempty" xml:"DestPortGroupPorts,omitempty" type:"Repeated"`
	// The type of the destination port in the access control policy. Valid values:
	//
	// *   **port**: port
	// *   **group**: port address book
	DestPortType *string `json:"DestPortType,omitempty" xml:"DestPortType,omitempty"`
	// The destination address in the access control policy. Valid values:
	//
	// *   If **DestinationType** is set to `net`, the value of this parameter is a CIDR block.
	// *   If **DestinationType** is set to `domain`, the value of this parameter is a domain name.
	// *   If **DestinationType** is set to `group`, the value of this parameter is an address book name.
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The CIDR blocks in the destination address book of the access control policy.
	DestinationGroupCidrs []*string `json:"DestinationGroupCidrs,omitempty" xml:"DestinationGroupCidrs,omitempty" type:"Repeated"`
	// The type of the destination address book in the access control policy. Valid values:
	//
	// *   **ip**: an address book that includes one or more CIDR blocks
	// *   **domain**: an address book that includes one or more domain names
	DestinationGroupType *string `json:"DestinationGroupType,omitempty" xml:"DestinationGroupType,omitempty"`
	// The type of the destination address in the access control policy. Valid values:
	//
	// *   **net**: CIDR block
	// *   **group**: address book
	// *   **domain**: domain name
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The time when the access control policy stops taking effect. The value is a UNIX timestamp. Unit: seconds. The value must be on the hour or on the half hour, and at least 30 minutes later than the value of StartTime.
	//
	// >  If RepeatType is set to Permanent, EndTime is left empty. If RepeatType is set to None, Daily, Weekly, or Monthly, EndTime must be specified.
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time when the access control policy was last hit. The value is a UNIX timestamp. Unit: seconds.
	HitLastTime *int64 `json:"HitLastTime,omitempty" xml:"HitLastTime,omitempty"`
	// The number of hits for the access control policy.
	HitTimes *int64 `json:"HitTimes,omitempty" xml:"HitTimes,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The time when the access control policy was modified. The value is a UNIX timestamp. Unit: seconds.
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The priority of the access control policy.
	//
	// The priority value starts from 1. A smaller priority value indicates a higher priority.
	Order *int32 `json:"Order,omitempty" xml:"Order,omitempty"`
	// The protocol type in the access control policy. Valid values:
	//
	// *   **TCP**
	// *   **UDP**
	// *   **ICMP**
	// *   **ANY**: all protocol types
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// Indicates whether the access control policy is enabled. By default, an access control policy is enabled after it is created. Valid values:
	//
	// *   **true**
	// *   **false**
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The days of a week or of a month on which the access control policy takes effect.
	//
	// *   If RepeatType is set to `Permanent`, `None`, or `Daily`, RepeatDays is left empty. Example: \[].
	// *   If RepeatType is set to Weekly, RepeatDays must be specified. Example: \[0, 6].
	//
	// >  If RepeatType is set to Weekly, the fields in the value of RepeatDays cannot be repeated.
	//
	// *   If RepeatType is set to `Monthly`, RepeatDays must be specified. Example: \[1, 31].
	//
	// >  If RepeatType is set to Monthly, the fields in the value of RepeatDays cannot be repeated.
	RepeatDays []*int64 `json:"RepeatDays,omitempty" xml:"RepeatDays,omitempty" type:"Repeated"`
	// The point in time when the recurrence ends. Example: 23:30. The value must be on the hour or on the half hour, and at least 30 minutes later than the value of RepeatStartTime.
	//
	// >  If RepeatType is set to Permanent or None, RepeatEndTime is left empty. If RepeatType is set to Daily, Weekly, or Monthly, RepeatEndTime must be specified.
	RepeatEndTime *string `json:"RepeatEndTime,omitempty" xml:"RepeatEndTime,omitempty"`
	// The point in time when the recurrence starts. Example: 08:00. The value must be on the hour or on the half hour, and at least 30 minutes earlier than the value of RepeatEndTime.
	//
	// >  If RepeatType is set to Permanent or None, RepeatStartTime is left empty. If RepeatType is set to Daily, Weekly, or Monthly, this parameter must be specified.
	RepeatStartTime *string `json:"RepeatStartTime,omitempty" xml:"RepeatStartTime,omitempty"`
	// The recurrence type for the access control policy to take effect. Valid values:
	//
	// *   **Permanent** (default): The policy always takes effect.
	// *   **None**: The policy takes effect for only once.
	// *   **Daily**: The policy takes effect on a daily basis.
	// *   **Weekly**: The policy takes effect on a weekly basis.
	// *   **Monthly**: The policy takes effect on a monthly basis.
	RepeatType *string `json:"RepeatType,omitempty" xml:"RepeatType,omitempty"`
	// The source address in the access control policy. Valid values:
	//
	// *   If **SourceType** is set to `net`, the value of this parameter is a CIDR block.
	// *   If **SourceType** is set to `group`, the value of this parameter is an address book name.
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The CIDR blocks in the source address book of the access control policy.
	SourceGroupCidrs []*string `json:"SourceGroupCidrs,omitempty" xml:"SourceGroupCidrs,omitempty" type:"Repeated"`
	// The type of the source address book in the access control policy. The value is fixed as **ip**. The value indicates an address book that includes one or more CIDR blocks.
	SourceGroupType *string `json:"SourceGroupType,omitempty" xml:"SourceGroupType,omitempty"`
	// The type of the source address in the access control policy. Valid values:
	//
	// *   **net**: CIDR block
	// *   **group**: address book
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The total quota consumed by the returned access control policies, which is the sum of the quota consumed by each policy. The quota that is consumed by an access control policy is calculated by using the following formula: Quota that is consumed by an access control policy = Number of source addresses × Number of destination addresses (number of CIDR blocks or domain names) × Number of applications × Number of port ranges.
	SpreadCnt *int64 `json:"SpreadCnt,omitempty" xml:"SpreadCnt,omitempty"`
	// The time when the access control policy starts to take effect. The value is a UNIX timestamp. Unit: seconds. The value must be on the hour or on the half hour, and at least 30 minutes earlier than the value of EndTime.
	//
	// >  If RepeatType is set to Permanent, StartTime is left empty. If RepeatType is set to None, Daily, Weekly, or Monthly, StartTime must be specified.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVpcFirewallControlPolicyResponseBodyPolicys) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallControlPolicyResponseBodyPolicys) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetAclAction(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.AclAction = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetAclUuid(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.AclUuid = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetApplicationId(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.ApplicationId = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetApplicationName(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.ApplicationName = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetApplicationNameList(v []*string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.ApplicationNameList = v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetCreateTime(v int64) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.CreateTime = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetDescription(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.Description = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetDestPort(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.DestPort = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetDestPortGroup(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.DestPortGroup = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetDestPortGroupPorts(v []*string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.DestPortGroupPorts = v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetDestPortType(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.DestPortType = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetDestination(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.Destination = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetDestinationGroupCidrs(v []*string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.DestinationGroupCidrs = v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetDestinationGroupType(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.DestinationGroupType = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetDestinationType(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.DestinationType = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetEndTime(v int64) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.EndTime = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetHitLastTime(v int64) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.HitLastTime = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetHitTimes(v int64) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.HitTimes = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetMemberUid(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.MemberUid = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetModifyTime(v int64) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.ModifyTime = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetOrder(v int32) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.Order = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetProto(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.Proto = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetRelease(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.Release = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetRepeatDays(v []*int64) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.RepeatDays = v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetRepeatEndTime(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.RepeatEndTime = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetRepeatStartTime(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.RepeatStartTime = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetRepeatType(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.RepeatType = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetSource(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.Source = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetSourceGroupCidrs(v []*string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.SourceGroupCidrs = v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetSourceGroupType(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.SourceGroupType = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetSourceType(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.SourceType = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetSpreadCnt(v int64) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.SpreadCnt = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetStartTime(v int64) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.StartTime = &v
	return s
}

type DescribeVpcFirewallControlPolicyResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcFirewallControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcFirewallControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallControlPolicyResponse) SetHeaders(v map[string]*string) *DescribeVpcFirewallControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponse) SetStatusCode(v int32) *DescribeVpcFirewallControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponse) SetBody(v *DescribeVpcFirewallControlPolicyResponseBody) *DescribeVpcFirewallControlPolicyResponse {
	s.Body = v
	return s
}

type DescribeVpcFirewallDefaultIPSConfigRequest struct {
	// The UID of the member that is managed by your Alibaba Cloud account.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The instance ID of the VPC firewall. Valid values:
	//
	// *   If the VPC firewall protects mutual access traffic between a VPC and a specified network instance that is attached to a Cloud Enterprise Network (CEN) instance, the value of this parameter is the ID of the CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance. You can call the [DescribeVpcFirewallCenList](~~345777~~) operation to query the IDs of CEN instances.
	// *   If the VPC firewall protects traffic between two VPCs that are connected by using an Express Connect circuit, the value of this parameter is the instance ID of the VPC firewall. You can call the [DescribeVpcFirewallList](~~342932~~) operation to query the instance IDs of VPC firewalls.
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s DescribeVpcFirewallDefaultIPSConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallDefaultIPSConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDefaultIPSConfigRequest) SetMemberUid(v string) *DescribeVpcFirewallDefaultIPSConfigRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeVpcFirewallDefaultIPSConfigRequest) SetVpcFirewallId(v string) *DescribeVpcFirewallDefaultIPSConfigRequest {
	s.VpcFirewallId = &v
	return s
}

type DescribeVpcFirewallDefaultIPSConfigResponseBody struct {
	// Indicates whether basic policies are enabled. Valid values:
	//
	// *   **1**: yes
	// *   **0**: no
	BasicRules *int32 `json:"BasicRules,omitempty" xml:"BasicRules,omitempty"`
	// Indicates whether virtual patching is enabled. Valid values:
	//
	// *   **1**: yes
	// *   **0**: no
	EnableAllPatch *int32 `json:"EnableAllPatch,omitempty" xml:"EnableAllPatch,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The mode of the intrusion prevention system (IPS). Valid values:
	//
	// *   **1**: block mode
	// *   **0**: monitor mode
	RunMode *int32 `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
}

func (s DescribeVpcFirewallDefaultIPSConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallDefaultIPSConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDefaultIPSConfigResponseBody) SetBasicRules(v int32) *DescribeVpcFirewallDefaultIPSConfigResponseBody {
	s.BasicRules = &v
	return s
}

func (s *DescribeVpcFirewallDefaultIPSConfigResponseBody) SetEnableAllPatch(v int32) *DescribeVpcFirewallDefaultIPSConfigResponseBody {
	s.EnableAllPatch = &v
	return s
}

func (s *DescribeVpcFirewallDefaultIPSConfigResponseBody) SetRequestId(v string) *DescribeVpcFirewallDefaultIPSConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallDefaultIPSConfigResponseBody) SetRunMode(v int32) *DescribeVpcFirewallDefaultIPSConfigResponseBody {
	s.RunMode = &v
	return s
}

type DescribeVpcFirewallDefaultIPSConfigResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcFirewallDefaultIPSConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcFirewallDefaultIPSConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallDefaultIPSConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDefaultIPSConfigResponse) SetHeaders(v map[string]*string) *DescribeVpcFirewallDefaultIPSConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcFirewallDefaultIPSConfigResponse) SetStatusCode(v int32) *DescribeVpcFirewallDefaultIPSConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcFirewallDefaultIPSConfigResponse) SetBody(v *DescribeVpcFirewallDefaultIPSConfigResponseBody) *DescribeVpcFirewallDefaultIPSConfigResponse {
	s.Body = v
	return s
}

type DescribeVpcFirewallDetailRequest struct {
	// The natural language of the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the local VPC.
	LocalVpcId *string `json:"LocalVpcId,omitempty" xml:"LocalVpcId,omitempty"`
	// The ID of the peer VPC.
	PeerVpcId *string `json:"PeerVpcId,omitempty" xml:"PeerVpcId,omitempty"`
	// The instance ID of the VPC firewall.
	//
	// >  You can call the [DescribeVpcFirewallList](~~342932~~) operation to query the instance IDs of VPC firewalls.
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s DescribeVpcFirewallDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDetailRequest) SetLang(v string) *DescribeVpcFirewallDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVpcFirewallDetailRequest) SetLocalVpcId(v string) *DescribeVpcFirewallDetailRequest {
	s.LocalVpcId = &v
	return s
}

func (s *DescribeVpcFirewallDetailRequest) SetPeerVpcId(v string) *DescribeVpcFirewallDetailRequest {
	s.PeerVpcId = &v
	return s
}

func (s *DescribeVpcFirewallDetailRequest) SetVpcFirewallId(v string) *DescribeVpcFirewallDetailRequest {
	s.VpcFirewallId = &v
	return s
}

type DescribeVpcFirewallDetailResponseBody struct {
	// The bandwidth of the Express Connect circuit. Unit: Mbit/s.
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The connection type of the VPC firewall. The value is fixed as **expressconnect**, which indicates Express Connect circuits.
	ConnectType *string `json:"ConnectType,omitempty" xml:"ConnectType,omitempty"`
	// The status of the VPC firewall. Valid values:
	//
	// *   **opened**: The VPC firewall is enabled.
	// *   **closed**: The VPC firewall is disabled.
	// *   **notconfigured**: The VPC firewall is not configured.
	// *   **configured**: The VPC firewall is configured.
	FirewallSwitchStatus *string `json:"FirewallSwitchStatus,omitempty" xml:"FirewallSwitchStatus,omitempty"`
	// The details about the local VPC.
	LocalVpc *DescribeVpcFirewallDetailResponseBodyLocalVpc `json:"LocalVpc,omitempty" xml:"LocalVpc,omitempty" type:"Struct"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The details about the peer VPC.
	PeerVpc *DescribeVpcFirewallDetailResponseBodyPeerVpc `json:"PeerVpc,omitempty" xml:"PeerVpc,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The instance ID of the VPC firewall.
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
	// The instance name of the VPC firewall.
	VpcFirewallName *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
}

func (s DescribeVpcFirewallDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDetailResponseBody) SetBandwidth(v int32) *DescribeVpcFirewallDetailResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBody) SetConnectType(v string) *DescribeVpcFirewallDetailResponseBody {
	s.ConnectType = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBody) SetFirewallSwitchStatus(v string) *DescribeVpcFirewallDetailResponseBody {
	s.FirewallSwitchStatus = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBody) SetLocalVpc(v *DescribeVpcFirewallDetailResponseBodyLocalVpc) *DescribeVpcFirewallDetailResponseBody {
	s.LocalVpc = v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBody) SetMemberUid(v string) *DescribeVpcFirewallDetailResponseBody {
	s.MemberUid = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBody) SetPeerVpc(v *DescribeVpcFirewallDetailResponseBodyPeerVpc) *DescribeVpcFirewallDetailResponseBody {
	s.PeerVpc = v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBody) SetRequestId(v string) *DescribeVpcFirewallDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBody) SetVpcFirewallId(v string) *DescribeVpcFirewallDetailResponseBody {
	s.VpcFirewallId = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBody) SetVpcFirewallName(v string) *DescribeVpcFirewallDetailResponseBody {
	s.VpcFirewallName = &v
	return s
}

type DescribeVpcFirewallDetailResponseBodyLocalVpc struct {
	// The ID of the ENI for the local VPC.
	EniId *string `json:"EniId,omitempty" xml:"EniId,omitempty"`
	// The private IP address of the elastic network interface (ENI) for the local VPC.
	EniPrivateIpAddress *string `json:"EniPrivateIpAddress,omitempty" xml:"EniPrivateIpAddress,omitempty"`
	// The region ID of the local VPC.
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The router interface ID of the local VPC.
	RouterInterfaceId *string `json:"RouterInterfaceId,omitempty" xml:"RouterInterfaceId,omitempty"`
	// The CIDR blocks of the local VPC.
	VpcCidrTableList []*DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableList `json:"VpcCidrTableList,omitempty" xml:"VpcCidrTableList,omitempty" type:"Repeated"`
	// The ID of the local VPC.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the local VPC.
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeVpcFirewallDetailResponseBodyLocalVpc) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallDetailResponseBodyLocalVpc) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpc) SetEniId(v string) *DescribeVpcFirewallDetailResponseBodyLocalVpc {
	s.EniId = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpc) SetEniPrivateIpAddress(v string) *DescribeVpcFirewallDetailResponseBodyLocalVpc {
	s.EniPrivateIpAddress = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpc) SetRegionNo(v string) *DescribeVpcFirewallDetailResponseBodyLocalVpc {
	s.RegionNo = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpc) SetRouterInterfaceId(v string) *DescribeVpcFirewallDetailResponseBodyLocalVpc {
	s.RouterInterfaceId = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpc) SetVpcCidrTableList(v []*DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableList) *DescribeVpcFirewallDetailResponseBodyLocalVpc {
	s.VpcCidrTableList = v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpc) SetVpcId(v string) *DescribeVpcFirewallDetailResponseBodyLocalVpc {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpc) SetVpcName(v string) *DescribeVpcFirewallDetailResponseBodyLocalVpc {
	s.VpcName = &v
	return s
}

type DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableList struct {
	// The route entries of the local VPC.
	RouteEntryList []*DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList `json:"RouteEntryList,omitempty" xml:"RouteEntryList,omitempty" type:"Repeated"`
	// The ID of the route table for the local VPC.
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableList) SetRouteEntryList(v []*DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList) *DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableList {
	s.RouteEntryList = v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableList) SetRouteTableId(v string) *DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableList {
	s.RouteTableId = &v
	return s
}

type DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList struct {
	// The destination CIDR block of the local VPC.
	DestinationCidr *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
	// The instance ID of the next hop for the local VPC.
	NextHopInstanceId *string `json:"NextHopInstanceId,omitempty" xml:"NextHopInstanceId,omitempty"`
}

func (s DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList) SetDestinationCidr(v string) *DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList {
	s.DestinationCidr = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList) SetNextHopInstanceId(v string) *DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList {
	s.NextHopInstanceId = &v
	return s
}

type DescribeVpcFirewallDetailResponseBodyPeerVpc struct {
	// The ID of the ENI for the peer VPC.
	EniId *string `json:"EniId,omitempty" xml:"EniId,omitempty"`
	// The private IP address of the ENI for the peer VPC.
	EniPrivateIpAddress *string `json:"EniPrivateIpAddress,omitempty" xml:"EniPrivateIpAddress,omitempty"`
	// The region ID of the peer VPC.
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The router interface ID of the peer VPC.
	RouterInterfaceId *string `json:"RouterInterfaceId,omitempty" xml:"RouterInterfaceId,omitempty"`
	// The CIDR blocks of the peer VPC.
	VpcCidrTableList []*DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableList `json:"VpcCidrTableList,omitempty" xml:"VpcCidrTableList,omitempty" type:"Repeated"`
	// The ID of the peer VPC.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the peer VPC.
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeVpcFirewallDetailResponseBodyPeerVpc) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallDetailResponseBodyPeerVpc) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpc) SetEniId(v string) *DescribeVpcFirewallDetailResponseBodyPeerVpc {
	s.EniId = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpc) SetEniPrivateIpAddress(v string) *DescribeVpcFirewallDetailResponseBodyPeerVpc {
	s.EniPrivateIpAddress = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpc) SetRegionNo(v string) *DescribeVpcFirewallDetailResponseBodyPeerVpc {
	s.RegionNo = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpc) SetRouterInterfaceId(v string) *DescribeVpcFirewallDetailResponseBodyPeerVpc {
	s.RouterInterfaceId = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpc) SetVpcCidrTableList(v []*DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableList) *DescribeVpcFirewallDetailResponseBodyPeerVpc {
	s.VpcCidrTableList = v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpc) SetVpcId(v string) *DescribeVpcFirewallDetailResponseBodyPeerVpc {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpc) SetVpcName(v string) *DescribeVpcFirewallDetailResponseBodyPeerVpc {
	s.VpcName = &v
	return s
}

type DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableList struct {
	// The route entries of the peer VPC.
	RouteEntryList []*DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableListRouteEntryList `json:"RouteEntryList,omitempty" xml:"RouteEntryList,omitempty" type:"Repeated"`
	// The ID of the route table for the peer VPC.
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableList) SetRouteEntryList(v []*DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableListRouteEntryList) *DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableList {
	s.RouteEntryList = v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableList) SetRouteTableId(v string) *DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableList {
	s.RouteTableId = &v
	return s
}

type DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableListRouteEntryList struct {
	// The destination CIDR block of the peer VPC.
	DestinationCidr *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
	// The instance ID of the next hop for the peer VPC.
	NextHopInstanceId *string `json:"NextHopInstanceId,omitempty" xml:"NextHopInstanceId,omitempty"`
}

func (s DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableListRouteEntryList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableListRouteEntryList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableListRouteEntryList) SetDestinationCidr(v string) *DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableListRouteEntryList {
	s.DestinationCidr = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableListRouteEntryList) SetNextHopInstanceId(v string) *DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableListRouteEntryList {
	s.NextHopInstanceId = &v
	return s
}

type DescribeVpcFirewallDetailResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcFirewallDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcFirewallDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDetailResponse) SetHeaders(v map[string]*string) *DescribeVpcFirewallDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcFirewallDetailResponse) SetStatusCode(v int32) *DescribeVpcFirewallDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponse) SetBody(v *DescribeVpcFirewallDetailResponseBody) *DescribeVpcFirewallDetailResponse {
	s.Body = v
	return s
}

type DescribeVpcFirewallIPSWhitelistRequest struct {
	// The language of the content within the request and response.
	//
	// Valid values:
	//
	// *   **zh** (default): Chinese
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member in Cloud Firewall.
	MemberUid *int64 `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The instance ID of the VPC firewall.
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s DescribeVpcFirewallIPSWhitelistRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallIPSWhitelistRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallIPSWhitelistRequest) SetLang(v string) *DescribeVpcFirewallIPSWhitelistRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVpcFirewallIPSWhitelistRequest) SetMemberUid(v int64) *DescribeVpcFirewallIPSWhitelistRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeVpcFirewallIPSWhitelistRequest) SetVpcFirewallId(v string) *DescribeVpcFirewallIPSWhitelistRequest {
	s.VpcFirewallId = &v
	return s
}

type DescribeVpcFirewallIPSWhitelistResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the IPS whitelist of the VPC firewall.
	Whitelists []*DescribeVpcFirewallIPSWhitelistResponseBodyWhitelists `json:"Whitelists,omitempty" xml:"Whitelists,omitempty" type:"Repeated"`
}

func (s DescribeVpcFirewallIPSWhitelistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallIPSWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallIPSWhitelistResponseBody) SetRequestId(v string) *DescribeVpcFirewallIPSWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallIPSWhitelistResponseBody) SetWhitelists(v []*DescribeVpcFirewallIPSWhitelistResponseBodyWhitelists) *DescribeVpcFirewallIPSWhitelistResponseBody {
	s.Whitelists = v
	return s
}

type DescribeVpcFirewallIPSWhitelistResponseBodyWhitelists struct {
	// The type of the list. Valid values:
	//
	// *   **1**: user-defined
	// *   **2**: address book
	ListType *int64 `json:"ListType,omitempty" xml:"ListType,omitempty"`
	// The entries in the list.
	ListValue *string `json:"ListValue,omitempty" xml:"ListValue,omitempty"`
	// The instance ID of the VPC firewall.
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
	// An array of entries in the list.
	WhiteListValue []*string `json:"WhiteListValue,omitempty" xml:"WhiteListValue,omitempty" type:"Repeated"`
	// The type of the whitelist. Valid values:
	//
	// *   **1**: destination
	// *   **2**: source
	WhiteType *int64 `json:"WhiteType,omitempty" xml:"WhiteType,omitempty"`
}

func (s DescribeVpcFirewallIPSWhitelistResponseBodyWhitelists) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallIPSWhitelistResponseBodyWhitelists) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallIPSWhitelistResponseBodyWhitelists) SetListType(v int64) *DescribeVpcFirewallIPSWhitelistResponseBodyWhitelists {
	s.ListType = &v
	return s
}

func (s *DescribeVpcFirewallIPSWhitelistResponseBodyWhitelists) SetListValue(v string) *DescribeVpcFirewallIPSWhitelistResponseBodyWhitelists {
	s.ListValue = &v
	return s
}

func (s *DescribeVpcFirewallIPSWhitelistResponseBodyWhitelists) SetVpcFirewallId(v string) *DescribeVpcFirewallIPSWhitelistResponseBodyWhitelists {
	s.VpcFirewallId = &v
	return s
}

func (s *DescribeVpcFirewallIPSWhitelistResponseBodyWhitelists) SetWhiteListValue(v []*string) *DescribeVpcFirewallIPSWhitelistResponseBodyWhitelists {
	s.WhiteListValue = v
	return s
}

func (s *DescribeVpcFirewallIPSWhitelistResponseBodyWhitelists) SetWhiteType(v int64) *DescribeVpcFirewallIPSWhitelistResponseBodyWhitelists {
	s.WhiteType = &v
	return s
}

type DescribeVpcFirewallIPSWhitelistResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcFirewallIPSWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcFirewallIPSWhitelistResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallIPSWhitelistResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallIPSWhitelistResponse) SetHeaders(v map[string]*string) *DescribeVpcFirewallIPSWhitelistResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcFirewallIPSWhitelistResponse) SetStatusCode(v int32) *DescribeVpcFirewallIPSWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcFirewallIPSWhitelistResponse) SetBody(v *DescribeVpcFirewallIPSWhitelistResponseBody) *DescribeVpcFirewallIPSWhitelistResponse {
	s.Body = v
	return s
}

type DescribeVpcFirewallListRequest struct {
	// The sub-type of the connection. Valid values:
	//
	// *   **vpc2vpc**: Express Connect connection
	// *   **vpcpeer**: peer connection
	ConnectSubType *string `json:"ConnectSubType,omitempty" xml:"ConnectSubType,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page **1**. Default value: **1**.
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The status of the VPC firewall. Valid values:
	//
	// *   **opened**: The VPC firewall is enabled.
	// *   **closed**: The VPC firewall is disabled.
	// *   **notconfigured**: The VPC firewall is not configured.
	// *   **configured**: The VPC firewall is configured.
	//
	// > If you do not specify this parameter, VPC firewalls in all states are queried.
	FirewallSwitchStatus *string `json:"FirewallSwitchStatus,omitempty" xml:"FirewallSwitchStatus,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The number of entries to return on each page.
	//
	// Default value: **10**. Maximum value: **50**.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The UID of the Alibaba Cloud account to which the peer VPC belongs.
	PeerUid *string `json:"PeerUid,omitempty" xml:"PeerUid,omitempty"`
	// The region ID of the VPC.
	//
	// > For more information about the regions, see [Supported regions](~~195657~~).
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The instance ID of the VPC firewall.
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
	// The instance name of the VPC firewall.
	VpcFirewallName *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
	// The ID of the VPC.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeVpcFirewallListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallListRequest) SetConnectSubType(v string) *DescribeVpcFirewallListRequest {
	s.ConnectSubType = &v
	return s
}

func (s *DescribeVpcFirewallListRequest) SetCurrentPage(v string) *DescribeVpcFirewallListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVpcFirewallListRequest) SetFirewallSwitchStatus(v string) *DescribeVpcFirewallListRequest {
	s.FirewallSwitchStatus = &v
	return s
}

func (s *DescribeVpcFirewallListRequest) SetLang(v string) *DescribeVpcFirewallListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVpcFirewallListRequest) SetMemberUid(v string) *DescribeVpcFirewallListRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeVpcFirewallListRequest) SetPageSize(v string) *DescribeVpcFirewallListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVpcFirewallListRequest) SetPeerUid(v string) *DescribeVpcFirewallListRequest {
	s.PeerUid = &v
	return s
}

func (s *DescribeVpcFirewallListRequest) SetRegionNo(v string) *DescribeVpcFirewallListRequest {
	s.RegionNo = &v
	return s
}

func (s *DescribeVpcFirewallListRequest) SetVpcFirewallId(v string) *DescribeVpcFirewallListRequest {
	s.VpcFirewallId = &v
	return s
}

func (s *DescribeVpcFirewallListRequest) SetVpcFirewallName(v string) *DescribeVpcFirewallListRequest {
	s.VpcFirewallName = &v
	return s
}

func (s *DescribeVpcFirewallListRequest) SetVpcId(v string) *DescribeVpcFirewallListRequest {
	s.VpcId = &v
	return s
}

type DescribeVpcFirewallListResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of VPC firewalls.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// An array that consists of the details about the VPC firewall.
	VpcFirewalls []*DescribeVpcFirewallListResponseBodyVpcFirewalls `json:"VpcFirewalls,omitempty" xml:"VpcFirewalls,omitempty" type:"Repeated"`
}

func (s DescribeVpcFirewallListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallListResponseBody) SetRequestId(v string) *DescribeVpcFirewallListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBody) SetTotalCount(v int32) *DescribeVpcFirewallListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBody) SetVpcFirewalls(v []*DescribeVpcFirewallListResponseBodyVpcFirewalls) *DescribeVpcFirewallListResponseBody {
	s.VpcFirewalls = v
	return s
}

type DescribeVpcFirewallListResponseBodyVpcFirewalls struct {
	// The bandwidth of the Express Connect circuit. Unit: Mbit/s.
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The sub-type of the connection. Valid values:
	//
	// *   **vpc2vpc**: Express Connect connection
	// *   **vpcpeer**: peer connection
	ConnectSubType *string `json:"ConnectSubType,omitempty" xml:"ConnectSubType,omitempty"`
	// The connection type of the VPC firewall. The value is fixed as **expressconnect**, which indicates an Express Connect connection.
	ConnectType *string `json:"ConnectType,omitempty" xml:"ConnectType,omitempty"`
	// The status of the VPC firewall. Valid values:
	//
	// *   **opened**: The VPC firewall is enabled.
	// *   **closed**: The VPC firewall is disabled.
	// *   **notconfigured**: The VPC firewall is not configured.
	FirewallSwitchStatus *string `json:"FirewallSwitchStatus,omitempty" xml:"FirewallSwitchStatus,omitempty"`
	// The information about the intrusion prevention system (IPS) configuration.
	IpsConfig *DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig `json:"IpsConfig,omitempty" xml:"IpsConfig,omitempty" type:"Struct"`
	// The details about the local VPC.
	LocalVpc *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc `json:"LocalVpc,omitempty" xml:"LocalVpc,omitempty" type:"Struct"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The details about the peer VPC.
	PeerVpc *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc `json:"PeerVpc,omitempty" xml:"PeerVpc,omitempty" type:"Struct"`
	// Indicates whether you can create a VPC firewall in a specified region. Valid values:
	//
	// *   **enable**: yes
	// *   **disable**: no
	RegionStatus *string `json:"RegionStatus,omitempty" xml:"RegionStatus,omitempty"`
	// The result code of the operation that creates the VPC firewall. Valid values:
	//
	// *   **Unauthorized**: Cloud Firewall is not authorized to access a VPC for which the VPC firewall is created, and the VPC firewall cannot be created.
	// *   **RegionDisable**: VPC Firewall is not supported in the region of a VPC for which the VPC firewall is created, and the VPC firewall cannot be created.
	// *   **Empty string**: You can create a VPC firewall for the network instance.
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// The instance ID of the VPC firewall.
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
	// The instance name of the VPC firewall.
	VpcFirewallName *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewalls) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewalls) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) SetBandwidth(v int32) *DescribeVpcFirewallListResponseBodyVpcFirewalls {
	s.Bandwidth = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) SetConnectSubType(v string) *DescribeVpcFirewallListResponseBodyVpcFirewalls {
	s.ConnectSubType = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) SetConnectType(v string) *DescribeVpcFirewallListResponseBodyVpcFirewalls {
	s.ConnectType = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) SetFirewallSwitchStatus(v string) *DescribeVpcFirewallListResponseBodyVpcFirewalls {
	s.FirewallSwitchStatus = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) SetIpsConfig(v *DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig) *DescribeVpcFirewallListResponseBodyVpcFirewalls {
	s.IpsConfig = v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) SetLocalVpc(v *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc) *DescribeVpcFirewallListResponseBodyVpcFirewalls {
	s.LocalVpc = v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) SetMemberUid(v string) *DescribeVpcFirewallListResponseBodyVpcFirewalls {
	s.MemberUid = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) SetPeerVpc(v *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc) *DescribeVpcFirewallListResponseBodyVpcFirewalls {
	s.PeerVpc = v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) SetRegionStatus(v string) *DescribeVpcFirewallListResponseBodyVpcFirewalls {
	s.RegionStatus = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) SetResultCode(v string) *DescribeVpcFirewallListResponseBodyVpcFirewalls {
	s.ResultCode = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) SetVpcFirewallId(v string) *DescribeVpcFirewallListResponseBodyVpcFirewalls {
	s.VpcFirewallId = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) SetVpcFirewallName(v string) *DescribeVpcFirewallListResponseBodyVpcFirewalls {
	s.VpcFirewallName = &v
	return s
}

type DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig struct {
	// Indicates whether basic protection is enabled. Valid values:
	//
	// *   **1**: yes
	// *   **0**: no
	BasicRules *int32 `json:"BasicRules,omitempty" xml:"BasicRules,omitempty"`
	// Indicates whether virtual patching is enabled. Valid values:
	//
	// *   **1**: yes
	// *   **0**: no
	EnableAllPatch *int32 `json:"EnableAllPatch,omitempty" xml:"EnableAllPatch,omitempty"`
	// The mode of the IPS. Valid values:
	//
	// *   **1**: block mode
	// *   **0**: monitor mode
	RunMode *int32 `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig) SetBasicRules(v int32) *DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig {
	s.BasicRules = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig) SetEnableAllPatch(v int32) *DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig {
	s.EnableAllPatch = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig) SetRunMode(v int32) *DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig {
	s.RunMode = &v
	return s
}

type DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc struct {
	// Indicates whether Cloud Firewall is authorized to access the local VPC. The value is fixed as authorized, which indicates that Cloud Firewall is authorized to access the local VPC.
	AuthorizationStatus *string `json:"AuthorizationStatus,omitempty" xml:"AuthorizationStatus,omitempty"`
	// The UID of the Alibaba Cloud account to which the local VPC belongs.
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the local VPC.
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// An array that consists of the CIDR blocks of the local VPC.
	VpcCidrTableList []*DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList `json:"VpcCidrTableList,omitempty" xml:"VpcCidrTableList,omitempty" type:"Repeated"`
	// The ID of the local VPC.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the local VPC.
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc) SetAuthorizationStatus(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc {
	s.AuthorizationStatus = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc) SetOwnerId(v int64) *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc {
	s.OwnerId = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc) SetRegionNo(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc {
	s.RegionNo = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc) SetVpcCidrTableList(v []*DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList) *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc {
	s.VpcCidrTableList = v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc) SetVpcId(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc) SetVpcName(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc {
	s.VpcName = &v
	return s
}

type DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList struct {
	// An array that consists of the route entries of the local VPC.
	RouteEntryList []*DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList `json:"RouteEntryList,omitempty" xml:"RouteEntryList,omitempty" type:"Repeated"`
	// The ID of the route table for the local VPC.
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList) SetRouteEntryList(v []*DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList) *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList {
	s.RouteEntryList = v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList) SetRouteTableId(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList {
	s.RouteTableId = &v
	return s
}

type DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList struct {
	// The destination CIDR block of the local VPC.
	DestinationCidr *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
	// The instance ID of the next hop for the local VPC.
	NextHopInstanceId *string `json:"NextHopInstanceId,omitempty" xml:"NextHopInstanceId,omitempty"`
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList) SetDestinationCidr(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList {
	s.DestinationCidr = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList) SetNextHopInstanceId(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList {
	s.NextHopInstanceId = &v
	return s
}

type DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc struct {
	// Indicates whether Cloud Firewall is authorized to access the peer VPC. The value is fixed as **authorized**, which indicates that Cloud Firewall is authorized to access the peer VPC.
	AuthorizationStatus *string `json:"AuthorizationStatus,omitempty" xml:"AuthorizationStatus,omitempty"`
	// The UID of the Alibaba Cloud account to which the peer VPC belongs.
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the peer VPC.
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// An array that consists of the CIDR blocks of the peer VPC.
	VpcCidrTableList []*DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableList `json:"VpcCidrTableList,omitempty" xml:"VpcCidrTableList,omitempty" type:"Repeated"`
	// The ID of the peer VPC.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the peer VPC.
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc) SetAuthorizationStatus(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc {
	s.AuthorizationStatus = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc) SetOwnerId(v int64) *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc {
	s.OwnerId = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc) SetRegionNo(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc {
	s.RegionNo = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc) SetVpcCidrTableList(v []*DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableList) *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc {
	s.VpcCidrTableList = v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc) SetVpcId(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc) SetVpcName(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc {
	s.VpcName = &v
	return s
}

type DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableList struct {
	// An array that consists of the route entries of the peer VPC.
	RouteEntryList []*DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableListRouteEntryList `json:"RouteEntryList,omitempty" xml:"RouteEntryList,omitempty" type:"Repeated"`
	// The ID of the route table for the peer VPC.
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableList) SetRouteEntryList(v []*DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableListRouteEntryList) *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableList {
	s.RouteEntryList = v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableList) SetRouteTableId(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableList {
	s.RouteTableId = &v
	return s
}

type DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableListRouteEntryList struct {
	// The destination CIDR block of the peer VPC.
	DestinationCidr *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
	// The instance ID of the next hop for the peer VPC.
	NextHopInstanceId *string `json:"NextHopInstanceId,omitempty" xml:"NextHopInstanceId,omitempty"`
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableListRouteEntryList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableListRouteEntryList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableListRouteEntryList) SetDestinationCidr(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableListRouteEntryList {
	s.DestinationCidr = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableListRouteEntryList) SetNextHopInstanceId(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableListRouteEntryList {
	s.NextHopInstanceId = &v
	return s
}

type DescribeVpcFirewallListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcFirewallListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcFirewallListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallListResponse) SetHeaders(v map[string]*string) *DescribeVpcFirewallListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcFirewallListResponse) SetStatusCode(v int32) *DescribeVpcFirewallListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcFirewallListResponse) SetBody(v *DescribeVpcFirewallListResponseBody) *DescribeVpcFirewallListResponse {
	s.Body = v
	return s
}

type DescribeVpcFirewallPolicyPriorUsedRequest struct {
	// The natural language of the request and response.
	//
	// Valid values:
	//
	// - **zh**: Chinese (default)
	// - **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the policy group to which the access control policy belongs. You can call the DescribeVpcFirewallAclGroupList operation to query the ID.
	//
	// Valid values:
	//
	// - If the VPC firewall is used to protect a Cloud Enterprise Network (CEN) instance, the value of this parameter is the ID of the CEN instance.
	//
	// Example: cen-ervw0g12b5jbw****
	// - If the VPC firewall is used to protect an Express Connect circuit, the value of this parameter is the ID of the VPC firewall instance.
	//
	// Example: vfw-a42bbb7b887148c9****
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s DescribeVpcFirewallPolicyPriorUsedRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallPolicyPriorUsedRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallPolicyPriorUsedRequest) SetLang(v string) *DescribeVpcFirewallPolicyPriorUsedRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVpcFirewallPolicyPriorUsedRequest) SetVpcFirewallId(v string) *DescribeVpcFirewallPolicyPriorUsedRequest {
	s.VpcFirewallId = &v
	return s
}

type DescribeVpcFirewallPolicyPriorUsedResponseBody struct {
	// The lowest priority for the access control policy.
	End *int32 `json:"End,omitempty" xml:"End,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The highest priority for the access control policy.
	Start *int32 `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s DescribeVpcFirewallPolicyPriorUsedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallPolicyPriorUsedResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallPolicyPriorUsedResponseBody) SetEnd(v int32) *DescribeVpcFirewallPolicyPriorUsedResponseBody {
	s.End = &v
	return s
}

func (s *DescribeVpcFirewallPolicyPriorUsedResponseBody) SetRequestId(v string) *DescribeVpcFirewallPolicyPriorUsedResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallPolicyPriorUsedResponseBody) SetStart(v int32) *DescribeVpcFirewallPolicyPriorUsedResponseBody {
	s.Start = &v
	return s
}

type DescribeVpcFirewallPolicyPriorUsedResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcFirewallPolicyPriorUsedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcFirewallPolicyPriorUsedResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallPolicyPriorUsedResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallPolicyPriorUsedResponse) SetHeaders(v map[string]*string) *DescribeVpcFirewallPolicyPriorUsedResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcFirewallPolicyPriorUsedResponse) SetStatusCode(v int32) *DescribeVpcFirewallPolicyPriorUsedResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcFirewallPolicyPriorUsedResponse) SetBody(v *DescribeVpcFirewallPolicyPriorUsedResponseBody) *DescribeVpcFirewallPolicyPriorUsedResponse {
	s.Body = v
	return s
}

type DescribeVpcListLiteRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// *   **zh** (default): Chinese
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region ID of the VPC.
	//
	// >  For more information about Cloud Firewall supported regions, see [Supported regions](~~195657~~).
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The source IP address of the request.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The ID of the VPC.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the VPC.
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeVpcListLiteRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcListLiteRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcListLiteRequest) SetLang(v string) *DescribeVpcListLiteRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVpcListLiteRequest) SetRegionNo(v string) *DescribeVpcListLiteRequest {
	s.RegionNo = &v
	return s
}

func (s *DescribeVpcListLiteRequest) SetSourceIp(v string) *DescribeVpcListLiteRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeVpcListLiteRequest) SetVpcId(v string) *DescribeVpcListLiteRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcListLiteRequest) SetVpcName(v string) *DescribeVpcListLiteRequest {
	s.VpcName = &v
	return s
}

type DescribeVpcListLiteResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the VPCs.
	VpcList []*DescribeVpcListLiteResponseBodyVpcList `json:"VpcList,omitempty" xml:"VpcList,omitempty" type:"Repeated"`
}

func (s DescribeVpcListLiteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcListLiteResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcListLiteResponseBody) SetRequestId(v string) *DescribeVpcListLiteResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcListLiteResponseBody) SetVpcList(v []*DescribeVpcListLiteResponseBodyVpcList) *DescribeVpcListLiteResponseBody {
	s.VpcList = v
	return s
}

type DescribeVpcListLiteResponseBodyVpcList struct {
	// The region ID of the VPC.
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The ID of the VPC.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the VPC.
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeVpcListLiteResponseBodyVpcList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcListLiteResponseBodyVpcList) GoString() string {
	return s.String()
}

func (s *DescribeVpcListLiteResponseBodyVpcList) SetRegionNo(v string) *DescribeVpcListLiteResponseBodyVpcList {
	s.RegionNo = &v
	return s
}

func (s *DescribeVpcListLiteResponseBodyVpcList) SetVpcId(v string) *DescribeVpcListLiteResponseBodyVpcList {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcListLiteResponseBodyVpcList) SetVpcName(v string) *DescribeVpcListLiteResponseBodyVpcList {
	s.VpcName = &v
	return s
}

type DescribeVpcListLiteResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcListLiteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcListLiteResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcListLiteResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcListLiteResponse) SetHeaders(v map[string]*string) *DescribeVpcListLiteResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcListLiteResponse) SetStatusCode(v int32) *DescribeVpcListLiteResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcListLiteResponse) SetBody(v *DescribeVpcListLiteResponseBody) *DescribeVpcListLiteResponse {
	s.Body = v
	return s
}

type DescribeVpcZoneRequest struct {
	// The environment. Valid values:
	//
	// *   **VPC**
	// *   **TransitRouter**
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// *   **zh** (default): Chinese
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member in Cloud Firewall.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The region ID.
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
}

func (s DescribeVpcZoneRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcZoneRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcZoneRequest) SetEnvironment(v string) *DescribeVpcZoneRequest {
	s.Environment = &v
	return s
}

func (s *DescribeVpcZoneRequest) SetLang(v string) *DescribeVpcZoneRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVpcZoneRequest) SetMemberUid(v string) *DescribeVpcZoneRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeVpcZoneRequest) SetRegionNo(v string) *DescribeVpcZoneRequest {
	s.RegionNo = &v
	return s
}

type DescribeVpcZoneResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The zones.
	ZoneList []*DescribeVpcZoneResponseBodyZoneList `json:"ZoneList,omitempty" xml:"ZoneList,omitempty" type:"Repeated"`
}

func (s DescribeVpcZoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcZoneResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcZoneResponseBody) SetRequestId(v string) *DescribeVpcZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcZoneResponseBody) SetZoneList(v []*DescribeVpcZoneResponseBodyZoneList) *DescribeVpcZoneResponseBody {
	s.ZoneList = v
	return s
}

type DescribeVpcZoneResponseBodyZoneList struct {
	// The name of the zone.
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The zone ID.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The zone type. Default value: AvailabilityZone. This value indicates Alibaba Cloud zones.
	ZoneType *string `json:"ZoneType,omitempty" xml:"ZoneType,omitempty"`
}

func (s DescribeVpcZoneResponseBodyZoneList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcZoneResponseBodyZoneList) GoString() string {
	return s.String()
}

func (s *DescribeVpcZoneResponseBodyZoneList) SetLocalName(v string) *DescribeVpcZoneResponseBodyZoneList {
	s.LocalName = &v
	return s
}

func (s *DescribeVpcZoneResponseBodyZoneList) SetZoneId(v string) *DescribeVpcZoneResponseBodyZoneList {
	s.ZoneId = &v
	return s
}

func (s *DescribeVpcZoneResponseBodyZoneList) SetZoneType(v string) *DescribeVpcZoneResponseBodyZoneList {
	s.ZoneType = &v
	return s
}

type DescribeVpcZoneResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpcZoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpcZoneResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcZoneResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcZoneResponse) SetHeaders(v map[string]*string) *DescribeVpcZoneResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcZoneResponse) SetStatusCode(v int32) *DescribeVpcZoneResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpcZoneResponse) SetBody(v *DescribeVpcZoneResponseBody) *DescribeVpcZoneResponse {
	s.Body = v
	return s
}

type DescribeVulnerabilityProtectedListRequest struct {
	// The attack type of the vulnerability prevention event. Valid values:
	//
	// *   **1**: suspicious connection
	// *   **2**: command execution
	// *   **3**: brute-force attack
	// *   **4**: scanning
	// *   **5**: others
	// *   **6**: information leakage
	// *   **7**: DoS attack
	// *   **8**: buffer overflow attack
	// *   **9**: web attack
	// *   **10**: webshell
	// *   **11**: computer worm
	// *   **12**: mining
	// *   **13**: reverse shell
	//
	// >  If you do not specify this parameter, the intrusion events of all attack types are queried.
	AttackType *string `json:"AttackType,omitempty" xml:"AttackType,omitempty"`
	// The edition of Cloud Firewall. If you use Cloud Firewall that uses the pay-as-you-go billing method, set the value to 10. You do not need to specify this parameter for other editions.
	BuyVersion *int64 `json:"BuyVersion,omitempty" xml:"BuyVersion,omitempty"`
	// The number of the page to return. Default value: 1.
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The order in which you want to sort the queried information. Valid values:
	//
	// *   **asc**: the ascending order.
	// *   **desc**: the descending order. This is the default value.
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The number of entries to return on each page. Maximum value: 50.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sorting basis. Set the value to **attackCnt**, which indicates the number of attacks.
	SortKey *string `json:"SortKey,omitempty" xml:"SortKey,omitempty"`
	// Deprecated
	// The IP address of the access source.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The type of the user. Set the value to **buy**, which indicates user of a paid edition of Cloud Firewall.
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
	// The Common Vulnerabilities and Exposures (CVE) ID of the vulnerability.
	VulnCveName *string `json:"VulnCveName,omitempty" xml:"VulnCveName,omitempty"`
	// The risk level of the vulnerability. Valid values:
	//
	// *   **high**
	// *   **medium**
	// *   **low**
	VulnLevel *string `json:"VulnLevel,omitempty" xml:"VulnLevel,omitempty"`
	// The number of assets that are affected by the vulnerability.
	VulnResource *string `json:"VulnResource,omitempty" xml:"VulnResource,omitempty"`
	// The status of vulnerability protection. Valid values:
	//
	// *   **partProtected**: partially protected
	// *   **protected**: protected
	// *   **unProtected**: unprotected
	VulnStatus *string `json:"VulnStatus,omitempty" xml:"VulnStatus,omitempty"`
	// The type of the vulnerability. Valid values:
	//
	// *   **App**: application vulnerability
	// *   **emg**: urgent vulnerability
	// *   **cms**: Web-CMS vulnerability
	VulnType *string `json:"VulnType,omitempty" xml:"VulnType,omitempty"`
}

func (s DescribeVulnerabilityProtectedListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulnerabilityProtectedListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVulnerabilityProtectedListRequest) SetAttackType(v string) *DescribeVulnerabilityProtectedListRequest {
	s.AttackType = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListRequest) SetBuyVersion(v int64) *DescribeVulnerabilityProtectedListRequest {
	s.BuyVersion = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListRequest) SetCurrentPage(v string) *DescribeVulnerabilityProtectedListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListRequest) SetEndTime(v string) *DescribeVulnerabilityProtectedListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListRequest) SetLang(v string) *DescribeVulnerabilityProtectedListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListRequest) SetMemberUid(v string) *DescribeVulnerabilityProtectedListRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListRequest) SetOrder(v string) *DescribeVulnerabilityProtectedListRequest {
	s.Order = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListRequest) SetPageSize(v string) *DescribeVulnerabilityProtectedListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListRequest) SetSortKey(v string) *DescribeVulnerabilityProtectedListRequest {
	s.SortKey = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListRequest) SetSourceIp(v string) *DescribeVulnerabilityProtectedListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListRequest) SetStartTime(v string) *DescribeVulnerabilityProtectedListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListRequest) SetUserType(v string) *DescribeVulnerabilityProtectedListRequest {
	s.UserType = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListRequest) SetVulnCveName(v string) *DescribeVulnerabilityProtectedListRequest {
	s.VulnCveName = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListRequest) SetVulnLevel(v string) *DescribeVulnerabilityProtectedListRequest {
	s.VulnLevel = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListRequest) SetVulnResource(v string) *DescribeVulnerabilityProtectedListRequest {
	s.VulnResource = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListRequest) SetVulnStatus(v string) *DescribeVulnerabilityProtectedListRequest {
	s.VulnStatus = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListRequest) SetVulnType(v string) *DescribeVulnerabilityProtectedListRequest {
	s.VulnType = &v
	return s
}

type DescribeVulnerabilityProtectedListResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of vulnerabilities that are detected by Cloud Firewall.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The vulnerabilities.
	VulnList []*DescribeVulnerabilityProtectedListResponseBodyVulnList `json:"VulnList,omitempty" xml:"VulnList,omitempty" type:"Repeated"`
	// The number of assets on which no vulnerabilities are detected.
	ZeroResourceCount *int32 `json:"ZeroResourceCount,omitempty" xml:"ZeroResourceCount,omitempty"`
}

func (s DescribeVulnerabilityProtectedListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulnerabilityProtectedListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVulnerabilityProtectedListResponseBody) SetRequestId(v string) *DescribeVulnerabilityProtectedListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBody) SetTotalCount(v int32) *DescribeVulnerabilityProtectedListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBody) SetVulnList(v []*DescribeVulnerabilityProtectedListResponseBodyVulnList) *DescribeVulnerabilityProtectedListResponseBody {
	s.VulnList = v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBody) SetZeroResourceCount(v int32) *DescribeVulnerabilityProtectedListResponseBody {
	s.ZeroResourceCount = &v
	return s
}

type DescribeVulnerabilityProtectedListResponseBodyVulnList struct {
	// The number of vulnerability attacks.
	AttackCnt *int32 `json:"AttackCnt,omitempty" xml:"AttackCnt,omitempty"`
	// The attack type of the vulnerability prevention event. Valid values:
	//
	// *   **1**: suspicious connection
	// *   **2**: command execution
	// *   **3**: brute-force attack
	// *   **4**: scanning
	// *   **5**: others
	// *   **6**: information leakage
	// *   **7**: DoS attack
	// *   **8**: buffer overflow attack
	// *   **9**: web attack
	// *   **10**: webshell
	// *   **11**: computer worm
	// *   **12**: mining
	// *   **13**: reverse shell
	AttackType *int32 `json:"AttackType,omitempty" xml:"AttackType,omitempty"`
	// The IDs of associated basic protection policies.
	BasicRuleIds *string `json:"BasicRuleIds,omitempty" xml:"BasicRuleIds,omitempty"`
	// The CVE IDs.
	CveId *string `json:"CveId,omitempty" xml:"CveId,omitempty"`
	// The time when the first attack was launched.
	FirstTime *int64 `json:"FirstTime,omitempty" xml:"FirstTime,omitempty"`
	// Indicates whether you need to pay special attention to the vulnerability. Valid values:
	//
	// *   **0**: no
	// *   **1**: yes
	HighlightTag *int32 `json:"HighlightTag,omitempty" xml:"HighlightTag,omitempty"`
	// The time when the last attack was launched.
	LastTime *int64 `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// Indicates whether the basic protection policy that related to the vulnerability is enabled. Valid values:
	//
	// *   **true**
	// *   **false**
	//
	// >  If the value of this parameter is true, you must set the action of the basic protection policy related to the vulnerability to Block.
	NeedOpenBasicRule *bool `json:"NeedOpenBasicRule,omitempty" xml:"NeedOpenBasicRule,omitempty"`
	// The UUIDs of the basic protection policies for which the action needs to be changed to Block.
	NeedOpenBasicRuleUuids *string `json:"NeedOpenBasicRuleUuids,omitempty" xml:"NeedOpenBasicRuleUuids,omitempty"`
	// Indicates whether Threat Engine Mode needs to be configured when you enable protection. Valid values:
	//
	// *   **true**
	// *   **false**
	NeedOpenRunMode *bool `json:"NeedOpenRunMode,omitempty" xml:"NeedOpenRunMode,omitempty"`
	// Indicates whether the virtual patching policy that related to the vulnerability is enabled. Valid values:
	//
	// *   **true**
	// *   **false**
	//
	// >  If the value of this parameter is true, you must set the action of the virtual patching policy that related to the vulnerability to Block.
	NeedOpenVirtualPatche *bool `json:"NeedOpenVirtualPatche,omitempty" xml:"NeedOpenVirtualPatche,omitempty"`
	// The UUIDs of the virtual patching policies for which the action needs to be changed to Block.
	NeedOpenVirtualPatcheUuids *string `json:"NeedOpenVirtualPatcheUuids,omitempty" xml:"NeedOpenVirtualPatcheUuids,omitempty"`
	// The type of the rule group. Valid values:
	//
	// *   **1** (default): loose
	// *   **2**: medium
	// *   **3**: strict
	NeedRuleClass *int32 `json:"NeedRuleClass,omitempty" xml:"NeedRuleClass,omitempty"`
	// The number of assets on which vulnerabilities are detected.
	ResourceCnt *int32 `json:"ResourceCnt,omitempty" xml:"ResourceCnt,omitempty"`
	// The assets on which the vulnerability is detected.
	ResourceList []*DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Repeated"`
	// The IDs of associated virtual patching policies.
	VirtualPatcheIds *string `json:"VirtualPatcheIds,omitempty" xml:"VirtualPatcheIds,omitempty"`
	// The code of the vulnerability.
	VulnKey *string `json:"VulnKey,omitempty" xml:"VulnKey,omitempty"`
	// The risk level of the vulnerability. Valid values:
	//
	// *   **high**
	// *   **medium**
	// *   **low**
	VulnLevel *string `json:"VulnLevel,omitempty" xml:"VulnLevel,omitempty"`
	// The name of the vulnerability.
	VulnName *string `json:"VulnName,omitempty" xml:"VulnName,omitempty"`
	// The status of the vulnerability prevention feature. Valid values:
	//
	// *   **partProtected**: enabled for partial assets
	// *   **protected**: enabled
	// *   **unProtected**: disabled
	VulnStatus *string `json:"VulnStatus,omitempty" xml:"VulnStatus,omitempty"`
	// The type of the vulnerability. Valid values:
	//
	// *   **emg**: urgent vulnerability
	// *   **webcms**: Web-CMS vulnerability
	// *   **app**: application vulnerability
	VulnType *string `json:"VulnType,omitempty" xml:"VulnType,omitempty"`
}

func (s DescribeVulnerabilityProtectedListResponseBodyVulnList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulnerabilityProtectedListResponseBodyVulnList) GoString() string {
	return s.String()
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetAttackCnt(v int32) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.AttackCnt = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetAttackType(v int32) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.AttackType = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetBasicRuleIds(v string) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.BasicRuleIds = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetCveId(v string) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.CveId = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetFirstTime(v int64) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.FirstTime = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetHighlightTag(v int32) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.HighlightTag = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetLastTime(v int64) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.LastTime = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetMemberUid(v string) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.MemberUid = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetNeedOpenBasicRule(v bool) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.NeedOpenBasicRule = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetNeedOpenBasicRuleUuids(v string) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.NeedOpenBasicRuleUuids = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetNeedOpenRunMode(v bool) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.NeedOpenRunMode = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetNeedOpenVirtualPatche(v bool) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.NeedOpenVirtualPatche = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetNeedOpenVirtualPatcheUuids(v string) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.NeedOpenVirtualPatcheUuids = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetNeedRuleClass(v int32) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.NeedRuleClass = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetResourceCnt(v int32) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.ResourceCnt = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetResourceList(v []*DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.ResourceList = v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetVirtualPatcheIds(v string) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.VirtualPatcheIds = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetVulnKey(v string) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.VulnKey = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetVulnLevel(v string) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.VulnLevel = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetVulnName(v string) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.VulnName = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetVulnStatus(v string) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.VulnStatus = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnList) SetVulnType(v string) *DescribeVulnerabilityProtectedListResponseBodyVulnList {
	s.VulnType = &v
	return s
}

type DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList struct {
	// The elastic IP address (EIP) that is associated with the instance.
	Eip *string `json:"Eip,omitempty" xml:"Eip,omitempty"`
	// The public IP address of the instance.
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the instance.
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The region ID of your Cloud Firewall.
	//
	// >  For more information about Cloud Firewall supported regions, see [Supported regions](~~195657~~).
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the instance.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the instance.
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The type of the asset. Valid values:
	//
	// *   **SLB**
	// *   **EIP**
	// *   **ECS**
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The status of the vulnerability prevention feature. Valid values:
	//
	// *   **partProtected**: enabled for partial assets
	// *   **protected**: enabled
	// *   **unProtected**: disabled
	VulnStatus *string `json:"VulnStatus,omitempty" xml:"VulnStatus,omitempty"`
}

func (s DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList) GoString() string {
	return s.String()
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList) SetEip(v string) *DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList {
	s.Eip = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList) SetInternetIp(v string) *DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList {
	s.InternetIp = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList) SetIntranetIp(v string) *DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList {
	s.IntranetIp = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList) SetRegionId(v string) *DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList {
	s.RegionId = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList) SetResourceId(v string) *DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList {
	s.ResourceId = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList) SetResourceName(v string) *DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList {
	s.ResourceName = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList) SetResourceType(v string) *DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList {
	s.ResourceType = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList) SetVulnStatus(v string) *DescribeVulnerabilityProtectedListResponseBodyVulnListResourceList {
	s.VulnStatus = &v
	return s
}

type DescribeVulnerabilityProtectedListResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVulnerabilityProtectedListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVulnerabilityProtectedListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVulnerabilityProtectedListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVulnerabilityProtectedListResponse) SetHeaders(v map[string]*string) *DescribeVulnerabilityProtectedListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponse) SetStatusCode(v int32) *DescribeVulnerabilityProtectedListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVulnerabilityProtectedListResponse) SetBody(v *DescribeVulnerabilityProtectedListResponseBody) *DescribeVulnerabilityProtectedListResponse {
	s.Body = v
	return s
}

type ModifyAddressBookRequest struct {
	// The addresses in the address book. Separate multiple addresses with commas (,). If you set GroupType to **ip**, **port**, or **domain**, you must specify this parameter.
	//
	// *   If you set GroupType to **ip**, you must specify IP addresses for the address book. Example: 1.2.XX.XX/32,1.2.XX.XX/24.
	// *   If you set GroupType to **port**, you must specify port numbers or port ranges for the address book. Example: 80/80,100/200.
	// *   If you set GroupType to **domain**, you must specify domain names for the address book. Example: demo1.aliyun.com,demo2.aliyun.com.
	AddressList *string `json:"AddressList,omitempty" xml:"AddressList,omitempty"`
	// Specifies whether to automatically add public IP addresses of Elastic Compute Service (ECS) instances to the address book if the instances match the specified tags. Valid values:
	//
	// *   **1**: yes
	// *   **0**: no
	AutoAddTagEcs *string `json:"AutoAddTagEcs,omitempty" xml:"AutoAddTagEcs,omitempty"`
	// The description of the address book.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the address book.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ID of the address book.
	//
	// >  To modify the address book, you must provide the ID of the address book. You can call the [DescribeAddressBook](~~138869~~) operation to query the ID.
	GroupUuid *string `json:"GroupUuid,omitempty" xml:"GroupUuid,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Deprecated
	// The source IP address of the request.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The ECS tags that you want to match.
	TagList []*ModifyAddressBookRequestTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The logical relationship among ECS tags. Valid values:
	//
	// *   **and**: Only the public IP addresses of ECS instances that match all the specified tags can be added to the address book.
	// *   **or**: The public IP addresses of ECS instances that match one of the specified tags can be added to the address book.
	TagRelation *string `json:"TagRelation,omitempty" xml:"TagRelation,omitempty"`
}

func (s ModifyAddressBookRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAddressBookRequest) GoString() string {
	return s.String()
}

func (s *ModifyAddressBookRequest) SetAddressList(v string) *ModifyAddressBookRequest {
	s.AddressList = &v
	return s
}

func (s *ModifyAddressBookRequest) SetAutoAddTagEcs(v string) *ModifyAddressBookRequest {
	s.AutoAddTagEcs = &v
	return s
}

func (s *ModifyAddressBookRequest) SetDescription(v string) *ModifyAddressBookRequest {
	s.Description = &v
	return s
}

func (s *ModifyAddressBookRequest) SetGroupName(v string) *ModifyAddressBookRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyAddressBookRequest) SetGroupUuid(v string) *ModifyAddressBookRequest {
	s.GroupUuid = &v
	return s
}

func (s *ModifyAddressBookRequest) SetLang(v string) *ModifyAddressBookRequest {
	s.Lang = &v
	return s
}

func (s *ModifyAddressBookRequest) SetSourceIp(v string) *ModifyAddressBookRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyAddressBookRequest) SetTagList(v []*ModifyAddressBookRequestTagList) *ModifyAddressBookRequest {
	s.TagList = v
	return s
}

func (s *ModifyAddressBookRequest) SetTagRelation(v string) *ModifyAddressBookRequest {
	s.TagRelation = &v
	return s
}

type ModifyAddressBookRequestTagList struct {
	// The key of ECS tag N that you want to match.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of ECS tag N that you want to match.
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ModifyAddressBookRequestTagList) String() string {
	return tea.Prettify(s)
}

func (s ModifyAddressBookRequestTagList) GoString() string {
	return s.String()
}

func (s *ModifyAddressBookRequestTagList) SetTagKey(v string) *ModifyAddressBookRequestTagList {
	s.TagKey = &v
	return s
}

func (s *ModifyAddressBookRequestTagList) SetTagValue(v string) *ModifyAddressBookRequestTagList {
	s.TagValue = &v
	return s
}

type ModifyAddressBookResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAddressBookResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAddressBookResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAddressBookResponseBody) SetRequestId(v string) *ModifyAddressBookResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAddressBookResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAddressBookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAddressBookResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAddressBookResponse) GoString() string {
	return s.String()
}

func (s *ModifyAddressBookResponse) SetHeaders(v map[string]*string) *ModifyAddressBookResponse {
	s.Headers = v
	return s
}

func (s *ModifyAddressBookResponse) SetStatusCode(v int32) *ModifyAddressBookResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAddressBookResponse) SetBody(v *ModifyAddressBookResponseBody) *ModifyAddressBookResponse {
	s.Body = v
	return s
}

type ModifyControlPolicyRequest struct {
	// The action that Cloud Firewall performs on the traffic. Valid values:
	//
	// *   **accept**: allows the traffic.
	// *   **drop**: denies the traffic.
	// *   **log**: monitors the traffic.
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The UUID of the access control policy.
	//
	// >  To modify an access control policy, you must specify the UUID of the policy. You can call the [DescribeControlPolicy](~~138866~~) interface to query the UUID.
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The type of the application that the access control policy supports. Valid values:
	//
	// *   **ANY**
	// *   **HTTP**
	// *   **HTTPS**
	// *   **MySQL**
	// *   **SMTP**
	// *   **SMTPS**
	// *   **RDP**
	// *   **VNC**
	// *   **SSH**
	// *   **Redis**
	// *   **MQTT**
	// *   **MongoDB**
	// *   **Memcache**
	// *   **SSL**
	//
	// >  The value *ANY* indicates all types of applications.
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The application names.
	ApplicationNameList []*string `json:"ApplicationNameList,omitempty" xml:"ApplicationNameList,omitempty" type:"Repeated"`
	// The description of the access control policy.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination port in the access control policy.
	DestPort *string `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	// The name of the destination port address book in the access control policy.
	DestPortGroup *string `json:"DestPortGroup,omitempty" xml:"DestPortGroup,omitempty"`
	// The type of the destination port in the access control policy. Valid values:
	//
	// *   **port**: port
	// *   **group**: port address book
	DestPortType *string `json:"DestPortType,omitempty" xml:"DestPortType,omitempty"`
	// The destination address in the access control policy.
	//
	// *   If **DestinationType** is set to net, the value of **Destination** is a CIDR block. Example: 1.2.XX.XX/24.
	// *   If **DestinationType** is set to group, the value of **Destination** is an address book. Example: db_group.
	// *   If **DestinationType** is set to domain, the value of **Destination** is a domain name. Example: \*.aliyuncs.com.
	// *   If **DestinationType** is set to location, the value of **Destination** is a location. For more information about the location codes, see the "AddControlPolicy" topic. Example: \["BJ11", "ZB"].
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The type of the destination address in the access control policy. Valid values:
	//
	// *   **net**: CIDR block
	// *   **group**: address book
	// *   **domain**: domain name
	// *   **location**: location
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The direction of the traffic to which the access control policy applies. Valid values:
	//
	// *   **in**: inbound traffic
	// *   **out**: outbound traffic
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The time when the access control policy stops taking effect. The value is a UNIX timestamp. Unit: seconds. The value must be on the hour or on the half hour, and at least 30 minutes later than the value of StartTime.
	//
	// >  If you set RepeatType to Permanent, leave this parameter empty. If you set RepeatType to None, Daily, Weekly, or Monthly, you must specify this parameter.
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the content within the request and the response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The protocol type supported by the access control policy. Valid values:
	//
	// *   **ANY**
	// *   **TCP**
	// *   **UDP**
	// *   **ICMP**
	//
	// >  The value *ANY* indicates all types of applications.
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// The status of the access control policy. Valid values:
	//
	// *   true: enabled
	// *   false: disabled
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The days of a week or of a month on which the access control policy takes effect.
	//
	// *   If you set RepeatType to `Permanent`, `None`, or `Daily`, the value of this parameter is an empty array. Example: \[]
	// *   If you set RepeatType to Weekly, you must specify this parameter. Example: \[0, 6]
	//
	// >  If you set RepeatType to Weekly, the fields in the value of this parameter cannot be repeated.
	//
	// *   If you set RepeatType to `Monthly`, you must specify this parameter. Example: \[1, 31]
	//
	// >  If you set RepeatType to Monthly, the fields in the value of this parameter cannot be repeated.
	RepeatDays []*int64 `json:"RepeatDays,omitempty" xml:"RepeatDays,omitempty" type:"Repeated"`
	// The point in time when the recurrence ends. Example: 23:30. The value must be on the hour or on the half hour, and at least 30 minutes later than the value of RepeatStartTime.
	//
	// >  If you set RepeatType to Permanent or None, leave this parameter empty. If you set RepeatType to Daily, Weekly, or Monthly, you must specify this parameter.
	RepeatEndTime *string `json:"RepeatEndTime,omitempty" xml:"RepeatEndTime,omitempty"`
	// The point in time when the recurrence starts. Example: 08:00. The value must be on the hour or on the half hour, and at least 30 minutes earlier than the value of RepeatEndTime.
	//
	// >  If you set RepeatType to Permanent or None, leave this parameter empty. If you set RepeatType to Daily, Weekly, or Monthly, you must specify this parameter.
	RepeatStartTime *string `json:"RepeatStartTime,omitempty" xml:"RepeatStartTime,omitempty"`
	// The recurrence type for the access control policy to take effect. Valid values:
	//
	// *   **Permanent** (default): The policy always takes effect.
	// *   **None**: The policy takes effect for only once.
	// *   **Daily**: The policy takes effect on a daily basis.
	// *   **Weekly**: The policy takes effect on a weekly basis.
	// *   **Monthly**: The policy takes effect on a monthly basis.
	RepeatType *string `json:"RepeatType,omitempty" xml:"RepeatType,omitempty"`
	// The source address in the access control policy.
	//
	// *   If **SourceType** is set to net, the value of **Source** is a CIDR block. Example: 1.2.XX.XX/24.
	// *   If **SourceType** is set to group, the value of **Source** is an address book. Example: db_group.
	// *   If **SourceType** is set to location, the value of **Source** is a location. For more information about the location codes, see the "AddControlPolicy" topic. Example: \["BJ11", "ZB"]
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The type of the source address in the access control policy. Valid values:
	//
	// *   **net**: CIDR block
	// *   **group**: address book
	// *   **location**: location
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The time when the access control policy starts to take effect. The value is a UNIX timestamp. Unit: seconds. The value must be on the hour or on the half hour, and at least 30 minutes earlier than the value of EndTime.
	//
	// >  If you set RepeatType to Permanent, leave this parameter empty. If you set RepeatType to None, Daily, Weekly, or Monthly, you must specify this parameter.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ModifyControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyControlPolicyRequest) SetAclAction(v string) *ModifyControlPolicyRequest {
	s.AclAction = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetAclUuid(v string) *ModifyControlPolicyRequest {
	s.AclUuid = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetApplicationName(v string) *ModifyControlPolicyRequest {
	s.ApplicationName = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetApplicationNameList(v []*string) *ModifyControlPolicyRequest {
	s.ApplicationNameList = v
	return s
}

func (s *ModifyControlPolicyRequest) SetDescription(v string) *ModifyControlPolicyRequest {
	s.Description = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetDestPort(v string) *ModifyControlPolicyRequest {
	s.DestPort = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetDestPortGroup(v string) *ModifyControlPolicyRequest {
	s.DestPortGroup = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetDestPortType(v string) *ModifyControlPolicyRequest {
	s.DestPortType = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetDestination(v string) *ModifyControlPolicyRequest {
	s.Destination = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetDestinationType(v string) *ModifyControlPolicyRequest {
	s.DestinationType = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetDirection(v string) *ModifyControlPolicyRequest {
	s.Direction = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetEndTime(v int64) *ModifyControlPolicyRequest {
	s.EndTime = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetLang(v string) *ModifyControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetProto(v string) *ModifyControlPolicyRequest {
	s.Proto = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetRelease(v string) *ModifyControlPolicyRequest {
	s.Release = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetRepeatDays(v []*int64) *ModifyControlPolicyRequest {
	s.RepeatDays = v
	return s
}

func (s *ModifyControlPolicyRequest) SetRepeatEndTime(v string) *ModifyControlPolicyRequest {
	s.RepeatEndTime = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetRepeatStartTime(v string) *ModifyControlPolicyRequest {
	s.RepeatStartTime = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetRepeatType(v string) *ModifyControlPolicyRequest {
	s.RepeatType = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetSource(v string) *ModifyControlPolicyRequest {
	s.Source = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetSourceType(v string) *ModifyControlPolicyRequest {
	s.SourceType = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetStartTime(v int64) *ModifyControlPolicyRequest {
	s.StartTime = &v
	return s
}

type ModifyControlPolicyResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyControlPolicyResponseBody) SetRequestId(v string) *ModifyControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

type ModifyControlPolicyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyControlPolicyResponse) SetHeaders(v map[string]*string) *ModifyControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyControlPolicyResponse) SetStatusCode(v int32) *ModifyControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyControlPolicyResponse) SetBody(v *ModifyControlPolicyResponseBody) *ModifyControlPolicyResponse {
	s.Body = v
	return s
}

type ModifyControlPolicyPositionRequest struct {
	// The direction of the traffic to which the IPv4 access control policy applies. Valid values:
	//
	// *   in: inbound traffic
	// *   out: outbound traffic
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// *   zh: Chinese (default)
	// *   en: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The new priority of the IPv4 access control policy. You must specify a numeric value for this parameter. The value 1 indicates the highest priority. A larger value indicates a lower priority.
	//
	// >  The new priority cannot exceed the priority range of the IPv4 access control policy. Otherwise, an error occurs when you call this operation. Before you call this operation, we recommend that you use the [DescribePolicyPriorUsed](~~138862~~) operation to query the priority range of the IPv4 access control policy in the specified direction.
	NewOrder *string `json:"NewOrder,omitempty" xml:"NewOrder,omitempty"`
	// The original priority of the IPv4 access control policy.
	OldOrder *string `json:"OldOrder,omitempty" xml:"OldOrder,omitempty"`
	// Deprecated
	// The source IP address of the request.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s ModifyControlPolicyPositionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyControlPolicyPositionRequest) GoString() string {
	return s.String()
}

func (s *ModifyControlPolicyPositionRequest) SetDirection(v string) *ModifyControlPolicyPositionRequest {
	s.Direction = &v
	return s
}

func (s *ModifyControlPolicyPositionRequest) SetLang(v string) *ModifyControlPolicyPositionRequest {
	s.Lang = &v
	return s
}

func (s *ModifyControlPolicyPositionRequest) SetNewOrder(v string) *ModifyControlPolicyPositionRequest {
	s.NewOrder = &v
	return s
}

func (s *ModifyControlPolicyPositionRequest) SetOldOrder(v string) *ModifyControlPolicyPositionRequest {
	s.OldOrder = &v
	return s
}

func (s *ModifyControlPolicyPositionRequest) SetSourceIp(v string) *ModifyControlPolicyPositionRequest {
	s.SourceIp = &v
	return s
}

type ModifyControlPolicyPositionResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyControlPolicyPositionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyControlPolicyPositionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyControlPolicyPositionResponseBody) SetRequestId(v string) *ModifyControlPolicyPositionResponseBody {
	s.RequestId = &v
	return s
}

type ModifyControlPolicyPositionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyControlPolicyPositionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyControlPolicyPositionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyControlPolicyPositionResponse) GoString() string {
	return s.String()
}

func (s *ModifyControlPolicyPositionResponse) SetHeaders(v map[string]*string) *ModifyControlPolicyPositionResponse {
	s.Headers = v
	return s
}

func (s *ModifyControlPolicyPositionResponse) SetStatusCode(v int32) *ModifyControlPolicyPositionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyControlPolicyPositionResponse) SetBody(v *ModifyControlPolicyPositionResponseBody) *ModifyControlPolicyPositionResponse {
	s.Body = v
	return s
}

type ModifyDefaultIPSConfigRequest struct {
	AiRules        *string `json:"AiRules,omitempty" xml:"AiRules,omitempty"`
	BasicRules     *string `json:"BasicRules,omitempty" xml:"BasicRules,omitempty"`
	CtiRules       *string `json:"CtiRules,omitempty" xml:"CtiRules,omitempty"`
	EnableAllPatch *string `json:"EnableAllPatch,omitempty" xml:"EnableAllPatch,omitempty"`
	EnableDefault  *string `json:"EnableDefault,omitempty" xml:"EnableDefault,omitempty"`
	Lang           *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PatchRules     *string `json:"PatchRules,omitempty" xml:"PatchRules,omitempty"`
	RuleClass      *string `json:"RuleClass,omitempty" xml:"RuleClass,omitempty"`
	RunMode        *string `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
	SourceIp       *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s ModifyDefaultIPSConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefaultIPSConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyDefaultIPSConfigRequest) SetAiRules(v string) *ModifyDefaultIPSConfigRequest {
	s.AiRules = &v
	return s
}

func (s *ModifyDefaultIPSConfigRequest) SetBasicRules(v string) *ModifyDefaultIPSConfigRequest {
	s.BasicRules = &v
	return s
}

func (s *ModifyDefaultIPSConfigRequest) SetCtiRules(v string) *ModifyDefaultIPSConfigRequest {
	s.CtiRules = &v
	return s
}

func (s *ModifyDefaultIPSConfigRequest) SetEnableAllPatch(v string) *ModifyDefaultIPSConfigRequest {
	s.EnableAllPatch = &v
	return s
}

func (s *ModifyDefaultIPSConfigRequest) SetEnableDefault(v string) *ModifyDefaultIPSConfigRequest {
	s.EnableDefault = &v
	return s
}

func (s *ModifyDefaultIPSConfigRequest) SetLang(v string) *ModifyDefaultIPSConfigRequest {
	s.Lang = &v
	return s
}

func (s *ModifyDefaultIPSConfigRequest) SetPatchRules(v string) *ModifyDefaultIPSConfigRequest {
	s.PatchRules = &v
	return s
}

func (s *ModifyDefaultIPSConfigRequest) SetRuleClass(v string) *ModifyDefaultIPSConfigRequest {
	s.RuleClass = &v
	return s
}

func (s *ModifyDefaultIPSConfigRequest) SetRunMode(v string) *ModifyDefaultIPSConfigRequest {
	s.RunMode = &v
	return s
}

func (s *ModifyDefaultIPSConfigRequest) SetSourceIp(v string) *ModifyDefaultIPSConfigRequest {
	s.SourceIp = &v
	return s
}

type ModifyDefaultIPSConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDefaultIPSConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefaultIPSConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDefaultIPSConfigResponseBody) SetRequestId(v string) *ModifyDefaultIPSConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDefaultIPSConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDefaultIPSConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDefaultIPSConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDefaultIPSConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyDefaultIPSConfigResponse) SetHeaders(v map[string]*string) *ModifyDefaultIPSConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyDefaultIPSConfigResponse) SetStatusCode(v int32) *ModifyDefaultIPSConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDefaultIPSConfigResponse) SetBody(v *ModifyDefaultIPSConfigResponseBody) *ModifyDefaultIPSConfigResponse {
	s.Body = v
	return s
}

type ModifyFirewallV2RoutePolicySwitchRequest struct {
	// The instance ID of the virtual private cloud (VPC) firewall.
	FirewallId *string `json:"FirewallId,omitempty" xml:"FirewallId,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// *   **zh** (default): Chinese
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Specifies whether to restore the traffic redirection configurations. Valid values:
	//
	// *   true: roll back
	// *   false: withdraw
	ShouldRecover *string `json:"ShouldRecover,omitempty" xml:"ShouldRecover,omitempty"`
	// The ID of the routing policy.
	TrFirewallRoutePolicyId *string `json:"TrFirewallRoutePolicyId,omitempty" xml:"TrFirewallRoutePolicyId,omitempty"`
	// The status of the routing policy. Valid values:
	//
	// *   open: enabled
	// *   close: disabled
	TrFirewallRoutePolicySwitchStatus *string `json:"TrFirewallRoutePolicySwitchStatus,omitempty" xml:"TrFirewallRoutePolicySwitchStatus,omitempty"`
}

func (s ModifyFirewallV2RoutePolicySwitchRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFirewallV2RoutePolicySwitchRequest) GoString() string {
	return s.String()
}

func (s *ModifyFirewallV2RoutePolicySwitchRequest) SetFirewallId(v string) *ModifyFirewallV2RoutePolicySwitchRequest {
	s.FirewallId = &v
	return s
}

func (s *ModifyFirewallV2RoutePolicySwitchRequest) SetLang(v string) *ModifyFirewallV2RoutePolicySwitchRequest {
	s.Lang = &v
	return s
}

func (s *ModifyFirewallV2RoutePolicySwitchRequest) SetShouldRecover(v string) *ModifyFirewallV2RoutePolicySwitchRequest {
	s.ShouldRecover = &v
	return s
}

func (s *ModifyFirewallV2RoutePolicySwitchRequest) SetTrFirewallRoutePolicyId(v string) *ModifyFirewallV2RoutePolicySwitchRequest {
	s.TrFirewallRoutePolicyId = &v
	return s
}

func (s *ModifyFirewallV2RoutePolicySwitchRequest) SetTrFirewallRoutePolicySwitchStatus(v string) *ModifyFirewallV2RoutePolicySwitchRequest {
	s.TrFirewallRoutePolicySwitchStatus = &v
	return s
}

type ModifyFirewallV2RoutePolicySwitchResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyFirewallV2RoutePolicySwitchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyFirewallV2RoutePolicySwitchResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFirewallV2RoutePolicySwitchResponseBody) SetRequestId(v string) *ModifyFirewallV2RoutePolicySwitchResponseBody {
	s.RequestId = &v
	return s
}

type ModifyFirewallV2RoutePolicySwitchResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyFirewallV2RoutePolicySwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyFirewallV2RoutePolicySwitchResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyFirewallV2RoutePolicySwitchResponse) GoString() string {
	return s.String()
}

func (s *ModifyFirewallV2RoutePolicySwitchResponse) SetHeaders(v map[string]*string) *ModifyFirewallV2RoutePolicySwitchResponse {
	s.Headers = v
	return s
}

func (s *ModifyFirewallV2RoutePolicySwitchResponse) SetStatusCode(v int32) *ModifyFirewallV2RoutePolicySwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyFirewallV2RoutePolicySwitchResponse) SetBody(v *ModifyFirewallV2RoutePolicySwitchResponseBody) *ModifyFirewallV2RoutePolicySwitchResponse {
	s.Body = v
	return s
}

type ModifyInstanceMemberAttributesRequest struct {
	// The members that to be modified.
	Members []*ModifyInstanceMemberAttributesRequestMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
}

func (s ModifyInstanceMemberAttributesRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceMemberAttributesRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMemberAttributesRequest) SetMembers(v []*ModifyInstanceMemberAttributesRequestMembers) *ModifyInstanceMemberAttributesRequest {
	s.Members = v
	return s
}

type ModifyInstanceMemberAttributesRequestMembers struct {
	// The remarks of the member in Cloud Firewall.
	MemberDesc *string `json:"MemberDesc,omitempty" xml:"MemberDesc,omitempty"`
	// The UID of the member in Cloud Firewall.
	MemberUid *int64 `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
}

func (s ModifyInstanceMemberAttributesRequestMembers) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceMemberAttributesRequestMembers) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMemberAttributesRequestMembers) SetMemberDesc(v string) *ModifyInstanceMemberAttributesRequestMembers {
	s.MemberDesc = &v
	return s
}

func (s *ModifyInstanceMemberAttributesRequestMembers) SetMemberUid(v int64) *ModifyInstanceMemberAttributesRequestMembers {
	s.MemberUid = &v
	return s
}

type ModifyInstanceMemberAttributesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceMemberAttributesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceMemberAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMemberAttributesResponseBody) SetRequestId(v string) *ModifyInstanceMemberAttributesResponseBody {
	s.RequestId = &v
	return s
}

type ModifyInstanceMemberAttributesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceMemberAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceMemberAttributesResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceMemberAttributesResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMemberAttributesResponse) SetHeaders(v map[string]*string) *ModifyInstanceMemberAttributesResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceMemberAttributesResponse) SetStatusCode(v int32) *ModifyInstanceMemberAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceMemberAttributesResponse) SetBody(v *ModifyInstanceMemberAttributesResponseBody) *ModifyInstanceMemberAttributesResponse {
	s.Body = v
	return s
}

type ModifyNatFirewallControlPolicyRequest struct {
	// The action that Cloud Firewall performs on the traffic. Valid values:
	//
	// *   **accept**: allows the traffic.
	// *   **drop**: denies the traffic.
	// *   **log**: monitors the traffic.
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The UUID of the access control policy.
	//
	// To modify the configurations of an access control policy, you must provide the UUID of the policy. You can call the DescribeNatFirewallControlPolicy operation to query the UUIDs of access control policies.
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The name of the application.
	ApplicationNameList []*string `json:"ApplicationNameList,omitempty" xml:"ApplicationNameList,omitempty" type:"Repeated"`
	// The description of the access control policy. Fuzzy match is supported.
	//
	// > If you do not specify this parameter, the descriptions of all policies are queried.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination port in the access control policy.
	//
	// > If **DestPortType** is set to `port`, you must specify this parameter.
	DestPort *string `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	// The name of the destination port address book in the access control policy.
	DestPortGroup *string `json:"DestPortGroup,omitempty" xml:"DestPortGroup,omitempty"`
	// The type of the destination port in the access control policy. Valid values:
	//
	// *   **port**: port
	// *   **group**: port address book
	DestPortType *string `json:"DestPortType,omitempty" xml:"DestPortType,omitempty"`
	// The destination address in the access control policy.
	//
	// *   If **DestinationType** is set to net, the value of **Destination** is a CIDR block. Example: 1.2.3.4/24
	// *   If **DestinationType** is set to group, the value of **Destination** is an address book. Example: db_group
	// *   If **DestinationType** is set to domain, the value of **Destination** is a domain name. Example: \*.aliyuncs.com
	// *   If **DestinationType** is set to location, the value of **Destination** is a location. For more information about the location codes, see the "AddControlPolicy" topic. Example: \["BJ11", "ZB"]
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The type of the destination address in the access control policy. Valid values:
	//
	// *   **net**: CIDR block
	// *   **group**: address book
	// *   **domain**: domain name
	// *   **location**
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The direction of the traffic to which the access control policy applies.
	//
	// *   Set the value to **out**.
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The domain name resolution method of the access control policy. By default, an access control policy is enabled after it is created. Valid values:
	//
	// *   **0**: Fully qualified domain name (FQDN)-based resolution
	// *   **1**: Domain Name System (DNS)-based dynamic resolution
	// *   **2**: FQDN and DNS-based dynamic resolution
	DomainResolveType *string `json:"DomainResolveType,omitempty" xml:"DomainResolveType,omitempty"`
	// The time when the access control policy stops taking effect. The value is a UNIX timestamp. Unit: seconds. The value must be on the hour or on the half hour, and at least 30 minutes later than the value of StartTime.
	//
	// >  If RepeatType is set to Permanent, EndTime is left empty. If RepeatType is set to None, Daily, Weekly, or Monthly, EndTime must be specified.
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the content within the request and the response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the NAT gateway.
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The protocol type supported by the access control policy. Valid values:
	//
	// *   **ANY**
	// *   **TCP**
	// *   **UDP**
	// *   **ICMP**
	//
	// > The value **ANY** indicates all types of protocols.
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// The status of the access control policy. Valid values:
	//
	// *   **true**: enabled
	// *   **false**: disabled
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The days of a week or of a month on which the access control policy takes effect.
	//
	// *   If RepeatType is set to `Permanent`, `None`, or `Daily`, RepeatDays is left empty. Example: \[].
	// *   If RepeatType is set to Weekly, RepeatDays must be specified. Example: \[0, 6].
	//
	// >  If RepeatType is set to Weekly, the fields in the value of RepeatDays cannot be repeated.
	//
	// *   If RepeatType is set to `Monthly`, RepeatDays must be specified. Example: \[1, 31].
	//
	// >  If RepeatType is set to Monthly, the fields in the value of RepeatDays cannot be repeated.
	RepeatDays []*int64 `json:"RepeatDays,omitempty" xml:"RepeatDays,omitempty" type:"Repeated"`
	// The point in time when the recurrence ends. Example: 23:30. The value must be on the hour or on the half hour, and at least 30 minutes later than the value of RepeatStartTime.
	//
	// >  If RepeatType is set to Permanent or None, RepeatEndTime is left empty. If RepeatType is set to Daily, Weekly, or Monthly, RepeatEndTime must be specified.
	RepeatEndTime *string `json:"RepeatEndTime,omitempty" xml:"RepeatEndTime,omitempty"`
	// The point in time when the recurrence starts. Example: 08:00. The value must be on the hour or on the half hour, and at least 30 minutes earlier than the value of RepeatEndTime.
	//
	// >  If RepeatType is set to Permanent or None, RepeatStartTime is left empty. If RepeatType is set to Daily, Weekly, or Monthly, this parameter must be specified.
	RepeatStartTime *string `json:"RepeatStartTime,omitempty" xml:"RepeatStartTime,omitempty"`
	// The recurrence type for the access control policy to take effect. Valid values:
	//
	// *   **Permanent** (default): The policy always takes effect.
	// *   **None**: The policy takes effect for only once.
	// *   **Daily**: The policy takes effect on a daily basis.
	// *   **Weekly**: The policy takes effect on a weekly basis.
	// *   **Monthly**: The policy takes effect on a monthly basis.
	RepeatType *string `json:"RepeatType,omitempty" xml:"RepeatType,omitempty"`
	// The source address in the access control policy. Valid values:
	//
	// *   If **SourceType** is set to `net`, the value of this parameter is a CIDR block. Example: 10.2.XX.XX/24.
	// *   If **SourceType** is set to `group`, the value of this parameter is an address book name. Example: db_group.
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The type of the source address in the access control policy. Valid values:
	//
	// *   **net**: CIDR block
	// *   **group**: address book
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The time when the access control policy starts to take effect. The value is a UNIX timestamp. Unit: seconds. The value must be on the hour or on the half hour, and at least 30 minutes earlier than the value of EndTime.
	//
	// >  If RepeatType is set to Permanent, StartTime is left empty. If RepeatType is set to None, Daily, Weekly, or Monthly, StartTime must be specified.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ModifyNatFirewallControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyNatFirewallControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyNatFirewallControlPolicyRequest) SetAclAction(v string) *ModifyNatFirewallControlPolicyRequest {
	s.AclAction = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetAclUuid(v string) *ModifyNatFirewallControlPolicyRequest {
	s.AclUuid = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetApplicationNameList(v []*string) *ModifyNatFirewallControlPolicyRequest {
	s.ApplicationNameList = v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetDescription(v string) *ModifyNatFirewallControlPolicyRequest {
	s.Description = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetDestPort(v string) *ModifyNatFirewallControlPolicyRequest {
	s.DestPort = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetDestPortGroup(v string) *ModifyNatFirewallControlPolicyRequest {
	s.DestPortGroup = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetDestPortType(v string) *ModifyNatFirewallControlPolicyRequest {
	s.DestPortType = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetDestination(v string) *ModifyNatFirewallControlPolicyRequest {
	s.Destination = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetDestinationType(v string) *ModifyNatFirewallControlPolicyRequest {
	s.DestinationType = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetDirection(v string) *ModifyNatFirewallControlPolicyRequest {
	s.Direction = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetDomainResolveType(v string) *ModifyNatFirewallControlPolicyRequest {
	s.DomainResolveType = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetEndTime(v int64) *ModifyNatFirewallControlPolicyRequest {
	s.EndTime = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetLang(v string) *ModifyNatFirewallControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetNatGatewayId(v string) *ModifyNatFirewallControlPolicyRequest {
	s.NatGatewayId = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetProto(v string) *ModifyNatFirewallControlPolicyRequest {
	s.Proto = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetRelease(v string) *ModifyNatFirewallControlPolicyRequest {
	s.Release = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetRepeatDays(v []*int64) *ModifyNatFirewallControlPolicyRequest {
	s.RepeatDays = v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetRepeatEndTime(v string) *ModifyNatFirewallControlPolicyRequest {
	s.RepeatEndTime = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetRepeatStartTime(v string) *ModifyNatFirewallControlPolicyRequest {
	s.RepeatStartTime = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetRepeatType(v string) *ModifyNatFirewallControlPolicyRequest {
	s.RepeatType = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetSource(v string) *ModifyNatFirewallControlPolicyRequest {
	s.Source = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetSourceType(v string) *ModifyNatFirewallControlPolicyRequest {
	s.SourceType = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyRequest) SetStartTime(v int64) *ModifyNatFirewallControlPolicyRequest {
	s.StartTime = &v
	return s
}

type ModifyNatFirewallControlPolicyResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyNatFirewallControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyNatFirewallControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNatFirewallControlPolicyResponseBody) SetRequestId(v string) *ModifyNatFirewallControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

type ModifyNatFirewallControlPolicyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyNatFirewallControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyNatFirewallControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyNatFirewallControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyNatFirewallControlPolicyResponse) SetHeaders(v map[string]*string) *ModifyNatFirewallControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyNatFirewallControlPolicyResponse) SetStatusCode(v int32) *ModifyNatFirewallControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyResponse) SetBody(v *ModifyNatFirewallControlPolicyResponseBody) *ModifyNatFirewallControlPolicyResponse {
	s.Body = v
	return s
}

type ModifyNatFirewallControlPolicyPositionRequest struct {
	// The UUID of the access control policy.
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The direction of the traffic to which the access control policy applies.
	//
	// *   Set the value to **out**.
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the NAT gateway.
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The new priority of the IPv4 access control policy. You must specify a numeric value for this parameter. The value 1 indicates the highest priority. A larger value indicates a lower priority.
	//
	// > Make sure that the value of this parameter is within the priority range of existing IPv4 access control policies. Otherwise, an error occurs when you call this operation.
	//
	// Before you call this operation, we recommend that you call the DescribeNatFirewallPolicyPriorUsed operation to query the priority range of the IPv4 access control policies in the specified traffic direction.
	NewOrder *int32 `json:"NewOrder,omitempty" xml:"NewOrder,omitempty"`
}

func (s ModifyNatFirewallControlPolicyPositionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyNatFirewallControlPolicyPositionRequest) GoString() string {
	return s.String()
}

func (s *ModifyNatFirewallControlPolicyPositionRequest) SetAclUuid(v string) *ModifyNatFirewallControlPolicyPositionRequest {
	s.AclUuid = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyPositionRequest) SetDirection(v string) *ModifyNatFirewallControlPolicyPositionRequest {
	s.Direction = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyPositionRequest) SetLang(v string) *ModifyNatFirewallControlPolicyPositionRequest {
	s.Lang = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyPositionRequest) SetNatGatewayId(v string) *ModifyNatFirewallControlPolicyPositionRequest {
	s.NatGatewayId = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyPositionRequest) SetNewOrder(v int32) *ModifyNatFirewallControlPolicyPositionRequest {
	s.NewOrder = &v
	return s
}

type ModifyNatFirewallControlPolicyPositionResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyNatFirewallControlPolicyPositionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyNatFirewallControlPolicyPositionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyNatFirewallControlPolicyPositionResponseBody) SetRequestId(v string) *ModifyNatFirewallControlPolicyPositionResponseBody {
	s.RequestId = &v
	return s
}

type ModifyNatFirewallControlPolicyPositionResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyNatFirewallControlPolicyPositionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyNatFirewallControlPolicyPositionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyNatFirewallControlPolicyPositionResponse) GoString() string {
	return s.String()
}

func (s *ModifyNatFirewallControlPolicyPositionResponse) SetHeaders(v map[string]*string) *ModifyNatFirewallControlPolicyPositionResponse {
	s.Headers = v
	return s
}

func (s *ModifyNatFirewallControlPolicyPositionResponse) SetStatusCode(v int32) *ModifyNatFirewallControlPolicyPositionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyNatFirewallControlPolicyPositionResponse) SetBody(v *ModifyNatFirewallControlPolicyPositionResponseBody) *ModifyNatFirewallControlPolicyPositionResponse {
	s.Body = v
	return s
}

type ModifyPolicyAdvancedConfigRequest struct {
	// Specifies whether to enable the strict mode for the access control policy. Valid values:
	//
	// *   **on**: enables the strict mode.
	// *   **off**: disables the strict mode.
	InternetSwitch *string `json:"InternetSwitch,omitempty" xml:"InternetSwitch,omitempty"`
	// The natural language of the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Deprecated
	// The source IP address of the request.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s ModifyPolicyAdvancedConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyAdvancedConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyPolicyAdvancedConfigRequest) SetInternetSwitch(v string) *ModifyPolicyAdvancedConfigRequest {
	s.InternetSwitch = &v
	return s
}

func (s *ModifyPolicyAdvancedConfigRequest) SetLang(v string) *ModifyPolicyAdvancedConfigRequest {
	s.Lang = &v
	return s
}

func (s *ModifyPolicyAdvancedConfigRequest) SetSourceIp(v string) *ModifyPolicyAdvancedConfigRequest {
	s.SourceIp = &v
	return s
}

type ModifyPolicyAdvancedConfigResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPolicyAdvancedConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyAdvancedConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPolicyAdvancedConfigResponseBody) SetRequestId(v string) *ModifyPolicyAdvancedConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyPolicyAdvancedConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPolicyAdvancedConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPolicyAdvancedConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyAdvancedConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyPolicyAdvancedConfigResponse) SetHeaders(v map[string]*string) *ModifyPolicyAdvancedConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyPolicyAdvancedConfigResponse) SetStatusCode(v int32) *ModifyPolicyAdvancedConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPolicyAdvancedConfigResponse) SetBody(v *ModifyPolicyAdvancedConfigResponseBody) *ModifyPolicyAdvancedConfigResponse {
	s.Body = v
	return s
}

type ModifyTrFirewallV2ConfigurationRequest struct {
	FirewallId   *string `json:"FirewallId,omitempty" xml:"FirewallId,omitempty"`
	FirewallName *string `json:"FirewallName,omitempty" xml:"FirewallName,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s ModifyTrFirewallV2ConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyTrFirewallV2ConfigurationRequest) GoString() string {
	return s.String()
}

func (s *ModifyTrFirewallV2ConfigurationRequest) SetFirewallId(v string) *ModifyTrFirewallV2ConfigurationRequest {
	s.FirewallId = &v
	return s
}

func (s *ModifyTrFirewallV2ConfigurationRequest) SetFirewallName(v string) *ModifyTrFirewallV2ConfigurationRequest {
	s.FirewallName = &v
	return s
}

func (s *ModifyTrFirewallV2ConfigurationRequest) SetLang(v string) *ModifyTrFirewallV2ConfigurationRequest {
	s.Lang = &v
	return s
}

type ModifyTrFirewallV2ConfigurationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTrFirewallV2ConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyTrFirewallV2ConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTrFirewallV2ConfigurationResponseBody) SetRequestId(v string) *ModifyTrFirewallV2ConfigurationResponseBody {
	s.RequestId = &v
	return s
}

type ModifyTrFirewallV2ConfigurationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyTrFirewallV2ConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyTrFirewallV2ConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyTrFirewallV2ConfigurationResponse) GoString() string {
	return s.String()
}

func (s *ModifyTrFirewallV2ConfigurationResponse) SetHeaders(v map[string]*string) *ModifyTrFirewallV2ConfigurationResponse {
	s.Headers = v
	return s
}

func (s *ModifyTrFirewallV2ConfigurationResponse) SetStatusCode(v int32) *ModifyTrFirewallV2ConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTrFirewallV2ConfigurationResponse) SetBody(v *ModifyTrFirewallV2ConfigurationResponseBody) *ModifyTrFirewallV2ConfigurationResponse {
	s.Body = v
	return s
}

type ModifyTrFirewallV2RoutePolicyScopeRequest struct {
	DestCandidateList       []*ModifyTrFirewallV2RoutePolicyScopeRequestDestCandidateList `json:"DestCandidateList,omitempty" xml:"DestCandidateList,omitempty" type:"Repeated"`
	FirewallId              *string                                                       `json:"FirewallId,omitempty" xml:"FirewallId,omitempty"`
	Lang                    *string                                                       `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ShouldRecover           *string                                                       `json:"ShouldRecover,omitempty" xml:"ShouldRecover,omitempty"`
	SrcCandidateList        []*ModifyTrFirewallV2RoutePolicyScopeRequestSrcCandidateList  `json:"SrcCandidateList,omitempty" xml:"SrcCandidateList,omitempty" type:"Repeated"`
	TrFirewallRoutePolicyId *string                                                       `json:"TrFirewallRoutePolicyId,omitempty" xml:"TrFirewallRoutePolicyId,omitempty"`
}

func (s ModifyTrFirewallV2RoutePolicyScopeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyTrFirewallV2RoutePolicyScopeRequest) GoString() string {
	return s.String()
}

func (s *ModifyTrFirewallV2RoutePolicyScopeRequest) SetDestCandidateList(v []*ModifyTrFirewallV2RoutePolicyScopeRequestDestCandidateList) *ModifyTrFirewallV2RoutePolicyScopeRequest {
	s.DestCandidateList = v
	return s
}

func (s *ModifyTrFirewallV2RoutePolicyScopeRequest) SetFirewallId(v string) *ModifyTrFirewallV2RoutePolicyScopeRequest {
	s.FirewallId = &v
	return s
}

func (s *ModifyTrFirewallV2RoutePolicyScopeRequest) SetLang(v string) *ModifyTrFirewallV2RoutePolicyScopeRequest {
	s.Lang = &v
	return s
}

func (s *ModifyTrFirewallV2RoutePolicyScopeRequest) SetShouldRecover(v string) *ModifyTrFirewallV2RoutePolicyScopeRequest {
	s.ShouldRecover = &v
	return s
}

func (s *ModifyTrFirewallV2RoutePolicyScopeRequest) SetSrcCandidateList(v []*ModifyTrFirewallV2RoutePolicyScopeRequestSrcCandidateList) *ModifyTrFirewallV2RoutePolicyScopeRequest {
	s.SrcCandidateList = v
	return s
}

func (s *ModifyTrFirewallV2RoutePolicyScopeRequest) SetTrFirewallRoutePolicyId(v string) *ModifyTrFirewallV2RoutePolicyScopeRequest {
	s.TrFirewallRoutePolicyId = &v
	return s
}

type ModifyTrFirewallV2RoutePolicyScopeRequestDestCandidateList struct {
	CandidateId   *string `json:"CandidateId,omitempty" xml:"CandidateId,omitempty"`
	CandidateType *string `json:"CandidateType,omitempty" xml:"CandidateType,omitempty"`
}

func (s ModifyTrFirewallV2RoutePolicyScopeRequestDestCandidateList) String() string {
	return tea.Prettify(s)
}

func (s ModifyTrFirewallV2RoutePolicyScopeRequestDestCandidateList) GoString() string {
	return s.String()
}

func (s *ModifyTrFirewallV2RoutePolicyScopeRequestDestCandidateList) SetCandidateId(v string) *ModifyTrFirewallV2RoutePolicyScopeRequestDestCandidateList {
	s.CandidateId = &v
	return s
}

func (s *ModifyTrFirewallV2RoutePolicyScopeRequestDestCandidateList) SetCandidateType(v string) *ModifyTrFirewallV2RoutePolicyScopeRequestDestCandidateList {
	s.CandidateType = &v
	return s
}

type ModifyTrFirewallV2RoutePolicyScopeRequestSrcCandidateList struct {
	CandidateId   *string `json:"CandidateId,omitempty" xml:"CandidateId,omitempty"`
	CandidateType *string `json:"CandidateType,omitempty" xml:"CandidateType,omitempty"`
}

func (s ModifyTrFirewallV2RoutePolicyScopeRequestSrcCandidateList) String() string {
	return tea.Prettify(s)
}

func (s ModifyTrFirewallV2RoutePolicyScopeRequestSrcCandidateList) GoString() string {
	return s.String()
}

func (s *ModifyTrFirewallV2RoutePolicyScopeRequestSrcCandidateList) SetCandidateId(v string) *ModifyTrFirewallV2RoutePolicyScopeRequestSrcCandidateList {
	s.CandidateId = &v
	return s
}

func (s *ModifyTrFirewallV2RoutePolicyScopeRequestSrcCandidateList) SetCandidateType(v string) *ModifyTrFirewallV2RoutePolicyScopeRequestSrcCandidateList {
	s.CandidateType = &v
	return s
}

type ModifyTrFirewallV2RoutePolicyScopeShrinkRequest struct {
	DestCandidateListShrink *string `json:"DestCandidateList,omitempty" xml:"DestCandidateList,omitempty"`
	FirewallId              *string `json:"FirewallId,omitempty" xml:"FirewallId,omitempty"`
	Lang                    *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ShouldRecover           *string `json:"ShouldRecover,omitempty" xml:"ShouldRecover,omitempty"`
	SrcCandidateListShrink  *string `json:"SrcCandidateList,omitempty" xml:"SrcCandidateList,omitempty"`
	TrFirewallRoutePolicyId *string `json:"TrFirewallRoutePolicyId,omitempty" xml:"TrFirewallRoutePolicyId,omitempty"`
}

func (s ModifyTrFirewallV2RoutePolicyScopeShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyTrFirewallV2RoutePolicyScopeShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyTrFirewallV2RoutePolicyScopeShrinkRequest) SetDestCandidateListShrink(v string) *ModifyTrFirewallV2RoutePolicyScopeShrinkRequest {
	s.DestCandidateListShrink = &v
	return s
}

func (s *ModifyTrFirewallV2RoutePolicyScopeShrinkRequest) SetFirewallId(v string) *ModifyTrFirewallV2RoutePolicyScopeShrinkRequest {
	s.FirewallId = &v
	return s
}

func (s *ModifyTrFirewallV2RoutePolicyScopeShrinkRequest) SetLang(v string) *ModifyTrFirewallV2RoutePolicyScopeShrinkRequest {
	s.Lang = &v
	return s
}

func (s *ModifyTrFirewallV2RoutePolicyScopeShrinkRequest) SetShouldRecover(v string) *ModifyTrFirewallV2RoutePolicyScopeShrinkRequest {
	s.ShouldRecover = &v
	return s
}

func (s *ModifyTrFirewallV2RoutePolicyScopeShrinkRequest) SetSrcCandidateListShrink(v string) *ModifyTrFirewallV2RoutePolicyScopeShrinkRequest {
	s.SrcCandidateListShrink = &v
	return s
}

func (s *ModifyTrFirewallV2RoutePolicyScopeShrinkRequest) SetTrFirewallRoutePolicyId(v string) *ModifyTrFirewallV2RoutePolicyScopeShrinkRequest {
	s.TrFirewallRoutePolicyId = &v
	return s
}

type ModifyTrFirewallV2RoutePolicyScopeResponseBody struct {
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TrFirewallRoutePolicyId *string `json:"TrFirewallRoutePolicyId,omitempty" xml:"TrFirewallRoutePolicyId,omitempty"`
}

func (s ModifyTrFirewallV2RoutePolicyScopeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyTrFirewallV2RoutePolicyScopeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTrFirewallV2RoutePolicyScopeResponseBody) SetRequestId(v string) *ModifyTrFirewallV2RoutePolicyScopeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTrFirewallV2RoutePolicyScopeResponseBody) SetTrFirewallRoutePolicyId(v string) *ModifyTrFirewallV2RoutePolicyScopeResponseBody {
	s.TrFirewallRoutePolicyId = &v
	return s
}

type ModifyTrFirewallV2RoutePolicyScopeResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyTrFirewallV2RoutePolicyScopeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyTrFirewallV2RoutePolicyScopeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyTrFirewallV2RoutePolicyScopeResponse) GoString() string {
	return s.String()
}

func (s *ModifyTrFirewallV2RoutePolicyScopeResponse) SetHeaders(v map[string]*string) *ModifyTrFirewallV2RoutePolicyScopeResponse {
	s.Headers = v
	return s
}

func (s *ModifyTrFirewallV2RoutePolicyScopeResponse) SetStatusCode(v int32) *ModifyTrFirewallV2RoutePolicyScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTrFirewallV2RoutePolicyScopeResponse) SetBody(v *ModifyTrFirewallV2RoutePolicyScopeResponseBody) *ModifyTrFirewallV2RoutePolicyScopeResponse {
	s.Body = v
	return s
}

type ModifyUserIPSWhitelistRequest struct {
	Direction *int64  `json:"Direction,omitempty" xml:"Direction,omitempty"`
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	ListType  *int64  `json:"ListType,omitempty" xml:"ListType,omitempty"`
	ListValue *string `json:"ListValue,omitempty" xml:"ListValue,omitempty"`
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	WhiteType *int64  `json:"WhiteType,omitempty" xml:"WhiteType,omitempty"`
}

func (s ModifyUserIPSWhitelistRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserIPSWhitelistRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserIPSWhitelistRequest) SetDirection(v int64) *ModifyUserIPSWhitelistRequest {
	s.Direction = &v
	return s
}

func (s *ModifyUserIPSWhitelistRequest) SetIpVersion(v string) *ModifyUserIPSWhitelistRequest {
	s.IpVersion = &v
	return s
}

func (s *ModifyUserIPSWhitelistRequest) SetLang(v string) *ModifyUserIPSWhitelistRequest {
	s.Lang = &v
	return s
}

func (s *ModifyUserIPSWhitelistRequest) SetListType(v int64) *ModifyUserIPSWhitelistRequest {
	s.ListType = &v
	return s
}

func (s *ModifyUserIPSWhitelistRequest) SetListValue(v string) *ModifyUserIPSWhitelistRequest {
	s.ListValue = &v
	return s
}

func (s *ModifyUserIPSWhitelistRequest) SetSourceIp(v string) *ModifyUserIPSWhitelistRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyUserIPSWhitelistRequest) SetWhiteType(v int64) *ModifyUserIPSWhitelistRequest {
	s.WhiteType = &v
	return s
}

type ModifyUserIPSWhitelistResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUserIPSWhitelistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserIPSWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUserIPSWhitelistResponseBody) SetRequestId(v string) *ModifyUserIPSWhitelistResponseBody {
	s.RequestId = &v
	return s
}

type ModifyUserIPSWhitelistResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyUserIPSWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyUserIPSWhitelistResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserIPSWhitelistResponse) GoString() string {
	return s.String()
}

func (s *ModifyUserIPSWhitelistResponse) SetHeaders(v map[string]*string) *ModifyUserIPSWhitelistResponse {
	s.Headers = v
	return s
}

func (s *ModifyUserIPSWhitelistResponse) SetStatusCode(v int32) *ModifyUserIPSWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyUserIPSWhitelistResponse) SetBody(v *ModifyUserIPSWhitelistResponseBody) *ModifyUserIPSWhitelistResponse {
	s.Body = v
	return s
}

type ModifyVpcFirewallCenConfigureRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The instance ID of the VPC firewall.
	//
	// > You can call the [DescribeVpcFirewallCenList](~~345777~~) operation to query the instance IDs of VPC firewalls.
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
	// The instance name of the VPC firewall.
	VpcFirewallName *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
}

func (s ModifyVpcFirewallCenConfigureRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallCenConfigureRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallCenConfigureRequest) SetLang(v string) *ModifyVpcFirewallCenConfigureRequest {
	s.Lang = &v
	return s
}

func (s *ModifyVpcFirewallCenConfigureRequest) SetMemberUid(v string) *ModifyVpcFirewallCenConfigureRequest {
	s.MemberUid = &v
	return s
}

func (s *ModifyVpcFirewallCenConfigureRequest) SetVpcFirewallId(v string) *ModifyVpcFirewallCenConfigureRequest {
	s.VpcFirewallId = &v
	return s
}

func (s *ModifyVpcFirewallCenConfigureRequest) SetVpcFirewallName(v string) *ModifyVpcFirewallCenConfigureRequest {
	s.VpcFirewallName = &v
	return s
}

type ModifyVpcFirewallCenConfigureResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVpcFirewallCenConfigureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallCenConfigureResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallCenConfigureResponseBody) SetRequestId(v string) *ModifyVpcFirewallCenConfigureResponseBody {
	s.RequestId = &v
	return s
}

type ModifyVpcFirewallCenConfigureResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVpcFirewallCenConfigureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVpcFirewallCenConfigureResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallCenConfigureResponse) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallCenConfigureResponse) SetHeaders(v map[string]*string) *ModifyVpcFirewallCenConfigureResponse {
	s.Headers = v
	return s
}

func (s *ModifyVpcFirewallCenConfigureResponse) SetStatusCode(v int32) *ModifyVpcFirewallCenConfigureResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVpcFirewallCenConfigureResponse) SetBody(v *ModifyVpcFirewallCenConfigureResponseBody) *ModifyVpcFirewallCenConfigureResponse {
	s.Body = v
	return s
}

type ModifyVpcFirewallCenSwitchStatusRequest struct {
	// Specifies whether to enable the VPC firewall. Valid values:
	//
	// *   **open**: yes
	// *   **close**: no
	FirewallSwitch *string `json:"FirewallSwitch,omitempty" xml:"FirewallSwitch,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The instance ID of the VPC firewall.
	//
	// > You can call the [DescribeVpcFirewallCenList](~~345777~~) operation to query the instance IDs of VPC firewalls.
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s ModifyVpcFirewallCenSwitchStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallCenSwitchStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallCenSwitchStatusRequest) SetFirewallSwitch(v string) *ModifyVpcFirewallCenSwitchStatusRequest {
	s.FirewallSwitch = &v
	return s
}

func (s *ModifyVpcFirewallCenSwitchStatusRequest) SetLang(v string) *ModifyVpcFirewallCenSwitchStatusRequest {
	s.Lang = &v
	return s
}

func (s *ModifyVpcFirewallCenSwitchStatusRequest) SetMemberUid(v string) *ModifyVpcFirewallCenSwitchStatusRequest {
	s.MemberUid = &v
	return s
}

func (s *ModifyVpcFirewallCenSwitchStatusRequest) SetVpcFirewallId(v string) *ModifyVpcFirewallCenSwitchStatusRequest {
	s.VpcFirewallId = &v
	return s
}

type ModifyVpcFirewallCenSwitchStatusResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVpcFirewallCenSwitchStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallCenSwitchStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallCenSwitchStatusResponseBody) SetRequestId(v string) *ModifyVpcFirewallCenSwitchStatusResponseBody {
	s.RequestId = &v
	return s
}

type ModifyVpcFirewallCenSwitchStatusResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVpcFirewallCenSwitchStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVpcFirewallCenSwitchStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallCenSwitchStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallCenSwitchStatusResponse) SetHeaders(v map[string]*string) *ModifyVpcFirewallCenSwitchStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyVpcFirewallCenSwitchStatusResponse) SetStatusCode(v int32) *ModifyVpcFirewallCenSwitchStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVpcFirewallCenSwitchStatusResponse) SetBody(v *ModifyVpcFirewallCenSwitchStatusResponseBody) *ModifyVpcFirewallCenSwitchStatusResponse {
	s.Body = v
	return s
}

type ModifyVpcFirewallConfigureRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The CIDR blocks of the local VPC. The value is a JSON string that contains the following parameters:
	//
	// *   **RouteTableId**: the ID of the route table for the local VPC.
	// *   **RouteEntryList**: The value is a JSON string that contains the DestinationCidr and NextHopInstanceId parameters. The DestinationCidr parameter indicates the destination CIDR block of the local VPC. The NextHopInstanceId parameter indicates the instance ID of the next hop for the local VPC.
	//
	// > You can call the [DescribeVpcFirewallDetail](~~342892~~) operation to query the CIDR blocks of local VPCs for VPC firewalls.
	LocalVpcCidrTableList *string `json:"LocalVpcCidrTableList,omitempty" xml:"LocalVpcCidrTableList,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The CIDR blocks of the peer VPC. The value is a JSON string that contains the following parameters:
	//
	// *   **RouteTableId**: the ID of the route table for the peer VPC.
	// *   **RouteEntryList**: The value is a JSON string that contains the DestinationCidr and NextHopInstanceId parameters. The DestinationCidr parameter indicates the destination CIDR block of the peer VPC. The NextHopInstanceId parameter indicates the instance ID of the next hop for the peer VPC.
	//
	// > You can call the [DescribeVpcFirewallDetail](~~342892~~) operation to query the CIDR blocks of peer VPCs for VPC firewalls.
	PeerVpcCidrTableList *string `json:"PeerVpcCidrTableList,omitempty" xml:"PeerVpcCidrTableList,omitempty"`
	// The instance ID of the VPC firewall.
	//
	// > You can call the [DescribeVpcFirewallList](~~342932~~) operation to query the instance IDs of VPC firewalls.
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
	// The instance name of the VPC firewall.
	VpcFirewallName *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
}

func (s ModifyVpcFirewallConfigureRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallConfigureRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallConfigureRequest) SetLang(v string) *ModifyVpcFirewallConfigureRequest {
	s.Lang = &v
	return s
}

func (s *ModifyVpcFirewallConfigureRequest) SetLocalVpcCidrTableList(v string) *ModifyVpcFirewallConfigureRequest {
	s.LocalVpcCidrTableList = &v
	return s
}

func (s *ModifyVpcFirewallConfigureRequest) SetMemberUid(v string) *ModifyVpcFirewallConfigureRequest {
	s.MemberUid = &v
	return s
}

func (s *ModifyVpcFirewallConfigureRequest) SetPeerVpcCidrTableList(v string) *ModifyVpcFirewallConfigureRequest {
	s.PeerVpcCidrTableList = &v
	return s
}

func (s *ModifyVpcFirewallConfigureRequest) SetVpcFirewallId(v string) *ModifyVpcFirewallConfigureRequest {
	s.VpcFirewallId = &v
	return s
}

func (s *ModifyVpcFirewallConfigureRequest) SetVpcFirewallName(v string) *ModifyVpcFirewallConfigureRequest {
	s.VpcFirewallName = &v
	return s
}

type ModifyVpcFirewallConfigureResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVpcFirewallConfigureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallConfigureResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallConfigureResponseBody) SetRequestId(v string) *ModifyVpcFirewallConfigureResponseBody {
	s.RequestId = &v
	return s
}

type ModifyVpcFirewallConfigureResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVpcFirewallConfigureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVpcFirewallConfigureResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallConfigureResponse) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallConfigureResponse) SetHeaders(v map[string]*string) *ModifyVpcFirewallConfigureResponse {
	s.Headers = v
	return s
}

func (s *ModifyVpcFirewallConfigureResponse) SetStatusCode(v int32) *ModifyVpcFirewallConfigureResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVpcFirewallConfigureResponse) SetBody(v *ModifyVpcFirewallConfigureResponseBody) *ModifyVpcFirewallConfigureResponse {
	s.Body = v
	return s
}

type ModifyVpcFirewallControlPolicyRequest struct {
	// The action that Cloud Firewall performs on the traffic.
	//
	// Valid values:
	//
	// *   **accept**: allows the traffic.
	// *   **drop**: blocks the traffic.
	// *   **log**: monitors the traffic.
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The unique ID of the access control policy.
	//
	// If you want to modify the configurations of an access control policy, you must provide the unique ID of the policy. You can call the [DescribeVpcFirewallControlPolicy](~~159758~~) operation to query the ID.
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The application type used in the access control policy.
	//
	// Valid values:
	//
	// *   ANY: all application types
	// *   FTP
	// *   HTTP
	// *   HTTPS
	// *   MySQL
	// *   SMTP
	// *   SMTPS
	// *   RDP
	// *   VNC
	// *   SSH
	// *   Redis
	// *   MQTT
	// *   MongoDB
	// *   Memcache
	// *   SSL
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The application names.
	ApplicationNameList []*string `json:"ApplicationNameList,omitempty" xml:"ApplicationNameList,omitempty" type:"Repeated"`
	// The description of the access control policy.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination port in the access control policy.
	DestPort *string `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	// The name of the destination port address book in the access control policy.
	DestPortGroup *string `json:"DestPortGroup,omitempty" xml:"DestPortGroup,omitempty"`
	// The type of the destination port in the access control policy.
	//
	// *   **port**: port
	// *   **group**: port address book
	DestPortType *string `json:"DestPortType,omitempty" xml:"DestPortType,omitempty"`
	// The destination address in the access control policy.
	//
	// *   If **DestinationType** is set to `net`, the value of this parameter must be a CIDR block.
	//
	//     Example: 10.2.3.0/24
	//
	// *   If **DestinationType** is set to `group`, the value of this parameter must be an address book name.
	//
	//     Example: db_group
	//
	// *   If **DestinationType** is set to `domain`, the value of this parameter must be a domain name.
	//
	//     Example: \*.aliyuncs.com
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The type of the destination address in the access control policy.
	//
	// Valid values:
	//
	// *   **net**: CIDR block
	// *   **group**: address book
	// *   **domain**: domain name
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The time when the access control policy stops taking effect. The value is a UNIX timestamp. Unit: seconds. The value must be on the hour or on the half hour, and at least 30 minutes later than the value of StartTime.
	//
	// >  If you set RepeatType to Permanent, leave this parameter empty. If you set RepeatType to None, Daily, Weekly, or Monthly, you must specify this parameter.
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the content within the response.
	//
	// Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The protocol type in the access control policy.
	//
	// Valid values:
	//
	// *   ANY: all protocol types
	// *   TCP
	// *   UDP
	// *   ICMP
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// Specifies whether to enable the access control policy. By default, an access control policy is enabled after the policy is created. Valid values:
	//
	// *   **true**: enables the access control policy.
	// *   **false**: disables the access control policy.
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The days of a week or of a month on which the access control policy takes effect.
	//
	// *   If you set RepeatType to `Permanent`, `None`, or `Daily`, the value of this parameter is an empty array. Example: \[].
	// *   If you set RepeatType to Weekly, you must specify this parameter. Example: \[0, 6].
	//
	// >  If you set RepeatType to Weekly, the fields in the value of this parameter cannot be repeated.
	//
	// *   If you set RepeatType to `Monthly`, you must specify this parameter. Example: \[1, 31].
	//
	// >  If you set RepeatType to Monthly, the fields in the value of this parameter cannot be repeated.
	RepeatDays []*int64 `json:"RepeatDays,omitempty" xml:"RepeatDays,omitempty" type:"Repeated"`
	// The point in time when the recurrence ends. Example: 23:30. The value must be on the hour or on the half hour, and at least 30 minutes later than the value of RepeatStartTime.
	//
	// >  If you set RepeatType to Permanent or None, leave this parameter empty. If you set RepeatType to Daily, Weekly, or Monthly, you must specify this parameter.
	RepeatEndTime *string `json:"RepeatEndTime,omitempty" xml:"RepeatEndTime,omitempty"`
	// The point in time when the recurrence starts. Example: 08:00. The value must be on the hour or on the half hour, and at least 30 minutes earlier than the value of RepeatEndTime.
	//
	// >  If you set RepeatType to Permanent or None, leave this parameter empty. If you set RepeatType to Daily, Weekly, or Monthly, you must specify this parameter.
	RepeatStartTime *string `json:"RepeatStartTime,omitempty" xml:"RepeatStartTime,omitempty"`
	// The recurrence type for the access control policy to take effect. Valid values:
	//
	// *   **Permanent** (default): The policy always takes effect.
	// *   **None**: The policy takes effect for only once.
	// *   **Daily**: The policy takes effect on a daily basis.
	// *   **Weekly**: The policy takes effect on a weekly basis.
	// *   **Monthly**: The policy takes effect on a monthly basis.
	RepeatType *string `json:"RepeatType,omitempty" xml:"RepeatType,omitempty"`
	// The source address in the access control policy.
	//
	// Valid values:
	//
	// *   If **SourceType** is set to `net`, the value of this parameter must be a CIDR block.
	//
	//     Example: 10.2.4.0/24
	//
	// *   If **SourceType** is set to `group`, the value of this parameter must be an address book name.
	//
	//     Example: db_group
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The type of the source address in the access control policy.
	//
	// Valid values:
	//
	// *   **net**: CIDR block
	// *   **group**: address book
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The time when the access control policy starts to take effect. The value is a UNIX timestamp. Unit: seconds. The value must be on the hour or on the half hour, and at least 30 minutes earlier than the value of EndTime.
	//
	// >  If you set RepeatType to Permanent, leave this parameter empty. If you set RepeatType to None, Daily, Weekly, or Monthly, you must specify this parameter.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The instance ID of the VPC firewall. You can call the [DescribeVpcFirewallAclGroupList](~~159760~~) operation to query the ID.
	//
	// *   If the VPC firewall is used to protect a CEN instance, the value of this parameter must be the ID of the CEN instance.
	//
	//     Example: cen-ervw0g12b5jbw\*\*\*\*
	//
	// *   If the VPC firewall is used to protect an Express Connect circuit, the value of this parameter must be the instance ID of the VPC firewall.
	//
	//     Example: vfw-a42bbb7b887148c9\*\*\*\*
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s ModifyVpcFirewallControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetAclAction(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.AclAction = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetAclUuid(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.AclUuid = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetApplicationName(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.ApplicationName = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetApplicationNameList(v []*string) *ModifyVpcFirewallControlPolicyRequest {
	s.ApplicationNameList = v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetDescription(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.Description = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetDestPort(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.DestPort = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetDestPortGroup(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.DestPortGroup = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetDestPortType(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.DestPortType = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetDestination(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.Destination = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetDestinationType(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.DestinationType = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetEndTime(v int64) *ModifyVpcFirewallControlPolicyRequest {
	s.EndTime = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetLang(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetProto(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.Proto = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetRelease(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.Release = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetRepeatDays(v []*int64) *ModifyVpcFirewallControlPolicyRequest {
	s.RepeatDays = v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetRepeatEndTime(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.RepeatEndTime = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetRepeatStartTime(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.RepeatStartTime = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetRepeatType(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.RepeatType = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetSource(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.Source = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetSourceType(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.SourceType = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetStartTime(v int64) *ModifyVpcFirewallControlPolicyRequest {
	s.StartTime = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetVpcFirewallId(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.VpcFirewallId = &v
	return s
}

type ModifyVpcFirewallControlPolicyResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVpcFirewallControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallControlPolicyResponseBody) SetRequestId(v string) *ModifyVpcFirewallControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

type ModifyVpcFirewallControlPolicyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVpcFirewallControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVpcFirewallControlPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallControlPolicyResponse) SetHeaders(v map[string]*string) *ModifyVpcFirewallControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyVpcFirewallControlPolicyResponse) SetStatusCode(v int32) *ModifyVpcFirewallControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyResponse) SetBody(v *ModifyVpcFirewallControlPolicyResponseBody) *ModifyVpcFirewallControlPolicyResponse {
	s.Body = v
	return s
}

type ModifyVpcFirewallControlPolicyPositionRequest struct {
	// The UUID of the access control policy.
	//
	// If you want to modify the configurations of an access control policy, you must provide the UUID of the policy. You can call the [DescribeVpcFirewallControlPolicy](~~159758~~) operation to query the UUID.
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The language of the content within the request and the response.
	//
	// Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The new priority of the access control policy.
	NewOrder *string `json:"NewOrder,omitempty" xml:"NewOrder,omitempty"`
	// The original priority of the access control policy.
	//
	// > This parameter is not recommended. We recommend that you use the AclUuid parameter to specify the policy that you want to modify.
	OldOrder *string `json:"OldOrder,omitempty" xml:"OldOrder,omitempty"`
	// The ID of the group to which the access control policy belongs. You can call the [DescribeVpcFirewallAclGroupList](~~159760~~) operation to query the ID.
	//
	// Valid values:
	//
	// *   If the VPC firewall is used to protect a CEN instance, the value of this parameter must be the ID of the CEN instance.
	//
	//     Example: cen-ervw0g12b5jbw\*\*\*\*
	//
	// *   If the VPC firewall is used to protect an Express Connect circuit, the value of this parameter must be the instance ID of the VPC firewall.
	//
	//     Example: vfw-a42bbb7b887148c9\*\*\*\*
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s ModifyVpcFirewallControlPolicyPositionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallControlPolicyPositionRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallControlPolicyPositionRequest) SetAclUuid(v string) *ModifyVpcFirewallControlPolicyPositionRequest {
	s.AclUuid = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyPositionRequest) SetLang(v string) *ModifyVpcFirewallControlPolicyPositionRequest {
	s.Lang = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyPositionRequest) SetNewOrder(v string) *ModifyVpcFirewallControlPolicyPositionRequest {
	s.NewOrder = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyPositionRequest) SetOldOrder(v string) *ModifyVpcFirewallControlPolicyPositionRequest {
	s.OldOrder = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyPositionRequest) SetVpcFirewallId(v string) *ModifyVpcFirewallControlPolicyPositionRequest {
	s.VpcFirewallId = &v
	return s
}

type ModifyVpcFirewallControlPolicyPositionResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVpcFirewallControlPolicyPositionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallControlPolicyPositionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallControlPolicyPositionResponseBody) SetRequestId(v string) *ModifyVpcFirewallControlPolicyPositionResponseBody {
	s.RequestId = &v
	return s
}

type ModifyVpcFirewallControlPolicyPositionResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVpcFirewallControlPolicyPositionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVpcFirewallControlPolicyPositionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallControlPolicyPositionResponse) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallControlPolicyPositionResponse) SetHeaders(v map[string]*string) *ModifyVpcFirewallControlPolicyPositionResponse {
	s.Headers = v
	return s
}

func (s *ModifyVpcFirewallControlPolicyPositionResponse) SetStatusCode(v int32) *ModifyVpcFirewallControlPolicyPositionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyPositionResponse) SetBody(v *ModifyVpcFirewallControlPolicyPositionResponseBody) *ModifyVpcFirewallControlPolicyPositionResponse {
	s.Body = v
	return s
}

type ModifyVpcFirewallDefaultIPSConfigRequest struct {
	// Specifies whether to enable basic protection. Valid values:
	//
	// *   **1**: yes
	// *   **0**: no
	BasicRules *string `json:"BasicRules,omitempty" xml:"BasicRules,omitempty"`
	// Specifies whether to enable virtual patching. Valid values:
	//
	// *   **1**: yes
	// *   **0**: no
	EnableAllPatch *string `json:"EnableAllPatch,omitempty" xml:"EnableAllPatch,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The mode of the intrusion prevention system (IPS). Valid values:
	//
	// *   **1**: block mode
	// *   **0**: monitor mode
	RunMode *string `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
	// Deprecated
	// The source IP address of the request.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The instance ID of the VPC firewall. Valid values:
	//
	// *   If the VPC firewall protects mutual access traffic between a VPC and a specified network instance that is attached to a Cloud Enterprise Network (CEN) instance, the value of this parameter is the ID of the CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance. You can call the [DescribeVpcFirewallCenList](~~345777~~) operation to query the IDs of CEN instances.
	// *   If the VPC firewall protects mutual access traffic between two VPCs that are connected by using an Express Connect circuit, the value of this parameter is the ID of the VPC firewall. You can call the [DescribeVpcFirewallList](~~342932~~) operation to query the instance IDs of VPC firewalls.
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s ModifyVpcFirewallDefaultIPSConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallDefaultIPSConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallDefaultIPSConfigRequest) SetBasicRules(v string) *ModifyVpcFirewallDefaultIPSConfigRequest {
	s.BasicRules = &v
	return s
}

func (s *ModifyVpcFirewallDefaultIPSConfigRequest) SetEnableAllPatch(v string) *ModifyVpcFirewallDefaultIPSConfigRequest {
	s.EnableAllPatch = &v
	return s
}

func (s *ModifyVpcFirewallDefaultIPSConfigRequest) SetLang(v string) *ModifyVpcFirewallDefaultIPSConfigRequest {
	s.Lang = &v
	return s
}

func (s *ModifyVpcFirewallDefaultIPSConfigRequest) SetMemberUid(v string) *ModifyVpcFirewallDefaultIPSConfigRequest {
	s.MemberUid = &v
	return s
}

func (s *ModifyVpcFirewallDefaultIPSConfigRequest) SetRunMode(v string) *ModifyVpcFirewallDefaultIPSConfigRequest {
	s.RunMode = &v
	return s
}

func (s *ModifyVpcFirewallDefaultIPSConfigRequest) SetSourceIp(v string) *ModifyVpcFirewallDefaultIPSConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyVpcFirewallDefaultIPSConfigRequest) SetVpcFirewallId(v string) *ModifyVpcFirewallDefaultIPSConfigRequest {
	s.VpcFirewallId = &v
	return s
}

type ModifyVpcFirewallDefaultIPSConfigResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVpcFirewallDefaultIPSConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallDefaultIPSConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallDefaultIPSConfigResponseBody) SetRequestId(v string) *ModifyVpcFirewallDefaultIPSConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyVpcFirewallDefaultIPSConfigResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVpcFirewallDefaultIPSConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVpcFirewallDefaultIPSConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallDefaultIPSConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallDefaultIPSConfigResponse) SetHeaders(v map[string]*string) *ModifyVpcFirewallDefaultIPSConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyVpcFirewallDefaultIPSConfigResponse) SetStatusCode(v int32) *ModifyVpcFirewallDefaultIPSConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVpcFirewallDefaultIPSConfigResponse) SetBody(v *ModifyVpcFirewallDefaultIPSConfigResponseBody) *ModifyVpcFirewallDefaultIPSConfigResponse {
	s.Body = v
	return s
}

type ModifyVpcFirewallIPSWhitelistRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// *   **zh** (default): Chinese
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The type of the list. Valid values:
	//
	// *   **1**: user-defined
	// *   **2**: address book
	ListType *int64 `json:"ListType,omitempty" xml:"ListType,omitempty"`
	// The entry in the list.
	ListValue *string `json:"ListValue,omitempty" xml:"ListValue,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	MemberUid *int64 `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The instance ID of the VPC firewall.
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
	// The type of the whitelist. Valid values:
	//
	// *   **1**: destination
	// *   **2**: source
	WhiteType *int64 `json:"WhiteType,omitempty" xml:"WhiteType,omitempty"`
}

func (s ModifyVpcFirewallIPSWhitelistRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallIPSWhitelistRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallIPSWhitelistRequest) SetLang(v string) *ModifyVpcFirewallIPSWhitelistRequest {
	s.Lang = &v
	return s
}

func (s *ModifyVpcFirewallIPSWhitelistRequest) SetListType(v int64) *ModifyVpcFirewallIPSWhitelistRequest {
	s.ListType = &v
	return s
}

func (s *ModifyVpcFirewallIPSWhitelistRequest) SetListValue(v string) *ModifyVpcFirewallIPSWhitelistRequest {
	s.ListValue = &v
	return s
}

func (s *ModifyVpcFirewallIPSWhitelistRequest) SetMemberUid(v int64) *ModifyVpcFirewallIPSWhitelistRequest {
	s.MemberUid = &v
	return s
}

func (s *ModifyVpcFirewallIPSWhitelistRequest) SetVpcFirewallId(v string) *ModifyVpcFirewallIPSWhitelistRequest {
	s.VpcFirewallId = &v
	return s
}

func (s *ModifyVpcFirewallIPSWhitelistRequest) SetWhiteType(v int64) *ModifyVpcFirewallIPSWhitelistRequest {
	s.WhiteType = &v
	return s
}

type ModifyVpcFirewallIPSWhitelistResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVpcFirewallIPSWhitelistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallIPSWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallIPSWhitelistResponseBody) SetRequestId(v string) *ModifyVpcFirewallIPSWhitelistResponseBody {
	s.RequestId = &v
	return s
}

type ModifyVpcFirewallIPSWhitelistResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVpcFirewallIPSWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVpcFirewallIPSWhitelistResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallIPSWhitelistResponse) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallIPSWhitelistResponse) SetHeaders(v map[string]*string) *ModifyVpcFirewallIPSWhitelistResponse {
	s.Headers = v
	return s
}

func (s *ModifyVpcFirewallIPSWhitelistResponse) SetStatusCode(v int32) *ModifyVpcFirewallIPSWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVpcFirewallIPSWhitelistResponse) SetBody(v *ModifyVpcFirewallIPSWhitelistResponseBody) *ModifyVpcFirewallIPSWhitelistResponse {
	s.Body = v
	return s
}

type ModifyVpcFirewallSwitchStatusRequest struct {
	// Specifies whether to enable the VPC firewall. Valid values:
	//
	// *   **open**: yes
	// *   **close**: no
	FirewallSwitch *string `json:"FirewallSwitch,omitempty" xml:"FirewallSwitch,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The instance ID of the VPC firewall.
	//
	// > You can call the [DescribeVpcFirewallList](~~342932~~) operation to query the instance IDs of VPC firewalls.
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s ModifyVpcFirewallSwitchStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallSwitchStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallSwitchStatusRequest) SetFirewallSwitch(v string) *ModifyVpcFirewallSwitchStatusRequest {
	s.FirewallSwitch = &v
	return s
}

func (s *ModifyVpcFirewallSwitchStatusRequest) SetLang(v string) *ModifyVpcFirewallSwitchStatusRequest {
	s.Lang = &v
	return s
}

func (s *ModifyVpcFirewallSwitchStatusRequest) SetMemberUid(v string) *ModifyVpcFirewallSwitchStatusRequest {
	s.MemberUid = &v
	return s
}

func (s *ModifyVpcFirewallSwitchStatusRequest) SetVpcFirewallId(v string) *ModifyVpcFirewallSwitchStatusRequest {
	s.VpcFirewallId = &v
	return s
}

type ModifyVpcFirewallSwitchStatusResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVpcFirewallSwitchStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallSwitchStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallSwitchStatusResponseBody) SetRequestId(v string) *ModifyVpcFirewallSwitchStatusResponseBody {
	s.RequestId = &v
	return s
}

type ModifyVpcFirewallSwitchStatusResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVpcFirewallSwitchStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVpcFirewallSwitchStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallSwitchStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallSwitchStatusResponse) SetHeaders(v map[string]*string) *ModifyVpcFirewallSwitchStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyVpcFirewallSwitchStatusResponse) SetStatusCode(v int32) *ModifyVpcFirewallSwitchStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVpcFirewallSwitchStatusResponse) SetBody(v *ModifyVpcFirewallSwitchStatusResponseBody) *ModifyVpcFirewallSwitchStatusResponse {
	s.Body = v
	return s
}

type PutDisableAllFwSwitchRequest struct {
	// The instance ID of your Cloud Firewall.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The natural language of the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Deprecated
	// The source IP address of the request.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s PutDisableAllFwSwitchRequest) String() string {
	return tea.Prettify(s)
}

func (s PutDisableAllFwSwitchRequest) GoString() string {
	return s.String()
}

func (s *PutDisableAllFwSwitchRequest) SetInstanceId(v string) *PutDisableAllFwSwitchRequest {
	s.InstanceId = &v
	return s
}

func (s *PutDisableAllFwSwitchRequest) SetLang(v string) *PutDisableAllFwSwitchRequest {
	s.Lang = &v
	return s
}

func (s *PutDisableAllFwSwitchRequest) SetSourceIp(v string) *PutDisableAllFwSwitchRequest {
	s.SourceIp = &v
	return s
}

type PutDisableAllFwSwitchResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutDisableAllFwSwitchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutDisableAllFwSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *PutDisableAllFwSwitchResponseBody) SetRequestId(v string) *PutDisableAllFwSwitchResponseBody {
	s.RequestId = &v
	return s
}

type PutDisableAllFwSwitchResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutDisableAllFwSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutDisableAllFwSwitchResponse) String() string {
	return tea.Prettify(s)
}

func (s PutDisableAllFwSwitchResponse) GoString() string {
	return s.String()
}

func (s *PutDisableAllFwSwitchResponse) SetHeaders(v map[string]*string) *PutDisableAllFwSwitchResponse {
	s.Headers = v
	return s
}

func (s *PutDisableAllFwSwitchResponse) SetStatusCode(v int32) *PutDisableAllFwSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *PutDisableAllFwSwitchResponse) SetBody(v *PutDisableAllFwSwitchResponseBody) *PutDisableAllFwSwitchResponse {
	s.Body = v
	return s
}

type PutDisableFwSwitchRequest struct {
	// The IP addresses.
	//
	// >  You must specify at least one of the IpaddrList, RegionList, and ResourceTypeList parameters.
	IpaddrList []*string `json:"IpaddrList,omitempty" xml:"IpaddrList,omitempty" type:"Repeated"`
	// The language of the content within the response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The regions.
	//
	// >  You must specify at least one of the IpaddrList, RegionList, and ResourceTypeList parameters.
	RegionList []*string `json:"RegionList,omitempty" xml:"RegionList,omitempty" type:"Repeated"`
	// The types of the assets.
	//
	// > You must specify at least one of the IpaddrList, RegionList, and ResourceTypeList parameters.
	ResourceTypeList []*string `json:"ResourceTypeList,omitempty" xml:"ResourceTypeList,omitempty" type:"Repeated"`
	// Deprecated
	// The source IP address of the request.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s PutDisableFwSwitchRequest) String() string {
	return tea.Prettify(s)
}

func (s PutDisableFwSwitchRequest) GoString() string {
	return s.String()
}

func (s *PutDisableFwSwitchRequest) SetIpaddrList(v []*string) *PutDisableFwSwitchRequest {
	s.IpaddrList = v
	return s
}

func (s *PutDisableFwSwitchRequest) SetLang(v string) *PutDisableFwSwitchRequest {
	s.Lang = &v
	return s
}

func (s *PutDisableFwSwitchRequest) SetRegionList(v []*string) *PutDisableFwSwitchRequest {
	s.RegionList = v
	return s
}

func (s *PutDisableFwSwitchRequest) SetResourceTypeList(v []*string) *PutDisableFwSwitchRequest {
	s.ResourceTypeList = v
	return s
}

func (s *PutDisableFwSwitchRequest) SetSourceIp(v string) *PutDisableFwSwitchRequest {
	s.SourceIp = &v
	return s
}

type PutDisableFwSwitchResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutDisableFwSwitchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutDisableFwSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *PutDisableFwSwitchResponseBody) SetRequestId(v string) *PutDisableFwSwitchResponseBody {
	s.RequestId = &v
	return s
}

type PutDisableFwSwitchResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutDisableFwSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutDisableFwSwitchResponse) String() string {
	return tea.Prettify(s)
}

func (s PutDisableFwSwitchResponse) GoString() string {
	return s.String()
}

func (s *PutDisableFwSwitchResponse) SetHeaders(v map[string]*string) *PutDisableFwSwitchResponse {
	s.Headers = v
	return s
}

func (s *PutDisableFwSwitchResponse) SetStatusCode(v int32) *PutDisableFwSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *PutDisableFwSwitchResponse) SetBody(v *PutDisableFwSwitchResponseBody) *PutDisableFwSwitchResponse {
	s.Body = v
	return s
}

type PutEnableAllFwSwitchRequest struct {
	// The instance ID of your Cloud Firewall.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// *   **zh**: Chinese (default)
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Deprecated
	// The source IP address of the request.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s PutEnableAllFwSwitchRequest) String() string {
	return tea.Prettify(s)
}

func (s PutEnableAllFwSwitchRequest) GoString() string {
	return s.String()
}

func (s *PutEnableAllFwSwitchRequest) SetInstanceId(v string) *PutEnableAllFwSwitchRequest {
	s.InstanceId = &v
	return s
}

func (s *PutEnableAllFwSwitchRequest) SetLang(v string) *PutEnableAllFwSwitchRequest {
	s.Lang = &v
	return s
}

func (s *PutEnableAllFwSwitchRequest) SetSourceIp(v string) *PutEnableAllFwSwitchRequest {
	s.SourceIp = &v
	return s
}

type PutEnableAllFwSwitchResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutEnableAllFwSwitchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutEnableAllFwSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *PutEnableAllFwSwitchResponseBody) SetRequestId(v string) *PutEnableAllFwSwitchResponseBody {
	s.RequestId = &v
	return s
}

type PutEnableAllFwSwitchResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutEnableAllFwSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutEnableAllFwSwitchResponse) String() string {
	return tea.Prettify(s)
}

func (s PutEnableAllFwSwitchResponse) GoString() string {
	return s.String()
}

func (s *PutEnableAllFwSwitchResponse) SetHeaders(v map[string]*string) *PutEnableAllFwSwitchResponse {
	s.Headers = v
	return s
}

func (s *PutEnableAllFwSwitchResponse) SetStatusCode(v int32) *PutEnableAllFwSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *PutEnableAllFwSwitchResponse) SetBody(v *PutEnableAllFwSwitchResponseBody) *PutEnableAllFwSwitchResponse {
	s.Body = v
	return s
}

type PutEnableFwSwitchRequest struct {
	// The IP addresses.
	//
	// > You must specify at least one of the IpaddrList, RegionList, and ResourceTypeList parameters.
	IpaddrList []*string `json:"IpaddrList,omitempty" xml:"IpaddrList,omitempty" type:"Repeated"`
	// The language of the content within the response.
	//
	// *   **zh**: Chinese
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The regions.
	//
	// > You must specify at least one of the IpaddrList, RegionList, and ResourceTypeList parameters.
	RegionList []*string `json:"RegionList,omitempty" xml:"RegionList,omitempty" type:"Repeated"`
	// The types of the assets.
	//
	// Valid values:
	//
	// *   BastionHostIP: the egress IP address of a bastion host
	// *   BastionHostIngressIP: the ingress IP address of a bastion host
	// *   EcsEIP: the elastic IP address (EIP) of an Elastic Compute Service (ECS) instance
	// *   EcsPublicIP: the public IP address of an ECS instance
	// *   EIP: the EIP
	// *   EniEIP: the EIP of an elastic network interface (ENI)
	// *   NatEIP: the EIP of a NAT gateway
	// *   SlbEIP: the EIP of a Server Load Balancer (SLB) instance
	// *   SlbPublicIP: the public IP address of an SLB instance
	// *   NatPublicIP: the public IP address of a NAT gateway
	// *   HAVIP: the high-availability virtual IP address (HAVIP)
	//
	// > You must specify at least one of the IpaddrList, RegionList, and ResourceTypeList parameters.
	ResourceTypeList []*string `json:"ResourceTypeList,omitempty" xml:"ResourceTypeList,omitempty" type:"Repeated"`
	// Deprecated
	// The source IP address of the request.
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s PutEnableFwSwitchRequest) String() string {
	return tea.Prettify(s)
}

func (s PutEnableFwSwitchRequest) GoString() string {
	return s.String()
}

func (s *PutEnableFwSwitchRequest) SetIpaddrList(v []*string) *PutEnableFwSwitchRequest {
	s.IpaddrList = v
	return s
}

func (s *PutEnableFwSwitchRequest) SetLang(v string) *PutEnableFwSwitchRequest {
	s.Lang = &v
	return s
}

func (s *PutEnableFwSwitchRequest) SetRegionList(v []*string) *PutEnableFwSwitchRequest {
	s.RegionList = v
	return s
}

func (s *PutEnableFwSwitchRequest) SetResourceTypeList(v []*string) *PutEnableFwSwitchRequest {
	s.ResourceTypeList = v
	return s
}

func (s *PutEnableFwSwitchRequest) SetSourceIp(v string) *PutEnableFwSwitchRequest {
	s.SourceIp = &v
	return s
}

type PutEnableFwSwitchResponseBody struct {
	// The status information of the asset when it is not synchronized to Cloud Firewall.
	AbnormalResourceStatusList []*PutEnableFwSwitchResponseBodyAbnormalResourceStatusList `json:"AbnormalResourceStatusList,omitempty" xml:"AbnormalResourceStatusList,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutEnableFwSwitchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutEnableFwSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *PutEnableFwSwitchResponseBody) SetAbnormalResourceStatusList(v []*PutEnableFwSwitchResponseBodyAbnormalResourceStatusList) *PutEnableFwSwitchResponseBody {
	s.AbnormalResourceStatusList = v
	return s
}

func (s *PutEnableFwSwitchResponseBody) SetRequestId(v string) *PutEnableFwSwitchResponseBody {
	s.RequestId = &v
	return s
}

type PutEnableFwSwitchResponseBodyAbnormalResourceStatusList struct {
	// The message displayed when the asset is not synchronized to Cloud Firewall. Valid values:
	//
	// *   cloudfirewall do not sync this ip address: This IP address is not synchronized to Cloud Firewall.
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The IP address of the asset.
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The status of the asset when it is not synchronized to Cloud Firewall. Valid values:
	//
	// *   ip_not_sync: The asset is not synchronized.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PutEnableFwSwitchResponseBodyAbnormalResourceStatusList) String() string {
	return tea.Prettify(s)
}

func (s PutEnableFwSwitchResponseBodyAbnormalResourceStatusList) GoString() string {
	return s.String()
}

func (s *PutEnableFwSwitchResponseBodyAbnormalResourceStatusList) SetMsg(v string) *PutEnableFwSwitchResponseBodyAbnormalResourceStatusList {
	s.Msg = &v
	return s
}

func (s *PutEnableFwSwitchResponseBodyAbnormalResourceStatusList) SetResource(v string) *PutEnableFwSwitchResponseBodyAbnormalResourceStatusList {
	s.Resource = &v
	return s
}

func (s *PutEnableFwSwitchResponseBodyAbnormalResourceStatusList) SetStatus(v string) *PutEnableFwSwitchResponseBodyAbnormalResourceStatusList {
	s.Status = &v
	return s
}

type PutEnableFwSwitchResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutEnableFwSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutEnableFwSwitchResponse) String() string {
	return tea.Prettify(s)
}

func (s PutEnableFwSwitchResponse) GoString() string {
	return s.String()
}

func (s *PutEnableFwSwitchResponse) SetHeaders(v map[string]*string) *PutEnableFwSwitchResponse {
	s.Headers = v
	return s
}

func (s *PutEnableFwSwitchResponse) SetStatusCode(v int32) *PutEnableFwSwitchResponse {
	s.StatusCode = &v
	return s
}

func (s *PutEnableFwSwitchResponse) SetBody(v *PutEnableFwSwitchResponseBody) *PutEnableFwSwitchResponse {
	s.Body = v
	return s
}

type ReleasePostInstanceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ReleasePostInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleasePostInstanceRequest) GoString() string {
	return s.String()
}

func (s *ReleasePostInstanceRequest) SetInstanceId(v string) *ReleasePostInstanceRequest {
	s.InstanceId = &v
	return s
}

type ReleasePostInstanceResponseBody struct {
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	ReleaseStatus  *bool   `json:"ReleaseStatus,omitempty" xml:"ReleaseStatus,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReleasePostInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleasePostInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ReleasePostInstanceResponseBody) SetHttpStatusCode(v int32) *ReleasePostInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ReleasePostInstanceResponseBody) SetReleaseStatus(v bool) *ReleasePostInstanceResponseBody {
	s.ReleaseStatus = &v
	return s
}

func (s *ReleasePostInstanceResponseBody) SetRequestId(v string) *ReleasePostInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleasePostInstanceResponseBody) SetSuccess(v bool) *ReleasePostInstanceResponseBody {
	s.Success = &v
	return s
}

type ReleasePostInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleasePostInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleasePostInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleasePostInstanceResponse) GoString() string {
	return s.String()
}

func (s *ReleasePostInstanceResponse) SetHeaders(v map[string]*string) *ReleasePostInstanceResponse {
	s.Headers = v
	return s
}

func (s *ReleasePostInstanceResponse) SetStatusCode(v int32) *ReleasePostInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleasePostInstanceResponse) SetBody(v *ReleasePostInstanceResponseBody) *ReleasePostInstanceResponse {
	s.Body = v
	return s
}

type ResetNatFirewallRuleHitCountRequest struct {
	// The UUID of the access control policy.
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// *   **zh** (default): Chinese
	// *   **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the NAT gateway.
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
}

func (s ResetNatFirewallRuleHitCountRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetNatFirewallRuleHitCountRequest) GoString() string {
	return s.String()
}

func (s *ResetNatFirewallRuleHitCountRequest) SetAclUuid(v string) *ResetNatFirewallRuleHitCountRequest {
	s.AclUuid = &v
	return s
}

func (s *ResetNatFirewallRuleHitCountRequest) SetLang(v string) *ResetNatFirewallRuleHitCountRequest {
	s.Lang = &v
	return s
}

func (s *ResetNatFirewallRuleHitCountRequest) SetNatGatewayId(v string) *ResetNatFirewallRuleHitCountRequest {
	s.NatGatewayId = &v
	return s
}

type ResetNatFirewallRuleHitCountResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetNatFirewallRuleHitCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetNatFirewallRuleHitCountResponseBody) GoString() string {
	return s.String()
}

func (s *ResetNatFirewallRuleHitCountResponseBody) SetRequestId(v string) *ResetNatFirewallRuleHitCountResponseBody {
	s.RequestId = &v
	return s
}

type ResetNatFirewallRuleHitCountResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetNatFirewallRuleHitCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetNatFirewallRuleHitCountResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetNatFirewallRuleHitCountResponse) GoString() string {
	return s.String()
}

func (s *ResetNatFirewallRuleHitCountResponse) SetHeaders(v map[string]*string) *ResetNatFirewallRuleHitCountResponse {
	s.Headers = v
	return s
}

func (s *ResetNatFirewallRuleHitCountResponse) SetStatusCode(v int32) *ResetNatFirewallRuleHitCountResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetNatFirewallRuleHitCountResponse) SetBody(v *ResetNatFirewallRuleHitCountResponseBody) *ResetNatFirewallRuleHitCountResponse {
	s.Body = v
	return s
}

type ResetVpcFirewallRuleHitCountRequest struct {
	// The ID of the access control policy.
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The natural language of the request and response.
	//
	// Valid values:
	//
	// - **zh**: Chinese (default)
	// - **en**: English
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s ResetVpcFirewallRuleHitCountRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetVpcFirewallRuleHitCountRequest) GoString() string {
	return s.String()
}

func (s *ResetVpcFirewallRuleHitCountRequest) SetAclUuid(v string) *ResetVpcFirewallRuleHitCountRequest {
	s.AclUuid = &v
	return s
}

func (s *ResetVpcFirewallRuleHitCountRequest) SetLang(v string) *ResetVpcFirewallRuleHitCountRequest {
	s.Lang = &v
	return s
}

type ResetVpcFirewallRuleHitCountResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetVpcFirewallRuleHitCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetVpcFirewallRuleHitCountResponseBody) GoString() string {
	return s.String()
}

func (s *ResetVpcFirewallRuleHitCountResponseBody) SetRequestId(v string) *ResetVpcFirewallRuleHitCountResponseBody {
	s.RequestId = &v
	return s
}

type ResetVpcFirewallRuleHitCountResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetVpcFirewallRuleHitCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetVpcFirewallRuleHitCountResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetVpcFirewallRuleHitCountResponse) GoString() string {
	return s.String()
}

func (s *ResetVpcFirewallRuleHitCountResponse) SetHeaders(v map[string]*string) *ResetVpcFirewallRuleHitCountResponse {
	s.Headers = v
	return s
}

func (s *ResetVpcFirewallRuleHitCountResponse) SetStatusCode(v int32) *ResetVpcFirewallRuleHitCountResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetVpcFirewallRuleHitCountResponse) SetBody(v *ResetVpcFirewallRuleHitCountResponseBody) *ResetVpcFirewallRuleHitCountResponse {
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
		"ap-southeast-1": tea.String("cloudfw.ap-southeast-1.aliyuncs.com"),
		"cn-hangzhou":    tea.String("cloudfw.cn-hangzhou.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("cloudfw"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

/**
 * You can call the AddAddressBook operation to create an address book for access control. The address book can be an IP address book, an ECS tag-based address book, a port address book, or a domain address book.
 * ## [](#qps)Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request AddAddressBookRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AddAddressBookResponse
 */
func (client *Client) AddAddressBookWithOptions(request *AddAddressBookRequest, runtime *util.RuntimeOptions) (_result *AddAddressBookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddressList)) {
		query["AddressList"] = request.AddressList
	}

	if !tea.BoolValue(util.IsUnset(request.AutoAddTagEcs)) {
		query["AutoAddTagEcs"] = request.AutoAddTagEcs
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupType)) {
		query["GroupType"] = request.GroupType
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.TagList)) {
		query["TagList"] = request.TagList
	}

	if !tea.BoolValue(util.IsUnset(request.TagRelation)) {
		query["TagRelation"] = request.TagRelation
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddAddressBook"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddAddressBookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the AddAddressBook operation to create an address book for access control. The address book can be an IP address book, an ECS tag-based address book, a port address book, or a domain address book.
 * ## [](#qps)Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request AddAddressBookRequest
 * @return AddAddressBookResponse
 */
func (client *Client) AddAddressBook(request *AddAddressBookRequest) (_result *AddAddressBookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddAddressBookResponse{}
	_body, _err := client.AddAddressBookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the AddControlPolicy operation to create an access control policy to allow, block, or monitor traffic that reaches Cloud Firewall.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request AddControlPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AddControlPolicyResponse
 */
func (client *Client) AddControlPolicyWithOptions(request *AddControlPolicyRequest, runtime *util.RuntimeOptions) (_result *AddControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclAction)) {
		query["AclAction"] = request.AclAction
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicationName)) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicationNameList)) {
		query["ApplicationNameList"] = request.ApplicationNameList
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DestPort)) {
		query["DestPort"] = request.DestPort
	}

	if !tea.BoolValue(util.IsUnset(request.DestPortGroup)) {
		query["DestPortGroup"] = request.DestPortGroup
	}

	if !tea.BoolValue(util.IsUnset(request.DestPortType)) {
		query["DestPortType"] = request.DestPortType
	}

	if !tea.BoolValue(util.IsUnset(request.Destination)) {
		query["Destination"] = request.Destination
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationType)) {
		query["DestinationType"] = request.DestinationType
	}

	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		query["Direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.IpVersion)) {
		query["IpVersion"] = request.IpVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.NewOrder)) {
		query["NewOrder"] = request.NewOrder
	}

	if !tea.BoolValue(util.IsUnset(request.Proto)) {
		query["Proto"] = request.Proto
	}

	if !tea.BoolValue(util.IsUnset(request.Release)) {
		query["Release"] = request.Release
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatDays)) {
		query["RepeatDays"] = request.RepeatDays
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatEndTime)) {
		query["RepeatEndTime"] = request.RepeatEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatStartTime)) {
		query["RepeatStartTime"] = request.RepeatStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatType)) {
		query["RepeatType"] = request.RepeatType
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddControlPolicy"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the AddControlPolicy operation to create an access control policy to allow, block, or monitor traffic that reaches Cloud Firewall.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request AddControlPolicyRequest
 * @return AddControlPolicyResponse
 */
func (client *Client) AddControlPolicy(request *AddControlPolicyRequest) (_result *AddControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddControlPolicyResponse{}
	_body, _err := client.AddControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the AddInstanceMembers operation to add members to Cloud Firewall.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request AddInstanceMembersRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AddInstanceMembersResponse
 */
func (client *Client) AddInstanceMembersWithOptions(request *AddInstanceMembersRequest, runtime *util.RuntimeOptions) (_result *AddInstanceMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Members)) {
		query["Members"] = request.Members
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddInstanceMembers"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddInstanceMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the AddInstanceMembers operation to add members to Cloud Firewall.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request AddInstanceMembersRequest
 * @return AddInstanceMembersResponse
 */
func (client *Client) AddInstanceMembers(request *AddInstanceMembersRequest) (_result *AddInstanceMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddInstanceMembersResponse{}
	_body, _err := client.AddInstanceMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the BatchCopyVpcFirewallControlPolicy operation to copy all access control policies from a policy group of a source VPC firewall to a policy group of a destination VPC firewall.
 * Before you call this operation, we recommend that you back up access control policies. For more information about how to back up an access control policy, see [Back up an access control policy](https://www.alibabacloud.com/help/en/cloud-firewall/latest/back-up-and-roll-back-an-access-control-policy).
 * After you call this operation, all the access control policies in the policy group of the destination VPC firewall are replaced.
 * The policy groups of the source VPC firewall and the destination VPC firewall must belong to the same Alibaba Cloud account.
 * ## Limits
 * You can call this operation up to 10 times per second per account. When the number of calls to this operation per second exceeds the limit, throttling is triggered. Throttling may affect your business. We recommend that you take note of the limit on this operation.
 *
 * @param request BatchCopyVpcFirewallControlPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return BatchCopyVpcFirewallControlPolicyResponse
 */
func (client *Client) BatchCopyVpcFirewallControlPolicyWithOptions(request *BatchCopyVpcFirewallControlPolicyRequest, runtime *util.RuntimeOptions) (_result *BatchCopyVpcFirewallControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.SourceVpcFirewallId)) {
		query["SourceVpcFirewallId"] = request.SourceVpcFirewallId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetVpcFirewallId)) {
		query["TargetVpcFirewallId"] = request.TargetVpcFirewallId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchCopyVpcFirewallControlPolicy"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchCopyVpcFirewallControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the BatchCopyVpcFirewallControlPolicy operation to copy all access control policies from a policy group of a source VPC firewall to a policy group of a destination VPC firewall.
 * Before you call this operation, we recommend that you back up access control policies. For more information about how to back up an access control policy, see [Back up an access control policy](https://www.alibabacloud.com/help/en/cloud-firewall/latest/back-up-and-roll-back-an-access-control-policy).
 * After you call this operation, all the access control policies in the policy group of the destination VPC firewall are replaced.
 * The policy groups of the source VPC firewall and the destination VPC firewall must belong to the same Alibaba Cloud account.
 * ## Limits
 * You can call this operation up to 10 times per second per account. When the number of calls to this operation per second exceeds the limit, throttling is triggered. Throttling may affect your business. We recommend that you take note of the limit on this operation.
 *
 * @param request BatchCopyVpcFirewallControlPolicyRequest
 * @return BatchCopyVpcFirewallControlPolicyResponse
 */
func (client *Client) BatchCopyVpcFirewallControlPolicy(request *BatchCopyVpcFirewallControlPolicyRequest) (_result *BatchCopyVpcFirewallControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchCopyVpcFirewallControlPolicyResponse{}
	_body, _err := client.BatchCopyVpcFirewallControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchDeleteVpcFirewallControlPolicyWithOptions(request *BatchDeleteVpcFirewallControlPolicyRequest, runtime *util.RuntimeOptions) (_result *BatchDeleteVpcFirewallControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclUuidList)) {
		query["AclUuidList"] = request.AclUuidList
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallId)) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchDeleteVpcFirewallControlPolicy"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchDeleteVpcFirewallControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchDeleteVpcFirewallControlPolicy(request *BatchDeleteVpcFirewallControlPolicyRequest) (_result *BatchDeleteVpcFirewallControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchDeleteVpcFirewallControlPolicyResponse{}
	_body, _err := client.BatchDeleteVpcFirewallControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDownloadTaskWithOptions(request *CreateDownloadTaskRequest, runtime *util.RuntimeOptions) (_result *CreateDownloadTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.TaskData)) {
		query["TaskData"] = request.TaskData
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDownloadTask"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDownloadTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDownloadTask(request *CreateDownloadTaskRequest) (_result *CreateDownloadTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDownloadTaskResponse{}
	_body, _err := client.CreateDownloadTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to create a policy that allows, denies, or monitors the traffic that passes through the NAT firewall.
 *
 * @param request CreateNatFirewallControlPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateNatFirewallControlPolicyResponse
 */
func (client *Client) CreateNatFirewallControlPolicyWithOptions(request *CreateNatFirewallControlPolicyRequest, runtime *util.RuntimeOptions) (_result *CreateNatFirewallControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclAction)) {
		query["AclAction"] = request.AclAction
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicationNameList)) {
		query["ApplicationNameList"] = request.ApplicationNameList
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DestPort)) {
		query["DestPort"] = request.DestPort
	}

	if !tea.BoolValue(util.IsUnset(request.DestPortGroup)) {
		query["DestPortGroup"] = request.DestPortGroup
	}

	if !tea.BoolValue(util.IsUnset(request.DestPortType)) {
		query["DestPortType"] = request.DestPortType
	}

	if !tea.BoolValue(util.IsUnset(request.Destination)) {
		query["Destination"] = request.Destination
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationType)) {
		query["DestinationType"] = request.DestinationType
	}

	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		query["Direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.DomainResolveType)) {
		query["DomainResolveType"] = request.DomainResolveType
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.IpVersion)) {
		query["IpVersion"] = request.IpVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.NatGatewayId)) {
		query["NatGatewayId"] = request.NatGatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.NewOrder)) {
		query["NewOrder"] = request.NewOrder
	}

	if !tea.BoolValue(util.IsUnset(request.Proto)) {
		query["Proto"] = request.Proto
	}

	if !tea.BoolValue(util.IsUnset(request.Release)) {
		query["Release"] = request.Release
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatDays)) {
		query["RepeatDays"] = request.RepeatDays
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatEndTime)) {
		query["RepeatEndTime"] = request.RepeatEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatStartTime)) {
		query["RepeatStartTime"] = request.RepeatStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatType)) {
		query["RepeatType"] = request.RepeatType
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateNatFirewallControlPolicy"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateNatFirewallControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to create a policy that allows, denies, or monitors the traffic that passes through the NAT firewall.
 *
 * @param request CreateNatFirewallControlPolicyRequest
 * @return CreateNatFirewallControlPolicyResponse
 */
func (client *Client) CreateNatFirewallControlPolicy(request *CreateNatFirewallControlPolicyRequest) (_result *CreateNatFirewallControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateNatFirewallControlPolicyResponse{}
	_body, _err := client.CreateNatFirewallControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSecurityProxyWithOptions(request *CreateSecurityProxyRequest, runtime *util.RuntimeOptions) (_result *CreateSecurityProxyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FirewallSwitch)) {
		query["FirewallSwitch"] = request.FirewallSwitch
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.NatGatewayId)) {
		query["NatGatewayId"] = request.NatGatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.NatRouteEntryList)) {
		query["NatRouteEntryList"] = request.NatRouteEntryList
	}

	if !tea.BoolValue(util.IsUnset(request.ProxyName)) {
		query["ProxyName"] = request.ProxyName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionNo)) {
		query["RegionNo"] = request.RegionNo
	}

	if !tea.BoolValue(util.IsUnset(request.StrictMode)) {
		query["StrictMode"] = request.StrictMode
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.VswitchAuto)) {
		query["VswitchAuto"] = request.VswitchAuto
	}

	if !tea.BoolValue(util.IsUnset(request.VswitchCidr)) {
		query["VswitchCidr"] = request.VswitchCidr
	}

	if !tea.BoolValue(util.IsUnset(request.VswitchId)) {
		query["VswitchId"] = request.VswitchId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSecurityProxy"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSecurityProxyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSecurityProxy(request *CreateSecurityProxyRequest) (_result *CreateSecurityProxyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSecurityProxyResponse{}
	_body, _err := client.CreateSecurityProxyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTrFirewallV2WithOptions(request *CreateTrFirewallV2Request, runtime *util.RuntimeOptions) (_result *CreateTrFirewallV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CenId)) {
		query["CenId"] = request.CenId
	}

	if !tea.BoolValue(util.IsUnset(request.FirewallDescription)) {
		query["FirewallDescription"] = request.FirewallDescription
	}

	if !tea.BoolValue(util.IsUnset(request.FirewallName)) {
		query["FirewallName"] = request.FirewallName
	}

	if !tea.BoolValue(util.IsUnset(request.FirewallSubnetCidr)) {
		query["FirewallSubnetCidr"] = request.FirewallSubnetCidr
	}

	if !tea.BoolValue(util.IsUnset(request.FirewallVpcCidr)) {
		query["FirewallVpcCidr"] = request.FirewallVpcCidr
	}

	if !tea.BoolValue(util.IsUnset(request.FirewallVpcId)) {
		query["FirewallVpcId"] = request.FirewallVpcId
	}

	if !tea.BoolValue(util.IsUnset(request.FirewallVswitchId)) {
		query["FirewallVswitchId"] = request.FirewallVswitchId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.RegionNo)) {
		query["RegionNo"] = request.RegionNo
	}

	if !tea.BoolValue(util.IsUnset(request.RouteMode)) {
		query["RouteMode"] = request.RouteMode
	}

	if !tea.BoolValue(util.IsUnset(request.TrAttachmentMasterCidr)) {
		query["TrAttachmentMasterCidr"] = request.TrAttachmentMasterCidr
	}

	if !tea.BoolValue(util.IsUnset(request.TrAttachmentMasterZone)) {
		query["TrAttachmentMasterZone"] = request.TrAttachmentMasterZone
	}

	if !tea.BoolValue(util.IsUnset(request.TrAttachmentSlaveCidr)) {
		query["TrAttachmentSlaveCidr"] = request.TrAttachmentSlaveCidr
	}

	if !tea.BoolValue(util.IsUnset(request.TrAttachmentSlaveZone)) {
		query["TrAttachmentSlaveZone"] = request.TrAttachmentSlaveZone
	}

	if !tea.BoolValue(util.IsUnset(request.TransitRouterId)) {
		query["TransitRouterId"] = request.TransitRouterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTrFirewallV2"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTrFirewallV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTrFirewallV2(request *CreateTrFirewallV2Request) (_result *CreateTrFirewallV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTrFirewallV2Response{}
	_body, _err := client.CreateTrFirewallV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTrFirewallV2RoutePolicyWithOptions(tmpReq *CreateTrFirewallV2RoutePolicyRequest, runtime *util.RuntimeOptions) (_result *CreateTrFirewallV2RoutePolicyResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateTrFirewallV2RoutePolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DestCandidateList)) {
		request.DestCandidateListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DestCandidateList, tea.String("DestCandidateList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SrcCandidateList)) {
		request.SrcCandidateListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SrcCandidateList, tea.String("SrcCandidateList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DestCandidateListShrink)) {
		query["DestCandidateList"] = request.DestCandidateListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.FirewallId)) {
		query["FirewallId"] = request.FirewallId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyDescription)) {
		query["PolicyDescription"] = request.PolicyDescription
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyName)) {
		query["PolicyName"] = request.PolicyName
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyType)) {
		query["PolicyType"] = request.PolicyType
	}

	if !tea.BoolValue(util.IsUnset(request.SrcCandidateListShrink)) {
		query["SrcCandidateList"] = request.SrcCandidateListShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTrFirewallV2RoutePolicy"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTrFirewallV2RoutePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTrFirewallV2RoutePolicy(request *CreateTrFirewallV2RoutePolicyRequest) (_result *CreateTrFirewallV2RoutePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTrFirewallV2RoutePolicyResponse{}
	_body, _err := client.CreateTrFirewallV2RoutePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the CreateVpcFirewallCenConfigure operation to create a VPC firewall. The VPC firewall protects mutual access traffic between a specified VPC and a network instance that is attached to a CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance. The VPC firewall cannot protect mutual access traffic between VBRs, between CCN instances, or between VBRs and CCN instances. For more information, see [VPC firewall limits](~~172295~~).
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request CreateVpcFirewallCenConfigureRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateVpcFirewallCenConfigureResponse
 */
func (client *Client) CreateVpcFirewallCenConfigureWithOptions(request *CreateVpcFirewallCenConfigureRequest, runtime *util.RuntimeOptions) (_result *CreateVpcFirewallCenConfigureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CenId)) {
		query["CenId"] = request.CenId
	}

	if !tea.BoolValue(util.IsUnset(request.FirewallSwitch)) {
		query["FirewallSwitch"] = request.FirewallSwitch
	}

	if !tea.BoolValue(util.IsUnset(request.FirewallVSwitchCidrBlock)) {
		query["FirewallVSwitchCidrBlock"] = request.FirewallVSwitchCidrBlock
	}

	if !tea.BoolValue(util.IsUnset(request.FirewallVpcCidrBlock)) {
		query["FirewallVpcCidrBlock"] = request.FirewallVpcCidrBlock
	}

	if !tea.BoolValue(util.IsUnset(request.FirewallVpcZoneId)) {
		query["FirewallVpcZoneId"] = request.FirewallVpcZoneId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkInstanceId)) {
		query["NetworkInstanceId"] = request.NetworkInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallName)) {
		query["VpcFirewallName"] = request.VpcFirewallName
	}

	if !tea.BoolValue(util.IsUnset(request.VpcRegion)) {
		query["VpcRegion"] = request.VpcRegion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVpcFirewallCenConfigure"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVpcFirewallCenConfigureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the CreateVpcFirewallCenConfigure operation to create a VPC firewall. The VPC firewall protects mutual access traffic between a specified VPC and a network instance that is attached to a CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance. The VPC firewall cannot protect mutual access traffic between VBRs, between CCN instances, or between VBRs and CCN instances. For more information, see [VPC firewall limits](~~172295~~).
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request CreateVpcFirewallCenConfigureRequest
 * @return CreateVpcFirewallCenConfigureResponse
 */
func (client *Client) CreateVpcFirewallCenConfigure(request *CreateVpcFirewallCenConfigureRequest) (_result *CreateVpcFirewallCenConfigureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVpcFirewallCenConfigureResponse{}
	_body, _err := client.CreateVpcFirewallCenConfigureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to create a VPC firewall. The VPC firewall controls traffic between two VPCs that are connected by using an Express Connect circuit. The VPC firewall does not control the mutual access traffic between VPCs that reside in different regions or belong to different Alibaba Cloud accounts. The firewall also does not control the mutual access traffic between VPCs and virtual border routers (VBRs). For more information, see [VPC firewall limits](~~172295~~).
 * ### [](#qps)QPS limit
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request CreateVpcFirewallConfigureRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateVpcFirewallConfigureResponse
 */
func (client *Client) CreateVpcFirewallConfigureWithOptions(request *CreateVpcFirewallConfigureRequest, runtime *util.RuntimeOptions) (_result *CreateVpcFirewallConfigureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FirewallSwitch)) {
		query["FirewallSwitch"] = request.FirewallSwitch
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.LocalVpcCidrTableList)) {
		query["LocalVpcCidrTableList"] = request.LocalVpcCidrTableList
	}

	if !tea.BoolValue(util.IsUnset(request.LocalVpcId)) {
		query["LocalVpcId"] = request.LocalVpcId
	}

	if !tea.BoolValue(util.IsUnset(request.LocalVpcRegion)) {
		query["LocalVpcRegion"] = request.LocalVpcRegion
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.PeerVpcCidrTableList)) {
		query["PeerVpcCidrTableList"] = request.PeerVpcCidrTableList
	}

	if !tea.BoolValue(util.IsUnset(request.PeerVpcId)) {
		query["PeerVpcId"] = request.PeerVpcId
	}

	if !tea.BoolValue(util.IsUnset(request.PeerVpcRegion)) {
		query["PeerVpcRegion"] = request.PeerVpcRegion
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallName)) {
		query["VpcFirewallName"] = request.VpcFirewallName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVpcFirewallConfigure"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVpcFirewallConfigureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to create a VPC firewall. The VPC firewall controls traffic between two VPCs that are connected by using an Express Connect circuit. The VPC firewall does not control the mutual access traffic between VPCs that reside in different regions or belong to different Alibaba Cloud accounts. The firewall also does not control the mutual access traffic between VPCs and virtual border routers (VBRs). For more information, see [VPC firewall limits](~~172295~~).
 * ### [](#qps)QPS limit
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request CreateVpcFirewallConfigureRequest
 * @return CreateVpcFirewallConfigureResponse
 */
func (client *Client) CreateVpcFirewallConfigure(request *CreateVpcFirewallConfigureRequest) (_result *CreateVpcFirewallConfigureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVpcFirewallConfigureResponse{}
	_body, _err := client.CreateVpcFirewallConfigureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the CreateVpcFirewallControlPolicy operation to create an access control policy in a specified policy group for a VPC firewall. Different access control policies are used when a VPC firewall is used to protect traffic between two VPCs that are connected by using a Cloud Enterprise Network (CEN) instance or an Express Connect circuit.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request CreateVpcFirewallControlPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateVpcFirewallControlPolicyResponse
 */
func (client *Client) CreateVpcFirewallControlPolicyWithOptions(request *CreateVpcFirewallControlPolicyRequest, runtime *util.RuntimeOptions) (_result *CreateVpcFirewallControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclAction)) {
		query["AclAction"] = request.AclAction
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicationName)) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicationNameList)) {
		query["ApplicationNameList"] = request.ApplicationNameList
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DestPort)) {
		query["DestPort"] = request.DestPort
	}

	if !tea.BoolValue(util.IsUnset(request.DestPortGroup)) {
		query["DestPortGroup"] = request.DestPortGroup
	}

	if !tea.BoolValue(util.IsUnset(request.DestPortType)) {
		query["DestPortType"] = request.DestPortType
	}

	if !tea.BoolValue(util.IsUnset(request.Destination)) {
		query["Destination"] = request.Destination
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationType)) {
		query["DestinationType"] = request.DestinationType
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.NewOrder)) {
		query["NewOrder"] = request.NewOrder
	}

	if !tea.BoolValue(util.IsUnset(request.Proto)) {
		query["Proto"] = request.Proto
	}

	if !tea.BoolValue(util.IsUnset(request.Release)) {
		query["Release"] = request.Release
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatDays)) {
		query["RepeatDays"] = request.RepeatDays
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatEndTime)) {
		query["RepeatEndTime"] = request.RepeatEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatStartTime)) {
		query["RepeatStartTime"] = request.RepeatStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatType)) {
		query["RepeatType"] = request.RepeatType
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallId)) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVpcFirewallControlPolicy"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVpcFirewallControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the CreateVpcFirewallControlPolicy operation to create an access control policy in a specified policy group for a VPC firewall. Different access control policies are used when a VPC firewall is used to protect traffic between two VPCs that are connected by using a Cloud Enterprise Network (CEN) instance or an Express Connect circuit.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request CreateVpcFirewallControlPolicyRequest
 * @return CreateVpcFirewallControlPolicyResponse
 */
func (client *Client) CreateVpcFirewallControlPolicy(request *CreateVpcFirewallControlPolicyRequest) (_result *CreateVpcFirewallControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVpcFirewallControlPolicyResponse{}
	_body, _err := client.CreateVpcFirewallControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DeleteAddressBook operation to delete an address book for access control.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DeleteAddressBookRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteAddressBookResponse
 */
func (client *Client) DeleteAddressBookWithOptions(request *DeleteAddressBookRequest, runtime *util.RuntimeOptions) (_result *DeleteAddressBookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupUuid)) {
		query["GroupUuid"] = request.GroupUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAddressBook"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAddressBookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DeleteAddressBook operation to delete an address book for access control.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DeleteAddressBookRequest
 * @return DeleteAddressBookResponse
 */
func (client *Client) DeleteAddressBook(request *DeleteAddressBookRequest) (_result *DeleteAddressBookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAddressBookResponse{}
	_body, _err := client.DeleteAddressBookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DeleteControlPolicy operation to delete an access control policy that applies to inbound or outbound traffic.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DeleteControlPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteControlPolicyResponse
 */
func (client *Client) DeleteControlPolicyWithOptions(request *DeleteControlPolicyRequest, runtime *util.RuntimeOptions) (_result *DeleteControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclUuid)) {
		query["AclUuid"] = request.AclUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		query["Direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteControlPolicy"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DeleteControlPolicy operation to delete an access control policy that applies to inbound or outbound traffic.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DeleteControlPolicyRequest
 * @return DeleteControlPolicyResponse
 */
func (client *Client) DeleteControlPolicy(request *DeleteControlPolicyRequest) (_result *DeleteControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteControlPolicyResponse{}
	_body, _err := client.DeleteControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteControlPolicyTemplateWithOptions(request *DeleteControlPolicyTemplateRequest, runtime *util.RuntimeOptions) (_result *DeleteControlPolicyTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteControlPolicyTemplate"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteControlPolicyTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteControlPolicyTemplate(request *DeleteControlPolicyTemplateRequest) (_result *DeleteControlPolicyTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteControlPolicyTemplateResponse{}
	_body, _err := client.DeleteControlPolicyTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to delete file download tasks and delete the files.
 * **
 * **Warning** Both tasks and involved files are deleted. You can no longer download the involved files by using the download links. This operation is irreversible. Proceed with caution.
 *
 * @param request DeleteDownloadTaskRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteDownloadTaskResponse
 */
func (client *Client) DeleteDownloadTaskWithOptions(request *DeleteDownloadTaskRequest, runtime *util.RuntimeOptions) (_result *DeleteDownloadTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDownloadTask"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDownloadTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to delete file download tasks and delete the files.
 * **
 * **Warning** Both tasks and involved files are deleted. You can no longer download the involved files by using the download links. This operation is irreversible. Proceed with caution.
 *
 * @param request DeleteDownloadTaskRequest
 * @return DeleteDownloadTaskResponse
 */
func (client *Client) DeleteDownloadTask(request *DeleteDownloadTaskRequest) (_result *DeleteDownloadTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDownloadTaskResponse{}
	_body, _err := client.DeleteDownloadTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFirewallV2RoutePoliciesWithOptions(request *DeleteFirewallV2RoutePoliciesRequest, runtime *util.RuntimeOptions) (_result *DeleteFirewallV2RoutePoliciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FirewallId)) {
		query["FirewallId"] = request.FirewallId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.TrFirewallRoutePolicyId)) {
		query["TrFirewallRoutePolicyId"] = request.TrFirewallRoutePolicyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFirewallV2RoutePolicies"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFirewallV2RoutePoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFirewallV2RoutePolicies(request *DeleteFirewallV2RoutePoliciesRequest) (_result *DeleteFirewallV2RoutePoliciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFirewallV2RoutePoliciesResponse{}
	_body, _err := client.DeleteFirewallV2RoutePoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DeleteInstanceMembers operation to remove members from Cloud Firewall.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DeleteInstanceMembersRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteInstanceMembersResponse
 */
func (client *Client) DeleteInstanceMembersWithOptions(request *DeleteInstanceMembersRequest, runtime *util.RuntimeOptions) (_result *DeleteInstanceMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MemberUids)) {
		query["MemberUids"] = request.MemberUids
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstanceMembers"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteInstanceMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DeleteInstanceMembers operation to remove members from Cloud Firewall.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DeleteInstanceMembersRequest
 * @return DeleteInstanceMembersResponse
 */
func (client *Client) DeleteInstanceMembers(request *DeleteInstanceMembersRequest) (_result *DeleteInstanceMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteInstanceMembersResponse{}
	_body, _err := client.DeleteInstanceMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can use this operation to delete an outbound access control policy that is created for a NAT firewall.
 *
 * @param request DeleteNatFirewallControlPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteNatFirewallControlPolicyResponse
 */
func (client *Client) DeleteNatFirewallControlPolicyWithOptions(request *DeleteNatFirewallControlPolicyRequest, runtime *util.RuntimeOptions) (_result *DeleteNatFirewallControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclUuid)) {
		query["AclUuid"] = request.AclUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		query["Direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.NatGatewayId)) {
		query["NatGatewayId"] = request.NatGatewayId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteNatFirewallControlPolicy"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteNatFirewallControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can use this operation to delete an outbound access control policy that is created for a NAT firewall.
 *
 * @param request DeleteNatFirewallControlPolicyRequest
 * @return DeleteNatFirewallControlPolicyResponse
 */
func (client *Client) DeleteNatFirewallControlPolicy(request *DeleteNatFirewallControlPolicyRequest) (_result *DeleteNatFirewallControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteNatFirewallControlPolicyResponse{}
	_body, _err := client.DeleteNatFirewallControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteNatFirewallControlPolicyBatchWithOptions(request *DeleteNatFirewallControlPolicyBatchRequest, runtime *util.RuntimeOptions) (_result *DeleteNatFirewallControlPolicyBatchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclUuidList)) {
		query["AclUuidList"] = request.AclUuidList
	}

	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		query["Direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.NatGatewayId)) {
		query["NatGatewayId"] = request.NatGatewayId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteNatFirewallControlPolicyBatch"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteNatFirewallControlPolicyBatchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteNatFirewallControlPolicyBatch(request *DeleteNatFirewallControlPolicyBatchRequest) (_result *DeleteNatFirewallControlPolicyBatchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteNatFirewallControlPolicyBatchResponse{}
	_body, _err := client.DeleteNatFirewallControlPolicyBatchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTrFirewallV2WithOptions(request *DeleteTrFirewallV2Request, runtime *util.RuntimeOptions) (_result *DeleteTrFirewallV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FirewallId)) {
		query["FirewallId"] = request.FirewallId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTrFirewallV2"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTrFirewallV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTrFirewallV2(request *DeleteTrFirewallV2Request) (_result *DeleteTrFirewallV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTrFirewallV2Response{}
	_body, _err := client.DeleteTrFirewallV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DeleteVpcFirewallCenConfigure operation to delete a VPC firewall. The VPC firewall protects mutual access traffic between a VPC and a specified network instance that is attached to a CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance. Before you call this operation, make sure that you have created a VPC firewall by calling the [CreateVpcFirewallCenConfigure](~~345772~~) operation.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DeleteVpcFirewallCenConfigureRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteVpcFirewallCenConfigureResponse
 */
func (client *Client) DeleteVpcFirewallCenConfigureWithOptions(request *DeleteVpcFirewallCenConfigureRequest, runtime *util.RuntimeOptions) (_result *DeleteVpcFirewallCenConfigureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallIdList)) {
		query["VpcFirewallIdList"] = request.VpcFirewallIdList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVpcFirewallCenConfigure"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVpcFirewallCenConfigureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DeleteVpcFirewallCenConfigure operation to delete a VPC firewall. The VPC firewall protects mutual access traffic between a VPC and a specified network instance that is attached to a CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance. Before you call this operation, make sure that you have created a VPC firewall by calling the [CreateVpcFirewallCenConfigure](~~345772~~) operation.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DeleteVpcFirewallCenConfigureRequest
 * @return DeleteVpcFirewallCenConfigureResponse
 */
func (client *Client) DeleteVpcFirewallCenConfigure(request *DeleteVpcFirewallCenConfigureRequest) (_result *DeleteVpcFirewallCenConfigureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVpcFirewallCenConfigureResponse{}
	_body, _err := client.DeleteVpcFirewallCenConfigureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DeleteVpcFirewallConfigure operation to delete a VPC firewall. The VPC firewall controls traffic between two VPCs that are connected by using an Express Connect circuit. Before you call the operation, make sure that you created a VPC firewall by calling the [CreateVpcFirewallConfigure](~~342893~~) operation.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DeleteVpcFirewallConfigureRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteVpcFirewallConfigureResponse
 */
func (client *Client) DeleteVpcFirewallConfigureWithOptions(request *DeleteVpcFirewallConfigureRequest, runtime *util.RuntimeOptions) (_result *DeleteVpcFirewallConfigureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallIdList)) {
		query["VpcFirewallIdList"] = request.VpcFirewallIdList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVpcFirewallConfigure"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVpcFirewallConfigureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DeleteVpcFirewallConfigure operation to delete a VPC firewall. The VPC firewall controls traffic between two VPCs that are connected by using an Express Connect circuit. Before you call the operation, make sure that you created a VPC firewall by calling the [CreateVpcFirewallConfigure](~~342893~~) operation.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DeleteVpcFirewallConfigureRequest
 * @return DeleteVpcFirewallConfigureResponse
 */
func (client *Client) DeleteVpcFirewallConfigure(request *DeleteVpcFirewallConfigureRequest) (_result *DeleteVpcFirewallConfigureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVpcFirewallConfigureResponse{}
	_body, _err := client.DeleteVpcFirewallConfigureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DeleteVpcFirewallControlPolicy operation to delete an access control policy from a specific policy group for a VPC firewall. Different access control policies are used for the VPC firewall that is used to protect each Cloud Enterprise Network (CEN) instance and the VPC firewall that is used to protect each Express Connect circuit.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DeleteVpcFirewallControlPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteVpcFirewallControlPolicyResponse
 */
func (client *Client) DeleteVpcFirewallControlPolicyWithOptions(request *DeleteVpcFirewallControlPolicyRequest, runtime *util.RuntimeOptions) (_result *DeleteVpcFirewallControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclUuid)) {
		query["AclUuid"] = request.AclUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallId)) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVpcFirewallControlPolicy"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVpcFirewallControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DeleteVpcFirewallControlPolicy operation to delete an access control policy from a specific policy group for a VPC firewall. Different access control policies are used for the VPC firewall that is used to protect each Cloud Enterprise Network (CEN) instance and the VPC firewall that is used to protect each Express Connect circuit.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DeleteVpcFirewallControlPolicyRequest
 * @return DeleteVpcFirewallControlPolicyResponse
 */
func (client *Client) DeleteVpcFirewallControlPolicy(request *DeleteVpcFirewallControlPolicyRequest) (_result *DeleteVpcFirewallControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVpcFirewallControlPolicyResponse{}
	_body, _err := client.DeleteVpcFirewallControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeACLProtectTrendWithOptions(request *DescribeACLProtectTrendRequest, runtime *util.RuntimeOptions) (_result *DescribeACLProtectTrendResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeACLProtectTrend"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeACLProtectTrendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeACLProtectTrend(request *DescribeACLProtectTrendRequest) (_result *DescribeACLProtectTrendResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeACLProtectTrendResponse{}
	_body, _err := client.DescribeACLProtectTrendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to query the details about an address book for an access control policy.
 * ## [](#qps)Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeAddressBookRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeAddressBookResponse
 */
func (client *Client) DescribeAddressBookWithOptions(request *DescribeAddressBookRequest, runtime *util.RuntimeOptions) (_result *DescribeAddressBookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContainPort)) {
		query["ContainPort"] = request.ContainPort
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.GroupType)) {
		query["GroupType"] = request.GroupType
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["Query"] = request.Query
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAddressBook"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAddressBookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to query the details about an address book for an access control policy.
 * ## [](#qps)Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeAddressBookRequest
 * @return DescribeAddressBookResponse
 */
func (client *Client) DescribeAddressBook(request *DescribeAddressBookRequest) (_result *DescribeAddressBookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAddressBookResponse{}
	_body, _err := client.DescribeAddressBookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DescribeAssetList operation to query the assets that are protected by Cloud Firewall.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeAssetListRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeAssetListResponse
 */
func (client *Client) DescribeAssetListWithOptions(request *DescribeAssetListRequest, runtime *util.RuntimeOptions) (_result *DescribeAssetListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.IpVersion)) {
		query["IpVersion"] = request.IpVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.NewResourceTag)) {
		query["NewResourceTag"] = request.NewResourceTag
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionNo)) {
		query["RegionNo"] = request.RegionNo
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.SearchItem)) {
		query["SearchItem"] = request.SearchItem
	}

	if !tea.BoolValue(util.IsUnset(request.SgStatus)) {
		query["SgStatus"] = request.SgStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.UserType)) {
		query["UserType"] = request.UserType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAssetList"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAssetListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DescribeAssetList operation to query the assets that are protected by Cloud Firewall.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeAssetListRequest
 * @return DescribeAssetListResponse
 */
func (client *Client) DescribeAssetList(request *DescribeAssetListRequest) (_result *DescribeAssetListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAssetListResponse{}
	_body, _err := client.DescribeAssetListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAssetRiskListWithOptions(request *DescribeAssetRiskListRequest, runtime *util.RuntimeOptions) (_result *DescribeAssetRiskListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IpAddrList)) {
		query["IpAddrList"] = request.IpAddrList
	}

	if !tea.BoolValue(util.IsUnset(request.IpVersion)) {
		query["IpVersion"] = request.IpVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAssetRiskList"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAssetRiskListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAssetRiskList(request *DescribeAssetRiskListRequest) (_result *DescribeAssetRiskListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAssetRiskListResponse{}
	_body, _err := client.DescribeAssetRiskListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCfwRiskLevelSummaryWithOptions(request *DescribeCfwRiskLevelSummaryRequest, runtime *util.RuntimeOptions) (_result *DescribeCfwRiskLevelSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
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
		Action:      tea.String("DescribeCfwRiskLevelSummary"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCfwRiskLevelSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCfwRiskLevelSummary(request *DescribeCfwRiskLevelSummaryRequest) (_result *DescribeCfwRiskLevelSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCfwRiskLevelSummaryResponse{}
	_body, _err := client.DescribeCfwRiskLevelSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DescribeControlPolicy operation to query the details about access control policies by page.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeControlPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeControlPolicyResponse
 */
func (client *Client) DescribeControlPolicyWithOptions(request *DescribeControlPolicyRequest, runtime *util.RuntimeOptions) (_result *DescribeControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclAction)) {
		query["AclAction"] = request.AclAction
	}

	if !tea.BoolValue(util.IsUnset(request.AclUuid)) {
		query["AclUuid"] = request.AclUuid
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Destination)) {
		query["Destination"] = request.Destination
	}

	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		query["Direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.IpVersion)) {
		query["IpVersion"] = request.IpVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Proto)) {
		query["Proto"] = request.Proto
	}

	if !tea.BoolValue(util.IsUnset(request.Release)) {
		query["Release"] = request.Release
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatType)) {
		query["RepeatType"] = request.RepeatType
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeControlPolicy"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DescribeControlPolicy operation to query the details about access control policies by page.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeControlPolicyRequest
 * @return DescribeControlPolicyResponse
 */
func (client *Client) DescribeControlPolicy(request *DescribeControlPolicyRequest) (_result *DescribeControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeControlPolicyResponse{}
	_body, _err := client.DescribeControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDefaultIPSConfigWithOptions(request *DescribeDefaultIPSConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeDefaultIPSConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDefaultIPSConfig"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDefaultIPSConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDefaultIPSConfig(request *DescribeDefaultIPSConfigRequest) (_result *DescribeDefaultIPSConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDefaultIPSConfigResponse{}
	_body, _err := client.DescribeDefaultIPSConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can use this operation to query the DNS record of a domain name. This operation can retrieve DNS records only from Alibaba Cloud DNS. Before you can call this operation, make sure that your domain name is hosted on Alibaba Cloud DNS.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeDomainResolveRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDomainResolveResponse
 */
func (client *Client) DescribeDomainResolveWithOptions(request *DescribeDomainResolveRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainResolveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.IpVersion)) {
		query["IpVersion"] = request.IpVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomainResolve"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDomainResolveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can use this operation to query the DNS record of a domain name. This operation can retrieve DNS records only from Alibaba Cloud DNS. Before you can call this operation, make sure that your domain name is hosted on Alibaba Cloud DNS.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeDomainResolveRequest
 * @return DescribeDomainResolveResponse
 */
func (client *Client) DescribeDomainResolve(request *DescribeDomainResolveRequest) (_result *DescribeDomainResolveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainResolveResponse{}
	_body, _err := client.DescribeDomainResolveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDownloadTaskWithOptions(request *DescribeDownloadTaskRequest, runtime *util.RuntimeOptions) (_result *DescribeDownloadTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		query["TaskType"] = request.TaskType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDownloadTask"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDownloadTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDownloadTask(request *DescribeDownloadTaskRequest) (_result *DescribeDownloadTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDownloadTaskResponse{}
	_body, _err := client.DescribeDownloadTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDownloadTaskTypeWithOptions(request *DescribeDownloadTaskTypeRequest, runtime *util.RuntimeOptions) (_result *DescribeDownloadTaskTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		query["TaskType"] = request.TaskType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDownloadTaskType"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDownloadTaskTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDownloadTaskType(request *DescribeDownloadTaskTypeRequest) (_result *DescribeDownloadTaskTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDownloadTaskTypeResponse{}
	_body, _err := client.DescribeDownloadTaskTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can use this operation to query the information about members in Cloud Firewall.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeInstanceMembersRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeInstanceMembersResponse
 */
func (client *Client) DescribeInstanceMembersWithOptions(request *DescribeInstanceMembersRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceMembersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.MemberDesc)) {
		query["MemberDesc"] = request.MemberDesc
	}

	if !tea.BoolValue(util.IsUnset(request.MemberDisplayName)) {
		query["MemberDisplayName"] = request.MemberDisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceMembers"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can use this operation to query the information about members in Cloud Firewall.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeInstanceMembersRequest
 * @return DescribeInstanceMembersResponse
 */
func (client *Client) DescribeInstanceMembers(request *DescribeInstanceMembersRequest) (_result *DescribeInstanceMembersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceMembersResponse{}
	_body, _err := client.DescribeInstanceMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceRiskLevelsWithOptions(request *DescribeInstanceRiskLevelsRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceRiskLevelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Instances)) {
		query["Instances"] = request.Instances
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceRiskLevels"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceRiskLevelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceRiskLevels(request *DescribeInstanceRiskLevelsRequest) (_result *DescribeInstanceRiskLevelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceRiskLevelsResponse{}
	_body, _err := client.DescribeInstanceRiskLevelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInternetOpenIpWithOptions(request *DescribeInternetOpenIpRequest, runtime *util.RuntimeOptions) (_result *DescribeInternetOpenIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssetsInstanceId)) {
		query["AssetsInstanceId"] = request.AssetsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.AssetsInstanceName)) {
		query["AssetsInstanceName"] = request.AssetsInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.AssetsType)) {
		query["AssetsType"] = request.AssetsType
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		query["Port"] = request.Port
	}

	if !tea.BoolValue(util.IsUnset(request.PublicIp)) {
		query["PublicIp"] = request.PublicIp
	}

	if !tea.BoolValue(util.IsUnset(request.RegionNo)) {
		query["RegionNo"] = request.RegionNo
	}

	if !tea.BoolValue(util.IsUnset(request.RiskLevel)) {
		query["RiskLevel"] = request.RiskLevel
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInternetOpenIp"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInternetOpenIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInternetOpenIp(request *DescribeInternetOpenIpRequest) (_result *DescribeInternetOpenIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInternetOpenIpResponse{}
	_body, _err := client.DescribeInternetOpenIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInternetTrafficTrendWithOptions(request *DescribeInternetTrafficTrendRequest, runtime *util.RuntimeOptions) (_result *DescribeInternetTrafficTrendResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		query["Direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceCode)) {
		query["SourceCode"] = request.SourceCode
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.SrcPrivateIP)) {
		query["SrcPrivateIP"] = request.SrcPrivateIP
	}

	if !tea.BoolValue(util.IsUnset(request.SrcPublicIP)) {
		query["SrcPublicIP"] = request.SrcPublicIP
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TrafficType)) {
		query["TrafficType"] = request.TrafficType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInternetTrafficTrend"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInternetTrafficTrendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInternetTrafficTrend(request *DescribeInternetTrafficTrendRequest) (_result *DescribeInternetTrafficTrendResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInternetTrafficTrendResponse{}
	_body, _err := client.DescribeInternetTrafficTrendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInvadeEventListWithOptions(request *DescribeInvadeEventListRequest, runtime *util.RuntimeOptions) (_result *DescribeInvadeEventListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AssetsIP)) {
		query["AssetsIP"] = request.AssetsIP
	}

	if !tea.BoolValue(util.IsUnset(request.AssetsInstanceId)) {
		query["AssetsInstanceId"] = request.AssetsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.AssetsInstanceName)) {
		query["AssetsInstanceName"] = request.AssetsInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.EventKey)) {
		query["EventKey"] = request.EventKey
	}

	if !tea.BoolValue(util.IsUnset(request.EventName)) {
		query["EventName"] = request.EventName
	}

	if !tea.BoolValue(util.IsUnset(request.EventUuid)) {
		query["EventUuid"] = request.EventUuid
	}

	if !tea.BoolValue(util.IsUnset(request.IsIgnore)) {
		query["IsIgnore"] = request.IsIgnore
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProcessStatusList)) {
		query["ProcessStatusList"] = request.ProcessStatusList
	}

	if !tea.BoolValue(util.IsUnset(request.RiskLevel)) {
		query["RiskLevel"] = request.RiskLevel
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInvadeEventList"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInvadeEventListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInvadeEventList(request *DescribeInvadeEventListRequest) (_result *DescribeInvadeEventListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInvadeEventListResponse{}
	_body, _err := client.DescribeInvadeEventListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNatAclPageStatusWithOptions(request *DescribeNatAclPageStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeNatAclPageStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNatAclPageStatus"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeNatAclPageStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNatAclPageStatus(request *DescribeNatAclPageStatusRequest) (_result *DescribeNatAclPageStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNatAclPageStatusResponse{}
	_body, _err := client.DescribeNatAclPageStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can use this operation to query the information about all access control policies that are created for NAT firewalls by page.
 *
 * @param request DescribeNatFirewallControlPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeNatFirewallControlPolicyResponse
 */
func (client *Client) DescribeNatFirewallControlPolicyWithOptions(request *DescribeNatFirewallControlPolicyRequest, runtime *util.RuntimeOptions) (_result *DescribeNatFirewallControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclAction)) {
		query["AclAction"] = request.AclAction
	}

	if !tea.BoolValue(util.IsUnset(request.AclUuid)) {
		query["AclUuid"] = request.AclUuid
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Destination)) {
		query["Destination"] = request.Destination
	}

	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		query["Direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.NatGatewayId)) {
		query["NatGatewayId"] = request.NatGatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Proto)) {
		query["Proto"] = request.Proto
	}

	if !tea.BoolValue(util.IsUnset(request.Release)) {
		query["Release"] = request.Release
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatType)) {
		query["RepeatType"] = request.RepeatType
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNatFirewallControlPolicy"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeNatFirewallControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can use this operation to query the information about all access control policies that are created for NAT firewalls by page.
 *
 * @param request DescribeNatFirewallControlPolicyRequest
 * @return DescribeNatFirewallControlPolicyResponse
 */
func (client *Client) DescribeNatFirewallControlPolicy(request *DescribeNatFirewallControlPolicyRequest) (_result *DescribeNatFirewallControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNatFirewallControlPolicyResponse{}
	_body, _err := client.DescribeNatFirewallControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can use this operation to query the priority range of access control policies that are created for a NAT firewall.
 *
 * @param request DescribeNatFirewallPolicyPriorUsedRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeNatFirewallPolicyPriorUsedResponse
 */
func (client *Client) DescribeNatFirewallPolicyPriorUsedWithOptions(request *DescribeNatFirewallPolicyPriorUsedRequest, runtime *util.RuntimeOptions) (_result *DescribeNatFirewallPolicyPriorUsedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		query["Direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.IpVersion)) {
		query["IpVersion"] = request.IpVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.NatGatewayId)) {
		query["NatGatewayId"] = request.NatGatewayId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNatFirewallPolicyPriorUsed"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeNatFirewallPolicyPriorUsedResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can use this operation to query the priority range of access control policies that are created for a NAT firewall.
 *
 * @param request DescribeNatFirewallPolicyPriorUsedRequest
 * @return DescribeNatFirewallPolicyPriorUsedResponse
 */
func (client *Client) DescribeNatFirewallPolicyPriorUsed(request *DescribeNatFirewallPolicyPriorUsedRequest) (_result *DescribeNatFirewallPolicyPriorUsedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNatFirewallPolicyPriorUsedResponse{}
	_body, _err := client.DescribeNatFirewallPolicyPriorUsedWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOutgoingDestinationIPWithOptions(request *DescribeOutgoingDestinationIPRequest, runtime *util.RuntimeOptions) (_result *DescribeOutgoingDestinationIPResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationName)) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		query["CategoryId"] = request.CategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DstIP)) {
		query["DstIP"] = request.DstIP
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		query["Port"] = request.Port
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateIP)) {
		query["PrivateIP"] = request.PrivateIP
	}

	if !tea.BoolValue(util.IsUnset(request.PublicIP)) {
		query["PublicIP"] = request.PublicIP
	}

	if !tea.BoolValue(util.IsUnset(request.Sort)) {
		query["Sort"] = request.Sort
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TagIdNew)) {
		query["TagIdNew"] = request.TagIdNew
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOutgoingDestinationIP"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOutgoingDestinationIPResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOutgoingDestinationIP(request *DescribeOutgoingDestinationIPRequest) (_result *DescribeOutgoingDestinationIPResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOutgoingDestinationIPResponse{}
	_body, _err := client.DescribeOutgoingDestinationIPWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOutgoingDomainWithOptions(request *DescribeOutgoingDomainRequest, runtime *util.RuntimeOptions) (_result *DescribeOutgoingDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		query["CategoryId"] = request.CategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PublicIP)) {
		query["PublicIP"] = request.PublicIP
	}

	if !tea.BoolValue(util.IsUnset(request.Sort)) {
		query["Sort"] = request.Sort
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TagIdNew)) {
		query["TagIdNew"] = request.TagIdNew
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOutgoingDomain"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOutgoingDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOutgoingDomain(request *DescribeOutgoingDomainRequest) (_result *DescribeOutgoingDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOutgoingDomainResponse{}
	_body, _err := client.DescribeOutgoingDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DescribePolicyAdvancedConfig operation to query whether the strict mode is enabled for an access control policy.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribePolicyAdvancedConfigRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribePolicyAdvancedConfigResponse
 */
func (client *Client) DescribePolicyAdvancedConfigWithOptions(request *DescribePolicyAdvancedConfigRequest, runtime *util.RuntimeOptions) (_result *DescribePolicyAdvancedConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePolicyAdvancedConfig"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePolicyAdvancedConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DescribePolicyAdvancedConfig operation to query whether the strict mode is enabled for an access control policy.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribePolicyAdvancedConfigRequest
 * @return DescribePolicyAdvancedConfigResponse
 */
func (client *Client) DescribePolicyAdvancedConfig(request *DescribePolicyAdvancedConfigRequest) (_result *DescribePolicyAdvancedConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePolicyAdvancedConfigResponse{}
	_body, _err := client.DescribePolicyAdvancedConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DescribePolicyPriorUsed operation to query the priority range of the access control policies that match specific query conditions.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribePolicyPriorUsedRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribePolicyPriorUsedResponse
 */
func (client *Client) DescribePolicyPriorUsedWithOptions(request *DescribePolicyPriorUsedRequest, runtime *util.RuntimeOptions) (_result *DescribePolicyPriorUsedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		query["Direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.IpVersion)) {
		query["IpVersion"] = request.IpVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePolicyPriorUsed"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePolicyPriorUsedResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DescribePolicyPriorUsed operation to query the priority range of the access control policies that match specific query conditions.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribePolicyPriorUsedRequest
 * @return DescribePolicyPriorUsedResponse
 */
func (client *Client) DescribePolicyPriorUsed(request *DescribePolicyPriorUsedRequest) (_result *DescribePolicyPriorUsedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePolicyPriorUsedResponse{}
	_body, _err := client.DescribePolicyPriorUsedWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePostpayTrafficDetailWithOptions(request *DescribePostpayTrafficDetailRequest, runtime *util.RuntimeOptions) (_result *DescribePostpayTrafficDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchItem)) {
		query["SearchItem"] = request.SearchItem
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TrafficType)) {
		query["TrafficType"] = request.TrafficType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePostpayTrafficDetail"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePostpayTrafficDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePostpayTrafficDetail(request *DescribePostpayTrafficDetailRequest) (_result *DescribePostpayTrafficDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePostpayTrafficDetailResponse{}
	_body, _err := client.DescribePostpayTrafficDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePostpayTrafficTotalWithOptions(request *DescribePostpayTrafficTotalRequest, runtime *util.RuntimeOptions) (_result *DescribePostpayTrafficTotalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePostpayTrafficTotal"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePostpayTrafficTotalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePostpayTrafficTotal(request *DescribePostpayTrafficTotalRequest) (_result *DescribePostpayTrafficTotalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePostpayTrafficTotalResponse{}
	_body, _err := client.DescribePostpayTrafficTotalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePrefixListsWithOptions(request *DescribePrefixListsRequest, runtime *util.RuntimeOptions) (_result *DescribePrefixListsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePrefixLists"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePrefixListsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePrefixLists(request *DescribePrefixListsRequest) (_result *DescribePrefixListsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePrefixListsResponse{}
	_body, _err := client.DescribePrefixListsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DescribeRiskEventGroup operation to query and download the details of intrusion events. We recommend that you query the details of 5 to 10 intrusion events at a time. If you do not need to query the geographical information about IP addresses, you can set the NoLocation parameter to true to prevent query timeout.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeRiskEventGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeRiskEventGroupResponse
 */
func (client *Client) DescribeRiskEventGroupWithOptions(request *DescribeRiskEventGroupRequest, runtime *util.RuntimeOptions) (_result *DescribeRiskEventGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AttackApp)) {
		query["AttackApp"] = request.AttackApp
	}

	if !tea.BoolValue(util.IsUnset(request.AttackType)) {
		query["AttackType"] = request.AttackType
	}

	if !tea.BoolValue(util.IsUnset(request.BuyVersion)) {
		query["BuyVersion"] = request.BuyVersion
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DataType)) {
		query["DataType"] = request.DataType
	}

	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		query["Direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.DstIP)) {
		query["DstIP"] = request.DstIP
	}

	if !tea.BoolValue(util.IsUnset(request.DstNetworkInstanceId)) {
		query["DstNetworkInstanceId"] = request.DstNetworkInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.EventName)) {
		query["EventName"] = request.EventName
	}

	if !tea.BoolValue(util.IsUnset(request.FirewallType)) {
		query["FirewallType"] = request.FirewallType
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.NoLocation)) {
		query["NoLocation"] = request.NoLocation
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RuleResult)) {
		query["RuleResult"] = request.RuleResult
	}

	if !tea.BoolValue(util.IsUnset(request.RuleSource)) {
		query["RuleSource"] = request.RuleSource
	}

	if !tea.BoolValue(util.IsUnset(request.Sort)) {
		query["Sort"] = request.Sort
	}

	if !tea.BoolValue(util.IsUnset(request.SrcIP)) {
		query["SrcIP"] = request.SrcIP
	}

	if !tea.BoolValue(util.IsUnset(request.SrcNetworkInstanceId)) {
		query["SrcNetworkInstanceId"] = request.SrcNetworkInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.VulLevel)) {
		query["VulLevel"] = request.VulLevel
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRiskEventGroup"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRiskEventGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DescribeRiskEventGroup operation to query and download the details of intrusion events. We recommend that you query the details of 5 to 10 intrusion events at a time. If you do not need to query the geographical information about IP addresses, you can set the NoLocation parameter to true to prevent query timeout.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeRiskEventGroupRequest
 * @return DescribeRiskEventGroupResponse
 */
func (client *Client) DescribeRiskEventGroup(request *DescribeRiskEventGroupRequest) (_result *DescribeRiskEventGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRiskEventGroupResponse{}
	_body, _err := client.DescribeRiskEventGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRiskEventPayloadWithOptions(request *DescribeRiskEventPayloadRequest, runtime *util.RuntimeOptions) (_result *DescribeRiskEventPayloadResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DstIP)) {
		query["DstIP"] = request.DstIP
	}

	if !tea.BoolValue(util.IsUnset(request.DstVpcId)) {
		query["DstVpcId"] = request.DstVpcId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.FirewallType)) {
		query["FirewallType"] = request.FirewallType
	}

	if !tea.BoolValue(util.IsUnset(request.PublicIP)) {
		query["PublicIP"] = request.PublicIP
	}

	if !tea.BoolValue(util.IsUnset(request.SrcIP)) {
		query["SrcIP"] = request.SrcIP
	}

	if !tea.BoolValue(util.IsUnset(request.SrcVpcId)) {
		query["SrcVpcId"] = request.SrcVpcId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.UUID)) {
		query["UUID"] = request.UUID
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRiskEventPayload"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRiskEventPayloadResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRiskEventPayload(request *DescribeRiskEventPayloadRequest) (_result *DescribeRiskEventPayloadResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRiskEventPayloadResponse{}
	_body, _err := client.DescribeRiskEventPayloadWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSignatureLibVersionWithOptions(runtime *util.RuntimeOptions) (_result *DescribeSignatureLibVersionResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeSignatureLibVersion"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSignatureLibVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSignatureLibVersion() (_result *DescribeSignatureLibVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSignatureLibVersionResponse{}
	_body, _err := client.DescribeSignatureLibVersionWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTrFirewallPolicyBackUpAssociationListWithOptions(tmpReq *DescribeTrFirewallPolicyBackUpAssociationListRequest, runtime *util.RuntimeOptions) (_result *DescribeTrFirewallPolicyBackUpAssociationListResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DescribeTrFirewallPolicyBackUpAssociationListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CandidateList)) {
		request.CandidateListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CandidateList, tea.String("CandidateList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CandidateListShrink)) {
		query["CandidateList"] = request.CandidateListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.FirewallId)) {
		query["FirewallId"] = request.FirewallId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.TrFirewallRoutePolicyId)) {
		query["TrFirewallRoutePolicyId"] = request.TrFirewallRoutePolicyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTrFirewallPolicyBackUpAssociationList"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTrFirewallPolicyBackUpAssociationListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTrFirewallPolicyBackUpAssociationList(request *DescribeTrFirewallPolicyBackUpAssociationListRequest) (_result *DescribeTrFirewallPolicyBackUpAssociationListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTrFirewallPolicyBackUpAssociationListResponse{}
	_body, _err := client.DescribeTrFirewallPolicyBackUpAssociationListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTrFirewallV2RoutePolicyListWithOptions(request *DescribeTrFirewallV2RoutePolicyListRequest, runtime *util.RuntimeOptions) (_result *DescribeTrFirewallV2RoutePolicyListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.FirewallId)) {
		query["FirewallId"] = request.FirewallId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		query["PolicyId"] = request.PolicyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTrFirewallV2RoutePolicyList"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTrFirewallV2RoutePolicyListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTrFirewallV2RoutePolicyList(request *DescribeTrFirewallV2RoutePolicyListRequest) (_result *DescribeTrFirewallV2RoutePolicyListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTrFirewallV2RoutePolicyListResponse{}
	_body, _err := client.DescribeTrFirewallV2RoutePolicyListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTrFirewallsV2DetailWithOptions(request *DescribeTrFirewallsV2DetailRequest, runtime *util.RuntimeOptions) (_result *DescribeTrFirewallsV2DetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FirewallId)) {
		query["FirewallId"] = request.FirewallId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTrFirewallsV2Detail"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTrFirewallsV2DetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTrFirewallsV2Detail(request *DescribeTrFirewallsV2DetailRequest) (_result *DescribeTrFirewallsV2DetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTrFirewallsV2DetailResponse{}
	_body, _err := client.DescribeTrFirewallsV2DetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTrFirewallsV2ListWithOptions(request *DescribeTrFirewallsV2ListRequest, runtime *util.RuntimeOptions) (_result *DescribeTrFirewallsV2ListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CenId)) {
		query["CenId"] = request.CenId
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.FirewallId)) {
		query["FirewallId"] = request.FirewallId
	}

	if !tea.BoolValue(util.IsUnset(request.FirewallName)) {
		query["FirewallName"] = request.FirewallName
	}

	if !tea.BoolValue(util.IsUnset(request.FirewallSwitchStatus)) {
		query["FirewallSwitchStatus"] = request.FirewallSwitchStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionNo)) {
		query["RegionNo"] = request.RegionNo
	}

	if !tea.BoolValue(util.IsUnset(request.RouteMode)) {
		query["RouteMode"] = request.RouteMode
	}

	if !tea.BoolValue(util.IsUnset(request.TransitRouterId)) {
		query["TransitRouterId"] = request.TransitRouterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTrFirewallsV2List"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTrFirewallsV2ListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTrFirewallsV2List(request *DescribeTrFirewallsV2ListRequest) (_result *DescribeTrFirewallsV2ListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTrFirewallsV2ListResponse{}
	_body, _err := client.DescribeTrFirewallsV2ListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTrFirewallsV2RouteListWithOptions(request *DescribeTrFirewallsV2RouteListRequest, runtime *util.RuntimeOptions) (_result *DescribeTrFirewallsV2RouteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.FirewallId)) {
		query["FirewallId"] = request.FirewallId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TrFirewallRoutePolicyId)) {
		query["TrFirewallRoutePolicyId"] = request.TrFirewallRoutePolicyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTrFirewallsV2RouteList"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTrFirewallsV2RouteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTrFirewallsV2RouteList(request *DescribeTrFirewallsV2RouteListRequest) (_result *DescribeTrFirewallsV2RouteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTrFirewallsV2RouteListResponse{}
	_body, _err := client.DescribeTrFirewallsV2RouteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserAssetIPTrafficInfoWithOptions(request *DescribeUserAssetIPTrafficInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeUserAssetIPTrafficInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUserAssetIPTrafficInfo"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUserAssetIPTrafficInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserAssetIPTrafficInfo(request *DescribeUserAssetIPTrafficInfoRequest) (_result *DescribeUserAssetIPTrafficInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserAssetIPTrafficInfoResponse{}
	_body, _err := client.DescribeUserAssetIPTrafficInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserIPSWhitelistWithOptions(request *DescribeUserIPSWhitelistRequest, runtime *util.RuntimeOptions) (_result *DescribeUserIPSWhitelistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUserIPSWhitelist"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUserIPSWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserIPSWhitelist(request *DescribeUserIPSWhitelistRequest) (_result *DescribeUserIPSWhitelistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserIPSWhitelistResponse{}
	_body, _err := client.DescribeUserIPSWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DescribeVpcFirewallAclGroupList operation to query the information about all policy groups of access control policies that are created for VPC firewalls.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeVpcFirewallAclGroupListRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeVpcFirewallAclGroupListResponse
 */
func (client *Client) DescribeVpcFirewallAclGroupListWithOptions(request *DescribeVpcFirewallAclGroupListRequest, runtime *util.RuntimeOptions) (_result *DescribeVpcFirewallAclGroupListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.FirewallConfigureStatus)) {
		query["FirewallConfigureStatus"] = request.FirewallConfigureStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVpcFirewallAclGroupList"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVpcFirewallAclGroupListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DescribeVpcFirewallAclGroupList operation to query the information about all policy groups of access control policies that are created for VPC firewalls.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeVpcFirewallAclGroupListRequest
 * @return DescribeVpcFirewallAclGroupListResponse
 */
func (client *Client) DescribeVpcFirewallAclGroupList(request *DescribeVpcFirewallAclGroupListRequest) (_result *DescribeVpcFirewallAclGroupListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVpcFirewallAclGroupListResponse{}
	_body, _err := client.DescribeVpcFirewallAclGroupListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DescribeVpcFirewallCenDetail operation to query the details about a VPC firewall. The VPC firewall protects access traffic between a specified VPC and a network instance that is attached to a CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeVpcFirewallCenDetailRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeVpcFirewallCenDetailResponse
 */
func (client *Client) DescribeVpcFirewallCenDetailWithOptions(request *DescribeVpcFirewallCenDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeVpcFirewallCenDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkInstanceId)) {
		query["NetworkInstanceId"] = request.NetworkInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallId)) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVpcFirewallCenDetail"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVpcFirewallCenDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DescribeVpcFirewallCenDetail operation to query the details about a VPC firewall. The VPC firewall protects access traffic between a specified VPC and a network instance that is attached to a CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeVpcFirewallCenDetailRequest
 * @return DescribeVpcFirewallCenDetailResponse
 */
func (client *Client) DescribeVpcFirewallCenDetail(request *DescribeVpcFirewallCenDetailRequest) (_result *DescribeVpcFirewallCenDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVpcFirewallCenDetailResponse{}
	_body, _err := client.DescribeVpcFirewallCenDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DescribeVpcFirewallCenList operation to query VPC firewalls. A VPC firewall protects mutual access traffic between a specified VPC and a network instance that is attached to a CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeVpcFirewallCenListRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeVpcFirewallCenListResponse
 */
func (client *Client) DescribeVpcFirewallCenListWithOptions(request *DescribeVpcFirewallCenListRequest, runtime *util.RuntimeOptions) (_result *DescribeVpcFirewallCenListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CenId)) {
		query["CenId"] = request.CenId
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.FirewallSwitchStatus)) {
		query["FirewallSwitchStatus"] = request.FirewallSwitchStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkInstanceId)) {
		query["NetworkInstanceId"] = request.NetworkInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionNo)) {
		query["RegionNo"] = request.RegionNo
	}

	if !tea.BoolValue(util.IsUnset(request.RouteMode)) {
		query["RouteMode"] = request.RouteMode
	}

	if !tea.BoolValue(util.IsUnset(request.TransitRouterType)) {
		query["TransitRouterType"] = request.TransitRouterType
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallId)) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallName)) {
		query["VpcFirewallName"] = request.VpcFirewallName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVpcFirewallCenList"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVpcFirewallCenListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DescribeVpcFirewallCenList operation to query VPC firewalls. A VPC firewall protects mutual access traffic between a specified VPC and a network instance that is attached to a CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeVpcFirewallCenListRequest
 * @return DescribeVpcFirewallCenListResponse
 */
func (client *Client) DescribeVpcFirewallCenList(request *DescribeVpcFirewallCenListRequest) (_result *DescribeVpcFirewallCenListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVpcFirewallCenListResponse{}
	_body, _err := client.DescribeVpcFirewallCenListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DescribeVpcFirewallControlPolicy operation to query the information about all access control policies that are created for a specified VPC firewall. Different access control policies are used when a VPC firewall is used to protect traffic between two VPCs that are connected by using a Cloud Enterprise Network (CEN) instance or an Express Connect circuit.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeVpcFirewallControlPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeVpcFirewallControlPolicyResponse
 */
func (client *Client) DescribeVpcFirewallControlPolicyWithOptions(request *DescribeVpcFirewallControlPolicyRequest, runtime *util.RuntimeOptions) (_result *DescribeVpcFirewallControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclAction)) {
		query["AclAction"] = request.AclAction
	}

	if !tea.BoolValue(util.IsUnset(request.AclUuid)) {
		query["AclUuid"] = request.AclUuid
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Destination)) {
		query["Destination"] = request.Destination
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Proto)) {
		query["Proto"] = request.Proto
	}

	if !tea.BoolValue(util.IsUnset(request.Release)) {
		query["Release"] = request.Release
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatType)) {
		query["RepeatType"] = request.RepeatType
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallId)) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVpcFirewallControlPolicy"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVpcFirewallControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DescribeVpcFirewallControlPolicy operation to query the information about all access control policies that are created for a specified VPC firewall. Different access control policies are used when a VPC firewall is used to protect traffic between two VPCs that are connected by using a Cloud Enterprise Network (CEN) instance or an Express Connect circuit.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeVpcFirewallControlPolicyRequest
 * @return DescribeVpcFirewallControlPolicyResponse
 */
func (client *Client) DescribeVpcFirewallControlPolicy(request *DescribeVpcFirewallControlPolicyRequest) (_result *DescribeVpcFirewallControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVpcFirewallControlPolicyResponse{}
	_body, _err := client.DescribeVpcFirewallControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DescribeVpcFirewallDefaultIPSConfig operation to query the intrusion prevention configurations of a VPC firewall.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeVpcFirewallDefaultIPSConfigRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeVpcFirewallDefaultIPSConfigResponse
 */
func (client *Client) DescribeVpcFirewallDefaultIPSConfigWithOptions(request *DescribeVpcFirewallDefaultIPSConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeVpcFirewallDefaultIPSConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallId)) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVpcFirewallDefaultIPSConfig"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVpcFirewallDefaultIPSConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DescribeVpcFirewallDefaultIPSConfig operation to query the intrusion prevention configurations of a VPC firewall.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeVpcFirewallDefaultIPSConfigRequest
 * @return DescribeVpcFirewallDefaultIPSConfigResponse
 */
func (client *Client) DescribeVpcFirewallDefaultIPSConfig(request *DescribeVpcFirewallDefaultIPSConfigRequest) (_result *DescribeVpcFirewallDefaultIPSConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVpcFirewallDefaultIPSConfigResponse{}
	_body, _err := client.DescribeVpcFirewallDefaultIPSConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DescribeVpcFirewallDetail operation to query the details about a VPC firewall. The VPC firewall controls traffic between two VPCs that are connected by using an Express Connect circuit.
 * Before you call the operation, make sure that you created a VPC firewall by calling the [CreateVpcFirewallConfigure](https://www.alibabacloud.com/help/en/cloud-firewall/latest/createvpcfirewallconfigure) operation.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeVpcFirewallDetailRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeVpcFirewallDetailResponse
 */
func (client *Client) DescribeVpcFirewallDetailWithOptions(request *DescribeVpcFirewallDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeVpcFirewallDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.LocalVpcId)) {
		query["LocalVpcId"] = request.LocalVpcId
	}

	if !tea.BoolValue(util.IsUnset(request.PeerVpcId)) {
		query["PeerVpcId"] = request.PeerVpcId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallId)) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVpcFirewallDetail"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVpcFirewallDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DescribeVpcFirewallDetail operation to query the details about a VPC firewall. The VPC firewall controls traffic between two VPCs that are connected by using an Express Connect circuit.
 * Before you call the operation, make sure that you created a VPC firewall by calling the [CreateVpcFirewallConfigure](https://www.alibabacloud.com/help/en/cloud-firewall/latest/createvpcfirewallconfigure) operation.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeVpcFirewallDetailRequest
 * @return DescribeVpcFirewallDetailResponse
 */
func (client *Client) DescribeVpcFirewallDetail(request *DescribeVpcFirewallDetailRequest) (_result *DescribeVpcFirewallDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVpcFirewallDetailResponse{}
	_body, _err := client.DescribeVpcFirewallDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVpcFirewallIPSWhitelistWithOptions(request *DescribeVpcFirewallIPSWhitelistRequest, runtime *util.RuntimeOptions) (_result *DescribeVpcFirewallIPSWhitelistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallId)) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVpcFirewallIPSWhitelist"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVpcFirewallIPSWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVpcFirewallIPSWhitelist(request *DescribeVpcFirewallIPSWhitelistRequest) (_result *DescribeVpcFirewallIPSWhitelistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVpcFirewallIPSWhitelistResponse{}
	_body, _err := client.DescribeVpcFirewallIPSWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DescribeVpcFirewallList operation to query the details about VPC firewalls by page. Each VPC firewall protects traffic between two VPCs that are connected by using an Express Connect circuit.
 * ### Limits
 * You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeVpcFirewallListRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeVpcFirewallListResponse
 */
func (client *Client) DescribeVpcFirewallListWithOptions(request *DescribeVpcFirewallListRequest, runtime *util.RuntimeOptions) (_result *DescribeVpcFirewallListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConnectSubType)) {
		query["ConnectSubType"] = request.ConnectSubType
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.FirewallSwitchStatus)) {
		query["FirewallSwitchStatus"] = request.FirewallSwitchStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PeerUid)) {
		query["PeerUid"] = request.PeerUid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionNo)) {
		query["RegionNo"] = request.RegionNo
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallId)) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallName)) {
		query["VpcFirewallName"] = request.VpcFirewallName
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVpcFirewallList"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVpcFirewallListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DescribeVpcFirewallList operation to query the details about VPC firewalls by page. Each VPC firewall protects traffic between two VPCs that are connected by using an Express Connect circuit.
 * ### Limits
 * You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeVpcFirewallListRequest
 * @return DescribeVpcFirewallListResponse
 */
func (client *Client) DescribeVpcFirewallList(request *DescribeVpcFirewallListRequest) (_result *DescribeVpcFirewallListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVpcFirewallListResponse{}
	_body, _err := client.DescribeVpcFirewallListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the DescribeVpcFirewallPolicyPriorUsed operation to query the priority range of access control policies that are created for a VPC firewall in a specific policy group.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeVpcFirewallPolicyPriorUsedRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeVpcFirewallPolicyPriorUsedResponse
 */
func (client *Client) DescribeVpcFirewallPolicyPriorUsedWithOptions(request *DescribeVpcFirewallPolicyPriorUsedRequest, runtime *util.RuntimeOptions) (_result *DescribeVpcFirewallPolicyPriorUsedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallId)) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVpcFirewallPolicyPriorUsed"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVpcFirewallPolicyPriorUsedResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the DescribeVpcFirewallPolicyPriorUsed operation to query the priority range of access control policies that are created for a VPC firewall in a specific policy group.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request DescribeVpcFirewallPolicyPriorUsedRequest
 * @return DescribeVpcFirewallPolicyPriorUsedResponse
 */
func (client *Client) DescribeVpcFirewallPolicyPriorUsed(request *DescribeVpcFirewallPolicyPriorUsedRequest) (_result *DescribeVpcFirewallPolicyPriorUsedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVpcFirewallPolicyPriorUsedResponse{}
	_body, _err := client.DescribeVpcFirewallPolicyPriorUsedWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVpcListLiteWithOptions(request *DescribeVpcListLiteRequest, runtime *util.RuntimeOptions) (_result *DescribeVpcListLiteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.RegionNo)) {
		query["RegionNo"] = request.RegionNo
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcName)) {
		query["VpcName"] = request.VpcName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVpcListLite"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVpcListLiteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVpcListLite(request *DescribeVpcListLiteRequest) (_result *DescribeVpcListLiteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVpcListLiteResponse{}
	_body, _err := client.DescribeVpcListLiteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVpcZoneWithOptions(request *DescribeVpcZoneRequest, runtime *util.RuntimeOptions) (_result *DescribeVpcZoneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Environment)) {
		query["Environment"] = request.Environment
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.RegionNo)) {
		query["RegionNo"] = request.RegionNo
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVpcZone"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVpcZoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVpcZone(request *DescribeVpcZoneRequest) (_result *DescribeVpcZoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVpcZoneResponse{}
	_body, _err := client.DescribeVpcZoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVulnerabilityProtectedListWithOptions(request *DescribeVulnerabilityProtectedListRequest, runtime *util.RuntimeOptions) (_result *DescribeVulnerabilityProtectedListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AttackType)) {
		query["AttackType"] = request.AttackType
	}

	if !tea.BoolValue(util.IsUnset(request.BuyVersion)) {
		query["BuyVersion"] = request.BuyVersion
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortKey)) {
		query["SortKey"] = request.SortKey
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.UserType)) {
		query["UserType"] = request.UserType
	}

	if !tea.BoolValue(util.IsUnset(request.VulnCveName)) {
		query["VulnCveName"] = request.VulnCveName
	}

	if !tea.BoolValue(util.IsUnset(request.VulnLevel)) {
		query["VulnLevel"] = request.VulnLevel
	}

	if !tea.BoolValue(util.IsUnset(request.VulnResource)) {
		query["VulnResource"] = request.VulnResource
	}

	if !tea.BoolValue(util.IsUnset(request.VulnStatus)) {
		query["VulnStatus"] = request.VulnStatus
	}

	if !tea.BoolValue(util.IsUnset(request.VulnType)) {
		query["VulnType"] = request.VulnType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVulnerabilityProtectedList"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVulnerabilityProtectedListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVulnerabilityProtectedList(request *DescribeVulnerabilityProtectedListRequest) (_result *DescribeVulnerabilityProtectedListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVulnerabilityProtectedListResponse{}
	_body, _err := client.DescribeVulnerabilityProtectedListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the ModifyAddressBook operation to modify the address book that is configured for access control.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyAddressBookRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyAddressBookResponse
 */
func (client *Client) ModifyAddressBookWithOptions(request *ModifyAddressBookRequest, runtime *util.RuntimeOptions) (_result *ModifyAddressBookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddressList)) {
		query["AddressList"] = request.AddressList
	}

	if !tea.BoolValue(util.IsUnset(request.AutoAddTagEcs)) {
		query["AutoAddTagEcs"] = request.AutoAddTagEcs
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupUuid)) {
		query["GroupUuid"] = request.GroupUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.TagList)) {
		query["TagList"] = request.TagList
	}

	if !tea.BoolValue(util.IsUnset(request.TagRelation)) {
		query["TagRelation"] = request.TagRelation
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAddressBook"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAddressBookResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the ModifyAddressBook operation to modify the address book that is configured for access control.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyAddressBookRequest
 * @return ModifyAddressBookResponse
 */
func (client *Client) ModifyAddressBook(request *ModifyAddressBookRequest) (_result *ModifyAddressBookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAddressBookResponse{}
	_body, _err := client.ModifyAddressBookWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to modify the configurations of an access control policy. The policy allows Cloud Firewall to allow, deny, or monitor the traffic that passes through Cloud Firewall.
 * ## [](#qps)Limit
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyControlPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyControlPolicyResponse
 */
func (client *Client) ModifyControlPolicyWithOptions(request *ModifyControlPolicyRequest, runtime *util.RuntimeOptions) (_result *ModifyControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclAction)) {
		query["AclAction"] = request.AclAction
	}

	if !tea.BoolValue(util.IsUnset(request.AclUuid)) {
		query["AclUuid"] = request.AclUuid
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicationName)) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicationNameList)) {
		query["ApplicationNameList"] = request.ApplicationNameList
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DestPort)) {
		query["DestPort"] = request.DestPort
	}

	if !tea.BoolValue(util.IsUnset(request.DestPortGroup)) {
		query["DestPortGroup"] = request.DestPortGroup
	}

	if !tea.BoolValue(util.IsUnset(request.DestPortType)) {
		query["DestPortType"] = request.DestPortType
	}

	if !tea.BoolValue(util.IsUnset(request.Destination)) {
		query["Destination"] = request.Destination
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationType)) {
		query["DestinationType"] = request.DestinationType
	}

	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		query["Direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Proto)) {
		query["Proto"] = request.Proto
	}

	if !tea.BoolValue(util.IsUnset(request.Release)) {
		query["Release"] = request.Release
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatDays)) {
		query["RepeatDays"] = request.RepeatDays
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatEndTime)) {
		query["RepeatEndTime"] = request.RepeatEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatStartTime)) {
		query["RepeatStartTime"] = request.RepeatStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatType)) {
		query["RepeatType"] = request.RepeatType
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyControlPolicy"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to modify the configurations of an access control policy. The policy allows Cloud Firewall to allow, deny, or monitor the traffic that passes through Cloud Firewall.
 * ## [](#qps)Limit
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyControlPolicyRequest
 * @return ModifyControlPolicyResponse
 */
func (client *Client) ModifyControlPolicy(request *ModifyControlPolicyRequest) (_result *ModifyControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyControlPolicyResponse{}
	_body, _err := client.ModifyControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can use this operation to modify the priority of an IPv4 access control policy for the Internet firewall. No API operations are provided for you to modify the priority of an IPv6 access control policy for the Internet firewall.
 * ## [](#qps)Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyControlPolicyPositionRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyControlPolicyPositionResponse
 */
func (client *Client) ModifyControlPolicyPositionWithOptions(request *ModifyControlPolicyPositionRequest, runtime *util.RuntimeOptions) (_result *ModifyControlPolicyPositionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		query["Direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.NewOrder)) {
		query["NewOrder"] = request.NewOrder
	}

	if !tea.BoolValue(util.IsUnset(request.OldOrder)) {
		query["OldOrder"] = request.OldOrder
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyControlPolicyPosition"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyControlPolicyPositionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can use this operation to modify the priority of an IPv4 access control policy for the Internet firewall. No API operations are provided for you to modify the priority of an IPv6 access control policy for the Internet firewall.
 * ## [](#qps)Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyControlPolicyPositionRequest
 * @return ModifyControlPolicyPositionResponse
 */
func (client *Client) ModifyControlPolicyPosition(request *ModifyControlPolicyPositionRequest) (_result *ModifyControlPolicyPositionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyControlPolicyPositionResponse{}
	_body, _err := client.ModifyControlPolicyPositionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDefaultIPSConfigWithOptions(request *ModifyDefaultIPSConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyDefaultIPSConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AiRules)) {
		query["AiRules"] = request.AiRules
	}

	if !tea.BoolValue(util.IsUnset(request.BasicRules)) {
		query["BasicRules"] = request.BasicRules
	}

	if !tea.BoolValue(util.IsUnset(request.CtiRules)) {
		query["CtiRules"] = request.CtiRules
	}

	if !tea.BoolValue(util.IsUnset(request.EnableAllPatch)) {
		query["EnableAllPatch"] = request.EnableAllPatch
	}

	if !tea.BoolValue(util.IsUnset(request.EnableDefault)) {
		query["EnableDefault"] = request.EnableDefault
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.PatchRules)) {
		query["PatchRules"] = request.PatchRules
	}

	if !tea.BoolValue(util.IsUnset(request.RuleClass)) {
		query["RuleClass"] = request.RuleClass
	}

	if !tea.BoolValue(util.IsUnset(request.RunMode)) {
		query["RunMode"] = request.RunMode
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDefaultIPSConfig"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDefaultIPSConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDefaultIPSConfig(request *ModifyDefaultIPSConfigRequest) (_result *ModifyDefaultIPSConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDefaultIPSConfigResponse{}
	_body, _err := client.ModifyDefaultIPSConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyFirewallV2RoutePolicySwitchWithOptions(request *ModifyFirewallV2RoutePolicySwitchRequest, runtime *util.RuntimeOptions) (_result *ModifyFirewallV2RoutePolicySwitchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FirewallId)) {
		query["FirewallId"] = request.FirewallId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.ShouldRecover)) {
		query["ShouldRecover"] = request.ShouldRecover
	}

	if !tea.BoolValue(util.IsUnset(request.TrFirewallRoutePolicyId)) {
		query["TrFirewallRoutePolicyId"] = request.TrFirewallRoutePolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.TrFirewallRoutePolicySwitchStatus)) {
		query["TrFirewallRoutePolicySwitchStatus"] = request.TrFirewallRoutePolicySwitchStatus
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyFirewallV2RoutePolicySwitch"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyFirewallV2RoutePolicySwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyFirewallV2RoutePolicySwitch(request *ModifyFirewallV2RoutePolicySwitchRequest) (_result *ModifyFirewallV2RoutePolicySwitchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyFirewallV2RoutePolicySwitchResponse{}
	_body, _err := client.ModifyFirewallV2RoutePolicySwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the ModifyInstanceMemberAttributes operation to update the information about members in Cloud Firewall.
 * ## Limits
 * You can call this operation up to 10 times per second for each account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyInstanceMemberAttributesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyInstanceMemberAttributesResponse
 */
func (client *Client) ModifyInstanceMemberAttributesWithOptions(request *ModifyInstanceMemberAttributesRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceMemberAttributesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Members)) {
		query["Members"] = request.Members
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyInstanceMemberAttributes"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyInstanceMemberAttributesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the ModifyInstanceMemberAttributes operation to update the information about members in Cloud Firewall.
 * ## Limits
 * You can call this operation up to 10 times per second for each account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyInstanceMemberAttributesRequest
 * @return ModifyInstanceMemberAttributesResponse
 */
func (client *Client) ModifyInstanceMemberAttributes(request *ModifyInstanceMemberAttributesRequest) (_result *ModifyInstanceMemberAttributesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceMemberAttributesResponse{}
	_body, _err := client.ModifyInstanceMemberAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can use this operation to modify the configurations of an access control policy. The policy is used to allow, deny, or monitor traffic that reaches a NAT firewall.
 *
 * @param request ModifyNatFirewallControlPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyNatFirewallControlPolicyResponse
 */
func (client *Client) ModifyNatFirewallControlPolicyWithOptions(request *ModifyNatFirewallControlPolicyRequest, runtime *util.RuntimeOptions) (_result *ModifyNatFirewallControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclAction)) {
		query["AclAction"] = request.AclAction
	}

	if !tea.BoolValue(util.IsUnset(request.AclUuid)) {
		query["AclUuid"] = request.AclUuid
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicationNameList)) {
		query["ApplicationNameList"] = request.ApplicationNameList
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DestPort)) {
		query["DestPort"] = request.DestPort
	}

	if !tea.BoolValue(util.IsUnset(request.DestPortGroup)) {
		query["DestPortGroup"] = request.DestPortGroup
	}

	if !tea.BoolValue(util.IsUnset(request.DestPortType)) {
		query["DestPortType"] = request.DestPortType
	}

	if !tea.BoolValue(util.IsUnset(request.Destination)) {
		query["Destination"] = request.Destination
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationType)) {
		query["DestinationType"] = request.DestinationType
	}

	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		query["Direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.DomainResolveType)) {
		query["DomainResolveType"] = request.DomainResolveType
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.NatGatewayId)) {
		query["NatGatewayId"] = request.NatGatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.Proto)) {
		query["Proto"] = request.Proto
	}

	if !tea.BoolValue(util.IsUnset(request.Release)) {
		query["Release"] = request.Release
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatDays)) {
		query["RepeatDays"] = request.RepeatDays
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatEndTime)) {
		query["RepeatEndTime"] = request.RepeatEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatStartTime)) {
		query["RepeatStartTime"] = request.RepeatStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatType)) {
		query["RepeatType"] = request.RepeatType
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyNatFirewallControlPolicy"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyNatFirewallControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can use this operation to modify the configurations of an access control policy. The policy is used to allow, deny, or monitor traffic that reaches a NAT firewall.
 *
 * @param request ModifyNatFirewallControlPolicyRequest
 * @return ModifyNatFirewallControlPolicyResponse
 */
func (client *Client) ModifyNatFirewallControlPolicy(request *ModifyNatFirewallControlPolicyRequest) (_result *ModifyNatFirewallControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyNatFirewallControlPolicyResponse{}
	_body, _err := client.ModifyNatFirewallControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyNatFirewallControlPolicyPositionWithOptions(request *ModifyNatFirewallControlPolicyPositionRequest, runtime *util.RuntimeOptions) (_result *ModifyNatFirewallControlPolicyPositionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclUuid)) {
		query["AclUuid"] = request.AclUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		query["Direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.NatGatewayId)) {
		query["NatGatewayId"] = request.NatGatewayId
	}

	if !tea.BoolValue(util.IsUnset(request.NewOrder)) {
		query["NewOrder"] = request.NewOrder
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyNatFirewallControlPolicyPosition"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyNatFirewallControlPolicyPositionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyNatFirewallControlPolicyPosition(request *ModifyNatFirewallControlPolicyPositionRequest) (_result *ModifyNatFirewallControlPolicyPositionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyNatFirewallControlPolicyPositionResponse{}
	_body, _err := client.ModifyNatFirewallControlPolicyPositionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the ModifyPolicyAdvancedConfig operation to enable or disable the strict mode for an access control policy.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyPolicyAdvancedConfigRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyPolicyAdvancedConfigResponse
 */
func (client *Client) ModifyPolicyAdvancedConfigWithOptions(request *ModifyPolicyAdvancedConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyPolicyAdvancedConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InternetSwitch)) {
		query["InternetSwitch"] = request.InternetSwitch
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyPolicyAdvancedConfig"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyPolicyAdvancedConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the ModifyPolicyAdvancedConfig operation to enable or disable the strict mode for an access control policy.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyPolicyAdvancedConfigRequest
 * @return ModifyPolicyAdvancedConfigResponse
 */
func (client *Client) ModifyPolicyAdvancedConfig(request *ModifyPolicyAdvancedConfigRequest) (_result *ModifyPolicyAdvancedConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyPolicyAdvancedConfigResponse{}
	_body, _err := client.ModifyPolicyAdvancedConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyTrFirewallV2ConfigurationWithOptions(request *ModifyTrFirewallV2ConfigurationRequest, runtime *util.RuntimeOptions) (_result *ModifyTrFirewallV2ConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FirewallId)) {
		query["FirewallId"] = request.FirewallId
	}

	if !tea.BoolValue(util.IsUnset(request.FirewallName)) {
		query["FirewallName"] = request.FirewallName
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyTrFirewallV2Configuration"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyTrFirewallV2ConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyTrFirewallV2Configuration(request *ModifyTrFirewallV2ConfigurationRequest) (_result *ModifyTrFirewallV2ConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyTrFirewallV2ConfigurationResponse{}
	_body, _err := client.ModifyTrFirewallV2ConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyTrFirewallV2RoutePolicyScopeWithOptions(tmpReq *ModifyTrFirewallV2RoutePolicyScopeRequest, runtime *util.RuntimeOptions) (_result *ModifyTrFirewallV2RoutePolicyScopeResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ModifyTrFirewallV2RoutePolicyScopeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DestCandidateList)) {
		request.DestCandidateListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DestCandidateList, tea.String("DestCandidateList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SrcCandidateList)) {
		request.SrcCandidateListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SrcCandidateList, tea.String("SrcCandidateList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DestCandidateListShrink)) {
		query["DestCandidateList"] = request.DestCandidateListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.FirewallId)) {
		query["FirewallId"] = request.FirewallId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.ShouldRecover)) {
		query["ShouldRecover"] = request.ShouldRecover
	}

	if !tea.BoolValue(util.IsUnset(request.SrcCandidateListShrink)) {
		query["SrcCandidateList"] = request.SrcCandidateListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TrFirewallRoutePolicyId)) {
		query["TrFirewallRoutePolicyId"] = request.TrFirewallRoutePolicyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyTrFirewallV2RoutePolicyScope"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyTrFirewallV2RoutePolicyScopeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyTrFirewallV2RoutePolicyScope(request *ModifyTrFirewallV2RoutePolicyScopeRequest) (_result *ModifyTrFirewallV2RoutePolicyScopeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyTrFirewallV2RoutePolicyScopeResponse{}
	_body, _err := client.ModifyTrFirewallV2RoutePolicyScopeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyUserIPSWhitelistWithOptions(request *ModifyUserIPSWhitelistRequest, runtime *util.RuntimeOptions) (_result *ModifyUserIPSWhitelistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Direction)) {
		query["Direction"] = request.Direction
	}

	if !tea.BoolValue(util.IsUnset(request.IpVersion)) {
		query["IpVersion"] = request.IpVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.ListType)) {
		query["ListType"] = request.ListType
	}

	if !tea.BoolValue(util.IsUnset(request.ListValue)) {
		query["ListValue"] = request.ListValue
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.WhiteType)) {
		query["WhiteType"] = request.WhiteType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyUserIPSWhitelist"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyUserIPSWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyUserIPSWhitelist(request *ModifyUserIPSWhitelistRequest) (_result *ModifyUserIPSWhitelistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyUserIPSWhitelistResponse{}
	_body, _err := client.ModifyUserIPSWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the ModifyVpcFirewallCenConfigure operation to modify the configurations of a VPC firewall. The VPC firewall protects mutual access traffic between a VPC and a specified network instance that is attached to a CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance. Before you call this operation, make sure that you have created a VPC firewall by calling the [CreateVpcFirewallCenConfigure](~~345772~~) operation.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyVpcFirewallCenConfigureRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyVpcFirewallCenConfigureResponse
 */
func (client *Client) ModifyVpcFirewallCenConfigureWithOptions(request *ModifyVpcFirewallCenConfigureRequest, runtime *util.RuntimeOptions) (_result *ModifyVpcFirewallCenConfigureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallId)) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallName)) {
		query["VpcFirewallName"] = request.VpcFirewallName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyVpcFirewallCenConfigure"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyVpcFirewallCenConfigureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the ModifyVpcFirewallCenConfigure operation to modify the configurations of a VPC firewall. The VPC firewall protects mutual access traffic between a VPC and a specified network instance that is attached to a CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance. Before you call this operation, make sure that you have created a VPC firewall by calling the [CreateVpcFirewallCenConfigure](~~345772~~) operation.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyVpcFirewallCenConfigureRequest
 * @return ModifyVpcFirewallCenConfigureResponse
 */
func (client *Client) ModifyVpcFirewallCenConfigure(request *ModifyVpcFirewallCenConfigureRequest) (_result *ModifyVpcFirewallCenConfigureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyVpcFirewallCenConfigureResponse{}
	_body, _err := client.ModifyVpcFirewallCenConfigureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the ModifyVpcFirewallCenSwitchStatus operation to enable or disable a VPC firewall. A VPC firewall protects mutual access traffic between a specified VPC and a network instance that is attached to a CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance. After you enable the VPC firewall, the VPC firewall protects mutual access traffic between a VPC and a specified network instance that is attached to a CEN instance. After you disable the VPC firewall, the VPC firewall no longer protect mutual access traffic between a VPC and a specified network instance that is attached to a CEN instance.
 * Before you call this operation, make sure that you have created a VPC firewall by calling the [CreateVpcFirewallCenConfigure](~~345772~~) operation.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyVpcFirewallCenSwitchStatusRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyVpcFirewallCenSwitchStatusResponse
 */
func (client *Client) ModifyVpcFirewallCenSwitchStatusWithOptions(request *ModifyVpcFirewallCenSwitchStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyVpcFirewallCenSwitchStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FirewallSwitch)) {
		query["FirewallSwitch"] = request.FirewallSwitch
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallId)) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyVpcFirewallCenSwitchStatus"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyVpcFirewallCenSwitchStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the ModifyVpcFirewallCenSwitchStatus operation to enable or disable a VPC firewall. A VPC firewall protects mutual access traffic between a specified VPC and a network instance that is attached to a CEN instance. The network instance can be a VPC, a virtual border router (VBR), or a Cloud Connect Network (CCN) instance. After you enable the VPC firewall, the VPC firewall protects mutual access traffic between a VPC and a specified network instance that is attached to a CEN instance. After you disable the VPC firewall, the VPC firewall no longer protect mutual access traffic between a VPC and a specified network instance that is attached to a CEN instance.
 * Before you call this operation, make sure that you have created a VPC firewall by calling the [CreateVpcFirewallCenConfigure](~~345772~~) operation.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyVpcFirewallCenSwitchStatusRequest
 * @return ModifyVpcFirewallCenSwitchStatusResponse
 */
func (client *Client) ModifyVpcFirewallCenSwitchStatus(request *ModifyVpcFirewallCenSwitchStatusRequest) (_result *ModifyVpcFirewallCenSwitchStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyVpcFirewallCenSwitchStatusResponse{}
	_body, _err := client.ModifyVpcFirewallCenSwitchStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the ModifyVpcFirewallConfigure operation to modify the configurations of a VPC firewall. The VPC firewall controls traffic between two VPCs that are connected by using an Express Connect circuit. Before you call the operation, make sure that you created a VPC firewall by calling the [CreateVpcFirewallConfigure](~~342893~~) operation.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyVpcFirewallConfigureRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyVpcFirewallConfigureResponse
 */
func (client *Client) ModifyVpcFirewallConfigureWithOptions(request *ModifyVpcFirewallConfigureRequest, runtime *util.RuntimeOptions) (_result *ModifyVpcFirewallConfigureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.LocalVpcCidrTableList)) {
		query["LocalVpcCidrTableList"] = request.LocalVpcCidrTableList
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.PeerVpcCidrTableList)) {
		query["PeerVpcCidrTableList"] = request.PeerVpcCidrTableList
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallId)) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallName)) {
		query["VpcFirewallName"] = request.VpcFirewallName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyVpcFirewallConfigure"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyVpcFirewallConfigureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the ModifyVpcFirewallConfigure operation to modify the configurations of a VPC firewall. The VPC firewall controls traffic between two VPCs that are connected by using an Express Connect circuit. Before you call the operation, make sure that you created a VPC firewall by calling the [CreateVpcFirewallConfigure](~~342893~~) operation.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyVpcFirewallConfigureRequest
 * @return ModifyVpcFirewallConfigureResponse
 */
func (client *Client) ModifyVpcFirewallConfigure(request *ModifyVpcFirewallConfigureRequest) (_result *ModifyVpcFirewallConfigureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyVpcFirewallConfigureResponse{}
	_body, _err := client.ModifyVpcFirewallConfigureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the ModifyVpcFirewallControlPolicy operation to modify the configurations of an access control policy that is created for a VPC firewall in a specified policy group. Different access control policies are used for the VPC firewalls that are used to protect each Cloud Enterprise Network (CEN) instance and the VPC firewalls that are used to protect each Express Connect circuit.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyVpcFirewallControlPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyVpcFirewallControlPolicyResponse
 */
func (client *Client) ModifyVpcFirewallControlPolicyWithOptions(request *ModifyVpcFirewallControlPolicyRequest, runtime *util.RuntimeOptions) (_result *ModifyVpcFirewallControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclAction)) {
		query["AclAction"] = request.AclAction
	}

	if !tea.BoolValue(util.IsUnset(request.AclUuid)) {
		query["AclUuid"] = request.AclUuid
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicationName)) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicationNameList)) {
		query["ApplicationNameList"] = request.ApplicationNameList
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DestPort)) {
		query["DestPort"] = request.DestPort
	}

	if !tea.BoolValue(util.IsUnset(request.DestPortGroup)) {
		query["DestPortGroup"] = request.DestPortGroup
	}

	if !tea.BoolValue(util.IsUnset(request.DestPortType)) {
		query["DestPortType"] = request.DestPortType
	}

	if !tea.BoolValue(util.IsUnset(request.Destination)) {
		query["Destination"] = request.Destination
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationType)) {
		query["DestinationType"] = request.DestinationType
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Proto)) {
		query["Proto"] = request.Proto
	}

	if !tea.BoolValue(util.IsUnset(request.Release)) {
		query["Release"] = request.Release
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatDays)) {
		query["RepeatDays"] = request.RepeatDays
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatEndTime)) {
		query["RepeatEndTime"] = request.RepeatEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatStartTime)) {
		query["RepeatStartTime"] = request.RepeatStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatType)) {
		query["RepeatType"] = request.RepeatType
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallId)) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyVpcFirewallControlPolicy"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyVpcFirewallControlPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the ModifyVpcFirewallControlPolicy operation to modify the configurations of an access control policy that is created for a VPC firewall in a specified policy group. Different access control policies are used for the VPC firewalls that are used to protect each Cloud Enterprise Network (CEN) instance and the VPC firewalls that are used to protect each Express Connect circuit.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyVpcFirewallControlPolicyRequest
 * @return ModifyVpcFirewallControlPolicyResponse
 */
func (client *Client) ModifyVpcFirewallControlPolicy(request *ModifyVpcFirewallControlPolicyRequest) (_result *ModifyVpcFirewallControlPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyVpcFirewallControlPolicyResponse{}
	_body, _err := client.ModifyVpcFirewallControlPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can use this operation to modify the priority of an access control policy that is created for a VPC firewall in a specific policy group.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyVpcFirewallControlPolicyPositionRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyVpcFirewallControlPolicyPositionResponse
 */
func (client *Client) ModifyVpcFirewallControlPolicyPositionWithOptions(request *ModifyVpcFirewallControlPolicyPositionRequest, runtime *util.RuntimeOptions) (_result *ModifyVpcFirewallControlPolicyPositionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclUuid)) {
		query["AclUuid"] = request.AclUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.NewOrder)) {
		query["NewOrder"] = request.NewOrder
	}

	if !tea.BoolValue(util.IsUnset(request.OldOrder)) {
		query["OldOrder"] = request.OldOrder
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallId)) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyVpcFirewallControlPolicyPosition"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyVpcFirewallControlPolicyPositionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can use this operation to modify the priority of an access control policy that is created for a VPC firewall in a specific policy group.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyVpcFirewallControlPolicyPositionRequest
 * @return ModifyVpcFirewallControlPolicyPositionResponse
 */
func (client *Client) ModifyVpcFirewallControlPolicyPosition(request *ModifyVpcFirewallControlPolicyPositionRequest) (_result *ModifyVpcFirewallControlPolicyPositionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyVpcFirewallControlPolicyPositionResponse{}
	_body, _err := client.ModifyVpcFirewallControlPolicyPositionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the ModifyVpcFirewallDefaultIPSConfig operation to modify the intrusion prevention configurations of a VPC firewall.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyVpcFirewallDefaultIPSConfigRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyVpcFirewallDefaultIPSConfigResponse
 */
func (client *Client) ModifyVpcFirewallDefaultIPSConfigWithOptions(request *ModifyVpcFirewallDefaultIPSConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyVpcFirewallDefaultIPSConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BasicRules)) {
		query["BasicRules"] = request.BasicRules
	}

	if !tea.BoolValue(util.IsUnset(request.EnableAllPatch)) {
		query["EnableAllPatch"] = request.EnableAllPatch
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.RunMode)) {
		query["RunMode"] = request.RunMode
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallId)) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyVpcFirewallDefaultIPSConfig"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyVpcFirewallDefaultIPSConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the ModifyVpcFirewallDefaultIPSConfig operation to modify the intrusion prevention configurations of a VPC firewall.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyVpcFirewallDefaultIPSConfigRequest
 * @return ModifyVpcFirewallDefaultIPSConfigResponse
 */
func (client *Client) ModifyVpcFirewallDefaultIPSConfig(request *ModifyVpcFirewallDefaultIPSConfigRequest) (_result *ModifyVpcFirewallDefaultIPSConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyVpcFirewallDefaultIPSConfigResponse{}
	_body, _err := client.ModifyVpcFirewallDefaultIPSConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyVpcFirewallIPSWhitelistWithOptions(request *ModifyVpcFirewallIPSWhitelistRequest, runtime *util.RuntimeOptions) (_result *ModifyVpcFirewallIPSWhitelistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.ListType)) {
		query["ListType"] = request.ListType
	}

	if !tea.BoolValue(util.IsUnset(request.ListValue)) {
		query["ListValue"] = request.ListValue
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallId)) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	if !tea.BoolValue(util.IsUnset(request.WhiteType)) {
		query["WhiteType"] = request.WhiteType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyVpcFirewallIPSWhitelist"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyVpcFirewallIPSWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyVpcFirewallIPSWhitelist(request *ModifyVpcFirewallIPSWhitelistRequest) (_result *ModifyVpcFirewallIPSWhitelistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyVpcFirewallIPSWhitelistResponse{}
	_body, _err := client.ModifyVpcFirewallIPSWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the ModifyVpcFirewallSwitchStatus operation to enable or disable a VPC firewall. The VPC firewall protects traffic between two VPCs that are connected by using an Express Connect circuit. After you enable the VPC firewall, the VPC firewall protects access traffic between two VPCs that are connected by using an Express Connect circuit. After you disable the VPC firewall, the VPC firewall no longer protects access traffic between two VPCs that are connected by using an Express Connect circuit.
 * Before you call the operation, make sure that you created a VPC firewall by calling the [CreateVpcFirewallConfigure](~~342893~~) operation.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyVpcFirewallSwitchStatusRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyVpcFirewallSwitchStatusResponse
 */
func (client *Client) ModifyVpcFirewallSwitchStatusWithOptions(request *ModifyVpcFirewallSwitchStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyVpcFirewallSwitchStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FirewallSwitch)) {
		query["FirewallSwitch"] = request.FirewallSwitch
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.VpcFirewallId)) {
		query["VpcFirewallId"] = request.VpcFirewallId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyVpcFirewallSwitchStatus"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyVpcFirewallSwitchStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the ModifyVpcFirewallSwitchStatus operation to enable or disable a VPC firewall. The VPC firewall protects traffic between two VPCs that are connected by using an Express Connect circuit. After you enable the VPC firewall, the VPC firewall protects access traffic between two VPCs that are connected by using an Express Connect circuit. After you disable the VPC firewall, the VPC firewall no longer protects access traffic between two VPCs that are connected by using an Express Connect circuit.
 * Before you call the operation, make sure that you created a VPC firewall by calling the [CreateVpcFirewallConfigure](~~342893~~) operation.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ModifyVpcFirewallSwitchStatusRequest
 * @return ModifyVpcFirewallSwitchStatusResponse
 */
func (client *Client) ModifyVpcFirewallSwitchStatus(request *ModifyVpcFirewallSwitchStatusRequest) (_result *ModifyVpcFirewallSwitchStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyVpcFirewallSwitchStatusResponse{}
	_body, _err := client.ModifyVpcFirewallSwitchStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the PutDisableAllFwSwitch operation to turn off all firewall switches.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request PutDisableAllFwSwitchRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return PutDisableAllFwSwitchResponse
 */
func (client *Client) PutDisableAllFwSwitchWithOptions(request *PutDisableAllFwSwitchRequest, runtime *util.RuntimeOptions) (_result *PutDisableAllFwSwitchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PutDisableAllFwSwitch"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PutDisableAllFwSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the PutDisableAllFwSwitch operation to turn off all firewall switches.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request PutDisableAllFwSwitchRequest
 * @return PutDisableAllFwSwitchResponse
 */
func (client *Client) PutDisableAllFwSwitch(request *PutDisableAllFwSwitchRequest) (_result *PutDisableAllFwSwitchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutDisableAllFwSwitchResponse{}
	_body, _err := client.PutDisableAllFwSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the PutDisableFwSwitch operation to disable a firewall for specific assets. After you disable the firewall, traffic does not pass through Cloud Firewall.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request PutDisableFwSwitchRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return PutDisableFwSwitchResponse
 */
func (client *Client) PutDisableFwSwitchWithOptions(request *PutDisableFwSwitchRequest, runtime *util.RuntimeOptions) (_result *PutDisableFwSwitchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IpaddrList)) {
		query["IpaddrList"] = request.IpaddrList
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.RegionList)) {
		query["RegionList"] = request.RegionList
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceTypeList)) {
		query["ResourceTypeList"] = request.ResourceTypeList
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PutDisableFwSwitch"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PutDisableFwSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the PutDisableFwSwitch operation to disable a firewall for specific assets. After you disable the firewall, traffic does not pass through Cloud Firewall.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request PutDisableFwSwitchRequest
 * @return PutDisableFwSwitchResponse
 */
func (client *Client) PutDisableFwSwitch(request *PutDisableFwSwitchRequest) (_result *PutDisableFwSwitchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutDisableFwSwitchResponse{}
	_body, _err := client.PutDisableFwSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the PutEnableAllFwSwitch operation to enable a firewall for all public IP addresses within your Alibaba Cloud account.
 * ## Limits
 * You can call this operation up to 10 times per second per account. You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request PutEnableAllFwSwitchRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return PutEnableAllFwSwitchResponse
 */
func (client *Client) PutEnableAllFwSwitchWithOptions(request *PutEnableAllFwSwitchRequest, runtime *util.RuntimeOptions) (_result *PutEnableAllFwSwitchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PutEnableAllFwSwitch"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PutEnableAllFwSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the PutEnableAllFwSwitch operation to enable a firewall for all public IP addresses within your Alibaba Cloud account.
 * ## Limits
 * You can call this operation up to 10 times per second per account. You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request PutEnableAllFwSwitchRequest
 * @return PutEnableAllFwSwitchResponse
 */
func (client *Client) PutEnableAllFwSwitch(request *PutEnableAllFwSwitchRequest) (_result *PutEnableAllFwSwitchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutEnableAllFwSwitchResponse{}
	_body, _err := client.PutEnableAllFwSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the PutEnableFwSwitch operation to enable a firewall. After you enable a firewall, traffic passes through Cloud Firewall.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds a limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limits when you call this operation.
 *
 * @param request PutEnableFwSwitchRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return PutEnableFwSwitchResponse
 */
func (client *Client) PutEnableFwSwitchWithOptions(request *PutEnableFwSwitchRequest, runtime *util.RuntimeOptions) (_result *PutEnableFwSwitchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IpaddrList)) {
		query["IpaddrList"] = request.IpaddrList
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.RegionList)) {
		query["RegionList"] = request.RegionList
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceTypeList)) {
		query["ResourceTypeList"] = request.ResourceTypeList
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PutEnableFwSwitch"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PutEnableFwSwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the PutEnableFwSwitch operation to enable a firewall. After you enable a firewall, traffic passes through Cloud Firewall.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds a limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limits when you call this operation.
 *
 * @param request PutEnableFwSwitchRequest
 * @return PutEnableFwSwitchResponse
 */
func (client *Client) PutEnableFwSwitch(request *PutEnableFwSwitchRequest) (_result *PutEnableFwSwitchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutEnableFwSwitchResponse{}
	_body, _err := client.PutEnableFwSwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleasePostInstanceWithOptions(request *ReleasePostInstanceRequest, runtime *util.RuntimeOptions) (_result *ReleasePostInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleasePostInstance"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleasePostInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleasePostInstance(request *ReleasePostInstanceRequest) (_result *ReleasePostInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleasePostInstanceResponse{}
	_body, _err := client.ReleasePostInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetNatFirewallRuleHitCountWithOptions(request *ResetNatFirewallRuleHitCountRequest, runtime *util.RuntimeOptions) (_result *ResetNatFirewallRuleHitCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclUuid)) {
		query["AclUuid"] = request.AclUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.NatGatewayId)) {
		query["NatGatewayId"] = request.NatGatewayId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetNatFirewallRuleHitCount"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetNatFirewallRuleHitCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetNatFirewallRuleHitCount(request *ResetNatFirewallRuleHitCountRequest) (_result *ResetNatFirewallRuleHitCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetNatFirewallRuleHitCountResponse{}
	_body, _err := client.ResetNatFirewallRuleHitCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the ResetVpcFirewallRuleHitCount operation to clear the count on hits of an access control policy that is created for a VPC firewall in a specific policy group.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ResetVpcFirewallRuleHitCountRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ResetVpcFirewallRuleHitCountResponse
 */
func (client *Client) ResetVpcFirewallRuleHitCountWithOptions(request *ResetVpcFirewallRuleHitCountRequest, runtime *util.RuntimeOptions) (_result *ResetVpcFirewallRuleHitCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclUuid)) {
		query["AclUuid"] = request.AclUuid
	}

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetVpcFirewallRuleHitCount"),
		Version:     tea.String("2017-12-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetVpcFirewallRuleHitCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the ResetVpcFirewallRuleHitCount operation to clear the count on hits of an access control policy that is created for a VPC firewall in a specific policy group.
 * ## Limits
 * You can call this operation up to 10 times per second per account. If the number of the calls per second exceeds the limit, throttling is triggered. As a result, your business may be affected. We recommend that you take note of the limit when you call this operation.
 *
 * @param request ResetVpcFirewallRuleHitCountRequest
 * @return ResetVpcFirewallRuleHitCountResponse
 */
func (client *Client) ResetVpcFirewallRuleHitCount(request *ResetVpcFirewallRuleHitCountRequest) (_result *ResetVpcFirewallRuleHitCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetVpcFirewallRuleHitCountResponse{}
	_body, _err := client.ResetVpcFirewallRuleHitCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
