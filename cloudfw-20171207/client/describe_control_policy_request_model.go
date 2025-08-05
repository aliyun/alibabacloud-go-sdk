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
	// The action that Cloud Firewall performs on the traffic. Valid values:
	//
	// 	- **accept**: allows the traffic.
	//
	// 	- **drop**: denies the traffic.
	//
	// 	- **log**: monitors the traffic.
	//
	// >  If you do not specify this parameter, access control policies of all action types are queried.
	//
	// example:
	//
	// accept
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The unique ID of the access control policy.
	//
	// example:
	//
	// 00281255-d220-4db1-8f4f-c4df221a****
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The number of the page to return.
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
	// >  If you do not specify this parameter, access control policies that have descriptions are queried.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination address in the access control policy. Fuzzy match is supported. The value of this parameter varies based on the value of the DestinationType parameter.
	//
	// 	- If you set DestinationType to `net`, the value of Destination is a CIDR block. Example: 10.0.3.0/24.
	//
	// 	- If you set DestinationType to `domain`, the value of Destination is a domain name. Example: aliyun.
	//
	// 	- If you set DestinationType to `group`, the value of Destination is the name of an address book. Example: db_group.
	//
	// 	- If you set DestinationType to `location`, the value of Destination is the name of a location. For more information about location codes, see AddControlPolicy. Example: ["BJ11", "ZB"].
	//
	// >  If you do not specify this parameter, access control policies of all destination address types are queried.
	//
	// example:
	//
	// 192.0.XX.XX
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The direction of the traffic to which the access control policies apply. Valid values:
	//
	// 	- **in**: inbound.
	//
	// 	- **out**: outbound.
	//
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The IP version of the address in the access control policy. Valid values:
	//
	// 	- **4**: IPv4 (default)
	//
	// 	- **6**: IPv6
	//
	// example:
	//
	// 6
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the protocol in the access control policy. Valid values:
	//
	// 	- **TCP**
	//
	// 	- **UDP**
	//
	// 	- **ICMP**
	//
	// 	- **ANY**: all types of protocols
	//
	// >  If you do not specify this parameter, access control policies of all protocol types are queried.
	//
	// example:
	//
	// TCP
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// Specifies whether the access control policy is enabled. By default, an access control policy is enabled after it is created. Valid values:
	//
	// 	- **true**: The access control policy is enabled.
	//
	// 	- **false**: The access control policy is disabled.
	//
	// example:
	//
	// true
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The recurrence type for the access control policy to take effect. Valid values:
	//
	// 	- **Permanent*	- (default): The policy always takes effect.
	//
	// 	- **None**: The policy takes effect for only once.
	//
	// 	- **Daily**: The policy takes effect on a daily basis.
	//
	// 	- **Weekly**: The policy takes effect on a weekly basis.
	//
	// 	- **Monthly**: The policy takes effect on a monthly basis.
	//
	// example:
	//
	// Permanent
	RepeatType *string `json:"RepeatType,omitempty" xml:"RepeatType,omitempty"`
	// The source address in the access control policy. Fuzzy match is supported. The value of this parameter depends on the value of the SourceType parameter.
	//
	// 	- If SourceType is set to `net`, the value of Source must be a CIDR block. Example: 192.0.XX.XX/24.
	//
	// 	- If SourceType is set to `group`, the value of Source must be the name of an address book. Example: db_group. If the db_group address book does not contain addresses, all source addresses are queried.
	//
	// 	- If SourceType is set to `location`, the value of Source must be a location. Example: beijing.
	//
	// >  If you do not specify this parameter, access control policies of all source address types are queried.
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
