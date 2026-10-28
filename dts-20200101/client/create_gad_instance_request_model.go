// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGadInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *CreateGadInstanceRequest
	GetInstanceName() *string
	SetInstanceType(v string) *CreateGadInstanceRequest
	GetInstanceType() *string
	SetMasterDbEngineType(v string) *CreateGadInstanceRequest
	GetMasterDbEngineType() *string
	SetMasterDbInstanceId(v string) *CreateGadInstanceRequest
	GetMasterDbInstanceId() *string
	SetMasterDbInstanceRegion(v string) *CreateGadInstanceRequest
	GetMasterDbInstanceRegion() *string
	SetMasterDbZoneId(v string) *CreateGadInstanceRequest
	GetMasterDbZoneId() *string
	SetOwnerId(v string) *CreateGadInstanceRequest
	GetOwnerId() *string
	SetRegionId(v string) *CreateGadInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateGadInstanceRequest
	GetResourceGroupId() *string
}

type CreateGadInstanceRequest struct {
	// The name of the group. The name must be globally unique.
	//
	// example:
	//
	// gad-sqlserver-hk-set4
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance type. Valid values: DR (disaster recovery) and GAD (active geo-redundancy).
	//
	// example:
	//
	// DR
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The database engine. Valid values: MySQL, SQL Server, and PostgreSQL.
	//
	// example:
	//
	// SQLServer
	MasterDbEngineType *string `json:"MasterDbEngineType,omitempty" xml:"MasterDbEngineType,omitempty"`
	// The ID of the primary ApsaraDB RDS instance.
	//
	// example:
	//
	// rm-j6c685vh0j0y1k***
	MasterDbInstanceId *string `json:"MasterDbInstanceId,omitempty" xml:"MasterDbInstanceId,omitempty"`
	// The region of the primary instance.
	//
	// example:
	//
	// cn-hongkong
	MasterDbInstanceRegion *string `json:"MasterDbInstanceRegion,omitempty" xml:"MasterDbInstanceRegion,omitempty"`
	// The zone of the primary instance.
	//
	// example:
	//
	// cn-hongkong-b
	MasterDbZoneId *string `json:"MasterDbZoneId,omitempty" xml:"MasterDbZoneId,omitempty"`
	OwnerId        *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region of the primary instance.
	//
	// example:
	//
	// cn-hongkong
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek3d****4mxya
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateGadInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGadInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateGadInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateGadInstanceRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateGadInstanceRequest) GetMasterDbEngineType() *string {
	return s.MasterDbEngineType
}

func (s *CreateGadInstanceRequest) GetMasterDbInstanceId() *string {
	return s.MasterDbInstanceId
}

func (s *CreateGadInstanceRequest) GetMasterDbInstanceRegion() *string {
	return s.MasterDbInstanceRegion
}

func (s *CreateGadInstanceRequest) GetMasterDbZoneId() *string {
	return s.MasterDbZoneId
}

func (s *CreateGadInstanceRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *CreateGadInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateGadInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateGadInstanceRequest) SetInstanceName(v string) *CreateGadInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateGadInstanceRequest) SetInstanceType(v string) *CreateGadInstanceRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateGadInstanceRequest) SetMasterDbEngineType(v string) *CreateGadInstanceRequest {
	s.MasterDbEngineType = &v
	return s
}

func (s *CreateGadInstanceRequest) SetMasterDbInstanceId(v string) *CreateGadInstanceRequest {
	s.MasterDbInstanceId = &v
	return s
}

func (s *CreateGadInstanceRequest) SetMasterDbInstanceRegion(v string) *CreateGadInstanceRequest {
	s.MasterDbInstanceRegion = &v
	return s
}

func (s *CreateGadInstanceRequest) SetMasterDbZoneId(v string) *CreateGadInstanceRequest {
	s.MasterDbZoneId = &v
	return s
}

func (s *CreateGadInstanceRequest) SetOwnerId(v string) *CreateGadInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateGadInstanceRequest) SetRegionId(v string) *CreateGadInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateGadInstanceRequest) SetResourceGroupId(v string) *CreateGadInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateGadInstanceRequest) Validate() error {
	return dara.Validate(s)
}
