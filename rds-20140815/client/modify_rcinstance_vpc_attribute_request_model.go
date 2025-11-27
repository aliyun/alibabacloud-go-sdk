// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCInstanceVpcAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyRCInstanceVpcAttributeRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyRCInstanceVpcAttributeRequest
	GetRegionId() *string
	SetVSwitchId(v string) *ModifyRCInstanceVpcAttributeRequest
	GetVSwitchId() *string
}

type ModifyRCInstanceVpcAttributeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// rc-u65p1qjg249t3a30738
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1s5fnvk4gn3tw12****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s ModifyRCInstanceVpcAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCInstanceVpcAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyRCInstanceVpcAttributeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyRCInstanceVpcAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyRCInstanceVpcAttributeRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ModifyRCInstanceVpcAttributeRequest) SetInstanceId(v string) *ModifyRCInstanceVpcAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyRCInstanceVpcAttributeRequest) SetRegionId(v string) *ModifyRCInstanceVpcAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyRCInstanceVpcAttributeRequest) SetVSwitchId(v string) *ModifyRCInstanceVpcAttributeRequest {
	s.VSwitchId = &v
	return s
}

func (s *ModifyRCInstanceVpcAttributeRequest) Validate() error {
	return dara.Validate(s)
}
