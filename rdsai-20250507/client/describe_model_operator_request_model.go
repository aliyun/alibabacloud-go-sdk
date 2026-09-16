// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelOperatorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeModelOperatorRequest
	GetInstanceId() *string
	SetRegion(v string) *DescribeModelOperatorRequest
	GetRegion() *string
}

type DescribeModelOperatorRequest struct {
	// The instance name.
	//
	// example:
	//
	// rds_copilot***_public_cn-*********6
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DescribeModelOperatorRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelOperatorRequest) GoString() string {
	return s.String()
}

func (s *DescribeModelOperatorRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeModelOperatorRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeModelOperatorRequest) SetInstanceId(v string) *DescribeModelOperatorRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeModelOperatorRequest) SetRegion(v string) *DescribeModelOperatorRequest {
	s.Region = &v
	return s
}

func (s *DescribeModelOperatorRequest) Validate() error {
	return dara.Validate(s)
}
