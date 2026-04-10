// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggTaskGroupRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetAggTaskGroupRequest struct {
}

func (s GetAggTaskGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAggTaskGroupRequest) GoString() string {
	return s.String()
}

func (s *GetAggTaskGroupRequest) Validate() error {
	return dara.Validate(s)
}
