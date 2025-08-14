// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRtcRobotInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeRtcRobotInstanceRequest
	GetInstanceId() *string
}

type DescribeRtcRobotInstanceRequest struct {
	// example:
	//
	// 727dc0e296014bb58670940a3da95592
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeRtcRobotInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcRobotInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeRtcRobotInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRtcRobotInstanceRequest) SetInstanceId(v string) *DescribeRtcRobotInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRtcRobotInstanceRequest) Validate() error {
	return dara.Validate(s)
}
