// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFileRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DescribeFileRequest struct {
}

func (s DescribeFileRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileRequest) GoString() string {
	return s.String()
}

func (s *DescribeFileRequest) Validate() error {
	return dara.Validate(s)
}
