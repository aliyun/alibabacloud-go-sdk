// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkAclAttributesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkAclAttribute(v *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute) *DescribeNetworkAclAttributesResponseBody
	GetNetworkAclAttribute() *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute
	SetRequestId(v string) *DescribeNetworkAclAttributesResponseBody
	GetRequestId() *string
}

type DescribeNetworkAclAttributesResponseBody struct {
	// The details of the network ACLs.
	NetworkAclAttribute *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute `json:"NetworkAclAttribute,omitempty" xml:"NetworkAclAttribute,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F5905F9C-0161-4E72-9CB1-1F3F3CF6268A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNetworkAclAttributesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAclAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAclAttributesResponseBody) GetNetworkAclAttribute() *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute {
	return s.NetworkAclAttribute
}

func (s *DescribeNetworkAclAttributesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNetworkAclAttributesResponseBody) SetNetworkAclAttribute(v *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute) *DescribeNetworkAclAttributesResponseBody {
	s.NetworkAclAttribute = v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBody) SetRequestId(v string) *DescribeNetworkAclAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBody) Validate() error {
	if s.NetworkAclAttribute != nil {
		if err := s.NetworkAclAttribute.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute struct {
	// The time when the network ACL was created.
	//
	// example:
	//
	// 2021-12-25 11:33:27
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the network ACL.
	//
	// example:
	//
	// This is my NetworkAcl.
	Description       *string                                                                       `json:"Description,omitempty" xml:"Description,omitempty"`
	EgressAclEntries  *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntries  `json:"EgressAclEntries,omitempty" xml:"EgressAclEntries,omitempty" type:"Struct"`
	IngressAclEntries *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntries `json:"IngressAclEntries,omitempty" xml:"IngressAclEntries,omitempty" type:"Struct"`
	// The ID of the network ACL.
	//
	// example:
	//
	// nacl-a2do9e413e0spnhmj****
	NetworkAclId *string `json:"NetworkAclId,omitempty" xml:"NetworkAclId,omitempty"`
	// The name of the network ACL.
	//
	// example:
	//
	// acl-1
	NetworkAclName *string `json:"NetworkAclName,omitempty" xml:"NetworkAclName,omitempty"`
	// The ID of the Alibaba Cloud account to which the network ACL belongs.
	//
	// example:
	//
	// 253460731706911258
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the network ACL.
	//
	// example:
	//
	// cn-hangzhou
	RegionId  *string                                                               `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Resources *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
	// The association status of the resource. Valid values:
	//
	// 	- **Available**
	//
	// 	- **Modifying**
	//
	// example:
	//
	// Available
	Status *string                                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The ID of the VPC to which the network ACL belongs.
	//
	// example:
	//
	// vpc-a2d33rfpl72k5defr****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute) GetDescription() *string {
	return s.Description
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute) GetEgressAclEntries() *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntries {
	return s.EgressAclEntries
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute) GetIngressAclEntries() *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntries {
	return s.IngressAclEntries
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute) GetNetworkAclId() *string {
	return s.NetworkAclId
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute) GetNetworkAclName() *string {
	return s.NetworkAclName
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute) GetResources() *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeResources {
	return s.Resources
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute) GetStatus() *string {
	return s.Status
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute) GetTags() *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeTags {
	return s.Tags
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute) SetCreationTime(v string) *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute {
	s.CreationTime = &v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute) SetDescription(v string) *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute {
	s.Description = &v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute) SetEgressAclEntries(v *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntries) *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute {
	s.EgressAclEntries = v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute) SetIngressAclEntries(v *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntries) *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute {
	s.IngressAclEntries = v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute) SetNetworkAclId(v string) *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute {
	s.NetworkAclId = &v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute) SetNetworkAclName(v string) *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute {
	s.NetworkAclName = &v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute) SetOwnerId(v int64) *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute {
	s.OwnerId = &v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute) SetRegionId(v string) *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute {
	s.RegionId = &v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute) SetResources(v *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeResources) *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute {
	s.Resources = v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute) SetStatus(v string) *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute {
	s.Status = &v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute) SetTags(v *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeTags) *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute {
	s.Tags = v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute) SetVpcId(v string) *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute {
	s.VpcId = &v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttribute) Validate() error {
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

type DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntries struct {
	EgressAclEntry []*DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry `json:"EgressAclEntry,omitempty" xml:"EgressAclEntry,omitempty" type:"Repeated"`
}

func (s DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntries) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntries) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntries) GetEgressAclEntry() []*DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry {
	return s.EgressAclEntry
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntries) SetEgressAclEntry(v []*DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry) *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntries {
	s.EgressAclEntry = v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntries) Validate() error {
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

type DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry struct {
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

func (s DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry) GetDescription() *string {
	return s.Description
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry) GetDestinationCidrIp() *string {
	return s.DestinationCidrIp
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry) GetEntryType() *string {
	return s.EntryType
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry) GetIpVersion() *string {
	return s.IpVersion
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry) GetNetworkAclEntryId() *string {
	return s.NetworkAclEntryId
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry) GetNetworkAclEntryName() *string {
	return s.NetworkAclEntryName
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry) GetPolicy() *string {
	return s.Policy
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry) GetPort() *string {
	return s.Port
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry) SetDescription(v string) *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry {
	s.Description = &v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry) SetDestinationCidrIp(v string) *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry {
	s.DestinationCidrIp = &v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry) SetEntryType(v string) *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry {
	s.EntryType = &v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry) SetIpVersion(v string) *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry {
	s.IpVersion = &v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry) SetNetworkAclEntryId(v string) *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry {
	s.NetworkAclEntryId = &v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry) SetNetworkAclEntryName(v string) *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry {
	s.NetworkAclEntryName = &v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry) SetPolicy(v string) *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry {
	s.Policy = &v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry) SetPort(v string) *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry {
	s.Port = &v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry) SetProtocol(v string) *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry {
	s.Protocol = &v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntries struct {
	IngressAclEntry []*DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry `json:"IngressAclEntry,omitempty" xml:"IngressAclEntry,omitempty" type:"Repeated"`
}

func (s DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntries) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntries) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntries) GetIngressAclEntry() []*DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry {
	return s.IngressAclEntry
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntries) SetIngressAclEntry(v []*DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry) *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntries {
	s.IngressAclEntry = v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntries) Validate() error {
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

type DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry struct {
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

func (s DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry) GetDescription() *string {
	return s.Description
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry) GetEntryType() *string {
	return s.EntryType
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry) GetIpVersion() *string {
	return s.IpVersion
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry) GetNetworkAclEntryId() *string {
	return s.NetworkAclEntryId
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry) GetNetworkAclEntryName() *string {
	return s.NetworkAclEntryName
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry) GetPolicy() *string {
	return s.Policy
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry) GetPort() *string {
	return s.Port
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry) SetDescription(v string) *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry {
	s.Description = &v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry) SetEntryType(v string) *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry {
	s.EntryType = &v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry) SetIpVersion(v string) *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry {
	s.IpVersion = &v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry) SetNetworkAclEntryId(v string) *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry {
	s.NetworkAclEntryId = &v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry) SetNetworkAclEntryName(v string) *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry {
	s.NetworkAclEntryName = &v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry) SetPolicy(v string) *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry {
	s.Policy = &v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry) SetPort(v string) *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry {
	s.Port = &v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry) SetProtocol(v string) *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry {
	s.Protocol = &v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry) SetSourceCidrIp(v string) *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry {
	s.SourceCidrIp = &v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeResources struct {
	Resource []*DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeResourcesResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
}

func (s DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeResources) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeResources) GetResource() []*DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeResourcesResource {
	return s.Resource
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeResources) SetResource(v []*DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeResourcesResource) *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeResources {
	s.Resource = v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeResources) Validate() error {
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

type DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeResourcesResource struct {
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeResourcesResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeResourcesResource) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeResourcesResource) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeResourcesResource) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeResourcesResource) GetStatus() *string {
	return s.Status
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeResourcesResource) SetResourceId(v string) *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeResourcesResource {
	s.ResourceId = &v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeResourcesResource) SetResourceType(v string) *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeResourcesResource {
	s.ResourceType = &v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeResourcesResource) SetStatus(v string) *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeResourcesResource {
	s.Status = &v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeResourcesResource) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeTags struct {
	Tag []*DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeTags) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeTags) GetTag() []*DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeTagsTag {
	return s.Tag
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeTags) SetTag(v []*DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeTagsTag) *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeTags {
	s.Tag = v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeTags) Validate() error {
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

type DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeTagsTag) SetKey(v string) *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeTagsTag) SetValue(v string) *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeNetworkAclAttributesResponseBodyNetworkAclAttributeTagsTag) Validate() error {
	return dara.Validate(s)
}
