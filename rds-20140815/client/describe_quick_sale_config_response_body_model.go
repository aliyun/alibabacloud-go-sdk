// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQuickSaleConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCommodity(v string) *DescribeQuickSaleConfigResponseBody
	GetCommodity() *string
	SetItems(v map[string]interface{}) *DescribeQuickSaleConfigResponseBody
	GetItems() map[string]interface{}
	SetRequestId(v string) *DescribeQuickSaleConfigResponseBody
	GetRequestId() *string
}

type DescribeQuickSaleConfigResponseBody struct {
	// The product code. Valid values:
	//
	// 	- rds: The instance is a subscription instance.
	//
	// 	- bards: The instance is a pay-as-you-go instance.
	//
	// example:
	//
	// rds
	Commodity *string `json:"Commodity,omitempty" xml:"Commodity,omitempty"`
	// The configuration details of the product.
	Items map[string]interface{} `json:"Items,omitempty" xml:"Items,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5DFFE9EC-3369-5937-A4E2-507C0C86A4C6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeQuickSaleConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeQuickSaleConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeQuickSaleConfigResponseBody) GetCommodity() *string {
	return s.Commodity
}

func (s *DescribeQuickSaleConfigResponseBody) GetItems() map[string]interface{} {
	return s.Items
}

func (s *DescribeQuickSaleConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeQuickSaleConfigResponseBody) SetCommodity(v string) *DescribeQuickSaleConfigResponseBody {
	s.Commodity = &v
	return s
}

func (s *DescribeQuickSaleConfigResponseBody) SetItems(v map[string]interface{}) *DescribeQuickSaleConfigResponseBody {
	s.Items = v
	return s
}

func (s *DescribeQuickSaleConfigResponseBody) SetRequestId(v string) *DescribeQuickSaleConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeQuickSaleConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
