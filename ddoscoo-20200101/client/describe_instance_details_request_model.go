// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *DescribeInstanceDetailsRequest
	GetInstanceIds() []*string
}

type DescribeInstanceDetailsRequest struct {
	// An array that consists of the IDs of instances to query.
	//
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s DescribeInstanceDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDetailsRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeInstanceDetailsRequest) SetInstanceIds(v []*string) *DescribeInstanceDetailsRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeInstanceDetailsRequest) Validate() error {
	return dara.Validate(s)
}
