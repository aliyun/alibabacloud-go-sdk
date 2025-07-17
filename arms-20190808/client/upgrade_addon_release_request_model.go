// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeAddonReleaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddonVersion(v string) *UpgradeAddonReleaseRequest
	GetAddonVersion() *string
	SetDryRun(v bool) *UpgradeAddonReleaseRequest
	GetDryRun() *bool
	SetEnvironmentId(v string) *UpgradeAddonReleaseRequest
	GetEnvironmentId() *string
	SetRegionId(v string) *UpgradeAddonReleaseRequest
	GetRegionId() *string
	SetReleaseName(v string) *UpgradeAddonReleaseRequest
	GetReleaseName() *string
	SetValues(v string) *UpgradeAddonReleaseRequest
	GetValues() *string
}

type UpgradeAddonReleaseRequest struct {
	// The version of the add-on.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.0.2
	AddonVersion *string `json:"AddonVersion,omitempty" xml:"AddonVersion,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The environment ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// env-xxx
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the release.
	//
	// This parameter is required.
	//
	// example:
	//
	// mysql-1695372983039
	ReleaseName *string `json:"ReleaseName,omitempty" xml:"ReleaseName,omitempty"`
	// The metadata information.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"host":"mysql-service.default","port":3306,"username":"root","password":"roots"}
	Values *string `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s UpgradeAddonReleaseRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeAddonReleaseRequest) GoString() string {
	return s.String()
}

func (s *UpgradeAddonReleaseRequest) GetAddonVersion() *string {
	return s.AddonVersion
}

func (s *UpgradeAddonReleaseRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpgradeAddonReleaseRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *UpgradeAddonReleaseRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpgradeAddonReleaseRequest) GetReleaseName() *string {
	return s.ReleaseName
}

func (s *UpgradeAddonReleaseRequest) GetValues() *string {
	return s.Values
}

func (s *UpgradeAddonReleaseRequest) SetAddonVersion(v string) *UpgradeAddonReleaseRequest {
	s.AddonVersion = &v
	return s
}

func (s *UpgradeAddonReleaseRequest) SetDryRun(v bool) *UpgradeAddonReleaseRequest {
	s.DryRun = &v
	return s
}

func (s *UpgradeAddonReleaseRequest) SetEnvironmentId(v string) *UpgradeAddonReleaseRequest {
	s.EnvironmentId = &v
	return s
}

func (s *UpgradeAddonReleaseRequest) SetRegionId(v string) *UpgradeAddonReleaseRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeAddonReleaseRequest) SetReleaseName(v string) *UpgradeAddonReleaseRequest {
	s.ReleaseName = &v
	return s
}

func (s *UpgradeAddonReleaseRequest) SetValues(v string) *UpgradeAddonReleaseRequest {
	s.Values = &v
	return s
}

func (s *UpgradeAddonReleaseRequest) Validate() error {
	return dara.Validate(s)
}
