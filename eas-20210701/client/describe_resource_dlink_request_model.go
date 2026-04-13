// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceDLinkRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DescribeResourceDLinkRequest struct {
}

func (s DescribeResourceDLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceDLinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourceDLinkRequest) Validate() error {
	return dara.Validate(s)
}
