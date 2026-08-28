// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPatrolConfigRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetPatrolConfigRequest struct {
}

func (s GetPatrolConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPatrolConfigRequest) GoString() string {
	return s.String()
}

func (s *GetPatrolConfigRequest) Validate() error {
	return dara.Validate(s)
}
