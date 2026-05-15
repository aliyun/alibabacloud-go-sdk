// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExecuteStatementEnabledRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetExecuteStatementEnabledRequest struct {
}

func (s GetExecuteStatementEnabledRequest) String() string {
	return dara.Prettify(s)
}

func (s GetExecuteStatementEnabledRequest) GoString() string {
	return s.String()
}

func (s *GetExecuteStatementEnabledRequest) Validate() error {
	return dara.Validate(s)
}
