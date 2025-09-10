// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceCronScalerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExcludeDates(v []*string) *UpdateServiceCronScalerRequest
	GetExcludeDates() []*string
	SetScaleJobs(v []*UpdateServiceCronScalerRequestScaleJobs) *UpdateServiceCronScalerRequest
	GetScaleJobs() []*UpdateServiceCronScalerRequestScaleJobs
}

type UpdateServiceCronScalerRequest struct {
	// The points in time that are excluded when you schedule a CronHPA job. The points in time must be specified by using a cron expression.
	ExcludeDates []*string `json:"ExcludeDates,omitempty" xml:"ExcludeDates,omitempty" type:"Repeated"`
	// The description of the CronHPA job.
	//
	// This parameter is required.
	ScaleJobs []*UpdateServiceCronScalerRequestScaleJobs `json:"ScaleJobs,omitempty" xml:"ScaleJobs,omitempty" type:"Repeated"`
}

func (s UpdateServiceCronScalerRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceCronScalerRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceCronScalerRequest) GetExcludeDates() []*string {
	return s.ExcludeDates
}

func (s *UpdateServiceCronScalerRequest) GetScaleJobs() []*UpdateServiceCronScalerRequestScaleJobs {
	return s.ScaleJobs
}

func (s *UpdateServiceCronScalerRequest) SetExcludeDates(v []*string) *UpdateServiceCronScalerRequest {
	s.ExcludeDates = v
	return s
}

func (s *UpdateServiceCronScalerRequest) SetScaleJobs(v []*UpdateServiceCronScalerRequestScaleJobs) *UpdateServiceCronScalerRequest {
	s.ScaleJobs = v
	return s
}

func (s *UpdateServiceCronScalerRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateServiceCronScalerRequestScaleJobs struct {
	// The name of the CronHPA job.
	//
	// example:
	//
	// scale-job-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The cron expression that is used to configure the execution time of the CronHPA job. For more information about how to configure cron expressions, see **Description of special characters*	- in this topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0 18 	- 	- 	- *
	Schedule *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	// The number of instances that you want to configure for the CronHPA job.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	TargetSize *int32  `json:"TargetSize,omitempty" xml:"TargetSize,omitempty"`
	TimeZone   *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s UpdateServiceCronScalerRequestScaleJobs) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceCronScalerRequestScaleJobs) GoString() string {
	return s.String()
}

func (s *UpdateServiceCronScalerRequestScaleJobs) GetName() *string {
	return s.Name
}

func (s *UpdateServiceCronScalerRequestScaleJobs) GetSchedule() *string {
	return s.Schedule
}

func (s *UpdateServiceCronScalerRequestScaleJobs) GetTargetSize() *int32 {
	return s.TargetSize
}

func (s *UpdateServiceCronScalerRequestScaleJobs) GetTimeZone() *string {
	return s.TimeZone
}

func (s *UpdateServiceCronScalerRequestScaleJobs) SetName(v string) *UpdateServiceCronScalerRequestScaleJobs {
	s.Name = &v
	return s
}

func (s *UpdateServiceCronScalerRequestScaleJobs) SetSchedule(v string) *UpdateServiceCronScalerRequestScaleJobs {
	s.Schedule = &v
	return s
}

func (s *UpdateServiceCronScalerRequestScaleJobs) SetTargetSize(v int32) *UpdateServiceCronScalerRequestScaleJobs {
	s.TargetSize = &v
	return s
}

func (s *UpdateServiceCronScalerRequestScaleJobs) SetTimeZone(v string) *UpdateServiceCronScalerRequestScaleJobs {
	s.TimeZone = &v
	return s
}

func (s *UpdateServiceCronScalerRequestScaleJobs) Validate() error {
	return dara.Validate(s)
}
