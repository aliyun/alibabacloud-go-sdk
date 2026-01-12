// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceAuthConfigShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigListShrink(v string) *ModifyInstanceAuthConfigShrinkRequest
	GetConfigListShrink() *string
	SetInstanceName(v string) *ModifyInstanceAuthConfigShrinkRequest
	GetInstanceName() *string
	SetRegionId(v string) *ModifyInstanceAuthConfigShrinkRequest
	GetRegionId() *string
}

type ModifyInstanceAuthConfigShrinkRequest struct {
	// The ID of the RDS Supabase instance.
	ConfigListShrink *string `json:"ConfigList,omitempty" xml:"ConfigList,omitempty"`
	// The region ID.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The operation that you want to perform. Set the value to **ModifyInstanceAuthConfig**.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyInstanceAuthConfigShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAuthConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAuthConfigShrinkRequest) GetConfigListShrink() *string {
	return s.ConfigListShrink
}

func (s *ModifyInstanceAuthConfigShrinkRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyInstanceAuthConfigShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInstanceAuthConfigShrinkRequest) SetConfigListShrink(v string) *ModifyInstanceAuthConfigShrinkRequest {
	s.ConfigListShrink = &v
	return s
}

func (s *ModifyInstanceAuthConfigShrinkRequest) SetInstanceName(v string) *ModifyInstanceAuthConfigShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyInstanceAuthConfigShrinkRequest) SetRegionId(v string) *ModifyInstanceAuthConfigShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceAuthConfigShrinkRequest) Validate() error {
	return dara.Validate(s)
}
