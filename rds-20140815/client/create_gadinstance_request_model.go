// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGADInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCentralDBInstanceId(v string) *CreateGADInstanceRequest
	GetCentralDBInstanceId() *string
	SetCentralRdsDtsAdminAccount(v string) *CreateGADInstanceRequest
	GetCentralRdsDtsAdminAccount() *string
	SetCentralRdsDtsAdminPassword(v string) *CreateGADInstanceRequest
	GetCentralRdsDtsAdminPassword() *string
	SetCentralRegionId(v string) *CreateGADInstanceRequest
	GetCentralRegionId() *string
	SetDBList(v string) *CreateGADInstanceRequest
	GetDBList() *string
	SetDescription(v string) *CreateGADInstanceRequest
	GetDescription() *string
	SetResourceGroupId(v string) *CreateGADInstanceRequest
	GetResourceGroupId() *string
	SetTag(v []*CreateGADInstanceRequestTag) *CreateGADInstanceRequest
	GetTag() []*CreateGADInstanceRequestTag
	SetUnitNode(v []*CreateGADInstanceRequestUnitNode) *CreateGADInstanceRequest
	GetUnitNode() []*CreateGADInstanceRequestUnitNode
}

type CreateGADInstanceRequest struct {
	// The ID of the primary instance. You can call the DescribeDBInstances operation to query the instance ID. This instance serves as the central node (primary node) of the GAD cluster.
	//
	// > 	- A primary instance ID can serve as the central node of only one GAD cluster.
	//
	// > 	- Only ApsaraDB RDS for MySQL primary instances in the China (Hangzhou), China (Shanghai), China (Qingdao), China (Beijing), China (Zhangjiakou), China (Shenzhen), and China (Chengdu) regions can serve as the central node of a GAD cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	CentralDBInstanceId *string `json:"CentralDBInstanceId,omitempty" xml:"CentralDBInstanceId,omitempty"`
	// The privileged account of the central node. You can call the DescribeAccounts operation to query the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	CentralRdsDtsAdminAccount *string `json:"CentralRdsDtsAdminAccount,omitempty" xml:"CentralRdsDtsAdminAccount,omitempty"`
	// The password of the privileged account for the central node.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test12345
	CentralRdsDtsAdminPassword *string `json:"CentralRdsDtsAdminPassword,omitempty" xml:"CentralRdsDtsAdminPassword,omitempty"`
	// The region ID of the central node. You can call the DescribeRegions operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	CentralRegionId *string `json:"CentralRegionId,omitempty" xml:"CentralRegionId,omitempty"`
	// A JSON array that contains the database information of the central node. All database information in this array is synchronized to the current unit node (secondary node). Parameter description:
	//
	// 	- **name**: the database name.
	//
	// 	- **all**: specifies whether to synchronize all data in the current database or table. Valid values: **true*	- | **false**.
	//
	// 	- **Table**: the table name. If the **all*	- parameter is set to **false**, you must also specify the names of the tables to be synchronized in the JSON array.
	//
	// Example: `{
	//
	//    "testdb": {
	//
	//     "name": "testdb",
	//
	//     "all": false,
	//
	//     "Table": {
	//
	//       "order": {
	//
	//         "name": "order",
	//
	//         "all": true
	//
	//       },
	//
	//       "ordernew": {
	//
	//         "name": "ordernew",
	//
	//         "all": true
	//
	//       }
	//
	//     }
	//
	//   }
	//
	// }`
	//
	// This parameter is required.
	//
	// example:
	//
	// {    "testdb": {     "name": "testdb",     "all": false,     "Table": {       "order": {         "name": "order",         "all": true       },       "ordernew": {         "name": "ordernew",         "all": true       }     }   } }
	DBList *string `json:"DBList,omitempty" xml:"DBList,omitempty"`
	// The name of the GAD cluster.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	Tag []*CreateGADInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The unit node information.
	//
	// This parameter is required.
	UnitNode []*CreateGADInstanceRequestUnitNode `json:"UnitNode,omitempty" xml:"UnitNode,omitempty" type:"Repeated"`
}

func (s CreateGADInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGADInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateGADInstanceRequest) GetCentralDBInstanceId() *string {
	return s.CentralDBInstanceId
}

func (s *CreateGADInstanceRequest) GetCentralRdsDtsAdminAccount() *string {
	return s.CentralRdsDtsAdminAccount
}

func (s *CreateGADInstanceRequest) GetCentralRdsDtsAdminPassword() *string {
	return s.CentralRdsDtsAdminPassword
}

func (s *CreateGADInstanceRequest) GetCentralRegionId() *string {
	return s.CentralRegionId
}

func (s *CreateGADInstanceRequest) GetDBList() *string {
	return s.DBList
}

func (s *CreateGADInstanceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateGADInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateGADInstanceRequest) GetTag() []*CreateGADInstanceRequestTag {
	return s.Tag
}

func (s *CreateGADInstanceRequest) GetUnitNode() []*CreateGADInstanceRequestUnitNode {
	return s.UnitNode
}

func (s *CreateGADInstanceRequest) SetCentralDBInstanceId(v string) *CreateGADInstanceRequest {
	s.CentralDBInstanceId = &v
	return s
}

func (s *CreateGADInstanceRequest) SetCentralRdsDtsAdminAccount(v string) *CreateGADInstanceRequest {
	s.CentralRdsDtsAdminAccount = &v
	return s
}

func (s *CreateGADInstanceRequest) SetCentralRdsDtsAdminPassword(v string) *CreateGADInstanceRequest {
	s.CentralRdsDtsAdminPassword = &v
	return s
}

func (s *CreateGADInstanceRequest) SetCentralRegionId(v string) *CreateGADInstanceRequest {
	s.CentralRegionId = &v
	return s
}

func (s *CreateGADInstanceRequest) SetDBList(v string) *CreateGADInstanceRequest {
	s.DBList = &v
	return s
}

func (s *CreateGADInstanceRequest) SetDescription(v string) *CreateGADInstanceRequest {
	s.Description = &v
	return s
}

func (s *CreateGADInstanceRequest) SetResourceGroupId(v string) *CreateGADInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateGADInstanceRequest) SetTag(v []*CreateGADInstanceRequestTag) *CreateGADInstanceRequest {
	s.Tag = v
	return s
}

func (s *CreateGADInstanceRequest) SetUnitNode(v []*CreateGADInstanceRequestUnitNode) *CreateGADInstanceRequest {
	s.UnitNode = v
	return s
}

func (s *CreateGADInstanceRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UnitNode != nil {
		for _, item := range s.UnitNode {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateGADInstanceRequestTag struct {
	// The tag key. You can create up to N tag keys at a time. Valid values of N: **1 to 20**. The tag key cannot be an empty string.
	//
	// example:
	//
	// testkey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value that corresponds to the tag key. You can create up to N tag values at a time. Valid values of N: **1 to 20**. The tag value can be an empty string.
	//
	// example:
	//
	// testvalue1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateGADInstanceRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateGADInstanceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateGADInstanceRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateGADInstanceRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateGADInstanceRequestTag) SetKey(v string) *CreateGADInstanceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateGADInstanceRequestTag) SetValue(v string) *CreateGADInstanceRequestTag {
	s.Value = &v
	return s
}

func (s *CreateGADInstanceRequestTag) Validate() error {
	return dara.Validate(s)
}

type CreateGADInstanceRequestUnitNode struct {
	// The name of the new unit node. The name must meet the following requirements:
	//
	// - The name must be **2 to 255*	- characters in length.
	//
	// - The name must start with a letter or a Chinese character. It can contain digits, Chinese characters, letters, underscores (_), and hyphens (-).
	//
	// - The name cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The storage capacity of the new unit node. Unit: GB. The value is incremented in 5 GB increments. For the value range, see [Primary instance types](https://help.aliyun.com/document_detail/26312.html). You can also call the DescribeAvailableResource operation to query the available storage capacity range for the target instance type.
	//
	// example:
	//
	// 20
	DBInstanceStorage *int64 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The instance storage type. Valid values:
	//
	// 	- **local_ssd**: Premium Local SSD (recommended).
	//
	// 	- **cloud_ssd**: standard SSD (not recommended because standard SSDs are no longer available for purchase in some regions).
	//
	// 	- **cloud_essd**: PL1 ESSD.
	//
	// 	- **cloud_essd2**: PL2 ESSD.
	//
	// 	- **cloud_essd3**: PL3 ESSD.
	//
	// The default value of this parameter is determined by the instance type specified in the **DBInstanceClass*	- parameter:
	//
	// - If the instance type is a Premium Local SSD instance type, the default value is **local_ssd**.
	//
	// - If the instance type is a cloud disk instance type, the default value is **cloud_essd**.
	//
	// example:
	//
	// cloud_essd2
	DBInstanceStorageType *string `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	// The instance type of the new unit node. For more information, see [Primary instance types](https://help.aliyun.com/document_detail/26312.html). You can also call the DescribeAvailableResource operation to query the available instance types in the target region.
	//
	// example:
	//
	// rds.mysql.t1.small
	DbInstanceClass *string `json:"DbInstanceClass,omitempty" xml:"DbInstanceClass,omitempty"`
	// The conflict resolution policy used when a primary key conflict occurs during data synchronization for the new unit node. Valid values:
	//
	// 	- **overwrite**: overwrites the conflicting primary key on the destination node.
	//
	// 	- **interrupt**: stops the synchronization task and reports an error.
	//
	// 	- **ignore**: ignores the conflicting primary key on the current node.
	//
	// This parameter is required.
	//
	// example:
	//
	// overwrite
	DtsConflict *string `json:"DtsConflict,omitempty" xml:"DtsConflict,omitempty"`
	// The specification of the data synchronization link for the new unit node. Valid values:
	//
	// 	- **small**
	//
	// 	- **medium**
	//
	// 	- **large**
	//
	// 	- **micro**
	//
	// >For more information about the differences between specifications, see [Data synchronization link specifications](https://help.aliyun.com/document_detail/26605.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// medium
	DtsInstanceClass *string `json:"DtsInstanceClass,omitempty" xml:"DtsInstanceClass,omitempty"`
	// The database engine of the new unit node. Only **MySQL*	- is supported.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version of the new unit node. Valid values:
	//
	// 	- **8.0**
	//
	// 	- **5.7**
	//
	// 	- **5.6**
	//
	// 	- **5.5**
	//
	// example:
	//
	// 8.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The billing method of the new unit node. Valid values:
	//
	// 	- **Postpaid**: pay-as-you-go.
	//
	// 	- **Prepaid**: subscription.
	//
	// >The system automatically generates and completes the payment for the order. You do not need to manually confirm the payment.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The region ID of the new unit node. You can call the DescribeRegions operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionID *string `json:"RegionID,omitempty" xml:"RegionID,omitempty"`
	// The [IP address whitelist](https://help.aliyun.com/document_detail/43185.html) of the new unit node. Separate multiple entries with commas (,). Entries cannot be duplicated. A maximum of 1,000 entries are allowed. The following two formats are supported:
	//
	// 	- IP address format, such as `10.10.10.10`.
	//
	// 	- CIDR format, such as `10.10.10.10/24` (Classless Inter-Domain Routing, where **24*	- indicates the length of the prefix, ranging from **1 to 32**).
	//
	// example:
	//
	// 10.10.10.10
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The vSwitch ID of the new unit node.
	//
	// example:
	//
	// vsw-bp1tg609m5j85****
	VSwitchID *string `json:"VSwitchID,omitempty" xml:"VSwitchID,omitempty"`
	// The virtual private cloud (VPC) ID of the new unit node.
	//
	// example:
	//
	// vpc-bp19ame5m1r3o****
	VpcID *string `json:"VpcID,omitempty" xml:"VpcID,omitempty"`
	// The zone ID of the new unit node. You can call the DescribeRegions operation to query the zone ID.
	//
	// example:
	//
	// cn-hangzhou-j
	ZoneID *string `json:"ZoneID,omitempty" xml:"ZoneID,omitempty"`
	// The zone ID of the secondary node for the new unit node. You can call the DescribeRegions operation to query the zone ID.
	//
	// 	- If this value is the same as the **ZoneId*	- of the current unit node, the single-zone deployment is used.
	//
	// 	- If this value is different from the **ZoneId*	- of the current unit node, the multi-zone deployment is used.
	//
	// example:
	//
	// cn-hangzhou-j
	ZoneIDSlave1 *string `json:"ZoneIDSlave1,omitempty" xml:"ZoneIDSlave1,omitempty"`
	// The zone ID of the logger node for the new unit node. You can call the DescribeRegions operation to query the zone ID.
	//
	// 	- If this value is the same as the **ZoneId*	- of the current unit node, the single-zone deployment is used.
	//
	// 	- If this value is different from the **ZoneId*	- of the current unit node, the multi-zone deployment is used.
	//
	// example:
	//
	// cn-hangzhou-j
	ZoneIDSlave2 *string `json:"ZoneIDSlave2,omitempty" xml:"ZoneIDSlave2,omitempty"`
}

func (s CreateGADInstanceRequestUnitNode) String() string {
	return dara.Prettify(s)
}

func (s CreateGADInstanceRequestUnitNode) GoString() string {
	return s.String()
}

func (s *CreateGADInstanceRequestUnitNode) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *CreateGADInstanceRequestUnitNode) GetDBInstanceStorage() *int64 {
	return s.DBInstanceStorage
}

func (s *CreateGADInstanceRequestUnitNode) GetDBInstanceStorageType() *string {
	return s.DBInstanceStorageType
}

func (s *CreateGADInstanceRequestUnitNode) GetDbInstanceClass() *string {
	return s.DbInstanceClass
}

func (s *CreateGADInstanceRequestUnitNode) GetDtsConflict() *string {
	return s.DtsConflict
}

func (s *CreateGADInstanceRequestUnitNode) GetDtsInstanceClass() *string {
	return s.DtsInstanceClass
}

func (s *CreateGADInstanceRequestUnitNode) GetEngine() *string {
	return s.Engine
}

func (s *CreateGADInstanceRequestUnitNode) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *CreateGADInstanceRequestUnitNode) GetPayType() *string {
	return s.PayType
}

func (s *CreateGADInstanceRequestUnitNode) GetRegionID() *string {
	return s.RegionID
}

func (s *CreateGADInstanceRequestUnitNode) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *CreateGADInstanceRequestUnitNode) GetVSwitchID() *string {
	return s.VSwitchID
}

func (s *CreateGADInstanceRequestUnitNode) GetVpcID() *string {
	return s.VpcID
}

func (s *CreateGADInstanceRequestUnitNode) GetZoneID() *string {
	return s.ZoneID
}

func (s *CreateGADInstanceRequestUnitNode) GetZoneIDSlave1() *string {
	return s.ZoneIDSlave1
}

func (s *CreateGADInstanceRequestUnitNode) GetZoneIDSlave2() *string {
	return s.ZoneIDSlave2
}

func (s *CreateGADInstanceRequestUnitNode) SetDBInstanceDescription(v string) *CreateGADInstanceRequestUnitNode {
	s.DBInstanceDescription = &v
	return s
}

func (s *CreateGADInstanceRequestUnitNode) SetDBInstanceStorage(v int64) *CreateGADInstanceRequestUnitNode {
	s.DBInstanceStorage = &v
	return s
}

func (s *CreateGADInstanceRequestUnitNode) SetDBInstanceStorageType(v string) *CreateGADInstanceRequestUnitNode {
	s.DBInstanceStorageType = &v
	return s
}

func (s *CreateGADInstanceRequestUnitNode) SetDbInstanceClass(v string) *CreateGADInstanceRequestUnitNode {
	s.DbInstanceClass = &v
	return s
}

func (s *CreateGADInstanceRequestUnitNode) SetDtsConflict(v string) *CreateGADInstanceRequestUnitNode {
	s.DtsConflict = &v
	return s
}

func (s *CreateGADInstanceRequestUnitNode) SetDtsInstanceClass(v string) *CreateGADInstanceRequestUnitNode {
	s.DtsInstanceClass = &v
	return s
}

func (s *CreateGADInstanceRequestUnitNode) SetEngine(v string) *CreateGADInstanceRequestUnitNode {
	s.Engine = &v
	return s
}

func (s *CreateGADInstanceRequestUnitNode) SetEngineVersion(v string) *CreateGADInstanceRequestUnitNode {
	s.EngineVersion = &v
	return s
}

func (s *CreateGADInstanceRequestUnitNode) SetPayType(v string) *CreateGADInstanceRequestUnitNode {
	s.PayType = &v
	return s
}

func (s *CreateGADInstanceRequestUnitNode) SetRegionID(v string) *CreateGADInstanceRequestUnitNode {
	s.RegionID = &v
	return s
}

func (s *CreateGADInstanceRequestUnitNode) SetSecurityIPList(v string) *CreateGADInstanceRequestUnitNode {
	s.SecurityIPList = &v
	return s
}

func (s *CreateGADInstanceRequestUnitNode) SetVSwitchID(v string) *CreateGADInstanceRequestUnitNode {
	s.VSwitchID = &v
	return s
}

func (s *CreateGADInstanceRequestUnitNode) SetVpcID(v string) *CreateGADInstanceRequestUnitNode {
	s.VpcID = &v
	return s
}

func (s *CreateGADInstanceRequestUnitNode) SetZoneID(v string) *CreateGADInstanceRequestUnitNode {
	s.ZoneID = &v
	return s
}

func (s *CreateGADInstanceRequestUnitNode) SetZoneIDSlave1(v string) *CreateGADInstanceRequestUnitNode {
	s.ZoneIDSlave1 = &v
	return s
}

func (s *CreateGADInstanceRequestUnitNode) SetZoneIDSlave2(v string) *CreateGADInstanceRequestUnitNode {
	s.ZoneIDSlave2 = &v
	return s
}

func (s *CreateGADInstanceRequestUnitNode) Validate() error {
	return dara.Validate(s)
}
