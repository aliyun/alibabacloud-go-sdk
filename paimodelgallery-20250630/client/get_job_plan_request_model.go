// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobPlanRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetJobPlanRequest struct {
}

func (s GetJobPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJobPlanRequest) GoString() string {
	return s.String()
}

func (s *GetJobPlanRequest) Validate() error {
	return dara.Validate(s)
}
