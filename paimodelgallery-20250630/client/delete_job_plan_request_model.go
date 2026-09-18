// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteJobPlanRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteJobPlanRequest struct {
}

func (s DeleteJobPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteJobPlanRequest) GoString() string {
	return s.String()
}

func (s *DeleteJobPlanRequest) Validate() error {
	return dara.Validate(s)
}
