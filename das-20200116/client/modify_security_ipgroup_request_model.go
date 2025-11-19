// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityIPGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGIpList(v string) *ModifySecurityIPGroupRequest
	GetGIpList() *string
	SetGlobalIgName(v string) *ModifySecurityIPGroupRequest
	GetGlobalIgName() *string
	SetGlobalSecurityGroupId(v string) *ModifySecurityIPGroupRequest
	GetGlobalSecurityGroupId() *string
	SetRegionName(v string) *ModifySecurityIPGroupRequest
	GetRegionName() *string
}

type ModifySecurityIPGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 192.168.0.1
	GIpList *string `json:"GIpList,omitempty" xml:"GIpList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_123
	GlobalIgName *string `json:"GlobalIgName,omitempty" xml:"GlobalIgName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// g-9ubyrqeoko****,g-pdxlrvn13k****
	GlobalSecurityGroupId *string `json:"GlobalSecurityGroupId,omitempty" xml:"GlobalSecurityGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
}

func (s ModifySecurityIPGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityIPGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifySecurityIPGroupRequest) GetGIpList() *string {
	return s.GIpList
}

func (s *ModifySecurityIPGroupRequest) GetGlobalIgName() *string {
	return s.GlobalIgName
}

func (s *ModifySecurityIPGroupRequest) GetGlobalSecurityGroupId() *string {
	return s.GlobalSecurityGroupId
}

func (s *ModifySecurityIPGroupRequest) GetRegionName() *string {
	return s.RegionName
}

func (s *ModifySecurityIPGroupRequest) SetGIpList(v string) *ModifySecurityIPGroupRequest {
	s.GIpList = &v
	return s
}

func (s *ModifySecurityIPGroupRequest) SetGlobalIgName(v string) *ModifySecurityIPGroupRequest {
	s.GlobalIgName = &v
	return s
}

func (s *ModifySecurityIPGroupRequest) SetGlobalSecurityGroupId(v string) *ModifySecurityIPGroupRequest {
	s.GlobalSecurityGroupId = &v
	return s
}

func (s *ModifySecurityIPGroupRequest) SetRegionName(v string) *ModifySecurityIPGroupRequest {
	s.RegionName = &v
	return s
}

func (s *ModifySecurityIPGroupRequest) Validate() error {
	return dara.Validate(s)
}
