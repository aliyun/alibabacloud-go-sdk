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

type AddAddressBookRequest struct {
	AddressList   *string                         `json:"AddressList,omitempty" xml:"AddressList,omitempty"`
	AutoAddTagEcs *string                         `json:"AutoAddTagEcs,omitempty" xml:"AutoAddTagEcs,omitempty"`
	Description   *string                         `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupName     *string                         `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	GroupType     *string                         `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	Lang          *string                         `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp      *string                         `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	TagList       []*AddAddressBookRequestTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	TagRelation   *string                         `json:"TagRelation,omitempty" xml:"TagRelation,omitempty"`
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
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
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
	GroupUuid *string `json:"GroupUuid,omitempty" xml:"GroupUuid,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddAddressBookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AclAction           *string   `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	ApplicationName     *string   `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	ApplicationNameList []*string `json:"ApplicationNameList,omitempty" xml:"ApplicationNameList,omitempty" type:"Repeated"`
	Description         *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	DestPort            *string   `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	DestPortGroup       *string   `json:"DestPortGroup,omitempty" xml:"DestPortGroup,omitempty"`
	DestPortType        *string   `json:"DestPortType,omitempty" xml:"DestPortType,omitempty"`
	Destination         *string   `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DestinationType     *string   `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	Direction           *string   `json:"Direction,omitempty" xml:"Direction,omitempty"`
	IpVersion           *string   `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	Lang                *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	NewOrder            *string   `json:"NewOrder,omitempty" xml:"NewOrder,omitempty"`
	Proto               *string   `json:"Proto,omitempty" xml:"Proto,omitempty"`
	Release             *string   `json:"Release,omitempty" xml:"Release,omitempty"`
	Source              *string   `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceIp            *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	SourceType          *string   `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
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

type AddControlPolicyResponseBody struct {
	AclUuid   *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	MemberDesc *string `json:"MemberDesc,omitempty" xml:"MemberDesc,omitempty"`
	MemberUid  *int64  `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddInstanceMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Lang                *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp            *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	SourceVpcFirewallId *string `json:"SourceVpcFirewallId,omitempty" xml:"SourceVpcFirewallId,omitempty"`
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
	// Id of the request
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
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchCopyVpcFirewallControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type CreateVpcFirewallCenConfigureRequest struct {
	CenId             *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	FirewallSwitch    *string `json:"FirewallSwitch,omitempty" xml:"FirewallSwitch,omitempty"`
	Lang              *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	MemberUid         *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	NetworkInstanceId *string `json:"NetworkInstanceId,omitempty" xml:"NetworkInstanceId,omitempty"`
	VpcFirewallName   *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
	VpcRegion         *string `json:"VpcRegion,omitempty" xml:"VpcRegion,omitempty"`
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

func (s *CreateVpcFirewallCenConfigureRequest) SetVpcFirewallName(v string) *CreateVpcFirewallCenConfigureRequest {
	s.VpcFirewallName = &v
	return s
}

func (s *CreateVpcFirewallCenConfigureRequest) SetVpcRegion(v string) *CreateVpcFirewallCenConfigureRequest {
	s.VpcRegion = &v
	return s
}

type CreateVpcFirewallCenConfigureResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateVpcFirewallCenConfigureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	FirewallSwitch        *string `json:"FirewallSwitch,omitempty" xml:"FirewallSwitch,omitempty"`
	Lang                  *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	LocalVpcCidrTableList *string `json:"LocalVpcCidrTableList,omitempty" xml:"LocalVpcCidrTableList,omitempty"`
	LocalVpcId            *string `json:"LocalVpcId,omitempty" xml:"LocalVpcId,omitempty"`
	LocalVpcRegion        *string `json:"LocalVpcRegion,omitempty" xml:"LocalVpcRegion,omitempty"`
	MemberUid             *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	PeerVpcCidrTableList  *string `json:"PeerVpcCidrTableList,omitempty" xml:"PeerVpcCidrTableList,omitempty"`
	PeerVpcId             *string `json:"PeerVpcId,omitempty" xml:"PeerVpcId,omitempty"`
	PeerVpcRegion         *string `json:"PeerVpcRegion,omitempty" xml:"PeerVpcRegion,omitempty"`
	VpcFirewallName       *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
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
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateVpcFirewallConfigureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AclAction       *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DestPort        *string `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	DestPortGroup   *string `json:"DestPortGroup,omitempty" xml:"DestPortGroup,omitempty"`
	DestPortType    *string `json:"DestPortType,omitempty" xml:"DestPortType,omitempty"`
	Destination     *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	MemberUid       *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	NewOrder        *string `json:"NewOrder,omitempty" xml:"NewOrder,omitempty"`
	Proto           *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	Release         *string `json:"Release,omitempty" xml:"Release,omitempty"`
	Source          *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceType      *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	VpcFirewallId   *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
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

func (s *CreateVpcFirewallControlPolicyRequest) SetSource(v string) *CreateVpcFirewallControlPolicyRequest {
	s.Source = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetSourceType(v string) *CreateVpcFirewallControlPolicyRequest {
	s.SourceType = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetVpcFirewallId(v string) *CreateVpcFirewallControlPolicyRequest {
	s.VpcFirewallId = &v
	return s
}

type CreateVpcFirewallControlPolicyResponseBody struct {
	AclUuid   *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
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
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateVpcFirewallControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	GroupUuid *string `json:"GroupUuid,omitempty" xml:"GroupUuid,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAddressBookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AclUuid   *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DeleteInstanceMembersRequest struct {
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteInstanceMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DeleteVpcFirewallCenConfigureRequest struct {
	Lang              *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	MemberUid         *string   `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteVpcFirewallCenConfigureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Lang              *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	MemberUid         *string   `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteVpcFirewallConfigureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AclUuid       *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	Lang          *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
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
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteVpcFirewallControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DescribeAddressBookRequest struct {
	ContainPort *string `json:"ContainPort,omitempty" xml:"ContainPort,omitempty"`
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	GroupType   *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageSize    *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Query       *string `json:"Query,omitempty" xml:"Query,omitempty"`
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
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

func (s *DescribeAddressBookRequest) SetSourceIp(v string) *DescribeAddressBookRequest {
	s.SourceIp = &v
	return s
}

type DescribeAddressBookResponseBody struct {
	Acls       []*DescribeAddressBookResponseBodyAcls `json:"Acls,omitempty" xml:"Acls,omitempty" type:"Repeated"`
	PageNo     *string                                `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize   *string                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *string                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	AddressList      []*string                                     `json:"AddressList,omitempty" xml:"AddressList,omitempty" type:"Repeated"`
	AddressListCount *int32                                        `json:"AddressListCount,omitempty" xml:"AddressListCount,omitempty"`
	AutoAddTagEcs    *int32                                        `json:"AutoAddTagEcs,omitempty" xml:"AutoAddTagEcs,omitempty"`
	Description      *string                                       `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupName        *string                                       `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	GroupType        *string                                       `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	GroupUuid        *string                                       `json:"GroupUuid,omitempty" xml:"GroupUuid,omitempty"`
	ReferenceCount   *int32                                        `json:"ReferenceCount,omitempty" xml:"ReferenceCount,omitempty"`
	TagList          []*DescribeAddressBookResponseBodyAclsTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	TagRelation      *string                                       `json:"TagRelation,omitempty" xml:"TagRelation,omitempty"`
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
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAddressBookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	CurrentPage  *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	IpVersion    *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	MemberUid    *int64  `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	PageSize     *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionNo     *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SearchItem   *string `json:"SearchItem,omitempty" xml:"SearchItem,omitempty"`
	SgStatus     *string `json:"SgStatus,omitempty" xml:"SgStatus,omitempty"`
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UserType     *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
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

func (s *DescribeAssetListRequest) SetSourceIp(v string) *DescribeAssetListRequest {
	s.SourceIp = &v
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
	Assets     []*DescribeAssetListResponseBodyAssets `json:"Assets,omitempty" xml:"Assets,omitempty" type:"Repeated"`
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	AliUid             *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	BindInstanceId     *string `json:"BindInstanceId,omitempty" xml:"BindInstanceId,omitempty"`
	BindInstanceName   *string `json:"BindInstanceName,omitempty" xml:"BindInstanceName,omitempty"`
	InternetAddress    *string `json:"InternetAddress,omitempty" xml:"InternetAddress,omitempty"`
	IntranetAddress    *string `json:"IntranetAddress,omitempty" xml:"IntranetAddress,omitempty"`
	IpVersion          *int32  `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	MemberUid          *int64  `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Note               *string `json:"Note,omitempty" xml:"Note,omitempty"`
	ProtectStatus      *string `json:"ProtectStatus,omitempty" xml:"ProtectStatus,omitempty"`
	RegionID           *string `json:"RegionID,omitempty" xml:"RegionID,omitempty"`
	RegionStatus       *string `json:"RegionStatus,omitempty" xml:"RegionStatus,omitempty"`
	ResourceInstanceId *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
	ResourceType       *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	RiskLevel          *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	SgStatus           *string `json:"SgStatus,omitempty" xml:"SgStatus,omitempty"`
	SgStatusTime       *int64  `json:"SgStatusTime,omitempty" xml:"SgStatusTime,omitempty"`
	SyncStatus         *string `json:"SyncStatus,omitempty" xml:"SyncStatus,omitempty"`
	Type               *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAssetListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DescribeControlPolicyRequest struct {
	AclAction   *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	AclUuid     *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	Direction   *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	IpVersion   *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	Lang        *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageSize    *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Proto       *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	Release     *string `json:"Release,omitempty" xml:"Release,omitempty"`
	Source      *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
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

func (s *DescribeControlPolicyRequest) SetSource(v string) *DescribeControlPolicyRequest {
	s.Source = &v
	return s
}

func (s *DescribeControlPolicyRequest) SetSourceIp(v string) *DescribeControlPolicyRequest {
	s.SourceIp = &v
	return s
}

type DescribeControlPolicyResponseBody struct {
	PageNo     *string                                     `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize   *string                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Policys    []*DescribeControlPolicyResponseBodyPolicys `json:"Policys,omitempty" xml:"Policys,omitempty" type:"Repeated"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *string                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	AclAction             *string   `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	AclUuid               *string   `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	ApplicationId         *string   `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	ApplicationName       *string   `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	ApplicationNameList   []*string `json:"ApplicationNameList,omitempty" xml:"ApplicationNameList,omitempty" type:"Repeated"`
	Description           *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	DestPort              *string   `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	DestPortGroup         *string   `json:"DestPortGroup,omitempty" xml:"DestPortGroup,omitempty"`
	DestPortGroupPorts    []*string `json:"DestPortGroupPorts,omitempty" xml:"DestPortGroupPorts,omitempty" type:"Repeated"`
	DestPortType          *string   `json:"DestPortType,omitempty" xml:"DestPortType,omitempty"`
	Destination           *string   `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DestinationGroupCidrs []*string `json:"DestinationGroupCidrs,omitempty" xml:"DestinationGroupCidrs,omitempty" type:"Repeated"`
	DestinationGroupType  *string   `json:"DestinationGroupType,omitempty" xml:"DestinationGroupType,omitempty"`
	DestinationType       *string   `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	Direction             *string   `json:"Direction,omitempty" xml:"Direction,omitempty"`
	DnsResult             *string   `json:"DnsResult,omitempty" xml:"DnsResult,omitempty"`
	DnsResultTime         *int64    `json:"DnsResultTime,omitempty" xml:"DnsResultTime,omitempty"`
	HitLastTime           *int64    `json:"HitLastTime,omitempty" xml:"HitLastTime,omitempty"`
	HitTimes              *int64    `json:"HitTimes,omitempty" xml:"HitTimes,omitempty"`
	IpVersion             *int32    `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	Order                 *int32    `json:"Order,omitempty" xml:"Order,omitempty"`
	Proto                 *string   `json:"Proto,omitempty" xml:"Proto,omitempty"`
	Release               *string   `json:"Release,omitempty" xml:"Release,omitempty"`
	Source                *string   `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceGroupCidrs      []*string `json:"SourceGroupCidrs,omitempty" xml:"SourceGroupCidrs,omitempty" type:"Repeated"`
	SourceGroupType       *string   `json:"SourceGroupType,omitempty" xml:"SourceGroupType,omitempty"`
	SourceType            *string   `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
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

type DescribeControlPolicyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DescribeDomainResolveRequest struct {
	Domain    *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
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
	RequestId     *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	IpAddrs    *string `json:"IpAddrs,omitempty" xml:"IpAddrs,omitempty"`
	UpdateTime *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDomainResolveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DescribeInstanceMembersRequest struct {
	CurrentPage       *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	MemberDesc        *string `json:"MemberDesc,omitempty" xml:"MemberDesc,omitempty"`
	MemberDisplayName *string `json:"MemberDisplayName,omitempty" xml:"MemberDisplayName,omitempty"`
	MemberUid         *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	PageSize          *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	Members   []*DescribeInstanceMembersResponseBodyMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	PageInfo  *DescribeInstanceMembersResponseBodyPageInfo  `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CreateTime        *int32  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	MemberDesc        *string `json:"MemberDesc,omitempty" xml:"MemberDesc,omitempty"`
	MemberDisplayName *string `json:"MemberDisplayName,omitempty" xml:"MemberDisplayName,omitempty"`
	MemberStatus      *string `json:"MemberStatus,omitempty" xml:"MemberStatus,omitempty"`
	MemberUid         *int64  `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	ModifyTime        *int32  `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
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
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInstanceMembersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DescribeInvadeEventListRequest struct {
	// 资产IP
	AssetsIP *string `json:"AssetsIP,omitempty" xml:"AssetsIP,omitempty"`
	// 实例ID
	AssetsInstanceId *string `json:"AssetsInstanceId,omitempty" xml:"AssetsInstanceId,omitempty"`
	// 实例名称
	AssetsInstanceName *string `json:"AssetsInstanceName,omitempty" xml:"AssetsInstanceName,omitempty"`
	// 当前页
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 结束时间
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 事件key
	EventKey *string `json:"EventKey,omitempty" xml:"EventKey,omitempty"`
	// 事件名称
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// 事件UUID
	EventUuid *string `json:"EventUuid,omitempty" xml:"EventUuid,omitempty"`
	// 是否忽略
	IsIgnore *string `json:"IsIgnore,omitempty" xml:"IsIgnore,omitempty"`
	// 语言
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// 成员账号UID
	MemberUid *int64 `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// 每页多少条
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 处理状态
	ProcessStatus *string `json:"ProcessStatus,omitempty" xml:"ProcessStatus,omitempty"`
	// 处理状态列表
	ProcessStatusList []*int32 `json:"ProcessStatusList,omitempty" xml:"ProcessStatusList,omitempty" type:"Repeated"`
	// 风险等级
	RiskLevel []*int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty" type:"Repeated"`
	// 源IP
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// 开始时间
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

func (s *DescribeInvadeEventListRequest) SetProcessStatus(v string) *DescribeInvadeEventListRequest {
	s.ProcessStatus = &v
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
	// 事件列表
	EventList []*DescribeInvadeEventListResponseBodyEventList `json:"EventList,omitempty" xml:"EventList,omitempty" type:"Repeated"`
	// 高风险等级比例
	HighLevelPercent *int32 `json:"HighLevelPercent,omitempty" xml:"HighLevelPercent,omitempty"`
	// 低风险等级比例
	LowLevelPercent *int32 `json:"LowLevelPercent,omitempty" xml:"LowLevelPercent,omitempty"`
	// 中风险等级比例
	MiddleLevelPercent *int32 `json:"MiddleLevelPercent,omitempty" xml:"MiddleLevelPercent,omitempty"`
	// 分页信息
	PageInfo  *DescribeInvadeEventListResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// 资产实例ID
	AssetsInstanceId *string `json:"AssetsInstanceId,omitempty" xml:"AssetsInstanceId,omitempty"`
	// 资产名称
	AssetsInstanceName *string `json:"AssetsInstanceName,omitempty" xml:"AssetsInstanceName,omitempty"`
	// 资产类型
	AssetsType *string `json:"AssetsType,omitempty" xml:"AssetsType,omitempty"`
	// CVE编号
	EventKey *string `json:"EventKey,omitempty" xml:"EventKey,omitempty"`
	// 事件名称
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// 事件来源
	EventSrc *string `json:"EventSrc,omitempty" xml:"EventSrc,omitempty"`
	// 事件UUID
	EventUuid *string `json:"EventUuid,omitempty" xml:"EventUuid,omitempty"`
	// 首次出现时间
	FirstTime *int32 `json:"FirstTime,omitempty" xml:"FirstTime,omitempty"`
	// 是否忽略
	IsIgnore *bool `json:"IsIgnore,omitempty" xml:"IsIgnore,omitempty"`
	// 最近一次时间
	LastTime *int32 `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	// 成员账号UID
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// 私网IP
	PrivateIP *string `json:"PrivateIP,omitempty" xml:"PrivateIP,omitempty"`
	// 处理状态
	ProcessStatus *int32 `json:"ProcessStatus,omitempty" xml:"ProcessStatus,omitempty"`
	// 公网IP
	PublicIP *string `json:"PublicIP,omitempty" xml:"PublicIP,omitempty"`
	// 公开类型
	PublicIpType *string `json:"PublicIpType,omitempty" xml:"PublicIpType,omitempty"`
	// 风险等级
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
	// 当前页
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 每页大小
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 总数
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInvadeEventListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DescribeOutgoingDestinationIPRequest struct {
	// ACL覆盖情况, 枚举值.
	// 默认值: 空
	// 可选值:
	// All (全部情况, 等同于空)
	// FullCoverage ( 已覆盖)
	// Uncovered (未覆盖)
	AclCoverage *string `json:"AclCoverage,omitempty" xml:"AclCoverage,omitempty"`
	// 应用名
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// 分类, 枚举值.
	// 默认值: 空
	// 可选值:
	// All (全部分类)
	// RiskDomain (风险域名分类)
	// RiskIP (风险IP分类)
	// AliYun (云产品分类)
	// NotAliYun (非云产品分类)
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// 当前页
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 目的IP
	DstIP *string `json:"DstIP,omitempty" xml:"DstIP,omitempty"`
	// 结束时间
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 语言
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// 排序字段, 枚举值.
	// 默认值: SessionCount
	// 可选值: InBytes, OutBytes,TotalBytes,SessionCount
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// 每页大小
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 端口号
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// 私网IP
	PrivateIP *string `json:"PrivateIP,omitempty" xml:"PrivateIP,omitempty"`
	// 公网IP
	PublicIP *string `json:"PublicIP,omitempty" xml:"PublicIP,omitempty"`
	// 安全建议, 枚举值: pass, alert, drop. 默认值为空
	SecuritySuggest *string `json:"SecuritySuggest,omitempty" xml:"SecuritySuggest,omitempty"`
	// 顺序, 枚举值, 可选:asc, desc
	Sort     *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// 开始时间,Unix timestamp, 精确到秒
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeOutgoingDestinationIPRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOutgoingDestinationIPRequest) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDestinationIPRequest) SetAclCoverage(v string) *DescribeOutgoingDestinationIPRequest {
	s.AclCoverage = &v
	return s
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

func (s *DescribeOutgoingDestinationIPRequest) SetSecuritySuggest(v string) *DescribeOutgoingDestinationIPRequest {
	s.SecuritySuggest = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetSort(v string) *DescribeOutgoingDestinationIPRequest {
	s.Sort = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetSourceIp(v string) *DescribeOutgoingDestinationIPRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeOutgoingDestinationIPRequest) SetStartTime(v string) *DescribeOutgoingDestinationIPRequest {
	s.StartTime = &v
	return s
}

type DescribeOutgoingDestinationIPResponseBody struct {
	// 外联IP列表
	DstIPList []*DescribeOutgoingDestinationIPResponseBodyDstIPList `json:"DstIPList,omitempty" xml:"DstIPList,omitempty" type:"Repeated"`
	RequestId *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 总数
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
	// ACL覆盖
	AclCoverage *string `json:"AclCoverage,omitempty" xml:"AclCoverage,omitempty"`
	// ACL推荐内容
	AclRecommendDetail *string `json:"AclRecommendDetail,omitempty" xml:"AclRecommendDetail,omitempty"`
	// ACL状态
	AclStatus *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	// 地址簿名称
	AddressGroupName *string `json:"AddressGroupName,omitempty" xml:"AddressGroupName,omitempty"`
	// 地址簿UUID
	AddressGroupUUID    *string                                                                  `json:"AddressGroupUUID,omitempty" xml:"AddressGroupUUID,omitempty"`
	ApplicationPortList []*DescribeOutgoingDestinationIPResponseBodyDstIPListApplicationPortList `json:"ApplicationPortList,omitempty" xml:"ApplicationPortList,omitempty" type:"Repeated"`
	// 分类ID
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// 分类名称
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// 目的IP
	DstIP *string `json:"DstIP,omitempty" xml:"DstIP,omitempty"`
	// 规则中的组名称
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// 是否有推荐ACL
	HasAclRecommend *bool `json:"HasAclRecommend,omitempty" xml:"HasAclRecommend,omitempty"`
	// 入方向流量
	InBytes *int64 `json:"InBytes,omitempty" xml:"InBytes,omitempty"`
	// 是否正常
	IsMarkNormal *bool `json:"IsMarkNormal,omitempty" xml:"IsMarkNormal,omitempty"`
	// 出流量
	OutBytes *int64 `json:"OutBytes,omitempty" xml:"OutBytes,omitempty"`
	// 规则UUID
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// 规则名称
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// 安全建议
	SecuritySuggest *string `json:"SecuritySuggest,omitempty" xml:"SecuritySuggest,omitempty"`
	// 会话数
	SessionCount *int64 `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
	// 标签列表
	TagList []*DescribeOutgoingDestinationIPResponseBodyDstIPListTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
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

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetAddressGroupName(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.AddressGroupName = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetAddressGroupUUID(v string) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.AddressGroupUUID = &v
	return s
}

func (s *DescribeOutgoingDestinationIPResponseBodyDstIPList) SetApplicationPortList(v []*DescribeOutgoingDestinationIPResponseBodyDstIPListApplicationPortList) *DescribeOutgoingDestinationIPResponseBodyDstIPList {
	s.ApplicationPortList = v
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

type DescribeOutgoingDestinationIPResponseBodyDstIPListApplicationPortList struct {
	// 应用名
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// 端口
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
	// 风险等级
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// 标签描述
	TagDescribe *string `json:"TagDescribe,omitempty" xml:"TagDescribe,omitempty"`
	// 标签ID
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// 标签名
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s DescribeOutgoingDestinationIPResponseBodyDstIPListTagList) String() string {
	return tea.Prettify(s)
}

func (s DescribeOutgoingDestinationIPResponseBodyDstIPListTagList) GoString() string {
	return s.String()
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeOutgoingDestinationIPResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// ACL覆盖情况, 枚举值.
	// 默认值: 空
	// 可选值:
	// All (全部情况, 等同于空)
	// FullCoverage ( 已覆盖)
	// Uncovered (未覆盖)
	AclCoverage *string `json:"AclCoverage,omitempty" xml:"AclCoverage,omitempty"`
	// 分类, 枚举值.
	// 默认值: 空
	// 可选值:
	// All (全部分类)
	// RiskDomain (风险域名分类)
	// RiskIP (风险IP分类)
	// AliYun (云产品分类)
	// NotAliYun (非云产品分类)
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// 当前页
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 域名
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// 结束时间,Unix timestamp, 精确到秒
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 语言, 枚举值.
	// 默认值: zh
	// 可选值: en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// 排序字段, 枚举值.
	// 默认值: SessionCount
	// 可选值: InBytes, OutBytes,TotalBytes,SessionCount
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// 每页条数, 不得超过100, 超过100会设置为100
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 公网IP
	PublicIP *string `json:"PublicIP,omitempty" xml:"PublicIP,omitempty"`
	// 安全建议, 枚举值: pass, alert, drop. 默认值为空
	SecuritySuggest *string `json:"SecuritySuggest,omitempty" xml:"SecuritySuggest,omitempty"`
	// 顺序, 枚举值, 可选:asc, desc
	Sort     *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// 开始时间,Unix timestamp, 精确到秒
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeOutgoingDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOutgoingDomainRequest) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDomainRequest) SetAclCoverage(v string) *DescribeOutgoingDomainRequest {
	s.AclCoverage = &v
	return s
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

func (s *DescribeOutgoingDomainRequest) SetSecuritySuggest(v string) *DescribeOutgoingDomainRequest {
	s.SecuritySuggest = &v
	return s
}

func (s *DescribeOutgoingDomainRequest) SetSort(v string) *DescribeOutgoingDomainRequest {
	s.Sort = &v
	return s
}

func (s *DescribeOutgoingDomainRequest) SetSourceIp(v string) *DescribeOutgoingDomainRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeOutgoingDomainRequest) SetStartTime(v string) *DescribeOutgoingDomainRequest {
	s.StartTime = &v
	return s
}

type DescribeOutgoingDomainResponseBody struct {
	DomainList []*DescribeOutgoingDomainResponseBodyDomainList `json:"DomainList,omitempty" xml:"DomainList,omitempty" type:"Repeated"`
	RequestId  *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 总数
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
	// ACL覆盖
	AclCoverage *string `json:"AclCoverage,omitempty" xml:"AclCoverage,omitempty"`
	// ACL推荐内容
	AclRecommendDetail *string `json:"AclRecommendDetail,omitempty" xml:"AclRecommendDetail,omitempty"`
	// ACL状态
	AclStatus *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	// 地址簿名称
	AddressGroupName *string `json:"AddressGroupName,omitempty" xml:"AddressGroupName,omitempty"`
	// 地址簿UUID
	AddressGroupUUID *string `json:"AddressGroupUUID,omitempty" xml:"AddressGroupUUID,omitempty"`
	// 分类ID
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// 分类名称
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// 域名
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// 规则中的组名称
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// 是否有ACL推荐规则
	HasAclRecommend *bool `json:"HasAclRecommend,omitempty" xml:"HasAclRecommend,omitempty"`
	// 入流量
	InBytes *int64 `json:"InBytes,omitempty" xml:"InBytes,omitempty"`
	// 是否正常
	IsMarkNormal *bool `json:"IsMarkNormal,omitempty" xml:"IsMarkNormal,omitempty"`
	// 出流量
	OutBytes *int64 `json:"OutBytes,omitempty" xml:"OutBytes,omitempty"`
	// ACL规则ID
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// 规则名称
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// 安全建议
	SecuritySuggest *string `json:"SecuritySuggest,omitempty" xml:"SecuritySuggest,omitempty"`
	// 会话数
	SessionCount *int64                                                 `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
	TagList      []*DescribeOutgoingDomainResponseBodyDomainListTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
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

type DescribeOutgoingDomainResponseBodyDomainListTagList struct {
	// 风险等级
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// 标签描述
	TagDescribe *string `json:"TagDescribe,omitempty" xml:"TagDescribe,omitempty"`
	// 标签ID
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// 标签名称
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s DescribeOutgoingDomainResponseBodyDomainListTagList) String() string {
	return tea.Prettify(s)
}

func (s DescribeOutgoingDomainResponseBodyDomainListTagList) GoString() string {
	return s.String()
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeOutgoingDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
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
	InternetSwitch *string `json:"InternetSwitch,omitempty" xml:"InternetSwitch,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePolicyAdvancedConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
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
	End       *int32  `json:"End,omitempty" xml:"End,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Start     *int32  `json:"Start,omitempty" xml:"Start,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePolicyPriorUsedResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DescribeRiskEventGroupRequest struct {
	AttackApp            []*string `json:"AttackApp,omitempty" xml:"AttackApp,omitempty" type:"Repeated"`
	AttackType           *string   `json:"AttackType,omitempty" xml:"AttackType,omitempty"`
	BuyVersion           *int64    `json:"BuyVersion,omitempty" xml:"BuyVersion,omitempty"`
	CurrentPage          *string   `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DataType             *string   `json:"DataType,omitempty" xml:"DataType,omitempty"`
	Direction            *string   `json:"Direction,omitempty" xml:"Direction,omitempty"`
	DstIP                *string   `json:"DstIP,omitempty" xml:"DstIP,omitempty"`
	DstNetworkInstanceId *string   `json:"DstNetworkInstanceId,omitempty" xml:"DstNetworkInstanceId,omitempty"`
	EndTime              *string   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FirewallType         *string   `json:"FirewallType,omitempty" xml:"FirewallType,omitempty"`
	Lang                 *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	NoLocation           *string   `json:"NoLocation,omitempty" xml:"NoLocation,omitempty"`
	PageSize             *string   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RuleResult           *string   `json:"RuleResult,omitempty" xml:"RuleResult,omitempty"`
	RuleSource           *string   `json:"RuleSource,omitempty" xml:"RuleSource,omitempty"`
	SrcIP                *string   `json:"SrcIP,omitempty" xml:"SrcIP,omitempty"`
	SrcNetworkInstanceId *string   `json:"SrcNetworkInstanceId,omitempty" xml:"SrcNetworkInstanceId,omitempty"`
	StartTime            *string   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	VulLevel             *string   `json:"VulLevel,omitempty" xml:"VulLevel,omitempty"`
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
	DataList   []*DescribeRiskEventGroupResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	RequestId  *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	AttackApp             *string                                                            `json:"AttackApp,omitempty" xml:"AttackApp,omitempty"`
	AttackType            *int32                                                             `json:"AttackType,omitempty" xml:"AttackType,omitempty"`
	Description           *string                                                            `json:"Description,omitempty" xml:"Description,omitempty"`
	Direction             *string                                                            `json:"Direction,omitempty" xml:"Direction,omitempty"`
	DstIP                 *string                                                            `json:"DstIP,omitempty" xml:"DstIP,omitempty"`
	EventCount            *int32                                                             `json:"EventCount,omitempty" xml:"EventCount,omitempty"`
	EventId               *string                                                            `json:"EventId,omitempty" xml:"EventId,omitempty"`
	EventName             *string                                                            `json:"EventName,omitempty" xml:"EventName,omitempty"`
	FirstEventTime        *int32                                                             `json:"FirstEventTime,omitempty" xml:"FirstEventTime,omitempty"`
	IPLocationInfo        *DescribeRiskEventGroupResponseBodyDataListIPLocationInfo          `json:"IPLocationInfo,omitempty" xml:"IPLocationInfo,omitempty" type:"Struct"`
	LastEventTime         *int32                                                             `json:"LastEventTime,omitempty" xml:"LastEventTime,omitempty"`
	ResourcePrivateIPList []*DescribeRiskEventGroupResponseBodyDataListResourcePrivateIPList `json:"ResourcePrivateIPList,omitempty" xml:"ResourcePrivateIPList,omitempty" type:"Repeated"`
	ResourceType          *string                                                            `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	RuleId                *string                                                            `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleResult            *int32                                                             `json:"RuleResult,omitempty" xml:"RuleResult,omitempty"`
	RuleSource            *int32                                                             `json:"RuleSource,omitempty" xml:"RuleSource,omitempty"`
	SrcIP                 *string                                                            `json:"SrcIP,omitempty" xml:"SrcIP,omitempty"`
	SrcPrivateIPList      []*string                                                          `json:"SrcPrivateIPList,omitempty" xml:"SrcPrivateIPList,omitempty" type:"Repeated"`
	Tag                   *string                                                            `json:"Tag,omitempty" xml:"Tag,omitempty"`
	VpcDstInfo            *DescribeRiskEventGroupResponseBodyDataListVpcDstInfo              `json:"VpcDstInfo,omitempty" xml:"VpcDstInfo,omitempty" type:"Struct"`
	VpcSrcInfo            *DescribeRiskEventGroupResponseBodyDataListVpcSrcInfo              `json:"VpcSrcInfo,omitempty" xml:"VpcSrcInfo,omitempty" type:"Struct"`
	VulLevel              *int32                                                             `json:"VulLevel,omitempty" xml:"VulLevel,omitempty"`
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
	CityId      *string `json:"CityId,omitempty" xml:"CityId,omitempty"`
	CityName    *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	CountryId   *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
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
	RegionNo             *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	ResourceInstanceId   *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
	ResourceInstanceName *string `json:"ResourceInstanceName,omitempty" xml:"ResourceInstanceName,omitempty"`
	ResourcePrivateIP    *string `json:"ResourcePrivateIP,omitempty" xml:"ResourcePrivateIP,omitempty"`
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
	EcsInstanceId       *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	EcsInstanceName     *string `json:"EcsInstanceName,omitempty" xml:"EcsInstanceName,omitempty"`
	NetworkInstanceId   *string `json:"NetworkInstanceId,omitempty" xml:"NetworkInstanceId,omitempty"`
	NetworkInstanceName *string `json:"NetworkInstanceName,omitempty" xml:"NetworkInstanceName,omitempty"`
	RegionNo            *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
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
	EcsInstanceId       *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	EcsInstanceName     *string `json:"EcsInstanceName,omitempty" xml:"EcsInstanceName,omitempty"`
	NetworkInstanceId   *string `json:"NetworkInstanceId,omitempty" xml:"NetworkInstanceId,omitempty"`
	NetworkInstanceName *string `json:"NetworkInstanceName,omitempty" xml:"NetworkInstanceName,omitempty"`
	RegionNo            *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRiskEventGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DescribeUserAssetIPTrafficInfoRequest struct {
	// 资产IP
	AssetIP *string `json:"AssetIP,omitempty" xml:"AssetIP,omitempty"`
	// 语言
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// 时间
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
	// 结束时间
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 入方向流量信息
	InTrafficInfo *DescribeUserAssetIPTrafficInfoResponseBodyInTrafficInfo `json:"InTrafficInfo,omitempty" xml:"InTrafficInfo,omitempty" type:"Struct"`
	// 出方向流量信息
	OutTrafficInfo *DescribeUserAssetIPTrafficInfoResponseBodyOutTrafficInfo `json:"OutTrafficInfo,omitempty" xml:"OutTrafficInfo,omitempty" type:"Struct"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 开始时间
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

func (s *DescribeUserAssetIPTrafficInfoResponseBody) SetInTrafficInfo(v *DescribeUserAssetIPTrafficInfoResponseBodyInTrafficInfo) *DescribeUserAssetIPTrafficInfoResponseBody {
	s.InTrafficInfo = v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponseBody) SetOutTrafficInfo(v *DescribeUserAssetIPTrafficInfoResponseBodyOutTrafficInfo) *DescribeUserAssetIPTrafficInfoResponseBody {
	s.OutTrafficInfo = v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponseBody) SetRequestId(v string) *DescribeUserAssetIPTrafficInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponseBody) SetStartTime(v int64) *DescribeUserAssetIPTrafficInfoResponseBody {
	s.StartTime = &v
	return s
}

type DescribeUserAssetIPTrafficInfoResponseBodyInTrafficInfo struct {
	// 入方向Bps
	InBps *int64 `json:"InBps,omitempty" xml:"InBps,omitempty"`
	// 入方向pps
	InPps *int64 `json:"InPps,omitempty" xml:"InPps,omitempty"`
	// 新建会话数
	NewConn *int64 `json:"NewConn,omitempty" xml:"NewConn,omitempty"`
	// 返回Bps
	OutBps *int64 `json:"OutBps,omitempty" xml:"OutBps,omitempty"`
	// 返回pps
	OutPps *int64 `json:"OutPps,omitempty" xml:"OutPps,omitempty"`
	// 会话数
	SessionCount *int64 `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
}

func (s DescribeUserAssetIPTrafficInfoResponseBodyInTrafficInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserAssetIPTrafficInfoResponseBodyInTrafficInfo) GoString() string {
	return s.String()
}

func (s *DescribeUserAssetIPTrafficInfoResponseBodyInTrafficInfo) SetInBps(v int64) *DescribeUserAssetIPTrafficInfoResponseBodyInTrafficInfo {
	s.InBps = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponseBodyInTrafficInfo) SetInPps(v int64) *DescribeUserAssetIPTrafficInfoResponseBodyInTrafficInfo {
	s.InPps = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponseBodyInTrafficInfo) SetNewConn(v int64) *DescribeUserAssetIPTrafficInfoResponseBodyInTrafficInfo {
	s.NewConn = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponseBodyInTrafficInfo) SetOutBps(v int64) *DescribeUserAssetIPTrafficInfoResponseBodyInTrafficInfo {
	s.OutBps = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponseBodyInTrafficInfo) SetOutPps(v int64) *DescribeUserAssetIPTrafficInfoResponseBodyInTrafficInfo {
	s.OutPps = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponseBodyInTrafficInfo) SetSessionCount(v int64) *DescribeUserAssetIPTrafficInfoResponseBodyInTrafficInfo {
	s.SessionCount = &v
	return s
}

type DescribeUserAssetIPTrafficInfoResponseBodyOutTrafficInfo struct {
	// 出方向接收流量Bps
	InBps *int64 `json:"InBps,omitempty" xml:"InBps,omitempty"`
	// 出方向接收流量Bps
	InPps *int64 `json:"InPps,omitempty" xml:"InPps,omitempty"`
	// 新建会话数
	NewConn *int64 `json:"NewConn,omitempty" xml:"NewConn,omitempty"`
	// 出方向流量Bps
	OutBps *int64 `json:"OutBps,omitempty" xml:"OutBps,omitempty"`
	// 出方向pps
	OutPps *int64 `json:"OutPps,omitempty" xml:"OutPps,omitempty"`
	// 会话数
	SessionCount *int64 `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
}

func (s DescribeUserAssetIPTrafficInfoResponseBodyOutTrafficInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserAssetIPTrafficInfoResponseBodyOutTrafficInfo) GoString() string {
	return s.String()
}

func (s *DescribeUserAssetIPTrafficInfoResponseBodyOutTrafficInfo) SetInBps(v int64) *DescribeUserAssetIPTrafficInfoResponseBodyOutTrafficInfo {
	s.InBps = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponseBodyOutTrafficInfo) SetInPps(v int64) *DescribeUserAssetIPTrafficInfoResponseBodyOutTrafficInfo {
	s.InPps = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponseBodyOutTrafficInfo) SetNewConn(v int64) *DescribeUserAssetIPTrafficInfoResponseBodyOutTrafficInfo {
	s.NewConn = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponseBodyOutTrafficInfo) SetOutBps(v int64) *DescribeUserAssetIPTrafficInfoResponseBodyOutTrafficInfo {
	s.OutBps = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponseBodyOutTrafficInfo) SetOutPps(v int64) *DescribeUserAssetIPTrafficInfoResponseBodyOutTrafficInfo {
	s.OutPps = &v
	return s
}

func (s *DescribeUserAssetIPTrafficInfoResponseBodyOutTrafficInfo) SetSessionCount(v int64) *DescribeUserAssetIPTrafficInfoResponseBodyOutTrafficInfo {
	s.SessionCount = &v
	return s
}

type DescribeUserAssetIPTrafficInfoResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeUserAssetIPTrafficInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DescribeVpcFirewallAclGroupListRequest struct {
	CurrentPage             *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	FirewallConfigureStatus *string `json:"FirewallConfigureStatus,omitempty" xml:"FirewallConfigureStatus,omitempty"`
	Lang                    *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageSize                *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	AclGroupList []*DescribeVpcFirewallAclGroupListResponseBodyAclGroupList `json:"AclGroupList,omitempty" xml:"AclGroupList,omitempty" type:"Repeated"`
	RequestId    *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int32                                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	AclGroupId   *string `json:"AclGroupId,omitempty" xml:"AclGroupId,omitempty"`
	AclGroupName *string `json:"AclGroupName,omitempty" xml:"AclGroupName,omitempty"`
	MemberUid    *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
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

func (s *DescribeVpcFirewallAclGroupListResponseBodyAclGroupList) SetMemberUid(v string) *DescribeVpcFirewallAclGroupListResponseBodyAclGroupList {
	s.MemberUid = &v
	return s
}

type DescribeVpcFirewallAclGroupListResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVpcFirewallAclGroupListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Lang              *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	NetworkInstanceId *string `json:"NetworkInstanceId,omitempty" xml:"NetworkInstanceId,omitempty"`
	VpcFirewallId     *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
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
	ConnectType          *string                                           `json:"ConnectType,omitempty" xml:"ConnectType,omitempty"`
	FirewallSwitchStatus *string                                           `json:"FirewallSwitchStatus,omitempty" xml:"FirewallSwitchStatus,omitempty"`
	LocalVpc             *DescribeVpcFirewallCenDetailResponseBodyLocalVpc `json:"LocalVpc,omitempty" xml:"LocalVpc,omitempty" type:"Struct"`
	RequestId            *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VpcFirewallId        *string                                           `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
	VpcFirewallName      *string                                           `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
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

type DescribeVpcFirewallCenDetailResponseBodyLocalVpc struct {
	AttachmentId        *string                                                             `json:"AttachmentId,omitempty" xml:"AttachmentId,omitempty"`
	AttachmentName      *string                                                             `json:"AttachmentName,omitempty" xml:"AttachmentName,omitempty"`
	DefendCidrList      []*string                                                           `json:"DefendCidrList,omitempty" xml:"DefendCidrList,omitempty" type:"Repeated"`
	EniList             []*DescribeVpcFirewallCenDetailResponseBodyLocalVpcEniList          `json:"EniList,omitempty" xml:"EniList,omitempty" type:"Repeated"`
	ManualVSwitchId     *string                                                             `json:"ManualVSwitchId,omitempty" xml:"ManualVSwitchId,omitempty"`
	NetworkInstanceId   *string                                                             `json:"NetworkInstanceId,omitempty" xml:"NetworkInstanceId,omitempty"`
	NetworkInstanceName *string                                                             `json:"NetworkInstanceName,omitempty" xml:"NetworkInstanceName,omitempty"`
	NetworkInstanceType *string                                                             `json:"NetworkInstanceType,omitempty" xml:"NetworkInstanceType,omitempty"`
	OwnerId             *string                                                             `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionNo            *string                                                             `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	RouteMode           *string                                                             `json:"RouteMode,omitempty" xml:"RouteMode,omitempty"`
	SupportManualMode   *string                                                             `json:"SupportManualMode,omitempty" xml:"SupportManualMode,omitempty"`
	TransitRouterId     *string                                                             `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	TransitRouterType   *string                                                             `json:"TransitRouterType,omitempty" xml:"TransitRouterType,omitempty"`
	VpcCidrTableList    []*DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableList `json:"VpcCidrTableList,omitempty" xml:"VpcCidrTableList,omitempty" type:"Repeated"`
	VpcId               *string                                                             `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcName             *string                                                             `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
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
	EniId               *string `json:"EniId,omitempty" xml:"EniId,omitempty"`
	EniPrivateIpAddress *string `json:"EniPrivateIpAddress,omitempty" xml:"EniPrivateIpAddress,omitempty"`
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

type DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableList struct {
	RouteEntryList []*DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList `json:"RouteEntryList,omitempty" xml:"RouteEntryList,omitempty" type:"Repeated"`
	RouteTableId   *string                                                                           `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
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
	DestinationCidr   *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVpcFirewallCenDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	CurrentPage          *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	FirewallSwitchStatus *string `json:"FirewallSwitchStatus,omitempty" xml:"FirewallSwitchStatus,omitempty"`
	Lang                 *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	MemberUid            *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	NetworkInstanceId    *string `json:"NetworkInstanceId,omitempty" xml:"NetworkInstanceId,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageSize             *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionNo             *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	RouteMode            *string `json:"RouteMode,omitempty" xml:"RouteMode,omitempty"`
	VpcFirewallId        *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
	VpcFirewallName      *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
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

func (s *DescribeVpcFirewallCenListRequest) SetVpcFirewallId(v string) *DescribeVpcFirewallCenListRequest {
	s.VpcFirewallId = &v
	return s
}

func (s *DescribeVpcFirewallCenListRequest) SetVpcFirewallName(v string) *DescribeVpcFirewallCenListRequest {
	s.VpcFirewallName = &v
	return s
}

type DescribeVpcFirewallCenListResponseBody struct {
	RequestId    *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int32                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	CenId                *string                                                      `json:"CenId,omitempty" xml:"CenId,omitempty"`
	CenName              *string                                                      `json:"CenName,omitempty" xml:"CenName,omitempty"`
	ConnectType          *string                                                      `json:"ConnectType,omitempty" xml:"ConnectType,omitempty"`
	FirewallSwitchStatus *string                                                      `json:"FirewallSwitchStatus,omitempty" xml:"FirewallSwitchStatus,omitempty"`
	IpsConfig            *DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig `json:"IpsConfig,omitempty" xml:"IpsConfig,omitempty" type:"Struct"`
	LocalVpc             *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc  `json:"LocalVpc,omitempty" xml:"LocalVpc,omitempty" type:"Struct"`
	MemberUid            *string                                                      `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	PrecheckStatus       *string                                                      `json:"PrecheckStatus,omitempty" xml:"PrecheckStatus,omitempty"`
	RegionStatus         *string                                                      `json:"RegionStatus,omitempty" xml:"RegionStatus,omitempty"`
	ResultCode           *string                                                      `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	VpcFirewallId        *string                                                      `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
	VpcFirewallName      *string                                                      `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
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
	BasicRules     *int32 `json:"BasicRules,omitempty" xml:"BasicRules,omitempty"`
	EnableAllPatch *int32 `json:"EnableAllPatch,omitempty" xml:"EnableAllPatch,omitempty"`
	RunMode        *int32 `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
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
	AuthorizationStatus *string                                                                       `json:"AuthorizationStatus,omitempty" xml:"AuthorizationStatus,omitempty"`
	DefendCidrList      []*string                                                                     `json:"DefendCidrList,omitempty" xml:"DefendCidrList,omitempty" type:"Repeated"`
	ManualVSwitchId     *string                                                                       `json:"ManualVSwitchId,omitempty" xml:"ManualVSwitchId,omitempty"`
	NetworkInstanceId   *string                                                                       `json:"NetworkInstanceId,omitempty" xml:"NetworkInstanceId,omitempty"`
	NetworkInstanceName *string                                                                       `json:"NetworkInstanceName,omitempty" xml:"NetworkInstanceName,omitempty"`
	NetworkInstanceType *string                                                                       `json:"NetworkInstanceType,omitempty" xml:"NetworkInstanceType,omitempty"`
	OwnerId             *int64                                                                        `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionNo            *string                                                                       `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	RouteMode           *string                                                                       `json:"RouteMode,omitempty" xml:"RouteMode,omitempty"`
	SupportManualMode   *string                                                                       `json:"SupportManualMode,omitempty" xml:"SupportManualMode,omitempty"`
	TransitRouterType   *string                                                                       `json:"TransitRouterType,omitempty" xml:"TransitRouterType,omitempty"`
	VpcCidrTableList    []*DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList `json:"VpcCidrTableList,omitempty" xml:"VpcCidrTableList,omitempty" type:"Repeated"`
	VpcId               *string                                                                       `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcName             *string                                                                       `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
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
	RouteEntryList []*DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList `json:"RouteEntryList,omitempty" xml:"RouteEntryList,omitempty" type:"Repeated"`
	RouteTableId   *string                                                                                     `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
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
	DestinationCidr   *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVpcFirewallCenListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AclAction     *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	AclUuid       *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	CurrentPage   *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Destination   *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	Lang          *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	MemberUid     *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	PageSize      *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Proto         *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	Release       *string `json:"Release,omitempty" xml:"Release,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
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

func (s *DescribeVpcFirewallControlPolicyRequest) SetSource(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.Source = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetVpcFirewallId(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.VpcFirewallId = &v
	return s
}

type DescribeVpcFirewallControlPolicyResponseBody struct {
	Policys    []*DescribeVpcFirewallControlPolicyResponseBodyPolicys `json:"Policys,omitempty" xml:"Policys,omitempty" type:"Repeated"`
	RequestId  *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *string                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	AclAction             *string   `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	AclUuid               *string   `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	ApplicationId         *string   `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	ApplicationName       *string   `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	Description           *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	DestPort              *string   `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	DestPortGroup         *string   `json:"DestPortGroup,omitempty" xml:"DestPortGroup,omitempty"`
	DestPortGroupPorts    []*string `json:"DestPortGroupPorts,omitempty" xml:"DestPortGroupPorts,omitempty" type:"Repeated"`
	DestPortType          *string   `json:"DestPortType,omitempty" xml:"DestPortType,omitempty"`
	Destination           *string   `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DestinationGroupCidrs []*string `json:"DestinationGroupCidrs,omitempty" xml:"DestinationGroupCidrs,omitempty" type:"Repeated"`
	DestinationType       *string   `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	HitTimes              *int32    `json:"HitTimes,omitempty" xml:"HitTimes,omitempty"`
	MemberUid             *string   `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	Order                 *int32    `json:"Order,omitempty" xml:"Order,omitempty"`
	Proto                 *string   `json:"Proto,omitempty" xml:"Proto,omitempty"`
	Release               *string   `json:"Release,omitempty" xml:"Release,omitempty"`
	Source                *string   `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceGroupCidrs      []*string `json:"SourceGroupCidrs,omitempty" xml:"SourceGroupCidrs,omitempty" type:"Repeated"`
	SourceType            *string   `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
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

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetDestinationType(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.DestinationType = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetHitTimes(v int32) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.HitTimes = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetMemberUid(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.MemberUid = &v
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

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetSource(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.Source = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetSourceGroupCidrs(v []*string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.SourceGroupCidrs = v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetSourceType(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.SourceType = &v
	return s
}

type DescribeVpcFirewallControlPolicyResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVpcFirewallControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	MemberUid     *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
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
	BasicRules     *int32  `json:"BasicRules,omitempty" xml:"BasicRules,omitempty"`
	EnableAllPatch *int32  `json:"EnableAllPatch,omitempty" xml:"EnableAllPatch,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RunMode        *int32  `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
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
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVpcFirewallDefaultIPSConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Lang          *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	LocalVpcId    *string `json:"LocalVpcId,omitempty" xml:"LocalVpcId,omitempty"`
	PeerVpcId     *string `json:"PeerVpcId,omitempty" xml:"PeerVpcId,omitempty"`
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
	Bandwidth            *int32                                         `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	ConnectType          *string                                        `json:"ConnectType,omitempty" xml:"ConnectType,omitempty"`
	FirewallSwitchStatus *string                                        `json:"FirewallSwitchStatus,omitempty" xml:"FirewallSwitchStatus,omitempty"`
	LocalVpc             *DescribeVpcFirewallDetailResponseBodyLocalVpc `json:"LocalVpc,omitempty" xml:"LocalVpc,omitempty" type:"Struct"`
	PeerVpc              *DescribeVpcFirewallDetailResponseBodyPeerVpc  `json:"PeerVpc,omitempty" xml:"PeerVpc,omitempty" type:"Struct"`
	RequestId            *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VpcFirewallId        *string                                        `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
	VpcFirewallName      *string                                        `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
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
	EniId               *string                                                          `json:"EniId,omitempty" xml:"EniId,omitempty"`
	EniPrivateIpAddress *string                                                          `json:"EniPrivateIpAddress,omitempty" xml:"EniPrivateIpAddress,omitempty"`
	RegionNo            *string                                                          `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	RouterInterfaceId   *string                                                          `json:"RouterInterfaceId,omitempty" xml:"RouterInterfaceId,omitempty"`
	VpcCidrTableList    []*DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableList `json:"VpcCidrTableList,omitempty" xml:"VpcCidrTableList,omitempty" type:"Repeated"`
	VpcId               *string                                                          `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcName             *string                                                          `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
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
	RouteEntryList []*DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList `json:"RouteEntryList,omitempty" xml:"RouteEntryList,omitempty" type:"Repeated"`
	RouteTableId   *string                                                                        `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
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
	DestinationCidr   *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
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
	EniId               *string                                                         `json:"EniId,omitempty" xml:"EniId,omitempty"`
	EniPrivateIpAddress *string                                                         `json:"EniPrivateIpAddress,omitempty" xml:"EniPrivateIpAddress,omitempty"`
	RegionNo            *string                                                         `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	RouterInterfaceId   *string                                                         `json:"RouterInterfaceId,omitempty" xml:"RouterInterfaceId,omitempty"`
	VpcCidrTableList    []*DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableList `json:"VpcCidrTableList,omitempty" xml:"VpcCidrTableList,omitempty" type:"Repeated"`
	VpcId               *string                                                         `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcName             *string                                                         `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
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
	RouteEntryList []*DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableListRouteEntryList `json:"RouteEntryList,omitempty" xml:"RouteEntryList,omitempty" type:"Repeated"`
	RouteTableId   *string                                                                       `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
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
	DestinationCidr   *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVpcFirewallDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DescribeVpcFirewallListRequest struct {
	CurrentPage          *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	FirewallSwitchStatus *string `json:"FirewallSwitchStatus,omitempty" xml:"FirewallSwitchStatus,omitempty"`
	Lang                 *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	MemberUid            *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	PageSize             *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionNo             *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	VpcFirewallId        *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
	VpcFirewallName      *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
	VpcId                *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeVpcFirewallListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallListRequest) GoString() string {
	return s.String()
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
	RequestId    *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int32                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	Bandwidth            *int32                                                    `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	ConnectType          *string                                                   `json:"ConnectType,omitempty" xml:"ConnectType,omitempty"`
	FirewallSwitchStatus *string                                                   `json:"FirewallSwitchStatus,omitempty" xml:"FirewallSwitchStatus,omitempty"`
	IpsConfig            *DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig `json:"IpsConfig,omitempty" xml:"IpsConfig,omitempty" type:"Struct"`
	LocalVpc             *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc  `json:"LocalVpc,omitempty" xml:"LocalVpc,omitempty" type:"Struct"`
	MemberUid            *string                                                   `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	PeerVpc              *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc   `json:"PeerVpc,omitempty" xml:"PeerVpc,omitempty" type:"Struct"`
	RegionStatus         *string                                                   `json:"RegionStatus,omitempty" xml:"RegionStatus,omitempty"`
	VpcFirewallId        *string                                                   `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
	VpcFirewallName      *string                                                   `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
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

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) SetVpcFirewallId(v string) *DescribeVpcFirewallListResponseBodyVpcFirewalls {
	s.VpcFirewallId = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) SetVpcFirewallName(v string) *DescribeVpcFirewallListResponseBodyVpcFirewalls {
	s.VpcFirewallName = &v
	return s
}

type DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig struct {
	BasicRules     *int32 `json:"BasicRules,omitempty" xml:"BasicRules,omitempty"`
	EnableAllPatch *int32 `json:"EnableAllPatch,omitempty" xml:"EnableAllPatch,omitempty"`
	RunMode        *int32 `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
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
	AuthorizationStatus *string                                                                    `json:"AuthorizationStatus,omitempty" xml:"AuthorizationStatus,omitempty"`
	OwnerId             *int64                                                                     `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionNo            *string                                                                    `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	VpcCidrTableList    []*DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList `json:"VpcCidrTableList,omitempty" xml:"VpcCidrTableList,omitempty" type:"Repeated"`
	VpcId               *string                                                                    `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcName             *string                                                                    `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
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
	RouteEntryList []*DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList `json:"RouteEntryList,omitempty" xml:"RouteEntryList,omitempty" type:"Repeated"`
	RouteTableId   *string                                                                                  `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
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
	DestinationCidr   *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
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
	AuthorizationStatus *string                                                                   `json:"AuthorizationStatus,omitempty" xml:"AuthorizationStatus,omitempty"`
	OwnerId             *int64                                                                    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionNo            *string                                                                   `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	VpcCidrTableList    []*DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableList `json:"VpcCidrTableList,omitempty" xml:"VpcCidrTableList,omitempty" type:"Repeated"`
	VpcId               *string                                                                   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcName             *string                                                                   `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
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
	RouteEntryList []*DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableListRouteEntryList `json:"RouteEntryList,omitempty" xml:"RouteEntryList,omitempty" type:"Repeated"`
	RouteTableId   *string                                                                                 `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
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
	DestinationCidr   *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVpcFirewallListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Lang          *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
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
	End       *int32  `json:"End,omitempty" xml:"End,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Start     *int32  `json:"Start,omitempty" xml:"Start,omitempty"`
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
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeVpcFirewallPolicyPriorUsedResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ModifyAddressBookRequest struct {
	AddressList   *string                            `json:"AddressList,omitempty" xml:"AddressList,omitempty"`
	AutoAddTagEcs *string                            `json:"AutoAddTagEcs,omitempty" xml:"AutoAddTagEcs,omitempty"`
	Description   *string                            `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupName     *string                            `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	GroupUuid     *string                            `json:"GroupUuid,omitempty" xml:"GroupUuid,omitempty"`
	Lang          *string                            `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp      *string                            `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	TagList       []*ModifyAddressBookRequestTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	TagRelation   *string                            `json:"TagRelation,omitempty" xml:"TagRelation,omitempty"`
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
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyAddressBookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AclAction           *string   `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	AclUuid             *string   `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	ApplicationName     *string   `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	ApplicationNameList []*string `json:"ApplicationNameList,omitempty" xml:"ApplicationNameList,omitempty" type:"Repeated"`
	Description         *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	DestPort            *string   `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	DestPortGroup       *string   `json:"DestPortGroup,omitempty" xml:"DestPortGroup,omitempty"`
	DestPortType        *string   `json:"DestPortType,omitempty" xml:"DestPortType,omitempty"`
	Destination         *string   `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DestinationType     *string   `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	Direction           *string   `json:"Direction,omitempty" xml:"Direction,omitempty"`
	Lang                *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	MessageType         *string   `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
	Proto               *string   `json:"Proto,omitempty" xml:"Proto,omitempty"`
	Release             *string   `json:"Release,omitempty" xml:"Release,omitempty"`
	Source              *string   `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceIp            *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	SourceType          *string   `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
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

func (s *ModifyControlPolicyRequest) SetLang(v string) *ModifyControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetMessageType(v string) *ModifyControlPolicyRequest {
	s.MessageType = &v
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

func (s *ModifyControlPolicyRequest) SetSource(v string) *ModifyControlPolicyRequest {
	s.Source = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetSourceIp(v string) *ModifyControlPolicyRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetSourceType(v string) *ModifyControlPolicyRequest {
	s.SourceType = &v
	return s
}

type ModifyControlPolicyResponseBody struct {
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	NewOrder  *string `json:"NewOrder,omitempty" xml:"NewOrder,omitempty"`
	OldOrder  *string `json:"OldOrder,omitempty" xml:"OldOrder,omitempty"`
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyControlPolicyPositionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ModifyInstanceMemberAttributesRequest struct {
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
	MemberDesc *string `json:"MemberDesc,omitempty" xml:"MemberDesc,omitempty"`
	MemberUid  *int64  `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
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
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyInstanceMemberAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ModifyPolicyAdvancedConfigRequest struct {
	InternetSwitch *string `json:"InternetSwitch,omitempty" xml:"InternetSwitch,omitempty"`
	Lang           *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp       *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyPolicyAdvancedConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ModifyVpcFirewallCenConfigureRequest struct {
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	MemberUid       *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	VpcFirewallId   *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyVpcFirewallCenConfigureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	FirewallSwitch *string `json:"FirewallSwitch,omitempty" xml:"FirewallSwitch,omitempty"`
	Lang           *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	MemberUid      *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	VpcFirewallId  *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
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
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyVpcFirewallCenSwitchStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Lang                  *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	LocalVpcCidrTableList *string `json:"LocalVpcCidrTableList,omitempty" xml:"LocalVpcCidrTableList,omitempty"`
	MemberUid             *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	PeerVpcCidrTableList  *string `json:"PeerVpcCidrTableList,omitempty" xml:"PeerVpcCidrTableList,omitempty"`
	VpcFirewallId         *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
	VpcFirewallName       *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyVpcFirewallConfigureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AclAction       *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	AclUuid         *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DestPort        *string `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	DestPortGroup   *string `json:"DestPortGroup,omitempty" xml:"DestPortGroup,omitempty"`
	DestPortType    *string `json:"DestPortType,omitempty" xml:"DestPortType,omitempty"`
	Destination     *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Proto           *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	Release         *string `json:"Release,omitempty" xml:"Release,omitempty"`
	Source          *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceType      *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	VpcFirewallId   *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
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

func (s *ModifyVpcFirewallControlPolicyRequest) SetSource(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.Source = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetSourceType(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.SourceType = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetVpcFirewallId(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.VpcFirewallId = &v
	return s
}

type ModifyVpcFirewallControlPolicyResponseBody struct {
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
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyVpcFirewallControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Lang          *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	NewOrder      *string `json:"NewOrder,omitempty" xml:"NewOrder,omitempty"`
	OldOrder      *string `json:"OldOrder,omitempty" xml:"OldOrder,omitempty"`
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s ModifyVpcFirewallControlPolicyPositionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallControlPolicyPositionRequest) GoString() string {
	return s.String()
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
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyVpcFirewallControlPolicyPositionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	BasicRules     *string `json:"BasicRules,omitempty" xml:"BasicRules,omitempty"`
	EnableAllPatch *string `json:"EnableAllPatch,omitempty" xml:"EnableAllPatch,omitempty"`
	Lang           *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	MemberUid      *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	RunMode        *string `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
	SourceIp       *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	VpcFirewallId  *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
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
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyVpcFirewallDefaultIPSConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ModifyVpcFirewallSwitchStatusRequest struct {
	FirewallSwitch *string `json:"FirewallSwitch,omitempty" xml:"FirewallSwitch,omitempty"`
	Lang           *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	MemberUid      *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	VpcFirewallId  *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyVpcFirewallSwitchStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PutDisableAllFwSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	IpaddrList       []*string `json:"IpaddrList,omitempty" xml:"IpaddrList,omitempty" type:"Repeated"`
	Lang             *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RegionList       []*string `json:"RegionList,omitempty" xml:"RegionList,omitempty" type:"Repeated"`
	ResourceTypeList []*string `json:"ResourceTypeList,omitempty" xml:"ResourceTypeList,omitempty" type:"Repeated"`
	SourceIp         *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PutDisableFwSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PutEnableAllFwSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	IpaddrList       []*string `json:"IpaddrList,omitempty" xml:"IpaddrList,omitempty" type:"Repeated"`
	Lang             *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RegionList       []*string `json:"RegionList,omitempty" xml:"RegionList,omitempty" type:"Repeated"`
	ResourceTypeList []*string `json:"ResourceTypeList,omitempty" xml:"ResourceTypeList,omitempty" type:"Repeated"`
	SourceIp         *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutEnableFwSwitchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutEnableFwSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *PutEnableFwSwitchResponseBody) SetRequestId(v string) *PutEnableFwSwitchResponseBody {
	s.RequestId = &v
	return s
}

type PutEnableFwSwitchResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PutEnableFwSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ResetVpcFirewallRuleHitCountRequest struct {
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	Lang    *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResetVpcFirewallRuleHitCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
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

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MemberUid)) {
		query["MemberUid"] = request.MemberUid
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkInstanceId)) {
		query["NetworkInstanceId"] = request.NetworkInstanceId
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

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
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

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
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

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
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

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIp)) {
		query["SourceIp"] = request.SourceIp
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

	if !tea.BoolValue(util.IsUnset(request.ProcessStatus)) {
		query["ProcessStatus"] = request.ProcessStatus
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

func (client *Client) DescribeOutgoingDestinationIPWithOptions(request *DescribeOutgoingDestinationIPRequest, runtime *util.RuntimeOptions) (_result *DescribeOutgoingDestinationIPResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclCoverage)) {
		query["AclCoverage"] = request.AclCoverage
	}

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

	if !tea.BoolValue(util.IsUnset(request.SecuritySuggest)) {
		query["SecuritySuggest"] = request.SecuritySuggest
	}

	if !tea.BoolValue(util.IsUnset(request.Sort)) {
		query["Sort"] = request.Sort
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
	if !tea.BoolValue(util.IsUnset(request.AclCoverage)) {
		query["AclCoverage"] = request.AclCoverage
	}

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

	if !tea.BoolValue(util.IsUnset(request.SecuritySuggest)) {
		query["SecuritySuggest"] = request.SecuritySuggest
	}

	if !tea.BoolValue(util.IsUnset(request.Sort)) {
		query["Sort"] = request.Sort
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

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RuleResult)) {
		query["RuleResult"] = request.RuleResult
	}

	if !tea.BoolValue(util.IsUnset(request.RuleSource)) {
		query["RuleSource"] = request.RuleSource
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

func (client *Client) DescribeVpcFirewallListWithOptions(request *DescribeVpcFirewallListRequest, runtime *util.RuntimeOptions) (_result *DescribeVpcFirewallListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.MessageType)) {
		query["MessageType"] = request.MessageType
	}

	if !tea.BoolValue(util.IsUnset(request.Proto)) {
		query["Proto"] = request.Proto
	}

	if !tea.BoolValue(util.IsUnset(request.Release)) {
		query["Release"] = request.Release
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

	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	if !tea.BoolValue(util.IsUnset(request.Proto)) {
		query["Proto"] = request.Proto
	}

	if !tea.BoolValue(util.IsUnset(request.Release)) {
		query["Release"] = request.Release
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
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

func (client *Client) ModifyVpcFirewallControlPolicyPositionWithOptions(request *ModifyVpcFirewallControlPolicyPositionRequest, runtime *util.RuntimeOptions) (_result *ModifyVpcFirewallControlPolicyPositionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
