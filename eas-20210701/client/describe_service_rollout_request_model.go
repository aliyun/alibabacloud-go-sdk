// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceRolloutRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DescribeServiceRolloutRequest struct {
}

func (s DescribeServiceRolloutRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceRolloutRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceRolloutRequest) Validate() error {
	return dara.Validate(s)
}
