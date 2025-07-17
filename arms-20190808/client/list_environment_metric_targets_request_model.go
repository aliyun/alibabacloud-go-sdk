// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvironmentMetricTargetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironmentId(v string) *ListEnvironmentMetricTargetsRequest
	GetEnvironmentId() *string
	SetJobName(v string) *ListEnvironmentMetricTargetsRequest
	GetJobName() *string
	SetRegionId(v string) *ListEnvironmentMetricTargetsRequest
	GetRegionId() *string
}

type ListEnvironmentMetricTargetsRequest struct {
	// The environment ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// env-xxx
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// The job name.
	//
	// example:
	//
	// blackbox
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListEnvironmentMetricTargetsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentMetricTargetsRequest) GoString() string {
	return s.String()
}

func (s *ListEnvironmentMetricTargetsRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *ListEnvironmentMetricTargetsRequest) GetJobName() *string {
	return s.JobName
}

func (s *ListEnvironmentMetricTargetsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEnvironmentMetricTargetsRequest) SetEnvironmentId(v string) *ListEnvironmentMetricTargetsRequest {
	s.EnvironmentId = &v
	return s
}

func (s *ListEnvironmentMetricTargetsRequest) SetJobName(v string) *ListEnvironmentMetricTargetsRequest {
	s.JobName = &v
	return s
}

func (s *ListEnvironmentMetricTargetsRequest) SetRegionId(v string) *ListEnvironmentMetricTargetsRequest {
	s.RegionId = &v
	return s
}

func (s *ListEnvironmentMetricTargetsRequest) Validate() error {
	return dara.Validate(s)
}
