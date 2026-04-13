// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartBenchmarkTaskRequest interface {
	dara.Model
	String() string
	GoString() string
}

type StartBenchmarkTaskRequest struct {
}

func (s StartBenchmarkTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s StartBenchmarkTaskRequest) GoString() string {
	return s.String()
}

func (s *StartBenchmarkTaskRequest) Validate() error {
	return dara.Validate(s)
}
