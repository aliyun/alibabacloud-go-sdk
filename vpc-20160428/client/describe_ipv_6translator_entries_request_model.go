// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIPv6TranslatorEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclId(v string) *DescribeIPv6TranslatorEntriesRequest
	GetAclId() *string
	SetAclStatus(v string) *DescribeIPv6TranslatorEntriesRequest
	GetAclStatus() *string
	SetAclType(v string) *DescribeIPv6TranslatorEntriesRequest
	GetAclType() *string
	SetAllocateIpv6Addr(v string) *DescribeIPv6TranslatorEntriesRequest
	GetAllocateIpv6Addr() *string
	SetAllocateIpv6Port(v int32) *DescribeIPv6TranslatorEntriesRequest
	GetAllocateIpv6Port() *int32
	SetBackendIpv4Addr(v string) *DescribeIPv6TranslatorEntriesRequest
	GetBackendIpv4Addr() *string
	SetBackendIpv4Port(v int32) *DescribeIPv6TranslatorEntriesRequest
	GetBackendIpv4Port() *int32
	SetClientToken(v string) *DescribeIPv6TranslatorEntriesRequest
	GetClientToken() *string
	SetEntryName(v string) *DescribeIPv6TranslatorEntriesRequest
	GetEntryName() *string
	SetIpv6TranslatorEntryId(v string) *DescribeIPv6TranslatorEntriesRequest
	GetIpv6TranslatorEntryId() *string
	SetIpv6TranslatorId(v string) *DescribeIPv6TranslatorEntriesRequest
	GetIpv6TranslatorId() *string
	SetOwnerAccount(v string) *DescribeIPv6TranslatorEntriesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeIPv6TranslatorEntriesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeIPv6TranslatorEntriesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeIPv6TranslatorEntriesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeIPv6TranslatorEntriesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeIPv6TranslatorEntriesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeIPv6TranslatorEntriesRequest
	GetResourceOwnerId() *int64
	SetTransProtocol(v string) *DescribeIPv6TranslatorEntriesRequest
	GetTransProtocol() *string
}

type DescribeIPv6TranslatorEntriesRequest struct {
	// The ID of the network ACL.
	//
	// example:
	//
	// ipv6transacl-bp1de2****
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// Specifies whether to enable access control lists (ACLs). Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// off
	AclStatus *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	// The ACL type. Valid values:
	//
	// 	- **white**: a whitelist. IPv6 addresses in the ACL are allowed to access backend services.
	//
	// 	- **black**: a blacklist. IPv6 addresses in the ACL are not allowed to access backend services.
	//
	// example:
	//
	// white
	AclType *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	// The IPv6 address allocated to the IPv6 Translation Service instance.
	//
	// example:
	//
	// 2400:3200:1600::XX
	AllocateIpv6Addr *string `json:"AllocateIpv6Addr,omitempty" xml:"AllocateIpv6Addr,omitempty"`
	// The port used by the IPv6 address allocated to the IPv6 Translation Service instance.
	//
	// example:
	//
	// 80
	AllocateIpv6Port *int32 `json:"AllocateIpv6Port,omitempty" xml:"AllocateIpv6Port,omitempty"`
	// The public IPv4 address that needs to provide IPv6 services.
	//
	// example:
	//
	// 47.99.XX.XX
	BackendIpv4Addr *string `json:"BackendIpv4Addr,omitempty" xml:"BackendIpv4Addr,omitempty"`
	// The port used by the public IPv4 address that needs to provide IPv6 services.
	//
	// example:
	//
	// 80
	BackendIpv4Port *int32 `json:"BackendIpv4Port,omitempty" xml:"BackendIpv4Port,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The name of the IPv6 mapping entry.
	//
	// example:
	//
	// entryname
	EntryName *string `json:"EntryName,omitempty" xml:"EntryName,omitempty"`
	// The ID of the IPv6 mapping entry.
	//
	// > If **Ipv6TranslatorId*	- and **Ipv6TranslatorEntryId*	- are empty, information about all IPv6 mapping entries is returned. If only **Ipv6TranslatorEntryId*	- is empty, information about the IPv6 mapping entries of the current IPv6 Translation Service instance is returned.
	//
	// example:
	//
	// ipv6transentry-bp1g8bhrde****
	Ipv6TranslatorEntryId *string `json:"Ipv6TranslatorEntryId,omitempty" xml:"Ipv6TranslatorEntryId,omitempty"`
	// The ID of the IPv6 Translation Service instance.
	//
	// example:
	//
	// ipv6trans-bp1858ysxx****
	Ipv6TranslatorId *string `json:"Ipv6TranslatorId,omitempty" xml:"Ipv6TranslatorId,omitempty"`
	OwnerAccount     *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	// The protocol used by the data to be forwarded.
	//
	// example:
	//
	// tcp
	TransProtocol *string `json:"TransProtocol,omitempty" xml:"TransProtocol,omitempty"`
}

func (s DescribeIPv6TranslatorEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIPv6TranslatorEntriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeIPv6TranslatorEntriesRequest) GetAclId() *string {
	return s.AclId
}

func (s *DescribeIPv6TranslatorEntriesRequest) GetAclStatus() *string {
	return s.AclStatus
}

func (s *DescribeIPv6TranslatorEntriesRequest) GetAclType() *string {
	return s.AclType
}

func (s *DescribeIPv6TranslatorEntriesRequest) GetAllocateIpv6Addr() *string {
	return s.AllocateIpv6Addr
}

func (s *DescribeIPv6TranslatorEntriesRequest) GetAllocateIpv6Port() *int32 {
	return s.AllocateIpv6Port
}

func (s *DescribeIPv6TranslatorEntriesRequest) GetBackendIpv4Addr() *string {
	return s.BackendIpv4Addr
}

func (s *DescribeIPv6TranslatorEntriesRequest) GetBackendIpv4Port() *int32 {
	return s.BackendIpv4Port
}

func (s *DescribeIPv6TranslatorEntriesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeIPv6TranslatorEntriesRequest) GetEntryName() *string {
	return s.EntryName
}

func (s *DescribeIPv6TranslatorEntriesRequest) GetIpv6TranslatorEntryId() *string {
	return s.Ipv6TranslatorEntryId
}

func (s *DescribeIPv6TranslatorEntriesRequest) GetIpv6TranslatorId() *string {
	return s.Ipv6TranslatorId
}

func (s *DescribeIPv6TranslatorEntriesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeIPv6TranslatorEntriesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeIPv6TranslatorEntriesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeIPv6TranslatorEntriesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeIPv6TranslatorEntriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeIPv6TranslatorEntriesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeIPv6TranslatorEntriesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeIPv6TranslatorEntriesRequest) GetTransProtocol() *string {
	return s.TransProtocol
}

func (s *DescribeIPv6TranslatorEntriesRequest) SetAclId(v string) *DescribeIPv6TranslatorEntriesRequest {
	s.AclId = &v
	return s
}

func (s *DescribeIPv6TranslatorEntriesRequest) SetAclStatus(v string) *DescribeIPv6TranslatorEntriesRequest {
	s.AclStatus = &v
	return s
}

func (s *DescribeIPv6TranslatorEntriesRequest) SetAclType(v string) *DescribeIPv6TranslatorEntriesRequest {
	s.AclType = &v
	return s
}

func (s *DescribeIPv6TranslatorEntriesRequest) SetAllocateIpv6Addr(v string) *DescribeIPv6TranslatorEntriesRequest {
	s.AllocateIpv6Addr = &v
	return s
}

func (s *DescribeIPv6TranslatorEntriesRequest) SetAllocateIpv6Port(v int32) *DescribeIPv6TranslatorEntriesRequest {
	s.AllocateIpv6Port = &v
	return s
}

func (s *DescribeIPv6TranslatorEntriesRequest) SetBackendIpv4Addr(v string) *DescribeIPv6TranslatorEntriesRequest {
	s.BackendIpv4Addr = &v
	return s
}

func (s *DescribeIPv6TranslatorEntriesRequest) SetBackendIpv4Port(v int32) *DescribeIPv6TranslatorEntriesRequest {
	s.BackendIpv4Port = &v
	return s
}

func (s *DescribeIPv6TranslatorEntriesRequest) SetClientToken(v string) *DescribeIPv6TranslatorEntriesRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeIPv6TranslatorEntriesRequest) SetEntryName(v string) *DescribeIPv6TranslatorEntriesRequest {
	s.EntryName = &v
	return s
}

func (s *DescribeIPv6TranslatorEntriesRequest) SetIpv6TranslatorEntryId(v string) *DescribeIPv6TranslatorEntriesRequest {
	s.Ipv6TranslatorEntryId = &v
	return s
}

func (s *DescribeIPv6TranslatorEntriesRequest) SetIpv6TranslatorId(v string) *DescribeIPv6TranslatorEntriesRequest {
	s.Ipv6TranslatorId = &v
	return s
}

func (s *DescribeIPv6TranslatorEntriesRequest) SetOwnerAccount(v string) *DescribeIPv6TranslatorEntriesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeIPv6TranslatorEntriesRequest) SetOwnerId(v int64) *DescribeIPv6TranslatorEntriesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeIPv6TranslatorEntriesRequest) SetPageNumber(v int32) *DescribeIPv6TranslatorEntriesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeIPv6TranslatorEntriesRequest) SetPageSize(v int32) *DescribeIPv6TranslatorEntriesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeIPv6TranslatorEntriesRequest) SetRegionId(v string) *DescribeIPv6TranslatorEntriesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeIPv6TranslatorEntriesRequest) SetResourceOwnerAccount(v string) *DescribeIPv6TranslatorEntriesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeIPv6TranslatorEntriesRequest) SetResourceOwnerId(v int64) *DescribeIPv6TranslatorEntriesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeIPv6TranslatorEntriesRequest) SetTransProtocol(v string) *DescribeIPv6TranslatorEntriesRequest {
	s.TransProtocol = &v
	return s
}

func (s *DescribeIPv6TranslatorEntriesRequest) Validate() error {
	return dara.Validate(s)
}
