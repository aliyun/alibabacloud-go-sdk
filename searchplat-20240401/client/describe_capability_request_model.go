// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCapabilityRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DescribeCapabilityRequest struct {
}

func (s DescribeCapabilityRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCapabilityRequest) GoString() string {
	return s.String()
}

func (s *DescribeCapabilityRequest) Validate() error {
	return dara.Validate(s)
}
