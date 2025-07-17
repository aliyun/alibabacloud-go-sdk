// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnvCustomJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomJobName(v string) *DeleteEnvCustomJobRequest
	GetCustomJobName() *string
	SetEnvironmentId(v string) *DeleteEnvCustomJobRequest
	GetEnvironmentId() *string
	SetRegionId(v string) *DeleteEnvCustomJobRequest
	GetRegionId() *string
}

type DeleteEnvCustomJobRequest struct {
	// The name of the custom job.
	//
	// This parameter is required.
	//
	// example:
	//
	// job1
	CustomJobName *string `json:"CustomJobName,omitempty" xml:"CustomJobName,omitempty"`
	// The ID of the environment instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// env-xxx
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

func (s DeleteEnvCustomJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnvCustomJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteEnvCustomJobRequest) GetCustomJobName() *string {
	return s.CustomJobName
}

func (s *DeleteEnvCustomJobRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *DeleteEnvCustomJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteEnvCustomJobRequest) SetCustomJobName(v string) *DeleteEnvCustomJobRequest {
	s.CustomJobName = &v
	return s
}

func (s *DeleteEnvCustomJobRequest) SetEnvironmentId(v string) *DeleteEnvCustomJobRequest {
	s.EnvironmentId = &v
	return s
}

func (s *DeleteEnvCustomJobRequest) SetRegionId(v string) *DeleteEnvCustomJobRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteEnvCustomJobRequest) Validate() error {
	return dara.Validate(s)
}
