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
	// This parameter is required.
	CentralDBInstanceId *string `json:"CentralDBInstanceId,omitempty" xml:"CentralDBInstanceId,omitempty"`
	// This parameter is required.
	CentralRdsDtsAdminAccount *string `json:"CentralRdsDtsAdminAccount,omitempty" xml:"CentralRdsDtsAdminAccount,omitempty"`
	// This parameter is required.
	CentralRdsDtsAdminPassword *string `json:"CentralRdsDtsAdminPassword,omitempty" xml:"CentralRdsDtsAdminPassword,omitempty"`
	// This parameter is required.
	CentralRegionId *string `json:"CentralRegionId,omitempty" xml:"CentralRegionId,omitempty"`
	// This parameter is required.
	DBList *string `json:"DBList,omitempty" xml:"DBList,omitempty"`
	// This parameter is required.
	GadInstanceId *string `json:"GadInstanceId,omitempty" xml:"GadInstanceId,omitempty"`
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
	// This parameter is required.
	RegionID       *string `json:"RegionID,omitempty" xml:"RegionID,omitempty"`
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// This parameter is required.
	VSwitchID *string `json:"VSwitchID,omitempty" xml:"VSwitchID,omitempty"`
	// This parameter is required.
	VpcID        *string `json:"VpcID,omitempty" xml:"VpcID,omitempty"`
	ZoneID       *string `json:"ZoneID,omitempty" xml:"ZoneID,omitempty"`
	ZoneIDSlave1 *string `json:"ZoneIDSlave1,omitempty" xml:"ZoneIDSlave1,omitempty"`
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
