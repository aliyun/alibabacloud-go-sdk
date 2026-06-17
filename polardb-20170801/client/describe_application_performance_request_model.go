// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationPerformanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribeApplicationPerformanceRequest
	GetApplicationId() *string
	SetConsumer(v string) *DescribeApplicationPerformanceRequest
	GetConsumer() *string
	SetConsumerGroup(v string) *DescribeApplicationPerformanceRequest
	GetConsumerGroup() *string
	SetEndTime(v string) *DescribeApplicationPerformanceRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeApplicationPerformanceRequest
	GetInterval() *string
	SetKey(v string) *DescribeApplicationPerformanceRequest
	GetKey() *string
	SetModelService(v string) *DescribeApplicationPerformanceRequest
	GetModelService() *string
	SetStartTime(v string) *DescribeApplicationPerformanceRequest
	GetStartTime() *string
}

type DescribeApplicationPerformanceRequest struct {
	// The ID of the application cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the consumer.
	//
	// example:
	//
	// c-xxxxxxx
	Consumer *string `json:"Consumer,omitempty" xml:"Consumer,omitempty"`
	// The ID of the consumer group.
	//
	// example:
	//
	// cg-xxxxxx
	ConsumerGroup *string `json:"ConsumerGroup,omitempty" xml:"ConsumerGroup,omitempty"`
	// The end time for the query. Specify the time in UTC in the `yyyy-MM-ddTHH:mmZ` format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-09-23T01:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The granularity of the performance data, in seconds. Valid values:
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
	// The performance metrics to query. Separate multiple metrics with commas (,).<br>You can specify up to five performance metrics.<br>
	//
	// This parameter is required.
	//
	// example:
	//
	// PolarDBSupabaseMemUsage
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The ID of the model service.
	//
	// example:
	//
	// ms-xxxxxx
	ModelService *string `json:"ModelService,omitempty" xml:"ModelService,omitempty"`
	// The start time for the query. Specify the time in UTC in the `yyyy-MM-ddTHH:mmZ` format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-09-23T01:01Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeApplicationPerformanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationPerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationPerformanceRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribeApplicationPerformanceRequest) GetConsumer() *string {
	return s.Consumer
}

func (s *DescribeApplicationPerformanceRequest) GetConsumerGroup() *string {
	return s.ConsumerGroup
}

func (s *DescribeApplicationPerformanceRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeApplicationPerformanceRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeApplicationPerformanceRequest) GetKey() *string {
	return s.Key
}

func (s *DescribeApplicationPerformanceRequest) GetModelService() *string {
	return s.ModelService
}

func (s *DescribeApplicationPerformanceRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeApplicationPerformanceRequest) SetApplicationId(v string) *DescribeApplicationPerformanceRequest {
	s.ApplicationId = &v
	return s
}

func (s *DescribeApplicationPerformanceRequest) SetConsumer(v string) *DescribeApplicationPerformanceRequest {
	s.Consumer = &v
	return s
}

func (s *DescribeApplicationPerformanceRequest) SetConsumerGroup(v string) *DescribeApplicationPerformanceRequest {
	s.ConsumerGroup = &v
	return s
}

func (s *DescribeApplicationPerformanceRequest) SetEndTime(v string) *DescribeApplicationPerformanceRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeApplicationPerformanceRequest) SetInterval(v string) *DescribeApplicationPerformanceRequest {
	s.Interval = &v
	return s
}

func (s *DescribeApplicationPerformanceRequest) SetKey(v string) *DescribeApplicationPerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribeApplicationPerformanceRequest) SetModelService(v string) *DescribeApplicationPerformanceRequest {
	s.ModelService = &v
	return s
}

func (s *DescribeApplicationPerformanceRequest) SetStartTime(v string) *DescribeApplicationPerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeApplicationPerformanceRequest) Validate() error {
	return dara.Validate(s)
}
