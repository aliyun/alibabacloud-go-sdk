// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DescribeServiceRequest struct {
}

func (s DescribeServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceRequest) Validate() error {
	return dara.Validate(s)
}
