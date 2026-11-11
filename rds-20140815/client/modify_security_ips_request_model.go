// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchMode(v bool) *ModifySecurityIpsRequest
	GetBatchMode() *bool
	SetDBInstanceIPArrayAttribute(v string) *ModifySecurityIpsRequest
	GetDBInstanceIPArrayAttribute() *string
	SetDBInstanceIPArrayName(v string) *ModifySecurityIpsRequest
	GetDBInstanceIPArrayName() *string
	SetDBInstanceId(v string) *ModifySecurityIpsRequest
	GetDBInstanceId() *string
	SetFreshWhiteListReadins(v string) *ModifySecurityIpsRequest
	GetFreshWhiteListReadins() *string
	SetGroupItems(v []*ModifySecurityIpsRequestGroupItems) *ModifySecurityIpsRequest
	GetGroupItems() []*ModifySecurityIpsRequestGroupItems
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
	// Specifies whether to enable the batch modification mode. Valid values:
	//
	// true: batch mode. Modifies multiple whitelist groups at a time based on the GroupItems parameter. In this case, the single-group parameters such as SecurityIps, DBInstanceIPArrayName, and DBInstanceIPArrayAttribute do not take effect.
	//
	// false: single-group mode. Modifies a single whitelist group based on parameters such as SecurityIps. This is the default value.
	//
	// example:
	//
	// false
	BatchMode *bool `json:"BatchMode,omitempty" xml:"BatchMode,omitempty"`
	// The attribute of the IP whitelist group.
	//
	// - If you do not specify this parameter, the group is a regular group. This is the default value.
	//
	// - If you set this parameter to `hidden`, the group is a system-generated group used by services such as DMS, DTS, and DAS. These groups are not displayed in the console. Deleting or modifying these groups causes services such as DMS, DTS, and DAS to fail to access the ApsaraDB RDS instance. Proceed with caution.
	//
	// example:
	//
	// hidden
	DBInstanceIPArrayAttribute *string `json:"DBInstanceIPArrayAttribute,omitempty" xml:"DBInstanceIPArrayAttribute,omitempty"`
	// The name of the IP whitelist group to be modified. Default value: Default. If the specified group does not exist, the system automatically creates the group.
	//
	// > A single instance supports up to 200 IP whitelist groups.
	//
	// example:
	//
	// test
	DBInstanceIPArrayName *string `json:"DBInstanceIPArrayName,omitempty" xml:"DBInstanceIPArrayName,omitempty"`
	// The instance ID of the target instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pgm-bp18n0c8zt45****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The list of read-only instances to which the whitelist is synchronized.
	//
	// - This parameter applies only to ApsaraDB RDS for PostgreSQL instances for which read-only instances are created.
	//
	// - If multiple read-only instances exist, separate their IDs with commas (,).
	//
	// example:
	//
	// pgr-bp17yuz4dn3d****,pgr-bp1vn2ph54u1****
	FreshWhiteListReadins *string `json:"FreshWhiteListReadins,omitempty" xml:"FreshWhiteListReadins,omitempty"`
	// The details of the whitelist groups. This parameter is required when BatchMode is set to true. In each group, GroupName is required and must be unique. You must specify at least one of SecurityIPList and GroupTag.
	//
	// example:
	//
	// [{"GroupName":"default","SecurityIPList":"127.0.0.1"},{"GroupName":"test","SecurityIPList":"10.23.XX.XX,10.24.XX.XX"}]
	GroupItems []*ModifySecurityIpsRequestGroupItems `json:"GroupItems,omitempty" xml:"GroupItems,omitempty" type:"Repeated"`
	// The modification mode. Valid values:
	//
	// 	- **Cover**: overwrites the original IP whitelist with the value of the **SecurityIps*	- parameter. This is the default value.
	//
	// 	- **Append**: appends the IP addresses specified in the **SecurityIps*	- parameter to the original IP whitelist.
	//
	// 	- **Delete**: deletes the IP addresses specified in the **SecurityIps*	- parameter from the original IP whitelist. You must retain at least one IP address.
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
	// The IP whitelist. Before you modify the IP whitelist, call the [DescribeDBInstanceIPArrayList](https://help.aliyun.com/document_detail/610518.html) operation to query the existing IP whitelist of the instance.
	//
	// <details>
	//
	// <summary>Configuration rules</summary>
	//
	// - IP addresses (such as 10.23.XX.XX) and CIDR blocks (such as 10.23.XX.XX/24) are supported.
	//
	// - Separate multiple IP addresses or CIDR blocks with commas (,). Do not add spaces before or after the commas.
	//
	// - You can add up to 1,000 IP addresses or CIDR blocks to a single instance. If you have a large number of IP addresses, we recommend that you merge scattered IP addresses into CIDR blocks, such as 10.23.XX.XX/24.
	//
	// </details>
	//
	// example:
	//
	// 10.23.XX.XX
	SecurityIps *string `json:"SecurityIps,omitempty" xml:"SecurityIps,omitempty"`
	// The network type of the whitelist. Valid values:
	//
	// 	- **MIX**: general mode. This is the default value.
	//
	// 	- **Classic**: classic network in enhanced whitelist mode.
	//
	// 	- **VPC**: VPC in enhanced whitelist mode.
	//
	// > 	- For ApsaraDB RDS for PostgreSQL instances with cloud disks, the network type is fixed to MIX. If you configure another mode, the system automatically converts it to MIX.
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

func (s *ModifySecurityIpsRequest) GetBatchMode() *bool {
	return s.BatchMode
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

func (s *ModifySecurityIpsRequest) GetGroupItems() []*ModifySecurityIpsRequestGroupItems {
	return s.GroupItems
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

func (s *ModifySecurityIpsRequest) SetBatchMode(v bool) *ModifySecurityIpsRequest {
	s.BatchMode = &v
	return s
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

func (s *ModifySecurityIpsRequest) SetGroupItems(v []*ModifySecurityIpsRequestGroupItems) *ModifySecurityIpsRequest {
	s.GroupItems = v
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
	if s.GroupItems != nil {
		for _, item := range s.GroupItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifySecurityIpsRequestGroupItems struct {
	// The name of the whitelist group. This parameter is required in batch mode and must be unique within the same GroupItems list.
	//
	// example:
	//
	// default
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The attribute of the whitelist group. If you set this parameter to hidden, the group is a system-generated group used by services such as DMS, DTS, and DAS. These groups are not displayed in the console. Proceed with caution. Generally, you do not need to specify this parameter. You must specify at least one of this parameter and SecurityIPList.
	//
	// example:
	//
	// hidden
	GroupTag *string `json:"GroupTag,omitempty" xml:"GroupTag,omitempty"`
	// The modified IP address whitelist of the group. Separate multiple IP addresses with commas (,). You must specify at least one of this parameter and GroupTag.
	//
	// example:
	//
	// 10.23.XX.XX,10.24.XX.XX
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
}

func (s ModifySecurityIpsRequestGroupItems) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityIpsRequestGroupItems) GoString() string {
	return s.String()
}

func (s *ModifySecurityIpsRequestGroupItems) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifySecurityIpsRequestGroupItems) GetGroupTag() *string {
	return s.GroupTag
}

func (s *ModifySecurityIpsRequestGroupItems) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *ModifySecurityIpsRequestGroupItems) SetGroupName(v string) *ModifySecurityIpsRequestGroupItems {
	s.GroupName = &v
	return s
}

func (s *ModifySecurityIpsRequestGroupItems) SetGroupTag(v string) *ModifySecurityIpsRequestGroupItems {
	s.GroupTag = &v
	return s
}

func (s *ModifySecurityIpsRequestGroupItems) SetSecurityIPList(v string) *ModifySecurityIpsRequestGroupItems {
	s.SecurityIPList = &v
	return s
}

func (s *ModifySecurityIpsRequestGroupItems) Validate() error {
	return dara.Validate(s)
}
