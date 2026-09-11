// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAddonReleaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddonName(v string) *DeleteAddonReleaseRequest
	GetAddonName() *string
	SetForce(v bool) *DeleteAddonReleaseRequest
	GetForce() *bool
	SetReleaseName(v string) *DeleteAddonReleaseRequest
	GetReleaseName() *string
}

type DeleteAddonReleaseRequest struct {
	// The addon name. If AddonName is specified, the ReleaseName parameter is ignored and all AddonRelease instances that belong to the same addon are uninstalled in a batch.
	//
	// example:
	//
	// cs-gpu
	AddonName *string `json:"addonName,omitempty" xml:"addonName,omitempty"`
	// Deprecated
	//
	// Specifies whether to force delete the addon release. Default value: false.
	//
	// example:
	//
	// false
	Force *bool `json:"force,omitempty" xml:"force,omitempty"`
	// The name of the AddonRelease.
	//
	// example:
	//
	// test-gpu-integration-name
	ReleaseName *string `json:"releaseName,omitempty" xml:"releaseName,omitempty"`
}

func (s DeleteAddonReleaseRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAddonReleaseRequest) GoString() string {
	return s.String()
}

func (s *DeleteAddonReleaseRequest) GetAddonName() *string {
	return s.AddonName
}

func (s *DeleteAddonReleaseRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteAddonReleaseRequest) GetReleaseName() *string {
	return s.ReleaseName
}

func (s *DeleteAddonReleaseRequest) SetAddonName(v string) *DeleteAddonReleaseRequest {
	s.AddonName = &v
	return s
}

func (s *DeleteAddonReleaseRequest) SetForce(v bool) *DeleteAddonReleaseRequest {
	s.Force = &v
	return s
}

func (s *DeleteAddonReleaseRequest) SetReleaseName(v string) *DeleteAddonReleaseRequest {
	s.ReleaseName = &v
	return s
}

func (s *DeleteAddonReleaseRequest) Validate() error {
	return dara.Validate(s)
}
