// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceRAGConfigShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyInstanceRAGConfigShrinkRequest
	GetClientToken() *string
	SetConfigListShrink(v string) *ModifyInstanceRAGConfigShrinkRequest
	GetConfigListShrink() *string
	SetInstanceName(v string) *ModifyInstanceRAGConfigShrinkRequest
	GetInstanceName() *string
	SetRegionId(v string) *ModifyInstanceRAGConfigShrinkRequest
	GetRegionId() *string
	SetStatus(v bool) *ModifyInstanceRAGConfigShrinkRequest
	GetStatus() *bool
}

type ModifyInstanceRAGConfigShrinkRequest struct {
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
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyInstanceRAGConfigShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceRAGConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceRAGConfigShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyInstanceRAGConfigShrinkRequest) GetConfigListShrink() *string {
	return s.ConfigListShrink
}

func (s *ModifyInstanceRAGConfigShrinkRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyInstanceRAGConfigShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInstanceRAGConfigShrinkRequest) GetStatus() *bool {
	return s.Status
}

func (s *ModifyInstanceRAGConfigShrinkRequest) SetClientToken(v string) *ModifyInstanceRAGConfigShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyInstanceRAGConfigShrinkRequest) SetConfigListShrink(v string) *ModifyInstanceRAGConfigShrinkRequest {
	s.ConfigListShrink = &v
	return s
}

func (s *ModifyInstanceRAGConfigShrinkRequest) SetInstanceName(v string) *ModifyInstanceRAGConfigShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyInstanceRAGConfigShrinkRequest) SetRegionId(v string) *ModifyInstanceRAGConfigShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceRAGConfigShrinkRequest) SetStatus(v bool) *ModifyInstanceRAGConfigShrinkRequest {
	s.Status = &v
	return s
}

func (s *ModifyInstanceRAGConfigShrinkRequest) Validate() error {
	return dara.Validate(s)
}
