// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityIPGroupRelationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGlobalSecurityGroupId(v string) *ModifySecurityIPGroupRelationRequest
	GetGlobalSecurityGroupId() *string
	SetInstanceId(v string) *ModifySecurityIPGroupRelationRequest
	GetInstanceId() *string
	SetRegionName(v string) *ModifySecurityIPGroupRelationRequest
	GetRegionName() *string
}

type ModifySecurityIPGroupRelationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// g-ggh7qvrrwikyxe****
	GlobalSecurityGroupId *string `json:"GlobalSecurityGroupId,omitempty" xml:"GlobalSecurityGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pc-bp1u5mas9exx7****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
}

func (s ModifySecurityIPGroupRelationRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityIPGroupRelationRequest) GoString() string {
	return s.String()
}

func (s *ModifySecurityIPGroupRelationRequest) GetGlobalSecurityGroupId() *string {
	return s.GlobalSecurityGroupId
}

func (s *ModifySecurityIPGroupRelationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifySecurityIPGroupRelationRequest) GetRegionName() *string {
	return s.RegionName
}

func (s *ModifySecurityIPGroupRelationRequest) SetGlobalSecurityGroupId(v string) *ModifySecurityIPGroupRelationRequest {
	s.GlobalSecurityGroupId = &v
	return s
}

func (s *ModifySecurityIPGroupRelationRequest) SetInstanceId(v string) *ModifySecurityIPGroupRelationRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifySecurityIPGroupRelationRequest) SetRegionName(v string) *ModifySecurityIPGroupRelationRequest {
	s.RegionName = &v
	return s
}

func (s *ModifySecurityIPGroupRelationRequest) Validate() error {
	return dara.Validate(s)
}
