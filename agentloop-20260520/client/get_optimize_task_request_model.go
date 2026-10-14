// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOptimizeTaskRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetOptimizeTaskRequest struct {
}

func (s GetOptimizeTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOptimizeTaskRequest) GoString() string {
	return s.String()
}

func (s *GetOptimizeTaskRequest) Validate() error {
	return dara.Validate(s)
}
