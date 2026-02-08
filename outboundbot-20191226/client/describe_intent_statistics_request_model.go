// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIntentStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeIntentStatisticsRequest
	GetInstanceId() *string
	SetJobGroupId(v string) *DescribeIntentStatisticsRequest
	GetJobGroupId() *string
	SetLimit(v int32) *DescribeIntentStatisticsRequest
	GetLimit() *int32
}

type DescribeIntentStatisticsRequest struct {
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// c3c92de8-e4bd-4db4-a962-50f8acce40bc
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Job group ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 040355a9-e80c-4308-b85c-aa5b9fd25246
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	// Number of statistics entries to display (Required)
	//
	// example:
	//
	// 5
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
}

func (s DescribeIntentStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIntentStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeIntentStatisticsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeIntentStatisticsRequest) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *DescribeIntentStatisticsRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *DescribeIntentStatisticsRequest) SetInstanceId(v string) *DescribeIntentStatisticsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeIntentStatisticsRequest) SetJobGroupId(v string) *DescribeIntentStatisticsRequest {
	s.JobGroupId = &v
	return s
}

func (s *DescribeIntentStatisticsRequest) SetLimit(v int32) *DescribeIntentStatisticsRequest {
	s.Limit = &v
	return s
}

func (s *DescribeIntentStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
