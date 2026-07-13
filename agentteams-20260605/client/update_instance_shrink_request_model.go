// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateInstanceShrinkRequest
	GetClientToken() *string
	SetInstanceId(v string) *UpdateInstanceShrinkRequest
	GetInstanceId() *string
	SetInstanceName(v string) *UpdateInstanceShrinkRequest
	GetInstanceName() *string
	SetNetworkType(v string) *UpdateInstanceShrinkRequest
	GetNetworkType() *string
	SetZonesShrink(v string) *UpdateInstanceShrinkRequest
	GetZonesShrink() *string
}

type UpdateInstanceShrinkRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// agentteams-abc123
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 新的实例名称
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	NetworkType  *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	ZonesShrink  *string `json:"Zones,omitempty" xml:"Zones,omitempty"`
}

func (s UpdateInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateInstanceShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateInstanceShrinkRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *UpdateInstanceShrinkRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *UpdateInstanceShrinkRequest) GetZonesShrink() *string {
	return s.ZonesShrink
}

func (s *UpdateInstanceShrinkRequest) SetClientToken(v string) *UpdateInstanceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateInstanceShrinkRequest) SetInstanceId(v string) *UpdateInstanceShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceShrinkRequest) SetInstanceName(v string) *UpdateInstanceShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *UpdateInstanceShrinkRequest) SetNetworkType(v string) *UpdateInstanceShrinkRequest {
	s.NetworkType = &v
	return s
}

func (s *UpdateInstanceShrinkRequest) SetZonesShrink(v string) *UpdateInstanceShrinkRequest {
	s.ZonesShrink = &v
	return s
}

func (s *UpdateInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
