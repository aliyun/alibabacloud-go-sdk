// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeInstanceConfigsRequest
	GetInstanceId() *string
}

type DescribeInstanceConfigsRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-123xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeInstanceConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceConfigsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceConfigsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceConfigsRequest) SetInstanceId(v string) *DescribeInstanceConfigsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceConfigsRequest) Validate() error {
	return dara.Validate(s)
}
