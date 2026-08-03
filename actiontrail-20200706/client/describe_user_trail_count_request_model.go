// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserTrailCountRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DescribeUserTrailCountRequest struct {
}

func (s DescribeUserTrailCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserTrailCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserTrailCountRequest) Validate() error {
	return dara.Validate(s)
}
