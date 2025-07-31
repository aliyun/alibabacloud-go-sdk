// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifySecurityIpsRequest
	GetDBInstanceId() *string
	SetModifyMode(v string) *ModifySecurityIpsRequest
	GetModifyMode() *string
	SetOwnerAccount(v string) *ModifySecurityIpsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySecurityIpsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifySecurityIpsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySecurityIpsRequest
	GetResourceOwnerId() *int64
	SetSecurityIpGroupAttribute(v string) *ModifySecurityIpsRequest
	GetSecurityIpGroupAttribute() *string
	SetSecurityIpGroupName(v string) *ModifySecurityIpsRequest
	GetSecurityIpGroupName() *string
	SetSecurityIps(v string) *ModifySecurityIpsRequest
	GetSecurityIps() *string
}

type ModifySecurityIpsRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp*****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The method that is used to modify the IP address whitelist. Valid values:
	//
	// 	- **Cover**: overwrites the original IP address whitelist.
	//
	// 	- **Append**: appends data to the IP address whitelist.
	//
	// 	- **Delete**: deletes the IP address whitelist.
	//
	// Default value: **Cover**.
	//
	// example:
	//
	// Append
	ModifyMode           *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The attribute of the IP address whitelist. It can contain a maximum of 120 characters in length and can contain uppercase letters, lowercase letters, and digits.
	//
	// This parameter is empty by default.
	//
	// example:
	//
	// test
	SecurityIpGroupAttribute *string `json:"SecurityIpGroupAttribute,omitempty" xml:"SecurityIpGroupAttribute,omitempty"`
	// The name of the IP address whitelist that you want to modify. Default value: **default**.
	//
	// example:
	//
	// allowserver
	SecurityIpGroupName *string `json:"SecurityIpGroupName,omitempty" xml:"SecurityIpGroupName,omitempty"`
	// The IP addresses in the IP address whitelist. Separate multiple IP addresses with commas (,). You can add a maximum of 1,000 different IP addresses to the IP address whitelist. The entries in the IP address whitelist must be in one of the following formats:
	//
	// 	- IP addresses, such as 127.0.0.1.
	//
	// 	- CIDR blocks, such as 127.0.0.1/24. In this example, 24 indicates that the prefix of each IP address in the IP address whitelist is 24 bits in length. You can replace 24 with a value within the range of 1 to 32.
	//
	// This parameter is required.
	//
	// example:
	//
	// 127.0.0.1/24,127.0.0.1
	SecurityIps *string `json:"SecurityIps,omitempty" xml:"SecurityIps,omitempty"`
}

func (s ModifySecurityIpsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityIpsRequest) GoString() string {
	return s.String()
}

func (s *ModifySecurityIpsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifySecurityIpsRequest) GetModifyMode() *string {
	return s.ModifyMode
}

func (s *ModifySecurityIpsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySecurityIpsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySecurityIpsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySecurityIpsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySecurityIpsRequest) GetSecurityIpGroupAttribute() *string {
	return s.SecurityIpGroupAttribute
}

func (s *ModifySecurityIpsRequest) GetSecurityIpGroupName() *string {
	return s.SecurityIpGroupName
}

func (s *ModifySecurityIpsRequest) GetSecurityIps() *string {
	return s.SecurityIps
}

func (s *ModifySecurityIpsRequest) SetDBInstanceId(v string) *ModifySecurityIpsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetModifyMode(v string) *ModifySecurityIpsRequest {
	s.ModifyMode = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetOwnerAccount(v string) *ModifySecurityIpsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetOwnerId(v int64) *ModifySecurityIpsRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetResourceOwnerAccount(v string) *ModifySecurityIpsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetResourceOwnerId(v int64) *ModifySecurityIpsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetSecurityIpGroupAttribute(v string) *ModifySecurityIpsRequest {
	s.SecurityIpGroupAttribute = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetSecurityIpGroupName(v string) *ModifySecurityIpsRequest {
	s.SecurityIpGroupName = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetSecurityIps(v string) *ModifySecurityIpsRequest {
	s.SecurityIps = &v
	return s
}

func (s *ModifySecurityIpsRequest) Validate() error {
	return dara.Validate(s)
}
