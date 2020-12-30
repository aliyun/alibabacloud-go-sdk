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

type AddAddressBookRequest struct {
	SourceIp      *string                         `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang          *string                         `json:"Lang,omitempty" xml:"Lang,omitempty"`
	AddressList   *string                         `json:"AddressList,omitempty" xml:"AddressList,omitempty"`
	Description   *string                         `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupName     *string                         `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	GroupType     *string                         `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	AutoAddTagEcs *string                         `json:"AutoAddTagEcs,omitempty" xml:"AutoAddTagEcs,omitempty"`
	TagRelation   *string                         `json:"TagRelation,omitempty" xml:"TagRelation,omitempty"`
	TagList       []*AddAddressBookRequestTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
}

func (s AddAddressBookRequest) String() string {
	return tea.Prettify(s)
}

func (s AddAddressBookRequest) GoString() string {
	return s.String()
}

func (s *AddAddressBookRequest) SetSourceIp(v string) *AddAddressBookRequest {
	s.SourceIp = &v
	return s
}

func (s *AddAddressBookRequest) SetLang(v string) *AddAddressBookRequest {
	s.Lang = &v
	return s
}

func (s *AddAddressBookRequest) SetAddressList(v string) *AddAddressBookRequest {
	s.AddressList = &v
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

func (s *AddAddressBookRequest) SetAutoAddTagEcs(v string) *AddAddressBookRequest {
	s.AutoAddTagEcs = &v
	return s
}

func (s *AddAddressBookRequest) SetTagRelation(v string) *AddAddressBookRequest {
	s.TagRelation = &v
	return s
}

func (s *AddAddressBookRequest) SetTagList(v []*AddAddressBookRequestTagList) *AddAddressBookRequest {
	s.TagList = v
	return s
}

type AddAddressBookRequestTagList struct {
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s AddAddressBookRequestTagList) String() string {
	return tea.Prettify(s)
}

func (s AddAddressBookRequestTagList) GoString() string {
	return s.String()
}

func (s *AddAddressBookRequestTagList) SetTagValue(v string) *AddAddressBookRequestTagList {
	s.TagValue = &v
	return s
}

func (s *AddAddressBookRequestTagList) SetTagKey(v string) *AddAddressBookRequestTagList {
	s.TagKey = &v
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddAddressBookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddAddressBookResponse) SetBody(v *AddAddressBookResponseBody) *AddAddressBookResponse {
	s.Body = v
	return s
}

type AddControlPolicyRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	AclAction       *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DestPort        *string `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	Destination     *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	Direction       *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	Proto           *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	Source          *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceType      *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	NewOrder        *string `json:"NewOrder,omitempty" xml:"NewOrder,omitempty"`
	DestPortType    *string `json:"DestPortType,omitempty" xml:"DestPortType,omitempty"`
	DestPortGroup   *string `json:"DestPortGroup,omitempty" xml:"DestPortGroup,omitempty"`
	Release         *string `json:"Release,omitempty" xml:"Release,omitempty"`
}

func (s AddControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s AddControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *AddControlPolicyRequest) SetSourceIp(v string) *AddControlPolicyRequest {
	s.SourceIp = &v
	return s
}

func (s *AddControlPolicyRequest) SetLang(v string) *AddControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *AddControlPolicyRequest) SetAclAction(v string) *AddControlPolicyRequest {
	s.AclAction = &v
	return s
}

func (s *AddControlPolicyRequest) SetApplicationName(v string) *AddControlPolicyRequest {
	s.ApplicationName = &v
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

func (s *AddControlPolicyRequest) SetProto(v string) *AddControlPolicyRequest {
	s.Proto = &v
	return s
}

func (s *AddControlPolicyRequest) SetSource(v string) *AddControlPolicyRequest {
	s.Source = &v
	return s
}

func (s *AddControlPolicyRequest) SetSourceType(v string) *AddControlPolicyRequest {
	s.SourceType = &v
	return s
}

func (s *AddControlPolicyRequest) SetNewOrder(v string) *AddControlPolicyRequest {
	s.NewOrder = &v
	return s
}

func (s *AddControlPolicyRequest) SetDestPortType(v string) *AddControlPolicyRequest {
	s.DestPortType = &v
	return s
}

func (s *AddControlPolicyRequest) SetDestPortGroup(v string) *AddControlPolicyRequest {
	s.DestPortGroup = &v
	return s
}

func (s *AddControlPolicyRequest) SetRelease(v string) *AddControlPolicyRequest {
	s.Release = &v
	return s
}

type AddControlPolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AclUuid   *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
}

func (s AddControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *AddControlPolicyResponseBody) SetRequestId(v string) *AddControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddControlPolicyResponseBody) SetAclUuid(v string) *AddControlPolicyResponseBody {
	s.AclUuid = &v
	return s
}

type AddControlPolicyResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddControlPolicyResponse) SetBody(v *AddControlPolicyResponseBody) *AddControlPolicyResponse {
	s.Body = v
	return s
}

type CreateVpcFirewallControlPolicyRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	AclAction       *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DestPort        *string `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	Destination     *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	VpcFirewallId   *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
	Proto           *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	Source          *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceType      *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	NewOrder        *string `json:"NewOrder,omitempty" xml:"NewOrder,omitempty"`
	DestPortType    *string `json:"DestPortType,omitempty" xml:"DestPortType,omitempty"`
	DestPortGroup   *string `json:"DestPortGroup,omitempty" xml:"DestPortGroup,omitempty"`
}

func (s CreateVpcFirewallControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcFirewallControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateVpcFirewallControlPolicyRequest) SetSourceIp(v string) *CreateVpcFirewallControlPolicyRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetLang(v string) *CreateVpcFirewallControlPolicyRequest {
	s.Lang = &v
	return s
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

func (s *CreateVpcFirewallControlPolicyRequest) SetDestination(v string) *CreateVpcFirewallControlPolicyRequest {
	s.Destination = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetDestinationType(v string) *CreateVpcFirewallControlPolicyRequest {
	s.DestinationType = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetVpcFirewallId(v string) *CreateVpcFirewallControlPolicyRequest {
	s.VpcFirewallId = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetProto(v string) *CreateVpcFirewallControlPolicyRequest {
	s.Proto = &v
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

func (s *CreateVpcFirewallControlPolicyRequest) SetNewOrder(v string) *CreateVpcFirewallControlPolicyRequest {
	s.NewOrder = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetDestPortType(v string) *CreateVpcFirewallControlPolicyRequest {
	s.DestPortType = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyRequest) SetDestPortGroup(v string) *CreateVpcFirewallControlPolicyRequest {
	s.DestPortGroup = &v
	return s
}

type CreateVpcFirewallControlPolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AclUuid   *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
}

func (s CreateVpcFirewallControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVpcFirewallControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVpcFirewallControlPolicyResponseBody) SetRequestId(v string) *CreateVpcFirewallControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVpcFirewallControlPolicyResponseBody) SetAclUuid(v string) *CreateVpcFirewallControlPolicyResponseBody {
	s.AclUuid = &v
	return s
}

type CreateVpcFirewallControlPolicyResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateVpcFirewallControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateVpcFirewallControlPolicyResponse) SetBody(v *CreateVpcFirewallControlPolicyResponseBody) *CreateVpcFirewallControlPolicyResponse {
	s.Body = v
	return s
}

type DeleteAddressBookRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	GroupUuid *string `json:"GroupUuid,omitempty" xml:"GroupUuid,omitempty"`
}

func (s DeleteAddressBookRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAddressBookRequest) GoString() string {
	return s.String()
}

func (s *DeleteAddressBookRequest) SetSourceIp(v string) *DeleteAddressBookRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteAddressBookRequest) SetLang(v string) *DeleteAddressBookRequest {
	s.Lang = &v
	return s
}

func (s *DeleteAddressBookRequest) SetGroupUuid(v string) *DeleteAddressBookRequest {
	s.GroupUuid = &v
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAddressBookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteAddressBookResponse) SetBody(v *DeleteAddressBookResponseBody) *DeleteAddressBookResponse {
	s.Body = v
	return s
}

type DeleteControlPolicyRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	AclUuid   *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
}

func (s DeleteControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteControlPolicyRequest) SetSourceIp(v string) *DeleteControlPolicyRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteControlPolicyRequest) SetLang(v string) *DeleteControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *DeleteControlPolicyRequest) SetAclUuid(v string) *DeleteControlPolicyRequest {
	s.AclUuid = &v
	return s
}

func (s *DeleteControlPolicyRequest) SetDirection(v string) *DeleteControlPolicyRequest {
	s.Direction = &v
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteControlPolicyResponse) SetBody(v *DeleteControlPolicyResponseBody) *DeleteControlPolicyResponse {
	s.Body = v
	return s
}

type DeleteVpcFirewallControlPolicyRequest struct {
	SourceIp      *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang          *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	AclUuid       *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s DeleteVpcFirewallControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVpcFirewallControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteVpcFirewallControlPolicyRequest) SetSourceIp(v string) *DeleteVpcFirewallControlPolicyRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteVpcFirewallControlPolicyRequest) SetLang(v string) *DeleteVpcFirewallControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *DeleteVpcFirewallControlPolicyRequest) SetAclUuid(v string) *DeleteVpcFirewallControlPolicyRequest {
	s.AclUuid = &v
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
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteVpcFirewallControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteVpcFirewallControlPolicyResponse) SetBody(v *DeleteVpcFirewallControlPolicyResponseBody) *DeleteVpcFirewallControlPolicyResponse {
	s.Body = v
	return s
}

type DescribePolicyPriorUsedRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
}

func (s DescribePolicyPriorUsedRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyPriorUsedRequest) GoString() string {
	return s.String()
}

func (s *DescribePolicyPriorUsedRequest) SetSourceIp(v string) *DescribePolicyPriorUsedRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribePolicyPriorUsedRequest) SetLang(v string) *DescribePolicyPriorUsedRequest {
	s.Lang = &v
	return s
}

func (s *DescribePolicyPriorUsedRequest) SetDirection(v string) *DescribePolicyPriorUsedRequest {
	s.Direction = &v
	return s
}

type DescribePolicyPriorUsedResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Start     *int32  `json:"Start,omitempty" xml:"Start,omitempty"`
	End       *int32  `json:"End,omitempty" xml:"End,omitempty"`
}

func (s DescribePolicyPriorUsedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePolicyPriorUsedResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolicyPriorUsedResponseBody) SetRequestId(v string) *DescribePolicyPriorUsedResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePolicyPriorUsedResponseBody) SetStart(v int32) *DescribePolicyPriorUsedResponseBody {
	s.Start = &v
	return s
}

func (s *DescribePolicyPriorUsedResponseBody) SetEnd(v int32) *DescribePolicyPriorUsedResponseBody {
	s.End = &v
	return s
}

type DescribePolicyPriorUsedResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePolicyPriorUsedResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribePolicyPriorUsedResponse) SetBody(v *DescribePolicyPriorUsedResponseBody) *DescribePolicyPriorUsedResponse {
	s.Body = v
	return s
}

type DescribeVpcFirewallControlPolicyRequest struct {
	SourceIp      *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang          *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
	CurrentPage   *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize      *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Destination   *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Proto         *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	AclAction     *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
}

func (s DescribeVpcFirewallControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetSourceIp(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetLang(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetVpcFirewallId(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.VpcFirewallId = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetCurrentPage(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetPageSize(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetSource(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.Source = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetDestination(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.Destination = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetDescription(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.Description = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetProto(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.Proto = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyRequest) SetAclAction(v string) *DescribeVpcFirewallControlPolicyRequest {
	s.AclAction = &v
	return s
}

type DescribeVpcFirewallControlPolicyResponseBody struct {
	TotalCount *string                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Policys    []*DescribeVpcFirewallControlPolicyResponseBodyPolicys `json:"Policys,omitempty" xml:"Policys,omitempty" type:"Repeated"`
}

func (s DescribeVpcFirewallControlPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallControlPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallControlPolicyResponseBody) SetTotalCount(v string) *DescribeVpcFirewallControlPolicyResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBody) SetRequestId(v string) *DescribeVpcFirewallControlPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBody) SetPolicys(v []*DescribeVpcFirewallControlPolicyResponseBodyPolicys) *DescribeVpcFirewallControlPolicyResponseBody {
	s.Policys = v
	return s
}

type DescribeVpcFirewallControlPolicyResponseBodyPolicys struct {
	Direction             *string   `json:"Direction,omitempty" xml:"Direction,omitempty"`
	Destination           *string   `json:"Destination,omitempty" xml:"Destination,omitempty"`
	Order                 *int32    `json:"Order,omitempty" xml:"Order,omitempty"`
	ApplicationName       *string   `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	SourceType            *string   `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	DestPortGroup         *string   `json:"DestPortGroup,omitempty" xml:"DestPortGroup,omitempty"`
	AclUuid               *string   `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	DestPortGroupPorts    []*string `json:"DestPortGroupPorts,omitempty" xml:"DestPortGroupPorts,omitempty" type:"Repeated"`
	DestPortType          *string   `json:"DestPortType,omitempty" xml:"DestPortType,omitempty"`
	Source                *string   `json:"Source,omitempty" xml:"Source,omitempty"`
	DestinationType       *string   `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	HitTimes              *int32    `json:"HitTimes,omitempty" xml:"HitTimes,omitempty"`
	SourceGroupCidrs      []*string `json:"SourceGroupCidrs,omitempty" xml:"SourceGroupCidrs,omitempty" type:"Repeated"`
	DestPort              *string   `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	DestinationGroupCidrs []*string `json:"DestinationGroupCidrs,omitempty" xml:"DestinationGroupCidrs,omitempty" type:"Repeated"`
	Description           *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	AclAction             *string   `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	ApplicationId         *string   `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	Proto                 *string   `json:"Proto,omitempty" xml:"Proto,omitempty"`
}

func (s DescribeVpcFirewallControlPolicyResponseBodyPolicys) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallControlPolicyResponseBodyPolicys) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetDirection(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.Direction = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetDestination(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.Destination = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetOrder(v int32) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.Order = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetApplicationName(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.ApplicationName = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetSourceType(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.SourceType = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetDestPortGroup(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.DestPortGroup = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetAclUuid(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.AclUuid = &v
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

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetSource(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.Source = &v
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

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetSourceGroupCidrs(v []*string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.SourceGroupCidrs = v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetDestPort(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.DestPort = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetDestinationGroupCidrs(v []*string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.DestinationGroupCidrs = v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetDescription(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.Description = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetAclAction(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.AclAction = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetApplicationId(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.ApplicationId = &v
	return s
}

func (s *DescribeVpcFirewallControlPolicyResponseBodyPolicys) SetProto(v string) *DescribeVpcFirewallControlPolicyResponseBodyPolicys {
	s.Proto = &v
	return s
}

type DescribeVpcFirewallControlPolicyResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVpcFirewallControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeVpcFirewallControlPolicyResponse) SetBody(v *DescribeVpcFirewallControlPolicyResponseBody) *DescribeVpcFirewallControlPolicyResponse {
	s.Body = v
	return s
}

type DescribeVpcFirewallPolicyPriorUsedRequest struct {
	SourceIp      *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang          *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
}

func (s DescribeVpcFirewallPolicyPriorUsedRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallPolicyPriorUsedRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallPolicyPriorUsedRequest) SetSourceIp(v string) *DescribeVpcFirewallPolicyPriorUsedRequest {
	s.SourceIp = &v
	return s
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Start     *int32  `json:"Start,omitempty" xml:"Start,omitempty"`
	End       *int32  `json:"End,omitempty" xml:"End,omitempty"`
}

func (s DescribeVpcFirewallPolicyPriorUsedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcFirewallPolicyPriorUsedResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallPolicyPriorUsedResponseBody) SetRequestId(v string) *DescribeVpcFirewallPolicyPriorUsedResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallPolicyPriorUsedResponseBody) SetStart(v int32) *DescribeVpcFirewallPolicyPriorUsedResponseBody {
	s.Start = &v
	return s
}

func (s *DescribeVpcFirewallPolicyPriorUsedResponseBody) SetEnd(v int32) *DescribeVpcFirewallPolicyPriorUsedResponseBody {
	s.End = &v
	return s
}

type DescribeVpcFirewallPolicyPriorUsedResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVpcFirewallPolicyPriorUsedResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeVpcFirewallPolicyPriorUsedResponse) SetBody(v *DescribeVpcFirewallPolicyPriorUsedResponseBody) *DescribeVpcFirewallPolicyPriorUsedResponse {
	s.Body = v
	return s
}

type ModifyAddressBookRequest struct {
	SourceIp      *string                            `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang          *string                            `json:"Lang,omitempty" xml:"Lang,omitempty"`
	AddressList   *string                            `json:"AddressList,omitempty" xml:"AddressList,omitempty"`
	Description   *string                            `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupName     *string                            `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	GroupUuid     *string                            `json:"GroupUuid,omitempty" xml:"GroupUuid,omitempty"`
	AutoAddTagEcs *string                            `json:"AutoAddTagEcs,omitempty" xml:"AutoAddTagEcs,omitempty"`
	TagRelation   *string                            `json:"TagRelation,omitempty" xml:"TagRelation,omitempty"`
	TagList       []*ModifyAddressBookRequestTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
}

func (s ModifyAddressBookRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAddressBookRequest) GoString() string {
	return s.String()
}

func (s *ModifyAddressBookRequest) SetSourceIp(v string) *ModifyAddressBookRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyAddressBookRequest) SetLang(v string) *ModifyAddressBookRequest {
	s.Lang = &v
	return s
}

func (s *ModifyAddressBookRequest) SetAddressList(v string) *ModifyAddressBookRequest {
	s.AddressList = &v
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

func (s *ModifyAddressBookRequest) SetAutoAddTagEcs(v string) *ModifyAddressBookRequest {
	s.AutoAddTagEcs = &v
	return s
}

func (s *ModifyAddressBookRequest) SetTagRelation(v string) *ModifyAddressBookRequest {
	s.TagRelation = &v
	return s
}

func (s *ModifyAddressBookRequest) SetTagList(v []*ModifyAddressBookRequestTagList) *ModifyAddressBookRequest {
	s.TagList = v
	return s
}

type ModifyAddressBookRequestTagList struct {
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ModifyAddressBookRequestTagList) String() string {
	return tea.Prettify(s)
}

func (s ModifyAddressBookRequestTagList) GoString() string {
	return s.String()
}

func (s *ModifyAddressBookRequestTagList) SetTagValue(v string) *ModifyAddressBookRequestTagList {
	s.TagValue = &v
	return s
}

func (s *ModifyAddressBookRequestTagList) SetTagKey(v string) *ModifyAddressBookRequestTagList {
	s.TagKey = &v
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyAddressBookResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyAddressBookResponse) SetBody(v *ModifyAddressBookResponseBody) *ModifyAddressBookResponse {
	s.Body = v
	return s
}

type ModifyControlPolicyRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	AclAction       *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DestPort        *string `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	Destination     *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	Direction       *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	Proto           *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	Source          *string `json:"Source,omitempty" xml:"Source,omitempty"`
	AclUuid         *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	SourceType      *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	DestPortType    *string `json:"DestPortType,omitempty" xml:"DestPortType,omitempty"`
	DestPortGroup   *string `json:"DestPortGroup,omitempty" xml:"DestPortGroup,omitempty"`
	Release         *string `json:"Release,omitempty" xml:"Release,omitempty"`
}

func (s ModifyControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyControlPolicyRequest) SetSourceIp(v string) *ModifyControlPolicyRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetLang(v string) *ModifyControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetAclAction(v string) *ModifyControlPolicyRequest {
	s.AclAction = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetApplicationName(v string) *ModifyControlPolicyRequest {
	s.ApplicationName = &v
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

func (s *ModifyControlPolicyRequest) SetProto(v string) *ModifyControlPolicyRequest {
	s.Proto = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetSource(v string) *ModifyControlPolicyRequest {
	s.Source = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetAclUuid(v string) *ModifyControlPolicyRequest {
	s.AclUuid = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetSourceType(v string) *ModifyControlPolicyRequest {
	s.SourceType = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetDestPortType(v string) *ModifyControlPolicyRequest {
	s.DestPortType = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetDestPortGroup(v string) *ModifyControlPolicyRequest {
	s.DestPortGroup = &v
	return s
}

func (s *ModifyControlPolicyRequest) SetRelease(v string) *ModifyControlPolicyRequest {
	s.Release = &v
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyControlPolicyResponse) SetBody(v *ModifyControlPolicyResponseBody) *ModifyControlPolicyResponse {
	s.Body = v
	return s
}

type ModifyControlPolicyPositionRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	NewOrder  *string `json:"NewOrder,omitempty" xml:"NewOrder,omitempty"`
	OldOrder  *string `json:"OldOrder,omitempty" xml:"OldOrder,omitempty"`
}

func (s ModifyControlPolicyPositionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyControlPolicyPositionRequest) GoString() string {
	return s.String()
}

func (s *ModifyControlPolicyPositionRequest) SetSourceIp(v string) *ModifyControlPolicyPositionRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyControlPolicyPositionRequest) SetLang(v string) *ModifyControlPolicyPositionRequest {
	s.Lang = &v
	return s
}

func (s *ModifyControlPolicyPositionRequest) SetDirection(v string) *ModifyControlPolicyPositionRequest {
	s.Direction = &v
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
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyControlPolicyPositionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyControlPolicyPositionResponse) SetBody(v *ModifyControlPolicyPositionResponseBody) *ModifyControlPolicyPositionResponse {
	s.Body = v
	return s
}

type ModifyPolicyAdvancedConfigRequest struct {
	SourceIp       *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang           *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	InternetSwitch *string `json:"InternetSwitch,omitempty" xml:"InternetSwitch,omitempty"`
}

func (s ModifyPolicyAdvancedConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPolicyAdvancedConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyPolicyAdvancedConfigRequest) SetSourceIp(v string) *ModifyPolicyAdvancedConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyPolicyAdvancedConfigRequest) SetLang(v string) *ModifyPolicyAdvancedConfigRequest {
	s.Lang = &v
	return s
}

func (s *ModifyPolicyAdvancedConfigRequest) SetInternetSwitch(v string) *ModifyPolicyAdvancedConfigRequest {
	s.InternetSwitch = &v
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
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyPolicyAdvancedConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyPolicyAdvancedConfigResponse) SetBody(v *ModifyPolicyAdvancedConfigResponseBody) *ModifyPolicyAdvancedConfigResponse {
	s.Body = v
	return s
}

type ModifyVpcFirewallControlPolicyRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang            *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	AclAction       *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DestPort        *string `json:"DestPort,omitempty" xml:"DestPort,omitempty"`
	Destination     *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	VpcFirewallId   *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
	Proto           *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	Source          *string `json:"Source,omitempty" xml:"Source,omitempty"`
	AclUuid         *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	SourceType      *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	DestPortType    *string `json:"DestPortType,omitempty" xml:"DestPortType,omitempty"`
	DestPortGroup   *string `json:"DestPortGroup,omitempty" xml:"DestPortGroup,omitempty"`
}

func (s ModifyVpcFirewallControlPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetSourceIp(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetLang(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.Lang = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetAclAction(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.AclAction = &v
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

func (s *ModifyVpcFirewallControlPolicyRequest) SetDestination(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.Destination = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetDestinationType(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.DestinationType = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetVpcFirewallId(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.VpcFirewallId = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetProto(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.Proto = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetSource(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.Source = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetAclUuid(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.AclUuid = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetSourceType(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.SourceType = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetDestPortType(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.DestPortType = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyRequest) SetDestPortGroup(v string) *ModifyVpcFirewallControlPolicyRequest {
	s.DestPortGroup = &v
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
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyVpcFirewallControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyVpcFirewallControlPolicyResponse) SetBody(v *ModifyVpcFirewallControlPolicyResponseBody) *ModifyVpcFirewallControlPolicyResponse {
	s.Body = v
	return s
}

type ModifyVpcFirewallControlPolicyPositionRequest struct {
	SourceIp      *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang          *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
	NewOrder      *string `json:"NewOrder,omitempty" xml:"NewOrder,omitempty"`
	OldOrder      *string `json:"OldOrder,omitempty" xml:"OldOrder,omitempty"`
}

func (s ModifyVpcFirewallControlPolicyPositionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyVpcFirewallControlPolicyPositionRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallControlPolicyPositionRequest) SetSourceIp(v string) *ModifyVpcFirewallControlPolicyPositionRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyPositionRequest) SetLang(v string) *ModifyVpcFirewallControlPolicyPositionRequest {
	s.Lang = &v
	return s
}

func (s *ModifyVpcFirewallControlPolicyPositionRequest) SetVpcFirewallId(v string) *ModifyVpcFirewallControlPolicyPositionRequest {
	s.VpcFirewallId = &v
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
	Headers map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyVpcFirewallControlPolicyPositionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyVpcFirewallControlPolicyPositionResponse) SetBody(v *ModifyVpcFirewallControlPolicyPositionResponseBody) *ModifyVpcFirewallControlPolicyPositionResponse {
	s.Body = v
	return s
}

type PutDisableAllFwSwitchRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s PutDisableAllFwSwitchRequest) String() string {
	return tea.Prettify(s)
}

func (s PutDisableAllFwSwitchRequest) GoString() string {
	return s.String()
}

func (s *PutDisableAllFwSwitchRequest) SetSourceIp(v string) *PutDisableAllFwSwitchRequest {
	s.SourceIp = &v
	return s
}

func (s *PutDisableAllFwSwitchRequest) SetLang(v string) *PutDisableAllFwSwitchRequest {
	s.Lang = &v
	return s
}

func (s *PutDisableAllFwSwitchRequest) SetInstanceId(v string) *PutDisableAllFwSwitchRequest {
	s.InstanceId = &v
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
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PutDisableAllFwSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *PutDisableAllFwSwitchResponse) SetBody(v *PutDisableAllFwSwitchResponseBody) *PutDisableAllFwSwitchResponse {
	s.Body = v
	return s
}

type PutDisableFwSwitchRequest struct {
	SourceIp         *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang             *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	IpaddrList       []*string `json:"IpaddrList,omitempty" xml:"IpaddrList,omitempty" type:"Repeated"`
	RegionList       []*string `json:"RegionList,omitempty" xml:"RegionList,omitempty" type:"Repeated"`
	ResourceTypeList []*string `json:"ResourceTypeList,omitempty" xml:"ResourceTypeList,omitempty" type:"Repeated"`
}

func (s PutDisableFwSwitchRequest) String() string {
	return tea.Prettify(s)
}

func (s PutDisableFwSwitchRequest) GoString() string {
	return s.String()
}

func (s *PutDisableFwSwitchRequest) SetSourceIp(v string) *PutDisableFwSwitchRequest {
	s.SourceIp = &v
	return s
}

func (s *PutDisableFwSwitchRequest) SetLang(v string) *PutDisableFwSwitchRequest {
	s.Lang = &v
	return s
}

func (s *PutDisableFwSwitchRequest) SetIpaddrList(v []*string) *PutDisableFwSwitchRequest {
	s.IpaddrList = v
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PutDisableFwSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *PutDisableFwSwitchResponse) SetBody(v *PutDisableFwSwitchResponseBody) *PutDisableFwSwitchResponse {
	s.Body = v
	return s
}

type PutEnableAllFwSwitchRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s PutEnableAllFwSwitchRequest) String() string {
	return tea.Prettify(s)
}

func (s PutEnableAllFwSwitchRequest) GoString() string {
	return s.String()
}

func (s *PutEnableAllFwSwitchRequest) SetSourceIp(v string) *PutEnableAllFwSwitchRequest {
	s.SourceIp = &v
	return s
}

func (s *PutEnableAllFwSwitchRequest) SetLang(v string) *PutEnableAllFwSwitchRequest {
	s.Lang = &v
	return s
}

func (s *PutEnableAllFwSwitchRequest) SetInstanceId(v string) *PutEnableAllFwSwitchRequest {
	s.InstanceId = &v
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
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PutEnableAllFwSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *PutEnableAllFwSwitchResponse) SetBody(v *PutEnableAllFwSwitchResponseBody) *PutEnableAllFwSwitchResponse {
	s.Body = v
	return s
}

type PutEnableFwSwitchRequest struct {
	SourceIp         *string   `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang             *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	IpaddrList       []*string `json:"IpaddrList,omitempty" xml:"IpaddrList,omitempty" type:"Repeated"`
	RegionList       []*string `json:"RegionList,omitempty" xml:"RegionList,omitempty" type:"Repeated"`
	ResourceTypeList []*string `json:"ResourceTypeList,omitempty" xml:"ResourceTypeList,omitempty" type:"Repeated"`
}

func (s PutEnableFwSwitchRequest) String() string {
	return tea.Prettify(s)
}

func (s PutEnableFwSwitchRequest) GoString() string {
	return s.String()
}

func (s *PutEnableFwSwitchRequest) SetSourceIp(v string) *PutEnableFwSwitchRequest {
	s.SourceIp = &v
	return s
}

func (s *PutEnableFwSwitchRequest) SetLang(v string) *PutEnableFwSwitchRequest {
	s.Lang = &v
	return s
}

func (s *PutEnableFwSwitchRequest) SetIpaddrList(v []*string) *PutEnableFwSwitchRequest {
	s.IpaddrList = v
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PutEnableFwSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *PutEnableFwSwitchResponse) SetBody(v *PutEnableFwSwitchResponseBody) *PutEnableFwSwitchResponse {
	s.Body = v
	return s
}

type ResetVpcFirewallRuleHitCountRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	AclUuid  *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
}

func (s ResetVpcFirewallRuleHitCountRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetVpcFirewallRuleHitCountRequest) GoString() string {
	return s.String()
}

func (s *ResetVpcFirewallRuleHitCountRequest) SetSourceIp(v string) *ResetVpcFirewallRuleHitCountRequest {
	s.SourceIp = &v
	return s
}

func (s *ResetVpcFirewallRuleHitCountRequest) SetLang(v string) *ResetVpcFirewallRuleHitCountRequest {
	s.Lang = &v
	return s
}

func (s *ResetVpcFirewallRuleHitCountRequest) SetAclUuid(v string) *ResetVpcFirewallRuleHitCountRequest {
	s.AclUuid = &v
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
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResetVpcFirewallRuleHitCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddAddressBookResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddAddressBook"), tea.String("2017-12-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddControlPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddControlPolicy"), tea.String("2017-12-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) CreateVpcFirewallControlPolicyWithOptions(request *CreateVpcFirewallControlPolicyRequest, runtime *util.RuntimeOptions) (_result *CreateVpcFirewallControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateVpcFirewallControlPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateVpcFirewallControlPolicy"), tea.String("2017-12-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAddressBookResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAddressBook"), tea.String("2017-12-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteControlPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteControlPolicy"), tea.String("2017-12-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DeleteVpcFirewallControlPolicyWithOptions(request *DeleteVpcFirewallControlPolicyRequest, runtime *util.RuntimeOptions) (_result *DeleteVpcFirewallControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteVpcFirewallControlPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteVpcFirewallControlPolicy"), tea.String("2017-12-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribePolicyPriorUsedWithOptions(request *DescribePolicyPriorUsedRequest, runtime *util.RuntimeOptions) (_result *DescribePolicyPriorUsedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePolicyPriorUsedResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePolicyPriorUsed"), tea.String("2017-12-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeVpcFirewallControlPolicyWithOptions(request *DescribeVpcFirewallControlPolicyRequest, runtime *util.RuntimeOptions) (_result *DescribeVpcFirewallControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVpcFirewallControlPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVpcFirewallControlPolicy"), tea.String("2017-12-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeVpcFirewallPolicyPriorUsedWithOptions(request *DescribeVpcFirewallPolicyPriorUsedRequest, runtime *util.RuntimeOptions) (_result *DescribeVpcFirewallPolicyPriorUsedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeVpcFirewallPolicyPriorUsedResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeVpcFirewallPolicyPriorUsed"), tea.String("2017-12-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyAddressBookResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyAddressBook"), tea.String("2017-12-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyControlPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyControlPolicy"), tea.String("2017-12-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyControlPolicyPositionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyControlPolicyPosition"), tea.String("2017-12-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ModifyPolicyAdvancedConfigWithOptions(request *ModifyPolicyAdvancedConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyPolicyAdvancedConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyPolicyAdvancedConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyPolicyAdvancedConfig"), tea.String("2017-12-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ModifyVpcFirewallControlPolicyWithOptions(request *ModifyVpcFirewallControlPolicyRequest, runtime *util.RuntimeOptions) (_result *ModifyVpcFirewallControlPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyVpcFirewallControlPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyVpcFirewallControlPolicy"), tea.String("2017-12-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyVpcFirewallControlPolicyPositionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyVpcFirewallControlPolicyPosition"), tea.String("2017-12-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) PutDisableAllFwSwitchWithOptions(request *PutDisableAllFwSwitchRequest, runtime *util.RuntimeOptions) (_result *PutDisableAllFwSwitchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PutDisableAllFwSwitchResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PutDisableAllFwSwitch"), tea.String("2017-12-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PutDisableFwSwitchResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PutDisableFwSwitch"), tea.String("2017-12-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PutEnableAllFwSwitchResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PutEnableAllFwSwitch"), tea.String("2017-12-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &PutEnableFwSwitchResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PutEnableFwSwitch"), tea.String("2017-12-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ResetVpcFirewallRuleHitCountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ResetVpcFirewallRuleHitCount"), tea.String("2017-12-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
