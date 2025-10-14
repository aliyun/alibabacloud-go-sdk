// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceTopologyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeDBInstanceTopologyResponseBodyData) *DescribeDBInstanceTopologyResponseBody
	GetData() *DescribeDBInstanceTopologyResponseBodyData
	SetRequestId(v string) *DescribeDBInstanceTopologyResponseBody
	GetRequestId() *string
}

type DescribeDBInstanceTopologyResponseBody struct {
	Data *DescribeDBInstanceTopologyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// C458B1E8-1683-3645-B154-6BA32080EEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceTopologyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceTopologyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceTopologyResponseBody) GetData() *DescribeDBInstanceTopologyResponseBodyData {
	return s.Data
}

func (s *DescribeDBInstanceTopologyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceTopologyResponseBody) SetData(v *DescribeDBInstanceTopologyResponseBodyData) *DescribeDBInstanceTopologyResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBody) SetRequestId(v string) *DescribeDBInstanceTopologyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstanceTopologyResponseBodyData struct {
	LogicInstanceTopology *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology `json:"LogicInstanceTopology,omitempty" xml:"LogicInstanceTopology,omitempty" type:"Struct"`
}

func (s DescribeDBInstanceTopologyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceTopologyResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceTopologyResponseBodyData) GetLogicInstanceTopology() *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology {
	return s.LogicInstanceTopology
}

func (s *DescribeDBInstanceTopologyResponseBodyData) SetLogicInstanceTopology(v *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) *DescribeDBInstanceTopologyResponseBodyData {
	s.LogicInstanceTopology = v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyData) Validate() error {
	if s.LogicInstanceTopology != nil {
		if err := s.LogicInstanceTopology.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology struct {
	// example:
	//
	// lvs
	DBInstanceConnType *string `json:"DBInstanceConnType,omitempty" xml:"DBInstanceConnType,omitempty"`
	// example:
	//
	// 2021-10-21T10:30:45Z 04:00:00
	DBInstanceCreateTime *string `json:"DBInstanceCreateTime,omitempty" xml:"DBInstanceCreateTime,omitempty"`
	// example:
	//
	// pxc-sprcym7g7wj7k
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// example:
	//
	// 304726047
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// pxc-sprcym7g7w****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// 8
	DBInstanceStatus *int32 `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	// example:
	//
	// TDE_MODIFYING
	DBInstanceStatusDescription *string `json:"DBInstanceStatusDescription,omitempty" xml:"DBInstanceStatusDescription,omitempty"`
	// example:
	//
	// 1
	DBInstanceStorage *int32 `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	// example:
	//
	// polarx
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// example:
	//
	// 2.0
	EngineVersion *string                                                                        `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	HistoryItems  []*DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems `json:"HistoryItems,omitempty" xml:"HistoryItems,omitempty" type:"Repeated"`
	Items         []*DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems        `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	LockMode   *int32  `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// example:
	//
	// 05:00:00
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// example:
	//
	// 04:00:00
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
}

func (s DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) GetDBInstanceConnType() *string {
	return s.DBInstanceConnType
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) GetDBInstanceCreateTime() *string {
	return s.DBInstanceCreateTime
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) GetDBInstanceStatus() *int32 {
	return s.DBInstanceStatus
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) GetDBInstanceStatusDescription() *string {
	return s.DBInstanceStatusDescription
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) GetDBInstanceStorage() *int32 {
	return s.DBInstanceStorage
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) GetHistoryItems() []*DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems {
	return s.HistoryItems
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) GetItems() []*DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	return s.Items
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) GetLockMode() *int32 {
	return s.LockMode
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) GetMaintainEndTime() *string {
	return s.MaintainEndTime
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) GetMaintainStartTime() *string {
	return s.MaintainStartTime
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) SetDBInstanceConnType(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology {
	s.DBInstanceConnType = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) SetDBInstanceCreateTime(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology {
	s.DBInstanceCreateTime = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) SetDBInstanceDescription(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) SetDBInstanceId(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) SetDBInstanceName(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) SetDBInstanceStatus(v int32) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) SetDBInstanceStatusDescription(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology {
	s.DBInstanceStatusDescription = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) SetDBInstanceStorage(v int32) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology {
	s.DBInstanceStorage = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) SetEngine(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) SetEngineVersion(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) SetHistoryItems(v []*DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology {
	s.HistoryItems = v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) SetItems(v []*DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology {
	s.Items = v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) SetLockMode(v int32) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) SetLockReason(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology {
	s.LockReason = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) SetMaintainEndTime(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology {
	s.MaintainEndTime = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) SetMaintainStartTime(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology {
	s.MaintainStartTime = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopology) Validate() error {
	if s.HistoryItems != nil {
		for _, item := range s.HistoryItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems struct {
	Activated       *bool   `json:"Activated,omitempty" xml:"Activated,omitempty"`
	Azone           *string `json:"Azone,omitempty" xml:"Azone,omitempty"`
	CharacterType   *string `json:"CharacterType,omitempty" xml:"CharacterType,omitempty"`
	DBInstanceId    *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceName  *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	PhyInstanceName *string `json:"PhyInstanceName,omitempty" xml:"PhyInstanceName,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Role            *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems) GetActivated() *bool {
	return s.Activated
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems) GetAzone() *string {
	return s.Azone
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems) GetCharacterType() *string {
	return s.CharacterType
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems) GetPhyInstanceName() *string {
	return s.PhyInstanceName
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems) GetRegion() *string {
	return s.Region
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems) GetRole() *string {
	return s.Role
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems) SetActivated(v bool) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems {
	s.Activated = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems) SetAzone(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems {
	s.Azone = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems) SetCharacterType(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems {
	s.CharacterType = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems) SetDBInstanceId(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems) SetDBInstanceName(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems) SetPhyInstanceName(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems {
	s.PhyInstanceName = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems) SetRegion(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems {
	s.Region = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems) SetRole(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems {
	s.Role = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyHistoryItems) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems struct {
	// example:
	//
	// true
	Activated *bool `json:"Activated,omitempty" xml:"Activated,omitempty"`
	// example:
	//
	// cn-hangzhou-a
	Azone              *string                                                                              `json:"Azone,omitempty" xml:"Azone,omitempty"`
	AzoneRoleList      []*DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsAzoneRoleList `json:"AzoneRoleList,omitempty" xml:"AzoneRoleList,omitempty" type:"Repeated"`
	CharacterType      *string                                                                              `json:"CharacterType,omitempty" xml:"CharacterType,omitempty"`
	ConnectionIp       []*DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsConnectionIp  `json:"ConnectionIp,omitempty" xml:"ConnectionIp,omitempty" type:"Repeated"`
	DBInstanceConnType *int32                                                                               `json:"DBInstanceConnType,omitempty" xml:"DBInstanceConnType,omitempty"`
	// example:
	//
	// 2021-10-21T10:30:45Z
	DBInstanceCreateTime  *string `json:"DBInstanceCreateTime,omitempty" xml:"DBInstanceCreateTime,omitempty"`
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// example:
	//
	// 304726049
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// pxc-i-tk6t4z****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// 8
	DBInstanceStatus            *int32  `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	DBInstanceStatusDescription *string `json:"DBInstanceStatusDescription,omitempty" xml:"DBInstanceStatusDescription,omitempty"`
	// example:
	//
	// 3145728
	DiskSize *int64 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// example:
	//
	// mysql
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// example:
	//
	// 5.7
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// example:
	//
	// 0
	LockMode          *int32  `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	LockReason        *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	MaintainEndTime   *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// example:
	//
	// 4000
	MaxConnections *int32 `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	// example:
	//
	// 7000
	MaxIops *int32 `json:"MaxIops,omitempty" xml:"MaxIops,omitempty"`
	// example:
	//
	// polarx.x4.large.2e
	NodeClass       *string `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	PhyInstanceName *string `json:"PhyInstanceName,omitempty" xml:"PhyInstanceName,omitempty"`
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// master
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// 0
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageUsed *string `json:"StorageUsed,omitempty" xml:"StorageUsed,omitempty"`
	// example:
	//
	// polarx-cdc-kernel-2.0.0-3985896
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) GetActivated() *bool {
	return s.Activated
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) GetAzone() *string {
	return s.Azone
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) GetAzoneRoleList() []*DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsAzoneRoleList {
	return s.AzoneRoleList
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) GetCharacterType() *string {
	return s.CharacterType
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) GetConnectionIp() []*DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsConnectionIp {
	return s.ConnectionIp
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) GetDBInstanceConnType() *int32 {
	return s.DBInstanceConnType
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) GetDBInstanceCreateTime() *string {
	return s.DBInstanceCreateTime
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) GetDBInstanceStatus() *int32 {
	return s.DBInstanceStatus
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) GetDBInstanceStatusDescription() *string {
	return s.DBInstanceStatusDescription
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) GetDiskSize() *int64 {
	return s.DiskSize
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) GetLockMode() *int32 {
	return s.LockMode
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) GetMaintainEndTime() *string {
	return s.MaintainEndTime
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) GetMaintainStartTime() *string {
	return s.MaintainStartTime
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) GetMaxConnections() *int32 {
	return s.MaxConnections
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) GetMaxIops() *int32 {
	return s.MaxIops
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) GetNodeClass() *string {
	return s.NodeClass
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) GetPhyInstanceName() *string {
	return s.PhyInstanceName
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) GetRegion() *string {
	return s.Region
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) GetRole() *string {
	return s.Role
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) GetStorageUsed() *string {
	return s.StorageUsed
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) GetVersion() *string {
	return s.Version
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetActivated(v bool) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.Activated = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetAzone(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.Azone = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetAzoneRoleList(v []*DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsAzoneRoleList) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.AzoneRoleList = v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetCharacterType(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.CharacterType = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetConnectionIp(v []*DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsConnectionIp) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.ConnectionIp = v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetDBInstanceConnType(v int32) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.DBInstanceConnType = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetDBInstanceCreateTime(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.DBInstanceCreateTime = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetDBInstanceDescription(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetDBInstanceId(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetDBInstanceName(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetDBInstanceStatus(v int32) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetDBInstanceStatusDescription(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.DBInstanceStatusDescription = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetDiskSize(v int64) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.DiskSize = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetEngine(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetEngineVersion(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetLockMode(v int32) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetLockReason(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.LockReason = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetMaintainEndTime(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.MaintainEndTime = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetMaintainStartTime(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.MaintainStartTime = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetMaxConnections(v int32) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.MaxConnections = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetMaxIops(v int32) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.MaxIops = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetNodeClass(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.NodeClass = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetPhyInstanceName(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.PhyInstanceName = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetRegion(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.Region = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetRole(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.Role = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetStatus(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.Status = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetStorageUsed(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.StorageUsed = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) SetVersion(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems {
	s.Version = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItems) Validate() error {
	if s.AzoneRoleList != nil {
		for _, item := range s.AzoneRoleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ConnectionIp != nil {
		for _, item := range s.ConnectionIp {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsAzoneRoleList struct {
	// example:
	//
	// cn-hangzhou-a
	Azone *string `json:"Azone,omitempty" xml:"Azone,omitempty"`
	// example:
	//
	// leader
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsAzoneRoleList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsAzoneRoleList) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsAzoneRoleList) GetAzone() *string {
	return s.Azone
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsAzoneRoleList) GetRole() *string {
	return s.Role
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsAzoneRoleList) SetAzone(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsAzoneRoleList {
	s.Azone = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsAzoneRoleList) SetRole(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsAzoneRoleList {
	s.Role = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsAzoneRoleList) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsConnectionIp struct {
	// example:
	//
	// pxc-xdb-m-pxcdym7g7w********.mysql.singapore.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// example:
	//
	// 1
	DBInstanceNetType *int32 `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsConnectionIp) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsConnectionIp) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsConnectionIp) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsConnectionIp) GetDBInstanceNetType() *int32 {
	return s.DBInstanceNetType
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsConnectionIp) GetPort() *string {
	return s.Port
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsConnectionIp) SetConnectionString(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsConnectionIp {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsConnectionIp) SetDBInstanceNetType(v int32) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsConnectionIp {
	s.DBInstanceNetType = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsConnectionIp) SetPort(v string) *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsConnectionIp {
	s.Port = &v
	return s
}

func (s *DescribeDBInstanceTopologyResponseBodyDataLogicInstanceTopologyItemsConnectionIp) Validate() error {
	return dara.Validate(s)
}
