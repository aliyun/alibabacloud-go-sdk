// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMemoryRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteMemoryRequest struct {
}

func (s DeleteMemoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMemoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteMemoryRequest) Validate() error {
	return dara.Validate(s)
}
