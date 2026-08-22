// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOpenSearchWhitelistGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *ModifyOpenSearchWhitelistGroupRequest
	GetDBInstanceName() *string
	SetGroupId(v string) *ModifyOpenSearchWhitelistGroupRequest
	GetGroupId() *string
	SetIPs(v string) *ModifyOpenSearchWhitelistGroupRequest
	GetIPs() *string
	SetRegionId(v string) *ModifyOpenSearchWhitelistGroupRequest
	GetRegionId() *string
	SetRemark(v string) *ModifyOpenSearchWhitelistGroupRequest
	GetRemark() *string
}

type ModifyOpenSearchWhitelistGroupRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-xxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_y6sqhtr6jyu52w4oublg3674du
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The list of source IP address restrictions.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.1.0/24
	IPs *string `json:"IPs,omitempty" xml:"IPs,omitempty"`
	// The region in which the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remarks.
	//
	// example:
	//
	// 000G2DJ1YGee321432
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s ModifyOpenSearchWhitelistGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyOpenSearchWhitelistGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyOpenSearchWhitelistGroupRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifyOpenSearchWhitelistGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ModifyOpenSearchWhitelistGroupRequest) GetIPs() *string {
	return s.IPs
}

func (s *ModifyOpenSearchWhitelistGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyOpenSearchWhitelistGroupRequest) GetRemark() *string {
	return s.Remark
}

func (s *ModifyOpenSearchWhitelistGroupRequest) SetDBInstanceName(v string) *ModifyOpenSearchWhitelistGroupRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyOpenSearchWhitelistGroupRequest) SetGroupId(v string) *ModifyOpenSearchWhitelistGroupRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyOpenSearchWhitelistGroupRequest) SetIPs(v string) *ModifyOpenSearchWhitelistGroupRequest {
	s.IPs = &v
	return s
}

func (s *ModifyOpenSearchWhitelistGroupRequest) SetRegionId(v string) *ModifyOpenSearchWhitelistGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyOpenSearchWhitelistGroupRequest) SetRemark(v string) *ModifyOpenSearchWhitelistGroupRequest {
	s.Remark = &v
	return s
}

func (s *ModifyOpenSearchWhitelistGroupRequest) Validate() error {
	return dara.Validate(s)
}
