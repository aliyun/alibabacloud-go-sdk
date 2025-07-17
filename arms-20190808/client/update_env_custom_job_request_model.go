// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEnvCustomJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunLang(v string) *UpdateEnvCustomJobRequest
	GetAliyunLang() *string
	SetConfigYaml(v string) *UpdateEnvCustomJobRequest
	GetConfigYaml() *string
	SetCustomJobName(v string) *UpdateEnvCustomJobRequest
	GetCustomJobName() *string
	SetEnvironmentId(v string) *UpdateEnvCustomJobRequest
	GetEnvironmentId() *string
	SetRegionId(v string) *UpdateEnvCustomJobRequest
	GetRegionId() *string
	SetStatus(v string) *UpdateEnvCustomJobRequest
	GetStatus() *string
}

type UpdateEnvCustomJobRequest struct {
	// The language. Valid values: zh and en. Default value: zh.
	//
	// example:
	//
	// zh
	AliyunLang *string `json:"AliyunLang,omitempty" xml:"AliyunLang,omitempty"`
	// The YAML configuration string.
	//
	// example:
	//
	// Refer to supplementary instructions.
	ConfigYaml *string `json:"ConfigYaml,omitempty" xml:"ConfigYaml,omitempty"`
	// The name of the custom job.
	//
	// This parameter is required.
	//
	// example:
	//
	// customJob1
	CustomJobName *string `json:"CustomJobName,omitempty" xml:"CustomJobName,omitempty"`
	// The environment ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// env-xxxxx
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the custom job. Valid values: run and stop.
	//
	// example:
	//
	// run
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateEnvCustomJobRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnvCustomJobRequest) GoString() string {
	return s.String()
}

func (s *UpdateEnvCustomJobRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *UpdateEnvCustomJobRequest) GetConfigYaml() *string {
	return s.ConfigYaml
}

func (s *UpdateEnvCustomJobRequest) GetCustomJobName() *string {
	return s.CustomJobName
}

func (s *UpdateEnvCustomJobRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *UpdateEnvCustomJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateEnvCustomJobRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateEnvCustomJobRequest) SetAliyunLang(v string) *UpdateEnvCustomJobRequest {
	s.AliyunLang = &v
	return s
}

func (s *UpdateEnvCustomJobRequest) SetConfigYaml(v string) *UpdateEnvCustomJobRequest {
	s.ConfigYaml = &v
	return s
}

func (s *UpdateEnvCustomJobRequest) SetCustomJobName(v string) *UpdateEnvCustomJobRequest {
	s.CustomJobName = &v
	return s
}

func (s *UpdateEnvCustomJobRequest) SetEnvironmentId(v string) *UpdateEnvCustomJobRequest {
	s.EnvironmentId = &v
	return s
}

func (s *UpdateEnvCustomJobRequest) SetRegionId(v string) *UpdateEnvCustomJobRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateEnvCustomJobRequest) SetStatus(v string) *UpdateEnvCustomJobRequest {
	s.Status = &v
	return s
}

func (s *UpdateEnvCustomJobRequest) Validate() error {
	return dara.Validate(s)
}
