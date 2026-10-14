// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOptimizeRunRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetOptimizeRunRequest struct {
}

func (s GetOptimizeRunRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOptimizeRunRequest) GoString() string {
	return s.String()
}

func (s *GetOptimizeRunRequest) Validate() error {
	return dara.Validate(s)
}
