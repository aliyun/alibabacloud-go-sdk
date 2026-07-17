// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateElasticPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomDates(v []*string) *CreateElasticPlanRequest
	GetCustomDates() []*string
	SetDescription(v string) *CreateElasticPlanRequest
	GetDescription() *string
	SetElasticLcu(v int32) *CreateElasticPlanRequest
	GetElasticLcu() *int32
	SetEndHour(v int32) *CreateElasticPlanRequest
	GetEndHour() *int32
	SetName(v string) *CreateElasticPlanRequest
	GetName() *string
	SetScheduleType(v string) *CreateElasticPlanRequest
	GetScheduleType() *string
	SetStartHour(v int32) *CreateElasticPlanRequest
	GetStartHour() *int32
	SetDryRun(v bool) *CreateElasticPlanRequest
	GetDryRun() *bool
}

type CreateElasticPlanRequest struct {
	CustomDates []*string `json:"customDates,omitempty" xml:"customDates,omitempty" type:"Repeated"`
	// example:
	//
	// my elastic plan
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 100
	ElasticLcu *int32 `json:"elasticLcu,omitempty" xml:"elasticLcu,omitempty"`
	// example:
	//
	// 13
	EndHour *int32 `json:"endHour,omitempty" xml:"endHour,omitempty"`
	// example:
	//
	// elastic_plan_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// WEEK
	ScheduleType *string `json:"scheduleType,omitempty" xml:"scheduleType,omitempty"`
	// example:
	//
	// 0
	StartHour *int32 `json:"startHour,omitempty" xml:"startHour,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateElasticPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateElasticPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateElasticPlanRequest) GetCustomDates() []*string {
	return s.CustomDates
}

func (s *CreateElasticPlanRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateElasticPlanRequest) GetElasticLcu() *int32 {
	return s.ElasticLcu
}

func (s *CreateElasticPlanRequest) GetEndHour() *int32 {
	return s.EndHour
}

func (s *CreateElasticPlanRequest) GetName() *string {
	return s.Name
}

func (s *CreateElasticPlanRequest) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *CreateElasticPlanRequest) GetStartHour() *int32 {
	return s.StartHour
}

func (s *CreateElasticPlanRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateElasticPlanRequest) SetCustomDates(v []*string) *CreateElasticPlanRequest {
	s.CustomDates = v
	return s
}

func (s *CreateElasticPlanRequest) SetDescription(v string) *CreateElasticPlanRequest {
	s.Description = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticLcu(v int32) *CreateElasticPlanRequest {
	s.ElasticLcu = &v
	return s
}

func (s *CreateElasticPlanRequest) SetEndHour(v int32) *CreateElasticPlanRequest {
	s.EndHour = &v
	return s
}

func (s *CreateElasticPlanRequest) SetName(v string) *CreateElasticPlanRequest {
	s.Name = &v
	return s
}

func (s *CreateElasticPlanRequest) SetScheduleType(v string) *CreateElasticPlanRequest {
	s.ScheduleType = &v
	return s
}

func (s *CreateElasticPlanRequest) SetStartHour(v int32) *CreateElasticPlanRequest {
	s.StartHour = &v
	return s
}

func (s *CreateElasticPlanRequest) SetDryRun(v bool) *CreateElasticPlanRequest {
	s.DryRun = &v
	return s
}

func (s *CreateElasticPlanRequest) Validate() error {
	return dara.Validate(s)
}
