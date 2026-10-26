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
	// The attribute of the whitelist group.
	//
	// - If you do not specify this parameter, the group is a regular group. This is the default value.
	//
	// - If you set this parameter to `hidden`, the group is a hidden group that is automatically created by the system for services such as DMS, DTS, and DAS. Hidden groups are not displayed in the console. If you delete or modify a hidden group, services such as DMS, DTS, and DAS cannot access the ApsaraDB RDS instance. Proceed with caution.
	//
	// example:
	//
	// hidden
	DBInstanceIPArrayAttribute *string `json:"DBInstanceIPArrayAttribute,omitempty" xml:"DBInstanceIPArrayAttribute,omitempty"`
	// The name of the IP address whitelist group to be modified. Default value: Default. If the specified group does not exist, the system automatically creates the group.
	//
	// >You can create up to 200 whitelist groups for a single instance.
	//
	// example:
	//
	// test
	DBInstanceIPArrayName *string `json:"DBInstanceIPArrayName,omitempty" xml:"DBInstanceIPArrayName,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pgm-bp18n0c8zt45****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The read-only instances to which the whitelist is synchronized.
	//
	// - This parameter is applicable only to ApsaraDB RDS for PostgreSQL instances that have read-only instances.
	//
	// - If multiple read-only instances exist, separate their IDs with commas (,).
	//
	// example:
	//
	// pgr-bp17yuz4dn3d****,pgr-bp1vn2ph54u1****
	FreshWhiteListReadins *string `json:"FreshWhiteListReadins,omitempty" xml:"FreshWhiteListReadins,omitempty"`
	// The modification mode. Valid values:
	//
	// 	- **Cover**: overwrites the original IP address whitelist with the IP addresses specified by the **SecurityIps*	- parameter. This is the default value.
	//
	// 	- **Append**: appends the IP addresses specified by the **SecurityIps*	- parameter to the original IP address whitelist.
	//
	// 	- **Delete**: deletes the IP addresses specified by the **SecurityIps*	- parameter from the original IP address whitelist. You must retain at least one IP address in the whitelist.
	//
	// example:
	//
	// Cover
	ModifyMode      *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the IP address. Set the value to IPv4. IPv6 is not supported.
	//
	// example:
	//
	// IPv4
	SecurityIPType *string `json:"SecurityIPType,omitempty" xml:"SecurityIPType,omitempty"`
	// The IP address whitelist. Before you modify the IP address whitelist, call the [DescribeDBInstanceIPArrayList](https://help.aliyun.com/document_detail/610518.html) operation to query the existing IP address whitelist of the instance.
	//
	// <details>
	//
	// <summary>Configuration rules</summary>
	//
	// - You can enter IP addresses (such as 10.23.XX.XX) or CIDR blocks (such as 10.23.XX.XX/24).
	//
	// - Separate multiple IP addresses or CIDR blocks with commas (,). Do not add spaces before or after the commas.
	//
	// - You can add up to 1,000 IP addresses or CIDR blocks to the whitelist of a single instance. If you have a large number of IP addresses, we recommend that you merge them into CIDR blocks, such as 10.23.XX.XX/24.
	//
	// </details>
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.23.XX.XX
	SecurityIps *string `json:"SecurityIps,omitempty" xml:"SecurityIps,omitempty"`
	// The network type of the whitelist. Valid values:
	//
	// 	- **MIX**: the general mode. This is the default value.
	//
	// 	- **Classic**: the classic network in the enhanced whitelist mode.
	//
	// 	- **VPC**: the VPC in the enhanced whitelist mode.
	//
	// > 	- For ApsaraDB RDS for PostgreSQL instances with cloud disks, the network type is fixed to the general mode (MIX). If you configure another mode, the system automatically changes it to the general mode.
	//
	// > 	- The enhanced whitelist mode is supported only by ApsaraDB RDS for MySQL 5.1, 5.5, 5.6, and 5.7 instances with premium local SSDs, and ApsaraDB RDS for PostgreSQL 9.4 and 10 instances with premium local SSDs.
	//
	// example:
	//
	// MIX
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
