// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *ModifySecurityIpsRequest
	GetDBInstanceName() *string
	SetGroupName(v string) *ModifySecurityIpsRequest
	GetGroupName() *string
	SetModifyMode(v string) *ModifySecurityIpsRequest
	GetModifyMode() *string
	SetRegionId(v string) *ModifySecurityIpsRequest
	GetRegionId() *string
	SetSecurityIPList(v string) *ModifySecurityIpsRequest
	GetSecurityIPList() *string
}

type ModifySecurityIpsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzjasd****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// default
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// 1
	ModifyMode *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 127.0.0.1,192.168.0.0
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
}

func (s ModifySecurityIpsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityIpsRequest) GoString() string {
	return s.String()
}

func (s *ModifySecurityIpsRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifySecurityIpsRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifySecurityIpsRequest) GetModifyMode() *string {
	return s.ModifyMode
}

func (s *ModifySecurityIpsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySecurityIpsRequest) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *ModifySecurityIpsRequest) SetDBInstanceName(v string) *ModifySecurityIpsRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetGroupName(v string) *ModifySecurityIpsRequest {
	s.GroupName = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetModifyMode(v string) *ModifySecurityIpsRequest {
	s.ModifyMode = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetRegionId(v string) *ModifySecurityIpsRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetSecurityIPList(v string) *ModifySecurityIpsRequest {
	s.SecurityIPList = &v
	return s
}

func (s *ModifySecurityIpsRequest) Validate() error {
	return dara.Validate(s)
}
