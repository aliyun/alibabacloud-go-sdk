// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnvCustomJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunLang(v string) *CreateEnvCustomJobRequest
	GetAliyunLang() *string
	SetConfigYaml(v string) *CreateEnvCustomJobRequest
	GetConfigYaml() *string
	SetCustomJobName(v string) *CreateEnvCustomJobRequest
	GetCustomJobName() *string
	SetEnvironmentId(v string) *CreateEnvCustomJobRequest
	GetEnvironmentId() *string
	SetRegionId(v string) *CreateEnvCustomJobRequest
	GetRegionId() *string
}

type CreateEnvCustomJobRequest struct {
	// The language. Valid values: zh and en. Default value: zh.
	//
	// example:
	//
	// zh
	AliyunLang *string `json:"AliyunLang,omitempty" xml:"AliyunLang,omitempty"`
	// The YAML configuration string of the custom job.
	//
	// This parameter is required.
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

func (s CreateEnvCustomJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEnvCustomJobRequest) GoString() string {
	return s.String()
}

func (s *CreateEnvCustomJobRequest) GetAliyunLang() *string {
	return s.AliyunLang
}

func (s *CreateEnvCustomJobRequest) GetConfigYaml() *string {
	return s.ConfigYaml
}

func (s *CreateEnvCustomJobRequest) GetCustomJobName() *string {
	return s.CustomJobName
}

func (s *CreateEnvCustomJobRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *CreateEnvCustomJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateEnvCustomJobRequest) SetAliyunLang(v string) *CreateEnvCustomJobRequest {
	s.AliyunLang = &v
	return s
}

func (s *CreateEnvCustomJobRequest) SetConfigYaml(v string) *CreateEnvCustomJobRequest {
	s.ConfigYaml = &v
	return s
}

func (s *CreateEnvCustomJobRequest) SetCustomJobName(v string) *CreateEnvCustomJobRequest {
	s.CustomJobName = &v
	return s
}

func (s *CreateEnvCustomJobRequest) SetEnvironmentId(v string) *CreateEnvCustomJobRequest {
	s.EnvironmentId = &v
	return s
}

func (s *CreateEnvCustomJobRequest) SetRegionId(v string) *CreateEnvCustomJobRequest {
	s.RegionId = &v
	return s
}

func (s *CreateEnvCustomJobRequest) Validate() error {
	return dara.Validate(s)
}
