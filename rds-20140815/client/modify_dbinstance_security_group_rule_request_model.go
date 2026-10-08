// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceSecurityGroupRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyDBInstanceSecurityGroupRuleRequest
	GetDBInstanceId() *string
	SetDescription(v string) *ModifyDBInstanceSecurityGroupRuleRequest
	GetDescription() *string
	SetIpProtocol(v string) *ModifyDBInstanceSecurityGroupRuleRequest
	GetIpProtocol() *string
	SetOwnerAccount(v string) *ModifyDBInstanceSecurityGroupRuleRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *ModifyDBInstanceSecurityGroupRuleRequest
	GetOwnerId() *string
	SetPortRange(v string) *ModifyDBInstanceSecurityGroupRuleRequest
	GetPortRange() *string
	SetResourceOwnerAccount(v string) *ModifyDBInstanceSecurityGroupRuleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceSecurityGroupRuleRequest
	GetResourceOwnerId() *int64
	SetSecurityGroupRuleId(v string) *ModifyDBInstanceSecurityGroupRuleRequest
	GetSecurityGroupRuleId() *string
	SetSourceCidrIp(v string) *ModifyDBInstanceSecurityGroupRuleRequest
	GetSourceCidrIp() *string
}

type ModifyDBInstanceSecurityGroupRuleRequest struct {
	// The instance ID. You can call [DescribeDBInstances](https://help.aliyun.com/document_detail/2628785.html) to obtain the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp15i4hn07r******
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The description of the security group rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// zht_test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The transport layer protocol type. Valid values:
	//
	// - TCP
	//
	// - UDP
	//
	// This parameter is required.
	//
	// example:
	//
	// TCP
	IpProtocol   *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The range of destination ports for the transport layer protocol (TCP/UDP) that the security group opens.
	//
	// Valid values: 1 to 65535. Separate the start port and end port with a forward slash (/). Example: 1/200.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1/200
	PortRange            *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The security group rule ID. You can call [DescribeDBInstanceSecurityGroupRule](https://help.aliyun.com/document_detail/2834044.html) to obtain the security group rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sgr-2ze17u******
	SecurityGroupRuleId *string `json:"SecurityGroupRuleId,omitempty" xml:"SecurityGroupRuleId,omitempty"`
	// The source IP address range. CIDR format and IPv4 format are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.XX.XX.100
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
}

func (s ModifyDBInstanceSecurityGroupRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceSecurityGroupRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceSecurityGroupRuleRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceSecurityGroupRuleRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyDBInstanceSecurityGroupRuleRequest) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *ModifyDBInstanceSecurityGroupRuleRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBInstanceSecurityGroupRuleRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ModifyDBInstanceSecurityGroupRuleRequest) GetPortRange() *string {
	return s.PortRange
}

func (s *ModifyDBInstanceSecurityGroupRuleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBInstanceSecurityGroupRuleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBInstanceSecurityGroupRuleRequest) GetSecurityGroupRuleId() *string {
	return s.SecurityGroupRuleId
}

func (s *ModifyDBInstanceSecurityGroupRuleRequest) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *ModifyDBInstanceSecurityGroupRuleRequest) SetDBInstanceId(v string) *ModifyDBInstanceSecurityGroupRuleRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceSecurityGroupRuleRequest) SetDescription(v string) *ModifyDBInstanceSecurityGroupRuleRequest {
	s.Description = &v
	return s
}

func (s *ModifyDBInstanceSecurityGroupRuleRequest) SetIpProtocol(v string) *ModifyDBInstanceSecurityGroupRuleRequest {
	s.IpProtocol = &v
	return s
}

func (s *ModifyDBInstanceSecurityGroupRuleRequest) SetOwnerAccount(v string) *ModifyDBInstanceSecurityGroupRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceSecurityGroupRuleRequest) SetOwnerId(v string) *ModifyDBInstanceSecurityGroupRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceSecurityGroupRuleRequest) SetPortRange(v string) *ModifyDBInstanceSecurityGroupRuleRequest {
	s.PortRange = &v
	return s
}

func (s *ModifyDBInstanceSecurityGroupRuleRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceSecurityGroupRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceSecurityGroupRuleRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceSecurityGroupRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceSecurityGroupRuleRequest) SetSecurityGroupRuleId(v string) *ModifyDBInstanceSecurityGroupRuleRequest {
	s.SecurityGroupRuleId = &v
	return s
}

func (s *ModifyDBInstanceSecurityGroupRuleRequest) SetSourceCidrIp(v string) *ModifyDBInstanceSecurityGroupRuleRequest {
	s.SourceCidrIp = &v
	return s
}

func (s *ModifyDBInstanceSecurityGroupRuleRequest) Validate() error {
	return dara.Validate(s)
}
