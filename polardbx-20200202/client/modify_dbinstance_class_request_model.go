// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceClassRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlignStoragePrimaryAzone(v bool) *ModifyDBInstanceClassRequest
	GetAlignStoragePrimaryAzone() *bool
	SetClientToken(v string) *ModifyDBInstanceClassRequest
	GetClientToken() *string
	SetCnClass(v string) *ModifyDBInstanceClassRequest
	GetCnClass() *string
	SetDBInstanceName(v string) *ModifyDBInstanceClassRequest
	GetDBInstanceName() *string
	SetDnClass(v string) *ModifyDBInstanceClassRequest
	GetDnClass() *string
	SetDnStorageSpace(v string) *ModifyDBInstanceClassRequest
	GetDnStorageSpace() *string
	SetRegionId(v string) *ModifyDBInstanceClassRequest
	GetRegionId() *string
	SetSpecifiedDNScale(v bool) *ModifyDBInstanceClassRequest
	GetSpecifiedDNScale() *bool
	SetSpecifiedDNSpecMapJson(v string) *ModifyDBInstanceClassRequest
	GetSpecifiedDNSpecMapJson() *string
	SetStorageType(v string) *ModifyDBInstanceClassRequest
	GetStorageType() *string
	SetSwitchTime(v string) *ModifyDBInstanceClassRequest
	GetSwitchTime() *string
	SetSwitchTimeMode(v string) *ModifyDBInstanceClassRequest
	GetSwitchTimeMode() *string
	SetTargetDBInstanceClass(v string) *ModifyDBInstanceClassRequest
	GetTargetDBInstanceClass() *string
}

type ModifyDBInstanceClassRequest struct {
	AlignStoragePrimaryAzone *bool `json:"AlignStoragePrimaryAzone,omitempty" xml:"AlignStoragePrimaryAzone,omitempty"`
	// The client token. It is an arbitrary unique string used to ensure the idempotence of the request.
	//
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// **Target specifications for Enterprise Edition compute nodes**
	//
	// **Primary instance compute node specifications (Enterprise Edition CN) (General-purpose)**
	//
	// - polarx.x4.medium.2e 2 vCPUs, 8 GB (General-purpose)
	//
	// - polarx.x4.large.2e 4 vCPUs, 16 GB (General-purpose)
	//
	// - polarx.x4.xlarge.2e 8 vCPUs, 32 GB (General-purpose)
	//
	// - polarx.x4.2xlarge.2e 16 vCPUs, 64 GB (General-purpose)
	//
	// **Primary instance compute node specifications (Enterprise Edition CN) (Dedicated)**
	//
	// - polarx.x8.large.2e 4 vCPUs, 32 GB (Dedicated)
	//
	// - polarx.x8.xlarge.2e 8 vCPUs, 64 GB (Dedicated)
	//
	// - polarx.x8.2xlarge.2e 16 vCPUs, 128 GB (Dedicated)
	//
	// - polarx.x4.4xlarge.2e 32 vCPUs, 128 GB (Dedicated)
	//
	// - polarx.x8.4xlarge.2e 32 vCPUs, 256 GB (Dedicated)
	//
	// - polarx.st.8xlarge.2e 60 vCPUs, 470 GB (Dedicated)
	//
	// - polarx.st.12xlarge.2e 90 vCPUs, 720 GB (Dedicated)
	//
	// **Read-only instance compute node specifications (Enterprise Edition CN) (General-purpose)**
	//
	// - polarxro.x4.medium.2e 2 vCPUs, 8 GB (General-purpose)
	//
	// - polarxro.x4.large.2e 4 vCPUs, 16 GB (General-purpose)
	//
	// - polarxro.x4.xlarge.2e 8 vCPUs, 32 GB (General-purpose)
	//
	// - polarxro.x4.2xlarge.2e 16 vCPUs, 64 GB (General-purpose)
	//
	// **Read-only instance compute node specifications (Enterprise Edition CN) (Dedicated)**
	//
	// - polarxro.x8.large.2e 4 vCPUs, 32 GB (Dedicated)
	//
	// - polarxro.x8.xlarge.2e 8 vCPUs, 64 GB (Dedicated)
	//
	// - polarxro.x8.2xlarge.2e 16 vCPUs, 128 GB (Dedicated)
	//
	// - polarxro.x4.4xlarge.2e 32 vCPUs, 128 GB (Dedicated)
	//
	// - polarxro.x8.4xlarge.2e 32 vCPUs, 256 GB (Dedicated)
	//
	// - polarxro.st.8xlarge.2e 60 vCPUs, 470 GB (Dedicated physical host)
	//
	// - polarxro.st.12xlarge.2e 90 vCPUs, 720 GB (Dedicated physical host)
	//
	// example:
	//
	// polarx.x4.medium.2e
	CnClass *string `json:"CnClass,omitempty" xml:"CnClass,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzjasd****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// **Target specifications for Enterprise Edition storage nodes**
	//
	// **Storage node specifications (Enterprise Edition DN) (General-purpose)**
	//
	// - mysql.n4.medium.25 2 vCPUs, 8 GB (General-purpose)
	//
	// - mysql.n4.large.25 4 vCPUs, 16 GB (General-purpose)
	//
	// - mysql.n4.xlarge.25 8 vCPUs, 32 GB (General-purpose)
	//
	// - mysql.n4.2xlarge.25 16 vCPUs, 64 GB (General-purpose)
	//
	// **Storage node specifications (Enterprise Edition DN) (Dedicated)**
	//
	// - mysql.x8.large.25 4 vCPUs, 32 GB (Dedicated)
	//
	// - mysql.x8.xlarge.25 8 vCPUs, 64 GB (Dedicated)
	//
	// - mysql.x8.2xlarge.25 16 vCPUs, 128 GB (Dedicated)
	//
	// - mysql.x4.4xlarge.25 32 vCPUs, 128 GB (Dedicated)
	//
	// - mysql.x8.4xlarge.25 32 vCPUs, 256 GB (Dedicated)
	//
	// - mysql.st.8xlarge.25 60 vCPUs, 470 GB (Dedicated)
	//
	// - mysql.st.12xlarge.25 90 vCPUs, 720 GB (Dedicated)
	//
	// **Read-only instance storage node specifications (Enterprise Edition DN) (General-purpose)**
	//
	// - rds.mysql.s2.xlarge 2 vCPUs, 8 GB (General-purpose)
	//
	// - mysqlro.x4.large.1 4 vCPUs, 16 GB (General-purpose)
	//
	// - mysqlro.x4.xlarge.1 8 vCPUs, 32 GB (General-purpose)
	//
	// - mysqlro.x4.2xlarge.1 16 vCPUs, 64 GB (General-purpose)
	//
	// **Read-only instance storage node specifications (Enterprise Edition DN) (Dedicated)**
	//
	// - mysqlro.x8.large.1 4 vCPUs, 32 GB (Dedicated)
	//
	// - mysqlro.x8.xlarge.1 8 vCPUs, 64 GB (Dedicated)
	//
	// - mysqlro.x8.2xlarge.1 16 vCPUs, 128 GB (Dedicated)
	//
	// - mysqlro.x4.4xlarge.1 32 vCPUs, 128 GB (Dedicated)
	//
	// - mysqlro.x8.4xlarge.1 32 vCPUs, 256 GB (Dedicated)
	//
	// - rds.mysql.st.h43 60 vCPUs, 470 GB (Dedicated physical host)
	//
	// - rds.mysql.st.v52 90 vCPUs, 720 GB (Dedicated physical host)
	//
	// example:
	//
	// mysql.n4.medium.25
	DnClass *string `json:"DnClass,omitempty" xml:"DnClass,omitempty"`
	// The target cloud disk size after the specification change.
	//
	// - If the instance uses a fixed-size cloud disk, you cannot change the cloud disk size. In this case, do not specify DnStorageSpace.
	//
	// - Valid values of DnStorageSpace: 20 to 3072.
	//
	// example:
	//
	// 50
	DnStorageSpace *string `json:"DnStorageSpace,omitempty" xml:"DnStorageSpace,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to change the DN nodes to different specifications.
	//
	// example:
	//
	// true
	SpecifiedDNScale *bool `json:"SpecifiedDNScale,omitempty" xml:"SpecifiedDNScale,omitempty"`
	// The target specification for each DN node when DN nodes are changed to different specifications.
	//
	// example:
	//
	// {"pxc-xdb-s-htr3rh44ki3s6a4354":"mysql.n4.medium.25","pxc-xdb-s-htr3rh44ki3s6a71c6":"mysql.n4.medium.25"}
	SpecifiedDNSpecMapJson *string `json:"SpecifiedDNSpecMapJson,omitempty" xml:"SpecifiedDNSpecMapJson,omitempty"`
	// The target storage type.
	//
	// example:
	//
	// cloud_auto
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The start time of the switchover. The switchover is performed within the period from the start time T to T+30 minutes. This parameter is not available.
	//
	// example:
	//
	// 2024-12-11T17:10:00Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// The switching time. Valid values:
	//
	// - 0: Immediately.
	//
	// - 1: Within the O&M window.
	//
	// example:
	//
	// 0
	SwitchTimeMode *string `json:"SwitchTimeMode,omitempty" xml:"SwitchTimeMode,omitempty"`
	// **Target specifications for Standard Edition**
	//
	// **Primary instance node specifications (Standard Edition) (General-purpose):**
	//
	// - mysql.n2.medium.25 2 vCPUs, 4 GB (General-purpose)
	//
	// - mysql.n4.medium.25 2 vCPUs, 8 GB (General-purpose)
	//
	// - mysql.n8.medium.25 2 vCPUs, 16 GB (General-purpose)
	//
	// - mysql.n2.large.25 4 vCPUs, 8 GB (General-purpose)
	//
	// - mysql.n4.large.25 4 vCPUs, 16 GB (General-purpose)
	//
	// - mysql.n8.large.25 4 vCPUs, 32 GB (General-purpose)
	//
	// - mysql.n2.xlarge.25 8 vCPUs, 16 GB (General-purpose)
	//
	// - mysql.n4.xlarge.25 8 vCPUs, 32 GB (General-purpose)
	//
	// - mysql.n8.xlarge.25 8 vCPUs, 64 GB (General-purpose)
	//
	// - mysql.n2.2xlarge.25 16 vCPUs, 32 GB (General-purpose)
	//
	// - mysql.n4.2xlarge.25 16 vCPUs, 64 GB (General-purpose)
	//
	// - mysql.n8.2xlarge.25 16 vCPUs, 128 GB (General-purpose)
	//
	// **Primary instance node specifications (Standard Edition) (Dedicated):**
	//
	// - mysql.x2.medium.25 2 vCPUs, 4 GB (Dedicated)
	//
	// - mysql.x4.medium.25 2 vCPUs, 8 GB (Dedicated)
	//
	// - mysql.x8.medium.25 2 vCPUs, 16 GB (Dedicated)
	//
	// - mysql.x2.large.25 4 vCPUs, 8 GB (Dedicated)
	//
	// - mysql.x4.large.25 4 vCPUs, 16 GB (Dedicated)
	//
	// - mysql.x8.large.25 4 vCPUs, 32 GB (Dedicated)
	//
	// - mysql.x2.xlarge.25 8 vCPUs, 16 GB (Dedicated)
	//
	// - mysql.x4.xlarge.25 8 vCPUs, 32 GB (Dedicated)
	//
	// - mysql.x8.xlarge.25 8 vCPUs, 64 GB (Dedicated)
	//
	// - mysql.x2.2xlarge.25 16 vCPUs, 32 GB (Dedicated)
	//
	// - mysql.x4.2xlarge.25 16 vCPUs, 64 GB (Dedicated)
	//
	// - mysql.x8.2xlarge.25 16 vCPUs, 128 GB (Dedicated)
	//
	// **Read-only instance node specifications (Standard Edition) (General-purpose)**
	//
	// - rds.mysql.s2.xlarge 2 vCPUs, 8 GB (General-purpose)
	//
	// - mysqlro.x4.large.1 4 vCPUs, 16 GB (General-purpose)
	//
	// - mysqlro.x4.xlarge.1 8 vCPUs, 32 GB (General-purpose)
	//
	// - mysqlro.x4.2xlarge.1 16 vCPUs, 64 GB (General-purpose)
	//
	// **Read-only instance node specifications (Standard Edition) (Dedicated)**
	//
	// - mysqlro.x8.large.1 4 vCPUs, 32 GB (Dedicated)
	//
	// - mysqlro.x8.xlarge.1 8 vCPUs, 64 GB (Dedicated)
	//
	// - mysqlro.x8.2xlarge.1 16 vCPUs, 128 GB (Dedicated)
	//
	// - mysqlro.x4.4xlarge.1 32 vCPUs, 128 GB (Dedicated)
	//
	// - mysqlro.x8.4xlarge.1 32 vCPUs, 256 GB (Dedicated)
	//
	// example:
	//
	// mysql.n2.medium.25
	TargetDBInstanceClass *string `json:"TargetDBInstanceClass,omitempty" xml:"TargetDBInstanceClass,omitempty"`
}

func (s ModifyDBInstanceClassRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceClassRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceClassRequest) GetAlignStoragePrimaryAzone() *bool {
	return s.AlignStoragePrimaryAzone
}

func (s *ModifyDBInstanceClassRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyDBInstanceClassRequest) GetCnClass() *string {
	return s.CnClass
}

func (s *ModifyDBInstanceClassRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifyDBInstanceClassRequest) GetDnClass() *string {
	return s.DnClass
}

func (s *ModifyDBInstanceClassRequest) GetDnStorageSpace() *string {
	return s.DnStorageSpace
}

func (s *ModifyDBInstanceClassRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDBInstanceClassRequest) GetSpecifiedDNScale() *bool {
	return s.SpecifiedDNScale
}

func (s *ModifyDBInstanceClassRequest) GetSpecifiedDNSpecMapJson() *string {
	return s.SpecifiedDNSpecMapJson
}

func (s *ModifyDBInstanceClassRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *ModifyDBInstanceClassRequest) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *ModifyDBInstanceClassRequest) GetSwitchTimeMode() *string {
	return s.SwitchTimeMode
}

func (s *ModifyDBInstanceClassRequest) GetTargetDBInstanceClass() *string {
	return s.TargetDBInstanceClass
}

func (s *ModifyDBInstanceClassRequest) SetAlignStoragePrimaryAzone(v bool) *ModifyDBInstanceClassRequest {
	s.AlignStoragePrimaryAzone = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) SetClientToken(v string) *ModifyDBInstanceClassRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) SetCnClass(v string) *ModifyDBInstanceClassRequest {
	s.CnClass = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) SetDBInstanceName(v string) *ModifyDBInstanceClassRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) SetDnClass(v string) *ModifyDBInstanceClassRequest {
	s.DnClass = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) SetDnStorageSpace(v string) *ModifyDBInstanceClassRequest {
	s.DnStorageSpace = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) SetRegionId(v string) *ModifyDBInstanceClassRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) SetSpecifiedDNScale(v bool) *ModifyDBInstanceClassRequest {
	s.SpecifiedDNScale = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) SetSpecifiedDNSpecMapJson(v string) *ModifyDBInstanceClassRequest {
	s.SpecifiedDNSpecMapJson = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) SetStorageType(v string) *ModifyDBInstanceClassRequest {
	s.StorageType = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) SetSwitchTime(v string) *ModifyDBInstanceClassRequest {
	s.SwitchTime = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) SetSwitchTimeMode(v string) *ModifyDBInstanceClassRequest {
	s.SwitchTimeMode = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) SetTargetDBInstanceClass(v string) *ModifyDBInstanceClassRequest {
	s.TargetDBInstanceClass = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) Validate() error {
	return dara.Validate(s)
}
