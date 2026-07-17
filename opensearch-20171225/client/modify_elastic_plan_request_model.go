// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyElasticPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomDates(v []*string) *ModifyElasticPlanRequest
	GetCustomDates() []*string
	SetDescription(v string) *ModifyElasticPlanRequest
	GetDescription() *string
	SetElasticLcu(v int32) *ModifyElasticPlanRequest
	GetElasticLcu() *int32
	SetEnabled(v bool) *ModifyElasticPlanRequest
	GetEnabled() *bool
	SetEndHour(v int32) *ModifyElasticPlanRequest
	GetEndHour() *int32
	SetScheduleType(v string) *ModifyElasticPlanRequest
	GetScheduleType() *string
	SetStartHour(v int32) *ModifyElasticPlanRequest
	GetStartHour() *int32
	SetDryRun(v bool) *ModifyElasticPlanRequest
	GetDryRun() *bool
}

type ModifyElasticPlanRequest struct {
	CustomDates []*string `json:"customDates,omitempty" xml:"customDates,omitempty" type:"Repeated"`
	// example:
	//
	// Updated description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 100
	ElasticLcu *int32 `json:"elasticLcu,omitempty" xml:"elasticLcu,omitempty"`
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// example:
	//
	// 16
	EndHour *int32 `json:"endHour,omitempty" xml:"endHour,omitempty"`
	// example:
	//
	// WEEK
	ScheduleType *string `json:"scheduleType,omitempty" xml:"scheduleType,omitempty"`
	// example:
	//
	// 9
	StartHour *int32 `json:"startHour,omitempty" xml:"startHour,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s ModifyElasticPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyElasticPlanRequest) GoString() string {
	return s.String()
}

func (s *ModifyElasticPlanRequest) GetCustomDates() []*string {
	return s.CustomDates
}

func (s *ModifyElasticPlanRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyElasticPlanRequest) GetElasticLcu() *int32 {
	return s.ElasticLcu
}

func (s *ModifyElasticPlanRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *ModifyElasticPlanRequest) GetEndHour() *int32 {
	return s.EndHour
}

func (s *ModifyElasticPlanRequest) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *ModifyElasticPlanRequest) GetStartHour() *int32 {
	return s.StartHour
}

func (s *ModifyElasticPlanRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyElasticPlanRequest) SetCustomDates(v []*string) *ModifyElasticPlanRequest {
	s.CustomDates = v
	return s
}

func (s *ModifyElasticPlanRequest) SetDescription(v string) *ModifyElasticPlanRequest {
	s.Description = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticLcu(v int32) *ModifyElasticPlanRequest {
	s.ElasticLcu = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetEnabled(v bool) *ModifyElasticPlanRequest {
	s.Enabled = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetEndHour(v int32) *ModifyElasticPlanRequest {
	s.EndHour = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetScheduleType(v string) *ModifyElasticPlanRequest {
	s.ScheduleType = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetStartHour(v int32) *ModifyElasticPlanRequest {
	s.StartHour = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetDryRun(v bool) *ModifyElasticPlanRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyElasticPlanRequest) Validate() error {
	return dara.Validate(s)
}
