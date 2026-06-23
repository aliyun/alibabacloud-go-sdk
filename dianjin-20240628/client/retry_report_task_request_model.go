// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryReportTaskRequest interface {
	dara.Model
	String() string
	GoString() string
}

type RetryReportTaskRequest struct {
}

func (s RetryReportTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s RetryReportTaskRequest) GoString() string {
	return s.String()
}

func (s *RetryReportTaskRequest) Validate() error {
	return dara.Validate(s)
}
