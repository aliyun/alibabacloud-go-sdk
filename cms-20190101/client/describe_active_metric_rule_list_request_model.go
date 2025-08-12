// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActiveMetricRuleListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProduct(v string) *DescribeActiveMetricRuleListRequest
	GetProduct() *string
}

type DescribeActiveMetricRuleListRequest struct {
	// The abbreviation of the cloud service that supports initiative alert rules.
	//
	// For more information about how to obtain the name of a cloud service, see [DescribeProductsOfActiveMetricRule](https://help.aliyun.com/document_detail/114930.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
}

func (s DescribeActiveMetricRuleListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveMetricRuleListRequest) GoString() string {
	return s.String()
}

func (s *DescribeActiveMetricRuleListRequest) GetProduct() *string {
	return s.Product
}

func (s *DescribeActiveMetricRuleListRequest) SetProduct(v string) *DescribeActiveMetricRuleListRequest {
	s.Product = &v
	return s
}

func (s *DescribeActiveMetricRuleListRequest) Validate() error {
	return dara.Validate(s)
}
