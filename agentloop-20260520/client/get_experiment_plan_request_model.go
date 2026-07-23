// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExperimentPlanRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetExperimentPlanRequest struct {
}

func (s GetExperimentPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s GetExperimentPlanRequest) GoString() string {
	return s.String()
}

func (s *GetExperimentPlanRequest) Validate() error {
	return dara.Validate(s)
}
