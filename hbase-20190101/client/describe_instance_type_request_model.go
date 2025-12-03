// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceType(v string) *DescribeInstanceTypeRequest
	GetInstanceType() *string
}

type DescribeInstanceTypeRequest struct {
	// example:
	//
	// hbase.n2.4xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeInstanceTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypeRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeInstanceTypeRequest) SetInstanceType(v string) *DescribeInstanceTypeRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceTypeRequest) Validate() error {
	return dara.Validate(s)
}
