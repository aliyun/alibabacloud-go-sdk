// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopSqlExecutionRequest interface {
	dara.Model
	String() string
	GoString() string
}

type StopSqlExecutionRequest struct {
}

func (s StopSqlExecutionRequest) String() string {
	return dara.Prettify(s)
}

func (s StopSqlExecutionRequest) GoString() string {
	return s.String()
}

func (s *StopSqlExecutionRequest) Validate() error {
	return dara.Validate(s)
}
