// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceMirrorRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DescribeServiceMirrorRequest struct {
}

func (s DescribeServiceMirrorRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMirrorRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceMirrorRequest) Validate() error {
	return dara.Validate(s)
}
