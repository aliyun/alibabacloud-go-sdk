// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableExecuteStatementRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DisableExecuteStatementRequest struct {
}

func (s DisableExecuteStatementRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableExecuteStatementRequest) GoString() string {
	return s.String()
}

func (s *DisableExecuteStatementRequest) Validate() error {
	return dara.Validate(s)
}
