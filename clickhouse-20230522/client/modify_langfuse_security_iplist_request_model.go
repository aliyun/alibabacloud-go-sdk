// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLangfuseSecurityIPListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyLangfuseSecurityIPListRequest
	GetDBInstanceId() *string
	SetGroupName(v string) *ModifyLangfuseSecurityIPListRequest
	GetGroupName() *string
	SetModifyMode(v string) *ModifyLangfuseSecurityIPListRequest
	GetModifyMode() *string
	SetRegionId(v string) *ModifyLangfuseSecurityIPListRequest
	GetRegionId() *string
	SetSecurityIPList(v string) *ModifyLangfuseSecurityIPListRequest
	GetSecurityIPList() *string
}

type ModifyLangfuseSecurityIPListRequest struct {
	// The Langfuse instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lfs-bp1*****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the whitelist group.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The modification mode. Valid values:
	//
	// - 0: overwrite
	//
	// - 1: increase
	//
	// - 2: delete
	//
	// > Specify 0 to use the overwrite mode.
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
	// The IP addresses to add to the instance whitelist. Separate multiple IP addresses with commas (,). For example, 192.168.0.0/24 indicates that all IP addresses in the 192.168.0.XX range are allowed to access the instance.
	//
	// example:
	//
	// 192.168.0.0/24,172.16.0.0/24
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
}

func (s ModifyLangfuseSecurityIPListRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLangfuseSecurityIPListRequest) GoString() string {
	return s.String()
}

func (s *ModifyLangfuseSecurityIPListRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyLangfuseSecurityIPListRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifyLangfuseSecurityIPListRequest) GetModifyMode() *string {
	return s.ModifyMode
}

func (s *ModifyLangfuseSecurityIPListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyLangfuseSecurityIPListRequest) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *ModifyLangfuseSecurityIPListRequest) SetDBInstanceId(v string) *ModifyLangfuseSecurityIPListRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyLangfuseSecurityIPListRequest) SetGroupName(v string) *ModifyLangfuseSecurityIPListRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyLangfuseSecurityIPListRequest) SetModifyMode(v string) *ModifyLangfuseSecurityIPListRequest {
	s.ModifyMode = &v
	return s
}

func (s *ModifyLangfuseSecurityIPListRequest) SetRegionId(v string) *ModifyLangfuseSecurityIPListRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyLangfuseSecurityIPListRequest) SetSecurityIPList(v string) *ModifyLangfuseSecurityIPListRequest {
	s.SecurityIPList = &v
	return s
}

func (s *ModifyLangfuseSecurityIPListRequest) Validate() error {
	return dara.Validate(s)
}
