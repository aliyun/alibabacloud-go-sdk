// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExperimentPlanRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteExperimentPlanRequest struct {
}

func (s DeleteExperimentPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteExperimentPlanRequest) GoString() string {
	return s.String()
}

func (s *DeleteExperimentPlanRequest) Validate() error {
	return dara.Validate(s)
}
