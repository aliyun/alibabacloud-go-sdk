// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeArtifactRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DescribeArtifactRequest struct {
}

func (s DescribeArtifactRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeArtifactRequest) GoString() string {
	return s.String()
}

func (s *DescribeArtifactRequest) Validate() error {
	return dara.Validate(s)
}
