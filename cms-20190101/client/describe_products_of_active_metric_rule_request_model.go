// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProductsOfActiveMetricRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeProductsOfActiveMetricRuleRequest
	GetRegionId() *string
}

type DescribeProductsOfActiveMetricRuleRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeProductsOfActiveMetricRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductsOfActiveMetricRuleRequest) GoString() string {
	return s.String()
}

func (s *DescribeProductsOfActiveMetricRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeProductsOfActiveMetricRuleRequest) SetRegionId(v string) *DescribeProductsOfActiveMetricRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeProductsOfActiveMetricRuleRequest) Validate() error {
	return dara.Validate(s)
}
