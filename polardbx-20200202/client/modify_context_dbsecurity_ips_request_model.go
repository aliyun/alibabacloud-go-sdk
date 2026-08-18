// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyContextDBSecurityIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *ModifyContextDBSecurityIpsRequest
	GetDBInstanceName() *string
	SetGroupName(v string) *ModifyContextDBSecurityIpsRequest
	GetGroupName() *string
	SetModifyMode(v string) *ModifyContextDBSecurityIpsRequest
	GetModifyMode() *string
	SetRegionId(v string) *ModifyContextDBSecurityIpsRequest
	GetRegionId() *string
	SetSecurityIPList(v string) *ModifyContextDBSecurityIpsRequest
	GetSecurityIPList() *string
}

type ModifyContextDBSecurityIpsRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-htri0****r4k9p
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
	// The region in which the instance resides.
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

func (s ModifyContextDBSecurityIpsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyContextDBSecurityIpsRequest) GoString() string {
	return s.String()
}

func (s *ModifyContextDBSecurityIpsRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifyContextDBSecurityIpsRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifyContextDBSecurityIpsRequest) GetModifyMode() *string {
	return s.ModifyMode
}

func (s *ModifyContextDBSecurityIpsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyContextDBSecurityIpsRequest) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *ModifyContextDBSecurityIpsRequest) SetDBInstanceName(v string) *ModifyContextDBSecurityIpsRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyContextDBSecurityIpsRequest) SetGroupName(v string) *ModifyContextDBSecurityIpsRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyContextDBSecurityIpsRequest) SetModifyMode(v string) *ModifyContextDBSecurityIpsRequest {
	s.ModifyMode = &v
	return s
}

func (s *ModifyContextDBSecurityIpsRequest) SetRegionId(v string) *ModifyContextDBSecurityIpsRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyContextDBSecurityIpsRequest) SetSecurityIPList(v string) *ModifyContextDBSecurityIpsRequest {
	s.SecurityIPList = &v
	return s
}

func (s *ModifyContextDBSecurityIpsRequest) Validate() error {
	return dara.Validate(s)
}
