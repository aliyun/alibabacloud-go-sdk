// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStackExecutionResultRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetStackExecutionResultRequest struct {
}

func (s GetStackExecutionResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStackExecutionResultRequest) GoString() string {
	return s.String()
}

func (s *GetStackExecutionResultRequest) Validate() error {
	return dara.Validate(s)
}
