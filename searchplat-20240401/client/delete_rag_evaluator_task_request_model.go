// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRagEvaluatorTaskRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteRagEvaluatorTaskRequest struct {
}

func (s DeleteRagEvaluatorTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRagEvaluatorTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteRagEvaluatorTaskRequest) Validate() error {
	return dara.Validate(s)
}
