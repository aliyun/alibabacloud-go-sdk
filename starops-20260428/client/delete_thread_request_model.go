// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteThreadRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteThreadRequest struct {
}

func (s DeleteThreadRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteThreadRequest) GoString() string {
	return s.String()
}

func (s *DeleteThreadRequest) Validate() error {
	return dara.Validate(s)
}
