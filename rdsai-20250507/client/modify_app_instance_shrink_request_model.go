// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyAppInstanceShrinkRequest
	GetClientToken() *string
	SetComponentsShrink(v string) *ModifyAppInstanceShrinkRequest
	GetComponentsShrink() *string
	SetInstanceName(v string) *ModifyAppInstanceShrinkRequest
	GetInstanceName() *string
	SetRegionId(v string) *ModifyAppInstanceShrinkRequest
	GetRegionId() *string
}

type ModifyAppInstanceShrinkRequest struct {
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ComponentsShrink *string `json:"Components,omitempty" xml:"Components,omitempty"`
	// example:
	//
	// ra-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyAppInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyAppInstanceShrinkRequest) GetComponentsShrink() *string {
	return s.ComponentsShrink
}

func (s *ModifyAppInstanceShrinkRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyAppInstanceShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyAppInstanceShrinkRequest) SetClientToken(v string) *ModifyAppInstanceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyAppInstanceShrinkRequest) SetComponentsShrink(v string) *ModifyAppInstanceShrinkRequest {
	s.ComponentsShrink = &v
	return s
}

func (s *ModifyAppInstanceShrinkRequest) SetInstanceName(v string) *ModifyAppInstanceShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyAppInstanceShrinkRequest) SetRegionId(v string) *ModifyAppInstanceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAppInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
