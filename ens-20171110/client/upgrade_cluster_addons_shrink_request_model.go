// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeClusterAddonsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddonsShrink(v string) *UpgradeClusterAddonsShrinkRequest
	GetAddonsShrink() *string
	SetClusterId(v string) *UpgradeClusterAddonsShrinkRequest
	GetClusterId() *string
}

type UpgradeClusterAddonsShrinkRequest struct {
	AddonsShrink *string `json:"Addons,omitempty" xml:"Addons,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eck-xxxxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s UpgradeClusterAddonsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeClusterAddonsShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpgradeClusterAddonsShrinkRequest) GetAddonsShrink() *string {
	return s.AddonsShrink
}

func (s *UpgradeClusterAddonsShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpgradeClusterAddonsShrinkRequest) SetAddonsShrink(v string) *UpgradeClusterAddonsShrinkRequest {
	s.AddonsShrink = &v
	return s
}

func (s *UpgradeClusterAddonsShrinkRequest) SetClusterId(v string) *UpgradeClusterAddonsShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *UpgradeClusterAddonsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
