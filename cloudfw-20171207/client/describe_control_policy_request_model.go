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
	// The action that Cloud Firewall performs on the traffic in the access control policy. Valid values:
	//
	// example:
	//
	// accept
	AclAction *string `json:"AclAction,omitempty" xml:"AclAction,omitempty"`
	// The unique ID of the access control policy. You must specify at least one of AclUuid and Direction. If AclUuid is specified, you can query the policy by its ID.
	//
	// example:
	//
	// 00281255-d220-4db1-8f4f-c4df221a****
	AclUuid *string `json:"AclUuid,omitempty" xml:"AclUuid,omitempty"`
	// The page number of the current page displayed in a paging query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The description of the access control policy. Fuzzy queries are supported.
	//
	// example:
	//
	// Allow access to office network segment
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination address in the access control policy. Fuzzy queries are supported. The value varies depending on the DestinationType (destination type).
	//
	// example:
	//
	// 192.0.XX.XX
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	// The traffic direction controlled by the access control policy. Valid values: in (inbound) or out (outbound). You must specify at least one of Direction and AclUuid. If AclUuid is not specified, you must specify a non-empty Direction. Otherwise, the ErrorParametersDirection error is returned.
	//
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The supported IP address version. Valid values:
	//
	// example:
	//
	// 6
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The language type for receiving messages. Valid values:
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of entries per page displayed in a paging query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The protocol type of the traffic in the access control policy. Valid values:
	//
	// example:
	//
	// TCP
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// The enabled status of the access control policy. Valid values:
	//
	// example:
	//
	// true
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The recurrence type of the policy validity period for the access control policy. Valid values:
	//
	// example:
	//
	// Permanent
	RepeatType *string `json:"RepeatType,omitempty" xml:"RepeatType,omitempty"`
	// The source address in the access control policy. Fuzzy queries are supported. The value varies depending on the SourceType (source type).
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
