// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupExecutingInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeGroupExecutingInfoRequest
	GetInstanceId() *string
	SetJobGroupId(v string) *DescribeGroupExecutingInfoRequest
	GetJobGroupId() *string
}

type DescribeGroupExecutingInfoRequest struct {
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// c46001bc-3ead-4bfd-9a69-4b5b66a4a3f4
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Task group ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 3640dda7-e5b1-4b3e-9ccf-da4fc5402e11
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
}

func (s DescribeGroupExecutingInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupExecutingInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeGroupExecutingInfoRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeGroupExecutingInfoRequest) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *DescribeGroupExecutingInfoRequest) SetInstanceId(v string) *DescribeGroupExecutingInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeGroupExecutingInfoRequest) SetJobGroupId(v string) *DescribeGroupExecutingInfoRequest {
	s.JobGroupId = &v
	return s
}

func (s *DescribeGroupExecutingInfoRequest) Validate() error {
	return dara.Validate(s)
}
