// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIPv6TranslatorEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *CreateIPv6TranslatorEntryRequest
	GetAclId() *string
	SetAclStatus(v string) *CreateIPv6TranslatorEntryRequest
	GetAclStatus() *string
	SetAclType(v string) *CreateIPv6TranslatorEntryRequest
	GetAclType() *string
	SetAllocateIpv6Port(v int32) *CreateIPv6TranslatorEntryRequest
	GetAllocateIpv6Port() *int32
	SetBackendIpv4Addr(v string) *CreateIPv6TranslatorEntryRequest
	GetBackendIpv4Addr() *string
	SetBackendIpv4Port(v int32) *CreateIPv6TranslatorEntryRequest
	GetBackendIpv4Port() *int32
	SetEntryBandwidth(v int32) *CreateIPv6TranslatorEntryRequest
	GetEntryBandwidth() *int32
	SetEntryDescription(v string) *CreateIPv6TranslatorEntryRequest
	GetEntryDescription() *string
	SetEntryName(v string) *CreateIPv6TranslatorEntryRequest
	GetEntryName() *string
	SetIpv6TranslatorId(v string) *CreateIPv6TranslatorEntryRequest
	GetIpv6TranslatorId() *string
	SetOwnerAccount(v string) *CreateIPv6TranslatorEntryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateIPv6TranslatorEntryRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateIPv6TranslatorEntryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateIPv6TranslatorEntryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateIPv6TranslatorEntryRequest
	GetResourceOwnerId() *int64
	SetTransProtocol(v string) *CreateIPv6TranslatorEntryRequest
	GetTransProtocol() *string
}

type CreateIPv6TranslatorEntryRequest struct {
	// The ID of the associated ACL.
	//
	// example:
	//
	// ipv6transacl-bp1g8bhrdexnrxxxx
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// Specifies whether to enable access control lists (ACLs). Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	AclStatus *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	// Specifies whether to enable ACLs. Valid values:
	//
	// 	- **white**: a whitelist. IPv6 addresses in the ACL are allowed to access backend services.
	//
	// 	- **black**: a blacklist. IPv6 addresses in the ACL are not allowed to access backend services.
	//
	// example:
	//
	// white
	AclType *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	// The port that is used by the IPv6 address allocated to the IPv6 Translation Service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	AllocateIpv6Port *int32 `json:"AllocateIpv6Port,omitempty" xml:"AllocateIpv6Port,omitempty"`
	// The public IPv4 address that needs to provide IPv6 services.
	//
	// This parameter is required.
	//
	// example:
	//
	// 46.22.xx.xx
	BackendIpv4Addr *string `json:"BackendIpv4Addr,omitempty" xml:"BackendIpv4Addr,omitempty"`
	// The port of the public IPv4 address that needs to provide IPv6 services.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	BackendIpv4Port *int32 `json:"BackendIpv4Port,omitempty" xml:"BackendIpv4Port,omitempty"`
	// The maximum bandwidth specified in the IPv6 mapping entry. Valid values:
	//
	// 	- \\-1 (default): does not limit the maximum bandwidth specified in the IPv6 mapping entry.
	//
	// 	- 1 to 200: the bandwidth value specified in the IPv6 mapping entry. Unit: Mbit/s.
	//
	// > The sum of the maximum bandwidth values specified in all IPv6 entries cannot exceed the maximum bandwidth supported by the instance.
	//
	// example:
	//
	// 2
	EntryBandwidth *int32 `json:"EntryBandwidth,omitempty" xml:"EntryBandwidth,omitempty"`
	// The description of the IPv6 mapping entry.
	//
	// example:
	//
	// description
	EntryDescription *string `json:"EntryDescription,omitempty" xml:"EntryDescription,omitempty"`
	// The name of the IPv6 mapping entry. It must be 2 to 100 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter. It cannot start with `http://` or `https://`.
	//
	// example:
	//
	// name1
	EntryName *string `json:"EntryName,omitempty" xml:"EntryName,omitempty"`
	// The ID of the IPv6 Translation Service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipv6trans-bp1858ys57xxxxxx
	Ipv6TranslatorId *string `json:"Ipv6TranslatorId,omitempty" xml:"Ipv6TranslatorId,omitempty"`
	OwnerAccount     *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region of the IPv6 Translation Service instance. You can call the **DescribeRegions*	- operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The protocol. Valid values:
	//
	// 	- **tcp**
	//
	// 	- **udp**
	//
	// This parameter is required.
	//
	// example:
	//
	// tcp
	TransProtocol *string `json:"TransProtocol,omitempty" xml:"TransProtocol,omitempty"`
}

func (s CreateIPv6TranslatorEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIPv6TranslatorEntryRequest) GoString() string {
	return s.String()
}

func (s *CreateIPv6TranslatorEntryRequest) GetAclId() *string {
	return s.AclId
}

func (s *CreateIPv6TranslatorEntryRequest) GetAclStatus() *string {
	return s.AclStatus
}

func (s *CreateIPv6TranslatorEntryRequest) GetAclType() *string {
	return s.AclType
}

func (s *CreateIPv6TranslatorEntryRequest) GetAllocateIpv6Port() *int32 {
	return s.AllocateIpv6Port
}

func (s *CreateIPv6TranslatorEntryRequest) GetBackendIpv4Addr() *string {
	return s.BackendIpv4Addr
}

func (s *CreateIPv6TranslatorEntryRequest) GetBackendIpv4Port() *int32 {
	return s.BackendIpv4Port
}

func (s *CreateIPv6TranslatorEntryRequest) GetEntryBandwidth() *int32 {
	return s.EntryBandwidth
}

func (s *CreateIPv6TranslatorEntryRequest) GetEntryDescription() *string {
	return s.EntryDescription
}

func (s *CreateIPv6TranslatorEntryRequest) GetEntryName() *string {
	return s.EntryName
}

func (s *CreateIPv6TranslatorEntryRequest) GetIpv6TranslatorId() *string {
	return s.Ipv6TranslatorId
}

func (s *CreateIPv6TranslatorEntryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateIPv6TranslatorEntryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateIPv6TranslatorEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateIPv6TranslatorEntryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateIPv6TranslatorEntryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateIPv6TranslatorEntryRequest) GetTransProtocol() *string {
	return s.TransProtocol
}

func (s *CreateIPv6TranslatorEntryRequest) SetAclId(v string) *CreateIPv6TranslatorEntryRequest {
	s.AclId = &v
	return s
}

func (s *CreateIPv6TranslatorEntryRequest) SetAclStatus(v string) *CreateIPv6TranslatorEntryRequest {
	s.AclStatus = &v
	return s
}

func (s *CreateIPv6TranslatorEntryRequest) SetAclType(v string) *CreateIPv6TranslatorEntryRequest {
	s.AclType = &v
	return s
}

func (s *CreateIPv6TranslatorEntryRequest) SetAllocateIpv6Port(v int32) *CreateIPv6TranslatorEntryRequest {
	s.AllocateIpv6Port = &v
	return s
}

func (s *CreateIPv6TranslatorEntryRequest) SetBackendIpv4Addr(v string) *CreateIPv6TranslatorEntryRequest {
	s.BackendIpv4Addr = &v
	return s
}

func (s *CreateIPv6TranslatorEntryRequest) SetBackendIpv4Port(v int32) *CreateIPv6TranslatorEntryRequest {
	s.BackendIpv4Port = &v
	return s
}

func (s *CreateIPv6TranslatorEntryRequest) SetEntryBandwidth(v int32) *CreateIPv6TranslatorEntryRequest {
	s.EntryBandwidth = &v
	return s
}

func (s *CreateIPv6TranslatorEntryRequest) SetEntryDescription(v string) *CreateIPv6TranslatorEntryRequest {
	s.EntryDescription = &v
	return s
}

func (s *CreateIPv6TranslatorEntryRequest) SetEntryName(v string) *CreateIPv6TranslatorEntryRequest {
	s.EntryName = &v
	return s
}

func (s *CreateIPv6TranslatorEntryRequest) SetIpv6TranslatorId(v string) *CreateIPv6TranslatorEntryRequest {
	s.Ipv6TranslatorId = &v
	return s
}

func (s *CreateIPv6TranslatorEntryRequest) SetOwnerAccount(v string) *CreateIPv6TranslatorEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateIPv6TranslatorEntryRequest) SetOwnerId(v int64) *CreateIPv6TranslatorEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateIPv6TranslatorEntryRequest) SetRegionId(v string) *CreateIPv6TranslatorEntryRequest {
	s.RegionId = &v
	return s
}

func (s *CreateIPv6TranslatorEntryRequest) SetResourceOwnerAccount(v string) *CreateIPv6TranslatorEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateIPv6TranslatorEntryRequest) SetResourceOwnerId(v int64) *CreateIPv6TranslatorEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateIPv6TranslatorEntryRequest) SetTransProtocol(v string) *CreateIPv6TranslatorEntryRequest {
	s.TransProtocol = &v
	return s
}

func (s *CreateIPv6TranslatorEntryRequest) Validate() error {
	return dara.Validate(s)
}
