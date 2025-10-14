// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkAclsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkAcls(v []*DescribeNetworkAclsResponseBodyNetworkAcls) *DescribeNetworkAclsResponseBody
	GetNetworkAcls() []*DescribeNetworkAclsResponseBodyNetworkAcls
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
	// Details of the network ACLs.
	NetworkAcls []*DescribeNetworkAclsResponseBodyNetworkAcls `json:"NetworkAcls,omitempty" xml:"NetworkAcls,omitempty" type:"Repeated"`
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
	// 2
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A1707FC0-430C-423A-B624-284046B20399
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 3
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeNetworkAclsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAclsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAclsResponseBody) GetNetworkAcls() []*DescribeNetworkAclsResponseBodyNetworkAcls {
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

func (s *DescribeNetworkAclsResponseBody) SetNetworkAcls(v []*DescribeNetworkAclsResponseBodyNetworkAcls) *DescribeNetworkAclsResponseBody {
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
		for _, item := range s.NetworkAcls {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNetworkAclsResponseBodyNetworkAcls struct {
	// The time when the network ACL was created. The time follows the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-11-01T06:08:46Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the network ACL.
	//
	// example:
	//
	// This is my NetworkAcl.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Details of the outbound rules.
	EgressAclEntries []*DescribeNetworkAclsResponseBodyNetworkAclsEgressAclEntries `json:"EgressAclEntries,omitempty" xml:"EgressAclEntries,omitempty" type:"Repeated"`
	// Details of the inbound rules.
	IngressAclEntries []*DescribeNetworkAclsResponseBodyNetworkAclsIngressAclEntries `json:"IngressAclEntries,omitempty" xml:"IngressAclEntries,omitempty" type:"Repeated"`
	// The ID of the network ACL.
	//
	// example:
	//
	// nacl-a2do9e413e0spxscd****
	NetworkAclId *string `json:"NetworkAclId,omitempty" xml:"NetworkAclId,omitempty"`
	// The name of the network ACL.
	//
	// example:
	//
	// acl-8
	NetworkAclName *string `json:"NetworkAclName,omitempty" xml:"NetworkAclName,omitempty"`
	// Details of the associated resources.
	Resources []*DescribeNetworkAclsResponseBodyNetworkAclsResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// The status of the network ACL. Valid values:
	//
	// 	- **Available**: The network ACL is available.
	//
	// 	- **Modifying**: The network ACL is being configured.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeNetworkAclsResponseBodyNetworkAcls) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAclsResponseBodyNetworkAcls) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAclsResponseBodyNetworkAcls) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeNetworkAclsResponseBodyNetworkAcls) GetDescription() *string {
	return s.Description
}

func (s *DescribeNetworkAclsResponseBodyNetworkAcls) GetEgressAclEntries() []*DescribeNetworkAclsResponseBodyNetworkAclsEgressAclEntries {
	return s.EgressAclEntries
}

func (s *DescribeNetworkAclsResponseBodyNetworkAcls) GetIngressAclEntries() []*DescribeNetworkAclsResponseBodyNetworkAclsIngressAclEntries {
	return s.IngressAclEntries
}

func (s *DescribeNetworkAclsResponseBodyNetworkAcls) GetNetworkAclId() *string {
	return s.NetworkAclId
}

func (s *DescribeNetworkAclsResponseBodyNetworkAcls) GetNetworkAclName() *string {
	return s.NetworkAclName
}

func (s *DescribeNetworkAclsResponseBodyNetworkAcls) GetResources() []*DescribeNetworkAclsResponseBodyNetworkAclsResources {
	return s.Resources
}

func (s *DescribeNetworkAclsResponseBodyNetworkAcls) GetStatus() *string {
	return s.Status
}

func (s *DescribeNetworkAclsResponseBodyNetworkAcls) SetCreationTime(v string) *DescribeNetworkAclsResponseBodyNetworkAcls {
	s.CreationTime = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAcls) SetDescription(v string) *DescribeNetworkAclsResponseBodyNetworkAcls {
	s.Description = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAcls) SetEgressAclEntries(v []*DescribeNetworkAclsResponseBodyNetworkAclsEgressAclEntries) *DescribeNetworkAclsResponseBodyNetworkAcls {
	s.EgressAclEntries = v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAcls) SetIngressAclEntries(v []*DescribeNetworkAclsResponseBodyNetworkAclsIngressAclEntries) *DescribeNetworkAclsResponseBodyNetworkAcls {
	s.IngressAclEntries = v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAcls) SetNetworkAclId(v string) *DescribeNetworkAclsResponseBodyNetworkAcls {
	s.NetworkAclId = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAcls) SetNetworkAclName(v string) *DescribeNetworkAclsResponseBodyNetworkAcls {
	s.NetworkAclName = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAcls) SetResources(v []*DescribeNetworkAclsResponseBodyNetworkAclsResources) *DescribeNetworkAclsResponseBodyNetworkAcls {
	s.Resources = v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAcls) SetStatus(v string) *DescribeNetworkAclsResponseBodyNetworkAcls {
	s.Status = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAcls) Validate() error {
	if s.EgressAclEntries != nil {
		for _, item := range s.EgressAclEntries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.IngressAclEntries != nil {
		for _, item := range s.IngressAclEntries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Resources != nil {
		for _, item := range s.Resources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNetworkAclsResponseBodyNetworkAclsEgressAclEntries struct {
	// The destination CIDR block.
	//
	// example:
	//
	// 10.0.0.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The description of the outbound rule.
	//
	// example:
	//
	// This is EgressAclEntries.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the outbound rule.
	//
	// example:
	//
	// nae-a2d447uw4tillfvgb****
	NetworkAclEntryId *string `json:"NetworkAclEntryId,omitempty" xml:"NetworkAclEntryId,omitempty"`
	// The name of the outbound rule.
	//
	// example:
	//
	// acl-2
	NetworkAclEntryName *string `json:"NetworkAclEntryName,omitempty" xml:"NetworkAclEntryName,omitempty"`
	// The action that is performed on network traffic that matches the rule. Valid values:
	//
	// 	- **accept**: allows the network traffic.
	//
	// 	- **drop**: blocks the network traffic.
	//
	// example:
	//
	// accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The destination port range of the outbound rule.
	//
	// 	- If **Protocol*	- of the outbound rule is set to **all*	- or **icmp*	- the port range is **-1/-1**, which indicates all ports.
	//
	// 	- If **Protocol*	- of the outbound rule is set to **tcp*	- or **udp**, the port range is in the following format: **1/200*	- or **80/80**. 1/200 indicates port 1 to port 200. 80/80 indicates port 80. Valid values for a port: **1 to 65535**.
	//
	// example:
	//
	// -1/-1
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// The priority of the rule. Valid values: **1 to 100**. Default value: **1**.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The protocol type. Valid values:
	//
	// 	- **icmp**: ICMP.
	//
	// 	- **tcp**: TCP.
	//
	// 	- **udp**: UDP.
	//
	// 	- **all**: all protocols.
	//
	// example:
	//
	// all
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The type of the rule. Valid values:
	//
	// 	- **system**: The rule is created by the system.
	//
	// 	- **custom**: The rule is created by a user.
	//
	// example:
	//
	// system
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeNetworkAclsResponseBodyNetworkAclsEgressAclEntries) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAclsResponseBodyNetworkAclsEgressAclEntries) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsEgressAclEntries) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsEgressAclEntries) GetDescription() *string {
	return s.Description
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsEgressAclEntries) GetNetworkAclEntryId() *string {
	return s.NetworkAclEntryId
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsEgressAclEntries) GetNetworkAclEntryName() *string {
	return s.NetworkAclEntryName
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsEgressAclEntries) GetPolicy() *string {
	return s.Policy
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsEgressAclEntries) GetPortRange() *string {
	return s.PortRange
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsEgressAclEntries) GetPriority() *int32 {
	return s.Priority
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsEgressAclEntries) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsEgressAclEntries) GetType() *string {
	return s.Type
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsEgressAclEntries) SetCidrBlock(v string) *DescribeNetworkAclsResponseBodyNetworkAclsEgressAclEntries {
	s.CidrBlock = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsEgressAclEntries) SetDescription(v string) *DescribeNetworkAclsResponseBodyNetworkAclsEgressAclEntries {
	s.Description = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsEgressAclEntries) SetNetworkAclEntryId(v string) *DescribeNetworkAclsResponseBodyNetworkAclsEgressAclEntries {
	s.NetworkAclEntryId = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsEgressAclEntries) SetNetworkAclEntryName(v string) *DescribeNetworkAclsResponseBodyNetworkAclsEgressAclEntries {
	s.NetworkAclEntryName = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsEgressAclEntries) SetPolicy(v string) *DescribeNetworkAclsResponseBodyNetworkAclsEgressAclEntries {
	s.Policy = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsEgressAclEntries) SetPortRange(v string) *DescribeNetworkAclsResponseBodyNetworkAclsEgressAclEntries {
	s.PortRange = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsEgressAclEntries) SetPriority(v int32) *DescribeNetworkAclsResponseBodyNetworkAclsEgressAclEntries {
	s.Priority = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsEgressAclEntries) SetProtocol(v string) *DescribeNetworkAclsResponseBodyNetworkAclsEgressAclEntries {
	s.Protocol = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsEgressAclEntries) SetType(v string) *DescribeNetworkAclsResponseBodyNetworkAclsEgressAclEntries {
	s.Type = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsEgressAclEntries) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkAclsResponseBodyNetworkAclsIngressAclEntries struct {
	// The source CIDR block.
	//
	// example:
	//
	// 10.0.0.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The description of the inbound rule.
	//
	// example:
	//
	// This is IngressAclEntries.
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	// The ID of the inbound rule.
	//
	// example:
	//
	// nae-5dk86arlydmezasw****
	NetworkAclEntryId *string `json:"NetworkAclEntryId,omitempty" xml:"NetworkAclEntryId,omitempty"`
	// The name of the inbound rule.
	//
	// example:
	//
	// acl-3
	NetworkAclEntryName *string `json:"NetworkAclEntryName,omitempty" xml:"NetworkAclEntryName,omitempty"`
	// The action that is performed on network traffic that matches the rule. Valid values:
	//
	// 	- **accept**: allows the network traffic.
	//
	// 	- **drop**: blocks the network traffic.
	//
	// example:
	//
	// accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The destination port range of the inbound rule.
	//
	// 	- If **Protocol*	- of the inbound rule is set to **all*	- or **icmp**, the port range is **-1/-1**, which indicates all ports.
	//
	// 	- If **Protocol*	- of the inbound rule is set to **tcp*	- or **udp**, the port range is in the following format: **1/200*	- or **80/80**. 1/200 indicates port 1 to port 200. 80/80 indicates port 80. Valid values for a port: **1 to 65535**.
	//
	// example:
	//
	// -1/-1
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// The priority of the rule. Valid values: **1 to 100**. Default value: **1**.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The protocol type. Valid values:
	//
	// 	- **icmp**: ICMP.
	//
	// 	- **tcp**: TCP.
	//
	// 	- **udp**: UDP.
	//
	// 	- **all**: all protocols.
	//
	// example:
	//
	// all
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The type of the rule. Valid values:
	//
	// 	- **system**: The rule is created by the system.
	//
	// 	- **custom**: The rule is created by a user.
	//
	// example:
	//
	// system
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeNetworkAclsResponseBodyNetworkAclsIngressAclEntries) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAclsResponseBodyNetworkAclsIngressAclEntries) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsIngressAclEntries) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsIngressAclEntries) GetDescription() *string {
	return s.Description
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsIngressAclEntries) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsIngressAclEntries) GetNetworkAclEntryId() *string {
	return s.NetworkAclEntryId
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsIngressAclEntries) GetNetworkAclEntryName() *string {
	return s.NetworkAclEntryName
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsIngressAclEntries) GetPolicy() *string {
	return s.Policy
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsIngressAclEntries) GetPortRange() *string {
	return s.PortRange
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsIngressAclEntries) GetPriority() *int32 {
	return s.Priority
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsIngressAclEntries) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsIngressAclEntries) GetType() *string {
	return s.Type
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsIngressAclEntries) SetCidrBlock(v string) *DescribeNetworkAclsResponseBodyNetworkAclsIngressAclEntries {
	s.CidrBlock = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsIngressAclEntries) SetDescription(v string) *DescribeNetworkAclsResponseBodyNetworkAclsIngressAclEntries {
	s.Description = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsIngressAclEntries) SetDestinationCidrBlock(v string) *DescribeNetworkAclsResponseBodyNetworkAclsIngressAclEntries {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsIngressAclEntries) SetNetworkAclEntryId(v string) *DescribeNetworkAclsResponseBodyNetworkAclsIngressAclEntries {
	s.NetworkAclEntryId = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsIngressAclEntries) SetNetworkAclEntryName(v string) *DescribeNetworkAclsResponseBodyNetworkAclsIngressAclEntries {
	s.NetworkAclEntryName = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsIngressAclEntries) SetPolicy(v string) *DescribeNetworkAclsResponseBodyNetworkAclsIngressAclEntries {
	s.Policy = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsIngressAclEntries) SetPortRange(v string) *DescribeNetworkAclsResponseBodyNetworkAclsIngressAclEntries {
	s.PortRange = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsIngressAclEntries) SetPriority(v int32) *DescribeNetworkAclsResponseBodyNetworkAclsIngressAclEntries {
	s.Priority = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsIngressAclEntries) SetProtocol(v string) *DescribeNetworkAclsResponseBodyNetworkAclsIngressAclEntries {
	s.Protocol = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsIngressAclEntries) SetType(v string) *DescribeNetworkAclsResponseBodyNetworkAclsIngressAclEntries {
	s.Type = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsIngressAclEntries) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkAclsResponseBodyNetworkAclsResources struct {
	// The ID of the edge node.
	//
	// example:
	//
	// cn-fuzhou-telecom
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The ID of the associated resource.
	//
	// example:
	//
	// n-****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the associated resource.
	//
	// example:
	//
	// Network
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The association status of the resource. Valid values:
	//
	// 	- **BINDED**: The resource is associated with the network ACL.
	//
	// 	- **BINDING**: The resource is being associated with the network ACL.
	//
	// 	- **UNBINDING**: The resource is being disassociated from the network ACL.
	//
	// example:
	//
	// BINDING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeNetworkAclsResponseBodyNetworkAclsResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAclsResponseBodyNetworkAclsResources) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsResources) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsResources) GetStatus() *string {
	return s.Status
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsResources) SetEnsRegionId(v string) *DescribeNetworkAclsResponseBodyNetworkAclsResources {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsResources) SetResourceId(v string) *DescribeNetworkAclsResponseBodyNetworkAclsResources {
	s.ResourceId = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsResources) SetResourceType(v string) *DescribeNetworkAclsResponseBodyNetworkAclsResources {
	s.ResourceType = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsResources) SetStatus(v string) *DescribeNetworkAclsResponseBodyNetworkAclsResources {
	s.Status = &v
	return s
}

func (s *DescribeNetworkAclsResponseBodyNetworkAclsResources) Validate() error {
	return dara.Validate(s)
}
