// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnvCustomJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomJobName(v string) *DescribeEnvCustomJobRequest
	GetCustomJobName() *string
	SetEncryptYaml(v bool) *DescribeEnvCustomJobRequest
	GetEncryptYaml() *bool
	SetEnvironmentId(v string) *DescribeEnvCustomJobRequest
	GetEnvironmentId() *string
	SetRegionId(v string) *DescribeEnvCustomJobRequest
	GetRegionId() *string
}

type DescribeEnvCustomJobRequest struct {
	// The name of the custom job.
	//
	// This parameter is required.
	//
	// example:
	//
	// customJob1
	CustomJobName *string `json:"CustomJobName,omitempty" xml:"CustomJobName,omitempty"`
	// Specifies whether to return an encrypted YAML string.
	//
	// example:
	//
	// true
	EncryptYaml *bool `json:"EncryptYaml,omitempty" xml:"EncryptYaml,omitempty"`
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

func (s DescribeEnvCustomJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnvCustomJobRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnvCustomJobRequest) GetCustomJobName() *string {
	return s.CustomJobName
}

func (s *DescribeEnvCustomJobRequest) GetEncryptYaml() *bool {
	return s.EncryptYaml
}

func (s *DescribeEnvCustomJobRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *DescribeEnvCustomJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEnvCustomJobRequest) SetCustomJobName(v string) *DescribeEnvCustomJobRequest {
	s.CustomJobName = &v
	return s
}

func (s *DescribeEnvCustomJobRequest) SetEncryptYaml(v bool) *DescribeEnvCustomJobRequest {
	s.EncryptYaml = &v
	return s
}

func (s *DescribeEnvCustomJobRequest) SetEnvironmentId(v string) *DescribeEnvCustomJobRequest {
	s.EnvironmentId = &v
	return s
}

func (s *DescribeEnvCustomJobRequest) SetRegionId(v string) *DescribeEnvCustomJobRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEnvCustomJobRequest) Validate() error {
	return dara.Validate(s)
}
