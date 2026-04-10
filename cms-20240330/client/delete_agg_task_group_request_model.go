// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAggTaskGroupRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteAggTaskGroupRequest struct {
}

func (s DeleteAggTaskGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAggTaskGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteAggTaskGroupRequest) Validate() error {
	return dara.Validate(s)
}
