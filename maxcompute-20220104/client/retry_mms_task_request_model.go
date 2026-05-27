// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryMmsTaskRequest interface {
	dara.Model
	String() string
	GoString() string
}

type RetryMmsTaskRequest struct {
}

func (s RetryMmsTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s RetryMmsTaskRequest) GoString() string {
	return s.String()
}

func (s *RetryMmsTaskRequest) Validate() error {
	return dara.Validate(s)
}
