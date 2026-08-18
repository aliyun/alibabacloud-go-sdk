// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPxfuseSecurityIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *ModifyPxfuseSecurityIpsRequest
	GetDBInstanceName() *string
	SetGroupName(v string) *ModifyPxfuseSecurityIpsRequest
	GetGroupName() *string
	SetModifyMode(v string) *ModifyPxfuseSecurityIpsRequest
	GetModifyMode() *string
	SetRegionId(v string) *ModifyPxfuseSecurityIpsRequest
	GetRegionId() *string
	SetSecurityIPList(v string) *ModifyPxfuseSecurityIpsRequest
	GetSecurityIPList() *string
}

type ModifyPxfuseSecurityIpsRequest struct {
	// The name of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxsp-*********
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

func (s ModifyPxfuseSecurityIpsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPxfuseSecurityIpsRequest) GoString() string {
	return s.String()
}

func (s *ModifyPxfuseSecurityIpsRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifyPxfuseSecurityIpsRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifyPxfuseSecurityIpsRequest) GetModifyMode() *string {
	return s.ModifyMode
}

func (s *ModifyPxfuseSecurityIpsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyPxfuseSecurityIpsRequest) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *ModifyPxfuseSecurityIpsRequest) SetDBInstanceName(v string) *ModifyPxfuseSecurityIpsRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyPxfuseSecurityIpsRequest) SetGroupName(v string) *ModifyPxfuseSecurityIpsRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyPxfuseSecurityIpsRequest) SetModifyMode(v string) *ModifyPxfuseSecurityIpsRequest {
	s.ModifyMode = &v
	return s
}

func (s *ModifyPxfuseSecurityIpsRequest) SetRegionId(v string) *ModifyPxfuseSecurityIpsRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyPxfuseSecurityIpsRequest) SetSecurityIPList(v string) *ModifyPxfuseSecurityIpsRequest {
	s.SecurityIPList = &v
	return s
}

func (s *ModifyPxfuseSecurityIpsRequest) Validate() error {
	return dara.Validate(s)
}
