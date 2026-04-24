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
}

type DescribeModelOperatorRequest struct {
	// example:
	//
	// rds_copilot***_public_cn-*********6
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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

func (s *DescribeModelOperatorRequest) SetInstanceId(v string) *DescribeModelOperatorRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeModelOperatorRequest) Validate() error {
	return dara.Validate(s)
}
