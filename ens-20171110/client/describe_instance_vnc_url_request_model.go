// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceVncUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeInstanceVncUrlRequest
	GetInstanceId() *string
}

type DescribeInstanceVncUrlRequest struct {
	// The ID of the instance
	//
	// This parameter is required.
	//
	// example:
	//
	// i-5bp1hzoinajzkh91h****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeInstanceVncUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceVncUrlRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceVncUrlRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceVncUrlRequest) SetInstanceId(v string) *DescribeInstanceVncUrlRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceVncUrlRequest) Validate() error {
	return dara.Validate(s)
}
