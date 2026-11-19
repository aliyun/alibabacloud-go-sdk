// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryMemoryTaskRequest interface {
	dara.Model
	String() string
	GoString() string
}

type RetryMemoryTaskRequest struct {
}

func (s RetryMemoryTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s RetryMemoryTaskRequest) GoString() string {
	return s.String()
}

func (s *RetryMemoryTaskRequest) Validate() error {
	return dara.Validate(s)
}
