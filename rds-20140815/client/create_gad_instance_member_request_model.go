// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGadInstanceMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCentralDBInstanceId(v string) *CreateGadInstanceMemberRequest
	GetCentralDBInstanceId() *string
	SetCentralRdsDtsAdminAccount(v string) *CreateGadInstanceMemberRequest
	GetCentralRdsDtsAdminAccount() *string
	SetCentralRdsDtsAdminPassword(v string) *CreateGadInstanceMemberRequest
	GetCentralRdsDtsAdminPassword() *string
	SetCentralRegionId(v string) *CreateGadInstanceMemberRequest
	GetCentralRegionId() *string
	SetDBList(v string) *CreateGadInstanceMemberRequest
	GetDBList() *string
	SetGadInstanceId(v string) *CreateGadInstanceMemberRequest
	GetGadInstanceId() *string
	SetUnitNode(v []*CreateGadInstanceMemberRequestUnitNode) *CreateGadInstanceMemberRequest
	GetUnitNode() []*CreateGadInstanceMemberRequestUnitNode
}

type CreateGadInstanceMemberRequest struct {
	// The ID of the central node . You can call the DescribeGadInstances operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gad-rm-bp1npi2j8****
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
	// A JSON array that consists of the information about the databases on the central node. All database information that you specify in this array is synchronized to the unit nodes of the global active database cluster. The JSON array contains the following fields:
	//
	// 	- **name**: the name of the database.
	//
	// 	- **all**: specifies whether to synchronize all data in the database or the table. Valid values: **true*	- and **false**.
	//
	// 	- **Table**: the name of the table. If you set the **all*	- field to **false**, you must nest the name of the table that you want to synchronize into the JSON array.
	//
	// Example: `{ "testdb": { "name": "testdb", "all": false, "Table": { "order": { "name": "order", "all": true }, "ordernew": { "name": "ordernew", "all": true } } } }`
	//
	// >  For more information, see [Objects of DTS tasks](https://help.aliyun.com/document_detail/209545.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// {    "testdb": {     "name": "testdb",     "all": false,     "Table": {       "order": {         "name": "order",         "all": true       },       "ordernew": {         "name": "ordernew",         "all": true       }     }   } }
	DBList *string `json:"DBList,omitempty" xml:"DBList,omitempty"`
	// The ID of the global active database cluster. You can call the DescribeGadInstances operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gad-rm-bp1npi2j8****
	GadInstanceId *string `json:"GadInstanceId,omitempty" xml:"GadInstanceId,omitempty"`
	// The information about the unit node.
	//
	// This parameter is required.
	UnitNode []*CreateGadInstanceMemberRequestUnitNode `json:"UnitNode,omitempty" xml:"UnitNode,omitempty" type:"Repeated"`
}

func (s CreateGadInstanceMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGadInstanceMemberRequest) GoString() string {
	return s.String()
}

func (s *CreateGadInstanceMemberRequest) GetCentralDBInstanceId() *string {
	return s.CentralDBInstanceId
}

func (s *CreateGadInstanceMemberRequest) GetCentralRdsDtsAdminAccount() *string {
	return s.CentralRdsDtsAdminAccount
}

func (s *CreateGadInstanceMemberRequest) GetCentralRdsDtsAdminPassword() *string {
	return s.CentralRdsDtsAdminPassword
}

func (s *CreateGadInstanceMemberRequest) GetCentralRegionId() *string {
	return s.CentralRegionId
}

func (s *CreateGadInstanceMemberRequest) GetDBList() *string {
	return s.DBList
}

func (s *CreateGadInstanceMemberRequest) GetGadInstanceId() *string {
	return s.GadInstanceId
}

func (s *CreateGadInstanceMemberRequest) GetUnitNode() []*CreateGadInstanceMemberRequestUnitNode {
	return s.UnitNode
}

func (s *CreateGadInstanceMemberRequest) SetCentralDBInstanceId(v string) *CreateGadInstanceMemberRequest {
	s.CentralDBInstanceId = &v
	return s
}

func (s *CreateGadInstanceMemberRequest) SetCentralRdsDtsAdminAccount(v string) *CreateGadInstanceMemberRequest {
	s.CentralRdsDtsAdminAccount = &v
	return s
}

func (s *CreateGadInstanceMemberRequest) SetCentralRdsDtsAdminPassword(v string) *CreateGadInstanceMemberRequest {
	s.CentralRdsDtsAdminPassword = &v
	return s
}

func (s *CreateGadInstanceMemberRequest) SetCentralRegionId(v string) *CreateGadInstanceMemberRequest {
	s.CentralRegionId = &v
	return s
}

func (s *CreateGadInstanceMemberRequest) SetDBList(v string) *CreateGadInstanceMemberRequest {
	s.DBList = &v
	return s
}

func (s *CreateGadInstanceMemberRequest) SetGadInstanceId(v string) *CreateGadInstanceMemberRequest {
	s.GadInstanceId = &v
	return s
}

func (s *CreateGadInstanceMemberRequest) SetUnitNode(v []*CreateGadInstanceMemberRequestUnitNode) *CreateGadInstanceMemberRequest {
	s.UnitNode = v
	return s
}

func (s *CreateGadInstanceMemberRequest) Validate() error {
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

type CreateGadInstanceMemberRequestUnitNode struct {
	// The name of the unit node that you want to create. The name must meet the following requirements:
	//
	// 	- The name must be **2 to 255*	- characters in length.
	//
	// 	- The name can contain letters, digits, underscores (_), and hyphens (-) and must start with a letter.
	//
	// 	- The name cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The storage capacity of the unit node that you want to create. Unit: GB The storage capacity increases in increments of 5 GB. For more information, see [Primary ApsaraDB RDS instance types](https://help.aliyun.com/document_detail/26312.html). You can also call the DescribeAvailableResource operation to query the storage capacity range that is supported by the new instance type.
	//
	// example:
	//
	// 20
	DBInstanceStorage *int64 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// The storage type of the instance. Valid values:
	//
	// 	- **local_ssd**: local SSD
	//
	// 	- **cloud_ssd**: standard SSD
	//
	// 	- **cloud_essd**: PL1 ESSD
	//
	// 	- **cloud_essd2**: PL2 ESSD
	//
	// 	- **cloud_essd3**: PL3 ESSD
	//
	// example:
	//
	// cloud_essd
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
	// 	- **ignore**: DTS overwrites the conflicting primary key on the logger node.
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
	// The region ID of the unit node or secondary node that you want to create. You can call the DescribeRegions operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionID *string `json:"RegionID,omitempty" xml:"RegionID,omitempty"`
	// The [IP address whitelist](https://help.aliyun.com/document_detail/43185.html) of the unit node that you want to create. If you want to add more than one entry to the IP address whitelist, separate the entries with commas (,). Each entry must be unique. The IP address whitelist can contain up to 1,000 entries. The entries in the IP address whitelist must be in one of the following formats:
	//
	// 	- IP addresses, such as `10.10.XX.XX`.
	//
	// 	- CIDR blocks, such as `10.10.XX.XX/24`. In this example, **24*	- indicates that the prefix of each IP address in the IP address whitelist is 24 bits in length. You can replace 24 with a value within the range of **1 to 32**.
	//
	// example:
	//
	// 10.10.XX.XX
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The vSwitch ID of the unit node that you want to create.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1tg609m5j85****
	VSwitchID *string `json:"VSwitchID,omitempty" xml:"VSwitchID,omitempty"`
	// The virtual private cloud (VPC) ID of the unit node that you want to create.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp19ame5m1r3o****
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

func (s CreateGadInstanceMemberRequestUnitNode) String() string {
	return dara.Prettify(s)
}

func (s CreateGadInstanceMemberRequestUnitNode) GoString() string {
	return s.String()
}

func (s *CreateGadInstanceMemberRequestUnitNode) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *CreateGadInstanceMemberRequestUnitNode) GetDBInstanceStorage() *int64 {
	return s.DBInstanceStorage
}

func (s *CreateGadInstanceMemberRequestUnitNode) GetDBInstanceStorageType() *string {
	return s.DBInstanceStorageType
}

func (s *CreateGadInstanceMemberRequestUnitNode) GetDbInstanceClass() *string {
	return s.DbInstanceClass
}

func (s *CreateGadInstanceMemberRequestUnitNode) GetDtsConflict() *string {
	return s.DtsConflict
}

func (s *CreateGadInstanceMemberRequestUnitNode) GetDtsInstanceClass() *string {
	return s.DtsInstanceClass
}

func (s *CreateGadInstanceMemberRequestUnitNode) GetEngine() *string {
	return s.Engine
}

func (s *CreateGadInstanceMemberRequestUnitNode) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *CreateGadInstanceMemberRequestUnitNode) GetRegionID() *string {
	return s.RegionID
}

func (s *CreateGadInstanceMemberRequestUnitNode) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *CreateGadInstanceMemberRequestUnitNode) GetVSwitchID() *string {
	return s.VSwitchID
}

func (s *CreateGadInstanceMemberRequestUnitNode) GetVpcID() *string {
	return s.VpcID
}

func (s *CreateGadInstanceMemberRequestUnitNode) GetZoneID() *string {
	return s.ZoneID
}

func (s *CreateGadInstanceMemberRequestUnitNode) GetZoneIDSlave1() *string {
	return s.ZoneIDSlave1
}

func (s *CreateGadInstanceMemberRequestUnitNode) GetZoneIDSlave2() *string {
	return s.ZoneIDSlave2
}

func (s *CreateGadInstanceMemberRequestUnitNode) SetDBInstanceDescription(v string) *CreateGadInstanceMemberRequestUnitNode {
	s.DBInstanceDescription = &v
	return s
}

func (s *CreateGadInstanceMemberRequestUnitNode) SetDBInstanceStorage(v int64) *CreateGadInstanceMemberRequestUnitNode {
	s.DBInstanceStorage = &v
	return s
}

func (s *CreateGadInstanceMemberRequestUnitNode) SetDBInstanceStorageType(v string) *CreateGadInstanceMemberRequestUnitNode {
	s.DBInstanceStorageType = &v
	return s
}

func (s *CreateGadInstanceMemberRequestUnitNode) SetDbInstanceClass(v string) *CreateGadInstanceMemberRequestUnitNode {
	s.DbInstanceClass = &v
	return s
}

func (s *CreateGadInstanceMemberRequestUnitNode) SetDtsConflict(v string) *CreateGadInstanceMemberRequestUnitNode {
	s.DtsConflict = &v
	return s
}

func (s *CreateGadInstanceMemberRequestUnitNode) SetDtsInstanceClass(v string) *CreateGadInstanceMemberRequestUnitNode {
	s.DtsInstanceClass = &v
	return s
}

func (s *CreateGadInstanceMemberRequestUnitNode) SetEngine(v string) *CreateGadInstanceMemberRequestUnitNode {
	s.Engine = &v
	return s
}

func (s *CreateGadInstanceMemberRequestUnitNode) SetEngineVersion(v string) *CreateGadInstanceMemberRequestUnitNode {
	s.EngineVersion = &v
	return s
}

func (s *CreateGadInstanceMemberRequestUnitNode) SetRegionID(v string) *CreateGadInstanceMemberRequestUnitNode {
	s.RegionID = &v
	return s
}

func (s *CreateGadInstanceMemberRequestUnitNode) SetSecurityIPList(v string) *CreateGadInstanceMemberRequestUnitNode {
	s.SecurityIPList = &v
	return s
}

func (s *CreateGadInstanceMemberRequestUnitNode) SetVSwitchID(v string) *CreateGadInstanceMemberRequestUnitNode {
	s.VSwitchID = &v
	return s
}

func (s *CreateGadInstanceMemberRequestUnitNode) SetVpcID(v string) *CreateGadInstanceMemberRequestUnitNode {
	s.VpcID = &v
	return s
}

func (s *CreateGadInstanceMemberRequestUnitNode) SetZoneID(v string) *CreateGadInstanceMemberRequestUnitNode {
	s.ZoneID = &v
	return s
}

func (s *CreateGadInstanceMemberRequestUnitNode) SetZoneIDSlave1(v string) *CreateGadInstanceMemberRequestUnitNode {
	s.ZoneIDSlave1 = &v
	return s
}

func (s *CreateGadInstanceMemberRequestUnitNode) SetZoneIDSlave2(v string) *CreateGadInstanceMemberRequestUnitNode {
	s.ZoneIDSlave2 = &v
	return s
}

func (s *CreateGadInstanceMemberRequestUnitNode) Validate() error {
	return dara.Validate(s)
}
