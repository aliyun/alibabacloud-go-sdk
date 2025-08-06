// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAuditSecurityIpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIps(v string) *SetAuditSecurityIpRequest
	GetIps() *string
	SetOperateMode(v string) *SetAuditSecurityIpRequest
	GetOperateMode() *string
	SetSecurityGroupName(v string) *SetAuditSecurityIpRequest
	GetSecurityGroupName() *string
}

type SetAuditSecurityIpRequest struct {
	// The IP addresses that you want to add to the review security group. You can add a maximum of 100 IP addresses to a review security group. Separate multiple IP addresses with commas (,). You can add IP addresses in the following formats to review security groups:
	//
	// 	- IP address: 192.168.0.1
	//
	// 	- CIDR block: 192.168.0.1/24. /24 indicates that the prefix of the CIDR block is 24 bits in length. You can replace 24 with a value that ranges `from 1 to 32`.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.0.1
	Ips *string `json:"Ips,omitempty" xml:"Ips,omitempty"`
	// The operation type. Valid values:
	//
	// 	- **Append*	- (default): adds the IP addresses to the original whitelist.
	//
	// 	- **Cover**: overwrites the original whitelist.
	//
	// 	- **Delete**: removes the IP addresses from the original whitelist.
	//
	// >  If the value that you specify is invalid, the default value is used.
	//
	// example:
	//
	// Cover
	OperateMode *string `json:"OperateMode,omitempty" xml:"OperateMode,omitempty"`
	// The name of the review security group. Default value: **Default**. You can specify a maximum of 10 review security groups.
	//
	// example:
	//
	// Default
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
}

func (s SetAuditSecurityIpRequest) String() string {
	return dara.Prettify(s)
}

func (s SetAuditSecurityIpRequest) GoString() string {
	return s.String()
}

func (s *SetAuditSecurityIpRequest) GetIps() *string {
	return s.Ips
}

func (s *SetAuditSecurityIpRequest) GetOperateMode() *string {
	return s.OperateMode
}

func (s *SetAuditSecurityIpRequest) GetSecurityGroupName() *string {
	return s.SecurityGroupName
}

func (s *SetAuditSecurityIpRequest) SetIps(v string) *SetAuditSecurityIpRequest {
	s.Ips = &v
	return s
}

func (s *SetAuditSecurityIpRequest) SetOperateMode(v string) *SetAuditSecurityIpRequest {
	s.OperateMode = &v
	return s
}

func (s *SetAuditSecurityIpRequest) SetSecurityGroupName(v string) *SetAuditSecurityIpRequest {
	s.SecurityGroupName = &v
	return s
}

func (s *SetAuditSecurityIpRequest) Validate() error {
	return dara.Validate(s)
}
