// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterAddonShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddonShrink(v string) *ModifyClusterAddonShrinkRequest
	GetAddonShrink() *string
	SetClusterId(v string) *ModifyClusterAddonShrinkRequest
	GetClusterId() *string
	SetComponentName(v string) *ModifyClusterAddonShrinkRequest
	GetComponentName() *string
}

type ModifyClusterAddonShrinkRequest struct {
	AddonShrink *string `json:"Addon,omitempty" xml:"Addon,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eck-xxxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// edge-csi-lite
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
}

func (s ModifyClusterAddonShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterAddonShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterAddonShrinkRequest) GetAddonShrink() *string {
	return s.AddonShrink
}

func (s *ModifyClusterAddonShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyClusterAddonShrinkRequest) GetComponentName() *string {
	return s.ComponentName
}

func (s *ModifyClusterAddonShrinkRequest) SetAddonShrink(v string) *ModifyClusterAddonShrinkRequest {
	s.AddonShrink = &v
	return s
}

func (s *ModifyClusterAddonShrinkRequest) SetClusterId(v string) *ModifyClusterAddonShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyClusterAddonShrinkRequest) SetComponentName(v string) *ModifyClusterAddonShrinkRequest {
	s.ComponentName = &v
	return s
}

func (s *ModifyClusterAddonShrinkRequest) Validate() error {
	return dara.Validate(s)
}
