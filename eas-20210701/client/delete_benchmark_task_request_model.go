// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBenchmarkTaskRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteBenchmarkTaskRequest struct {
}

func (s DeleteBenchmarkTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBenchmarkTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteBenchmarkTaskRequest) Validate() error {
	return dara.Validate(s)
}
