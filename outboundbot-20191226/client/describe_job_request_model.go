// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeJobRequest
	GetInstanceId() *string
	SetJobId(v string) *DescribeJobRequest
	GetJobId() *string
	SetWithScript(v bool) *DescribeJobRequest
	GetWithScript() *bool
}

type DescribeJobRequest struct {
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c8a2b7f2-ad1a-4865-b872-d0080d9802d9
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// Indicates whether to return job scenario information.
	//
	// example:
	//
	// false
	WithScript *bool `json:"WithScript,omitempty" xml:"WithScript,omitempty"`
}

func (s DescribeJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobRequest) GoString() string {
	return s.String()
}

func (s *DescribeJobRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *DescribeJobRequest) GetWithScript() *bool {
	return s.WithScript
}

func (s *DescribeJobRequest) SetInstanceId(v string) *DescribeJobRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeJobRequest) SetJobId(v string) *DescribeJobRequest {
	s.JobId = &v
	return s
}

func (s *DescribeJobRequest) SetWithScript(v bool) *DescribeJobRequest {
	s.WithScript = &v
	return s
}

func (s *DescribeJobRequest) Validate() error {
	return dara.Validate(s)
}
