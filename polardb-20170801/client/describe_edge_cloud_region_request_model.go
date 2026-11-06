// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEdgeCloudRegionRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DescribeEdgeCloudRegionRequest struct {
}

func (s DescribeEdgeCloudRegionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEdgeCloudRegionRequest) GoString() string {
	return s.String()
}

func (s *DescribeEdgeCloudRegionRequest) Validate() error {
	return dara.Validate(s)
}
