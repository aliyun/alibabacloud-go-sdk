// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNetworkAclEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateNetworkAclEntriesRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateNetworkAclEntriesRequest
	GetDryRun() *bool
	SetEgressAclEntries(v []*UpdateNetworkAclEntriesRequestEgressAclEntries) *UpdateNetworkAclEntriesRequest
	GetEgressAclEntries() []*UpdateNetworkAclEntriesRequestEgressAclEntries
	SetIngressAclEntries(v []*UpdateNetworkAclEntriesRequestIngressAclEntries) *UpdateNetworkAclEntriesRequest
	GetIngressAclEntries() []*UpdateNetworkAclEntriesRequestIngressAclEntries
	SetNetworkAclId(v string) *UpdateNetworkAclEntriesRequest
	GetNetworkAclId() *string
	SetOwnerAccount(v string) *UpdateNetworkAclEntriesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateNetworkAclEntriesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateNetworkAclEntriesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UpdateNetworkAclEntriesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateNetworkAclEntriesRequest
	GetResourceOwnerId() *int64
	SetUpdateEgressAclEntries(v bool) *UpdateNetworkAclEntriesRequest
	GetUpdateEgressAclEntries() *bool
	SetUpdateIngressAclEntries(v bool) *UpdateNetworkAclEntriesRequest
	GetUpdateIngressAclEntries() *bool
}

type UpdateNetworkAclEntriesRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs a dry run. The system checks the request for potential issues, including the AccessKey pair, the permissions of the RAM user, and the required parameters. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The information about the outbound rules.
	EgressAclEntries []*UpdateNetworkAclEntriesRequestEgressAclEntries `json:"EgressAclEntries,omitempty" xml:"EgressAclEntries,omitempty" type:"Repeated"`
	// The information about the inbound rule.
	IngressAclEntries []*UpdateNetworkAclEntriesRequestIngressAclEntries `json:"IngressAclEntries,omitempty" xml:"IngressAclEntries,omitempty" type:"Repeated"`
	// The ID of the network ACL.
	//
	// This parameter is required.
	//
	// example:
	//
	// nacl-bp1lhl0taikrzxsc****
	NetworkAclId *string `json:"NetworkAclId,omitempty" xml:"NetworkAclId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the network ACL.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to update outbound rules. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// >  This parameter cannot be used to add outbound rules to ACLs. If you want to add more outbound rules to ACLs, specify both the existing rule and the rule that you want to add when you call this API operation. If you specify only the rule that you want to add, it overwrites the existing rule.
	//
	// example:
	//
	// false
	UpdateEgressAclEntries *bool `json:"UpdateEgressAclEntries,omitempty" xml:"UpdateEgressAclEntries,omitempty"`
	// Specifies whether to update inbound rules. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// >  This parameter cannot be used to add inbound rules to ACLs. If you want to add more inbound rules to ACLs, you must specify both the existing rule and the rule that you want to add when you call this API operation. If you specify only the rule that you want to add, it overwrites the existing rule.
	//
	// example:
	//
	// false
	UpdateIngressAclEntries *bool `json:"UpdateIngressAclEntries,omitempty" xml:"UpdateIngressAclEntries,omitempty"`
}

func (s UpdateNetworkAclEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNetworkAclEntriesRequest) GoString() string {
	return s.String()
}

func (s *UpdateNetworkAclEntriesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateNetworkAclEntriesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateNetworkAclEntriesRequest) GetEgressAclEntries() []*UpdateNetworkAclEntriesRequestEgressAclEntries {
	return s.EgressAclEntries
}

func (s *UpdateNetworkAclEntriesRequest) GetIngressAclEntries() []*UpdateNetworkAclEntriesRequestIngressAclEntries {
	return s.IngressAclEntries
}

func (s *UpdateNetworkAclEntriesRequest) GetNetworkAclId() *string {
	return s.NetworkAclId
}

func (s *UpdateNetworkAclEntriesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateNetworkAclEntriesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateNetworkAclEntriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateNetworkAclEntriesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateNetworkAclEntriesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateNetworkAclEntriesRequest) GetUpdateEgressAclEntries() *bool {
	return s.UpdateEgressAclEntries
}

func (s *UpdateNetworkAclEntriesRequest) GetUpdateIngressAclEntries() *bool {
	return s.UpdateIngressAclEntries
}

func (s *UpdateNetworkAclEntriesRequest) SetClientToken(v string) *UpdateNetworkAclEntriesRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateNetworkAclEntriesRequest) SetDryRun(v bool) *UpdateNetworkAclEntriesRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateNetworkAclEntriesRequest) SetEgressAclEntries(v []*UpdateNetworkAclEntriesRequestEgressAclEntries) *UpdateNetworkAclEntriesRequest {
	s.EgressAclEntries = v
	return s
}

func (s *UpdateNetworkAclEntriesRequest) SetIngressAclEntries(v []*UpdateNetworkAclEntriesRequestIngressAclEntries) *UpdateNetworkAclEntriesRequest {
	s.IngressAclEntries = v
	return s
}

func (s *UpdateNetworkAclEntriesRequest) SetNetworkAclId(v string) *UpdateNetworkAclEntriesRequest {
	s.NetworkAclId = &v
	return s
}

func (s *UpdateNetworkAclEntriesRequest) SetOwnerAccount(v string) *UpdateNetworkAclEntriesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateNetworkAclEntriesRequest) SetOwnerId(v int64) *UpdateNetworkAclEntriesRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateNetworkAclEntriesRequest) SetRegionId(v string) *UpdateNetworkAclEntriesRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateNetworkAclEntriesRequest) SetResourceOwnerAccount(v string) *UpdateNetworkAclEntriesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateNetworkAclEntriesRequest) SetResourceOwnerId(v int64) *UpdateNetworkAclEntriesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateNetworkAclEntriesRequest) SetUpdateEgressAclEntries(v bool) *UpdateNetworkAclEntriesRequest {
	s.UpdateEgressAclEntries = &v
	return s
}

func (s *UpdateNetworkAclEntriesRequest) SetUpdateIngressAclEntries(v bool) *UpdateNetworkAclEntriesRequest {
	s.UpdateIngressAclEntries = &v
	return s
}

func (s *UpdateNetworkAclEntriesRequest) Validate() error {
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
	return nil
}

type UpdateNetworkAclEntriesRequestEgressAclEntries struct {
	// The description of the outbound rule.
	//
	// The description must be 1 to 256 characters in length, and cannot start with `http://` or `https://`.
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
	// The type of the rule. Set the value to **custom**, which specifies custom rules.
	//
	// example:
	//
	// custom
	EntryType *string `json:"EntryType,omitempty" xml:"EntryType,omitempty"`
	// The IP version. Valid values:
	//
	// 	- **IPv4*	- (default)
	//
	// 	- **IPv6**
	//
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The ID of the outbound rule.
	//
	// Valid values of **N**: **0*	- to **99**. You can specify at most 100 outbound rules.
	//
	// example:
	//
	// nae-2zecs97e0brcge46****
	NetworkAclEntryId *string `json:"NetworkAclEntryId,omitempty" xml:"NetworkAclEntryId,omitempty"`
	// The name of the outbound rule.
	//
	// The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
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

func (s UpdateNetworkAclEntriesRequestEgressAclEntries) String() string {
	return dara.Prettify(s)
}

func (s UpdateNetworkAclEntriesRequestEgressAclEntries) GoString() string {
	return s.String()
}

func (s *UpdateNetworkAclEntriesRequestEgressAclEntries) GetDescription() *string {
	return s.Description
}

func (s *UpdateNetworkAclEntriesRequestEgressAclEntries) GetDestinationCidrIp() *string {
	return s.DestinationCidrIp
}

func (s *UpdateNetworkAclEntriesRequestEgressAclEntries) GetEntryType() *string {
	return s.EntryType
}

func (s *UpdateNetworkAclEntriesRequestEgressAclEntries) GetIpVersion() *string {
	return s.IpVersion
}

func (s *UpdateNetworkAclEntriesRequestEgressAclEntries) GetNetworkAclEntryId() *string {
	return s.NetworkAclEntryId
}

func (s *UpdateNetworkAclEntriesRequestEgressAclEntries) GetNetworkAclEntryName() *string {
	return s.NetworkAclEntryName
}

func (s *UpdateNetworkAclEntriesRequestEgressAclEntries) GetPolicy() *string {
	return s.Policy
}

func (s *UpdateNetworkAclEntriesRequestEgressAclEntries) GetPort() *string {
	return s.Port
}

func (s *UpdateNetworkAclEntriesRequestEgressAclEntries) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateNetworkAclEntriesRequestEgressAclEntries) SetDescription(v string) *UpdateNetworkAclEntriesRequestEgressAclEntries {
	s.Description = &v
	return s
}

func (s *UpdateNetworkAclEntriesRequestEgressAclEntries) SetDestinationCidrIp(v string) *UpdateNetworkAclEntriesRequestEgressAclEntries {
	s.DestinationCidrIp = &v
	return s
}

func (s *UpdateNetworkAclEntriesRequestEgressAclEntries) SetEntryType(v string) *UpdateNetworkAclEntriesRequestEgressAclEntries {
	s.EntryType = &v
	return s
}

func (s *UpdateNetworkAclEntriesRequestEgressAclEntries) SetIpVersion(v string) *UpdateNetworkAclEntriesRequestEgressAclEntries {
	s.IpVersion = &v
	return s
}

func (s *UpdateNetworkAclEntriesRequestEgressAclEntries) SetNetworkAclEntryId(v string) *UpdateNetworkAclEntriesRequestEgressAclEntries {
	s.NetworkAclEntryId = &v
	return s
}

func (s *UpdateNetworkAclEntriesRequestEgressAclEntries) SetNetworkAclEntryName(v string) *UpdateNetworkAclEntriesRequestEgressAclEntries {
	s.NetworkAclEntryName = &v
	return s
}

func (s *UpdateNetworkAclEntriesRequestEgressAclEntries) SetPolicy(v string) *UpdateNetworkAclEntriesRequestEgressAclEntries {
	s.Policy = &v
	return s
}

func (s *UpdateNetworkAclEntriesRequestEgressAclEntries) SetPort(v string) *UpdateNetworkAclEntriesRequestEgressAclEntries {
	s.Port = &v
	return s
}

func (s *UpdateNetworkAclEntriesRequestEgressAclEntries) SetProtocol(v string) *UpdateNetworkAclEntriesRequestEgressAclEntries {
	s.Protocol = &v
	return s
}

func (s *UpdateNetworkAclEntriesRequestEgressAclEntries) Validate() error {
	return dara.Validate(s)
}

type UpdateNetworkAclEntriesRequestIngressAclEntries struct {
	// The description of the inbound rule.
	//
	// The description must be 1 to 256 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// This is IngressAclEntries.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The type of the rule. Set the value to **custom**, which specifies custom rules.
	//
	// example:
	//
	// custom
	EntryType *string `json:"EntryType,omitempty" xml:"EntryType,omitempty"`
	// The IP version. Valid values:
	//
	// 	- **IPv4*	- (default)
	//
	// 	- **IPv6**
	//
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The ID of the inbound rule.
	//
	// Valid values of **N**: **0*	- to **99**. You can specify at most 100 inbound rules.
	//
	// example:
	//
	// nae-2zepn32de59j8m4****
	NetworkAclEntryId *string `json:"NetworkAclEntryId,omitempty" xml:"NetworkAclEntryId,omitempty"`
	// The name of the inbound rule.
	//
	// The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
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
	// The source port range of the inbound rule.
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

func (s UpdateNetworkAclEntriesRequestIngressAclEntries) String() string {
	return dara.Prettify(s)
}

func (s UpdateNetworkAclEntriesRequestIngressAclEntries) GoString() string {
	return s.String()
}

func (s *UpdateNetworkAclEntriesRequestIngressAclEntries) GetDescription() *string {
	return s.Description
}

func (s *UpdateNetworkAclEntriesRequestIngressAclEntries) GetEntryType() *string {
	return s.EntryType
}

func (s *UpdateNetworkAclEntriesRequestIngressAclEntries) GetIpVersion() *string {
	return s.IpVersion
}

func (s *UpdateNetworkAclEntriesRequestIngressAclEntries) GetNetworkAclEntryId() *string {
	return s.NetworkAclEntryId
}

func (s *UpdateNetworkAclEntriesRequestIngressAclEntries) GetNetworkAclEntryName() *string {
	return s.NetworkAclEntryName
}

func (s *UpdateNetworkAclEntriesRequestIngressAclEntries) GetPolicy() *string {
	return s.Policy
}

func (s *UpdateNetworkAclEntriesRequestIngressAclEntries) GetPort() *string {
	return s.Port
}

func (s *UpdateNetworkAclEntriesRequestIngressAclEntries) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateNetworkAclEntriesRequestIngressAclEntries) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *UpdateNetworkAclEntriesRequestIngressAclEntries) SetDescription(v string) *UpdateNetworkAclEntriesRequestIngressAclEntries {
	s.Description = &v
	return s
}

func (s *UpdateNetworkAclEntriesRequestIngressAclEntries) SetEntryType(v string) *UpdateNetworkAclEntriesRequestIngressAclEntries {
	s.EntryType = &v
	return s
}

func (s *UpdateNetworkAclEntriesRequestIngressAclEntries) SetIpVersion(v string) *UpdateNetworkAclEntriesRequestIngressAclEntries {
	s.IpVersion = &v
	return s
}

func (s *UpdateNetworkAclEntriesRequestIngressAclEntries) SetNetworkAclEntryId(v string) *UpdateNetworkAclEntriesRequestIngressAclEntries {
	s.NetworkAclEntryId = &v
	return s
}

func (s *UpdateNetworkAclEntriesRequestIngressAclEntries) SetNetworkAclEntryName(v string) *UpdateNetworkAclEntriesRequestIngressAclEntries {
	s.NetworkAclEntryName = &v
	return s
}

func (s *UpdateNetworkAclEntriesRequestIngressAclEntries) SetPolicy(v string) *UpdateNetworkAclEntriesRequestIngressAclEntries {
	s.Policy = &v
	return s
}

func (s *UpdateNetworkAclEntriesRequestIngressAclEntries) SetPort(v string) *UpdateNetworkAclEntriesRequestIngressAclEntries {
	s.Port = &v
	return s
}

func (s *UpdateNetworkAclEntriesRequestIngressAclEntries) SetProtocol(v string) *UpdateNetworkAclEntriesRequestIngressAclEntries {
	s.Protocol = &v
	return s
}

func (s *UpdateNetworkAclEntriesRequestIngressAclEntries) SetSourceCidrIp(v string) *UpdateNetworkAclEntriesRequestIngressAclEntries {
	s.SourceCidrIp = &v
	return s
}

func (s *UpdateNetworkAclEntriesRequestIngressAclEntries) Validate() error {
	return dara.Validate(s)
}
