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
	SetResourceOwnerId(v int64) *ModifySecurityIPListRequest
	GetResourceOwnerId() *int64
	SetSecurityIPList(v string) *ModifySecurityIPListRequest
	GetSecurityIPList() *string
}

type ModifySecurityIPListRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the whitelist. Default value: **Default**.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The mode in which you want to modify the whitelist. Valid values:
	//
	// 	- **0**: overwrites the IP addresses in the whitelist.
	//
	// 	- **1**: adds IP addresses to the whitelist.
	//
	// 	- **2**: removes IP addresses from the whitelist.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	ModifyMode *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The IP addresses in the whitelist of the instance. Separate multiple IP addresses with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 127.0.XX.XX,127.2.XX.XX
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

func (s *ModifySecurityIPListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
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

func (s *ModifySecurityIPListRequest) SetResourceOwnerId(v int64) *ModifySecurityIPListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySecurityIPListRequest) SetSecurityIPList(v string) *ModifySecurityIPListRequest {
	s.SecurityIPList = &v
	return s
}

func (s *ModifySecurityIPListRequest) Validate() error {
	return dara.Validate(s)
}
