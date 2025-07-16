// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceClassRequest interface {
	dara.Model
	String() string
	GoString() string
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
	SetSwitchTime(v string) *ModifyDBInstanceClassRequest
	GetSwitchTime() *string
	SetSwitchTimeMode(v string) *ModifyDBInstanceClassRequest
	GetSwitchTimeMode() *string
	SetTargetDBInstanceClass(v string) *ModifyDBInstanceClassRequest
	GetTargetDBInstanceClass() *string
}

type ModifyDBInstanceClassRequest struct {
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CnClass     *string `json:"CnClass,omitempty" xml:"CnClass,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzjasd****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	DnClass        *string `json:"DnClass,omitempty" xml:"DnClass,omitempty"`
	DnStorageSpace *string `json:"DnStorageSpace,omitempty" xml:"DnStorageSpace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SpecifiedDNScale       *bool   `json:"SpecifiedDNScale,omitempty" xml:"SpecifiedDNScale,omitempty"`
	SpecifiedDNSpecMapJson *string `json:"SpecifiedDNSpecMapJson,omitempty" xml:"SpecifiedDNSpecMapJson,omitempty"`
	SwitchTime             *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	SwitchTimeMode         *string `json:"SwitchTimeMode,omitempty" xml:"SwitchTimeMode,omitempty"`
	// example:
	//
	// polarx.x4.xlarge.2e
	TargetDBInstanceClass *string `json:"TargetDBInstanceClass,omitempty" xml:"TargetDBInstanceClass,omitempty"`
}

func (s ModifyDBInstanceClassRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceClassRequest) GoString() string {
	return s.String()
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

func (s *ModifyDBInstanceClassRequest) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *ModifyDBInstanceClassRequest) GetSwitchTimeMode() *string {
	return s.SwitchTimeMode
}

func (s *ModifyDBInstanceClassRequest) GetTargetDBInstanceClass() *string {
	return s.TargetDBInstanceClass
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
