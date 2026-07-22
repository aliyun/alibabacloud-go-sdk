// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScheduledTaskExecutionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeScheduledTaskExecutionsRequest
	GetEndTime() *string
	SetInstanceId(v string) *DescribeScheduledTaskExecutionsRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *DescribeScheduledTaskExecutionsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeScheduledTaskExecutionsRequest
	GetNextToken() *string
	SetScheduledId(v string) *DescribeScheduledTaskExecutionsRequest
	GetScheduledId() *string
	SetStartTime(v string) *DescribeScheduledTaskExecutionsRequest
	GetStartTime() *string
	SetStatus(v string) *DescribeScheduledTaskExecutionsRequest
	GetStatus() *string
}

type DescribeScheduledTaskExecutionsRequest struct {
	// The end time of the time range in ISO-8601 format.
	//
	// example:
	//
	// 2026-06-12T23:59:59
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// acp-axxkuuxahbu1*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of results to return per request. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Leave this parameter empty for the first request.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kU****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the scheduled task.
	//
	// This parameter is required.
	//
	// example:
	//
	// sch-260705-agb*****
	ScheduledId *string `json:"ScheduledId,omitempty" xml:"ScheduledId,omitempty"`
	// The start time of the time range in ISO-8601 format.
	//
	// example:
	//
	// 2026-06-01T00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the scheduled task.
	//
	// example:
	//
	// COMPLETED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeScheduledTaskExecutionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeScheduledTaskExecutionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeScheduledTaskExecutionsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeScheduledTaskExecutionsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeScheduledTaskExecutionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeScheduledTaskExecutionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeScheduledTaskExecutionsRequest) GetScheduledId() *string {
	return s.ScheduledId
}

func (s *DescribeScheduledTaskExecutionsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeScheduledTaskExecutionsRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeScheduledTaskExecutionsRequest) SetEndTime(v string) *DescribeScheduledTaskExecutionsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeScheduledTaskExecutionsRequest) SetInstanceId(v string) *DescribeScheduledTaskExecutionsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeScheduledTaskExecutionsRequest) SetMaxResults(v int32) *DescribeScheduledTaskExecutionsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeScheduledTaskExecutionsRequest) SetNextToken(v string) *DescribeScheduledTaskExecutionsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeScheduledTaskExecutionsRequest) SetScheduledId(v string) *DescribeScheduledTaskExecutionsRequest {
	s.ScheduledId = &v
	return s
}

func (s *DescribeScheduledTaskExecutionsRequest) SetStartTime(v string) *DescribeScheduledTaskExecutionsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeScheduledTaskExecutionsRequest) SetStatus(v string) *DescribeScheduledTaskExecutionsRequest {
	s.Status = &v
	return s
}

func (s *DescribeScheduledTaskExecutionsRequest) Validate() error {
	return dara.Validate(s)
}
