// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoScalingConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeAutoScalingConfigRequest
	GetInstanceId() *string
}

type DescribeAutoScalingConfigRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-2ze1prap1k46r****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeAutoScalingConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoScalingConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutoScalingConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAutoScalingConfigRequest) SetInstanceId(v string) *DescribeAutoScalingConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAutoScalingConfigRequest) Validate() error {
	return dara.Validate(s)
}
