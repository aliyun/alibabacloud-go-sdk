// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOrderForIsvRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *DescribeOrderForIsvRequest
	GetOrderId() *string
}

type DescribeOrderForIsvRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 202*********415
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s DescribeOrderForIsvRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOrderForIsvRequest) GoString() string {
	return s.String()
}

func (s *DescribeOrderForIsvRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *DescribeOrderForIsvRequest) SetOrderId(v string) *DescribeOrderForIsvRequest {
	s.OrderId = &v
	return s
}

func (s *DescribeOrderForIsvRequest) Validate() error {
	return dara.Validate(s)
}
