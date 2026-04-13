// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceCronScalerRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DescribeServiceCronScalerRequest struct {
}

func (s DescribeServiceCronScalerRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceCronScalerRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceCronScalerRequest) Validate() error {
	return dara.Validate(s)
}
