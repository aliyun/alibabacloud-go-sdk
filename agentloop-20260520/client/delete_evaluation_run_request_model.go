// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEvaluationRunRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteEvaluationRunRequest struct {
}

func (s DeleteEvaluationRunRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEvaluationRunRequest) GoString() string {
	return s.String()
}

func (s *DeleteEvaluationRunRequest) Validate() error {
	return dara.Validate(s)
}
