// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyContext0SecurityIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *ModifyContext0SecurityIpsRequest
	GetDBInstanceName() *string
	SetGroupName(v string) *ModifyContext0SecurityIpsRequest
	GetGroupName() *string
	SetModifyMode(v string) *ModifyContext0SecurityIpsRequest
	GetModifyMode() *string
	SetRegionId(v string) *ModifyContext0SecurityIpsRequest
	GetRegionId() *string
	SetSecurityIPList(v string) *ModifyContext0SecurityIpsRequest
	GetSecurityIPList() *string
}

type ModifyContext0SecurityIpsRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
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
	// The ID of the region where the instance resides.
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

func (s ModifyContext0SecurityIpsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyContext0SecurityIpsRequest) GoString() string {
	return s.String()
}

func (s *ModifyContext0SecurityIpsRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifyContext0SecurityIpsRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifyContext0SecurityIpsRequest) GetModifyMode() *string {
	return s.ModifyMode
}

func (s *ModifyContext0SecurityIpsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyContext0SecurityIpsRequest) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *ModifyContext0SecurityIpsRequest) SetDBInstanceName(v string) *ModifyContext0SecurityIpsRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyContext0SecurityIpsRequest) SetGroupName(v string) *ModifyContext0SecurityIpsRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyContext0SecurityIpsRequest) SetModifyMode(v string) *ModifyContext0SecurityIpsRequest {
	s.ModifyMode = &v
	return s
}

func (s *ModifyContext0SecurityIpsRequest) SetRegionId(v string) *ModifyContext0SecurityIpsRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyContext0SecurityIpsRequest) SetSecurityIPList(v string) *ModifyContext0SecurityIpsRequest {
	s.SecurityIPList = &v
	return s
}

func (s *ModifyContext0SecurityIpsRequest) Validate() error {
	return dara.Validate(s)
}
