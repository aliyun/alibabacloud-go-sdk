// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExperimentRunRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetExperimentRunRequest struct {
}

func (s GetExperimentRunRequest) String() string {
	return dara.Prettify(s)
}

func (s GetExperimentRunRequest) GoString() string {
	return s.String()
}

func (s *GetExperimentRunRequest) Validate() error {
	return dara.Validate(s)
}
