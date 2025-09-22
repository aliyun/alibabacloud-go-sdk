// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceStorageConfigShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyInstanceStorageConfigShrinkRequest
	GetClientToken() *string
	SetConfigListShrink(v string) *ModifyInstanceStorageConfigShrinkRequest
	GetConfigListShrink() *string
	SetInstanceName(v string) *ModifyInstanceStorageConfigShrinkRequest
	GetInstanceName() *string
	SetRegionId(v string) *ModifyInstanceStorageConfigShrinkRequest
	GetRegionId() *string
}

type ModifyInstanceStorageConfigShrinkRequest struct {
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConfigListShrink *string `json:"ConfigList,omitempty" xml:"ConfigList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyInstanceStorageConfigShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceStorageConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceStorageConfigShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyInstanceStorageConfigShrinkRequest) GetConfigListShrink() *string {
	return s.ConfigListShrink
}

func (s *ModifyInstanceStorageConfigShrinkRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyInstanceStorageConfigShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInstanceStorageConfigShrinkRequest) SetClientToken(v string) *ModifyInstanceStorageConfigShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyInstanceStorageConfigShrinkRequest) SetConfigListShrink(v string) *ModifyInstanceStorageConfigShrinkRequest {
	s.ConfigListShrink = &v
	return s
}

func (s *ModifyInstanceStorageConfigShrinkRequest) SetInstanceName(v string) *ModifyInstanceStorageConfigShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyInstanceStorageConfigShrinkRequest) SetRegionId(v string) *ModifyInstanceStorageConfigShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceStorageConfigShrinkRequest) Validate() error {
	return dara.Validate(s)
}
