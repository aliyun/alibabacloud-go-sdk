// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWarmPoolRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DescribeWarmPoolRequest struct {
}

func (s DescribeWarmPoolRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWarmPoolRequest) GoString() string {
	return s.String()
}

func (s *DescribeWarmPoolRequest) Validate() error {
	return dara.Validate(s)
}
