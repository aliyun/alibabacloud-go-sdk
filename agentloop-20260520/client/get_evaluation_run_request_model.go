// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEvaluationRunRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetEvaluationRunRequest struct {
}

func (s GetEvaluationRunRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEvaluationRunRequest) GoString() string {
	return s.String()
}

func (s *GetEvaluationRunRequest) Validate() error {
	return dara.Validate(s)
}
