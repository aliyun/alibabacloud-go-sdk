// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstanceSecurityGroupRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *CreateDBInstanceSecurityGroupRuleRequest
	GetDBInstanceId() *string
	SetDescription(v string) *CreateDBInstanceSecurityGroupRuleRequest
	GetDescription() *string
	SetIpProtocol(v string) *CreateDBInstanceSecurityGroupRuleRequest
	GetIpProtocol() *string
	SetOwnerAccount(v string) *CreateDBInstanceSecurityGroupRuleRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *CreateDBInstanceSecurityGroupRuleRequest
	GetOwnerId() *string
	SetPortRange(v string) *CreateDBInstanceSecurityGroupRuleRequest
	GetPortRange() *string
	SetResourceOwnerAccount(v string) *CreateDBInstanceSecurityGroupRuleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDBInstanceSecurityGroupRuleRequest
	GetResourceOwnerId() *int64
	SetSourceCidrIp(v string) *CreateDBInstanceSecurityGroupRuleRequest
	GetSourceCidrIp() *string
}

type CreateDBInstanceSecurityGroupRuleRequest struct {
	// The ID of the instance. You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/2628785.html) operation to query the IDs of instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp15i4hn07r******
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The description of the security group rule.
	//
	// example:
	//
	// zht_test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The type of the transport layer protocol. Valid values:
	//
	// 	- TCP
	//
	// 	- UDP
	//
	// example:
	//
	// TCP
	IpProtocol   *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The range of destination ports over which TCP and UDP traffic is allowed in the security group rule.
	//
	// Valid values: 1 to 65535. Separate the start port number and the end port number with a forward slash (/). Example: 1/200.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1/200
	PortRange            *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The range of source IP addresses. CIDR blocks and IPv4 addresses are supported.
	//
	// example:
	//
	// 192.XX.XX.100
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
}

func (s CreateDBInstanceSecurityGroupRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceSecurityGroupRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceSecurityGroupRuleRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateDBInstanceSecurityGroupRuleRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDBInstanceSecurityGroupRuleRequest) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *CreateDBInstanceSecurityGroupRuleRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateDBInstanceSecurityGroupRuleRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *CreateDBInstanceSecurityGroupRuleRequest) GetPortRange() *string {
	return s.PortRange
}

func (s *CreateDBInstanceSecurityGroupRuleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateDBInstanceSecurityGroupRuleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDBInstanceSecurityGroupRuleRequest) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *CreateDBInstanceSecurityGroupRuleRequest) SetDBInstanceId(v string) *CreateDBInstanceSecurityGroupRuleRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDBInstanceSecurityGroupRuleRequest) SetDescription(v string) *CreateDBInstanceSecurityGroupRuleRequest {
	s.Description = &v
	return s
}

func (s *CreateDBInstanceSecurityGroupRuleRequest) SetIpProtocol(v string) *CreateDBInstanceSecurityGroupRuleRequest {
	s.IpProtocol = &v
	return s
}

func (s *CreateDBInstanceSecurityGroupRuleRequest) SetOwnerAccount(v string) *CreateDBInstanceSecurityGroupRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDBInstanceSecurityGroupRuleRequest) SetOwnerId(v string) *CreateDBInstanceSecurityGroupRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBInstanceSecurityGroupRuleRequest) SetPortRange(v string) *CreateDBInstanceSecurityGroupRuleRequest {
	s.PortRange = &v
	return s
}

func (s *CreateDBInstanceSecurityGroupRuleRequest) SetResourceOwnerAccount(v string) *CreateDBInstanceSecurityGroupRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDBInstanceSecurityGroupRuleRequest) SetResourceOwnerId(v int64) *CreateDBInstanceSecurityGroupRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDBInstanceSecurityGroupRuleRequest) SetSourceCidrIp(v string) *CreateDBInstanceSecurityGroupRuleRequest {
	s.SourceCidrIp = &v
	return s
}

func (s *CreateDBInstanceSecurityGroupRuleRequest) Validate() error {
	return dara.Validate(s)
}
