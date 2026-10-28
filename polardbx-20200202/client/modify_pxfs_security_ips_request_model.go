// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPxfsSecurityIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *ModifyPxfsSecurityIpsRequest
	GetDBInstanceName() *string
	SetGroupName(v string) *ModifyPxfsSecurityIpsRequest
	GetGroupName() *string
	SetModifyMode(v string) *ModifyPxfsSecurityIpsRequest
	GetModifyMode() *string
	SetRegionId(v string) *ModifyPxfsSecurityIpsRequest
	GetRegionId() *string
	SetSecurityIPList(v string) *ModifyPxfsSecurityIpsRequest
	GetSecurityIPList() *string
}

type ModifyPxfsSecurityIpsRequest struct {
	// The ID of the PolarDB-X instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-xxxxxxxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The name of the whitelist group.
	//
	// example:
	//
	// default
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The modification mode of the whitelist.
	//
	// example:
	//
	// 0
	ModifyMode *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IP whitelist.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.1.0/24,10.0.0.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
}

func (s ModifyPxfsSecurityIpsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPxfsSecurityIpsRequest) GoString() string {
	return s.String()
}

func (s *ModifyPxfsSecurityIpsRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifyPxfsSecurityIpsRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifyPxfsSecurityIpsRequest) GetModifyMode() *string {
	return s.ModifyMode
}

func (s *ModifyPxfsSecurityIpsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyPxfsSecurityIpsRequest) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *ModifyPxfsSecurityIpsRequest) SetDBInstanceName(v string) *ModifyPxfsSecurityIpsRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyPxfsSecurityIpsRequest) SetGroupName(v string) *ModifyPxfsSecurityIpsRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyPxfsSecurityIpsRequest) SetModifyMode(v string) *ModifyPxfsSecurityIpsRequest {
	s.ModifyMode = &v
	return s
}

func (s *ModifyPxfsSecurityIpsRequest) SetRegionId(v string) *ModifyPxfsSecurityIpsRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyPxfsSecurityIpsRequest) SetSecurityIPList(v string) *ModifyPxfsSecurityIpsRequest {
	s.SecurityIPList = &v
	return s
}

func (s *ModifyPxfsSecurityIpsRequest) Validate() error {
	return dara.Validate(s)
}
