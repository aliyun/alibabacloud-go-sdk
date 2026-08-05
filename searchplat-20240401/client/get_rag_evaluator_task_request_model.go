// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRagEvaluatorTaskRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetRagEvaluatorTaskRequest struct {
}

func (s GetRagEvaluatorTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRagEvaluatorTaskRequest) GoString() string {
	return s.String()
}

func (s *GetRagEvaluatorTaskRequest) Validate() error {
	return dara.Validate(s)
}
