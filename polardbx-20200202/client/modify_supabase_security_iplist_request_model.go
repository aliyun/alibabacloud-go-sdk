// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySupabaseSecurityIPListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *ModifySupabaseSecurityIPListRequest
	GetDBInstanceName() *string
	SetGroupName(v string) *ModifySupabaseSecurityIPListRequest
	GetGroupName() *string
	SetModifyMode(v string) *ModifySupabaseSecurityIPListRequest
	GetModifyMode() *string
	SetRegionId(v string) *ModifySupabaseSecurityIPListRequest
	GetRegionId() *string
	SetSecurityIPList(v string) *ModifySupabaseSecurityIPListRequest
	GetSecurityIPList() *string
}

type ModifySupabaseSecurityIPListRequest struct {
	// The instance name.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxsp-*********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The whitelist group name.
	//
	// example:
	//
	// default
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The modification mode. Valid values:
	//
	// - Cover: overwrites the existing whitelist.
	//
	// - Append: appends entries to the existing whitelist.
	//
	// - Delete: deletes entries from the existing whitelist.
	//
	// example:
	//
	// Append
	ModifyMode *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IP whitelist. Separate multiple IP addresses with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 127.0.0.1,192.168.0.0
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
}

func (s ModifySupabaseSecurityIPListRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySupabaseSecurityIPListRequest) GoString() string {
	return s.String()
}

func (s *ModifySupabaseSecurityIPListRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifySupabaseSecurityIPListRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifySupabaseSecurityIPListRequest) GetModifyMode() *string {
	return s.ModifyMode
}

func (s *ModifySupabaseSecurityIPListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySupabaseSecurityIPListRequest) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *ModifySupabaseSecurityIPListRequest) SetDBInstanceName(v string) *ModifySupabaseSecurityIPListRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifySupabaseSecurityIPListRequest) SetGroupName(v string) *ModifySupabaseSecurityIPListRequest {
	s.GroupName = &v
	return s
}

func (s *ModifySupabaseSecurityIPListRequest) SetModifyMode(v string) *ModifySupabaseSecurityIPListRequest {
	s.ModifyMode = &v
	return s
}

func (s *ModifySupabaseSecurityIPListRequest) SetRegionId(v string) *ModifySupabaseSecurityIPListRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySupabaseSecurityIPListRequest) SetSecurityIPList(v string) *ModifySupabaseSecurityIPListRequest {
	s.SecurityIPList = &v
	return s
}

func (s *ModifySupabaseSecurityIPListRequest) Validate() error {
	return dara.Validate(s)
}
