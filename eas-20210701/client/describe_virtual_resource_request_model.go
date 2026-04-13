// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVirtualResourceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DescribeVirtualResourceRequest struct {
}

func (s DescribeVirtualResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVirtualResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeVirtualResourceRequest) Validate() error {
	return dara.Validate(s)
}
