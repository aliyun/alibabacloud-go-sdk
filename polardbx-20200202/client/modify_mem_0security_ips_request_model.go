// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMem0SecurityIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *ModifyMem0SecurityIpsRequest
	GetDBInstanceName() *string
	SetGroupName(v string) *ModifyMem0SecurityIpsRequest
	GetGroupName() *string
	SetModifyMode(v string) *ModifyMem0SecurityIpsRequest
	GetModifyMode() *string
	SetRegionId(v string) *ModifyMem0SecurityIpsRequest
	GetRegionId() *string
	SetSecurityIPList(v string) *ModifyMem0SecurityIpsRequest
	GetSecurityIPList() *string
}

type ModifyMem0SecurityIpsRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-*********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The name of the whitelist group for the instance.
	//
	// example:
	//
	// default
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The modification mode of the whitelist. Valid values:
	//
	// - 0: overwrites the whitelist group.
	//
	// - 1: adds a whitelist group.
	//
	// - 2: deletes a whitelist group.
	//
	// example:
	//
	// 1
	ModifyMode *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IP addresses in the whitelist group. Separate multiple IP addresses with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 127.0.0.1,192.168.0.0
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
}

func (s ModifyMem0SecurityIpsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyMem0SecurityIpsRequest) GoString() string {
	return s.String()
}

func (s *ModifyMem0SecurityIpsRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifyMem0SecurityIpsRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifyMem0SecurityIpsRequest) GetModifyMode() *string {
	return s.ModifyMode
}

func (s *ModifyMem0SecurityIpsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyMem0SecurityIpsRequest) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *ModifyMem0SecurityIpsRequest) SetDBInstanceName(v string) *ModifyMem0SecurityIpsRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyMem0SecurityIpsRequest) SetGroupName(v string) *ModifyMem0SecurityIpsRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyMem0SecurityIpsRequest) SetModifyMode(v string) *ModifyMem0SecurityIpsRequest {
	s.ModifyMode = &v
	return s
}

func (s *ModifyMem0SecurityIpsRequest) SetRegionId(v string) *ModifyMem0SecurityIpsRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyMem0SecurityIpsRequest) SetSecurityIPList(v string) *ModifyMem0SecurityIpsRequest {
	s.SecurityIPList = &v
	return s
}

func (s *ModifyMem0SecurityIpsRequest) Validate() error {
	return dara.Validate(s)
}
