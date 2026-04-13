// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBenchmarkTaskRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DescribeBenchmarkTaskRequest struct {
}

func (s DescribeBenchmarkTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBenchmarkTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeBenchmarkTaskRequest) Validate() error {
	return dara.Validate(s)
}
