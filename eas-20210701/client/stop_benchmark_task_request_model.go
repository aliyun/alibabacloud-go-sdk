// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopBenchmarkTaskRequest interface {
	dara.Model
	String() string
	GoString() string
}

type StopBenchmarkTaskRequest struct {
}

func (s StopBenchmarkTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s StopBenchmarkTaskRequest) GoString() string {
	return s.String()
}

func (s *StopBenchmarkTaskRequest) Validate() error {
	return dara.Validate(s)
}
