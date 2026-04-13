// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMemoryNodeRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteMemoryNodeRequest struct {
}

func (s DeleteMemoryNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMemoryNodeRequest) GoString() string {
	return s.String()
}

func (s *DeleteMemoryNodeRequest) Validate() error {
	return dara.Validate(s)
}
