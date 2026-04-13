// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceEndpointsRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DescribeServiceEndpointsRequest struct {
}

func (s DescribeServiceEndpointsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceEndpointsRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceEndpointsRequest) Validate() error {
	return dara.Validate(s)
}
