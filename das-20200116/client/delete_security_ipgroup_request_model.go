// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityIPGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGlobalSecurityGroupId(v string) *DeleteSecurityIPGroupRequest
	GetGlobalSecurityGroupId() *string
	SetRegionName(v string) *DeleteSecurityIPGroupRequest
	GetRegionName() *string
}

type DeleteSecurityIPGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// g-rcge12hbfooad3m****
	GlobalSecurityGroupId *string `json:"GlobalSecurityGroupId,omitempty" xml:"GlobalSecurityGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
}

func (s DeleteSecurityIPGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityIPGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecurityIPGroupRequest) GetGlobalSecurityGroupId() *string {
	return s.GlobalSecurityGroupId
}

func (s *DeleteSecurityIPGroupRequest) GetRegionName() *string {
	return s.RegionName
}

func (s *DeleteSecurityIPGroupRequest) SetGlobalSecurityGroupId(v string) *DeleteSecurityIPGroupRequest {
	s.GlobalSecurityGroupId = &v
	return s
}

func (s *DeleteSecurityIPGroupRequest) SetRegionName(v string) *DeleteSecurityIPGroupRequest {
	s.RegionName = &v
	return s
}

func (s *DeleteSecurityIPGroupRequest) Validate() error {
	return dara.Validate(s)
}
