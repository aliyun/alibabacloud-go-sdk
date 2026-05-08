// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContextRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteContextRequest struct {
}

func (s DeleteContextRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteContextRequest) GoString() string {
	return s.String()
}

func (s *DeleteContextRequest) Validate() error {
	return dara.Validate(s)
}
