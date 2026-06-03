// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDialogueNodeStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeDialogueNodeStatisticsRequest
	GetInstanceId() *string
	SetJobGroupId(v string) *DescribeDialogueNodeStatisticsRequest
	GetJobGroupId() *string
	SetLimit(v int32) *DescribeDialogueNodeStatisticsRequest
	GetLimit() *int32
}

type DescribeDialogueNodeStatisticsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// aeff669b-388f-4619-82af-81e177df5628
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a3c670d1-01bf-491d-b9aa-759b1a82f47c
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
}

func (s DescribeDialogueNodeStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDialogueNodeStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDialogueNodeStatisticsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDialogueNodeStatisticsRequest) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *DescribeDialogueNodeStatisticsRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *DescribeDialogueNodeStatisticsRequest) SetInstanceId(v string) *DescribeDialogueNodeStatisticsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDialogueNodeStatisticsRequest) SetJobGroupId(v string) *DescribeDialogueNodeStatisticsRequest {
	s.JobGroupId = &v
	return s
}

func (s *DescribeDialogueNodeStatisticsRequest) SetLimit(v int32) *DescribeDialogueNodeStatisticsRequest {
	s.Limit = &v
	return s
}

func (s *DescribeDialogueNodeStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
