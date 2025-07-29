// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRouteStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *CreateRouteStrategyRequest
	GetGroupId() *string
	SetJobId(v int64) *CreateRouteStrategyRequest
	GetJobId() *int64
	SetName(v string) *CreateRouteStrategyRequest
	GetName() *string
	SetNamespace(v string) *CreateRouteStrategyRequest
	GetNamespace() *string
	SetRegionId(v string) *CreateRouteStrategyRequest
	GetRegionId() *string
	SetStatus(v int32) *CreateRouteStrategyRequest
	GetStatus() *int32
	SetStrategyContent(v string) *CreateRouteStrategyRequest
	GetStrategyContent() *string
	SetType(v int32) *CreateRouteStrategyRequest
	GetType() *int32
}

type CreateRouteStrategyRequest struct {
	// The ID of the application group. You can obtain the ID on the **Application Management*	- page in the SchedulerX console.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSchedulerx.defaultGroup
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The job ID. You can obtain the ID on the **Task Management*	- page in the SchedulerX console.
	//
	// example:
	//
	// 54978
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The name of the routing policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-strategy
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace ID. You can obtain the namespace ID on the **Namespace*	- page in the SchedulerX console.
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
	// Specifies whether to enable the routing policy. Valid values:
	//
	// 	- **0**: disables the routing policy.
	//
	// 	- **1**: enables the routing policy.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The details of the routing policy. The value is a JSON string. For more information about this parameter, see **the additional information about request parameters*	- below this table.
	//
	// example:
	//
	// [{"percentage":20,"target":"[\\"version1\\"]","targetType":"label"}]
	StrategyContent *string `json:"StrategyContent,omitempty" xml:"StrategyContent,omitempty"`
	// The type of the routing policy. Valid value:
	//
	// 	- **3**: routes by proportion.
	//
	// example:
	//
	// 3
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateRouteStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRouteStrategyRequest) GoString() string {
	return s.String()
}

func (s *CreateRouteStrategyRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateRouteStrategyRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *CreateRouteStrategyRequest) GetName() *string {
	return s.Name
}

func (s *CreateRouteStrategyRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateRouteStrategyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateRouteStrategyRequest) GetStatus() *int32 {
	return s.Status
}

func (s *CreateRouteStrategyRequest) GetStrategyContent() *string {
	return s.StrategyContent
}

func (s *CreateRouteStrategyRequest) GetType() *int32 {
	return s.Type
}

func (s *CreateRouteStrategyRequest) SetGroupId(v string) *CreateRouteStrategyRequest {
	s.GroupId = &v
	return s
}

func (s *CreateRouteStrategyRequest) SetJobId(v int64) *CreateRouteStrategyRequest {
	s.JobId = &v
	return s
}

func (s *CreateRouteStrategyRequest) SetName(v string) *CreateRouteStrategyRequest {
	s.Name = &v
	return s
}

func (s *CreateRouteStrategyRequest) SetNamespace(v string) *CreateRouteStrategyRequest {
	s.Namespace = &v
	return s
}

func (s *CreateRouteStrategyRequest) SetRegionId(v string) *CreateRouteStrategyRequest {
	s.RegionId = &v
	return s
}

func (s *CreateRouteStrategyRequest) SetStatus(v int32) *CreateRouteStrategyRequest {
	s.Status = &v
	return s
}

func (s *CreateRouteStrategyRequest) SetStrategyContent(v string) *CreateRouteStrategyRequest {
	s.StrategyContent = &v
	return s
}

func (s *CreateRouteStrategyRequest) SetType(v int32) *CreateRouteStrategyRequest {
	s.Type = &v
	return s
}

func (s *CreateRouteStrategyRequest) Validate() error {
	return dara.Validate(s)
}
