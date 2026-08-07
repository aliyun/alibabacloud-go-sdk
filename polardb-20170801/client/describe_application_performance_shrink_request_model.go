// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationPerformanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribeApplicationPerformanceShrinkRequest
	GetApplicationId() *string
	SetConsumer(v string) *DescribeApplicationPerformanceShrinkRequest
	GetConsumer() *string
	SetConsumerGroup(v string) *DescribeApplicationPerformanceShrinkRequest
	GetConsumerGroup() *string
	SetDownsample(v string) *DescribeApplicationPerformanceShrinkRequest
	GetDownsample() *string
	SetEndStep(v int64) *DescribeApplicationPerformanceShrinkRequest
	GetEndStep() *int64
	SetEndTime(v string) *DescribeApplicationPerformanceShrinkRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeApplicationPerformanceShrinkRequest
	GetInterval() *string
	SetKey(v string) *DescribeApplicationPerformanceShrinkRequest
	GetKey() *string
	SetMaxPoints(v int32) *DescribeApplicationPerformanceShrinkRequest
	GetMaxPoints() *int32
	SetModelService(v string) *DescribeApplicationPerformanceShrinkRequest
	GetModelService() *string
	SetStartStep(v int64) *DescribeApplicationPerformanceShrinkRequest
	GetStartStep() *int64
	SetStartTime(v string) *DescribeApplicationPerformanceShrinkRequest
	GetStartTime() *string
	SetFilterShrink(v string) *DescribeApplicationPerformanceShrinkRequest
	GetFilterShrink() *string
}

type DescribeApplicationPerformanceShrinkRequest struct {
	// The application cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The user.
	//
	// example:
	//
	// c-xxxxxxx
	Consumer *string `json:"Consumer,omitempty" xml:"Consumer,omitempty"`
	// The user group.
	//
	// example:
	//
	// cg-xxxxxx
	ConsumerGroup *string `json:"ConsumerGroup,omitempty" xml:"ConsumerGroup,omitempty"`
	// The downsampling policy.
	//
	// example:
	//
	// raw_sample
	Downsample *string `json:"Downsample,omitempty" xml:"Downsample,omitempty"`
	// The end step number.
	//
	// example:
	//
	// 100
	EndStep *int64 `json:"EndStep,omitempty" xml:"EndStep,omitempty"`
	// The end of the time range to query. Specify the time in the yyyy-MM-ddTHH:mmZ format (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-09-23T01:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The data granularity of performance data. Valid values:
	//
	// - 5
	//
	// - 30
	//
	// - 60
	//
	// - 600
	//
	// - 1800
	//
	// - 3600
	//
	// - 86400
	//
	// example:
	//
	// 5
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The performance metrics to query. Separate multiple values with commas (,).
	//
	// > **Note*	- You can specify up to 5 performance metrics.
	//
	// This parameter is required.
	//
	// example:
	//
	// PolarDBSupabaseMemUsage
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The maximum number of data points to return.
	//
	// example:
	//
	// 1000
	MaxPoints *int32 `json:"MaxPoints,omitempty" xml:"MaxPoints,omitempty"`
	// The model service.
	//
	// example:
	//
	// ms-xxxxxx
	ModelService *string `json:"ModelService,omitempty" xml:"ModelService,omitempty"`
	// The start step number.
	//
	// example:
	//
	// 1
	StartStep *int64 `json:"StartStep,omitempty" xml:"StartStep,omitempty"`
	// The beginning of the time range to query. Specify the time in the yyyy-MM-ddTHH:mmZ format (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-09-23T01:01Z
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	FilterShrink *string `json:"filter,omitempty" xml:"filter,omitempty"`
}

func (s DescribeApplicationPerformanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationPerformanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationPerformanceShrinkRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribeApplicationPerformanceShrinkRequest) GetConsumer() *string {
	return s.Consumer
}

func (s *DescribeApplicationPerformanceShrinkRequest) GetConsumerGroup() *string {
	return s.ConsumerGroup
}

func (s *DescribeApplicationPerformanceShrinkRequest) GetDownsample() *string {
	return s.Downsample
}

func (s *DescribeApplicationPerformanceShrinkRequest) GetEndStep() *int64 {
	return s.EndStep
}

func (s *DescribeApplicationPerformanceShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeApplicationPerformanceShrinkRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeApplicationPerformanceShrinkRequest) GetKey() *string {
	return s.Key
}

func (s *DescribeApplicationPerformanceShrinkRequest) GetMaxPoints() *int32 {
	return s.MaxPoints
}

func (s *DescribeApplicationPerformanceShrinkRequest) GetModelService() *string {
	return s.ModelService
}

func (s *DescribeApplicationPerformanceShrinkRequest) GetStartStep() *int64 {
	return s.StartStep
}

func (s *DescribeApplicationPerformanceShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeApplicationPerformanceShrinkRequest) GetFilterShrink() *string {
	return s.FilterShrink
}

func (s *DescribeApplicationPerformanceShrinkRequest) SetApplicationId(v string) *DescribeApplicationPerformanceShrinkRequest {
	s.ApplicationId = &v
	return s
}

func (s *DescribeApplicationPerformanceShrinkRequest) SetConsumer(v string) *DescribeApplicationPerformanceShrinkRequest {
	s.Consumer = &v
	return s
}

func (s *DescribeApplicationPerformanceShrinkRequest) SetConsumerGroup(v string) *DescribeApplicationPerformanceShrinkRequest {
	s.ConsumerGroup = &v
	return s
}

func (s *DescribeApplicationPerformanceShrinkRequest) SetDownsample(v string) *DescribeApplicationPerformanceShrinkRequest {
	s.Downsample = &v
	return s
}

func (s *DescribeApplicationPerformanceShrinkRequest) SetEndStep(v int64) *DescribeApplicationPerformanceShrinkRequest {
	s.EndStep = &v
	return s
}

func (s *DescribeApplicationPerformanceShrinkRequest) SetEndTime(v string) *DescribeApplicationPerformanceShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeApplicationPerformanceShrinkRequest) SetInterval(v string) *DescribeApplicationPerformanceShrinkRequest {
	s.Interval = &v
	return s
}

func (s *DescribeApplicationPerformanceShrinkRequest) SetKey(v string) *DescribeApplicationPerformanceShrinkRequest {
	s.Key = &v
	return s
}

func (s *DescribeApplicationPerformanceShrinkRequest) SetMaxPoints(v int32) *DescribeApplicationPerformanceShrinkRequest {
	s.MaxPoints = &v
	return s
}

func (s *DescribeApplicationPerformanceShrinkRequest) SetModelService(v string) *DescribeApplicationPerformanceShrinkRequest {
	s.ModelService = &v
	return s
}

func (s *DescribeApplicationPerformanceShrinkRequest) SetStartStep(v int64) *DescribeApplicationPerformanceShrinkRequest {
	s.StartStep = &v
	return s
}

func (s *DescribeApplicationPerformanceShrinkRequest) SetStartTime(v string) *DescribeApplicationPerformanceShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeApplicationPerformanceShrinkRequest) SetFilterShrink(v string) *DescribeApplicationPerformanceShrinkRequest {
	s.FilterShrink = &v
	return s
}

func (s *DescribeApplicationPerformanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
