// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouteStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *DeleteRouteStrategyRequest
	GetGroupId() *string
	SetJobId(v int64) *DeleteRouteStrategyRequest
	GetJobId() *int64
	SetNamespace(v string) *DeleteRouteStrategyRequest
	GetNamespace() *string
	SetRegionId(v string) *DeleteRouteStrategyRequest
	GetRegionId() *string
}

type DeleteRouteStrategyRequest struct {
	// The application ID. You can obtain the application ID on the **Application Management*	- page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The job ID. You can obtain the job ID on the **Task Management*	- page in the SchedulerX console.
	//
	// example:
	//
	// 92583
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The namespace ID. You can obtain the ID of the namespace on the **Namespace*	- page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// adcfc35d-e2fe-4fe9-bbaa-20e90ffc****
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteRouteStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteStrategyRequest) GoString() string {
	return s.String()
}

func (s *DeleteRouteStrategyRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DeleteRouteStrategyRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *DeleteRouteStrategyRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteRouteStrategyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteRouteStrategyRequest) SetGroupId(v string) *DeleteRouteStrategyRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteRouteStrategyRequest) SetJobId(v int64) *DeleteRouteStrategyRequest {
	s.JobId = &v
	return s
}

func (s *DeleteRouteStrategyRequest) SetNamespace(v string) *DeleteRouteStrategyRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteRouteStrategyRequest) SetRegionId(v string) *DeleteRouteStrategyRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteRouteStrategyRequest) Validate() error {
	return dara.Validate(s)
}
