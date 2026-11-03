// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeControlPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclAction(v string) *DescribeControlPolicyRequest
	GetAclAction() *string
	SetAclUuid(v string) *DescribeControlPolicyRequest
	GetAclUuid() *string
	SetCurrentPage(v string) *DescribeControlPolicyRequest
	GetCurrentPage() *string
	SetDescription(v string) *DescribeControlPolicyRequest
	GetDescription() *string
	SetDestination(v string) *DescribeControlPolicyRequest
	GetDestination() *string
	SetDirection(v string) *DescribeControlPolicyRequest
	GetDirection() *string
	SetIpVersion(v string) *DescribeControlPolicyRequest
	GetIpVersion() *string
	SetLang(v string) *DescribeControlPolicyRequest
	GetLang() *string
	SetPageSize(v string) *DescribeControlPolicyRequest
	GetPageSize() *string
	SetProto(v string) *DescribeControlPolicyRequest
	GetProto() *string
	SetRelease(v string) *DescribeControlPolicyRequest
	GetRelease() *string
	SetRepeatType(v string) *DescribeControlPolicyRequest
	GetRepeatType() *string
	SetSource(v string) *DescribeControlPolicyRequest
	GetSource() *string
}

type DescribeControlPolicyRequest struct {
	// The action that Cloud Firewall performs on the traffic that matches the access control policy. Valid values:
	//
	// - **accept**: Allow
	//
	// - **drop**: Deny
	//
	// - **log**: Monitor
	//
	// > If you do not specify this parameter, all action types are queried.
	//
	// example:
	//
	// accept
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The unique ID of the access control policy. You must specify at least one of AclUuid and Direction. If you specify AclUuid, you can query the policy by its ID.
	//
	// example:
	//
	// 00281255-d220-4db1-8f4f-c4df221a****
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The page number of the current page returned for a paged query during paging.
	//
	// Default value: 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The description of the access control policy. Fuzzy match is supported.
	//
	// > If you do not specify this parameter, the descriptions of all policies are queried.
	//
	// example:
	//
	// Allow access to the office CIDR block
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination address in the access control policy. Fuzzy match is supported. The value of Destination varies based on the destination type (DestinationType).
	//
	// - If DestinationType is set to `net`, the destination address is a CIDR block. Example: 10.0.3.0/24.
	//
	// - If DestinationType is set to `domain`, the destination address is a domain name. Example: aliyun.
	//
	// - If DestinationType is set to `group`, the destination address is the name of an address book. Example: db_group.
	//
	// - If DestinationType is set to `location`, the destination address is an area name. For more information about the area positional encoding, see AddControlPolicy. Example: ["BJ11", "ZB"].
	//
	// > If you do not specify this parameter in the Settings, all types of destination addresses are queried.
	//
	// example:
	//
	// 192.0.XX.XX
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The traffic direction controlled by the access control policy. Valid values: in (inbound) and out (outbound). You must specify at least one of Direction and AclUuid. If AclUuid is not specified, you must specify a non-empty Direction. Otherwise, ErrorParametersDirection is returned.
	//
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The IP address version. Valid values:
	//
	// - **4*	- (default): IPv4
	//
	// - **6**: IPv6
	//
	// example:
	//
	// 6
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The language of the response message. Valid values:
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of entries to return on each page for a paged query during paging.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The protocol type of the traffic in the access control policy. Valid values:
	//
	// - **TCP**
	//
	// - **UDP**
	//
	// - **ICMP**
	//
	// - **ANY*	- (all protocol types)
	//
	// > If you do not specify this parameter, all protocol types are queried.
	//
	// example:
	//
	// TCP
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// Specifies whether the access control policy is enabled. Valid values:
	//
	// - **true**: The access control policy is enabled.
	//
	// - **false**: The access control policy is disabled.
	//
	// example:
	//
	// true
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The recurrence type of the Policy Validity Period for the access control policy. Valid values:
	//
	// - **Permanent*	- (default): Always
	//
	// - **None**: One-time
	//
	// - **Daily**: Daily
	//
	// - **Weekly**: Weekly
	//
	// - **Monthly**: Monthly
	//
	// example:
	//
	// Permanent
	RepeatType *string `json:"RepeatType,omitempty" xml:"RepeatType,omitempty"`
	// The source address in the access control policy. Fuzzy match is supported. The value of Source varies based on the value of SourceType.
	//
	// - If SourceType is set to `net`, the source address is a CIDR block. Example: 192.0.XX.XX/24.
	//
	// - If SourceType is set to `group`, the source address is the name of a source address book. Example: db_group. If this parameter is left empty, all access sources are queried.
	//
	// - If SourceType is set to `location`, the source address is a source region. Example: Beijing or beijing. Both Chinese and English region names are supported.
	//
	// > If you do not specify this parameter, all types of access sources are queried.
	//
	// example:
	//
	// 192.0.XX.XX
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s DescribeControlPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeControlPolicyRequest) GetAclAction() *string {
	return s.AclAction
}

func (s *DescribeControlPolicyRequest) GetAclUuid() *string {
	return s.AclUuid
}

func (s *DescribeControlPolicyRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeControlPolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *DescribeControlPolicyRequest) GetDestination() *string {
	return s.Destination
}

func (s *DescribeControlPolicyRequest) GetDirection() *string {
	return s.Direction
}

func (s *DescribeControlPolicyRequest) GetIpVersion() *string {
	return s.IpVersion
}

func (s *DescribeControlPolicyRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeControlPolicyRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeControlPolicyRequest) GetProto() *string {
	return s.Proto
}

func (s *DescribeControlPolicyRequest) GetRelease() *string {
	return s.Release
}

func (s *DescribeControlPolicyRequest) GetRepeatType() *string {
	return s.RepeatType
}

func (s *DescribeControlPolicyRequest) GetSource() *string {
	return s.Source
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

func (s *DescribeControlPolicyRequest) Validate() error {
	return dara.Validate(s)
}
