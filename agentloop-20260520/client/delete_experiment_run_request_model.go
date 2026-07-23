// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExperimentRunRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteExperimentRunRequest struct {
}

func (s DeleteExperimentRunRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteExperimentRunRequest) GoString() string {
	return s.String()
}

func (s *DeleteExperimentRunRequest) Validate() error {
	return dara.Validate(s)
}
