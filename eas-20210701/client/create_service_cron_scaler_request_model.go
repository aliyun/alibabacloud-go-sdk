// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceCronScalerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExcludeDates(v []*string) *CreateServiceCronScalerRequest
	GetExcludeDates() []*string
	SetScaleJobs(v []*CreateServiceCronScalerRequestScaleJobs) *CreateServiceCronScalerRequest
	GetScaleJobs() []*CreateServiceCronScalerRequestScaleJobs
}

type CreateServiceCronScalerRequest struct {
	// The points in time that are excluded when you schedule a CronHPA job. The points in time must be specified by using a cron expression.
	ExcludeDates []*string `json:"ExcludeDates,omitempty" xml:"ExcludeDates,omitempty" type:"Repeated"`
	// The description of the CronHPA job.
	//
	// This parameter is required.
	ScaleJobs []*CreateServiceCronScalerRequestScaleJobs `json:"ScaleJobs,omitempty" xml:"ScaleJobs,omitempty" type:"Repeated"`
}

func (s CreateServiceCronScalerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceCronScalerRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceCronScalerRequest) GetExcludeDates() []*string {
	return s.ExcludeDates
}

func (s *CreateServiceCronScalerRequest) GetScaleJobs() []*CreateServiceCronScalerRequestScaleJobs {
	return s.ScaleJobs
}

func (s *CreateServiceCronScalerRequest) SetExcludeDates(v []*string) *CreateServiceCronScalerRequest {
	s.ExcludeDates = v
	return s
}

func (s *CreateServiceCronScalerRequest) SetScaleJobs(v []*CreateServiceCronScalerRequestScaleJobs) *CreateServiceCronScalerRequest {
	s.ScaleJobs = v
	return s
}

func (s *CreateServiceCronScalerRequest) Validate() error {
	if s.ScaleJobs != nil {
		for _, item := range s.ScaleJobs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateServiceCronScalerRequestScaleJobs struct {
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
	// 1
	TargetSize *int32  `json:"TargetSize,omitempty" xml:"TargetSize,omitempty"`
	TimeZone   *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s CreateServiceCronScalerRequestScaleJobs) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceCronScalerRequestScaleJobs) GoString() string {
	return s.String()
}

func (s *CreateServiceCronScalerRequestScaleJobs) GetName() *string {
	return s.Name
}

func (s *CreateServiceCronScalerRequestScaleJobs) GetSchedule() *string {
	return s.Schedule
}

func (s *CreateServiceCronScalerRequestScaleJobs) GetTargetSize() *int32 {
	return s.TargetSize
}

func (s *CreateServiceCronScalerRequestScaleJobs) GetTimeZone() *string {
	return s.TimeZone
}

func (s *CreateServiceCronScalerRequestScaleJobs) SetName(v string) *CreateServiceCronScalerRequestScaleJobs {
	s.Name = &v
	return s
}

func (s *CreateServiceCronScalerRequestScaleJobs) SetSchedule(v string) *CreateServiceCronScalerRequestScaleJobs {
	s.Schedule = &v
	return s
}

func (s *CreateServiceCronScalerRequestScaleJobs) SetTargetSize(v int32) *CreateServiceCronScalerRequestScaleJobs {
	s.TargetSize = &v
	return s
}

func (s *CreateServiceCronScalerRequestScaleJobs) SetTimeZone(v string) *CreateServiceCronScalerRequestScaleJobs {
	s.TimeZone = &v
	return s
}

func (s *CreateServiceCronScalerRequestScaleJobs) Validate() error {
	return dara.Validate(s)
}
