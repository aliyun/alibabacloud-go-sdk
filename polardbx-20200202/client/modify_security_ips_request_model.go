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
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzjasd****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The name of the whitelist group of the instance.
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
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of IP addresses in the whitelist group. Separate multiple IP addresses with commas (,).
	//
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
