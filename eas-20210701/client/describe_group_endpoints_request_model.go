// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupEndpointsRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DescribeGroupEndpointsRequest struct {
}

func (s DescribeGroupEndpointsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupEndpointsRequest) GoString() string {
	return s.String()
}

func (s *DescribeGroupEndpointsRequest) Validate() error {
	return dara.Validate(s)
}
