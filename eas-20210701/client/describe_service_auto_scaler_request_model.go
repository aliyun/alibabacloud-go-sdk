// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceAutoScalerRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DescribeServiceAutoScalerRequest struct {
}

func (s DescribeServiceAutoScalerRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceAutoScalerRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceAutoScalerRequest) Validate() error {
	return dara.Validate(s)
}
