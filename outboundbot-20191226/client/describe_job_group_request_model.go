// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBriefTypes(v []*string) *DescribeJobGroupRequest
	GetBriefTypes() []*string
	SetInstanceId(v string) *DescribeJobGroupRequest
	GetInstanceId() *string
	SetJobGroupId(v string) *DescribeJobGroupRequest
	GetJobGroupId() *string
}

type DescribeJobGroupRequest struct {
	// Filter condition (historical parameter, deprecated).
	BriefTypes []*string `json:"BriefTypes,omitempty" xml:"BriefTypes,omitempty" type:"Repeated"`
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Task group ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 46a9ad0c-3e11-44da-a9a7-2c21bf5ce185
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
}

func (s DescribeJobGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeJobGroupRequest) GetBriefTypes() []*string {
	return s.BriefTypes
}

func (s *DescribeJobGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeJobGroupRequest) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *DescribeJobGroupRequest) SetBriefTypes(v []*string) *DescribeJobGroupRequest {
	s.BriefTypes = v
	return s
}

func (s *DescribeJobGroupRequest) SetInstanceId(v string) *DescribeJobGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeJobGroupRequest) SetJobGroupId(v string) *DescribeJobGroupRequest {
	s.JobGroupId = &v
	return s
}

func (s *DescribeJobGroupRequest) Validate() error {
	return dara.Validate(s)
}
