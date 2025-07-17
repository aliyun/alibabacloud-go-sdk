// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnvServiceMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunLang(v string) *CreateEnvServiceMonitorRequest
	GetAliyunLang() *string
	SetConfigYaml(v string) *CreateEnvServiceMonitorRequest
	GetConfigYaml() *string
	SetDryRun(v bool) *CreateEnvServiceMonitorRequest
	GetDryRun() *bool
	SetEnvironmentId(v string) *CreateEnvServiceMonitorRequest
	GetEnvironmentId() *string
	SetRegionId(v string) *CreateEnvServiceMonitorRequest
	GetRegionId() *string
}

type CreateEnvServiceMonitorRequest struct {
	// The language. Valid values:
	//
	// 	- zh (default): Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AliyunLang *string `json:"AliyunLang,omitempty" xml:"AliyunLang,omitempty"`
	// The YAML configuration file of the ServiceMonitor.
	//
	// This parameter is required.
	//
	// example:
	//
	// Refer to supplementary instructions.
	ConfigYaml *string `json:"ConfigYaml,omitempty" xml:"ConfigYaml,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. The system checks whether the format is valid and whether targets are matched.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the environment instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// env-xxxxxx
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateEnvServiceMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEnvServiceMonitorRequest) GoString() string {
	return s.String()
}

func (s *CreateEnvServiceMonitorRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *CreateEnvServiceMonitorRequest) GetConfigYaml() *string {
	return s.ConfigYaml
}

func (s *CreateEnvServiceMonitorRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateEnvServiceMonitorRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *CreateEnvServiceMonitorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateEnvServiceMonitorRequest) SetAliyunLang(v string) *CreateEnvServiceMonitorRequest {
	s.AliyunLang = &v
	return s
}

func (s *CreateEnvServiceMonitorRequest) SetConfigYaml(v string) *CreateEnvServiceMonitorRequest {
	s.ConfigYaml = &v
	return s
}

func (s *CreateEnvServiceMonitorRequest) SetDryRun(v bool) *CreateEnvServiceMonitorRequest {
	s.DryRun = &v
	return s
}

func (s *CreateEnvServiceMonitorRequest) SetEnvironmentId(v string) *CreateEnvServiceMonitorRequest {
	s.EnvironmentId = &v
	return s
}

func (s *CreateEnvServiceMonitorRequest) SetRegionId(v string) *CreateEnvServiceMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *CreateEnvServiceMonitorRequest) Validate() error {
	return dara.Validate(s)
}
