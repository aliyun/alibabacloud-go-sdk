// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTaskRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteTaskRequest struct {
}

func (s DeleteTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteTaskRequest) Validate() error {
	return dara.Validate(s)
}
