// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceLogRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DescribeResourceLogRequest struct {
}

func (s DescribeResourceLogRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourceLogRequest) Validate() error {
	return dara.Validate(s)
}
