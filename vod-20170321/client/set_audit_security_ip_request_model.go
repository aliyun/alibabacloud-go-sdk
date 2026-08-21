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
	// The list of security IP addresses for review. Each group supports a maximum of 100 IP addresses. Separate multiple IP addresses with commas (,). The following formats are supported:
	//
	// - Exact IP address: 192.168.0.1
	//
	// - CIDR block: 192.168.0.1/24 (Classless Inter-Domain Routing. /24 specifies the length of the prefix in the address. Valid values: `[1,32]`.)
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.0.1
	Ips *string `json:"Ips,omitempty" xml:"Ips,omitempty"`
	// The operation mode. Valid values:
	//
	// - **Append**: default value. Appends IP addresses to the IP address whitelist.
	//
	// - **Cover**: overwrites the existing IP address whitelist.
	//
	// - **Delete**: deletes IP addresses from the IP address whitelist.
	//
	// > If the specified value is not within the valid values, the default value (Append) is used.
	//
	// example:
	//
	// Cover
	OperateMode *string `json:"OperateMode,omitempty" xml:"OperateMode,omitempty"`
	// The name of the security group for review. Default value: **Default**. A maximum of 10 security groups are supported.
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
