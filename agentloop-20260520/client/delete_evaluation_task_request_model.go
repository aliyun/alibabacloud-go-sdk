// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEvaluationTaskRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteEvaluationTaskRequest struct {
}

func (s DeleteEvaluationTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEvaluationTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteEvaluationTaskRequest) Validate() error {
	return dara.Validate(s)
}
