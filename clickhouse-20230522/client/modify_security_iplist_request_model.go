// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityIPListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifySecurityIPListRequest
	GetDBInstanceId() *string
	SetGroupName(v string) *ModifySecurityIPListRequest
	GetGroupName() *string
	SetModifyMode(v string) *ModifySecurityIPListRequest
	GetModifyMode() *string
	SetRegionId(v string) *ModifySecurityIPListRequest
	GetRegionId() *string
	SetSecurityIPList(v string) *ModifySecurityIPListRequest
	GetSecurityIPList() *string
}

type ModifySecurityIPListRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-xxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the whitelist whose settings you want to modify.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The modification mode.
	//
	// 	- 0: overwrites the original IP addresses and CIDR blocks in the whitelist.
	//
	// 	- 1: adds the IP addresses and CIDR blocks to the whitelist.
	//
	// 	- 2: removes the IP addresses and CIDR blocks from the whitelist.
	//
	// >  We recommend that you set the value to 0.
	//
	// example:
	//
	// 0
	ModifyMode *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IP addresses and CIDR blocks in the whitelist.
	//
	// example:
	//
	// 192.168.0.0/24,172.16.0.0/24
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
}

func (s ModifySecurityIPListRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityIPListRequest) GoString() string {
	return s.String()
}

func (s *ModifySecurityIPListRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifySecurityIPListRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifySecurityIPListRequest) GetModifyMode() *string {
	return s.ModifyMode
}

func (s *ModifySecurityIPListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySecurityIPListRequest) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *ModifySecurityIPListRequest) SetDBInstanceId(v string) *ModifySecurityIPListRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifySecurityIPListRequest) SetGroupName(v string) *ModifySecurityIPListRequest {
	s.GroupName = &v
	return s
}

func (s *ModifySecurityIPListRequest) SetModifyMode(v string) *ModifySecurityIPListRequest {
	s.ModifyMode = &v
	return s
}

func (s *ModifySecurityIPListRequest) SetRegionId(v string) *ModifySecurityIPListRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySecurityIPListRequest) SetSecurityIPList(v string) *ModifySecurityIPListRequest {
	s.SecurityIPList = &v
	return s
}

func (s *ModifySecurityIPListRequest) Validate() error {
	return dara.Validate(s)
}
