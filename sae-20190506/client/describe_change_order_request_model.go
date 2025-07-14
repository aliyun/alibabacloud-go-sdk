// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChangeOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChangeOrderId(v string) *DescribeChangeOrderRequest
	GetChangeOrderId() *string
}

type DescribeChangeOrderRequest struct {
	// The ID of the change order. You can call the [ListChangeOrders](https://help.aliyun.com/document_detail/126615.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 76fa5c0-9ebb-4bb4-b383-1f885447****
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s DescribeChangeOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeChangeOrderRequest) GoString() string {
	return s.String()
}

func (s *DescribeChangeOrderRequest) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *DescribeChangeOrderRequest) SetChangeOrderId(v string) *DescribeChangeOrderRequest {
	s.ChangeOrderId = &v
	return s
}

func (s *DescribeChangeOrderRequest) Validate() error {
	return dara.Validate(s)
}
