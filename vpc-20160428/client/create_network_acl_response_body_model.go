// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkAclResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkAclAttribute(v *CreateNetworkAclResponseBodyNetworkAclAttribute) *CreateNetworkAclResponseBody
	GetNetworkAclAttribute() *CreateNetworkAclResponseBodyNetworkAclAttribute
	SetNetworkAclId(v string) *CreateNetworkAclResponseBody
	GetNetworkAclId() *string
	SetRequestId(v string) *CreateNetworkAclResponseBody
	GetRequestId() *string
}

type CreateNetworkAclResponseBody struct {
	// The attributes of the network ACL.
	NetworkAclAttribute *CreateNetworkAclResponseBodyNetworkAclAttribute `json:"NetworkAclAttribute,omitempty" xml:"NetworkAclAttribute,omitempty" type:"Struct"`
	// The ID of the network ACL.
	//
	// example:
	//
	// nacl-a2do9e413e0spzasx****
	NetworkAclId *string `json:"NetworkAclId,omitempty" xml:"NetworkAclId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNetworkAclResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkAclResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNetworkAclResponseBody) GetNetworkAclAttribute() *CreateNetworkAclResponseBodyNetworkAclAttribute {
	return s.NetworkAclAttribute
}

func (s *CreateNetworkAclResponseBody) GetNetworkAclId() *string {
	return s.NetworkAclId
}

func (s *CreateNetworkAclResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNetworkAclResponseBody) SetNetworkAclAttribute(v *CreateNetworkAclResponseBodyNetworkAclAttribute) *CreateNetworkAclResponseBody {
	s.NetworkAclAttribute = v
	return s
}

func (s *CreateNetworkAclResponseBody) SetNetworkAclId(v string) *CreateNetworkAclResponseBody {
	s.NetworkAclId = &v
	return s
}

func (s *CreateNetworkAclResponseBody) SetRequestId(v string) *CreateNetworkAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNetworkAclResponseBody) Validate() error {
	if s.NetworkAclAttribute != nil {
		if err := s.NetworkAclAttribute.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateNetworkAclResponseBodyNetworkAclAttribute struct {
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
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The outbound rules.
	EgressAclEntries *CreateNetworkAclResponseBodyNetworkAclAttributeEgressAclEntries `json:"EgressAclEntries,omitempty" xml:"EgressAclEntries,omitempty" type:"Struct"`
	// The inbound rules.
	IngressAclEntries *CreateNetworkAclResponseBodyNetworkAclAttributeIngressAclEntries `json:"IngressAclEntries,omitempty" xml:"IngressAclEntries,omitempty" type:"Struct"`
	// The ID of the network ACL.
	//
	// example:
	//
	// nacl-a2do9e413e0spdefr****
	NetworkAclId *string `json:"NetworkAclId,omitempty" xml:"NetworkAclId,omitempty"`
	// The name of the network ACL.
	//
	// example:
	//
	// acl-1
	NetworkAclName *string `json:"NetworkAclName,omitempty" xml:"NetworkAclName,omitempty"`
	// The region ID of the network ACL.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The information about the associated resources.
	Resources *CreateNetworkAclResponseBodyNetworkAclAttributeResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
	// The status of the network ACL. Valid values:
	//
	// 	- **Available**
	//
	// 	- **Modifying**
	//
	// example:
	//
	// Modifying
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the VPC to which the network ACL belongs.
	//
	// example:
	//
	// vpc-a2d33rfpl72k5xsscd****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateNetworkAclResponseBodyNetworkAclAttribute) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkAclResponseBodyNetworkAclAttribute) GoString() string {
	return s.String()
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttribute) GetCreationTime() *string {
	return s.CreationTime
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttribute) GetDescription() *string {
	return s.Description
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttribute) GetEgressAclEntries() *CreateNetworkAclResponseBodyNetworkAclAttributeEgressAclEntries {
	return s.EgressAclEntries
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttribute) GetIngressAclEntries() *CreateNetworkAclResponseBodyNetworkAclAttributeIngressAclEntries {
	return s.IngressAclEntries
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttribute) GetNetworkAclId() *string {
	return s.NetworkAclId
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttribute) GetNetworkAclName() *string {
	return s.NetworkAclName
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttribute) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttribute) GetResources() *CreateNetworkAclResponseBodyNetworkAclAttributeResources {
	return s.Resources
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttribute) GetStatus() *string {
	return s.Status
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttribute) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttribute) SetCreationTime(v string) *CreateNetworkAclResponseBodyNetworkAclAttribute {
	s.CreationTime = &v
	return s
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttribute) SetDescription(v string) *CreateNetworkAclResponseBodyNetworkAclAttribute {
	s.Description = &v
	return s
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttribute) SetEgressAclEntries(v *CreateNetworkAclResponseBodyNetworkAclAttributeEgressAclEntries) *CreateNetworkAclResponseBodyNetworkAclAttribute {
	s.EgressAclEntries = v
	return s
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttribute) SetIngressAclEntries(v *CreateNetworkAclResponseBodyNetworkAclAttributeIngressAclEntries) *CreateNetworkAclResponseBodyNetworkAclAttribute {
	s.IngressAclEntries = v
	return s
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttribute) SetNetworkAclId(v string) *CreateNetworkAclResponseBodyNetworkAclAttribute {
	s.NetworkAclId = &v
	return s
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttribute) SetNetworkAclName(v string) *CreateNetworkAclResponseBodyNetworkAclAttribute {
	s.NetworkAclName = &v
	return s
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttribute) SetRegionId(v string) *CreateNetworkAclResponseBodyNetworkAclAttribute {
	s.RegionId = &v
	return s
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttribute) SetResources(v *CreateNetworkAclResponseBodyNetworkAclAttributeResources) *CreateNetworkAclResponseBodyNetworkAclAttribute {
	s.Resources = v
	return s
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttribute) SetStatus(v string) *CreateNetworkAclResponseBodyNetworkAclAttribute {
	s.Status = &v
	return s
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttribute) SetVpcId(v string) *CreateNetworkAclResponseBodyNetworkAclAttribute {
	s.VpcId = &v
	return s
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttribute) Validate() error {
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
	return nil
}

type CreateNetworkAclResponseBodyNetworkAclAttributeEgressAclEntries struct {
	EgressAclEntry []*CreateNetworkAclResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry `json:"EgressAclEntry,omitempty" xml:"EgressAclEntry,omitempty" type:"Repeated"`
}

func (s CreateNetworkAclResponseBodyNetworkAclAttributeEgressAclEntries) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkAclResponseBodyNetworkAclAttributeEgressAclEntries) GoString() string {
	return s.String()
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeEgressAclEntries) GetEgressAclEntry() []*CreateNetworkAclResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry {
	return s.EgressAclEntry
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeEgressAclEntries) SetEgressAclEntry(v []*CreateNetworkAclResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry) *CreateNetworkAclResponseBodyNetworkAclAttributeEgressAclEntries {
	s.EgressAclEntry = v
	return s
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeEgressAclEntries) Validate() error {
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

type CreateNetworkAclResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry struct {
	// The description of the outbound rule.
	//
	// example:
	//
	// This is EgressAclEntries.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination CIDR block.
	//
	// example:
	//
	// 10.0.0.0/24
	DestinationCidrIp *string `json:"DestinationCidrIp,omitempty" xml:"DestinationCidrIp,omitempty"`
	// The ID of the outbound rule.
	//
	// example:
	//
	// nae-a2d447uw4tillxsdc****
	NetworkAclEntryId *string `json:"NetworkAclEntryId,omitempty" xml:"NetworkAclEntryId,omitempty"`
	// The name of the outbound rule.
	//
	// example:
	//
	// acl-2
	NetworkAclEntryName *string `json:"NetworkAclEntryName,omitempty" xml:"NetworkAclEntryName,omitempty"`
	// The action to be performed on network traffic that matches the rule. Valid values:
	//
	// 	- **accept**
	//
	// 	- **drop**
	//
	// example:
	//
	// accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The destination port range of the outbound traffic.
	//
	// 	- If the **protocol*	- of the outbound rule is set to **all**, **icmp**, or **gre**, the port range is -1/-1, which specified all ports.
	//
	// 	- If the **protocol*	- of the outbound rule is set to **tcp*	- or **udp**, set the port range in the following format: **1/200*	- or **80/80**, which specifies port 1 to port 200 or port 80. Valid values for a port: **1*	- to **65535**.
	//
	// example:
	//
	// -1/-1
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol. Valid values:
	//
	// 	- **icmp**
	//
	// 	- **gre**
	//
	// 	- **tcp**
	//
	// 	- **udp**
	//
	// 	- **all**
	//
	// example:
	//
	// all
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s CreateNetworkAclResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkAclResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry) GoString() string {
	return s.String()
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry) GetDescription() *string {
	return s.Description
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry) GetDestinationCidrIp() *string {
	return s.DestinationCidrIp
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry) GetNetworkAclEntryId() *string {
	return s.NetworkAclEntryId
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry) GetNetworkAclEntryName() *string {
	return s.NetworkAclEntryName
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry) GetPolicy() *string {
	return s.Policy
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry) GetPort() *string {
	return s.Port
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry) SetDescription(v string) *CreateNetworkAclResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry {
	s.Description = &v
	return s
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry) SetDestinationCidrIp(v string) *CreateNetworkAclResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry {
	s.DestinationCidrIp = &v
	return s
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry) SetNetworkAclEntryId(v string) *CreateNetworkAclResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry {
	s.NetworkAclEntryId = &v
	return s
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry) SetNetworkAclEntryName(v string) *CreateNetworkAclResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry {
	s.NetworkAclEntryName = &v
	return s
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry) SetPolicy(v string) *CreateNetworkAclResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry {
	s.Policy = &v
	return s
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry) SetPort(v string) *CreateNetworkAclResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry {
	s.Port = &v
	return s
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry) SetProtocol(v string) *CreateNetworkAclResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry {
	s.Protocol = &v
	return s
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeEgressAclEntriesEgressAclEntry) Validate() error {
	return dara.Validate(s)
}

type CreateNetworkAclResponseBodyNetworkAclAttributeIngressAclEntries struct {
	IngressAclEntry []*CreateNetworkAclResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry `json:"IngressAclEntry,omitempty" xml:"IngressAclEntry,omitempty" type:"Repeated"`
}

func (s CreateNetworkAclResponseBodyNetworkAclAttributeIngressAclEntries) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkAclResponseBodyNetworkAclAttributeIngressAclEntries) GoString() string {
	return s.String()
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeIngressAclEntries) GetIngressAclEntry() []*CreateNetworkAclResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry {
	return s.IngressAclEntry
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeIngressAclEntries) SetIngressAclEntry(v []*CreateNetworkAclResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry) *CreateNetworkAclResponseBodyNetworkAclAttributeIngressAclEntries {
	s.IngressAclEntry = v
	return s
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeIngressAclEntries) Validate() error {
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

type CreateNetworkAclResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry struct {
	// The description of the inbound rule.
	//
	// example:
	//
	// This is IngressAclEntries.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the inbound rule.
	//
	// example:
	//
	// nae-a2dk86arlydmexscd****
	NetworkAclEntryId *string `json:"NetworkAclEntryId,omitempty" xml:"NetworkAclEntryId,omitempty"`
	// The name of the inbound rule.
	//
	// example:
	//
	// acl-3
	NetworkAclEntryName *string `json:"NetworkAclEntryName,omitempty" xml:"NetworkAclEntryName,omitempty"`
	// The action to be performed on network traffic that matches the rule. Valid values:
	//
	// 	- **accept**
	//
	// 	- **drop**
	//
	// example:
	//
	// accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The destination port range of the inbound traffic.
	//
	// 	- If the **protocol*	- of the inbound rule is set to **all**, **icmp**, or **gre**, the port range is -1/-1, which specifies all ports.
	//
	// 	- If the **protocol*	- of the inbound rule is set to **tcp*	- or **udp**, set the port range in the following format: **1/200*	- or **80/80**, which specifies port 1 to port 200 or port 80. Valid ports: **1*	- to **65535**.
	//
	// example:
	//
	// -1/-1
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol. Valid values:
	//
	// 	- **icmp**
	//
	// 	- **gre**
	//
	// 	- **tcp**
	//
	// 	- **udp**
	//
	// 	- **all**
	//
	// example:
	//
	// all
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The source CIDR block.
	//
	// example:
	//
	// 10.0.0.0/24
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
}

func (s CreateNetworkAclResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkAclResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry) GoString() string {
	return s.String()
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry) GetDescription() *string {
	return s.Description
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry) GetNetworkAclEntryId() *string {
	return s.NetworkAclEntryId
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry) GetNetworkAclEntryName() *string {
	return s.NetworkAclEntryName
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry) GetPolicy() *string {
	return s.Policy
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry) GetPort() *string {
	return s.Port
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry) SetDescription(v string) *CreateNetworkAclResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry {
	s.Description = &v
	return s
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry) SetNetworkAclEntryId(v string) *CreateNetworkAclResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry {
	s.NetworkAclEntryId = &v
	return s
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry) SetNetworkAclEntryName(v string) *CreateNetworkAclResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry {
	s.NetworkAclEntryName = &v
	return s
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry) SetPolicy(v string) *CreateNetworkAclResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry {
	s.Policy = &v
	return s
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry) SetPort(v string) *CreateNetworkAclResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry {
	s.Port = &v
	return s
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry) SetProtocol(v string) *CreateNetworkAclResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry {
	s.Protocol = &v
	return s
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry) SetSourceCidrIp(v string) *CreateNetworkAclResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry {
	s.SourceCidrIp = &v
	return s
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeIngressAclEntriesIngressAclEntry) Validate() error {
	return dara.Validate(s)
}

type CreateNetworkAclResponseBodyNetworkAclAttributeResources struct {
	Resource []*CreateNetworkAclResponseBodyNetworkAclAttributeResourcesResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
}

func (s CreateNetworkAclResponseBodyNetworkAclAttributeResources) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkAclResponseBodyNetworkAclAttributeResources) GoString() string {
	return s.String()
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeResources) GetResource() []*CreateNetworkAclResponseBodyNetworkAclAttributeResourcesResource {
	return s.Resource
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeResources) SetResource(v []*CreateNetworkAclResponseBodyNetworkAclAttributeResourcesResource) *CreateNetworkAclResponseBodyNetworkAclAttributeResources {
	s.Resource = v
	return s
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeResources) Validate() error {
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

type CreateNetworkAclResponseBodyNetworkAclAttributeResourcesResource struct {
	// The ID of the associated resource.
	//
	// example:
	//
	// vsw-bp1de348lntdwgthy****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of resource with which you want to associate the network ACL.
	//
	// example:
	//
	// VSwitch
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The association status of the resource. Valid values:
	//
	// 	- **BINDED**
	//
	// 	- **BINDING**
	//
	// 	- **UNBINDING**
	//
	// example:
	//
	// BINDED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateNetworkAclResponseBodyNetworkAclAttributeResourcesResource) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkAclResponseBodyNetworkAclAttributeResourcesResource) GoString() string {
	return s.String()
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeResourcesResource) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeResourcesResource) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeResourcesResource) GetStatus() *string {
	return s.Status
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeResourcesResource) SetResourceId(v string) *CreateNetworkAclResponseBodyNetworkAclAttributeResourcesResource {
	s.ResourceId = &v
	return s
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeResourcesResource) SetResourceType(v string) *CreateNetworkAclResponseBodyNetworkAclAttributeResourcesResource {
	s.ResourceType = &v
	return s
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeResourcesResource) SetStatus(v string) *CreateNetworkAclResponseBodyNetworkAclAttributeResourcesResource {
	s.Status = &v
	return s
}

func (s *CreateNetworkAclResponseBodyNetworkAclAttributeResourcesResource) Validate() error {
	return dara.Validate(s)
}
