// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecurityIPGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGIpList(v string) *CreateSecurityIPGroupRequest
	GetGIpList() *string
	SetGlobalIgName(v string) *CreateSecurityIPGroupRequest
	GetGlobalIgName() *string
	SetRegionName(v string) *CreateSecurityIPGroupRequest
	GetRegionName() *string
}

type CreateSecurityIPGroupRequest struct {
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
	// cn-beijing
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
}

func (s CreateSecurityIPGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityIPGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateSecurityIPGroupRequest) GetGIpList() *string {
	return s.GIpList
}

func (s *CreateSecurityIPGroupRequest) GetGlobalIgName() *string {
	return s.GlobalIgName
}

func (s *CreateSecurityIPGroupRequest) GetRegionName() *string {
	return s.RegionName
}

func (s *CreateSecurityIPGroupRequest) SetGIpList(v string) *CreateSecurityIPGroupRequest {
	s.GIpList = &v
	return s
}

func (s *CreateSecurityIPGroupRequest) SetGlobalIgName(v string) *CreateSecurityIPGroupRequest {
	s.GlobalIgName = &v
	return s
}

func (s *CreateSecurityIPGroupRequest) SetRegionName(v string) *CreateSecurityIPGroupRequest {
	s.RegionName = &v
	return s
}

func (s *CreateSecurityIPGroupRequest) Validate() error {
	return dara.Validate(s)
}
