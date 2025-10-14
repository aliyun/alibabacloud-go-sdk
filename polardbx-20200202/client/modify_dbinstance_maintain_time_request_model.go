// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceMaintainTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyDBInstanceMaintainTimeRequest
	GetClientToken() *string
	SetDBInstanceName(v string) *ModifyDBInstanceMaintainTimeRequest
	GetDBInstanceName() *string
	SetMaintainTime(v string) *ModifyDBInstanceMaintainTimeRequest
	GetMaintainTime() *string
	SetRegionId(v string) *ModifyDBInstanceMaintainTimeRequest
	GetRegionId() *string
}

type ModifyDBInstanceMaintainTimeRequest struct {
	// example:
	//
	// 73559800-3c8c-11ec-bd40-99cfcff3fe1e
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-htri0ori2r4k9p
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 19:00Z-20:00Z
	MaintainTime *string `json:"MaintainTime,omitempty" xml:"MaintainTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDBInstanceMaintainTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceMaintainTimeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceMaintainTimeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyDBInstanceMaintainTimeRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifyDBInstanceMaintainTimeRequest) GetMaintainTime() *string {
	return s.MaintainTime
}

func (s *ModifyDBInstanceMaintainTimeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetClientToken(v string) *ModifyDBInstanceMaintainTimeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetDBInstanceName(v string) *ModifyDBInstanceMaintainTimeRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetMaintainTime(v string) *ModifyDBInstanceMaintainTimeRequest {
	s.MaintainTime = &v
	return s
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetRegionId(v string) *ModifyDBInstanceMaintainTimeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBInstanceMaintainTimeRequest) Validate() error {
	return dara.Validate(s)
}
