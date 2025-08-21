// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceSpecsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *DescribeInstanceSpecsRequest
	GetInstanceIds() []*string
}

type DescribeInstanceSpecsRequest struct {
	// An array that consists of the IDs of instances to query.
	//
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s DescribeInstanceSpecsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSpecsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeInstanceSpecsRequest) SetInstanceIds(v []*string) *DescribeInstanceSpecsRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeInstanceSpecsRequest) Validate() error {
	return dara.Validate(s)
}
