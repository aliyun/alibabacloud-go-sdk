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
	SetEnvironmentId(v string) *DeleteAddonReleaseRequest
	GetEnvironmentId() *string
	SetForce(v bool) *DeleteAddonReleaseRequest
	GetForce() *bool
	SetRegionId(v string) *DeleteAddonReleaseRequest
	GetRegionId() *string
	SetReleaseName(v string) *DeleteAddonReleaseRequest
	GetReleaseName() *string
}

type DeleteAddonReleaseRequest struct {
	// The name of the add-on. If you assign a value to AddonName, the ReleaseName parameter is ignored and all AddonReleases that belong to the same add-on are deleted.
	//
	// example:
	//
	// mysql
	AddonName *string `json:"AddonName,omitempty" xml:"AddonName,omitempty"`
	// Environment ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// env-xxx
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// Whether to be forcibly deleted. The default value is false.
	//
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Name of Release.
	//
	// This parameter is required.
	//
	// example:
	//
	// agent-822567d4-2449
	ReleaseName *string `json:"ReleaseName,omitempty" xml:"ReleaseName,omitempty"`
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

func (s *DeleteAddonReleaseRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *DeleteAddonReleaseRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteAddonReleaseRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteAddonReleaseRequest) GetReleaseName() *string {
	return s.ReleaseName
}

func (s *DeleteAddonReleaseRequest) SetAddonName(v string) *DeleteAddonReleaseRequest {
	s.AddonName = &v
	return s
}

func (s *DeleteAddonReleaseRequest) SetEnvironmentId(v string) *DeleteAddonReleaseRequest {
	s.EnvironmentId = &v
	return s
}

func (s *DeleteAddonReleaseRequest) SetForce(v bool) *DeleteAddonReleaseRequest {
	s.Force = &v
	return s
}

func (s *DeleteAddonReleaseRequest) SetRegionId(v string) *DeleteAddonReleaseRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAddonReleaseRequest) SetReleaseName(v string) *DeleteAddonReleaseRequest {
	s.ReleaseName = &v
	return s
}

func (s *DeleteAddonReleaseRequest) Validate() error {
	return dara.Validate(s)
}
