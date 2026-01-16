// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnInstallClusterAddonsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddonsShrink(v string) *UnInstallClusterAddonsShrinkRequest
	GetAddonsShrink() *string
	SetClusterId(v string) *UnInstallClusterAddonsShrinkRequest
	GetClusterId() *string
}

type UnInstallClusterAddonsShrinkRequest struct {
	AddonsShrink *string `json:"Addons,omitempty" xml:"Addons,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eck-xxxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s UnInstallClusterAddonsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UnInstallClusterAddonsShrinkRequest) GoString() string {
	return s.String()
}

func (s *UnInstallClusterAddonsShrinkRequest) GetAddonsShrink() *string {
	return s.AddonsShrink
}

func (s *UnInstallClusterAddonsShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UnInstallClusterAddonsShrinkRequest) SetAddonsShrink(v string) *UnInstallClusterAddonsShrinkRequest {
	s.AddonsShrink = &v
	return s
}

func (s *UnInstallClusterAddonsShrinkRequest) SetClusterId(v string) *UnInstallClusterAddonsShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *UnInstallClusterAddonsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
