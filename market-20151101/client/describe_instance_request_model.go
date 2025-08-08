// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeInstanceRequest
	GetInstanceId() *string
	SetOrderType(v string) *DescribeInstanceRequest
	GetOrderType() *string
	SetOwnerId(v int64) *DescribeInstanceRequest
	GetOwnerId() *int64
}

type DescribeInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 155****11
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// NEW
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribeInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeInstanceRequest) SetInstanceId(v string) *DescribeInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceRequest) SetOrderType(v string) *DescribeInstanceRequest {
	s.OrderType = &v
	return s
}

func (s *DescribeInstanceRequest) SetOwnerId(v int64) *DescribeInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstanceRequest) Validate() error {
	return dara.Validate(s)
}
