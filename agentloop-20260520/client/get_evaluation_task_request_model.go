// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEvaluationTaskRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetEvaluationTaskRequest struct {
}

func (s GetEvaluationTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEvaluationTaskRequest) GoString() string {
	return s.String()
}

func (s *GetEvaluationTaskRequest) Validate() error {
	return dara.Validate(s)
}
