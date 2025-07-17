// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallAddonRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddonVersion(v string) *InstallAddonRequest
	GetAddonVersion() *string
	SetAliyunLang(v string) *InstallAddonRequest
	GetAliyunLang() *string
	SetDryRun(v bool) *InstallAddonRequest
	GetDryRun() *bool
	SetEnvironmentId(v string) *InstallAddonRequest
	GetEnvironmentId() *string
	SetName(v string) *InstallAddonRequest
	GetName() *string
	SetRegionId(v string) *InstallAddonRequest
	GetRegionId() *string
	SetReleaseName(v string) *InstallAddonRequest
	GetReleaseName() *string
	SetValues(v string) *InstallAddonRequest
	GetValues() *string
}

type InstallAddonRequest struct {
	// The version of the add-on.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.0.1
	AddonVersion *string `json:"AddonVersion,omitempty" xml:"AddonVersion,omitempty"`
	// The language. Valid values: zh and en. Default value: zh.
	//
	// example:
	//
	// zh
	AliyunLang *string `json:"AliyunLang,omitempty" xml:"AliyunLang,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Default value: false.
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
	// The name of the add-on.
	//
	// This parameter is required.
	//
	// example:
	//
	// mysql
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the add-on after it is installed. If you do not specify this parameter, a default rule name is generated.
	//
	// example:
	//
	// mysql-xxxxx
	ReleaseName *string `json:"ReleaseName,omitempty" xml:"ReleaseName,omitempty"`
	// The metadata.
	//
	// example:
	//
	// {"host":"mysql-service.default","port":3306,"username":"root","password":"roots"}
	Values *string `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s InstallAddonRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallAddonRequest) GoString() string {
	return s.String()
}

func (s *InstallAddonRequest) GetAddonVersion() *string {
	return s.AddonVersion
}

func (s *InstallAddonRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *InstallAddonRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *InstallAddonRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *InstallAddonRequest) GetName() *string {
	return s.Name
}

func (s *InstallAddonRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *InstallAddonRequest) GetReleaseName() *string {
	return s.ReleaseName
}

func (s *InstallAddonRequest) GetValues() *string {
	return s.Values
}

func (s *InstallAddonRequest) SetAddonVersion(v string) *InstallAddonRequest {
	s.AddonVersion = &v
	return s
}

func (s *InstallAddonRequest) SetAliyunLang(v string) *InstallAddonRequest {
	s.AliyunLang = &v
	return s
}

func (s *InstallAddonRequest) SetDryRun(v bool) *InstallAddonRequest {
	s.DryRun = &v
	return s
}

func (s *InstallAddonRequest) SetEnvironmentId(v string) *InstallAddonRequest {
	s.EnvironmentId = &v
	return s
}

func (s *InstallAddonRequest) SetName(v string) *InstallAddonRequest {
	s.Name = &v
	return s
}

func (s *InstallAddonRequest) SetRegionId(v string) *InstallAddonRequest {
	s.RegionId = &v
	return s
}

func (s *InstallAddonRequest) SetReleaseName(v string) *InstallAddonRequest {
	s.ReleaseName = &v
	return s
}

func (s *InstallAddonRequest) SetValues(v string) *InstallAddonRequest {
	s.Values = &v
	return s
}

func (s *InstallAddonRequest) Validate() error {
	return dara.Validate(s)
}
