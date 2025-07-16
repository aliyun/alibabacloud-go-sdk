// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceCronScalerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExcludeDates(v []*string) *DescribeServiceCronScalerResponseBody
	GetExcludeDates() []*string
	SetRequestId(v string) *DescribeServiceCronScalerResponseBody
	GetRequestId() *string
	SetScaleJobs(v []*DescribeServiceCronScalerResponseBodyScaleJobs) *DescribeServiceCronScalerResponseBody
	GetScaleJobs() []*DescribeServiceCronScalerResponseBodyScaleJobs
	SetServiceName(v string) *DescribeServiceCronScalerResponseBody
	GetServiceName() *string
}

type DescribeServiceCronScalerResponseBody struct {
	// The points in time that are excluded when you schedule a CronHPA job. The points in time must be specified by using a cron expression.
	ExcludeDates []*string `json:"ExcludeDates,omitempty" xml:"ExcludeDates,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The CronHPA jobs.
	ScaleJobs []*DescribeServiceCronScalerResponseBodyScaleJobs `json:"ScaleJobs,omitempty" xml:"ScaleJobs,omitempty" type:"Repeated"`
	// The service name.
	//
	// example:
	//
	// foo
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s DescribeServiceCronScalerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceCronScalerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceCronScalerResponseBody) GetExcludeDates() []*string {
	return s.ExcludeDates
}

func (s *DescribeServiceCronScalerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeServiceCronScalerResponseBody) GetScaleJobs() []*DescribeServiceCronScalerResponseBodyScaleJobs {
	return s.ScaleJobs
}

func (s *DescribeServiceCronScalerResponseBody) GetServiceName() *string {
	return s.ServiceName
}

func (s *DescribeServiceCronScalerResponseBody) SetExcludeDates(v []*string) *DescribeServiceCronScalerResponseBody {
	s.ExcludeDates = v
	return s
}

func (s *DescribeServiceCronScalerResponseBody) SetRequestId(v string) *DescribeServiceCronScalerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceCronScalerResponseBody) SetScaleJobs(v []*DescribeServiceCronScalerResponseBodyScaleJobs) *DescribeServiceCronScalerResponseBody {
	s.ScaleJobs = v
	return s
}

func (s *DescribeServiceCronScalerResponseBody) SetServiceName(v string) *DescribeServiceCronScalerResponseBody {
	s.ServiceName = &v
	return s
}

func (s *DescribeServiceCronScalerResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeServiceCronScalerResponseBodyScaleJobs struct {
	// The time when the most recent CronHPA job was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-06-24T02:11:30Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the most recent CronHPA job ran. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-02-24T06:31:00Z
	LastProbeTime *string `json:"LastProbeTime,omitempty" xml:"LastProbeTime,omitempty"`
	// The returned message.
	//
	// example:
	//
	// "cron hpa job scale-jobs-0 executed successfully. current replicas:3, desired replicas:2."
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the CronHPA job.
	//
	// example:
	//
	// scale-job-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The cron expression that is used to configure the execution time of the CronHPA job.
	//
	// example:
	//
	// 0 18 	- 	- 	- *
	Schedule *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	// The status of the most recent CronHPA job.
	//
	// example:
	//
	// Succeed
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The number of instances that you expect to configure for the CronHPA job.
	//
	// example:
	//
	// 1
	TargetSize *int32 `json:"TargetSize,omitempty" xml:"TargetSize,omitempty"`
}

func (s DescribeServiceCronScalerResponseBodyScaleJobs) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceCronScalerResponseBodyScaleJobs) GoString() string {
	return s.String()
}

func (s *DescribeServiceCronScalerResponseBodyScaleJobs) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeServiceCronScalerResponseBodyScaleJobs) GetLastProbeTime() *string {
	return s.LastProbeTime
}

func (s *DescribeServiceCronScalerResponseBodyScaleJobs) GetMessage() *string {
	return s.Message
}

func (s *DescribeServiceCronScalerResponseBodyScaleJobs) GetName() *string {
	return s.Name
}

func (s *DescribeServiceCronScalerResponseBodyScaleJobs) GetSchedule() *string {
	return s.Schedule
}

func (s *DescribeServiceCronScalerResponseBodyScaleJobs) GetState() *string {
	return s.State
}

func (s *DescribeServiceCronScalerResponseBodyScaleJobs) GetTargetSize() *int32 {
	return s.TargetSize
}

func (s *DescribeServiceCronScalerResponseBodyScaleJobs) SetCreateTime(v string) *DescribeServiceCronScalerResponseBodyScaleJobs {
	s.CreateTime = &v
	return s
}

func (s *DescribeServiceCronScalerResponseBodyScaleJobs) SetLastProbeTime(v string) *DescribeServiceCronScalerResponseBodyScaleJobs {
	s.LastProbeTime = &v
	return s
}

func (s *DescribeServiceCronScalerResponseBodyScaleJobs) SetMessage(v string) *DescribeServiceCronScalerResponseBodyScaleJobs {
	s.Message = &v
	return s
}

func (s *DescribeServiceCronScalerResponseBodyScaleJobs) SetName(v string) *DescribeServiceCronScalerResponseBodyScaleJobs {
	s.Name = &v
	return s
}

func (s *DescribeServiceCronScalerResponseBodyScaleJobs) SetSchedule(v string) *DescribeServiceCronScalerResponseBodyScaleJobs {
	s.Schedule = &v
	return s
}

func (s *DescribeServiceCronScalerResponseBodyScaleJobs) SetState(v string) *DescribeServiceCronScalerResponseBodyScaleJobs {
	s.State = &v
	return s
}

func (s *DescribeServiceCronScalerResponseBodyScaleJobs) SetTargetSize(v int32) *DescribeServiceCronScalerResponseBodyScaleJobs {
	s.TargetSize = &v
	return s
}

func (s *DescribeServiceCronScalerResponseBodyScaleJobs) Validate() error {
	return dara.Validate(s)
}
