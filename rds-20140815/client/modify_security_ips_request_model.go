// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceIPArrayAttribute(v string) *ModifySecurityIpsRequest
	GetDBInstanceIPArrayAttribute() *string
	SetDBInstanceIPArrayName(v string) *ModifySecurityIpsRequest
	GetDBInstanceIPArrayName() *string
	SetDBInstanceId(v string) *ModifySecurityIpsRequest
	GetDBInstanceId() *string
	SetFreshWhiteListReadins(v string) *ModifySecurityIpsRequest
	GetFreshWhiteListReadins() *string
	SetModifyMode(v string) *ModifySecurityIpsRequest
	GetModifyMode() *string
	SetResourceOwnerId(v int64) *ModifySecurityIpsRequest
	GetResourceOwnerId() *int64
	SetSecurityIPType(v string) *ModifySecurityIpsRequest
	GetSecurityIPType() *string
	SetSecurityIps(v string) *ModifySecurityIpsRequest
	GetSecurityIps() *string
	SetWhitelistNetworkType(v string) *ModifySecurityIpsRequest
	GetWhitelistNetworkType() *string
}

type ModifySecurityIpsRequest struct {
	// The attribute of the IP address whitelist. By default, this parameter is empty.
	//
	// > The IP address whitelists that have the hidden attribute are not displayed in the ApsaraDB RDS console. These IP address whitelists are used to access Alibaba Cloud services, such as Data Transmission Service (DTS).
	//
	// example:
	//
	// hidden
	DBInstanceIPArrayAttribute *string `json:"DBInstanceIPArrayAttribute,omitempty" xml:"DBInstanceIPArrayAttribute,omitempty"`
	// The name of the IP address whitelist that you want to modify. Default value: **Default**.
	//
	// > A maximum of 200 IP address whitelists can be configured for each instance.
	//
	// example:
	//
	// test
	DBInstanceIPArrayName *string `json:"DBInstanceIPArrayName,omitempty" xml:"DBInstanceIPArrayName,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pgm-bp18n0c8zt45****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The read-only instances to which you want to synchronize the IP address whitelist.
	//
	// 	- This parameter applies only to ApsaraDB RDS for PostgreSQL instances.
	//
	// 	- If the instance is attached with a read-only instance, you can use this parameter to synchronize the IP address whitelist to the read-only instance. If the instance is attached with multiple read-only instances, separate the read-only instances with commas (,).
	//
	// 	- If the instance is not attached with a read-only instance, leave this parameter empty.
	//
	// example:
	//
	// pgr-bp17yuz4dn3d****,pgr-bp1vn2ph54u1****
	FreshWhiteListReadins *string `json:"FreshWhiteListReadins,omitempty" xml:"FreshWhiteListReadins,omitempty"`
	// The method that is used to modify the whitelist. Valid values:
	//
	// 	- **Cover**: Use the IP addresses and CIDR blocks that are specified in the **SecurityIps*	- parameter to overwrite the existing IP addresses and CIDR blocks in the IP address whitelist.
	//
	// 	- **Append**: Add the IP addresses and CIDR blocks that are specified in the **SecurityIps*	- parameter to the IP address whitelist.
	//
	// 	- **Delete**: Delete the IP addresses and CIDR blocks that are specified in the **SecurityIps*	- parameter from the IP address whitelist. You must retain at least one IP address or CIDR block.
	//
	// Default value: **Cover**.
	//
	// example:
	//
	// Cover
	ModifyMode      *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The IP address type. The value is fixed as IPv4.
	//
	// example:
	//
	// IPv4
	SecurityIPType *string `json:"SecurityIPType,omitempty" xml:"SecurityIPType,omitempty"`
	// The IP addresses in an IP address whitelist. Separate multiple IP addresses with commas (,). Each IP address in the IP address whitelist must be unique. The entries in the IP address whitelist must be in one of the following formats:
	//
	// 	- IP addresses, such as 10.23.XX.XX.
	//
	// 	- CIDR blocks, such as 10.23.XX.XX/24. In this example, 24 indicates that the prefix of each IP address in the IP address whitelist is 24 bits in length. You can replace 24 with a value within the range of 1 to 32.
	//
	// > A maximum of 1,000 IP addresses or CIDR blocks can be added for each instance. If you want to add a large number of IP addresses, we recommend that you merge them into CIDR blocks, such as 10.23.XX.XX/24.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.23.XX.XX
	SecurityIps *string `json:"SecurityIps,omitempty" xml:"SecurityIps,omitempty"`
	// The network type of the IP address whitelist. Valid values:
	//
	// 	- **Classic**: classic network in enhanced whitelist mode
	//
	// 	- **VPC**: virtual private cloud (VPC) network type in enhanced whitelist mode.
	//
	// 	- **MIX**: standard whitelist mode
	//
	// Default value: **MIX**.
	//
	// >
	//
	// 	- In standard whitelist mode, IP addresses and CIDR blocks are added only to the default IP address whitelist. In enhanced whitelist mode, IP addresses and CIDR blocks are added to the IP address whitelists of the classic network type and the VPC network type.
	//
	// 	- If your RDS instance runs PostgreSQL and uses cloud disks, set this parameter to MIX. If you set it to another value, the system automatically changes the value to MIX.
	//
	// example:
	//
	// Classic
	WhitelistNetworkType *string `json:"WhitelistNetworkType,omitempty" xml:"WhitelistNetworkType,omitempty"`
}

func (s ModifySecurityIpsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityIpsRequest) GoString() string {
	return s.String()
}

func (s *ModifySecurityIpsRequest) GetDBInstanceIPArrayAttribute() *string {
	return s.DBInstanceIPArrayAttribute
}

func (s *ModifySecurityIpsRequest) GetDBInstanceIPArrayName() *string {
	return s.DBInstanceIPArrayName
}

func (s *ModifySecurityIpsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifySecurityIpsRequest) GetFreshWhiteListReadins() *string {
	return s.FreshWhiteListReadins
}

func (s *ModifySecurityIpsRequest) GetModifyMode() *string {
	return s.ModifyMode
}

func (s *ModifySecurityIpsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySecurityIpsRequest) GetSecurityIPType() *string {
	return s.SecurityIPType
}

func (s *ModifySecurityIpsRequest) GetSecurityIps() *string {
	return s.SecurityIps
}

func (s *ModifySecurityIpsRequest) GetWhitelistNetworkType() *string {
	return s.WhitelistNetworkType
}

func (s *ModifySecurityIpsRequest) SetDBInstanceIPArrayAttribute(v string) *ModifySecurityIpsRequest {
	s.DBInstanceIPArrayAttribute = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetDBInstanceIPArrayName(v string) *ModifySecurityIpsRequest {
	s.DBInstanceIPArrayName = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetDBInstanceId(v string) *ModifySecurityIpsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetFreshWhiteListReadins(v string) *ModifySecurityIpsRequest {
	s.FreshWhiteListReadins = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetModifyMode(v string) *ModifySecurityIpsRequest {
	s.ModifyMode = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetResourceOwnerId(v int64) *ModifySecurityIpsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetSecurityIPType(v string) *ModifySecurityIpsRequest {
	s.SecurityIPType = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetSecurityIps(v string) *ModifySecurityIpsRequest {
	s.SecurityIps = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetWhitelistNetworkType(v string) *ModifySecurityIpsRequest {
	s.WhitelistNetworkType = &v
	return s
}

func (s *ModifySecurityIpsRequest) Validate() error {
	return dara.Validate(s)
}
