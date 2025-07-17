// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvCustomJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEncryptYaml(v bool) *ListEnvCustomJobsRequest
	GetEncryptYaml() *bool
	SetEnvironmentId(v string) *ListEnvCustomJobsRequest
	GetEnvironmentId() *string
	SetRegionId(v string) *ListEnvCustomJobsRequest
	GetRegionId() *string
}

type ListEnvCustomJobsRequest struct {
	// Specifies whether to return the encrypted YAML string.
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
}

func (s ListEnvCustomJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEnvCustomJobsRequest) GoString() string {
	return s.String()
}

func (s *ListEnvCustomJobsRequest) GetEncryptYaml() *bool {
	return s.EncryptYaml
}

func (s *ListEnvCustomJobsRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *ListEnvCustomJobsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEnvCustomJobsRequest) SetEncryptYaml(v bool) *ListEnvCustomJobsRequest {
	s.EncryptYaml = &v
	return s
}

func (s *ListEnvCustomJobsRequest) SetEnvironmentId(v string) *ListEnvCustomJobsRequest {
	s.EnvironmentId = &v
	return s
}

func (s *ListEnvCustomJobsRequest) SetRegionId(v string) *ListEnvCustomJobsRequest {
	s.RegionId = &v
	return s
}

func (s *ListEnvCustomJobsRequest) Validate() error {
	return dara.Validate(s)
}
