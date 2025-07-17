// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEnvPodMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunLang(v string) *UpdateEnvPodMonitorRequest
	GetAliyunLang() *string
	SetConfigYaml(v string) *UpdateEnvPodMonitorRequest
	GetConfigYaml() *string
	SetDryRun(v bool) *UpdateEnvPodMonitorRequest
	GetDryRun() *bool
	SetEnvironmentId(v string) *UpdateEnvPodMonitorRequest
	GetEnvironmentId() *string
	SetNamespace(v string) *UpdateEnvPodMonitorRequest
	GetNamespace() *string
	SetPodMonitorName(v string) *UpdateEnvPodMonitorRequest
	GetPodMonitorName() *string
	SetRegionId(v string) *UpdateEnvPodMonitorRequest
	GetRegionId() *string
}

type UpdateEnvPodMonitorRequest struct {
	// The language. Valid values:
	//
	// 	- zh (default value): Chinese
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
	// Checks whether the format is valid and whether targets are matched.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The environment ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// env-xxxxx
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// The namespace where the PodMonitor resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// arms-prom
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The name of the PodMonitor.
	//
	// This parameter is required.
	//
	// example:
	//
	// arms-admin-pm1
	PodMonitorName *string `json:"PodMonitorName,omitempty" xml:"PodMonitorName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateEnvPodMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnvPodMonitorRequest) GoString() string {
	return s.String()
}

func (s *UpdateEnvPodMonitorRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *UpdateEnvPodMonitorRequest) GetConfigYaml() *string {
	return s.ConfigYaml
}

func (s *UpdateEnvPodMonitorRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateEnvPodMonitorRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *UpdateEnvPodMonitorRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateEnvPodMonitorRequest) GetPodMonitorName() *string {
	return s.PodMonitorName
}

func (s *UpdateEnvPodMonitorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateEnvPodMonitorRequest) SetAliyunLang(v string) *UpdateEnvPodMonitorRequest {
	s.AliyunLang = &v
	return s
}

func (s *UpdateEnvPodMonitorRequest) SetConfigYaml(v string) *UpdateEnvPodMonitorRequest {
	s.ConfigYaml = &v
	return s
}

func (s *UpdateEnvPodMonitorRequest) SetDryRun(v bool) *UpdateEnvPodMonitorRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateEnvPodMonitorRequest) SetEnvironmentId(v string) *UpdateEnvPodMonitorRequest {
	s.EnvironmentId = &v
	return s
}

func (s *UpdateEnvPodMonitorRequest) SetNamespace(v string) *UpdateEnvPodMonitorRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateEnvPodMonitorRequest) SetPodMonitorName(v string) *UpdateEnvPodMonitorRequest {
	s.PodMonitorName = &v
	return s
}

func (s *UpdateEnvPodMonitorRequest) SetRegionId(v string) *UpdateEnvPodMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateEnvPodMonitorRequest) Validate() error {
	return dara.Validate(s)
}
