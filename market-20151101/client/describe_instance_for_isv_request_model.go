// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceForIsvRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeInstanceForIsvRequest
	GetInstanceId() *string
}

type DescribeInstanceForIsvRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 155****11
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeInstanceForIsvRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceForIsvRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceForIsvRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceForIsvRequest) SetInstanceId(v string) *DescribeInstanceForIsvRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceForIsvRequest) Validate() error {
	return dara.Validate(s)
}
