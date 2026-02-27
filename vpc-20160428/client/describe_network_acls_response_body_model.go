// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkAclsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkAcls(v *DescribeNetworkAclsResponseBodyNetworkAcls) *DescribeNetworkAclsResponseBody
	GetNetworkAcls() *DescribeNetworkAclsResponseBodyNetworkAcls
	SetPageNumber(v string) *DescribeNetworkAclsResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *DescribeNetworkAclsResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeNetworkAclsResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeNetworkAclsResponseBody
	GetTotalCount() *string
}

type DescribeNetworkAclsResponseBody struct {
	NetworkAcls *DescribeNetworkAclsResponseBodyNetworkAcls `json:"NetworkAcls,omitempty" xml:"NetworkAcls,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F7DDDC17-FA06-4AC2-8F35-59D2470FCFC1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeNetworkAclsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAclsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAclsResponseBody) GetNetworkAcls() *DescribeNetworkAclsResponseBodyNetworkAcls {
	return s.NetworkAcls
}

func (s *DescribeNetworkAclsResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeNetworkAclsResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeNetworkAclsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNetworkAclsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeNetworkAclsResponseBody) SetNetworkAcls(v *DescribeNetworkAclsResponseBodyNetworkAcls) *DescribeNetworkAclsResponseBody {
	s.NetworkAcls = v
	return s
}

func (s *DescribeNetworkAclsResponseBody) SetPageNumber(v string) *DescribeNetworkAclsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeNetworkAclsResponseBody) SetPageSize(v string) *DescribeNetworkAclsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeNetworkAclsResponseBody) SetRequestId(v string) *DescribeNetworkAclsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNetworkAclsResponseBody) SetTotalCount(v string) *DescribeNetworkAclsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeNetworkAclsResponseBody) Validate() error {
	if s.NetworkAcls != nil {
		if err := s.NetworkAcls.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeNetworkAclsResponseBodyNetworkAcls struct {
	NetworkAcl []*DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl `json:"NetworkAcl,omitempty" xml:"NetworkAcl,omitempty" type:"Repeated"`
}

func (s DescribeNetworkAclsResponseBodyNetworkAcls) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAclsResponseBodyNetworkAcls) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAclsResponseBodyNetworkAcls) GetNetworkAcl() []*DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl {
	return s.NetworkAcl
}

func (s *DescribeNetworkAclsResponseBodyNetworkAcls) SetNetworkAcl(v []*DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl) *DescribeNetworkAclsResponseBodyNetworkAcls {
	s.NetworkAcl = v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAcls) Validate() error {
	if s.NetworkAcl != nil {
		for _, item := range s.NetworkAcl {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl struct {
	CreationTime      *string                                                                `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description       *string                                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	EgressAclEntries  *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntries  `json:"EgressAclEntries,omitempty" xml:"EgressAclEntries,omitempty" type:"Struct"`
	IngressAclEntries *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntries `json:"IngressAclEntries,omitempty" xml:"IngressAclEntries,omitempty" type:"Struct"`
	NetworkAclId      *string                                                                `json:"NetworkAclId,omitempty" xml:"NetworkAclId,omitempty"`
	NetworkAclName    *string                                                                `json:"NetworkAclName,omitempty" xml:"NetworkAclName,omitempty"`
	OwnerId           *int64                                                                 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId          *string                                                                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Resources         *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclResources         `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
	Status            *string                                                                `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags              *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclTags              `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	VpcId             *string                                                                `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl) GetDescription() *string {
	return s.Description
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl) GetEgressAclEntries() *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntries {
	return s.EgressAclEntries
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl) GetIngressAclEntries() *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntries {
	return s.IngressAclEntries
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl) GetNetworkAclId() *string {
	return s.NetworkAclId
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl) GetNetworkAclName() *string {
	return s.NetworkAclName
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl) GetResources() *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclResources {
	return s.Resources
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl) GetStatus() *string {
	return s.Status
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl) GetTags() *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclTags {
	return s.Tags
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl) SetCreationTime(v string) *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl {
	s.CreationTime = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl) SetDescription(v string) *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl {
	s.Description = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl) SetEgressAclEntries(v *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntries) *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl {
	s.EgressAclEntries = v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl) SetIngressAclEntries(v *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntries) *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl {
	s.IngressAclEntries = v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl) SetNetworkAclId(v string) *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl {
	s.NetworkAclId = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl) SetNetworkAclName(v string) *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl {
	s.NetworkAclName = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl) SetOwnerId(v int64) *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl {
	s.OwnerId = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl) SetRegionId(v string) *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl {
	s.RegionId = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl) SetResources(v *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclResources) *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl {
	s.Resources = v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl) SetStatus(v string) *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl {
	s.Status = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl) SetTags(v *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclTags) *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl {
	s.Tags = v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl) SetVpcId(v string) *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl {
	s.VpcId = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl) Validate() error {
	if s.EgressAclEntries != nil {
		if err := s.EgressAclEntries.Validate(); err != nil {
			return err
		}
	}
	if s.IngressAclEntries != nil {
		if err := s.IngressAclEntries.Validate(); err != nil {
			return err
		}
	}
	if s.Resources != nil {
		if err := s.Resources.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntries struct {
	EgressAclEntry []*DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntriesEgressAclEntry `json:"EgressAclEntry,omitempty" xml:"EgressAclEntry,omitempty" type:"Repeated"`
}

func (s DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntries) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntries) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntries) GetEgressAclEntry() []*DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntriesEgressAclEntry {
	return s.EgressAclEntry
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntries) SetEgressAclEntry(v []*DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntriesEgressAclEntry) *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntries {
	s.EgressAclEntry = v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntries) Validate() error {
	if s.EgressAclEntry != nil {
		for _, item := range s.EgressAclEntry {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntriesEgressAclEntry struct {
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DestinationCidrIp   *string `json:"DestinationCidrIp,omitempty" xml:"DestinationCidrIp,omitempty"`
	EntryType           *string `json:"EntryType,omitempty" xml:"EntryType,omitempty"`
	IpVersion           *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	NetworkAclEntryId   *string `json:"NetworkAclEntryId,omitempty" xml:"NetworkAclEntryId,omitempty"`
	NetworkAclEntryName *string `json:"NetworkAclEntryName,omitempty" xml:"NetworkAclEntryName,omitempty"`
	Policy              *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Port                *string `json:"Port,omitempty" xml:"Port,omitempty"`
	Protocol            *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntriesEgressAclEntry) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntriesEgressAclEntry) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntriesEgressAclEntry) GetDescription() *string {
	return s.Description
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntriesEgressAclEntry) GetDestinationCidrIp() *string {
	return s.DestinationCidrIp
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntriesEgressAclEntry) GetEntryType() *string {
	return s.EntryType
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntriesEgressAclEntry) GetIpVersion() *string {
	return s.IpVersion
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntriesEgressAclEntry) GetNetworkAclEntryId() *string {
	return s.NetworkAclEntryId
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntriesEgressAclEntry) GetNetworkAclEntryName() *string {
	return s.NetworkAclEntryName
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntriesEgressAclEntry) GetPolicy() *string {
	return s.Policy
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntriesEgressAclEntry) GetPort() *string {
	return s.Port
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntriesEgressAclEntry) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntriesEgressAclEntry) SetDescription(v string) *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntriesEgressAclEntry {
	s.Description = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntriesEgressAclEntry) SetDestinationCidrIp(v string) *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntriesEgressAclEntry {
	s.DestinationCidrIp = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntriesEgressAclEntry) SetEntryType(v string) *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntriesEgressAclEntry {
	s.EntryType = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntriesEgressAclEntry) SetIpVersion(v string) *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntriesEgressAclEntry {
	s.IpVersion = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntriesEgressAclEntry) SetNetworkAclEntryId(v string) *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntriesEgressAclEntry {
	s.NetworkAclEntryId = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntriesEgressAclEntry) SetNetworkAclEntryName(v string) *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntriesEgressAclEntry {
	s.NetworkAclEntryName = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntriesEgressAclEntry) SetPolicy(v string) *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntriesEgressAclEntry {
	s.Policy = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntriesEgressAclEntry) SetPort(v string) *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntriesEgressAclEntry {
	s.Port = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntriesEgressAclEntry) SetProtocol(v string) *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntriesEgressAclEntry {
	s.Protocol = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntriesEgressAclEntry) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntries struct {
	IngressAclEntry []*DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntriesIngressAclEntry `json:"IngressAclEntry,omitempty" xml:"IngressAclEntry,omitempty" type:"Repeated"`
}

func (s DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntries) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntries) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntries) GetIngressAclEntry() []*DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntriesIngressAclEntry {
	return s.IngressAclEntry
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntries) SetIngressAclEntry(v []*DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntriesIngressAclEntry) *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntries {
	s.IngressAclEntry = v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntries) Validate() error {
	if s.IngressAclEntry != nil {
		for _, item := range s.IngressAclEntry {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntriesIngressAclEntry struct {
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EntryType           *string `json:"EntryType,omitempty" xml:"EntryType,omitempty"`
	IpVersion           *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	NetworkAclEntryId   *string `json:"NetworkAclEntryId,omitempty" xml:"NetworkAclEntryId,omitempty"`
	NetworkAclEntryName *string `json:"NetworkAclEntryName,omitempty" xml:"NetworkAclEntryName,omitempty"`
	Policy              *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Port                *string `json:"Port,omitempty" xml:"Port,omitempty"`
	Protocol            *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	SourceCidrIp        *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
}

func (s DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntriesIngressAclEntry) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntriesIngressAclEntry) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntriesIngressAclEntry) GetDescription() *string {
	return s.Description
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntriesIngressAclEntry) GetEntryType() *string {
	return s.EntryType
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntriesIngressAclEntry) GetIpVersion() *string {
	return s.IpVersion
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntriesIngressAclEntry) GetNetworkAclEntryId() *string {
	return s.NetworkAclEntryId
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntriesIngressAclEntry) GetNetworkAclEntryName() *string {
	return s.NetworkAclEntryName
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntriesIngressAclEntry) GetPolicy() *string {
	return s.Policy
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntriesIngressAclEntry) GetPort() *string {
	return s.Port
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntriesIngressAclEntry) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntriesIngressAclEntry) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntriesIngressAclEntry) SetDescription(v string) *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntriesIngressAclEntry {
	s.Description = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntriesIngressAclEntry) SetEntryType(v string) *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntriesIngressAclEntry {
	s.EntryType = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntriesIngressAclEntry) SetIpVersion(v string) *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntriesIngressAclEntry {
	s.IpVersion = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntriesIngressAclEntry) SetNetworkAclEntryId(v string) *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntriesIngressAclEntry {
	s.NetworkAclEntryId = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntriesIngressAclEntry) SetNetworkAclEntryName(v string) *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntriesIngressAclEntry {
	s.NetworkAclEntryName = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntriesIngressAclEntry) SetPolicy(v string) *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntriesIngressAclEntry {
	s.Policy = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntriesIngressAclEntry) SetPort(v string) *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntriesIngressAclEntry {
	s.Port = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntriesIngressAclEntry) SetProtocol(v string) *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntriesIngressAclEntry {
	s.Protocol = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntriesIngressAclEntry) SetSourceCidrIp(v string) *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntriesIngressAclEntry {
	s.SourceCidrIp = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntriesIngressAclEntry) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclResources struct {
	Resource []*DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclResourcesResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
}

func (s DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclResources) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclResources) GetResource() []*DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclResourcesResource {
	return s.Resource
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclResources) SetResource(v []*DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclResourcesResource) *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclResources {
	s.Resource = v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclResources) Validate() error {
	if s.Resource != nil {
		for _, item := range s.Resource {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclResourcesResource struct {
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclResourcesResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclResourcesResource) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclResourcesResource) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclResourcesResource) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclResourcesResource) GetStatus() *string {
	return s.Status
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclResourcesResource) SetResourceId(v string) *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclResourcesResource {
	s.ResourceId = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclResourcesResource) SetResourceType(v string) *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclResourcesResource {
	s.ResourceType = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclResourcesResource) SetStatus(v string) *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclResourcesResource {
	s.Status = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclResourcesResource) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclTags struct {
	Tag []*DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclTags) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclTags) GetTag() []*DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclTagsTag {
	return s.Tag
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclTags) SetTag(v []*DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclTagsTag) *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclTags {
	s.Tag = v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclTags) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclTagsTag) SetKey(v string) *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclTagsTag) SetValue(v string) *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclTagsTag) Validate() error {
	return dara.Validate(s)
}
