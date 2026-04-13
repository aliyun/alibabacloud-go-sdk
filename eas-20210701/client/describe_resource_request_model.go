// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DescribeResourceRequest struct {
}

func (s DescribeResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourceRequest) Validate() error {
	return dara.Validate(s)
}
