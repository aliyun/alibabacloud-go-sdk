// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallClusterAddonsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddonsShrink(v string) *InstallClusterAddonsShrinkRequest
	GetAddonsShrink() *string
	SetClusterId(v string) *InstallClusterAddonsShrinkRequest
	GetClusterId() *string
}

type InstallClusterAddonsShrinkRequest struct {
	// This parameter is required.
	AddonsShrink *string `json:"Addons,omitempty" xml:"Addons,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eck-xxxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s InstallClusterAddonsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallClusterAddonsShrinkRequest) GoString() string {
	return s.String()
}

func (s *InstallClusterAddonsShrinkRequest) GetAddonsShrink() *string {
	return s.AddonsShrink
}

func (s *InstallClusterAddonsShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *InstallClusterAddonsShrinkRequest) SetAddonsShrink(v string) *InstallClusterAddonsShrinkRequest {
	s.AddonsShrink = &v
	return s
}

func (s *InstallClusterAddonsShrinkRequest) SetClusterId(v string) *InstallClusterAddonsShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *InstallClusterAddonsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
