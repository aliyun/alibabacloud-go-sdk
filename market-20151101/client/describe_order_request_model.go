// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *DescribeOrderRequest
	GetOrderId() *string
}

type DescribeOrderRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 202*********415
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s DescribeOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOrderRequest) GoString() string {
	return s.String()
}

func (s *DescribeOrderRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *DescribeOrderRequest) SetOrderId(v string) *DescribeOrderRequest {
	s.OrderId = &v
	return s
}

func (s *DescribeOrderRequest) Validate() error {
	return dara.Validate(s)
}
