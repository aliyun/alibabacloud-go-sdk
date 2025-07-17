// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEnvServiceMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunLang(v string) *UpdateEnvServiceMonitorRequest
	GetAliyunLang() *string
	SetConfigYaml(v string) *UpdateEnvServiceMonitorRequest
	GetConfigYaml() *string
	SetDryRun(v bool) *UpdateEnvServiceMonitorRequest
	GetDryRun() *bool
	SetEnvironmentId(v string) *UpdateEnvServiceMonitorRequest
	GetEnvironmentId() *string
	SetNamespace(v string) *UpdateEnvServiceMonitorRequest
	GetNamespace() *string
	SetRegionId(v string) *UpdateEnvServiceMonitorRequest
	GetRegionId() *string
	SetServiceMonitorName(v string) *UpdateEnvServiceMonitorRequest
	GetServiceMonitorName() *string
}

type UpdateEnvServiceMonitorRequest struct {
	// The language. Valid values: zh and en. Default value: zh.
	//
	// example:
	//
	// zh
	AliyunLang *string `json:"AliyunLang,omitempty" xml:"AliyunLang,omitempty"`
	// The YAML configuration string.
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
	// The environment ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// env-xxxxx
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// The namespace where the ServiceMonitor is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// arms-prom
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the ServiceMonitor.
	//
	// This parameter is required.
	//
	// example:
	//
	// serviceMonitor1
	ServiceMonitorName *string `json:"ServiceMonitorName,omitempty" xml:"ServiceMonitorName,omitempty"`
}

func (s UpdateEnvServiceMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnvServiceMonitorRequest) GoString() string {
	return s.String()
}

func (s *UpdateEnvServiceMonitorRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *UpdateEnvServiceMonitorRequest) GetConfigYaml() *string {
	return s.ConfigYaml
}

func (s *UpdateEnvServiceMonitorRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateEnvServiceMonitorRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *UpdateEnvServiceMonitorRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateEnvServiceMonitorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateEnvServiceMonitorRequest) GetServiceMonitorName() *string {
	return s.ServiceMonitorName
}

func (s *UpdateEnvServiceMonitorRequest) SetAliyunLang(v string) *UpdateEnvServiceMonitorRequest {
	s.AliyunLang = &v
	return s
}

func (s *UpdateEnvServiceMonitorRequest) SetConfigYaml(v string) *UpdateEnvServiceMonitorRequest {
	s.ConfigYaml = &v
	return s
}

func (s *UpdateEnvServiceMonitorRequest) SetDryRun(v bool) *UpdateEnvServiceMonitorRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateEnvServiceMonitorRequest) SetEnvironmentId(v string) *UpdateEnvServiceMonitorRequest {
	s.EnvironmentId = &v
	return s
}

func (s *UpdateEnvServiceMonitorRequest) SetNamespace(v string) *UpdateEnvServiceMonitorRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateEnvServiceMonitorRequest) SetRegionId(v string) *UpdateEnvServiceMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateEnvServiceMonitorRequest) SetServiceMonitorName(v string) *UpdateEnvServiceMonitorRequest {
	s.ServiceMonitorName = &v
	return s
}

func (s *UpdateEnvServiceMonitorRequest) Validate() error {
	return dara.Validate(s)
}
