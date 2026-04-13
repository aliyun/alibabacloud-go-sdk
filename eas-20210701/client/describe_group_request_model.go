// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DescribeGroupRequest struct {
}

func (s DescribeGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeGroupRequest) Validate() error {
	return dara.Validate(s)
}
