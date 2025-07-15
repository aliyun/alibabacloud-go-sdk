// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIPv6TranslatorEntriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIpv6TranslatorEntries(v *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntries) *DescribeIPv6TranslatorEntriesResponseBody
	GetIpv6TranslatorEntries() *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntries
	SetPageNumber(v int32) *DescribeIPv6TranslatorEntriesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeIPv6TranslatorEntriesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeIPv6TranslatorEntriesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeIPv6TranslatorEntriesResponseBody
	GetTotalCount() *int32
}

type DescribeIPv6TranslatorEntriesResponseBody struct {
	// The IPv6 mapping entries that are queried.
	Ipv6TranslatorEntries *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntries `json:"Ipv6TranslatorEntries,omitempty" xml:"Ipv6TranslatorEntries,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeIPv6TranslatorEntriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIPv6TranslatorEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIPv6TranslatorEntriesResponseBody) GetIpv6TranslatorEntries() *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntries {
	return s.Ipv6TranslatorEntries
}

func (s *DescribeIPv6TranslatorEntriesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeIPv6TranslatorEntriesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeIPv6TranslatorEntriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIPv6TranslatorEntriesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeIPv6TranslatorEntriesResponseBody) SetIpv6TranslatorEntries(v *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntries) *DescribeIPv6TranslatorEntriesResponseBody {
	s.Ipv6TranslatorEntries = v
	return s
}

func (s *DescribeIPv6TranslatorEntriesResponseBody) SetPageNumber(v int32) *DescribeIPv6TranslatorEntriesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeIPv6TranslatorEntriesResponseBody) SetPageSize(v int32) *DescribeIPv6TranslatorEntriesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeIPv6TranslatorEntriesResponseBody) SetRequestId(v string) *DescribeIPv6TranslatorEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIPv6TranslatorEntriesResponseBody) SetTotalCount(v int32) *DescribeIPv6TranslatorEntriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeIPv6TranslatorEntriesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntries struct {
	Ipv6TranslatorEntry []*DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry `json:"Ipv6TranslatorEntry,omitempty" xml:"Ipv6TranslatorEntry,omitempty" type:"Repeated"`
}

func (s DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntries) String() string {
	return dara.Prettify(s)
}

func (s DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntries) GoString() string {
	return s.String()
}

func (s *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntries) GetIpv6TranslatorEntry() []*DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry {
	return s.Ipv6TranslatorEntry
}

func (s *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntries) SetIpv6TranslatorEntry(v []*DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry) *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntries {
	s.Ipv6TranslatorEntry = v
	return s
}

func (s *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntries) Validate() error {
	return dara.Validate(s)
}

type DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry struct {
	// The ID of the associated ACL.
	//
	// example:
	//
	// ipv6transacl-bp1de2****
	AclId *string `json:"AclId,omitempty" xml:"AclId,omitempty"`
	// Indicates whether ACLs are enabled.
	//
	// example:
	//
	// on
	AclStatus *string `json:"AclStatus,omitempty" xml:"AclStatus,omitempty"`
	// The ACL type.
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
	// The public IP address of the backend IPv4 server.
	//
	// example:
	//
	// 47.99.XX.XX
	BackendIpv4Addr *string `json:"BackendIpv4Addr,omitempty" xml:"BackendIpv4Addr,omitempty"`
	// The public IPv4 port used by the IPv4 server that needs to provide IPv6 access.
	//
	// example:
	//
	// 80
	BackendIpv4Port *string `json:"BackendIpv4Port,omitempty" xml:"BackendIpv4Port,omitempty"`
	// The bandwidth specified in the IPv6 mapping entry.
	//
	// example:
	//
	// 1
	EntryBandwidth *string `json:"EntryBandwidth,omitempty" xml:"EntryBandwidth,omitempty"`
	// The description of the IPv6 mapping entry.
	//
	// example:
	//
	// description
	EntryDescription *string `json:"EntryDescription,omitempty" xml:"EntryDescription,omitempty"`
	// The name of the IPv6 mapping entry.
	//
	// example:
	//
	// name
	EntryName *string `json:"EntryName,omitempty" xml:"EntryName,omitempty"`
	// The status of the IPv6 mapping entry.
	//
	// example:
	//
	// active
	EntryStatus *string `json:"EntryStatus,omitempty" xml:"EntryStatus,omitempty"`
	// The ID of the IPv6 mapping entry.
	//
	// example:
	//
	// ipv6transentry-bp1g8bhrde****
	Ipv6TranslatorEntryId *string `json:"Ipv6TranslatorEntryId,omitempty" xml:"Ipv6TranslatorEntryId,omitempty"`
	// The ID of the IPv6 Translation Service instance to which the IPv6 mapping entry belongs.
	//
	// example:
	//
	// ipv6trans-bp1858ys****
	Ipv6TranslatorId *string `json:"Ipv6TranslatorId,omitempty" xml:"Ipv6TranslatorId,omitempty"`
	// The region of the IPv6 Translation Service instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The protocol.
	//
	// example:
	//
	// tcp
	TransProtocol *string `json:"TransProtocol,omitempty" xml:"TransProtocol,omitempty"`
}

func (s DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry) String() string {
	return dara.Prettify(s)
}

func (s DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry) GoString() string {
	return s.String()
}

func (s *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry) GetAclId() *string {
	return s.AclId
}

func (s *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry) GetAclStatus() *string {
	return s.AclStatus
}

func (s *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry) GetAclType() *string {
	return s.AclType
}

func (s *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry) GetAllocateIpv6Addr() *string {
	return s.AllocateIpv6Addr
}

func (s *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry) GetAllocateIpv6Port() *int32 {
	return s.AllocateIpv6Port
}

func (s *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry) GetBackendIpv4Addr() *string {
	return s.BackendIpv4Addr
}

func (s *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry) GetBackendIpv4Port() *string {
	return s.BackendIpv4Port
}

func (s *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry) GetEntryBandwidth() *string {
	return s.EntryBandwidth
}

func (s *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry) GetEntryDescription() *string {
	return s.EntryDescription
}

func (s *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry) GetEntryName() *string {
	return s.EntryName
}

func (s *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry) GetEntryStatus() *string {
	return s.EntryStatus
}

func (s *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry) GetIpv6TranslatorEntryId() *string {
	return s.Ipv6TranslatorEntryId
}

func (s *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry) GetIpv6TranslatorId() *string {
	return s.Ipv6TranslatorId
}

func (s *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry) GetTransProtocol() *string {
	return s.TransProtocol
}

func (s *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry) SetAclId(v string) *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry {
	s.AclId = &v
	return s
}

func (s *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry) SetAclStatus(v string) *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry {
	s.AclStatus = &v
	return s
}

func (s *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry) SetAclType(v string) *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry {
	s.AclType = &v
	return s
}

func (s *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry) SetAllocateIpv6Addr(v string) *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry {
	s.AllocateIpv6Addr = &v
	return s
}

func (s *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry) SetAllocateIpv6Port(v int32) *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry {
	s.AllocateIpv6Port = &v
	return s
}

func (s *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry) SetBackendIpv4Addr(v string) *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry {
	s.BackendIpv4Addr = &v
	return s
}

func (s *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry) SetBackendIpv4Port(v string) *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry {
	s.BackendIpv4Port = &v
	return s
}

func (s *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry) SetEntryBandwidth(v string) *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry {
	s.EntryBandwidth = &v
	return s
}

func (s *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry) SetEntryDescription(v string) *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry {
	s.EntryDescription = &v
	return s
}

func (s *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry) SetEntryName(v string) *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry {
	s.EntryName = &v
	return s
}

func (s *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry) SetEntryStatus(v string) *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry {
	s.EntryStatus = &v
	return s
}

func (s *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry) SetIpv6TranslatorEntryId(v string) *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry {
	s.Ipv6TranslatorEntryId = &v
	return s
}

func (s *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry) SetIpv6TranslatorId(v string) *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry {
	s.Ipv6TranslatorId = &v
	return s
}

func (s *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry) SetRegionId(v string) *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry {
	s.RegionId = &v
	return s
}

func (s *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry) SetTransProtocol(v string) *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry {
	s.TransProtocol = &v
	return s
}

func (s *DescribeIPv6TranslatorEntriesResponseBodyIpv6TranslatorEntriesIpv6TranslatorEntry) Validate() error {
	return dara.Validate(s)
}
