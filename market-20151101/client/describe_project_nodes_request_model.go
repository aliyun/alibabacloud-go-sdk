// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProjectNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeProjectNodesRequest
	GetInstanceId() *string
}

type DescribeProjectNodesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4****89
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeProjectNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeProjectNodesRequest) GoString() string {
	return s.String()
}

func (s *DescribeProjectNodesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeProjectNodesRequest) SetInstanceId(v string) *DescribeProjectNodesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeProjectNodesRequest) Validate() error {
	return dara.Validate(s)
}
