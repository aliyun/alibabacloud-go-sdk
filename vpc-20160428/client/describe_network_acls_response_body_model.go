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
	// The details of the network ACLs.
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
	return dara.Validate(s)
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
	return dara.Validate(s)
}

type DescribeNetworkAclsResponseBodyNetworkAclsNetworkAcl struct {
	// The time when the network ACL was created.
	//
	// example:
	//
	// 2021-12-25 11:44:17
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the network ACL.
	//
	// example:
	//
	// This is my NetworkAcl.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The outbound rules.
	EgressAclEntries *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntries `json:"EgressAclEntries,omitempty" xml:"EgressAclEntries,omitempty" type:"Struct"`
	// The configurations of the inbound rules.
	IngressAclEntries *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntries `json:"IngressAclEntries,omitempty" xml:"IngressAclEntries,omitempty" type:"Struct"`
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
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resources that are associated with the network ACL.
	Resources *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
	// The status of the network ACL. Valid values:
	//
	// 	- **Available**
	//
	// 	- **Modifying**
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The information about the tags.
	Tags *DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The ID of the associated VPC.
	//
	// example:
	//
	// vpc-m5ebpc2xh64mqm27e****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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
	return dara.Validate(s)
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
	return dara.Validate(s)
}

type DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclEgressAclEntriesEgressAclEntry struct {
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
	// The type of the inbound rule.
	//
	// - **custom**
	//
	// - **system**
	//
	// example:
	//
	// custom
	EntryType *string `json:"EntryType,omitempty" xml:"EntryType,omitempty"`
	// The IP version.
	//
	// 	- **IPv4**
	//
	// 	- **IPv6**
	//
	// example:
	//
	// IPV4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
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
	return dara.Validate(s)
}

type DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclIngressAclEntriesIngressAclEntry struct {
	// The description of the inbound rule.
	//
	// example:
	//
	// This is IngressAclEntries.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The type of the inbound rule.
	//
	// - **custom**
	//
	// - **system**
	//
	// example:
	//
	// custom
	EntryType *string `json:"EntryType,omitempty" xml:"EntryType,omitempty"`
	// The IP version.
	//
	// 	- **IPv4**
	//
	// 	- **IPv6**
	//
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The ID of the inbound rule.
	//
	// example:
	//
	// nae-a2dk86arlydmezasw****
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
	return dara.Validate(s)
}

type DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclResourcesResource struct {
	// The ID of the associated resource.
	//
	// example:
	//
	// vsw-bp1de348lntdwcdf****
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
	return dara.Validate(s)
}

type DescribeNetworkAclsResponseBodyNetworkAclsNetworkAclTagsTag struct {
	// The key of tag N added to the resource.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N added to the resource.
	//
	// example:
	//
	// FinanceJoshua
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
