// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExperimentRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteExperimentRequest struct {
}

func (s DeleteExperimentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteExperimentRequest) GoString() string {
	return s.String()
}

func (s *DeleteExperimentRequest) Validate() error {
	return dara.Validate(s)
}
