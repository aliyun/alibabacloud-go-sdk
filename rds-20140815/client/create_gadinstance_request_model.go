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
	// This parameter is required.
	CentralDBInstanceId *string `json:"CentralDBInstanceId,omitempty" xml:"CentralDBInstanceId,omitempty"`
	// This parameter is required.
	CentralRdsDtsAdminAccount *string `json:"CentralRdsDtsAdminAccount,omitempty" xml:"CentralRdsDtsAdminAccount,omitempty"`
	// This parameter is required.
	CentralRdsDtsAdminPassword *string `json:"CentralRdsDtsAdminPassword,omitempty" xml:"CentralRdsDtsAdminPassword,omitempty"`
	// This parameter is required.
	CentralRegionId *string `json:"CentralRegionId,omitempty" xml:"CentralRegionId,omitempty"`
	// This parameter is required.
	DBList          *string                        `json:"DBList,omitempty" xml:"DBList,omitempty"`
	Description     *string                        `json:"Description,omitempty" xml:"Description,omitempty"`
	ResourceGroupId *string                        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag             []*CreateGADInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	DBInstanceStorage     *int64  `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty"`
	DBInstanceStorageType *string `json:"DBInstanceStorageType,omitempty" xml:"DBInstanceStorageType,omitempty"`
	DbInstanceClass       *string `json:"DbInstanceClass,omitempty" xml:"DbInstanceClass,omitempty"`
	// This parameter is required.
	DtsConflict *string `json:"DtsConflict,omitempty" xml:"DtsConflict,omitempty"`
	// This parameter is required.
	DtsInstanceClass *string `json:"DtsInstanceClass,omitempty" xml:"DtsInstanceClass,omitempty"`
	Engine           *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	EngineVersion    *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	PayType          *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// This parameter is required.
	RegionID       *string `json:"RegionID,omitempty" xml:"RegionID,omitempty"`
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	VSwitchID      *string `json:"VSwitchID,omitempty" xml:"VSwitchID,omitempty"`
	VpcID          *string `json:"VpcID,omitempty" xml:"VpcID,omitempty"`
	ZoneID         *string `json:"ZoneID,omitempty" xml:"ZoneID,omitempty"`
	ZoneIDSlave1   *string `json:"ZoneIDSlave1,omitempty" xml:"ZoneIDSlave1,omitempty"`
	ZoneIDSlave2   *string `json:"ZoneIDSlave2,omitempty" xml:"ZoneIDSlave2,omitempty"`
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
