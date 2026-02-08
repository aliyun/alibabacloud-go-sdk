// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobDataParsingTaskProgressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeJobDataParsingTaskProgressRequest
	GetInstanceId() *string
	SetJobDataParsingTaskId(v string) *DescribeJobDataParsingTaskProgressRequest
	GetJobDataParsingTaskId() *string
}

type DescribeJobDataParsingTaskProgressRequest struct {
	// instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// c209abb3-6804-4a75-b2c7-dd55c8c61b6a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Job data analytics task ID
	//
	// > You can obtain this parameter value by calling the CreateJobDataParsingTask API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 50d5e164-9365-4261-980e-3d979c2c948c
	JobDataParsingTaskId *string `json:"JobDataParsingTaskId,omitempty" xml:"JobDataParsingTaskId,omitempty"`
}

func (s DescribeJobDataParsingTaskProgressRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobDataParsingTaskProgressRequest) GoString() string {
	return s.String()
}

func (s *DescribeJobDataParsingTaskProgressRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeJobDataParsingTaskProgressRequest) GetJobDataParsingTaskId() *string {
	return s.JobDataParsingTaskId
}

func (s *DescribeJobDataParsingTaskProgressRequest) SetInstanceId(v string) *DescribeJobDataParsingTaskProgressRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeJobDataParsingTaskProgressRequest) SetJobDataParsingTaskId(v string) *DescribeJobDataParsingTaskProgressRequest {
	s.JobDataParsingTaskId = &v
	return s
}

func (s *DescribeJobDataParsingTaskProgressRequest) Validate() error {
	return dara.Validate(s)
}
