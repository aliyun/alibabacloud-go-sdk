// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnvPodMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunLang(v string) *CreateEnvPodMonitorRequest
	GetAliyunLang() *string
	SetConfigYaml(v string) *CreateEnvPodMonitorRequest
	GetConfigYaml() *string
	SetDryRun(v bool) *CreateEnvPodMonitorRequest
	GetDryRun() *bool
	SetEnvironmentId(v string) *CreateEnvPodMonitorRequest
	GetEnvironmentId() *string
	SetRegionId(v string) *CreateEnvPodMonitorRequest
	GetRegionId() *string
}

type CreateEnvPodMonitorRequest struct {
	// The language. Valid values: zh and en. Default value: zh.
	//
	// example:
	//
	// zh
	AliyunLang *string `json:"AliyunLang,omitempty" xml:"AliyunLang,omitempty"`
	// The YAML configuration snippet for PodMonitor.
	//
	// This parameter is required.
	//
	// example:
	//
	// Refer to supplementary instructions.
	ConfigYaml *string `json:"ConfigYaml,omitempty" xml:"ConfigYaml,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request.
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

func (s CreateEnvPodMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEnvPodMonitorRequest) GoString() string {
	return s.String()
}

func (s *CreateEnvPodMonitorRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *CreateEnvPodMonitorRequest) GetConfigYaml() *string {
	return s.ConfigYaml
}

func (s *CreateEnvPodMonitorRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateEnvPodMonitorRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *CreateEnvPodMonitorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateEnvPodMonitorRequest) SetAliyunLang(v string) *CreateEnvPodMonitorRequest {
	s.AliyunLang = &v
	return s
}

func (s *CreateEnvPodMonitorRequest) SetConfigYaml(v string) *CreateEnvPodMonitorRequest {
	s.ConfigYaml = &v
	return s
}

func (s *CreateEnvPodMonitorRequest) SetDryRun(v bool) *CreateEnvPodMonitorRequest {
	s.DryRun = &v
	return s
}

func (s *CreateEnvPodMonitorRequest) SetEnvironmentId(v string) *CreateEnvPodMonitorRequest {
	s.EnvironmentId = &v
	return s
}

func (s *CreateEnvPodMonitorRequest) SetRegionId(v string) *CreateEnvPodMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *CreateEnvPodMonitorRequest) Validate() error {
	return dara.Validate(s)
}
