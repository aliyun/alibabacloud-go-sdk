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
	// The ID of the primary instance. You can call the DescribeDBInstances operation to query the instance ID. The primary instance serves as the central node of the global active database cluster.
	//
	// > 	- A primary instance can serve only as the central node of a single global active database cluster.
	//
	// > 	- The primary instance can serve as the central node of the global active database cluster only in the following regions: China (Hangzhou), China (Shanghai), China (Qingdao), China (Beijing), China (Zhangjiakou), China (Shenzhen), and China (Chengdu).
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5*******
	CentralDBInstanceId *string `json:"CentralDBInstanceId,omitempty" xml:"CentralDBInstanceId,omitempty"`
	// The username of the privileged account of the central node. You can call the DescribeAccounts operation to query the privileged account of the central node.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	CentralRdsDtsAdminAccount *string `json:"CentralRdsDtsAdminAccount,omitempty" xml:"CentralRdsDtsAdminAccount,omitempty"`
	// The password of the privileged account of the central node.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test12345
	CentralRdsDtsAdminPassword *string `json:"CentralRdsDtsAdminPassword,omitempty" xml:"CentralRdsDtsAdminPassword,omitempty"`
	// The region ID of the central node. You can call the DescribeRegions operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	CentralRegionId *string `json:"CentralRegionId,omitempty" xml:"CentralRegionId,omitempty"`
	// A JSON array that consists of the information about a specified database on the central node. All database information that you specify in this array is synchronized to the unit nodes of the global active database cluster. The JSON array contains the following fields:
	//
	// 	- **name**: the name of the database.
	//
	// 	- **all**: specifies whether to synchronize all data in the database or the table. Valid values: **true*	- and **false**.
	//
	// 	- **Table**: the name of the table. If you set the **all*	- field to **false**, you must nest the name of the table that you want to synchronize into the JSON array.
	//
	// Example: `{ "testdb": { "name": "testdb", "all": false, "Table": { "order": { "name": "order", "all": true }, "ordernew": { "name": "ordernew", "all": true } } } }`
	//
	// This parameter is required.
	//
	// example:
	//
	// {    "testdb": {     "name": "testdb",     "all": false,     "Table": {       "order": {         "name": "order",         "all": true       },       "ordernew": {         "name": "ordernew",         "all": true       }     }   } }
	DBList *string `json:"DBList,omitempty" xml:"DBList,omitempty"`
	// The name of the global active database cluster.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// An array that consists of the details about the tag.
	Tag []*CreateGADInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The information about the unit node.
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
	// The key of the tag. You can create N tag keys at a time. Valid values of N: **1 to 20**. The value of this parameter cannot be an empty string.
	//
	// example:
	//
	// testkey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag. You can create N tag values at a time. Valid values of N: **1 to 20**. The value of this parameter can be an empty string.
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
	// The name of the unit node that you want to create. The name must meet the following requirements:
	//
	// 	- The name must be **2 to 255*	- characters in length.
	//
	// 	- The name can contain letters, digits, underscores (_), and hyphens (-) and must start with a letter.
	//
	// 	- Does not start with `http://` or `https://`.
	//
	// example:
	//
	// test
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The storage capacity of the unit node that you want to create. Unit: GB. You can adjust the storage capacity in increments of 5 GB. For more information, see [Primary ApsaraDB RDS instance types](https://help.aliyun.com/document_detail/26312.html). You can also call the DescribeAvailableResource operation to query the storage capacity range that is supported by the new instance type.
	//
	// example:
	//
	// 20
	DBInstanceStorage *int64 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The storage type of the new instance. Valid values:
	//
	// 	- **local_ssd**: Premium Local SSD (recommended)
	//
	// 	- **cloud_ssd**: standard SSD. This storage type is not recommended. Standard SSDs are no longer available for purchase in specific Alibaba Cloud regions.
	//
	// 	- **cloud_essd**: Enterprise SSD (ESSD) of performance level 1 (PL1).
	//
	// 	- **cloud_essd2**: ESSD of PL2.
	//
	// 	- **cloud_essd3**: ESSD of PL3.
	//
	// The default value of this parameter is determined by the instance type specified by the **DBInstanceClass*	- parameter.
	//
	// 	- If the instance type specifies the Premium Local SSD storage type, the default value of this parameter is **local_ssd**.
	//
	// 	- If the instance type specifies the cloud disk storage type, the default value of this parameter is **cloud_essd**.
	//
	// example:
	//
	// cloud_essd2
	DBInstanceStorageType *string `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	// The instance type of the unit node that you want to create. For more information, see [Primary ApsaraDB RDS instance types](https://help.aliyun.com/document_detail/26312.html). You can call the DescribeAvailableResource operation to query the available instance types in a region.
	//
	// example:
	//
	// rds.mysql.t1.small
	DbInstanceClass *string `json:"DbInstanceClass,omitempty" xml:"DbInstanceClass,omitempty"`
	// The conflict resolution policy based on which Data Transmission Service (DTS) responds to primary key conflicts during data synchronization to the unit node that you want to create. Valid values:
	//
	// 	- **overwrite**: DTS overwrites the conflicting primary key on the destination node.
	//
	// 	- **interrupt**: DTS stops the synchronization task, reports an error, and then exits.
	//
	// 	- **ignore**: DTS hides the conflicting primary key on the node.
	//
	// This parameter is required.
	//
	// example:
	//
	// overwrite
	DtsConflict *string `json:"DtsConflict,omitempty" xml:"DtsConflict,omitempty"`
	// The specifications of the data synchronization task for the unit node that you want to create. Valid values:
	//
	// 	- **small**
	//
	// 	- **medium**
	//
	// 	- **large**
	//
	// 	- **micro**
	//
	// >  For more information, see [Specifications of data synchronization tasks](https://help.aliyun.com/document_detail/26605.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// medium
	DtsInstanceClass *string `json:"DtsInstanceClass,omitempty" xml:"DtsInstanceClass,omitempty"`
	// The database engine of the unit node that you want to create. Set the value to **MySQL**.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version of the unit node that you want to create. Valid values:
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
	// The billing method of the unit node that you want to create. Valid values:
	//
	// 	- **Postpaid**: pay-as-you-go
	//
	// 	- **Prepaid**: subscription
	//
	// >  The system automatically generates a purchase order and completes the payment. You do not need to manually confirm the purchase order or complete the payment.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The region ID of the unit node that you want to create. You can call the DescribeRegions operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionID *string `json:"RegionID,omitempty" xml:"RegionID,omitempty"`
	// The [IP address whitelist](https://help.aliyun.com/document_detail/43185.html) of the unit node that you want to create. If you want to add more than one entry to the IP address whitelist, separate the entries with commas (,). Each entry must be unique. The IP address whitelist can contain up to 1,000 entries. The entries in the IP address whitelist must be in one of the following formats:
	//
	// 	- IP addresses, such as `10.10.10.10`.
	//
	// 	- CIDR blocks, such as `10.10.10.10/24`. In this example, **24*	- indicates that the prefix of the IP address in the whitelist is 24 bits in length. You can replace 24 with a value within the range of **1 to 32**.
	//
	// example:
	//
	// 10.10.10.10
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The vSwitch ID of the unit node that you want to create.
	//
	// example:
	//
	// vsw-bp1tg609m5j85********
	VSwitchID *string `json:"VSwitchID,omitempty" xml:"VSwitchID,omitempty"`
	// The virtual private cloud (VPC) ID of the unit node that you want to create.
	//
	// example:
	//
	// vpc-bp19ame5m1r3o********
	VpcID *string `json:"VpcID,omitempty" xml:"VpcID,omitempty"`
	// The zone ID of the unit node that you want to create. You can call the DescribeRegions operation to query the zone ID.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneID *string `json:"ZoneID,omitempty" xml:"ZoneID,omitempty"`
	// The zone ID of the secondary node of the unit node that you want to create. You can call the DescribeRegions operation to query the zone ID.
	//
	// 	- If the value of this parameter is the same as the **zone ID*	- of the unit node that you want to create, the single-zone deployment method is used.
	//
	// 	- If the value of this parameter is different from the **zone ID*	- of the unit node that you want to create, the multiple-zone deployment method is used.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneIDSlave1 *string `json:"ZoneIDSlave1,omitempty" xml:"ZoneIDSlave1,omitempty"`
	// The zone ID of the logger node of the unit node that you want to create. You can call the DescribeRegions operation to query the zone ID.
	//
	// 	- If the value of this parameter is the same as the **zone ID*	- of the unit node that you want to create, the single-zone deployment method is used.
	//
	// 	- If the value of this parameter is different from the **zone ID*	- of the unit node that you want to create, the multiple-zone deployment method is used.
	//
	// example:
	//
	// cn-hangzhou-h
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
